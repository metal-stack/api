// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"sync"

	compress "github.com/klauspost/connect-compress/v2"

	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
)

type (
	Client interface {
		Adminv2() Adminv2
		Apiv2() Apiv2
		Infrav2() Infrav2
	}
	client struct {
		config DialConfig

		sync.Mutex
	}
	Adminv2 interface {
		Filesystem() adminv2connect.FilesystemServiceClient
		Image() adminv2connect.ImageServiceClient
		IP() adminv2connect.IPServiceClient
		Network() adminv2connect.NetworkServiceClient
		Partition() adminv2connect.PartitionServiceClient
		Tenant() adminv2connect.TenantServiceClient
		Token() adminv2connect.TokenServiceClient
	}

	adminv2 struct {
		filesystemservice adminv2connect.FilesystemServiceClient
		imageservice      adminv2connect.ImageServiceClient
		ipservice         adminv2connect.IPServiceClient
		networkservice    adminv2connect.NetworkServiceClient
		partitionservice  adminv2connect.PartitionServiceClient
		tenantservice     adminv2connect.TenantServiceClient
		tokenservice      adminv2connect.TokenServiceClient
	}

	Apiv2 interface {
		Filesystem() apiv2connect.FilesystemServiceClient
		Health() apiv2connect.HealthServiceClient
		Image() apiv2connect.ImageServiceClient
		IP() apiv2connect.IPServiceClient
		Method() apiv2connect.MethodServiceClient
		Network() apiv2connect.NetworkServiceClient
		Partition() apiv2connect.PartitionServiceClient
		Project() apiv2connect.ProjectServiceClient
		Tenant() apiv2connect.TenantServiceClient
		Token() apiv2connect.TokenServiceClient
		User() apiv2connect.UserServiceClient
		Version() apiv2connect.VersionServiceClient
	}

	apiv2 struct {
		filesystemservice apiv2connect.FilesystemServiceClient
		healthservice     apiv2connect.HealthServiceClient
		imageservice      apiv2connect.ImageServiceClient
		ipservice         apiv2connect.IPServiceClient
		methodservice     apiv2connect.MethodServiceClient
		networkservice    apiv2connect.NetworkServiceClient
		partitionservice  apiv2connect.PartitionServiceClient
		projectservice    apiv2connect.ProjectServiceClient
		tenantservice     apiv2connect.TenantServiceClient
		tokenservice      apiv2connect.TokenServiceClient
		userservice       apiv2connect.UserServiceClient
		versionservice    apiv2connect.VersionServiceClient
	}

	Infrav2 interface {
		BMC() infrav2connect.BMCServiceClient
	}

	infrav2 struct {
		bmcservice infrav2connect.BMCServiceClient
	}
)

func New(config DialConfig) (Client, error) {
	exp, err := getExpiresAt(config.Token)
	if err != nil {
		return nil, err
	}
	config.expiresAt = exp
	c := &client{
		config: config,
	}

	go c.renewToken()

	return c, nil
}

func (c *client) Adminv2() Adminv2 {
	a := &adminv2{
		filesystemservice: adminv2connect.NewFilesystemServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		imageservice: adminv2connect.NewImageServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		ipservice: adminv2connect.NewIPServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		networkservice: adminv2connect.NewNetworkServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		partitionservice: adminv2connect.NewPartitionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: adminv2connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tokenservice: adminv2connect.NewTokenServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
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
func (c *adminv2) Network() adminv2connect.NetworkServiceClient {
	return c.networkservice
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

func (c *client) Apiv2() Apiv2 {
	a := &apiv2{
		filesystemservice: apiv2connect.NewFilesystemServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		healthservice: apiv2connect.NewHealthServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		imageservice: apiv2connect.NewImageServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		ipservice: apiv2connect.NewIPServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		methodservice: apiv2connect.NewMethodServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		networkservice: apiv2connect.NewNetworkServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		partitionservice: apiv2connect.NewPartitionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		projectservice: apiv2connect.NewProjectServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: apiv2connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tokenservice: apiv2connect.NewTokenServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		userservice: apiv2connect.NewUserServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		versionservice: apiv2connect.NewVersionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
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

func (c *client) Infrav2() Infrav2 {
	a := &infrav2{
		bmcservice: infrav2connect.NewBMCServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *infrav2) BMC() infrav2connect.BMCServiceClient {
	return c.bmcservice
}
