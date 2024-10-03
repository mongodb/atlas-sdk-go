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
			AddOrganizationRole Add One Organization Role to One MongoDB Cloud User

			Adds one organization level role to the MongoDB Cloud user. You can add a role to an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

		**Note**: This operation is atomic.

		**Note**: This resource cannot be used to add a role to users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
			@return AddOrganizationRoleApiRequest
	*/
	AddOrganizationRole(ctx context.Context, orgId string, userId string) AddOrganizationRoleApiRequest
	/*
		AddOrganizationRole Add One Organization Role to One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddOrganizationRoleApiParams - Parameters for the request
		@return AddOrganizationRoleApiRequest
	*/
	AddOrganizationRoleWithParams(ctx context.Context, args *AddOrganizationRoleApiParams) AddOrganizationRoleApiRequest

	// Method available only for mocking purposes
	AddOrganizationRoleExecute(r AddOrganizationRoleApiRequest) (*OrgUserResponse, *http.Response, error)

	/*
			CreateOrganizationUser Create One MongoDB Cloud User in One Organization

			Invites one new or existing MongoDB Cloud user to join the organization. The invitation to join the organization will be sent to the username provided and must be accepted within 30 days. To use this resource, the requesting API Key must have the Organization Owner role.

		**Note**: If the user does not have an existing MongoDB Cloud account, they will be prompted to finish setting up an account upon accepting the invitation. If the user already has an account, they will still receive an invitation to access the organization.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param orgUserRequest Represents the MongoDB Cloud user to be created within the organization.
			@return CreateOrganizationUserApiRequest
	*/
	CreateOrganizationUser(ctx context.Context, orgId string, orgUserRequest *OrgUserRequest) CreateOrganizationUserApiRequest
	/*
		CreateOrganizationUser Create One MongoDB Cloud User in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrganizationUserApiParams - Parameters for the request
		@return CreateOrganizationUserApiRequest
	*/
	CreateOrganizationUserWithParams(ctx context.Context, args *CreateOrganizationUserApiParams) CreateOrganizationUserApiRequest

	// Method available only for mocking purposes
	CreateOrganizationUserExecute(r CreateOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error)

	/*
			CreateUser Create One MongoDB Cloud User

			Creates one MongoDB Cloud user account. A MongoDB Cloud user account grants access to only the MongoDB Cloud application. To grant database access, create a database user. MongoDB Cloud sends an email to the users you specify, inviting them to join the project. Invited users don't have access to the project until they accept the invitation. Invitations expire after 30 days.

		 MongoDB Cloud limits MongoDB Cloud user membership to a maximum of 250 MongoDB Cloud users per team. MongoDB Cloud limits MongoDB Cloud user membership to 500 MongoDB Cloud users per project and 500 MongoDB Cloud users per organization, which includes the combined membership of all projects in the organization. MongoDB Cloud raises an error if an operation exceeds these limits. For example, if you have an organization with five projects, and each project has 100 MongoDB Cloud users, and each MongoDB Cloud user belongs to only one project, you can't add any MongoDB Cloud users to this organization without first removing existing MongoDB Cloud users from the organization.

		 To use this resource, the requesting API Key can have any role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param cloudAppUser MongoDB Cloud user account to create.
			@return CreateUserApiRequest
	*/
	CreateUser(ctx context.Context, cloudAppUser *CloudAppUser) CreateUserApiRequest
	/*
		CreateUser Create One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateUserApiParams - Parameters for the request
		@return CreateUserApiRequest
	*/
	CreateUserWithParams(ctx context.Context, args *CreateUserApiParams) CreateUserApiRequest

	// Method available only for mocking purposes
	CreateUserExecute(r CreateUserApiRequest) (*CloudAppUser, *http.Response, error)

	/*
			GetOrganizationUser Return One MongoDB Cloud User in One Organization

			Returns information about the specified MongoDB Cloud user within the context of the specified organization. To use this resource, the requesting API Key must have the Organization Read Only role.

		**Note**: This resource can only be used to fetch information about MongoDB Cloud human users. To return information about an API Key, use the [Return One Organization API Key](#tag/Programmatic-API-Keys/operation/getApiKey) endpoint.

		**Note**: This resource does not return information about pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
			@return GetOrganizationUserApiRequest
	*/
	GetOrganizationUser(ctx context.Context, orgId string, userId string) GetOrganizationUserApiRequest
	/*
		GetOrganizationUser Return One MongoDB Cloud User in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationUserApiParams - Parameters for the request
		@return GetOrganizationUserApiRequest
	*/
	GetOrganizationUserWithParams(ctx context.Context, args *GetOrganizationUserApiParams) GetOrganizationUserApiRequest

	// Method available only for mocking purposes
	GetOrganizationUserExecute(r GetOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error)

	/*
		GetUser Return One MongoDB Cloud User using Its ID

		Returns the details for one MongoDB Cloud user account with the specified unique identifier for the user. You can't use this endpoint to return information on an API Key. To return information about an API Key, use the [Return One Organization](#tag/Organizations/operation/getOrganization) API Key endpoint. You can always retrieve your own user account. If you are the owner of a MongoDB Cloud organization or project, you can also retrieve the user profile for any user with membership in that organization or project. To use this resource, the requesting API Key can have any role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId Unique 24-hexadecimal digit string that identifies this user.
		@return GetUserApiRequest
	*/
	GetUser(ctx context.Context, userId string) GetUserApiRequest
	/*
		GetUser Return One MongoDB Cloud User using Its ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetUserApiParams - Parameters for the request
		@return GetUserApiRequest
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
	*/
	GetUserByUsername(ctx context.Context, userName string) GetUserByUsernameApiRequest
	/*
		GetUserByUsername Return One MongoDB Cloud User using Their Username


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetUserByUsernameApiParams - Parameters for the request
		@return GetUserByUsernameApiRequest
	*/
	GetUserByUsernameWithParams(ctx context.Context, args *GetUserByUsernameApiParams) GetUserByUsernameApiRequest

	// Method available only for mocking purposes
	GetUserByUsernameExecute(r GetUserByUsernameApiRequest) (*CloudAppUser, *http.Response, error)

	/*
			UpdateOrganizationUser Update One MongoDB Cloud User in One Organization

			Updates one MongoDB Cloud user in the specified organization. You can update an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

		**Note**: Only include the fields you wish to update in the request body. Supplying a field with an empty value will reset that field on the user.

		**Note**: This resource cannot be used to update pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
			@param orgUserUpdateRequest Represents the roles and teams to assign the MongoDB Cloud user.
			@return UpdateOrganizationUserApiRequest
	*/
	UpdateOrganizationUser(ctx context.Context, orgId string, userId string, orgUserUpdateRequest *OrgUserUpdateRequest) UpdateOrganizationUserApiRequest
	/*
		UpdateOrganizationUser Update One MongoDB Cloud User in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationUserApiParams - Parameters for the request
		@return UpdateOrganizationUserApiRequest
	*/
	UpdateOrganizationUserWithParams(ctx context.Context, args *UpdateOrganizationUserApiParams) UpdateOrganizationUserApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationUserExecute(r UpdateOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error)
}

