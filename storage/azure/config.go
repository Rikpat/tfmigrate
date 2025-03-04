package azure

import (
	"errors"
	"os"

	"github.com/minamijoyo/tfmigrate/storage"
)

type Config struct {
	// Storage Config
	// ResourceGroup string `hcl:"resource_group_name"`
	AccountName   string `hcl:"storage_account_name,optional"`
	ContainerName string `hcl:"container_name,optional"`
	Key           string `hcl:"key,optional"`
	AccessKey     string `hcl:"access_key,optional"`
	// Access
	// UseMSI         bool `hcl:use_msi,optional`
	// ClientID       bool `hcl:client_id,optional` // Principal, Managed Identity, Workload Identity
}

// Config implements a storage.Config.
var _ storage.Config = (*Config)(nil)

// NewStorage returns a new instance of storage.Storage.
func (c *Config) NewStorage() (storage.Storage, error) {
	// As this doesn't support cli params for storage, load everything from env
	if c.AccountName == "" {
		c.AccountName = os.Getenv("TFMIGRATE_AZURERM_STORAGE_ACCOUNT_NAME")
	}
	if c.ContainerName == "" {
		c.ContainerName = os.Getenv("TFMIGRATE_AZURERM_CONTAINER_NAME")
	}
	if c.Key == "" {
		c.Key = os.Getenv("TFMIGRATE_AZURERM_FILE_NAME")
	}
	if c.AccessKey == "" {
		c.AccessKey = os.Getenv("TFMIGRATE_AZURERM_ACCESS_KEY")
	}
	if c.AccountName == "" || c.ContainerName == "" || c.Key == "" {
		return nil, errors.New("missing required input")
	}
	return NewStorage(c, nil)
}
