package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the request header.
	header := r.Header.Get("Authorization")

	// Check if the header contains the auth token.
	if !strings.Contains(header, "Bearer") {
		fmt.Println("Invalid Auth Header")
		return
	}

	// Get the secret string from the header.
	secretString := strings.Split(header, "Bearer ")[1]

	// Verify the secret string.
	if secretString != os.Getenv("AUTH_HEADER") {
		fmt.Println("Invalid Auth Header")
		return
	}

	auditLogs := AuditLogs{}

	err := json.NewDecoder(r.Body).Decode(&auditLogs)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func CreateResponse(body []byte) struct {
	Status    string `json:"status"`
	AuditLogs []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		Author  struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		Timestamp string `json:"timestamp"`
	} `json:"auditLogs"`
} {
	var response struct {
		Status    string `json:"status"`
		AuditLogs []struct {
			ID      string `json:"id"`
			Message string `json:"message"`
			Author  struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}
			Timestamp string `json:"timestamp"`
		} `json:"auditLogs"`
	}

	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return response
	}

	return response
}
