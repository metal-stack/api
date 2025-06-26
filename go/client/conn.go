package client

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type (
	// DialConfig is the configuration to create a api-server connection
	DialConfig struct {
		BaseURL string
		Token   string
		Debug   bool

		UserAgent string
		// TokenRenewal defines if and how the token should be renewed
		TokenRenewal *TokenRenewal
	}

	TokenRenewal struct {
		// Interval replaces the token with a new one on the given interval
		Interval *time.Duration
		// PersistTokenFn is called to persist the newly fetched token
		PersistTokenFn PersistTokenFn
	}

	PersistTokenFn func(token string) error
)

// TODO implement token refresh call based on the distance to expiresAt of the current token.
// The token must be parsed to known the exp timestamp.
// Token refresh is either called on interval if given, or on every request a cache is asked.
// PersistTokenFn is called after a new token was issued.

func (d *DialConfig) HttpClient() *http.Client {
	return &http.Client{
		Transport: &AddHeaderTransport{
			debug: d.Debug,
			t:     http.DefaultTransport,
			token: d.Token,
		},
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
