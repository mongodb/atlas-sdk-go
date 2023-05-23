package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/admin"
)

/*
* MongoDB Atlas Go SDK Example
*
* Example performs following actions:
* - Finds first project in the user organization
* - Creates new MongoDB cluster backed by AWS provider in that project
* - Creates database user for connection
* - Creates IP access to enable connection
 */
func main() {
	ctx := context.Background()
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/authentication/api-key/
	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL("https://cloud-dev.mongodb.com"),
		admin.UseDebug(false))
	handleErr(err, nil)

	// -- 1. Get first project
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	handleErr(err, response)

	if projects.GetTotalCount() == 0 {
		log.Fatal("account should have at least single project")
	}

	projectId := projects.GetResults()[0].GetId()
	providers := []string{"AWS","GCP","AZURE"}
	regions,response, err :=sdk.ClustersApi.ListCloudProviderRegions(ctx,projectId).Providers(providers).Execute()
	handleErr(err,response)
	fmt.Println(regions);
}

func handleErr(err error, resp *http.Response) {
	if err == nil {
		return
	}

	if resp != nil {
		fmt.Println(resp.Body)
		// Printing generic message
		fmt.Println(err.Error())
	}
	apiErr, _ := admin.AsError(err)
	log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())

}

func getIpAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
