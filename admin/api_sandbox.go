// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SandboxApi interface {

	/*
		CreateOrgSandboxConfig Create One Sandbox Configuration for an Organization

		Create one sandbox configuration for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param sandboxConfigRequest Settings to configure the sandbox feature for an organization.
		@return CreateOrgSandboxConfigApiRequest
	*/
	CreateOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigRequest *SandboxConfigRequest) CreateOrgSandboxConfigApiRequest
	/*
		CreateOrgSandboxConfig Create One Sandbox Configuration for an Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrgSandboxConfigApiParams - Parameters for the request
		@return CreateOrgSandboxConfigApiRequest
	*/
	CreateOrgSandboxConfigWithParams(ctx context.Context, args *CreateOrgSandboxConfigApiParams) CreateOrgSandboxConfigApiRequest

	// Method available only for mocking purposes
	CreateOrgSandboxConfigExecute(r CreateOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error)

	/*
		DeleteOrgSandboxConfig Delete the Default Sandbox Configuration and Disable Sandbox

		Delete the default sandbox configuration and disable sandbox for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
		@return DeleteOrgSandboxConfigApiRequest
	*/
	DeleteOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string) DeleteOrgSandboxConfigApiRequest
	/*
		DeleteOrgSandboxConfig Delete the Default Sandbox Configuration and Disable Sandbox


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrgSandboxConfigApiParams - Parameters for the request
		@return DeleteOrgSandboxConfigApiRequest
	*/
	DeleteOrgSandboxConfigWithParams(ctx context.Context, args *DeleteOrgSandboxConfigApiParams) DeleteOrgSandboxConfigApiRequest

	// Method available only for mocking purposes
	DeleteOrgSandboxConfigExecute(r DeleteOrgSandboxConfigApiRequest) (*http.Response, error)

	/*
		GenerateSandboxClusterDescription Return Cluster Description from Sandbox Template Configuration

		Return cluster description from sandbox template configuration for M0 and M10 clusters only. Results from this endpoint can be put directly into the /clusters endpoint. View https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-creategroupcluster for more information. For Flex clusters, please visit https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-creategroupflexcluster for an example.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
		@return GenerateSandboxClusterDescriptionApiRequest
	*/
	GenerateSandboxClusterDescription(ctx context.Context, groupId string, sandboxConfigId string) GenerateSandboxClusterDescriptionApiRequest
	/*
		GenerateSandboxClusterDescription Return Cluster Description from Sandbox Template Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GenerateSandboxClusterDescriptionApiParams - Parameters for the request
		@return GenerateSandboxClusterDescriptionApiRequest
	*/
	GenerateSandboxClusterDescriptionWithParams(ctx context.Context, args *GenerateSandboxClusterDescriptionApiParams) GenerateSandboxClusterDescriptionApiRequest

	// Method available only for mocking purposes
	GenerateSandboxClusterDescriptionExecute(r GenerateSandboxClusterDescriptionApiRequest) (*ClusterDescription20240805, *http.Response, error)

	/*
		GetOrgSandboxConfig Return Default Sandbox Configuration for One Organization

		Return the default sandbox configuration for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
		@return GetOrgSandboxConfigApiRequest
	*/
	GetOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string) GetOrgSandboxConfigApiRequest
	/*
		GetOrgSandboxConfig Return Default Sandbox Configuration for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgSandboxConfigApiParams - Parameters for the request
		@return GetOrgSandboxConfigApiRequest
	*/
	GetOrgSandboxConfigWithParams(ctx context.Context, args *GetOrgSandboxConfigApiParams) GetOrgSandboxConfigApiRequest

	// Method available only for mocking purposes
	GetOrgSandboxConfigExecute(r GetOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error)

	/*
		ListOrgSandboxConfig Return Sandbox Configuration IDs for an Organization

		Return an array of sandbox configuration IDs for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrgSandboxConfigApiRequest
	*/
	ListOrgSandboxConfig(ctx context.Context, orgId string) ListOrgSandboxConfigApiRequest
	/*
		ListOrgSandboxConfig Return Sandbox Configuration IDs for an Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgSandboxConfigApiParams - Parameters for the request
		@return ListOrgSandboxConfigApiRequest
	*/
	ListOrgSandboxConfigWithParams(ctx context.Context, args *ListOrgSandboxConfigApiParams) ListOrgSandboxConfigApiRequest

	// Method available only for mocking purposes
	ListOrgSandboxConfigExecute(r ListOrgSandboxConfigApiRequest) (*PaginatedSandboxConfigs, *http.Response, error)

	/*
		UpdateOrgSandboxConfig Update an Existing Sandbox Configuration

		Update an existing sandbox configuration for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
		@param sandboxConfigUpdateRequest Settings to update the sandbox configuration for an organization.
		@return UpdateOrgSandboxConfigApiRequest
	*/
	UpdateOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string, sandboxConfigUpdateRequest *SandboxConfigUpdateRequest) UpdateOrgSandboxConfigApiRequest
	/*
		UpdateOrgSandboxConfig Update an Existing Sandbox Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrgSandboxConfigApiParams - Parameters for the request
		@return UpdateOrgSandboxConfigApiRequest
	*/
	UpdateOrgSandboxConfigWithParams(ctx context.Context, args *UpdateOrgSandboxConfigApiParams) UpdateOrgSandboxConfigApiRequest

	// Method available only for mocking purposes
	UpdateOrgSandboxConfigExecute(r UpdateOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error)
}

