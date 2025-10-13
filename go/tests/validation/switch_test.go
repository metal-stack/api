package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"google.golang.org/protobuf/proto"
)

func TestValidateSwitch(t *testing.T) {
	tests := prototests{
		{
			name: "SwitchNic with invalid MAC",
			msg: &apiv2.SwitchNic{
				Name:       "eth0",
				Identifier: "swp1",
				Mac:        "abc",
				Vrf:        proto.String("10"),
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
			},
			wantErr: false,
		},
		{
			name: "Switch with valid id",
			msg: &apiv2.Switch{
				Id:             "leaf01",
				Partition:      "p1",
				ManagementIp:   "1.2.3.4",
				ManagementUser: proto.String("admin"),
				ConsoleCommand: proto.String("ssh"),
			},
			wantErr: false,
		},
		{
			name: "Switch with invalid id",
			msg: &apiv2.Switch{
				Id:             "_1",
				Partition:      "p1",
				ManagementIp:   "1.2.3.4",
				ManagementUser: proto.String("admin"),
				ConsoleCommand: proto.String("ssh"),
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - id: value must be a valid hostname [string.hostname]`,
		},
	}
	validateProtos(t, tests)
}
