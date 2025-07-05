package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/metal-stack/api/go/client"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	vs := &mockService{}
	mux := http.NewServeMux()
	mux.Handle(apiv2connect.NewVersionServiceHandler(vs))
	server := httptest.NewTLSServer(mux)
	server.EnableHTTP2 = true
	defer func() {
		server.Close()
	}()

	server.Client()
	c := client.New(client.DialConfig{
		BaseURL:   server.URL,
		Token:     "",
		Transport: server.Client().Transport,
	})
	v, err := c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)
}

type mockService struct {
}

func (m *mockService) Get(ctx context.Context, req *connect.Request[apiv2.VersionServiceGetRequest]) (*connect.Response[apiv2.VersionServiceGetResponse], error) {
	return connect.NewResponse(&apiv2.VersionServiceGetResponse{Version: &apiv2.Version{Version: "1.0"}}), nil
}
