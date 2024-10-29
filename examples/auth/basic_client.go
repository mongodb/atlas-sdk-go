package main

import (
	"context"
	"fmt"
	"go.mongodb.org/atlas-sdk/v20241023001/admin"
	"log"
	"os"
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

	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	projects, _, err := request.IncludeCount(true).PageNum(1).Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if projects.Results == nil {
		fmt.Printf("projects should not be empty:  %v", projects)
	}
}
