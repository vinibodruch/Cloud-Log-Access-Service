# variables.tf
variable "aws_region" {
  description = "A região da AWS onde o bucket S3 será criado."
  type        = string
  default     = "us-east-1"
}

variable "aws_bucket_name_prefix" {
  description = "Prefixo para o nome do bucket S3 na AWS. Deve ser globalmente único."
  type        = string
  default     = "my-unique-app-aws-bucket" # Altere para algo único
}

variable "gcp_project_id" {
  description = "O ID do projeto GCP onde o bucket GCS será criado."
  type        = string
  # Substitua pelo seu Project ID real
  default     = "your-gcp-project-id" 
}

variable "gcp_region" {
  description = "A região GCP (para fins de provedor, não para bucket que é global/multi-regional)."
  type        = string
  default     = "us-central1"
}

variable "gcp_bucket_name_prefix" {
  description = "Prefixo para o nome do bucket GCS no GCP. Deve ser globalmente único."
  type        = string
  default     = "my-unique-app-gcp-bucket" # Altere para algo único
}

variable "azure_resource_group_name" {
  description = "Nome do Resource Group Azure onde o storage account será criado."
  type        = string
  default     = "my-terraform-rg"
}

variable "azure_location" {
  description = "Localização Azure (região) para o Resource Group e Storage Account."
  type        = string
  default     = "East US"
}

variable "azure_storage_account_name_prefix" {
  description = "Prefixo para o nome da Storage Account Azure. Deve ser globalmente único e min 3, max 24 caracteres, apenas minúsculas e números."
  type        = string
  default     = "mystorageapp" # Altere para algo único
}

variable "azure_container_name" {
  description = "Nome do container de blob dentro da Storage Account Azure."
  type        = string
  default     = "mycontainer"
}