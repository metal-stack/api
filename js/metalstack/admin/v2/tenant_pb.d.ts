import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Paging } from "../../api/v2/common_pb";
import type { Tenant } from "../../api/v2/tenant_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/tenant.proto.
 */
export declare const file_metalstack_admin_v2_tenant: GenFile;
/**
 * TenantServiceCreateRequest is the request payload of the tenant create request
 *
 * @generated from message metalstack.admin.v2.TenantServiceCreateRequest
 */
export type TenantServiceCreateRequest = Message<"metalstack.admin.v2.TenantServiceCreateRequest"> & {
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
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceCreateRequest.
 * Use `create(TenantServiceCreateRequestSchema)` to create a new message.
 */
export declare const TenantServiceCreateRequestSchema: GenMessage<TenantServiceCreateRequest>;
/**
 * TenantServiceCreateResponse is the response payload of the tenant create request
 *
 * @generated from message metalstack.admin.v2.TenantServiceCreateResponse
 */
export type TenantServiceCreateResponse = Message<"metalstack.admin.v2.TenantServiceCreateResponse"> & {
    /**
     * Tenant is the tenant
     *
     * @generated from field: metalstack.api.v2.Tenant tenant = 1;
     */
    tenant?: Tenant;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceCreateResponse.
 * Use `create(TenantServiceCreateResponseSchema)` to create a new message.
 */
export declare const TenantServiceCreateResponseSchema: GenMessage<TenantServiceCreateResponse>;
/**
 * TenantServiceListRequest is the request payload for a tenant list request
 *
 * @generated from message metalstack.admin.v2.TenantServiceListRequest
 */
export type TenantServiceListRequest = Message<"metalstack.admin.v2.TenantServiceListRequest"> & {
    /**
     * Login of the tenant to list
     *
     * @generated from field: optional string login = 1;
     */
    login?: string;
    /**
     * Name of the tenant to list
     *
     * @generated from field: optional string name = 2;
     */
    name?: string;
    /**
     * Email of the tenant to list
     *
     * @generated from field: optional string email = 3;
     */
    email?: string;
    /**
     * Paging details for the list request
     *
     * @generated from field: metalstack.api.v2.Paging paging = 7;
     */
    paging?: Paging;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceListRequest.
 * Use `create(TenantServiceListRequestSchema)` to create a new message.
 */
export declare const TenantServiceListRequestSchema: GenMessage<TenantServiceListRequest>;
/**
 * TenantServiceListResponse is the response payload for a tenant list request
 *
 * @generated from message metalstack.admin.v2.TenantServiceListResponse
 */
export type TenantServiceListResponse = Message<"metalstack.admin.v2.TenantServiceListResponse"> & {
    /**
     * Tenants are the list of tenants
     *
     * @generated from field: repeated metalstack.api.v2.Tenant tenants = 1;
     */
    tenants: Tenant[];
    /**
     * NextPage is used for pagination, returns the next page to be fetched and must then be provided in the list request.
     *
     * @generated from field: optional uint64 next_page = 2;
     */
    nextPage?: bigint;
};
/**
 * Describes the message metalstack.admin.v2.TenantServiceListResponse.
 * Use `create(TenantServiceListResponseSchema)` to create a new message.
 */
export declare const TenantServiceListResponseSchema: GenMessage<TenantServiceListResponse>;
/**
 * TenantService serves tenant related functions
 *
 * @generated from service metalstack.admin.v2.TenantService
 */
export declare const TenantService: GenService<{
    /**
     * Create a tenant
     *
     * @generated from rpc metalstack.admin.v2.TenantService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TenantServiceCreateRequestSchema;
        output: typeof TenantServiceCreateResponseSchema;
    };
    /**
     * List all tenants
     *
     * @generated from rpc metalstack.admin.v2.TenantService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TenantServiceListRequestSchema;
        output: typeof TenantServiceListResponseSchema;
    };
}>;
