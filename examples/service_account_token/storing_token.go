package main

import (
	"context"
	"fmt"
	"go.mongodb.org/atlas-sdk/v20241023001/auth/credentials"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023001/admin"
)

// Example for reusing Service Account OAuth Token across application restarts
// Covers creation of the HTTP client with Service Account OAuth authentication
// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
func main() {
	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}
	ctx := context.Background()
	// 1. Simulate retrieving Token from filesystem
	// In this example we are fetching token from backend
	config := credentials.NewConfig(clientID, clientSecret)
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Create Admin API Client with OAuth credentials.
	// 2. Providing cached token as parameter
	sdk, err := admin.NewClient(
		admin.UseOAuthAuth(clientID, clientSecret, token),
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 3. Make API call
	projects, _, err := sdk.ProjectsApi.ListProjectsWithParams(ctx, &admin.ListProjectsApiParams{}).Execute()

	// 4. Cache token
	// ?

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if projects.Results == nil {
		fmt.Printf("projects should not be empty:  %v", projects)
	}
}