// MongoDBCloudUsersApiService MongoDBCloudUsersApi service
type MongoDBCloudUsersApiService service

type AddOrganizationRoleApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	orgId      string
	userId     string
	role       *string
}

type AddOrganizationRoleApiParams struct {
	OrgId  string
	UserId string
	Role   *string
}

func (a *MongoDBCloudUsersApiService) AddOrganizationRoleWithParams(ctx context.Context, args *AddOrganizationRoleApiParams) AddOrganizationRoleApiRequest {
	return AddOrganizationRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		userId:     args.UserId,
		role:       args.Role,
	}
}

// Organization level role to assign to the MongoDB Cloud user.
func (r AddOrganizationRoleApiRequest) Role(role string) AddOrganizationRoleApiRequest {
	r.role = &role
	return r
}

func (r AddOrganizationRoleApiRequest) Execute() (*OrgUserResponse, *http.Response, error) {
	return r.ApiService.AddOrganizationRoleExecute(r)
}

/*
AddOrganizationRole Add One Organization Role to One MongoDB Cloud User

Adds one organization level role to the MongoDB Cloud user. You can add a role to an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

**Note**: This operation is atomic.

**Note**: This resource cannot be used to add a role to users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
	@return AddOrganizationRoleApiRequest
*/
func (a *MongoDBCloudUsersApiService) AddOrganizationRole(ctx context.Context, orgId string, userId string) AddOrganizationRoleApiRequest {
	return AddOrganizationRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		userId:     userId,
	}
}

// AddOrganizationRoleExecute executes the request
//
//	@return OrgUserResponse
func (a *MongoDBCloudUsersApiService) AddOrganizationRoleExecute(r AddOrganizationRoleApiRequest) (*OrgUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrgUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.AddOrganizationRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}:addRole"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.role == nil {
		return localVarReturnValue, nil, reportError("role is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type CreateOrganizationUserApiRequest struct {
	ctx            context.Context
	ApiService     MongoDBCloudUsersApi
	orgId          string
	orgUserRequest *OrgUserRequest
}

type CreateOrganizationUserApiParams struct {
	OrgId          string
	OrgUserRequest *OrgUserRequest
}

func (a *MongoDBCloudUsersApiService) CreateOrganizationUserWithParams(ctx context.Context, args *CreateOrganizationUserApiParams) CreateOrganizationUserApiRequest {
	return CreateOrganizationUserApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          args.OrgId,
		orgUserRequest: args.OrgUserRequest,
	}
}

func (r CreateOrganizationUserApiRequest) Execute() (*OrgUserResponse, *http.Response, error) {
	return r.ApiService.CreateOrganizationUserExecute(r)
}

