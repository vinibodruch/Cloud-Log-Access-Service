package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// AppConfig contém todas as configurações globais da aplicação.
type AppConfig struct {
	AWS    AWSConfig
	Azure  AzureConfig
	GCP    GCPConfig // Placeholder para futura integração GCP
	APIVersion string
	Port       string
}

// LoadAppConfig carrega todas as configurações da aplicação a partir de variáveis de ambiente ou .env.
func LoadAppConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Aviso: Não foi possível carregar o arquivo .env. As variáveis de ambiente do sistema serão usadas.")
	}

	awsCfg := LoadAWSConfig()
	azureCfg := LoadAzureConfig()
	gcpCfg := GCPConfig{} // Placeholder: Implemente LoadGCPConfig se necessário

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1"
		log.Printf("Aviso: Variável de ambiente API_VERSION não definida. Usando a versão padrão: %s", apiVersion)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Aviso: Variável de ambiente PORT não definida. Usando a porta padrão: %s", port)
	}

	return AppConfig{
		AWS:    awsCfg,
		Azure:  azureCfg,
		GCP:    gcpCfg,
		APIVersion: apiVersion,
		Port:       port,
	}
}
