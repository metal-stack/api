import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Machine, MachineBMCCommand, MachineBMCQuery, MachineBMCReport, MachineQuery } from "../../api/v2/machine_pb";
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
 * MachineServiceBMCCommandRequest is the request payload for a machine bmc command
 *
 * @generated from message metalstack.admin.v2.MachineServiceBMCCommandRequest
 */
export type MachineServiceBMCCommandRequest = Message<"metalstack.admin.v2.MachineServiceBMCCommandRequest"> & {
    /**
     * UUID of the machine to send the command to
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Command to send to the bmc of the machine
     *
     * @generated from field: metalstack.api.v2.MachineBMCCommand command = 2;
     */
    command: MachineBMCCommand;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceBMCCommandRequest.
 * Use `create(MachineServiceBMCCommandRequestSchema)` to create a new message.
 */
export declare const MachineServiceBMCCommandRequestSchema: GenMessage<MachineServiceBMCCommandRequest>;
/**
 * MachineServiceBMCCommandResponse is the response payload for a machine bmc command
 *
 * @generated from message metalstack.admin.v2.MachineServiceBMCCommandResponse
 */
export type MachineServiceBMCCommandResponse = Message<"metalstack.admin.v2.MachineServiceBMCCommandResponse"> & {};
/**
 * Describes the message metalstack.admin.v2.MachineServiceBMCCommandResponse.
 * Use `create(MachineServiceBMCCommandResponseSchema)` to create a new message.
 */
export declare const MachineServiceBMCCommandResponseSchema: GenMessage<MachineServiceBMCCommandResponse>;
/**
 * MachineServiceGetBMCRequest is the request payload for a machine getbmc request
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetBMCRequest
 */
export type MachineServiceGetBMCRequest = Message<"metalstack.admin.v2.MachineServiceGetBMCRequest"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetBMCRequest.
 * Use `create(MachineServiceGetBMCRequestSchema)` to create a new message.
 */
export declare const MachineServiceGetBMCRequestSchema: GenMessage<MachineServiceGetBMCRequest>;
/**
 * MachineServiceGetBMCResponse is the response payload for a machine getbmc request
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetBMCResponse
 */
export type MachineServiceGetBMCResponse = Message<"metalstack.admin.v2.MachineServiceGetBMCResponse"> & {
    /**
     * UUID of the machine
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * BMC contains the BMC details of this machine
     *
     * @generated from field: metalstack.api.v2.MachineBMCReport bmc = 2;
     */
    bmc?: MachineBMCReport;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetBMCResponse.
 * Use `create(MachineServiceGetBMCResponseSchema)` to create a new message.
 */
export declare const MachineServiceGetBMCResponseSchema: GenMessage<MachineServiceGetBMCResponse>;
/**
 * MachineServiceListBMCRequest is the request payload for a machine listbmc request
 *
 * @generated from message metalstack.admin.v2.MachineServiceListBMCRequest
 */
export type MachineServiceListBMCRequest = Message<"metalstack.admin.v2.MachineServiceListBMCRequest"> & {
    /**
     * Query to list one ore more bmcs of more machines
     *
     * @generated from field: metalstack.api.v2.MachineBMCQuery query = 1;
     */
    query?: MachineBMCQuery;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListBMCRequest.
 * Use `create(MachineServiceListBMCRequestSchema)` to create a new message.
 */
export declare const MachineServiceListBMCRequestSchema: GenMessage<MachineServiceListBMCRequest>;
/**
 * MachineServiceListBMCResponse is the response payload for a machine listbmc request
 *
 * @generated from message metalstack.admin.v2.MachineServiceListBMCResponse
 */
export type MachineServiceListBMCResponse = Message<"metalstack.admin.v2.MachineServiceListBMCResponse"> & {
    /**
     * BMCReports maps the bmc report per machine uuid
     *
     * @generated from field: map<string, metalstack.api.v2.MachineBMCReport> bmc_reports = 1;
     */
    bmcReports: {
        [key: string]: MachineBMCReport;
    };
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListBMCResponse.
 * Use `create(MachineServiceListBMCResponseSchema)` to create a new message.
 */
export declare const MachineServiceListBMCResponseSchema: GenMessage<MachineServiceListBMCResponse>;
/**
 * MachineServiceConsolePasswordRequest is the request to get the console password
 *
 * @generated from message metalstack.admin.v2.MachineServiceConsolePasswordRequest
 */
export type MachineServiceConsolePasswordRequest = Message<"metalstack.admin.v2.MachineServiceConsolePasswordRequest"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Reason must be provided why access to the console is requested.
     * Reason is only forwarded to an audit sink
     *
     * @generated from field: string reason = 2;
     */
    reason: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceConsolePasswordRequest.
 * Use `create(MachineServiceConsolePasswordRequestSchema)` to create a new message.
 */
export declare const MachineServiceConsolePasswordRequestSchema: GenMessage<MachineServiceConsolePasswordRequest>;
/**
 * MachineServiceConsolePasswordResponse is the response to the console password request
 *
 * @generated from message metalstack.admin.v2.MachineServiceConsolePasswordResponse
 */
export type MachineServiceConsolePasswordResponse = Message<"metalstack.admin.v2.MachineServiceConsolePasswordResponse"> & {
    /**
     * UUID of the machine to get
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Password to access the console
     *
     * @generated from field: string password = 2;
     */
    password: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceConsolePasswordResponse.
 * Use `create(MachineServiceConsolePasswordResponseSchema)` to create a new message.
 */
export declare const MachineServiceConsolePasswordResponseSchema: GenMessage<MachineServiceConsolePasswordResponse>;
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
    /**
     * BMCCommand send a command to the bmc of a machine
     *
     * @generated from rpc metalstack.admin.v2.MachineService.BMCCommand
     */
    bMCCommand: {
        methodKind: "unary";
        input: typeof MachineServiceBMCCommandRequestSchema;
        output: typeof MachineServiceBMCCommandResponseSchema;
    };
    /**
     * GetBMC returns the BMC details of a machine
     *
     * @generated from rpc metalstack.admin.v2.MachineService.GetBMC
     */
    getBMC: {
        methodKind: "unary";
        input: typeof MachineServiceGetBMCRequestSchema;
        output: typeof MachineServiceGetBMCResponseSchema;
    };
    /**
     * ListBMC returns the BMC details of many machines
     *
     * @generated from rpc metalstack.admin.v2.MachineService.ListBMC
     */
    listBMC: {
        methodKind: "unary";
        input: typeof MachineServiceListBMCRequestSchema;
        output: typeof MachineServiceListBMCResponseSchema;
    };
    /**
     * ConsolePassword returns the password to access the serial console of the machine
     *
     * @generated from rpc metalstack.admin.v2.MachineService.ConsolePassword
     */
    consolePassword: {
        methodKind: "unary";
        input: typeof MachineServiceConsolePasswordRequestSchema;
        output: typeof MachineServiceConsolePasswordResponseSchema;
    };
}>;
