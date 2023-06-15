package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/admin"
	"go.mongodb.org/atlas-sdk/examples"
)

/*
* MongoDB Atlas Go SDK Example for fetching cluster logs file content
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
	var out io.Writer
	projectId := projects.GetResults()[0].GetId()

	// -- 2. Get first Process
	hosts, response, err := sdk.MonitoringAndLogsApi.ListAtlasProcesses(ctx, projectId).Execute()
	examples.HandleErr(err, response)
	if len(hosts.GetResults()) == 0 {
		log.Fatal("MongoDB atlas hosts are missing. Please review your cluster details in Atlas CLI or UI")
	}
	host := hosts.GetResults()[0].GetHostname()
	params := &admin.GetHostLogsApiParams{
		GroupId:  projectId,
		HostName: host,
		LogName:  "mongos",
	}

	logs, response, err := sdk.MonitoringAndLogsApi.GetHostLogsWithParams(ctx, params).Execute()
	examples.HandleErr(err, response)
	_, err = io.Copy(out, logs)
	fmt.Println(err)
}
