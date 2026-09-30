import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Labels, Meta, Paging, TenantRole, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/tenant.proto.
 */
export declare const file_metalstack_api_v2_tenant: GenFile;
/**
 * Tenant is a customer of the platform.
 *
 * @generated from message metalstack.api.v2.Tenant
 */
export type Tenant = Message<"metalstack.api.v2.Tenant"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Meta for this tenant.
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta | undefined;
    /**
     * Name of the tenant.
     *
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * Email of the tenant.
     *
     * @generated from field: string email = 4;
     */
    email: string;
    /**
     * Description of this tenant.
     *
     * @generated from field: string description = 5;
     */
    description: string;
    /**
     * AvatarUrl of the tenant.
     *
     * @generated from field: string avatar_url = 6;
     */
    avatarUrl: string;
    /**
     * CreatedBy stores who created this tenant.
     *
     * @generated from field: string created_by = 7;
     */
    createdBy: string;
};
/**
 * Describes the message metalstack.api.v2.Tenant.
 * Use `create(TenantSchema)` to create a new message.
 */
export declare const TenantSchema: GenMessage<Tenant>;
/**
 * TenantInvite defines invite to tenant.
 *
 * @generated from message metalstack.api.v2.TenantInvite
 */
export type TenantInvite = Message<"metalstack.api.v2.TenantInvite"> & {
    /**
     * Secret is the secret part of the invite, typically part of the url.
     *
     * @generated from field: string secret = 1;
     */
    secret: string;
    /**
     * TargetTenant is the tenant id for which this invite was created.
     *
     * @generated from field: string target_tenant = 2;
     */
    targetTenant: string;
    /**
     * Role is the role in this tenant the user will get after accepting the invitation.
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 3;
     */
    role: TenantRole;
    /**
     * Joined is false as long as a user has not accepted the invite.
     *
     * @generated from field: bool joined = 4;
     */
    joined: boolean;
    /**
     * TargetTenantName is the tenant name for which this invite was created.
     *
     * @generated from field: string target_tenant_name = 5;
     */
    targetTenantName: string;
    /**
     * Tenant is the login of the tenant inviting another user to join this tenant.
     *
     * @generated from field: string tenant = 6;
     */
    tenant: string;
    /**
     * TenantName is the name of the tenant inviting another user to join this tenant.
     *
     * @generated from field: string tenant_name = 7;
     */
    tenantName: string;
    /**
     * ExpiresAt the date when this invite expires.
     *
     * @generated from field: google.protobuf.Timestamp expires_at = 8;
     */
    expiresAt?: Timestamp | undefined;
    /**
     * JoinedAt the date when the member accepted this invite.
     *
     * @generated from field: google.protobuf.Timestamp joined_at = 9;
     */
    joinedAt?: Timestamp | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantInvite.
 * Use `create(TenantInviteSchema)` to create a new message.
 */
export declare const TenantInviteSchema: GenMessage<TenantInvite>;
/**
 * TenantServiceListRequest is the request payload of the tenant list request.
 *
 * @generated from message metalstack.api.v2.TenantServiceListRequest
 */
export type TenantServiceListRequest = Message<"metalstack.api.v2.TenantServiceListRequest"> & {
    /**
     * Query for tenants.
     *
     * @generated from field: metalstack.api.v2.TenantQuery query = 1;
     */
    query?: TenantQuery | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceListRequest.
 * Use `create(TenantServiceListRequestSchema)` to create a new message.
 */
export declare const TenantServiceListRequestSchema: GenMessage<TenantServiceListRequest>;
/**
 * TenantQuery is used to search tenants.
 *
 * @generated from message metalstack.api.v2.TenantQuery
 */
export type TenantQuery = Message<"metalstack.api.v2.TenantQuery"> & {
    /**
     * Id filters tenants by login.
     *
     * @generated from field: optional string login = 1;
     */
    login?: string | undefined;
    /**
     * Name filters tenants by name.
     *
     * @generated from field: optional string name = 2;
     */
    name?: string | undefined;
    /**
     * Labels lists only projects containing the given labels.
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 3;
     */
    labels?: Labels | undefined;
    /**
     * Paging details for the list request.
     *
     * @generated from field: metalstack.api.v2.Paging paging = 4;
     */
    paging?: Paging | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantQuery.
 * Use `create(TenantQuerySchema)` to create a new message.
 */
export declare const TenantQuerySchema: GenMessage<TenantQuery>;
/**
 * TenantServiceGetRequest is the request payload of the tenant get request.
 *
 * @generated from message metalstack.api.v2.TenantServiceGetRequest
 */
export type TenantServiceGetRequest = Message<"metalstack.api.v2.TenantServiceGetRequest"> & {
    /**
     * Login of the tenant.
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
 * TenantServiceCreateRequest is the request payload of the tenant create request.
 *
 * @generated from message metalstack.api.v2.TenantServiceCreateRequest
 */
export type TenantServiceCreateRequest = Message<"metalstack.api.v2.TenantServiceCreateRequest"> & {
    /**
     * Name of this tenant.
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Description of this tenant.
     *
     * @generated from field: optional string description = 2;
     */
    description?: string | undefined;
    /**
     * Email of the tenant, if not set will be inherited from the creator.
     *
     * @generated from field: optional string email = 3;
     */
    email?: string | undefined;
    /**
     * AvatarUrl of the tenant.
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the tenant.
     *
     * @generated from field: metalstack.api.v2.Labels labels = 5;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceCreateRequest.
 * Use `create(TenantServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantServiceCreateRequestSchema: GenMessage<TenantServiceCreateRequest>;
/**
 * TenantServiceUpdateRequest is the request payload of the tenant update request.
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateRequest
 */
export type TenantServiceUpdateRequest = Message<"metalstack.api.v2.TenantServiceUpdateRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * Name of the tenant.
     *
     * @generated from field: optional string name = 3;
     */
    name?: string | undefined;
    /**
     * Email of the tenant.
     *
     * @generated from field: optional string email = 4;
     */
    email?: string | undefined;
    /**
     * Description of this tenant.
     *
     * @generated from field: optional string description = 5;
     */
    description?: string | undefined;
    /**
     * AvatarUrl of the tenant.
     *
     * @generated from field: optional string avatar_url = 6;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the tenant.
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 7;
     */
    labels?: UpdateLabels | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateRequest.
 * Use `create(TenantServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TenantServiceUpdateRequestSchema: GenMessage<TenantServiceUpdateRequest>;
/**
 * TenantServiceDeleteRequest is the request payload of the tenant delete request.
 *
 * @generated from message metalstack.api.v2.TenantServiceDeleteRequest
 */
export type TenantServiceDeleteRequest = Message<"metalstack.api.v2.TenantServiceDeleteRequest"> & {
    /**
     * Login of the tenant.
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
 * TenantServiceGetResponse is the response payload of the tenant get request.
 *
 * @generated from message metalstack.api.v2.TenantServiceGetResponse
 */
export type TenantServiceGetResponse = Message<"metalstack.api.v2.TenantServiceGetResponse"> & {
    /**
     * Tenant is the tenant.
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceGetResponse.
 * Use `create(TenantServiceGetResponseSchema)` to create a new message.
 */
export declare const TenantServiceGetResponseSchema: GenMessage<TenantServiceGetResponse>;
/**
 * TenantServiceListResponse is the response payload of the tenant list request.
 *
 * @generated from message metalstack.api.v2.TenantServiceListResponse
 */
export type TenantServiceListResponse = Message<"metalstack.api.v2.TenantServiceListResponse"> & {
    /**
     * Tenants is the list of tenants.
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
 * TenantServiceCreateResponse is the response payload of the tenant create request.
 *
 * @generated from message metalstack.api.v2.TenantServiceCreateResponse
 */
export type TenantServiceCreateResponse = Message<"metalstack.api.v2.TenantServiceCreateResponse"> & {
    /**
     * Tenant is the tenant.
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceCreateResponse.
 * Use `create(TenantServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantServiceCreateResponseSchema: GenMessage<TenantServiceCreateResponse>;
/**
 * TenantServiceUpdateResponse is the response payload of the tenant update request.
 *
 * @generated from message metalstack.api.v2.TenantServiceUpdateResponse
 */
export type TenantServiceUpdateResponse = Message<"metalstack.api.v2.TenantServiceUpdateResponse"> & {
    /**
     * Tenant is the tenant.
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceUpdateResponse.
 * Use `create(TenantServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TenantServiceUpdateResponseSchema: GenMessage<TenantServiceUpdateResponse>;
/**
 * TenantServiceDeleteResponse is the response payload of the tenant delete request.
 *
 * @generated from message metalstack.api.v2.TenantServiceDeleteResponse
 */
export type TenantServiceDeleteResponse = Message<"metalstack.api.v2.TenantServiceDeleteResponse"> & {
    /**
     * Tenant is the tenant.
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceDeleteResponse.
 * Use `create(TenantServiceDeleteResponseSchema)` to create a new message.
 */
export declare const TenantServiceDeleteResponseSchema: GenMessage<TenantServiceDeleteResponse>;
/**
 * TenantServiceInviteRequest is used to invite a member to a tenant.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteRequest
 */
export type TenantServiceInviteRequest = Message<"metalstack.api.v2.TenantServiceInviteRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Role of this user in this tenant.
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
 * TenantServiceInviteRequest is the response payload to a invite member request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteResponse
 */
export type TenantServiceInviteResponse = Message<"metalstack.api.v2.TenantServiceInviteResponse"> & {
    /**
     * Invite contains a secret which can be sent to a potential user.
     *
     * @generated from field: metalstack.api.v2.TenantInvite invite = 1;
     */
    invite?: TenantInvite | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteResponse.
 * Use `create(TenantServiceInviteResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteResponseSchema: GenMessage<TenantServiceInviteResponse>;
/**
 * TenantServiceInvitesListRequest is the request payload to a list invites request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInvitesListRequest
 */
export type TenantServiceInvitesListRequest = Message<"metalstack.api.v2.TenantServiceInvitesListRequest"> & {
    /**
     * Login of the tenant.
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
 * TenantServiceInvitesListResponse is the response payload to a list invites request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInvitesListResponse
 */
export type TenantServiceInvitesListResponse = Message<"metalstack.api.v2.TenantServiceInvitesListResponse"> & {
    /**
     * Invites that have not yet accepted the invitation to this tenant.
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
 * TenantServiceInviteGetRequest is the request payload to get a invite.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteGetRequest
 */
export type TenantServiceInviteGetRequest = Message<"metalstack.api.v2.TenantServiceInviteGetRequest"> & {
    /**
     * Secret of the invite to get.
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
 * TenantServiceInviteGetResponse is the response payload to a get invite request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteGetResponse
 */
export type TenantServiceInviteGetResponse = Message<"metalstack.api.v2.TenantServiceInviteGetResponse"> & {
    /**
     * Invite is the invite.
     *
     * @generated from field: metalstack.api.v2.TenantInvite invite = 1;
     */
    invite?: TenantInvite | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantServiceInviteGetResponse.
 * Use `create(TenantServiceInviteGetResponseSchema)` to create a new message.
 */
export declare const TenantServiceInviteGetResponseSchema: GenMessage<TenantServiceInviteGetResponse>;
/**
 * TenantServiceInviteAcceptRequest is the request payload to a accept invite request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteAcceptRequest
 */
export type TenantServiceInviteAcceptRequest = Message<"metalstack.api.v2.TenantServiceInviteAcceptRequest"> & {
    /**
     * Secret is the invitation secret part of the invitation url.
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
 * TenantServiceInviteAcceptResponse is the response payload to a accept invite request.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteAcceptResponse
 */
export type TenantServiceInviteAcceptResponse = Message<"metalstack.api.v2.TenantServiceInviteAcceptResponse"> & {
    /**
     * Tenant ID of the joined tenant.
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * TenantName of the joined tenant.
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
 * TenantServiceInviteDeleteRequest is the request payload to a delete invite.
 *
 * @generated from message metalstack.api.v2.TenantServiceInviteDeleteRequest
 */
export type TenantServiceInviteDeleteRequest = Message<"metalstack.api.v2.TenantServiceInviteDeleteRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Secret of the invite to delete.
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
 * TenantServiceInviteDeleteResponse is the response payload of a delete invite request.
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
 * TenantService provides tenant management operations.
 *
 * @generated from service metalstack.api.v2.TenantService
 */
export declare const TenantService: GenService<{
    /**
     * Creates a new tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantServiceCreateRequestSchema;
        output: typeof TenantServiceCreateResponseSchema;
    };
    /**
     * Returns the list of tenants.
     *
     * @generated from rpc metalstack.api.v2.TenantService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantServiceListRequestSchema;
        output: typeof TenantServiceListResponseSchema;
    };
    /**
     * Get a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TenantServiceGetRequestSchema;
        output: typeof TenantServiceGetResponseSchema;
    };
    /**
     * Update a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TenantServiceUpdateRequestSchema;
        output: typeof TenantServiceUpdateResponseSchema;
    };
    /**
     * Delete a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof TenantServiceDeleteRequestSchema;
        output: typeof TenantServiceDeleteResponseSchema;
    };
    /**
     * Invite a user to a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.Invite
     */
    invite: {
        methodKind: "unary";
        input: typeof TenantServiceInviteRequestSchema;
        output: typeof TenantServiceInviteResponseSchema;
    };
    /**
     * InviteAccept is called from a user to accept an invitation.
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteAccept
     */
    inviteAccept: {
        methodKind: "unary";
        input: typeof TenantServiceInviteAcceptRequestSchema;
        output: typeof TenantServiceInviteAcceptResponseSchema;
    };
    /**
     * InviteDelete deletes a pending invitation.
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteDelete
     */
    inviteDelete: {
        methodKind: "unary";
        input: typeof TenantServiceInviteDeleteRequestSchema;
        output: typeof TenantServiceInviteDeleteResponseSchema;
    };
    /**
     * InvitesList list all invites to a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantService.InvitesList
     */
    invitesList: {
        methodKind: "unary";
        input: typeof TenantServiceInvitesListRequestSchema;
        output: typeof TenantServiceInvitesListResponseSchema;
    };
    /**
     * InviteGet get an invite.
     *
     * @generated from rpc metalstack.api.v2.TenantService.InviteGet
     */
    inviteGet: {
        methodKind: "unary";
        input: typeof TenantServiceInviteGetRequestSchema;
        output: typeof TenantServiceInviteGetResponseSchema;
    };
}>;
