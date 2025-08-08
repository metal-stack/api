package client_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/metal-stack/api/go/client"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	"github.com/stretchr/testify/require"
)

func TestNewFilesystemTokenPersiter(t *testing.T) {
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

	tokenDir, err := os.MkdirTemp("/tmp", "token-")
	require.NoError(t, err)
	tokenPath := path.Join(tokenDir, "token")

	_, err = os.OpenFile(tokenPath, os.O_RDONLY|os.O_CREATE, 0600)
	require.NoError(t, err)

	filesystemPersister, err := client.NewFilesystemTokenPersiter(tokenPath)
	require.NoError(t, err)

	c, err := client.New(&client.DialConfig{
		BaseURL:   server.URL,
		Token:     tokenString,
		Transport: server.Client().Transport,
		TokenRenewal: &client.TokenRenewal{
			PersistTokenFn: filesystemPersister,
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

	time.Sleep(2 * time.Second)
	v, err = c.Apiv2().Version().Get(t.Context(), connect.NewRequest(&apiv2.VersionServiceGetRequest{}))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, "1.0", v.Msg.Version.Version)
	require.True(t, ts.wasCalled)
	require.NotEqual(t, tokenString, vs.token, "token must have changed")

	tokenBytes, err := os.ReadFile(tokenPath)
	require.NoError(t, err)
	require.Equal(t, string(tokenBytes), vs.token)

}
