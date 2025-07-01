package routes

import (
	"github.com/gin-gonic/gin"
)

// S3Routes configures the routes related to AWS S3.
func S3Routes(router *gin.RouterGroup, h *AvailableHandlers) {
	s3Handler := h.S3 // Gets the S3Handler interface from the collection

	s3 := router.Group("/s3")
	{
		s3.GET("/bucket/:bucketName/objects", s3Handler.ListBucketObjects)
		s3.GET("/bucket/:bucketName/object/:objectKey", s3Handler.GetObjectFromBucket)
	}
}
