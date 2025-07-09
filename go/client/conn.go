package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"connectrpc.com/connect"

	"github.com/golang-jwt/jwt/v5"
	api "github.com/metal-stack/api/go/metalstack/api/v2"
)

var defaultReplaceBefore = time.Minute

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

		expiresAt time.Time
	}

	TokenRenewal struct {
		// ReplaceBefore replaces the token with a new one if existing token only has this duration left until expiresAt
		// defaults to 1 minute if not specified
		ReplaceBefore *time.Duration
		// PersistTokenFn is called to persist the newly fetched token
		// token will not be persisted if not specified
		PersistTokenFn PersistTokenFn
	}

	PersistTokenFn func(token string) error
)

// TODO implement token refresh call based on the distance to expiresAt of the current token.
// The token must be parsed to known the exp timestamp.
// Token refresh is either called on interval if given, or on every request a cache is asked.
// PersistTokenFn is called after a new token was issued.

func (d *DialConfig) HttpClient() *http.Client {
	transport := http.DefaultTransport
	if d.Transport != nil {
		transport = d.Transport
	}

	if d.TokenRenewal != nil && d.TokenRenewal.ReplaceBefore == nil {
		d.TokenRenewal.ReplaceBefore = &defaultReplaceBefore
	}

	return &http.Client{
		Transport: &AddHeaderTransport{
			debug: d.Debug,
			t:     transport,
			token: d.Token,
		},
	}
}

func getExpiresAt(token string) (*time.Time, error) {
	parsed, err := jwt.Parse(token, nil)
	if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
		return nil, fmt.Errorf("unable to parse token:%w", err)
	}
	expiresAt, err := parsed.Claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("unable to extract expiresAt from token:%w", err)
	}
	return &expiresAt.Time, nil
}

func (c client) renewToken() {
	if c.config.TokenRenewal == nil {
		return
	}
	if c.config.expiresAt.IsZero() {
		return
	}
	if time.Since(c.config.expiresAt).Abs() > *c.config.TokenRenewal.ReplaceBefore {
		return
	}

	resp, err := c.Apiv2().Token().Refresh(context.Background(), connect.NewRequest(&api.TokenServiceRefreshRequest{}))
	if err != nil {
		// TODO howto handle errors
		fmt.Printf("unable to refresh token:%v\n", err)
	}
	token := resp.Msg.Secret
	exp, err := getExpiresAt(token)
	if err != nil {
		fmt.Printf("unable to parse token:%v\n", err)
	}
	c.config.Token = token
	c.config.expiresAt = *exp

	if c.config.TokenRenewal.PersistTokenFn == nil {
		return
	}

	err = c.config.TokenRenewal.PersistTokenFn(token)
	if err != nil {
		fmt.Printf("unable to persist token:%v\n", err)
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
