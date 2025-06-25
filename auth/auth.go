package auth

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"Cloud-Log-Access-Service/utils"

	"github.com/Nerzal/gocloak/v13" // Import gocloak for Keycloak integration
)

// Define a custom type for context keys to avoid collisions.
type contextKey string

const (
	providerContextKey contextKey = "cloudProvider" // Key for storing the cloud provider in context
)

// KeycloakConfig holds Keycloak client configuration.
type KeycloakConfig struct {
	URL      string
	Realm    string
	ClientID string
}

var (
	keycloakClient gocloak.GoCloak // Keycloak client instance
	kcConfig       KeycloakConfig  // Keycloak configuration
)

// InitKeycloak initializes the Keycloak client. This function MUST be called once at application startup.
func InitKeycloak(config KeycloakConfig) {
	kcConfig = config
	// The gocloak.NewClient expects the base Keycloak URL without the realm path.
	// The realm is passed separately in method calls (e.g., RetrospectToken, DecodeAccessToken).
	keycloakClient = gocloak.NewClient(kcConfig.URL)
	log.Printf("Keycloak client initialized for URL: %s, Realm: %s, ClientID: %s", kcConfig.URL, kcConfig.Realm, kcConfig.ClientID)
}

// KeycloakAuthMiddleware validates the JWT token with Keycloak's introspection endpoint.
// It expects an Authorization header in the format: "Bearer <JWT_TOKEN>".
// It extracts realm roles from the token and maps them to cloud providers (aws, gcp, azure).
// The detected cloud provider is then injected into the request context for subsequent handlers.
// If no valid token, no sufficient roles, or Keycloak is unreachable, it returns appropriate HTTP errors.
func KeycloakAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the Keycloak client has been initialized.
		if keycloakClient == nil {
			log.Println("Internal Server Error: Keycloak client not initialized. Call InitKeycloak at startup.")
			utils.SendError(w, http.StatusInternalServerError, "Authentication service not configured")
			return
		}

		// Extract the Authorization header.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Unauthorized: Missing Authorization header")
			utils.SendError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Validate header format.
		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println("Unauthorized: Invalid Authorization header format, expected 'Bearer <token>'")
			utils.SendError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		// Extract the JWT access token.
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Perform token introspection with Keycloak to verify its active status.
		// A context with a timeout is used to prevent blocking indefinitely.
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// The client secret is usually required for confidential clients; for public clients, it's often empty.
		// Assuming a public client for this example.
		rptResult, err := keycloakClient.RetrospectToken(ctx, accessToken, kcConfig.ClientID, "", kcConfig.Realm)
		if err != nil {
			log.Printf("Error introspecting token with Keycloak: %v", err)
			utils.SendError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Check if the token is active.
		if rptResult == nil || !*rptResult.Active {
			log.Println("Unauthorized: Token is inactive or invalid")
			utils.SendError(w, http.StatusUnauthorized, "Token is inactive or invalid")
			return
		}

		// Decode the access token to get its claims, including roles.
		tokenClaims, err := keycloakClient.DecodeAccessToken(ctx, accessToken, kcConfig.Realm)
		if err != nil {
			log.Printf("Error decoding access token: %v", err)
			utils.SendError(w, http.StatusInternalServerError, "Failed to process token claims")
			return
		}

		// Determine the cloud provider type based on the user's realm roles.
		// This logic maps Keycloak roles to specific cloud provider access.
		var providerType string
		if realmAccess, ok := tokenClaims.RealmAccess["roles"].([]interface{}); ok {
			for _, role := range realmAccess {
				roleStr, isString := role.(string)
				if !isString {
					continue // Skip if role is not a string
				}
				// Map specific Keycloak roles to cloud provider types.
				switch roleStr {
				case "aws-access": // User has role "aws-access"
					providerType = "aws"
					break
				case "gcp-access": // User has role "gcp-access"
					providerType = "gcp"
					break
				case "azure-access": // User has role "azure-access"
					providerType = "azure"
					break
				}
				if providerType != "" {
					break // Once a cloud access role is found, no need to check further
				}
			}
		}

		// If no recognized cloud access role is found, deny access.
		if providerType == "" {
			log.Println("Forbidden: No authorized cloud access roles found in token")
			utils.SendError(w, http.StatusForbidden, "Forbidden: No authorized cloud access")
			return
		}

		log.Printf("Authenticated user with %s access for provider: %s", providerType, providerType)

		// Add the determined provider type to the request context for subsequent handlers.
		ctx = context.WithValue(r.Context(), providerContextKey, providerType)
		// Call the next handler in the chain with the new context.
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetProviderFromContext retrieves the cloud provider string from the request context.
// It returns the provider string and a boolean indicating if it was found.
func GetProviderFromContext(r *http.Request) (string, bool) {
	provider, ok := r.Context().Value(providerContextKey).(string)
	return provider, ok
}
