import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { UpdateMeta } from "../../api/v2/common_pb";
import type { MachineQuery } from "../../api/v2/machine_pb";
import type { Switch, SwitchNic, SwitchOS, SwitchPortStatus, SwitchQuery, SwitchReplaceMode, SwitchWithMachines } from "../../api/v2/switch_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/switch.proto.
 */
export declare const file_metalstack_admin_v2_switch: GenFile;
/**
 * SwitchServiceGetRequest is the request payload for getting a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceGetRequest
 */
export type SwitchServiceGetRequest = Message<"metalstack.admin.v2.SwitchServiceGetRequest"> & {
    /**
     * Id of the switch to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceGetRequest.
 * Use `create(SwitchServiceGetRequestSchema)` to create a new message.
 */
export declare const SwitchServiceGetRequestSchema: GenMessage<SwitchServiceGetRequest>;
/**
 * SwitchServiceGetResponse is the response payload for getting a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceGetResponse
 */
export type SwitchServiceGetResponse = Message<"metalstack.admin.v2.SwitchServiceGetResponse"> & {
    /**
     * Switch contains the requested switch
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceGetResponse.
 * Use `create(SwitchServiceGetResponseSchema)` to create a new message.
 */
export declare const SwitchServiceGetResponseSchema: GenMessage<SwitchServiceGetResponse>;
/**
 * SwitchServiceListRequest is the request payload for listing switches.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceListRequest
 */
export type SwitchServiceListRequest = Message<"metalstack.admin.v2.SwitchServiceListRequest"> & {
    /**
     * Query to filter the results
     *
     * @generated from field: metalstack.api.v2.SwitchQuery query = 1;
     */
    query?: SwitchQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceListRequest.
 * Use `create(SwitchServiceListRequestSchema)` to create a new message.
 */
export declare const SwitchServiceListRequestSchema: GenMessage<SwitchServiceListRequest>;
/**
 * SwitchServiceListResponse is the response payload for listing switches.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceListResponse
 */
export type SwitchServiceListResponse = Message<"metalstack.admin.v2.SwitchServiceListResponse"> & {
    /**
     * Switches that match the request query
     *
     * @generated from field: repeated metalstack.api.v2.Switch switches = 1;
     */
    switches: Switch[];
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceListResponse.
 * Use `create(SwitchServiceListResponseSchema)` to create a new message.
 */
export declare const SwitchServiceListResponseSchema: GenMessage<SwitchServiceListResponse>;
/**
 * SwitchServiceUpdateRequest is the request payload for updating a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceUpdateRequest
 */
export type SwitchServiceUpdateRequest = Message<"metalstack.admin.v2.SwitchServiceUpdateRequest"> & {
    /**
     * ID of the switch
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta | undefined;
    /**
     * UpdatedAt is the date when this entity was updated
     * Must be part of the update request to ensure optimistic locking
     *
     * @generated from field: google.protobuf.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp | undefined;
    /**
     * Description of the switch
     *
     * @generated from field: optional string description = 4;
     */
    description?: string | undefined;
    /**
     * ReplaceMode is used to mark a switch ready for replacement
     *
     * @generated from field: optional metalstack.api.v2.SwitchReplaceMode replace_mode = 5;
     */
    replaceMode?: SwitchReplaceMode | undefined;
    /**
     * ManagementIp is the switch's IP for management access
     *
     * @generated from field: optional string management_ip = 6;
     */
    managementIp?: string | undefined;
    /**
     * ManagementUser is the user name to use for management access
     *
     * @generated from field: optional string management_user = 7;
     */
    managementUser?: string | undefined;
    /**
     * ConsoleCommand is the command for accessing the switch's console
     *
     * @generated from field: optional string console_command = 8;
     */
    consoleCommand?: string | undefined;
    /**
     * Nics are the front panel ports of the switch
     *
     * @generated from field: repeated metalstack.api.v2.SwitchNic nics = 9;
     */
    nics: SwitchNic[];
    /**
     * SwitchOs is the OS running on the switch
     *
     * @generated from field: optional metalstack.api.v2.SwitchOS os = 10;
     */
    os?: SwitchOS | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceUpdateRequest.
 * Use `create(SwitchServiceUpdateRequestSchema)` to create a new message.
 */
export declare const SwitchServiceUpdateRequestSchema: GenMessage<SwitchServiceUpdateRequest>;
/**
 * SwitchServiceUpdateResponse is the response payload for updating a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceUpdateResponse
 */
export type SwitchServiceUpdateResponse = Message<"metalstack.admin.v2.SwitchServiceUpdateResponse"> & {
    /**
     * Switch contains the updated switch
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceUpdateResponse.
 * Use `create(SwitchServiceUpdateResponseSchema)` to create a new message.
 */
export declare const SwitchServiceUpdateResponseSchema: GenMessage<SwitchServiceUpdateResponse>;
/**
 * SwitchServiceDeleteRequest is the request payload for deleting a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceDeleteRequest
 */
export type SwitchServiceDeleteRequest = Message<"metalstack.admin.v2.SwitchServiceDeleteRequest"> & {
    /**
     * Id of the switch
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Force will allow switch deletion despite existing machine connections
     *
     * @generated from field: bool force = 2;
     */
    force: boolean;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceDeleteRequest.
 * Use `create(SwitchServiceDeleteRequestSchema)` to create a new message.
 */
export declare const SwitchServiceDeleteRequestSchema: GenMessage<SwitchServiceDeleteRequest>;
/**
 * SwitchServiceDeleteResponse is the response payload for deleting a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceDeleteResponse
 */
export type SwitchServiceDeleteResponse = Message<"metalstack.admin.v2.SwitchServiceDeleteResponse"> & {
    /**
     * Switch contains the deleted switch
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceDeleteResponse.
 * Use `create(SwitchServiceDeleteResponseSchema)` to create a new message.
 */
export declare const SwitchServiceDeleteResponseSchema: GenMessage<SwitchServiceDeleteResponse>;
/**
 * SwitchServiceMigrateRequest is the request payload for migrating a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceMigrateRequest
 */
export type SwitchServiceMigrateRequest = Message<"metalstack.admin.v2.SwitchServiceMigrateRequest"> & {
    /**
     * OldSwitch is the switch to migrate away from
     *
     * @generated from field: string old_switch = 1;
     */
    oldSwitch: string;
    /**
     * NewSwitch is the switch to migrate to
     *
     * @generated from field: string new_switch = 2;
     */
    newSwitch: string;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceMigrateRequest.
 * Use `create(SwitchServiceMigrateRequestSchema)` to create a new message.
 */
export declare const SwitchServiceMigrateRequestSchema: GenMessage<SwitchServiceMigrateRequest>;
/**
 * SwitchServiceMigrateResponse is the response payload for migrating a switch.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceMigrateResponse
 */
export type SwitchServiceMigrateResponse = Message<"metalstack.admin.v2.SwitchServiceMigrateResponse"> & {
    /**
     * Switch contains the switch was migrated to
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceMigrateResponse.
 * Use `create(SwitchServiceMigrateResponseSchema)` to create a new message.
 */
export declare const SwitchServiceMigrateResponseSchema: GenMessage<SwitchServiceMigrateResponse>;
/**
 * SwitchServicePortRequest is the request payload for setting the port status of a switch port.
 *
 * @generated from message metalstack.admin.v2.SwitchServicePortRequest
 */
export type SwitchServicePortRequest = Message<"metalstack.admin.v2.SwitchServicePortRequest"> & {
    /**
     * Id of the switch
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * NicName of the port whose status should be changed
     *
     * @generated from field: string nic_name = 2;
     */
    nicName: string;
    /**
     * Status that the port should have
     *
     * @generated from field: metalstack.api.v2.SwitchPortStatus status = 3;
     */
    status: SwitchPortStatus;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServicePortRequest.
 * Use `create(SwitchServicePortRequestSchema)` to create a new message.
 */
export declare const SwitchServicePortRequestSchema: GenMessage<SwitchServicePortRequest>;
/**
 * SwitchServicePortResponse is the response payload for setting the port status of a switch port.
 *
 * @generated from message metalstack.admin.v2.SwitchServicePortResponse
 */
export type SwitchServicePortResponse = Message<"metalstack.admin.v2.SwitchServicePortResponse"> & {
    /**
     * Switch after the port status toggle
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServicePortResponse.
 * Use `create(SwitchServicePortResponseSchema)` to create a new message.
 */
export declare const SwitchServicePortResponseSchema: GenMessage<SwitchServicePortResponse>;
/**
 * SwitchServiceConnectedMachinesRequest is the request payload for listing switches with machine connections.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceConnectedMachinesRequest
 */
export type SwitchServiceConnectedMachinesRequest = Message<"metalstack.admin.v2.SwitchServiceConnectedMachinesRequest"> & {
    /**
     * Query to filter the switch results
     *
     * @generated from field: metalstack.api.v2.SwitchQuery query = 1;
     */
    query?: SwitchQuery | undefined;
    /**
     * MachineQuery to filter the machine results
     *
     * @generated from field: metalstack.api.v2.MachineQuery machine_query = 2;
     */
    machineQuery?: MachineQuery | undefined;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceConnectedMachinesRequest.
 * Use `create(SwitchServiceConnectedMachinesRequestSchema)` to create a new message.
 */
export declare const SwitchServiceConnectedMachinesRequestSchema: GenMessage<SwitchServiceConnectedMachinesRequest>;
/**
 * SwitchServiceConnectedMachinesResponse is the response payload for listing switches with machine connections.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceConnectedMachinesResponse
 */
export type SwitchServiceConnectedMachinesResponse = Message<"metalstack.admin.v2.SwitchServiceConnectedMachinesResponse"> & {
    /**
     * SwitchesWithMachines contains all switches with their machine connections
     *
     * @generated from field: repeated metalstack.api.v2.SwitchWithMachines switches_with_machines = 1;
     */
    switchesWithMachines: SwitchWithMachines[];
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceConnectedMachinesResponse.
 * Use `create(SwitchServiceConnectedMachinesResponseSchema)` to create a new message.
 */
export declare const SwitchServiceConnectedMachinesResponseSchema: GenMessage<SwitchServiceConnectedMachinesResponse>;
/**
 * SwitchService provides switch management operations.
 *
 * @generated from service metalstack.admin.v2.SwitchService
 */
export declare const SwitchService: GenService<{
    /**
     * Returns the switch with the specified ID.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SwitchServiceGetRequestSchema;
        output: typeof SwitchServiceGetResponseSchema;
    };
    /**
     * Returns the list of switches.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SwitchServiceListRequestSchema;
        output: typeof SwitchServiceListResponseSchema;
    };
    /**
     * Updates a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof SwitchServiceUpdateRequestSchema;
        output: typeof SwitchServiceUpdateResponseSchema;
    };
    /**
     * Deletes a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof SwitchServiceDeleteRequestSchema;
        output: typeof SwitchServiceDeleteResponseSchema;
    };
    /**
     * Migrates a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Migrate
     */
    migrate: {
        methodKind: "unary";
        input: typeof SwitchServiceMigrateRequestSchema;
        output: typeof SwitchServiceMigrateResponseSchema;
    };
    /**
     * Sets the port status of a switch port.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Port
     */
    port: {
        methodKind: "unary";
        input: typeof SwitchServicePortRequestSchema;
        output: typeof SwitchServicePortResponseSchema;
    };
    /**
     * Returns all switches with their machine connections.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.ConnectedMachines
     */
    connectedMachines: {
        methodKind: "unary";
        input: typeof SwitchServiceConnectedMachinesRequestSchema;
        output: typeof SwitchServiceConnectedMachinesResponseSchema;
    };
}>;
