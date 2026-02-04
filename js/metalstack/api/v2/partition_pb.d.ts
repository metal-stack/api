import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Labels, Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/partition.proto.
 */
export declare const file_metalstack_api_v2_partition: GenFile;
/**
 * Partition is a failure domain with machines and switches
 *
 * @generated from message metalstack.api.v2.Partition
 */
export type Partition = Message<"metalstack.api.v2.Partition"> & {
    /**
     * ID of this partition
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this ip
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
    /**
     * Description of this partition
     *
     * @generated from field: string description = 3;
     */
    description: string;
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
};
/**
 * Describes the message metalstack.api.v2.Partition.
 * Use `create(PartitionSchema)` to create a new message.
 */
export declare const PartitionSchema: GenMessage<Partition>;
/**
 * PartitionQuery is used to search partitions
 *
 * @generated from message metalstack.api.v2.PartitionQuery
 */
export type PartitionQuery = Message<"metalstack.api.v2.PartitionQuery"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
    /**
     * Labels lists only partitions containing the given labels
     *
     * @generated from field: optional metalstack.api.v2.Labels labels = 2;
     */
    labels?: Labels;
};
/**
 * Describes the message metalstack.api.v2.PartitionQuery.
 * Use `create(PartitionQuerySchema)` to create a new message.
 */
export declare const PartitionQuerySchema: GenMessage<PartitionQuery>;
/**
 * PartitionBootConfiguration defines how metal-hammer boots
 *
 * @generated from message metalstack.api.v2.PartitionBootConfiguration
 */
export type PartitionBootConfiguration = Message<"metalstack.api.v2.PartitionBootConfiguration"> & {
    /**
     * ImageURL the url to download the initrd for the boot image
     *
     * @generated from field: string image_url = 1;
     */
    imageUrl: string;
    /**
     * KernelURL the url to download the kernel for the boot image
     *
     * @generated from field: string kernel_url = 2;
     */
    kernelUrl: string;
    /**
     * Commandline the cmdline to the kernel for the boot image
     *
     * @generated from field: string commandline = 3;
     */
    commandline: string;
};
/**
 * Describes the message metalstack.api.v2.PartitionBootConfiguration.
 * Use `create(PartitionBootConfigurationSchema)` to create a new message.
 */
export declare const PartitionBootConfigurationSchema: GenMessage<PartitionBootConfiguration>;
/**
 * DNSServer
 *
 * @generated from message metalstack.api.v2.DNSServer
 */
export type DNSServer = Message<"metalstack.api.v2.DNSServer"> & {
    /**
     * IP address of this dns server
     *
     * @generated from field: string ip = 1;
     */
    ip: string;
};
/**
 * Describes the message metalstack.api.v2.DNSServer.
 * Use `create(DNSServerSchema)` to create a new message.
 */
export declare const DNSServerSchema: GenMessage<DNSServer>;
/**
 * NTPServer
 *
 * @generated from message metalstack.api.v2.NTPServer
 */
export type NTPServer = Message<"metalstack.api.v2.NTPServer"> & {
    /**
     * Address either as ip or hostname
     *
     * @generated from field: string address = 1;
     */
    address: string;
};
/**
 * Describes the message metalstack.api.v2.NTPServer.
 * Use `create(NTPServerSchema)` to create a new message.
 */
export declare const NTPServerSchema: GenMessage<NTPServer>;
/**
 * PartitionServiceGetRequest is the request payload for a partition get request
 *
 * @generated from message metalstack.api.v2.PartitionServiceGetRequest
 */
export type PartitionServiceGetRequest = Message<"metalstack.api.v2.PartitionServiceGetRequest"> & {
    /**
     * ID of the partition to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.api.v2.PartitionServiceGetRequest.
 * Use `create(PartitionServiceGetRequestSchema)` to create a new message.
 */
export declare const PartitionServiceGetRequestSchema: GenMessage<PartitionServiceGetRequest>;
/**
 * PartitionServiceListRequest is the request payload for a partition list request
 *
 * @generated from message metalstack.api.v2.PartitionServiceListRequest
 */
export type PartitionServiceListRequest = Message<"metalstack.api.v2.PartitionServiceListRequest"> & {
    /**
     * Query for partitions
     *
     * @generated from field: metalstack.api.v2.PartitionQuery query = 1;
     */
    query?: PartitionQuery;
};
/**
 * Describes the message metalstack.api.v2.PartitionServiceListRequest.
 * Use `create(PartitionServiceListRequestSchema)` to create a new message.
 */
export declare const PartitionServiceListRequestSchema: GenMessage<PartitionServiceListRequest>;
/**
 * PartitionServiceGetResponse is the response payload for a partition get request
 *
 * @generated from message metalstack.api.v2.PartitionServiceGetResponse
 */
export type PartitionServiceGetResponse = Message<"metalstack.api.v2.PartitionServiceGetResponse"> & {
    /**
     * Ip the partition
     *
     * @generated from field: metalstack.api.v2.Partition partition = 1;
     */
    partition?: Partition;
};
/**
 * Describes the message metalstack.api.v2.PartitionServiceGetResponse.
 * Use `create(PartitionServiceGetResponseSchema)` to create a new message.
 */
export declare const PartitionServiceGetResponseSchema: GenMessage<PartitionServiceGetResponse>;
/**
 * PartitionServiceListResponse is the response payload for a partition list request
 *
 * @generated from message metalstack.api.v2.PartitionServiceListResponse
 */
export type PartitionServiceListResponse = Message<"metalstack.api.v2.PartitionServiceListResponse"> & {
    /**
     * Ips the partitions
     *
     * @generated from field: repeated metalstack.api.v2.Partition partitions = 1;
     */
    partitions: Partition[];
};
/**
 * Describes the message metalstack.api.v2.PartitionServiceListResponse.
 * Use `create(PartitionServiceListResponseSchema)` to create a new message.
 */
export declare const PartitionServiceListResponseSchema: GenMessage<PartitionServiceListResponse>;
/**
 * PartitionService serves partition address related functions
 *
 * @generated from service metalstack.api.v2.PartitionService
 */
export declare const PartitionService: GenService<{
    /**
     * Get a partition
     *
     * @generated from rpc metalstack.api.v2.PartitionService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof PartitionServiceGetRequestSchema;
        output: typeof PartitionServiceGetResponseSchema;
    };
    /**
     * List all partitions
     *
     * @generated from rpc metalstack.api.v2.PartitionService.List
     */
    list: {
        methodKind: "unary";
        input: typeof PartitionServiceListRequestSchema;
        output: typeof PartitionServiceListResponseSchema;
    };
}>;
