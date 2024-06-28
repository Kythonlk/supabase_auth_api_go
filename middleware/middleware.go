package middleware

import (
	"github.com/supabase-community/gotrue-go"
	"log"
	"net/http"
)

func AccessTokenMiddleware(next http.HandlerFunc, client gotrue.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// token := strings.TrimPrefix(authHeader, "Bearer ")

		// Get the user
		_, err := client.GetUser()
		if err != nil {
			log.Printf("Invalid access token: %v", err)
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
