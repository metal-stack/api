import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Labels, Meta, TenantRole, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/tenant.proto.
 */
export declare const file_metalstack_api_v2_tenant: GenFile;
/**
 * Tenant is a customer of the platform
 *
 * @generated from message metalstack.api.v2.Tenant
 */
export type Tenant = Message<"metalstack.api.v2.Tenant"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Meta for this tenant
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Name of the tenant
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Email of the tenant
     *
     * @generated from field: string email = 4;
     */
    email: string;
    /**
     * Description of this tenant
     *
     * @generated from field: string description = 5;
     */
    description: string;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: string avatar_url = 6;
     */
    avatarUrl: string;
    /**
     * CreatedBy stores who created this tenant
     *
     * @generated from field: string created_by = 15;
     */
    createdBy: string;
};
/**
 * Describes the message metalstack.api.v2.Tenant.
 * Use `create(TenantSchema)` to create a new message.
 */
export declare const TenantSchema: GenMessage<Tenant>;
/**
 * TenantMember defines a user that participates at a tenant
 *
 * @generated from message metalstack.api.v2.TenantMember
 */
export type TenantMember = Message<"metalstack.api.v2.TenantMember"> & {
    /**
     * Id is the user id of the member
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Role is the role of the member
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 2;
     */
    role: TenantRole;
    /**
     * Projects for the projects in which a user is a direct member
     *
     * @generated from field: repeated string projects = 4;
     */
    projects: string[];
    /**
     * CreatedAt the date when the member was added to the tenant
     *
     * @generated from field: google.protobuf.Timestamp created_at = 10;
     */
    createdAt?: Timestamp;
};
/**
 * Describes the message metalstack.api.v2.TenantMember.
 * Use `create(TenantMemberSchema)` to create a new message.
 */
export declare const TenantMemberSchema: GenMessage<TenantMember>;
/**
 * TenantInvite defines invite to tenant
 *
 * @generated from message metalstack.api.v2.TenantInvite
 */
export type TenantInvite = Message<"metalstack.api.v2.TenantInvite"> & {
    /**
     * Secret is the secret part of the invite, typically part of the url
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
    /**
     * TargetTenant is the tenant id for which this invite was created
     *
     * @generated from field: string target_tenant = 2;
     */
    targetTenant: string;
    /**
     * Role is the role in this tenant the user will get after accepting the invitation
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 3;
     */
    role: TenantRole;
    /**
     * Joined is false as long as a user has not accepted the invite
     *
     * @generated from field: bool joined = 4;
     */
    joined: boolean;
    /**
     * TargetTenantName is the tenant name for which this invite was created
     *
     * @generated from field: string target_tenant_name = 5;
     */
    targetTenantName: string;
    /**
     * Tenant is the login of tenant who invites to join this tenant
     *
     * @generated from field: string tenant = 6;
     */
    tenant: string;
    /**
     * TenantName is the name of tenant who invites to join this tenant
     *
     * @generated from field: string tenant_name = 7;
     */
    tenantName: string;
    /**
     * ExpiresAt the date when this invite expires
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 10;
     */
    expiresAt?: Timestamp;
    /**
     * JoinedAt the date when the member accepted this invite
     *
     * @generated from field: google.protobuf.Timestamp joined_at = 11;
     */
    joinedAt?: Timestamp;
};
/**
 * Describes the message metalstack.api.v2.TenantInvite.
 * Use `create(TenantInviteSchema)` to create a new message.
 */
export declare const TenantInviteSchema: GenMessage<TenantInvite>;
/**
 * TenantServiceListRequest is the request payload of the tenant list request
 *
 * @generated from message metalstack.api.v2.TenantServiceListRequest
 */
