package app

import (
	"log"

	awsHandlers "Cloud-Log-Access-Service/aws/handlers"
	awsServices "Cloud-Log-Access-Service/aws/services"
	azureHandlers "Cloud-Log-Access-Service/azure/handlers"
	azureServices "Cloud-Log-Access-Service/azure/services"
	"Cloud-Log-Access-Service/config"
	"Cloud-Log-Access-Service/routes"

	"github.com/gin-gonic/gin"
)

// Application represents the Go application and its main components.
type Application struct {
	Config   config.AppConfig
	Handlers *routes.AvailableHandlers
	Router   *gin.Engine
}

// NewApplication initializes all application components (configuration, services, handlers, routes).
func NewApplication() *Application {
	// 1. Load the global application configuration
	appCfg := config.LoadAppConfig()

	// 2. Initialize Services (Low-level dependencies, e.g., SDK clients)
	s3Service := awsServices.NewS3Service(appCfg.AWS.Config)
	azureBlobService := azureServices.NewBlobService(appCfg.Azure.Client)

	// 3. Initialize Handlers (Depend on Services and encapsulate HTTP logic)
	s3Handler := awsHandlers.NewS3Handler(s3Service)
	azureBlobHandler := azureHandlers.NewBlobHandler(azureBlobService)

	// 4. Group all available Handlers into a single struct for injection into routes
	allHandlers := &routes.AvailableHandlers{
		S3:        s3Handler,
		AzureBlob: azureBlobHandler,
		// Add other handlers here (GCPStorage: gcpStorageHandler)
	}

	// 5. Initialize the Gin router
	router := gin.Default()

	// 6. Define the route setup functions to be loaded for the current API version.
	var apiRoutes []routes.RouteSetupFunc

	switch appCfg.APIVersion {
	case "v1":
		apiRoutes = []routes.RouteSetupFunc{
			routes.S3Routes,    // Routes for AWS S3
			routes.AzureRoutes, // Routes for Azure Blob Storage
			// routes.GCPRoutes,   // Add routes for GCP here if available
		}
	case "v2":
		// Example: specific routes for API v2
		// apiRoutes = []routes.RouteSetupFunc{
		// 	routes.S3RoutesV2,
		// 	routes.AzureRoutesV2,
		// }
	default:
		log.Fatalf("API version '%s' not supported. Check the API_VERSION environment variable.", appCfg.APIVersion)
	}

	// 7. Set up all application routes
	routes.SetupRoutes(router, appCfg.APIVersion, allHandlers, apiRoutes)

	return &Application{
		Config:   appCfg,
		Handlers: allHandlers,
		Router:   router,
	}
}

// Run starts the application's HTTP server.
func (a *Application) Run() error {
	log.Printf("Server started on port %s (API Version: %s)", a.Config.Port, a.Config.APIVersion)
	return a.Router.Run(":" + a.Config.Port)
}
