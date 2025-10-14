package validation

import (
	"testing"

	infrav2 "github.com/metal-stack/api/go/metalstack/infra/v2"
)

func TestValidateBoot(t *testing.T) {
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
 - mac: this string must be a valid macaddress [string.macaddress]`,
		},
	}
	validateProtos(t, tests)
}
