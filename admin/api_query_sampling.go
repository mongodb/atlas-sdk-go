// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type QuerySamplingAPI interface {

	/*
		CreateClusterQuerySampling Create Query Sampling for One Namespace

		Starts query sampling on one namespace of the specified cluster. Sampled queries feed the shard key analysis for the namespace. The response reports the observed sampling state, in which `active` may still be false: the cluster applies the sampling configuration asynchronously.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param querySamplingCreateRequest Namespace on which to start query sampling.
		@return CreateClusterQuerySamplingApiRequest
	*/
	CreateClusterQuerySampling(ctx context.Context, groupId string, clusterName string, querySamplingCreateRequest *QuerySamplingCreateRequest) CreateClusterQuerySamplingApiRequest
	/*
		CreateClusterQuerySampling Create Query Sampling for One Namespace


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateClusterQuerySamplingApiParams - Parameters for the request
		@return CreateClusterQuerySamplingApiRequest
	*/
	CreateClusterQuerySamplingWithParams(ctx context.Context, args *CreateClusterQuerySamplingApiParams) CreateClusterQuerySamplingApiRequest

	// Method available only for mocking purposes
	CreateClusterQuerySamplingExecute(r CreateClusterQuerySamplingApiRequest) (*QuerySamplingResponse, *http.Response, error)

	/*
		DeleteClusterQuerySampling Remove Query Sampling from One Namespace

		Stops query sampling on one namespace of the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param namespace Human-readable label that identifies the namespace (`<database>.<collection>`) on which queries are sampled.
		@return DeleteClusterQuerySamplingApiRequest
	*/
	DeleteClusterQuerySampling(ctx context.Context, groupId string, clusterName string, namespace string) DeleteClusterQuerySamplingApiRequest
	/*
		DeleteClusterQuerySampling Remove Query Sampling from One Namespace


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteClusterQuerySamplingApiParams - Parameters for the request
		@return DeleteClusterQuerySamplingApiRequest
	*/
	DeleteClusterQuerySamplingWithParams(ctx context.Context, args *DeleteClusterQuerySamplingApiParams) DeleteClusterQuerySamplingApiRequest

	// Method available only for mocking purposes
	DeleteClusterQuerySamplingExecute(r DeleteClusterQuerySamplingApiRequest) (*http.Response, error)

	/*
		GetClusterQuerySampling Return Query Sampling State for One Namespace

		Returns the query sampling state of one namespace of the specified cluster, including counters for the operations sampled so far. Returns 404 if query sampling is not configured on the namespace. `active` becomes true asynchronously after query sampling starts.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param namespace Human-readable label that identifies the namespace (`<database>.<collection>`) on which queries are sampled.
		@return GetClusterQuerySamplingApiRequest
	*/
	GetClusterQuerySampling(ctx context.Context, groupId string, clusterName string, namespace string) GetClusterQuerySamplingApiRequest
	/*
		GetClusterQuerySampling Return Query Sampling State for One Namespace


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterQuerySamplingApiParams - Parameters for the request
		@return GetClusterQuerySamplingApiRequest
	*/
	GetClusterQuerySamplingWithParams(ctx context.Context, args *GetClusterQuerySamplingApiParams) GetClusterQuerySamplingApiRequest

	// Method available only for mocking purposes
	GetClusterQuerySamplingExecute(r GetClusterQuerySamplingApiRequest) (*QuerySamplingResponse, *http.Response, error)
}

// QuerySamplingAPIService QuerySamplingAPI service
type QuerySamplingAPIService service

type CreateClusterQuerySamplingApiRequest struct {
	ctx                        context.Context
	ApiService                 QuerySamplingAPI
	groupId                    string
	clusterName                string
	querySamplingCreateRequest *QuerySamplingCreateRequest
}

type CreateClusterQuerySamplingApiParams struct {
	GroupId                    string
	ClusterName                string
	QuerySamplingCreateRequest *QuerySamplingCreateRequest
}

func (a *QuerySamplingAPIService) CreateClusterQuerySamplingWithParams(ctx context.Context, args *CreateClusterQuerySamplingApiParams) CreateClusterQuerySamplingApiRequest {
	return CreateClusterQuerySamplingApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		querySamplingCreateRequest: args.QuerySamplingCreateRequest,
	}
}

func (r CreateClusterQuerySamplingApiRequest) Execute() (*QuerySamplingResponse, *http.Response, error) {
	return r.ApiService.CreateClusterQuerySamplingExecute(r)
}

