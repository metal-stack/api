import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
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
export type UpdateBMCInfoRequest = Message<"metalstack.infra.v2.UpdateBMCInfoRequest"> & {};
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
export type UpdateBMCInfoResponse = Message<"metalstack.infra.v2.UpdateBMCInfoResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.UpdateBMCInfoResponse.
 * Use `create(UpdateBMCInfoResponseSchema)` to create a new message.
 */
export declare const UpdateBMCInfoResponseSchema: GenMessage<UpdateBMCInfoResponse>;
/**
 * WaitForMachineEventRequest
 *
 * @generated from message metalstack.infra.v2.WaitForMachineEventRequest
 */
export type WaitForMachineEventRequest = Message<"metalstack.infra.v2.WaitForMachineEventRequest"> & {};
/**
 * Describes the message metalstack.infra.v2.WaitForMachineEventRequest.
 * Use `create(WaitForMachineEventRequestSchema)` to create a new message.
 */
export declare const WaitForMachineEventRequestSchema: GenMessage<WaitForMachineEventRequest>;
/**
 * WaitForMachineEventResponse
 *
 * @generated from message metalstack.infra.v2.WaitForMachineEventResponse
 */
export type WaitForMachineEventResponse = Message<"metalstack.infra.v2.WaitForMachineEventResponse"> & {};
/**
 * Describes the message metalstack.infra.v2.WaitForMachineEventResponse.
 * Use `create(WaitForMachineEventResponseSchema)` to create a new message.
 */
export declare const WaitForMachineEventResponseSchema: GenMessage<WaitForMachineEventResponse>;
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
     * WaitForMachineEvent is called by the metal-bmc and is returned with a bmc command to execute.
     *
     * @generated from rpc metalstack.infra.v2.BMCService.WaitForMachineEvent
     */
    waitForMachineEvent: {
        methodKind: "server_streaming";
        input: typeof WaitForMachineEventRequestSchema;
        output: typeof WaitForMachineEventResponseSchema;
    };
}>;
