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

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve a single AI model rate limit for the given group.

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
		GetGroupModelLimits Return AI Model Rate Limits for One Group

		Retrieve AI model rate limits for the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetGroupModelLimitsApiRequest
	*/
	GetGroupModelLimits(ctx context.Context, groupId string) GetGroupModelLimitsApiRequest
	/*
		GetGroupModelLimits Return AI Model Rate Limits for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupModelLimitsApiParams - Parameters for the request
		@return GetGroupModelLimitsApiRequest
	*/
	GetGroupModelLimitsWithParams(ctx context.Context, args *GetGroupModelLimitsApiParams) GetGroupModelLimitsApiRequest

	// Method available only for mocking purposes
	GetGroupModelLimitsExecute(r GetGroupModelLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error)

	/*
		GetGroupRateLimits Return Single AI Model Rate Limit for One Group

		Retrieve a single scoped AI model rate limit for the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param modelGroupName The name of the model group to be retrieved.
		@return GetGroupRateLimitsApiRequest
	*/
	GetGroupRateLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string) GetGroupRateLimitsApiRequest
	/*
		GetGroupRateLimits Return Single AI Model Rate Limit for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupRateLimitsApiParams - Parameters for the request
		@return GetGroupRateLimitsApiRequest
	*/
	GetGroupRateLimitsWithParams(ctx context.Context, args *GetGroupRateLimitsApiParams) GetGroupRateLimitsApiRequest

	// Method available only for mocking purposes
	GetGroupRateLimitsExecute(r GetGroupRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error)

	/*
			GetOrgModelLimit Return Single AI Model Rate Limit for One Organization

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

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
		GetOrgModelLimits Return AI Model Rate Limits for One Organization

		Retrieve AI model rate limits for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetOrgModelLimitsApiRequest
	*/
	GetOrgModelLimits(ctx context.Context, orgId string) GetOrgModelLimitsApiRequest
	/*
		GetOrgModelLimits Return AI Model Rate Limits for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgModelLimitsApiParams - Parameters for the request
		@return GetOrgModelLimitsApiRequest
	*/
	GetOrgModelLimitsWithParams(ctx context.Context, args *GetOrgModelLimitsApiParams) GetOrgModelLimitsApiRequest

	// Method available only for mocking purposes
	GetOrgModelLimitsExecute(r GetOrgModelLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error)

	/*
		GetOrgRateLimits Return Single AI Model Rate Limit for One Organization

		Retrieve a single scoped AI model rate limit for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param modelGroupName The name of the model group to be retrieved.
		@return GetOrgRateLimitsApiRequest
	*/
	GetOrgRateLimits(ctx context.Context, orgId string, cloud string, geography string, modelGroupName string) GetOrgRateLimitsApiRequest
	/*
		GetOrgRateLimits Return Single AI Model Rate Limit for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgRateLimitsApiParams - Parameters for the request
		@return GetOrgRateLimitsApiRequest
	*/
	GetOrgRateLimitsWithParams(ctx context.Context, args *GetOrgRateLimitsApiParams) GetOrgRateLimitsApiRequest

	// Method available only for mocking purposes
	GetOrgRateLimitsExecute(r GetOrgRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error)

	/*
			ListGroupModelLimits Return AI Model Rate Limits for One Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve AI model rate limits for the given group.

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

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

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
			ResetGroupApiLimits Reset AI Model Rate Limits for Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Reset the AI Model rate limits for the given group to default values.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ResetGroupApiLimitsApiRequest
	*/
	ResetGroupApiLimits(ctx context.Context, groupId string) ResetGroupApiLimitsApiRequest
	/*
		ResetGroupApiLimits Reset AI Model Rate Limits for Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetGroupApiLimitsApiParams - Parameters for the request
		@return ResetGroupApiLimitsApiRequest
	*/
	ResetGroupApiLimitsWithParams(ctx context.Context, args *ResetGroupApiLimitsApiParams) ResetGroupApiLimitsApiRequest

	// Method available only for mocking purposes
	ResetGroupApiLimitsExecute(r ResetGroupApiLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error)

	/*
			ResetGroupModelLimit Reset AI Model Rate Limit for One Model Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Reset the AI model rate limit for the given model group to default values.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param modelGroupName The name of the model group to be reset to default rate limits.
			@return ResetGroupModelLimitApiRequest
	*/
	ResetGroupModelLimit(ctx context.Context, groupId string, modelGroupName string) ResetGroupModelLimitApiRequest
	/*
		ResetGroupModelLimit Reset AI Model Rate Limit for One Model Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetGroupModelLimitApiParams - Parameters for the request
		@return ResetGroupModelLimitApiRequest
	*/
	ResetGroupModelLimitWithParams(ctx context.Context, args *ResetGroupModelLimitApiParams) ResetGroupModelLimitApiRequest

	// Method available only for mocking purposes
	ResetGroupModelLimitExecute(r ResetGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error)

	/*
		ResetGroupModelLimits Reset AI Model Rate Limit for One Model Group

		Reset the scoped AI model rate limit for the given model group to default values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param modelGroupName The name of the model group to be reset to default rate limits.
		@return ResetGroupModelLimitsApiRequest
	*/
	ResetGroupModelLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string) ResetGroupModelLimitsApiRequest
	/*
		ResetGroupModelLimits Reset AI Model Rate Limit for One Model Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetGroupModelLimitsApiParams - Parameters for the request
		@return ResetGroupModelLimitsApiRequest
	*/
	ResetGroupModelLimitsWithParams(ctx context.Context, args *ResetGroupModelLimitsApiParams) ResetGroupModelLimitsApiRequest

	// Method available only for mocking purposes
	ResetGroupModelLimitsExecute(r ResetGroupModelLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error)

	/*
		ResetGroupRateLimits Reset AI Model Rate Limits for Group

		Reset the AI Model rate limits for the given group to default values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ResetGroupRateLimitsApiRequest
	*/
	ResetGroupRateLimits(ctx context.Context, groupId string) ResetGroupRateLimitsApiRequest
	/*
		ResetGroupRateLimits Reset AI Model Rate Limits for Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResetGroupRateLimitsApiParams - Parameters for the request
		@return ResetGroupRateLimitsApiRequest
	*/
	ResetGroupRateLimitsWithParams(ctx context.Context, args *ResetGroupRateLimitsApiParams) ResetGroupRateLimitsApiRequest

	// Method available only for mocking purposes
	ResetGroupRateLimitsExecute(r ResetGroupRateLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error)

	/*
			UpdateGroupModelLimit Update AI Model Rate Limit

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Update an AI model rate limit for the given model group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param modelGroupName The name of the model group to be updated.
			@param aiModelsRateLimitUpdateRequest A request containing the new rate limits for the given model group.
			@return UpdateGroupModelLimitApiRequest
	*/
	UpdateGroupModelLimit(ctx context.Context, groupId string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateGroupModelLimitApiRequest
	/*
		UpdateGroupModelLimit Update AI Model Rate Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupModelLimitApiParams - Parameters for the request
		@return UpdateGroupModelLimitApiRequest
	*/
	UpdateGroupModelLimitWithParams(ctx context.Context, args *UpdateGroupModelLimitApiParams) UpdateGroupModelLimitApiRequest

	// Method available only for mocking purposes
	UpdateGroupModelLimitExecute(r UpdateGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error)

	/*
		UpdateGroupRateLimits Update AI Model Rate Limit

		Update a scoped AI model rate limit for the given model group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
		@param modelGroupName The name of the model group to be updated.
		@param aiModelsRateLimitUpdateRequest A request containing the new rate limits for the given model group.
		@return UpdateGroupRateLimitsApiRequest
	*/
	UpdateGroupRateLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateGroupRateLimitsApiRequest
	/*
		UpdateGroupRateLimits Update AI Model Rate Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupRateLimitsApiParams - Parameters for the request
		@return UpdateGroupRateLimitsApiRequest
	*/
	UpdateGroupRateLimitsWithParams(ctx context.Context, args *UpdateGroupRateLimitsApiParams) UpdateGroupRateLimitsApiRequest

	// Method available only for mocking purposes
	UpdateGroupRateLimitsExecute(r UpdateGroupRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error)
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

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve a single AI model rate limit for the given group.

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

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiLimits/{modelGroupName}"
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

type GetGroupModelLimitsApiRequest struct {
	ctx          context.Context
	ApiService   AIModelRateLimitsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type GetGroupModelLimitsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelRateLimitsApiService) GetGroupModelLimitsWithParams(ctx context.Context, args *GetGroupModelLimitsApiParams) GetGroupModelLimitsApiRequest {
	return GetGroupModelLimitsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r GetGroupModelLimitsApiRequest) ItemsPerPage(itemsPerPage int) GetGroupModelLimitsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetGroupModelLimitsApiRequest) PageNum(pageNum int) GetGroupModelLimitsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetGroupModelLimitsApiRequest) Execute() (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	return r.ApiService.GetGroupModelLimitsExecute(r)
}

