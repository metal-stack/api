package enum

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestGetEnum(t *testing.T) {
	got, err := GetEnum[apiv2.NetworkType]("shared")
	require.NoError(t, err)
	require.Equal(t, got, apiv2.NetworkType_NETWORK_TYPE_SHARED)

	got2, err := GetEnum[apiv2.Format]("ext4")
	require.NoError(t, err)
	require.Equal(t, got2, apiv2.Format_FORMAT_EXT4)

	got3, err := GetEnum[apiv2.Format]("invalid")
	require.Error(t, err)
	require.EqualError(t, err, `no matching enum of type:metalstack.api.v2.Format for "invalid" found`)
	require.Equal(t, got3, apiv2.Format_FORMAT_UNSPECIFIED)
}

func TestGetStringValue(t *testing.T) {
	got, err := GetStringValue(apiv2.NetworkType_NETWORK_TYPE_SHARED)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "shared", *got)

	got, err = GetStringValue(apiv2.Format_FORMAT_EXT4)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "ext4", *got)

	got, err = GetStringValue(protoreflect.Enum(apiv2.Format_FORMAT_UNSPECIFIED))
	require.Error(t, err)
	require.EqualError(t, err, "unable to fetch stringvalue from FORMAT_UNSPECIFIED")
	require.Nil(t, got)
}
