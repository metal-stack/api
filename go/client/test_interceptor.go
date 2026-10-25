package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/testing/protocmp"
)

type TestClientInterceptor struct {
	t     *testing.T
	calls []ClientCall
	count int
}

type ClientCall struct {
	WantRequest  proto.Message
	WantResponse func() connect.AnyResponse
	WantError    *connect.Error

	WantStreamResponses func() []connect.AnyResponse
	BlockingStream      bool
}

func NewTestInterceptor(t *testing.T, calls []ClientCall) *TestClientInterceptor {
	return &TestClientInterceptor{
		t:     t,
		calls: calls,
	}
}

func (t *TestClientInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		defer func() { t.count++ }()

		if t.count >= len(t.calls) {
			t.t.Errorf("received an unexpected client call of type %T: %v", ar.Any(), ar.Any())
			t.t.FailNow()
		}

		call := t.calls[t.count]

		if diff := cmp.Diff(call.WantRequest, ar.Any(), protocmp.Transform(), IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
			t.t.Errorf("request diff (+got -want):\n %s", diff)
			t.t.FailNow()
		}

		if call.WantError != nil {
			return nil, call.WantError
		}

		return call.WantResponse(), nil
	}
}

func (t *TestClientInterceptor) WrapStreamingClient(connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		if t.count >= len(t.calls) {
			t.t.Errorf("received an unexpected streaming client call for procedure %s", spec.Procedure)
			t.t.FailNow()
		}

		call := t.calls[t.count]
		t.count++

		return &testStreamingClientConn{t: t.t, ctx: ctx, spec: spec, call: call}
	}
}

func (t *TestClientInterceptor) WrapStreamingHandler(connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	t.t.Errorf("streaming handler not supported")
	return nil
}

type testStreamingClientConn struct {
	t    *testing.T
	ctx  context.Context
	spec connect.Spec
	call ClientCall

	responses   []connect.AnyResponse
	recvIndex   int
	errReturned bool
}

func (c *testStreamingClientConn) Spec() connect.Spec { return c.spec }
func (c *testStreamingClientConn) Peer() connect.Peer { return connect.Peer{} }

func (c *testStreamingClientConn) Send(msg any) error {
	got, ok := msg.(proto.Message)
	if !ok {
		c.t.Errorf("unexpected send message of type %T", msg)
		return connect.NewError(connect.CodeInternal, errors.New("unexpected send message"))
	}

	if diff := cmp.Diff(c.call.WantRequest, got, protocmp.Transform(), IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
		c.t.Errorf("request diff (+got -want):\n %s", diff)
		return connect.NewError(connect.CodeInvalidArgument, errors.New("unexpected request"))
	}

	if c.call.WantStreamResponses != nil {
		c.responses = c.call.WantStreamResponses()
	}

	return nil
}

func (c *testStreamingClientConn) RequestHeader() http.Header  { return http.Header{} }
func (c *testStreamingClientConn) CloseRequest() error         { return nil }
func (c *testStreamingClientConn) ResponseHeader() http.Header { return http.Header{} }
func (c *testStreamingClientConn) ResponseTrailer() http.Header {
	return http.Header{}
}
func (c *testStreamingClientConn) CloseResponse() error { return nil }

func (c *testStreamingClientConn) Receive(msg any) error {
	if c.recvIndex < len(c.responses) {
		target, ok := msg.(proto.Message)
		if !ok {
			c.t.Errorf("unexpected receive message of type %T", msg)
			return connect.NewError(connect.CodeInternal, errors.New("unexpected receive message"))
		}

		source, ok := c.responses[c.recvIndex].Any().(proto.Message)
		if !ok {
			c.t.Errorf("unexpected response message of type %T", c.responses[c.recvIndex].Any())
			return connect.NewError(connect.CodeInternal, errors.New("unexpected response message"))
		}

		proto.Reset(target)
		proto.Merge(target, source)
		c.recvIndex++
		return nil
	}

	if c.call.WantError != nil && !c.errReturned {
		c.errReturned = true
		return c.call.WantError
	}

	if c.call.BlockingStream {
		<-c.ctx.Done()
	}

	return io.EOF
}

func IgnoreUnexported() cmp.Option {
	// the exporter opt allows all unexported fields: https://github.com/google/go-cmp/pull/176
	return cmp.Exporter(func(reflect.Type) bool { return true })
}
