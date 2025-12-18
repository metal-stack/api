// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"context"
	"sync"

	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"

{{ range $name, $api := . -}}
	"github.com/metal-stack/api/go{{ $api.Path }}/{{ $api.Name }}connect"
{{ end }}
)

type (
	Client interface {
{{ range $name, $api := . -}}
	{{ $name | title }}() {{ $name | title }}
{{ end }}
	}
	client struct {
		config *DialConfig

		interceptors []connect.Interceptor

		sync.Mutex
	}
{{ range $name, $api := . -}}
	{{ $name | title }} interface {
{{ range $svc := $api.Services -}}
	{{ $svc | trimSuffix "Service" }}() {{ $name }}connect.{{ $svc }}Client
{{ end }}
	}

    {{ $name }} struct {
{{ range $svc := $api.Services -}}
	{{ $svc | lower }} {{ $name }}connect.{{ $svc }}Client
{{ end }}
    }

{{ end }}
)

func New(config *DialConfig) (Client, error) {
	err := config.parse()
	if err != nil {
		return nil, err
	}

	c := &client{
		config:       config,
		interceptors: []connect.Interceptor{},
	}

	authInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			request.Header().Add("Authorization", "Bearer "+config.Token)
			return next(ctx, request)
		})
	})

	loggingInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "body", request.Any())
			response, err := next(ctx, request)
			if err != nil {
				return nil, err
			}
			config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "response", response.Any())
			return response, err
		})
	})

	if config.Token != "" {
		c.interceptors = append(c.interceptors, authInterceptor)
	}
	if config.Log != nil {
		c.interceptors = append(c.interceptors, loggingInterceptor)
	}
	c.interceptors = append(c.interceptors, config.Interceptors...)

	go c.startTokenRenewal()

	return c, nil
}

{{ range $name, $api := . -}}
func (c *client) {{ $name | title }}() {{ $name | title }} {
	a := &{{ $name }}{
{{ range $svc := $api.Services -}}
	{{ $svc | lower }}:  {{ $name }}connect.New{{ $svc }}Client(
		c.config.HttpClient(),
		c.config.BaseURL,
		connect.WithInterceptors(c.interceptors...),
		compress.WithAll(compress.LevelBalanced),
	),
{{ end }}
	}
	return a
}

{{ range $svc := $api.Services -}}
func (c  *{{ $name }} ) {{ $svc | trimSuffix "Service" }}() {{ $name }}connect.{{ $svc }}Client {
	return c.{{ $svc | lower }}
}
{{ end }}

{{ end }}
