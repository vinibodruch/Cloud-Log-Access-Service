# modules/azure-blob-storage/variables.tf
variable "resource_group_name" {
  description = "Nome do Resource Group Azure."
  type        = string
}

variable "location" {
  description = "Localização Azure (região)."
  type        = string
}

variable "storage_account_name_prefix" {
  description = "Prefixo para o nome da Storage Account Azure. Deve ser globalmente único, min 3, max 24, apenas minúsculas e números."
  type        = string
}

variable "container_name" {
  description = "Nome do container de blob dentro da Storage Account."
  type        = string
}

variable "random_suffix" {
  description = "Sufixo aleatório para garantir unicidade do nome da Storage Account."
  type        = string
}