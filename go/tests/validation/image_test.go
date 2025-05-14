package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateImage(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid Image minimal config",
			msg: &apiv2.Image{
				Id:             "debian:12.0.20250101",
				Url:            "http://download.org/debian:12.0",
				Features:       []apiv2.ImageFeature{apiv2.ImageFeature_IMAGE_FEATURE_MACHINE},
				Classification: apiv2.ImageClassification_IMAGE_CLASSIFICATION_PREVIEW,
			},
			wantErr: false,
		},
		{
			name: "Invalid Image, url wrong",
			msg: &apiv2.Image{
				Id:       "debian:12.0.20250101",
				Url:      "not-a-uri",
				Features: []apiv2.ImageFeature{apiv2.ImageFeature_IMAGE_FEATURE_MACHINE},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - url: url must be a valid URI [valid_url]`,
		},
		{
			name: "Invalid Image, no features",
			msg: &apiv2.Image{
				Id:       "debian:12.0.20250101",
				Url:      "http://download.org/debian:12.0",
				Features: []apiv2.ImageFeature{3},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - features[0]: feature must be valid [features]`,
		},
	}

	validateProtos(t, tests, validator)
}
