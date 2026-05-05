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
		CreateModelApiKey Create New AI Model API Key

		Create a new AI model API key for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param aiModelsApiKeyCreateRequest A request containing the name of the new API key.
		@return CreateModelApiKeyApiRequest
	*/
	CreateModelApiKey(ctx context.Context, groupId string, aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest) CreateModelApiKeyApiRequest
	/*
		CreateModelApiKey Create New AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateModelApiKeyApiParams - Parameters for the request
		@return CreateModelApiKeyApiRequest
	*/
	CreateModelApiKeyWithParams(ctx context.Context, args *CreateModelApiKeyApiParams) CreateModelApiKeyApiRequest

	// Method available only for mocking purposes
	CreateModelApiKeyExecute(r CreateModelApiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		DeleteModelApiKey Delete Existing AI Model API Key

		Delete an existing AI model API key in the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiKeyId The id of the API key to be deleted.
		@return DeleteModelApiKeyApiRequest
	*/
	DeleteModelApiKey(ctx context.Context, groupId string, apiKeyId string) DeleteModelApiKeyApiRequest
	/*
		DeleteModelApiKey Delete Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteModelApiKeyApiParams - Parameters for the request
		@return DeleteModelApiKeyApiRequest
	*/
	DeleteModelApiKeyWithParams(ctx context.Context, args *DeleteModelApiKeyApiParams) DeleteModelApiKeyApiRequest

	// Method available only for mocking purposes
	DeleteModelApiKeyExecute(r DeleteModelApiKeyApiRequest) (*http.Response, error)

	/*
		GetGroupModelKey Return Single AI Model API Key for One Group

		Retrieve a single AI model API key for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

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
	GetGroupModelKeyExecute(r GetGroupModelKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

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
	GetOrgModelKeyExecute(r GetOrgModelKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)

	/*
		ListGroupModelKeys Return AI Model API Keys for One Group

		Retrieve AI model API keys for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

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
	ListGroupModelKeysExecute(r ListGroupModelKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error)

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
	ListOrgModelKeysExecute(r ListOrgModelKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error)

	/*
		UpdateModelApiKey Update Existing AI Model API Key

		Update an existing AI model API key in the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiKeyId The id of the API key to be updated.
		@param aiModelsApiKeyUpdateRequest A request containing the new name for the API key.
		@return UpdateModelApiKeyApiRequest
	*/
	UpdateModelApiKey(ctx context.Context, groupId string, apiKeyId string, aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest) UpdateModelApiKeyApiRequest
	/*
		UpdateModelApiKey Update Existing AI Model API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateModelApiKeyApiParams - Parameters for the request
		@return UpdateModelApiKeyApiRequest
	*/
	UpdateModelApiKeyWithParams(ctx context.Context, args *UpdateModelApiKeyApiParams) UpdateModelApiKeyApiRequest

	// Method available only for mocking purposes
	UpdateModelApiKeyExecute(r UpdateModelApiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error)
}

// AIModelAPIKeysApiService AIModelAPIKeysApi service
type AIModelAPIKeysApiService service

type CreateModelApiKeyApiRequest struct {
	ctx                         context.Context
	ApiService                  AIModelAPIKeysApi
	groupId                     string
	aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest
}

type CreateModelApiKeyApiParams struct {
	GroupId                     string
	AiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest
}

func (a *AIModelAPIKeysApiService) CreateModelApiKeyWithParams(ctx context.Context, args *CreateModelApiKeyApiParams) CreateModelApiKeyApiRequest {
	return CreateModelApiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		aiModelsApiKeyCreateRequest: args.AiModelsApiKeyCreateRequest,
	}
}

func (r CreateModelApiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.CreateModelApiKeyExecute(r)
}

