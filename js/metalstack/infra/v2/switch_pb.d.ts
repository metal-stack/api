import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration } from "@bufbuild/protobuf/wkt";
import type { Switch, SwitchBGPPortState, SwitchPortStatus, SwitchSync } from "../../api/v2/switch_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/switch.proto.
 */
export declare const file_metalstack_infra_v2_switch: GenFile;
/**
 * SwitchServiceGetRequest.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceGetRequest
 */
export type SwitchServiceGetRequest = Message<"metalstack.infra.v2.SwitchServiceGetRequest"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceGetRequest.
 * Use `create(SwitchServiceGetRequestSchema)` to create a new message.
 */
export declare const SwitchServiceGetRequestSchema: GenMessage<SwitchServiceGetRequest>;
/**
 * SwitchServiceGetResponse.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceGetResponse
 */
export type SwitchServiceGetResponse = Message<"metalstack.infra.v2.SwitchServiceGetResponse"> & {
    /**
     * Switch that was requested.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceGetResponse.
 * Use `create(SwitchServiceGetResponseSchema)` to create a new message.
 */
export declare const SwitchServiceGetResponseSchema: GenMessage<SwitchServiceGetResponse>;
/**
 * SwitchServiceRegisterRequest.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceRegisterRequest
 */
export type SwitchServiceRegisterRequest = Message<"metalstack.infra.v2.SwitchServiceRegisterRequest"> & {
    /**
     * Switch to register.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceRegisterRequest.
 * Use `create(SwitchServiceRegisterRequestSchema)` to create a new message.
 */
export declare const SwitchServiceRegisterRequestSchema: GenMessage<SwitchServiceRegisterRequest>;
/**
 * SwitchServiceRegisterResponse.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceRegisterResponse
 */
export type SwitchServiceRegisterResponse = Message<"metalstack.infra.v2.SwitchServiceRegisterResponse"> & {
    /**
     * Switch that was registered.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceRegisterResponse.
 * Use `create(SwitchServiceRegisterResponseSchema)` to create a new message.
 */
export declare const SwitchServiceRegisterResponseSchema: GenMessage<SwitchServiceRegisterResponse>;
/**
 * SwitchServiceHeartbeatRequest.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceHeartbeatRequest
 */
export type SwitchServiceHeartbeatRequest = Message<"metalstack.infra.v2.SwitchServiceHeartbeatRequest"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Duration of the sync.
     *
     * @generated from field: google.protobuf.Duration duration = 2;
     */
    duration?: Duration;
    /**
     * Error if any occurred during the sync.
     *
     * @generated from field: optional string error = 3;
     */
    error?: string;
    /**
     * PortStates maps port identifiers to the respective port's operational state.
     *
     * @generated from field: map<string, metalstack.api.v2.SwitchPortStatus> port_states = 4;
     */
    portStates: {
        [key: string]: SwitchPortStatus;
    };
    /**
     * BgpPortStates maps port identifiers to the respective port's BGP state.
     *
     * @generated from field: map<string, metalstack.api.v2.SwitchBGPPortState> bgp_port_states = 5;
     */
    bgpPortStates: {
        [key: string]: SwitchBGPPortState;
    };
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceHeartbeatRequest.
 * Use `create(SwitchServiceHeartbeatRequestSchema)` to create a new message.
 */
export declare const SwitchServiceHeartbeatRequestSchema: GenMessage<SwitchServiceHeartbeatRequest>;
/**
 * SwitchServiceHeartbeatResponse.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceHeartbeatResponse
 */
export type SwitchServiceHeartbeatResponse = Message<"metalstack.infra.v2.SwitchServiceHeartbeatResponse"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * LastSync holds information about the last sync.
     *
     * @generated from field: metalstack.api.v2.SwitchSync last_sync = 2;
     */
    lastSync?: SwitchSync;
    /**
     * LastSyncError holds information about the last erroneous sync.
     *
     * @generated from field: metalstack.api.v2.SwitchSync last_sync_error = 3;
     */
    lastSyncError?: SwitchSync;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceHeartbeatResponse.
 * Use `create(SwitchServiceHeartbeatResponseSchema)` to create a new message.
 */
export declare const SwitchServiceHeartbeatResponseSchema: GenMessage<SwitchServiceHeartbeatResponse>;
/**
 * SwitchService serves switch related functions.
 *
 * @generated from service metalstack.infra.v2.SwitchService
 */
export declare const SwitchService: GenService<{
    /**
     * Get a switch by ID.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SwitchServiceGetRequestSchema;
        output: typeof SwitchServiceGetResponseSchema;
    };
    /**
     * Register a switch.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Register
     */
    register: {
        methodKind: "unary";
        input: typeof SwitchServiceRegisterRequestSchema;
        output: typeof SwitchServiceRegisterResponseSchema;
    };
    /**
     * Heartbeat a switch.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Heartbeat
     */
    heartbeat: {
        methodKind: "unary";
        input: typeof SwitchServiceHeartbeatRequestSchema;
        output: typeof SwitchServiceHeartbeatResponseSchema;
    };
}>;
