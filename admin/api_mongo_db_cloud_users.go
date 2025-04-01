// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MongoDBCloudUsersApi interface {

	/*
			CreateUser Create One MongoDB Cloud User

			Creates one MongoDB Cloud user account. A MongoDB Cloud user account grants access to only the MongoDB Cloud application. To grant database access, create a database user. MongoDB Cloud sends an email to the users you specify, inviting them to join the project. Invited users don't have access to the project until they accept the invitation. Invitations expire after 30 days.

		 MongoDB Cloud limits MongoDB Cloud user membership to a maximum of 250 MongoDB Cloud users per team. MongoDB Cloud limits MongoDB Cloud user membership to 500 MongoDB Cloud users per project and 500 MongoDB Cloud users per organization, which includes the combined membership of all projects in the organization. MongoDB Cloud raises an error if an operation exceeds these limits. For example, if you have an organization with five projects, and each project has 100 MongoDB Cloud users, and each MongoDB Cloud user belongs to only one project, you can't add any MongoDB Cloud users to this organization without first removing existing MongoDB Cloud users from the organization.

		 To use this resource, the requesting API Key can have any role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param cloudAppUser MongoDB Cloud user account to create.
			@return CreateUserApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	CreateUser(ctx context.Context, cloudAppUser *CloudAppUser) CreateUserApiRequest
	/*
		CreateUser Create One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateUserApiParams - Parameters for the request
		@return CreateUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	CreateUserWithParams(ctx context.Context, args *CreateUserApiParams) CreateUserApiRequest

	// Method available only for mocking purposes
	CreateUserExecute(r CreateUserApiRequest) (*CloudAppUser, *http.Response, error)

	/*
		GetUser Return One MongoDB Cloud User using Its ID

		Returns the details for one MongoDB Cloud user account with the specified unique identifier for the user. You can't use this endpoint to return information on an API Key. To return information about an API Key, use the [Return One Organization](#tag/Organizations/operation/getOrganization) API Key endpoint. You can always retrieve your own user account. If you are the owner of a MongoDB Cloud organization or project, you can also retrieve the user profile for any user with membership in that organization or project. To use this resource, the requesting API Key can have any role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId Unique 24-hexadecimal digit string that identifies this user.
		@return GetUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	GetUser(ctx context.Context, userId string) GetUserApiRequest
	/*
		GetUser Return One MongoDB Cloud User using Its ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetUserApiParams - Parameters for the request
		@return GetUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	GetUserWithParams(ctx context.Context, args *GetUserApiParams) GetUserApiRequest

	// Method available only for mocking purposes
	GetUserExecute(r GetUserApiRequest) (*CloudAppUser, *http.Response, error)

	/*
		GetUserByUsername Return One MongoDB Cloud User using Their Username

		Returns the details for one MongoDB Cloud user account with the specified username. You can't use this endpoint to return information about an API Key. To return information about an API Key, use the [Return One Organization](#tag/Organizations/operation/getOrganization) API Key endpoint. To use this resource, the requesting API Key can have any role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userName Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.
		@return GetUserByUsernameApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	GetUserByUsername(ctx context.Context, userName string) GetUserByUsernameApiRequest
	/*
		GetUserByUsername Return One MongoDB Cloud User using Their Username


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetUserByUsernameApiParams - Parameters for the request
		@return GetUserByUsernameApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	GetUserByUsernameWithParams(ctx context.Context, args *GetUserByUsernameApiParams) GetUserByUsernameApiRequest

	// Method available only for mocking purposes
	GetUserByUsernameExecute(r GetUserByUsernameApiRequest) (*CloudAppUser, *http.Response, error)

	/*
			ListOrganizationUsers Return All MongoDB Cloud Users in One Organization

			Returns details about the pending and active MongoDB Cloud users associated with the specified organization. To use this resource, the requesting API Key must have the Organization Member role.

		**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

		**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrganizationUsersApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListOrganizationUsers(ctx context.Context, orgId string) ListOrganizationUsersApiRequest
	/*
		ListOrganizationUsers Return All MongoDB Cloud Users in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationUsersApiParams - Parameters for the request
		@return ListOrganizationUsersApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListOrganizationUsersWithParams(ctx context.Context, args *ListOrganizationUsersApiParams) ListOrganizationUsersApiRequest

	// Method available only for mocking purposes
	ListOrganizationUsersExecute(r ListOrganizationUsersApiRequest) (*PaginatedAppUser, *http.Response, error)

	/*
			ListProjectUsers Return All MongoDB Cloud Users in One Project

			Returns details about the pending and active MongoDB Cloud users associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

		**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListProjectUsersApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListProjectUsers(ctx context.Context, groupId string) ListProjectUsersApiRequest
	/*
		ListProjectUsers Return All MongoDB Cloud Users in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectUsersApiParams - Parameters for the request
		@return ListProjectUsersApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListProjectUsersWithParams(ctx context.Context, args *ListProjectUsersApiParams) ListProjectUsersApiRequest

	// Method available only for mocking purposes
	ListProjectUsersExecute(r ListProjectUsersApiRequest) (*PaginatedAppUser, *http.Response, error)

	/*
			ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team

			Returns details about the pending and active MongoDB Cloud users associated with the specified team in the organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role.

		**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

		**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param teamId Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.
			@return ListTeamUsersApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListTeamUsers(ctx context.Context, orgId string, teamId string) ListTeamUsersApiRequest
	/*
		ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListTeamUsersApiParams - Parameters for the request
		@return ListTeamUsersApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	ListTeamUsersWithParams(ctx context.Context, args *ListTeamUsersApiParams) ListTeamUsersApiRequest

	// Method available only for mocking purposes
	ListTeamUsersExecute(r ListTeamUsersApiRequest) (*PaginatedAppUser, *http.Response, error)

	/*
			RemoveOrganizationUser Remove One MongoDB Cloud User From One Organization

			Removes one MongoDB Cloud user in the specified organization. You can remove an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

		**Note**: This resource cannot be used to remove pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

		**Note**: To remove pending or active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users can be removed.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Organization](#tag/MongoDB-Cloud-Users/operation/listOrganizationUsers) resource and filter by username.
			@return RemoveOrganizationUserApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	RemoveOrganizationUser(ctx context.Context, orgId string, userId string) RemoveOrganizationUserApiRequest
	/*
		RemoveOrganizationUser Remove One MongoDB Cloud User From One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveOrganizationUserApiParams - Parameters for the request
		@return RemoveOrganizationUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	RemoveOrganizationUserWithParams(ctx context.Context, args *RemoveOrganizationUserApiParams) RemoveOrganizationUserApiRequest

	// Method available only for mocking purposes
	RemoveOrganizationUserExecute(r RemoveOrganizationUserApiRequest) (any, *http.Response, error)

	/*
			RemoveProjectUser Remove One MongoDB Cloud User from One Project

			Removes one MongoDB Cloud user from the specified project. You can remove an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Project Owner role.

		**Note**: This resource cannot be used to remove pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

		**Note**: To remove pending or active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users can be removed.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Project](#tag/MongoDB-Cloud-Users/operation/listProjectUsers) resource and filter by username.
			@return RemoveProjectUserApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	RemoveProjectUser(ctx context.Context, groupId string, userId string) RemoveProjectUserApiRequest
	/*
		RemoveProjectUser Remove One MongoDB Cloud User from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveProjectUserApiParams - Parameters for the request
		@return RemoveProjectUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for MongoDBCloudUsersApi
	*/
	RemoveProjectUserWithParams(ctx context.Context, args *RemoveProjectUserApiParams) RemoveProjectUserApiRequest

	// Method available only for mocking purposes
	RemoveProjectUserExecute(r RemoveProjectUserApiRequest) (*http.Response, error)
}

