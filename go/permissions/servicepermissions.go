// Code generated discover.go. DO NOT EDIT.
package permissions

import (
	"connectrpc.com/connect"
)

func GetServices() []string {
	return []string{
		"metalstack.admin.v2.FilesystemService",
		"metalstack.admin.v2.IPService",
		"metalstack.admin.v2.ImageService",
		"metalstack.admin.v2.NetworkService",
		"metalstack.admin.v2.PartitionService",
		"metalstack.admin.v2.SizeService",
		"metalstack.admin.v2.TenantService",
		"metalstack.admin.v2.TokenService",
		"metalstack.api.v2.FilesystemService",
		"metalstack.api.v2.HealthService",
		"metalstack.api.v2.IPService",
		"metalstack.api.v2.ImageService",
		"metalstack.api.v2.MethodService",
		"metalstack.api.v2.NetworkService",
		"metalstack.api.v2.PartitionService",
		"metalstack.api.v2.ProjectService",
		"metalstack.api.v2.SizeService",
		"metalstack.api.v2.TenantService",
		"metalstack.api.v2.TokenService",
		"metalstack.api.v2.UserService",
		"metalstack.api.v2.VersionService",
		"metalstack.infra.v2.BMCService",
	}
}

func GetServicePermissions() *ServicePermissions {
	return &ServicePermissions{
		Roles: Roles{
			Admin: Admin{
				"ADMIN_ROLE_EDITOR": []string{
					"/metalstack.admin.v2.FilesystemService/Create",
					"/metalstack.admin.v2.FilesystemService/Update",
					"/metalstack.admin.v2.FilesystemService/Delete",
					"/metalstack.admin.v2.ImageService/Create",
					"/metalstack.admin.v2.ImageService/Update",
					"/metalstack.admin.v2.ImageService/Delete",
					"/metalstack.admin.v2.ImageService/Usage",
					"/metalstack.admin.v2.IPService/List",
					"/metalstack.admin.v2.IPService/Issues",
					"/metalstack.admin.v2.NetworkService/Get",
					"/metalstack.admin.v2.NetworkService/Create",
					"/metalstack.admin.v2.NetworkService/Update",
					"/metalstack.admin.v2.NetworkService/Delete",
					"/metalstack.admin.v2.NetworkService/List",
					"/metalstack.admin.v2.PartitionService/Create",
					"/metalstack.admin.v2.PartitionService/Update",
					"/metalstack.admin.v2.PartitionService/Delete",
					"/metalstack.admin.v2.PartitionService/Capacity",
					"/metalstack.admin.v2.SizeService/Create",
					"/metalstack.admin.v2.SizeService/Update",
					"/metalstack.admin.v2.SizeService/Delete",
					"/metalstack.admin.v2.TenantService/Create",
					"/metalstack.admin.v2.TenantService/List",
					"/metalstack.admin.v2.TokenService/List",
					"/metalstack.admin.v2.TokenService/Revoke",
				},
				"ADMIN_ROLE_VIEWER": []string{
					"/metalstack.admin.v2.ImageService/Usage",
					"/metalstack.admin.v2.IPService/List",
					"/metalstack.admin.v2.IPService/Issues",
					"/metalstack.admin.v2.NetworkService/Get",
					"/metalstack.admin.v2.NetworkService/List",
					"/metalstack.admin.v2.PartitionService/Capacity",
					"/metalstack.admin.v2.SizeService/Create",
					"/metalstack.admin.v2.SizeService/Update",
					"/metalstack.admin.v2.SizeService/Delete",
					"/metalstack.admin.v2.TenantService/List",
					"/metalstack.admin.v2.TokenService/List",
				},
			},
			Infra: Infra{
				"INFRA_ROLE_EDITOR": []string{
					"/metalstack.infra.v2.BMCService/UpdateBMCInfo",
				},
				"INFRA_ROLE_VIEWER": []string{
					"/metalstack.infra.v2.BMCService/UpdateBMCInfo",
				},
			},
			Tenant: Tenant{
				"TENANT_ROLE_EDITOR": []string{
					"/metalstack.api.v2.ProjectService/Create",
					"/metalstack.api.v2.TenantService/Get",
					"/metalstack.api.v2.TenantService/Update",
					"/metalstack.api.v2.TenantService/Delete",
				},
				"TENANT_ROLE_GUEST": []string{
					"/metalstack.api.v2.TenantService/Get",
				},
				"TENANT_ROLE_OWNER": []string{
					"/metalstack.api.v2.ProjectService/Create",
					"/metalstack.api.v2.TenantService/Get",
					"/metalstack.api.v2.TenantService/Update",
					"/metalstack.api.v2.TenantService/Delete",
					"/metalstack.api.v2.TenantService/RemoveMember",
					"/metalstack.api.v2.TenantService/UpdateMember",
					"/metalstack.api.v2.TenantService/Invite",
					"/metalstack.api.v2.TenantService/InviteDelete",
					"/metalstack.api.v2.TenantService/InvitesList",
				},
				"TENANT_ROLE_VIEWER": []string{
					"/metalstack.api.v2.TenantService/Get",
				},
			},
			Project: Project{
				"PROJECT_ROLE_EDITOR": []string{
					"/metalstack.api.v2.IPService/Get",
					"/metalstack.api.v2.IPService/Create",
					"/metalstack.api.v2.IPService/Update",
					"/metalstack.api.v2.IPService/List",
					"/metalstack.api.v2.IPService/Delete",
					"/metalstack.api.v2.NetworkService/Get",
					"/metalstack.api.v2.NetworkService/Create",
					"/metalstack.api.v2.NetworkService/Update",
					"/metalstack.api.v2.NetworkService/List",
					"/metalstack.api.v2.NetworkService/ListBaseNetworks",
					"/metalstack.api.v2.NetworkService/Delete",
					"/metalstack.api.v2.ProjectService/Get",
					"/metalstack.api.v2.ProjectService/Update",
				},
				"PROJECT_ROLE_OWNER": []string{
					"/metalstack.api.v2.IPService/Get",
					"/metalstack.api.v2.IPService/Create",
					"/metalstack.api.v2.IPService/Update",
					"/metalstack.api.v2.IPService/List",
					"/metalstack.api.v2.IPService/Delete",
					"/metalstack.api.v2.NetworkService/Get",
					"/metalstack.api.v2.NetworkService/Create",
					"/metalstack.api.v2.NetworkService/Update",
					"/metalstack.api.v2.NetworkService/List",
					"/metalstack.api.v2.NetworkService/ListBaseNetworks",
					"/metalstack.api.v2.NetworkService/Delete",
					"/metalstack.api.v2.ProjectService/Get",
					"/metalstack.api.v2.ProjectService/Delete",
					"/metalstack.api.v2.ProjectService/Update",
					"/metalstack.api.v2.ProjectService/RemoveMember",
					"/metalstack.api.v2.ProjectService/UpdateMember",
					"/metalstack.api.v2.ProjectService/Invite",
					"/metalstack.api.v2.ProjectService/InviteDelete",
					"/metalstack.api.v2.ProjectService/InvitesList",
				},
				"PROJECT_ROLE_VIEWER": []string{
					"/metalstack.api.v2.IPService/Get",
					"/metalstack.api.v2.IPService/List",
					"/metalstack.api.v2.NetworkService/Get",
					"/metalstack.api.v2.NetworkService/List",
					"/metalstack.api.v2.NetworkService/ListBaseNetworks",
					"/metalstack.api.v2.ProjectService/Get",
				},
			},
		},
		Methods: map[string]bool{
			"/metalstack.admin.v2.FilesystemService/Create":      true,
			"/metalstack.admin.v2.FilesystemService/Delete":      true,
			"/metalstack.admin.v2.FilesystemService/Update":      true,
			"/metalstack.admin.v2.IPService/Issues":              true,
			"/metalstack.admin.v2.IPService/List":                true,
			"/metalstack.admin.v2.ImageService/Create":           true,
			"/metalstack.admin.v2.ImageService/Delete":           true,
			"/metalstack.admin.v2.ImageService/Update":           true,
			"/metalstack.admin.v2.ImageService/Usage":            true,
			"/metalstack.admin.v2.NetworkService/Create":         true,
			"/metalstack.admin.v2.NetworkService/Delete":         true,
			"/metalstack.admin.v2.NetworkService/Get":            true,
			"/metalstack.admin.v2.NetworkService/List":           true,
			"/metalstack.admin.v2.NetworkService/Update":         true,
			"/metalstack.admin.v2.PartitionService/Capacity":     true,
			"/metalstack.admin.v2.PartitionService/Create":       true,
			"/metalstack.admin.v2.PartitionService/Delete":       true,
			"/metalstack.admin.v2.PartitionService/Update":       true,
			"/metalstack.admin.v2.SizeService/Create":            true,
			"/metalstack.admin.v2.SizeService/Delete":            true,
			"/metalstack.admin.v2.SizeService/Update":            true,
			"/metalstack.admin.v2.TenantService/Create":          true,
			"/metalstack.admin.v2.TenantService/List":            true,
			"/metalstack.admin.v2.TokenService/List":             true,
			"/metalstack.admin.v2.TokenService/Revoke":           true,
			"/metalstack.api.v2.FilesystemService/Get":           true,
			"/metalstack.api.v2.FilesystemService/List":          true,
			"/metalstack.api.v2.FilesystemService/Match":         true,
			"/metalstack.api.v2.FilesystemService/Try":           true,
			"/metalstack.api.v2.HealthService/Get":               true,
			"/metalstack.api.v2.IPService/Create":                true,
			"/metalstack.api.v2.IPService/Delete":                true,
			"/metalstack.api.v2.IPService/Get":                   true,
			"/metalstack.api.v2.IPService/List":                  true,
			"/metalstack.api.v2.IPService/Update":                true,
			"/metalstack.api.v2.ImageService/Get":                true,
			"/metalstack.api.v2.ImageService/Latest":             true,
			"/metalstack.api.v2.ImageService/List":               true,
			"/metalstack.api.v2.MethodService/List":              true,
			"/metalstack.api.v2.MethodService/TokenScopedList":   true,
			"/metalstack.api.v2.NetworkService/Create":           true,
			"/metalstack.api.v2.NetworkService/Delete":           true,
			"/metalstack.api.v2.NetworkService/Get":              true,
			"/metalstack.api.v2.NetworkService/List":             true,
			"/metalstack.api.v2.NetworkService/ListBaseNetworks": true,
			"/metalstack.api.v2.NetworkService/Update":           true,
			"/metalstack.api.v2.PartitionService/Get":            true,
			"/metalstack.api.v2.PartitionService/List":           true,
			"/metalstack.api.v2.ProjectService/Create":           true,
			"/metalstack.api.v2.ProjectService/Delete":           true,
			"/metalstack.api.v2.ProjectService/Get":              true,
			"/metalstack.api.v2.ProjectService/Invite":           true,
			"/metalstack.api.v2.ProjectService/InviteAccept":     true,
			"/metalstack.api.v2.ProjectService/InviteDelete":     true,
			"/metalstack.api.v2.ProjectService/InviteGet":        true,
			"/metalstack.api.v2.ProjectService/InvitesList":      true,
			"/metalstack.api.v2.ProjectService/List":             true,
			"/metalstack.api.v2.ProjectService/RemoveMember":     true,
			"/metalstack.api.v2.ProjectService/Update":           true,
			"/metalstack.api.v2.ProjectService/UpdateMember":     true,
			"/metalstack.api.v2.SizeService/Get":                 true,
			"/metalstack.api.v2.SizeService/List":                true,
			"/metalstack.api.v2.TenantService/Create":            true,
			"/metalstack.api.v2.TenantService/Delete":            true,
			"/metalstack.api.v2.TenantService/Get":               true,
			"/metalstack.api.v2.TenantService/Invite":            true,
			"/metalstack.api.v2.TenantService/InviteAccept":      true,
			"/metalstack.api.v2.TenantService/InviteDelete":      true,
			"/metalstack.api.v2.TenantService/InviteGet":         true,
			"/metalstack.api.v2.TenantService/InvitesList":       true,
			"/metalstack.api.v2.TenantService/List":              true,
			"/metalstack.api.v2.TenantService/RemoveMember":      true,
			"/metalstack.api.v2.TenantService/Update":            true,
			"/metalstack.api.v2.TenantService/UpdateMember":      true,
			"/metalstack.api.v2.TokenService/Create":             true,
			"/metalstack.api.v2.TokenService/Get":                true,
			"/metalstack.api.v2.TokenService/List":               true,
			"/metalstack.api.v2.TokenService/Refresh":            true,
			"/metalstack.api.v2.TokenService/Revoke":             true,
			"/metalstack.api.v2.TokenService/Update":             true,
			"/metalstack.api.v2.UserService/Get":                 true,
			"/metalstack.api.v2.VersionService/Get":              true,
			"/metalstack.infra.v2.BMCService/UpdateBMCInfo":      true,
		},
		Visibility: Visibility{
			Public: map[string]bool{
				"/grpc.reflection.v1.ServerReflection/ServerReflectionInfo":      true,
				"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo": true,
				"/metalstack.api.v2.HealthService/Get":                           true,
				"/metalstack.api.v2.MethodService/List":                          true,
				"/metalstack.api.v2.VersionService/Get":                          true,
			},
			Self: map[string]bool{
				"/metalstack.api.v2.FilesystemService/Get":         true,
				"/metalstack.api.v2.FilesystemService/List":        true,
				"/metalstack.api.v2.FilesystemService/Match":       true,
				"/metalstack.api.v2.FilesystemService/Try":         true,
				"/metalstack.api.v2.ImageService/Get":              true,
				"/metalstack.api.v2.ImageService/Latest":           true,
				"/metalstack.api.v2.ImageService/List":             true,
				"/metalstack.api.v2.MethodService/TokenScopedList": true,
				"/metalstack.api.v2.PartitionService/Get":          true,
				"/metalstack.api.v2.PartitionService/List":         true,
				"/metalstack.api.v2.ProjectService/InviteAccept":   true,
				"/metalstack.api.v2.ProjectService/InviteGet":      true,
				"/metalstack.api.v2.ProjectService/List":           true,
				"/metalstack.api.v2.SizeService/Get":               true,
				"/metalstack.api.v2.SizeService/List":              true,
				"/metalstack.api.v2.TenantService/Create":          true,
				"/metalstack.api.v2.TenantService/InviteAccept":    true,
				"/metalstack.api.v2.TenantService/InviteGet":       true,
				"/metalstack.api.v2.TenantService/List":            true,
				"/metalstack.api.v2.TokenService/Create":           true,
				"/metalstack.api.v2.TokenService/Get":              true,
				"/metalstack.api.v2.TokenService/List":             true,
				"/metalstack.api.v2.TokenService/Refresh":          true,
				"/metalstack.api.v2.TokenService/Revoke":           true,
				"/metalstack.api.v2.TokenService/Update":           true,
				"/metalstack.api.v2.UserService/Get":               true,
			},
			Admin: map[string]bool{
				"/metalstack.admin.v2.FilesystemService/Create":  true,
				"/metalstack.admin.v2.FilesystemService/Delete":  true,
				"/metalstack.admin.v2.FilesystemService/Update":  true,
				"/metalstack.admin.v2.IPService/Issues":          true,
				"/metalstack.admin.v2.IPService/List":            true,
				"/metalstack.admin.v2.ImageService/Create":       true,
				"/metalstack.admin.v2.ImageService/Delete":       true,
				"/metalstack.admin.v2.ImageService/Update":       true,
				"/metalstack.admin.v2.ImageService/Usage":        true,
				"/metalstack.admin.v2.NetworkService/Create":     true,
				"/metalstack.admin.v2.NetworkService/Delete":     true,
				"/metalstack.admin.v2.NetworkService/Get":        true,
				"/metalstack.admin.v2.NetworkService/List":       true,
				"/metalstack.admin.v2.NetworkService/Update":     true,
				"/metalstack.admin.v2.PartitionService/Capacity": true,
				"/metalstack.admin.v2.PartitionService/Create":   true,
				"/metalstack.admin.v2.PartitionService/Delete":   true,
				"/metalstack.admin.v2.PartitionService/Update":   true,
				"/metalstack.admin.v2.SizeService/Create":        true,
				"/metalstack.admin.v2.SizeService/Delete":        true,
				"/metalstack.admin.v2.SizeService/Update":        true,
				"/metalstack.admin.v2.TenantService/Create":      true,
				"/metalstack.admin.v2.TenantService/List":        true,
				"/metalstack.admin.v2.TokenService/List":         true,
				"/metalstack.admin.v2.TokenService/Revoke":       true,
			},
			Infra: map[string]bool{
				"/metalstack.infra.v2.BMCService/UpdateBMCInfo": true,
			},
			Tenant: map[string]bool{
				"/metalstack.api.v2.ProjectService/Create":      true,
				"/metalstack.api.v2.TenantService/Delete":       true,
				"/metalstack.api.v2.TenantService/Get":          true,
				"/metalstack.api.v2.TenantService/Invite":       true,
				"/metalstack.api.v2.TenantService/InviteDelete": true,
				"/metalstack.api.v2.TenantService/InvitesList":  true,
				"/metalstack.api.v2.TenantService/RemoveMember": true,
				"/metalstack.api.v2.TenantService/Update":       true,
				"/metalstack.api.v2.TenantService/UpdateMember": true,
			},
			Project: map[string]bool{
				"/metalstack.api.v2.IPService/Create":                true,
				"/metalstack.api.v2.IPService/Delete":                true,
				"/metalstack.api.v2.IPService/Get":                   true,
				"/metalstack.api.v2.IPService/List":                  true,
				"/metalstack.api.v2.IPService/Update":                true,
				"/metalstack.api.v2.NetworkService/Create":           true,
				"/metalstack.api.v2.NetworkService/Delete":           true,
				"/metalstack.api.v2.NetworkService/Get":              true,
				"/metalstack.api.v2.NetworkService/List":             true,
				"/metalstack.api.v2.NetworkService/ListBaseNetworks": true,
				"/metalstack.api.v2.NetworkService/Update":           true,
				"/metalstack.api.v2.ProjectService/Delete":           true,
				"/metalstack.api.v2.ProjectService/Get":              true,
				"/metalstack.api.v2.ProjectService/Invite":           true,
				"/metalstack.api.v2.ProjectService/InviteDelete":     true,
				"/metalstack.api.v2.ProjectService/InvitesList":      true,
				"/metalstack.api.v2.ProjectService/RemoveMember":     true,
				"/metalstack.api.v2.ProjectService/Update":           true,
				"/metalstack.api.v2.ProjectService/UpdateMember":     true,
			},
		},
		Auditable: map[string]bool{
			"/metalstack.admin.v2.FilesystemService/Create":      true,
			"/metalstack.admin.v2.FilesystemService/Delete":      true,
			"/metalstack.admin.v2.FilesystemService/Update":      true,
			"/metalstack.admin.v2.IPService/Issues":              false,
			"/metalstack.admin.v2.IPService/List":                false,
			"/metalstack.admin.v2.ImageService/Create":           true,
			"/metalstack.admin.v2.ImageService/Delete":           true,
			"/metalstack.admin.v2.ImageService/Update":           true,
			"/metalstack.admin.v2.ImageService/Usage":            true,
			"/metalstack.admin.v2.NetworkService/Create":         true,
			"/metalstack.admin.v2.NetworkService/Delete":         true,
			"/metalstack.admin.v2.NetworkService/Get":            false,
			"/metalstack.admin.v2.NetworkService/List":           false,
			"/metalstack.admin.v2.NetworkService/Update":         true,
			"/metalstack.admin.v2.PartitionService/Capacity":     false,
			"/metalstack.admin.v2.PartitionService/Create":       true,
			"/metalstack.admin.v2.PartitionService/Delete":       true,
			"/metalstack.admin.v2.PartitionService/Update":       true,
			"/metalstack.admin.v2.SizeService/Create":            true,
			"/metalstack.admin.v2.SizeService/Delete":            true,
			"/metalstack.admin.v2.SizeService/Update":            true,
			"/metalstack.admin.v2.TenantService/Create":          true,
			"/metalstack.admin.v2.TenantService/List":            true,
			"/metalstack.admin.v2.TokenService/List":             true,
			"/metalstack.admin.v2.TokenService/Revoke":           true,
			"/metalstack.api.v2.FilesystemService/Get":           false,
			"/metalstack.api.v2.FilesystemService/List":          false,
			"/metalstack.api.v2.FilesystemService/Match":         false,
			"/metalstack.api.v2.FilesystemService/Try":           false,
			"/metalstack.api.v2.HealthService/Get":               false,
			"/metalstack.api.v2.IPService/Create":                true,
			"/metalstack.api.v2.IPService/Delete":                true,
			"/metalstack.api.v2.IPService/Get":                   false,
			"/metalstack.api.v2.IPService/List":                  false,
			"/metalstack.api.v2.IPService/Update":                true,
			"/metalstack.api.v2.ImageService/Get":                false,
			"/metalstack.api.v2.ImageService/Latest":             false,
			"/metalstack.api.v2.ImageService/List":               false,
			"/metalstack.api.v2.MethodService/List":              true,
			"/metalstack.api.v2.MethodService/TokenScopedList":   true,
			"/metalstack.api.v2.NetworkService/Create":           true,
			"/metalstack.api.v2.NetworkService/Delete":           true,
			"/metalstack.api.v2.NetworkService/Get":              false,
			"/metalstack.api.v2.NetworkService/List":             false,
			"/metalstack.api.v2.NetworkService/ListBaseNetworks": false,
			"/metalstack.api.v2.NetworkService/Update":           true,
			"/metalstack.api.v2.PartitionService/Get":            false,
			"/metalstack.api.v2.PartitionService/List":           false,
			"/metalstack.api.v2.ProjectService/Create":           true,
			"/metalstack.api.v2.ProjectService/Delete":           true,
			"/metalstack.api.v2.ProjectService/Get":              false,
			"/metalstack.api.v2.ProjectService/Invite":           true,
			"/metalstack.api.v2.ProjectService/InviteAccept":     true,
			"/metalstack.api.v2.ProjectService/InviteDelete":     true,
			"/metalstack.api.v2.ProjectService/InviteGet":        false,
			"/metalstack.api.v2.ProjectService/InvitesList":      false,
			"/metalstack.api.v2.ProjectService/List":             false,
			"/metalstack.api.v2.ProjectService/RemoveMember":     true,
			"/metalstack.api.v2.ProjectService/Update":           true,
			"/metalstack.api.v2.ProjectService/UpdateMember":     true,
			"/metalstack.api.v2.SizeService/Get":                 false,
			"/metalstack.api.v2.SizeService/List":                false,
			"/metalstack.api.v2.TenantService/Create":            true,
			"/metalstack.api.v2.TenantService/Delete":            true,
			"/metalstack.api.v2.TenantService/Get":               false,
			"/metalstack.api.v2.TenantService/Invite":            true,
			"/metalstack.api.v2.TenantService/InviteAccept":      true,
			"/metalstack.api.v2.TenantService/InviteDelete":      true,
			"/metalstack.api.v2.TenantService/InviteGet":         false,
			"/metalstack.api.v2.TenantService/InvitesList":       false,
			"/metalstack.api.v2.TenantService/List":              false,
			"/metalstack.api.v2.TenantService/RemoveMember":      true,
			"/metalstack.api.v2.TenantService/Update":            true,
			"/metalstack.api.v2.TenantService/UpdateMember":      true,
			"/metalstack.api.v2.TokenService/Create":             true,
			"/metalstack.api.v2.TokenService/Get":                true,
			"/metalstack.api.v2.TokenService/List":               true,
			"/metalstack.api.v2.TokenService/Refresh":            true,
			"/metalstack.api.v2.TokenService/Revoke":             true,
			"/metalstack.api.v2.TokenService/Update":             true,
			"/metalstack.api.v2.UserService/Get":                 true,
			"/metalstack.api.v2.VersionService/Get":              false,
			"/metalstack.infra.v2.BMCService/UpdateBMCInfo":      false,
		},
	}
}

func IsPublicScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Public[req.Spec().Procedure]
	return ok
}

func IsSelfScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Self[req.Spec().Procedure]
	return ok
}

func IsAdminScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Admin[req.Spec().Procedure]
	return ok
}

func IsInfraScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Infra[req.Spec().Procedure]
	return ok
}

func IsTenantScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Tenant[req.Spec().Procedure]
	return ok
}

func IsProjectScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Visibility.Project[req.Spec().Procedure]
	return ok
}

func IsAuditable(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().Auditable[req.Spec().Procedure]
	return ok
}

func GetTenantFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsTenantScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetLogin() string }:
		return rq.GetLogin(), true
	}
	return "", false
}

func GetProjectFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsProjectScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetProject() string }:
		return rq.GetProject(), true
	}
	return "", false
}
