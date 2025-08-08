package client_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	var (
		vs  = &mockVersionService{}
		ts  = &mockTokenService{}
		mux = http.NewServeMux()
	)

	mux.Handle(apiv2connect.NewVersionServiceHandler(vs))
	mux.Handle(apiv2connect.NewTokenServiceHandler(ts))
	server := httptest.NewTLSServer(mux)
	server.EnableHTTP2 = true
	defer func() {
		server.Close()
	}()

	tokenString, err := generateToken(1 * time.Second)
	require.NoError(t, err)

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	c, err := client.New(&client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
		TokenRenewal: &client.TokenRenewal{
			PersistTokenFn: func(token string) error {
				t.Log("token persisted:", token)
				return nil
			},
		},
		Log: log,
	})
	require.NoError(t, err)
	v, err := c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)
	require.False(t, ts.wasCalled)
	require.Equal(t, tokenString, vs.token)

	time.Sleep(300 * time.Millisecond)
	v, err = c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)
	require.False(t, ts.wasCalled)
	require.Equal(t, tokenString, vs.token)

	time.Sleep(1 * time.Second)
	v, err = c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)

	require.True(t, ts.wasCalled)
	require.NotEqual(t, tokenString, vs.token, "token must have changed")
}

func generateToken(duration time.Duration) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type mockVersionService struct {
	token string
}

func (m *mockVersionService) Get(ctx context.Context, req *connect.Request[apiv2.VersionServiceGetRequest]) (*connect.Response[apiv2.VersionServiceGetResponse], error) {
	authHeader := req.Header().Get("Authorization")
	_, token, found := strings.Cut(authHeader, "Bearer ")

	if !found {
		return nil, fmt.Errorf("unable to extract token from header:%s", authHeader)
	}

	m.token = token
	return connect.NewResponse(&apiv2.VersionServiceGetResponse{Version: &apiv2.Version{Version: "1.0"}}), nil
}

type mockTokenService struct {
	wasCalled bool
}

// Create implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Create(context.Context, *connect.Request[apiv2.TokenServiceCreateRequest]) (*connect.Response[apiv2.TokenServiceCreateResponse], error) {
	panic("unimplemented")
}

// Get implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Get(context.Context, *connect.Request[apiv2.TokenServiceGetRequest]) (*connect.Response[apiv2.TokenServiceGetResponse], error) {
	panic("unimplemented")
}

// List implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) List(context.Context, *connect.Request[apiv2.TokenServiceListRequest]) (*connect.Response[apiv2.TokenServiceListResponse], error) {
	panic("unimplemented")
}

// Refresh implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Refresh(ctx context.Context, _ *connect.Request[apiv2.TokenServiceRefreshRequest]) (*connect.Response[apiv2.TokenServiceRefreshResponse], error) {
	token, err := generateToken(2 * time.Second)
	if err != nil {
		return nil, err
	}
	m.wasCalled = true
	return connect.NewResponse(&apiv2.TokenServiceRefreshResponse{Token: &apiv2.Token{}, Secret: token}), nil
}

// Revoke implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Revoke(context.Context, *connect.Request[apiv2.TokenServiceRevokeRequest]) (*connect.Response[apiv2.TokenServiceRevokeResponse], error) {
	panic("unimplemented")
}

// Update implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Update(context.Context, *connect.Request[apiv2.TokenServiceUpdateRequest]) (*connect.Response[apiv2.TokenServiceUpdateResponse], error) {
	panic("unimplemented")
}
