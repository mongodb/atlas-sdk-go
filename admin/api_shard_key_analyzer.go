// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ShardKeyAnalyzerAPI interface {

	/*
		CreateShardKeyAnalysis Create One Shard Key Analysis for One Cluster

		Analyzes how well one candidate shard key would distribute one collection across the shards of the specified cluster. This endpoint accepts the request and returns immediately: the `Location` header carries the URI of the operation that tracks the analysis. Poll that operation until its `status` is terminal, then follow its `resultHref` to read the analysis. MongoDB Cloud validates and authorizes the request before accepting it, so a rejected submission returns a `4xx` status and never an operation in a failed state. A namespace that is well formed but absent cannot be detected before the analysis runs and surfaces as a failed operation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that contains the collection to analyze.
		@param shardKeyAnalysisRequest Shard key to analyze, and which metrics to compute for it.
		@return CreateShardKeyAnalysisApiRequest
	*/
	CreateShardKeyAnalysis(ctx context.Context, groupId string, clusterName string, shardKeyAnalysisRequest *ShardKeyAnalysisRequest) CreateShardKeyAnalysisApiRequest
	/*
		CreateShardKeyAnalysis Create One Shard Key Analysis for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateShardKeyAnalysisApiParams - Parameters for the request
		@return CreateShardKeyAnalysisApiRequest
	*/
	CreateShardKeyAnalysisWithParams(ctx context.Context, args *CreateShardKeyAnalysisApiParams) CreateShardKeyAnalysisApiRequest

	// Method available only for mocking purposes
	CreateShardKeyAnalysisExecute(r CreateShardKeyAnalysisApiRequest) (*http.Response, error)

	/*
		GetAnalysisOperation Return One Shard Key Analysis Operation for One Cluster

		Returns the progress of one shard key analysis for the specified cluster. Poll this resource until `status` reaches a terminal value. On success it carries `resultHref`, the URI of the analysis it created; on failure it carries `error`. While the status is not terminal it carries `retryAfterSeconds`, how long to wait before polling again. MongoDB Cloud retains an operation for 30 days past its terminal state.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the shard key analysis operation you want to return.
		@param operationId Unique 24-hexadecimal digit string that identifies the shard key analysis operation to return.
		@return GetAnalysisOperationApiRequest
	*/
	GetAnalysisOperation(ctx context.Context, groupId string, clusterName string, operationId string) GetAnalysisOperationApiRequest
	/*
		GetAnalysisOperation Return One Shard Key Analysis Operation for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAnalysisOperationApiParams - Parameters for the request
		@return GetAnalysisOperationApiRequest
	*/
	GetAnalysisOperationWithParams(ctx context.Context, args *GetAnalysisOperationApiParams) GetAnalysisOperationApiRequest

	// Method available only for mocking purposes
	GetAnalysisOperationExecute(r GetAnalysisOperationApiRequest) (*ShardKeyAnalysisOperationResponse, *http.Response, error)

	/*
		GetShardKeyAnalysis Return One Shard Key Analysis for One Cluster

		Returns one completed shard key analysis for the specified cluster. This resource exists only once the operation that produced it has succeeded and reports no operation status of its own. MongoDB Cloud retains an analysis for 30 days.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the shard key analysis you want to return.
		@param analysisId Unique 24-hexadecimal digit string that identifies the shard key analysis to return.
		@return GetShardKeyAnalysisApiRequest
	*/
	GetShardKeyAnalysis(ctx context.Context, groupId string, clusterName string, analysisId string) GetShardKeyAnalysisApiRequest
	/*
		GetShardKeyAnalysis Return One Shard Key Analysis for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetShardKeyAnalysisApiParams - Parameters for the request
		@return GetShardKeyAnalysisApiRequest
	*/
	GetShardKeyAnalysisWithParams(ctx context.Context, args *GetShardKeyAnalysisApiParams) GetShardKeyAnalysisApiRequest

	// Method available only for mocking purposes
	GetShardKeyAnalysisExecute(r GetShardKeyAnalysisApiRequest) (*ShardKeyAnalysisResponse, *http.Response, error)

	/*
		ListAnalysisOperations Return All Shard Key Analysis Operations for One Cluster

		Returns every shard key analysis attempt for the specified cluster, newest first, including attempts that are still running and attempts that failed. Neither of those produces an analysis, so this sub-resource is the only place they appear. MongoDB Cloud retains an operation for 30 days past its terminal state.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the shard key analysis operations you want to return.
		@return ListAnalysisOperationsApiRequest
	*/
	ListAnalysisOperations(ctx context.Context, groupId string, clusterName string) ListAnalysisOperationsApiRequest
	/*
		ListAnalysisOperations Return All Shard Key Analysis Operations for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAnalysisOperationsApiParams - Parameters for the request
		@return ListAnalysisOperationsApiRequest
	*/
	ListAnalysisOperationsWithParams(ctx context.Context, args *ListAnalysisOperationsApiParams) ListAnalysisOperationsApiRequest

	// Method available only for mocking purposes
	ListAnalysisOperationsExecute(r ListAnalysisOperationsApiRequest) (*PaginatedShardKeyAnalysisOperationResponse, *http.Response, error)

	/*
		ListShardKeyAnalyses Return All Shard Key Analyses for One Cluster

		Returns all completed shard key analyses for the specified cluster, newest first. An analysis appears here only once the operation that produced it has succeeded, so attempts that failed or are still running are absent; read the operations sub-resource to see those. MongoDB Cloud retains an analysis for 30 days.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the shard key analyses you want to return.
		@return ListShardKeyAnalysesApiRequest
	*/
	ListShardKeyAnalyses(ctx context.Context, groupId string, clusterName string) ListShardKeyAnalysesApiRequest
	/*
		ListShardKeyAnalyses Return All Shard Key Analyses for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListShardKeyAnalysesApiParams - Parameters for the request
		@return ListShardKeyAnalysesApiRequest
	*/
	ListShardKeyAnalysesWithParams(ctx context.Context, args *ListShardKeyAnalysesApiParams) ListShardKeyAnalysesApiRequest

	// Method available only for mocking purposes
	ListShardKeyAnalysesExecute(r ListShardKeyAnalysesApiRequest) (*PaginatedShardKeyAnalysisResponse, *http.Response, error)
}

