package main

import (
	"log"
	"net/http"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20231115008/admin"
	"go.mongodb.org/atlas-sdk/v20231115008/examples"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/mongodb-forks/digest"
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
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/authentication/api-key/
	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	url := os.Getenv("MONGODB_ATLAS_URL")

	// Using custom client 
	// This example relies on https://pkg.go.dev/github.com/hashicorp/go-retryablehttp
	// retryablehttp performs automatic retries under certain conditions. 
	// Mainly, if an error is returned by the client (connection errors etc), 
	/// or if a 500-range response is received, then a retry is invoked.
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3

	retryableClient := createRetryableClient(retryClient, apiKey, apiSecret)
	sdk, err := admin.NewClient(
		admin.UseHTTPClient(retryableClient),
		admin.UseBaseURL(url),
		admin.UseDebug(false))
	examples.HandleErr(err, nil)

	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})
	projects, response, err := request.Execute()
	examples.HandleErr(err, response)

	if projects.GetTotalCount() == 0 {
		log.Fatal("account should have at least single project")
	}

}

func createRetryableClient(retryClient *retryablehttp.Client, apiKey string, apiSecret string) *http.Client {
	var transport http.RoundTripper = &retryablehttp.RoundTripper{Client: retryClient}
	digestRetryAbleTransport := digest.NewTransportWithHTTPRoundTripper(apiKey, apiSecret,transport)
	retryableClient, err := digestRetryAbleTransport.Client()
	if err != nil {
		log.Fatal("Cannot instantiate client")
	}
	return retryableClient
}
