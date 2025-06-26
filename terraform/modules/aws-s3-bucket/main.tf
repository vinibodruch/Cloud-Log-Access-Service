resource "aws_s3_bucket" "this" {
  bucket = var.bucket_name
  tags = {
    Environment = "Dev"
    ManagedBy   = "Terraform"
  }
}

resource "aws_s3_bucket_public_access_block" "this" {
  bucket = aws_s3_bucket.this.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# --- Carregar sample_application_log.json ---
resource "local_file" "application_log_file" {
  content  = file("${path.module}/logs/sample_application_log.json")
  filename = "${path.module}/logs/sample_application_log.json" 
}

resource "aws_s3_object" "application_log_object" {
  bucket      = aws_s3_bucket.this.id 
  key  = "application-logs/${formatdate("YYYY/MM/DD", timestamp())}/sample_application_log.json"
  source      = local_file.application_log_file.filename

  content_type = "application/json"
}

