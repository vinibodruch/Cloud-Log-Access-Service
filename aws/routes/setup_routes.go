package routes

import (
	"Cloud-Log-Access-Service/aws/handlers"

	"github.com/gin-gonic/gin"
)

func S3Routes(router *gin.RouterGroup, s3Handler handlers.S3Handler) {
	s3 := router.Group("/s3/bucket/:bucketName")
	{
		s3.GET("/objects", s3Handler.ListBucketObjects)
		s3.GET("/object/:objectKey", s3Handler.GetBucketObject)
	}
}