// ShardKeyAnalyzerAPIService ShardKeyAnalyzerAPI service
type ShardKeyAnalyzerAPIService service

type CreateShardKeyAnalysisApiRequest struct {
	ctx                     context.Context
	ApiService              ShardKeyAnalyzerAPI
	groupId                 string
	clusterName             string
	shardKeyAnalysisRequest *ShardKeyAnalysisRequest
}

type CreateShardKeyAnalysisApiParams struct {
	GroupId                 string
	ClusterName             string
	ShardKeyAnalysisRequest *ShardKeyAnalysisRequest
}

func (a *ShardKeyAnalyzerAPIService) CreateShardKeyAnalysisWithParams(ctx context.Context, args *CreateShardKeyAnalysisApiParams) CreateShardKeyAnalysisApiRequest {
	return CreateShardKeyAnalysisApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		clusterName:             args.ClusterName,
		shardKeyAnalysisRequest: args.ShardKeyAnalysisRequest,
	}
}

func (r CreateShardKeyAnalysisApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateShardKeyAnalysisExecute(r)
}

/*
CreateShardKeyAnalysis Create One Shard Key Analysis for One Cluster

Analyzes how well one candidate shard key would distribute one collection across the shards of the specified cluster. This endpoint accepts the request and returns immediately: the `Location` header carries the URI of the operation that tracks the analysis. Poll that operation until its `status` is terminal, then follow its `resultHref` to read the analysis. MongoDB Cloud validates and authorizes the request before accepting it, so a rejected submission returns a `4xx` status and never an operation in a failed state. A namespace that is well formed but absent cannot be detected before the analysis runs and surfaces as a failed operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that contains the collection to analyze.
	@return CreateShardKeyAnalysisApiRequest
*/
func (a *ShardKeyAnalyzerAPIService) CreateShardKeyAnalysis(ctx context.Context, groupId string, clusterName string, shardKeyAnalysisRequest *ShardKeyAnalysisRequest) CreateShardKeyAnalysisApiRequest {
	return CreateShardKeyAnalysisApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		clusterName:             clusterName,
		shardKeyAnalysisRequest: shardKeyAnalysisRequest,
	}
}

