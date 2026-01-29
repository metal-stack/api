import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Machine, MachineQuery } from "../../api/v2/machine_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/machine.proto.
 */
export declare const file_metalstack_admin_v2_machine: GenFile;
/**
 * MachineServiceGetRequest is the request payload for a machine get request
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetRequest
 */
export type MachineServiceGetRequest = Message<"metalstack.admin.v2.MachineServiceGetRequest"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetRequest.
 * Use `create(MachineServiceGetRequestSchema)` to create a new message.
 */
export declare const MachineServiceGetRequestSchema: GenMessage<MachineServiceGetRequest>;
/**
 * MachineServiceGetResponse is the request payload for a machine get response
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetResponse
 */
export type MachineServiceGetResponse = Message<"metalstack.admin.v2.MachineServiceGetResponse"> & {
    /**
     * Machine is the machine requested
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetResponse.
 * Use `create(MachineServiceGetResponseSchema)` to create a new message.
 */
export declare const MachineServiceGetResponseSchema: GenMessage<MachineServiceGetResponse>;
/**
 * MachineServiceListRequest is the request payload for a machine list request
 *
 * @generated from message metalstack.admin.v2.MachineServiceListRequest
 */
export type MachineServiceListRequest = Message<"metalstack.admin.v2.MachineServiceListRequest"> & {
    /**
     * Query to list one ore more machines
     *
     * @generated from field: metalstack.api.v2.MachineQuery query = 1;
     */
    query?: MachineQuery;
    /**
     * Partition for which machines should be listed, could be left empty if only one partition is present
     * otherwise an error is thrown that the partition must be specified
     *
     * @generated from field: optional string partition = 2;
     */
    partition?: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListRequest.
 * Use `create(MachineServiceListRequestSchema)` to create a new message.
 */
export declare const MachineServiceListRequestSchema: GenMessage<MachineServiceListRequest>;
/**
 * MachineServiceListResponse is the request payload for a machine list response
 *
 * @generated from message metalstack.admin.v2.MachineServiceListResponse
 */
export type MachineServiceListResponse = Message<"metalstack.admin.v2.MachineServiceListResponse"> & {
    /**
     * Machines are the machines requested by a list request
     *
     * @generated from field: repeated metalstack.api.v2.Machine machines = 1;
     */
    machines: Machine[];
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListResponse.
 * Use `create(MachineServiceListResponseSchema)` to create a new message.
 */
export declare const MachineServiceListResponseSchema: GenMessage<MachineServiceListResponse>;
/**
 * MachineService serves machine related functions
 *
 * @generated from service metalstack.admin.v2.MachineService
 */
export declare const MachineService: GenService<{
    /**
     * Get a machine
     *
     * @generated from rpc metalstack.admin.v2.MachineService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof MachineServiceGetRequestSchema;
        output: typeof MachineServiceGetResponseSchema;
    };
    /**
     * List all machines
     *
     * @generated from rpc metalstack.admin.v2.MachineService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MachineServiceListRequestSchema;
        output: typeof MachineServiceListResponseSchema;
    };
}>;
