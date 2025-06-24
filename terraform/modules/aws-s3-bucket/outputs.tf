# modules/aws-s3-bucket/outputs.tf
output "bucket_name" {
  description = "Nome completo do bucket S3 criado."
  value       = aws_s3_bucket.this.id
}

output "bucket_arn" {
  description = "ARN do bucket S3 criado."
  value       = aws_s3_bucket.this.arn
}