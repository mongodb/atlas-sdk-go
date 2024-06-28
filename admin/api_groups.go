// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type GroupsApi interface {

	/*
		AddProjectServiceAccount Assign One Service Account to One Project

		Assigns the specified organization Service Account to the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceAccountId Id of the service account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountRoleAssignment Service Account Roles to be assigned to the specified project.
		@return AddProjectServiceAccountApiRequest
	*/
	AddProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string, groupServiceAccountRoleAssignment *[]GroupServiceAccountRoleAssignment) AddProjectServiceAccountApiRequest
	/*
		AddProjectServiceAccount Assign One Service Account to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddProjectServiceAccountApiParams - Parameters for the request
		@return AddProjectServiceAccountApiRequest
	*/
	AddProjectServiceAccountWithParams(ctx context.Context, args *AddProjectServiceAccountApiParams) AddProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	AddProjectServiceAccountExecute(r AddProjectServiceAccountApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		CreateProjectServiceAccount Create and Assign one Service Account to one project

		Creates and assigns the specified Service Account to the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountRequest Details to create service account and add to the specified project.
		@return CreateProjectServiceAccountApiRequest
	*/
	CreateProjectServiceAccount(ctx context.Context, groupId string, groupServiceAccountRequest *GroupServiceAccountRequest) CreateProjectServiceAccountApiRequest
	/*
		CreateProjectServiceAccount Create and Assign one Service Account to one project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectServiceAccountApiParams - Parameters for the request
		@return CreateProjectServiceAccountApiRequest
	*/
	CreateProjectServiceAccountWithParams(ctx context.Context, args *CreateProjectServiceAccountApiParams) CreateProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	CreateProjectServiceAccountExecute(r CreateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		DeleteProjectServiceAccount Unassign One Service Account from One Project.

		Removes one Service Account from the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceAccountId Id of the service account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectServiceAccountApiRequest
	*/
	DeleteProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string) DeleteProjectServiceAccountApiRequest
	/*
		DeleteProjectServiceAccount Unassign One Service Account from One Project.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectServiceAccountApiParams - Parameters for the request
		@return DeleteProjectServiceAccountApiRequest
	*/
	DeleteProjectServiceAccountWithParams(ctx context.Context, args *DeleteProjectServiceAccountApiParams) DeleteProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	DeleteProjectServiceAccountExecute(r DeleteProjectServiceAccountApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetProjectServiceAccount Service Account Fetching

		Get project Service Account Details.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param serviceAccountId Id of the service account.
		@return GetProjectServiceAccountApiRequest
	*/
	GetProjectServiceAccount(ctx context.Context, groupId string, serviceAccountId string) GetProjectServiceAccountApiRequest
	/*
		GetProjectServiceAccount Service Account Fetching


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectServiceAccountApiParams - Parameters for the request
		@return GetProjectServiceAccountApiRequest
	*/
	GetProjectServiceAccountWithParams(ctx context.Context, args *GetProjectServiceAccountApiParams) GetProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	GetProjectServiceAccountExecute(r GetProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		ListProjectServiceAccounts Return All project Service Accounts

		Returns list of Service Accounts for a project with service account details.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectServiceAccountsApiRequest
	*/
	ListProjectServiceAccounts(ctx context.Context, groupId string) ListProjectServiceAccountsApiRequest
	/*
		ListProjectServiceAccounts Return All project Service Accounts


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectServiceAccountsApiParams - Parameters for the request
		@return ListProjectServiceAccountsApiRequest
	*/
	ListProjectServiceAccountsWithParams(ctx context.Context, args *ListProjectServiceAccountsApiParams) ListProjectServiceAccountsApiRequest

	// Method available only for mocking purposes
	ListProjectServiceAccountsExecute(r ListProjectServiceAccountsApiRequest) (*PaginatedGroupServiceAccounts, *http.Response, error)

	/*
		UpdateProjectServiceAccount Service Account Update in Project

		Update Service Account in Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceAccountId Id of the service account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountUpdateRequest Details to update service account in the specified Project.
		@return UpdateProjectServiceAccountApiRequest
	*/
	UpdateProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string, groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest) UpdateProjectServiceAccountApiRequest
	/*
		UpdateProjectServiceAccount Service Account Update in Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectServiceAccountApiParams - Parameters for the request
		@return UpdateProjectServiceAccountApiRequest
	*/
	UpdateProjectServiceAccountWithParams(ctx context.Context, args *UpdateProjectServiceAccountApiParams) UpdateProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	UpdateProjectServiceAccountExecute(r UpdateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)
}

// GroupsApiService GroupsApi service
type GroupsApiService service

type AddProjectServiceAccountApiRequest struct {
	ctx                               context.Context
	ApiService                        GroupsApi
	serviceAccountId                  string
	groupId                           string
	groupServiceAccountRoleAssignment *[]GroupServiceAccountRoleAssignment
}

type AddProjectServiceAccountApiParams struct {
	ServiceAccountId                  string
	GroupId                           string
	GroupServiceAccountRoleAssignment *[]GroupServiceAccountRoleAssignment
}

func (a *GroupsApiService) AddProjectServiceAccountWithParams(ctx context.Context, args *AddProjectServiceAccountApiParams) AddProjectServiceAccountApiRequest {
	return AddProjectServiceAccountApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		serviceAccountId:                  args.ServiceAccountId,
		groupId:                           args.GroupId,
		groupServiceAccountRoleAssignment: args.GroupServiceAccountRoleAssignment,
	}
}

