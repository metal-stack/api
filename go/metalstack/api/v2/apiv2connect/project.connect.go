// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metalstack/api/v2/project.proto

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
	// ProjectServiceName is the fully-qualified name of the ProjectService service.
	ProjectServiceName = "metalstack.api.v2.ProjectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProjectServiceListProcedure is the fully-qualified name of the ProjectService's List RPC.
	ProjectServiceListProcedure = "/metalstack.api.v2.ProjectService/List"
	// ProjectServiceGetProcedure is the fully-qualified name of the ProjectService's Get RPC.
	ProjectServiceGetProcedure = "/metalstack.api.v2.ProjectService/Get"
	// ProjectServiceCreateProcedure is the fully-qualified name of the ProjectService's Create RPC.
	ProjectServiceCreateProcedure = "/metalstack.api.v2.ProjectService/Create"
	// ProjectServiceDeleteProcedure is the fully-qualified name of the ProjectService's Delete RPC.
	ProjectServiceDeleteProcedure = "/metalstack.api.v2.ProjectService/Delete"
	// ProjectServiceUpdateProcedure is the fully-qualified name of the ProjectService's Update RPC.
	ProjectServiceUpdateProcedure = "/metalstack.api.v2.ProjectService/Update"
	// ProjectServiceRemoveMemberProcedure is the fully-qualified name of the ProjectService's
	// RemoveMember RPC.
	ProjectServiceRemoveMemberProcedure = "/metalstack.api.v2.ProjectService/RemoveMember"
	// ProjectServiceUpdateMemberProcedure is the fully-qualified name of the ProjectService's
	// UpdateMember RPC.
	ProjectServiceUpdateMemberProcedure = "/metalstack.api.v2.ProjectService/UpdateMember"
	// ProjectServiceInviteProcedure is the fully-qualified name of the ProjectService's Invite RPC.
	ProjectServiceInviteProcedure = "/metalstack.api.v2.ProjectService/Invite"
	// ProjectServiceInviteAcceptProcedure is the fully-qualified name of the ProjectService's
	// InviteAccept RPC.
	ProjectServiceInviteAcceptProcedure = "/metalstack.api.v2.ProjectService/InviteAccept"
	// ProjectServiceInviteDeleteProcedure is the fully-qualified name of the ProjectService's
	// InviteDelete RPC.
	ProjectServiceInviteDeleteProcedure = "/metalstack.api.v2.ProjectService/InviteDelete"
	// ProjectServiceInvitesListProcedure is the fully-qualified name of the ProjectService's
	// InvitesList RPC.
	ProjectServiceInvitesListProcedure = "/metalstack.api.v2.ProjectService/InvitesList"
	// ProjectServiceInviteGetProcedure is the fully-qualified name of the ProjectService's InviteGet
	// RPC.
	ProjectServiceInviteGetProcedure = "/metalstack.api.v2.ProjectService/InviteGet"
)

