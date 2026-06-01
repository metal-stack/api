import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Project } from "./project_pb";
import type { Tenant } from "./tenant_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/user.proto.
 */
export declare const file_metalstack_api_v2_user: GenFile;
/**
 * User represents an end user of the platform.
 *
 * @generated from message metalstack.api.v2.User
 */
export type User = Message<"metalstack.api.v2.User"> & {
    /**
     * Login is the login at the provider
     *
     * @generated from field: string login = 1;
     */
    login: string;
    /**
     * Name of the user
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Email of the user
     *
     * @generated from field: string email = 3;
     */
    email: string;
    /**
     * AvatarUrl of the user
     *
     * @generated from field: string avatar_url = 4;
     */
    avatarUrl: string;
    /**
     * Tenants the user belongs to
     *
     * @generated from field: repeated metalstack.api.v2.Tenant tenants = 5;
     */
    tenants: Tenant[];
    /**
     * Projects the user belongs to
     *
     * @generated from field: repeated metalstack.api.v2.Project projects = 6;
     */
    projects: Project[];
    /**
     * DefaultTenant this user belongs to
     *
     * @generated from field: metalstack.api.v2.Tenant default_tenant = 7;
     */
    defaultTenant?: Tenant | undefined;
};
/**
 * Describes the message metalstack.api.v2.User.
 * Use `create(UserSchema)` to create a new message.
 */
export declare const UserSchema: GenMessage<User>;
/**
 * UserServiceGetRequest is the request payload for getting the user.
 *
 * @generated from message metalstack.api.v2.UserServiceGetRequest
 */
export type UserServiceGetRequest = Message<"metalstack.api.v2.UserServiceGetRequest"> & {};
/**
 * Describes the message metalstack.api.v2.UserServiceGetRequest.
 * Use `create(UserServiceGetRequestSchema)` to create a new message.
 */
export declare const UserServiceGetRequestSchema: GenMessage<UserServiceGetRequest>;
/**
 * UserServiceGetResponse is the response payload for getting the user.
 *
 * @generated from message metalstack.api.v2.UserServiceGetResponse
 */
export type UserServiceGetResponse = Message<"metalstack.api.v2.UserServiceGetResponse"> & {
    /**
     * User contains the authenticated user
     *
     * @generated from field: metalstack.api.v2.User user = 1;
     */
    user?: User | undefined;
};
/**
 * Describes the message metalstack.api.v2.UserServiceGetResponse.
 * Use `create(UserServiceGetResponseSchema)` to create a new message.
 */
export declare const UserServiceGetResponseSchema: GenMessage<UserServiceGetResponse>;
/**
 * UserService provides user information operations.
 *
 * @generated from service metalstack.api.v2.UserService
 */
export declare const UserService: GenService<{
    /**
     * Returns the authenticated user.
     *
     * @generated from rpc metalstack.api.v2.UserService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof UserServiceGetRequestSchema;
        output: typeof UserServiceGetResponseSchema;
    };
}>;