export type TenantServiceListRequest = Message<"metalstack.api.v2.TenantServiceListRequest"> & {
    /**
     * Id filters tenants by id
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Name filters tenants by name
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Labels lists only projects containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 3;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceListRequest.
 * Use `create(TenantServiceListRequestSchema)` to create a new message.
 */
export declare const TenantServiceListRequestSchema: GenMessage<TenantServiceListRequest>;
/**
 * TenantServiceGetRequest is the request payload of the tenant get request
 *
 * @generated from message metalstack.api.v2.TenantServiceGetRequest
 */
export type TenantServiceGetRequest = Message<"metalstack.api.v2.TenantServiceGetRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceGetRequest.
 * Use `create(TenantServiceGetRequestSchema)` to create a new message.
 */
export declare const TenantServiceGetRequestSchema: GenMessage<TenantServiceGetRequest>;
/**
 * TenantServiceCreateRequest is the request payload of the tenant create request
 *
 * @generated from message metalstack.api.v2.TenantServiceCreateRequest
 */
export type TenantServiceCreateRequest = Message<"metalstack.api.v2.TenantServiceCreateRequest"> & {
    /**
     * Name of this tenant
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Description of this tenant
     *
     * @generated from field: optional string description = 2;
     */
    description?: string;
    /**
     * Email of the tenant, if not set will be inherited from the creator
     *
     * @generated from field: optional string email = 3;
     */
    email?: string;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string;
    /**
     * Labels on the tenant
     *
     * @generated from field: metalstack.api.v2.Labels labels = 5;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceCreateRequest.
 * Use `create(TenantServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantServiceCreateRequestSchema: GenMessage<TenantServiceCreateRequest>;
/**
 * TenantServiceUpdateRequest is the request payload of the tenant update request
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateRequest
 */
export type TenantServiceUpdateRequest = Message<"metalstack.api.v2.TenantServiceUpdateRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Name of the tenant
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Email of the tenant
     *
     * @generated from field: optional string email = 4;
     */
    email?: string;
    /**
     * Description of this tenant
     *
     * @generated from field: optional string description = 5;
     */
    description?: string;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: optional string avatar_url = 6;
     */
    avatarUrl?: string;
    /**
     * Labels on the tenant
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 7;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateRequest.
 * Use `create(TenantServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TenantServiceUpdateRequestSchema: GenMessage<TenantServiceUpdateRequest>;
/**
 * TenantServiceDeleteRequest is the request payload of the tenant delete request
 *
 * @generated from message metalstack.api.v2.TenantServiceDeleteRequest
 */
export type TenantServiceDeleteRequest = Message<"metalstack.api.v2.TenantServiceDeleteRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceDeleteRequest.
 * Use `create(TenantServiceDeleteRequestSchema)` to create a new message.
 */
export declare const TenantServiceDeleteRequestSchema: GenMessage<TenantServiceDeleteRequest>;
/**
 * TenantServiceGetResponse is the response payload of the tenant get request
 *
 * @generated from message metalstack.api.v2.TenantServiceGetResponse
 */
export type TenantServiceGetResponse = Message<"metalstack.api.v2.TenantServiceGetResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant;
    /**
     * TenantMembers of this tenant
     *
     * @generated from field: repeated metalstack.api.v2.TenantMember tenant_members = 2;
     */
    tenantMembers: TenantMember[];
};
/**
 * Describes the message metalstack.api.v2.TenantServiceGetResponse.
 * Use `create(TenantServiceGetResponseSchema)` to create a new message.
 */
export declare const TenantServiceGetResponseSchema: GenMessage<TenantServiceGetResponse>;
/**
 * TenantServiceListResponse is the response payload of the tenant list request
 *
 * @generated from message metalstack.api.v2.TenantServiceListResponse
 */
export type TenantServiceListResponse = Message<"metalstack.api.v2.TenantServiceListResponse"> & {
    /**
     * Tenants is the list of tenants
     *
     * @generated from field: repeated metalstack.api.v2.Tenant tenants = 1;
     */
    tenants: Tenant[];
};
/**
 * Describes the message metalstack.api.v2.TenantServiceListResponse.
 * Use `create(TenantServiceListResponseSchema)` to create a new message.
 */
export declare const TenantServiceListResponseSchema: GenMessage<TenantServiceListResponse>;
/**
 * TenantServiceCreateResponse is the response payload of the tenant create request
 *
 * @generated from message metalstack.api.v2.TenantServiceCreateResponse
 */
export type TenantServiceCreateResponse = Message<"metalstack.api.v2.TenantServiceCreateResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceCreateResponse.
 * Use `create(TenantServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantServiceCreateResponseSchema: GenMessage<TenantServiceCreateResponse>;
/**
 * TenantServiceUpdateResponse is the response payload of the tenant update request
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateResponse
 */
export type TenantServiceUpdateResponse = Message<"metalstack.api.v2.TenantServiceUpdateResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateResponse.
 * Use `create(TenantServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TenantServiceUpdateResponseSchema: GenMessage<TenantServiceUpdateResponse>;
/**
 * TenantServiceDeleteResponse is the response payload of the tenant delete request
 *
 * @generated from message metalstack.api.v2.TenantServiceDeleteResponse
 */
export type TenantServiceDeleteResponse = Message<"metalstack.api.v2.TenantServiceDeleteResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceDeleteResponse.
 * Use `create(TenantServiceDeleteResponseSchema)` to create a new message.
 */
export declare const TenantServiceDeleteResponseSchema: GenMessage<TenantServiceDeleteResponse>;
/**
 * TenantServiceInviteRequest is used to invite a member to a tenant
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteRequest
 */
export type TenantServiceInviteRequest = Message<"metalstack.api.v2.TenantServiceInviteRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Role of this user in this tenant
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 2;
     */
    role: TenantRole;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteRequest.
 * Use `create(TenantServiceInviteRequestSchema)` to create a new message.
 */
export declare const TenantServiceInviteRequestSchema: GenMessage<TenantServiceInviteRequest>;
/**
 * TenantServiceInviteRequest is the response payload to a invite member request
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteResponse
 */
export type TenantServiceInviteResponse = Message<"metalstack.api.v2.TenantServiceInviteResponse"> & {
    /**
     * Invite contains a secret which can be sent to a potential user
     * can be appended to the invitation endpoint at our cloud console like
     * console.metalstack.cloud/invite/<secret>
     *
     * @generated from field: metalstack.api.v2.TenantInvite invite = 1;
     */
    invite?: TenantInvite;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteResponse.
 * Use `create(TenantServiceInviteResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteResponseSchema: GenMessage<TenantServiceInviteResponse>;
/**
 * TenantServiceInvitesListRequest is the request payload to a list invites request
 *
 * @generated from message metalstack.api.v2.TenantServiceInvitesListRequest
 */
export type TenantServiceInvitesListRequest = Message<"metalstack.api.v2.TenantServiceInvitesListRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInvitesListRequest.
 * Use `create(TenantServiceInvitesListRequestSchema)` to create a new message.
 */
export declare const TenantServiceInvitesListRequestSchema: GenMessage<TenantServiceInvitesListRequest>;
/**
 * TenantServiceInvitesListResponse is the response payload to a list invites request
 *
 * @generated from message metalstack.api.v2.TenantServiceInvitesListResponse
 */
export type TenantServiceInvitesListResponse = Message<"metalstack.api.v2.TenantServiceInvitesListResponse"> & {
    /**
     * Invites not already accepted the invitation to this tenant
     *
     * @generated from field: repeated metalstack.api.v2.TenantInvite invites = 1;
     */
    invites: TenantInvite[];
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInvitesListResponse.
 * Use `create(TenantServiceInvitesListResponseSchema)` to create a new message.
 */
export declare const TenantServiceInvitesListResponseSchema: GenMessage<TenantServiceInvitesListResponse>;
/**
 * TenantServiceInviteGetRequest is the request payload to get a invite
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteGetRequest
 */
export type TenantServiceInviteGetRequest = Message<"metalstack.api.v2.TenantServiceInviteGetRequest"> & {
    /**
     * Secret of the invite to get
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteGetRequest.
 * Use `create(TenantServiceInviteGetRequestSchema)` to create a new message.
 */
export declare const TenantServiceInviteGetRequestSchema: GenMessage<TenantServiceInviteGetRequest>;
/**
 * TenantServiceInviteGetResponse is the response payload to a get invite request
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteGetResponse
 */
export type TenantServiceInviteGetResponse = Message<"metalstack.api.v2.TenantServiceInviteGetResponse"> & {
    /**
     * Invite is the invite
     *
     * @generated from field: metalstack.api.v2.TenantInvite invite = 1;
     */
    invite?: TenantInvite;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteGetResponse.
 * Use `create(TenantServiceInviteGetResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteGetResponseSchema: GenMessage<TenantServiceInviteGetResponse>;
/**
 * TenantServiceRemoveMemberRequest is used to remove a member from a tenant
 *
 * @generated from message metalstack.api.v2.TenantServiceRemoveMemberRequest
 */
export type TenantServiceRemoveMemberRequest = Message<"metalstack.api.v2.TenantServiceRemoveMemberRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Member is the id of the member to remove from this tenant
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceRemoveMemberRequest.
 * Use `create(TenantServiceRemoveMemberRequestSchema)` to create a new message.
 */
export declare const TenantServiceRemoveMemberRequestSchema: GenMessage<TenantServiceRemoveMemberRequest>;
/**
 * TenantServiceLeaveTenantRequest is used to leave a tenant
 *
 * @generated from message metalstack.api.v2.TenantServiceLeaveRequest
 */
export type TenantServiceLeaveRequest = Message<"metalstack.api.v2.TenantServiceLeaveRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceLeaveRequest.
 * Use `create(TenantServiceLeaveRequestSchema)` to create a new message.
 */
export declare const TenantServiceLeaveRequestSchema: GenMessage<TenantServiceLeaveRequest>;
/**
 * TenantServiceLeaveTenantResponse is the response payload to a leave tenant request
 *
 * @generated from message metalstack.api.v2.TenantServiceLeaveResponse
 */
export type TenantServiceLeaveResponse = Message<"metalstack.api.v2.TenantServiceLeaveResponse"> & {};
/**
 * Describes the message metalstack.api.v2.TenantServiceLeaveResponse.
 * Use `create(TenantServiceLeaveResponseSchema)` to create a new message.
 */
export declare const TenantServiceLeaveResponseSchema: GenMessage<TenantServiceLeaveResponse>;
/**
 * TenantServiceRemoveMemberResponse is the response payload to a remove member request
 *
 * @generated from message metalstack.api.v2.TenantServiceRemoveMemberResponse
 */
export type TenantServiceRemoveMemberResponse = Message<"metalstack.api.v2.TenantServiceRemoveMemberResponse"> & {};
/**
 * Describes the message metalstack.api.v2.TenantServiceRemoveMemberResponse.
 * Use `create(TenantServiceRemoveMemberResponseSchema)` to create a new message.
 */
export declare const TenantServiceRemoveMemberResponseSchema: GenMessage<TenantServiceRemoveMemberResponse>;
/**
 * TenantServiceInviteAcceptRequest is the request payload to a accept invite request
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteAcceptRequest
 */
export type TenantServiceInviteAcceptRequest = Message<"metalstack.api.v2.TenantServiceInviteAcceptRequest"> & {
    /**
     * Secret is the invitation secret part of the invitation url
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteAcceptRequest.
 * Use `create(TenantServiceInviteAcceptRequestSchema)` to create a new message.
 */
export declare const TenantServiceInviteAcceptRequestSchema: GenMessage<TenantServiceInviteAcceptRequest>;
/**
 * TenantServiceInvitesListResponse is the response payload to a accept invite request
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteAcceptResponse
 */
export type TenantServiceInviteAcceptResponse = Message<"metalstack.api.v2.TenantServiceInviteAcceptResponse"> & {
    /**
     * Tenant ID of the joined tenant
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * TenantName of the joined tenant
     *
     * @generated from field: string tenant_name = 2;
     */
    tenantName: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteAcceptResponse.
 * Use `create(TenantServiceInviteAcceptResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteAcceptResponseSchema: GenMessage<TenantServiceInviteAcceptResponse>;
/**
 * TenantServiceInviteDeleteRequest is the request payload to a delete invite
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteDeleteRequest
 */
export type TenantServiceInviteDeleteRequest = Message<"metalstack.api.v2.TenantServiceInviteDeleteRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Secret of the invite to delete
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteDeleteRequest.
 * Use `create(TenantServiceInviteDeleteRequestSchema)` to create a new message.
 */
export declare const TenantServiceInviteDeleteRequestSchema: GenMessage<TenantServiceInviteDeleteRequest>;
/**
 * TenantServiceInviteDeleteResponse is the response payload of a delete invite request
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteDeleteResponse
 */
export type TenantServiceInviteDeleteResponse = Message<"metalstack.api.v2.TenantServiceInviteDeleteResponse"> & {};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteDeleteResponse.
 * Use `create(TenantServiceInviteDeleteResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteDeleteResponseSchema: GenMessage<TenantServiceInviteDeleteResponse>;
/**
 * TenantServiceUpdateMemberRequest is used to update a member from a tenant
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateMemberRequest
 */
export type TenantServiceUpdateMemberRequest = Message<"metalstack.api.v2.TenantServiceUpdateMemberRequest"> & {
    /**
     * Login of the tenant
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Member is the id of the member to update in this tenant
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role of this user in this tenant
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 3;
     */
    role: TenantRole;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateMemberRequest.
 * Use `create(TenantServiceUpdateMemberRequestSchema)` to create a new message.
 */
export declare const TenantServiceUpdateMemberRequestSchema: GenMessage<TenantServiceUpdateMemberRequest>;
/**
 * TenantServiceUpdateMemberResponse is the response payload to a update member request
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateMemberResponse
 */
export type TenantServiceUpdateMemberResponse = Message<"metalstack.api.v2.TenantServiceUpdateMemberResponse"> & {
    /**
     * TenantMember is the updated membership
     *
     * @generated from field: metalstack.api.v2.TenantMember tenant_member = 1;
     */
    tenantMember?: TenantMember;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateMemberResponse.
 * Use `create(TenantServiceUpdateMemberResponseSchema)` to create a new message.
 */
export declare const TenantServiceUpdateMemberResponseSchema: GenMessage<TenantServiceUpdateMemberResponse>;
/**
 * TenantService serves tenant related functions
 *
 * @generated from service metalstack.api.v2.TenantService
 */
export declare const TenantService: GenService<{
    /**
     * Create a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantServiceCreateRequestSchema;
        output: typeof TenantServiceCreateResponseSchema;
    };
    /**
     * List tenants
     *
     * @generated from rpc metalstack.api.v2.TenantService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantServiceListRequestSchema;
        output: typeof TenantServiceListResponseSchema;
    };
    /**
     * Get a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TenantServiceGetRequestSchema;
        output: typeof TenantServiceGetResponseSchema;
    };
    /**
     * Update a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TenantServiceUpdateRequestSchema;
        output: typeof TenantServiceUpdateResponseSchema;
    };
    /**
     * Delete a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof TenantServiceDeleteRequestSchema;
        output: typeof TenantServiceDeleteResponseSchema;
    };
    /**
     * Leave remove a member of a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Leave
     */
    leave: {
        methodKind: "unary";
        input: typeof TenantServiceLeaveRequestSchema;
        output: typeof TenantServiceLeaveResponseSchema;
    };
    /**
     * RemoveMember remove a member of a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.RemoveMember
     */
    removeMember: {
        methodKind: "unary";
        input: typeof TenantServiceRemoveMemberRequestSchema;
        output: typeof TenantServiceRemoveMemberResponseSchema;
    };
    /**
     * UpdateMember update a member of a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.UpdateMember
     */
    updateMember: {
        methodKind: "unary";
        input: typeof TenantServiceUpdateMemberRequestSchema;
        output: typeof TenantServiceUpdateMemberResponseSchema;
    };
    /**
     * Invite a user to a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.Invite
     */
    invite: {
        methodKind: "unary";
        input: typeof TenantServiceInviteRequestSchema;
        output: typeof TenantServiceInviteResponseSchema;
    };
    /**
     * InviteAccept is called from a user to accept an invitation
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteAccept
     */
    inviteAccept: {
        methodKind: "unary";
        input: typeof TenantServiceInviteAcceptRequestSchema;
        output: typeof TenantServiceInviteAcceptResponseSchema;
    };
    /**
     * InviteDelete deletes a pending invitation
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteDelete
     */
    inviteDelete: {
        methodKind: "unary";
        input: typeof TenantServiceInviteDeleteRequestSchema;
        output: typeof TenantServiceInviteDeleteResponseSchema;
    };
    /**
     * InvitesList list all invites to a tenant
     *
     * @generated from rpc metalstack.api.v2.TenantService.InvitesList
     */
    invitesList: {
        methodKind: "unary";
        input: typeof TenantServiceInvitesListRequestSchema;
        output: typeof TenantServiceInvitesListResponseSchema;
    };
    /**
     * InviteGet get an invite
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteGet
     */
    inviteGet: {
        methodKind: "unary";
        input: typeof TenantServiceInviteGetRequestSchema;
        output: typeof TenantServiceInviteGetResponseSchema;
    };
}>;
