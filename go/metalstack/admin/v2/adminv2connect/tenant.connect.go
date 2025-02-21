// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/admin/v2/tenant.proto

package adminv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "github.com/metal-stack/api/go/metalstack/admin/v2"
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
	// TenantServiceName is the fully-qualified name of the TenantService service.
	TenantServiceName = "metalstack.admin.v2.TenantService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TenantServiceCreateProcedure is the fully-qualified name of the TenantService's Create RPC.
	TenantServiceCreateProcedure = "/metalstack.admin.v2.TenantService/Create"
	// TenantServiceListProcedure is the fully-qualified name of the TenantService's List RPC.
	TenantServiceListProcedure = "/metalstack.admin.v2.TenantService/List"
)

// TenantServiceClient is a client for the metalstack.admin.v2.TenantService service.
type TenantServiceClient interface {
	// Create a tenant
	Create(context.Context, *connect.Request[v2.TenantServiceCreateRequest]) (*connect.Response[v2.TenantServiceCreateResponse], error)
	// List all tenants
	List(context.Context, *connect.Request[v2.TenantServiceListRequest]) (*connect.Response[v2.TenantServiceListResponse], error)
}

// NewTenantServiceClient constructs a client for the metalstack.admin.v2.TenantService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTenantServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TenantServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	tenantServiceMethods := v2.File_metalstack_admin_v2_tenant_proto.Services().ByName("TenantService").Methods()
	return &tenantServiceClient{
		create: connect.NewClient[v2.TenantServiceCreateRequest, v2.TenantServiceCreateResponse](
			httpClient,
			baseURL+TenantServiceCreateProcedure,
			connect.WithSchema(tenantServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v2.TenantServiceListRequest, v2.TenantServiceListResponse](
			httpClient,
			baseURL+TenantServiceListProcedure,
			connect.WithSchema(tenantServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
	}
}

// tenantServiceClient implements TenantServiceClient.
type tenantServiceClient struct {
	create *connect.Client[v2.TenantServiceCreateRequest, v2.TenantServiceCreateResponse]
	list   *connect.Client[v2.TenantServiceListRequest, v2.TenantServiceListResponse]
}

// Create calls metalstack.admin.v2.TenantService.Create.
func (c *tenantServiceClient) Create(ctx context.Context, req *connect.Request[v2.TenantServiceCreateRequest]) (*connect.Response[v2.TenantServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// List calls metalstack.admin.v2.TenantService.List.
func (c *tenantServiceClient) List(ctx context.Context, req *connect.Request[v2.TenantServiceListRequest]) (*connect.Response[v2.TenantServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// TenantServiceHandler is an implementation of the metalstack.admin.v2.TenantService service.
type TenantServiceHandler interface {
	// Create a tenant
	Create(context.Context, *connect.Request[v2.TenantServiceCreateRequest]) (*connect.Response[v2.TenantServiceCreateResponse], error)
	// List all tenants
	List(context.Context, *connect.Request[v2.TenantServiceListRequest]) (*connect.Response[v2.TenantServiceListResponse], error)
}

// NewTenantServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTenantServiceHandler(svc TenantServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tenantServiceMethods := v2.File_metalstack_admin_v2_tenant_proto.Services().ByName("TenantService").Methods()
	tenantServiceCreateHandler := connect.NewUnaryHandler(
		TenantServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(tenantServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	tenantServiceListHandler := connect.NewUnaryHandler(
		TenantServiceListProcedure,
		svc.List,
		connect.WithSchema(tenantServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.admin.v2.TenantService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TenantServiceCreateProcedure:
			tenantServiceCreateHandler.ServeHTTP(w, r)
		case TenantServiceListProcedure:
			tenantServiceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTenantServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTenantServiceHandler struct{}

func (UnimplementedTenantServiceHandler) Create(context.Context, *connect.Request[v2.TenantServiceCreateRequest]) (*connect.Response[v2.TenantServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.TenantService.Create is not implemented"))
}

func (UnimplementedTenantServiceHandler) List(context.Context, *connect.Request[v2.TenantServiceListRequest]) (*connect.Response[v2.TenantServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.TenantService.List is not implemented"))
}
