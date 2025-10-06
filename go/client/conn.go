package client

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	api "github.com/metal-stack/api/go/metalstack/api/v2"
)

const tokenRenewChecksDuringLifetime = 4

type (
	// DialConfig is the configuration to create a api-server connection
	DialConfig struct {
		BaseURL string
		Token   string
		Debug   bool

		UserAgent string
		// TokenRenewal defines if and how the token should be renewed
		TokenRenewal *TokenRenewal

		Transport http.RoundTripper

		Log *slog.Logger

		expiresAt time.Time
		issuedAt  time.Time
	}

	TokenRenewal struct {
		// PersistTokenFn is called to persist the newly fetched token
		// token will not be persisted if not specified
		PersistTokenFn PersistTokenFn
	}

	PersistTokenFn func(token string) error
)

func (d *DialConfig) HttpClient() *http.Client {
	transport := http.DefaultTransport
	if d.Transport != nil {
		transport = d.Transport
	}

	return &http.Client{
		Transport: &AddHeaderTransport{
			debug: d.Debug,
			t:     transport,
			token: d.Token,
		},
	}
}

func (dc *DialConfig) parse() error {
	if dc.Token == "" {
		return nil
	}
	parsed, err := jwt.Parse(dc.Token, nil)
	if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
		return fmt.Errorf("unable to parse token:%w", err)
	}
	expiresAt, err := parsed.Claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("unable to extract expiresAt from token:%w", err)
	}
	if expiresAt != nil {
		dc.expiresAt = expiresAt.Time
	}

	issuedAt, err := parsed.Claims.GetIssuedAt()
	if err != nil {
		return fmt.Errorf("unable to extract issuedAt from token:%w", err)
	}
	if issuedAt != nil {
		dc.issuedAt = issuedAt.Time
	}
	if dc.issuedAt.IsZero() {
		dc.issuedAt = time.Now()
	}
	return nil
}

func (c *client) startTokenRenewal() {
	if c.config.TokenRenewal == nil {
		return
	}
	if c.config.expiresAt.IsZero() {
		return
	}
	if c.config.Log == nil {
		c.config.Log = slog.Default()
	}

	replaceBefore := c.config.expiresAt.Sub(c.config.issuedAt) / tokenRenewChecksDuringLifetime

	err := c.renewTokenIfNeeded(replaceBefore)
	if err != nil {
		c.config.Log.Error("unable to renew token", "error", err)
	}

	ticker := time.NewTicker(replaceBefore)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.renewTokenIfNeeded(replaceBefore)
			if err != nil {
				c.config.Log.Error("unable to renew token", "error", err)
			}
		}
	}
}

func (c *client) renewTokenIfNeeded(replaceBefore time.Duration) error {
	if time.Until(c.config.expiresAt) > replaceBefore {
		return nil
	}
	c.config.Log.Info("call token refresh, current token expires soon", "expires", c.config.expiresAt.String())

	c.Lock()
	defer c.Unlock()

	resp, err := c.Apiv2().Token().Refresh(context.Background(), &api.TokenServiceRefreshRequest{})
	if err != nil {
		return fmt.Errorf("unable to refresh token %w", err)
	}

	c.config.Token = resp.Secret
	err = c.config.parse()
	if err != nil {
		return fmt.Errorf("unable to parse token %w", err)
	}

	if c.config.TokenRenewal.PersistTokenFn == nil {
		return nil
	}

	err = c.config.TokenRenewal.PersistTokenFn(c.config.Token)
	if err != nil {
		return fmt.Errorf("unable to persist token %w", err)
	}

	c.config.Log.Info("token refreshed, new token expires in", "expires", c.config.expiresAt.String())
	return nil
}

type AddHeaderTransport struct {
	debug bool

	token string
	t     http.RoundTripper
}

func (a *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	_, callInfo := connect.NewClientContext(req.Context())
	callInfo.RequestHeader().Add("Authorization", "Bearer "+a.token)

	if a.debug {
		reqDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Printf("DEBUG ERROR: %s\n", err)
		} else {
			fmt.Printf("DEBUG REQUEST:\n%s\n", string(reqDump))
		}
	}

	resp, err := a.t.RoundTrip(req)

	if a.debug && resp != nil {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Printf("DEBUG ERROR: %s\n", err)
		} else {
			fmt.Printf("DEBUG RESPONSE:\n%s\n", string(respDump))
		}
	}

	return resp, err
}
