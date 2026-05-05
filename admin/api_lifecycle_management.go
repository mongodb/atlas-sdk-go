// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type LifecycleManagementApi interface {

	/*
		CreateLifecycleManagementPolicy Create One Lifecycle Management Policy

		Creates a lifecycle management policy for the specified cluster. The policy defines rules for taking lifecycle management actions against a cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param clusterName Human-readable label that identifies the cluster.
		@param apiLifecycleManagementPolicyCreateRequest Lifecycle management policy to create.
		@return CreateLifecycleManagementPolicyApiRequest
	*/
	CreateLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, apiLifecycleManagementPolicyCreateRequest *ApiLifecycleManagementPolicyCreateRequest) CreateLifecycleManagementPolicyApiRequest
	/*
		CreateLifecycleManagementPolicy Create One Lifecycle Management Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateLifecycleManagementPolicyApiParams - Parameters for the request
		@return CreateLifecycleManagementPolicyApiRequest
	*/
	CreateLifecycleManagementPolicyWithParams(ctx context.Context, args *CreateLifecycleManagementPolicyApiParams) CreateLifecycleManagementPolicyApiRequest

	// Method available only for mocking purposes
	CreateLifecycleManagementPolicyExecute(r CreateLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error)

	/*
		DeleteLifecycleManagementPolicy Delete One Lifecycle Management Policy

		Deletes the lifecycle management policy with the specified identifier. Policies can be deleted from any operational state.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param clusterName Human-readable label that identifies the cluster.
		@param policyId Unique 24-hexadecimal digit string that identifies the policy.
		@return DeleteLifecycleManagementPolicyApiRequest
	*/
	DeleteLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string) DeleteLifecycleManagementPolicyApiRequest
	/*
		DeleteLifecycleManagementPolicy Delete One Lifecycle Management Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLifecycleManagementPolicyApiParams - Parameters for the request
		@return DeleteLifecycleManagementPolicyApiRequest
	*/
	DeleteLifecycleManagementPolicyWithParams(ctx context.Context, args *DeleteLifecycleManagementPolicyApiParams) DeleteLifecycleManagementPolicyApiRequest

	// Method available only for mocking purposes
	DeleteLifecycleManagementPolicyExecute(r DeleteLifecycleManagementPolicyApiRequest) (*http.Response, error)

	/*
		GetLifecycleManagementPolicy Return One Lifecycle Management Policy

		Returns the lifecycle management policy with the specified identifier.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param clusterName Human-readable label that identifies the cluster.
		@param policyId Unique 24-hexadecimal digit string that identifies the policy.
		@return GetLifecycleManagementPolicyApiRequest
	*/
	GetLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string) GetLifecycleManagementPolicyApiRequest
	/*
		GetLifecycleManagementPolicy Return One Lifecycle Management Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLifecycleManagementPolicyApiParams - Parameters for the request
		@return GetLifecycleManagementPolicyApiRequest
	*/
	GetLifecycleManagementPolicyWithParams(ctx context.Context, args *GetLifecycleManagementPolicyApiParams) GetLifecycleManagementPolicyApiRequest

	// Method available only for mocking purposes
	GetLifecycleManagementPolicyExecute(r GetLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error)

	/*
		ListLifecycleManagementPolicies Return All Lifecycle Management Policies for One Cluster

		Returns all lifecycle management policies for the specified cluster. Each policy represents automated data transition rules between collections.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListLifecycleManagementPoliciesApiRequest
	*/
	ListLifecycleManagementPolicies(ctx context.Context, groupId string, clusterName string) ListLifecycleManagementPoliciesApiRequest
	/*
		ListLifecycleManagementPolicies Return All Lifecycle Management Policies for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListLifecycleManagementPoliciesApiParams - Parameters for the request
		@return ListLifecycleManagementPoliciesApiRequest
	*/
	ListLifecycleManagementPoliciesWithParams(ctx context.Context, args *ListLifecycleManagementPoliciesApiParams) ListLifecycleManagementPoliciesApiRequest

	// Method available only for mocking purposes
	ListLifecycleManagementPoliciesExecute(r ListLifecycleManagementPoliciesApiRequest) (*PaginatedApiLifecycleManagementPolicyResponse, *http.Response, error)

	/*
		UpdateLifecycleManagementPolicy Update One Lifecycle Management Policy

		Updates the lifecycle management policy with the specified identifier. Supports partial updates via PATCH semantics. Can update action fields, transition policy state, or both.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param clusterName Human-readable label that identifies the cluster.
		@param policyId Unique 24-hexadecimal digit string that identifies the policy.
		@param apiLifecycleManagementPolicyUpdateRequest Lifecycle management policy fields to update.
		@return UpdateLifecycleManagementPolicyApiRequest
	*/
	UpdateLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string, apiLifecycleManagementPolicyUpdateRequest *ApiLifecycleManagementPolicyUpdateRequest) UpdateLifecycleManagementPolicyApiRequest
	/*
		UpdateLifecycleManagementPolicy Update One Lifecycle Management Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateLifecycleManagementPolicyApiParams - Parameters for the request
		@return UpdateLifecycleManagementPolicyApiRequest
	*/
	UpdateLifecycleManagementPolicyWithParams(ctx context.Context, args *UpdateLifecycleManagementPolicyApiParams) UpdateLifecycleManagementPolicyApiRequest

	// Method available only for mocking purposes
	UpdateLifecycleManagementPolicyExecute(r UpdateLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error)
}

