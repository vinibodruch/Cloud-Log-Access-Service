package routes

import (
	"github.com/gin-gonic/gin"
)

// AzureRoutes configura as rotas relacionadas ao Azure Blob Storage.
func AzureRoutes(router *gin.RouterGroup, h *AvailableHandlers) {
	azureBlobHandler := h.AzureBlob // Obtém a interface BlobHandler da coleção

	azure := router.Group("/azure")
	{
		azure.GET("/containers", azureBlobHandler.ListContainers)
		azure.GET("/containers/:containerName/blobs", azureBlobHandler.ListBlobsInContainer)
	}
}
