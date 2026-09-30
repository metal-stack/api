import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Meta, TenantRole, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/tenant_member.proto.
 */
export declare const file_metalstack_api_v2_tenant_member: GenFile;
/**
 * TenantMember defines a user that participates in a tenant.
 *
 * @generated from message metalstack.api.v2.TenantMember
 */
export type TenantMember = Message<"metalstack.api.v2.TenantMember"> & {
    /**
     * Member is the user id of the member.
     *
     * @generated from field: string member = 1;
     */
    member: string;
    /**
     * Role is the role of the member.
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 2;
     */
    role: TenantRole;
    /**
     * Projects in which a user is a direct member.
     *
     * @generated from field: repeated string projects = 3;
     */
    projects: string[];
    /**
     * CreatedAt the date when the member was added to the tenant.
     *
     * @generated from field: google.protobuf.Timestamp created_at = 4;
     */
    createdAt?: Timestamp | undefined;
    /**
     * Meta for this tenant member.
     *
     * @generated from field: metalstack.api.v2.Meta meta = 5;
     */
    meta?: Meta | undefined;
    /**
     * Tenant is the tenant that this member participates in.
     *
     * @generated from field: string tenant = 6;
     */
    tenant: string;
};
/**
 * Describes the message metalstack.api.v2.TenantMember.
 * Use `create(TenantMemberSchema)` to create a new message.
 */
export declare const TenantMemberSchema: GenMessage<TenantMember>;
/**
 * TenantMemberServiceListRequest is the request payload of the tenant member list request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceListRequest
 */
