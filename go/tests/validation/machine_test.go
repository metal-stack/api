package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateMachine(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "MachineNicQuery with invalid MACs",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"abc"},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - macs[0]: this string must be a valid macaddress [string.macaddress]`,
		},
		{
			name: "MachineNicQuery with invalid and valid MACs mixed",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"abc", "00:80:41:ae:fd:7e"},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - macs[0]: this string must be a valid macaddress [string.macaddress]`,
		},
		{
			name: "MachineNicQuery with valid MACs mixed",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"00:80:41:ae:fd:7e", "00:80:41:ae:fd:7f"},
			},
			wantErr: false,
		},
	}
	validateProtos(t, tests, validator)
}
