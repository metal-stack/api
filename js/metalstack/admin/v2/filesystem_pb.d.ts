import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { UpdateMeta } from "../../api/v2/common_pb";
import type { Disk, Filesystem, FilesystemLayout, FilesystemLayoutConstraints, LogicalVolume, Raid, VolumeGroup } from "../../api/v2/filesystem_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/admin/v2/filesystem.proto.
 */
export declare const file_metalstack_admin_v2_filesystem: GenFile;
/**
 * FilesystemServiceCreateRequest
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceCreateRequest
 */
export type FilesystemServiceCreateRequest = Message<"metalstack.admin.v2.FilesystemServiceCreateRequest"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceCreateRequest.
 * Use `create(FilesystemServiceCreateRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceCreateRequestSchema: GenMessage<FilesystemServiceCreateRequest>;
/**
 * FilesystemServiceCreateResponse
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceCreateResponse
 */
export type FilesystemServiceCreateResponse = Message<"metalstack.admin.v2.FilesystemServiceCreateResponse"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceCreateResponse.
 * Use `create(FilesystemServiceCreateResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceCreateResponseSchema: GenMessage<FilesystemServiceCreateResponse>;
/**
 * FilesystemServiceUpdateRequest
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceUpdateRequest
 */
export type FilesystemServiceUpdateRequest = Message<"metalstack.admin.v2.FilesystemServiceUpdateRequest"> & {
    /**
     * Id of this filesystemLayout
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * UpdateMeta contains the timestamp and strategy to be used in this update request
     *
     * @generated from field: metalstack.api.v2.UpdateMeta update_meta = 2;
     */
    updateMeta?: UpdateMeta;
    /**
     * Name of this filesystemLayout
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Description of this filesystemLayout
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Filesystems is a list of filesystems to create on a machine
     *
     * @generated from field: repeated metalstack.api.v2.Filesystem filesystems = 5;
     */
    filesystems: Filesystem[];
    /**
     * Disks list of disks that belong to this layout
     *
     * @generated from field: repeated metalstack.api.v2.Disk disks = 6;
     */
    disks: Disk[];
    /**
     * raid arrays to create
     *
     * @generated from field: repeated metalstack.api.v2.Raid raid = 7;
     */
    raid: Raid[];
    /**
     * VolumeGroups list of volumegroups to create
     *
     * @generated from field: repeated metalstack.api.v2.VolumeGroup volume_groups = 8;
     */
    volumeGroups: VolumeGroup[];
    /**
     * LogicalVolumes list of logicalvolumes to create
     *
     * @generated from field: repeated metalstack.api.v2.LogicalVolume logical_volumes = 9;
     */
    logicalVolumes: LogicalVolume[];
    /**
     * Constraints which must match that this layout is taken, if sizes and images are empty these are develop layouts
     *
     * @generated from field: metalstack.api.v2.FilesystemLayoutConstraints constraints = 10;
     */
    constraints?: FilesystemLayoutConstraints;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceUpdateRequest.
 * Use `create(FilesystemServiceUpdateRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceUpdateRequestSchema: GenMessage<FilesystemServiceUpdateRequest>;
/**
 * FilesystemServiceUpdateResponse
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceUpdateResponse
 */
export type FilesystemServiceUpdateResponse = Message<"metalstack.admin.v2.FilesystemServiceUpdateResponse"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceUpdateResponse.
 * Use `create(FilesystemServiceUpdateResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceUpdateResponseSchema: GenMessage<FilesystemServiceUpdateResponse>;
/**
 * FilesystemServiceDeleteRequest
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceDeleteRequest
 */
export type FilesystemServiceDeleteRequest = Message<"metalstack.admin.v2.FilesystemServiceDeleteRequest"> & {
    /**
     * ID of the filesystem to delete
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceDeleteRequest.
 * Use `create(FilesystemServiceDeleteRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceDeleteRequestSchema: GenMessage<FilesystemServiceDeleteRequest>;
/**
 * message FilesystemServiceDeleteResponse {
 *
 * @generated from message metalstack.admin.v2.FilesystemServiceDeleteResponse
 */
export type FilesystemServiceDeleteResponse = Message<"metalstack.admin.v2.FilesystemServiceDeleteResponse"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.admin.v2.FilesystemServiceDeleteResponse.
 * Use `create(FilesystemServiceDeleteResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceDeleteResponseSchema: GenMessage<FilesystemServiceDeleteResponse>;
/**
 * FilesystemService serves filesystem related functions
 *
 * @generated from service metalstack.admin.v2.FilesystemService
 */
export declare const FilesystemService: GenService<{
    /**
     * Create a filesystem
     *
     * @generated from rpc metalstack.admin.v2.FilesystemService.Create
     */
    create: {
        methodKind: "unary";
        input: typeof FilesystemServiceCreateRequestSchema;
        output: typeof FilesystemServiceCreateResponseSchema;
    };
    /**
     * Update a filesystem
     *
     * @generated from rpc metalstack.admin.v2.FilesystemService.Update
     */
    update: {
        methodKind: "unary";
        input: typeof FilesystemServiceUpdateRequestSchema;
        output: typeof FilesystemServiceUpdateResponseSchema;
    };
    /**
     * Delete a filesystem
     *
     * @generated from rpc metalstack.admin.v2.FilesystemService.Delete
     */
    delete: {
        methodKind: "unary";
        input: typeof FilesystemServiceDeleteRequestSchema;
        output: typeof FilesystemServiceDeleteResponseSchema;
    };
}>;
