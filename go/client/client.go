// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"context"
	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2connect"
	"github.com/metal-stack/api/go/metalstack/infra/v2/infrav2connect"
)

type (
	Client interface {
		Adminv2() Adminv2
		Apiv2() Apiv2
		Infrav2() Infrav2

		Ping(context.Context, *PingConfig)
	}
	Adminv2 interface {
		Audit() adminv2connect.AuditServiceClient
		Component() adminv2connect.ComponentServiceClient
		Filesystem() adminv2connect.FilesystemServiceClient
		Image() adminv2connect.ImageServiceClient
		IP() adminv2connect.IPServiceClient
		Machine() adminv2connect.MachineServiceClient
		Network() adminv2connect.NetworkServiceClient
		Partition() adminv2connect.PartitionServiceClient
		Project() adminv2connect.ProjectServiceClient
		Size() adminv2connect.SizeServiceClient
		SizeImageConstraint() adminv2connect.SizeImageConstraintServiceClient
		SizeReservation() adminv2connect.SizeReservationServiceClient
		Switch() adminv2connect.SwitchServiceClient
		Task() adminv2connect.TaskServiceClient
		Tenant() adminv2connect.TenantServiceClient
		Token() adminv2connect.TokenServiceClient
		VPN() adminv2connect.VPNServiceClient
	}

	adminv2 struct {
		auditservice               adminv2connect.AuditServiceClient
		componentservice           adminv2connect.ComponentServiceClient
		filesystemservice          adminv2connect.FilesystemServiceClient
		imageservice               adminv2connect.ImageServiceClient
		ipservice                  adminv2connect.IPServiceClient
		machineservice             adminv2connect.MachineServiceClient
		networkservice             adminv2connect.NetworkServiceClient
		partitionservice           adminv2connect.PartitionServiceClient
		projectservice             adminv2connect.ProjectServiceClient
		sizeservice                adminv2connect.SizeServiceClient
		sizeimageconstraintservice adminv2connect.SizeImageConstraintServiceClient
		sizereservationservice     adminv2connect.SizeReservationServiceClient
		switchservice              adminv2connect.SwitchServiceClient
		taskservice                adminv2connect.TaskServiceClient
		tenantservice              adminv2connect.TenantServiceClient
		tokenservice               adminv2connect.TokenServiceClient
		vpnservice                 adminv2connect.VPNServiceClient
	}

	Apiv2 interface {
		Audit() apiv2connect.AuditServiceClient
		Filesystem() apiv2connect.FilesystemServiceClient
		Health() apiv2connect.HealthServiceClient
		Image() apiv2connect.ImageServiceClient
		IP() apiv2connect.IPServiceClient
		Machine() apiv2connect.MachineServiceClient
		Method() apiv2connect.MethodServiceClient
		Network() apiv2connect.NetworkServiceClient
		Partition() apiv2connect.PartitionServiceClient
		Project() apiv2connect.ProjectServiceClient
		Size() apiv2connect.SizeServiceClient
		SizeImageConstraint() apiv2connect.SizeImageConstraintServiceClient
		SizeReservation() apiv2connect.SizeReservationServiceClient
		Tenant() apiv2connect.TenantServiceClient
		Token() apiv2connect.TokenServiceClient
		User() apiv2connect.UserServiceClient
		Version() apiv2connect.VersionServiceClient
	}

	apiv2 struct {
		auditservice               apiv2connect.AuditServiceClient
		filesystemservice          apiv2connect.FilesystemServiceClient
		healthservice              apiv2connect.HealthServiceClient
		imageservice               apiv2connect.ImageServiceClient
		ipservice                  apiv2connect.IPServiceClient
		machineservice             apiv2connect.MachineServiceClient
		methodservice              apiv2connect.MethodServiceClient
		networkservice             apiv2connect.NetworkServiceClient
		partitionservice           apiv2connect.PartitionServiceClient
		projectservice             apiv2connect.ProjectServiceClient
		sizeservice                apiv2connect.SizeServiceClient
		sizeimageconstraintservice apiv2connect.SizeImageConstraintServiceClient
		sizereservationservice     apiv2connect.SizeReservationServiceClient
		tenantservice              apiv2connect.TenantServiceClient
		tokenservice               apiv2connect.TokenServiceClient
		userservice                apiv2connect.UserServiceClient
		versionservice             apiv2connect.VersionServiceClient
	}

	Infrav2 interface {
		BMC() infrav2connect.BMCServiceClient
		Boot() infrav2connect.BootServiceClient
		Component() infrav2connect.ComponentServiceClient
		Event() infrav2connect.EventServiceClient
		Switch() infrav2connect.SwitchServiceClient
	}

	infrav2 struct {
		bmcservice       infrav2connect.BMCServiceClient
		bootservice      infrav2connect.BootServiceClient
		componentservice infrav2connect.ComponentServiceClient
		eventservice     infrav2connect.EventServiceClient
		switchservice    infrav2connect.SwitchServiceClient
	}
)

