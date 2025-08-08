package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateFirewallRules(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

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
			wantErrorMessage: `validation error:
 - protocol: value must be one of the defined enum values [enum.defined_only]
 - ports[0]: value must be less than or equal to 65532 [uint32.lte]
 - to[0]: to prefixes must be valid [valid_to]`,
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
			wantErrorMessage: "validation error:\n - comment: value does not match regex pattern `^[a-z_ -]*$` [string.pattern]",
		},
	}

	validateProtos(t, tests, validator)
}
