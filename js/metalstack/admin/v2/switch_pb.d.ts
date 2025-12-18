import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { UpdateMeta } from "../../api/v2/common_pb";
import type { Switch, SwitchNic, SwitchOS, SwitchPortStatus, SwitchQuery, SwitchReplaceMode } from "../../api/v2/switch_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/switch.proto.
 */
export declare const file_metalstack_admin_v2_switch: GenFile;
/**
 * SwitchServiceGetRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceGetRequest
 */
export type SwitchServiceGetRequest = Message<"metalstack.admin.v2.SwitchServiceGetRequest"> & {
    /**
     * Id of the switch to get.
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
 * SwitchServiceGetResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceGetResponse
 */
export type SwitchServiceGetResponse = Message<"metalstack.admin.v2.SwitchServiceGetResponse"> & {
    /**
     * Switch that was requested.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceGetResponse.
 * Use `create(SwitchServiceGetResponseSchema)` to create a new message.
 */
export declare const SwitchServiceGetResponseSchema: GenMessage<SwitchServiceGetResponse>;
/**
 * SwitchServiceListRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceListRequest
 */
export type SwitchServiceListRequest = Message<"metalstack.admin.v2.SwitchServiceListRequest"> & {
    /**
     * Query to filter the results.
     *
     * @generated from field: metalstack.api.v2.SwitchQuery query = 1;
     */
    query?: SwitchQuery;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceListRequest.
 * Use `create(SwitchServiceListRequestSchema)` to create a new message.
 */
export declare const SwitchServiceListRequestSchema: GenMessage<SwitchServiceListRequest>;
/**
 * SwitchServiceListResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceListResponse
 */
export type SwitchServiceListResponse = Message<"metalstack.admin.v2.SwitchServiceListResponse"> & {
    /**
     * Switches that match the request query.
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
 * SwitchServiceUpdateRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceUpdateRequest
 */
export type SwitchServiceUpdateRequest = Message<"metalstack.admin.v2.SwitchServiceUpdateRequest"> & {
    /**
     * ID of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request.
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * UpdatedAt is the date when this entity was updated.
     * must be part of the update request to ensure optimistic locking.
     *
     * @generated from field: google.protobuf.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * Description of the switch.
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Replace mode is used to mark a switch ready for replacement.
     *
     * @generated from field: optional metalstack.api.v2.SwitchReplaceMode replace_mode = 5;
     */
    replaceMode?: SwitchReplaceMode;
    /**
     * Management IP is the switch's IP for management access.
     *
     * @generated from field: optional string management_ip = 6;
     */
    managementIp?: string;
    /**
     * Management user is the user name to use for management access.
     *
     * @generated from field: optional string management_user = 7;
     */
    managementUser?: string;
    /**
     * Console command is the command for accessing the switch's console.
     *
     * @generated from field: optional string console_command = 8;
     */
    consoleCommand?: string;
    /**
     * Nics are the front panel ports of the switch.
     *
     * @generated from field: repeated metalstack.api.v2.SwitchNic nics = 9;
     */
    nics: SwitchNic[];
    /**
     * SwitchOs is the OS running on the switch.
     *
     * @generated from field: optional metalstack.api.v2.SwitchOS os = 10;
     */
    os?: SwitchOS;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceUpdateRequest.
 * Use `create(SwitchServiceUpdateRequestSchema)` to create a new message.
 */
export declare const SwitchServiceUpdateRequestSchema: GenMessage<SwitchServiceUpdateRequest>;
/**
 * SwitchServiceUpdateResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceUpdateResponse
 */
export type SwitchServiceUpdateResponse = Message<"metalstack.admin.v2.SwitchServiceUpdateResponse"> & {
    /**
     * Switch that was updated.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceUpdateResponse.
 * Use `create(SwitchServiceUpdateResponseSchema)` to create a new message.
 */
export declare const SwitchServiceUpdateResponseSchema: GenMessage<SwitchServiceUpdateResponse>;
/**
 * SwitchServiceDeleteRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceDeleteRequest
 */
export type SwitchServiceDeleteRequest = Message<"metalstack.admin.v2.SwitchServiceDeleteRequest"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Force will allow switch deletion despite existing machine connections.
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
 * SwitchServiceDeleteResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceDeleteResponse
 */
export type SwitchServiceDeleteResponse = Message<"metalstack.admin.v2.SwitchServiceDeleteResponse"> & {
    /**
     * Switch that has been deleted.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceDeleteResponse.
 * Use `create(SwitchServiceDeleteResponseSchema)` to create a new message.
 */
export declare const SwitchServiceDeleteResponseSchema: GenMessage<SwitchServiceDeleteResponse>;
/**
 * SwitchServiceMigrateRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceMigrateRequest
 */
export type SwitchServiceMigrateRequest = Message<"metalstack.admin.v2.SwitchServiceMigrateRequest"> & {
    /**
     * OldSwitch which to migrate away from.
     *
     * @generated from field: string old_switch = 1;
     */
    oldSwitch: string;
    /**
     * NewSwitch which to migrate to.
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
 * SwitchServiceMigrateResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServiceMigrateResponse
 */
export type SwitchServiceMigrateResponse = Message<"metalstack.admin.v2.SwitchServiceMigrateResponse"> & {
    /**
     * Switch that was migrated to.
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServiceMigrateResponse.
 * Use `create(SwitchServiceMigrateResponseSchema)` to create a new message.
 */
export declare const SwitchServiceMigrateResponseSchema: GenMessage<SwitchServiceMigrateResponse>;
/**
 * SwitchServicePortRequest.
 *
 * @generated from message metalstack.admin.v2.SwitchServicePortRequest
 */
export type SwitchServicePortRequest = Message<"metalstack.admin.v2.SwitchServicePortRequest"> & {
    /**
     * Id of the switch.
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * NicName of the port whose status should be changed.
     *
     * @generated from field: string nic_name = 2;
     */
    nicName: string;
    /**
     * Status that the port should have.
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
 * SwitchServicePortResponse.
 *
 * @generated from message metalstack.admin.v2.SwitchServicePortResponse
 */
export type SwitchServicePortResponse = Message<"metalstack.admin.v2.SwitchServicePortResponse"> & {
    /**
     * Switch after the port status toggle..
     *
     * @generated from field: metalstack.api.v2.Switch switch = 1;
     */
    switch?: Switch;
};
/**
 * Describes the message metalstack.admin.v2.SwitchServicePortResponse.
 * Use `create(SwitchServicePortResponseSchema)` to create a new message.
 */
export declare const SwitchServicePortResponseSchema: GenMessage<SwitchServicePortResponse>;
/**
 * SwitchService serves switch related functions.
 *
 * @generated from service metalstack.admin.v2.SwitchService
 */
export declare const SwitchService: GenService<{
    /**
     * Get a switch by ID.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof SwitchServiceGetRequestSchema;
        output: typeof SwitchServiceGetResponseSchema;
    };
    /**
     * List switches.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.List
     */
    list: {
        methodKind: "unary";
        input: typeof SwitchServiceListRequestSchema;
        output: typeof SwitchServiceListResponseSchema;
    };
    /**
     * Update a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof SwitchServiceUpdateRequestSchema;
        output: typeof SwitchServiceUpdateResponseSchema;
    };
    /**
     * Delete a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof SwitchServiceDeleteRequestSchema;
        output: typeof SwitchServiceDeleteResponseSchema;
    };
    /**
     * Migrate a switch.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Migrate
     */
    migrate: {
        methodKind: "unary";
        input: typeof SwitchServiceMigrateRequestSchema;
        output: typeof SwitchServiceMigrateResponseSchema;
    };
    /**
     * Port set the port status of a switch port.
     *
     * @generated from rpc metalstack.admin.v2.SwitchService.Port
     */
    port: {
        methodKind: "unary";
        input: typeof SwitchServicePortRequestSchema;
        output: typeof SwitchServicePortResponseSchema;
    };
}>;
