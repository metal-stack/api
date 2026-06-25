import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, TenantRole } from "../../api/v2/common_pb";
import type { Tenant, TenantQuery } from "../../api/v2/tenant_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/tenant.proto.
 */
export declare const file_metalstack_admin_v2_tenant: GenFile;
/**
 * TenantServiceCreateRequest is the request payload for creating a tenant.
 *
 * @generated from message metalstack.admin.v2.TenantServiceCreateRequest
 */
export type TenantServiceCreateRequest = Message<"metalstack.admin.v2.TenantServiceCreateRequest"> & {
    /**
     * Name of the tenant
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Description of the tenant
     *
     * @generated from field: optional string description = 2;
     */
    description?: string | undefined;
    /**
     * Email of the tenant, if not set will be inherited from the creator
     *
     * @generated from field: optional string email = 3;
     */
    email?: string | undefined;
    /**
     * AvatarUrl of the tenant
     *
     * @generated from field: optional string avatar_url = 4;
     */
    avatarUrl?: string | undefined;
    /**
     * Labels on the tenant
     *
     * @generated from field: metalstack.api.v2.Labels labels = 5;
     */
    labels?: Labels | undefined;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceCreateRequest.
 * Use `create(TenantServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantServiceCreateRequestSchema: GenMessage<TenantServiceCreateRequest>;
/**
 * TenantServiceCreateResponse is the response payload for creating a tenant.
 *
 * @generated from message metalstack.admin.v2.TenantServiceCreateResponse
 */
export type TenantServiceCreateResponse = Message<"metalstack.admin.v2.TenantServiceCreateResponse"> & {
    /**
     * Tenant contains the created tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceCreateResponse.
 * Use `create(TenantServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantServiceCreateResponseSchema: GenMessage<TenantServiceCreateResponse>;
/**
 * TenantServiceListRequest is the request payload for listing tenants.
 *
 * @generated from message metalstack.admin.v2.TenantServiceListRequest
 */
export type TenantServiceListRequest = Message<"metalstack.admin.v2.TenantServiceListRequest"> & {
    /**
     * Query for tenants
     *
     * @generated from field: metalstack.api.v2.TenantQuery query = 1;
     */
    query?: TenantQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceListRequest.
 * Use `create(TenantServiceListRequestSchema)` to create a new message.
 */
export declare const TenantServiceListRequestSchema: GenMessage<TenantServiceListRequest>;
/**
 * TenantServiceListResponse is the response payload for listing tenants.
 *
 * @generated from message metalstack.admin.v2.TenantServiceListResponse
 */
export type TenantServiceListResponse = Message<"metalstack.admin.v2.TenantServiceListResponse"> & {
    /**
     * Tenants contains the list of tenants
     *
     * @generated from field: repeated metalstack.api.v2.Tenant tenants = 1;
     */
    tenants: Tenant[];
    /**
     * NextPage is used for pagination, returns the next page to be fetched and must then be provided in the list request
     *
     * @generated from field: optional uint64 next_page = 2;
     */
    nextPage?: bigint | undefined;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceListResponse.
 * Use `create(TenantServiceListResponseSchema)` to create a new message.
 */
export declare const TenantServiceListResponseSchema: GenMessage<TenantServiceListResponse>;
/**
 * TenantServiceAddMemberRequest is the request payload for adding a member to a tenant
 *
 * @generated from message metalstack.admin.v2.TenantServiceAddMemberRequest
 */
export type TenantServiceAddMemberRequest = Message<"metalstack.admin.v2.TenantServiceAddMemberRequest"> & {
    /**
     * Login of the tenant to which the member will be added
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Login of the member to add
     *
     * @generated from field: string member = 2;
     */
    member: string;
    /**
     * Role to assign to the new member
     *
     * @generated from field: metalstack.api.v2.TenantRole role = 3;
     */
    role: TenantRole;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceAddMemberRequest.
 * Use `create(TenantServiceAddMemberRequestSchema)` to create a new message.
 */
export declare const TenantServiceAddMemberRequestSchema: GenMessage<TenantServiceAddMemberRequest>;
/**
 * TenantServiceAddMemberResponse is the response payload for the add member request
 *
 * @generated from message metalstack.admin.v2.TenantServiceAddMemberResponse
 */
export type TenantServiceAddMemberResponse = Message<"metalstack.admin.v2.TenantServiceAddMemberResponse"> & {};
/**
 * Describes the message metalstack.admin.v2.TenantServiceAddMemberResponse.
 * Use `create(TenantServiceAddMemberResponseSchema)` to create a new message.
 */
export declare const TenantServiceAddMemberResponseSchema: GenMessage<TenantServiceAddMemberResponse>;
/**
 * TenantServiceRemoveMemberRequest is the request payload for removing a member from a tenant
 *
 * @generated from message metalstack.admin.v2.TenantServiceRemoveMemberRequest
 */
export type TenantServiceRemoveMemberRequest = Message<"metalstack.admin.v2.TenantServiceRemoveMemberRequest"> & {
    /**
     * Login of the tenant from which the member will be removed
     *
     * @generated from field: string tenant = 1;
     */
    tenant: string;
    /**
     * Login of the member to remove
     *
     * @generated from field: string member = 2;
     */
    member: string;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceRemoveMemberRequest.
 * Use `create(TenantServiceRemoveMemberRequestSchema)` to create a new message.
 */
export declare const TenantServiceRemoveMemberRequestSchema: GenMessage<TenantServiceRemoveMemberRequest>;
/**
 * TenantServiceRemoveMemberResponse is the response payload for the remove member request
 *
 * @generated from message metalstack.admin.v2.TenantServiceRemoveMemberResponse
 */
export type TenantServiceRemoveMemberResponse = Message<"metalstack.admin.v2.TenantServiceRemoveMemberResponse"> & {};
/**
 * Describes the message metalstack.admin.v2.TenantServiceRemoveMemberResponse.
 * Use `create(TenantServiceRemoveMemberResponseSchema)` to create a new message.
 */
export declare const TenantServiceRemoveMemberResponseSchema: GenMessage<TenantServiceRemoveMemberResponse>;
/**
 * TenantService provides tenant management operations.
 *
 * @generated from service metalstack.admin.v2.TenantService
 */
export declare const TenantService: GenService<{
    /**
     * Creates a new tenant.
     *
     * @generated from rpc metalstack.admin.v2.TenantService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantServiceCreateRequestSchema;
        output: typeof TenantServiceCreateResponseSchema;
    };
    /**
     * Returns the list of all tenants.
     *
     * @generated from rpc metalstack.admin.v2.TenantService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantServiceListRequestSchema;
        output: typeof TenantServiceListResponseSchema;
    };
    /**
     * Add a member to a tenant
     *
     * @generated from rpc metalstack.admin.v2.TenantService.AddMember
     */
    addMember: {
        methodKind: "unary";
        input: typeof TenantServiceAddMemberRequestSchema;
        output: typeof TenantServiceAddMemberResponseSchema;
    };
    /**
     * RemoveMember remove a member of a tenant
     *
     * @generated from rpc metalstack.admin.v2.TenantService.RemoveMember
     */
    removeMember: {
        methodKind: "unary";
        input: typeof TenantServiceRemoveMemberRequestSchema;
        output: typeof TenantServiceRemoveMemberResponseSchema;
    };
}>;
