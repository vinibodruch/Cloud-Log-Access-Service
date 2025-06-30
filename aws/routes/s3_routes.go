package routes

import (
	"Cloud-Log-Access-Service/aws/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, s3Handler handlers.S3Handler) {
	api := router.Group("/api/cloud-log-access-services/v1")
	S3Routes(api, s3Handler)
}
