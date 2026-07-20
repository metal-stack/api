package validation

import (
	"testing"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func Test_validateToken(t *testing.T) {
	tests := prototests{
		{
			name: "valid token create request",
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
			name: "invalid token create request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"p42": 5,
				},
			},
			wantErr: true,
			wantErrorMessage: `validation errors:
 - project_roles: subject must be a '*', or a valid UUID
 - project_roles["p42"]: value must be one of the defined enum values`,
		},
		{
			name: "invalid token create request",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"00000000-0000-0000-0000-000000000000": 5,
				},
			},
			wantErr:          true,
			wantErrorMessage: `validation error: project_roles["00000000-0000-0000-0000-000000000000"]: value must be one of the defined enum values`,
		},
		{
			name: "valid token create request, api token with permissions",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				Permissions: []*apiv2.TypedMethodPermission{
					{
						Permissiontype: &apiv2.TypedMethodPermission_Project{
							Project: &apiv2.ProjectPermissions{
								Project: "08538119-edd3-4734-84a4-8681da17da5b",
								Methods: []string{"/metalstack.api.v2.IPService/List"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid token create request, api token with wildcard permissions",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				Permissions: []*apiv2.TypedMethodPermission{
					{
						Permissiontype: &apiv2.TypedMethodPermission_Project{
							Project: &apiv2.ProjectPermissions{
								Project: "*",
								Methods: []string{"/metalstack.api.v2.IPService/List"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid token create request, api token with project permission that has an invalid uuid",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "B Token",
				Permissions: []*apiv2.TypedMethodPermission{
					{
						Permissiontype: &apiv2.TypedMethodPermission_Project{
							Project: &apiv2.ProjectPermissions{
								Project: "foo",
								Methods: []string{"/metalstack.api.v2.IPService/List"},
							},
						},
					},
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: permissions[0].project.project: must be '*' or a uuid",
		},
		{
			name: "invalid token, user token with permissions",
			msg: &apiv2.Token{
				Uuid:        "00000000-0000-0000-0000-000000000000",
				User:        "user-a",
				Description: "B Token",
				Permissions: []*apiv2.MethodPermission{
					{
						Subject: "08538119-edd3-4734-84a4-8681da17da5b",
						Methods: []string{"/metalstack.api.v2.IPService/List"},
					},
				},
				TokenType: apiv2.TokenType_TOKEN_TYPE_USER,
			},
			wantErr:          true,
			wantErrorMessage: `validation error: token type user must not have permissions`,
		},
		{
			name: "valid token, api token with permissions",
			msg: &apiv2.Token{
				Uuid:        "00000000-0000-0000-0000-000000000000",
				User:        "user-a",
				Description: "B Token",
				Permissions: []*apiv2.MethodPermission{
					{
						Subject: "08538119-edd3-4734-84a4-8681da17da5b",
						Methods: []string{"/metalstack.api.v2.IPService/List"},
					},
				},
				TokenType: apiv2.TokenType_TOKEN_TYPE_API,
			},
			wantErr: false,
		},
		{
			name: "valid token, api token with machineroles",
			msg: &apiv2.TokenServiceUpdateRequest{
				Uuid: "00000000-0000-0000-0000-000000000000",
				MachineRoles: map[string]apiv2.MachineRole{
					"*":                                    apiv2.MachineRole_MACHINE_ROLE_EDITOR,
					"00000000-0000-0000-0000-000000000000": apiv2.MachineRole_MACHINE_ROLE_EDITOR,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid token, api token with machineroles with invalid uuid",
			msg: &apiv2.TokenServiceUpdateRequest{
				Uuid: "00000000-0000-0000-0000-000000000000",
				MachineRoles: map[string]apiv2.MachineRole{
					"*":                                    apiv2.MachineRole_MACHINE_ROLE_EDITOR,
					"y0000000-0000-0000-0000-000000000000": apiv2.MachineRole_MACHINE_ROLE_EDITOR,
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: machine_roles: subject must be empty string, '*', or a valid UUID",
		},
		{
			name: "invalid token, api token with machineroles with invalid key",
			msg: &apiv2.TokenServiceUpdateRequest{
				Uuid: "00000000-0000-0000-0000-000000000000",
				MachineRoles: map[string]apiv2.MachineRole{
					"A":                                    apiv2.MachineRole_MACHINE_ROLE_EDITOR,
					"*":                                    apiv2.MachineRole_MACHINE_ROLE_EDITOR,
					"00000000-0000-0000-0000-000000000000": apiv2.MachineRole_MACHINE_ROLE_EDITOR,
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: machine_roles: subject must be empty string, '*', or a valid UUID",
		},
		{
			name: "invalid token create request, token role not trimmed",
			msg: &apiv2.TokenServiceCreateRequest{
				Description: "A Token",
				ProjectRoles: map[string]apiv2.ProjectRole{
					"00000000-0000-0000-0000-000000000000": apiv2.ProjectRole_PROJECT_ROLE_EDITOR,
					"00000000-0000-0000-0000-000000000001": apiv2.ProjectRole_PROJECT_ROLE_VIEWER,
					"00000000-0000-0000-0000-000000000002": apiv2.ProjectRole_PROJECT_ROLE_OWNER,
				},
				TenantRoles: map[string]apiv2.TenantRole{
					" t12": apiv2.TenantRole_TENANT_ROLE_EDITOR,
				},
			},
			wantErr:          true,
			wantErrorMessage: "validation error: tenant_roles: keys must not start or end with whitespace",
		},
	}

	validateProtos(t, tests)
}
