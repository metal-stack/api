import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Meta } from "./common_pb";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file metalstack/api/v2/filesystem.proto.
 */
export declare const file_metalstack_api_v2_filesystem: GenFile;
/**
 * FilesystemServiceGetRequest is the request payload for a filesystem get request
 *
 * @generated from message metalstack.api.v2.FilesystemServiceGetRequest
 */
export type FilesystemServiceGetRequest = Message<"metalstack.api.v2.FilesystemServiceGetRequest"> & {
    /**
     * ID of the filesystem to get
     *
     * @generated from field: string id = 1;
     */
    id: string;
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceGetRequest.
 * Use `create(FilesystemServiceGetRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceGetRequestSchema: GenMessage<FilesystemServiceGetRequest>;
/**
 * FilesystemServiceListRequest is the request payload for a filesystem list request
 *
 * @generated from message metalstack.api.v2.FilesystemServiceListRequest
 */
export type FilesystemServiceListRequest = Message<"metalstack.api.v2.FilesystemServiceListRequest"> & {
    /**
     * ID of the filesystem to get
     *
     * @generated from field: optional string id = 1;
     */
    id?: string;
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceListRequest.
 * Use `create(FilesystemServiceListRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceListRequestSchema: GenMessage<FilesystemServiceListRequest>;
/**
 * FilesystemServiceGetResponse is the response payload for a filesystem get request
 *
 * @generated from message metalstack.api.v2.FilesystemServiceGetResponse
 */
export type FilesystemServiceGetResponse = Message<"metalstack.api.v2.FilesystemServiceGetResponse"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceGetResponse.
 * Use `create(FilesystemServiceGetResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceGetResponseSchema: GenMessage<FilesystemServiceGetResponse>;
/**
 * FilesystemServiceListResponse is the response payload for a filesystem list request
 *
 * @generated from message metalstack.api.v2.FilesystemServiceListResponse
 */
export type FilesystemServiceListResponse = Message<"metalstack.api.v2.FilesystemServiceListResponse"> & {
    /**
     * FilesystemLayouts the filesystemlayouts
     *
     * @generated from field: repeated metalstack.api.v2.FilesystemLayout filesystem_layouts = 1;
     */
    filesystemLayouts: FilesystemLayout[];
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceListResponse.
 * Use `create(FilesystemServiceListResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceListResponseSchema: GenMessage<FilesystemServiceListResponse>;
/**
 * FilesystemServiceMatchRequest
 *
 * @generated from message metalstack.api.v2.FilesystemServiceMatchRequest
 */
export type FilesystemServiceMatchRequest = Message<"metalstack.api.v2.FilesystemServiceMatchRequest"> & {
    /**
     * Match either size and image to a filesystemlayout
     * or if a machine matches to a filesystemlayout
     *
     * @generated from oneof metalstack.api.v2.FilesystemServiceMatchRequest.match
     */
    match: {
        /**
         * SizeAndImage
         *
         * @generated from field: metalstack.api.v2.MatchImageAndSize size_and_image = 1;
         */
        value: MatchImageAndSize;
        case: "sizeAndImage";
    } | {
        /**
         * MachineAndFilesystemlayout
         *
         * @generated from field: metalstack.api.v2.MatchMachine machine_and_filesystemlayout = 2;
         */
        value: MatchMachine;
        case: "machineAndFilesystemlayout";
    } | {
        case: undefined;
        value?: undefined;
    };
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceMatchRequest.
 * Use `create(FilesystemServiceMatchRequestSchema)` to create a new message.
 */
export declare const FilesystemServiceMatchRequestSchema: GenMessage<FilesystemServiceMatchRequest>;
/**
 * MatchImageAndSize
 *
 * @generated from message metalstack.api.v2.MatchImageAndSize
 */
export type MatchImageAndSize = Message<"metalstack.api.v2.MatchImageAndSize"> & {
    /**
     * Size, machine size to try
     *
     * @generated from field: string size = 1;
     */
    size: string;
    /**
     * Image machine image to try
     *
     * @generated from field: string image = 2;
     */
    image: string;
};
/**
 * Describes the message metalstack.api.v2.MatchImageAndSize.
 * Use `create(MatchImageAndSizeSchema)` to create a new message.
 */
export declare const MatchImageAndSizeSchema: GenMessage<MatchImageAndSize>;
/**
 * MatchMachine
 *
 * @generated from message metalstack.api.v2.MatchMachine
 */
export type MatchMachine = Message<"metalstack.api.v2.MatchMachine"> & {
    /**
     * Machine to check
     *
     * @generated from field: string machine = 1;
     */
    machine: string;
    /**
     * FilesystemLayout to check
     *
     * @generated from field: string filesystem_layout = 2;
     */
    filesystemLayout: string;
};
/**
 * Describes the message metalstack.api.v2.MatchMachine.
 * Use `create(MatchMachineSchema)` to create a new message.
 */
export declare const MatchMachineSchema: GenMessage<MatchMachine>;
/**
 * FilesystemServiceMatchResponse
 *
 * @generated from message metalstack.api.v2.FilesystemServiceMatchResponse
 */
export type FilesystemServiceMatchResponse = Message<"metalstack.api.v2.FilesystemServiceMatchResponse"> & {
    /**
     * FilesystemLayout the filesystemlayout
     *
     * @generated from field: metalstack.api.v2.FilesystemLayout filesystem_layout = 1;
     */
    filesystemLayout?: FilesystemLayout;
};
/**
 * Describes the message metalstack.api.v2.FilesystemServiceMatchResponse.
 * Use `create(FilesystemServiceMatchResponseSchema)` to create a new message.
 */
export declare const FilesystemServiceMatchResponseSchema: GenMessage<FilesystemServiceMatchResponse>;
/**
 * FilesystemLayout
 *
 * @generated from message metalstack.api.v2.FilesystemLayout
 */
export type FilesystemLayout = Message<"metalstack.api.v2.FilesystemLayout"> & {
    /**
     * Id of this filesystemLayout
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * Meta for this filesystemLayout
     *
     * @generated from field: metalstack.api.v2.Meta meta = 2;
     */
    meta?: Meta;
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
 * Describes the message metalstack.api.v2.FilesystemLayout.
 * Use `create(FilesystemLayoutSchema)` to create a new message.
 */
export declare const FilesystemLayoutSchema: GenMessage<FilesystemLayout>;
/**
 * FilesystemLayoutConstraints
 *
 * @generated from message metalstack.api.v2.FilesystemLayoutConstraints
 */
export type FilesystemLayoutConstraints = Message<"metalstack.api.v2.FilesystemLayoutConstraints"> & {
    /**
     * Sizes list of sizes this layout applies to
     *
     * @generated from field: repeated string sizes = 1;
     */
    sizes: string[];
    /**
     * Images list of images this layout applies to
     *
     * @generated from field: map<string, string> images = 2;
     */
    images: {
        [key: string]: string;
    };
};
/**
 * Describes the message metalstack.api.v2.FilesystemLayoutConstraints.
 * Use `create(FilesystemLayoutConstraintsSchema)` to create a new message.
 */
export declare const FilesystemLayoutConstraintsSchema: GenMessage<FilesystemLayoutConstraints>;
/**
 * Filesystem
 *
 * @generated from message metalstack.api.v2.Filesystem
 */
export type Filesystem = Message<"metalstack.api.v2.Filesystem"> & {
    /**
     * Device the underlying device where this filesystem should be created
     *
     * @generated from field: string device = 1;
     */
    device: string;
    /**
     * Format of the filesystem format
     *
     * @generated from field: metalstack.api.v2.Format format = 2;
     */
    format: Format;
    /**
     * Name of this filesystem
     *
     * @generated from field: optional string name = 3;
     */
    name?: string;
    /**
     * Description of this filesystem
     *
     * @generated from field: optional string description = 4;
     */
    description?: string;
    /**
     * Path the mountpoint where this filesystem should be mounted on
     *
     * @generated from field: optional string path = 5;
     */
    path?: string;
    /**
     * Label optional label for this this filesystem
     *
     * @generated from field: optional string label = 6;
     */
    label?: string;
    /**
     * MountOptions the options to use to mount this filesystem
     *
     * @generated from field: repeated string mount_options = 7;
     */
    mountOptions: string[];
    /**
     * CreateOptions the options to use to create (mkfs) this filesystem
     *
     * @generated from field: repeated string create_options = 8;
     */
    createOptions: string[];
};
/**
 * Describes the message metalstack.api.v2.Filesystem.
 * Use `create(FilesystemSchema)` to create a new message.
 */
export declare const FilesystemSchema: GenMessage<Filesystem>;
/**
 * Disk
 *
 * @generated from message metalstack.api.v2.Disk
 */
export type Disk = Message<"metalstack.api.v2.Disk"> & {
    /**
     * Device the device to create the partitions
     *
     * @generated from field: string device = 1;
     */
    device: string;
    /**
     * Partitions list of partitions to create on this disk
     *
     * @generated from field: repeated metalstack.api.v2.DiskPartition partitions = 2;
     */
    partitions: DiskPartition[];
};
/**
 * Describes the message metalstack.api.v2.Disk.
 * Use `create(DiskSchema)` to create a new message.
 */
export declare const DiskSchema: GenMessage<Disk>;
/**
 * Raid
 *
 * @generated from message metalstack.api.v2.Raid
 */
export type Raid = Message<"metalstack.api.v2.Raid"> & {
    /**
     * ArrayName the name of the resulting array device
     *
     * @generated from field: string array_name = 1;
     */
    arrayName: string;
    /**
     * Devices list of devices to form the raid array from
     *
     * @generated from field: repeated string devices = 2;
     */
    devices: string[];
    /**
     * Level raid level to create, should be 0 or 1
     *
     * @generated from field: metalstack.api.v2.RaidLevel level = 3;
     */
    level: RaidLevel;
    /**
     * CreateOptions the options to use to create the raid array
     *
     * @generated from field: repeated string create_options = 4;
     */
    createOptions: string[];
    /**
     * Spares number of spares for the raid array
     *
     * @generated from field: int32 spares = 5;
     */
    spares: number;
};
/**
 * Describes the message metalstack.api.v2.Raid.
 * Use `create(RaidSchema)` to create a new message.
 */
export declare const RaidSchema: GenMessage<Raid>;
/**
 * DiskPartition
 *
 * @generated from message metalstack.api.v2.DiskPartition
 */
export type DiskPartition = Message<"metalstack.api.v2.DiskPartition"> & {
    /**
     * Number partition number, will be appended to partitionprefix to create the final devicename
     *
     * @generated from field: uint32 number = 1;
     */
    number: number;
    /**
     * optional label for this this partition
     *
     * @generated from field: optional string label = 2;
     */
    label?: string;
    /**
     * Size size in mebibytes (MiB) of this partition"
     *
     * @generated from field: uint64 size = 3;
     */
    size: bigint;
    /**
     * GPTType the gpt partition table type of this partition
     *
     * @generated from field: optional metalstack.api.v2.GPTType gpt_type = 4;
     */
    gptType?: GPTType;
};
/**
 * Describes the message metalstack.api.v2.DiskPartition.
 * Use `create(DiskPartitionSchema)` to create a new message.
 */
export declare const DiskPartitionSchema: GenMessage<DiskPartition>;
/**
 * VolumeGroup
 *
 * @generated from message metalstack.api.v2.VolumeGroup
 */
export type VolumeGroup = Message<"metalstack.api.v2.VolumeGroup"> & {
    /**
     * Name the name of the resulting volume group
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * Devices list of devices to form the volume group from
     *
     * @generated from field: repeated string devices = 2;
     */
    devices: string[];
    /**
     * Tags list of tags to add to the volume group
     *
     * @generated from field: repeated string tags = 3;
     */
    tags: string[];
};
/**
 * Describes the message metalstack.api.v2.VolumeGroup.
 * Use `create(VolumeGroupSchema)` to create a new message.
 */
export declare const VolumeGroupSchema: GenMessage<VolumeGroup>;
/**
 * LogicalVolume
 *
 * @generated from message metalstack.api.v2.LogicalVolume
 */
export type LogicalVolume = Message<"metalstack.api.v2.LogicalVolume"> & {
    /**
     * Name the name of the logical volume
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * VolumeGroup the name of the volume group where to create the logical volume onto
     *
     * @generated from field: string volume_group = 2;
     */
    volumeGroup: string;
    /**
     * Size size in mebibytes (MiB) of this volume
     *
     * @generated from field: uint64 size = 3;
     */
    size: bigint;
    /**
     * LVMType the type of this logical volume can be either linear|striped|raid1
     *
     * @generated from field: metalstack.api.v2.LVMType lvm_type = 4;
     */
    lvmType: LVMType;
};
/**
 * Describes the message metalstack.api.v2.LogicalVolume.
 * Use `create(LogicalVolumeSchema)` to create a new message.
 */
export declare const LogicalVolumeSchema: GenMessage<LogicalVolume>;
/**
 * LVMType
 *
 * @generated from enum metalstack.api.v2.LVMType
 */
export declare enum LVMType {
    /**
     * LVM_TYPE_UNSPECIFIED is not specified
     *
     * @generated from enum value: LVM_TYPE_UNSPECIFIED = 0;
     */
    LVM_TYPE_UNSPECIFIED = 0,
    /**
     * LVM_TYPE_LINEAR append across all physical volumes
     *
     * @generated from enum value: LVM_TYPE_LINEAR = 1;
     */
    LVM_TYPE_LINEAR = 1,
    /**
     * LVM_TYPE_STRIPED stripe across all physical volumes
     *
     * @generated from enum value: LVM_TYPE_STRIPED = 2;
     */
    LVM_TYPE_STRIPED = 2,
    /**
     * LVM_TYPE_RAID1 mirror with raid across all physical volumes
     *
     * @generated from enum value: LVM_TYPE_RAID1 = 3;
     */
    LVM_TYPE_RAID1 = 3
}
/**
 * Describes the enum metalstack.api.v2.LVMType.
 */
export declare const LVMTypeSchema: GenEnum<LVMType>;
/**
 * Format specifies the filesystem to use on a volume
 *
 * @generated from enum metalstack.api.v2.Format
 */
export declare enum Format {
    /**
     * FORMAT_UNSPECIFIED
     *
     * @generated from enum value: FORMAT_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * FORMAT_VFAT is used for the UEFI boot partition
     *
     * @generated from enum value: FORMAT_VFAT = 1;
     */
    VFAT = 1,
    /**
     * FORMAT_EXT3 is usually only used for /boot
     *
     * @generated from enum value: FORMAT_EXT3 = 2;
     */
    EXT3 = 2,
    /**
     * FORMAT_EXT4 is the default fs
     *
     * @generated from enum value: FORMAT_EXT4 = 3;
     */
    EXT4 = 3,
    /**
     * FORMAT_SWAP is for the swap partition
     *
     * @generated from enum value: FORMAT_SWAP = 4;
     */
    SWAP = 4,
    /**
     * FORMAT_TMPFS is used for a memory filesystem typically /tmp
     *
     * @generated from enum value: FORMAT_TMPFS = 5;
     */
    TMPFS = 5,
    /**
     * FORMAT_NONE
     *
     * @generated from enum value: FORMAT_NONE = 6;
     */
    NONE = 6
}
/**
 * Describes the enum metalstack.api.v2.Format.
 */
export declare const FormatSchema: GenEnum<Format>;
/**
 * GPTType specifies the partition type in uefi systems
 *
 * @generated from enum metalstack.api.v2.GPTType
 */
export declare enum GPTType {
    /**
     * GPT_TYPE_UNSPECIFIED is no specified
     *
     * @generated from enum value: GPT_TYPE_UNSPECIFIED = 0;
     */
    GPT_TYPE_UNSPECIFIED = 0,
    /**
     * GPT_TYPE_BOOT EFI Boot Partition
     *
     * @generated from enum value: GPT_TYPE_BOOT = 1;
     */
    GPT_TYPE_BOOT = 1,
    /**
     * GPT_TYPE_LINUX Linux Partition
     *
     * @generated from enum value: GPT_TYPE_LINUX = 2;
     */
    GPT_TYPE_LINUX = 2,
    /**
     * GPT_TYPE_LINUX_RAID Linux Raid Partition
     *
     * @generated from enum value: GPT_TYPE_LINUX_RAID = 3;
     */
    GPT_TYPE_LINUX_RAID = 3,
    /**
     * GPT_TYPE_LINUX_LVM Linux LVM Partition
     *
     * @generated from enum value: GPT_TYPE_LINUX_LVM = 4;
     */
    GPT_TYPE_LINUX_LVM = 4
}
/**
 * Describes the enum metalstack.api.v2.GPTType.
 */
export declare const GPTTypeSchema: GenEnum<GPTType>;
/**
 * RaidLevel defines howto mirror two or more block devices
 *
 * @generated from enum metalstack.api.v2.RaidLevel
 */
export declare enum RaidLevel {
    /**
     * RAID_LEVEL_UNSPECIFIED is not specified
     *
     * @generated from enum value: RAID_LEVEL_UNSPECIFIED = 0;
     */
    RAID_LEVEL_UNSPECIFIED = 0,
    /**
     * RAID_LEVEL_0 is a stripe of two or more disks
     *
     * @generated from enum value: RAID_LEVEL_0 = 1;
     */
    RAID_LEVEL_0 = 1,
    /**
     * RAID_LEVEL_1 is a mirror of two disks
     *
     * @generated from enum value: RAID_LEVEL_1 = 2;
     */
    RAID_LEVEL_1 = 2
}
/**
 * Describes the enum metalstack.api.v2.RaidLevel.
 */
export declare const RaidLevelSchema: GenEnum<RaidLevel>;
/**
 * FilesystemService serves filesystem related functions
 *
 * @generated from service metalstack.api.v2.FilesystemService
 */
export declare const FilesystemService: GenService<{
    /**
     * Get a filesystem
     *
     * @generated from rpc metalstack.api.v2.FilesystemService.Get
     */
    get: {
        methodKind: "unary";
        input: typeof FilesystemServiceGetRequestSchema;
        output: typeof FilesystemServiceGetResponseSchema;
    };
    /**
     * List all filesystems
     *
     * @generated from rpc metalstack.api.v2.FilesystemService.List
     */
    list: {
        methodKind: "unary";
        input: typeof FilesystemServiceListRequestSchema;
        output: typeof FilesystemServiceListResponseSchema;
    };
    /**
     * Match a filesystems
     *
     * @generated from rpc metalstack.api.v2.FilesystemService.Match
     */
    match: {
        methodKind: "unary";
        input: typeof FilesystemServiceMatchRequestSchema;
        output: typeof FilesystemServiceMatchResponseSchema;
    };
}>;
