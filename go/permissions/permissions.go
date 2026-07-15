package permissions

import (
	_ "embed"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

type (
	ServicePermissions struct {
		Roles      Roles      `json:"roles"`
		Methods    Methods    `json:"methods"`
		Visibility Visibility `json:"visibility"`
		Auditable  Auditable  `json:"auditable,omitempty"`
		Services   []string   `json:"services,omitempty"`
	}

	Methods map[string]struct{}

	Chargeable map[string]bool
	Auditable  map[string]bool

	Admin   map[apiv2.AdminRole]Methods
	Infra   map[apiv2.InfraRole]Methods
	Machine map[apiv2.MachineRole]Methods
	Tenant  map[apiv2.TenantRole]Methods
	Project map[apiv2.ProjectRole]Methods

	// Roles
	Roles struct {
		Admin   Admin   `json:"admin,omitempty"`
		Infra   Infra   `json:"infra,omitempty"`
		Machine Machine `json:"machine,omitempty"`
		Tenant  Tenant  `json:"tenant,omitempty"`
		Project Project `json:"project,omitempty"`
	}

	Visibility struct {
		Public  map[string]bool `json:"public,omitempty"`
		Self    map[string]bool `json:"self,omitempty"`
		Admin   map[string]bool `json:"admin,omitempty"`
		Infra   map[string]bool `json:"infra,omitempty"`
		Machine map[string]bool `json:"machine,omitempty"`
		Tenant  map[string]bool `json:"tenant,omitempty"`
		Project map[string]bool `json:"project,omitempty"`
	}
)
