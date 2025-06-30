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

// Application representa a aplicação Go e seus componentes principais.
type Application struct {
	Config    config.AppConfig
	Handlers  *routes.AvailableHandlers
	Router    *gin.Engine
}

// NewApplication inicializa todos os componentes da aplicação (configuração, serviços, handlers, rotas).
func NewApplication() *Application {
	// 1. Carregar a configuração global da aplicação
	appCfg := config.LoadAppConfig()

	// 2. Inicializar Serviços (Dependências de baixo nível, ex: clientes de SDK)
	s3Service := awsServices.NewS3Service(appCfg.AWS.Config)
	azureBlobService := azureServices.NewBlobService(appCfg.Azure.Client)

	// 3. Inicializar Handlers (Dependem dos Serviços e encapsulam a lógica HTTP)
	s3Handler := awsHandlers.NewS3Handler(s3Service)
	azureBlobHandler := azureHandlers.NewBlobHandler(azureBlobService)

	// 4. Agrupar todos os Handlers Disponíveis em uma única struct para injeção nas rotas
	allHandlers := &routes.AvailableHandlers{
		S3:        s3Handler,
		AzureBlob: azureBlobHandler,
		// Adicione outros handlers aqui (GCPStorage: gcpStorageHandler)
	}

	// 5. Inicializar o roteador Gin
	router := gin.Default()

	// 6. Definir as funções de configuração de rota que serão carregadas para a versão da API atual.
	var apiRoutes []routes.RouteSetupFunc

	switch appCfg.APIVersion {
	case "v1":
		apiRoutes = []routes.RouteSetupFunc{
			routes.S3Routes,    // Rotas para AWS S3
			routes.AzureRoutes, // Rotas para Azure Blob Storage
			// routes.GCPRoutes,   // Adicione rotas para GCP aqui se tiver
		}
	case "v2":
		// Exemplo: rotas específicas para a v2 da API
		// apiRoutes = []routes.RouteSetupFunc{
		// 	routes.S3RoutesV2,
		// 	routes.AzureRoutesV2,
		// }
	default:
		log.Fatalf("Versão da API '%s' não suportada. Verifique a variável de ambiente API_VERSION.", appCfg.APIVersion)
	}

	// 7. Configurar todas as rotas da aplicação
	routes.SetupRoutes(router, appCfg.APIVersion, allHandlers, apiRoutes)

	return &Application{
		Config:    appCfg,
		Handlers:  allHandlers,
		Router:    router,
	}
}

// Run inicia o servidor HTTP da aplicação.
func (a *Application) Run() error {
	log.Printf("Servidor iniciado na porta %s (Versão da API: %s)", a.Config.Port, a.Config.APIVersion)
	return a.Router.Run(":" + a.Config.Port)
}