// MongoDBCloudUsersApiService MongoDBCloudUsersApi service
type MongoDBCloudUsersApiService service

type CreateUserApiRequest struct {
	ctx          context.Context
	ApiService   MongoDBCloudUsersApi
	cloudAppUser *CloudAppUser
}

type CreateUserApiParams struct {
	CloudAppUser *CloudAppUser
}

func (a *MongoDBCloudUsersApiService) CreateUserWithParams(ctx context.Context, args *CreateUserApiParams) CreateUserApiRequest {
	return CreateUserApiRequest{
		ApiService:   a,
		ctx:          ctx,
		cloudAppUser: args.CloudAppUser,
	}
}

func (r CreateUserApiRequest) Execute() (*CloudAppUser, *http.Response, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
CreateUser Create One MongoDB Cloud User

Creates one MongoDB Cloud user account. A MongoDB Cloud user account grants access to only the MongoDB Cloud application. To grant database access, create a database user. MongoDB Cloud sends an email to the users you specify, inviting them to join the project. Invited users don't have access to the project until they accept the invitation. Invitations expire after 30 days.

	MongoDB Cloud limits MongoDB Cloud user membership to a maximum of 250 MongoDB Cloud users per team. MongoDB Cloud limits MongoDB Cloud user membership to 500 MongoDB Cloud users per project and 500 MongoDB Cloud users per organization, which includes the combined membership of all projects in the organization. MongoDB Cloud raises an error if an operation exceeds these limits. For example, if you have an organization with five projects, and each project has 100 MongoDB Cloud users, and each MongoDB Cloud user belongs to only one project, you can't add any MongoDB Cloud users to this organization without first removing existing MongoDB Cloud users from the organization.

	To use this resource, the requesting API Key can have any role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateUserApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) CreateUser(ctx context.Context, cloudAppUser *CloudAppUser) CreateUserApiRequest {
	return CreateUserApiRequest{
		ApiService:   a,
		ctx:          ctx,
		cloudAppUser: cloudAppUser,
	}
}

