resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
  numeric = true
}

# --- Module AWS S3 Bucket ---
module "aws_s3_bucket" {
  source = "./modules/aws-s3-bucket" # Caminho para o módulo local

  bucket_name = "${var.aws_bucket_name_prefix}-${random_string.suffix.result}"
  aws_region  = var.aws_region
}

# module "gcp_gcs_bucket" {
#   source = "./modules/gcp-gcs-bucket" # Caminho para o módulo local

#   project_id  = var.gcp_project_id
#   bucket_name = "${var.gcp_bucket_name_prefix}-${random_string.suffix.result}"
# }

# module "azure_blob_storage" {
#   source = "./modules/az-blob-storage" # Caminho para o módulo local

#   resource_group_name         = var.azure_resource_group_name
#   location                    = var.azure_location
#   storage_account_name_prefix = var.azure_storage_account_name_prefix
#   container_name              = var.azure_container_name
#   random_suffix               = random_string.suffix.result # Passa o sufixo aleatório
# }