// SandboxApiService SandboxApi service
type SandboxApiService service

type CreateOrgSandboxConfigApiRequest struct {
	ctx                  context.Context
	ApiService           SandboxApi
	orgId                string
	sandboxConfigRequest *SandboxConfigRequest
}

type CreateOrgSandboxConfigApiParams struct {
	OrgId                string
	SandboxConfigRequest *SandboxConfigRequest
}

func (a *SandboxApiService) CreateOrgSandboxConfigWithParams(ctx context.Context, args *CreateOrgSandboxConfigApiParams) CreateOrgSandboxConfigApiRequest {
	return CreateOrgSandboxConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                args.OrgId,
		sandboxConfigRequest: args.SandboxConfigRequest,
	}
}

func (r CreateOrgSandboxConfigApiRequest) Execute() (*SandboxConfigResponse, *http.Response, error) {
	return r.ApiService.CreateOrgSandboxConfigExecute(r)
}

/*
CreateOrgSandboxConfig Create One Sandbox Configuration for an Organization

Create one sandbox configuration for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateOrgSandboxConfigApiRequest
*/
func (a *SandboxApiService) CreateOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigRequest *SandboxConfigRequest) CreateOrgSandboxConfigApiRequest {
	return CreateOrgSandboxConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                orgId,
		sandboxConfigRequest: sandboxConfigRequest,
	}
}

// CreateOrgSandboxConfigExecute executes the request
//
//	@return SandboxConfigResponse
func (a *SandboxApiService) CreateOrgSandboxConfigExecute(r CreateOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SandboxConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.CreateOrgSandboxConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/sandboxConfig"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sandboxConfigRequest == nil {
		return localVarReturnValue, nil, reportError("sandboxConfigRequest is required and must be specified")
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
	localVarPostBody = r.sandboxConfigRequest
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

type DeleteOrgSandboxConfigApiRequest struct {
	ctx             context.Context
	ApiService      SandboxApi
	orgId           string
	sandboxConfigId string
}

type DeleteOrgSandboxConfigApiParams struct {
	OrgId           string
	SandboxConfigId string
}

func (a *SandboxApiService) DeleteOrgSandboxConfigWithParams(ctx context.Context, args *DeleteOrgSandboxConfigApiParams) DeleteOrgSandboxConfigApiRequest {
	return DeleteOrgSandboxConfigApiRequest{
		ApiService:      a,
		ctx:             ctx,
		orgId:           args.OrgId,
		sandboxConfigId: args.SandboxConfigId,
	}
}

func (r DeleteOrgSandboxConfigApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrgSandboxConfigExecute(r)
}

/*
DeleteOrgSandboxConfig Delete the Default Sandbox Configuration and Disable Sandbox

Delete the default sandbox configuration and disable sandbox for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
	@return DeleteOrgSandboxConfigApiRequest
*/
func (a *SandboxApiService) DeleteOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string) DeleteOrgSandboxConfigApiRequest {
	return DeleteOrgSandboxConfigApiRequest{
		ApiService:      a,
		ctx:             ctx,
		orgId:           orgId,
		sandboxConfigId: sandboxConfigId,
	}
}

