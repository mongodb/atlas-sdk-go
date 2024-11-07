package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb/atlas-sdk-go/admin"
)

// Basic example for Service Account OAuth Authentication
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
	// Create Admin API Client with OAuth credentials.
	// Skips optional tokenCache parameter
	sdk, err := admin.NewClient(
		admin.UseOAuthAuth(clientID, clientSecret, nil),
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	projects, _, err := sdk.ProjectsApi.ListProjectsWithParams(ctx,&admin.ListProjectsApiParams{}).Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if projects.Results == nil {
		fmt.Printf("projects should not be empty:  %v", projects)
	}
}
