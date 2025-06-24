# outputs.tf
output "aws_s3_bucket_name" {
  description = "Nome completo do bucket S3 da AWS."
  value       = module.aws_s3_bucket.bucket_name
}

output "aws_s3_bucket_arn" {
  description = "ARN do bucket S3 da AWS."
  value       = module.aws_s3_bucket.bucket_arn
}

output "gcp_gcs_bucket_name" {
  description = "Nome completo do bucket GCS do GCP."
  value       = module.gcp_gcs_bucket.bucket_name
}

output "gcp_gcs_bucket_url" {
  description = "URL do bucket GCS do GCP."
  value       = module.gcp_gcs_bucket.bucket_url
}

output "azure_storage_account_name" {
  description = "Nome da Storage Account Azure criada."
  value       = module.azure_blob_storage.storage_account_name
}

output "azure_blob_container_url" {
  description = "URL do container de blob Azure."
  value       = module.azure_blob_storage.blob_container_url
}