package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/atlas-sdk/v20241023001/admin"
	"go.mongodb.org/atlas-sdk/v20241023001/auth/credentials"
)

// Advanced Example for Service Account OAuth Authentication
// Provides:
// 1. Custom Token Cache in order to store OAuth Token for further reuse across application restarts
// 2. Usage of logout (Token Revocation) method.
// 3. Using custom Context and HttpTransports
//
// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
// export MONGODB_ATLAS_BASE_URL=https://cloud.mongodb.com
func main() {
	host := os.Getenv("MONGODB_ATLAS_BASE_URL")
	if host == "" {
		log.Fatal("Missing MONGODB_ATLAS_BASE_URL")
	}

	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET environment variables")
	}

	// Initialize the OAuth client with memory
	tokenSource := credentials.NewTokenSource(clientID, clientSecret)
	// Create an HTTP client with the custom transport (injecting the token)
	httpClient := credentials.NewHTTPClientWithOAuthToken(tokenSource)

	fileTokenCache := MyTokenCache{}
	// Initialize tokenSource
	tokenSource = credentials.NewTokenSourceWithOptions(
		credentials.AtlasTokenSourceOptions{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenCache:   &fileTokenCache,
			BaseURL:      &host,
		})
	httpClient = credentials.NewHTTPClientWithOAuthToken(tokenSource)

	ctx := context.Background()

	sdk, err := admin.NewClient(
		admin.UseHTTPClient(httpClient),
		admin.UseBaseURL(host),
		admin.UseDebug(true))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	projects, _, err := sdk.ProjectsApi.ListProjectsWithParams(ctx, &admin.ListProjectsApiParams{}).Execute()
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	if projects.Results == nil {
		fmt.Printf("projects should not be empty:  %v", projects)
	}

	// Revoke Token in the client (logout)
	err = tokenSource.RevokeToken()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// MyTokenCache is an implementation of LocalTokenCache that stores the token in a file.
type MyTokenCache struct {
	fileContent []byte
	mu          sync.Mutex
}

func (s *MyTokenCache) RetrieveToken(_ context.Context) (string, error) {
	// Locking added to ensure no race condition happens with SaveToken
	s.mu.Lock()
	defer s.mu.Unlock()

	// Example: token is saved as string.
	// Please replace it with your logic to read the token from filesystem
	var tkn string
	err := json.Unmarshal(s.fileContent, &tkn)
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func (s *MyTokenCache) SaveToken(_ context.Context, tkn string) error {
	// Locking added to ensure no race condition happens with RetrieveToken
	s.mu.Lock()
	defer s.mu.Unlock()

	fileData, err := json.Marshal(tkn)
	if err != nil {
		return err
	}
	// Example: Replace with logic to write file to filesystem
	s.fileContent = fileData
	return nil
}
