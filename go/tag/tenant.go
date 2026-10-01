package tag

const (
	// ProviderTenant if attached to a tenant with a non empty value,
	// this tenant is treated as the provider tenant.
	ProviderTenant = "tenant.metal-stack.io/provider"
	// MachineBootstrapperTenant if attached to a tenant (empty value is fine),
	// this tenant can be used for issuing machine tokens through the infra api.
	// this is intended for pixie to create tokens for the metal-hammer.
	MachineBootstrapperTenant = "tenant.metal-stack.io/machine-bootstrapper"
)
