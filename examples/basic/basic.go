package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20240805001/admin"
	"go.mongodb.org/atlas-sdk/v20240805001/examples"
)

/*
* MongoDB Atlas Go SDK Basic Example
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
		admin.UseBaseURL(url))
	examples.HandleErr(err, nil)

	// -- 1. Get first project
	request := sdk.ProjectsApi.ListProjectsWithParams(ctx,
		// 2. We passing struct with all parameters to the request
		&admin.ListProjectsApiParams{
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
