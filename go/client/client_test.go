package client_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
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

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	require.NoError(t, err)

	server.Client()
	c, err := client.New(client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
	})
	require.NoError(t, err)
	v, err := c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)

	// TODO test token refresh
}

type mockService struct {
}

func (m *mockService) Get(ctx context.Context, req *connect.Request[apiv2.VersionServiceGetRequest]) (*connect.Response[apiv2.VersionServiceGetResponse], error) {
	return connect.NewResponse(&apiv2.VersionServiceGetResponse{Version: &apiv2.Version{Version: "1.0"}}), nil
}
