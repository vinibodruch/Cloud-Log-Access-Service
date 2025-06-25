package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondJSON sends a JSON response with the given status code and payload.
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}

// SendError sends a JSON error response with the given status code and message.
func SendError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	RespondJSON(w, status, errorResponse)
}
