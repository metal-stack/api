package validation

import (
	"testing"

	apiv1 "github.com/metal-stack/api/go/metalstack/api/v2"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/require"
)

func TestValidateIP(t *testing.T) {

	validator, err := protovalidate.New()
	require.NoError(t, err)

	ip := &apiv1.IP{
		Uuid: "abc",
	}

	err = validator.Validate(ip)
	require.Error(t, err)
	require.EqualError(t, err, "validation error:\n - uuid: value must be a valid UUID [string.uuid]\n - ip: value is empty, which is not a valid IP address [string.ip_empty]\n - name: value length must be at least 2 characters [string.min_len]\n - network: value length must be at least 2 characters [string.min_len]\n - project: value is empty, which is not a valid UUID [string.uuid_empty]")
}
