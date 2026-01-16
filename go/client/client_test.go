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
	infrav2 "github.com/metal-stack/api/go/metalstack/infra/v2"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	var (
		vs  = &mockVersionService{}
		ts  = &mockTokenService{}
		mux = http.NewServeMux()
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	)

	mux.Handle(apiv2connect.NewVersionServiceHandler(vs))
	mux.Handle(apiv2connect.NewTokenServiceHandler(ts))
	server := httptest.NewTLSServer(mux)
	server.EnableHTTP2 = true
	defer func() {
		server.Close()
	}()

	tokenString, err := generateToken(2 * time.Second)
	require.NoError(t, err)

	c, err := client.New(&client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
		TokenRenewal: &client.TokenRenewal{
			PersistTokenFn: func(token string) error {
				ts.token = token
				t.Log("token persisted:", token)
				return nil
			},
		},
		Log: log,
	})
	require.NoError(t, err)
	v, err := c.Apiv2().Version().Get(t.Context(), &apiv2.VersionServiceGetRequest{})
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Version.Version)
	require.False(t, ts.wasCalled)
	require.Equal(t, tokenString, vs.token)

	time.Sleep(1 * time.Second)
	v, err = c.Apiv2().Version().Get(t.Context(), &apiv2.VersionServiceGetRequest{})
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Version.Version)
	require.False(t, ts.wasCalled)
	require.Equal(t, tokenString, vs.token)

	time.Sleep(1 * time.Second)
	v, err = c.Apiv2().Version().Get(t.Context(), &apiv2.VersionServiceGetRequest{})
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Version.Version)

	require.True(t, ts.wasCalled)
	require.NotEqual(t, tokenString, ts.token, "token must have changed")
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

func (m *mockVersionService) Get(ctx context.Context, req *apiv2.VersionServiceGetRequest) (*apiv2.VersionServiceGetResponse, error) {
	callinfo, _ := connect.CallInfoForHandlerContext(ctx)
	authHeader := callinfo.RequestHeader().Get("Authorization")

	_, token, found := strings.Cut(authHeader, "Bearer ")

	if !found {
		return nil, fmt.Errorf("unable to extract token from header:%s", authHeader)
	}

	m.token = token
	return &apiv2.VersionServiceGetResponse{Version: &apiv2.Version{Version: "1.0"}}, nil
}

type mockTokenService struct {
	wasCalled bool
	token     string
}

// Create implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Create(context.Context, *apiv2.TokenServiceCreateRequest) (*apiv2.TokenServiceCreateResponse, error) {
	panic("unimplemented")
}

// Get implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Get(context.Context, *apiv2.TokenServiceGetRequest) (*apiv2.TokenServiceGetResponse, error) {
	panic("unimplemented")
}

// List implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) List(context.Context, *apiv2.TokenServiceListRequest) (*apiv2.TokenServiceListResponse, error) {
	panic("unimplemented")
}

// Refresh implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Refresh(ctx context.Context, _ *apiv2.TokenServiceRefreshRequest) (*apiv2.TokenServiceRefreshResponse, error) {
	token, err := generateToken(2 * time.Second)
	if err != nil {
		return nil, err
	}
	m.wasCalled = true
	return &apiv2.TokenServiceRefreshResponse{Token: &apiv2.Token{}, Secret: token}, nil
}

// Revoke implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Revoke(context.Context, *apiv2.TokenServiceRevokeRequest) (*apiv2.TokenServiceRevokeResponse, error) {
	panic("unimplemented")
}

// Update implements apiv2connect.TokenServiceHandler.
func (m *mockTokenService) Update(context.Context, *apiv2.TokenServiceUpdateRequest) (*apiv2.TokenServiceUpdateResponse, error) {
	panic("unimplemented")
}

func Test_ClientInterceptors(t *testing.T) {
	var (
		bs  = &mockBMCService{}
		mux = http.NewServeMux()
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		ctx = t.Context()
	)

	mux.Handle(infrav2connect.NewBMCServiceHandler(bs))
	server := httptest.NewTLSServer(mux)
	server.EnableHTTP2 = true
	defer func() {
		server.Close()
	}()

	tokenString, err := generateToken(1 * time.Second)
	require.NoError(t, err)

	c, err := client.New(&client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
		Log:       log,
	})
	require.NoError(t, err)

	// Synchronous call has authheader set
	resp, err := c.Infrav2().BMC().UpdateBMCInfo(ctx, &infrav2.UpdateBMCInfoRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, tokenString, bs.token)
	bs.token = ""

	// Asynchronous call has authheader set
	stream, err := c.Infrav2().BMC().WaitForMachineEvent(ctx, &infrav2.WaitForMachineEventRequest{})
	require.NoError(t, err)
	require.NotNil(t, stream)
	require.Equal(t, tokenString, bs.token)
}

type mockBMCService struct {
	token string
}

func (m *mockBMCService) UpdateBMCInfo(ctx context.Context, _ *infrav2.UpdateBMCInfoRequest) (*infrav2.UpdateBMCInfoResponse, error) {
	callinfo, _ := connect.CallInfoForHandlerContext(ctx)
	authHeader := callinfo.RequestHeader().Get("Authorization")

	_, token, found := strings.Cut(authHeader, "Bearer ")

	if !found {
		return nil, fmt.Errorf("unable to extract token from header:%s", authHeader)
	}
	m.token = token
	return &infrav2.UpdateBMCInfoResponse{}, nil
}

func (m *mockBMCService) WaitForMachineEvent(ctx context.Context, _ *infrav2.WaitForMachineEventRequest, stream *connect.ServerStream[infrav2.WaitForMachineEventResponse]) error {
	callinfo, _ := connect.CallInfoForHandlerContext(ctx)
	authHeader := callinfo.RequestHeader().Get("Authorization")

	_, token, found := strings.Cut(authHeader, "Bearer ")

	if !found {
		return fmt.Errorf("unable to extract token from header:%s", authHeader)
	}

	m.token = token
	return stream.Send(&infrav2.WaitForMachineEventResponse{})
}
