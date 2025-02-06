// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	apiclient "github.com/metal-stack/api/go/client"
	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	adminv2mocks "github.com/metal-stack/api/go/tests/mocks/metalstack/admin/v2/adminv2connect"
	apiv2mocks "github.com/metal-stack/api/go/tests/mocks/metalstack/api/v2/apiv2connect"

	"github.com/stretchr/testify/mock"
)

type (
	client struct {
		adminv2service *adminv2
		apiv2service   *apiv2
	}

	ClientMockFns struct {
		Adminv2Mocks *Adminv2MockFns
		Apiv2Mocks   *Apiv2MockFns
	}

	wrapper struct {
		t *testing.T
	}
	adminv2 struct {
		partitionservice *adminv2mocks.PartitionServiceClient
		tenantservice    *adminv2mocks.TenantServiceClient
		tokenservice     *adminv2mocks.TokenServiceClient
	}

	Adminv2MockFns struct {
		Partition func(m *mock.Mock)
		Tenant    func(m *mock.Mock)
		Token     func(m *mock.Mock)
	}
	apiv2 struct {
		healthservice    *apiv2mocks.HealthServiceClient
		ipservice        *apiv2mocks.IPServiceClient
		methodservice    *apiv2mocks.MethodServiceClient
		partitionservice *apiv2mocks.PartitionServiceClient
		projectservice   *apiv2mocks.ProjectServiceClient
		tenantservice    *apiv2mocks.TenantServiceClient
		tokenservice     *apiv2mocks.TokenServiceClient
		userservice      *apiv2mocks.UserServiceClient
		versionservice   *apiv2mocks.VersionServiceClient
	}

	Apiv2MockFns struct {
		Health    func(m *mock.Mock)
		IP        func(m *mock.Mock)
		Method    func(m *mock.Mock)
		Partition func(m *mock.Mock)
		Project   func(m *mock.Mock)
		Tenant    func(m *mock.Mock)
		Token     func(m *mock.Mock)
		User      func(m *mock.Mock)
		Version   func(m *mock.Mock)
	}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
		adminv2service: w.Adminv2(fns.Adminv2Mocks),
		apiv2service:   w.Apiv2(fns.Apiv2Mocks),
	}
}

func (c *client) Adminv2() apiclient.Adminv2 {
	return c.adminv2service
}
func (c *client) Apiv2() apiclient.Apiv2 {
	return c.apiv2service
}

func (w wrapper) Adminv2(fns *Adminv2MockFns) *adminv2 {
	return newadminv2(w.t, fns)
}

func newadminv2(t *testing.T, fns *Adminv2MockFns) *adminv2 {
	a := &adminv2{
		partitionservice: adminv2mocks.NewPartitionServiceClient(t),
		tenantservice:    adminv2mocks.NewTenantServiceClient(t),
		tokenservice:     adminv2mocks.NewTokenServiceClient(t),
	}

	if fns != nil {
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}

	}

	return a
}

func (c *adminv2) Partition() adminv2connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *adminv2) Tenant() adminv2connect.TenantServiceClient {
	return c.tenantservice
}
func (c *adminv2) Token() adminv2connect.TokenServiceClient {
	return c.tokenservice
}

func (w wrapper) Apiv2(fns *Apiv2MockFns) *apiv2 {
	return newapiv2(w.t, fns)
}

func newapiv2(t *testing.T, fns *Apiv2MockFns) *apiv2 {
	a := &apiv2{
		healthservice:    apiv2mocks.NewHealthServiceClient(t),
		ipservice:        apiv2mocks.NewIPServiceClient(t),
		methodservice:    apiv2mocks.NewMethodServiceClient(t),
		partitionservice: apiv2mocks.NewPartitionServiceClient(t),
		projectservice:   apiv2mocks.NewProjectServiceClient(t),
		tenantservice:    apiv2mocks.NewTenantServiceClient(t),
		tokenservice:     apiv2mocks.NewTokenServiceClient(t),
		userservice:      apiv2mocks.NewUserServiceClient(t),
		versionservice:   apiv2mocks.NewVersionServiceClient(t),
	}

	if fns != nil {
		if fns.Health != nil {
			fns.Health(&a.healthservice.Mock)
		}
		if fns.IP != nil {
			fns.IP(&a.ipservice.Mock)
		}
		if fns.Method != nil {
			fns.Method(&a.methodservice.Mock)
		}
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}
		if fns.User != nil {
			fns.User(&a.userservice.Mock)
		}
		if fns.Version != nil {
			fns.Version(&a.versionservice.Mock)
		}

	}

	return a
}

func (c *apiv2) Health() apiv2connect.HealthServiceClient {
	return c.healthservice
}
func (c *apiv2) IP() apiv2connect.IPServiceClient {
	return c.ipservice
}
func (c *apiv2) Method() apiv2connect.MethodServiceClient {
	return c.methodservice
}
func (c *apiv2) Partition() apiv2connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *apiv2) Project() apiv2connect.ProjectServiceClient {
	return c.projectservice
}
func (c *apiv2) Tenant() apiv2connect.TenantServiceClient {
	return c.tenantservice
}
func (c *apiv2) Token() apiv2connect.TokenServiceClient {
	return c.tokenservice
}
func (c *apiv2) User() apiv2connect.UserServiceClient {
	return c.userservice
}
func (c *apiv2) Version() apiv2connect.VersionServiceClient {
	return c.versionservice
}
