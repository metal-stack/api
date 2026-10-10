// Code generated generate.go. DO NOT EDIT.
package permissions

import (
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"connectrpc.com/connect/v2"
	"google.golang.org/protobuf/proto"
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
			Admin:    map[string]bool{
{{- range $key, $value := .Visibility.Admin }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Infra:    map[string]bool{
{{- range $key, $value := .Visibility.Infra }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Machine:    map[string]bool{
{{- range $key, $value := .Visibility.Machine }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Tenant:    map[string]bool{
{{- range $key, $value := .Visibility.Tenant }}
	"{{ $key }}": {{ $value }} ,
{{- end }}
			},
			Project:    map[string]bool{
{{- range $key, $value := .Visibility.Project }}
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

func IsPublicScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Public[spec.Procedure]
	return ok
}

func IsSelfScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Self[spec.Procedure]
	return ok
}

func IsAdminScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Admin[spec.Procedure]
	return ok
}

func IsInfraScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Infra[spec.Procedure]
	return ok
}

func IsMachineScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Machine[spec.Procedure]
	return ok
}

func IsTenantScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Tenant[spec.Procedure]
	return ok
}

func IsProjectScope(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Visibility.Project[spec.Procedure]
	return ok
}

func IsAuditable(spec connect.Spec) bool {
	_, ok := GetServicePermissions().Auditable[spec.Procedure]
	return ok
}

func GetTenantFromRequest(spec connect.Spec, req proto.Message) (string, bool) {
	if !IsTenantScope(spec) {
		return "", false
	}
	switch rq := req.(type) {
	case interface{ GetLogin() string }:
		return rq.GetLogin(), true
	}
	return "", false
}

func GetProjectFromRequest(spec connect.Spec, req proto.Message) (string, bool) {
	if !IsProjectScope(spec) {
		return "", false
	}
	switch rq := req.(type) {
	case interface{ GetProject() string }:
		return rq.GetProject(), true
	}
	return "", false
}

func GetMachineIdFromRequest(spec connect.Spec, req proto.Message) (string, bool) {
	if !IsMachineScope(spec) {
		return "", false
	}
	switch rq := req.(type) {
	case interface{ GetUuid() string }:
		return rq.GetUuid(), true
	}
	return "", false
}