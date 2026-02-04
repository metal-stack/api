import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { AdminRole, InfraRole, ProjectRole, TenantRole } from "./common_pb";
import type { MethodPermission } from "./token_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/method.proto.
 */
export declare const file_metalstack_api_v2_method: GenFile;
/**
 * MethodServiceListRequest is the request payload to list all public methods
 *
 * @generated from message metalstack.api.v2.MethodServiceListRequest
 */
export type MethodServiceListRequest = Message<"metalstack.api.v2.MethodServiceListRequest"> & {};
/**
 * Describes the message metalstack.api.v2.MethodServiceListRequest.
 * Use `create(MethodServiceListRequestSchema)` to create a new message.
 */
export declare const MethodServiceListRequestSchema: GenMessage<MethodServiceListRequest>;
/**
 * MethodServiceListResponse is the response payload with all public visible methods
 *
 * @generated from message metalstack.api.v2.MethodServiceListResponse
 */
export type MethodServiceListResponse = Message<"metalstack.api.v2.MethodServiceListResponse"> & {
    /**
     * Methods is a list of methods public callable
     *
     * @generated from field: repeated string methods = 1;
     */
    methods: string[];
};
/**
 * Describes the message metalstack.api.v2.MethodServiceListResponse.
 * Use `create(MethodServiceListResponseSchema)` to create a new message.
 */
export declare const MethodServiceListResponseSchema: GenMessage<MethodServiceListResponse>;
/**
 * MethodServiceTokenScopedListRequest is the request payload to list all methods callable with the token present in the request
 *
 * @generated from message metalstack.api.v2.MethodServiceTokenScopedListRequest
 */
export type MethodServiceTokenScopedListRequest = Message<"metalstack.api.v2.MethodServiceTokenScopedListRequest"> & {};
/**
 * Describes the message metalstack.api.v2.MethodServiceTokenScopedListRequest.
 * Use `create(MethodServiceTokenScopedListRequestSchema)` to create a new message.
 */
export declare const MethodServiceTokenScopedListRequestSchema: GenMessage<MethodServiceTokenScopedListRequest>;
/**
 * MethodServiceTokenScopedListResponse is the response payload which contains all methods which are callable with the given token
 *
 * @generated from message metalstack.api.v2.MethodServiceTokenScopedListResponse
 */
export type MethodServiceTokenScopedListResponse = Message<"metalstack.api.v2.MethodServiceTokenScopedListResponse"> & {
    /**
     * Permissions a list of methods which can be called
     *
     * @generated from field: repeated metalstack.api.v2.MethodPermission permissions = 1;
     */
    permissions: MethodPermission[];
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.ProjectRole> project_roles = 3;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.TenantRole> tenant_roles = 4;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.AdminRole admin_role = 5;
     */
    adminRole?: AdminRole;
    /**
     * InfraRole defines the infrastructure role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.InfraRole infra_role = 6;
     */
    infraRole?: InfraRole;
};
/**
 * Describes the message metalstack.api.v2.MethodServiceTokenScopedListResponse.
 * Use `create(MethodServiceTokenScopedListResponseSchema)` to create a new message.
 */
export declare const MethodServiceTokenScopedListResponseSchema: GenMessage<MethodServiceTokenScopedListResponse>;
/**
 * MethodService serves method related functions
 * methods are functions in services
 *
 * @generated from service metalstack.api.v2.MethodService
 */
export declare const MethodService: GenService<{
    /**
     * List all public visible methods
     *
     * @generated from rpc metalstack.api.v2.MethodService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MethodServiceListRequestSchema;
        output: typeof MethodServiceListResponseSchema;
    };
    /**
     * TokenScopedList all methods callable with the token present in the request
     *
     * @generated from rpc metalstack.api.v2.MethodService.TokenScopedList
     */
    tokenScopedList: {
        methodKind: "unary";
        input: typeof MethodServiceTokenScopedListRequestSchema;
        output: typeof MethodServiceTokenScopedListResponseSchema;
    };
}>;
