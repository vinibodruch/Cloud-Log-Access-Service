package routes

import (
	"github.com/gin-gonic/gin"
)

// RouteSetupFunc define a assinatura para funções que configuram grupos de rotas.
// Ela recebe o grupo de roteador do Gin e uma struct com todos os handlers disponíveis.
type RouteSetupFunc func(routerGroup *gin.RouterGroup, h *AvailableHandlers)
