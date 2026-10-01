package client_test

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/metal-stack/api/go/client"
	infrav2 "github.com/metal-stack/api/go/metalstack/infra/v2"
	"github.com/stretchr/testify/require"
)

func Test_ReconnectingStreamRead_Reconnects(t *testing.T) {
	req := &infrav2.WaitForBMCCommandRequest{Partition: "p1"}

	c, err := client.New(&client.DialConfig{
		BaseURL: "http://this-is-just-for-testing",
		Interceptors: []connect.Interceptor{
			client.NewTestInterceptor(t, []client.ClientCall{
				{
					WantRequest: req,
					WantStreamResponses: func() []connect.AnyResponse {
						return []connect.AnyResponse{
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "1"}),
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "2"}),
						}
					},
				},
				{
					WantRequest: req,
					WantStreamResponses: func() []connect.AnyResponse {
						return []connect.AnyResponse{
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "3"}),
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "4"}),
						}
					},
				},
				{
					WantRequest: req,
					WantStreamResponses: func() []connect.AnyResponse {
						return []connect.AnyResponse{connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "5"})}
					},
					BlockingStream: true,
				},
			}),
		},
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	messages, errs := client.ReconnectingStreamRead(ctx, func(ctx context.Context) (*connect.ServerStreamForClient[infrav2.WaitForBMCCommandResponse], error) {
		return c.Infrav2().BMC().WaitForBMCCommand(ctx, &infrav2.WaitForBMCCommandRequest{Partition: "p1"})
	}, client.WithStreamBackoff(0), client.WithStreamLogger(slog.Default()))

	for i := range 5 {
		select {
		case msg := <-messages:
			require.Equal(t, strconv.Itoa(i+1), msg.CommandId)
		case err := <-errs:
			t.Fatalf("unexpected error: %v", err)
		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for message")
		}
	}

	cancel()
	requireChannelsClosed(t, messages, errs)
}

func Test_ReconnectingStreamRead_ReportsErrorsAndReconnects(t *testing.T) {
	req := &infrav2.WaitForBMCCommandRequest{Partition: "p1"}

	c, err := client.New(&client.DialConfig{
		BaseURL: "http://this-is-just-for-testing",
		Interceptors: []connect.Interceptor{
			client.NewTestInterceptor(t, []client.ClientCall{
				{
					WantRequest: req,
					WantError:   connect.NewError(connect.CodeInternal, errors.New("stream failure")),
				},
				{
					WantRequest: req,
					WantStreamResponses: func() []connect.AnyResponse {
						return []connect.AnyResponse{
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "1"}),
							connect.NewResponse(&infrav2.WaitForBMCCommandResponse{CommandId: "2"}),
						}
					},
					BlockingStream: true,
				},
			}),
		},
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	messages, errs := client.ReconnectingStreamRead(ctx, func(ctx context.Context) (*connect.ServerStreamForClient[infrav2.WaitForBMCCommandResponse], error) {
		return c.Infrav2().BMC().WaitForBMCCommand(ctx, &infrav2.WaitForBMCCommandRequest{Partition: "p1"})
	}, client.WithStreamBackoff(0))

	select {
	case err := <-errs:
		require.ErrorContains(t, err, "stream failure")
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for error")
	}

	for i := range 2 {
		select {
		case msg := <-messages:
			require.Equal(t, strconv.Itoa(i+1), msg.CommandId)
		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for message after reconnect")
		}
	}

	cancel()
	requireChannelsClosed(t, messages, errs)
}

func Test_ReconnectingStreamRead_ClosesOnContextCancel(t *testing.T) {
	req := &infrav2.WaitForBMCCommandRequest{Partition: "p1"}

	c, err := client.New(&client.DialConfig{
		BaseURL: "http://this-is-just-for-testing",
		Interceptors: []connect.Interceptor{
			client.NewTestInterceptor(t, []client.ClientCall{
				{
					WantRequest:    req,
					BlockingStream: true,
				},
			}),
		},
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())

	messages, errs := client.ReconnectingStreamRead(ctx, func(ctx context.Context) (*connect.ServerStreamForClient[infrav2.WaitForBMCCommandResponse], error) {
		return c.Infrav2().BMC().WaitForBMCCommand(ctx, &infrav2.WaitForBMCCommandRequest{Partition: "p1"})
	}, client.WithStreamBackoff(0))

	cancel()
	requireChannelsClosed(t, messages, errs)
}

func requireChannelsClosed[T any](t *testing.T, messages <-chan *T, errs <-chan error) {
	t.Helper()

	deadline := time.After(5 * time.Second)
	messagesClosed, errsClosed := false, false
	for !messagesClosed || !errsClosed {
		select {
		case _, ok := <-messages:
			if !ok {
				messagesClosed = true
				messages = nil
			}
		case _, ok := <-errs:
			if !ok {
				errsClosed = true
				errs = nil
			}
		case <-deadline:
			t.Fatal("both channels should be closed after cancel")
		}
	}
}
