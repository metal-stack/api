package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestValidateSwitch(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "SwitchNic with invalid MAC",
			msg: &apiv2.SwitchNic{
				Name:       "eth0",
				Identifier: "swp1",
				Mac:        "abc",
				Vrf:        proto.String("10"),
				Actual:     apiv2.SwitchPortStatus_SWITCH_PORT_STATUS_UP,
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - mac: this string must be a valid macaddress [string.macaddress]`,
		},
		{
			name: "SwitchNic with valid lowercase MAC",
			msg: &apiv2.SwitchNic{
				Name:       "eth0",
				Identifier: "swp1",
				Mac:        "00:80:41:ae:fd:7e",
				Vrf:        proto.String("10"),
				Actual:     apiv2.SwitchPortStatus_SWITCH_PORT_STATUS_UP,
			},
			wantErr: false,
		},
		{
			name: "SwitchNic with valid uppercase MAC",
			msg: &apiv2.SwitchNic{
				Name:       "eth0",
				Identifier: "swp1",
				Mac:        "00:80:41:AE:FD:7E",
				Vrf:        proto.String("10"),
				Actual:     apiv2.SwitchPortStatus_SWITCH_PORT_STATUS_UP,
			},
			wantErr: false,
		},
	}
	validateProtos(t, tests, validator)
}
