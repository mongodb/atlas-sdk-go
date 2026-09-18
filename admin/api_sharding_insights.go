// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ShardingInsightsAPI interface {

	/*
		GetClusterShardingInsights Return Sharding Insights for One Cluster

		Returns the cluster-level sharding state of one sharded cluster: the number of shards, the balancer state, the availability of the config servers, and the configured balancing window.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterShardingInsightsApiRequest
	*/
	GetClusterShardingInsights(ctx context.Context, groupId string, clusterName string) GetClusterShardingInsightsApiRequest
	/*
		GetClusterShardingInsights Return Sharding Insights for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterShardingInsightsApiParams - Parameters for the request
		@return GetClusterShardingInsightsApiRequest
	*/
	GetClusterShardingInsightsWithParams(ctx context.Context, args *GetClusterShardingInsightsApiParams) GetClusterShardingInsightsApiRequest

	// Method available only for mocking purposes
	GetClusterShardingInsightsExecute(r GetClusterShardingInsightsApiRequest) (*ShardingInsightsResponse, *http.Response, error)
}

// ShardingInsightsAPIService ShardingInsightsAPI service
type ShardingInsightsAPIService service

type GetClusterShardingInsightsApiRequest struct {
	ctx         context.Context
	ApiService  ShardingInsightsAPI
	groupId     string
	clusterName string
}

type GetClusterShardingInsightsApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ShardingInsightsAPIService) GetClusterShardingInsightsWithParams(ctx context.Context, args *GetClusterShardingInsightsApiParams) GetClusterShardingInsightsApiRequest {
	return GetClusterShardingInsightsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterShardingInsightsApiRequest) Execute() (*ShardingInsightsResponse, *http.Response, error) {
	return r.ApiService.GetClusterShardingInsightsExecute(r)
}

/*
GetClusterShardingInsights Return Sharding Insights for One Cluster

Returns the cluster-level sharding state of one sharded cluster: the number of shards, the balancer state, the availability of the config servers, and the configured balancing window.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterShardingInsightsApiRequest
*/
func (a *ShardingInsightsAPIService) GetClusterShardingInsights(ctx context.Context, groupId string, clusterName string) GetClusterShardingInsightsApiRequest {
	return GetClusterShardingInsightsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetClusterShardingInsightsExecute executes the request
//
//	@return ShardingInsightsResponse
func (a *ShardingInsightsAPIService) GetClusterShardingInsightsExecute(r GetClusterShardingInsightsApiRequest) (*ShardingInsightsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ShardingInsightsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardingInsightsAPIService.GetClusterShardingInsights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardingInsights"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	if r.groupId == "." || r.groupId == ".." {
		return localVarReturnValue, nil, reportError("groupId must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	if r.clusterName == "." || r.clusterName == ".." {
		return localVarReturnValue, nil, reportError("clusterName must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
