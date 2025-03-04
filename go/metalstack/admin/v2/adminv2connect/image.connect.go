// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/admin/v2/image.proto

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
	// ImageServiceName is the fully-qualified name of the ImageService service.
	ImageServiceName = "metalstack.admin.v2.ImageService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ImageServiceCreateProcedure is the fully-qualified name of the ImageService's Create RPC.
	ImageServiceCreateProcedure = "/metalstack.admin.v2.ImageService/Create"
	// ImageServiceUpdateProcedure is the fully-qualified name of the ImageService's Update RPC.
	ImageServiceUpdateProcedure = "/metalstack.admin.v2.ImageService/Update"
	// ImageServiceDeleteProcedure is the fully-qualified name of the ImageService's Delete RPC.
	ImageServiceDeleteProcedure = "/metalstack.admin.v2.ImageService/Delete"
	// ImageServiceUsageProcedure is the fully-qualified name of the ImageService's Usage RPC.
	ImageServiceUsageProcedure = "/metalstack.admin.v2.ImageService/Usage"
)

// ImageServiceClient is a client for the metalstack.admin.v2.ImageService service.
type ImageServiceClient interface {
	// Create a image
	Create(context.Context, *connect.Request[v2.ImageServiceCreateRequest]) (*connect.Response[v2.ImageServiceCreateResponse], error)
	// Update a image
	Update(context.Context, *connect.Request[v2.ImageServiceUpdateRequest]) (*connect.Response[v2.ImageServiceUpdateResponse], error)
	// Delete a image
	Delete(context.Context, *connect.Request[v2.ImageServiceDeleteRequest]) (*connect.Response[v2.ImageServiceDeleteResponse], error)
	// Usage of images
	Usage(context.Context, *connect.Request[v2.ImageServiceUsageRequest]) (*connect.Response[v2.ImageServiceUsageResponse], error)
}

// NewImageServiceClient constructs a client for the metalstack.admin.v2.ImageService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewImageServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ImageServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	imageServiceMethods := v2.File_metalstack_admin_v2_image_proto.Services().ByName("ImageService").Methods()
	return &imageServiceClient{
		create: connect.NewClient[v2.ImageServiceCreateRequest, v2.ImageServiceCreateResponse](
			httpClient,
			baseURL+ImageServiceCreateProcedure,
			connect.WithSchema(imageServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v2.ImageServiceUpdateRequest, v2.ImageServiceUpdateResponse](
			httpClient,
			baseURL+ImageServiceUpdateProcedure,
			connect.WithSchema(imageServiceMethods.ByName("Update")),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v2.ImageServiceDeleteRequest, v2.ImageServiceDeleteResponse](
			httpClient,
			baseURL+ImageServiceDeleteProcedure,
			connect.WithSchema(imageServiceMethods.ByName("Delete")),
			connect.WithClientOptions(opts...),
		),
		usage: connect.NewClient[v2.ImageServiceUsageRequest, v2.ImageServiceUsageResponse](
			httpClient,
			baseURL+ImageServiceUsageProcedure,
			connect.WithSchema(imageServiceMethods.ByName("Usage")),
			connect.WithClientOptions(opts...),
		),
	}
}

// imageServiceClient implements ImageServiceClient.
type imageServiceClient struct {
	create *connect.Client[v2.ImageServiceCreateRequest, v2.ImageServiceCreateResponse]
	update *connect.Client[v2.ImageServiceUpdateRequest, v2.ImageServiceUpdateResponse]
	delete *connect.Client[v2.ImageServiceDeleteRequest, v2.ImageServiceDeleteResponse]
	usage  *connect.Client[v2.ImageServiceUsageRequest, v2.ImageServiceUsageResponse]
}

// Create calls metalstack.admin.v2.ImageService.Create.
func (c *imageServiceClient) Create(ctx context.Context, req *connect.Request[v2.ImageServiceCreateRequest]) (*connect.Response[v2.ImageServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls metalstack.admin.v2.ImageService.Update.
func (c *imageServiceClient) Update(ctx context.Context, req *connect.Request[v2.ImageServiceUpdateRequest]) (*connect.Response[v2.ImageServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls metalstack.admin.v2.ImageService.Delete.
func (c *imageServiceClient) Delete(ctx context.Context, req *connect.Request[v2.ImageServiceDeleteRequest]) (*connect.Response[v2.ImageServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Usage calls metalstack.admin.v2.ImageService.Usage.
func (c *imageServiceClient) Usage(ctx context.Context, req *connect.Request[v2.ImageServiceUsageRequest]) (*connect.Response[v2.ImageServiceUsageResponse], error) {
	return c.usage.CallUnary(ctx, req)
}

// ImageServiceHandler is an implementation of the metalstack.admin.v2.ImageService service.
type ImageServiceHandler interface {
	// Create a image
	Create(context.Context, *connect.Request[v2.ImageServiceCreateRequest]) (*connect.Response[v2.ImageServiceCreateResponse], error)
	// Update a image
	Update(context.Context, *connect.Request[v2.ImageServiceUpdateRequest]) (*connect.Response[v2.ImageServiceUpdateResponse], error)
	// Delete a image
	Delete(context.Context, *connect.Request[v2.ImageServiceDeleteRequest]) (*connect.Response[v2.ImageServiceDeleteResponse], error)
	// Usage of images
	Usage(context.Context, *connect.Request[v2.ImageServiceUsageRequest]) (*connect.Response[v2.ImageServiceUsageResponse], error)
}

// NewImageServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewImageServiceHandler(svc ImageServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	imageServiceMethods := v2.File_metalstack_admin_v2_image_proto.Services().ByName("ImageService").Methods()
	imageServiceCreateHandler := connect.NewUnaryHandler(
		ImageServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(imageServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	imageServiceUpdateHandler := connect.NewUnaryHandler(
		ImageServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(imageServiceMethods.ByName("Update")),
		connect.WithHandlerOptions(opts...),
	)
	imageServiceDeleteHandler := connect.NewUnaryHandler(
		ImageServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(imageServiceMethods.ByName("Delete")),
		connect.WithHandlerOptions(opts...),
	)
	imageServiceUsageHandler := connect.NewUnaryHandler(
		ImageServiceUsageProcedure,
		svc.Usage,
		connect.WithSchema(imageServiceMethods.ByName("Usage")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.admin.v2.ImageService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ImageServiceCreateProcedure:
			imageServiceCreateHandler.ServeHTTP(w, r)
		case ImageServiceUpdateProcedure:
			imageServiceUpdateHandler.ServeHTTP(w, r)
		case ImageServiceDeleteProcedure:
			imageServiceDeleteHandler.ServeHTTP(w, r)
		case ImageServiceUsageProcedure:
			imageServiceUsageHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedImageServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedImageServiceHandler struct{}

func (UnimplementedImageServiceHandler) Create(context.Context, *connect.Request[v2.ImageServiceCreateRequest]) (*connect.Response[v2.ImageServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.ImageService.Create is not implemented"))
}

func (UnimplementedImageServiceHandler) Update(context.Context, *connect.Request[v2.ImageServiceUpdateRequest]) (*connect.Response[v2.ImageServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.ImageService.Update is not implemented"))
}

func (UnimplementedImageServiceHandler) Delete(context.Context, *connect.Request[v2.ImageServiceDeleteRequest]) (*connect.Response[v2.ImageServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.ImageService.Delete is not implemented"))
}

func (UnimplementedImageServiceHandler) Usage(context.Context, *connect.Request[v2.ImageServiceUsageRequest]) (*connect.Response[v2.ImageServiceUsageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.admin.v2.ImageService.Usage is not implemented"))
}