/*
CreateOrganizationUser Create One MongoDB Cloud User in One Organization

Invites one new or existing MongoDB Cloud user to join the organization. The invitation to join the organization will be sent to the username provided and must be accepted within 30 days. To use this resource, the requesting API Key must have the Organization Owner role.

**Note**: If the user does not have an existing MongoDB Cloud account, they will be prompted to finish setting up an account upon accepting the invitation. If the user already has an account, they will still receive an invitation to access the organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateOrganizationUserApiRequest
*/
func (a *MongoDBCloudUsersApiService) CreateOrganizationUser(ctx context.Context, orgId string, orgUserRequest *OrgUserRequest) CreateOrganizationUserApiRequest {
	return CreateOrganizationUserApiRequest{
		ApiService:     a,
		ctx:            ctx,
		orgId:          orgId,
		orgUserRequest: orgUserRequest,
	}
}

// CreateOrganizationUserExecute executes the request
//
//	@return OrgUserResponse
func (a *MongoDBCloudUsersApiService) CreateOrganizationUserExecute(r CreateOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrgUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.CreateOrganizationUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgUserRequest == nil {
		return localVarReturnValue, nil, reportError("orgUserRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orgUserRequest
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
func (a *MongoDBCloudUsersApiService) CreateUserExecute(r CreateUserApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type GetOrganizationUserApiRequest struct {
	ctx        context.Context
	ApiService MongoDBCloudUsersApi
	orgId      string
	userId     string
}

type GetOrganizationUserApiParams struct {
	OrgId  string
	UserId string
}

func (a *MongoDBCloudUsersApiService) GetOrganizationUserWithParams(ctx context.Context, args *GetOrganizationUserApiParams) GetOrganizationUserApiRequest {
	return GetOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		userId:     args.UserId,
	}
}

func (r GetOrganizationUserApiRequest) Execute() (*OrgUserResponse, *http.Response, error) {
	return r.ApiService.GetOrganizationUserExecute(r)
}

/*
GetOrganizationUser Return One MongoDB Cloud User in One Organization

Returns information about the specified MongoDB Cloud user within the context of the specified organization. To use this resource, the requesting API Key must have the Organization Read Only role.

**Note**: This resource can only be used to fetch information about MongoDB Cloud human users. To return information about an API Key, use the [Return One Organization API Key](#tag/Programmatic-API-Keys/operation/getApiKey) endpoint.

**Note**: This resource does not return information about pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
	@return GetOrganizationUserApiRequest
*/
func (a *MongoDBCloudUsersApiService) GetOrganizationUser(ctx context.Context, orgId string, userId string) GetOrganizationUserApiRequest {
	return GetOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		userId:     userId,
	}
}

// GetOrganizationUserExecute executes the request
//
//	@return OrgUserResponse
func (a *MongoDBCloudUsersApiService) GetOrganizationUserExecute(r GetOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrgUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.GetOrganizationUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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
func (a *MongoDBCloudUsersApiService) GetUserExecute(r GetUserApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
func (a *MongoDBCloudUsersApiService) GetUserByUsernameExecute(r GetUserByUsernameApiRequest) (*CloudAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.GetUserByUsername")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/users/byName/{userName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userName"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type UpdateOrganizationUserApiRequest struct {
	ctx                  context.Context
	ApiService           MongoDBCloudUsersApi
	orgId                string
	userId               string
	orgUserUpdateRequest *OrgUserUpdateRequest
}

type UpdateOrganizationUserApiParams struct {
	OrgId                string
	UserId               string
	OrgUserUpdateRequest *OrgUserUpdateRequest
}

func (a *MongoDBCloudUsersApiService) UpdateOrganizationUserWithParams(ctx context.Context, args *UpdateOrganizationUserApiParams) UpdateOrganizationUserApiRequest {
	return UpdateOrganizationUserApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                args.OrgId,
		userId:               args.UserId,
		orgUserUpdateRequest: args.OrgUserUpdateRequest,
	}
}

func (r UpdateOrganizationUserApiRequest) Execute() (*OrgUserResponse, *http.Response, error) {
	return r.ApiService.UpdateOrganizationUserExecute(r)
}

/*
UpdateOrganizationUser Update One MongoDB Cloud User in One Organization

Updates one MongoDB Cloud user in the specified organization. You can update an active user or a user that has not yet accepted the invitation to join the organization. To use this resource, the requesting API Key must have the Organization Owner role.

**Note**: Only include the fields you wish to update in the request body. Supplying a field with an empty value will reset that field on the user.

**Note**: This resource cannot be used to update pending users invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.
	@return UpdateOrganizationUserApiRequest
*/
func (a *MongoDBCloudUsersApiService) UpdateOrganizationUser(ctx context.Context, orgId string, userId string, orgUserUpdateRequest *OrgUserUpdateRequest) UpdateOrganizationUserApiRequest {
	return UpdateOrganizationUserApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                orgId,
		userId:               userId,
		orgUserUpdateRequest: orgUserUpdateRequest,
	}
}

// UpdateOrganizationUserExecute executes the request
//
//	@return OrgUserResponse
func (a *MongoDBCloudUsersApiService) UpdateOrganizationUserExecute(r UpdateOrganizationUserApiRequest) (*OrgUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrgUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MongoDBCloudUsersApiService.UpdateOrganizationUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgUserUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("orgUserUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orgUserUpdateRequest
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
