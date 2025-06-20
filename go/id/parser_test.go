package id

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestGet(t *testing.T) {

	tur := &apiv2.TenantServiceUpdateMemberRequest{
		Login: "john@google",
	}
	nur := &apiv2.NetworkServiceUpdateRequest{
		Id: "internet",
	}
	pur := &apiv2.ProjectServiceUpdateRequest{
		Project: "p1",
	}
	tests := []struct {
		name    string
		m       proto.Message
		want    string
		wantErr bool
	}{
		{
			name: "network",
			m:    nur,
			want: "internet",
		},
		{
			name: "tenant",
			m:    tur,
			want: "john@google",
		},
		{
			name: "project",
			m:    pur,
			want: "p1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrimaryKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPrimaryKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	tur := &apiv2.TenantServiceUpdateMemberRequest{
		Login: "john@google",
	}
	for b.Loop() {
		got, err := Get(tur)
		require.NoError(b, err)
		require.NotNil(b, got)
	}
}
