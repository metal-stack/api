import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration, Timestamp } from "@bufbuild/protobuf/wkt";
import type { ComponentType } from "../../api/v2/component_pb";
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
     * @generated from field: metalstack.api.v2.ComponentType type = 1;
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
     * StartedAt is the timestamp this service was started.
     *
     * @generated from field: google.protobuf.Timestamp started_at = 3;
     */
    startedAt?: Timestamp;
    /**
     * Interval at which the ping is scheduled, must be between 5 seconds and 1 hour.
     * Also gets validated in the same way in go/client/ping.go.
     *
     * @generated from field: google.protobuf.Duration interval = 4;
     */
    interval?: Duration;
    /**
     * Version of this service
     *
     * @generated from field: metalstack.api.v2.Version version = 5;
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
 * ComponentService serves component, e.g. microservices related functions.
 *
 * @generated from service metalstack.infra.v2.ComponentService
 */
export declare const ComponentService: GenService<{
    /**
     * Ping must be called from every connected microservice in a recurring manner
     * to get visibility of all registered microservices.
     *
     * @generated from rpc metalstack.infra.v2.ComponentService.Ping
     */
    ping: {
        methodKind: "unary";
        input: typeof ComponentServicePingRequestSchema;
        output: typeof ComponentServicePingResponseSchema;
    };
}>;
