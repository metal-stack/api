import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Token, TokenServiceCreateRequest as TokenServiceCreateRequest$1 } from "../../api/v2/token_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/token.proto.
 */
export declare const file_metalstack_admin_v2_token: GenFile;
/**
 * TokenServiceListRequest is the request payload for the token list request
 *
 * @generated from message metalstack.admin.v2.TokenServiceListRequest
 */
export type TokenServiceListRequest = Message<"metalstack.admin.v2.TokenServiceListRequest"> & {
    /**
     * User is the id of the user for which the tokens should be listed
     *
     * @generated from field: optional string user = 1;
     */
    user?: string;
};
/**
 * Describes the message metalstack.admin.v2.TokenServiceListRequest.
 * Use `create(TokenServiceListRequestSchema)` to create a new message.
 */
export declare const TokenServiceListRequestSchema: GenMessage<TokenServiceListRequest>;
/**
 * TokenServiceListResponse is the response payload for the token list request
 *
 * @generated from message metalstack.admin.v2.TokenServiceListResponse
 */
export type TokenServiceListResponse = Message<"metalstack.admin.v2.TokenServiceListResponse"> & {
    /**
     * Tokens is the list of tokens
     *
     * @generated from field: repeated metalstack.api.v2.Token tokens = 1;
     */
    tokens: Token[];
};
/**
 * Describes the message metalstack.admin.v2.TokenServiceListResponse.
 * Use `create(TokenServiceListResponseSchema)` to create a new message.
 */
export declare const TokenServiceListResponseSchema: GenMessage<TokenServiceListResponse>;
/**
 * TokenServiceRevokeRequest is the request payload for the token revoke request
 *
 * @generated from message metalstack.admin.v2.TokenServiceRevokeRequest
 */
export type TokenServiceRevokeRequest = Message<"metalstack.admin.v2.TokenServiceRevokeRequest"> & {
    /**
     * Uuid is the uuid of the token which should be revoked
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * User is the id of the user for which the token should be revoked
     *
     * @generated from field: string user = 2;
     */
    user: string;
};
/**
 * Describes the message metalstack.admin.v2.TokenServiceRevokeRequest.
 * Use `create(TokenServiceRevokeRequestSchema)` to create a new message.
 */
export declare const TokenServiceRevokeRequestSchema: GenMessage<TokenServiceRevokeRequest>;
/**
 * TokenServiceRevokeResponse is the response payload for the token revoke request
 *
 * @generated from message metalstack.admin.v2.TokenServiceRevokeResponse
 */
export type TokenServiceRevokeResponse = Message<"metalstack.admin.v2.TokenServiceRevokeResponse"> & {};
/**
 * Describes the message metalstack.admin.v2.TokenServiceRevokeResponse.
 * Use `create(TokenServiceRevokeResponseSchema)` to create a new message.
 */
export declare const TokenServiceRevokeResponseSchema: GenMessage<TokenServiceRevokeResponse>;
/**
 * TokenServiceCreateRequest is the request payload to create a token
 *
 * @generated from message metalstack.admin.v2.TokenServiceCreateRequest
 */
export type TokenServiceCreateRequest = Message<"metalstack.admin.v2.TokenServiceCreateRequest"> & {
    /**
     * User this token should be created for, if omitted, user is derived from caller
     *
     * @generated from field: optional string user = 1;
     */
    user?: string;
    /**
     * TokenCreateRequest which should be created
     *
     * @generated from field: metalstack.api.v2.TokenServiceCreateRequest token_create_request = 2;
     */
    tokenCreateRequest?: TokenServiceCreateRequest$1;
};
/**
 * Describes the message metalstack.admin.v2.TokenServiceCreateRequest.
 * Use `create(TokenServiceCreateRequestSchema)` to create a new message.
 */
export declare const TokenServiceCreateRequestSchema: GenMessage<TokenServiceCreateRequest>;
/**
 * TokenServiceCreateResponse is the response payload of a token create request
 *
 * @generated from message metalstack.admin.v2.TokenServiceCreateResponse
 */
export type TokenServiceCreateResponse = Message<"metalstack.admin.v2.TokenServiceCreateResponse"> & {
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
 * Describes the message metalstack.admin.v2.TokenServiceCreateResponse.
 * Use `create(TokenServiceCreateResponseSchema)` to create a new message.
 */
export declare const TokenServiceCreateResponseSchema: GenMessage<TokenServiceCreateResponse>;
/**
 * TokenService serves token related functions
 *
 * @generated from service metalstack.admin.v2.TokenService
 */
export declare const TokenService: GenService<{
    /**
     * List tokens
     *
     * @generated from rpc metalstack.admin.v2.TokenService.List
     */
    list: {
        methodKind: "unary";
        input: typeof TokenServiceListRequestSchema;
        output: typeof TokenServiceListResponseSchema;
    };
    /**
     * Revoke a token
     *
     * @generated from rpc metalstack.admin.v2.TokenService.Revoke
     */
    revoke: {
        methodKind: "unary";
        input: typeof TokenServiceRevokeRequestSchema;
        output: typeof TokenServiceRevokeResponseSchema;
    };
    /**
     * Create a token to authenticate against the platform, the secret will be only visible in the response.
     * This service is suitable to create tokens for other users instead of deriving users from tokens directly.
     *
     * @generated from rpc metalstack.admin.v2.TokenService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof TokenServiceCreateRequestSchema;
        output: typeof TokenServiceCreateResponseSchema;
    };
}>;
