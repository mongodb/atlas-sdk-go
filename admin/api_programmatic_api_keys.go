// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ProgrammaticAPIKeysApi interface {

	/*
		AddProjectApiKey Assign One Organization API Key to One Project

		Assigns the specified organization API key to the specified project. Users with the Project Owner role in the project associated with the API key can then use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.
		@param userAccessRoleAssignment Organization API key to be assigned to the specified project.
		@return AddProjectApiKeyApiRequest
	*/
	AddProjectApiKey(ctx context.Context, groupId string, apiUserId string, userAccessRoleAssignment *[]UserAccessRoleAssignment) AddProjectApiKeyApiRequest
	/*
		AddProjectApiKey Assign One Organization API Key to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddProjectApiKeyApiParams - Parameters for the request
		@return AddProjectApiKeyApiRequest
	*/
	AddProjectApiKeyWithParams(ctx context.Context, args *AddProjectApiKeyApiParams) AddProjectApiKeyApiRequest

	// Method available only for mocking purposes
	AddProjectApiKeyExecute(r AddProjectApiKeyApiRequest) (*http.Response, error)

	/*
		CreateApiKey Create One Organization API Key

		Creates one API key for the specified organization. An organization API key grants programmatic access to an organization. You can't use the API key to log into the console. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param createAtlasOrganizationApiKey Organization API Key to be created.
		@return CreateApiKeyApiRequest
	*/
	CreateApiKey(ctx context.Context, orgId string, createAtlasOrganizationApiKey *CreateAtlasOrganizationApiKey) CreateApiKeyApiRequest
	/*
		CreateApiKey Create One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateApiKeyApiParams - Parameters for the request
		@return CreateApiKeyApiRequest
	*/
	CreateApiKeyWithParams(ctx context.Context, args *CreateApiKeyApiParams) CreateApiKeyApiRequest

	// Method available only for mocking purposes
	CreateApiKeyExecute(r CreateApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error)

	/*
		CreateApiKeyAccessList Create One Access List Entry for One Organization API Key

		Creates the access list entries for the specified organization API key. Resources require all API requests originate from IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.
		@param userAccessListRequest Access list entries to be created for the specified organization API key.
		@return CreateApiKeyAccessListApiRequest
	*/
	CreateApiKeyAccessList(ctx context.Context, orgId string, apiUserId string, userAccessListRequest *[]UserAccessListRequest) CreateApiKeyAccessListApiRequest
	/*
		CreateApiKeyAccessList Create One Access List Entry for One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateApiKeyAccessListApiParams - Parameters for the request
		@return CreateApiKeyAccessListApiRequest
	*/
	CreateApiKeyAccessListWithParams(ctx context.Context, args *CreateApiKeyAccessListApiParams) CreateApiKeyAccessListApiRequest

	// Method available only for mocking purposes
	CreateApiKeyAccessListExecute(r CreateApiKeyAccessListApiRequest) (*PaginatedApiUserAccessListResponse, *http.Response, error)

	/*
		CreateProjectApiKey Create and Assign One Organization API Key to One Project

		Creates and assigns the specified organization API key to the specified project. Users with the Project Owner role in the project associated with the API key can use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param createAtlasProjectApiKey Organization API key to be created and assigned to the specified project.
		@return CreateProjectApiKeyApiRequest
	*/
	CreateProjectApiKey(ctx context.Context, groupId string, createAtlasProjectApiKey *CreateAtlasProjectApiKey) CreateProjectApiKeyApiRequest
	/*
		CreateProjectApiKey Create and Assign One Organization API Key to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectApiKeyApiParams - Parameters for the request
		@return CreateProjectApiKeyApiRequest
	*/
	CreateProjectApiKeyWithParams(ctx context.Context, args *CreateProjectApiKeyApiParams) CreateProjectApiKeyApiRequest

	// Method available only for mocking purposes
	CreateProjectApiKeyExecute(r CreateProjectApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error)

	/*
		DeleteApiKey Remove One Organization API Key

		Removes one organization API key from the specified organization. When you remove an API key from an organization, MongoDB Cloud also removes that key from any projects that use that key. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key.
		@return DeleteApiKeyApiRequest
	*/
	DeleteApiKey(ctx context.Context, orgId string, apiUserId string) DeleteApiKeyApiRequest
	/*
		DeleteApiKey Remove One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteApiKeyApiParams - Parameters for the request
		@return DeleteApiKeyApiRequest
	*/
	DeleteApiKeyWithParams(ctx context.Context, args *DeleteApiKeyApiParams) DeleteApiKeyApiRequest

	// Method available only for mocking purposes
	DeleteApiKeyExecute(r DeleteApiKeyApiRequest) (*http.Response, error)

	/*
		DeleteApiKeyAccessListEntry Remove One Access List Entry for One Organization API Key

		Removes the specified access list entry from the specified organization API key. Resources require all API requests originate from the IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Read Write role. In addition, you cannot remove the requesting IP address from the requesting organization API key.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.
		@param ipAddress One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
		@return DeleteApiKeyAccessListEntryApiRequest
	*/
	DeleteApiKeyAccessListEntry(ctx context.Context, orgId string, apiUserId string, ipAddress string) DeleteApiKeyAccessListEntryApiRequest
	/*
		DeleteApiKeyAccessListEntry Remove One Access List Entry for One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteApiKeyAccessListEntryApiParams - Parameters for the request
		@return DeleteApiKeyAccessListEntryApiRequest
	*/
	DeleteApiKeyAccessListEntryWithParams(ctx context.Context, args *DeleteApiKeyAccessListEntryApiParams) DeleteApiKeyAccessListEntryApiRequest

	// Method available only for mocking purposes
	DeleteApiKeyAccessListEntryExecute(r DeleteApiKeyAccessListEntryApiRequest) (*http.Response, error)

	/*
		GetApiKey Return One Organization API Key

		Returns one organization API key. The organization API keys grant programmatic access to an organization. You can't use the API key to log into MongoDB Cloud through the user interface. To use this resource, the requesting Service Account or API Key must have the  Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.
		@return GetApiKeyApiRequest
	*/
	GetApiKey(ctx context.Context, orgId string, apiUserId string) GetApiKeyApiRequest
	/*
		GetApiKey Return One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetApiKeyApiParams - Parameters for the request
		@return GetApiKeyApiRequest
	*/
	GetApiKeyWithParams(ctx context.Context, args *GetApiKeyApiParams) GetApiKeyApiRequest

	// Method available only for mocking purposes
	GetApiKeyExecute(r GetApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error)

	/*
		GetApiKeyAccessList Return One Access List Entry for One Organization API Key

		Returns one access list entry for the specified organization API key. Resources require  all API requests originate from IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param ipAddress One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.
		@return GetApiKeyAccessListApiRequest
	*/
	GetApiKeyAccessList(ctx context.Context, orgId string, ipAddress string, apiUserId string) GetApiKeyAccessListApiRequest
	/*
		GetApiKeyAccessList Return One Access List Entry for One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetApiKeyAccessListApiParams - Parameters for the request
		@return GetApiKeyAccessListApiRequest
	*/
	GetApiKeyAccessListWithParams(ctx context.Context, args *GetApiKeyAccessListApiParams) GetApiKeyAccessListApiRequest

	// Method available only for mocking purposes
	GetApiKeyAccessListExecute(r GetApiKeyAccessListApiRequest) (*UserAccessListResponse, *http.Response, error)

	/*
		ListApiKeyAccessListsEntries Return All Access List Entries for One Organization API Key

		Returns all access list entries that you configured for the specified organization API key. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.
		@return ListApiKeyAccessListsEntriesApiRequest
	*/
	ListApiKeyAccessListsEntries(ctx context.Context, orgId string, apiUserId string) ListApiKeyAccessListsEntriesApiRequest
	/*
		ListApiKeyAccessListsEntries Return All Access List Entries for One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListApiKeyAccessListsEntriesApiParams - Parameters for the request
		@return ListApiKeyAccessListsEntriesApiRequest
	*/
	ListApiKeyAccessListsEntriesWithParams(ctx context.Context, args *ListApiKeyAccessListsEntriesApiParams) ListApiKeyAccessListsEntriesApiRequest

	// Method available only for mocking purposes
	ListApiKeyAccessListsEntriesExecute(r ListApiKeyAccessListsEntriesApiRequest) (*PaginatedApiUserAccessListResponse, *http.Response, error)

	/*
		ListApiKeys Return All Organization API Keys

		Returns all organization API keys for the specified organization. The organization API keys grant programmatic access to an organization. You can't use the API key to log into MongoDB Cloud through the console. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListApiKeysApiRequest
	*/
	ListApiKeys(ctx context.Context, orgId string) ListApiKeysApiRequest
	/*
		ListApiKeys Return All Organization API Keys


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListApiKeysApiParams - Parameters for the request
		@return ListApiKeysApiRequest
	*/
	ListApiKeysWithParams(ctx context.Context, args *ListApiKeysApiParams) ListApiKeysApiRequest

	// Method available only for mocking purposes
	ListApiKeysExecute(r ListApiKeysApiRequest) (*PaginatedApiApiUser, *http.Response, error)

	/*
		ListProjectApiKeys Return All Organization API Keys Assigned to One Project

		Returns all organization API keys that you assigned to the specified project. Users with the Project Owner role in the project associated with the API key can use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectApiKeysApiRequest
	*/
	ListProjectApiKeys(ctx context.Context, groupId string) ListProjectApiKeysApiRequest
	/*
		ListProjectApiKeys Return All Organization API Keys Assigned to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectApiKeysApiParams - Parameters for the request
		@return ListProjectApiKeysApiRequest
	*/
	ListProjectApiKeysWithParams(ctx context.Context, args *ListProjectApiKeysApiParams) ListProjectApiKeysApiRequest

	// Method available only for mocking purposes
	ListProjectApiKeysExecute(r ListProjectApiKeysApiRequest) (*PaginatedApiApiUser, *http.Response, error)

	/*
		RemoveProjectApiKey Unassign One Organization API Key from One Project

		Removes one organization API key from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
		@return RemoveProjectApiKeyApiRequest
	*/
	RemoveProjectApiKey(ctx context.Context, groupId string, apiUserId string) RemoveProjectApiKeyApiRequest
	/*
		RemoveProjectApiKey Unassign One Organization API Key from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveProjectApiKeyApiParams - Parameters for the request
		@return RemoveProjectApiKeyApiRequest
	*/
	RemoveProjectApiKeyWithParams(ctx context.Context, args *RemoveProjectApiKeyApiParams) RemoveProjectApiKeyApiRequest

	// Method available only for mocking purposes
	RemoveProjectApiKeyExecute(r RemoveProjectApiKeyApiRequest) (*http.Response, error)

	/*
		UpdateApiKey Update One Organization API Key

		Updates one organization API key in the specified organization. The organization API keys  grant programmatic access to an organization. To use this resource, the requesting  API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.
		@param updateAtlasOrganizationApiKey Organization API key to be updated. This request requires a minimum of one of the two body parameters.
		@return UpdateApiKeyApiRequest
	*/
	UpdateApiKey(ctx context.Context, orgId string, apiUserId string, updateAtlasOrganizationApiKey *UpdateAtlasOrganizationApiKey) UpdateApiKeyApiRequest
	/*
		UpdateApiKey Update One Organization API Key


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateApiKeyApiParams - Parameters for the request
		@return UpdateApiKeyApiRequest
	*/
	UpdateApiKeyWithParams(ctx context.Context, args *UpdateApiKeyApiParams) UpdateApiKeyApiRequest

	// Method available only for mocking purposes
	UpdateApiKeyExecute(r UpdateApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error)

	/*
		UpdateApiKeyRoles Update Organization API Key Roles for One Project

		Updates the roles of the organization API key that you specify for the project that you specify. You must specify at least one valid role for the project. The application removes any roles that you do not include in this request if they were previously set in the organization API key that you specify for the project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
		@param updateAtlasProjectApiKey Organization API Key to be updated. This request requires a minimum of one of the two body parameters.
		@return UpdateApiKeyRolesApiRequest
	*/
	UpdateApiKeyRoles(ctx context.Context, groupId string, apiUserId string, updateAtlasProjectApiKey *UpdateAtlasProjectApiKey) UpdateApiKeyRolesApiRequest
	/*
		UpdateApiKeyRoles Update Organization API Key Roles for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateApiKeyRolesApiParams - Parameters for the request
		@return UpdateApiKeyRolesApiRequest
	*/
	UpdateApiKeyRolesWithParams(ctx context.Context, args *UpdateApiKeyRolesApiParams) UpdateApiKeyRolesApiRequest

	// Method available only for mocking purposes
	UpdateApiKeyRolesExecute(r UpdateApiKeyRolesApiRequest) (*ApiKeyUserDetails, *http.Response, error)
}

