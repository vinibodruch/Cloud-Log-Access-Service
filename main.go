package main

import (
	"log"

	"Cloud-Log-Access-Service/aws/config"
	"Cloud-Log-Access-Service/aws/handlers"
	"Cloud-Log-Access-Service/aws/routes"
	"Cloud-Log-Access-Service/aws/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadAWSConfig()

	s3Service := services.NewS3Service(cfg)

	s3Handler := handlers.NewS3Handler(s3Service)

	router := gin.Default()

	routes.SetupRoutes(router, s3Handler)

	port := "8080"
	log.Printf("Server started on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