/*
CreateModelApiKey Create New AI Model API Key

Create a new AI model API key for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateModelApiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) CreateModelApiKey(ctx context.Context, groupId string, aiModelsApiKeyCreateRequest *AiModelsApiKeyCreateRequest) CreateModelApiKeyApiRequest {
	return CreateModelApiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		aiModelsApiKeyCreateRequest: aiModelsApiKeyCreateRequest,
	}
}

// CreateModelApiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) CreateModelApiKeyExecute(r CreateModelApiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.CreateModelApiKey")
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

type DeleteModelApiKeyApiRequest struct {
	ctx        context.Context
	ApiService AIModelAPIKeysApi
	groupId    string
	apiKeyId   string
}

type DeleteModelApiKeyApiParams struct {
	GroupId  string
	ApiKeyId string
}

func (a *AIModelAPIKeysApiService) DeleteModelApiKeyWithParams(ctx context.Context, args *DeleteModelApiKeyApiParams) DeleteModelApiKeyApiRequest {
	return DeleteModelApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiKeyId:   args.ApiKeyId,
	}
}

func (r DeleteModelApiKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteModelApiKeyExecute(r)
}

/*
DeleteModelApiKey Delete Existing AI Model API Key

Delete an existing AI model API key in the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be deleted.
	@return DeleteModelApiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) DeleteModelApiKey(ctx context.Context, groupId string, apiKeyId string) DeleteModelApiKeyApiRequest {
	return DeleteModelApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiKeyId:   apiKeyId,
	}
}

// DeleteModelApiKeyExecute executes the request
func (a *AIModelAPIKeysApiService) DeleteModelApiKeyExecute(r DeleteModelApiKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.DeleteModelApiKey")
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

func (r GetGroupModelKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.GetGroupModelKeyExecute(r)
}

/*
GetGroupModelKey Return Single AI Model API Key for One Group

Retrieve a single AI model API key for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

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
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) GetGroupModelKeyExecute(r GetGroupModelKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
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

func (r GetOrgModelKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
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
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) GetOrgModelKeyExecute(r GetOrgModelKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
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

func (r ListGroupModelKeysApiRequest) Execute() (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	return r.ApiService.ListGroupModelKeysExecute(r)
}

/*
ListGroupModelKeys Return AI Model API Keys for One Group

Retrieve AI model API keys for the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner, Project Model Owner, or Project Read Only role.

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
//	@return PaginatedAtlasAiModelsApiKeysResponse
func (a *AIModelAPIKeysApiService) ListGroupModelKeysExecute(r ListGroupModelKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsApiKeysResponse
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

func (r ListOrgModelKeysApiRequest) Execute() (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
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
//	@return PaginatedAtlasAiModelsApiKeysResponse
func (a *AIModelAPIKeysApiService) ListOrgModelKeysExecute(r ListOrgModelKeysApiRequest) (*PaginatedAtlasAiModelsApiKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasAiModelsApiKeysResponse
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

type UpdateModelApiKeyApiRequest struct {
	ctx                         context.Context
	ApiService                  AIModelAPIKeysApi
	groupId                     string
	apiKeyId                    string
	aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest
}

type UpdateModelApiKeyApiParams struct {
	GroupId                     string
	ApiKeyId                    string
	AiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest
}

func (a *AIModelAPIKeysApiService) UpdateModelApiKeyWithParams(ctx context.Context, args *UpdateModelApiKeyApiParams) UpdateModelApiKeyApiRequest {
	return UpdateModelApiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		apiKeyId:                    args.ApiKeyId,
		aiModelsApiKeyUpdateRequest: args.AiModelsApiKeyUpdateRequest,
	}
}

func (r UpdateModelApiKeyApiRequest) Execute() (*AiModelsApiKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateModelApiKeyExecute(r)
}

/*
UpdateModelApiKey Update Existing AI Model API Key

Update an existing AI model API key in the given group. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Model Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiKeyId The id of the API key to be updated.
	@return UpdateModelApiKeyApiRequest
*/
func (a *AIModelAPIKeysApiService) UpdateModelApiKey(ctx context.Context, groupId string, apiKeyId string, aiModelsApiKeyUpdateRequest *AiModelsApiKeyUpdateRequest) UpdateModelApiKeyApiRequest {
	return UpdateModelApiKeyApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		apiKeyId:                    apiKeyId,
		aiModelsApiKeyUpdateRequest: aiModelsApiKeyUpdateRequest,
	}
}

// UpdateModelApiKeyExecute executes the request
//
//	@return AiModelsApiKeyResponse
func (a *AIModelAPIKeysApiService) UpdateModelApiKeyExecute(r UpdateModelApiKeyApiRequest) (*AiModelsApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AiModelsApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIModelAPIKeysApiService.UpdateModelApiKey")
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