export type TenantMemberServiceListRequest = Message<"metalstack.api.v2.TenantMemberServiceListRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Query for tenant members.
     *
     * @generated from field: metalstack.api.v2.TenantMemberQuery query = 2;
     */
    query?: TenantMemberQuery | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceListRequest.
 * Use `create(TenantMemberServiceListRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceListRequestSchema: GenMessage<TenantMemberServiceListRequest>;
/**
 * TenantMemberQuery is used to search tenant members.
 *
 * @generated from message metalstack.api.v2.TenantMemberQuery
 */
export type TenantMemberQuery = Message<"metalstack.api.v2.TenantMemberQuery"> & {
    /**
     * Member is the user id of the member.
     *
     * @generated from field: optional string member = 1;
     */
    member?: string | undefined;
    /**
     * Role is the role of the member.
     *
     * @generated from field: optional metalstack.api.v2.TenantRole role = 2;
     */
    role?: TenantRole | undefined;
    /**
     * Projects in which a user is a direct member.
     *
     * @generated from field: repeated string projects = 3;
     */
    projects: string[];
};
/**
 * Describes the message metalstack.api.v2.TenantMemberQuery.
 * Use `create(TenantMemberQuerySchema)` to create a new message.
 */
export declare const TenantMemberQuerySchema: GenMessage<TenantMemberQuery>;
/**
 * TenantMemberServiceGetRequest is the request payload of the tenant member get request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceGetRequest
 */
export type TenantMemberServiceGetRequest = Message<"metalstack.api.v2.TenantMemberServiceGetRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Member is the user id of the member.
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceGetRequest.
 * Use `create(TenantMemberServiceGetRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceGetRequestSchema: GenMessage<TenantMemberServiceGetRequest>;
/**
 * TenantMemberServiceDeleteRequest is used to remove a member from a tenant.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceDeleteRequest
 */
export type TenantMemberServiceDeleteRequest = Message<"metalstack.api.v2.TenantMemberServiceDeleteRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Member is the id of the member to remove from this tenant.
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceDeleteRequest.
 * Use `create(TenantMemberServiceDeleteRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceDeleteRequestSchema: GenMessage<TenantMemberServiceDeleteRequest>;
/**
 * TenantMemberServiceCreateRequest is the request payload for adding a member to a tenant.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceCreateRequest
 */
export type TenantMemberServiceCreateRequest = Message<"metalstack.api.v2.TenantMemberServiceCreateRequest"> & {
    /**
     * Login of the tenant to which the member will be added.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Login of the member to add.
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role to assign to the new member.
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 3;
     */
    role: TenantRole;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceCreateRequest.
 * Use `create(TenantMemberServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceCreateRequestSchema: GenMessage<TenantMemberServiceCreateRequest>;
/**
 * TenantMemberServiceUpdateRequest is used to update a member from a tenant.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceUpdateRequest
 */
export type TenantMemberServiceUpdateRequest = Message<"metalstack.api.v2.TenantMemberServiceUpdateRequest"> & {
    /**
     * Login of the tenant.
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Member is the id of the member to update in this tenant.
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role of this user in this tenant.
     *
     * @generated from field: optional metalstack.api.v2.TenantRole role = 3;
     */
    role?: TenantRole | undefined;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 4;
     */
    updateMeta?: UpdateMeta | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceUpdateRequest.
 * Use `create(TenantMemberServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceUpdateRequestSchema: GenMessage<TenantMemberServiceUpdateRequest>;
/**
 * TenantMemberServiceGetResponse is the response payload of the tenant member get request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceGetResponse
 */
export type TenantMemberServiceGetResponse = Message<"metalstack.api.v2.TenantMemberServiceGetResponse"> & {
    /**
     * Member is the member of this tenant.
     *
     * @generated from field: metalstack.api.v2.TenantMember member = 1;
     */
    member?: TenantMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceGetResponse.
 * Use `create(TenantMemberServiceGetResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceGetResponseSchema: GenMessage<TenantMemberServiceGetResponse>;
/**
 * TenantMemberServiceListResponse is the response payload of the tenant member list request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceListResponse
 */
export type TenantMemberServiceListResponse = Message<"metalstack.api.v2.TenantMemberServiceListResponse"> & {
    /**
     * Members is the list of tenant members.
     *
     * @generated from field: repeated metalstack.api.v2.TenantMember members = 1;
     */
    members: TenantMember[];
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceListResponse.
 * Use `create(TenantMemberServiceListResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceListResponseSchema: GenMessage<TenantMemberServiceListResponse>;
/**
 * TenantMemberServiceCreateResponse is the response payload of the tenant member create request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceCreateResponse
 */
export type TenantMemberServiceCreateResponse = Message<"metalstack.api.v2.TenantMemberServiceCreateResponse"> & {
    /**
     * Member is the member of this tenant.
     *
     * @generated from field: metalstack.api.v2.TenantMember member = 1;
     */
    member?: TenantMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceCreateResponse.
 * Use `create(TenantMemberServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceCreateResponseSchema: GenMessage<TenantMemberServiceCreateResponse>;
/**
 * TenantMemberServiceUpdateResponse is the response payload of the tenant member update request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceUpdateResponse
 */
export type TenantMemberServiceUpdateResponse = Message<"metalstack.api.v2.TenantMemberServiceUpdateResponse"> & {
    /**
     * Member is the member of this tenant.
     *
     * @generated from field: metalstack.api.v2.TenantMember member = 1;
     */
    member?: TenantMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceUpdateResponse.
 * Use `create(TenantMemberServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceUpdateResponseSchema: GenMessage<TenantMemberServiceUpdateResponse>;
/**
 * TenantMemberServiceDeleteResponse is the response payload of the tenant member delete request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceDeleteResponse
 */
export type TenantMemberServiceDeleteResponse = Message<"metalstack.api.v2.TenantMemberServiceDeleteResponse"> & {
    /**
     * Member is the member of this tenant.
     *
     * @generated from field: metalstack.api.v2.TenantMember member = 1;
     */
    member?: TenantMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceDeleteResponse.
 * Use `create(TenantMemberServiceDeleteResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceDeleteResponseSchema: GenMessage<TenantMemberServiceDeleteResponse>;
/**
 * TenantMemberServiceLeaveRequest is used to leave a tenant. This way a member is not required to ask a tenant owner to remove him.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceLeaveRequest
 */
export type TenantMemberServiceLeaveRequest = Message<"metalstack.api.v2.TenantMemberServiceLeaveRequest"> & {
    /**
     * Login of the tenant that the user wants to leave.
     *
     * @generated from field: string login = 1;
     */
    login: string;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceLeaveRequest.
 * Use `create(TenantMemberServiceLeaveRequestSchema)` to create a new message.
 */
export declare const TenantMemberServiceLeaveRequestSchema: GenMessage<TenantMemberServiceLeaveRequest>;
/**
 * TenantMemberServiceLeaveResponse is the response payload to a leave tenant request.
 *
 * @generated from message metalstack.api.v2.TenantMemberServiceLeaveResponse
 */
export type TenantMemberServiceLeaveResponse = Message<"metalstack.api.v2.TenantMemberServiceLeaveResponse"> & {
    /**
     * Member is the member of this tenant.
     *
     * @generated from field: metalstack.api.v2.TenantMember member = 1;
     */
    member?: TenantMember | undefined;
};
/**
 * Describes the message metalstack.api.v2.TenantMemberServiceLeaveResponse.
 * Use `create(TenantMemberServiceLeaveResponseSchema)` to create a new message.
 */
export declare const TenantMemberServiceLeaveResponseSchema: GenMessage<TenantMemberServiceLeaveResponse>;
/**
 * TenantMemberService provides tenant member management operations.
 *
 * @generated from service metalstack.api.v2.TenantMemberService
 */
export declare const TenantMemberService: GenService<{
    /**
     * Creates a new tenant member.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantMemberServiceCreateRequestSchema;
        output: typeof TenantMemberServiceCreateResponseSchema;
    };
    /**
     * Returns the list of tenant members.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantMemberServiceListRequestSchema;
        output: typeof TenantMemberServiceListResponseSchema;
    };
    /**
     * Get a tenant member.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TenantMemberServiceGetRequestSchema;
        output: typeof TenantMemberServiceGetResponseSchema;
    };
    /**
     * Update a tenant member.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TenantMemberServiceUpdateRequestSchema;
        output: typeof TenantMemberServiceUpdateResponseSchema;
    };
    /**
     * Delete a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof TenantMemberServiceDeleteRequestSchema;
        output: typeof TenantMemberServiceDeleteResponseSchema;
    };
    /**
     * Leave removes a member from a tenant.
     *
     * @generated from rpc metalstack.api.v2.TenantMemberService.Leave
     */
    leave: {
        methodKind: "unary";
        input: typeof TenantMemberServiceLeaveRequestSchema;
        output: typeof TenantMemberServiceLeaveResponseSchema;
    };
}>;
