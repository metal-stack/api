// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/api/v2/token.proto

package apiv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "github.com/metal-stack/api/go/metalstack/api/v2"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TokenServiceName is the fully-qualified name of the TokenService service.
	TokenServiceName = "metalstack.api.v2.TokenService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TokenServiceGetProcedure is the fully-qualified name of the TokenService's Get RPC.
	TokenServiceGetProcedure = "/metalstack.api.v2.TokenService/Get"
	// TokenServiceCreateProcedure is the fully-qualified name of the TokenService's Create RPC.
	TokenServiceCreateProcedure = "/metalstack.api.v2.TokenService/Create"
	// TokenServiceUpdateProcedure is the fully-qualified name of the TokenService's Update RPC.
	TokenServiceUpdateProcedure = "/metalstack.api.v2.TokenService/Update"
	// TokenServiceListProcedure is the fully-qualified name of the TokenService's List RPC.
	TokenServiceListProcedure = "/metalstack.api.v2.TokenService/List"
	// TokenServiceRevokeProcedure is the fully-qualified name of the TokenService's Revoke RPC.
	TokenServiceRevokeProcedure = "/metalstack.api.v2.TokenService/Revoke"
	// TokenServiceRefreshProcedure is the fully-qualified name of the TokenService's Refresh RPC.
	TokenServiceRefreshProcedure = "/metalstack.api.v2.TokenService/Refresh"
)

// TokenServiceClient is a client for the metalstack.api.v2.TokenService service.
type TokenServiceClient interface {
	// Get a token
	Get(context.Context, *connect.Request[v2.TokenServiceGetRequest]) (*connect.Response[v2.TokenServiceGetResponse], error)
	// Create a token to authenticate against the platform, the secret will be only visible in the response
	Create(context.Context, *connect.Request[v2.TokenServiceCreateRequest]) (*connect.Response[v2.TokenServiceCreateResponse], error)
	// Update a token
	Update(context.Context, *connect.Request[v2.TokenServiceUpdateRequest]) (*connect.Response[v2.TokenServiceUpdateResponse], error)
	// List all your tokens
	List(context.Context, *connect.Request[v2.TokenServiceListRequest]) (*connect.Response[v2.TokenServiceListResponse], error)
	// Revoke a token, no further usage is possible afterwards
	Revoke(context.Context, *connect.Request[v2.TokenServiceRevokeRequest]) (*connect.Response[v2.TokenServiceRevokeResponse], error)
	// Refresh a token, this will create a new token with the exact same permissions as the calling token contains
	Refresh(context.Context, *connect.Request[v2.TokenServiceRefreshRequest]) (*connect.Response[v2.TokenServiceRefreshResponse], error)
}

// NewTokenServiceClient constructs a client for the metalstack.api.v2.TokenService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTokenServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TokenServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	tokenServiceMethods := v2.File_metalstack_api_v2_token_proto.Services().ByName("TokenService").Methods()
	return &tokenServiceClient{
		get: connect.NewClient[v2.TokenServiceGetRequest, v2.TokenServiceGetResponse](
			httpClient,
			baseURL+TokenServiceGetProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
		create: connect.NewClient[v2.TokenServiceCreateRequest, v2.TokenServiceCreateResponse](
			httpClient,
			baseURL+TokenServiceCreateProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v2.TokenServiceUpdateRequest, v2.TokenServiceUpdateResponse](
			httpClient,
			baseURL+TokenServiceUpdateProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("Update")),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v2.TokenServiceListRequest, v2.TokenServiceListResponse](
			httpClient,
			baseURL+TokenServiceListProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		revoke: connect.NewClient[v2.TokenServiceRevokeRequest, v2.TokenServiceRevokeResponse](
			httpClient,
			baseURL+TokenServiceRevokeProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("Revoke")),
			connect.WithClientOptions(opts...),
		),
		refresh: connect.NewClient[v2.TokenServiceRefreshRequest, v2.TokenServiceRefreshResponse](
			httpClient,
			baseURL+TokenServiceRefreshProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("Refresh")),
			connect.WithClientOptions(opts...),
		),
	}
}

// tokenServiceClient implements TokenServiceClient.
type tokenServiceClient struct {
	get     *connect.Client[v2.TokenServiceGetRequest, v2.TokenServiceGetResponse]
	create  *connect.Client[v2.TokenServiceCreateRequest, v2.TokenServiceCreateResponse]
	update  *connect.Client[v2.TokenServiceUpdateRequest, v2.TokenServiceUpdateResponse]
	list    *connect.Client[v2.TokenServiceListRequest, v2.TokenServiceListResponse]
	revoke  *connect.Client[v2.TokenServiceRevokeRequest, v2.TokenServiceRevokeResponse]
	refresh *connect.Client[v2.TokenServiceRefreshRequest, v2.TokenServiceRefreshResponse]
}

