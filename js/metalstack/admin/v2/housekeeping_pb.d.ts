import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Machine, MachineLiveliness } from "../../api/v2/machine_pb";
import type { VPNNode } from "../../api/v2/vpn_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/housekeeping.proto.
 */
export declare const file_metalstack_admin_v2_housekeeping: GenFile;
/**
 * HousekeepingServiceMachineLivelinessRequest is the request payload for a machine liveliness housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineLivelinessRequest
 */
export type HousekeepingServiceMachineLivelinessRequest = Message<"metalstack.admin.v2.HousekeepingServiceMachineLivelinessRequest"> & {};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineLivelinessRequest.
 * Use `create(HousekeepingServiceMachineLivelinessRequestSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineLivelinessRequestSchema: GenMessage<HousekeepingServiceMachineLivelinessRequest>;
/**
 * HousekeepingServiceMachineLivelinessResponse is the response payload for a machine liveliness housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineLivelinessResponse
 */
export type HousekeepingServiceMachineLivelinessResponse = Message<"metalstack.admin.v2.HousekeepingServiceMachineLivelinessResponse"> & {
    /**
     * Modified Machines modified during this housekeeping action.
     *
     * @generated from field: repeated metalstack.admin.v2.MachineWithLiveliness modified_machines = 1;
     */
    modifiedMachines: MachineWithLiveliness[];
};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineLivelinessResponse.
 * Use `create(HousekeepingServiceMachineLivelinessResponseSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineLivelinessResponseSchema: GenMessage<HousekeepingServiceMachineLivelinessResponse>;
/**
 * MachineWithLiveliness
 *
 * @generated from message metalstack.admin.v2.MachineWithLiveliness
 */
export type MachineWithLiveliness = Message<"metalstack.admin.v2.MachineWithLiveliness"> & {
    /**
     * UUID of this machine.
     *
     * @generated from field: string uuid = 1;
     */
    uuid: string;
    /**
     * Liveliness of this machine.
     *
     * @generated from field: metalstack.api.v2.MachineLiveliness liveliness = 2;
     */
    liveliness: MachineLiveliness;
};
/**
 * Describes the message metalstack.admin.v2.MachineWithLiveliness.
 * Use `create(MachineWithLivelinessSchema)` to create a new message.
 */
export declare const MachineWithLivelinessSchema: GenMessage<MachineWithLiveliness>;
/**
 * HousekeepingServiceMachineVPNConnectedRequest is the request payload for a machine vpn connected housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedRequest
 */
export type HousekeepingServiceMachineVPNConnectedRequest = Message<"metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedRequest"> & {};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedRequest.
 * Use `create(HousekeepingServiceMachineVPNConnectedRequestSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineVPNConnectedRequestSchema: GenMessage<HousekeepingServiceMachineVPNConnectedRequest>;
/**
 * HousekeepingServiceMachineVPNConnectedResponse is the response payload for a machine vpn connected housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedResponse
 */
export type HousekeepingServiceMachineVPNConnectedResponse = Message<"metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedResponse"> & {
    /**
     * Nodes connected to the VPN.
     *
     * @generated from field: repeated metalstack.api.v2.VPNNode modified_machines = 1;
     */
    modifiedMachines: VPNNode[];
};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineVPNConnectedResponse.
 * Use `create(HousekeepingServiceMachineVPNConnectedResponseSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineVPNConnectedResponseSchema: GenMessage<HousekeepingServiceMachineVPNConnectedResponse>;
/**
 * HousekeepingServiceMachineResurrectRequest is the request payload for a machine resurrect housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineResurrectRequest
 */
export type HousekeepingServiceMachineResurrectRequest = Message<"metalstack.admin.v2.HousekeepingServiceMachineResurrectRequest"> & {};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineResurrectRequest.
 * Use `create(HousekeepingServiceMachineResurrectRequestSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineResurrectRequestSchema: GenMessage<HousekeepingServiceMachineResurrectRequest>;
/**
 * HousekeepingServiceMachineResurrectResponse is the response payload for a machine resurrect housekeeping request.
 *
 * @generated from message metalstack.admin.v2.HousekeepingServiceMachineResurrectResponse
 */
export type HousekeepingServiceMachineResurrectResponse = Message<"metalstack.admin.v2.HousekeepingServiceMachineResurrectResponse"> & {
    /**
     * Machines which have been resurrected
     *
     * @generated from field: repeated metalstack.api.v2.Machine modified_machines = 1;
     */
    modifiedMachines: Machine[];
};
/**
 * Describes the message metalstack.admin.v2.HousekeepingServiceMachineResurrectResponse.
 * Use `create(HousekeepingServiceMachineResurrectResponseSchema)` to create a new message.
 */
export declare const HousekeepingServiceMachineResurrectResponseSchema: GenMessage<HousekeepingServiceMachineResurrectResponse>;
/**
 * HousekeepingService provides housekeeping operations.
 *
 * @generated from service metalstack.admin.v2.HousekeepingService
 */
export declare const HousekeepingService: GenService<{
    /**
     * Machine liveliness
     *
     * @generated from rpc metalstack.admin.v2.HousekeepingService.MachineLiveliness
     */
    machineLiveliness: {
        methodKind: "unary";
        input: typeof HousekeepingServiceMachineLivelinessRequestSchema;
        output: typeof HousekeepingServiceMachineLivelinessResponseSchema;
    };
    /**
     * Machine vpn connected
     *
     * @generated from rpc metalstack.admin.v2.HousekeepingService.MachineVPNConnected
     */
    machineVPNConnected: {
        methodKind: "unary";
        input: typeof HousekeepingServiceMachineVPNConnectedRequestSchema;
        output: typeof HousekeepingServiceMachineVPNConnectedResponseSchema;
    };
    /**
     * Machine resurrect dead machines
     *
     * @generated from rpc metalstack.admin.v2.HousekeepingService.MachineResurrect
     */
    machineResurrect: {
        methodKind: "unary";
        input: typeof HousekeepingServiceMachineResurrectRequestSchema;
        output: typeof HousekeepingServiceMachineResurrectResponseSchema;
    };
}>;
