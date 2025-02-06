// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/admin/v2/partition.proto

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
	// PartitionServiceName is the fully-qualified name of the PartitionService service.
	PartitionServiceName = "metalstack.admin.v2.PartitionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PartitionServiceCreateProcedure is the fully-qualified name of the PartitionService's Create RPC.
	PartitionServiceCreateProcedure = "/metalstack.admin.v2.PartitionService/Create"
	// PartitionServiceUpdateProcedure is the fully-qualified name of the PartitionService's Update RPC.
	PartitionServiceUpdateProcedure = "/metalstack.admin.v2.PartitionService/Update"
	// PartitionServiceDeleteProcedure is the fully-qualified name of the PartitionService's Delete RPC.
	PartitionServiceDeleteProcedure = "/metalstack.admin.v2.PartitionService/Delete"
	// PartitionServiceCapacityProcedure is the fully-qualified name of the PartitionService's Capacity
	// RPC.
	PartitionServiceCapacityProcedure = "/metalstack.admin.v2.PartitionService/Capacity"
)

// PartitionServiceClient is a client for the metalstack.admin.v2.PartitionService service.
type PartitionServiceClient interface {
	// Create a partition
	Create(context.Context, *connect.Request[v2.PartitionServiceCreateRequest]) (*connect.Response[v2.PartitionServiceCreateResponse], error)
	// Update a partition
	Update(context.Context, *connect.Request[v2.PartitionServiceUpdateRequest]) (*connect.Response[v2.PartitionServiceUpdateResponse], error)
	// Delete a partition
	Delete(context.Context, *connect.Request[v2.PartitionServiceDeleteRequest]) (*connect.Response[v2.PartitionServiceDeleteResponse], error)
	// Capacity of a partitions
	Capacity(context.Context, *connect.Request[v2.PartitionServiceCapacityRequest]) (*connect.Response[v2.PartitionServiceCapacityResponse], error)
}

// NewPartitionServiceClient constructs a client for the metalstack.admin.v2.PartitionService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPartitionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PartitionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	partitionServiceMethods := v2.File_metalstack_admin_v2_partition_proto.Services().ByName("PartitionService").Methods()
	return &partitionServiceClient{
		create: connect.NewClient[v2.PartitionServiceCreateRequest, v2.PartitionServiceCreateResponse](
			httpClient,
			baseURL+PartitionServiceCreateProcedure,
			connect.WithSchema(partitionServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v2.PartitionServiceUpdateRequest, v2.PartitionServiceUpdateResponse](
			httpClient,
			baseURL+PartitionServiceUpdateProcedure,
			connect.WithSchema(partitionServiceMethods.ByName("Update")),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v2.PartitionServiceDeleteRequest, v2.PartitionServiceDeleteResponse](
			httpClient,
			baseURL+PartitionServiceDeleteProcedure,
			connect.WithSchema(partitionServiceMethods.ByName("Delete")),
			connect.WithClientOptions(opts...),
		),
		capacity: connect.NewClient[v2.PartitionServiceCapacityRequest, v2.PartitionServiceCapacityResponse](
			httpClient,
			baseURL+PartitionServiceCapacityProcedure,
			connect.WithSchema(partitionServiceMethods.ByName("Capacity")),
			connect.WithClientOptions(opts...),
		),
	}
}

// partitionServiceClient implements PartitionServiceClient.
type partitionServiceClient struct {
	create   *connect.Client[v2.PartitionServiceCreateRequest, v2.PartitionServiceCreateResponse]
	update   *connect.Client[v2.PartitionServiceUpdateRequest, v2.PartitionServiceUpdateResponse]
	delete   *connect.Client[v2.PartitionServiceDeleteRequest, v2.PartitionServiceDeleteResponse]
	capacity *connect.Client[v2.PartitionServiceCapacityRequest, v2.PartitionServiceCapacityResponse]
}

