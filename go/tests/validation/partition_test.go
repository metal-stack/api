package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateNTPServer(t *testing.T) {
	tests := prototests{
		{
			name: "Valid NTPServer with ip as address",
			msg: &apiv2.NTPServer{
				Address: "172.10.0.1",
			},
			wantErr: false,
		},
		{
			name: "Valid NTPServer with hostname as address",
			msg: &apiv2.NTPServer{
				Address: "ntp.company.com",
			},
			wantErr: false,
		},
		{
			name: "Invalid NTPServer address",
			msg: &apiv2.NTPServer{
				Address: "1.1.1.1.1",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: address: must be a valid IP address or hostname",
		},
		{
			name: "Invalid NTPServer address",
			msg: &apiv2.NTPServer{
				Address: "-bla.com#",
			},
			wantErr:          true,
			wantErrorMessage: "validation error: address: must be a valid IP address or hostname",
		},
	}
	validateProtos(t, tests)
}
