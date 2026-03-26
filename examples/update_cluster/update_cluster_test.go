package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20250312016/admin"
	"go.mongodb.org/atlas-sdk/v20250312016/mockadmin"
)

// changeClusterTier demonstrates the recommended way to resize a cluster:
// GET the current config, update only the instance size in ReplicationSpecs,
// then PATCH with only that root-level field.
func changeClusterTier(ctx context.Context, api admin.ClustersApi, projectID, clusterName, instanceSize string) (*admin.ClusterDescription20240805, error) {
	cluster, _, err := api.GetCluster(ctx, projectID, clusterName).Execute()
	if err != nil {
		return nil, err
	}
	if cluster.ReplicationSpecs != nil {
		for i := range *cluster.ReplicationSpecs {
			spec := &(*cluster.ReplicationSpecs)[i]
			if spec.RegionConfigs == nil {
				continue
			}
			for j := range *spec.RegionConfigs {
				regionConfig := &(*spec.RegionConfigs)[j]
				if regionConfig.ElectableSpecs != nil {
					regionConfig.ElectableSpecs.InstanceSize = &instanceSize
				}
			}
		}
	}
	payload := &admin.ClusterDescription20240805{
		ReplicationSpecs: cluster.ReplicationSpecs,
	}
	result, _, err := api.UpdateCluster(ctx, projectID, clusterName, payload).Execute()
	return result, err
}

// setClusterPaused demonstrates pausing or unpausing a cluster by sending
// only the Paused root-level field in the PATCH payload.
func setClusterPaused(ctx context.Context, api admin.ClustersApi, projectID, clusterName string, paused bool) (*admin.ClusterDescription20240805, error) {
	payload := &admin.ClusterDescription20240805{
		Paused: admin.PtrBool(paused),
	}
	result, _, err := api.UpdateCluster(ctx, projectID, clusterName, payload).Execute()
	return result, err
}

func TestChangeClusterTier(t *testing.T) {
	clusterAPI := mockadmin.NewClustersApi(t)

	providerName := "AWS"
	regionName := "US_EAST_1"
	priority := 7
	nodeCount := 3
	currentSize := "M10"

	fakeCluster := &admin.ClusterDescription20240805{
		Name:      admin.PtrString("my-cluster"),
		StateName: admin.PtrString("IDLE"),
		ReplicationSpecs: &[]admin.ReplicationSpec20240805{
			{
				RegionConfigs: &[]admin.CloudRegionConfig20240805{
					{
						ProviderName: &providerName,
						RegionName:   &regionName,
						Priority:     &priority,
						ElectableSpecs: &admin.HardwareSpec20240805{
							InstanceSize: &currentSize,
							NodeCount:    &nodeCount,
						},
					},
				},
			},
		},
	}
	updatingCluster := &admin.ClusterDescription20240805{
		Name:      admin.PtrString("my-cluster"),
		StateName: admin.PtrString("UPDATING"),
	}

	clusterAPI.EXPECT().GetCluster(mock.Anything, "my-project", "my-cluster").Return(admin.GetClusterApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().GetClusterExecute(mock.Anything).Return(fakeCluster, &http.Response{StatusCode: 200}, nil)
	clusterAPI.EXPECT().
		UpdateCluster(mock.Anything, "my-project", "my-cluster",
			mock.MatchedBy(func(payload *admin.ClusterDescription20240805) bool {
				if payload.ReplicationSpecs == nil || payload.Paused != nil || payload.Name != nil {
					return false
				}
				specs := *payload.ReplicationSpecs
				if len(specs) == 0 || specs[0].RegionConfigs == nil {
					return false
				}
				regions := *specs[0].RegionConfigs
				if len(regions) == 0 || regions[0].ElectableSpecs == nil {
					return false
				}
				return regions[0].ElectableSpecs.GetInstanceSize() == "M20"
			}),
		).
		Return(admin.UpdateClusterApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().UpdateClusterExecute(mock.Anything).Return(updatingCluster, &http.Response{StatusCode: 200}, nil)

	result, err := changeClusterTier(context.Background(), clusterAPI, "my-project", "my-cluster", "M20")

	assert.NoError(t, err)
	assert.Equal(t, "UPDATING", result.GetStateName())
}

func TestChangeClusterTierGetError(t *testing.T) {
	clusterAPI := mockadmin.NewClustersApi(t)

	apiErr := admin.GenericOpenAPIError{}
	apiErr.SetError("cluster not found")

	clusterAPI.EXPECT().GetCluster(mock.Anything, mock.Anything, mock.Anything).Return(admin.GetClusterApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().GetClusterExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, apiErr)

	result, err := changeClusterTier(context.Background(), clusterAPI, "my-project", "my-cluster", "M20")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSetClusterPaused(t *testing.T) {
	clusterAPI := mockadmin.NewClustersApi(t)

	pausedCluster := &admin.ClusterDescription20240805{
		Name:   admin.PtrString("my-cluster"),
		Paused: admin.PtrBool(true),
	}

	clusterAPI.EXPECT().
		UpdateCluster(mock.Anything, "my-project", "my-cluster",
			mock.MatchedBy(func(payload *admin.ClusterDescription20240805) bool {
				return payload.Paused != nil && *payload.Paused == true && payload.ReplicationSpecs == nil
			}),
		).
		Return(admin.UpdateClusterApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().UpdateClusterExecute(mock.Anything).Return(pausedCluster, &http.Response{StatusCode: 200}, nil)

	result, err := setClusterPaused(context.Background(), clusterAPI, "my-project", "my-cluster", true)

	assert.NoError(t, err)
	assert.Equal(t, true, result.GetPaused())
}

func TestSetClusterUnpaused(t *testing.T) {
	clusterAPI := mockadmin.NewClustersApi(t)

	unpausedCluster := &admin.ClusterDescription20240805{
		Name:   admin.PtrString("my-cluster"),
		Paused: admin.PtrBool(false),
	}

	clusterAPI.EXPECT().
		UpdateCluster(mock.Anything, "my-project", "my-cluster",
			mock.MatchedBy(func(payload *admin.ClusterDescription20240805) bool {
				return payload.Paused != nil && *payload.Paused == false && payload.ReplicationSpecs == nil
			}),
		).
		Return(admin.UpdateClusterApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().UpdateClusterExecute(mock.Anything).Return(unpausedCluster, &http.Response{StatusCode: 200}, nil)

	result, err := setClusterPaused(context.Background(), clusterAPI, "my-project", "my-cluster", false)

	assert.NoError(t, err)
	assert.Equal(t, false, result.GetPaused())
}
