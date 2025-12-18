import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/version.proto.
 */
export declare const file_metalstack_api_v2_version: GenFile;
/**
 * Version of the application
 *
 * @generated from message metalstack.api.v2.Version
 */
export type Version = Message<"metalstack.api.v2.Version"> & {
    /**
     * Version of the application
     *
     * @generated from field: string version = 1;
     */
    version: string;
    /**
     * Revision of the application
     *
     * @generated from field: string revision = 2;
     */
    revision: string;
    /**
     * GitSHA1 of the application
     *
     * @generated from field: string git_sha1 = 3;
     */
    gitSha1: string;
    /**
     * BuildDate of the application
     *
     * @generated from field: string build_date = 4;
     */
    buildDate: string;
};
/**
 * Describes the message metalstack.api.v2.Version.
 * Use `create(VersionSchema)` to create a new message.
 */
export declare const VersionSchema: GenMessage<Version>;
/**
 * VersionServiceGetRequest is the request payload to get the version
 *
 * @generated from message metalstack.api.v2.VersionServiceGetRequest
 */
export type VersionServiceGetRequest = Message<"metalstack.api.v2.VersionServiceGetRequest"> & {};
/**
 * Describes the message metalstack.api.v2.VersionServiceGetRequest.
 * Use `create(VersionServiceGetRequestSchema)` to create a new message.
 */
export declare const VersionServiceGetRequestSchema: GenMessage<VersionServiceGetRequest>;
/**
 * VersionServiceGetResponse is the response payload with the version
 *
 * @generated from message metalstack.api.v2.VersionServiceGetResponse
 */
export type VersionServiceGetResponse = Message<"metalstack.api.v2.VersionServiceGetResponse"> & {
    /**
     * Version of the application
     *
     * @generated from field: metalstack.api.v2.Version version = 1;
     */
    version?: Version;
};
/**
 * Describes the message metalstack.api.v2.VersionServiceGetResponse.
 * Use `create(VersionServiceGetResponseSchema)` to create a new message.
 */
export declare const VersionServiceGetResponseSchema: GenMessage<VersionServiceGetResponse>;
/**
 * VersionService serves version related functions
 *
 * @generated from service metalstack.api.v2.VersionService
 */
export declare const VersionService: GenService<{
    /**
     * Get the version
     *
     * @generated from rpc metalstack.api.v2.VersionService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof VersionServiceGetRequestSchema;
        output: typeof VersionServiceGetResponseSchema;
    };
}>;
