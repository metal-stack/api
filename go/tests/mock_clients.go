// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	apiclient "github.com/metal-stack/api/go/client"
	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
	adminv2mocks "github.com/metal-stack/api/go/tests/mocks/metalstack/admin/v2/adminv2connect"
	apiv2mocks "github.com/metal-stack/api/go/tests/mocks/metalstack/api/v2/apiv2connect"
	infrav2mocks "github.com/metal-stack/api/go/tests/mocks/metalstack/infra/v2/infrav2connect"

	"github.com/stretchr/testify/mock"
)

type (
	client struct {
		adminv2service *adminv2
		apiv2service   *apiv2
		infrav2service *infrav2
	}

	ClientMockFns struct {
		Adminv2Mocks *Adminv2MockFns
		Apiv2Mocks   *Apiv2MockFns
		Infrav2Mocks *Infrav2MockFns
	}

	wrapper struct {
		t *testing.T
	}
	adminv2 struct {
		filesystemservice *adminv2mocks.FilesystemServiceClient
		imageservice      *adminv2mocks.ImageServiceClient
		ipservice         *adminv2mocks.IPServiceClient
		machineservice    *adminv2mocks.MachineServiceClient
		networkservice    *adminv2mocks.NetworkServiceClient
		partitionservice  *adminv2mocks.PartitionServiceClient
		projectservice    *adminv2mocks.ProjectServiceClient
		sizeservice       *adminv2mocks.SizeServiceClient
		switchservice     *adminv2mocks.SwitchServiceClient
		tenantservice     *adminv2mocks.TenantServiceClient
		tokenservice      *adminv2mocks.TokenServiceClient
		vpnservice        *adminv2mocks.VPNServiceClient
	}

	Adminv2MockFns struct {
		Filesystem func(m *mock.Mock)
		Image      func(m *mock.Mock)
		IP         func(m *mock.Mock)
		Machine    func(m *mock.Mock)
		Network    func(m *mock.Mock)
		Partition  func(m *mock.Mock)
		Project    func(m *mock.Mock)
		Size       func(m *mock.Mock)
		Switch     func(m *mock.Mock)
		Tenant     func(m *mock.Mock)
		Token      func(m *mock.Mock)
		VPN        func(m *mock.Mock)
	}
	apiv2 struct {
		filesystemservice *apiv2mocks.FilesystemServiceClient
		healthservice     *apiv2mocks.HealthServiceClient
		imageservice      *apiv2mocks.ImageServiceClient
		ipservice         *apiv2mocks.IPServiceClient
		machineservice    *apiv2mocks.MachineServiceClient
		methodservice     *apiv2mocks.MethodServiceClient
		networkservice    *apiv2mocks.NetworkServiceClient
		partitionservice  *apiv2mocks.PartitionServiceClient
		projectservice    *apiv2mocks.ProjectServiceClient
		sizeservice       *apiv2mocks.SizeServiceClient
		tenantservice     *apiv2mocks.TenantServiceClient
		tokenservice      *apiv2mocks.TokenServiceClient
		userservice       *apiv2mocks.UserServiceClient
		versionservice    *apiv2mocks.VersionServiceClient
	}

	Apiv2MockFns struct {
		Filesystem func(m *mock.Mock)
		Health     func(m *mock.Mock)
		Image      func(m *mock.Mock)
		IP         func(m *mock.Mock)
		Machine    func(m *mock.Mock)
		Method     func(m *mock.Mock)
		Network    func(m *mock.Mock)
		Partition  func(m *mock.Mock)
		Project    func(m *mock.Mock)
		Size       func(m *mock.Mock)
		Tenant     func(m *mock.Mock)
		Token      func(m *mock.Mock)
		User       func(m *mock.Mock)
		Version    func(m *mock.Mock)
	}
	infrav2 struct {
		bmcservice    *infrav2mocks.BMCServiceClient
		switchservice *infrav2mocks.SwitchServiceClient
	}

	Infrav2MockFns struct {
		BMC    func(m *mock.Mock)
		Switch func(m *mock.Mock)
	}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
		adminv2service: w.Adminv2(fns.Adminv2Mocks),
		apiv2service:   w.Apiv2(fns.Apiv2Mocks),
		infrav2service: w.Infrav2(fns.Infrav2Mocks),
	}
}

func (c *client) Adminv2() apiclient.Adminv2 {
	return c.adminv2service
}
func (c *client) Apiv2() apiclient.Apiv2 {
	return c.apiv2service
}
func (c *client) Infrav2() apiclient.Infrav2 {
	return c.infrav2service
}

func (w wrapper) Adminv2(fns *Adminv2MockFns) *adminv2 {
	return newadminv2(w.t, fns)
}

