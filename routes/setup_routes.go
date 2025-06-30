package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas da aplicação para uma dada versão da API.
// Recebe o roteador principal do Gin, a versão da API, a struct com todos os handlers disponíveis
// e um slice de funções de configuração de rota.
func SetupRoutes(
	router *gin.Engine,
	apiVersion string,
	appHandlers *AvailableHandlers,
	routeSetupFunctions []RouteSetupFunc,
) {
	// Define o grupo base da API com a versão dinâmica, ex: /api/cloud-log-access-services/v1
	apiGroup := router.Group("/api/cloud-log-access-services/" + apiVersion)

	// Itera sobre as funções de configuração de rota e as executa, injetando os handlers.
	for _, setupFunc := range routeSetupFunctions {
		setupFunc(apiGroup, appHandlers)
	}
}
