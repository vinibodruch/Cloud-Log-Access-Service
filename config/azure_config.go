package config

import (
	//"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// AzureConfig contains the configuration and client for Azure Blob Storage.
type AzureConfig struct {
	AccountName string
	Client      *azblob.Client
}

// LoadAzureConfig loads the configuration and initializes the Azure Blob client.
func LoadAzureConfig() AzureConfig {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	if accountName == "" {
		log.Printf("AZURE_STORAGE_ACCOUNT_NAME not set. Skipping Azure Blob configuration.")
		return AzureConfig{}
	}

	// DefaultAzureCredential tries to use several credential sources (environment variables, Azure CLI, Managed Identity, etc.)
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain Azure credentials: %v", err)
	}

	serviceURL := "https://" + accountName + ".blob.core.windows.net/"
	client, err := azblob.NewClient(serviceURL, credential, nil)
	if err != nil {
		log.Fatalf("Failed to create Azure Blob client: %v", err)
	}

	return AzureConfig{
		AccountName: accountName,
		Client:      client,
	}
}

// GCPConfig is a placeholder for GCP configuration.
type GCPConfig struct {
	// Define GCP configurations here
}

// Implement func LoadGCPConfig() GCPConfig to load GCP configurations.
