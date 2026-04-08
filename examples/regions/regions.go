package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20250312018/admin"
	"go.mongodb.org/atlas-sdk/v20250312018/examples"
)

/*
* MongoDB Atlas Go SDK Example for fetching cloud provider regions
 */
func main() {
	ctx := context.Background()
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/service-accounts/
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("MONGODB_ATLAS_CLIENT_ID and MONGODB_ATLAS_CLIENT_SECRET must be set")
	}

	sdk, err := admin.NewClient(
		admin.UseBaseURL(url),
		admin.UseOAuthAuth(ctx, clientID, clientSecret),
		admin.UseDebug(false))
	examples.HandleErr(err, nil)

	// -- 1. Get first project
	projects, response, err := sdk.ProjectsApi.ListGroupsWithParams(ctx,
		&admin.ListGroupsApiParams{
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
	regions, response, err := sdk.ClustersApi.ListClusterProviderRegions(ctx, projectId).Providers(providers).Execute()
	examples.HandleErr(err, response)
	fmt.Println(regions)
}
