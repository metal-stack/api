import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Duration } from "@bufbuild/protobuf/wkt";
import type { Meta } from "../../api/v2/common_pb";
import type { NicState, Switch, SwitchBGPPortState, SwitchOS, SwitchPortStatus, SwitchSync } from "../../api/v2/switch_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/infra/v2/switch.proto.
 */
export declare const file_metalstack_infra_v2_switch: GenFile;
/**
 * SwitchServiceGetRequest is the request payload for getting a switch.
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
 * SwitchServiceGetResponse is the response payload for getting a switch.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceGetResponse
 */
export type SwitchServiceGetResponse = Message<"metalstack.infra.v2.SwitchServiceGetResponse"> & {
    /**
     * Switch contains the requested switch.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceGetResponse.
 * Use `create(SwitchServiceGetResponseSchema)` to create a new message.
 */
export declare const SwitchServiceGetResponseSchema: GenMessage<SwitchServiceGetResponse>;
/**
 * SwitchServiceRegisterRequest is the request payload for registering a switch.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceRegisterRequest
 */
export type SwitchServiceRegisterRequest = Message<"metalstack.infra.v2.SwitchServiceRegisterRequest"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 2;
     */
    id: string;
    /**
     * Meta for this switch.
     *
     * @generated from field: metalstack.api.v2.Meta meta = 3;
     */
    meta?: Meta | undefined;
    /**
     * Rack ID if the switch resides in a rack.
     *
     * @generated from field: optional string rack = 4;
     */
    rack?: string | undefined;
    /**
     * Room ID if the switch resides in a room.
     *
     * @generated from field: optional string room = 5;
     */
    room?: string | undefined;
    /**
     * Partition the switch belongs to.
     *
     * @generated from field: string partition = 6;
     */
    partition: string;
    /**
     * ManagementIp is the switch's IP for management access.
     *
     * @generated from field: string management_ip = 7;
     */
    managementIp: string;
    /**
     * ManagementUser is the user name to use for management access.
     *
     * @generated from field: optional string management_user = 8;
     */
    managementUser?: string | undefined;
    /**
     * Nics are the front panel ports of the switch.
     *
     * @generated from field: repeated metalstack.infra.v2.SwitchRegisterNic nics = 9;
     */
    nics: SwitchRegisterNic[];
    /**
     * SwitchOs is the OS running on the switch.
     *
     * @generated from field: metalstack.api.v2.SwitchOS os = 10;
     */
    os?: SwitchOS | undefined;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceRegisterRequest.
 * Use `create(SwitchServiceRegisterRequestSchema)` to create a new message.
 */
export declare const SwitchServiceRegisterRequestSchema: GenMessage<SwitchServiceRegisterRequest>;
/**
 * SwitchRegisterNic is the switch nic used when registering a switch at the api.
 *
 * @generated from message metalstack.infra.v2.SwitchRegisterNic
 */
export type SwitchRegisterNic = Message<"metalstack.infra.v2.SwitchRegisterNic"> & {
    /**
     * Name of the switch port.
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Identifier of the port.
     *
     * @generated from field: string identifier = 2;
     */
    identifier: string;
    /**
     * NicState describes the current state of the switch port.
     *
     * @generated from field: metalstack.api.v2.NicState state = 3;
     */
    state?: NicState | undefined;
};
/**
 * Describes the message metalstack.infra.v2.SwitchRegisterNic.
 * Use `create(SwitchRegisterNicSchema)` to create a new message.
 */
export declare const SwitchRegisterNicSchema: GenMessage<SwitchRegisterNic>;
/**
 * SwitchServiceRegisterResponse is the response payload for registering a switch.
 *
 * @generated from message metalstack.infra.v2.SwitchServiceRegisterResponse
 */
export type SwitchServiceRegisterResponse = Message<"metalstack.infra.v2.SwitchServiceRegisterResponse"> & {
    /**
     * Switch contains the registered switch.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceRegisterResponse.
 * Use `create(SwitchServiceRegisterResponseSchema)` to create a new message.
 */
export declare const SwitchServiceRegisterResponseSchema: GenMessage<SwitchServiceRegisterResponse>;
/**
 * SwitchServiceHeartbeatRequest is the request payload for sending a switch heartbeat.
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
    duration?: Duration | undefined;
    /**
     * Error if any occurred during the sync.
     *
     * @generated from field: optional string error = 3;
     */
    error?: string | undefined;
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
 * SwitchServiceHeartbeatResponse is the response payload for sending a switch heartbeat.
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
    lastSync?: SwitchSync | undefined;
    /**
     * LastSyncError holds information about the last erroneous sync.
     *
     * @generated from field: metalstack.api.v2.SwitchSync last_sync_error = 3;
     */
    lastSyncError?: SwitchSync | undefined;
};
/**
 * Describes the message metalstack.infra.v2.SwitchServiceHeartbeatResponse.
 * Use `create(SwitchServiceHeartbeatResponseSchema)` to create a new message.
 */
export declare const SwitchServiceHeartbeatResponseSchema: GenMessage<SwitchServiceHeartbeatResponse>;
/**
 * SwitchService provides infrastructure switch management operations.
 *
 * @generated from service metalstack.infra.v2.SwitchService
 */
export declare const SwitchService: GenService<{
    /**
     * Returns the switch with the specified ID.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SwitchServiceGetRequestSchema;
        output: typeof SwitchServiceGetResponseSchema;
    };
    /**
     * Registers a switch.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Register
     */
    register: {
        methodKind: "unary";
        input: typeof SwitchServiceRegisterRequestSchema;
        output: typeof SwitchServiceRegisterResponseSchema;
    };
    /**
     * Heartbeat sends a heartbeat from a switch.
     *
     * @generated from rpc metalstack.infra.v2.SwitchService.Heartbeat
     */
    heartbeat: {
        methodKind: "unary";
        input: typeof SwitchServiceHeartbeatRequestSchema;
        output: typeof SwitchServiceHeartbeatResponseSchema;
    };
}>;
