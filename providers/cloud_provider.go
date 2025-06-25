package providers

import (
	"fmt"
	"time"
)

// CloudProvider is an interface that defines the operations for interacting with cloud storage.
// Each specific cloud provider (AWS, GCP, Azure) will implement this interface.
type CloudProvider interface {
	ListFiles(bucket string) ([]string, error)                                       // Lists files in a given bucket.
	DownloadFile(bucket, filename string) ([]byte, error)                            // Downloads a specific file from a bucket.
	GenerateSignedURL(bucket, filename string, expiry time.Duration) (string, error) // Generates a temporary access URL for a file.
}

// ProviderFactory is a function type that creates a new CloudProvider instance.
type ProviderFactory func() CloudProvider

// registeredProviders is a map to store factories for different cloud providers.
// This allows us to dynamically create provider instances based on a string identifier.
var registeredProviders = make(map[string]ProviderFactory)

// RegisterProvider allows external packages to register their CloudProvider implementations.
// This is part of the Abstract Factory pattern, enabling extensibility.
func RegisterProvider(name string, factory ProviderFactory) {
	registeredProviders[name] = factory
}

// GetCloudProvider creates and returns a CloudProvider instance based on the provided name.
// It returns an error if the provider name is not recognized (i.e., not registered).
func GetCloudProvider(providerType string) (CloudProvider, error) {
	factory, ok := registeredProviders[providerType]
	if !ok {
		return nil, fmt.Errorf("unsupported cloud provider: %s", providerType)
	}
	return factory(), nil
}
