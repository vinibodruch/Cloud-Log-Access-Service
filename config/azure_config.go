package config

import (
	//"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// AzureConfig contém a configuração e o cliente para Azure Blob Storage.
type AzureConfig struct {
	AccountName string
	Client      *azblob.Client
}

// LoadAzureConfig carrega a configuração e inicializa o cliente Azure Blob.
func LoadAzureConfig() AzureConfig {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	if accountName == "" {
		log.Printf("AZURE_STORAGE_ACCOUNT_NAME não definida. Pulando configuração Azure Blob.")
		return AzureConfig{}
	}

	// DefaultAzureCredential tenta usar várias formas de credenciais (variáveis de ambiente, Azure CLI, Managed Identity, etc.)
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Falha ao obter credenciais do Azure: %v", err)
	}

	serviceURL := "https://" + accountName + ".blob.core.windows.net/"
	client, err := azblob.NewClient(serviceURL, credential, nil)
	if err != nil {
		log.Fatalf("Falha ao criar cliente Azure Blob: %v", err)
	}

	return AzureConfig{
		AccountName: accountName,
		Client:      client,
	}
}

// GCPConfig é um placeholder para a configuração do GCP.
type GCPConfig struct {
	// Defina aqui as configurações para GCP
}

// Implemente func LoadGCPConfig() GCPConfig para carregar configurações do GCP.
