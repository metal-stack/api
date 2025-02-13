package validation

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	apiv1 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateFilesystem(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid Filesystem minimal config",
			msg: &apiv1.Filesystem{
				Device: "/dev/sda3",
				Format: apiv1.Format_FORMAT_EXT4,
			},
			wantErr: false,
		},
		{
			name: "Invalid Filesystem, device to short, format required",
			msg: &apiv1.Filesystem{
				Device: "/",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - device: value length must be at least 2 characters [string.min_len]
 - format: value is required [required]`,
		},
		{
			name: "Invalid Filesystem, device to short, format invalid",
			msg: &apiv1.Filesystem{
				Device: "/dev/sda3",
				Format: apiv1.Format(99),
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - format: value must be one of the defined enum values [enum.defined_only]`,
		},
	}

	validateProtos(t, tests, validator)
}
