package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateMachine(t *testing.T) {
	tests := prototests{
		{
			name: "MachineNicQuery with invalid MACs",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"abc"},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: macs[0]: this string must be a valid macaddress`,
		},
		{
			name: "MachineNicQuery with invalid and valid MACs mixed",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"abc", "00:80:41:ae:fd:7e"},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: macs[0]: this string must be a valid macaddress`,
		},
		{
			name: "MachineNicQuery with valid MACs mixed",
			msg: &apiv2.MachineNicQuery{
				Macs: []string{"00:80:41:ae:fd:7e", "00:80:41:ae:fd:7f"},
			},
			wantErr: false,
		},
		{
			name: "MachineNetworkQuery with valid IPs",
			msg: &apiv2.MachineNetworkQuery{
				Ips: []string{"1.2.3.4"},
			},
			wantErr: false,
		},
		{
			name: "MachineNetworkQuery with invalid IPs",
			msg: &apiv2.MachineNetworkQuery{
				Ips: []string{"1.2.3.4.5"},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: ips: given ips must be valid`,
		},
	}
	validateProtos(t, tests)
}
