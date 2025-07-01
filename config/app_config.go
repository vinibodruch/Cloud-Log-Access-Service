package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// AppConfig contains all global application configurations.
type AppConfig struct {
	AWS        AWSConfig
	Azure      AzureConfig
	GCP        GCPConfig // Placeholder for future GCP integration
	APIVersion string
	Port       string
}

// LoadAppConfig loads all application configurations from environment variables or .env file.
func LoadAppConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file. System environment variables will be used.")
	}

	awsCfg := LoadAWSConfig()
	azureCfg := LoadAzureConfig()
	gcpCfg := GCPConfig{} // Placeholder: Implement LoadGCPConfig if needed

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1"
		log.Printf("Warning: API_VERSION environment variable not set. Using default version: %s", apiVersion)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Warning: PORT environment variable not set. Using default port: %s", port)
	}

	return AppConfig{
		AWS:        awsCfg,
		Azure:      azureCfg,
		GCP:        gcpCfg,
		APIVersion: apiVersion,
		Port:       port,
	}
}