// CreateUserExecute executes the request
//
//	@return CloudAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) CreateUserExecute(r CreateUserApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.CreateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudAppUser == nil {
		return localVarReturnValue, nil, reportError("cloudAppUser is required and must be specified")
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
	localVarPostBody = r.cloudAppUser
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

type GetUserApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	userId     string
}

type GetUserApiParams struct {
	UserId string
}

func (a *MongoDBCloudUsersApiService) GetUserWithParams(ctx context.Context, args *GetUserApiParams) GetUserApiRequest {
	return GetUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     args.UserId,
	}
}

func (r GetUserApiRequest) Execute() (*CloudAppUser, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Return One MongoDB Cloud User using Its ID

Returns the details for one MongoDB Cloud user account with the specified unique identifier for the user. You can't use this endpoint to return information on an API Key. To return information about an API Key, use the [Return One Organization](#tag/Organizations/operation/getOrganization) API Key endpoint. You can always retrieve your own user account. If you are the owner of a MongoDB Cloud organization or project, you can also retrieve the user profile for any user with membership in that organization or project. To use this resource, the requesting API Key can have any role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId Unique 24-hexadecimal digit string that identifies this user.
	@return GetUserApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) GetUser(ctx context.Context, userId string) GetUserApiRequest {
	return GetUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// GetUserExecute executes the request
//
//	@return CloudAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) GetUserExecute(r GetUserApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(r.userId), -1)

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

type GetUserByUsernameApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	userName   string
}

type GetUserByUsernameApiParams struct {
	UserName string
}

func (a *MongoDBCloudUsersApiService) GetUserByUsernameWithParams(ctx context.Context, args *GetUserByUsernameApiParams) GetUserByUsernameApiRequest {
	return GetUserByUsernameApiRequest{
		ApiService: a,
		ctx:        ctx,
		userName:   args.UserName,
	}
}

func (r GetUserByUsernameApiRequest) Execute() (*CloudAppUser, *http.Response, error) {
	return r.ApiService.GetUserByUsernameExecute(r)
}

