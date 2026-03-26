package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/atlas-sdk/v20250312016/admin"
	"go.mongodb.org/atlas-sdk/v20250312016/examples"
)

/*
 * MongoDB Atlas Go SDK Example - Update an Existing Cluster
 *
 * This example demonstrates the recommended approach for updating an Atlas
 * cluster via the PATCH API. It covers the following scenarios:
 *
 *   A. Change the cluster tier (instance size) for all electable nodes.
 *   B. Pause / unpause a cluster.
 *
 * Required env variables:
 *   `MONGODB_ATLAS_CLIENT_ID`, `MONGODB_ATLAS_CLIENT_SECRET`
 *   `MONGODB_ATLAS_PROJECT_ID`, `MONGODB_ATLAS_CLUSTER_NAME`
 * Optional:
 *   `MONGODB_ATLAS_BASE_URL` - override the default Atlas API base URL
 *
 * Considerations when using the PATCH endpoint:
 *
 *   -  The Atlas PATCH endpoint operates on `ROOT-LEVEL` fields only.
 *      To change a nested value such as `ReplicationSpecs`, you must send the
 *      ENTIRE `ReplicationSpecs` array, not just the element you changed.
 *      For any root-level field you do NOT want to change, simply omit it.
 *
 *   -  The following root-level fields returned by `GetCluster` are READ-ONLY and must
 *      NOT be sent in the PATCH payload: `ConnectionStrings`, `CreateDate`,
 *      `Id`, `MongoDBVersion`, `StateName`, and `GroupId`. The API returns HTTP 400
 *      if any of these are included.
 *
 *   -  Within `ReplicationSpecs`, the following sub-fields are read-only:
 *      `ZoneId`, `StateName`, and `CreateDate`. The API does not reject them but
 *      ignores them. Passing them has no effect. It is best practice to copy only
 *      the fields you need.
 *
 *   -  `UpdateCluster` is ASYNCHRONOUS. The API returns immediately while the
 *      cluster transitions through states (e.g., UPDATING, PAUSING). Do not
 *      issue a second `UpdateCluster` call until the cluster returns to IDLE.
 *      Poll `GetCluster` and check `StateName` before proceeding.
 *
 * For mock-based tests demonstrating these scenarios, see examples/update_cluster/update_cluster_test.go.
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
		admin.UseOAuthAuth(ctx, clientID, clientSecret),
		admin.UseBaseURL(url),
		admin.UseDebug(false))
	examples.HandleErr(err, nil)

	// Run each scenario independently. Do not call both in sequence against a
	// real cluster without polling for IDLE state in between (see async note above).
	changeClusterTierExample(ctx, sdk, projectID, clusterName)
	// pauseClusterExample(ctx, sdk, projectID, clusterName)
	// unpauseClusterExample(ctx, sdk, projectID, clusterName)
}

// changeClusterTierExample resizes all electable nodes to a new instance size.
//
// Scenario A: Change the Cluster Tier (Instance Size)
func changeClusterTierExample(ctx context.Context, sdk *admin.APIClient, projectID, clusterName string) {
	// Step 1 - GET the current cluster configuration.
	cluster, resp, err := sdk.ClustersApi.GetCluster(ctx, projectID, clusterName).Execute()
	examples.HandleErr(err, resp)

	// Step 2 - Define the new cluster tier.
	newInstanceSize := "M20"

	// Step 3 - Update instance size across all shards and regions in ReplicationSpecs.
	if cluster.ReplicationSpecs != nil {
		for i := range *cluster.ReplicationSpecs {
			spec := &(*cluster.ReplicationSpecs)[i]
			if spec.RegionConfigs == nil {
				continue
			}
			for j := range *spec.RegionConfigs {
				regionConfig := &(*spec.RegionConfigs)[j]
				if regionConfig.ElectableSpecs != nil {
					regionConfig.ElectableSpecs.InstanceSize = &newInstanceSize
				}
			}
		}
	}

	// Step 4 - Build a minimal PATCH payload.
	// Include ONLY the root-level fields that must change.
	updatePayload := &admin.ClusterDescription20240805{
		ReplicationSpecs: cluster.ReplicationSpecs,
	}

	updatedCluster, resp, err := sdk.ClustersApi.UpdateCluster(ctx, projectID, clusterName, updatePayload).Execute()
	examples.HandleErr(err, resp)

	fmt.Printf("Cluster %q tier update initiated. Current state: %s\n",
		updatedCluster.GetName(), updatedCluster.GetStateName())
}

// pauseClusterExample pauses a cluster. The cluster must be in IDLE state before calling.
//
// Scenario B: Pause a Cluster
func pauseClusterExample(ctx context.Context, sdk *admin.APIClient, projectID, clusterName string) {
	pausePayload := &admin.ClusterDescription20240805{
		Paused: admin.PtrBool(true),
	}

	pausedCluster, resp, err := sdk.ClustersApi.UpdateCluster(ctx, projectID, clusterName, pausePayload).Execute()
	examples.HandleErr(err, resp)

	fmt.Printf("Cluster %q pause initiated. Paused: %v\n",
		pausedCluster.GetName(), pausedCluster.GetPaused())
}

// unpauseClusterExample unpauses a cluster. The cluster must be in IDLE state before calling.
//
// Scenario B: Unpause a Cluster
func unpauseClusterExample(ctx context.Context, sdk *admin.APIClient, projectID, clusterName string) {
	unpausePayload := &admin.ClusterDescription20240805{
		Paused: admin.PtrBool(false),
	}

	unpausedCluster, resp, err := sdk.ClustersApi.UpdateCluster(ctx, projectID, clusterName, unpausePayload).Execute()
	examples.HandleErr(err, resp)

	fmt.Printf("Cluster %q unpause initiated. Paused: %v\n",
		unpausedCluster.GetName(), unpausedCluster.GetPaused())
}
