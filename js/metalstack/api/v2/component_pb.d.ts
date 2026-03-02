import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { Token } from "./token_pb";
import type { Version } from "./version_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/component.proto.
 */
export declare const file_metalstack_api_v2_component: GenFile;
/**
 * Component represents a microservice connected to our apiserver
 *
 * @generated from message metalstack.api.v2.Component
 */
export type Component = Message<"metalstack.api.v2.Component"> & {
    /**
     * UUID identifies this component event
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Type defines which service is actually pinging
     *
     * @generated from field: metalstack.api.v2.ComponentType type = 2;
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
     * ReportedAt is the timestamp this service sent this ping.
     *
     * @generated from field: google.protobuf.Timestamp reported_at = 5;
     */
    reportedAt?: Timestamp;
    /**
     * Interval at which the ping is scheduled.
     *
     * @generated from field: google.protobuf.Duration interval = 6;
     */
    interval?: Duration;
    /**
     * Version of this service
     *
     * @generated from field: metalstack.api.v2.Version version = 7;
     */
    version?: Version;
    /**
     * Token is the token which is actually used by this microservice.
     *
     * @generated from field: metalstack.api.v2.Token token = 8;
     */
    token?: Token;
};
/**
 * Describes the message metalstack.api.v2.Component.
 * Use `create(ComponentSchema)` to create a new message.
 */
export declare const ComponentSchema: GenMessage<Component>;
/**
 * ComponentQuery to query components
 *
 * @generated from message metalstack.api.v2.ComponentQuery
 */
export type ComponentQuery = Message<"metalstack.api.v2.ComponentQuery"> & {
    /**
     * UUID identifies this component
     *
     * @generated from field: optional string uuid = 1;
     */
    uuid?: string;
    /**
     * Type defines which service is actually pinging
     *
     * @generated from field: optional metalstack.api.v2.ComponentType type = 2;
     */
    type?: ComponentType;
    /**
     * Identifier is a unique identifier of this service, e.g. if two instance are running, this might be the pod id.
     * micro_service and identifier guarantee uniqueness.
     *
     * @generated from field: optional string identifier = 3;
     */
    identifier?: string;
};
/**
 * Describes the message metalstack.api.v2.ComponentQuery.
 * Use `create(ComponentQuerySchema)` to create a new message.
 */
export declare const ComponentQuerySchema: GenMessage<ComponentQuery>;
/**
 * ComponentType defines which service is actually pinging
 *
 * @generated from enum metalstack.api.v2.ComponentType
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
     * @generated from enum value: COMPONENT_TYPE_METAL_METRICS_EXPORTER = 6;
     */
    METAL_METRICS_EXPORTER = 6
}
/**
 * Describes the enum metalstack.api.v2.ComponentType.
 */
export declare const ComponentTypeSchema: GenEnum<ComponentType>;
