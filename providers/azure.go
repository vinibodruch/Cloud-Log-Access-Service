package providers

import (
	"fmt"
	"time"
)

// init function is called automatically when the package is imported.
// It registers the AzureCloudProvider with the main provider factory.
func init() {
	RegisterProvider("azure", func() CloudProvider {
		return &AzureCloudProvider{}
	})
}

// AzureCloudProvider is a dummy implementation of CloudProvider for Azure Blob Storage.
// In a real application, this would use the Azure SDK (e.g., github.com/Azure/azure-sdk-for-go/sdk/storage/azblob).
type AzureCloudProvider struct{}

// ListFiles simulates listing files in an Azure Blob Storage container.
func (p *AzureCloudProvider) ListFiles(bucket string) ([]string, error) {
	// Simulate fetching files from a dummy data source.
	if bucket == "log-saas-azure" {
		return []string{"azure_web_log_1.log", "azure_func_log_2.log", "azure_sys.log"}, nil
	}
	return nil, fmt.Errorf("Azure container '%s' not found or accessible", bucket)
}

// DownloadFile simulates downloading a file from an Azure Blob Storage container.
func (p *AzureCloudProvider) DownloadFile(bucket, filename string) ([]byte, error) {
	// Simulate file content.
	if bucket == "log-saas-azure" {
		switch filename {
		case "azure_web_log_1.log":
			return []byte("Content of Azure web log 1."), nil
		case "azure_func_log_2.log":
			return []byte("Content of Azure function log 2."), nil
		case "azure_sys.log":
			return []byte("Sample Azure system log: [WARNING] Disk space low on VM."), nil
		default:
			return nil, fmt.Errorf("Azure file '%s' not found in container '%s'", filename, bucket)
		}
	}
	return nil, fmt.Errorf("Azure container '%s' not found or accessible", bucket)
}

// GenerateSignedURL simulates generating a pre-signed URL (SAS token) for an Azure Blob.
func (p *AzureCloudProvider) GenerateSignedURL(bucket, filename string, expiry time.Duration) (string, error) {
	if bucket == "log-saas-azure" {
		// In a real scenario, this would involve creating a SAS token using Azure SDK.
		return fmt.Sprintf("https://mock-azure-blob.core.windows.net/%s/%s?sig=mock_azure_sas&se=%d", bucket, filename, time.Now().Add(expiry).Unix()), nil
	}
	return "", fmt.Errorf("Azure container '%s' not found or accessible", bucket)
}
