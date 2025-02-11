package validation

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/google/go-cmp/cmp"
	apiv1 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestValidateIP(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := []struct {
		name             string
		msg              proto.Message
		wantErr          bool
		wantErrorMessage string
	}{
		{
			name: "Invalid IP",
			msg: &apiv1.IP{
				Uuid: "abc",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - uuid: value must be a valid UUID [string.uuid]
 - ip: value is empty, which is not a valid IP address [string.ip_empty]
 - name: value length must be at least 2 characters [string.min_len]
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP, but valid uuid",
			msg: &apiv1.IP{
				Uuid: "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - ip: value is empty, which is not a valid IP address [string.ip_empty]
 - name: value length must be at least 2 characters [string.min_len]
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP, but valid uuid and ipv4",
			msg: &apiv1.IP{
				Uuid: "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Ip:   "1.2.3.4",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - name: value length must be at least 2 characters [string.min_len]
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP, but valid uuid and ipv6",
			msg: &apiv1.IP{
				Uuid: "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Ip:   "fe80:db8::1",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - name: value length must be at least 2 characters [string.min_len]
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP, but valid uuid, name and ipv6",
			msg: &apiv1.IP{
				Uuid: "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name: "Test IPv6",
				Ip:   "fe80:db8::1",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP, but valid uuid, name, network and ipv6",
			msg: &apiv1.IP{
				Uuid:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:    "Test IPv6",
				Ip:      "fe80:db8::1",
				Network: "Internet",
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Invalid IP with invalid type",
			msg: &apiv1.IP{
				Uuid:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:    "Test IPv6",
				Ip:      "fe80:db8::1",
				Network: "Internet",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Type:    apiv1.IPType(99),
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - type: value must be one of the defined enum values [enum.defined_only]`,
		},
		{
			name: "Valid IP",
			msg: &apiv1.IP{
				Uuid:    "e266fcc6-f6de-4ee1-ba26-baa17bf47b13",
				Name:    "Test IPv6",
				Ip:      "fe80:db8::1",
				Network: "Internet",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
				Type:    apiv1.IPType_IP_TYPE_EPHEMERAL,
			},
			wantErr: false,
		},
		// IPServiceGetRequest
		{
			name:    "invalid IPServiceGetRequest",
			msg:     &apiv1.IPServiceGetRequest{},
			wantErr: true,
			wantErrorMessage: `validation error:
 - ip: value is empty, which is not a valid IP address [string.ip_empty]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Valid IPServiceGetRequest",
			msg: &apiv1.IPServiceGetRequest{
				Ip:      "fe80:db8::1",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
			},
			wantErr: false,
		},
		// IPServiceCreateRequest
		{
			name:    "invalid IPServiceCreateRequest",
			msg:     &apiv1.IPServiceCreateRequest{},
			wantErr: true,
			wantErrorMessage: `validation error:
 - network: value length must be at least 2 characters [string.min_len]
 - project: value is empty, which is not a valid UUID [string.uuid_empty]`,
		},
		{
			name: "Valid IPServiceCreateRequest",
			msg: &apiv1.IPServiceCreateRequest{
				Network:      "Internet",
				Project: "57cd8678-9ff0-4f8c-a34a-43d8f16caadf",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.msg)
			if err != nil && !tt.wantErr {
				t.Errorf("validate error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErr {
				if diff := cmp.Diff(err.Error(), tt.wantErrorMessage); diff != "" {
					t.Errorf("validate error = %v, diff %v", err, diff)
				}
			}
		})
	}
}
