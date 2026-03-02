import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { MachineBMC, MachineBMCCommand, MachineBMCReport } from "../../api/v2/machine_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/bmc.proto.
 */
export declare const file_metalstack_infra_v2_bmc: GenFile;
/**
 * UpdateBMCInfoRequest
 *
 * @generated from message metalstack.infra.v2.UpdateBMCInfoRequest
 */
export type UpdateBMCInfoRequest = Message<"metalstack.infra.v2.UpdateBMCInfoRequest"> & {
    /**
     * Partition the partition id where metal-bmc wants to receive events
     *
     * @generated from field: string partition = 1;
     */
    partition: string;
    /**
     * BmcReports contains maps the bmc report per machine uuid
     *
     * @generated from field: map<string, metalstack.api.v2.MachineBMCReport> bmc_reports = 2;
     */
    bmcReports: {
        [key: string]: MachineBMCReport;
    };
};
/**
 * Describes the message metalstack.infra.v2.UpdateBMCInfoRequest.
 * Use `create(UpdateBMCInfoRequestSchema)` to create a new message.
 */
export declare const UpdateBMCInfoRequestSchema: GenMessage<UpdateBMCInfoRequest>;
/**
 * UpdateBMCInfoResponse
 *
 * @generated from message metalstack.infra.v2.UpdateBMCInfoResponse
 */
export type UpdateBMCInfoResponse = Message<"metalstack.infra.v2.UpdateBMCInfoResponse"> & {
    /**
     * UpdatedMachines is a slice of machine uuids which were updated
     *
     * @generated from field: repeated string updated_machines = 1;
     */
    updatedMachines: string[];
    /**
     * CreatedMachines is a slice of machine uuids which were created
     *
     * @generated from field: repeated string created_machines = 2;
     */
    createdMachines: string[];
};
/**
 * Describes the message metalstack.infra.v2.UpdateBMCInfoResponse.
 * Use `create(UpdateBMCInfoResponseSchema)` to create a new message.
 */
export declare const UpdateBMCInfoResponseSchema: GenMessage<UpdateBMCInfoResponse>;
/**
 * WaitForBMCCommandRequest
 *
 * @generated from message metalstack.infra.v2.WaitForBMCCommandRequest
 */
export type WaitForBMCCommandRequest = Message<"metalstack.infra.v2.WaitForBMCCommandRequest"> & {
    /**
     * Partition the partition id where metal-bmc wants to receive bmc commands
     *
     * @generated from field: string partition = 1;
     */
    partition: string;
};
/**
 * Describes the message metalstack.infra.v2.WaitForBMCCommandRequest.
 * Use `create(WaitForBMCCommandRequestSchema)` to create a new message.
 */
export declare const WaitForBMCCommandRequestSchema: GenMessage<WaitForBMCCommandRequest>;
/**
 * WaitForBMCCommandResponse
 *
 * @generated from message metalstack.infra.v2.WaitForBMCCommandResponse
 */
export type WaitForBMCCommandResponse = Message<"metalstack.infra.v2.WaitForBMCCommandResponse"> & {
    /**
     * UUID of the machine to send the command to
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * BmcCommand to execute against the bmc of the machine
     *
     * @generated from field: metalstack.api.v2.MachineBMCCommand bmc_command = 2;
     */
    bmcCommand: MachineBMCCommand;
    /**
     * MachineBmc contains connection details of the machine to issue the bmcCommand to
     *
     * @generated from field: metalstack.api.v2.MachineBMC machine_bmc = 3;
     */
    machineBmc?: MachineBMC;
    /**
     * CommandId is a unique ID which must be sent back after execution
     * it is usually in the form: <machine-uuid>:machine-bmc-command:<command>
     *
     * @generated from field: string command_id = 4;
     */
    commandId: string;
};
/**
 * Describes the message metalstack.infra.v2.WaitForBMCCommandResponse.
 * Use `create(WaitForBMCCommandResponseSchema)` to create a new message.
 */
export declare const WaitForBMCCommandResponseSchema: GenMessage<WaitForBMCCommandResponse>;
/**
 * BMCCommandDoneRequest must be returned after command execution
 *
 * @generated from message metalstack.infra.v2.BMCCommandDoneRequest
 */
export type BMCCommandDoneRequest = Message<"metalstack.infra.v2.BMCCommandDoneRequest"> & {
    /**
     * CommandId is a unique ID which must be sent back after execution
     * it is usually in the form: <machine-uuid>:machine-bmc-command
     *
     * @generated from field: string command_id = 1;
     */
    commandId: string;
    /**
     * Error of the command execution, nil if it was successful
     *
     * @generated from field: optional string error = 2;
     */
    error?: string;
};
/**
 * Describes the message metalstack.infra.v2.BMCCommandDoneRequest.
 * Use `create(BMCCommandDoneRequestSchema)` to create a new message.
 */
export declare const BMCCommandDoneRequestSchema: GenMessage<BMCCommandDoneRequest>;
/**
 * BMCCommandDoneResponse
 *
 * @generated from message metalstack.infra.v2.BMCCommandDoneResponse
 */
export type BMCCommandDoneResponse = Message<"metalstack.infra.v2.BMCCommandDoneResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.BMCCommandDoneResponse.
 * Use `create(BMCCommandDoneResponseSchema)` to create a new message.
 */
export declare const BMCCommandDoneResponseSchema: GenMessage<BMCCommandDoneResponse>;
/**
 * BMCService serves bmc related functions
 *
 * @generated from service metalstack.infra.v2.BMCService
 */
export declare const BMCService: GenService<{
    /**
     * UpdateBMCInfo
     *
     * @generated from rpc metalstack.infra.v2.BMCService.UpdateBMCInfo
     */
    updateBMCInfo: {
        methodKind: "unary";
        input: typeof UpdateBMCInfoRequestSchema;
        output: typeof UpdateBMCInfoResponseSchema;
    };
    /**
     * WaitForBMCCommand is called by the metal-bmc and is returned with a bmc command to execute.
     *
     * @generated from rpc metalstack.infra.v2.BMCService.WaitForBMCCommand
     */
    waitForBMCCommand: {
        methodKind: "server_streaming";
        input: typeof WaitForBMCCommandRequestSchema;
        output: typeof WaitForBMCCommandResponseSchema;
    };
    /**
     * BMCCommandDone must be called from metal-bmc after the command execution
     *
     * @generated from rpc metalstack.infra.v2.BMCService.BMCCommandDone
     */
    bMCCommandDone: {
        methodKind: "unary";
        input: typeof BMCCommandDoneRequestSchema;
        output: typeof BMCCommandDoneResponseSchema;
    };
}>;