func (r AddProjectServiceAccountApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AddProjectServiceAccountExecute(r)
}

/*
AddProjectServiceAccount Assign One Service Account to One Project

Assigns the specified organization Service Account to the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceAccountId Id of the service account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AddProjectServiceAccountApiRequest
*/
func (a *GroupsApiService) AddProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string, groupServiceAccountRoleAssignment *[]GroupServiceAccountRoleAssignment) AddProjectServiceAccountApiRequest {
	return AddProjectServiceAccountApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		serviceAccountId:                  serviceAccountId,
		groupId:                           groupId,
		groupServiceAccountRoleAssignment: groupServiceAccountRoleAssignment,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *GroupsApiService) AddProjectServiceAccountExecute(r AddProjectServiceAccountApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.AddProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{serviceAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceAccountId"+"}", url.PathEscape(parameterValueToString(r.serviceAccountId, "serviceAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountRoleAssignment == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountRoleAssignment is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-10-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupServiceAccountRoleAssignment
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

type CreateProjectServiceAccountApiRequest struct {
	ctx                        context.Context
	ApiService                 GroupsApi
	groupId                    string
	groupServiceAccountRequest *GroupServiceAccountRequest
}

type CreateProjectServiceAccountApiParams struct {
	GroupId                    string
	GroupServiceAccountRequest *GroupServiceAccountRequest
}

func (a *GroupsApiService) CreateProjectServiceAccountWithParams(ctx context.Context, args *CreateProjectServiceAccountApiParams) CreateProjectServiceAccountApiRequest {
	return CreateProjectServiceAccountApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		groupServiceAccountRequest: args.GroupServiceAccountRequest,
	}
}

func (r CreateProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.CreateProjectServiceAccountExecute(r)
}

/*
CreateProjectServiceAccount Create and Assign one Service Account to one project

Creates and assigns the specified Service Account to the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectServiceAccountApiRequest
*/
func (a *GroupsApiService) CreateProjectServiceAccount(ctx context.Context, groupId string, groupServiceAccountRequest *GroupServiceAccountRequest) CreateProjectServiceAccountApiRequest {
	return CreateProjectServiceAccountApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		groupServiceAccountRequest: groupServiceAccountRequest,
	}
}

// Execute executes the request
//
//	@return GroupServiceAccount
func (a *GroupsApiService) CreateProjectServiceAccountExecute(r CreateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.CreateProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountRequest == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-10-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupServiceAccountRequest
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

type DeleteProjectServiceAccountApiRequest struct {
	ctx              context.Context
	ApiService       GroupsApi
	serviceAccountId string
	groupId          string
}

type DeleteProjectServiceAccountApiParams struct {
	ServiceAccountId string
	GroupId          string
}

func (a *GroupsApiService) DeleteProjectServiceAccountWithParams(ctx context.Context, args *DeleteProjectServiceAccountApiParams) DeleteProjectServiceAccountApiRequest {
	return DeleteProjectServiceAccountApiRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: args.ServiceAccountId,
		groupId:          args.GroupId,
	}
}

func (r DeleteProjectServiceAccountApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteProjectServiceAccountExecute(r)
}

