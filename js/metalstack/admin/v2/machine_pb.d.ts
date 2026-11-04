import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Machine, MachineBMCCommand, MachineBMCDetails, MachineIssues, MachineIssuesQuery, MachineQuery, MachineState } from "../../api/v2/machine_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/machine.proto.
 */
export declare const file_metalstack_admin_v2_machine: GenFile;
/**
 * MachineServiceGetRequest is the request payload for a machine get request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetRequest
 */
export type MachineServiceGetRequest = Message<"metalstack.admin.v2.MachineServiceGetRequest"> & {
    /**
     * UUID of the machine to get.
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
 * MachineServiceGetResponse is the response payload for a machine get request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetResponse
 */
export type MachineServiceGetResponse = Message<"metalstack.admin.v2.MachineServiceGetResponse"> & {
    /**
     * Machine is the machine requested.
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetResponse.
 * Use `create(MachineServiceGetResponseSchema)` to create a new message.
 */
export declare const MachineServiceGetResponseSchema: GenMessage<MachineServiceGetResponse>;
/**
 * MachineServiceListRequest is the request payload for a machine list request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceListRequest
 */
export type MachineServiceListRequest = Message<"metalstack.admin.v2.MachineServiceListRequest"> & {
    /**
     * Query to list one or more machines.
     *
     * @generated from field: metalstack.api.v2.MachineQuery query = 1;
     */
    query?: MachineQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListRequest.
 * Use `create(MachineServiceListRequestSchema)` to create a new message.
 */
export declare const MachineServiceListRequestSchema: GenMessage<MachineServiceListRequest>;
/**
 * MachineServiceListResponse is the response payload for a machine list request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceListResponse
 */
export type MachineServiceListResponse = Message<"metalstack.admin.v2.MachineServiceListResponse"> & {
    /**
     * Machines are the machines requested by a list request.
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
 * MachineServiceDeleteRequest is the request payload for a machine delete request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceDeleteRequest
 */
export type MachineServiceDeleteRequest = Message<"metalstack.admin.v2.MachineServiceDeleteRequest"> & {
    /**
     * UUID of the machine to permanently delete.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceDeleteRequest.
 * Use `create(MachineServiceDeleteRequestSchema)` to create a new message.
 */
export declare const MachineServiceDeleteRequestSchema: GenMessage<MachineServiceDeleteRequest>;
/**
 * MachineServiceDeleteResponse is the response payload for a machine delete request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceDeleteResponse
 */
export type MachineServiceDeleteResponse = Message<"metalstack.admin.v2.MachineServiceDeleteResponse"> & {
    /**
     * Machine is the machine permanently deleted.
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceDeleteResponse.
 * Use `create(MachineServiceDeleteResponseSchema)` to create a new message.
 */
export declare const MachineServiceDeleteResponseSchema: GenMessage<MachineServiceDeleteResponse>;
/**
 * MachineServiceBMCCommandRequest is the request payload for a machine bmc command.
 *
 * @generated from message metalstack.admin.v2.MachineServiceBMCCommandRequest
 */
export type MachineServiceBMCCommandRequest = Message<"metalstack.admin.v2.MachineServiceBMCCommandRequest"> & {
    /**
     * UUID of the machine to send the command to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Command to send to the bmc of the machine.
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
 * MachineServiceBMCCommandResponse is the response payload for a machine bmc command.
 *
 * @generated from message metalstack.admin.v2.MachineServiceBMCCommandResponse
 */
export type MachineServiceBMCCommandResponse = Message<"metalstack.admin.v2.MachineServiceBMCCommandResponse"> & {
    /**
     * TaskId is the identifier of the task that managed the bmc command execution.
     *
     * @generated from field: string task_id = 1;
     */
    taskId: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceBMCCommandResponse.
 * Use `create(MachineServiceBMCCommandResponseSchema)` to create a new message.
 */
export declare const MachineServiceBMCCommandResponseSchema: GenMessage<MachineServiceBMCCommandResponse>;
/**
 * MachineServiceGetBMCRequest is the request payload for a machine getbmc request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetBMCRequest
 */
export type MachineServiceGetBMCRequest = Message<"metalstack.admin.v2.MachineServiceGetBMCRequest"> & {
    /**
     * UUID of the machine to get.
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
 * MachineServiceGetBMCResponse is the response payload for a machine getbmc request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceGetBMCResponse
 */
export type MachineServiceGetBMCResponse = Message<"metalstack.admin.v2.MachineServiceGetBMCResponse"> & {
    /**
     * BMCDetails of the machine.
     *
     * @generated from field: metalstack.api.v2.MachineBMCDetails bmc_details = 1;
     */
    bmcDetails?: MachineBMCDetails | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceGetBMCResponse.
 * Use `create(MachineServiceGetBMCResponseSchema)` to create a new message.
 */
export declare const MachineServiceGetBMCResponseSchema: GenMessage<MachineServiceGetBMCResponse>;
/**
 * MachineServiceListBMCRequest is the request payload for a machine listbmc request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceListBMCRequest
 */
export type MachineServiceListBMCRequest = Message<"metalstack.admin.v2.MachineServiceListBMCRequest"> & {
    /**
     * Query to list one or more BMCs of many machines.
     *
     * @generated from field: metalstack.api.v2.MachineQuery query = 1;
     */
    query?: MachineQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListBMCRequest.
 * Use `create(MachineServiceListBMCRequestSchema)` to create a new message.
 */
export declare const MachineServiceListBMCRequestSchema: GenMessage<MachineServiceListBMCRequest>;
/**
 * MachineServiceListBMCResponse is the response payload for a machine listbmc request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceListBMCResponse
 */
export type MachineServiceListBMCResponse = Message<"metalstack.admin.v2.MachineServiceListBMCResponse"> & {
    /**
     * BMCDetails of the machines.
     *
     * @generated from field: repeated metalstack.api.v2.MachineBMCDetails bmc_details = 1;
     */
    bmcDetails: MachineBMCDetails[];
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceListBMCResponse.
 * Use `create(MachineServiceListBMCResponseSchema)` to create a new message.
 */
export declare const MachineServiceListBMCResponseSchema: GenMessage<MachineServiceListBMCResponse>;
/**
 * MachineServiceConsolePasswordRequest is the request payload for getting the console password.
 *
 * @generated from message metalstack.admin.v2.MachineServiceConsolePasswordRequest
 */
export type MachineServiceConsolePasswordRequest = Message<"metalstack.admin.v2.MachineServiceConsolePasswordRequest"> & {
    /**
     * UUID of the machine to get.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Reason must be provided why access to the console is requested.
     * Reason is only forwarded to an audit sink.
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
 * MachineServiceConsolePasswordResponse is the response payload for getting the console password.
 *
 * @generated from message metalstack.admin.v2.MachineServiceConsolePasswordResponse
 */
export type MachineServiceConsolePasswordResponse = Message<"metalstack.admin.v2.MachineServiceConsolePasswordResponse"> & {
    /**
     * UUID of the machine to get.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Password to access the console.
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
 * MachineServiceSetStateRequest is the request payload for a machine set state request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceSetStateRequest
 */
export type MachineServiceSetStateRequest = Message<"metalstack.admin.v2.MachineServiceSetStateRequest"> & {
    /**
     * UUID of the machine to set the state to.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * State the state of this machine.
     * If State set to available, description is also cleared.
     *
     * @generated from field: metalstack.api.v2.MachineState state = 2;
     */
    state: MachineState;
    /**
     * Description a description why this machine is in the given state.
     *
     * @generated from field: string description = 3;
     */
    description: string;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceSetStateRequest.
 * Use `create(MachineServiceSetStateRequestSchema)` to create a new message.
 */
export declare const MachineServiceSetStateRequestSchema: GenMessage<MachineServiceSetStateRequest>;
/**
 * MachineServiceSetStateResponse is the response payload for a machine set state request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceSetStateResponse
 */
export type MachineServiceSetStateResponse = Message<"metalstack.admin.v2.MachineServiceSetStateResponse"> & {
    /**
     * Machine with the changed state.
     *
     * @generated from field: metalstack.api.v2.Machine machine = 1;
     */
    machine?: Machine | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceSetStateResponse.
 * Use `create(MachineServiceSetStateResponseSchema)` to create a new message.
 */
export declare const MachineServiceSetStateResponseSchema: GenMessage<MachineServiceSetStateResponse>;
/**
 * MachineServiceIssuesRequest is the request payload for a machine issues request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceIssuesRequest
 */
export type MachineServiceIssuesRequest = Message<"metalstack.admin.v2.MachineServiceIssuesRequest"> & {
    /**
     * Query which machine issues to include.
     *
     * @generated from field: metalstack.api.v2.MachineIssuesQuery query = 1;
     */
    query?: MachineIssuesQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceIssuesRequest.
 * Use `create(MachineServiceIssuesRequestSchema)` to create a new message.
 */
export declare const MachineServiceIssuesRequestSchema: GenMessage<MachineServiceIssuesRequest>;
/**
 * MachineServiceIssuesResponse is the response payload for a machine issues request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceIssuesResponse
 */
export type MachineServiceIssuesResponse = Message<"metalstack.admin.v2.MachineServiceIssuesResponse"> & {
    /**
     * Issues is the result of the issues query.
     *
     * @generated from field: repeated metalstack.api.v2.MachineIssues issues = 1;
     */
    issues: MachineIssues[];
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceIssuesResponse.
 * Use `create(MachineServiceIssuesResponseSchema)` to create a new message.
 */
export declare const MachineServiceIssuesResponseSchema: GenMessage<MachineServiceIssuesResponse>;
/**
 * MachineServiceHousekeepingRequest is the request payload for a machine housekeeping request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceHousekeepingRequest
 */
export type MachineServiceHousekeepingRequest = Message<"metalstack.admin.v2.MachineServiceHousekeepingRequest"> & {
    /**
     * Housekeeping Action to perform.
     *
     * @generated from field: metalstack.admin.v2.MachineHousekeepingAction action = 1;
     */
    action: MachineHousekeepingAction;
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceHousekeepingRequest.
 * Use `create(MachineServiceHousekeepingRequestSchema)` to create a new message.
 */
export declare const MachineServiceHousekeepingRequestSchema: GenMessage<MachineServiceHousekeepingRequest>;
/**
 * MachineServiceHousekeepingResponse is the response payload for a machine housekeeping request.
 *
 * @generated from message metalstack.admin.v2.MachineServiceHousekeepingResponse
 */
export type MachineServiceHousekeepingResponse = Message<"metalstack.admin.v2.MachineServiceHousekeepingResponse"> & {
    /**
     * Housekeeping Action which was performed.
     *
     * @generated from field: metalstack.admin.v2.MachineHousekeepingAction action = 1;
     */
    action: MachineHousekeepingAction;
    /**
     * Modified Machine IDs modified during this housekeeping action.
     *
     * @generated from field: repeated string modified_machine_ids = 2;
     */
    modifiedMachineIds: string[];
};
/**
 * Describes the message metalstack.admin.v2.MachineServiceHousekeepingResponse.
 * Use `create(MachineServiceHousekeepingResponseSchema)` to create a new message.
 */
export declare const MachineServiceHousekeepingResponseSchema: GenMessage<MachineServiceHousekeepingResponse>;
/**
 * HousekeepingAction defines which housekeeping action should be triggered.
 *
 * @generated from enum metalstack.admin.v2.MachineHousekeepingAction
 */
export declare enum MachineHousekeepingAction {
    /**
     * MACHINE_HOUSEKEEPING_ACTION_UNSPECIFIED is not defined.
     *
     * @generated from enum value: MACHINE_HOUSEKEEPING_ACTION_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * MACHINE_HOUSEKEEPING_ACTION_VPN_CONNECTED updates machines which have vpn configured if they are connected or not.
     *
     * @generated from enum value: MACHINE_HOUSEKEEPING_ACTION_VPN_CONNECTED = 1;
     */
    VPN_CONNECTED = 1,
    /**
     * MACHINE_HOUSEKEEPING_ACTION_LIVELINESS reports the liveliness of machines.
     *
     * @generated from enum value: MACHINE_HOUSEKEEPING_ACTION_LIVELINESS = 2;
     */
    LIVELINESS = 2,
    /**
     * MACHINE_HOUSEKEEPING_ACTION_RESURRECT resurrect dead machines.
     *
     * @generated from enum value: MACHINE_HOUSEKEEPING_ACTION_RESURRECT = 3;
     */
    RESURRECT = 3
}
/**
 * Describes the enum metalstack.admin.v2.MachineHousekeepingAction.
 */
export declare const MachineHousekeepingActionSchema: GenEnum<MachineHousekeepingAction>;
/**
 * MachineService provides machine lifecycle management operations.
 *
 * @generated from service metalstack.admin.v2.MachineService
 */
export declare const MachineService: GenService<{
    /**
     * Returns the machine with the specified UUID.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof MachineServiceGetRequestSchema;
        output: typeof MachineServiceGetResponseSchema;
    };
    /**
     * Returns the list of all machines.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.List
     */
    list: {
        methodKind: "unary";
        input: typeof MachineServiceListRequestSchema;
        output: typeof MachineServiceListResponseSchema;
    };
    /**
     * Delete a machine from the database. This can only be done if the machine is offline and dead.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof MachineServiceDeleteRequestSchema;
        output: typeof MachineServiceDeleteResponseSchema;
    };
    /**
     * BMCCommand sends a command to the BMC of a machine.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.BMCCommand
     */
    bMCCommand: {
        methodKind: "unary";
        input: typeof MachineServiceBMCCommandRequestSchema;
        output: typeof MachineServiceBMCCommandResponseSchema;
    };
    /**
     * Returns the BMC details of a machine.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.GetBMC
     */
    getBMC: {
        methodKind: "unary";
        input: typeof MachineServiceGetBMCRequestSchema;
        output: typeof MachineServiceGetBMCResponseSchema;
    };
    /**
     * ListBMC returns the BMC details of many machines.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.ListBMC
     */
    listBMC: {
        methodKind: "unary";
        input: typeof MachineServiceListBMCRequestSchema;
        output: typeof MachineServiceListBMCResponseSchema;
    };
    /**
     * GetConsolePassword returns the password to access the serial console of the machine.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.ConsolePassword
     */
    consolePassword: {
        methodKind: "unary";
        input: typeof MachineServiceConsolePasswordRequestSchema;
        output: typeof MachineServiceConsolePasswordResponseSchema;
    };
    /**
     * SetState set the state of a machine.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.SetState
     */
    setState: {
        methodKind: "unary";
        input: typeof MachineServiceSetStateRequestSchema;
        output: typeof MachineServiceSetStateResponseSchema;
    };
    /**
     * Issues allows to query issues of machines.
     *
     * @generated from rpc metalstack.admin.v2.MachineService.Issues
     */
    issues: {
        methodKind: "unary";
        input: typeof MachineServiceIssuesRequestSchema;
        output: typeof MachineServiceIssuesResponseSchema;
    };
    /**
     * Housekeeping
     *
     * @generated from rpc metalstack.admin.v2.MachineService.Housekeeping
     */
    housekeeping: {
        methodKind: "unary";
        input: typeof MachineServiceHousekeepingRequestSchema;
        output: typeof MachineServiceHousekeepingResponseSchema;
    };
}>;
