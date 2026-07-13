// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AIModelAPIKeysApi interface {

	/*
			CreateGroupAiKey Create New AI Model API Key

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Create a new AI model API key for the given group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param aiModelsApiKeyCreateRequest A request containing the name of the new API key.
			@return CreateGroupAiKeyApiRequest
	*/
	CreateGroupAiKey(ctx context.Context, groupId string, aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest) CreateGroupAiKeyApiRequest
	/*
		CreateGroupAiKey Create New AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupAiKeyApiParams - Parameters for the request
		@return CreateGroupAiKeyApiRequest
	*/
	CreateGroupAiKeyWithParams(ctx context.Context, args *CreateGroupAiKeyApiParams) CreateGroupAiKeyApiRequest

	// Method available only for mocking purposes
	CreateGroupAiKeyExecute(r CreateGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		CreateGroupModelKey Create New AI Model API Key

		Create a new AI model API key for the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param aiModelApiKeyCreateRequest A request containing the name, cloud, and geography of the new API key.
		@return CreateGroupModelKeyApiRequest
	*/
	CreateGroupModelKey(ctx context.Context, groupId string, aiModelApiKeyCreateRequest *AiModelApiKeyCreateRequest) CreateGroupModelKeyApiRequest
	/*
		CreateGroupModelKey Create New AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupModelKeyApiParams - Parameters for the request
		@return CreateGroupModelKeyApiRequest
	*/
	CreateGroupModelKeyWithParams(ctx context.Context, args *CreateGroupModelKeyApiParams) CreateGroupModelKeyApiRequest

	// Method available only for mocking purposes
	CreateGroupModelKeyExecute(r CreateGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error)

	/*
			DeleteGroupAiKey Delete Existing AI Model API Key

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Delete an existing AI model API key in the given group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param apiKeyId The id of the API key to be deleted.
			@return DeleteGroupAiKeyApiRequest
	*/
	DeleteGroupAiKey(ctx context.Context, groupId string, apiKeyId string) DeleteGroupAiKeyApiRequest
	/*
		DeleteGroupAiKey Delete Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupAiKeyApiParams - Parameters for the request
		@return DeleteGroupAiKeyApiRequest
	*/
	DeleteGroupAiKeyWithParams(ctx context.Context, args *DeleteGroupAiKeyApiParams) DeleteGroupAiKeyApiRequest

	// Method available only for mocking purposes
	DeleteGroupAiKeyExecute(r DeleteGroupAiKeyApiRequest) (*http.Response, error)

	/*
		DeleteGroupModelKey Delete Existing AI Model API Key

		Delete an existing AI model API key in the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiKeyId The id of the API key to be deleted.
		@return DeleteGroupModelKeyApiRequest
	*/
	DeleteGroupModelKey(ctx context.Context, groupId string, apiKeyId string) DeleteGroupModelKeyApiRequest
	/*
		DeleteGroupModelKey Delete Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupModelKeyApiParams - Parameters for the request
		@return DeleteGroupModelKeyApiRequest
	*/
	DeleteGroupModelKeyWithParams(ctx context.Context, args *DeleteGroupModelKeyApiParams) DeleteGroupModelKeyApiRequest

	// Method available only for mocking purposes
	DeleteGroupModelKeyExecute(r DeleteGroupModelKeyApiRequest) (*http.Response, error)

	/*
			GetGroupAiKey Return Single AI Model API Key for One Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve a single AI model API key for the given group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param apiKeyId The id of the API key to be retrieved.
			@return GetGroupAiKeyApiRequest
	*/
	GetGroupAiKey(ctx context.Context, groupId string, apiKeyId string) GetGroupAiKeyApiRequest
	/*
		GetGroupAiKey Return Single AI Model API Key for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupAiKeyApiParams - Parameters for the request
		@return GetGroupAiKeyApiRequest
	*/
	GetGroupAiKeyWithParams(ctx context.Context, args *GetGroupAiKeyApiParams) GetGroupAiKeyApiRequest

	// Method available only for mocking purposes
	GetGroupAiKeyExecute(r GetGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		GetGroupModelKey Return Single AI Model API Key for One Group

		Retrieve a single AI model API key for the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiKeyId The id of the API key to be retrieved.
		@return GetGroupModelKeyApiRequest
	*/
	GetGroupModelKey(ctx context.Context, groupId string, apiKeyId string) GetGroupModelKeyApiRequest
	/*
		GetGroupModelKey Return Single AI Model API Key for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupModelKeyApiParams - Parameters for the request
		@return GetGroupModelKeyApiRequest
	*/
	GetGroupModelKeyWithParams(ctx context.Context, args *GetGroupModelKeyApiParams) GetGroupModelKeyApiRequest

	// Method available only for mocking purposes
	GetGroupModelKeyExecute(r GetGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error)

	/*
			GetOrgAiKey Return Single AI Model API Key for One Organization

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve a single AI model API key for the given organization.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param apiKeyId The id of the API key to be retrieved.
			@return GetOrgAiKeyApiRequest
	*/
	GetOrgAiKey(ctx context.Context, orgId string, apiKeyId string) GetOrgAiKeyApiRequest
	/*
		GetOrgAiKey Return Single AI Model API Key for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgAiKeyApiParams - Parameters for the request
		@return GetOrgAiKeyApiRequest
	*/
	GetOrgAiKeyWithParams(ctx context.Context, args *GetOrgAiKeyApiParams) GetOrgAiKeyApiRequest

	// Method available only for mocking purposes
	GetOrgAiKeyExecute(r GetOrgAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		GetOrgModelKey Return Single AI Model API Key for One Organization

		Retrieve a single AI model API key for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiKeyId The id of the API key to be retrieved.
		@return GetOrgModelKeyApiRequest
	*/
	GetOrgModelKey(ctx context.Context, orgId string, apiKeyId string) GetOrgModelKeyApiRequest
	/*
		GetOrgModelKey Return Single AI Model API Key for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgModelKeyApiParams - Parameters for the request
		@return GetOrgModelKeyApiRequest
	*/
	GetOrgModelKeyWithParams(ctx context.Context, args *GetOrgModelKeyApiParams) GetOrgModelKeyApiRequest

	// Method available only for mocking purposes
	GetOrgModelKeyExecute(r GetOrgModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error)

	/*
			ListGroupAiKeys Return AI Model API Keys for One Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve AI model API keys for the given group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListGroupAiKeysApiRequest
	*/
	ListGroupAiKeys(ctx context.Context, groupId string) ListGroupAiKeysApiRequest
	/*
		ListGroupAiKeys Return AI Model API Keys for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupAiKeysApiParams - Parameters for the request
		@return ListGroupAiKeysApiRequest
	*/
	ListGroupAiKeysWithParams(ctx context.Context, args *ListGroupAiKeysApiParams) ListGroupAiKeysApiRequest

	// Method available only for mocking purposes
	ListGroupAiKeysExecute(r ListGroupAiKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error)

	/*
		ListGroupModelKeys Return AI Model API Keys for One Group

		Retrieve AI model API keys for the given group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListGroupModelKeysApiRequest
	*/
	ListGroupModelKeys(ctx context.Context, groupId string) ListGroupModelKeysApiRequest
	/*
		ListGroupModelKeys Return AI Model API Keys for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupModelKeysApiParams - Parameters for the request
		@return ListGroupModelKeysApiRequest
	*/
	ListGroupModelKeysWithParams(ctx context.Context, args *ListGroupModelKeysApiParams) ListGroupModelKeysApiRequest

	// Method available only for mocking purposes
	ListGroupModelKeysExecute(r ListGroupModelKeysApiRequest) (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error)

	/*
			ListOrgAiKeys Return AI Model API Keys for One Organization

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve AI model API keys for the given organization.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrgAiKeysApiRequest
	*/
	ListOrgAiKeys(ctx context.Context, orgId string) ListOrgAiKeysApiRequest
	/*
		ListOrgAiKeys Return AI Model API Keys for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgAiKeysApiParams - Parameters for the request
		@return ListOrgAiKeysApiRequest
	*/
	ListOrgAiKeysWithParams(ctx context.Context, args *ListOrgAiKeysApiParams) ListOrgAiKeysApiRequest

	// Method available only for mocking purposes
	ListOrgAiKeysExecute(r ListOrgAiKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error)

	/*
		ListOrgModelKeys Return AI Model API Keys for One Organization

		Retrieve AI model API keys for the given organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrgModelKeysApiRequest
	*/
	ListOrgModelKeys(ctx context.Context, orgId string) ListOrgModelKeysApiRequest
	/*
		ListOrgModelKeys Return AI Model API Keys for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgModelKeysApiParams - Parameters for the request
		@return ListOrgModelKeysApiRequest
	*/
	ListOrgModelKeysWithParams(ctx context.Context, args *ListOrgModelKeysApiParams) ListOrgModelKeysApiRequest

	// Method available only for mocking purposes
	ListOrgModelKeysExecute(r ListOrgModelKeysApiRequest) (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error)

	/*
			UpdateGroupAiKey Update Existing AI Model API Key

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Update an existing AI model API key in the given group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param apiKeyId The id of the API key to be updated.
			@param aiModelsApiKeyUpdateRequest A request containing the new name for the API key.
			@return UpdateGroupAiKeyApiRequest
	*/
	UpdateGroupAiKey(ctx context.Context, groupId string, apiKeyId string, aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest) UpdateGroupAiKeyApiRequest
	/*
		UpdateGroupAiKey Update Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupAiKeyApiParams - Parameters for the request
		@return UpdateGroupAiKeyApiRequest
	*/
	UpdateGroupAiKeyWithParams(ctx context.Context, args *UpdateGroupAiKeyApiParams) UpdateGroupAiKeyApiRequest

	// Method available only for mocking purposes
	UpdateGroupAiKeyExecute(r UpdateGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		UpdateGroupModelKey Update Existing AI Model API Key

		Update an existing AI model API key in the given group. Only the name can be updated; scope is immutable after creation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiKeyId The id of the API key to be updated.
		@param aiModelApiKeyUpdateRequest A request containing the new name for the API key.
		@return UpdateGroupModelKeyApiRequest
	*/
	UpdateGroupModelKey(ctx context.Context, groupId string, apiKeyId string, aiModelApiKeyUpdateRequest *AiModelApiKeyUpdateRequest) UpdateGroupModelKeyApiRequest
	/*
		UpdateGroupModelKey Update Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupModelKeyApiParams - Parameters for the request
		@return UpdateGroupModelKeyApiRequest
	*/
	UpdateGroupModelKeyWithParams(ctx context.Context, args *UpdateGroupModelKeyApiParams) UpdateGroupModelKeyApiRequest

	// Method available only for mocking purposes
	UpdateGroupModelKeyExecute(r UpdateGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error)
}

// AIModelAPIKeysApiService AIModelAPIKeysApi service
type AIModelAPIKeysApiService service

type CreateGroupAiKeyApiRequest struct {
	ctx                         context.Context
	ApiService                  AIModelAPIKeysApi
	groupId                     string
	aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest
}

type CreateGroupAiKeyApiParams struct {
	GroupId                     string
	AiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest
}

func (a *AIModelAPIKeysApiService) CreateGroupAiKeyWithParams(ctx context.Context, args *CreateGroupAiKeyApiParams) CreateGroupAiKeyApiRequest {
	return CreateGroupAiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		aiModelsApiKeyCreateRequest: args.AiModelsApiKeyCreateRequest,
	}
}

func (r CreateGroupAiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.CreateGroupAiKeyExecute(r)
}

/*
CreateGroupAiKey Create New AI Model API Key

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Create a new AI model API key for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateGroupAiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) CreateGroupAiKey(ctx context.Context, groupId string, aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest) CreateGroupAiKeyApiRequest {
	return CreateGroupAiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		aiModelsApiKeyCreateRequest: aiModelsApiKeyCreateRequest,
	}
}

// CreateGroupAiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) CreateGroupAiKeyExecute(r CreateGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.CreateGroupAiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelKeys"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiModelsApiKeyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("aiModelsApiKeyCreateRequest is required and must be specified")
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
	localVarPostBody = r.aiModelsApiKeyCreateRequest
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

type CreateGroupModelKeyApiRequest struct {
	ctx                        context.Context
	ApiService                 AIModelAPIKeysApi
	groupId                    string
	aiModelApiKeyCreateRequest *AiModelApiKeyCreateRequest
}

type CreateGroupModelKeyApiParams struct {
	GroupId                    string
	AiModelApiKeyCreateRequest *AiModelApiKeyCreateRequest
}

func (a *AIModelAPIKeysApiService) CreateGroupModelKeyWithParams(ctx context.Context, args *CreateGroupModelKeyApiParams) CreateGroupModelKeyApiRequest {
	return CreateGroupModelKeyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		aiModelApiKeyCreateRequest: args.AiModelApiKeyCreateRequest,
	}
}

func (r CreateGroupModelKeyApiRequest) Execute() (*AiModelApiKeyResponse, *http.Response, error) {
	return r.ApiService.CreateGroupModelKeyExecute(r)
}

/*
CreateGroupModelKey Create New AI Model API Key

Create a new AI model API key for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateGroupModelKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) CreateGroupModelKey(ctx context.Context, groupId string, aiModelApiKeyCreateRequest *AiModelApiKeyCreateRequest) CreateGroupModelKeyApiRequest {
	return CreateGroupModelKeyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		aiModelApiKeyCreateRequest: aiModelApiKeyCreateRequest,
	}
}

// CreateGroupModelKeyExecute executes the request
//
//	@return AiModelApiKeyResponse
func (a *AIModelAPIKeysApiService) CreateGroupModelKeyExecute(r CreateGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.CreateGroupModelKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiKeys"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiModelApiKeyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("aiModelApiKeyCreateRequest is required and must be specified")
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
	localVarPostBody = r.aiModelApiKeyCreateRequest
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

type DeleteGroupAiKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	groupId    string
	apiKeyId   string
}

type DeleteGroupAiKeyApiParams struct {
	GroupId  string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) DeleteGroupAiKeyWithParams(ctx context.Context, args *DeleteGroupAiKeyApiParams) DeleteGroupAiKeyApiRequest {
	return DeleteGroupAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r DeleteGroupAiKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupAiKeyExecute(r)
}

/*
DeleteGroupAiKey Delete Existing AI Model API Key

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Delete an existing AI model API key in the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be deleted.
	@return DeleteGroupAiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) DeleteGroupAiKey(ctx context.Context, groupId string, apiKeyId string) DeleteGroupAiKeyApiRequest {
	return DeleteGroupAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiKeyId:   apiKeyId,
	}
}

// DeleteGroupAiKeyExecute executes the request
func (a *AIModelAPIKeysApiService) DeleteGroupAiKeyExecute(r DeleteGroupAiKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.DeleteGroupAiKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelKeys/{apiKeyId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type DeleteGroupModelKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	groupId    string
	apiKeyId   string
}

type DeleteGroupModelKeyApiParams struct {
	GroupId  string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) DeleteGroupModelKeyWithParams(ctx context.Context, args *DeleteGroupModelKeyApiParams) DeleteGroupModelKeyApiRequest {
	return DeleteGroupModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r DeleteGroupModelKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupModelKeyExecute(r)
}

/*
DeleteGroupModelKey Delete Existing AI Model API Key

Delete an existing AI model API key in the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be deleted.
	@return DeleteGroupModelKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) DeleteGroupModelKey(ctx context.Context, groupId string, apiKeyId string) DeleteGroupModelKeyApiRequest {
	return DeleteGroupModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiKeyId:   apiKeyId,
	}
}

// DeleteGroupModelKeyExecute executes the request
func (a *AIModelAPIKeysApiService) DeleteGroupModelKeyExecute(r DeleteGroupModelKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.DeleteGroupModelKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type GetGroupAiKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	groupId    string
	apiKeyId   string
}

type GetGroupAiKeyApiParams struct {
	GroupId  string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) GetGroupAiKeyWithParams(ctx context.Context, args *GetGroupAiKeyApiParams) GetGroupAiKeyApiRequest {
	return GetGroupAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r GetGroupAiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.GetGroupAiKeyExecute(r)
}

/*
GetGroupAiKey Return Single AI Model API Key for One Group

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve a single AI model API key for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be retrieved.
	@return GetGroupAiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) GetGroupAiKey(ctx context.Context, groupId string, apiKeyId string) GetGroupAiKeyApiRequest {
	return GetGroupAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiKeyId:   apiKeyId,
	}
}

// GetGroupAiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) GetGroupAiKeyExecute(r GetGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.GetGroupAiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelKeys/{apiKeyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type GetGroupModelKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	groupId    string
	apiKeyId   string
}

type GetGroupModelKeyApiParams struct {
	GroupId  string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) GetGroupModelKeyWithParams(ctx context.Context, args *GetGroupModelKeyApiParams) GetGroupModelKeyApiRequest {
	return GetGroupModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r GetGroupModelKeyApiRequest) Execute() (*AiModelApiKeyResponse, *http.Response, error) {
	return r.ApiService.GetGroupModelKeyExecute(r)
}

/*
GetGroupModelKey Return Single AI Model API Key for One Group

Retrieve a single AI model API key for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be retrieved.
	@return GetGroupModelKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) GetGroupModelKey(ctx context.Context, groupId string, apiKeyId string) GetGroupModelKeyApiRequest {
	return GetGroupModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiKeyId:   apiKeyId,
	}
}

// GetGroupModelKeyExecute executes the request
//
//	@return AiModelApiKeyResponse
func (a *AIModelAPIKeysApiService) GetGroupModelKeyExecute(r GetGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.GetGroupModelKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type GetOrgAiKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	orgId      string
	apiKeyId   string
}

type GetOrgAiKeyApiParams struct {
	OrgId    string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) GetOrgAiKeyWithParams(ctx context.Context, args *GetOrgAiKeyApiParams) GetOrgAiKeyApiRequest {
	return GetOrgAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r GetOrgAiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.GetOrgAiKeyExecute(r)
}

/*
GetOrgAiKey Return Single AI Model API Key for One Organization

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve a single AI model API key for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiKeyId The id of the API key to be retrieved.
	@return GetOrgAiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) GetOrgAiKey(ctx context.Context, orgId string, apiKeyId string) GetOrgAiKeyApiRequest {
	return GetOrgAiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiKeyId:   apiKeyId,
	}
}

// GetOrgAiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) GetOrgAiKeyExecute(r GetOrgAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.GetOrgAiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelKeys/{apiKeyId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type GetOrgModelKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	orgId      string
	apiKeyId   string
}

type GetOrgModelKeyApiParams struct {
	OrgId    string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) GetOrgModelKeyWithParams(ctx context.Context, args *GetOrgModelKeyApiParams) GetOrgModelKeyApiRequest {
	return GetOrgModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r GetOrgModelKeyApiRequest) Execute() (*AiModelApiKeyResponse, *http.Response, error) {
	return r.ApiService.GetOrgModelKeyExecute(r)
}

/*
GetOrgModelKey Return Single AI Model API Key for One Organization

Retrieve a single AI model API key for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiKeyId The id of the API key to be retrieved.
	@return GetOrgModelKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) GetOrgModelKey(ctx context.Context, orgId string, apiKeyId string) GetOrgModelKeyApiRequest {
	return GetOrgModelKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiKeyId:   apiKeyId,
	}
}

// GetOrgModelKeyExecute executes the request
//
//	@return AiModelApiKeyResponse
func (a *AIModelAPIKeysApiService) GetOrgModelKeyExecute(r GetOrgModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.GetOrgModelKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiKeys/{apiKeyId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

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

type ListGroupAiKeysApiRequest struct {
	ctx          context.Context
	ApiService   AIModelAPIKeysApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListGroupAiKeysApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelAPIKeysApiService) ListGroupAiKeysWithParams(ctx context.Context, args *ListGroupAiKeysApiParams) ListGroupAiKeysApiRequest {
	return ListGroupAiKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupAiKeysApiRequest) ItemsPerPage(itemsPerPage int) ListGroupAiKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupAiKeysApiRequest) PageNum(pageNum int) ListGroupAiKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupAiKeysApiRequest) Execute() (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	return r.ApiService.ListGroupAiKeysExecute(r)
}

/*
ListGroupAiKeys Return AI Model API Keys for One Group

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve AI model API keys for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListGroupAiKeysApiRequest
*/
func (a *AIModelAPIKeysApiService) ListGroupAiKeys(ctx context.Context, groupId string) ListGroupAiKeysApiRequest {
	return ListGroupAiKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupAiKeysExecute executes the request
//
//	@return PaginatedAtlasAiModelsApiKeysResponse
func (a *AIModelAPIKeysApiService) ListGroupAiKeysExecute(r ListGroupAiKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsApiKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.ListGroupAiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelKeys"
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

type ListGroupModelKeysApiRequest struct {
	ctx          context.Context
	ApiService   AIModelAPIKeysApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListGroupModelKeysApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelAPIKeysApiService) ListGroupModelKeysWithParams(ctx context.Context, args *ListGroupModelKeysApiParams) ListGroupModelKeysApiRequest {
	return ListGroupModelKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupModelKeysApiRequest) ItemsPerPage(itemsPerPage int) ListGroupModelKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupModelKeysApiRequest) PageNum(pageNum int) ListGroupModelKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupModelKeysApiRequest) Execute() (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error) {
	return r.ApiService.ListGroupModelKeysExecute(r)
}

/*
ListGroupModelKeys Return AI Model API Keys for One Group

Retrieve AI model API keys for the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListGroupModelKeysApiRequest
*/
func (a *AIModelAPIKeysApiService) ListGroupModelKeys(ctx context.Context, groupId string) ListGroupModelKeysApiRequest {
	return ListGroupModelKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupModelKeysExecute executes the request
//
//	@return PaginatedAtlasAiModelApiKeysResponse
func (a *AIModelAPIKeysApiService) ListGroupModelKeysExecute(r ListGroupModelKeysApiRequest) (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelApiKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.ListGroupModelKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiKeys"
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

type ListOrgAiKeysApiRequest struct {
	ctx          context.Context
	ApiService   AIModelAPIKeysApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type ListOrgAiKeysApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelAPIKeysApiService) ListOrgAiKeysWithParams(ctx context.Context, args *ListOrgAiKeysApiParams) ListOrgAiKeysApiRequest {
	return ListOrgAiKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgAiKeysApiRequest) ItemsPerPage(itemsPerPage int) ListOrgAiKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgAiKeysApiRequest) PageNum(pageNum int) ListOrgAiKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgAiKeysApiRequest) Execute() (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	return r.ApiService.ListOrgAiKeysExecute(r)
}

/*
ListOrgAiKeys Return AI Model API Keys for One Organization

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve AI model API keys for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrgAiKeysApiRequest
*/
func (a *AIModelAPIKeysApiService) ListOrgAiKeys(ctx context.Context, orgId string) ListOrgAiKeysApiRequest {
	return ListOrgAiKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrgAiKeysExecute executes the request
//
//	@return PaginatedAtlasAiModelsApiKeysResponse
func (a *AIModelAPIKeysApiService) ListOrgAiKeysExecute(r ListOrgAiKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsApiKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.ListOrgAiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelKeys"
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

type ListOrgModelKeysApiRequest struct {
	ctx          context.Context
	ApiService   AIModelAPIKeysApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type ListOrgModelKeysApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *AIModelAPIKeysApiService) ListOrgModelKeysWithParams(ctx context.Context, args *ListOrgModelKeysApiParams) ListOrgModelKeysApiRequest {
	return ListOrgModelKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgModelKeysApiRequest) ItemsPerPage(itemsPerPage int) ListOrgModelKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgModelKeysApiRequest) PageNum(pageNum int) ListOrgModelKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgModelKeysApiRequest) Execute() (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error) {
	return r.ApiService.ListOrgModelKeysExecute(r)
}

/*
ListOrgModelKeys Return AI Model API Keys for One Organization

Retrieve AI model API keys for the given organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrgModelKeysApiRequest
*/
func (a *AIModelAPIKeysApiService) ListOrgModelKeys(ctx context.Context, orgId string) ListOrgModelKeysApiRequest {
	return ListOrgModelKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrgModelKeysExecute executes the request
//
//	@return PaginatedAtlasAiModelApiKeysResponse
func (a *AIModelAPIKeysApiService) ListOrgModelKeysExecute(r ListOrgModelKeysApiRequest) (*PaginatedAtlasAiModelApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelApiKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.ListOrgModelKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/aiModelApiKeys"
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

type UpdateGroupAiKeyApiRequest struct {
	ctx                         context.Context
	ApiService                  AIModelAPIKeysApi
	groupId                     string
	apiKeyId                    string
	aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest
}

type UpdateGroupAiKeyApiParams struct {
	GroupId                     string
	ApiKeyId                    string
	AiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest
}

func (a *AIModelAPIKeysApiService) UpdateGroupAiKeyWithParams(ctx context.Context, args *UpdateGroupAiKeyApiParams) UpdateGroupAiKeyApiRequest {
	return UpdateGroupAiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		apiKeyId:                    args.ApiKeyId,
		aiModelsApiKeyUpdateRequest: args.AiModelsApiKeyUpdateRequest,
	}
}

func (r UpdateGroupAiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupAiKeyExecute(r)
}

/*
UpdateGroupAiKey Update Existing AI Model API Key

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Update an existing AI model API key in the given group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be updated.
	@return UpdateGroupAiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) UpdateGroupAiKey(ctx context.Context, groupId string, apiKeyId string, aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest) UpdateGroupAiKeyApiRequest {
	return UpdateGroupAiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		apiKeyId:                    apiKeyId,
		aiModelsApiKeyUpdateRequest: aiModelsApiKeyUpdateRequest,
	}
}

// UpdateGroupAiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) UpdateGroupAiKeyExecute(r UpdateGroupAiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.UpdateGroupAiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelKeys/{apiKeyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiModelsApiKeyUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("aiModelsApiKeyUpdateRequest is required and must be specified")
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
	localVarPostBody = r.aiModelsApiKeyUpdateRequest
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

type UpdateGroupModelKeyApiRequest struct {
	ctx                        context.Context
	ApiService                 AIModelAPIKeysApi
	groupId                    string
	apiKeyId                   string
	aiModelApiKeyUpdateRequest *AiModelApiKeyUpdateRequest
}

type UpdateGroupModelKeyApiParams struct {
	GroupId                    string
	ApiKeyId                   string
	AiModelApiKeyUpdateRequest *AiModelApiKeyUpdateRequest
}

func (a *AIModelAPIKeysApiService) UpdateGroupModelKeyWithParams(ctx context.Context, args *UpdateGroupModelKeyApiParams) UpdateGroupModelKeyApiRequest {
	return UpdateGroupModelKeyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		apiKeyId:                   args.ApiKeyId,
		aiModelApiKeyUpdateRequest: args.AiModelApiKeyUpdateRequest,
	}
}

func (r UpdateGroupModelKeyApiRequest) Execute() (*AiModelApiKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupModelKeyExecute(r)
}

/*
UpdateGroupModelKey Update Existing AI Model API Key

Update an existing AI model API key in the given group. Only the name can be updated; scope is immutable after creation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be updated.
	@return UpdateGroupModelKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) UpdateGroupModelKey(ctx context.Context, groupId string, apiKeyId string, aiModelApiKeyUpdateRequest *AiModelApiKeyUpdateRequest) UpdateGroupModelKeyApiRequest {
	return UpdateGroupModelKeyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		apiKeyId:                   apiKeyId,
		aiModelApiKeyUpdateRequest: aiModelApiKeyUpdateRequest,
	}
}

// UpdateGroupModelKeyExecute executes the request
//
//	@return AiModelApiKeyResponse
func (a *AIModelAPIKeysApiService) UpdateGroupModelKeyExecute(r UpdateGroupModelKeyApiRequest) (*AiModelApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.UpdateGroupModelKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/aiModelApiKeys/{apiKeyId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiKeyId == "" {
		return localVarReturnValue, nil, reportError("apiKeyId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", url.PathEscape(r.apiKeyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiModelApiKeyUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("aiModelApiKeyUpdateRequest is required and must be specified")
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
	localVarPostBody = r.aiModelApiKeyUpdateRequest
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