// CreateShardKeyAnalysisExecute executes the request
func (a *ShardKeyAnalyzerAPIService) CreateShardKeyAnalysisExecute(r CreateShardKeyAnalysisApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardKeyAnalyzerAPIService.CreateShardKeyAnalysis")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	if r.groupId == "." || r.groupId == ".." {
		return nil, reportError("groupId must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	if r.clusterName == "." || r.clusterName == ".." {
		return nil, reportError("clusterName must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shardKeyAnalysisRequest == nil {
		return nil, reportError("shardKeyAnalysisRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

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
	// body params
	localVarPostBody = r.shardKeyAnalysisRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetAnalysisOperationApiRequest struct {
	ctx         context.Context
	ApiService  ShardKeyAnalyzerAPI
	groupId     string
	clusterName string
	operationId string
}

type GetAnalysisOperationApiParams struct {
	GroupId     string
	ClusterName string
	OperationId string
}

func (a *ShardKeyAnalyzerAPIService) GetAnalysisOperationWithParams(ctx context.Context, args *GetAnalysisOperationApiParams) GetAnalysisOperationApiRequest {
	return GetAnalysisOperationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		operationId: args.OperationId,
	}
}

func (r GetAnalysisOperationApiRequest) Execute() (*ShardKeyAnalysisOperationResponse, *http.Response, error) {
	return r.ApiService.GetAnalysisOperationExecute(r)
}

/*
GetAnalysisOperation Return One Shard Key Analysis Operation for One Cluster

Returns the progress of one shard key analysis for the specified cluster. Poll this resource until `status` reaches a terminal value. On success it carries `resultHref`, the URI of the analysis it created; on failure it carries `error`. While the status is not terminal it carries `retryAfterSeconds`, how long to wait before polling again. MongoDB Cloud retains an operation for 30 days past its terminal state.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the shard key analysis operation you want to return.
	@param operationId Unique 24-hexadecimal digit string that identifies the shard key analysis operation to return.
	@return GetAnalysisOperationApiRequest
*/
func (a *ShardKeyAnalyzerAPIService) GetAnalysisOperation(ctx context.Context, groupId string, clusterName string, operationId string) GetAnalysisOperationApiRequest {
	return GetAnalysisOperationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		operationId: operationId,
	}
}

// GetAnalysisOperationExecute executes the request
//
//	@return ShardKeyAnalysisOperationResponse
func (a *ShardKeyAnalyzerAPIService) GetAnalysisOperationExecute(r GetAnalysisOperationApiRequest) (*ShardKeyAnalysisOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ShardKeyAnalysisOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardKeyAnalyzerAPIService.GetAnalysisOperation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/operations/{operationId}"
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
	if r.operationId == "" {
		return localVarReturnValue, nil, reportError("operationId is empty and must be specified")
	}
	if r.operationId == "." || r.operationId == ".." {
		return localVarReturnValue, nil, reportError("operationId must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"operationId"+"}", url.PathEscape(r.operationId), -1)

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

type GetShardKeyAnalysisApiRequest struct {
	ctx         context.Context
	ApiService  ShardKeyAnalyzerAPI
	groupId     string
	clusterName string
	analysisId  string
}

type GetShardKeyAnalysisApiParams struct {
	GroupId     string
	ClusterName string
	AnalysisId  string
}

func (a *ShardKeyAnalyzerAPIService) GetShardKeyAnalysisWithParams(ctx context.Context, args *GetShardKeyAnalysisApiParams) GetShardKeyAnalysisApiRequest {
	return GetShardKeyAnalysisApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		analysisId:  args.AnalysisId,
	}
}

func (r GetShardKeyAnalysisApiRequest) Execute() (*ShardKeyAnalysisResponse, *http.Response, error) {
	return r.ApiService.GetShardKeyAnalysisExecute(r)
}

/*
GetShardKeyAnalysis Return One Shard Key Analysis for One Cluster

Returns one completed shard key analysis for the specified cluster. This resource exists only once the operation that produced it has succeeded and reports no operation status of its own. MongoDB Cloud retains an analysis for 30 days.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the shard key analysis you want to return.
	@param analysisId Unique 24-hexadecimal digit string that identifies the shard key analysis to return.
	@return GetShardKeyAnalysisApiRequest
*/
func (a *ShardKeyAnalyzerAPIService) GetShardKeyAnalysis(ctx context.Context, groupId string, clusterName string, analysisId string) GetShardKeyAnalysisApiRequest {
	return GetShardKeyAnalysisApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		analysisId:  analysisId,
	}
}

// GetShardKeyAnalysisExecute executes the request
//
//	@return ShardKeyAnalysisResponse
func (a *ShardKeyAnalyzerAPIService) GetShardKeyAnalysisExecute(r GetShardKeyAnalysisApiRequest) (*ShardKeyAnalysisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ShardKeyAnalysisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardKeyAnalyzerAPIService.GetShardKeyAnalysis")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/{analysisId}"
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
	if r.analysisId == "" {
		return localVarReturnValue, nil, reportError("analysisId is empty and must be specified")
	}
	if r.analysisId == "." || r.analysisId == ".." {
		return localVarReturnValue, nil, reportError("analysisId must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"analysisId"+"}", url.PathEscape(r.analysisId), -1)

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

type ListAnalysisOperationsApiRequest struct {
	ctx          context.Context
	ApiService   ShardKeyAnalyzerAPI
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListAnalysisOperationsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ShardKeyAnalyzerAPIService) ListAnalysisOperationsWithParams(ctx context.Context, args *ListAnalysisOperationsApiParams) ListAnalysisOperationsApiRequest {
	return ListAnalysisOperationsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether MongoDB Cloud calculates the total number of items for the response. When set to &#x60;false&#x60;, MongoDB Cloud may skip an additional count operation. The response may still include &#x60;totalCount&#x60; when the count is available without additional calculation.
func (r ListAnalysisOperationsApiRequest) IncludeCount(includeCount bool) ListAnalysisOperationsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAnalysisOperationsApiRequest) ItemsPerPage(itemsPerPage int) ListAnalysisOperationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAnalysisOperationsApiRequest) PageNum(pageNum int) ListAnalysisOperationsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListAnalysisOperationsApiRequest) Execute() (*PaginatedShardKeyAnalysisOperationResponse, *http.Response, error) {
	return r.ApiService.ListAnalysisOperationsExecute(r)
}

/*
ListAnalysisOperations Return All Shard Key Analysis Operations for One Cluster

Returns every shard key analysis attempt for the specified cluster, newest first, including attempts that are still running and attempts that failed. Neither of those produces an analysis, so this sub-resource is the only place they appear. MongoDB Cloud retains an operation for 30 days past its terminal state.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the shard key analysis operations you want to return.
	@return ListAnalysisOperationsApiRequest
*/
func (a *ShardKeyAnalyzerAPIService) ListAnalysisOperations(ctx context.Context, groupId string, clusterName string) ListAnalysisOperationsApiRequest {
	return ListAnalysisOperationsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListAnalysisOperationsExecute executes the request
//
//	@return PaginatedShardKeyAnalysisOperationResponse
func (a *ShardKeyAnalyzerAPIService) ListAnalysisOperationsExecute(r ListAnalysisOperationsApiRequest) (*PaginatedShardKeyAnalysisOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedShardKeyAnalysisOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardKeyAnalyzerAPIService.ListAnalysisOperations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/operations"
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

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
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

type ListShardKeyAnalysesApiRequest struct {
	ctx          context.Context
	ApiService   ShardKeyAnalyzerAPI
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListShardKeyAnalysesApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ShardKeyAnalyzerAPIService) ListShardKeyAnalysesWithParams(ctx context.Context, args *ListShardKeyAnalysesApiParams) ListShardKeyAnalysesApiRequest {
	return ListShardKeyAnalysesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether MongoDB Cloud calculates the total number of items for the response. When set to &#x60;false&#x60;, MongoDB Cloud may skip an additional count operation. The response may still include &#x60;totalCount&#x60; when the count is available without additional calculation.
func (r ListShardKeyAnalysesApiRequest) IncludeCount(includeCount bool) ListShardKeyAnalysesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListShardKeyAnalysesApiRequest) ItemsPerPage(itemsPerPage int) ListShardKeyAnalysesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListShardKeyAnalysesApiRequest) PageNum(pageNum int) ListShardKeyAnalysesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListShardKeyAnalysesApiRequest) Execute() (*PaginatedShardKeyAnalysisResponse, *http.Response, error) {
	return r.ApiService.ListShardKeyAnalysesExecute(r)
}

/*
ListShardKeyAnalyses Return All Shard Key Analyses for One Cluster

Returns all completed shard key analyses for the specified cluster, newest first. An analysis appears here only once the operation that produced it has succeeded, so attempts that failed or are still running are absent; read the operations sub-resource to see those. MongoDB Cloud retains an analysis for 30 days.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the shard key analyses you want to return.
	@return ListShardKeyAnalysesApiRequest
*/
func (a *ShardKeyAnalyzerAPIService) ListShardKeyAnalyses(ctx context.Context, groupId string, clusterName string) ListShardKeyAnalysesApiRequest {
	return ListShardKeyAnalysesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListShardKeyAnalysesExecute executes the request
//
//	@return PaginatedShardKeyAnalysisResponse
func (a *ShardKeyAnalyzerAPIService) ListShardKeyAnalysesExecute(r ListShardKeyAnalysesApiRequest) (*PaginatedShardKeyAnalysisResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedShardKeyAnalysisResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShardKeyAnalyzerAPIService.ListShardKeyAnalyses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses"
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

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
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
