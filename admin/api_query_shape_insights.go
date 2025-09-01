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

type QueryShapeInsightsApi interface {

	/*
		GetGroupClusterQueryShapeInsightDetails Return Query Shape Details

		Returns the metadata and statistics summary for a given query shape hash.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param queryShapeHash A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs.
		@return GetGroupClusterQueryShapeInsightDetailsApiRequest
	*/
	GetGroupClusterQueryShapeInsightDetails(ctx context.Context, groupId string, clusterName string, queryShapeHash string) GetGroupClusterQueryShapeInsightDetailsApiRequest
	/*
		GetGroupClusterQueryShapeInsightDetails Return Query Shape Details


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupClusterQueryShapeInsightDetailsApiParams - Parameters for the request
		@return GetGroupClusterQueryShapeInsightDetailsApiRequest
	*/
	GetGroupClusterQueryShapeInsightDetailsWithParams(ctx context.Context, args *GetGroupClusterQueryShapeInsightDetailsApiParams) GetGroupClusterQueryShapeInsightDetailsApiRequest

	// Method available only for mocking purposes
	GetGroupClusterQueryShapeInsightDetailsExecute(r GetGroupClusterQueryShapeInsightDetailsApiRequest) (*QueryStatsDetailsResponse, *http.Response, error)

	/*
		GetGroupClusterQueryShapeInsightSummaries Return Query Statistic Summaries

		Returns a list of query shape statistics summaries for a given cluster. Query shape statistics provide performance insights about MongoDB queries, helping users identify problematic query patterns and potential optimizations.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetGroupClusterQueryShapeInsightSummariesApiRequest
	*/
	GetGroupClusterQueryShapeInsightSummaries(ctx context.Context, groupId string, clusterName string) GetGroupClusterQueryShapeInsightSummariesApiRequest
	/*
		GetGroupClusterQueryShapeInsightSummaries Return Query Statistic Summaries


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupClusterQueryShapeInsightSummariesApiParams - Parameters for the request
		@return GetGroupClusterQueryShapeInsightSummariesApiRequest
	*/
	GetGroupClusterQueryShapeInsightSummariesWithParams(ctx context.Context, args *GetGroupClusterQueryShapeInsightSummariesApiParams) GetGroupClusterQueryShapeInsightSummariesApiRequest

	// Method available only for mocking purposes
	GetGroupClusterQueryShapeInsightSummariesExecute(r GetGroupClusterQueryShapeInsightSummariesApiRequest) (*QueryStatsSummaryListResponse, *http.Response, error)
}

// QueryShapeInsightsApiService QueryShapeInsightsApi service
type QueryShapeInsightsApiService service

type GetGroupClusterQueryShapeInsightDetailsApiRequest struct {
	ctx            context.Context
	ApiService     QueryShapeInsightsApi
	groupId        string
	clusterName    string
	queryShapeHash string
	since          *int64
	until          *int64
	processIds     *[]string
}

type GetGroupClusterQueryShapeInsightDetailsApiParams struct {
	GroupId        string
	ClusterName    string
	QueryShapeHash string
	Since          *int64
	Until          *int64
	ProcessIds     *[]string
}

func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightDetailsWithParams(ctx context.Context, args *GetGroupClusterQueryShapeInsightDetailsApiParams) GetGroupClusterQueryShapeInsightDetailsApiRequest {
	return GetGroupClusterQueryShapeInsightDetailsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		queryShapeHash: args.QueryShapeHash,
		since:          args.Since,
		until:          args.Until,
		processIds:     args.ProcessIds,
	}
}

// Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.
func (r GetGroupClusterQueryShapeInsightDetailsApiRequest) Since(since int64) GetGroupClusterQueryShapeInsightDetailsApiRequest {
	r.since = &since
	return r
}

// Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you specify the **until** parameter, you must specify the **since** parameter. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.
func (r GetGroupClusterQueryShapeInsightDetailsApiRequest) Until(until int64) GetGroupClusterQueryShapeInsightDetailsApiRequest {
	r.until = &until
	return r
}

// ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each processId.
func (r GetGroupClusterQueryShapeInsightDetailsApiRequest) ProcessIds(processIds []string) GetGroupClusterQueryShapeInsightDetailsApiRequest {
	r.processIds = &processIds
	return r
}

func (r GetGroupClusterQueryShapeInsightDetailsApiRequest) Execute() (*QueryStatsDetailsResponse, *http.Response, error) {
	return r.ApiService.GetGroupClusterQueryShapeInsightDetailsExecute(r)
}

/*
GetGroupClusterQueryShapeInsightDetails Return Query Shape Details

Returns the metadata and statistics summary for a given query shape hash.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param queryShapeHash A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs.
	@return GetGroupClusterQueryShapeInsightDetailsApiRequest
*/
func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightDetails(ctx context.Context, groupId string, clusterName string, queryShapeHash string) GetGroupClusterQueryShapeInsightDetailsApiRequest {
	return GetGroupClusterQueryShapeInsightDetailsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		queryShapeHash: queryShapeHash,
	}
}

