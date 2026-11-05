import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/housekeeping.proto.
 */
export declare const file_metalstack_admin_v2_housekeeping: GenFile;
/**
 * HousekeepingServiceMachineRequest is the request payload for a machine housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineRequest
 */
export type HousekeepingServiceMachineRequest = Message<"metalstack.admin.v2.HousekeepingServiceMachineRequest"> & {
    /**
     * Housekeeping Action to perform.
     *
     * @generated from field: metalstack.admin.v2.MachineHousekeepingAction action = 1;
     */
    action: MachineHousekeepingAction;
};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineRequest.
 * Use `create(HousekeepingServiceMachineRequestSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineRequestSchema: GenMessage<HousekeepingServiceMachineRequest>;
/**
 * HousekeepingServiceMachineResponse is the response payload for a machine housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineResponse
 */
export type HousekeepingServiceMachineResponse = Message<"metalstack.admin.v2.HousekeepingServiceMachineResponse"> & {
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
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineResponse.
 * Use `create(HousekeepingServiceMachineResponseSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineResponseSchema: GenMessage<HousekeepingServiceMachineResponse>;
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
 * HousekeepingService provides housekeeping operations.
 *
 * @generated from service metalstack.admin.v2.HousekeepingService
 */
export declare const HousekeepingService: GenService<{
    /**
     * Machine housekeeping
     *
     * @generated from rpc metalstack.admin.v2.HousekeepingService.Machine
     */
    machine: {
        methodKind: "unary";
        input: typeof HousekeepingServiceMachineRequestSchema;
        output: typeof HousekeepingServiceMachineResponseSchema;
    };
}>;
