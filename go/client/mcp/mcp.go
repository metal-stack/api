package mcp

import (
	"github.com/metal-stack/api/go/client"
	"github.com/metal-stack/api/go/metalstack/admin/v2/adminv2mcp"
	"github.com/metal-stack/api/go/metalstack/api/v2/apiv2mcp"
	"github.com/redpanda-data/protoc-gen-go-mcp/pkg/runtime"
)

// TODO Generate

func ForwardToApiClients(c client.Apiv2, mcps runtime.MCPServer, opts []runtime.Option) {
	// API Clients
	apiv2mcp.ForwardToAuditServiceClient(mcps, c.Audit(), opts...)
	apiv2mcp.ForwardToFilesystemServiceClient(mcps, c.Filesystem(), opts...)
	apiv2mcp.ForwardToHealthServiceClient(mcps, c.Health(), opts...)
	apiv2mcp.ForwardToImageServiceClient(mcps, c.Image(), opts...)
	apiv2mcp.ForwardToIPServiceClient(mcps, c.IP(), opts...)
	apiv2mcp.ForwardToMachineServiceClient(mcps, c.Machine(), opts...)
	apiv2mcp.ForwardToMethodServiceClient(mcps, c.Method(), opts...)
	apiv2mcp.ForwardToNetworkServiceClient(mcps, c.Network(), opts...)
	apiv2mcp.ForwardToPartitionServiceClient(mcps, c.Partition(), opts...)
	apiv2mcp.ForwardToProjectServiceClient(mcps, c.Project(), opts...)
	apiv2mcp.ForwardToSizeImageConstraintServiceClient(mcps, c.SizeImageConstraint(), opts...)
	apiv2mcp.ForwardToSizeReservationServiceClient(mcps, c.SizeReservation(), opts...)
	apiv2mcp.ForwardToSizeServiceClient(mcps, c.Size(), opts...)
	apiv2mcp.ForwardToTenantServiceClient(mcps, c.Tenant(), opts...)
	apiv2mcp.ForwardToTokenServiceClient(mcps, c.Token(), opts...)
	apiv2mcp.ForwardToUserServiceClient(mcps, c.User(), opts...)
	apiv2mcp.ForwardToVersionServiceClient(mcps, c.Version(), opts...)

}
func ForwardToAdminClients(c client.Adminv2, mcps runtime.MCPServer, opts []runtime.Option) {
	// Admin Clients
	adminv2mcp.ForwardToAuditServiceClient(mcps, c.Audit(), opts...)
	adminv2mcp.ForwardToComponentServiceClient(mcps, c.Component(), opts...)
	adminv2mcp.ForwardToFilesystemServiceClient(mcps, c.Filesystem(), opts...)
	adminv2mcp.ForwardToImageServiceClient(mcps, c.Image(), opts...)
	adminv2mcp.ForwardToIPServiceClient(mcps, c.IP(), opts...)
	adminv2mcp.ForwardToMachineServiceClient(mcps, c.Machine(), opts...)
	adminv2mcp.ForwardToNetworkServiceClient(mcps, c.Network(), opts...)
	adminv2mcp.ForwardToPartitionServiceClient(mcps, c.Partition(), opts...)
	adminv2mcp.ForwardToProjectServiceClient(mcps, c.Project(), opts...)
	adminv2mcp.ForwardToSizeImageConstraintServiceClient(mcps, c.SizeImageConstraint(), opts...)
	adminv2mcp.ForwardToSizeReservationServiceClient(mcps, c.SizeReservation(), opts...)
	adminv2mcp.ForwardToSizeServiceClient(mcps, c.Size(), opts...)
	adminv2mcp.ForwardToSwitchServiceClient(mcps, c.Switch(), opts...)
	adminv2mcp.ForwardToTaskServiceClient(mcps, c.Task(), opts...)
	adminv2mcp.ForwardToTenantServiceClient(mcps, c.Tenant(), opts...)
	adminv2mcp.ForwardToTokenServiceClient(mcps, c.Token(), opts...)
	adminv2mcp.ForwardToVPNServiceClient(mcps, c.VPN(), opts...)
}
