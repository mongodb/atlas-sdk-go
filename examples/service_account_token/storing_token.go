package main

import (
	"context"
	"fmt"
	"go.mongodb.org/atlas-sdk/v20241023002/auth/credentials"
	"golang.org/x/oauth2"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
)

// Example for reusing Service Account OAuth Token across application restarts
// Covers creation of the HTTP client with Service Account OAuth authentication
// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
func main() {
	host := os.Getenv("MONGODB_ATLAS_URL")
	if host == "" {
		host = "https://cloud.mongodb.com"
	}
	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}
	ctx := context.Background()
	// 1. Simulate retrieving Token from filesystem
	token, err := getCachedToken(clientID, clientSecret, host, ctx)

	// Create Admin API Client enabling token and cache.
	var tokenCache = func(token string) error {
		fmt.Println("Cache was called. Use it to save token in secure location")
		return nil
	}

	sdk, err := admin.NewClient(
		admin.UseBaseURL(host),
		admin.UseOAuthAuthWithCache(clientID, clientSecret, token, &tokenCache),
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 3. Make API call
	projects, _, err := sdk.ProjectsApi.ListProjectsWithParams(ctx, &admin.ListProjectsApiParams{}).Execute()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if projects.Results == nil {
		fmt.Printf("projects should not be empty:  %v", projects)
	}
}

func getCachedToken(clientID string, clientSecret string, host string, ctx context.Context) (*oauth2.Token, error) {
	// In this example we are fetching token from backend as we do not want to save token
	config := credentials.NewConfig(clientID, clientSecret)
	config.TokenURL = host + "/api/oauth/token"
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return token, err
}