// GetGroupClusterQueryShapeInsightDetailsExecute executes the request
//
//	@return QueryStatsDetailsResponse
func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightDetailsExecute(r GetGroupClusterQueryShapeInsightDetailsApiRequest) (*QueryStatsDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *QueryStatsDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryShapeInsightsApiService.GetGroupClusterQueryShapeInsightDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/{queryShapeHash}/details"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.queryShapeHash == "" {
		return localVarReturnValue, nil, reportError("queryShapeHash is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"queryShapeHash"+"}", url.PathEscape(r.queryShapeHash), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "")
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "")
	}
	if r.processIds != nil {
		t := *r.processIds
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "processIds", t, "multi")

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

type GetGroupClusterQueryShapeInsightSummariesApiRequest struct {
	ctx              context.Context
	ApiService       QueryShapeInsightsApi
	groupId          string
	clusterName      string
	since            *int64
	until            *int64
	processIds       *[]string
	namespaces       *[]string
	commands         *[]string
	nSummaries       *int64
	series           *[]string
	queryShapeHashes *[]string
}

type GetGroupClusterQueryShapeInsightSummariesApiParams struct {
	GroupId          string
	ClusterName      string
	Since            *int64
	Until            *int64
	ProcessIds       *[]string
	Namespaces       *[]string
	Commands         *[]string
	NSummaries       *int64
	Series           *[]string
	QueryShapeHashes *[]string
}

func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightSummariesWithParams(ctx context.Context, args *GetGroupClusterQueryShapeInsightSummariesApiParams) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	return GetGroupClusterQueryShapeInsightSummariesApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		clusterName:      args.ClusterName,
		since:            args.Since,
		until:            args.Until,
		processIds:       args.ProcessIds,
		namespaces:       args.Namespaces,
		commands:         args.Commands,
		nSummaries:       args.NSummaries,
		series:           args.Series,
		queryShapeHashes: args.QueryShapeHashes,
	}
}

// Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Since(since int64) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.since = &since
	return r
}

// Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you specify the **until** parameter, you must specify the **since** parameter. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Until(until int64) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.until = &until
	return r
}

// ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each processId.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) ProcessIds(processIds []string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.processIds = &processIds
	return r
}

// Namespaces from which to retrieve query shape statistics. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Namespaces(namespaces []string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.namespaces = &namespaces
	return r
}

// Retrieve query shape statistics matching specified MongoDB commands. To include multiple commands, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each command. The currently supported parameters are find, distinct, and aggregate. Omit this parameter to return results for all supported commands.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Commands(commands []string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.commands = &commands
	return r
}

// Maximum number of query statistic summaries to return.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) NSummaries(nSummaries int64) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.nSummaries = &nSummaries
	return r
}

// Query shape statistics data series to retrieve. A series represents a specific metric about query execution. To include multiple series, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each series. Omit this parameter to return results for all available series.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Series(series []string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.series = &series
	return r
}

// A list of SHA256 hashes of desired query shapes, output by MongoDB commands like $queryStats and $explain or slow query logs. To include multiple series, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each series. Omit this parameter to return results for all available series.
func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) QueryShapeHashes(queryShapeHashes []string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	r.queryShapeHashes = &queryShapeHashes
	return r
}

func (r GetGroupClusterQueryShapeInsightSummariesApiRequest) Execute() (*QueryStatsSummaryListResponse, *http.Response, error) {
	return r.ApiService.GetGroupClusterQueryShapeInsightSummariesExecute(r)
}

/*
GetGroupClusterQueryShapeInsightSummaries Return Query Statistic Summaries

Returns a list of query shape statistics summaries for a given cluster. Query shape statistics provide performance insights about MongoDB queries, helping users identify problematic query patterns and potential optimizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetGroupClusterQueryShapeInsightSummariesApiRequest
*/
func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightSummaries(ctx context.Context, groupId string, clusterName string) GetGroupClusterQueryShapeInsightSummariesApiRequest {
	return GetGroupClusterQueryShapeInsightSummariesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetGroupClusterQueryShapeInsightSummariesExecute executes the request
//
//	@return QueryStatsSummaryListResponse
func (a *QueryShapeInsightsApiService) GetGroupClusterQueryShapeInsightSummariesExecute(r GetGroupClusterQueryShapeInsightSummariesApiRequest) (*QueryStatsSummaryListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *QueryStatsSummaryListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryShapeInsightsApiService.GetGroupClusterQueryShapeInsightSummaries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/summaries"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "")
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "")
	}
	if r.processIds != nil {
		t := *r.processIds
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "processIds", t, "multi")

	}
	if r.namespaces != nil {
		t := *r.namespaces
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespaces", t, "multi")

	}
	if r.commands != nil {
		t := *r.commands
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "commands", t, "multi")

	}
	if r.nSummaries != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nSummaries", r.nSummaries, "")
	} else {
		var defaultValue int64 = 100
		r.nSummaries = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nSummaries", r.nSummaries, "")
	}
	if r.series != nil {
		t := *r.series
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "series", t, "multi")

	}
	if r.queryShapeHashes != nil {
		t := *r.queryShapeHashes
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "queryShapeHashes", t, "multi")

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
