// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"connectrpc.com/connect"
)

func GetServices() []string {
	return []string{
{{- range $s := .Services }}
	"{{ $s }}",
{{- end }}
	}
}

func GetServicePermissions() *ServicePermissions {
	return &ServicePermissions{
		Roles:      Roles{
			Admin:   Admin{
				{{- range $role, $methods := .Roles.Admin }}
					apiv2.AdminRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Infra:   Infra{
				{{- range $role, $methods := .Roles.Infra }}
					apiv2.InfraRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Machine:   Machine{
				{{- range $role, $methods := .Roles.Machine }}
					apiv2.MachineRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Tenant:  Tenant{
				{{- range $role, $methods := .Roles.Tenant }}
					apiv2.TenantRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
			Project: Project{
				{{- range $role, $methods := .Roles.Project }}
					apiv2.ProjectRole_{{ $role }}: map[string]struct{}{
						{{- range $method,$v := $methods }}
							"{{ $method }}":{{ $v }},
						{{- end }}
					},
				{{- end }}
			},
		},
		Methods:    map[string]struct{}{
{{- range $key, $value := .Methods }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
		},
		MethodsByScope: map[MethodScope]Methods{
{{- range $scope, $methods := .MethodsByScope }}

{{- if eq $scope "public"}}
			PublicScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "admin"}}
			AdminScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "self"}}
			SelfScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "project"}}
			ProjectScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "tenant"}}
			TenantScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "infra"}}
			InfraScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}
{{- if eq $scope "machine"}}
			MachineScope:   Methods{
{{- range $key, $value := $methods }}
				"{{ $key }}": {{ $value }} ,
{{- end }}
			},
{{- end }}

{{- end }}
		},
		Visibility: Visibility{
			Public:  map[string]bool{
{{- range $key, $value := .Visibility.Public }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Self:    map[string]bool{
{{- range $key, $value := .Visibility.Self }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
		},
		Auditable:  map[string]bool{
{{- range $key, $value := .Auditable }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
		},
	}
}

func IsPublicScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[PublicScope][req.Spec().Procedure]
	return ok
}

func IsSelfScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[SelfScope][req.Spec().Procedure]
	return ok
}

func IsAdminScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[AdminScope][req.Spec().Procedure]
	return ok
}

func IsInfraScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[InfraScope][req.Spec().Procedure]
	return ok
}

func IsMachineScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[MachineScope][req.Spec().Procedure]
	return ok
}

func IsTenantScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[TenantScope][req.Spec().Procedure]
	return ok
}

func IsProjectScope(req connect.AnyRequest) bool {
	_, ok := GetServicePermissions().MethodsByScope[ProjectScope][req.Spec().Procedure]
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

func GetMachineIdFromRequest(req connect.AnyRequest) (string, bool) {
	if !IsMachineScope(req) {
		return "", false
	}
	switch rq := req.Any().(type) {
	case interface{ GetUuid() string }:
		return rq.GetUuid(), true
	}
	return "", false
}
