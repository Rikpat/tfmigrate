package azure

import "github.com/minamijoyo/tfmigrate/storage"

type Config struct {
	// Storage Config
	// ResourceGroup string `hcl:"resource_group_name"` // Not Needed
	AccountName   string `hcl:"storage_account_name"`
	ContainerName string `hcl:"container_name"`
	Key           string `hcl:"key"`
	// Access
	// UseAzureADAuth bool `hcl:use_azuread_auth,optional`
	// UseMSI         bool `hcl:use_msi,optional`
	// ClientID       bool `hcl:client_id,optional` // Principal, Managed Identity, Workload Identity
}

// Config implements a storage.Config.
var _ storage.Config = (*Config)(nil)

// NewStorage returns a new instance of storage.Storage.
func (c *Config) NewStorage() (storage.Storage, error) {
	return NewStorage(c, nil)
}
