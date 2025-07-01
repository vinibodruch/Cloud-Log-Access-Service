package routes

import (
	"github.com/gin-gonic/gin"
)

// RouteSetupFunc defines the signature for functions that set up route groups.
// It receives the Gin router group and a struct with all available handlers.
type RouteSetupFunc func(routerGroup *gin.RouterGroup, h *AvailableHandlers)