// DeleteOrgSandboxConfigExecute executes the request
func (a *SandboxApiService) DeleteOrgSandboxConfigExecute(r DeleteOrgSandboxConfigApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.DeleteOrgSandboxConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/sandboxConfig/{sandboxConfigId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.sandboxConfigId == "" {
		return nil, reportError("sandboxConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"sandboxConfigId"+"}", url.PathEscape(r.sandboxConfigId), -1)

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

type GenerateSandboxClusterDescriptionApiRequest struct {
	ctx             context.Context
	ApiService      SandboxApi
	groupId         string
	sandboxConfigId string
}

type GenerateSandboxClusterDescriptionApiParams struct {
	GroupId         string
	SandboxConfigId string
}

func (a *SandboxApiService) GenerateSandboxClusterDescriptionWithParams(ctx context.Context, args *GenerateSandboxClusterDescriptionApiParams) GenerateSandboxClusterDescriptionApiRequest {
	return GenerateSandboxClusterDescriptionApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		sandboxConfigId: args.SandboxConfigId,
	}
}

func (r GenerateSandboxClusterDescriptionApiRequest) Execute() (*ClusterDescription20240805, *http.Response, error) {
	return r.ApiService.GenerateSandboxClusterDescriptionExecute(r)
}

/*
GenerateSandboxClusterDescription Return Cluster Description from Sandbox Template Configuration

Return cluster description from sandbox template configuration for M0 and M10 clusters only. Results from this endpoint can be put directly into the /clusters endpoint. View https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-creategroupcluster for more information. For Flex clusters, please visit https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-creategroupflexcluster for an example.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
	@return GenerateSandboxClusterDescriptionApiRequest
*/
func (a *SandboxApiService) GenerateSandboxClusterDescription(ctx context.Context, groupId string, sandboxConfigId string) GenerateSandboxClusterDescriptionApiRequest {
	return GenerateSandboxClusterDescriptionApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		sandboxConfigId: sandboxConfigId,
	}
}

// GenerateSandboxClusterDescriptionExecute executes the request
//
//	@return ClusterDescription20240805
func (a *SandboxApiService) GenerateSandboxClusterDescriptionExecute(r GenerateSandboxClusterDescriptionApiRequest) (*ClusterDescription20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescription20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.GenerateSandboxClusterDescription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/sandbox/{sandboxConfigId}:generateClusterDescription"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.sandboxConfigId == "" {
		return localVarReturnValue, nil, reportError("sandboxConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"sandboxConfigId"+"}", url.PathEscape(r.sandboxConfigId), -1)

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

type GetOrgSandboxConfigApiRequest struct {
	ctx             context.Context
	ApiService      SandboxApi
	orgId           string
	sandboxConfigId string
}

type GetOrgSandboxConfigApiParams struct {
	OrgId           string
	SandboxConfigId string
}

func (a *SandboxApiService) GetOrgSandboxConfigWithParams(ctx context.Context, args *GetOrgSandboxConfigApiParams) GetOrgSandboxConfigApiRequest {
	return GetOrgSandboxConfigApiRequest{
		ApiService:      a,
		ctx:             ctx,
		orgId:           args.OrgId,
		sandboxConfigId: args.SandboxConfigId,
	}
}

func (r GetOrgSandboxConfigApiRequest) Execute() (*SandboxConfigResponse, *http.Response, error) {
	return r.ApiService.GetOrgSandboxConfigExecute(r)
}

/*
GetOrgSandboxConfig Return Default Sandbox Configuration for One Organization

Return the default sandbox configuration for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
	@return GetOrgSandboxConfigApiRequest
*/
func (a *SandboxApiService) GetOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string) GetOrgSandboxConfigApiRequest {
	return GetOrgSandboxConfigApiRequest{
		ApiService:      a,
		ctx:             ctx,
		orgId:           orgId,
		sandboxConfigId: sandboxConfigId,
	}
}

// GetOrgSandboxConfigExecute executes the request
//
//	@return SandboxConfigResponse
func (a *SandboxApiService) GetOrgSandboxConfigExecute(r GetOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SandboxConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.GetOrgSandboxConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/sandboxConfig/{sandboxConfigId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.sandboxConfigId == "" {
		return localVarReturnValue, nil, reportError("sandboxConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"sandboxConfigId"+"}", url.PathEscape(r.sandboxConfigId), -1)

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

type ListOrgSandboxConfigApiRequest struct {
	ctx          context.Context
	ApiService   SandboxApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type ListOrgSandboxConfigApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *SandboxApiService) ListOrgSandboxConfigWithParams(ctx context.Context, args *ListOrgSandboxConfigApiParams) ListOrgSandboxConfigApiRequest {
	return ListOrgSandboxConfigApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgSandboxConfigApiRequest) ItemsPerPage(itemsPerPage int) ListOrgSandboxConfigApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgSandboxConfigApiRequest) PageNum(pageNum int) ListOrgSandboxConfigApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgSandboxConfigApiRequest) Execute() (*PaginatedSandboxConfigs, *http.Response, error) {
	return r.ApiService.ListOrgSandboxConfigExecute(r)
}

/*
ListOrgSandboxConfig Return Sandbox Configuration IDs for an Organization

Return an array of sandbox configuration IDs for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrgSandboxConfigApiRequest
*/
func (a *SandboxApiService) ListOrgSandboxConfig(ctx context.Context, orgId string) ListOrgSandboxConfigApiRequest {
	return ListOrgSandboxConfigApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrgSandboxConfigExecute executes the request
//
//	@return PaginatedSandboxConfigs
func (a *SandboxApiService) ListOrgSandboxConfigExecute(r ListOrgSandboxConfigApiRequest) (*PaginatedSandboxConfigs, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedSandboxConfigs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.ListOrgSandboxConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/sandboxConfig"
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

type UpdateOrgSandboxConfigApiRequest struct {
	ctx                        context.Context
	ApiService                 SandboxApi
	orgId                      string
	sandboxConfigId            string
	sandboxConfigUpdateRequest *SandboxConfigUpdateRequest
}

type UpdateOrgSandboxConfigApiParams struct {
	OrgId                      string
	SandboxConfigId            string
	SandboxConfigUpdateRequest *SandboxConfigUpdateRequest
}

func (a *SandboxApiService) UpdateOrgSandboxConfigWithParams(ctx context.Context, args *UpdateOrgSandboxConfigApiParams) UpdateOrgSandboxConfigApiRequest {
	return UpdateOrgSandboxConfigApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		orgId:                      args.OrgId,
		sandboxConfigId:            args.SandboxConfigId,
		sandboxConfigUpdateRequest: args.SandboxConfigUpdateRequest,
	}
}

func (r UpdateOrgSandboxConfigApiRequest) Execute() (*SandboxConfigResponse, *http.Response, error) {
	return r.ApiService.UpdateOrgSandboxConfigExecute(r)
}

/*
UpdateOrgSandboxConfig Update an Existing Sandbox Configuration

Update an existing sandbox configuration for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param sandboxConfigId Unique 24-hexadecimal digit string that identifies the sandbox configuration.
	@return UpdateOrgSandboxConfigApiRequest
*/
func (a *SandboxApiService) UpdateOrgSandboxConfig(ctx context.Context, orgId string, sandboxConfigId string, sandboxConfigUpdateRequest *SandboxConfigUpdateRequest) UpdateOrgSandboxConfigApiRequest {
	return UpdateOrgSandboxConfigApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		orgId:                      orgId,
		sandboxConfigId:            sandboxConfigId,
		sandboxConfigUpdateRequest: sandboxConfigUpdateRequest,
	}
}

// UpdateOrgSandboxConfigExecute executes the request
//
//	@return SandboxConfigResponse
func (a *SandboxApiService) UpdateOrgSandboxConfigExecute(r UpdateOrgSandboxConfigApiRequest) (*SandboxConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SandboxConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxApiService.UpdateOrgSandboxConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/sandboxConfig/{sandboxConfigId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.sandboxConfigId == "" {
		return localVarReturnValue, nil, reportError("sandboxConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"sandboxConfigId"+"}", url.PathEscape(r.sandboxConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sandboxConfigUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("sandboxConfigUpdateRequest is required and must be specified")
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
	localVarPostBody = r.sandboxConfigUpdateRequest
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
