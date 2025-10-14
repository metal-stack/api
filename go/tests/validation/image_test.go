package validation

import (
	"testing"

	adminv2 "github.com/metal-stack/api/go/metalstack/admin/v2"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"google.golang.org/protobuf/proto"
)

func TestValidateImage(t *testing.T) {
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
 - url: given uri must be valid [string.uri]`,
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
 - features[0]: value must be one of the defined enum values [enum.defined_only]`,
		},
		{
			name: "Valid ImageUpdate minimal config",
			msg: &adminv2.ImageServiceUpdateRequest{
				Id:             "debian:12.0.20250101",
				Name:           proto.String("debian"),
				UpdateMeta:     &apiv2.UpdateMeta{},
				Features:       []apiv2.ImageFeature{apiv2.ImageFeature_IMAGE_FEATURE_MACHINE},
				Classification: apiv2.ImageClassification_IMAGE_CLASSIFICATION_PREVIEW,
			},
			wantErr: false,
		},
		{
			name: "InValid ImageUpdate duplicate Features",
			msg: &adminv2.ImageServiceUpdateRequest{
				Id:             "debian:12.0.20250101",
				Name:           proto.String("debian"),
				UpdateMeta:     &apiv2.UpdateMeta{},
				Features:       []apiv2.ImageFeature{apiv2.ImageFeature_IMAGE_FEATURE_MACHINE, apiv2.ImageFeature_IMAGE_FEATURE_MACHINE},
				Classification: apiv2.ImageClassification_IMAGE_CLASSIFICATION_PREVIEW,
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - features: repeated value must contain unique items [repeated.unique]`,
		},
		{
			name: "InValid ImageUpdate invalid Features",
			msg: &adminv2.ImageServiceUpdateRequest{
				Id:             "debian:12.0.20250101",
				Name:           proto.String("debian"),
				UpdateMeta:     &apiv2.UpdateMeta{},
				Features:       []apiv2.ImageFeature{apiv2.ImageFeature_IMAGE_FEATURE_MACHINE, 3},
				Classification: apiv2.ImageClassification_IMAGE_CLASSIFICATION_PREVIEW,
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - features[1]: value must be one of the defined enum values [enum.defined_only]`,
		},
	}

	validateProtos(t, tests)
}
