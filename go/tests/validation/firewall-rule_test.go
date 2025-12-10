package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateFirewallRules(t *testing.T) {
	tests := prototests{
		{
			name: "Invalid Rule with valid comment",
			msg: &apiv2.FirewallEgressRule{
				Protocol: 5,
				Ports:    []uint32{100000},
				To:       []string{"a.192"},
				Comment:  "a rule",
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - protocol: value must be one of the defined enum values
 - ports[0]: value must be less than or equal to 65532
 - to: given prefixes must be valid`,
		},
		{
			name: "Invalid Rule with invalid comment",
			msg: &apiv2.FirewallEgressRule{
				Protocol: apiv2.IPProtocol_IP_PROTOCOL_TCP,
				Ports:    []uint32{1200},
				To:       []string{"1.2.3.4/32"},
				Comment:  "a # invalid 3 rule",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: comment: value does not match regex pattern `^[a-z_ -]*$`",
		},
	}

	validateProtos(t, tests)
}
