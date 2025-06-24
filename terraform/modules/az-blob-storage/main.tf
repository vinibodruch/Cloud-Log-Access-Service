# modules/azure-blob-storage/main.tf
resource "azurerm_resource_group" "this" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_storage_account" "this" {
  name                     = "${var.storage_account_name_prefix}${var.random_suffix}" # Sufixo para unicidade
  resource_group_name      = azurerm_resource_group.this.name
  location                 = azurerm_resource_group.this.location
  account_tier             = "Standard"
  account_replication_type = "LRS" # Locally-redundant storage

  # Boa prática: Desabilitar acesso público a blobs/containers por padrão
  public_network_access_enabled = true # Necessário para acessar do Terraform, mas controle de acesso ao blob
  # Para desabilitar completamente o acesso público, você usaria network_rules
  # Ou definir account_kind para "BlockBlobStorage" e then block_public_access_enabled = true (para contas v2)

  tags = {
    environment = "dev"
    managed_by  = "Terraform"
  }
}

resource "azurerm_storage_container" "this" {
  name                  = var.container_name
  storage_account_name  = azurerm_storage_account.this.name
  container_access_type = "private" # Ninguém pode ler ou escrever blobs sem autorização
}