// LifecycleManagementApiService LifecycleManagementApi service
type LifecycleManagementApiService service

type CreateLifecycleManagementPolicyApiRequest struct {
	ctx                                       context.Context
	ApiService                                LifecycleManagementApi
	groupId                                   string
	clusterName                               string
	apiLifecycleManagementPolicyCreateRequest *ApiLifecycleManagementPolicyCreateRequest
}

type CreateLifecycleManagementPolicyApiParams struct {
	GroupId                                   string
	ClusterName                               string
	ApiLifecycleManagementPolicyCreateRequest *ApiLifecycleManagementPolicyCreateRequest
}

func (a *LifecycleManagementApiService) CreateLifecycleManagementPolicyWithParams(ctx context.Context, args *CreateLifecycleManagementPolicyApiParams) CreateLifecycleManagementPolicyApiRequest {
	return CreateLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		apiLifecycleManagementPolicyCreateRequest: args.ApiLifecycleManagementPolicyCreateRequest,
	}
}

func (r CreateLifecycleManagementPolicyApiRequest) Execute() (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	return r.ApiService.CreateLifecycleManagementPolicyExecute(r)
}

/*
CreateLifecycleManagementPolicy Create One Lifecycle Management Policy

Creates a lifecycle management policy for the specified cluster. The policy defines rules for taking lifecycle management actions against a cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param clusterName Human-readable label that identifies the cluster.
	@return CreateLifecycleManagementPolicyApiRequest
*/
func (a *LifecycleManagementApiService) CreateLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, apiLifecycleManagementPolicyCreateRequest *ApiLifecycleManagementPolicyCreateRequest) CreateLifecycleManagementPolicyApiRequest {
	return CreateLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		apiLifecycleManagementPolicyCreateRequest: apiLifecycleManagementPolicyCreateRequest,
	}
}

