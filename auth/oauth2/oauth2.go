package main

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

type CustomTokenSource struct {
	oauth2.TokenSource
	tokenFile string
	mu        sync.Mutex
	token     *oauth2.Token
	expiry    time.Time
}

func (ts *CustomTokenSource) Token() (*oauth2.Token, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if ts.token != nil && !ts.token.Valid() {
		return ts.token, nil
	}

	// If token is expired or not cached, fetch a new one
	token, err := ts.TokenSource.Token()
	if err != nil {
		return nil, err
	}

	// Save the new token for future use
	ts.token = token
	err = ts.saveToken(token)
	if err != nil {
		return nil, err
	}

	return ts.token, nil
}

func (ts *CustomTokenSource) loadToken() (*oauth2.Token, error) {
	return nil, nil
}

func (ts *CustomTokenSource) saveToken(token *oauth2.Token) error {
	return nil
}

func main() {
	// Define OAuth2 configuration
	cfg := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://cloud.mongodb.com/api/oauth/token",
		},
	}

	// Fetch initial token using client_credentials flow
	ctx := context.Background()
	token, err := cfg.PasswordCredentialsToken(ctx, "username", "password")
	if err != nil {
		log.Fatalf("Error fetching token: %v", err)
	}

	// Create an HTTP client that automatically injects and refreshes the token
	client := cfg.Client(ctx, token)

	// Use the client to make authenticated requests
	resp, err := client.Get("https://localhost:8080/api/atlas/v2/info")
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
}
