package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312017/admin"
	"go.mongodb.org/atlas-sdk/v20250312017/examples"
)

/*
 * MongoDB Atlas Go SDK Example - Pause and Unpause a Cluster
 *
 * This example demonstrates how to pause and unpause an Atlas cluster
 * via the PATCH API.
 *
 * Required env variables:
 *   `MONGODB_ATLAS_CLIENT_ID`, `MONGODB_ATLAS_CLIENT_SECRET`
 *   `MONGODB_ATLAS_PROJECT_ID`, `MONGODB_ATLAS_CLUSTER_NAME`
 * Optional:
 *   `MONGODB_ATLAS_BASE_URL` - override the default Atlas API base URL
 *
 * Note: `UseBaseURL` must be called before `UseOAuthAuth` when constructing the
 * client so that the OAuth token endpoint is set to the correct base URL.
 * See the `admin.NewClient` call in main() for the correct ordering.
 *
 * Considerations:
 *
 *   -  Pausing only requires setting the `Paused` root-level field to true
 *      in a minimal PATCH payload. Omit all other fields.
 *
 *   -  `UpdateCluster` is ASYNCHRONOUS. The API returns immediately while the
 *      cluster transitions through states (e.g., PAUSING, IDLE). Do not
 *      issue a second `UpdateCluster` call until the cluster returns to IDLE.
 *      Poll `GetCluster` and check `StateName` before proceeding.
 */
func main() {
	ctx := context.Background()

	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")

	projectID := os.Getenv("MONGODB_ATLAS_PROJECT_ID")
	clusterName := os.Getenv("MONGODB_ATLAS_CLUSTER_NAME")

	if clientID == "" || clientSecret == "" {
		log.Fatal("MONGODB_ATLAS_CLIENT_ID and MONGODB_ATLAS_CLIENT_SECRET must be set")
	}
	if projectID == "" || clusterName == "" {
		log.Fatal("MONGODB_ATLAS_PROJECT_ID and MONGODB_ATLAS_CLUSTER_NAME must be set")
	}

	sdk, err := admin.NewClient(
		admin.UseBaseURL(url),
		admin.UseOAuthAuth(ctx, clientID, clientSecret),
		admin.UseDebug(false))
	examples.HandleErr(err, nil)

	// Pause then unpause the cluster, with a short wait in between to
	// illustrate that the cluster is doing work between the two operations.
	// Note: in practice you should poll GetCluster until StateName == IDLE
	// before issuing the next UpdateCluster call (see async note above).
	pauseCluster(ctx, sdk, projectID, clusterName)

	// To keep the cluster paused longer, comment out the unpauseCluster call
	// and the lines below it.
	fmt.Println("Waiting 5 seconds before unpausing...")
	time.Sleep(5 * time.Second)

	unpauseCluster(ctx, sdk, projectID, clusterName)
}

// pauseCluster pauses a cluster. The cluster must be in IDLE state before calling.
func pauseCluster(ctx context.Context, sdk *admin.APIClient, projectID, clusterName string) {
	payload := &admin.ClusterDescription20240805{
		Paused: admin.PtrBool(true),
	}

	result, resp, err := sdk.ClustersApi.UpdateCluster(ctx, projectID, clusterName, payload).Execute()
	examples.HandleErr(err, resp)

	fmt.Printf("Cluster %q pause initiated. Paused: %v\n",
		result.GetName(), result.GetPaused())
}

// unpauseCluster unpauses a cluster. The cluster must be in IDLE state before calling.
func unpauseCluster(ctx context.Context, sdk *admin.APIClient, projectID, clusterName string) {
	payload := &admin.ClusterDescription20240805{
		Paused: admin.PtrBool(false),
	}

	result, resp, err := sdk.ClustersApi.UpdateCluster(ctx, projectID, clusterName, payload).Execute()
	examples.HandleErr(err, resp)

	fmt.Printf("Cluster %q unpause initiated. Paused: %v\n",
		result.GetName(), result.GetPaused())
}
