package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes for a given API version.
// Receives the main Gin router, the API version, the struct with all available handlers,
// and a slice of route setup functions.
func SetupRoutes(
	router *gin.Engine,
	apiVersion string,
	appHandlers *AvailableHandlers,
	routeSetupFunctions []RouteSetupFunc,
) {
	// Defines the base API group with the dynamic version, e.g.: /api/cloud-log-access-services/v1
	apiGroup := router.Group("/api/cloud-log-access-services/" + apiVersion)

	// Iterates over the route setup functions and executes them, injecting the handlers.
	for _, setupFunc := range routeSetupFunctions {
		setupFunc(apiGroup, appHandlers)
	}
}
