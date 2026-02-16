import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Version } from "../../api/v2/version_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/component.proto.
 */
export declare const file_metalstack_infra_v2_component: GenFile;
/**
 * ComponentServicePingRequest is sent from a microservice to report its state regularly
 *
 * @generated from message metalstack.infra.v2.ComponentServicePingRequest
 */
export type ComponentServicePingRequest = Message<"metalstack.infra.v2.ComponentServicePingRequest"> & {
    /**
     * Type defines which service is actually pinging
     *
     * @generated from field: metalstack.infra.v2.ComponentType type = 1;
     */
    type: ComponentType;
    /**
     * Identifier is a unique identifier of this service, e.g. if two instance are running, this might be the pod id.
     * micro_service and identifier guarantee uniqueness.
     *
     * @generated from field: string identifier = 2;
     */
    identifier: string;
    /**
     * StartedAt is the timestamp this service was started
     *
     * @generated from field: google.protobuf.Timestamp started_at = 3;
     */
    startedAt?: Timestamp;
    /**
     * Version of this service
     *
     * @generated from field: metalstack.api.v2.Version version = 4;
     */
    version?: Version;
};
/**
 * Describes the message metalstack.infra.v2.ComponentServicePingRequest.
 * Use `create(ComponentServicePingRequestSchema)` to create a new message.
 */
export declare const ComponentServicePingRequestSchema: GenMessage<ComponentServicePingRequest>;
/**
 * ComponentServicePingResponse is the response to a ping request
 *
 * @generated from message metalstack.infra.v2.ComponentServicePingResponse
 */
export type ComponentServicePingResponse = Message<"metalstack.infra.v2.ComponentServicePingResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.ComponentServicePingResponse.
 * Use `create(ComponentServicePingResponseSchema)` to create a new message.
 */
export declare const ComponentServicePingResponseSchema: GenMessage<ComponentServicePingResponse>;
/**
 * ComponentType defines which service is actually pinging
 *
 * @generated from enum metalstack.infra.v2.ComponentType
 */
export declare enum ComponentType {
    /**
     * COMPONENT_TYPE_UNSPECIFIED is unspecified
     *
     * @generated from enum value: COMPONENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * COMPONENT_TYPE_PIXIECORE is pixiecore
     *
     * @generated from enum value: COMPONENT_TYPE_PIXIECORE = 1;
     */
    PIXIECORE = 1,
    /**
     * COMPONENT_TYPE_METAL_CORE is metal-core
     *
     * @generated from enum value: COMPONENT_TYPE_METAL_CORE = 2;
     */
    METAL_CORE = 2,
    /**
     * COMPONENT_TYPE_METAL_BMC is metal-bmc
     *
     * @generated from enum value: COMPONENT_TYPE_METAL_BMC = 3;
     */
    METAL_BMC = 3,
    /**
     * COMPONENT_TYPE_METAL_IMAGE_CACHE_SYNC is metal-image-cache-sync
     *
     * @generated from enum value: COMPONENT_TYPE_METAL_IMAGE_CACHE_SYNC = 4;
     */
    METAL_IMAGE_CACHE_SYNC = 4,
    /**
     * COMPONENT_TYPE_METAL_CONSOLE is metal-console
     *
     * @generated from enum value: COMPONENT_TYPE_METAL_CONSOLE = 5;
     */
    METAL_CONSOLE = 5,
    /**
     * COMPONENT_TYPE_METAL_METRICS_EXPORTER is metal-metrics-exporter
     *
     * TODO what about gepm, ccm, etc. I would allow them to call this service but would not introduce
     * enums for all of them as they are not strictly metal-stack components.
     *
     * @generated from enum value: COMPONENT_TYPE_METAL_METRICS_EXPORTER = 6;
     */
    METAL_METRICS_EXPORTER = 6
}
/**
 * Describes the enum metalstack.infra.v2.ComponentType.
 */
export declare const ComponentTypeSchema: GenEnum<ComponentType>;
/**
 * ComponentService serves component, e.g. microservices related functions.
 *
 * @generated from service metalstack.infra.v2.ComponentService
 */
export declare const ComponentService: GenService<{
    /**
     * Ping must be called from every connected microservice in a recurring manner
     * to get visibility of all required microservices.
     *
     * @generated from rpc metalstack.infra.v2.ComponentService.Ping
     */
    ping: {
        methodKind: "unary";
        input: typeof ComponentServicePingRequestSchema;
        output: typeof ComponentServicePingResponseSchema;
    };
}>;