// Get calls metalstack.api.v2.TokenService.Get.
func (c *tokenServiceClient) Get(ctx context.Context, req *connect.Request[v2.TokenServiceGetRequest]) (*connect.Response[v2.TokenServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Create calls metalstack.api.v2.TokenService.Create.
func (c *tokenServiceClient) Create(ctx context.Context, req *connect.Request[v2.TokenServiceCreateRequest]) (*connect.Response[v2.TokenServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls metalstack.api.v2.TokenService.Update.
func (c *tokenServiceClient) Update(ctx context.Context, req *connect.Request[v2.TokenServiceUpdateRequest]) (*connect.Response[v2.TokenServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// List calls metalstack.api.v2.TokenService.List.
func (c *tokenServiceClient) List(ctx context.Context, req *connect.Request[v2.TokenServiceListRequest]) (*connect.Response[v2.TokenServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Revoke calls metalstack.api.v2.TokenService.Revoke.
func (c *tokenServiceClient) Revoke(ctx context.Context, req *connect.Request[v2.TokenServiceRevokeRequest]) (*connect.Response[v2.TokenServiceRevokeResponse], error) {
	return c.revoke.CallUnary(ctx, req)
}

// Refresh calls metalstack.api.v2.TokenService.Refresh.
func (c *tokenServiceClient) Refresh(ctx context.Context, req *connect.Request[v2.TokenServiceRefreshRequest]) (*connect.Response[v2.TokenServiceRefreshResponse], error) {
	return c.refresh.CallUnary(ctx, req)
}

// TokenServiceHandler is an implementation of the metalstack.api.v2.TokenService service.
type TokenServiceHandler interface {
	// Get a token
	Get(context.Context, *connect.Request[v2.TokenServiceGetRequest]) (*connect.Response[v2.TokenServiceGetResponse], error)
	// Create a token to authenticate against the platform, the secret will be only visible in the response
	Create(context.Context, *connect.Request[v2.TokenServiceCreateRequest]) (*connect.Response[v2.TokenServiceCreateResponse], error)
	// Update a token
	Update(context.Context, *connect.Request[v2.TokenServiceUpdateRequest]) (*connect.Response[v2.TokenServiceUpdateResponse], error)
	// List all your tokens
	List(context.Context, *connect.Request[v2.TokenServiceListRequest]) (*connect.Response[v2.TokenServiceListResponse], error)
	// Revoke a token, no further usage is possible afterwards
	Revoke(context.Context, *connect.Request[v2.TokenServiceRevokeRequest]) (*connect.Response[v2.TokenServiceRevokeResponse], error)
	// Refresh a token, this will create a new token with the exact same permissions as the calling token contains
	Refresh(context.Context, *connect.Request[v2.TokenServiceRefreshRequest]) (*connect.Response[v2.TokenServiceRefreshResponse], error)
}

// NewTokenServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTokenServiceHandler(svc TokenServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tokenServiceMethods := v2.File_metalstack_api_v2_token_proto.Services().ByName("TokenService").Methods()
	tokenServiceGetHandler := connect.NewUnaryHandler(
		TokenServiceGetProcedure,
		svc.Get,
		connect.WithSchema(tokenServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceCreateHandler := connect.NewUnaryHandler(
		TokenServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(tokenServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceUpdateHandler := connect.NewUnaryHandler(
		TokenServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(tokenServiceMethods.ByName("Update")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceListHandler := connect.NewUnaryHandler(
		TokenServiceListProcedure,
		svc.List,
		connect.WithSchema(tokenServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceRevokeHandler := connect.NewUnaryHandler(
		TokenServiceRevokeProcedure,
		svc.Revoke,
		connect.WithSchema(tokenServiceMethods.ByName("Revoke")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceRefreshHandler := connect.NewUnaryHandler(
		TokenServiceRefreshProcedure,
		svc.Refresh,
		connect.WithSchema(tokenServiceMethods.ByName("Refresh")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.api.v2.TokenService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TokenServiceGetProcedure:
			tokenServiceGetHandler.ServeHTTP(w, r)
		case TokenServiceCreateProcedure:
			tokenServiceCreateHandler.ServeHTTP(w, r)
		case TokenServiceUpdateProcedure:
			tokenServiceUpdateHandler.ServeHTTP(w, r)
		case TokenServiceListProcedure:
			tokenServiceListHandler.ServeHTTP(w, r)
		case TokenServiceRevokeProcedure:
			tokenServiceRevokeHandler.ServeHTTP(w, r)
		case TokenServiceRefreshProcedure:
			tokenServiceRefreshHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTokenServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTokenServiceHandler struct{}

func (UnimplementedTokenServiceHandler) Get(context.Context, *connect.Request[v2.TokenServiceGetRequest]) (*connect.Response[v2.TokenServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.Get is not implemented"))
}

func (UnimplementedTokenServiceHandler) Create(context.Context, *connect.Request[v2.TokenServiceCreateRequest]) (*connect.Response[v2.TokenServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.Create is not implemented"))
}

func (UnimplementedTokenServiceHandler) Update(context.Context, *connect.Request[v2.TokenServiceUpdateRequest]) (*connect.Response[v2.TokenServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.Update is not implemented"))
}

func (UnimplementedTokenServiceHandler) List(context.Context, *connect.Request[v2.TokenServiceListRequest]) (*connect.Response[v2.TokenServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.List is not implemented"))
}

func (UnimplementedTokenServiceHandler) Revoke(context.Context, *connect.Request[v2.TokenServiceRevokeRequest]) (*connect.Response[v2.TokenServiceRevokeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.Revoke is not implemented"))
}

func (UnimplementedTokenServiceHandler) Refresh(context.Context, *connect.Request[v2.TokenServiceRefreshRequest]) (*connect.Response[v2.TokenServiceRefreshResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.TokenService.Refresh is not implemented"))
}
