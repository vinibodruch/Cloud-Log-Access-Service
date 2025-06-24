# modules/aws-s3-bucket/variables.tf
variable "bucket_name" {
  description = "Nome único para o bucket S3."
  type        = string
}

variable "aws_region" {
  description = "Região da AWS para o bucket (herdado do módulo raiz)."
  type        = string
}