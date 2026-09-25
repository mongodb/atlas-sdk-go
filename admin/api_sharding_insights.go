// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
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

	/*
		GetShardMetricMeasurements Return Measurements for One Shard

		Returns time series of the requested measurements for one shard of the specified sharded cluster. The series average the readings of the shard's member processes selected by `nodeType`. Responds with a 404 only when no shard with the given `replicaSetName` exists in the cluster; a shard that exists but has no readings yet returns series of `null` data points. Each series carries one data point per `granularity` interval, starting at `start` and not passing `end`. A data point's `value` is `null` when MongoDB Cloud has no reading for that interval; this includes the first minutes after a cluster or shard is created, before its metrics are collected, so `null` must not be read as zero load. Set `period`, or set both `start` and `end`; when neither is set, the response covers the last hour. A request whose window and `granularity` would produce more than 1000 data points per series is rejected.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param replicaSetName Human-readable label that identifies the shard, as the replica set name of its members.
		@return GetShardMetricMeasurementsApiRequest
	*/
	GetShardMetricMeasurements(ctx context.Context, groupId string, clusterName string, replicaSetName string) GetShardMetricMeasurementsApiRequest
	/*
		GetShardMetricMeasurements Return Measurements for One Shard


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetShardMetricMeasurementsApiParams - Parameters for the request
		@return GetShardMetricMeasurementsApiRequest
	*/
	GetShardMetricMeasurementsWithParams(ctx context.Context, args *GetShardMetricMeasurementsApiParams) GetShardMetricMeasurementsApiRequest

	// Method available only for mocking purposes
	GetShardMetricMeasurementsExecute(r GetShardMetricMeasurementsApiRequest) (*ShardMeasurementsResponse, *http.Response, error)

	/*
		ListShardMetricMeasurements Return Measurements for All Shards of One Cluster

		Returns time series of the requested measurements for every shard of the specified sharded cluster, one series per shard and measurement. Each shard's series averages the readings of its member processes selected by `nodeType`. Each series carries one data point per `granularity` interval, starting at `start` and not passing `end`. A data point's `value` is `null` when MongoDB Cloud has no reading for that interval; this includes the first minutes after a cluster or shard is created, before its metrics are collected, so `null` must not be read as zero load. Set `period`, or set both `start` and `end`; when neither is set, the response covers the last hour. A request whose window and `granularity` would produce more than 1000 data points per series is rejected.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListShardMetricMeasurementsApiRequest
	*/
	ListShardMetricMeasurements(ctx context.Context, groupId string, clusterName string) ListShardMetricMeasurementsApiRequest
	/*
		ListShardMetricMeasurements Return Measurements for All Shards of One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListShardMetricMeasurementsApiParams - Parameters for the request
		@return ListShardMetricMeasurementsApiRequest
	*/
	ListShardMetricMeasurementsWithParams(ctx context.Context, args *ListShardMetricMeasurementsApiParams) ListShardMetricMeasurementsApiRequest

	// Method available only for mocking purposes
	ListShardMetricMeasurementsExecute(r ListShardMetricMeasurementsApiRequest) (*ClusterShardMeasurementsResponse, *http.Response, error)
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

type GetShardMetricMeasurementsApiRequest struct {
	ctx            context.Context
	ApiService     ShardingInsightsAPI
	groupId        string
	clusterName    string
	replicaSetName string
	m              *[]string
	granularity    *string
	period         *string
	start          *time.Time
	end            *time.Time
	nodeType       *string
}

type GetShardMetricMeasurementsApiParams struct {
	GroupId        string
	ClusterName    string
	ReplicaSetName string
	M              *[]string
	Granularity    *string
	Period         *string
	Start          *time.Time
	End            *time.Time
	NodeType       *string
}

func (a *ShardingInsightsAPIService) GetShardMetricMeasurementsWithParams(ctx context.Context, args *GetShardMetricMeasurementsApiParams) GetShardMetricMeasurementsApiRequest {
	return GetShardMetricMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		replicaSetName: args.ReplicaSetName,
		m:              args.M,
		granularity:    args.Granularity,
		period:         args.Period,
		start:          args.Start,
		end:            args.End,
		nodeType:       args.NodeType,
	}
}

// One or more measurements to return. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value.
func (r GetShardMetricMeasurementsApiRequest) M(m []string) GetShardMetricMeasurementsApiRequest {
	r.m = &m
	return r
}

// Interval between consecutive data points, in ISO 8601 duration format. If omitted, MongoDB Cloud chooses the resolution from the width of the window.
func (r GetShardMetricMeasurementsApiRequest) Granularity(granularity string) GetShardMetricMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetShardMetricMeasurementsApiRequest) Period(period string) GetShardMetricMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetShardMetricMeasurementsApiRequest) Start(start time.Time) GetShardMetricMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetShardMetricMeasurementsApiRequest) End(end time.Time) GetShardMetricMeasurementsApiRequest {
	r.end = &end
	return r
}

// Members of the shard whose readings are averaged into the shard&#39;s series.
func (r GetShardMetricMeasurementsApiRequest) NodeType(nodeType string) GetShardMetricMeasurementsApiRequest {
	r.nodeType = &nodeType
	return r
}

func (r GetShardMetricMeasurementsApiRequest) Execute() (*ShardMeasurementsResponse, *http.Response, error) {
	return r.ApiService.GetShardMetricMeasurementsExecute(r)
}

