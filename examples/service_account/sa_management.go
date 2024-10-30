package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20241023001/admin"
)

// Example for Service Account Management API
// Example uses Service Account to create Service Account.
// Please ensure that Service Account has organizational admin permission.

// Required env variables to run example:
// export MONGODB_ATLAS_CLIENT_ID="your_client_id"
// export MONGODB_ATLAS_CLIENT_SECRET="your_client_secret"
// export MONGODB_ATLAS_ORG="your_org_id"
func main() {
	host := os.Getenv("MONGODB_ATLAS_URL")
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
		admin.UseDebug(true),
		admin.UseBaseURL(host),
		admin.UseOAuthAuth(clientID, clientSecret, nil),
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	request := sdk.ServiceAccountsApi.CreateServiceAccount(ctx, org, admin.NewOrgServiceAccountRequest("SA created by sdk-example",
		"example", []string{"ORG_OWNER"}, 365*24))
	sa, _, err := request.Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_, _, err = sdk.ServiceAccountsApi.CreateServiceAccountSecret(ctx,  org, *sa.ClientId, &admin.ServiceAccountSecretRequest{
		SecretExpiresAfterHours: 365*24,
	}).Execute();
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_,  err = sdk.ServiceAccountsApi.DeleteServiceAccountSecret(ctx, *sa.ClientId, sa.GetSecrets()[0].Id, org).Execute();
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sdk.ServiceAccountsApi.DeleteServiceAccount(ctx, *sa.ClientId, org)
}
