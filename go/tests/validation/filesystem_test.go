package validation

import (
	"testing"

	"github.com/metal-stack/api/go/enum"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateFilesystem(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Filesystem minimal config",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with all optional fields set",
			msg: &apiv2.Filesystem{
				Device:        "/dev/sda3",
				Format:        apiv2.Format_FORMAT_EXT4,
				Name:          new("rootfs"),
				Description:   new("root filesystem for production servers"),
				Path:          new("/"),
				Label:         new("root-label"),
				MountOptions:  []string{"noatime", "nodiratime"},
				CreateOptions: []string{"E", "lazy_itable_init"},
			},
			wantErr: false,
		},
		{
			name: "Invalid Filesystem, device too short, format required",
			msg: &apiv2.Filesystem{
				Device: "/",
			},
			wantErr:          true,
			wantErrorMessage: "validation errors:\n - device: must be within 2 and 128 characters\n - format: value is required",
		},
		{
			name: "Invalid Filesystem, device too short with invalid format",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format(99),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: format: value must be one of the defined enum values",
		},
		{
			name: "Invalid Filesystem, device with leading whitespace",
			msg: &apiv2.Filesystem{
				Device: " /dev/sda",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem, device with trailing whitespace",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda ",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem, device too long (129 chars)",
			msg: &apiv2.Filesystem{
				Device: createString(129),
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem, path too long",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new(createString(4097)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: path: must be at most 4096 characters",
		},
		{
			name: "Invalid Filesystem, path empty string",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new(""),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: path: must be at least 1 characters",
		},
		{
			name: "Invalid Filesystem, label with leading whitespace",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(" leading-label"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem, label with trailing whitespace",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("trailing-label "),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem, label too short",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("a"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem, label too long",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(createString(129)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem, label both with whitespace",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(" label "),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem, mount_options has duplicates",
			msg: &apiv2.Filesystem{
				Device:       "/dev/sda3",
				Format:       apiv2.Format_FORMAT_EXT4,
				MountOptions: []string{"noatime", "noatime"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mount_options: repeated value must contain unique items",
		},
		{
			name: "Invalid Filesystem, mount_options item too long",
			msg: &apiv2.Filesystem{
				Device:       "/dev/sda3",
				Format:       apiv2.Format_FORMAT_EXT4,
				MountOptions: []string{createString(129)},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mount_options[0]: must be within 1 and 128 characters",
		},
		{
			name: "Invalid Filesystem, mount_options item too short",
			msg: &apiv2.Filesystem{
				Device:       "/dev/sda3",
				Format:       apiv2.Format_FORMAT_EXT4,
				MountOptions: []string{""},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mount_options[0]: must be within 1 and 128 characters",
		},
		{
			name: "Invalid Filesystem, create_options has duplicates",
			msg: &apiv2.Filesystem{
				Device:        "/dev/sda3",
				Format:        apiv2.Format_FORMAT_EXT4,
				CreateOptions: []string{"E", "E"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: create_options: repeated value must contain unique items",
		},
		{
			name: "Invalid Filesystem, create_options item too long",
			msg: &apiv2.Filesystem{
				Device:        "/dev/sda3",
				Format:        apiv2.Format_FORMAT_EXT4,
				CreateOptions: []string{createString(129)},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: create_options[0]: must be within 1 and 128 characters",
		},
		{
			name: "Invalid Filesystem, create_options item too short",
			msg: &apiv2.Filesystem{
				Device:        "/dev/sda3",
				Format:        apiv2.Format_FORMAT_EXT4,
				CreateOptions: []string{""},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: create_options[0]: must be within 1 and 128 characters",
		},
		{
			name: "Valid FilesystemLayout minimal config",
			msg: &apiv2.FilesystemLayout{
				Id: "c1-large",
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemLayout full config",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        new("C1 Large Machine"),
				Description: new("Large machine with 64GB RAM and 4 CPUs"),
				Filesystems: []*apiv2.Filesystem{
					{
						Device: "/dev/sda1",
						Format: apiv2.Format_FORMAT_VFAT,
						Path:   new("/boot/efi"),
						Label:  new("efi"),
					},
				},
				Disks: []*apiv2.Disk{
					{
						Device: "/dev/sda",
						Partitions: []*apiv2.DiskPartition{
							{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT)},
							{Number: 2, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
						},
					},
				},
				Raid: []*apiv2.Raid{
					{
						ArrayName: "md0",
						Devices:   []string{"/dev/sda", "/dev/sdb"},
						Level:     apiv2.RaidLevel_RAID_LEVEL_1,
					},
				},
				VolumeGroups: []*apiv2.VolumeGroup{
					{
						Name:    "vg0",
						Devices: []string{"/dev/md0"},
					},
				},
				LogicalVolumes: []*apiv2.LogicalVolume{
					{
						Name:        "lv-root",
						VolumeGroup: "vg0",
						Size:        102400,
						LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
					},
				},
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Sizes: []string{"c1-large", "c2-large"},
					Images: map[string]string{
						"debian": ">= 12.0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemLayout empty name and description is ok",
			msg: &apiv2.FilesystemLayout{
				Id:          "l1-small",
				Description: new(""),
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemLayout with empty mount and create options",
			msg: &apiv2.Filesystem{
				Device:        "/dev/sda3",
				Format:        apiv2.Format_FORMAT_EXT4,
				MountOptions:  []string{},
				CreateOptions: []string{},
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 2 char name boundary",
			msg: &apiv2.FilesystemLayout{
				Id:   "l1-medium",
				Name: new("ab"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 2 char label boundary",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("ab"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 128 char label boundary",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(createString(128)),
			},
			wantErr: false,
		},
		{
			name: "Invalid FilesystemLayout, name too short",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        new("c"),
				Description: new("c1-large"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid FilesystemLayout, name too long",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        new(createString(129)),
				Description: new("c1-large"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid FilesystemLayout, description too long",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Description: new(createString(257)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: description: must be shorter than 256 characters",
		},
		{
			name: "Valid FilesystemLayout, description at max length (256 chars)",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Description: new(createString(256)),
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemServiceGetRequest(t *testing.T) {
	tests := prototests{
		{
			name: "Valid FilesystemServiceGetRequest",
			msg: &apiv2.FilesystemServiceGetRequest{
				Id: "c1-large",
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemServiceGetRequest with 128 char id",
			msg: &apiv2.FilesystemServiceGetRequest{
				Id: createString(128),
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemServiceGetRequest with 2 char id minimum",
			msg: &apiv2.FilesystemServiceGetRequest{
				Id: "ab",
			},
			wantErr: false,
		},
		{
			name: "Invalid FilesystemServiceGetRequest, id too short",
			msg: &apiv2.FilesystemServiceGetRequest{
				Id: "a",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: id: must be within 2 and 128 characters",
		},
		{
			name: "Invalid FilesystemServiceGetRequest, id too long",
			msg: &apiv2.FilesystemServiceGetRequest{
				Id: createString(129),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: id: must be within 2 and 128 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemServiceListRequest(t *testing.T) {
	tests := prototests{
		{
			name: "Valid FilesystemServiceListRequest with nil id",
			msg: &apiv2.FilesystemServiceListRequest{
				Id: nil,
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemServiceListRequest with valid id",
			msg: &apiv2.FilesystemServiceListRequest{
				Id: new("c1-large"),
			},
			wantErr: false,
		},
		{
			name: "Invalid FilesystemServiceListRequest, id too short",
			msg: &apiv2.FilesystemServiceListRequest{
				Id: new("x"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: id: must be within 2 and 128 characters",
		},
		{
			name: "Invalid FilesystemServiceListRequest, id too long",
			msg: &apiv2.FilesystemServiceListRequest{
				Id: new(createString(129)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: id: must be within 2 and 128 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateDisk(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Disk minimal config",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
			},
			wantErr: false,
		},
		{
			name: "Valid Disk with multiple partitions",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT)},
					{Number: 2, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Disk with partition label",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_BOOT),
						Label:   new("boot-efi"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Disk with partition label at max length",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_LINUX),
						Label:   new(createString(128)),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Disk, device with leading whitespace",
			msg: &apiv2.Disk{
				Device: " /dev/sda",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: value must not start or end with whitespace",
		},
		{
			name: "Invalid Disk, device with trailing whitespace",
			msg: &apiv2.Disk{
				Device: "/dev/sda ",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: value must not start or end with whitespace",
		},
		{
			name: "Invalid Disk, device too short",
			msg: &apiv2.Disk{
				Device: "/",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Disk, device too long",
			msg: &apiv2.Disk{
				Device: createString(129),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Disk, partition gpt_type unspecified",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_UNSPECIFIED)},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: gpt_type: value must be one of the defined enum values",
		},
		{
			name: "Invalid Disk, partition label with leading whitespace",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_LINUX),
						Label:   new(" label"),
					},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: partitions[0].label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Disk, partition label with trailing whitespace",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_LINUX),
						Label:   new("label "),
					},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: partitions[0].label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Disk, partition label too short",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_LINUX),
						Label:   new("a"),
					},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: partitions[0].label: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Disk, partition label too long",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{
						Number:  1,
						Size:    512,
						GptType: new(apiv2.GPTType_GPT_TYPE_LINUX),
						Label:   new(createString(129)),
					},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: partitions[0].label: must be within 2 and 128 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateRaid(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Raid minimal config",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr: false,
		},
		{
			name: "Valid Raid with create_options and spares",
			msg: &apiv2.Raid{
				ArrayName:     "md0",
				Devices:       []string{"/dev/sda", "/dev/sdb", "/dev/sdc"},
				Level:         apiv2.RaidLevel_RAID_LEVEL_0,
				CreateOptions: []string{"--chunk=512", "--level=0"},
				Spares:        1,
			},
			wantErr: false,
		},
		{
			name: "Valid Raid with exactly 128 devices",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   createDiskDevices(128),
				Level:     apiv2.RaidLevel_RAID_LEVEL_0,
			},
			wantErr: false,
		},
		{
			name: "Valid Raid with empty create_options",
			msg: &apiv2.Raid{
				ArrayName:     "md0",
				Devices:       []string{"/dev/sda", "/dev/sdb"},
				Level:         apiv2.RaidLevel_RAID_LEVEL_1,
				CreateOptions: []string{},
				Spares:        0,
			},
			wantErr: false,
		},
		{
			name: "Invalid Raid, array_name with leading whitespace",
			msg: &apiv2.Raid{
				ArrayName: " md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: array_name: value must not start or end with whitespace",
		},
		{
			name: "Invalid Raid, array_name with trailing whitespace",
			msg: &apiv2.Raid{
				ArrayName: "md0 ",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: array_name: value must not start or end with whitespace",
		},
		{
			name: "Invalid Raid, array_name too short",
			msg: &apiv2.Raid{
				ArrayName: "m",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: array_name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Raid, array_name too long",
			msg: &apiv2.Raid{
				ArrayName: createString(129),
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: array_name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Raid, level unspecified",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_UNSPECIFIED,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: level: value is required",
		},
		{
			name: "Invalid Raid, devices with leading whitespace",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda ", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: given values must not start or end with whitespace",
		},
		{
			name: "Invalid Raid, devices with trailing whitespace",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb "},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: given values must not start or end with whitespace",
		},
		{
			name: "Invalid Raid, devices too many",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   createDiskDevices(129),
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: must contain no more than 128 item(s)",
		},
		{
			name: "Invalid Raid, devices not unique",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sda"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: must contain unique values",
		},
		{
			name: "Invalid Raid, create_options duplicates",
			msg: &apiv2.Raid{
				ArrayName:     "md0",
				Devices:       []string{"/dev/sda", "/dev/sdb"},
				Level:         apiv2.RaidLevel_RAID_LEVEL_1,
				CreateOptions: []string{"--chunk=512", "--chunk=512"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: create_options: repeated value must contain unique items",
		},
	}

	validateProtos(t, tests)
}

func TestValidateVolumeGroup(t *testing.T) {
	tests := prototests{
		{
			name: "Valid VolumeGroup minimal config",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1", "/dev/sdb1"},
			},
			wantErr: false,
		},
		{
			name: "Valid VolumeGroup with tags",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1", "/dev/sdb1"},
				Tags:    []string{"prod", "storage-fast"},
			},
			wantErr: false,
		},
		{
			name: "Valid VolumeGroup name max length",
			msg: &apiv2.VolumeGroup{
				Name:    createString(128),
				Devices: []string{"/dev/sda1"},
			},
			wantErr: false,
		},
		{
			name: "Valid VolumeGroup with empty devices",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{},
			},
			wantErr: false,
		},
		{
			name: "Valid VolumeGroup with empty tags",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1"},
				Tags:    []string{},
			},
			wantErr: false,
		},
		{
			name: "Invalid VolumeGroup, name too short",
			msg: &apiv2.VolumeGroup{
				Name:    "a",
				Devices: []string{"/dev/sda1"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid VolumeGroup, name too long",
			msg: &apiv2.VolumeGroup{
				Name:    createString(129),
				Devices: []string{"/dev/sda1"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid VolumeGroup, devices too many",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: createDiskDevices(129),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: must contain no more than 128 item(s)",
		},
		{
			name: "Invalid VolumeGroup, devices with whitespace",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1 ", "/dev/sdb1"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices: given values must not start or end with whitespace",
		},
		{
			name: "Invalid VolumeGroup, tags too many",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1"},
				Tags:    createStringSlice(129),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: tags: must contain no more than 128 item(s)",
		},
		{
			name: "Invalid VolumeGroup, tags with whitespace",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1"},
				Tags:    []string{" valid-tag", "prod"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: tags: given values must not start or end with whitespace",
		},
		{
			name: "Invalid VolumeGroup, device name too short",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/x"},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices[0]: must be within 2 and 128 characters",
		},
		{
			name: "Invalid VolumeGroup, device name too long",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{createString(129)},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: devices[0]: must be within 2 and 128 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateLogicalVolume(t *testing.T) {
	tests := prototests{
		{
			name: "Valid LogicalVolume with linear type",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        102400,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr: false,
		},
		{
			name: "Valid LogicalVolume with striped type",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-data",
				VolumeGroup: "vg0",
				Size:        204800,
				LvmType:     apiv2.LVMType_LVM_TYPE_STRIPED,
			},
			wantErr: false,
		},
		{
			name: "Valid LogicalVolume with raid1 type",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-backup",
				VolumeGroup: "vg0",
				Size:        512000,
				LvmType:     apiv2.LVMType_LVM_TYPE_RAID1,
			},
			wantErr: false,
		},
		{
			name: "Valid LogicalVolume with zero size",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        0,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr: false,
		},
		{
			name: "Invalid LogicalVolume, name too short",
			msg: &apiv2.LogicalVolume{
				Name:        "x",
				VolumeGroup: "vg0",
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid LogicalVolume, name too long",
			msg: &apiv2.LogicalVolume{
				Name:        createString(129),
				VolumeGroup: "vg0",
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: name: must be within 2 and 128 characters",
		},
		{
			name: "Invalid LogicalVolume, volume_group too short",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "x",
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: volume_group: must be within 2 and 128 characters",
		},
		{
			name: "Invalid LogicalVolume, volume_group too long",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: createString(129),
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: volume_group: must be within 2 and 128 characters",
		},
		{
			name: "Invalid LogicalVolume, lvm_type unspecified",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_UNSPECIFIED,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: lvm_type: value is required",
		},
		{
			name: "Invalid LogicalVolume, lvm_type invalid value",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        1024,
				LvmType:     apiv2.LVMType(99),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: lvm_type: value must be one of the defined enum values",
		},
	}

	validateProtos(t, tests)
}

func TestValidateFormatEnum(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Format VFAT for UEFI",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda1",
				Format: apiv2.Format_FORMAT_VFAT,
			},
			wantErr: false,
		},
		{
			name: "Valid Format EXT3 for /boot",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda1",
				Format: apiv2.Format_FORMAT_EXT3,
			},
			wantErr: false,
		},
		{
			name: "Valid Format EXT4 default",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda1",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Format SWAP",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda2",
				Format: apiv2.Format_FORMAT_SWAP,
			},
			wantErr: false,
		},
		{
			name: "Valid Format TMPFS for /tmp",
			msg: &apiv2.Filesystem{
				Device: "/tmp",
				Format: apiv2.Format_FORMAT_TMPFS,
			},
			wantErr: false,
		},
		{
			name: "Valid Format NONE",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_NONE,
			},
			wantErr: false,
		},
		{
			name: "Invalid Format UNSPECIFIED",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda1",
				Format: apiv2.Format_FORMAT_UNSPECIFIED,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: format: value is required",
		},
		{
			name: "Invalid Format invalid value",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda1",
				Format: apiv2.Format(99),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: format: value must be one of the defined enum values",
		},
	}

	validateProtos(t, tests)
}

func TestValidateGPTTypeEnum(t *testing.T) {
	tests := prototests{
		{
			name: "Valid GPTType BOOT for EFI",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid GPTType LINUX",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid GPTType LINUX_RAID",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_RAID)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid GPTType LINUX_LVM",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_LVM)},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid GPTType UNSPECIFIED",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_UNSPECIFIED)},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: gpt_type: value must be one of the defined enum values",
		},
		{
			name: "Invalid GPTType invalid value",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType(99))},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: partitions[0].gpt_type: value must be one of the defined enum values",
		},
	}

	validateProtos(t, tests)
}

func TestValidateRaidLevelEnum(t *testing.T) {
	tests := prototests{
		{
			name: "Valid RaidLevel 0 for stripe",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_0,
			},
			wantErr: false,
		},
		{
			name: "Valid RaidLevel 1 for mirror",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
			},
			wantErr: false,
		},
		{
			name: "Invalid RaidLevel UNSPECIFIED",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_UNSPECIFIED,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: level: value is required",
		},
		{
			name: "Invalid RaidLevel invalid value",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel(99),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: level: value must be one of the defined enum values",
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemLayoutConstraints(t *testing.T) {
	tests := prototests{
		{
			name: "Valid constraints with sizes and images",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        new("C1 Large"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Sizes:  []string{"c1-large", "c2-large"},
					Images: map[string]string{"debian": ">= 12.0"},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with asterisk size to match all",
			msg: &apiv2.FilesystemLayout{
				Id:          "all-sizes",
				Name:        new("All Sizes"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Sizes:  []string{"*"},
					Images: map[string]string{"*": ">= 0.0"},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with only sizes",
			msg: &apiv2.FilesystemLayout{
				Id:          "sizes-only",
				Name:        new("Sizes Only"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Sizes: []string{"c1-small", "c1-medium", "c1-large"},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with only images",
			msg: &apiv2.FilesystemLayout{
				Id:          "images-only",
				Name:        new("Images Only"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Images: map[string]string{"debian": ">= 11", "ubuntu": ">= 22.04"},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints empty is valid",
			msg: &apiv2.FilesystemLayout{
				Id:          "empty-constraints",
				Name:        new("Empty Constraints"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with no constraints (nil)",
			msg: &apiv2.FilesystemLayout{
				Id:          "no-constraints",
				Name:        new("No Constraints"),
				Description: new(""),
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with many image entries",
			msg: &apiv2.FilesystemLayout{
				Id:          "multi-image",
				Name:        new("Multi Image"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Images: map[string]string{
						"debian":    ">= 11.0",
						"ubuntu":    ">= 20.04",
						"alpine":    ">= 3.16",
						"centos":    ">= 8.0",
						"rhel":      ">= 8.0",
						"rocky":     ">= 8.0",
						"almalinux": ">= 8.0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid constraints with complex semver expressions",
			msg: &apiv2.FilesystemLayout{
				Id:          "complex-images",
				Name:        new("Complex Images"),
				Description: new(""),
				Constraints: &apiv2.FilesystemLayoutConstraints{
					Images: map[string]string{
						"debian": ">= 12.0 < 14.0",
						"ubuntu": "~ 22.04",
					},
				},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateDiskPartition(t *testing.T) {
	tests := prototests{
		{
			name: "Valid DiskPartition with boot type",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with linux type",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 2, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with linux_raid type",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 3, Size: 51200, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_RAID)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with linux_lvm type",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 4, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_LVM)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with partition number 1",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 1024, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with large partition number",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 16, Size: 1024, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid DiskPartition with large size",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 9223372036854775807, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid DiskPartition with all unspecified types",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_UNSPECIFIED)},
					{Number: 2, Size: 1024, GptType: new(apiv2.GPTType_GPT_TYPE_UNSPECIFIED)},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: gpt_type: value must be one of the defined enum values",
		},
		{
			name: "Valid Disk with multiple partitions of different types",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT)},
					{Number: 2, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX)},
					{Number: 3, Size: 51200, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_RAID)},
					{Number: 4, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_LVM)},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid empty partitions list",
			msg: &apiv2.Disk{
				Device:     "/dev/sda",
				Partitions: []*apiv2.DiskPartition{},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateLVMType(t *testing.T) {
	tests := prototests{
		{
			name: "Valid LVMType_LINEAR for append across volumes",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        102400,
				LvmType:     apiv2.LVMType_LVM_TYPE_LINEAR,
			},
			wantErr: false,
		},
		{
			name: "Valid LVMType_STRIPED for striping",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-data",
				VolumeGroup: "vg0",
				Size:        204800,
				LvmType:     apiv2.LVMType_LVM_TYPE_STRIPED,
			},
			wantErr: false,
		},
		{
			name: "Valid LVMType_RAID1 for mirroring",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-backup",
				VolumeGroup: "vg0",
				Size:        512000,
				LvmType:     apiv2.LVMType_LVM_TYPE_RAID1,
			},
			wantErr: false,
		},
		{
			name: "Invalid LVMType_UNSPECIFIED",
			msg: &apiv2.LogicalVolume{
				Name:        "lv-root",
				VolumeGroup: "vg0",
				Size:        1024,
				LvmType:     apiv2.LVMType_LVM_TYPE_UNSPECIFIED,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: lvm_type: value must be one of the defined enum values",
		},
		{
			name: "Valid all LVM types on same volume group",
			msg: &apiv2.FilesystemLayout{
				Id:   "lvm-full",
				Name: new("Full LVM"),
				VolumeGroups: []*apiv2.VolumeGroup{
					{Name: "vg0", Devices: []string{"/dev/sda1", "/dev/sdb1"}},
				},
				LogicalVolumes: []*apiv2.LogicalVolume{
					{Name: "lv-linear", VolumeGroup: "vg0", Size: 1024, LvmType: apiv2.LVMType_LVM_TYPE_LINEAR},
					{Name: "lv-striped", VolumeGroup: "vg0", Size: 2048, LvmType: apiv2.LVMType_LVM_TYPE_STRIPED},
					{Name: "lv-raid1", VolumeGroup: "vg0", Size: 4096, LvmType: apiv2.LVMType_LVM_TYPE_RAID1},
				},
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateRaidSpares(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Raid with zero spares",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
				Spares:    0,
			},
			wantErr: false,
		},
		{
			name: "Valid Raid with spares set",
			msg: &apiv2.Raid{
				ArrayName: "md0",
				Devices:   []string{"/dev/sda", "/dev/sdb", "/dev/sdc"},
				Level:     apiv2.RaidLevel_RAID_LEVEL_1,
				Spares:    1,
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemLabels(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Filesystem with empty label",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(""),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: must be within 2 and 128 characters",
		},
		{
			name: "Valid Filesystem with valid label",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("my-label-01"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 2 char label minimum",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("ab"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 128 char label maximum",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(createString(128)),
			},
			wantErr: false,
		},
		{
			name: "Invalid Filesystem label with both leading and trailing whitespace",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(" label "),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: value must not start or end with whitespace",
		},
		{
			name: "Invalid Filesystem label too short single char",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new("a"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem label too long 129 chars",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Label:  new(createString(129)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: label: must be within 2 and 128 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemPath(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Filesystem with root path",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new("/"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with nested path",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new("/var/lib/docker"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with long path",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new(createString(4096)),
			},
			wantErr: false,
		},
		{
			name: "Invalid Filesystem path too long 4097 chars",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new(createString(4097)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: path: must be at most 4096 characters",
		},
		{
			name: "Invalid Filesystem path empty string",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
				Path:   new(""),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: path: must be at least 1 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateDescription(t *testing.T) {
	tests := prototests{
		{
			name: "Valid FilesystemLayout with short description",
			msg: &apiv2.FilesystemLayout{
				Id:          "l1-small",
				Name:        new("Small"),
				Description: new("Small machine"),
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemLayout with max length description",
			msg: &apiv2.FilesystemLayout{
				Id:          "l1-small",
				Name:        new("Small"),
				Description: new(createString(256)),
			},
			wantErr: false,
		},
		{
			name: "Valid FilesystemLayout with empty description",
			msg: &apiv2.FilesystemLayout{
				Id:          "l1-small",
				Name:        new("Small"),
				Description: new(""),
			},
			wantErr: false,
		},
		{
			name: "Invalid FilesystemLayout description too long",
			msg: &apiv2.FilesystemLayout{
				Id:          "l1-small",
				Name:        new("Small"),
				Description: new(createString(257)),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: description: must be shorter than 256 characters",
		},
	}

	validateProtos(t, tests)
}

func TestValidateFilesystemDevice(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Filesystem with standard device",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with nvme device",
			msg: &apiv2.Filesystem{
				Device: "/dev/nvme0n1p1",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with md device",
			msg: &apiv2.Filesystem{
				Device: "/dev/md0",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 2 char device minimum",
			msg: &apiv2.Filesystem{
				Device: "/d",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem with 128 char device maximum",
			msg: &apiv2.Filesystem{
				Device: createString(128),
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Invalid Filesystem device too short",
			msg: &apiv2.Filesystem{
				Device: "/",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem device too long",
			msg: &apiv2.Filesystem{
				Device: createString(129),
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: must be within 2 and 128 characters",
		},
		{
			name: "Invalid Filesystem device with whitespace",
			msg: &apiv2.Filesystem{
				Device: " /dev/sda ",
				Format: apiv2.Format_FORMAT_EXT4,
			},
			wantErr:          true,
			wantErrorMessage: "validation error: device: value must not start or end with whitespace",
		},
	}

	validateProtos(t, tests)
}

func TestValidateDiskWithPartitions(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Disk with multiple partitions having labels",
			msg: &apiv2.Disk{
				Device: "/dev/sda",
				Partitions: []*apiv2.DiskPartition{
					{Number: 1, Size: 512, GptType: new(apiv2.GPTType_GPT_TYPE_BOOT), Label: new("efi")},
					{Number: 2, Size: 102400, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX), Label: new("root")},
					{Number: 3, Size: 8192, GptType: new(apiv2.GPTType_GPT_TYPE_LINUX_LVM), Label: new("swap")},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid empty Disk without partitions",
			msg: &apiv2.Disk{
				Device:     "/dev/sda",
				Partitions: []*apiv2.DiskPartition{},
			},
			wantErr: false,
		},
		{
			name: "Valid nil Disk partitions",
			msg: &apiv2.Disk{
				Device:     "/dev/sda",
				Partitions: nil,
			},
			wantErr: false,
		},
	}

	validateProtos(t, tests)
}

func TestValidateVolumeGroupDevicesAndTags(t *testing.T) {
	tests := prototests{
		{
			name: "Valid VolumeGroup with devices at boundary 128",
			msg: &apiv2.VolumeGroup{
				Name:    "vg-max-devices",
				Devices: createDiskDevices(128),
				Tags:    []string{"tag1", "tag2"},
			},
			wantErr: false,
		},
		{
			name: "Valid VolumeGroup with max tags 128",
			msg: &apiv2.VolumeGroup{
				Name:    "vg-max-tags",
				Devices: []string{"/dev/sda1"},
				Tags:    createStringSlice(128),
			},
			wantErr: false,
		},
		{
			name: "Invalid VolumeGroup devices and tags both with whitespace",
			msg: &apiv2.VolumeGroup{
				Name:    "vg0",
				Devices: []string{"/dev/sda1 "},
				Tags:    []string{" tag "},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - devices: given values must not start or end with whitespace
 - tags: given values must not start or end with whitespace`,
		},
	}

	validateProtos(t, tests)
}

func TestGetStringValue(t *testing.T) {
	value, err := enum.GetStringValue(apiv2.Format_FORMAT_EXT4)
	require.NoError(t, err)
	ext4 := "ext4"
	require.Equal(t, &ext4, value)

	value, err = enum.GetStringValue(apiv2.LVMType_LVM_TYPE_LINEAR)
	require.NoError(t, err)
	linear := "linear"
	require.Equal(t, &linear, value)

	value, err = enum.GetStringValue(apiv2.GPTType_GPT_TYPE_BOOT)
	require.NoError(t, err)
	ef := "ef00"
	require.Equal(t, &ef, value)

	_, err = enum.GetStringValue(apiv2.Format_FORMAT_UNSPECIFIED)
	require.Error(t, err)
	require.EqualError(t, err, "unable to fetch stringvalue from FORMAT_UNSPECIFIED")
}
