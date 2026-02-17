package client_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/metal-stack/api/go/client"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	infra "github.com/metal-stack/api/go/metalstack/infra/v2"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Ping(t *testing.T) {
	var (
		mux = http.NewServeMux()
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		cs  = &mockComponentService{log: log}
	)

	mux.Handle(infrav2connect.NewComponentServiceHandler(cs))
	server := httptest.NewTLSServer(mux)
	server.EnableHTTP2 = true
	defer server.Close()

	ctx := t.Context()

	tokenString, err := generateToken(2 * time.Second)
	require.NoError(t, err)

	c, err := client.New(&client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
		Log:       log,
	})
	require.NoError(t, err)

	start := time.Now()
	config := &client.PingConfig{
		ComponentType: apiv2.ComponentType_COMPONENT_TYPE_METAL_CORE,
		Identifier:    new("server01"),
		StartedAt:     start,
		Interval:      5 * time.Second,
	}

	want := []*infra.ComponentServicePingRequest{
		{Type: apiv2.ComponentType_COMPONENT_TYPE_METAL_CORE, Identifier: "server01", StartedAt: timestamppb.New(start), Interval: durationpb.New(5 * time.Second), Version: &apiv2.Version{}},
		{Type: apiv2.ComponentType_COMPONENT_TYPE_METAL_CORE, Identifier: "server01", StartedAt: timestamppb.New(start), Interval: durationpb.New(5 * time.Second), Version: &apiv2.Version{}},
	}

	c.Ping(ctx, config)
	time.Sleep(3 * config.Interval)
	if diff := cmp.Diff(
		cs.pings, want,
		protocmp.Transform(),
		cmpopts.IgnoreUnexported(),
	); diff != "" {
		t.Errorf("%v, want %v diff: %s", cs.pings, want, diff)
	}
}

type mockComponentService struct {
	log   *slog.Logger
	pings []*infra.ComponentServicePingRequest
}

func (m *mockComponentService) Ping(_ context.Context, req *infra.ComponentServicePingRequest) (*infra.ComponentServicePingResponse, error) {
	m.log.Debug("ping", "req", req)
	m.pings = append(m.pings, req)
	return nil, nil
}
