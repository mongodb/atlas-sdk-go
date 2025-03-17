package examples

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
	"go.mongodb.org/atlas-sdk/v20250312001/mockadmin"
)

func myFunctionCallingListClusters(clusterAPI admin.ClustersApi) (int, error) {
	clusters, _, err := clusterAPI.ListClusters(context.Background(), "my_group_id").Execute()
	if err != nil {
		return 0, err
	}
	return clusters.GetTotalCount(), err
}

func TestListClusters(t *testing.T) {
	// Create mock API.
	clusterAPI := mockadmin.NewClustersApi(t)

	// Program expectations.
	list := &admin.PaginatedClusterDescription20240805{
		TotalCount: admin.PtrInt(2),
		Results: &[]admin.ClusterDescription20240805{
			{StateName: admin.PtrString("IDLE")},
			{StateName: admin.PtrString("DELETING")},
		},
	}
	clusterAPI.EXPECT().ListClusters(mock.Anything, mock.Anything).Return(admin.ListClustersApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().ListClustersExecute(mock.Anything).Return(list, &http.Response{StatusCode: 200}, nil)

	// Call functions using the API, they won't make real calls to Atlas API.
	clusterCount, err := myFunctionCallingListClusters(clusterAPI)

	// Assert expectations.
	assert.Nil(t, err)
	assert.Equal(t, 2, clusterCount)

	clusterAPI.EXPECT().ListClustersExecute(mock.Anything).Return(list, &http.Response{StatusCode: 200}, nil)
}

func TestListClustersErrorMocks(t *testing.T) {
	// Create mock API.
	clusterAPI := mockadmin.NewClustersApi(t)
	apiError := admin.GenericOpenAPIError{}

	apiError.SetModel(admin.ApiError{
		Detail:    admin.PtrString("Error when listing clusters"),
		Error:     400,
		ErrorCode: "CLUSTERS_UNREACHABLE",
		Reason:    admin.PtrString("Clusters unreachable"),
	})
	apiError.SetError("Mocked error")
	clusterAPI.EXPECT().ListClusters(mock.Anything, mock.Anything).Return(admin.ListClustersApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().ListClustersExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 400}, apiError)

	// Call functions using the API, they won't make real calls to Atlas API.
	_, err := myFunctionCallingListClusters(clusterAPI)

	// Assert expectations.
	assert.Equal(t, err, apiError)
}
