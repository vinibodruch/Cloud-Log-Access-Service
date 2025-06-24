# modules/gcp-gcs-bucket/outputs.tf
output "bucket_name" {
  description = "Nome completo do bucket GCS criado."
  value       = google_storage_bucket.this.name
}

output "bucket_url" {
  description = "URL do bucket GCS (gs://)."
  value       = google_storage_bucket.this.self_link
}