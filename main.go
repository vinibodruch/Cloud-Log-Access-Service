package main

import (
	"flag"
	"log"
	"net/http"

	"Cloud-Log-Access-Service/auth"
	"Cloud-Log-Access-Service/handlers"
	_ "Cloud-Log-Access-Service/providers/aws"   // Import providers to register them via their init() functions
	_ "Cloud-Log-Access-Service/providers/azure" //
	_ "Cloud-Log-Access-Service/providers/gcp"   // This is crucial for the Abstract Factory to know about them
)

func main() {
	// Initialize Keycloak client with your Keycloak URL, realm, and client ID.
	// Ensure these match your Keycloak setup.
	keycloakURL := flag.String("keycloak-url", "http://localhost:8082", "Base URL of Keycloak (e.g., http://localhost:8082)")
	keycloakRealm := flag.String("keycloak-realm", "master", "Your Keycloak realm (e.g., 'master' or 'SAP')")
	keycloakClientID := flag.String("keycloak-client-id", "cloud-log-access-service-client", "Your Keycloak client ID")
	serverPort := flag.String("port", ":8080", "Port for the HTTP server to listen on")

	flag.Parse()

	auth.InitKeycloak(auth.KeycloakConfig{
		URL:      *keycloakURL,
		Realm:    *keycloakRealm,
		ClientID: *keycloakClientID,
	})

	// Create a new HTTP multiplexer
	mux := http.NewServeMux()

	// Register the handler functions with their respective routes.
	// The auth.KeycloakAuthMiddleware now performs actual token validation
	// and injects the detected cloud provider into the request context based on Keycloak roles.
	mux.HandleFunc("/api/logs/list", auth.KeycloakAuthMiddleware(handlers.ListLogsHandler))
	mux.HandleFunc("/api/logs/download", auth.KeycloakAuthMiddleware(handlers.DownloadLogHandler))
	mux.HandleFunc("/api/logs/signed-url", auth.KeycloakAuthMiddleware(handlers.GenerateSignedURLHandler))

	// Define the address the server will listen on
	log.Printf("Server starting on port %s", *serverPort)

	// Start the HTTP server. log.Fatal will log any error and then exit the program.
	if err := http.ListenAndServe(*serverPort, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
