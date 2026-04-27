import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/health.proto.
 */
export declare const file_metalstack_api_v2_health: GenFile;
/**
 * Health reports the health status of all services.
 *
 * @generated from message metalstack.api.v2.Health
 */
export type Health = Message<"metalstack.api.v2.Health"> & {
    /**
     * Services contains the health of all individual services
     *
     * @generated from field: repeated metalstack.api.v2.HealthStatus services = 1;
     */
    services: HealthStatus[];
};
/**
 * Describes the message metalstack.api.v2.Health.
 * Use `create(HealthSchema)` to create a new message.
 */
export declare const HealthSchema: GenMessage<Health>;
/**
 * HealthStatus represents the health status of one service.
 *
 * @generated from message metalstack.api.v2.HealthStatus
 */
export type HealthStatus = Message<"metalstack.api.v2.HealthStatus"> & {
    /**
     * Name is the name of the service
     *
     * @generated from field: metalstack.api.v2.Service name = 1;
     */
    name: Service;
    /**
     * Status is the status of this service
     *
     * @generated from field: metalstack.api.v2.ServiceStatus status = 2;
     */
    status: ServiceStatus;
    /**
     * Message describes the reason for the unhealthy status if possible
     *
     * @generated from field: string message = 3;
     */
    message: string;
    /**
     * Partitions describes the health of the service by partition
     *
     * @generated from field: map<string, metalstack.api.v2.PartitionHealth> partitions = 4;
     */
    partitions: {
        [key: string]: PartitionHealth;
    };
};
/**
 * Describes the message metalstack.api.v2.HealthStatus.
 * Use `create(HealthStatusSchema)` to create a new message.
 */
export declare const HealthStatusSchema: GenMessage<HealthStatus>;
/**
 * PartitionHealth represents the status of a specific service in this partition.
 *
 * @generated from message metalstack.api.v2.PartitionHealth
 */
export type PartitionHealth = Message<"metalstack.api.v2.PartitionHealth"> & {
    /**
     * Status is the health status of the service in this partition
     *
     * @generated from field: metalstack.api.v2.ServiceStatus status = 1;
     */
    status: ServiceStatus;
    /**
     * Message describes the reason for the unhealthy status if possible
     *
     * @generated from field: string message = 2;
     */
    message: string;
};
/**
 * Describes the message metalstack.api.v2.PartitionHealth.
 * Use `create(PartitionHealthSchema)` to create a new message.
 */
export declare const PartitionHealthSchema: GenMessage<PartitionHealth>;
/**
 * HealthServiceGetRequest is the request payload for getting the health of the system.
 *
 * @generated from message metalstack.api.v2.HealthServiceGetRequest
 */
export type HealthServiceGetRequest = Message<"metalstack.api.v2.HealthServiceGetRequest"> & {};
/**
 * Describes the message metalstack.api.v2.HealthServiceGetRequest.
 * Use `create(HealthServiceGetRequestSchema)` to create a new message.
 */
export declare const HealthServiceGetRequestSchema: GenMessage<HealthServiceGetRequest>;
/**
 * HealthServiceGetResponse is the response payload for getting the health of the system.
 *
 * @generated from message metalstack.api.v2.HealthServiceGetResponse
 */
export type HealthServiceGetResponse = Message<"metalstack.api.v2.HealthServiceGetResponse"> & {
    /**
     * Health contains the overall health of the system
     *
     * @generated from field: metalstack.api.v2.Health health = 1;
     */
    health?: Health | undefined;
};
/**
 * Describes the message metalstack.api.v2.HealthServiceGetResponse.
 * Use `create(HealthServiceGetResponseSchema)` to create a new message.
 */
export declare const HealthServiceGetResponseSchema: GenMessage<HealthServiceGetResponse>;
/**
 * ServiceStatus defines the status of a service.
 *
 * @generated from enum metalstack.api.v2.ServiceStatus
 */
export declare enum ServiceStatus {
    /**
     * SERVICE_STATUS_UNSPECIFIED service status is not known or unspecified
     *
     * @generated from enum value: SERVICE_STATUS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SERVICE_STATUS_DEGRADED the service is in degraded status, not the whole functionality is available
     *
     * @generated from enum value: SERVICE_STATUS_DEGRADED = 1;
     */
    DEGRADED = 1,
    /**
     * SERVICE_STATUS_UNHEALTHY the service is in unhealthy status, serious impact is expected
     *
     * @generated from enum value: SERVICE_STATUS_UNHEALTHY = 2;
     */
    UNHEALTHY = 2,
    /**
     * SERVICE_STATUS_HEALTHY the service is in healthy status, for example fully functional
     *
     * @generated from enum value: SERVICE_STATUS_HEALTHY = 3;
     */
    HEALTHY = 3
}
/**
 * Describes the enum metalstack.api.v2.ServiceStatus.
 */
export declare const ServiceStatusSchema: GenEnum<ServiceStatus>;
/**
 * Service defines the service for which the healthiness is reported.
 *
 * @generated from enum metalstack.api.v2.Service
 */
export declare enum Service {
    /**
     * SERVICE_UNSPECIFIED is a unknown service
     *
     * @generated from enum value: SERVICE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * SERVICE_IPAM the IPAM service
     *
     * @generated from enum value: SERVICE_IPAM = 1;
     */
    IPAM = 1,
    /**
     * SERVICE_RETHINK the RethinkDB
     *
     * @generated from enum value: SERVICE_RETHINK = 2;
     */
    RETHINK = 2,
    /**
     * SERVICE_MASTERDATA the masterdata API
     *
     * @generated from enum value: SERVICE_MASTERDATA = 3;
     */
    MASTERDATA = 3,
    /**
     * SERVICE_MACHINES the machine service
     *
     * @generated from enum value: SERVICE_MACHINES = 4;
     */
    MACHINES = 4,
    /**
     * SERVICE_AUDIT the auditing
     *
     * @generated from enum value: SERVICE_AUDIT = 5;
     */
    AUDIT = 5,
    /**
     * SERVICE_VPN the VPN service
     *
     * @generated from enum value: SERVICE_VPN = 6;
     */
    VPN = 6,
    /**
     * SERVICE_REDIS the Redis service
     *
     * @generated from enum value: SERVICE_REDIS = 7;
     */
    REDIS = 7,
    /**
     * SERVICE_TASKS the tasks service
     *
     * @generated from enum value: SERVICE_TASKS = 8;
     */
    TASKS = 8
}
/**
 * Describes the enum metalstack.api.v2.Service.
 */
export declare const ServiceSchema: GenEnum<Service>;
/**
 * HealthService provides health status monitoring operations.
 *
 * @generated from service metalstack.api.v2.HealthService
 */
export declare const HealthService: GenService<{
    /**
     * Returns the health status of the platform.
     *
     * @generated from rpc metalstack.api.v2.HealthService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof HealthServiceGetRequestSchema;
        output: typeof HealthServiceGetResponseSchema;
    };
}>;