// ProjectServiceClient is a client for the metalstack.api.v2.ProjectService service.
type ProjectServiceClient interface {
	// List all accessible projects
	List(context.Context, *connect.Request[v2.ProjectServiceListRequest]) (*connect.Response[v2.ProjectServiceListResponse], error)
	// Get a project
	Get(context.Context, *connect.Request[v2.ProjectServiceGetRequest]) (*connect.Response[v2.ProjectServiceGetResponse], error)
	// Create a project
	Create(context.Context, *connect.Request[v2.ProjectServiceCreateRequest]) (*connect.Response[v2.ProjectServiceCreateResponse], error)
	// Delete a project
	Delete(context.Context, *connect.Request[v2.ProjectServiceDeleteRequest]) (*connect.Response[v2.ProjectServiceDeleteResponse], error)
	// Update a project
	Update(context.Context, *connect.Request[v2.ProjectServiceUpdateRequest]) (*connect.Response[v2.ProjectServiceUpdateResponse], error)
	// RemoveMember remove a user from a project
	RemoveMember(context.Context, *connect.Request[v2.ProjectServiceRemoveMemberRequest]) (*connect.Response[v2.ProjectServiceRemoveMemberResponse], error)
	// UpdateMember update a user for a project
	UpdateMember(context.Context, *connect.Request[v2.ProjectServiceUpdateMemberRequest]) (*connect.Response[v2.ProjectServiceUpdateMemberResponse], error)
	// Invite a user to a project
	Invite(context.Context, *connect.Request[v2.ProjectServiceInviteRequest]) (*connect.Response[v2.ProjectServiceInviteResponse], error)
	// InviteAccept is called from a user to accept a invitation
	InviteAccept(context.Context, *connect.Request[v2.ProjectServiceInviteAcceptRequest]) (*connect.Response[v2.ProjectServiceInviteAcceptResponse], error)
	// InviteDelete deletes a pending invitation
	InviteDelete(context.Context, *connect.Request[v2.ProjectServiceInviteDeleteRequest]) (*connect.Response[v2.ProjectServiceInviteDeleteResponse], error)
	// InvitesList list all invites to a project
	InvitesList(context.Context, *connect.Request[v2.ProjectServiceInvitesListRequest]) (*connect.Response[v2.ProjectServiceInvitesListResponse], error)
	// InviteGet get an invite
	InviteGet(context.Context, *connect.Request[v2.ProjectServiceInviteGetRequest]) (*connect.Response[v2.ProjectServiceInviteGetResponse], error)
}

