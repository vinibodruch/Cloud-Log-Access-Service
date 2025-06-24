# modules/gcp-gcs-bucket/main.tf
resource "google_storage_bucket" "this" {
  name          = var.bucket_name
  project       = var.project_id
  location      = "US" # Pode ser "US", "EU", "ASIA" ou uma multi-região específica
  storage_class = "STANDARD" # Opções: STANDARD, NEARLINE, COLDLINE, ARCHIVE
  uniform_bucket_level_access = true # Boa prática para consistência de permissões

  labels = {
    environment = "dev"
    managed_by  = "terraform"
  }

  # Desabilita o acesso público (recomendado)
  public_access_prevention = "enforced"
}