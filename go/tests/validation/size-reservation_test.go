package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateSizeReservation(t *testing.T) {
	tests := prototests{
		{
			name: "Invalid SizeReservation",
			msg: &apiv2.SizeReservation{
				Partitions: []string{"partition-1"},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - name: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID
 - size: must be within 2 and 128 characters
 - amount: value must be greater than 0`,
		},
		{
			name: "Invalid SizeReservation id not a uuid",
			msg: &apiv2.SizeReservation{
				Id:         "abc",
				Partitions: []string{"partition-1"},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - id: value must be a valid UUID
 - name: must be within 2 and 128 characters
 - project: value is empty, which is not a valid UUID
 - size: must be within 2 and 128 characters
 - amount: value must be greater than 0`,
		},
		{
			name: "Invalid SizeReservation partitions not unique",
			msg: &apiv2.SizeReservation{
				Id:         "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:       "Simple",
				Project:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Size:       "c1-xlarge",
				Amount:     2,
				Partitions: []string{"partition-1", "partition-1"},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: partitions: repeated value must contain unique items`,
		},
		{
			name: "Invalid SizeReservation partitions not unique and to short",
			msg: &apiv2.SizeReservation{
				Id:         "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:       "Simple",
				Project:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Size:       "c1-xlarge",
				Amount:     3,
				Partitions: []string{"partition-1", "partition-1", "p"},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - partitions: repeated value must contain unique items
 - partitions[2]: must be within 2 and 128 characters`,
		},
		{
			name: "Valid SizeReservation",
			msg: &apiv2.SizeReservation{
				Id:         "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:       "Simple",
				Project:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Size:       "c1-xlarge",
				Amount:     3,
				Partitions: []string{"partition-1", "partition-2"},
			},
			wantErr: false,
		},
		{
			name: "Valid SizeReservation empty id",
			msg: &apiv2.SizeReservation{
				Name:       "Simple",
				Project:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Size:       "c1-xlarge",
				Amount:     3,
				Partitions: []string{"partition-1", "partition-2"},
			},
			wantErr: false,
		},
	}
	validateProtos(t, tests)
}