// NewProjectServiceClient constructs a client for the metalstack.api.v2.ProjectService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	projectServiceMethods := v2.File_metalstack_api_v2_project_proto.Services().ByName("ProjectService").Methods()
	return &projectServiceClient{
		list: connect.NewClient[v2.ProjectServiceListRequest, v2.ProjectServiceListResponse](
			httpClient,
			baseURL+ProjectServiceListProcedure,
			connect.WithSchema(projectServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v2.ProjectServiceGetRequest, v2.ProjectServiceGetResponse](
			httpClient,
			baseURL+ProjectServiceGetProcedure,
			connect.WithSchema(projectServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
		create: connect.NewClient[v2.ProjectServiceCreateRequest, v2.ProjectServiceCreateResponse](
			httpClient,
			baseURL+ProjectServiceCreateProcedure,
			connect.WithSchema(projectServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v2.ProjectServiceDeleteRequest, v2.ProjectServiceDeleteResponse](
			httpClient,
			baseURL+ProjectServiceDeleteProcedure,
			connect.WithSchema(projectServiceMethods.ByName("Delete")),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v2.ProjectServiceUpdateRequest, v2.ProjectServiceUpdateResponse](
			httpClient,
			baseURL+ProjectServiceUpdateProcedure,
			connect.WithSchema(projectServiceMethods.ByName("Update")),
			connect.WithClientOptions(opts...),
		),
		removeMember: connect.NewClient[v2.ProjectServiceRemoveMemberRequest, v2.ProjectServiceRemoveMemberResponse](
			httpClient,
			baseURL+ProjectServiceRemoveMemberProcedure,
			connect.WithSchema(projectServiceMethods.ByName("RemoveMember")),
			connect.WithClientOptions(opts...),
		),
		updateMember: connect.NewClient[v2.ProjectServiceUpdateMemberRequest, v2.ProjectServiceUpdateMemberResponse](
			httpClient,
			baseURL+ProjectServiceUpdateMemberProcedure,
			connect.WithSchema(projectServiceMethods.ByName("UpdateMember")),
			connect.WithClientOptions(opts...),
		),
		invite: connect.NewClient[v2.ProjectServiceInviteRequest, v2.ProjectServiceInviteResponse](
			httpClient,
			baseURL+ProjectServiceInviteProcedure,
			connect.WithSchema(projectServiceMethods.ByName("Invite")),
			connect.WithClientOptions(opts...),
		),
		inviteAccept: connect.NewClient[v2.ProjectServiceInviteAcceptRequest, v2.ProjectServiceInviteAcceptResponse](
			httpClient,
			baseURL+ProjectServiceInviteAcceptProcedure,
			connect.WithSchema(projectServiceMethods.ByName("InviteAccept")),
			connect.WithClientOptions(opts...),
		),
		inviteDelete: connect.NewClient[v2.ProjectServiceInviteDeleteRequest, v2.ProjectServiceInviteDeleteResponse](
			httpClient,
			baseURL+ProjectServiceInviteDeleteProcedure,
			connect.WithSchema(projectServiceMethods.ByName("InviteDelete")),
			connect.WithClientOptions(opts...),
		),
		invitesList: connect.NewClient[v2.ProjectServiceInvitesListRequest, v2.ProjectServiceInvitesListResponse](
			httpClient,
			baseURL+ProjectServiceInvitesListProcedure,
			connect.WithSchema(projectServiceMethods.ByName("InvitesList")),
			connect.WithClientOptions(opts...),
		),
		inviteGet: connect.NewClient[v2.ProjectServiceInviteGetRequest, v2.ProjectServiceInviteGetResponse](
			httpClient,
			baseURL+ProjectServiceInviteGetProcedure,
			connect.WithSchema(projectServiceMethods.ByName("InviteGet")),
			connect.WithClientOptions(opts...),
		),
	}
}

// projectServiceClient implements ProjectServiceClient.
type projectServiceClient struct {
	list         *connect.Client[v2.ProjectServiceListRequest, v2.ProjectServiceListResponse]
	get          *connect.Client[v2.ProjectServiceGetRequest, v2.ProjectServiceGetResponse]
	create       *connect.Client[v2.ProjectServiceCreateRequest, v2.ProjectServiceCreateResponse]
	delete       *connect.Client[v2.ProjectServiceDeleteRequest, v2.ProjectServiceDeleteResponse]
	update       *connect.Client[v2.ProjectServiceUpdateRequest, v2.ProjectServiceUpdateResponse]
	removeMember *connect.Client[v2.ProjectServiceRemoveMemberRequest, v2.ProjectServiceRemoveMemberResponse]
	updateMember *connect.Client[v2.ProjectServiceUpdateMemberRequest, v2.ProjectServiceUpdateMemberResponse]
	invite       *connect.Client[v2.ProjectServiceInviteRequest, v2.ProjectServiceInviteResponse]
	inviteAccept *connect.Client[v2.ProjectServiceInviteAcceptRequest, v2.ProjectServiceInviteAcceptResponse]
	inviteDelete *connect.Client[v2.ProjectServiceInviteDeleteRequest, v2.ProjectServiceInviteDeleteResponse]
	invitesList  *connect.Client[v2.ProjectServiceInvitesListRequest, v2.ProjectServiceInvitesListResponse]
	inviteGet    *connect.Client[v2.ProjectServiceInviteGetRequest, v2.ProjectServiceInviteGetResponse]
}

// List calls metalstack.api.v2.ProjectService.List.
func (c *projectServiceClient) List(ctx context.Context, req *connect.Request[v2.ProjectServiceListRequest]) (*connect.Response[v2.ProjectServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Get calls metalstack.api.v2.ProjectService.Get.
func (c *projectServiceClient) Get(ctx context.Context, req *connect.Request[v2.ProjectServiceGetRequest]) (*connect.Response[v2.ProjectServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Create calls metalstack.api.v2.ProjectService.Create.
func (c *projectServiceClient) Create(ctx context.Context, req *connect.Request[v2.ProjectServiceCreateRequest]) (*connect.Response[v2.ProjectServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Delete calls metalstack.api.v2.ProjectService.Delete.
func (c *projectServiceClient) Delete(ctx context.Context, req *connect.Request[v2.ProjectServiceDeleteRequest]) (*connect.Response[v2.ProjectServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Update calls metalstack.api.v2.ProjectService.Update.
func (c *projectServiceClient) Update(ctx context.Context, req *connect.Request[v2.ProjectServiceUpdateRequest]) (*connect.Response[v2.ProjectServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// RemoveMember calls metalstack.api.v2.ProjectService.RemoveMember.
func (c *projectServiceClient) RemoveMember(ctx context.Context, req *connect.Request[v2.ProjectServiceRemoveMemberRequest]) (*connect.Response[v2.ProjectServiceRemoveMemberResponse], error) {
	return c.removeMember.CallUnary(ctx, req)
}

// UpdateMember calls metalstack.api.v2.ProjectService.UpdateMember.
func (c *projectServiceClient) UpdateMember(ctx context.Context, req *connect.Request[v2.ProjectServiceUpdateMemberRequest]) (*connect.Response[v2.ProjectServiceUpdateMemberResponse], error) {
	return c.updateMember.CallUnary(ctx, req)
}

// Invite calls metalstack.api.v2.ProjectService.Invite.
func (c *projectServiceClient) Invite(ctx context.Context, req *connect.Request[v2.ProjectServiceInviteRequest]) (*connect.Response[v2.ProjectServiceInviteResponse], error) {
	return c.invite.CallUnary(ctx, req)
}

// InviteAccept calls metalstack.api.v2.ProjectService.InviteAccept.
func (c *projectServiceClient) InviteAccept(ctx context.Context, req *connect.Request[v2.ProjectServiceInviteAcceptRequest]) (*connect.Response[v2.ProjectServiceInviteAcceptResponse], error) {
	return c.inviteAccept.CallUnary(ctx, req)
}

// InviteDelete calls metalstack.api.v2.ProjectService.InviteDelete.
func (c *projectServiceClient) InviteDelete(ctx context.Context, req *connect.Request[v2.ProjectServiceInviteDeleteRequest]) (*connect.Response[v2.ProjectServiceInviteDeleteResponse], error) {
	return c.inviteDelete.CallUnary(ctx, req)
}

// InvitesList calls metalstack.api.v2.ProjectService.InvitesList.
func (c *projectServiceClient) InvitesList(ctx context.Context, req *connect.Request[v2.ProjectServiceInvitesListRequest]) (*connect.Response[v2.ProjectServiceInvitesListResponse], error) {
	return c.invitesList.CallUnary(ctx, req)
}

// InviteGet calls metalstack.api.v2.ProjectService.InviteGet.
func (c *projectServiceClient) InviteGet(ctx context.Context, req *connect.Request[v2.ProjectServiceInviteGetRequest]) (*connect.Response[v2.ProjectServiceInviteGetResponse], error) {
	return c.inviteGet.CallUnary(ctx, req)
}

// ProjectServiceHandler is an implementation of the metalstack.api.v2.ProjectService service.
type ProjectServiceHandler interface {
	// List all accessible projects
	List(context.Context, *connect.Request[v2.ProjectServiceListRequest]) (*connect.Response[v2.ProjectServiceListResponse], error)
	// Get a project
	Get(context.Context, *connect.Request[v2.ProjectServiceGetRequest]) (*connect.Response[v2.ProjectServiceGetResponse], error)
	// Create a project
	Create(context.Context, *connect.Request[v2.ProjectServiceCreateRequest]) (*connect.Response[v2.ProjectServiceCreateResponse], error)
	// Delete a project
	Delete(context.Context, *connect.Request[v2.ProjectServiceDeleteRequest]) (*connect.Response[v2.ProjectServiceDeleteResponse], error)
	// Update a project
	Update(context.Context, *connect.Request[v2.ProjectServiceUpdateRequest]) (*connect.Response[v2.ProjectServiceUpdateResponse], error)
	// RemoveMember remove a user from a project
	RemoveMember(context.Context, *connect.Request[v2.ProjectServiceRemoveMemberRequest]) (*connect.Response[v2.ProjectServiceRemoveMemberResponse], error)
	// UpdateMember update a user for a project
	UpdateMember(context.Context, *connect.Request[v2.ProjectServiceUpdateMemberRequest]) (*connect.Response[v2.ProjectServiceUpdateMemberResponse], error)
	// Invite a user to a project
	Invite(context.Context, *connect.Request[v2.ProjectServiceInviteRequest]) (*connect.Response[v2.ProjectServiceInviteResponse], error)
	// InviteAccept is called from a user to accept a invitation
	InviteAccept(context.Context, *connect.Request[v2.ProjectServiceInviteAcceptRequest]) (*connect.Response[v2.ProjectServiceInviteAcceptResponse], error)
	// InviteDelete deletes a pending invitation
	InviteDelete(context.Context, *connect.Request[v2.ProjectServiceInviteDeleteRequest]) (*connect.Response[v2.ProjectServiceInviteDeleteResponse], error)
	// InvitesList list all invites to a project
	InvitesList(context.Context, *connect.Request[v2.ProjectServiceInvitesListRequest]) (*connect.Response[v2.ProjectServiceInvitesListResponse], error)
	// InviteGet get an invite
	InviteGet(context.Context, *connect.Request[v2.ProjectServiceInviteGetRequest]) (*connect.Response[v2.ProjectServiceInviteGetResponse], error)
}

// NewProjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectServiceHandler(svc ProjectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	projectServiceMethods := v2.File_metalstack_api_v2_project_proto.Services().ByName("ProjectService").Methods()
	projectServiceListHandler := connect.NewUnaryHandler(
		ProjectServiceListProcedure,
		svc.List,
		connect.WithSchema(projectServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceGetHandler := connect.NewUnaryHandler(
		ProjectServiceGetProcedure,
		svc.Get,
		connect.WithSchema(projectServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceCreateHandler := connect.NewUnaryHandler(
		ProjectServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(projectServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceDeleteHandler := connect.NewUnaryHandler(
		ProjectServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(projectServiceMethods.ByName("Delete")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceUpdateHandler := connect.NewUnaryHandler(
		ProjectServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(projectServiceMethods.ByName("Update")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceRemoveMemberHandler := connect.NewUnaryHandler(
		ProjectServiceRemoveMemberProcedure,
		svc.RemoveMember,
		connect.WithSchema(projectServiceMethods.ByName("RemoveMember")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceUpdateMemberHandler := connect.NewUnaryHandler(
		ProjectServiceUpdateMemberProcedure,
		svc.UpdateMember,
		connect.WithSchema(projectServiceMethods.ByName("UpdateMember")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteHandler := connect.NewUnaryHandler(
		ProjectServiceInviteProcedure,
		svc.Invite,
		connect.WithSchema(projectServiceMethods.ByName("Invite")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteAcceptHandler := connect.NewUnaryHandler(
		ProjectServiceInviteAcceptProcedure,
		svc.InviteAccept,
		connect.WithSchema(projectServiceMethods.ByName("InviteAccept")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteDeleteHandler := connect.NewUnaryHandler(
		ProjectServiceInviteDeleteProcedure,
		svc.InviteDelete,
		connect.WithSchema(projectServiceMethods.ByName("InviteDelete")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInvitesListHandler := connect.NewUnaryHandler(
		ProjectServiceInvitesListProcedure,
		svc.InvitesList,
		connect.WithSchema(projectServiceMethods.ByName("InvitesList")),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteGetHandler := connect.NewUnaryHandler(
		ProjectServiceInviteGetProcedure,
		svc.InviteGet,
		connect.WithSchema(projectServiceMethods.ByName("InviteGet")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metalstack.api.v2.ProjectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProjectServiceListProcedure:
			projectServiceListHandler.ServeHTTP(w, r)
		case ProjectServiceGetProcedure:
			projectServiceGetHandler.ServeHTTP(w, r)
		case ProjectServiceCreateProcedure:
			projectServiceCreateHandler.ServeHTTP(w, r)
		case ProjectServiceDeleteProcedure:
			projectServiceDeleteHandler.ServeHTTP(w, r)
		case ProjectServiceUpdateProcedure:
			projectServiceUpdateHandler.ServeHTTP(w, r)
		case ProjectServiceRemoveMemberProcedure:
			projectServiceRemoveMemberHandler.ServeHTTP(w, r)
		case ProjectServiceUpdateMemberProcedure:
			projectServiceUpdateMemberHandler.ServeHTTP(w, r)
		case ProjectServiceInviteProcedure:
			projectServiceInviteHandler.ServeHTTP(w, r)
		case ProjectServiceInviteAcceptProcedure:
			projectServiceInviteAcceptHandler.ServeHTTP(w, r)
		case ProjectServiceInviteDeleteProcedure:
			projectServiceInviteDeleteHandler.ServeHTTP(w, r)
		case ProjectServiceInvitesListProcedure:
			projectServiceInvitesListHandler.ServeHTTP(w, r)
		case ProjectServiceInviteGetProcedure:
			projectServiceInviteGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectServiceHandler struct{}

func (UnimplementedProjectServiceHandler) List(context.Context, *connect.Request[v2.ProjectServiceListRequest]) (*connect.Response[v2.ProjectServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.List is not implemented"))
}

func (UnimplementedProjectServiceHandler) Get(context.Context, *connect.Request[v2.ProjectServiceGetRequest]) (*connect.Response[v2.ProjectServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.Get is not implemented"))
}

func (UnimplementedProjectServiceHandler) Create(context.Context, *connect.Request[v2.ProjectServiceCreateRequest]) (*connect.Response[v2.ProjectServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.Create is not implemented"))
}

func (UnimplementedProjectServiceHandler) Delete(context.Context, *connect.Request[v2.ProjectServiceDeleteRequest]) (*connect.Response[v2.ProjectServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.Delete is not implemented"))
}

func (UnimplementedProjectServiceHandler) Update(context.Context, *connect.Request[v2.ProjectServiceUpdateRequest]) (*connect.Response[v2.ProjectServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.Update is not implemented"))
}

func (UnimplementedProjectServiceHandler) RemoveMember(context.Context, *connect.Request[v2.ProjectServiceRemoveMemberRequest]) (*connect.Response[v2.ProjectServiceRemoveMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.RemoveMember is not implemented"))
}

func (UnimplementedProjectServiceHandler) UpdateMember(context.Context, *connect.Request[v2.ProjectServiceUpdateMemberRequest]) (*connect.Response[v2.ProjectServiceUpdateMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.UpdateMember is not implemented"))
}

func (UnimplementedProjectServiceHandler) Invite(context.Context, *connect.Request[v2.ProjectServiceInviteRequest]) (*connect.Response[v2.ProjectServiceInviteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.Invite is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteAccept(context.Context, *connect.Request[v2.ProjectServiceInviteAcceptRequest]) (*connect.Response[v2.ProjectServiceInviteAcceptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.InviteAccept is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteDelete(context.Context, *connect.Request[v2.ProjectServiceInviteDeleteRequest]) (*connect.Response[v2.ProjectServiceInviteDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.InviteDelete is not implemented"))
}

func (UnimplementedProjectServiceHandler) InvitesList(context.Context, *connect.Request[v2.ProjectServiceInvitesListRequest]) (*connect.Response[v2.ProjectServiceInvitesListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.InvitesList is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteGet(context.Context, *connect.Request[v2.ProjectServiceInviteGetRequest]) (*connect.Response[v2.ProjectServiceInviteGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metalstack.api.v2.ProjectService.InviteGet is not implemented"))
}
