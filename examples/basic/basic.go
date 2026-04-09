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
* MongoDB Atlas Go SDK Basic Example
 */
func main() {
	ctx := context.Background()
	// Service Accounts are the recommended way to authenticate programmatically with the MongoDB Atlas API
	// See: https://www.mongodb.com/docs/cloud-manager/tutorial/manage-programmatic-api-keys/
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("MONGODB_ATLAS_CLIENT_ID and MONGODB_ATLAS_CLIENT_SECRET must be set")
	}

	sdk, err := admin.NewClient(
		admin.UseBaseURL(url),
		admin.UseOAuthAuth(ctx, clientID, clientSecret))
	examples.HandleErr(err, nil)

	// -- 1. Get first project
	request := sdk.ProjectsApi.ListGroupsWithParams(ctx,
		// 2. We are passing a struct with all parameters to the request
		&admin.ListGroupsApiParams{
			ItemsPerPage: admin.PtrInt(1),
			IncludeCount: admin.PtrBool(true),
			PageNum:      admin.PtrInt(1),
		})

	// 3. We can also use builder pattern to construct request
	projects, response, err := request.IncludeCount(true).PageNum(1).Execute()
	examples.HandleErr(err, response)

	if projects.GetTotalCount() == 0 {
		log.Fatal("account should have at least single project")
	}

	projectId := projects.GetResults()[0].GetId()
	fmt.Printf("Project we use %v", projectId)
}
