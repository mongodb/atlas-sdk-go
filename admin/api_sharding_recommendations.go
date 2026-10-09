// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ShardingRecommendationsAPI interface {

	/*
		GetClusterShardingRecommendations Return Sharding Recommendations for One Cluster

		Returns the active Sharding Advisor recommendations for one cluster, together with the recommendation types that were evaluated. Recommendations describe changes that may improve the cluster's data distribution, such as adding shards or sharding a collection, together with the monitored conditions that produced them.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterShardingRecommendationsApiRequest
	*/
	GetClusterShardingRecommendations(ctx context.Context, groupId string, clusterName string) GetClusterShardingRecommendationsApiRequest
	/*
		GetClusterShardingRecommendations Return Sharding Recommendations for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterShardingRecommendationsApiParams - Parameters for the request
		@return GetClusterShardingRecommendationsApiRequest
	*/
	GetClusterShardingRecommendationsWithParams(ctx context.Context, args *GetClusterShardingRecommendationsApiParams) GetClusterShardingRecommendationsApiRequest

	// Method available only for mocking purposes
	GetClusterShardingRecommendationsExecute(r GetClusterShardingRecommendationsApiRequest) (*ShardingRecommendationsResponse, *http.Response, error)
}

// ShardingRecommendationsAPIService ShardingRecommendationsAPI service
type ShardingRecommendationsAPIService service

type GetClusterShardingRecommendationsApiRequest struct {
	ctx                 context.Context
	ApiService          ShardingRecommendationsAPI
	groupId             string
	clusterName         string
	recommendationTypes *[]string
	shards              *[]string
	namespaces          *[]string
}

type GetClusterShardingRecommendationsApiParams struct {
	GroupId             string
	ClusterName         string
	RecommendationTypes *[]string
	Shards              *[]string
	Namespaces          *[]string
}

func (a *ShardingRecommendationsAPIService) GetClusterShardingRecommendationsWithParams(ctx context.Context, args *GetClusterShardingRecommendationsApiParams) GetClusterShardingRecommendationsApiRequest {
	return GetClusterShardingRecommendationsApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		clusterName:         args.ClusterName,
		recommendationTypes: args.RecommendationTypes,
		shards:              args.Shards,
		namespaces:          args.Namespaces,
	}
}

// Recommendation types to evaluate. To include multiple types, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each type. Omit this parameter to evaluate every type that applies to the cluster.
func (r GetClusterShardingRecommendationsApiRequest) RecommendationTypes(recommendationTypes []string) GetClusterShardingRecommendationsApiRequest {
	r.recommendationTypes = &recommendationTypes
	return r
}

// Shards to which to limit shard-level recommendations. To include multiple shards, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each shard. Setting this parameter skips the cluster-level recommendation types. Omit this parameter to evaluate every shard.
func (r GetClusterShardingRecommendationsApiRequest) Shards(shards []string) GetClusterShardingRecommendationsApiRequest {
	r.shards = &shards
	return r
}

// Namespaces to which to limit collection-level recommendations. A namespace consists of one database and one collection resource written as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to evaluate every namespace.
func (r GetClusterShardingRecommendationsApiRequest) Namespaces(namespaces []string) GetClusterShardingRecommendationsApiRequest {
	r.namespaces = &namespaces
	return r
}

func (r GetClusterShardingRecommendationsApiRequest) Execute() (*ShardingRecommendationsResponse, *http.Response, error) {
	return r.ApiService.GetClusterShardingRecommendationsExecute(r)
}

/*
GetClusterShardingRecommendations Return Sharding Recommendations for One Cluster

Returns the active Sharding Advisor recommendations for one cluster, together with the recommendation types that were evaluated. Recommendations describe changes that may improve the cluster's data distribution, such as adding shards or sharding a collection, together with the monitored conditions that produced them.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterShardingRecommendationsApiRequest
*/
func (a *ShardingRecommendationsAPIService) GetClusterShardingRecommendations(ctx context.Context, groupId string, clusterName string) GetClusterShardingRecommendationsApiRequest {
	return GetClusterShardingRecommendationsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetClusterShardingRecommendationsExecute executes the request
//
//	@return ShardingRecommendationsResponse
func (a *ShardingRecommendationsAPIService) GetClusterShardingRecommendationsExecute(r GetClusterShardingRecommendationsApiRequest) (*ShardingRecommendationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ShardingRecommendationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardingRecommendationsAPIService.GetClusterShardingRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardingRecommendations"
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

	if r.recommendationTypes != nil {
		t := *r.recommendationTypes
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "recommendationTypes", t, "multi")

	}
	if r.shards != nil {
		t := *r.shards
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "shards", t, "multi")

	}
	if r.namespaces != nil {
		t := *r.namespaces
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespaces", t, "multi")

	}
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
