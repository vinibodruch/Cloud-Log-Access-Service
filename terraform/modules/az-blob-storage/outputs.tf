# modules/azure-blob-storage/outputs.tf
output "resource_group_name" {
  description = "Nome do Resource Group Azure criado."
  value       = azurerm_resource_group.this.name
}

output "storage_account_name" {
  description = "Nome da Storage Account Azure criada."
  value       = azurerm_storage_account.this.name
}

output "blob_container_url" {
  description = "URL do container de blob Azure."
  value       = azurerm_storage_container.this.resource_manager_id
}