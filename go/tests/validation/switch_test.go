package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateMACAddress(t *testing.T) {
	tests := prototests{
		{
			name: "Valid MAC address",
			msg: &apiv2.SwitchNic{
				Mac:        "11:22:33:44:55:66",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr: false,
		},
		{
			name: "Too long",
			msg: &apiv2.SwitchNic{
				Mac:        "11:22:33:44:55:66:77",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac: this string must be a valid macaddress",
		},
		// FIXME does not fail anymore
		// {
		// 	name: "Too short",
		// 	msg: &apiv2.SwitchNic{
		// 		Mac:        "11:22:33:44:55",
		// 		Name:       "Ethernet0",
		// 		Identifier: "Eth1/1",
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "SwitchNic with valid uppercase MAC",
			msg: &apiv2.SwitchNic{
				Name:       "eth0",
				Identifier: "swp1",
				Mac:        "00:80:41:AE:FD:7E",
				Vrf:        new("10"),
			},
			wantErr: false,
		},
		{
			name: "Switch with valid id",
			msg: &apiv2.Switch{
				Id:             "leaf01",
				Partition:      "p1",
				ManagementIp:   "1.2.3.4",
				ManagementUser: new("admin"),
				ConsoleCommand: new("ssh"),
			},
			wantErr: false,
		},
		{
			name: "Switch with invalid id",
			msg: &apiv2.Switch{
				Id:             "_1",
				Partition:      "p1",
				ManagementIp:   "1.2.3.4",
				ManagementUser: new("admin"),
				ConsoleCommand: new("ssh"),
			},
			wantErr:          true,
			wantErrorMessage: "validation error: id: value must be a valid hostname",
		},
		{
			name: "Invalid separator",
			msg: &apiv2.SwitchNic{
				Mac:        "11.22.33.44.55.66",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac: this string must be a valid macaddress",
		},
		{
			name: "Invalid character",
			msg: &apiv2.SwitchNic{
				Mac:        "11:22:33:44:55:gg",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac: this string must be a valid macaddress",
		},
		{
			name: "Uppercase and lowercase allowed",
			msg: &apiv2.SwitchNic{
				Mac:        "11:22:33:44:55:aA",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: mac: this string must be a valid macaddress",
		},
	}

	validateProtos(t, tests)
}
