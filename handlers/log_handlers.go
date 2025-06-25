package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"Cloud-Log-Access-Service/auth"
	"Cloud-Log-Access-Service/providers"
	"Cloud-Log-Access-Service/utils"
)

// getCloudProviderInstance is a helper function to retrieve the correct cloud provider
// based on the authenticated user's context.
func getCloudProviderInstance(w http.ResponseWriter, r *http.Request) (providers.CloudProvider, error) {
	// Get the cloud provider type from the request context, which was set by the auth middleware.
	providerType, ok := auth.GetProviderFromContext(r)
	if !ok || providerType == "" {
		log.Println("Internal Error: Cloud provider not found in context")
		utils.SendError(w, http.StatusInternalServerError, "Internal server error: provider information missing")
		return nil, fmt.Errorf("provider information missing in context")
	}

	// Use the abstract factory to get the concrete CloudProvider implementation.
	provider, err := providers.GetCloudProvider(providerType)
	if err != nil {
		log.Printf("Error getting cloud provider for type '%s': %v", providerType, err)
		utils.SendError(w, http.StatusBadRequest, err.Error()) // Bad request if provider type is not supported
		return nil, err
	}
	return provider, nil
}

// ListLogsHandler handles requests to list log files in a specified bucket.
// Expects a query parameter: `bucket`.
// Example: GET /api/logs/list?bucket=log-saas-aws
func ListLogsHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET.
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Extract the 'bucket' query parameter.
	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		utils.SendError(w, http.StatusBadRequest, "Missing 'bucket' query parameter")
		return
	}

	// Get the appropriate cloud provider instance.
	provider, err := getCloudProviderInstance(w, r)
	if err != nil {
		return // Error already sent by getCloudProviderInstance
	}

	// List files using the selected provider.
	// Goroutines could be introduced here if you need to perform multiple concurrent
	// listing operations or other background tasks for a single request.
	files, err := provider.ListFiles(bucket)
	if err != nil {
		log.Printf("Error listing files for bucket '%s': %v", bucket, err)
		utils.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to list files: %v", err))
		return
	}

	// Respond with the list of files in JSON format.
	utils.RespondJSON(w, http.StatusOK, map[string][]string{"files": files})
}

// DownloadLogHandler handles requests to download a specific log file.
// Expects query parameters: `bucket` and `filename`.
// Example: GET /api/logs/download?bucket=log-saas-aws&filename=aws_log_2023-01-01.txt
func DownloadLogHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET.
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Extract query parameters.
	bucket := r.URL.Query().Get("bucket")
	filename := r.URL.Query().Get("filename")

	if bucket == "" || filename == "" {
		utils.SendError(w, http.StatusBadRequest, "Missing 'bucket' or 'filename' query parameters")
		return
	}

	// Get the appropriate cloud provider instance.
	provider, err := getCloudProviderInstance(w, r)
	if err != nil {
		return // Error already sent by getCloudProviderInstance
	}

	// Download the file content.
	content, err := provider.DownloadFile(bucket, filename)
	if err != nil {
		log.Printf("Error downloading file '%s' from bucket '%s': %v", filename, bucket, err)
		utils.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to download file: %v", err))
		return
	}

	// Set appropriate headers for file download.
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream") // Generic binary file type
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))

	// Write the file content to the response body.
	_, err = w.Write(content)
	if err != nil {
		log.Printf("Error writing file content to response for '%s': %v", filename, err)
		// No need to send error to client as headers might already be sent.
	}
}

// GenerateSignedURLHandler handles requests to generate a temporary access link (pre-signed URL).
// Expects query parameters: `bucket`, `filename`, and optionally `expiry` (in seconds).
// Example: GET /api/logs/signed-url?bucket=log-saas-aws&filename=aws_access.log&expiry=3600
func GenerateSignedURLHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET.
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Extract query parameters.
	bucket := r.URL.Query().Get("bucket")
	filename := r.URL.Query().Get("filename")
	expiryStr := r.URL.Query().Get("expiry")

	if bucket == "" || filename == "" {
		utils.SendError(w, http.StatusBadRequest, "Missing 'bucket' or 'filename' query parameters")
		return
	}

	// Parse expiry duration, default to 1 hour (3600 seconds) if not provided or invalid.
	expirySeconds, err := strconv.Atoi(expiryStr)
	if err != nil || expirySeconds <= 0 {
		expirySeconds = 3600 // Default to 1 hour
		log.Printf("Invalid or missing 'expiry' parameter, defaulting to %d seconds.", expirySeconds)
	}
	expiryDuration := time.Duration(expirySeconds) * time.Second

	// Get the appropriate cloud provider instance.
	provider, err := getCloudProviderInstance(w, r)
	if err != nil {
		return // Error already sent by getCloudProviderInstance
	}

	// Generate the signed URL.
	signedURL, err := provider.GenerateSignedURL(bucket, filename, expiryDuration)
	if err != nil {
		log.Printf("Error generating signed URL for '%s/%s': %v", bucket, filename, err)
		utils.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate signed URL: %v", err))
		return
	}

	// Respond with the signed URL in JSON format.
	utils.RespondJSON(w, http.StatusOK, map[string]string{"signedURL": signedURL, "expires_in_seconds": fmt.Sprintf("%d", expirySeconds)})
}
