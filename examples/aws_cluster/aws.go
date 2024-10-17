package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/preview/admin"
	"go.mongodb.org/atlas-sdk/preview/examples"
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
	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	url := os.Getenv("MONGODB_ATLAS_URL")

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL(url),
		admin.UseDebug(false))
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

	// -- 2. Create Cluster
	cluster := createClusterRequest(projectId)
	createdCluster, resp, err := sdk.ClustersApi.CreateCluster(ctx, projectId, cluster).Execute()
	examples.HandleErr(err, resp)

	// -- 3. Create Database User
	user := createDatabaseUserRequest(sdk, projectId)
	if user != nil {
		_, resp, err = sdk.DatabaseUsersApi.
			CreateDatabaseUser(ctx, projectId, user).
			Execute()
		examples.HandleErr(err, resp)
	}

	ipAddress := getIpAddress()
	ipAddressEntry := []admin.NetworkPermissionEntry{{IpAddress: &ipAddress}}
	// -- 4. Enable IP access
	_, resp, err = sdk.ProjectIPAccessListApi.
		CreateProjectIpAccessList(context.Background(), projectId, &ipAddressEntry).
		Execute()
	examples.HandleErr(err, resp)

	// -- Print connection details to user
	fmt.Printf("SDK: Created new MongoDB cluster: %v \n", createdCluster.GetId())
	fmt.Println("Please wait up to 10 minutes for cluster to provision.")
}

func createDatabaseUserRequest(sdk *admin.APIClient, groupId string) *admin.CloudDatabaseUser {
	username := "sdk-example"
	databaseName := "admin"

	currentUser, _, _ := sdk.DatabaseUsersApi.GetDatabaseUser(context.Background(),
		groupId, databaseName, username).Execute()

	if currentUser != nil {
		return nil
	}

	collectionName := "sdk-test"
	password, err := randomASCIIString(10)
	if err != nil {
		log.Fatal(err)
	}

	return &admin.CloudDatabaseUser{
		Username:     username,
		Password:     &password,
		DatabaseName: databaseName,
		Roles: &[]admin.DatabaseUserRole{
			{
				DatabaseName:   databaseName,
				CollectionName: &collectionName,
				RoleName:       "read",
			},
		},
	}
}

func createClusterRequest(projectId string) *admin.ClusterDescription20240805 {
	// Input arguments used for creation of the cluster
	clusterName, _ := uniqueName("example-aws-cluster")

	// Location
	providerName := "AWS"
	clusterType := "SHARDED"
	regionName := "US_EAST_1"

	// Size
	priority := int(7)
	nodeCount := int(3)
	instanceSize := "M10"

	return &admin.ClusterDescription20240805{
		Name:        &clusterName,
		ClusterType: &clusterType,
		ReplicationSpecs: &[]admin.ReplicationSpec20240805{
			{
				RegionConfigs: &[]admin.CloudRegionConfig20240805{
					{
						ProviderName: &providerName,
						Priority:     &priority,
						RegionName:   &regionName,
						ElectableSpecs: &admin.HardwareSpec20240805{
							InstanceSize: &instanceSize,
							NodeCount:    &nodeCount,
						},
					},
				},
			},
		},
	}

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

func randomASCIIString(length int) (string, error) {
	const asciiMax = 127
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(asciiMax)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 64 && n < asciiMax {
			result += string(rune(n))
		}
	}
}

func uniqueName(prefix string) (string, error) {
	const asciiMax = 122
	const length = 5
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(asciiMax)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 97 && n < asciiMax {
			prefix += string(rune(n))
		}
	}
	return prefix, nil
}