// Create calls metalstack.admin.v2.PartitionService.Create.
func (c *partitionServiceClient) Create(ctx context.Context, req *connect.Request[v2.PartitionServiceCreateRequest]) (*connect.Response[v2.PartitionServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls metalstack.admin.v2.PartitionService.Update.
func (c *partitionServiceClient) Update(ctx context.Context, req *connect.Request[v2.PartitionServiceUpdateRequest]) (*connect.Response[v2.PartitionServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls metalstack.admin.v2.PartitionService.Delete.
func (c *partitionServiceClient) Delete(ctx context.Context, req *connect.Request[v2.PartitionServiceDeleteRequest]) (*connect.Response[v2.PartitionServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Capacity calls metalstack.admin.v2.PartitionService.Capacity.
func (c *partitionServiceClient) Capacity(ctx context.Context, req *connect.Request[v2.PartitionServiceCapacityRequest]) (*connect.Response[v2.PartitionServiceCapacityResponse], error) {
	return c.capacity.CallUnary(ctx, req)
}

// PartitionServiceHandler is an implementation of the metalstack.admin.v2.PartitionService service.
type PartitionServiceHandler interface {
	// Create a partition
	Create(context.Context, *connect.Request[v2.PartitionServiceCreateRequest]) (*connect.Response[v2.PartitionServiceCreateResponse], error)
	// Update a partition
	Update(context.Context, *connect.Request[v2.PartitionServiceUpdateRequest]) (*connect.Response[v2.PartitionServiceUpdateResponse], error)
	// Delete a partition
	Delete(context.Context, *connect.Request[v2.PartitionServiceDeleteRequest]) (*connect.Response[v2.PartitionServiceDeleteResponse], error)
	// Capacity of a partitions
	Capacity(context.Context, *connect.Request[v2.PartitionServiceCapacityRequest]) (*connect.Response[v2.PartitionServiceCapacityResponse], error)
}

// NewPartitionServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPartitionServiceHandler(svc PartitionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	partitionServiceMethods := v2.File_metalstack_admin_v2_partition_proto.Services().ByName("PartitionService").Methods()
	partitionServiceCreateHandler := connect.NewUnaryHandler(
		PartitionServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(partitionServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	partitionServiceUpdateHandler := connect.NewUnaryHandler(
		PartitionServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(partitionServiceMethods.ByName("Update")),
		connect.WithHandlerOptions(opts...),
	)
	partitionServiceDeleteHandler := connect.NewUnaryHandler(
		PartitionServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(partitionServiceMethods.ByName("Delete")),
		connect.WithHandlerOptions(opts...),
	)
	partitionServiceCapacityHandler := connect.NewUnaryHandler(
		PartitionServiceCapacityProcedure,
		svc.Capacity,
		connect.WithSchema(partitionServiceMethods.ByName("Capacity")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.admin.v2.PartitionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PartitionServiceCreateProcedure:
			partitionServiceCreateHandler.ServeHTTP(w, r)
		case PartitionServiceUpdateProcedure:
			partitionServiceUpdateHandler.ServeHTTP(w, r)
		case PartitionServiceDeleteProcedure:
			partitionServiceDeleteHandler.ServeHTTP(w, r)
		case PartitionServiceCapacityProcedure:
			partitionServiceCapacityHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPartitionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPartitionServiceHandler struct{}

func (UnimplementedPartitionServiceHandler) Create(context.Context, *connect.Request[v2.PartitionServiceCreateRequest]) (*connect.Response[v2.PartitionServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.PartitionService.Create is not implemented"))
}

func (UnimplementedPartitionServiceHandler) Update(context.Context, *connect.Request[v2.PartitionServiceUpdateRequest]) (*connect.Response[v2.PartitionServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.PartitionService.Update is not implemented"))
}

func (UnimplementedPartitionServiceHandler) Delete(context.Context, *connect.Request[v2.PartitionServiceDeleteRequest]) (*connect.Response[v2.PartitionServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.PartitionService.Delete is not implemented"))
}

func (UnimplementedPartitionServiceHandler) Capacity(context.Context, *connect.Request[v2.PartitionServiceCapacityRequest]) (*connect.Response[v2.PartitionServiceCapacityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.PartitionService.Capacity is not implemented"))
}
