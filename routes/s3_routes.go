package routes

import (
	"github.com/gin-gonic/gin"
)

// S3Routes configura as rotas relacionadas ao AWS S3.
func S3Routes(router *gin.RouterGroup, h *AvailableHandlers) {
	s3Handler := h.S3 // Obtém a interface S3Handler da coleção

	s3 := router.Group("/s3")
	{
		s3.GET("/bucket/:bucketName/objects", s3Handler.ListBucketObjects)
		s3.GET("/bucket/:bucketName/object/:objectKey", s3Handler.GetObjectFromBucket)
	}
}
