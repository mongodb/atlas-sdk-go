package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20240805004/admin"
	"go.mongodb.org/atlas-sdk/v20240805004/auth/credentials"
)

// Example for Service Account Authentication
// Required env variables
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
// export MONGODB_ATLAS_ACCESS_TOKEN="optional_existing_token" # Optional
// export MONGODB_ATLAS_URL=https://cloud.mongodb.com
func main() {
	host := os.Getenv("MONGODB_ATLAS_URL")
	if host == "" {
		log.Fatal("Missing MONGODB_ATLAS_URL")
	}

	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}

	// Optionally, you can pass an existing token from the environment
	accessToken := os.Getenv("MONGODB_ATLAS_ACCESS_TOKEN")

	// Initialize the OAuth client
	client, err := credentials.NewServiceAccountOAuthClient(clientID, clientSecret, accessToken, nil)
	if err != nil {
		log.Fatalf("Error initializing OAuth client: %v", err)
	}

	// Create an HTTP client with the custom transport (injecting the token)
	httpClient := credentials.NewHTTPClientWithServiceAccountAuth(client)

	ctx := context.Background()
	url := os.Getenv("MONGODB_ATLAS_URL")

	sdk, err := admin.NewClient(
		admin.UseHTTPClient(httpClient),
		admin.UseBaseURL(url),
		admin.UseDebug(false))

	if err != nil{
		log.Fatal("Error: %e", err);
	}

	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})

	if err != nil {
		log.Fatalf("Error making request: %e", err)
	}
	projects, _, err := request.IncludeCount(true).PageNum(1).Execute()
	if err != nil {
		log.Fatalf("Error making request: %e", err)
	}

	if projects.Results != nil {
		log.Fatal("projects should not be empty")
	}
}
