package main

import (
	"log"

	"github.com/supabase-community/gotrue-go"
)

const (
	projectReference = ""
	apiKey           = ""
)

func main() {
	// Initialise client``
	client := gotrue.New(
		projectReference,
		apiKey,
	)

	tokenResponse, err := client.SignInWithEmailPassword("email@gmail.com", "12345678")
	if err != nil {
		log.Printf("Failed to sign in: %v", err)
	}

	login, err := client.RefreshToken(tokenResponse.RefreshToken)
	if err != nil {
		log.Printf("Failed to sign in: %v", err)
	}
	// Use the token response
	log.Printf("Access Token: %s\n", tokenResponse.RefreshToken)
	log.Printf("Refresh Token: %s\n", login.RefreshToken)
}
