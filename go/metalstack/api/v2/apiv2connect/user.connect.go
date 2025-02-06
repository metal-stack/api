// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/api/v2/user.proto

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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "metalstack.api.v2.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetProcedure is the fully-qualified name of the UserService's Get RPC.
	UserServiceGetProcedure = "/metalstack.api.v2.UserService/Get"
)

// UserServiceClient is a client for the metalstack.api.v2.UserService service.
type UserServiceClient interface {
	// Get a User
	Get(context.Context, *connect.Request[v2.UserServiceGetRequest]) (*connect.Response[v2.UserServiceGetResponse], error)
}

// NewUserServiceClient constructs a client for the metalstack.api.v2.UserService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	userServiceMethods := v2.File_metalstack_api_v2_user_proto.Services().ByName("UserService").Methods()
	return &userServiceClient{
		get: connect.NewClient[v2.UserServiceGetRequest, v2.UserServiceGetResponse](
			httpClient,
			baseURL+UserServiceGetProcedure,
			connect.WithSchema(userServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	get *connect.Client[v2.UserServiceGetRequest, v2.UserServiceGetResponse]
}

// Get calls metalstack.api.v2.UserService.Get.
func (c *userServiceClient) Get(ctx context.Context, req *connect.Request[v2.UserServiceGetRequest]) (*connect.Response[v2.UserServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the metalstack.api.v2.UserService service.
type UserServiceHandler interface {
	// Get a User
	Get(context.Context, *connect.Request[v2.UserServiceGetRequest]) (*connect.Response[v2.UserServiceGetResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceMethods := v2.File_metalstack_api_v2_user_proto.Services().ByName("UserService").Methods()
	userServiceGetHandler := connect.NewUnaryHandler(
		UserServiceGetProcedure,
		svc.Get,
		connect.WithSchema(userServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.api.v2.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetProcedure:
			userServiceGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) Get(context.Context, *connect.Request[v2.UserServiceGetRequest]) (*connect.Response[v2.UserServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.UserService.Get is not implemented"))
}
