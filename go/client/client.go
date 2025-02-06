// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	compress "github.com/klauspost/connect-compress/v2"

	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
)

type (
	Client interface {
		Adminv2() Adminv2
		Apiv2() Apiv2
	}
	client struct {
		config DialConfig
	}
	Adminv2 interface {
		Partition() adminv2connect.PartitionServiceClient
		Tenant() adminv2connect.TenantServiceClient
		Token() adminv2connect.TokenServiceClient
	}

	adminv2 struct {
		partitionservice adminv2connect.PartitionServiceClient
		tenantservice    adminv2connect.TenantServiceClient
		tokenservice     adminv2connect.TokenServiceClient
	}

	Apiv2 interface {
		Health() apiv2connect.HealthServiceClient
		IP() apiv2connect.IPServiceClient
		Method() apiv2connect.MethodServiceClient
		Partition() apiv2connect.PartitionServiceClient
		Project() apiv2connect.ProjectServiceClient
		Tenant() apiv2connect.TenantServiceClient
		Token() apiv2connect.TokenServiceClient
		User() apiv2connect.UserServiceClient
		Version() apiv2connect.VersionServiceClient
	}

	apiv2 struct {
		healthservice    apiv2connect.HealthServiceClient
		ipservice        apiv2connect.IPServiceClient
		methodservice    apiv2connect.MethodServiceClient
		partitionservice apiv2connect.PartitionServiceClient
		projectservice   apiv2connect.ProjectServiceClient
		tenantservice    apiv2connect.TenantServiceClient
		tokenservice     apiv2connect.TokenServiceClient
		userservice      apiv2connect.UserServiceClient
		versionservice   apiv2connect.VersionServiceClient
	}
)

func New(config DialConfig) Client {
	return &client{
		config: config,
	}
}

func (c client) Adminv2() Adminv2 {
	a := &adminv2{
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

func (c *adminv2) Partition() adminv2connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *adminv2) Tenant() adminv2connect.TenantServiceClient {
	return c.tenantservice
}
func (c *adminv2) Token() adminv2connect.TokenServiceClient {
	return c.tokenservice
}

func (c client) Apiv2() Apiv2 {
	a := &apiv2{
		healthservice: apiv2connect.NewHealthServiceClient(
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
