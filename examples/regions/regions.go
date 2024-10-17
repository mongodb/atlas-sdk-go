package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"github.com/mongodb/atlas-sdk-go/admin"
	"github.com/mongodb/atlas-sdk-go/examples"
)

/*
* MongoDB Atlas Go SDK Example for fetching cloud provider regions
 */
func main() {
	ctx := context.Background()
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/authentication/api-key/
	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	url := os.Getenv("MONGODB_ATLAS_URL")

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL(url),
		admin.UseDebug(false))
	examples.HandleErr(err, nil)

	// -- 1. Get first project
	projects, response, err := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		&admin.ListProjectsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		}).Execute()
	examples.HandleErr(err, response)

	if projects.GetTotalCount() == 0 {
		log.Fatal("account should have at least single project")
	}

	projectId := projects.GetResults()[0].GetId()
	providers := []string{"AWS", "GCP", "AZURE"}
	regions, response, err := sdk.ClustersApi.ListCloudProviderRegions(ctx, projectId).Providers(providers).Execute()
	examples.HandleErr(err, response)
	fmt.Println(regions)
}
