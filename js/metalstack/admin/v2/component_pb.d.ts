import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Token } from "../../api/v2/token_pb";
import type { Version } from "../../api/v2/version_pb";
import type { ComponentType } from "../../infra/v2/component_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/component.proto.
 */
export declare const file_metalstack_admin_v2_component: GenFile;
/**
 * ComponentServiceListRequest
 *
 * @generated from message metalstack.admin.v2.ComponentServiceListRequest
 */
export type ComponentServiceListRequest = Message<"metalstack.admin.v2.ComponentServiceListRequest"> & {};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceListRequest.
 * Use `create(ComponentServiceListRequestSchema)` to create a new message.
 */
export declare const ComponentServiceListRequestSchema: GenMessage<ComponentServiceListRequest>;
/**
 * ComponentServiceListResponse
 *
 * @generated from message metalstack.admin.v2.ComponentServiceListResponse
 */
export type ComponentServiceListResponse = Message<"metalstack.admin.v2.ComponentServiceListResponse"> & {
    /**
     * Components
     *
     * @generated from field: repeated metalstack.admin.v2.Component components = 1;
     */
    components: Component[];
};
/**
 * Describes the message metalstack.admin.v2.ComponentServiceListResponse.
 * Use `create(ComponentServiceListResponseSchema)` to create a new message.
 */
export declare const ComponentServiceListResponseSchema: GenMessage<ComponentServiceListResponse>;
/**
 * Component represents a microservice connected to our apiserver
 *
 * @generated from message metalstack.admin.v2.Component
 */
export type Component = Message<"metalstack.admin.v2.Component"> & {
    /**
     * UUID identifies this component event
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Type defines which service is actually pinging
     *
     * @generated from field: metalstack.infra.v2.ComponentType type = 2;
     */
    type: ComponentType;
    /**
     * Identifier is a unique identifier of this service, e.g. if two instance are running, this might be the pod id.
     * micro_service and identifier guarantee uniqueness.
     *
     * @generated from field: string identifier = 3;
     */
    identifier: string;
    /**
     * StartedAt is the timestamp this service was started
     *
     * @generated from field: google.protobuf.Timestamp started_at = 4;
     */
    startedAt?: Timestamp;
    /**
     * Version of this service
     *
     * @generated from field: metalstack.api.v2.Version version = 5;
     */
    version?: Version;
    /**
     * Token is the token which is actually used by this microservice.
     *
     * @generated from field: metalstack.api.v2.Token token = 6;
     */
    token?: Token;
};
/**
 * Describes the message metalstack.admin.v2.Component.
 * Use `create(ComponentSchema)` to create a new message.
 */
export declare const ComponentSchema: GenMessage<Component>;
/**
 * ComponentService serves microservice related functions
 *
 * @generated from service metalstack.admin.v2.ComponentService
 */
export declare const ComponentService: GenService<{
    /**
     * List all components with their status
     *
     * @generated from rpc metalstack.admin.v2.ComponentService.List
     */
    list: {
        methodKind: "unary";
        input: typeof ComponentServiceListRequestSchema;
        output: typeof ComponentServiceListResponseSchema;
    };
}>;
