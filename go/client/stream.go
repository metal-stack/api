package client

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"time"
)

const defaultStreamBackoff = time.Second

type (
	// Stream is the client-side view of a server-streaming RPC.
	Stream[T any] interface {
		Receive() (*T, error)
		Close() error
	}

	StreamFunc[T any] func(ctx context.Context) (Stream[T], error)

	StreamOption func(*streamOptions)

	streamOptions struct {
		backoff time.Duration
		log     *slog.Logger
	}
)

func WithStreamBackoff(d time.Duration) StreamOption {
	return func(o *streamOptions) {
		o.backoff = d
	}
}

func WithStreamLogger(log *slog.Logger) StreamOption {
	return func(o *streamOptions) {
		o.log = log
	}
}

func (o *streamOptions) debug(msg string, args ...any) {
	if o.log == nil {
		return
	}
	o.log.Debug(msg, args...)
}

func (o *streamOptions) error(msg string, err error) {
	if o.log == nil {
		return
	}
	o.log.Error(msg, "error", err, "backoff", o.backoff)
}

func ReconnectingStreamRead[T any](ctx context.Context, open StreamFunc[T], opts ...StreamOption) (<-chan *T, <-chan error) {
	var (
		options    = &streamOptions{backoff: defaultStreamBackoff}
		messages   = make(chan *T)
		errorsChan = make(chan error)
	)
	for _, opt := range opts {
		opt(options)
	}

	go func() {
		defer close(messages)
		defer close(errorsChan)
		defer func() { options.debug("stream read stopped", "reason", ctx.Err()) }()

		for {
			if ctx.Err() != nil {
				return
			}

			options.debug("opening stream")

			stream, err := open(ctx)
			if err != nil {
				options.error("opening stream failed, reconnecting", err)

				if !send(ctx, errorsChan, err) {
					return
				}
				if !wait(ctx, options.backoff) {
					return
				}
				continue
			}

			reconnect := consumeStream(ctx, stream, messages, errorsChan, options)
			_ = stream.Close()
			if !reconnect {
				return
			}

			options.debug("stream ended, reconnecting")

			if !wait(ctx, options.backoff) {
				return
			}
		}
	}()

	return messages, errorsChan
}

func consumeStream[T any](ctx context.Context, stream Stream[T], messages chan<- *T, errorsChan chan<- error, options *streamOptions) bool {
	for {
		msg, err := stream.Receive()
		if err != nil {
			if !errors.Is(err, io.EOF) && !errors.Is(err, context.Canceled) && ctx.Err() == nil {
				options.error("reading stream failed, reconnecting", err)
				if !send(ctx, errorsChan, err) {
					return false
				}
			}
			break
		}

		if !send(ctx, messages, msg) {
			return false
		}
	}

	return ctx.Err() == nil
}

func send[T any](ctx context.Context, ch chan<- T, v T) bool {
	select {
	case <-ctx.Done():
		return false
	case ch <- v:
		return true
	}
}

func wait(ctx context.Context, d time.Duration) bool {
	if d <= 0 {
		return ctx.Err() == nil
	}

	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case <-ctx.Done():
		return false
	case <-t.C:
		return true
	}
}
