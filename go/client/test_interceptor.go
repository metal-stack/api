package client

import (
	"context"
	"io"
	"reflect"
	"testing"

	"connectrpc.com/connect/v2"
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
	WantResponse func() proto.Message
	WantError    *connect.Error

	WantStreamResponses func() []proto.Message
	BlockingStream      bool
}

func NewTestInterceptor(t *testing.T, calls []ClientCall) connect.ClientInterceptor {
	interceptor := &TestClientInterceptor{
		t:     t,
		calls: calls,
	}
	return interceptor.intercept
}

func (t *TestClientInterceptor) intercept(next connect.ClientFunc) connect.ClientFunc {
	return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
		if t.count >= len(t.calls) {
			t.t.Errorf("received an unexpected client call for procedure %s", spec.Procedure)
			t.t.FailNow()
		}

		call := t.calls[t.count]
		t.count++

		return &testClientStream{t: t.t, ctx: ctx, spec: spec, call: call}, nil
	}
}

type testClientStream struct {
	t    *testing.T
	ctx  context.Context
	spec connect.Spec
	call ClientCall

	responses   []proto.Message
	recvIndex   int
	errReturned bool
}

func (c *testClientStream) SendHeaders() error { return nil }

func (c *testClientStream) Send(msg any) error {
	got, ok := msg.(proto.Message)
	if !ok {
		c.t.Errorf("unexpected send message of type %T", msg)
		return connect.NewError(connect.CodeInternal, "unexpected send message")
	}

	if diff := cmp.Diff(c.call.WantRequest, got, protocmp.Transform(), IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
		c.t.Errorf("request diff (+got -want):\n %s", diff)
		return connect.NewError(connect.CodeInvalidArgument, "unexpected request")
	}

	if c.call.WantStreamResponses != nil {
		c.responses = c.call.WantStreamResponses()
	}

	return nil
}

func (c *testClientStream) CloseSend() error { return nil }

func (c *testClientStream) Receive(msg any) error {
	if c.spec.StreamType == connect.StreamTypeUnary {
		return c.receiveUnary(msg)
	}
	return c.receiveStream(msg)
}

func (c *testClientStream) receiveUnary(msg any) error {
	if c.call.WantError != nil {
		return c.call.WantError
	}
	if c.call.WantResponse == nil {
		return nil
	}
	return c.copyMessage(msg, c.call.WantResponse())
}

func (c *testClientStream) receiveStream(msg any) error {
	if c.recvIndex < len(c.responses) {
		err := c.copyMessage(msg, c.responses[c.recvIndex])
		c.recvIndex++
		return err
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

func (c *testClientStream) Close() error { return nil }

func (c *testClientStream) copyMessage(dst any, src proto.Message) error {
	target, ok := dst.(proto.Message)
	if !ok {
		c.t.Errorf("unexpected receive message of type %T", dst)
		return connect.NewError(connect.CodeInternal, "unexpected receive message")
	}

	proto.Reset(target)
	proto.Merge(target, src)
	return nil
}

func IgnoreUnexported() cmp.Option {
	// the exporter opt allows all unexported fields: https://github.com/google/go-cmp/pull/176
	return cmp.Exporter(func(reflect.Type) bool { return true })
}
