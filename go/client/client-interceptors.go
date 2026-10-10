package client

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"connectrpc.com/connect/v2"
	apiv2models "github.com/metal-stack/api/go/metalstack/api/v2"
)

// authInterceptor adds the required auth headers
type authInterceptor struct {
	config *DialConfig
}

func (i *authInterceptor) intercept(next connect.ClientFunc) connect.ClientFunc {
	return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
		if callInfo, ok := connect.CallInfoForClientContext(ctx); ok {
			callInfo.RequestHeader().Set("Authorization", "Bearer "+i.config.Token)
		}
		return next(ctx, spec)
	}
}

type loggingInterceptor struct {
	config *DialConfig
}

func (i *loggingInterceptor) intercept(next connect.ClientFunc) connect.ClientFunc {
	return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
		stream, err := next(ctx, spec)
		if err != nil {
			return nil, err
		}
		return &loggingStream{config: i.config, spec: spec, stream: stream}, nil
	}
}

type loggingStream struct {
	config *DialConfig
	spec   connect.Spec
	stream connect.ClientStream
}

func (s *loggingStream) SendHeaders() error {
	return s.stream.SendHeaders()
}

func (s *loggingStream) Send(msg any) error {
	s.config.Log.Debug("intercept", "request procedure", s.spec.Procedure, "body", msg)
	return s.stream.Send(msg)
}

func (s *loggingStream) CloseSend() error {
	return s.stream.CloseSend()
}

func (s *loggingStream) Receive(msg any) error {
	err := s.stream.Receive(msg)
	if err != nil {
		return err
	}
	s.config.Log.Debug("intercept", "request procedure", s.spec.Procedure, "response", msg)
	return nil
}

func (s *loggingStream) Close() error {
	return s.stream.Close()
}

type tokenRenewingInterceptor struct {
	config *DialConfig
	client *client

	renewing atomic.Bool

	sync.Mutex
}

func (i *tokenRenewingInterceptor) intercept(next connect.ClientFunc) connect.ClientFunc {
	return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
		err := i.renewTokenIfNeeded()
		if err != nil {
			return nil, err
		}
		return next(ctx, spec)
	}
}

func (i *tokenRenewingInterceptor) renewTokenIfNeeded() error {
	if i.renewing.Load() {
		return nil
	}
	if i.config.Log == nil {
		i.config.Log = slog.Default()
	}

	if i.config.TokenFile != "" {
		return i.rereadTokenFile()
	}

	return i.renewToken()
}

func (i *tokenRenewingInterceptor) rereadTokenFile() error {
	// The token is refreshed by a sidecar, this means we should periodically read the tokenfile
	// and store the token in config instead und update the tokenFileLastRead afterwards
	if time.Since(i.config.tokenFileLastRead) < i.config.TokenFileRereadDuration {
		return nil
	}
	i.config.Log.Info("tokenfile specified, re-reading content")

	content, err := os.ReadFile(i.config.TokenFile)
	if err != nil {
		return fmt.Errorf("unable to read tokenfile %w", err)
	}
	newToken := string(content)
	if i.config.Token == newToken {
		return nil
	}

	i.renewing.Store(true)
	defer i.renewing.Store(false)

	i.Lock()
	defer i.Unlock()

	i.config.Token = newToken
	err = i.config.parseTokenClaims()
	if err != nil {
		return fmt.Errorf("unable to parse token %w", err)
	}
	i.config.tokenFileLastRead = time.Now()
	return nil
}

func (i *tokenRenewingInterceptor) renewToken() error {
	if i.config.expiresAt.IsZero() {
		return nil
	}

	replaceBefore := i.config.expiresAt.Sub(i.config.issuedAt) / tokenRenewChecksDuringLifetime

	if time.Until(i.config.expiresAt) > replaceBefore {
		return nil
	}

	i.renewing.Store(true)
	defer i.renewing.Store(false)

	i.config.Log.Info("call token refresh, current token expires soon", "expires", i.config.expiresAt.String())

	i.Lock()
	defer i.Unlock()

	resp, err := i.client.Apiv2().Token().Refresh(context.Background(), &apiv2models.TokenServiceRefreshRequest{})
	if err != nil {
		return fmt.Errorf("unable to refresh token %w", err)
	}

	i.config.Token = resp.Secret
	err = i.config.parse()
	if err != nil {
		return fmt.Errorf("unable to parse token %w", err)
	}

	if i.config.TokenRenewal.PersistTokenFn == nil {
		return nil
	}

	err = i.config.TokenRenewal.PersistTokenFn(i.config.Token)
	if err != nil {
		return fmt.Errorf("unable to persist token %w", err)
	}

	i.config.Log.Info("token refreshed, new token expires in", "expires", i.config.expiresAt.String())
	return nil
}
