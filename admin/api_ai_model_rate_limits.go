// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AIModelRateLimitsApi interface {

	/*
		GetGroupModelLimit Return Single AI Model Rate Limit for One Group

		Retrieve a single AI model rate limit for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param modelGroupName The name of the model group to be retrieved.
		@return GetGroupModelLimitApiRequest
	*/
	GetGroupModelLimit(ctx context.Context, groupId string, modelGroupName string) GetGroupModelLimitApiRequest
	/*
		GetGroupModelLimit Return Single AI Model Rate Limit for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupModelLimitApiParams - Parameters for the request
		@return GetGroupModelLimitApiRequest
	*/
	GetGroupModelLimitWithParams(ctx context.Context, args *GetGroupModelLimitApiParams) GetGroupModelLimitApiRequest

	// Method available only for mocking purposes
	GetGroupModelLimitExecute(r GetGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error)

	/*
		GetOrgModelLimit Return Single AI Model Rate Limit for One Organization

		Retrieve a single AI model rate limit for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param modelGroupName The name of the model group to be retrieved.
		@return GetOrgModelLimitApiRequest
	*/
	GetOrgModelLimit(ctx context.Context, orgId string, modelGroupName string) GetOrgModelLimitApiRequest
	/*
		GetOrgModelLimit Return Single AI Model Rate Limit for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgModelLimitApiParams - Parameters for the request
		@return GetOrgModelLimitApiRequest
	*/
	GetOrgModelLimitWithParams(ctx context.Context, args *GetOrgModelLimitApiParams) GetOrgModelLimitApiRequest

	// Method available only for mocking purposes
	GetOrgModelLimitExecute(r GetOrgModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error)

	/*
		ListGroupModelLimits Return AI Model Rate Limits for One Group

		Retrieve AI model rate limits for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListGroupModelLimitsApiRequest
	*/
	ListGroupModelLimits(ctx context.Context, groupId string) ListGroupModelLimitsApiRequest
	/*
		ListGroupModelLimits Return AI Model Rate Limits for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupModelLimitsApiParams - Parameters for the request
		@return ListGroupModelLimitsApiRequest
	*/
	ListGroupModelLimitsWithParams(ctx context.Context, args *ListGroupModelLimitsApiParams) ListGroupModelLimitsApiRequest

	// Method available only for mocking purposes
	ListGroupModelLimitsExecute(r ListGroupModelLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error)

	/*
		ListOrgModelLimits Return AI Model Rate Limits for One Organization

		Retrieve AI model rate limits for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrgModelLimitsApiRequest
	*/
	ListOrgModelLimits(ctx context.Context, orgId string) ListOrgModelLimitsApiRequest
	/*
		ListOrgModelLimits Return AI Model Rate Limits for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgModelLimitsApiParams - Parameters for the request
		@return ListOrgModelLimitsApiRequest
	*/
	ListOrgModelLimitsWithParams(ctx context.Context, args *ListOrgModelLimitsApiParams) ListOrgModelLimitsApiRequest

	// Method available only for mocking purposes
	ListOrgModelLimitsExecute(r ListOrgModelLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error)

	/*
		ResetModelRateLimit Reset AI Model Rate Limit for One Model Group

		Reset the AI model rate limit for the given model group to default values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param modelGroupName The name of the model group to be reset to default rate limits.
		@return ResetModelRateLimitApiRequest
	*/
	ResetModelRateLimit(ctx context.Context, groupId string, modelGroupName string) ResetModelRateLimitApiRequest
	/*
		ResetModelRateLimit Reset AI Model Rate Limit for One Model Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetModelRateLimitApiParams - Parameters for the request
		@return ResetModelRateLimitApiRequest
	*/
	ResetModelRateLimitWithParams(ctx context.Context, args *ResetModelRateLimitApiParams) ResetModelRateLimitApiRequest

	// Method available only for mocking purposes
	ResetModelRateLimitExecute(r ResetModelRateLimitApiRequest) (*http.Response, error)

	/*
		ResetModelRateLimits Reset AI Model Rate Limits for Group

		Reset the AI Model rate limits for the given group to default values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ResetModelRateLimitsApiRequest
	*/
	ResetModelRateLimits(ctx context.Context, groupId string) ResetModelRateLimitsApiRequest
	/*
		ResetModelRateLimits Reset AI Model Rate Limits for Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetModelRateLimitsApiParams - Parameters for the request
		@return ResetModelRateLimitsApiRequest
	*/
	ResetModelRateLimitsWithParams(ctx context.Context, args *ResetModelRateLimitsApiParams) ResetModelRateLimitsApiRequest

	// Method available only for mocking purposes
	ResetModelRateLimitsExecute(r ResetModelRateLimitsApiRequest) (*http.Response, error)

	/*
		UpdateModelRateLimit Update AI Model Rate Limit

		Update an AI model rate limit for the given model group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param modelGroupName The name of the model group to be updated.
		@param aiModelsRateLimitUpdateRequest A request containing the new rate limits for the given model group.
		@return UpdateModelRateLimitApiRequest
	*/
	UpdateModelRateLimit(ctx context.Context, groupId string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateModelRateLimitApiRequest
	/*
		UpdateModelRateLimit Update AI Model Rate Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateModelRateLimitApiParams - Parameters for the request
		@return UpdateModelRateLimitApiRequest
	*/
	UpdateModelRateLimitWithParams(ctx context.Context, args *UpdateModelRateLimitApiParams) UpdateModelRateLimitApiRequest

	// Method available only for mocking purposes
	UpdateModelRateLimitExecute(r UpdateModelRateLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error)
}