/*
CreateClusterQuerySampling Create Query Sampling for One Namespace

Starts query sampling on one namespace of the specified cluster. Sampled queries feed the shard key analysis for the namespace. The response reports the observed sampling state, in which `active` may still be false: the cluster applies the sampling configuration asynchronously.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return CreateClusterQuerySamplingApiRequest
*/
func (a *QuerySamplingAPIService) CreateClusterQuerySampling(ctx context.Context, groupId string, clusterName string, querySamplingCreateRequest *QuerySamplingCreateRequest) CreateClusterQuerySamplingApiRequest {
	return CreateClusterQuerySamplingApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		querySamplingCreateRequest: querySamplingCreateRequest,
	}
}

// CreateClusterQuerySamplingExecute executes the request
//
//	@return QuerySamplingResponse
func (a *QuerySamplingAPIService) CreateClusterQuerySamplingExecute(r CreateClusterQuerySamplingApiRequest) (*QuerySamplingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *QuerySamplingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuerySamplingAPIService.CreateClusterQuerySampling")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling"
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
	if r.querySamplingCreateRequest == nil {
		return localVarReturnValue, nil, reportError("querySamplingCreateRequest is required and must be specified")
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
	localVarPostBody = r.querySamplingCreateRequest
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

type DeleteClusterQuerySamplingApiRequest struct {
	ctx         context.Context
	ApiService  QuerySamplingAPI
	groupId     string
	clusterName string
	namespace   string
}

type DeleteClusterQuerySamplingApiParams struct {
	GroupId     string
	ClusterName string
	Namespace   string
}

func (a *QuerySamplingAPIService) DeleteClusterQuerySamplingWithParams(ctx context.Context, args *DeleteClusterQuerySamplingApiParams) DeleteClusterQuerySamplingApiRequest {
	return DeleteClusterQuerySamplingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		namespace:   args.Namespace,
	}
}

func (r DeleteClusterQuerySamplingApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClusterQuerySamplingExecute(r)
}

/*
DeleteClusterQuerySampling Remove Query Sampling from One Namespace

Stops query sampling on one namespace of the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param namespace Human-readable label that identifies the namespace (`<database>.<collection>`) on which queries are sampled.
	@return DeleteClusterQuerySamplingApiRequest
*/
func (a *QuerySamplingAPIService) DeleteClusterQuerySampling(ctx context.Context, groupId string, clusterName string, namespace string) DeleteClusterQuerySamplingApiRequest {
	return DeleteClusterQuerySamplingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		namespace:   namespace,
	}
}

// DeleteClusterQuerySamplingExecute executes the request
func (a *QuerySamplingAPIService) DeleteClusterQuerySamplingExecute(r DeleteClusterQuerySamplingApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuerySamplingAPIService.DeleteClusterQuerySampling")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling/{namespace}"
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
	if r.namespace == "" {
		return nil, reportError("namespace is empty and must be specified")
	}
	if r.namespace == "." || r.namespace == ".." {
		return nil, reportError("namespace must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(r.namespace), -1)

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

type GetClusterQuerySamplingApiRequest struct {
	ctx         context.Context
	ApiService  QuerySamplingAPI
	groupId     string
	clusterName string
	namespace   string
}

type GetClusterQuerySamplingApiParams struct {
	GroupId     string
	ClusterName string
	Namespace   string
}

func (a *QuerySamplingAPIService) GetClusterQuerySamplingWithParams(ctx context.Context, args *GetClusterQuerySamplingApiParams) GetClusterQuerySamplingApiRequest {
	return GetClusterQuerySamplingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		namespace:   args.Namespace,
	}
}

func (r GetClusterQuerySamplingApiRequest) Execute() (*QuerySamplingResponse, *http.Response, error) {
	return r.ApiService.GetClusterQuerySamplingExecute(r)
}

/*
GetClusterQuerySampling Return Query Sampling State for One Namespace

Returns the query sampling state of one namespace of the specified cluster, including counters for the operations sampled so far. Returns 404 if query sampling is not configured on the namespace. `active` becomes true asynchronously after query sampling starts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param namespace Human-readable label that identifies the namespace (`<database>.<collection>`) on which queries are sampled.
	@return GetClusterQuerySamplingApiRequest
*/
func (a *QuerySamplingAPIService) GetClusterQuerySampling(ctx context.Context, groupId string, clusterName string, namespace string) GetClusterQuerySamplingApiRequest {
	return GetClusterQuerySamplingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		namespace:   namespace,
	}
}

// GetClusterQuerySamplingExecute executes the request
//
//	@return QuerySamplingResponse
func (a *QuerySamplingAPIService) GetClusterQuerySamplingExecute(r GetClusterQuerySamplingApiRequest) (*QuerySamplingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *QuerySamplingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuerySamplingAPIService.GetClusterQuerySampling")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling/{namespace}"
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
	if r.namespace == "" {
		return localVarReturnValue, nil, reportError("namespace is empty and must be specified")
	}
	if r.namespace == "." || r.namespace == ".." {
		return localVarReturnValue, nil, reportError("namespace must not be a dot-segment path parameter")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(r.namespace), -1)

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