// CreateLifecycleManagementPolicyExecute executes the request
//
//	@return ApiLifecycleManagementPolicyResponse
func (a *LifecycleManagementApiService) CreateLifecycleManagementPolicyExecute(r CreateLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiLifecycleManagementPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleManagementApiService.CreateLifecycleManagementPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/lifecycleManagementPolicies"
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
	if r.apiLifecycleManagementPolicyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("apiLifecycleManagementPolicyCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiLifecycleManagementPolicyCreateRequest
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

type DeleteLifecycleManagementPolicyApiRequest struct {
	ctx         context.Context
	ApiService  LifecycleManagementApi
	groupId     string
	clusterName string
	policyId    string
}

type DeleteLifecycleManagementPolicyApiParams struct {
	GroupId     string
	ClusterName string
	PolicyId    string
}

func (a *LifecycleManagementApiService) DeleteLifecycleManagementPolicyWithParams(ctx context.Context, args *DeleteLifecycleManagementPolicyApiParams) DeleteLifecycleManagementPolicyApiRequest {
	return DeleteLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		policyId:    args.PolicyId,
	}
}

func (r DeleteLifecycleManagementPolicyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLifecycleManagementPolicyExecute(r)
}

/*
DeleteLifecycleManagementPolicy Delete One Lifecycle Management Policy

Deletes the lifecycle management policy with the specified identifier. Policies can be deleted from any operational state.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param clusterName Human-readable label that identifies the cluster.
	@param policyId Unique 24-hexadecimal digit string that identifies the policy.
	@return DeleteLifecycleManagementPolicyApiRequest
*/
func (a *LifecycleManagementApiService) DeleteLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string) DeleteLifecycleManagementPolicyApiRequest {
	return DeleteLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		policyId:    policyId,
	}
}

