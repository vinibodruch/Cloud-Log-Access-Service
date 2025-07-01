package routes

import (
	"github.com/gin-gonic/gin"
)

// AzureRoutes configures the routes related to Azure Blob Storage.
func AzureRoutes(router *gin.RouterGroup, h *AvailableHandlers) {
	azureBlobHandler := h.AzureBlob // Gets the BlobHandler interface from the collection

	azure := router.Group("/azure")
	{
		azure.GET("/containers", azureBlobHandler.ListContainers)
		azure.GET("/containers/:containerName/blobs", azureBlobHandler.ListBlobsInContainer)
	}
}
