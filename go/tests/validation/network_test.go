package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestValidateNetwork(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid NetworkCreateRequest minimal config",
			msg: &apiv2.NetworkServiceCreateRequest{
				Id:       proto.String("internet"),
				Options:  &apiv2.NetworkOptions{Shared: true},
				Prefixes: []string{"1.2.0.0/16"},
			},
			wantErr: false,
		},
		{
			name: "InValid NetworkCreateRequest prefixes malformed",
			msg: &apiv2.NetworkServiceCreateRequest{
				Id:                  proto.String("internet"),
				Options:             &apiv2.NetworkOptions{Shared: true},
				Prefixes:            []string{"1.2.3.4.5/99"},
				DestinationPrefixes: []string{"0.0.0.0.0/0"},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - given prefixes must be valid [prefixes]
 - given destination_prefixes must be valid [destination_prefixes]`,
		},
	}
	validateProtos(t, tests, validator)
}
