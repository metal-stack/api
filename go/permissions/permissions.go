package permissions

import (
	_ "embed"
)

type ServicePermissions struct {
	Roles      Roles      `json:"roles"`
	Methods    Methods    `json:"methods"`
	Visibility Visibility `json:"visibility"`
	Auditable  Auditable  `json:"auditable,omitempty"`
	Services   []string   `json:"services,omitempty"`
}
type (
	Methods map[string]bool

	Chargeable map[string]bool
	Auditable  map[string]bool

	Admin   map[string][]string
	Infra   map[string][]string
	Tenant  map[string][]string
	Project map[string][]string
)

// Roles
type Roles struct {
	Admin   Admin   `json:"admin,omitempty"`
	Infra   Infra   `json:"infra,omitempty"`
	Tenant  Tenant  `json:"tenant,omitempty"`
	Project Project `json:"project,omitempty"`
}

type Visibility struct {
	Public  map[string]bool `json:"public,omitempty"`
	Self    map[string]bool `json:"self,omitempty"`
	Admin   map[string]bool `json:"admin,omitempty"`
	Infra   map[string]bool `json:"infra,omitempty"`
	Tenant  map[string]bool `json:"tenant,omitempty"`
	Project map[string]bool `json:"project,omitempty"`
}
