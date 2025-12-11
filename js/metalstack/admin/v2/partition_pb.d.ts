import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { UpdateLabels, UpdateMeta } from "../../api/v2/common_pb";
import type { DNSServer, NTPServer, Partition, PartitionBootConfiguration } from "../../api/v2/partition_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/partition.proto.
 */
export declare const file_metalstack_admin_v2_partition: GenFile;
/**
 * PartitionServiceCreateRequest is the request payload for a partition create request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceCreateRequest
 */
export type PartitionServiceCreateRequest = Message<"metalstack.admin.v2.PartitionServiceCreateRequest"> & {
    /**
     * Partition the partition
     *
     * @generated from field: metalstack.api.v2.Partition partition = 1;
     */
    partition?: Partition;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceCreateRequest.
 * Use `create(PartitionServiceCreateRequestSchema)` to create a new message.
 */
export declare const PartitionServiceCreateRequestSchema: GenMessage<PartitionServiceCreateRequest>;
/**
 * PartitionServiceUpdateRequest is the request payload for a partition update request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceUpdateRequest
 */
export type PartitionServiceUpdateRequest = Message<"metalstack.admin.v2.PartitionServiceUpdateRequest"> & {
    /**
     * ID of this partition
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Description of this partition
     *
     * @generated from field: optional string description = 3;
     */
    description?: string;
    /**
     * PartitionBootConfiguration defines how metal-hammer boots
     *
     * @generated from field: metalstack.api.v2.PartitionBootConfiguration boot_configuration = 4;
     */
    bootConfiguration?: PartitionBootConfiguration;
    /**
     * DNSServers for this partition
     *
     * @generated from field: repeated metalstack.api.v2.DNSServer dns_server = 5;
     */
    dnsServer: DNSServer[];
    /**
     * NTPServers for this partition
     *
     * @generated from field: repeated metalstack.api.v2.NTPServer ntp_server = 6;
     */
    ntpServer: NTPServer[];
    /**
     * ManagementServiceAddresses defines where the management is reachable
     * should be in the form <ip|host>:<port>
     *
     * @generated from field: repeated string mgmt_service_addresses = 7;
     */
    mgmtServiceAddresses: string[];
    /**
     * Labels to update on this network
     *
     * @generated from field: optional metalstack.api.v2.UpdateLabels labels = 8;
     */
    labels?: UpdateLabels;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceUpdateRequest.
 * Use `create(PartitionServiceUpdateRequestSchema)` to create a new message.
 */
export declare const PartitionServiceUpdateRequestSchema: GenMessage<PartitionServiceUpdateRequest>;
/**
 * PartitionServiceDeleteRequest is the request payload for a partition delete request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceDeleteRequest
 */
export type PartitionServiceDeleteRequest = Message<"metalstack.admin.v2.PartitionServiceDeleteRequest"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceDeleteRequest.
 * Use `create(PartitionServiceDeleteRequestSchema)` to create a new message.
 */
export declare const PartitionServiceDeleteRequestSchema: GenMessage<PartitionServiceDeleteRequest>;
/**
 * PartitionServiceCreateResponse is the response payload for a partition create request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceCreateResponse
 */
export type PartitionServiceCreateResponse = Message<"metalstack.admin.v2.PartitionServiceCreateResponse"> & {
    /**
     * Partition the partition
     *
     * @generated from field: metalstack.api.v2.Partition partition = 1;
     */
    partition?: Partition;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceCreateResponse.
 * Use `create(PartitionServiceCreateResponseSchema)` to create a new message.
 */
export declare const PartitionServiceCreateResponseSchema: GenMessage<PartitionServiceCreateResponse>;
/**
 * PartitionServiceUpdateResponse is the response payload for a partition update request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceUpdateResponse
 */
export type PartitionServiceUpdateResponse = Message<"metalstack.admin.v2.PartitionServiceUpdateResponse"> & {
    /**
     * Partition the partition
     *
     * @generated from field: metalstack.api.v2.Partition partition = 1;
     */
    partition?: Partition;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceUpdateResponse.
 * Use `create(PartitionServiceUpdateResponseSchema)` to create a new message.
 */
export declare const PartitionServiceUpdateResponseSchema: GenMessage<PartitionServiceUpdateResponse>;
/**
 * PartitionServiceCapacityResponse is the response payload for a partition delete request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceDeleteResponse
 */
export type PartitionServiceDeleteResponse = Message<"metalstack.admin.v2.PartitionServiceDeleteResponse"> & {
    /**
     * Partition the partition
     *
     * @generated from field: metalstack.api.v2.Partition partition = 1;
     */
    partition?: Partition;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceDeleteResponse.
 * Use `create(PartitionServiceDeleteResponseSchema)` to create a new message.
 */
export declare const PartitionServiceDeleteResponseSchema: GenMessage<PartitionServiceDeleteResponse>;
/**
 * PartitionServiceListRequest is the request payload for a partition capacity request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceCapacityRequest
 */
export type PartitionServiceCapacityRequest = Message<"metalstack.admin.v2.PartitionServiceCapacityRequest"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Size of machines to show the capacity
     *
     * @generated from field: optional string size = 2;
     */
    size?: string;
    /**
     * Project of machines to show the capacity
     *
     * @generated from field: optional string project = 3;
     */
    project?: string;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceCapacityRequest.
 * Use `create(PartitionServiceCapacityRequestSchema)` to create a new message.
 */
export declare const PartitionServiceCapacityRequestSchema: GenMessage<PartitionServiceCapacityRequest>;
/**
 * PartitionServiceCapacityResponse is the response payload for a partition capacity request
 *
 * @generated from message metalstack.admin.v2.PartitionServiceCapacityResponse
 */
export type PartitionServiceCapacityResponse = Message<"metalstack.admin.v2.PartitionServiceCapacityResponse"> & {
    /**
     * Size is the size id correlating to all counts in this server capacity.
     *
     * @generated from field: string size = 1;
     */
    size: string;
    /**
     * Total is the total amount of machines for this size.
     *
     * @generated from field: int64 total = 2;
     */
    total: bigint;
    /**
     * PhonedHome is the amount of machines that are currently in the provisioning state "phoned home".
     *
     * @generated from field: int64 phoned_home = 3;
     */
    phonedHome: bigint;
    /**
     * Waiting is the amount of machines that are currently in the provisioning state "waiting".
     *
     * @generated from field: int64 waiting = 4;
     */
    waiting: bigint;
    /**
     * Other is the amount of machines that are neither in the provisioning state waiting nor in phoned home but in another provisioning state.
     *
     * @generated from field: int64 other = 5;
     */
    other: bigint;
    /**
     * OtherMachines contains the machine IDs for machines that were classified into "Other".
     *
     * @generated from field: repeated string other_machines = 6;
     */
    otherMachines: string[];
    /**
     * Allocated is the amount of machines that are currently allocated.
     *
     * @generated from field: int64 allocated = 7;
     */
    allocated: bigint;
    /**
     * Allocatable is the amount of machines in a partition is the amount of machines that can be allocated.
     * Effectively this is the amount of waiting machines minus the machines that are unavailable due to machine state or un-allocatable. Size reservations are not considered in this count.
     *
     * @generated from field: int64 allocatable = 8;
     */
    allocatable: bigint;
    /**
     * Free is the amount of machines in a partition that can be freely allocated at any given moment by a project.
     * Effectively this is the amount of waiting machines minus the machines that are unavailable due to machine state or un-allocatable due to size reservations.
     *
     * @generated from field: int64 free = 9;
     */
    free: bigint;
    /**
     * Unavailable is the amount of machine in a partition that are currently not allocatable because they are not waiting or
     * not in the machine state "available", e.g. locked or reserved.
     *
     * @generated from field: int64 unavailable = 10;
     */
    unavailable: bigint;
    /**
     * Faulty is the amount of machines that are neither allocated nor in the pool of available machines because they report an error.
     *
     * @generated from field: int64 faulty = 11;
     */
    faulty: bigint;
    /**
     * FaultyMachines contains the machine IDs for machines that were classified into "Faulty".
     *
     * @generated from field: repeated string faulty_machines = 12;
     */
    faultyMachines: string[];
    /**
     * Reservations is the amount of reservations made for this size.
     *
     * @generated from field: int64 reservations = 13;
     */
    reservations: bigint;
    /**
     * UsedReservations is the amount of reservations already used up for this size.
     *
     * @generated from field: int64 used_reservations = 14;
     */
    usedReservations: bigint;
    /**
     * RemainingReservations is the amount of reservations remaining for this size.
     *
     * @generated from field: int64 remaining_reservations = 15;
     */
    remainingReservations: bigint;
};
/**
 * Describes the message metalstack.admin.v2.PartitionServiceCapacityResponse.
 * Use `create(PartitionServiceCapacityResponseSchema)` to create a new message.
 */
export declare const PartitionServiceCapacityResponseSchema: GenMessage<PartitionServiceCapacityResponse>;
/**
 * PartitionService serves partition address related functions
 *
 * @generated from service metalstack.admin.v2.PartitionService
 */
export declare const PartitionService: GenService<{
    /**
     * Create a partition
     *
     * @generated from rpc metalstack.admin.v2.PartitionService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof PartitionServiceCreateRequestSchema;
        output: typeof PartitionServiceCreateResponseSchema;
    };
    /**
     * Update a partition
     *
     * @generated from rpc metalstack.admin.v2.PartitionService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof PartitionServiceUpdateRequestSchema;
        output: typeof PartitionServiceUpdateResponseSchema;
    };
    /**
     * Delete a partition
     *
     * @generated from rpc metalstack.admin.v2.PartitionService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof PartitionServiceDeleteRequestSchema;
        output: typeof PartitionServiceDeleteResponseSchema;
    };
    /**
     * Capacity of a partitions
     *
     * @generated from rpc metalstack.admin.v2.PartitionService.Capacity
     */
    capacity: {
        methodKind: "unary";
        input: typeof PartitionServiceCapacityRequestSchema;
        output: typeof PartitionServiceCapacityResponseSchema;
    };
}>;
