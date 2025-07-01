package handlers

import (
	"net/http"

	"Cloud-Log-Access-Service/azure/services"

	"github.com/gin-gonic/gin"
)

// BlobHandler is the interface for Azure Blob Storage handlers.
type BlobHandler interface {
	ListContainers(c *gin.Context)
	ListBlobsInContainer(c *gin.Context)
}

// blobHandlerImpl implements the BlobHandler interface.
type blobHandlerImpl struct {
	blobService services.BlobService
}

// NewBlobHandler creates a new instance of BlobHandler.
func NewBlobHandler(blobService services.BlobService) BlobHandler {
	return &blobHandlerImpl{
		blobService: blobService,
	}
}

// ListContainers lists all containers in Azure Blob Storage.
func (h *blobHandlerImpl) ListContainers(c *gin.Context) {
	containers, err := h.blobService.ListContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list containers", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List containers (Azure)", "containers": containers})
}

// ListBlobsInContainer lists all blobs in a specific Azure Blob Storage container.
func (h *blobHandlerImpl) ListBlobsInContainer(c *gin.Context) {
	containerName := c.Param("containerName")
	if containerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container name is required"})
		return
	}
	blobs, err := h.blobService.ListBlobsInContainer(containerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list blobs", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List blobs in container " + containerName + " (Azure)", "blobs": blobs})
}
