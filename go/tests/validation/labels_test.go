package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateLabels(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Labels minimal possible",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"k": "",
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Labels keys are trimmed",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"k ": "",
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: labels: keys and values must not start or end with whitespace",
		},
		{
			name: "Valid Labels keys are trimmed",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"k ": " asdf",
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: labels: keys and values must not start or end with whitespace",
		},
		{
			name: "Valid Labels keys are trimmed",
			msg: &apiv2.Labels{
				Labels: map[string]string{
					"k ":    " asdf",
					"color": "blue",
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: labels: keys and values must not start or end with whitespace",
		},
	}
	validateProtos(t, tests)
}
