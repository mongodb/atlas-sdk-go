package main

import (
	"io"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20250312019/admin"
	"go.mongodb.org/atlas-sdk/v20250312019/examples"
)

/*
* MongoDB Atlas Go SDK Example for fetching cluster logs file content
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

	// -- 2. Get first Process
	hosts, response, err := sdk.MonitoringAndLogsApi.ListGroupProcesses(ctx, projectId).Execute()
	examples.HandleErr(err, response)
	if len(hosts.GetResults()) == 0 {
		log.Fatal("your cluster should have at least single host. Are you running Atlas M0?")
	}
	host := hosts.GetResults()[0].GetHostname()
	params := &admin.DownloadClusterLogApiParams{
		GroupId:  projectId,
		HostName: host,
		LogName:  "mongos",
	}

	logs, response, err := sdk.MonitoringAndLogsApi.DownloadClusterLogWithParams(ctx, params).Execute()
	examples.HandleErr(err, response)
	defer func() {
		_ = logs.Close()
	}()
	_, err = io.Copy(os.Stdout, logs)
	examples.HandleErr(err, nil)
}