// AIModelRateLimitsApiService AIModelRateLimitsApi service
type AIModelRateLimitsApiService service

type GetGroupModelLimitApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	groupId        string
	modelGroupName string
}

type GetGroupModelLimitApiParams struct {
	GroupId        string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) GetGroupModelLimitWithParams(ctx context.Context, args *GetGroupModelLimitApiParams) GetGroupModelLimitApiRequest {
	return GetGroupModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		modelGroupName: args.ModelGroupName,
	}
}

func (r GetGroupModelLimitApiRequest) Execute() (*AiModelsRateLimitResponse, *http.Response, error) {
	return r.ApiService.GetGroupModelLimitExecute(r)
}

/*
GetGroupModelLimit Return Single AI Model Rate Limit for One Group

Retrieve a single AI model rate limit for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param modelGroupName The name of the model group to be retrieved.
	@return GetGroupModelLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) GetGroupModelLimit(ctx context.Context, groupId string, modelGroupName string) GetGroupModelLimitApiRequest {
	return GetGroupModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		modelGroupName: modelGroupName,
	}
}

// GetGroupModelLimitExecute executes the request
//
//	@return AiModelsRateLimitResponse
func (a *AIModelRateLimitsApiService) GetGroupModelLimitExecute(r GetGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetGroupModelLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelRateLimits/{modelGroupName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.modelGroupName == "" {
		return localVarReturnValue, nil, reportError("modelGroupName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"modelGroupName"+"}", url.PathEscape(r.modelGroupName), -1)

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

type GetOrgModelLimitApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	orgId          string
	modelGroupName string
}

type GetOrgModelLimitApiParams struct {
	OrgId          string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) GetOrgModelLimitWithParams(ctx context.Context, args *GetOrgModelLimitApiParams) GetOrgModelLimitApiRequest {
	return GetOrgModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          args.OrgId,
		modelGroupName: args.ModelGroupName,
	}
}

func (r GetOrgModelLimitApiRequest) Execute() (*AiModelsRateLimitResponse, *http.Response, error) {
	return r.ApiService.GetOrgModelLimitExecute(r)
}

/*
GetOrgModelLimit Return Single AI Model Rate Limit for One Organization

Retrieve a single AI model rate limit for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param modelGroupName The name of the model group to be retrieved.
	@return GetOrgModelLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) GetOrgModelLimit(ctx context.Context, orgId string, modelGroupName string) GetOrgModelLimitApiRequest {
	return GetOrgModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          orgId,
		modelGroupName: modelGroupName,
	}
}

// GetOrgModelLimitExecute executes the request
//
//	@return AiModelsRateLimitResponse
func (a *AIModelRateLimitsApiService) GetOrgModelLimitExecute(r GetOrgModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetOrgModelLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelRateLimits/{modelGroupName}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.modelGroupName == "" {
		return localVarReturnValue, nil, reportError("modelGroupName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"modelGroupName"+"}", url.PathEscape(r.modelGroupName), -1)

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

type ListGroupModelLimitsApiRequest struct {
	ctx          context.Context
	ApiService   AIModelRateLimitsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListGroupModelLimitsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelRateLimitsApiService) ListGroupModelLimitsWithParams(ctx context.Context, args *ListGroupModelLimitsApiParams) ListGroupModelLimitsApiRequest {
	return ListGroupModelLimitsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupModelLimitsApiRequest) ItemsPerPage(itemsPerPage int) ListGroupModelLimitsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupModelLimitsApiRequest) PageNum(pageNum int) ListGroupModelLimitsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupModelLimitsApiRequest) Execute() (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	return r.ApiService.ListGroupModelLimitsExecute(r)
}

/*
ListGroupModelLimits Return AI Model Rate Limits for One Group

Retrieve AI model rate limits for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListGroupModelLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ListGroupModelLimits(ctx context.Context, groupId string) ListGroupModelLimitsApiRequest {
	return ListGroupModelLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupModelLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelsRateLimitsResponse
func (a *AIModelRateLimitsApiService) ListGroupModelLimitsExecute(r ListGroupModelLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ListGroupModelLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelRateLimits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type ListOrgModelLimitsApiRequest struct {
	ctx          context.Context
	ApiService   AIModelRateLimitsApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type ListOrgModelLimitsApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelRateLimitsApiService) ListOrgModelLimitsWithParams(ctx context.Context, args *ListOrgModelLimitsApiParams) ListOrgModelLimitsApiRequest {
	return ListOrgModelLimitsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgModelLimitsApiRequest) ItemsPerPage(itemsPerPage int) ListOrgModelLimitsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgModelLimitsApiRequest) PageNum(pageNum int) ListOrgModelLimitsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgModelLimitsApiRequest) Execute() (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	return r.ApiService.ListOrgModelLimitsExecute(r)
}

/*
ListOrgModelLimits Return AI Model Rate Limits for One Organization

Retrieve AI model rate limits for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrgModelLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ListOrgModelLimits(ctx context.Context, orgId string) ListOrgModelLimitsApiRequest {
	return ListOrgModelLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrgModelLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelsRateLimitsResponse
func (a *AIModelRateLimitsApiService) ListOrgModelLimitsExecute(r ListOrgModelLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ListOrgModelLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelRateLimits"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

type ResetModelRateLimitApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	groupId        string
	modelGroupName string
}

type ResetModelRateLimitApiParams struct {
	GroupId        string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) ResetModelRateLimitWithParams(ctx context.Context, args *ResetModelRateLimitApiParams) ResetModelRateLimitApiRequest {
	return ResetModelRateLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		modelGroupName: args.ModelGroupName,
	}
}

func (r ResetModelRateLimitApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResetModelRateLimitExecute(r)
}

/*
ResetModelRateLimit Reset AI Model Rate Limit for One Model Group

Reset the AI model rate limit for the given model group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param modelGroupName The name of the model group to be reset to default rate limits.
	@return ResetModelRateLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetModelRateLimit(ctx context.Context, groupId string, modelGroupName string) ResetModelRateLimitApiRequest {
	return ResetModelRateLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		modelGroupName: modelGroupName,
	}
}

// ResetModelRateLimitExecute executes the request
func (a *AIModelRateLimitsApiService) ResetModelRateLimitExecute(r ResetModelRateLimitApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetModelRateLimit")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelRateLimits/{modelGroupName}:reset"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.modelGroupName == "" {
		return nil, reportError("modelGroupName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"modelGroupName"+"}", url.PathEscape(r.modelGroupName), -1)

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

type ResetModelRateLimitsApiRequest struct {
	ctx        context.Context
	ApiService AIModelRateLimitsApi
	groupId    string
}

type ResetModelRateLimitsApiParams struct {
	GroupId string
}

func (a *AIModelRateLimitsApiService) ResetModelRateLimitsWithParams(ctx context.Context, args *ResetModelRateLimitsApiParams) ResetModelRateLimitsApiRequest {
	return ResetModelRateLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ResetModelRateLimitsApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResetModelRateLimitsExecute(r)
}

/*
ResetModelRateLimits Reset AI Model Rate Limits for Group

Reset the AI Model rate limits for the given group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ResetModelRateLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetModelRateLimits(ctx context.Context, groupId string) ResetModelRateLimitsApiRequest {
	return ResetModelRateLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ResetModelRateLimitsExecute executes the request
func (a *AIModelRateLimitsApiService) ResetModelRateLimitsExecute(r ResetModelRateLimitsApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetModelRateLimits")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelRateLimits:reset"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type UpdateModelRateLimitApiRequest struct {
	ctx                            context.Context
	ApiService                     AIModelRateLimitsApi
	groupId                        string
	modelGroupName                 string
	aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

type UpdateModelRateLimitApiParams struct {
	GroupId                        string
	ModelGroupName                 string
	AiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

func (a *AIModelRateLimitsApiService) UpdateModelRateLimitWithParams(ctx context.Context, args *UpdateModelRateLimitApiParams) UpdateModelRateLimitApiRequest {
	return UpdateModelRateLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		modelGroupName:                 args.ModelGroupName,
		aiModelsRateLimitUpdateRequest: args.AiModelsRateLimitUpdateRequest,
	}
}

func (r UpdateModelRateLimitApiRequest) Execute() (*AiModelsRateLimitResponse, *http.Response, error) {
	return r.ApiService.UpdateModelRateLimitExecute(r)
}

/*
UpdateModelRateLimit Update AI Model Rate Limit

Update an AI model rate limit for the given model group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param modelGroupName The name of the model group to be updated.
	@return UpdateModelRateLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) UpdateModelRateLimit(ctx context.Context, groupId string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateModelRateLimitApiRequest {
	return UpdateModelRateLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		modelGroupName:                 modelGroupName,
		aiModelsRateLimitUpdateRequest: aiModelsRateLimitUpdateRequest,
	}
}

// UpdateModelRateLimitExecute executes the request
//
//	@return AiModelsRateLimitResponse
func (a *AIModelRateLimitsApiService) UpdateModelRateLimitExecute(r UpdateModelRateLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.UpdateModelRateLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelRateLimits/{modelGroupName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.modelGroupName == "" {
		return localVarReturnValue, nil, reportError("modelGroupName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"modelGroupName"+"}", url.PathEscape(r.modelGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiModelsRateLimitUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("aiModelsRateLimitUpdateRequest is required and must be specified")
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
	localVarPostBody = r.aiModelsRateLimitUpdateRequest
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
