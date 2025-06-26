# versions.tf
terraform {
  required_version = ">= 1.0" # Versão mínima do Terraform

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0" # Use a versão do provedor AWS
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0" # Use a versão do provedor GCP
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0" # Use a versão do provedor Azure
    }
  }
}

# Configuração dos provedores (AWS, GCP, Azure)
# As credenciais e regiões/locais geralmente vêm de variáveis de ambiente
# ou do provedor de credenciais padrão (AWS CLI, gcloud CLI, Azure CLI)

provider "aws" {
  region = var.aws_region
}

# provider "google" {
#   project = var.gcp_project_id
#   region  = var.gcp_region # Para recursos regionais, como alguns serviços, não para buckets globais
# }

# provider "azurerm" {
#   features {} # Bloco vazio necessário para o provedor Azurerm
#   # As credenciais serão buscadas automaticamente se você estiver logado via Azure CLI
#   # subscription_id = var.azure_subscription_id # Opcional, pode ser inferido
#   # client_id       = var.azure_client_id       # Opcional
#   # client_secret   = var.azure_client_secret   # Opcional
#   # tenant_id       = var.azure_tenant_id       # Opcional
# }