/*
GetUserByUsername Return One MongoDB Cloud User using Their Username

Returns the details for one MongoDB Cloud user account with the specified username. You can't use this endpoint to return information about an API Key. To return information about an API Key, use the [Return One Organization](#tag/Organizations/operation/getOrganization) API Key endpoint. To use this resource, the requesting API Key can have any role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userName Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.
	@return GetUserByUsernameApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) GetUserByUsername(ctx context.Context, userName string) GetUserByUsernameApiRequest {
	return GetUserByUsernameApiRequest{
		ApiService: a,
		ctx:        ctx,
		userName:   userName,
	}
}

// GetUserByUsernameExecute executes the request
//
//	@return CloudAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) GetUserByUsernameExecute(r GetUserByUsernameApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.GetUserByUsername")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/users/byName/{userName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userName"+"}", url.PathEscape(r.userName), -1)

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

type ListOrganizationUsersApiRequest struct {
	ctx                 context.Context
	ApiService          MongoDBCloudUsersApi
	orgId               string
	includeCount        *bool
	itemsPerPage        *int
	pageNum             *int
	username            *string
	orgMembershipStatus *string
}

type ListOrganizationUsersApiParams struct {
	OrgId               string
	IncludeCount        *bool
	ItemsPerPage        *int
	PageNum             *int
	Username            *string
	OrgMembershipStatus *string
}

func (a *MongoDBCloudUsersApiService) ListOrganizationUsersWithParams(ctx context.Context, args *ListOrganizationUsersApiParams) ListOrganizationUsersApiRequest {
	return ListOrganizationUsersApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		orgId:               args.OrgId,
		includeCount:        args.IncludeCount,
		itemsPerPage:        args.ItemsPerPage,
		pageNum:             args.PageNum,
		username:            args.Username,
		orgMembershipStatus: args.OrgMembershipStatus,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOrganizationUsersApiRequest) IncludeCount(includeCount bool) ListOrganizationUsersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListOrganizationUsersApiRequest) ItemsPerPage(itemsPerPage int) ListOrganizationUsersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrganizationUsersApiRequest) PageNum(pageNum int) ListOrganizationUsersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Email address to filter users by. Not supported in deprecated versions.
func (r ListOrganizationUsersApiRequest) Username(username string) ListOrganizationUsersApiRequest {
	r.username = &username
	return r
}

// Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.
func (r ListOrganizationUsersApiRequest) OrgMembershipStatus(orgMembershipStatus string) ListOrganizationUsersApiRequest {
	r.orgMembershipStatus = &orgMembershipStatus
	return r
}

func (r ListOrganizationUsersApiRequest) Execute() (*PaginatedAppUser, *http.Response, error) {
	return r.ApiService.ListOrganizationUsersExecute(r)
}

/*
ListOrganizationUsers Return All MongoDB Cloud Users in One Organization

Returns details about the pending and active MongoDB Cloud users associated with the specified organization. To use this resource, the requesting API Key must have the Organization Member role.

**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationUsersApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) ListOrganizationUsers(ctx context.Context, orgId string) ListOrganizationUsersApiRequest {
	return ListOrganizationUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrganizationUsersExecute executes the request
//
//	@return PaginatedAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) ListOrganizationUsersExecute(r ListOrganizationUsersApiRequest) (*PaginatedAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.ListOrganizationUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users"
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
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
	}
	if r.orgMembershipStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orgMembershipStatus", r.orgMembershipStatus, "")
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

type ListProjectUsersApiRequest struct {
	ctx                 context.Context
	ApiService          MongoDBCloudUsersApi
	groupId             string
	includeCount        *bool
	itemsPerPage        *int
	pageNum             *int
	flattenTeams        *bool
	includeOrgUsers     *bool
	orgMembershipStatus *string
	username            *string
}

type ListProjectUsersApiParams struct {
	GroupId             string
	IncludeCount        *bool
	ItemsPerPage        *int
	PageNum             *int
	FlattenTeams        *bool
	IncludeOrgUsers     *bool
	OrgMembershipStatus *string
	Username            *string
}

func (a *MongoDBCloudUsersApiService) ListProjectUsersWithParams(ctx context.Context, args *ListProjectUsersApiParams) ListProjectUsersApiRequest {
	return ListProjectUsersApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		includeCount:        args.IncludeCount,
		itemsPerPage:        args.ItemsPerPage,
		pageNum:             args.PageNum,
		flattenTeams:        args.FlattenTeams,
		includeOrgUsers:     args.IncludeOrgUsers,
		orgMembershipStatus: args.OrgMembershipStatus,
		username:            args.Username,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectUsersApiRequest) IncludeCount(includeCount bool) ListProjectUsersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectUsersApiRequest) ItemsPerPage(itemsPerPage int) ListProjectUsersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectUsersApiRequest) PageNum(pageNum int) ListProjectUsersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the returned list should include users who belong to a team with a role in this project. You might not have assigned the individual users a role in this project. If &#x60;\&quot;flattenTeams\&quot; : false&#x60;, this resource returns only users with a role in the project.  If &#x60;\&quot;flattenTeams\&quot; : true&#x60;, this resource returns both users with roles in the project and users who belong to teams with roles in the project.
func (r ListProjectUsersApiRequest) FlattenTeams(flattenTeams bool) ListProjectUsersApiRequest {
	r.flattenTeams = &flattenTeams
	return r
}

// Flag that indicates whether the returned list should include users with implicit access to the project, the Organization Owner or Organization Read Only role. You might not have assigned the individual users a role in this project. If &#x60;\&quot;includeOrgUsers\&quot;: false&#x60;, this resource returns only users with a role in the project. If &#x60;\&quot;includeOrgUsers\&quot;: true&#x60;, this resource returns both users with roles in the project and users who have implicit access to the project through their organization role.
func (r ListProjectUsersApiRequest) IncludeOrgUsers(includeOrgUsers bool) ListProjectUsersApiRequest {
	r.includeOrgUsers = &includeOrgUsers
	return r
}

// Flag that indicates whether to filter the returned list by users organization membership status. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.
func (r ListProjectUsersApiRequest) OrgMembershipStatus(orgMembershipStatus string) ListProjectUsersApiRequest {
	r.orgMembershipStatus = &orgMembershipStatus
	return r
}

// Email address to filter users by. Not supported in deprecated versions.
func (r ListProjectUsersApiRequest) Username(username string) ListProjectUsersApiRequest {
	r.username = &username
	return r
}

func (r ListProjectUsersApiRequest) Execute() (*PaginatedAppUser, *http.Response, error) {
	return r.ApiService.ListProjectUsersExecute(r)
}

/*
ListProjectUsers Return All MongoDB Cloud Users in One Project

Returns details about the pending and active MongoDB Cloud users associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectUsersApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) ListProjectUsers(ctx context.Context, groupId string) ListProjectUsersApiRequest {
	return ListProjectUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectUsersExecute executes the request
//
//	@return PaginatedAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) ListProjectUsersExecute(r ListProjectUsersApiRequest) (*PaginatedAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.ListProjectUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users"
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
	if r.flattenTeams != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flattenTeams", r.flattenTeams, "")
	} else {
		var defaultValue bool = false
		r.flattenTeams = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "flattenTeams", r.flattenTeams, "")
	}
	if r.includeOrgUsers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOrgUsers", r.includeOrgUsers, "")
	} else {
		var defaultValue bool = false
		r.includeOrgUsers = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOrgUsers", r.includeOrgUsers, "")
	}
	if r.orgMembershipStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orgMembershipStatus", r.orgMembershipStatus, "")
	}
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
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

type ListTeamUsersApiRequest struct {
	ctx                 context.Context
	ApiService          MongoDBCloudUsersApi
	orgId               string
	teamId              string
	itemsPerPage        *int
	pageNum             *int
	username            *string
	orgMembershipStatus *string
}

type ListTeamUsersApiParams struct {
	OrgId               string
	TeamId              string
	ItemsPerPage        *int
	PageNum             *int
	Username            *string
	OrgMembershipStatus *string
}

func (a *MongoDBCloudUsersApiService) ListTeamUsersWithParams(ctx context.Context, args *ListTeamUsersApiParams) ListTeamUsersApiRequest {
	return ListTeamUsersApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		orgId:               args.OrgId,
		teamId:              args.TeamId,
		itemsPerPage:        args.ItemsPerPage,
		pageNum:             args.PageNum,
		username:            args.Username,
		orgMembershipStatus: args.OrgMembershipStatus,
	}
}

// Number of items that the response returns per page.
func (r ListTeamUsersApiRequest) ItemsPerPage(itemsPerPage int) ListTeamUsersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListTeamUsersApiRequest) PageNum(pageNum int) ListTeamUsersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Email address to filter users by. Not supported in deprecated versions.
func (r ListTeamUsersApiRequest) Username(username string) ListTeamUsersApiRequest {
	r.username = &username
	return r
}

// Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.
func (r ListTeamUsersApiRequest) OrgMembershipStatus(orgMembershipStatus string) ListTeamUsersApiRequest {
	r.orgMembershipStatus = &orgMembershipStatus
	return r
}

func (r ListTeamUsersApiRequest) Execute() (*PaginatedAppUser, *http.Response, error) {
	return r.ApiService.ListTeamUsersExecute(r)
}

/*
ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team

Returns details about the pending and active MongoDB Cloud users associated with the specified team in the organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role.

**Note**: This resource cannot be used to view details about users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

**Note**: To return both pending and active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users will be returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.
	@return ListTeamUsersApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) ListTeamUsers(ctx context.Context, orgId string, teamId string) ListTeamUsersApiRequest {
	return ListTeamUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
	}
}

// ListTeamUsersExecute executes the request
//
//	@return PaginatedAppUser
//
// Deprecated
func (a *MongoDBCloudUsersApiService) ListTeamUsersExecute(r ListTeamUsersApiRequest) (*PaginatedAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.ListTeamUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

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
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
	}
	if r.orgMembershipStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orgMembershipStatus", r.orgMembershipStatus, "")
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

type RemoveOrganizationUserApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	orgId      string
	userId     string
}

type RemoveOrganizationUserApiParams struct {
	OrgId  string
	UserId string
}

func (a *MongoDBCloudUsersApiService) RemoveOrganizationUserWithParams(ctx context.Context, args *RemoveOrganizationUserApiParams) RemoveOrganizationUserApiRequest {
	return RemoveOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		userId:     args.UserId,
	}
}

func (r RemoveOrganizationUserApiRequest) Execute() (any, *http.Response, error) {
	return r.ApiService.RemoveOrganizationUserExecute(r)
}

/*
RemoveOrganizationUser Remove One MongoDB Cloud User From One Organization

Removes one MongoDB Cloud user in the specified organization. You can remove an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

**Note**: This resource cannot be used to remove pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

**Note**: To remove pending or active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users can be removed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Organization](#tag/MongoDB-Cloud-Users/operation/listOrganizationUsers) resource and filter by username.
	@return RemoveOrganizationUserApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) RemoveOrganizationUser(ctx context.Context, orgId string, userId string) RemoveOrganizationUserApiRequest {
	return RemoveOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		userId:     userId,
	}
}

// RemoveOrganizationUserExecute executes the request
//
//	@return any
//
// Deprecated
func (a *MongoDBCloudUsersApiService) RemoveOrganizationUserExecute(r RemoveOrganizationUserApiRequest) (any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.RemoveOrganizationUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(r.userId), -1)

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

type RemoveProjectUserApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	groupId    string
	userId     string
}

type RemoveProjectUserApiParams struct {
	GroupId string
	UserId  string
}

func (a *MongoDBCloudUsersApiService) RemoveProjectUserWithParams(ctx context.Context, args *RemoveProjectUserApiParams) RemoveProjectUserApiRequest {
	return RemoveProjectUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		userId:     args.UserId,
	}
}

func (r RemoveProjectUserApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveProjectUserExecute(r)
}

/*
RemoveProjectUser Remove One MongoDB Cloud User from One Project

Removes one MongoDB Cloud user from the specified project. You can remove an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Project Owner role.

**Note**: This resource cannot be used to remove pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

**Note**: To remove pending or active users, use v2-{2025-02-19} or later. If using a deprecated version, only active users can be removed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Project](#tag/MongoDB-Cloud-Users/operation/listProjectUsers) resource and filter by username.
	@return RemoveProjectUserApiRequest

Deprecated
*/
func (a *MongoDBCloudUsersApiService) RemoveProjectUser(ctx context.Context, groupId string, userId string) RemoveProjectUserApiRequest {
	return RemoveProjectUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		userId:     userId,
	}
}

// RemoveProjectUserExecute executes the request
// Deprecated
func (a *MongoDBCloudUsersApiService) RemoveProjectUserExecute(r RemoveProjectUserApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.RemoveProjectUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(r.userId), -1)

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
