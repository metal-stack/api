package permissions

import (
	_ "embed"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

const (
	AdminScope   MethodScope = "admin"
	InfraScope   MethodScope = "infra"
	MachineScope MethodScope = "machine"
	ProjectScope MethodScope = "project"
	PublicScope  MethodScope = "public"
	SelfScope    MethodScope = "self"
	TenantScope  MethodScope = "tenant"
)

type (
	ServicePermissions struct {
		Auditable      Auditable               `json:"auditable,omitempty"`
		Methods        Methods                 `json:"methods"`
		MethodsByScope map[MethodScope]Methods `json:"methodsByScope"`
		Roles          Roles                   `json:"roles"`
		Services       []string                `json:"services,omitempty"`
		Visibility     Visibility              `json:"visibility"`
	}

	MethodScope string
	Methods     map[string]struct{}

	Chargeable map[string]bool
	Auditable  map[string]bool

	Admin   map[apiv2.AdminRole]Methods
	Infra   map[apiv2.InfraRole]Methods
	Machine map[apiv2.MachineRole]Methods
	Project map[apiv2.ProjectRole]Methods
	Tenant  map[apiv2.TenantRole]Methods

	// Roles
	Roles struct {
		Admin   Admin   `json:"admin,omitempty"`
		Infra   Infra   `json:"infra,omitempty"`
		Machine Machine `json:"machine,omitempty"`
		Project Project `json:"project,omitempty"`
		Tenant  Tenant  `json:"tenant,omitempty"`
	}

	Visibility struct {
		Public map[string]bool `json:"public,omitempty"`
		Self   map[string]bool `json:"self,omitempty"`
	}
)
