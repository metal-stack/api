package validation

import (
	"testing"

	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateNetwork(t *testing.T) {
	tests := prototests{
		{
			name: "Valid NetworkCreateRequest minimal config",
			msg: &adminv2.NetworkServiceCreateRequest{
				Id:       new("internet"),
				NatType:  apiv2.NATType_NAT_TYPE_IPV4_MASQUERADE.Enum(),
				Prefixes: []string{"1.2.0.0/16"},
			},
			wantErr: false,
		},
		{
			name: "InValid NetworkCreateRequest prefixes malformed",
			msg: &adminv2.NetworkServiceCreateRequest{
				Id:                  new("internet"),
				NatType:             apiv2.NATType_NAT_TYPE_IPV4_MASQUERADE.Enum(),
				Prefixes:            []string{"1.2.3.4.5/99"},
				DestinationPrefixes: []string{"0.0.0.0.0/0"},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - prefixes: given prefixes must be valid
 - destination_prefixes: given prefixes must be valid`,
		},
		{
			name: "Valid Network minimal config",
			msg: &apiv2.Network{
				Id:       "internet",
				Prefixes: []string{"1.2.0.0/16"},
			},
			wantErr: false,
		},
		{
			name: "InValid Network minimal config",
			msg: &apiv2.Network{
				Id:                  "internet",
				Prefixes:            []string{"1.2.3.4.5/99"},
				DestinationPrefixes: []string{"0.0.0.0.0/0"}},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - prefixes: given prefixes must be valid
 - destination_prefixes: given prefixes must be valid`,
		},
	}
	validateProtos(t, tests)
}
