variable "bucket_name" {
  description = "Unique name for the S3 bucket."
  type        = string
}

variable "aws_region" {
  description = "AWS region for the bucket (inherited from root module)."
  type        = string
}