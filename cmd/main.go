package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Kythonlk/supabase_auth_api_go/types"
	"github.com/supabase-community/gotrue-go"
	"io"
)

func main() {

	projectReference := os.Getenv("PR_REF")
	apiKey := os.Getenv("API_KEY")

	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	// Initialise client
	client := gotrue.New(
		projectReference,
		apiKey,
	)

	http.HandleFunc("/api/v01/refresh_token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req types.RefreshTokenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		login, err := client.RefreshToken(req.RefreshToken)
		if err != nil {
			log.Printf("Failed to refresh token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		response := types.RefreshTokenResponse{
			AccessToken:  login.AccessToken,
			RefreshToken: login.RefreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3202"
	}

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
