package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateToken(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := prototests{
		{
			name: "Valid Token Create Request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "A Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"p12": apiv2.ProjectRole_PROJECT_ROLE_EDITOR,
					"p22": apiv2.ProjectRole_PROJECT_ROLE_VIEWER,
					"p32": apiv2.ProjectRole_PROJECT_ROLE_OWNER,
				},
				TenantRoles: map[string]apiv2.TenantRole{
					"t12": apiv2.TenantRole_TENANT_ROLE_EDITOR,
				},
			},
			wantErr: false,
		},
		{
			name: "InValid Token Create Request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"p42": 5,
				},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - project_roles["p42"]: value must be one of the defined enum values [enum.defined_only]`,
		},
	}

	validateProtos(t, tests, validator)
}
