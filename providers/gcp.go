package providers

import (
	"fmt"
	"time"
)

// init function is called automatically when the package is imported.
// It registers the GCPCloudProvider with the main provider factory.
func init() {
	RegisterProvider("gcp", func() CloudProvider {
		return &GCPCloudProvider{}
	})
}

// GCPCloudProvider is a dummy implementation of CloudProvider for GCP Cloud Storage.
// In a real application, this would use the GCP Cloud Storage SDK (e.g., cloud.google.com/go/storage).
type GCPCloudProvider struct{}

// ListFiles simulates listing files in a GCP Cloud Storage bucket.
func (p *GCPCloudProvider) ListFiles(bucket string) ([]string, error) {
	// Simulate fetching files from a dummy data source.
	if bucket == "log-saas-gcp" {
		return []string{"gcp_audit_log_alpha.json", "gcp_error_log_beta.json", "gcp_app.log"}, nil
	}
	return nil, fmt.Errorf("GCP bucket '%s' not found or accessible", bucket)
}

// DownloadFile simulates downloading a file from a GCP Cloud Storage bucket.
func (p *GCPCloudProvider) DownloadFile(bucket, filename string) ([]byte, error) {
	// Simulate file content.
	if bucket == "log-saas-gcp" {
		switch filename {
		case "gcp_audit_log_alpha.json":
			return []byte(`{"timestamp": "2023-11-10T14:15:00Z", "event": "user_login", "user": "alice"}`), nil
		case "gcp_error_log_beta.json":
			return []byte(`{"timestamp": "2023-11-10T14:15:05Z", "error": "resource_not_found", "path": "/data/config"}`), nil
		case "gcp_app.log":
			return []byte("Sample GCP application log: [INFO] Request received for /api/users."), nil
		default:
			return nil, fmt.Errorf("GCP file '%s' not found in bucket '%s'", filename, bucket)
		}
	}
	return nil, fmt.Errorf("GCP bucket '%s' not found or accessible", bucket)
}

// GenerateSignedURL simulates generating a pre-signed URL for a GCP Cloud Storage object.
func (p *GCPCloudProvider) GenerateSignedURL(bucket, filename string, expiry time.Duration) (string, error) {
	if bucket == "log-saas-gcp" {
		// In a real scenario, this would involve creating a signed URL using GCP Storage client.
		return fmt.Sprintf("https://mock-gcp-storage.com/%s/%s?Expires=%d&Signature=mock_gcp_signature", bucket, filename, time.Now().Add(expiry).Unix()), nil
	}
	return "", fmt.Errorf("GCP bucket '%s' not found or accessible", bucket)
}