/*
GetShardMetricMeasurements Return Measurements for One Shard

Returns time series of the requested measurements for one shard of the specified sharded cluster. The series average the readings of the shard's member processes selected by `nodeType`. Responds with a 404 only when no shard with the given `replicaSetName` exists in the cluster; a shard that exists but has no readings yet returns series of `null` data points. Each series carries one data point per `granularity` interval, starting at `start` and not passing `end`. A data point's `value` is `null` when MongoDB Cloud has no reading for that interval; this includes the first minutes after a cluster or shard is created, before its metrics are collected, so `null` must not be read as zero load. Set `period`, or set both `start` and `end`; when neither is set, the response covers the last hour. A request whose window and `granularity` would produce more than 1000 data points per series is rejected.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param replicaSetName Human-readable label that identifies the shard, as the replica set name of its members.
	@return GetShardMetricMeasurementsApiRequest
*/
func (a *ShardingInsightsAPIService) GetShardMetricMeasurements(ctx context.Context, groupId string, clusterName string, replicaSetName string) GetShardMetricMeasurementsApiRequest {
	return GetShardMetricMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		replicaSetName: replicaSetName,
	}
}

// GetShardMetricMeasurementsExecute executes the request
//
//	@return ShardMeasurementsResponse
func (a *ShardingInsightsAPIService) GetShardMetricMeasurementsExecute(r GetShardMetricMeasurementsApiRequest) (*ShardMeasurementsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ShardMeasurementsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardingInsightsAPIService.GetShardMetricMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics/{replicaSetName}/measurements"
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
	if r.replicaSetName == "" {
		return localVarReturnValue, nil, reportError("replicaSetName is empty and must be specified")
	}
	if r.replicaSetName == "." || r.replicaSetName == ".." {
		return localVarReturnValue, nil, reportError("replicaSetName must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"replicaSetName"+"}", url.PathEscape(r.replicaSetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.m != nil {
		t := *r.m
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "m", t, "multi")

	}
	if r.granularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.nodeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nodeType", r.nodeType, "")
	} else {
		var defaultValue string = "ALL"
		r.nodeType = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nodeType", r.nodeType, "")
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

type ListShardMetricMeasurementsApiRequest struct {
	ctx         context.Context
	ApiService  ShardingInsightsAPI
	groupId     string
	clusterName string
	m           *[]string
	granularity *string
	period      *string
	start       *time.Time
	end         *time.Time
	nodeType    *string
}

type ListShardMetricMeasurementsApiParams struct {
	GroupId     string
	ClusterName string
	M           *[]string
	Granularity *string
	Period      *string
	Start       *time.Time
	End         *time.Time
	NodeType    *string
}

func (a *ShardingInsightsAPIService) ListShardMetricMeasurementsWithParams(ctx context.Context, args *ListShardMetricMeasurementsApiParams) ListShardMetricMeasurementsApiRequest {
	return ListShardMetricMeasurementsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		m:           args.M,
		granularity: args.Granularity,
		period:      args.Period,
		start:       args.Start,
		end:         args.End,
		nodeType:    args.NodeType,
	}
}

// One or more measurements to return. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value.
func (r ListShardMetricMeasurementsApiRequest) M(m []string) ListShardMetricMeasurementsApiRequest {
	r.m = &m
	return r
}

// Interval between consecutive data points, in ISO 8601 duration format. If omitted, MongoDB Cloud chooses the resolution from the width of the window.
func (r ListShardMetricMeasurementsApiRequest) Granularity(granularity string) ListShardMetricMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r ListShardMetricMeasurementsApiRequest) Period(period string) ListShardMetricMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r ListShardMetricMeasurementsApiRequest) Start(start time.Time) ListShardMetricMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r ListShardMetricMeasurementsApiRequest) End(end time.Time) ListShardMetricMeasurementsApiRequest {
	r.end = &end
	return r
}

// Members of each shard whose readings are averaged into the shard&#39;s series.
func (r ListShardMetricMeasurementsApiRequest) NodeType(nodeType string) ListShardMetricMeasurementsApiRequest {
	r.nodeType = &nodeType
	return r
}

func (r ListShardMetricMeasurementsApiRequest) Execute() (*ClusterShardMeasurementsResponse, *http.Response, error) {
	return r.ApiService.ListShardMetricMeasurementsExecute(r)
}

/*
ListShardMetricMeasurements Return Measurements for All Shards of One Cluster

Returns time series of the requested measurements for every shard of the specified sharded cluster, one series per shard and measurement. Each shard's series averages the readings of its member processes selected by `nodeType`. Each series carries one data point per `granularity` interval, starting at `start` and not passing `end`. A data point's `value` is `null` when MongoDB Cloud has no reading for that interval; this includes the first minutes after a cluster or shard is created, before its metrics are collected, so `null` must not be read as zero load. Set `period`, or set both `start` and `end`; when neither is set, the response covers the last hour. A request whose window and `granularity` would produce more than 1000 data points per series is rejected.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListShardMetricMeasurementsApiRequest
*/
func (a *ShardingInsightsAPIService) ListShardMetricMeasurements(ctx context.Context, groupId string, clusterName string) ListShardMetricMeasurementsApiRequest {
	return ListShardMetricMeasurementsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListShardMetricMeasurementsExecute executes the request
//
//	@return ClusterShardMeasurementsResponse
func (a *ShardingInsightsAPIService) ListShardMetricMeasurementsExecute(r ListShardMetricMeasurementsApiRequest) (*ClusterShardMeasurementsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterShardMeasurementsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardingInsightsAPIService.ListShardMetricMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics:listMeasurements"
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

	if r.m != nil {
		t := *r.m
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "m", t, "multi")

	}
	if r.granularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.nodeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nodeType", r.nodeType, "")
	} else {
		var defaultValue string = "ALL"
		r.nodeType = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nodeType", r.nodeType, "")
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
