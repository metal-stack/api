package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateLabels(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Labels",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"color": "green",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Labels",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					" gradient ": " 34 ",
				},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - labels[" gradient "]: keys must be trimmed
 - labels[" gradient "]: values must be trimmed`,
		},
		{
			name: "Invalid Labels",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"color": "green ",
				},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: labels["color"]: values must be trimmed`,
		},
		{
			name: "Invalid Labels",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					" color": "green",
				},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: labels[" color"]: keys must be trimmed`,
		},
	}
	validateProtos(t, tests)
}
