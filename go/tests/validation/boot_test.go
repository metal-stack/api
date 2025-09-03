package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	infrav2 "github.com/metal-stack/api/go/metalstack/infra/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateBoot(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid BootServiceBootRequest",
			msg: &infrav2.BootServiceBootRequest{
				Partition: "partition-1",
				Mac:       "00:00:00:00:00:01",
			},
			wantErr: false,
		},
		{
			name: "InValid BootServiceBootRequest",
			msg: &infrav2.BootServiceBootRequest{
				Partition: "p1",
				Mac:       "X0-00:00:00:00:00:01",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - mac: mac must be a valid mac address [mac]`,
		},
	}
	validateProtos(t, tests, validator)
}
