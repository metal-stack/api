package validation

import (
	"testing"

	"github.com/metal-stack/api/go/enum"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
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
			name: "Invalid Filesystem, device to short, format required",
			msg: &apiv2.Filesystem{
				Device: "/",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - device: value length must be at least 2 characters [string.min_len]
 - format: value is required [required]`,
		},
		{
			name: "Invalid Filesystem, device to short, format invalid",
			msg: &apiv2.Filesystem{
				Device: "/dev/sda3",
				Format: apiv2.Format(99),
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - format: value must be one of the defined enum values [enum.defined_only]`,
		},
		{
			name: "Valid FilesystemLayout minimal config",
			msg: &apiv2.FilesystemLayout{
				Id: "c1-large",
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem minimal config with valid name and description",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        proto.String("c1-large"),
				Description: proto.String("c1-large"),
			},
			wantErr: false,
		},
		{
			name: "Valid Filesystem minimal config with invalid name",
			msg: &apiv2.FilesystemLayout{
				Id:          "c1-large",
				Name:        proto.String("c"),
				Description: proto.String("c1-large"),
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - name: must be within 2 and 128 characters [string.is_name]`},
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