// DeleteLifecycleManagementPolicyExecute executes the request
func (a *LifecycleManagementApiService) DeleteLifecycleManagementPolicyExecute(r DeleteLifecycleManagementPolicyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleManagementApiService.DeleteLifecycleManagementPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/lifecycleManagementPolicies/{policyId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.policyId == "" {
		return nil, reportError("policyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(r.policyId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

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

type GetLifecycleManagementPolicyApiRequest struct {
	ctx         context.Context
	ApiService  LifecycleManagementApi
	groupId     string
	clusterName string
	policyId    string
}

type GetLifecycleManagementPolicyApiParams struct {
	GroupId     string
	ClusterName string
	PolicyId    string
}

func (a *LifecycleManagementApiService) GetLifecycleManagementPolicyWithParams(ctx context.Context, args *GetLifecycleManagementPolicyApiParams) GetLifecycleManagementPolicyApiRequest {
	return GetLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		policyId:    args.PolicyId,
	}
}

func (r GetLifecycleManagementPolicyApiRequest) Execute() (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	return r.ApiService.GetLifecycleManagementPolicyExecute(r)
}

/*
GetLifecycleManagementPolicy Return One Lifecycle Management Policy

Returns the lifecycle management policy with the specified identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param clusterName Human-readable label that identifies the cluster.
	@param policyId Unique 24-hexadecimal digit string that identifies the policy.
	@return GetLifecycleManagementPolicyApiRequest
*/
func (a *LifecycleManagementApiService) GetLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string) GetLifecycleManagementPolicyApiRequest {
	return GetLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		policyId:    policyId,
	}
}

// GetLifecycleManagementPolicyExecute executes the request
//
//	@return ApiLifecycleManagementPolicyResponse
func (a *LifecycleManagementApiService) GetLifecycleManagementPolicyExecute(r GetLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiLifecycleManagementPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleManagementApiService.GetLifecycleManagementPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/lifecycleManagementPolicies/{policyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.policyId == "" {
		return localVarReturnValue, nil, reportError("policyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(r.policyId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

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

type ListLifecycleManagementPoliciesApiRequest struct {
	ctx          context.Context
	ApiService   LifecycleManagementApi
	groupId      string
	clusterName  string
	itemsPerPage *int
	pageNum      *int
}

type ListLifecycleManagementPoliciesApiParams struct {
	GroupId      string
	ClusterName  string
	ItemsPerPage *int
	PageNum      *int
}

func (a *LifecycleManagementApiService) ListLifecycleManagementPoliciesWithParams(ctx context.Context, args *ListLifecycleManagementPoliciesApiParams) ListLifecycleManagementPoliciesApiRequest {
	return ListLifecycleManagementPoliciesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListLifecycleManagementPoliciesApiRequest) ItemsPerPage(itemsPerPage int) ListLifecycleManagementPoliciesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListLifecycleManagementPoliciesApiRequest) PageNum(pageNum int) ListLifecycleManagementPoliciesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListLifecycleManagementPoliciesApiRequest) Execute() (*PaginatedApiLifecycleManagementPolicyResponse, *http.Response, error) {
	return r.ApiService.ListLifecycleManagementPoliciesExecute(r)
}

/*
ListLifecycleManagementPolicies Return All Lifecycle Management Policies for One Cluster

Returns all lifecycle management policies for the specified cluster. Each policy represents automated data transition rules between collections.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListLifecycleManagementPoliciesApiRequest
*/
func (a *LifecycleManagementApiService) ListLifecycleManagementPolicies(ctx context.Context, groupId string, clusterName string) ListLifecycleManagementPoliciesApiRequest {
	return ListLifecycleManagementPoliciesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListLifecycleManagementPoliciesExecute executes the request
//
//	@return PaginatedApiLifecycleManagementPolicyResponse
func (a *LifecycleManagementApiService) ListLifecycleManagementPoliciesExecute(r ListLifecycleManagementPoliciesApiRequest) (*PaginatedApiLifecycleManagementPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiLifecycleManagementPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleManagementApiService.ListLifecycleManagementPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/lifecycleManagementPolicies"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

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

type UpdateLifecycleManagementPolicyApiRequest struct {
	ctx                                       context.Context
	ApiService                                LifecycleManagementApi
	groupId                                   string
	clusterName                               string
	policyId                                  string
	apiLifecycleManagementPolicyUpdateRequest *ApiLifecycleManagementPolicyUpdateRequest
}

type UpdateLifecycleManagementPolicyApiParams struct {
	GroupId                                   string
	ClusterName                               string
	PolicyId                                  string
	ApiLifecycleManagementPolicyUpdateRequest *ApiLifecycleManagementPolicyUpdateRequest
}

func (a *LifecycleManagementApiService) UpdateLifecycleManagementPolicyWithParams(ctx context.Context, args *UpdateLifecycleManagementPolicyApiParams) UpdateLifecycleManagementPolicyApiRequest {
	return UpdateLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		policyId:    args.PolicyId,
		apiLifecycleManagementPolicyUpdateRequest: args.ApiLifecycleManagementPolicyUpdateRequest,
	}
}

func (r UpdateLifecycleManagementPolicyApiRequest) Execute() (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateLifecycleManagementPolicyExecute(r)
}

/*
UpdateLifecycleManagementPolicy Update One Lifecycle Management Policy

Updates the lifecycle management policy with the specified identifier. Supports partial updates via PATCH semantics. Can update action fields, transition policy state, or both.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param clusterName Human-readable label that identifies the cluster.
	@param policyId Unique 24-hexadecimal digit string that identifies the policy.
	@return UpdateLifecycleManagementPolicyApiRequest
*/
func (a *LifecycleManagementApiService) UpdateLifecycleManagementPolicy(ctx context.Context, groupId string, clusterName string, policyId string, apiLifecycleManagementPolicyUpdateRequest *ApiLifecycleManagementPolicyUpdateRequest) UpdateLifecycleManagementPolicyApiRequest {
	return UpdateLifecycleManagementPolicyApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		policyId:    policyId,
		apiLifecycleManagementPolicyUpdateRequest: apiLifecycleManagementPolicyUpdateRequest,
	}
}

// UpdateLifecycleManagementPolicyExecute executes the request
//
//	@return ApiLifecycleManagementPolicyResponse
func (a *LifecycleManagementApiService) UpdateLifecycleManagementPolicyExecute(r UpdateLifecycleManagementPolicyApiRequest) (*ApiLifecycleManagementPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiLifecycleManagementPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LifecycleManagementApiService.UpdateLifecycleManagementPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/lifecycleManagementPolicies/{policyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.policyId == "" {
		return localVarReturnValue, nil, reportError("policyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(r.policyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiLifecycleManagementPolicyUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("apiLifecycleManagementPolicyUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiLifecycleManagementPolicyUpdateRequest
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