func (c *client) Adminv2() Adminv2 {
	a := &adminv2{
		auditservice:               adminv2connect.NewAuditServiceClient(c.httpClient),
		componentservice:           adminv2connect.NewComponentServiceClient(c.httpClient),
		filesystemservice:          adminv2connect.NewFilesystemServiceClient(c.httpClient),
		imageservice:               adminv2connect.NewImageServiceClient(c.httpClient),
		ipservice:                  adminv2connect.NewIPServiceClient(c.httpClient),
		machineservice:             adminv2connect.NewMachineServiceClient(c.httpClient),
		networkservice:             adminv2connect.NewNetworkServiceClient(c.httpClient),
		partitionservice:           adminv2connect.NewPartitionServiceClient(c.httpClient),
		projectservice:             adminv2connect.NewProjectServiceClient(c.httpClient),
		sizeservice:                adminv2connect.NewSizeServiceClient(c.httpClient),
		sizeimageconstraintservice: adminv2connect.NewSizeImageConstraintServiceClient(c.httpClient),
		sizereservationservice:     adminv2connect.NewSizeReservationServiceClient(c.httpClient),
		switchservice:              adminv2connect.NewSwitchServiceClient(c.httpClient),
		taskservice:                adminv2connect.NewTaskServiceClient(c.httpClient),
		tenantservice:              adminv2connect.NewTenantServiceClient(c.httpClient),
		tokenservice:               adminv2connect.NewTokenServiceClient(c.httpClient),
		vpnservice:                 adminv2connect.NewVPNServiceClient(c.httpClient),
	}
	return a
}

func (c *adminv2) Audit() adminv2connect.AuditServiceClient {
	return c.auditservice
}
func (c *adminv2) Component() adminv2connect.ComponentServiceClient {
	return c.componentservice
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
func (c *adminv2) SizeImageConstraint() adminv2connect.SizeImageConstraintServiceClient {
	return c.sizeimageconstraintservice
}
func (c *adminv2) SizeReservation() adminv2connect.SizeReservationServiceClient {
	return c.sizereservationservice
}
func (c *adminv2) Switch() adminv2connect.SwitchServiceClient {
	return c.switchservice
}
func (c *adminv2) Task() adminv2connect.TaskServiceClient {
	return c.taskservice
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

func (c *client) Apiv2() Apiv2 {
	a := &apiv2{
		auditservice:               apiv2connect.NewAuditServiceClient(c.httpClient),
		filesystemservice:          apiv2connect.NewFilesystemServiceClient(c.httpClient),
		healthservice:              apiv2connect.NewHealthServiceClient(c.httpClient),
		imageservice:               apiv2connect.NewImageServiceClient(c.httpClient),
		ipservice:                  apiv2connect.NewIPServiceClient(c.httpClient),
		machineservice:             apiv2connect.NewMachineServiceClient(c.httpClient),
		methodservice:              apiv2connect.NewMethodServiceClient(c.httpClient),
		networkservice:             apiv2connect.NewNetworkServiceClient(c.httpClient),
		partitionservice:           apiv2connect.NewPartitionServiceClient(c.httpClient),
		projectservice:             apiv2connect.NewProjectServiceClient(c.httpClient),
		sizeservice:                apiv2connect.NewSizeServiceClient(c.httpClient),
		sizeimageconstraintservice: apiv2connect.NewSizeImageConstraintServiceClient(c.httpClient),
		sizereservationservice:     apiv2connect.NewSizeReservationServiceClient(c.httpClient),
		tenantservice:              apiv2connect.NewTenantServiceClient(c.httpClient),
		tokenservice:               apiv2connect.NewTokenServiceClient(c.httpClient),
		userservice:                apiv2connect.NewUserServiceClient(c.httpClient),
		versionservice:             apiv2connect.NewVersionServiceClient(c.httpClient),
	}
	return a
}

func (c *apiv2) Audit() apiv2connect.AuditServiceClient {
	return c.auditservice
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
func (c *apiv2) SizeImageConstraint() apiv2connect.SizeImageConstraintServiceClient {
	return c.sizeimageconstraintservice
}
func (c *apiv2) SizeReservation() apiv2connect.SizeReservationServiceClient {
	return c.sizereservationservice
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
		bmcservice:       infrav2connect.NewBMCServiceClient(c.httpClient),
		bootservice:      infrav2connect.NewBootServiceClient(c.httpClient),
		componentservice: infrav2connect.NewComponentServiceClient(c.httpClient),
		eventservice:     infrav2connect.NewEventServiceClient(c.httpClient),
		switchservice:    infrav2connect.NewSwitchServiceClient(c.httpClient),
	}
	return a
}

func (c *infrav2) BMC() infrav2connect.BMCServiceClient {
	return c.bmcservice
}
func (c *infrav2) Boot() infrav2connect.BootServiceClient {
	return c.bootservice
}
func (c *infrav2) Component() infrav2connect.ComponentServiceClient {
	return c.componentservice
}
func (c *infrav2) Event() infrav2connect.EventServiceClient {
	return c.eventservice
}
func (c *infrav2) Switch() infrav2connect.SwitchServiceClient {
	return c.switchservice
}
