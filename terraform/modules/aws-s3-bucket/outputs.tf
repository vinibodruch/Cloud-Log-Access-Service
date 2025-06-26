output "bucket_name" {
  description = "Full name of the created S3 bucket."
  value       = aws_s3_bucket.this.id
}

output "bucket_arn" {
  description = "ARN of the created S3 bucket."
  value       = aws_s3_bucket.this.arn
}

output "application_log_s3_path" {
  value       = "s3://${aws_s3_bucket.this.id}/${aws_s3_object.application_log_object.key}"
  description = "S3 path to the sample application log file."
}