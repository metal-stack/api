// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/api/v2/methods.proto

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
	// MethodServiceName is the fully-qualified name of the MethodService service.
	MethodServiceName = "metalstack.api.v2.MethodService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MethodServiceListProcedure is the fully-qualified name of the MethodService's List RPC.
	MethodServiceListProcedure = "/metalstack.api.v2.MethodService/List"
	// MethodServiceTokenScopedListProcedure is the fully-qualified name of the MethodService's
	// TokenScopedList RPC.
	MethodServiceTokenScopedListProcedure = "/metalstack.api.v2.MethodService/TokenScopedList"
)

// MethodServiceClient is a client for the metalstack.api.v2.MethodService service.
type MethodServiceClient interface {
	// List all public visible methods
	List(context.Context, *connect.Request[v2.MethodServiceListRequest]) (*connect.Response[v2.MethodServiceListResponse], error)
	// TokenScopedList all methods callable with the token present in the request
	TokenScopedList(context.Context, *connect.Request[v2.MethodServiceTokenScopedListRequest]) (*connect.Response[v2.MethodServiceTokenScopedListResponse], error)
}

// NewMethodServiceClient constructs a client for the metalstack.api.v2.MethodService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMethodServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MethodServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	methodServiceMethods := v2.File_metalstack_api_v2_methods_proto.Services().ByName("MethodService").Methods()
	return &methodServiceClient{
		list: connect.NewClient[v2.MethodServiceListRequest, v2.MethodServiceListResponse](
			httpClient,
			baseURL+MethodServiceListProcedure,
			connect.WithSchema(methodServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		tokenScopedList: connect.NewClient[v2.MethodServiceTokenScopedListRequest, v2.MethodServiceTokenScopedListResponse](
			httpClient,
			baseURL+MethodServiceTokenScopedListProcedure,
			connect.WithSchema(methodServiceMethods.ByName("TokenScopedList")),
			connect.WithClientOptions(opts...),
		),
	}
}

// methodServiceClient implements MethodServiceClient.
type methodServiceClient struct {
	list            *connect.Client[v2.MethodServiceListRequest, v2.MethodServiceListResponse]
	tokenScopedList *connect.Client[v2.MethodServiceTokenScopedListRequest, v2.MethodServiceTokenScopedListResponse]
}

// List calls metalstack.api.v2.MethodService.List.
func (c *methodServiceClient) List(ctx context.Context, req *connect.Request[v2.MethodServiceListRequest]) (*connect.Response[v2.MethodServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// TokenScopedList calls metalstack.api.v2.MethodService.TokenScopedList.
func (c *methodServiceClient) TokenScopedList(ctx context.Context, req *connect.Request[v2.MethodServiceTokenScopedListRequest]) (*connect.Response[v2.MethodServiceTokenScopedListResponse], error) {
	return c.tokenScopedList.CallUnary(ctx, req)
}

// MethodServiceHandler is an implementation of the metalstack.api.v2.MethodService service.
type MethodServiceHandler interface {
	// List all public visible methods
	List(context.Context, *connect.Request[v2.MethodServiceListRequest]) (*connect.Response[v2.MethodServiceListResponse], error)
	// TokenScopedList all methods callable with the token present in the request
	TokenScopedList(context.Context, *connect.Request[v2.MethodServiceTokenScopedListRequest]) (*connect.Response[v2.MethodServiceTokenScopedListResponse], error)
}

// NewMethodServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMethodServiceHandler(svc MethodServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	methodServiceMethods := v2.File_metalstack_api_v2_methods_proto.Services().ByName("MethodService").Methods()
	methodServiceListHandler := connect.NewUnaryHandler(
		MethodServiceListProcedure,
		svc.List,
		connect.WithSchema(methodServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	methodServiceTokenScopedListHandler := connect.NewUnaryHandler(
		MethodServiceTokenScopedListProcedure,
		svc.TokenScopedList,
		connect.WithSchema(methodServiceMethods.ByName("TokenScopedList")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.api.v2.MethodService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MethodServiceListProcedure:
			methodServiceListHandler.ServeHTTP(w, r)
		case MethodServiceTokenScopedListProcedure:
			methodServiceTokenScopedListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMethodServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMethodServiceHandler struct{}

func (UnimplementedMethodServiceHandler) List(context.Context, *connect.Request[v2.MethodServiceListRequest]) (*connect.Response[v2.MethodServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.MethodService.List is not implemented"))
}

func (UnimplementedMethodServiceHandler) TokenScopedList(context.Context, *connect.Request[v2.MethodServiceTokenScopedListRequest]) (*connect.Response[v2.MethodServiceTokenScopedListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.MethodService.TokenScopedList is not implemented"))
}
