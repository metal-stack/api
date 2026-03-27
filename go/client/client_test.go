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
	"testing/synctest"

	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/metal-stack/api/go/client"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	infrav2 "github.com/metal-stack/api/go/metalstack/infra/v2"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	var (
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	)

	synctest.Test(t, func(t *testing.T) {
		tokenString, err := generateToken(2 * time.Second)
		require.NoError(t, err)
		var renewedToken string

		c, err := client.New(&client.DialConfig{
			BaseURL: "http://localhost",
			Token:   tokenString,

			Interceptors: []connect.Interceptor{
				client.NewTestInterceptor(t, []client.ClientCall{
					{
						WantRequest: &apiv2.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv2.VersionServiceGetResponse{
								Version: &apiv2.Version{Version: "1.0"},
							})
						},
					},
					{
						WantRequest: &apiv2.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv2.VersionServiceGetResponse{
								Version: &apiv2.Version{Version: "1.0"},
							})
						},
					},
					{
						WantRequest: &apiv2.TokenServiceRefreshRequest{},
						WantResponse: func() connect.AnyResponse {
							tokenString, err := generateToken(2 * time.Second)
							require.NoError(t, err)

							return connect.NewResponse(&apiv2.TokenServiceRefreshResponse{
								Secret: tokenString,
							})
						},
					},
					{
						WantRequest: &apiv2.VersionServiceGetRequest{},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&apiv2.VersionServiceGetResponse{
								Version: &apiv2.Version{Version: "1.0"},
							})
						},
					},
				}),
			},
			TokenRenewal: &client.TokenRenewal{
				PersistTokenFn: func(token string) error {
					renewedToken = token
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
		require.Empty(t, renewedToken)

		time.Sleep(1 * time.Second)
		v, err = c.Apiv2().Version().Get(t.Context(), &apiv2.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
		require.Empty(t, renewedToken)

		time.Sleep(3 * time.Second)
		v, err = c.Apiv2().Version().Get(t.Context(), &apiv2.VersionServiceGetRequest{})
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "1.0", v.Version.Version)
		require.NotEmpty(t, renewedToken)
		require.NotEqual(t, renewedToken, tokenString, "haven't changed")
	})
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
	stream, err := c.Infrav2().BMC().WaitForBMCCommand(ctx, &infrav2.WaitForBMCCommandRequest{})
	require.NoError(t, err)
	require.NotNil(t, stream)
	require.Equal(t, tokenString, bs.token)
}

type mockBMCService struct {
	token string
}

func (m *mockBMCService) BMCCommandDone(context.Context, *infrav2.BMCCommandDoneRequest) (*infrav2.BMCCommandDoneResponse, error) {
	panic("unimplemented")
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

func (m *mockBMCService) WaitForBMCCommand(ctx context.Context, _ *infrav2.WaitForBMCCommandRequest, stream *connect.ServerStream[infrav2.WaitForBMCCommandResponse]) error {
	callinfo, _ := connect.CallInfoForHandlerContext(ctx)
	authHeader := callinfo.RequestHeader().Get("Authorization")

	_, token, found := strings.Cut(authHeader, "Bearer ")

	if !found {
		return fmt.Errorf("unable to extract token from header:%s", authHeader)
	}

	m.token = token
	return stream.Send(&infrav2.WaitForBMCCommandResponse{})
}
