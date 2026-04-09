package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20250312018/admin"
	"go.mongodb.org/atlas-sdk/v20250312018/auth"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

/*
* MongoDB Atlas Go SDK Retryable Request Example
*
* Example using custom http client that handles rate limiting and 500 Http errors by retrying requests automatically.
* Example uses https://pkg.go.dev/github.com/hashicorp/go-retryablehttp.
* Please refer to the package documentation for more information.
 */
func main() {
	ctx := context.Background()
	// Service Accounts are the recommended way to authenticate programmatically with the MongoDB Atlas API
	// See: See: https://www.mongodb.com/docs/cloud-manager/tutorial/manage-programmatic-api-keys/
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("MONGODB_ATLAS_CLIENT_ID and MONGODB_ATLAS_CLIENT_SECRET must be set")
	}

	// Using custom client
	// This example relies on https://pkg.go.dev/github.com/hashicorp/go-retryablehttp
	// retryablehttp performs automatic retries under certain conditions.
	// Mainly, if an error is returned by the client (connection errors etc),
	// or if a 500-range response is received, then a retry is invoked.
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3

	// Inject the retryable client into the context so that UseOAuthAuth wraps it
	// rather than http.DefaultClient. This ensures retries apply to all API requests.
	ctx = context.WithValue(ctx, auth.HTTPClient, retryClient.StandardClient())

	sdk, err := admin.NewClient(
		admin.UseBaseURL(url),
		admin.UseOAuthAuth(ctx, clientID, clientSecret),
		admin.UseDebug(false))
	if err != nil {
		log.Fatal(err)
	}

	request := sdk.ProjectsApi.ListGroupsWithParams(ctx,
		&admin.ListGroupsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})
	projects, _, err := request.Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Projects", projects.GetTotalCount())

}
