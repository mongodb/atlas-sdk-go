package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
)

// Example for Service Account Management API
// Example uses Service Account to create Service Account.
// Please ensure that Service Account has organizational admin permission.

// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
// export MONGODB_ATLAS_ORG="your_org_id"
func main() {
	host := os.Getenv("MONGODB_ATLAS_BASE_URL")
	if host == "" {
		host = "https://cloud.mongodb.com"
	}

	// Fetch clientID and clientSecret from environment variables
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	org := os.Getenv("MONGODB_ATLAS_ORG")

	if clientID == "" || clientSecret == "" || org == "" {
		log.Fatal("Missing required environment variables")
	}

	ctx := context.Background()
	sdk, err := admin.NewClient(
		admin.UseBaseURL(host),
		admin.UseOAuthAuth(clientID, clientSecret, nil),
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 1. Create Service Account
	request := sdk.ServiceAccountsApi.CreateServiceAccount(ctx, org, admin.NewOrgServiceAccountRequest("SA created by sdk-example",
		"example", []string{"ORG_OWNER"}, 365*24))
	sa, _, err := request.Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 2. Rotate secret
	newSecret, _, err := sdk.ServiceAccountsApi.CreateServiceAccountSecret(ctx, org, *sa.ClientId, &admin.ServiceAccountSecretRequest{
		SecretExpiresAfterHours: 365 * 24,
	}).Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 3. Delete rotated secret
	_, err = sdk.ServiceAccountsApi.DeleteServiceAccountSecret(ctx, *sa.ClientId, sa.GetSecrets()[0].Id, org).Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 4. Create new SDK client using New Service Account
	newSDK, err := admin.NewClient(
		admin.UseBaseURL(host),
		// 4.1 Using ClientId and Secret returned by API
		// API might have up to 2 secrets
		admin.UseOAuthAuth(*sa.ClientId, *newSecret.Secret, nil),
	)

	// 5. Make request using new Service Account
	projects, _, err := newSDK.ProjectsApi.ListProjectsWithParams(ctx,
		&admin.ListProjectsApiParams{}).Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Projects size: %v", *projects.TotalCount)

	// 6. Remove created Service Account. We would not be able to use it afterwards without access to Secret value.
	sdk.ServiceAccountsApi.DeleteServiceAccount(ctx, *sa.ClientId, org)
}
