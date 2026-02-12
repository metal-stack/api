import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { AdminRole, InfraRole, Labels, MachineRole, Meta, ProjectRole, TenantRole, UpdateLabels, UpdateMeta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/token.proto.
 */
export declare const file_metalstack_api_v2_token: GenFile;
/**
 * Token generates a jwt authentication token to access the api
 *
 * There are two different types of tokens, api- and user- tokens
 *
 * A user token is used to authenticate end user requests for example from a cli.
 * The configured roles in a user token are expanded in the api server
 * based on the memberships in other projects and tenants based on the role granted there.
 * User tokens will never contain permissions.
 * Permissions are always derived from the tenant and project roles and memberships.
 *
 * The api token should be used to authenticate services.
 * In contrast to a user token, the api token permissions and roles apply as configured during the token create process.
 *
 * @generated from message metalstack.api.v2.Token
 */
export type Token = Message<"metalstack.api.v2.Token"> & {
    /**
     * Uuid of the jwt token, used to reference it by revoke
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * User who created this token
     *
     * @generated from field: string user = 2;
     */
    user: string;
    /**
     * Meta for this token
     *
     * @generated from field: metalstack.api.v2.Meta meta = 3;
     */
    meta?: Meta;
    /**
     * Description is a user given description of this token.
     *
     * @generated from field: string description = 4;
     */
    description: string;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated metalstack.api.v2.MethodPermission permissions = 5;
     */
    permissions: MethodPermission[];
    /**
     * Expires gives the date in the future after which this token can not be used anymore
     *
     * @generated from field: google.protobuf.Timestamp expires = 6;
     */
    expires?: Timestamp;
    /**
     * IssuedAt gives the date when this token was created
     *
     * @generated from field: google.protobuf.Timestamp issued_at = 7;
     */
    issuedAt?: Timestamp;
    /**
     * TokenType describes the type of this token
     *
     * @generated from field: metalstack.api.v2.TokenType token_type = 8;
     */
    tokenType: TokenType;
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.ProjectRole> project_roles = 9;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.TenantRole> tenant_roles = 10;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.AdminRole admin_role = 11;
     */
    adminRole?: AdminRole;
    /**
     * InfraRole defines the infrastructure role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.InfraRole infra_role = 12;
     */
    infraRole?: InfraRole;
    /**
     * MachineRoles associates a machine uuid with the corresponding role of the token owner
     * TODO: decide if we need this map from machine.uuid->role, we could instead use the subject in the token instead
     *
     * @generated from field: map<string, metalstack.api.v2.MachineRole> machine_roles = 13;
     */
    machineRoles: {
        [key: string]: MachineRole;
    };
};
/**
 * Describes the message metalstack.api.v2.Token.
 * Use `create(TokenSchema)` to create a new message.
 */
export declare const TokenSchema: GenMessage<Token>;
/**
 * TokenServiceCreateRequest is the request payload to create a token
 *
 * @generated from message metalstack.api.v2.TokenServiceCreateRequest
 */
export type TokenServiceCreateRequest = Message<"metalstack.api.v2.TokenServiceCreateRequest"> & {
    /**
     * Description of the token
     *
     * @generated from field: string description = 1;
     */
    description: string;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated metalstack.api.v2.MethodPermission permissions = 2;
     */
    permissions: MethodPermission[];
    /**
     * Expires gives the duration since now, after which this token can not be used anymore
     *
     * @generated from field: google.protobuf.Duration expires = 4;
     */
    expires?: Duration;
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.ProjectRole> project_roles = 5;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.TenantRole> tenant_roles = 6;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.AdminRole admin_role = 7;
     */
    adminRole?: AdminRole;
    /**
     * InfraRole defines the infrastructure role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.InfraRole infra_role = 8;
     */
    infraRole?: InfraRole;
    /**
     * MachineRoles associates a machine uuid with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.MachineRole> machine_roles = 9;
     */
    machineRoles: {
        [key: string]: MachineRole;
    };
    /**
     * Labels on this token
     *
     * @generated from field: metalstack.api.v2.Labels labels = 10;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceCreateRequest.
 * Use `create(TokenServiceCreateRequestSchema)` to create a new message.
 */
export declare const TokenServiceCreateRequestSchema: GenMessage<TokenServiceCreateRequest>;
/**
 * MethodPermission is a mapping from a subject/project to a service method
 *
 * @generated from message metalstack.api.v2.MethodPermission
 */
export type MethodPermission = Message<"metalstack.api.v2.MethodPermission"> & {
    /**
     * Subject maybe either the project or the tenant
     * for which the methods should be allowed
     *
     * asterisk (*) can be specified to match any subject
     * empty string ("") can be specified for requests that do not require a subject, e.g. partition list
     * otherwise either a projectid or a tenant login should be specified
     *
     * @generated from field: string subject = 1;
     */
    subject: string;
    /**
     * Methods which should be accessible
     *
     * @generated from field: repeated string methods = 2;
     */
    methods: string[];
};
/**
 * Describes the message metalstack.api.v2.MethodPermission.
 * Use `create(MethodPermissionSchema)` to create a new message.
 */
export declare const MethodPermissionSchema: GenMessage<MethodPermission>;
/**
 * TokenServiceCreateResponse is the response payload of a token create request
 *
 * @generated from message metalstack.api.v2.TokenServiceCreateResponse
 */
export type TokenServiceCreateResponse = Message<"metalstack.api.v2.TokenServiceCreateResponse"> & {
    /**
     * Token which was created
     *
     * @generated from field: metalstack.api.v2.Token token = 1;
     */
    token?: Token;
    /**
     * Secret is the body if the jwt token, should be used in api requests as bearer token
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceCreateResponse.
 * Use `create(TokenServiceCreateResponseSchema)` to create a new message.
 */
export declare const TokenServiceCreateResponseSchema: GenMessage<TokenServiceCreateResponse>;
/**
 * TokenServiceListRequest is the request payload to list tokens
 *
 * @generated from message metalstack.api.v2.TokenServiceListRequest
 */
export type TokenServiceListRequest = Message<"metalstack.api.v2.TokenServiceListRequest"> & {};
/**
 * Describes the message metalstack.api.v2.TokenServiceListRequest.
 * Use `create(TokenServiceListRequestSchema)` to create a new message.
 */
export declare const TokenServiceListRequestSchema: GenMessage<TokenServiceListRequest>;
/**
 * TokenServiceListResponse is the response payload of a token list request
 *
 * @generated from message metalstack.api.v2.TokenServiceListResponse
 */
export type TokenServiceListResponse = Message<"metalstack.api.v2.TokenServiceListResponse"> & {
    /**
     * Tokens is a list of tokens without the secrets
     *
     * @generated from field: repeated metalstack.api.v2.Token tokens = 1;
     */
    tokens: Token[];
};
/**
 * Describes the message metalstack.api.v2.TokenServiceListResponse.
 * Use `create(TokenServiceListResponseSchema)` to create a new message.
 */
export declare const TokenServiceListResponseSchema: GenMessage<TokenServiceListResponse>;
/**
 * TokenServiceRevokeRequest is the request payload of a token revoke request
 *
 * @generated from message metalstack.api.v2.TokenServiceRevokeRequest
 */
export type TokenServiceRevokeRequest = Message<"metalstack.api.v2.TokenServiceRevokeRequest"> & {
    /**
     * Uuid of the token to revoke
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceRevokeRequest.
 * Use `create(TokenServiceRevokeRequestSchema)` to create a new message.
 */
export declare const TokenServiceRevokeRequestSchema: GenMessage<TokenServiceRevokeRequest>;
/**
 * TokenServiceRevokeResponse is the response payload of a token revoke request
 *
 * @generated from message metalstack.api.v2.TokenServiceRevokeResponse
 */
export type TokenServiceRevokeResponse = Message<"metalstack.api.v2.TokenServiceRevokeResponse"> & {};
/**
 * Describes the message metalstack.api.v2.TokenServiceRevokeResponse.
 * Use `create(TokenServiceRevokeResponseSchema)` to create a new message.
 */
export declare const TokenServiceRevokeResponseSchema: GenMessage<TokenServiceRevokeResponse>;
/**
 * TokenServiceUpdateRequest is the request payload of a token update request
 *
 * @generated from message metalstack.api.v2.TokenServiceUpdateRequest
 */
export type TokenServiceUpdateRequest = Message<"metalstack.api.v2.TokenServiceUpdateRequest"> & {
    /**
     * Uuid of the token to update
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     * TokenUpdate is not guarded with optlock in the backend
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Description is a user given description of this token.
     *
     * @generated from field: optional string description = 3;
     */
    description?: string;
    /**
     * Permissions is a list of service methods this token can be used for
     *
     * @generated from field: repeated metalstack.api.v2.MethodPermission permissions = 4;
     */
    permissions: MethodPermission[];
    /**
     * ProjectRoles associates a project id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.ProjectRole> project_roles = 5;
     */
    projectRoles: {
        [key: string]: ProjectRole;
    };
    /**
     * TenantRoles associates a tenant id with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.TenantRole> tenant_roles = 6;
     */
    tenantRoles: {
        [key: string]: TenantRole;
    };
    /**
     * AdminRole defines the admin role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.AdminRole admin_role = 7;
     */
    adminRole?: AdminRole;
    /**
     * InfraRole defines the infrastructure role of the token owner
     *
     * @generated from field: optional metalstack.api.v2.InfraRole infra_role = 8;
     */
    infraRole?: InfraRole;
    /**
     * MachineRoles associates a machine uuid with the corresponding role of the token owner
     *
     * @generated from field: map<string, metalstack.api.v2.MachineRole> machine_roles = 9;
     */
    machineRoles: {
        [key: string]: MachineRole;
    };
    /**
     * Labels on this token
     *
     * @generated from field: metalstack.api.v2.UpdateLabels labels = 10;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceUpdateRequest.
 * Use `create(TokenServiceUpdateRequestSchema)` to create a new message.
 */
export declare const TokenServiceUpdateRequestSchema: GenMessage<TokenServiceUpdateRequest>;
/**
 * TokenServiceUpdateResponse is the response payload of a token update request
 *
 * @generated from message metalstack.api.v2.TokenServiceUpdateResponse
 */
export type TokenServiceUpdateResponse = Message<"metalstack.api.v2.TokenServiceUpdateResponse"> & {
    /**
     * Token is the updated token
     *
     * @generated from field: metalstack.api.v2.Token token = 1;
     */
    token?: Token;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceUpdateResponse.
 * Use `create(TokenServiceUpdateResponseSchema)` to create a new message.
 */
export declare const TokenServiceUpdateResponseSchema: GenMessage<TokenServiceUpdateResponse>;
/**
 * TokenServiceGetRequest is the request payload of a token get request
 *
 * @generated from message metalstack.api.v2.TokenServiceGetRequest
 */
export type TokenServiceGetRequest = Message<"metalstack.api.v2.TokenServiceGetRequest"> & {
    /**
     * Uuid of the token to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceGetRequest.
 * Use `create(TokenServiceGetRequestSchema)` to create a new message.
 */
export declare const TokenServiceGetRequestSchema: GenMessage<TokenServiceGetRequest>;
/**
 * TokenServiceGetResponse is the response payload of a token get request
 *
 * @generated from message metalstack.api.v2.TokenServiceGetResponse
 */
export type TokenServiceGetResponse = Message<"metalstack.api.v2.TokenServiceGetResponse"> & {
    /**
     * Token is the token
     *
     * @generated from field: metalstack.api.v2.Token token = 1;
     */
    token?: Token;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceGetResponse.
 * Use `create(TokenServiceGetResponseSchema)` to create a new message.
 */
export declare const TokenServiceGetResponseSchema: GenMessage<TokenServiceGetResponse>;
/**
 * TokenServiceRefreshRequest is the request payload of a token refresh request
 * Permissions, Roles and Expiration duration and all other properties are inherited from the calling token.
 * The expiration duration will be calculated from the existing token (exp - iat)
 *
 * @generated from message metalstack.api.v2.TokenServiceRefreshRequest
 */
export type TokenServiceRefreshRequest = Message<"metalstack.api.v2.TokenServiceRefreshRequest"> & {};
/**
 * Describes the message metalstack.api.v2.TokenServiceRefreshRequest.
 * Use `create(TokenServiceRefreshRequestSchema)` to create a new message.
 */
export declare const TokenServiceRefreshRequestSchema: GenMessage<TokenServiceRefreshRequest>;
/**
 * TokenServiceRefreshResponse is the response payload of a token refresh request
 *
 * @generated from message metalstack.api.v2.TokenServiceRefreshResponse
 */
export type TokenServiceRefreshResponse = Message<"metalstack.api.v2.TokenServiceRefreshResponse"> & {
    /**
     * Token which was refreshed
     *
     * @generated from field: metalstack.api.v2.Token token = 1;
     */
    token?: Token;
    /**
     * Secret is the body if the jwt token, should be used in api requests as bearer token
     *
     * @generated from field: string secret = 2;
     */
    secret: string;
};
/**
 * Describes the message metalstack.api.v2.TokenServiceRefreshResponse.
 * Use `create(TokenServiceRefreshResponseSchema)` to create a new message.
 */
export declare const TokenServiceRefreshResponseSchema: GenMessage<TokenServiceRefreshResponse>;
/**
 * TokenType specifies different use cases of tokens
 *
 * @generated from enum metalstack.api.v2.TokenType
 */
export declare enum TokenType {
    /**
     * TOKEN_TYPE_UNSPECIFIED is not specified
     *
     * @generated from enum value: TOKEN_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * TOKEN_TYPE_API is a token for api usage
     *
     * @generated from enum value: TOKEN_TYPE_API = 1;
     */
    API = 1,
    /**
     * TOKEN_TYPE_USER is a token to access the api with cli, a web application or other user induced actions.
     *
     * @generated from enum value: TOKEN_TYPE_USER = 2;
     */
    USER = 2
}
/**
 * Describes the enum metalstack.api.v2.TokenType.
 */
export declare const TokenTypeSchema: GenEnum<TokenType>;
/**
 * TokenService serves token related functions
 *
 * @generated from service metalstack.api.v2.TokenService
 */
export declare const TokenService: GenService<{
    /**
     * Get a token
     *
     * @generated from rpc metalstack.api.v2.TokenService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof TokenServiceGetRequestSchema;
        output: typeof TokenServiceGetResponseSchema;
    };
    /**
     * Create a token to authenticate against the platform, the secret will be only visible in the response.
     *
     * @generated from rpc metalstack.api.v2.TokenService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TokenServiceCreateRequestSchema;
        output: typeof TokenServiceCreateResponseSchema;
    };
    /**
     * Update a token
     *
     * @generated from rpc metalstack.api.v2.TokenService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof TokenServiceUpdateRequestSchema;
        output: typeof TokenServiceUpdateResponseSchema;
    };
    /**
     * List all your tokens
     *
     * @generated from rpc metalstack.api.v2.TokenService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TokenServiceListRequestSchema;
        output: typeof TokenServiceListResponseSchema;
    };
    /**
     * Revoke a token, no further usage is possible afterwards
     *
     * @generated from rpc metalstack.api.v2.TokenService.Revoke
     */
    revoke: {
        methodKind: "unary";
        input: typeof TokenServiceRevokeRequestSchema;
        output: typeof TokenServiceRevokeResponseSchema;
    };
    /**
     * Refresh a token, this will create a new token with the exact same permissions as the calling token contains
     *
     * @generated from rpc metalstack.api.v2.TokenService.Refresh
     */
    refresh: {
        methodKind: "unary";
        input: typeof TokenServiceRefreshRequestSchema;
        output: typeof TokenServiceRefreshResponseSchema;
    };
}>;
