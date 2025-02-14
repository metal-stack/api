package validation

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/metal-stack/api/go/enum"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateFilesystem(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

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
	}

	validateProtos(t, tests, validator)
}

func TestGetStringValue(t *testing.T) {

	value, err := enum.GetFormatStringValue(apiv2.Format_FORMAT_EXT4)
	require.NoError(t, err)
	require.Equal(t, "ext4", value)

	_, err = enum.GetFormatStringValue(apiv2.Format_FORMAT_UNSPECIFIED)
	require.Error(t, err)
	require.EqualError(t, err, "unable to fetch stringvalue from FORMAT_UNSPECIFIED")

	value, err = enum.GetLVMTypeStringValue(apiv2.LVMType_LVM_TYPE_LINEAR)
	require.NoError(t, err)
	require.Equal(t, "linear", value)

	value, err = enum.GetGPTTypeStringValue(apiv2.GPTType_GPT_TYPE_BOOT)
	require.NoError(t, err)
	require.Equal(t, "ef00", value)
}
