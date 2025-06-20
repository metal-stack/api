package enum

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestGetEnum(t *testing.T) {
	got, err := GetEnum[apiv2.Format]("ext4")
	require.NoError(t, err)
	require.Equal(t, got, apiv2.Format_FORMAT_EXT4)

	got2, err := GetEnum[apiv2.Format]("ext4")
	require.NoError(t, err)
	require.Equal(t, got2, apiv2.Format_FORMAT_EXT4)

	got3, err := GetEnum[apiv2.Format]("invalid")
	require.Error(t, err)
	require.EqualError(t, err, `no matching enum of type:metalstack.api.v2.Format for "invalid" found`)
	require.Equal(t, got3, apiv2.Format_FORMAT_UNSPECIFIED)
}

func TestGetStringValue(t *testing.T) {
	got, err := GetStringValue(protoreflect.Enum(apiv2.Format_FORMAT_EXT4))
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "ext4", *got)

	got, err = GetStringValue(apiv2.Format_FORMAT_EXT4)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "ext4", *got)

	got, err = GetStringValue(protoreflect.Enum(apiv2.Format_FORMAT_UNSPECIFIED))
	require.Error(t, err)
	require.EqualError(t, err, "unable to fetch stringvalue from FORMAT_UNSPECIFIED")
	require.Nil(t, got)

	var e *apiv2.Format
	got, err = GetStringValue(e)
	require.Error(t, err)
	require.EqualError(t, err, "given enum is a nil pointer")
	require.Nil(t, got)

	e = apiv2.Format_FORMAT_EXT3.Enum()
	got, err = GetStringValue(e)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "ext3", *got)
}

func TestGetPrimaryKey(t *testing.T) {

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
		m       protoreflect.Message
		want    string
		wantErr bool
	}{
		{
			name: "network",
			m:    nur.ProtoReflect(),
			want: "internet",
		},
		{
			name: "tenant",
			m:    tur.ProtoReflect(),
			want: "john@google",
		},
		{
			name: "project",
			m:    pur.ProtoReflect(),
			want: "p1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPrimaryKey(tt.m)
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
