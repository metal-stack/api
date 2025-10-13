package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func TestValidateToken(t *testing.T) {
	tests := prototests{
		{
			name: "Valid Token Create Request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "A Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"00000000-0000-0000-0000-000000000000": apiv2.ProjectRole_PROJECT_ROLE_EDITOR,
					"00000000-0000-0000-0000-000000000001": apiv2.ProjectRole_PROJECT_ROLE_VIEWER,
					"00000000-0000-0000-0000-000000000002": apiv2.ProjectRole_PROJECT_ROLE_OWNER,
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
 - project_roles["p42"]: value must be a valid UUID [string.uuid]
 - project_roles["p42"]: value must be one of the defined enum values [enum.defined_only]`,
		},
		{
			name: "InValid Token Create Request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"00000000-0000-0000-0000-000000000000": 5,
				},
			},
			wantErr: true,
			wantErrorMessage: `validation error:
 - project_roles["00000000-0000-0000-0000-000000000000"]: value must be one of the defined enum values [enum.defined_only]`,
		},
	}

	validateProtos(t, tests)
}
