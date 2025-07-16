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

const defaultReplaceBefore = time.Hour

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
	}

	TokenRenewal struct {
		// ReplaceBefore replaces the token with a new one if existing token only has this duration left until expiresAt
		// defaults to 1 hour if not specified
		ReplaceBefore time.Duration
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

	if d.TokenRenewal != nil && d.TokenRenewal.ReplaceBefore == 0 {
		d.TokenRenewal.ReplaceBefore = defaultReplaceBefore
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
	ticker := time.NewTicker(c.config.TokenRenewal.ReplaceBefore / 2)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if time.Until(c.config.expiresAt) > c.config.TokenRenewal.ReplaceBefore {
				continue
			}

			c.Lock()
			defer c.Unlock()

			resp, err := c.Apiv2().Token().Refresh(context.Background(), connect.NewRequest(&api.TokenServiceRefreshRequest{}))
			if err != nil {
				c.config.Log.Error("unable to refresh token", "error", err)
				continue
			}

			c.config.Token = resp.Msg.Secret
			err = c.config.parse()
			if err != nil {
				c.config.Log.Error("unable to parse token", "error", err)
				continue
			}

			if c.config.TokenRenewal.PersistTokenFn == nil {
				return
			}

			err = c.config.TokenRenewal.PersistTokenFn(c.config.Token)
			if err != nil {
				c.config.Log.Error("unable to persist token", "error", err)
				continue
			}
		}
	}
}

type AddHeaderTransport struct {
	debug bool

	token string
	t     http.RoundTripper
}

func (a *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+a.token)

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
