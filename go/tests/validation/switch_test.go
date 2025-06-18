package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateMACAddress(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid MAC address",
			msg: &apiv2.SwitchNic{
				MacAddress: "11:22:33:44:55:66",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr: false,
		},
		{
			name: "Too long",
			msg: &apiv2.SwitchNic{
				MacAddress: "11:22:33:44:55:66:77",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error:\n - mac_address: value does not match regex pattern `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$` [string.pattern]",
		},
		{
			name: "Too short",
			msg: &apiv2.SwitchNic{
				MacAddress: "11:22:33:44:55",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error:\n - mac_address: value does not match regex pattern `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$` [string.pattern]",
		},
		{
			name: "Invalid separator",
			msg: &apiv2.SwitchNic{
				MacAddress: "11.22.33.44.55.66",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error:\n - mac_address: value does not match regex pattern `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$` [string.pattern]",
		},
		{
			name: "Invalid character",
			msg: &apiv2.SwitchNic{
				MacAddress: "11:22:33:44:55:gg",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error:\n - mac_address: value does not match regex pattern `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$` [string.pattern]",
		},
		{
			name: "Uppercase and lowercase allowed",
			msg: &apiv2.SwitchNic{
				MacAddress: "11:22:33:44:55:aA",
				Name:       "Ethernet0",
				Identifier: "Eth1/1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error:\n - mac_address: value does not match regex pattern `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$` [string.pattern]",
		},
	}

	validateProtos(t, tests, validator)
}
