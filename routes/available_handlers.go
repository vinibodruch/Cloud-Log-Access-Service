package routes

import (
	awsHandlers "Cloud-Log-Access-Service/aws/handlers"
	azureHandlers "Cloud-Log-Access-Service/azure/handlers"
	// gcpHandlers "Cloud-Log-Access-Service/gcp/handlers" // Placeholder for GCP
)

// AvailableHandlers is a struct that groups all available handler interfaces in the application.
type AvailableHandlers struct {
	S3        awsHandlers.S3Handler
	AzureBlob azureHandlers.BlobHandler
	// GCPStorage gcpHandlers.StorageHandler // Example: Add the GCP handler here
}