func newadminv2(t *testing.T, fns *Adminv2MockFns) *adminv2 {
	a := &adminv2{
		filesystemservice: adminv2mocks.NewFilesystemServiceClient(t),
		imageservice:      adminv2mocks.NewImageServiceClient(t),
		ipservice:         adminv2mocks.NewIPServiceClient(t),
		machineservice:    adminv2mocks.NewMachineServiceClient(t),
		networkservice:    adminv2mocks.NewNetworkServiceClient(t),
		partitionservice:  adminv2mocks.NewPartitionServiceClient(t),
		projectservice:    adminv2mocks.NewProjectServiceClient(t),
		sizeservice:       adminv2mocks.NewSizeServiceClient(t),
		switchservice:     adminv2mocks.NewSwitchServiceClient(t),
		tenantservice:     adminv2mocks.NewTenantServiceClient(t),
		tokenservice:      adminv2mocks.NewTokenServiceClient(t),
		vpnservice:        adminv2mocks.NewVPNServiceClient(t),
	}

	if fns != nil {
		if fns.Filesystem != nil {
			fns.Filesystem(&a.filesystemservice.Mock)
		}
		if fns.Image != nil {
			fns.Image(&a.imageservice.Mock)
		}
		if fns.IP != nil {
			fns.IP(&a.ipservice.Mock)
		}
		if fns.Machine != nil {
			fns.Machine(&a.machineservice.Mock)
		}
		if fns.Network != nil {
			fns.Network(&a.networkservice.Mock)
		}
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.Size != nil {
			fns.Size(&a.sizeservice.Mock)
		}
		if fns.Switch != nil {
			fns.Switch(&a.switchservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}
		if fns.VPN != nil {
			fns.VPN(&a.vpnservice.Mock)
		}

	}

	return a
}

func (c *adminv2) Filesystem() adminv2connect.FilesystemServiceClient {
	return c.filesystemservice
}
func (c *adminv2) Image() adminv2connect.ImageServiceClient {
	return c.imageservice
}
func (c *adminv2) IP() adminv2connect.IPServiceClient {
	return c.ipservice
}
func (c *adminv2) Machine() adminv2connect.MachineServiceClient {
	return c.machineservice
}
func (c *adminv2) Network() adminv2connect.NetworkServiceClient {
	return c.networkservice
}
func (c *adminv2) Partition() adminv2connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *adminv2) Project() adminv2connect.ProjectServiceClient {
	return c.projectservice
}
func (c *adminv2) Size() adminv2connect.SizeServiceClient {
	return c.sizeservice
}
func (c *adminv2) Switch() adminv2connect.SwitchServiceClient {
	return c.switchservice
}
func (c *adminv2) Tenant() adminv2connect.TenantServiceClient {
	return c.tenantservice
}
func (c *adminv2) Token() adminv2connect.TokenServiceClient {
	return c.tokenservice
}
func (c *adminv2) VPN() adminv2connect.VPNServiceClient {
	return c.vpnservice
}

func (w wrapper) Apiv2(fns *Apiv2MockFns) *apiv2 {
	return newapiv2(w.t, fns)
}

func newapiv2(t *testing.T, fns *Apiv2MockFns) *apiv2 {
	a := &apiv2{
		filesystemservice: apiv2mocks.NewFilesystemServiceClient(t),
		healthservice:     apiv2mocks.NewHealthServiceClient(t),
		imageservice:      apiv2mocks.NewImageServiceClient(t),
		ipservice:         apiv2mocks.NewIPServiceClient(t),
		machineservice:    apiv2mocks.NewMachineServiceClient(t),
		methodservice:     apiv2mocks.NewMethodServiceClient(t),
		networkservice:    apiv2mocks.NewNetworkServiceClient(t),
		partitionservice:  apiv2mocks.NewPartitionServiceClient(t),
		projectservice:    apiv2mocks.NewProjectServiceClient(t),
		sizeservice:       apiv2mocks.NewSizeServiceClient(t),
		tenantservice:     apiv2mocks.NewTenantServiceClient(t),
		tokenservice:      apiv2mocks.NewTokenServiceClient(t),
		userservice:       apiv2mocks.NewUserServiceClient(t),
		versionservice:    apiv2mocks.NewVersionServiceClient(t),
	}

	if fns != nil {
		if fns.Filesystem != nil {
			fns.Filesystem(&a.filesystemservice.Mock)
		}
		if fns.Health != nil {
			fns.Health(&a.healthservice.Mock)
		}
		if fns.Image != nil {
			fns.Image(&a.imageservice.Mock)
		}
		if fns.IP != nil {
			fns.IP(&a.ipservice.Mock)
		}
		if fns.Machine != nil {
			fns.Machine(&a.machineservice.Mock)
		}
		if fns.Method != nil {
			fns.Method(&a.methodservice.Mock)
		}
		if fns.Network != nil {
			fns.Network(&a.networkservice.Mock)
		}
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.Size != nil {
			fns.Size(&a.sizeservice.Mock)
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

func (c *apiv2) Filesystem() apiv2connect.FilesystemServiceClient {
	return c.filesystemservice
}
func (c *apiv2) Health() apiv2connect.HealthServiceClient {
	return c.healthservice
}
func (c *apiv2) Image() apiv2connect.ImageServiceClient {
	return c.imageservice
}
func (c *apiv2) IP() apiv2connect.IPServiceClient {
	return c.ipservice
}
func (c *apiv2) Machine() apiv2connect.MachineServiceClient {
	return c.machineservice
}
func (c *apiv2) Method() apiv2connect.MethodServiceClient {
	return c.methodservice
}
func (c *apiv2) Network() apiv2connect.NetworkServiceClient {
	return c.networkservice
}
func (c *apiv2) Partition() apiv2connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *apiv2) Project() apiv2connect.ProjectServiceClient {
	return c.projectservice
}
func (c *apiv2) Size() apiv2connect.SizeServiceClient {
	return c.sizeservice
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

func (w wrapper) Infrav2(fns *Infrav2MockFns) *infrav2 {
	return newinfrav2(w.t, fns)
}

func newinfrav2(t *testing.T, fns *Infrav2MockFns) *infrav2 {
	a := &infrav2{
		bmcservice:    infrav2mocks.NewBMCServiceClient(t),
		switchservice: infrav2mocks.NewSwitchServiceClient(t),
	}

	if fns != nil {
		if fns.BMC != nil {
			fns.BMC(&a.bmcservice.Mock)
		}
		if fns.Switch != nil {
			fns.Switch(&a.switchservice.Mock)
		}

	}

	return a
}

func (c *infrav2) BMC() infrav2connect.BMCServiceClient {
	return c.bmcservice
}
func (c *infrav2) Switch() infrav2connect.SwitchServiceClient {
	return c.switchservice
}
