# modules/gcp-gcs-bucket/variables.tf
variable "project_id" {
  description = "O ID do projeto GCP."
  type        = string
}

variable "bucket_name" {
  description = "Nome único para o bucket GCS."
  type        = string
}