/*
GetGroupModelLimits Return AI Model Rate Limits for One Group

Retrieve AI model rate limits for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetGroupModelLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) GetGroupModelLimits(ctx context.Context, groupId string) GetGroupModelLimitsApiRequest {
	return GetGroupModelLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetGroupModelLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelRateLimitsResponse
func (a *AIModelRateLimitsApiService) GetGroupModelLimitsExecute(r GetGroupModelLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetGroupModelLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiRateLimits"
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

type GetGroupRateLimitsApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	groupId        string
	cloud          string
	geography      string
	modelGroupName string
}

type GetGroupRateLimitsApiParams struct {
	GroupId        string
	Cloud          string
	Geography      string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) GetGroupRateLimitsWithParams(ctx context.Context, args *GetGroupRateLimitsApiParams) GetGroupRateLimitsApiRequest {
	return GetGroupRateLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		cloud:          args.Cloud,
		geography:      args.Geography,
		modelGroupName: args.ModelGroupName,
	}
}

func (r GetGroupRateLimitsApiRequest) Execute() (*AiModelRateLimitResponse, *http.Response, error) {
	return r.ApiService.GetGroupRateLimitsExecute(r)
}

/*
GetGroupRateLimits Return Single AI Model Rate Limit for One Group

Retrieve a single scoped AI model rate limit for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param modelGroupName The name of the model group to be retrieved.
	@return GetGroupRateLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) GetGroupRateLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string) GetGroupRateLimitsApiRequest {
	return GetGroupRateLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		cloud:          cloud,
		geography:      geography,
		modelGroupName: modelGroupName,
	}
}

// GetGroupRateLimitsExecute executes the request
//
//	@return AiModelRateLimitResponse
func (a *AIModelRateLimitsApiService) GetGroupRateLimitsExecute(r GetGroupRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetGroupRateLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloud == "" {
		return localVarReturnValue, nil, reportError("cloud is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloud"+"}", url.PathEscape(r.cloud), -1)
	if r.geography == "" {
		return localVarReturnValue, nil, reportError("geography is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"geography"+"}", url.PathEscape(r.geography), -1)
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

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

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

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiLimits/{modelGroupName}"
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

type GetOrgModelLimitsApiRequest struct {
	ctx          context.Context
	ApiService   AIModelRateLimitsApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type GetOrgModelLimitsApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelRateLimitsApiService) GetOrgModelLimitsWithParams(ctx context.Context, args *GetOrgModelLimitsApiParams) GetOrgModelLimitsApiRequest {
	return GetOrgModelLimitsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r GetOrgModelLimitsApiRequest) ItemsPerPage(itemsPerPage int) GetOrgModelLimitsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetOrgModelLimitsApiRequest) PageNum(pageNum int) GetOrgModelLimitsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetOrgModelLimitsApiRequest) Execute() (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	return r.ApiService.GetOrgModelLimitsExecute(r)
}

/*
GetOrgModelLimits Return AI Model Rate Limits for One Organization

Retrieve AI model rate limits for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetOrgModelLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) GetOrgModelLimits(ctx context.Context, orgId string) GetOrgModelLimitsApiRequest {
	return GetOrgModelLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetOrgModelLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelRateLimitsResponse
func (a *AIModelRateLimitsApiService) GetOrgModelLimitsExecute(r GetOrgModelLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetOrgModelLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiRateLimits"
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

type GetOrgRateLimitsApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	orgId          string
	cloud          string
	geography      string
	modelGroupName string
}

type GetOrgRateLimitsApiParams struct {
	OrgId          string
	Cloud          string
	Geography      string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) GetOrgRateLimitsWithParams(ctx context.Context, args *GetOrgRateLimitsApiParams) GetOrgRateLimitsApiRequest {
	return GetOrgRateLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          args.OrgId,
		cloud:          args.Cloud,
		geography:      args.Geography,
		modelGroupName: args.ModelGroupName,
	}
}

func (r GetOrgRateLimitsApiRequest) Execute() (*AiModelRateLimitResponse, *http.Response, error) {
	return r.ApiService.GetOrgRateLimitsExecute(r)
}

/*
GetOrgRateLimits Return Single AI Model Rate Limit for One Organization

Retrieve a single scoped AI model rate limit for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param modelGroupName The name of the model group to be retrieved.
	@return GetOrgRateLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) GetOrgRateLimits(ctx context.Context, orgId string, cloud string, geography string, modelGroupName string) GetOrgRateLimitsApiRequest {
	return GetOrgRateLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          orgId,
		cloud:          cloud,
		geography:      geography,
		modelGroupName: modelGroupName,
	}
}

// GetOrgRateLimitsExecute executes the request
//
//	@return AiModelRateLimitResponse
func (a *AIModelRateLimitsApiService) GetOrgRateLimitsExecute(r GetOrgRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.GetOrgRateLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.cloud == "" {
		return localVarReturnValue, nil, reportError("cloud is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloud"+"}", url.PathEscape(r.cloud), -1)
	if r.geography == "" {
		return localVarReturnValue, nil, reportError("geography is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"geography"+"}", url.PathEscape(r.geography), -1)
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

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve AI model rate limits for the given group.

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

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiLimits"
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

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

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

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiLimits"
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

type ResetGroupApiLimitsApiRequest struct {
	ctx        context.Context
	ApiService AIModelRateLimitsApi
	groupId    string
}

type ResetGroupApiLimitsApiParams struct {
	GroupId string
}

func (a *AIModelRateLimitsApiService) ResetGroupApiLimitsWithParams(ctx context.Context, args *ResetGroupApiLimitsApiParams) ResetGroupApiLimitsApiRequest {
	return ResetGroupApiLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ResetGroupApiLimitsApiRequest) Execute() (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	return r.ApiService.ResetGroupApiLimitsExecute(r)
}

/*
ResetGroupApiLimits Reset AI Model Rate Limits for Group

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Reset the AI Model rate limits for the given group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ResetGroupApiLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetGroupApiLimits(ctx context.Context, groupId string) ResetGroupApiLimitsApiRequest {
	return ResetGroupApiLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ResetGroupApiLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelsRateLimitsResponse
func (a *AIModelRateLimitsApiService) ResetGroupApiLimitsExecute(r ResetGroupApiLimitsApiRequest) (*PaginatedAtlasAiModelsRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetGroupApiLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiLimits:reset"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
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

type ResetGroupModelLimitApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	groupId        string
	modelGroupName string
}

type ResetGroupModelLimitApiParams struct {
	GroupId        string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) ResetGroupModelLimitWithParams(ctx context.Context, args *ResetGroupModelLimitApiParams) ResetGroupModelLimitApiRequest {
	return ResetGroupModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		modelGroupName: args.ModelGroupName,
	}
}

func (r ResetGroupModelLimitApiRequest) Execute() (*AiModelsRateLimitResponse, *http.Response, error) {
	return r.ApiService.ResetGroupModelLimitExecute(r)
}

/*
ResetGroupModelLimit Reset AI Model Rate Limit for One Model Group

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Reset the AI model rate limit for the given model group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param modelGroupName The name of the model group to be reset to default rate limits.
	@return ResetGroupModelLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetGroupModelLimit(ctx context.Context, groupId string, modelGroupName string) ResetGroupModelLimitApiRequest {
	return ResetGroupModelLimitApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		modelGroupName: modelGroupName,
	}
}

// ResetGroupModelLimitExecute executes the request
//
//	@return AiModelsRateLimitResponse
func (a *AIModelRateLimitsApiService) ResetGroupModelLimitExecute(r ResetGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetGroupModelLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiLimits/{modelGroupName}:reset"
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

type ResetGroupModelLimitsApiRequest struct {
	ctx            context.Context
	ApiService     AIModelRateLimitsApi
	groupId        string
	cloud          string
	geography      string
	modelGroupName string
}

type ResetGroupModelLimitsApiParams struct {
	GroupId        string
	Cloud          string
	Geography      string
	ModelGroupName string
}

func (a *AIModelRateLimitsApiService) ResetGroupModelLimitsWithParams(ctx context.Context, args *ResetGroupModelLimitsApiParams) ResetGroupModelLimitsApiRequest {
	return ResetGroupModelLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		cloud:          args.Cloud,
		geography:      args.Geography,
		modelGroupName: args.ModelGroupName,
	}
}

func (r ResetGroupModelLimitsApiRequest) Execute() (*AiModelRateLimitResponse, *http.Response, error) {
	return r.ApiService.ResetGroupModelLimitsExecute(r)
}

/*
ResetGroupModelLimits Reset AI Model Rate Limit for One Model Group

Reset the scoped AI model rate limit for the given model group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param modelGroupName The name of the model group to be reset to default rate limits.
	@return ResetGroupModelLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetGroupModelLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string) ResetGroupModelLimitsApiRequest {
	return ResetGroupModelLimitsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		cloud:          cloud,
		geography:      geography,
		modelGroupName: modelGroupName,
	}
}

// ResetGroupModelLimitsExecute executes the request
//
//	@return AiModelRateLimitResponse
func (a *AIModelRateLimitsApiService) ResetGroupModelLimitsExecute(r ResetGroupModelLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetGroupModelLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits:reset"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloud == "" {
		return localVarReturnValue, nil, reportError("cloud is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloud"+"}", url.PathEscape(r.cloud), -1)
	if r.geography == "" {
		return localVarReturnValue, nil, reportError("geography is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"geography"+"}", url.PathEscape(r.geography), -1)
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

type ResetGroupRateLimitsApiRequest struct {
	ctx        context.Context
	ApiService AIModelRateLimitsApi
	groupId    string
}

type ResetGroupRateLimitsApiParams struct {
	GroupId string
}

func (a *AIModelRateLimitsApiService) ResetGroupRateLimitsWithParams(ctx context.Context, args *ResetGroupRateLimitsApiParams) ResetGroupRateLimitsApiRequest {
	return ResetGroupRateLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ResetGroupRateLimitsApiRequest) Execute() (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	return r.ApiService.ResetGroupRateLimitsExecute(r)
}

/*
ResetGroupRateLimits Reset AI Model Rate Limits for Group

Reset the AI Model rate limits for the given group to default values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ResetGroupRateLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) ResetGroupRateLimits(ctx context.Context, groupId string) ResetGroupRateLimitsApiRequest {
	return ResetGroupRateLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ResetGroupRateLimitsExecute executes the request
//
//	@return PaginatedAtlasAiModelRateLimitsResponse
func (a *AIModelRateLimitsApiService) ResetGroupRateLimitsExecute(r ResetGroupRateLimitsApiRequest) (*PaginatedAtlasAiModelRateLimitsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelRateLimitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.ResetGroupRateLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiRateLimits:reset"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
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

type UpdateGroupModelLimitApiRequest struct {
	ctx                            context.Context
	ApiService                     AIModelRateLimitsApi
	groupId                        string
	modelGroupName                 string
	aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

type UpdateGroupModelLimitApiParams struct {
	GroupId                        string
	ModelGroupName                 string
	AiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

func (a *AIModelRateLimitsApiService) UpdateGroupModelLimitWithParams(ctx context.Context, args *UpdateGroupModelLimitApiParams) UpdateGroupModelLimitApiRequest {
	return UpdateGroupModelLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		modelGroupName:                 args.ModelGroupName,
		aiModelsRateLimitUpdateRequest: args.AiModelsRateLimitUpdateRequest,
	}
}

func (r UpdateGroupModelLimitApiRequest) Execute() (*AiModelsRateLimitResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupModelLimitExecute(r)
}

/*
UpdateGroupModelLimit Update AI Model Rate Limit

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Update an AI model rate limit for the given model group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param modelGroupName The name of the model group to be updated.
	@return UpdateGroupModelLimitApiRequest
*/
func (a *AIModelRateLimitsApiService) UpdateGroupModelLimit(ctx context.Context, groupId string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateGroupModelLimitApiRequest {
	return UpdateGroupModelLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		modelGroupName:                 modelGroupName,
		aiModelsRateLimitUpdateRequest: aiModelsRateLimitUpdateRequest,
	}
}

