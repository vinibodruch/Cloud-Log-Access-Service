package providers

import (
	"fmt"
	"time"
)

// init function is called automatically when the package is imported.
// It registers the AWSCloudProvider with the main provider factory.
func init() {
	RegisterProvider("aws", func() CloudProvider {
		return &AWSCloudProvider{}
	})
}

// AWSCloudProvider is a dummy implementation of CloudProvider for AWS S3.
// In a real application, this would use the AWS SDK (e.g., github.com/aws/aws-sdk-go-v2/service/s3).
type AWSCloudProvider struct{}

// ListFiles simulates listing files in an AWS S3 bucket.
func (p *AWSCloudProvider) ListFiles(bucket string) ([]string, error) {
	// Simulate fetching files from a dummy data source.
	if bucket == "log-saas-aws" {
		return []string{"aws_log_2023-01-01.txt", "aws_log_2023-01-02.txt", "aws_access.log"}, nil
	}
	return nil, fmt.Errorf("AWS bucket '%s' not found or accessible", bucket)
}

// DownloadFile simulates downloading a file from an AWS S3 bucket.
func (p *AWSCloudProvider) DownloadFile(bucket, filename string) ([]byte, error) {
	// Simulate file content.
	if bucket == "log-saas-aws" {
		switch filename {
		case "aws_log_2023-01-01.txt":
			return []byte("Content of AWS log 2023-01-01."), nil
		case "aws_log_2023-01-02.txt":
			return []byte("Content of AWS log 2023-01-02."), nil
		case "aws_access.log":
			return []byte("Sample AWS access log entry: [10/Nov/2023:14:12:01 +0000] \"GET /index.html\" 200 1234"), nil
		default:
			return nil, fmt.Errorf("AWS file '%s' not found in bucket '%s'", filename, bucket)
		}
	}
	return nil, fmt.Errorf("AWS bucket '%s' not found or accessible", bucket)
}

// GenerateSignedURL simulates generating a pre-signed URL for an AWS S3 object.
func (p *AWSCloudProvider) GenerateSignedURL(bucket, filename string, expiry time.Duration) (string, error) {
	if bucket == "log-saas-aws" {
		// In a real scenario, this would involve creating a pre-signed URL using AWS SDK's PresignClient.
		// The URL would be temporary and grant specific permissions.
		return fmt.Sprintf("https://mock-aws-s3.com/%s/%s?Expires=%d&Signature=mock_aws_signature", bucket, filename, time.Now().Add(expiry).Unix()), nil
	}
	return "", fmt.Errorf("AWS bucket '%s' not found or accessible", bucket)
}
