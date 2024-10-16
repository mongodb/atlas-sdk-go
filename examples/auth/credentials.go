package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"go.mongodb.org/atlas-sdk/v20240805004/admin"
	"go.mongodb.org/atlas-sdk/v20240805004/auth/credentials"
)

// Example for Service Account Authentication
// Required env variables
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
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
	url := os.Getenv("MONGODB_ATLAS_URL")

	// Initialize the OAuth client in memory
	client := credentials.NewServiceAccountOAuthClient(clientID, clientSecret, &url)
	// Create an HTTP client with the custom transport (injecting the token)
	httpClient := credentials.NewHTTPClientWithServiceAccountAuth(client)

	fileTokenSource := FileTokenSource{
	}
	// Initialize the OAuth client using example FileTokenSource
	client = credentials.NewServiceAccountOAuthClientWithTokenSource(clientID, clientSecret, &fileTokenSource, &url)
	httpClient = credentials.NewHTTPClientWithServiceAccountAuth(client)

	ctx := context.Background()


	sdk, err := admin.NewClient(
		admin.UseHTTPClient(httpClient),
		admin.UseBaseURL(url),
		admin.UseDebug(true))

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

	if projects.Results == nil {
		log.Fatal("projects should not be empty:  %+v", projects)
	}
}

// FileTokenSource is an implementation of TokenSource that stores the token in a file.
type FileTokenSource struct {
	fileContent []byte
	mu       sync.Mutex
}

func (s *FileTokenSource) RetrieveToken() (*string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Replace with logic to read file
	var tkn string
	err := json.Unmarshal(s.fileContent, &tkn)
	if err != nil {
		return nil, err
	}

	return &tkn, nil
}

func (s *FileTokenSource) SaveToken(tkn string) error {
	s.mu.Lock()
	defer s.mu.Unlock()


	fileData, err := json.Marshal(tkn)
	if err != nil {
		return err
	}
	// Replace with logic to write file
	s.fileContent = fileData;
	return nil;
}
