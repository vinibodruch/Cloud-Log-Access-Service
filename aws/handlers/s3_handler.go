package handlers

import (
	"net/http"

	"Cloud-Log-Access-Service/aws/services"

	"github.com/gin-gonic/gin"
)

// S3Handler is the interface for handlers that interact with AWS S3.
type S3Handler interface {
	ListBucketObjects(c *gin.Context)
	GetObjectFromBucket(c *gin.Context)
}

// s3HandlerImpl implements the S3Handler interface.
type s3HandlerImpl struct {
	s3Service services.S3Service
}

// NewS3Handler creates a new instance of S3Handler.
func NewS3Handler(s3Service services.S3Service) S3Handler {
	return &s3HandlerImpl{
		s3Service: s3Service,
	}
}

// ListBucketObjects godoc
// @Summary List objects in an S3 bucket
// @Description Returns a list of objects in a specified S3 bucket
// @Tags s3
// @Accept json
// @Produce json
// @Param bucketName path string true "S3 bucket name"
// @Success 200 {array} object "List of S3 objects"
// @Failure 400 {object} map[string]string "Invalid request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /s3/bucket/{bucketName}/objects [get]
func (h *s3HandlerImpl) ListBucketObjects(c *gin.Context) {
	bucketName := c.Param("bucketName")
	if bucketName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bucket name is required"})
		return
	}

	objects, err := h.s3Service.ListObjectsInBucket(bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing bucket objects", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objects)
}

// GetObjectFromBucket godoc
// @Summary Get an object from an S3 bucket
// @Description Returns the content of a specific object from an S3 bucket
// @Tags s3
// @Accept json
// @Produce octet-stream
// @Param bucketName path string true "S3 bucket name"
// @Param objectKey path string true "S3 object key"
// @Success 200 {string} string "Object content"
// @Failure 400 {object} map[string]string "Invalid request error"
// @Failure 404 {object} map[string]string "Object not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /s3/bucket/{bucketName}/object/{objectKey} [get]
func (h *s3HandlerImpl) GetObjectFromBucket(c *gin.Context) {
	bucketName := c.Param("bucketName")
	objectKey := c.Param("objectKey")

	if bucketName == "" || objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bucket name and object key are required"})
		return
	}

	content, err := h.s3Service.GetObjectFromBucket(bucketName, objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting object from bucket", "details": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", content)
}