/*
DeleteProjectServiceAccount Unassign One Service Account from One Project.

Removes one Service Account from the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceAccountId Id of the service account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectServiceAccountApiRequest
*/
func (a *GroupsApiService) DeleteProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string) DeleteProjectServiceAccountApiRequest {
	return DeleteProjectServiceAccountApiRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		groupId:          groupId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *GroupsApiService) DeleteProjectServiceAccountExecute(r DeleteProjectServiceAccountApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.DeleteProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{serviceAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceAccountId"+"}", url.PathEscape(parameterValueToString(r.serviceAccountId, "serviceAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

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

type GetProjectServiceAccountApiRequest struct {
	ctx              context.Context
	ApiService       GroupsApi
	groupId          string
	serviceAccountId string
}

type GetProjectServiceAccountApiParams struct {
	GroupId          string
	ServiceAccountId string
}

func (a *GroupsApiService) GetProjectServiceAccountWithParams(ctx context.Context, args *GetProjectServiceAccountApiParams) GetProjectServiceAccountApiRequest {
	return GetProjectServiceAccountApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		serviceAccountId: args.ServiceAccountId,
	}
}

func (r GetProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.GetProjectServiceAccountExecute(r)
}

/*
GetProjectServiceAccount Service Account Fetching

Get project Service Account Details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param serviceAccountId Id of the service account.
	@return GetProjectServiceAccountApiRequest
*/
func (a *GroupsApiService) GetProjectServiceAccount(ctx context.Context, groupId string, serviceAccountId string) GetProjectServiceAccountApiRequest {
	return GetProjectServiceAccountApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		serviceAccountId: serviceAccountId,
	}
}

// Execute executes the request
//
//	@return GroupServiceAccount
func (a *GroupsApiService) GetProjectServiceAccountExecute(r GetProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{serviceAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceAccountId"+"}", url.PathEscape(parameterValueToString(r.serviceAccountId, "serviceAccountId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

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

type ListProjectServiceAccountsApiRequest struct {
	ctx          context.Context
	ApiService   GroupsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListProjectServiceAccountsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *GroupsApiService) ListProjectServiceAccountsWithParams(ctx context.Context, args *ListProjectServiceAccountsApiParams) ListProjectServiceAccountsApiRequest {
	return ListProjectServiceAccountsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListProjectServiceAccountsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectServiceAccountsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectServiceAccountsApiRequest) PageNum(pageNum int) ListProjectServiceAccountsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectServiceAccountsApiRequest) Execute() (*PaginatedGroupServiceAccounts, *http.Response, error) {
	return r.ApiService.ListProjectServiceAccountsExecute(r)
}

/*
ListProjectServiceAccounts Return All project Service Accounts

Returns list of Service Accounts for a project with service account details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectServiceAccountsApiRequest
*/
func (a *GroupsApiService) ListProjectServiceAccounts(ctx context.Context, groupId string) ListProjectServiceAccountsApiRequest {
	return ListProjectServiceAccountsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedGroupServiceAccounts
func (a *GroupsApiService) ListProjectServiceAccountsExecute(r ListProjectServiceAccountsApiRequest) (*PaginatedGroupServiceAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedGroupServiceAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.ListProjectServiceAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

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

type UpdateProjectServiceAccountApiRequest struct {
	ctx                              context.Context
	ApiService                       GroupsApi
	serviceAccountId                 string
	groupId                          string
	groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest
}

type UpdateProjectServiceAccountApiParams struct {
	ServiceAccountId                 string
	GroupId                          string
	GroupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest
}

func (a *GroupsApiService) UpdateProjectServiceAccountWithParams(ctx context.Context, args *UpdateProjectServiceAccountApiParams) UpdateProjectServiceAccountApiRequest {
	return UpdateProjectServiceAccountApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		serviceAccountId:                 args.ServiceAccountId,
		groupId:                          args.GroupId,
		groupServiceAccountUpdateRequest: args.GroupServiceAccountUpdateRequest,
	}
}

func (r UpdateProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.UpdateProjectServiceAccountExecute(r)
}

/*
UpdateProjectServiceAccount Service Account Update in Project

Update Service Account in Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceAccountId Id of the service account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectServiceAccountApiRequest
*/
func (a *GroupsApiService) UpdateProjectServiceAccount(ctx context.Context, serviceAccountId string, groupId string, groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest) UpdateProjectServiceAccountApiRequest {
	return UpdateProjectServiceAccountApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		serviceAccountId:                 serviceAccountId,
		groupId:                          groupId,
		groupServiceAccountUpdateRequest: groupServiceAccountUpdateRequest,
	}
}

// Execute executes the request
//
//	@return GroupServiceAccount
func (a *GroupsApiService) UpdateProjectServiceAccountExecute(r UpdateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.UpdateProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{serviceAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceAccountId"+"}", url.PathEscape(parameterValueToString(r.serviceAccountId, "serviceAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-10-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupServiceAccountUpdateRequest
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