// ProgrammaticAPIKeysApiService ProgrammaticAPIKeysApi service
type ProgrammaticAPIKeysApiService service

type AddProjectApiKeyApiRequest struct {
	ctx                      context.Context
	ApiService               ProgrammaticAPIKeysApi
	groupId                  string
	apiUserId                string
	userAccessRoleAssignment *[]UserAccessRoleAssignment
}

type AddProjectApiKeyApiParams struct {
	GroupId                  string
	ApiUserId                string
	UserAccessRoleAssignment *[]UserAccessRoleAssignment
}

func (a *ProgrammaticAPIKeysApiService) AddProjectApiKeyWithParams(ctx context.Context, args *AddProjectApiKeyApiParams) AddProjectApiKeyApiRequest {
	return AddProjectApiKeyApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		apiUserId:                args.ApiUserId,
		userAccessRoleAssignment: args.UserAccessRoleAssignment,
	}
}

func (r AddProjectApiKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddProjectApiKeyExecute(r)
}

/*
AddProjectApiKey Assign One Organization API Key to One Project

Assigns the specified organization API key to the specified project. Users with the Project Owner role in the project associated with the API key can then use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.
	@return AddProjectApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) AddProjectApiKey(ctx context.Context, groupId string, apiUserId string, userAccessRoleAssignment *[]UserAccessRoleAssignment) AddProjectApiKeyApiRequest {
	return AddProjectApiKeyApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		apiUserId:                apiUserId,
		userAccessRoleAssignment: userAccessRoleAssignment,
	}
}

// AddProjectApiKeyExecute executes the request
func (a *ProgrammaticAPIKeysApiService) AddProjectApiKeyExecute(r AddProjectApiKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.AddProjectApiKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiUserId == "" {
		return nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userAccessRoleAssignment == nil {
		return nil, reportError("userAccessRoleAssignment is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userAccessRoleAssignment
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

type CreateApiKeyApiRequest struct {
	ctx                           context.Context
	ApiService                    ProgrammaticAPIKeysApi
	orgId                         string
	createAtlasOrganizationApiKey *CreateAtlasOrganizationApiKey
}

type CreateApiKeyApiParams struct {
	OrgId                         string
	CreateAtlasOrganizationApiKey *CreateAtlasOrganizationApiKey
}

func (a *ProgrammaticAPIKeysApiService) CreateApiKeyWithParams(ctx context.Context, args *CreateApiKeyApiParams) CreateApiKeyApiRequest {
	return CreateApiKeyApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		createAtlasOrganizationApiKey: args.CreateAtlasOrganizationApiKey,
	}
}

func (r CreateApiKeyApiRequest) Execute() (*ApiKeyUserDetails, *http.Response, error) {
	return r.ApiService.CreateApiKeyExecute(r)
}

/*
CreateApiKey Create One Organization API Key

Creates one API key for the specified organization. An organization API key grants programmatic access to an organization. You can't use the API key to log into the console. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) CreateApiKey(ctx context.Context, orgId string, createAtlasOrganizationApiKey *CreateAtlasOrganizationApiKey) CreateApiKeyApiRequest {
	return CreateApiKeyApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		createAtlasOrganizationApiKey: createAtlasOrganizationApiKey,
	}
}

// CreateApiKeyExecute executes the request
//
//	@return ApiKeyUserDetails
func (a *ProgrammaticAPIKeysApiService) CreateApiKeyExecute(r CreateApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiKeyUserDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.CreateApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAtlasOrganizationApiKey == nil {
		return localVarReturnValue, nil, reportError("createAtlasOrganizationApiKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createAtlasOrganizationApiKey
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

type CreateApiKeyAccessListApiRequest struct {
	ctx                   context.Context
	ApiService            ProgrammaticAPIKeysApi
	orgId                 string
	apiUserId             string
	userAccessListRequest *[]UserAccessListRequest
	includeCount          *bool
	itemsPerPage          *int
	pageNum               *int
}

type CreateApiKeyAccessListApiParams struct {
	OrgId                 string
	ApiUserId             string
	UserAccessListRequest *[]UserAccessListRequest
	IncludeCount          *bool
	ItemsPerPage          *int
	PageNum               *int
}

func (a *ProgrammaticAPIKeysApiService) CreateApiKeyAccessListWithParams(ctx context.Context, args *CreateApiKeyAccessListApiParams) CreateApiKeyAccessListApiRequest {
	return CreateApiKeyAccessListApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		orgId:                 args.OrgId,
		apiUserId:             args.ApiUserId,
		userAccessListRequest: args.UserAccessListRequest,
		includeCount:          args.IncludeCount,
		itemsPerPage:          args.ItemsPerPage,
		pageNum:               args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r CreateApiKeyAccessListApiRequest) IncludeCount(includeCount bool) CreateApiKeyAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r CreateApiKeyAccessListApiRequest) ItemsPerPage(itemsPerPage int) CreateApiKeyAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r CreateApiKeyAccessListApiRequest) PageNum(pageNum int) CreateApiKeyAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r CreateApiKeyAccessListApiRequest) Execute() (*PaginatedApiUserAccessListResponse, *http.Response, error) {
	return r.ApiService.CreateApiKeyAccessListExecute(r)
}

/*
CreateApiKeyAccessList Create One Access List Entry for One Organization API Key

Creates the access list entries for the specified organization API key. Resources require all API requests originate from IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.
	@return CreateApiKeyAccessListApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) CreateApiKeyAccessList(ctx context.Context, orgId string, apiUserId string, userAccessListRequest *[]UserAccessListRequest) CreateApiKeyAccessListApiRequest {
	return CreateApiKeyAccessListApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		orgId:                 orgId,
		apiUserId:             apiUserId,
		userAccessListRequest: userAccessListRequest,
	}
}

// CreateApiKeyAccessListExecute executes the request
//
//	@return PaginatedApiUserAccessListResponse
func (a *ProgrammaticAPIKeysApiService) CreateApiKeyAccessListExecute(r CreateApiKeyAccessListApiRequest) (*PaginatedApiUserAccessListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiUserAccessListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.CreateApiKeyAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userAccessListRequest == nil {
		return localVarReturnValue, nil, reportError("userAccessListRequest is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userAccessListRequest
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

type CreateProjectApiKeyApiRequest struct {
	ctx                      context.Context
	ApiService               ProgrammaticAPIKeysApi
	groupId                  string
	createAtlasProjectApiKey *CreateAtlasProjectApiKey
}

type CreateProjectApiKeyApiParams struct {
	GroupId                  string
	CreateAtlasProjectApiKey *CreateAtlasProjectApiKey
}

func (a *ProgrammaticAPIKeysApiService) CreateProjectApiKeyWithParams(ctx context.Context, args *CreateProjectApiKeyApiParams) CreateProjectApiKeyApiRequest {
	return CreateProjectApiKeyApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		createAtlasProjectApiKey: args.CreateAtlasProjectApiKey,
	}
}

func (r CreateProjectApiKeyApiRequest) Execute() (*ApiKeyUserDetails, *http.Response, error) {
	return r.ApiService.CreateProjectApiKeyExecute(r)
}

/*
CreateProjectApiKey Create and Assign One Organization API Key to One Project

Creates and assigns the specified organization API key to the specified project. Users with the Project Owner role in the project associated with the API key can use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) CreateProjectApiKey(ctx context.Context, groupId string, createAtlasProjectApiKey *CreateAtlasProjectApiKey) CreateProjectApiKeyApiRequest {
	return CreateProjectApiKeyApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		createAtlasProjectApiKey: createAtlasProjectApiKey,
	}
}

// CreateProjectApiKeyExecute executes the request
//
//	@return ApiKeyUserDetails
func (a *ProgrammaticAPIKeysApiService) CreateProjectApiKeyExecute(r CreateProjectApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiKeyUserDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.CreateProjectApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/apiKeys"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAtlasProjectApiKey == nil {
		return localVarReturnValue, nil, reportError("createAtlasProjectApiKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createAtlasProjectApiKey
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

type DeleteApiKeyApiRequest struct {
	ctx        context.Context
	ApiService ProgrammaticAPIKeysApi
	orgId      string
	apiUserId  string
}

type DeleteApiKeyApiParams struct {
	OrgId     string
	ApiUserId string
}

func (a *ProgrammaticAPIKeysApiService) DeleteApiKeyWithParams(ctx context.Context, args *DeleteApiKeyApiParams) DeleteApiKeyApiRequest {
	return DeleteApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		apiUserId:  args.ApiUserId,
	}
}

func (r DeleteApiKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApiKeyExecute(r)
}

/*
DeleteApiKey Remove One Organization API Key

Removes one organization API key from the specified organization. When you remove an API key from an organization, MongoDB Cloud also removes that key from any projects that use that key. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key.
	@return DeleteApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) DeleteApiKey(ctx context.Context, orgId string, apiUserId string) DeleteApiKeyApiRequest {
	return DeleteApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiUserId:  apiUserId,
	}
}

// DeleteApiKeyExecute executes the request
func (a *ProgrammaticAPIKeysApiService) DeleteApiKeyExecute(r DeleteApiKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.DeleteApiKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type DeleteApiKeyAccessListEntryApiRequest struct {
	ctx        context.Context
	ApiService ProgrammaticAPIKeysApi
	orgId      string
	apiUserId  string
	ipAddress  string
}

type DeleteApiKeyAccessListEntryApiParams struct {
	OrgId     string
	ApiUserId string
	IpAddress string
}

func (a *ProgrammaticAPIKeysApiService) DeleteApiKeyAccessListEntryWithParams(ctx context.Context, args *DeleteApiKeyAccessListEntryApiParams) DeleteApiKeyAccessListEntryApiRequest {
	return DeleteApiKeyAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		apiUserId:  args.ApiUserId,
		ipAddress:  args.IpAddress,
	}
}

func (r DeleteApiKeyAccessListEntryApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApiKeyAccessListEntryExecute(r)
}

/*
DeleteApiKeyAccessListEntry Remove One Access List Entry for One Organization API Key

Removes the specified access list entry from the specified organization API key. Resources require all API requests originate from the IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Read Write role. In addition, you cannot remove the requesting IP address from the requesting organization API key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.
	@param ipAddress One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
	@return DeleteApiKeyAccessListEntryApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) DeleteApiKeyAccessListEntry(ctx context.Context, orgId string, apiUserId string, ipAddress string) DeleteApiKeyAccessListEntryApiRequest {
	return DeleteApiKeyAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiUserId:  apiUserId,
		ipAddress:  ipAddress,
	}
}

// DeleteApiKeyAccessListEntryExecute executes the request
func (a *ProgrammaticAPIKeysApiService) DeleteApiKeyAccessListEntryExecute(r DeleteApiKeyAccessListEntryApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.DeleteApiKeyAccessListEntry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)
	if r.ipAddress == "" {
		return nil, reportError("ipAddress is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", url.PathEscape(r.ipAddress), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type GetApiKeyApiRequest struct {
	ctx        context.Context
	ApiService ProgrammaticAPIKeysApi
	orgId      string
	apiUserId  string
}

type GetApiKeyApiParams struct {
	OrgId     string
	ApiUserId string
}

func (a *ProgrammaticAPIKeysApiService) GetApiKeyWithParams(ctx context.Context, args *GetApiKeyApiParams) GetApiKeyApiRequest {
	return GetApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		apiUserId:  args.ApiUserId,
	}
}

func (r GetApiKeyApiRequest) Execute() (*ApiKeyUserDetails, *http.Response, error) {
	return r.ApiService.GetApiKeyExecute(r)
}

/*
GetApiKey Return One Organization API Key

Returns one organization API key. The organization API keys grant programmatic access to an organization. You can't use the API key to log into MongoDB Cloud through the user interface. To use this resource, the requesting Service Account or API Key must have the  Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.
	@return GetApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) GetApiKey(ctx context.Context, orgId string, apiUserId string) GetApiKeyApiRequest {
	return GetApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiUserId:  apiUserId,
	}
}

// GetApiKeyExecute executes the request
//
//	@return ApiKeyUserDetails
func (a *ProgrammaticAPIKeysApiService) GetApiKeyExecute(r GetApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiKeyUserDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.GetApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type GetApiKeyAccessListApiRequest struct {
	ctx        context.Context
	ApiService ProgrammaticAPIKeysApi
	orgId      string
	ipAddress  string
	apiUserId  string
}

type GetApiKeyAccessListApiParams struct {
	OrgId     string
	IpAddress string
	ApiUserId string
}

func (a *ProgrammaticAPIKeysApiService) GetApiKeyAccessListWithParams(ctx context.Context, args *GetApiKeyAccessListApiParams) GetApiKeyAccessListApiRequest {
	return GetApiKeyAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		ipAddress:  args.IpAddress,
		apiUserId:  args.ApiUserId,
	}
}

func (r GetApiKeyAccessListApiRequest) Execute() (*UserAccessListResponse, *http.Response, error) {
	return r.ApiService.GetApiKeyAccessListExecute(r)
}

/*
GetApiKeyAccessList Return One Access List Entry for One Organization API Key

Returns one access list entry for the specified organization API key. Resources require  all API requests originate from IP addresses on the API access list. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param ipAddress One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.
	@return GetApiKeyAccessListApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) GetApiKeyAccessList(ctx context.Context, orgId string, ipAddress string, apiUserId string) GetApiKeyAccessListApiRequest {
	return GetApiKeyAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		ipAddress:  ipAddress,
		apiUserId:  apiUserId,
	}
}

// GetApiKeyAccessListExecute executes the request
//
//	@return UserAccessListResponse
func (a *ProgrammaticAPIKeysApiService) GetApiKeyAccessListExecute(r GetApiKeyAccessListApiRequest) (*UserAccessListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserAccessListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.GetApiKeyAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.ipAddress == "" {
		return localVarReturnValue, nil, reportError("ipAddress is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", url.PathEscape(r.ipAddress), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListApiKeyAccessListsEntriesApiRequest struct {
	ctx          context.Context
	ApiService   ProgrammaticAPIKeysApi
	orgId        string
	apiUserId    string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListApiKeyAccessListsEntriesApiParams struct {
	OrgId        string
	ApiUserId    string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProgrammaticAPIKeysApiService) ListApiKeyAccessListsEntriesWithParams(ctx context.Context, args *ListApiKeyAccessListsEntriesApiParams) ListApiKeyAccessListsEntriesApiRequest {
	return ListApiKeyAccessListsEntriesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		apiUserId:    args.ApiUserId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListApiKeyAccessListsEntriesApiRequest) IncludeCount(includeCount bool) ListApiKeyAccessListsEntriesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListApiKeyAccessListsEntriesApiRequest) ItemsPerPage(itemsPerPage int) ListApiKeyAccessListsEntriesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListApiKeyAccessListsEntriesApiRequest) PageNum(pageNum int) ListApiKeyAccessListsEntriesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListApiKeyAccessListsEntriesApiRequest) Execute() (*PaginatedApiUserAccessListResponse, *http.Response, error) {
	return r.ApiService.ListApiKeyAccessListsEntriesExecute(r)
}

/*
ListApiKeyAccessListsEntries Return All Access List Entries for One Organization API Key

Returns all access list entries that you configured for the specified organization API key. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.
	@return ListApiKeyAccessListsEntriesApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) ListApiKeyAccessListsEntries(ctx context.Context, orgId string, apiUserId string) ListApiKeyAccessListsEntriesApiRequest {
	return ListApiKeyAccessListsEntriesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		apiUserId:  apiUserId,
	}
}

// ListApiKeyAccessListsEntriesExecute executes the request
//
//	@return PaginatedApiUserAccessListResponse
func (a *ProgrammaticAPIKeysApiService) ListApiKeyAccessListsEntriesExecute(r ListApiKeyAccessListsEntriesApiRequest) (*PaginatedApiUserAccessListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiUserAccessListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.ListApiKeyAccessListsEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListApiKeysApiRequest struct {
	ctx          context.Context
	ApiService   ProgrammaticAPIKeysApi
	orgId        string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListApiKeysApiParams struct {
	OrgId        string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProgrammaticAPIKeysApiService) ListApiKeysWithParams(ctx context.Context, args *ListApiKeysApiParams) ListApiKeysApiRequest {
	return ListApiKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListApiKeysApiRequest) IncludeCount(includeCount bool) ListApiKeysApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListApiKeysApiRequest) ItemsPerPage(itemsPerPage int) ListApiKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListApiKeysApiRequest) PageNum(pageNum int) ListApiKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListApiKeysApiRequest) Execute() (*PaginatedApiApiUser, *http.Response, error) {
	return r.ApiService.ListApiKeysExecute(r)
}

/*
ListApiKeys Return All Organization API Keys

Returns all organization API keys for the specified organization. The organization API keys grant programmatic access to an organization. You can't use the API key to log into MongoDB Cloud through the console. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListApiKeysApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) ListApiKeys(ctx context.Context, orgId string) ListApiKeysApiRequest {
	return ListApiKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListApiKeysExecute executes the request
//
//	@return PaginatedApiApiUser
func (a *ProgrammaticAPIKeysApiService) ListApiKeysExecute(r ListApiKeysApiRequest) (*PaginatedApiApiUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiApiUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.ListApiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListProjectApiKeysApiRequest struct {
	ctx          context.Context
	ApiService   ProgrammaticAPIKeysApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectApiKeysApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProgrammaticAPIKeysApiService) ListProjectApiKeysWithParams(ctx context.Context, args *ListProjectApiKeysApiParams) ListProjectApiKeysApiRequest {
	return ListProjectApiKeysApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectApiKeysApiRequest) IncludeCount(includeCount bool) ListProjectApiKeysApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectApiKeysApiRequest) ItemsPerPage(itemsPerPage int) ListProjectApiKeysApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectApiKeysApiRequest) PageNum(pageNum int) ListProjectApiKeysApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectApiKeysApiRequest) Execute() (*PaginatedApiApiUser, *http.Response, error) {
	return r.ApiService.ListProjectApiKeysExecute(r)
}

/*
ListProjectApiKeys Return All Organization API Keys Assigned to One Project

Returns all organization API keys that you assigned to the specified project. Users with the Project Owner role in the project associated with the API key can use the organization API key to access the resources. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectApiKeysApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) ListProjectApiKeys(ctx context.Context, groupId string) ListProjectApiKeysApiRequest {
	return ListProjectApiKeysApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectApiKeysExecute executes the request
//
//	@return PaginatedApiApiUser
func (a *ProgrammaticAPIKeysApiService) ListProjectApiKeysExecute(r ListProjectApiKeysApiRequest) (*PaginatedApiApiUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiApiUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.ListProjectApiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/apiKeys"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type RemoveProjectApiKeyApiRequest struct {
	ctx        context.Context
	ApiService ProgrammaticAPIKeysApi
	groupId    string
	apiUserId  string
}

type RemoveProjectApiKeyApiParams struct {
	GroupId   string
	ApiUserId string
}

func (a *ProgrammaticAPIKeysApiService) RemoveProjectApiKeyWithParams(ctx context.Context, args *RemoveProjectApiKeyApiParams) RemoveProjectApiKeyApiRequest {
	return RemoveProjectApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiUserId:  args.ApiUserId,
	}
}

func (r RemoveProjectApiKeyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveProjectApiKeyExecute(r)
}

/*
RemoveProjectApiKey Unassign One Organization API Key from One Project

Removes one organization API key from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
	@return RemoveProjectApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) RemoveProjectApiKey(ctx context.Context, groupId string, apiUserId string) RemoveProjectApiKeyApiRequest {
	return RemoveProjectApiKeyApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiUserId:  apiUserId,
	}
}

// RemoveProjectApiKeyExecute executes the request
func (a *ProgrammaticAPIKeysApiService) RemoveProjectApiKeyExecute(r RemoveProjectApiKeyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.RemoveProjectApiKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiUserId == "" {
		return nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type UpdateApiKeyApiRequest struct {
	ctx                           context.Context
	ApiService                    ProgrammaticAPIKeysApi
	orgId                         string
	apiUserId                     string
	updateAtlasOrganizationApiKey *UpdateAtlasOrganizationApiKey
}

type UpdateApiKeyApiParams struct {
	OrgId                         string
	ApiUserId                     string
	UpdateAtlasOrganizationApiKey *UpdateAtlasOrganizationApiKey
}

func (a *ProgrammaticAPIKeysApiService) UpdateApiKeyWithParams(ctx context.Context, args *UpdateApiKeyApiParams) UpdateApiKeyApiRequest {
	return UpdateApiKeyApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		apiUserId:                     args.ApiUserId,
		updateAtlasOrganizationApiKey: args.UpdateAtlasOrganizationApiKey,
	}
}

func (r UpdateApiKeyApiRequest) Execute() (*ApiKeyUserDetails, *http.Response, error) {
	return r.ApiService.UpdateApiKeyExecute(r)
}

/*
UpdateApiKey Update One Organization API Key

Updates one organization API key in the specified organization. The organization API keys  grant programmatic access to an organization. To use this resource, the requesting  API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.
	@return UpdateApiKeyApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) UpdateApiKey(ctx context.Context, orgId string, apiUserId string, updateAtlasOrganizationApiKey *UpdateAtlasOrganizationApiKey) UpdateApiKeyApiRequest {
	return UpdateApiKeyApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		apiUserId:                     apiUserId,
		updateAtlasOrganizationApiKey: updateAtlasOrganizationApiKey,
	}
}

// UpdateApiKeyExecute executes the request
//
//	@return ApiKeyUserDetails
func (a *ProgrammaticAPIKeysApiService) UpdateApiKeyExecute(r UpdateApiKeyApiRequest) (*ApiKeyUserDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiKeyUserDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.UpdateApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateAtlasOrganizationApiKey == nil {
		return localVarReturnValue, nil, reportError("updateAtlasOrganizationApiKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateAtlasOrganizationApiKey
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

type UpdateApiKeyRolesApiRequest struct {
	ctx                      context.Context
	ApiService               ProgrammaticAPIKeysApi
	groupId                  string
	apiUserId                string
	updateAtlasProjectApiKey *UpdateAtlasProjectApiKey
	pageNum                  *int
	itemsPerPage             *int
	includeCount             *bool
}

type UpdateApiKeyRolesApiParams struct {
	GroupId                  string
	ApiUserId                string
	UpdateAtlasProjectApiKey *UpdateAtlasProjectApiKey
	PageNum                  *int
	ItemsPerPage             *int
	IncludeCount             *bool
}

func (a *ProgrammaticAPIKeysApiService) UpdateApiKeyRolesWithParams(ctx context.Context, args *UpdateApiKeyRolesApiParams) UpdateApiKeyRolesApiRequest {
	return UpdateApiKeyRolesApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		apiUserId:                args.ApiUserId,
		updateAtlasProjectApiKey: args.UpdateAtlasProjectApiKey,
		pageNum:                  args.PageNum,
		itemsPerPage:             args.ItemsPerPage,
		includeCount:             args.IncludeCount,
	}
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r UpdateApiKeyRolesApiRequest) PageNum(pageNum int) UpdateApiKeyRolesApiRequest {
	r.pageNum = &pageNum
	return r
}

// Number of items that the response returns per page.
func (r UpdateApiKeyRolesApiRequest) ItemsPerPage(itemsPerPage int) UpdateApiKeyRolesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r UpdateApiKeyRolesApiRequest) IncludeCount(includeCount bool) UpdateApiKeyRolesApiRequest {
	r.includeCount = &includeCount
	return r
}

func (r UpdateApiKeyRolesApiRequest) Execute() (*ApiKeyUserDetails, *http.Response, error) {
	return r.ApiService.UpdateApiKeyRolesExecute(r)
}

/*
UpdateApiKeyRoles Update Organization API Key Roles for One Project

Updates the roles of the organization API key that you specify for the project that you specify. You must specify at least one valid role for the project. The application removes any roles that you do not include in this request if they were previously set in the organization API key that you specify for the project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param apiUserId Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
	@return UpdateApiKeyRolesApiRequest
*/
func (a *ProgrammaticAPIKeysApiService) UpdateApiKeyRoles(ctx context.Context, groupId string, apiUserId string, updateAtlasProjectApiKey *UpdateAtlasProjectApiKey) UpdateApiKeyRolesApiRequest {
	return UpdateApiKeyRolesApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		apiUserId:                apiUserId,
		updateAtlasProjectApiKey: updateAtlasProjectApiKey,
	}
}

// UpdateApiKeyRolesExecute executes the request
//
//	@return ApiKeyUserDetails
func (a *ProgrammaticAPIKeysApiService) UpdateApiKeyRolesExecute(r UpdateApiKeyRolesApiRequest) (*ApiKeyUserDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiKeyUserDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgrammaticAPIKeysApiService.UpdateApiKeyRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.apiUserId == "" {
		return localVarReturnValue, nil, reportError("apiUserId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"apiUserId"+"}", url.PathEscape(r.apiUserId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateAtlasProjectApiKey == nil {
		return localVarReturnValue, nil, reportError("updateAtlasProjectApiKey is required and must be specified")
	}

	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateAtlasProjectApiKey
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