// UpdateGroupModelLimitExecute executes the request
//
//	@return AiModelsRateLimitResponse
func (a *AIModelRateLimitsApiService) UpdateGroupModelLimitExecute(r UpdateGroupModelLimitApiRequest) (*AiModelsRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.UpdateGroupModelLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiLimits/{modelGroupName}"
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

type UpdateGroupRateLimitsApiRequest struct {
	ctx                            context.Context
	ApiService                     AIModelRateLimitsApi
	groupId                        string
	cloud                          string
	geography                      string
	modelGroupName                 string
	aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

type UpdateGroupRateLimitsApiParams struct {
	GroupId                        string
	Cloud                          string
	Geography                      string
	ModelGroupName                 string
	AiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest
}

func (a *AIModelRateLimitsApiService) UpdateGroupRateLimitsWithParams(ctx context.Context, args *UpdateGroupRateLimitsApiParams) UpdateGroupRateLimitsApiRequest {
	return UpdateGroupRateLimitsApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		cloud:                          args.Cloud,
		geography:                      args.Geography,
		modelGroupName:                 args.ModelGroupName,
		aiModelsRateLimitUpdateRequest: args.AiModelsRateLimitUpdateRequest,
	}
}

func (r UpdateGroupRateLimitsApiRequest) Execute() (*AiModelRateLimitResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupRateLimitsExecute(r)
}

/*
UpdateGroupRateLimits Update AI Model Rate Limit

Update a scoped AI model rate limit for the given model group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloud Cloud provider scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param geography Geography scope. Must be \"ANY\". Additional values will be supported in future API versions.
	@param modelGroupName The name of the model group to be updated.
	@return UpdateGroupRateLimitsApiRequest
*/
func (a *AIModelRateLimitsApiService) UpdateGroupRateLimits(ctx context.Context, groupId string, cloud string, geography string, modelGroupName string, aiModelsRateLimitUpdateRequest *AiModelsRateLimitUpdateRequest) UpdateGroupRateLimitsApiRequest {
	return UpdateGroupRateLimitsApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		cloud:                          cloud,
		geography:                      geography,
		modelGroupName:                 modelGroupName,
		aiModelsRateLimitUpdateRequest: aiModelsRateLimitUpdateRequest,
	}
}

// UpdateGroupRateLimitsExecute executes the request
//
//	@return AiModelRateLimitResponse
func (a *AIModelRateLimitsApiService) UpdateGroupRateLimitsExecute(r UpdateGroupRateLimitsApiRequest) (*AiModelRateLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelRateLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelRateLimitsApiService.UpdateGroupRateLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiClouds/{cloud}/geographies/{geography}/modelGroupNames/{modelGroupName}/rateLimits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloud == "" {
		return localVarReturnValue, nil, reportError("cloud is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloud"+"}", url.PathEscape(r.cloud), -1)
	if r.geography == "" {
		return localVarReturnValue, nil, reportError("geography is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"geography"+"}", url.PathEscape(r.geography), -1)
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
