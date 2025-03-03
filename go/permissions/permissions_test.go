package permissions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetServicePermissions(t *testing.T) {
	perms := GetServicePermissions()
	require.NotNil(t, perms)
	// TODO more coverage
	require.Contains(t, perms.Methods, "/metalstack.api.v2.IPService/List")
	require.Contains(t, perms.Visibility.Self, "/metalstack.api.v2.TokenService/Create")
	require.Contains(t, perms.Visibility.Admin, "/metalstack.admin.v2.TenantService/List")
	require.Contains(t, perms.Visibility.Project, "/metalstack.api.v2.IPService/List")
}
