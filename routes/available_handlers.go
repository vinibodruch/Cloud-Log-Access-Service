package routes

import (
	awsHandlers "Cloud-Log-Access-Service/aws/handlers"
	azureHandlers "Cloud-Log-Access-Service/azure/handlers"
	// gcpHandlers "Cloud-Log-Access-Service/gcp/handlers" // Placeholder para GCP
)

// AvailableHandlers é uma struct que agrupa todas as interfaces de handler disponíveis na aplicação.
type AvailableHandlers struct {
	S3        awsHandlers.S3Handler
	AzureBlob azureHandlers.BlobHandler
	// GCPStorage gcpHandlers.StorageHandler // Exemplo: Adicione o handler GCP aqui
}
