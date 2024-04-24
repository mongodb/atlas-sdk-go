package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20231115011/admin"
	"go.mongodb.org/atlas-sdk/v20231115011/mockadmin"
)

func TestListClusters(t *testing.T) {
	// Create mock API.
	clusterAPI := mockadmin.NewClustersApi(t)

	// Program expectations.
	list := &admin.PaginatedAdvancedClusterDescription{
		TotalCount: admin.PtrInt(2),
		Results: &[]admin.AdvancedClusterDescription{
			{StateName: admin.PtrString("IDLE")},
			{StateName: admin.PtrString("DELETING")},
		},
	}
	clusterAPI.EXPECT().ListClusters(mock.Anything, mock.Anything).Return(admin.ListClustersApiRequest{ApiService: clusterAPI})
	clusterAPI.EXPECT().ListClustersExecute(mock.Anything).Return(list, &http.Response{StatusCode: 200}, nil)

	// Call functions using the API, they won't make real calls to Atlas API.
	clusterCount, err := MyFunctionCallingListClusters(clusterAPI)

	// Assert expectations.
	assert.Nil(t, err)
	assert.Equal(t, 2, clusterCount)
}

func MyFunctionCallingListClusters(clusterAPI admin.ClustersApi) (int, error) {
	clusters, _, err := clusterAPI.ListClusters(context.Background(), "my_group_id").Execute()
	return clusters.GetTotalCount(), err
}
