package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Cloud-Log-Access-Service/azure/services"
)

// BlobHandler é a interface para handlers do Azure Blob Storage.
type BlobHandler interface {
	ListContainers(c *gin.Context)
	ListBlobsInContainer(c *gin.Context)
}

// blobHandlerImpl implementa a interface BlobHandler.
type blobHandlerImpl struct {
	blobService services.BlobService
}

// NewBlobHandler cria uma nova instância de BlobHandler.
func NewBlobHandler(blobService services.BlobService) BlobHandler {
	return &blobHandlerImpl{
		blobService: blobService,
	}
}

// ListContainers lista todos os containers no Azure Blob Storage.
func (h *blobHandlerImpl) ListContainers(c *gin.Context) {
	containers, err := h.blobService.ListContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list containers", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List containers (Azure)", "containers": containers})
}

// ListBlobsInContainer lista todos os blobs em um container específico do Azure Blob Storage.
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
