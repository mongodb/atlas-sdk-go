// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OrganizationsApi interface {

	/*
		CreateOrganization Create One Organization

		Creates one organization in MongoDB Cloud and links it to the requesting Service Account's or API Key's organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role. The requesting Service Account's or API Key's organization must be a paying organization. To learn more, see Configure a Paying Organization in the MongoDB Atlas documentation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param createOrganizationRequest Organization that you want to create.
		@return CreateOrganizationApiRequest
	*/
	CreateOrganization(ctx context.Context, createOrganizationRequest *CreateOrganizationRequest) CreateOrganizationApiRequest
	/*
		CreateOrganization Create One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrganizationApiParams - Parameters for the request
		@return CreateOrganizationApiRequest
	*/
	CreateOrganizationWithParams(ctx context.Context, args *CreateOrganizationApiParams) CreateOrganizationApiRequest

	// Method available only for mocking purposes
	CreateOrganizationExecute(r CreateOrganizationApiRequest) (*CreateOrganizationResponse, *http.Response, error)

	/*
			CreateOrganizationInvitation Invite One MongoDB Cloud User to One Atlas Organization

			Invites one MongoDB Cloud user to join the specified organization. The user must accept the invitation to access information within the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: Invitation management APIs are deprecated. Use Add One MongoDB Cloud User to One Organization to invite a user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param organizationInvitationRequest Invites one MongoDB Cloud user to join the specified organization.
			@return CreateOrganizationInvitationApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	CreateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) CreateOrganizationInvitationApiRequest
	/*
		CreateOrganizationInvitation Invite One MongoDB Cloud User to One Atlas Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrganizationInvitationApiParams - Parameters for the request
		@return CreateOrganizationInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	CreateOrganizationInvitationWithParams(ctx context.Context, args *CreateOrganizationInvitationApiParams) CreateOrganizationInvitationApiRequest

	// Method available only for mocking purposes
	CreateOrganizationInvitationExecute(r CreateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
			DeleteOrganization Remove One Organization

			Removes one specified organization. MongoDB Cloud imposes the following limits on this resource:

		 - Organizations with active projects cannot be removed.
		 - All projects in the organization must be removed before you can remove the organization.
		 To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return DeleteOrganizationApiRequest
	*/
	DeleteOrganization(ctx context.Context, orgId string) DeleteOrganizationApiRequest
	/*
		DeleteOrganization Remove One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrganizationApiParams - Parameters for the request
		@return DeleteOrganizationApiRequest
	*/
	DeleteOrganizationWithParams(ctx context.Context, args *DeleteOrganizationApiParams) DeleteOrganizationApiRequest

	// Method available only for mocking purposes
	DeleteOrganizationExecute(r DeleteOrganizationApiRequest) (*http.Response, error)

	/*
			DeleteOrganizationInvitation Remove One Organization Invitation

			Cancels one pending invitation sent to the specified MongoDB Cloud user to join an organization. You can't cancel an invitation that the user accepted. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: Invitation management APIs are deprecated. Use Remove One MongoDB Cloud User From One Organization to remove a pending user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
			@return DeleteOrganizationInvitationApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	DeleteOrganizationInvitation(ctx context.Context, orgId string, invitationId string) DeleteOrganizationInvitationApiRequest
	/*
		DeleteOrganizationInvitation Remove One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrganizationInvitationApiParams - Parameters for the request
		@return DeleteOrganizationInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	DeleteOrganizationInvitationWithParams(ctx context.Context, args *DeleteOrganizationInvitationApiParams) DeleteOrganizationInvitationApiRequest

	// Method available only for mocking purposes
	DeleteOrganizationInvitationExecute(r DeleteOrganizationInvitationApiRequest) (*http.Response, error)

	/*
		GetOrganization Return One Organization

		Returns one organization to which the requesting Service Account or API Key has access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetOrganizationApiRequest
	*/
	GetOrganization(ctx context.Context, orgId string) GetOrganizationApiRequest
	/*
		GetOrganization Return One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationApiParams - Parameters for the request
		@return GetOrganizationApiRequest
	*/
	GetOrganizationWithParams(ctx context.Context, args *GetOrganizationApiParams) GetOrganizationApiRequest

	// Method available only for mocking purposes
	GetOrganizationExecute(r GetOrganizationApiRequest) (*AtlasOrganization, *http.Response, error)

	/*
			GetOrganizationInvitation Return One Organization Invitation

			Returns the details of one pending invitation to the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: Invitation management APIs are deprecated. Use Return One MongoDB Cloud User in One Organization to return a pending user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
			@return GetOrganizationInvitationApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	GetOrganizationInvitation(ctx context.Context, orgId string, invitationId string) GetOrganizationInvitationApiRequest
	/*
		GetOrganizationInvitation Return One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationInvitationApiParams - Parameters for the request
		@return GetOrganizationInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	GetOrganizationInvitationWithParams(ctx context.Context, args *GetOrganizationInvitationApiParams) GetOrganizationInvitationApiRequest

	// Method available only for mocking purposes
	GetOrganizationInvitationExecute(r GetOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		GetOrganizationSettings Return Settings for One Organization

		Returns details about the specified organization's settings. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetOrganizationSettingsApiRequest
	*/
	GetOrganizationSettings(ctx context.Context, orgId string) GetOrganizationSettingsApiRequest
	/*
		GetOrganizationSettings Return Settings for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationSettingsApiParams - Parameters for the request
		@return GetOrganizationSettingsApiRequest
	*/
	GetOrganizationSettingsWithParams(ctx context.Context, args *GetOrganizationSettingsApiParams) GetOrganizationSettingsApiRequest

	// Method available only for mocking purposes
	GetOrganizationSettingsExecute(r GetOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error)

	/*
			ListOrganizationInvitations Return All Organization Invitations

			Returns all pending invitations to the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: Invitation management APIs are deprecated. Use Return All MongoDB Cloud Users in One Organization and filter by `orgMembershipStatus` to return all pending users.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrganizationInvitationsApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	ListOrganizationInvitations(ctx context.Context, orgId string) ListOrganizationInvitationsApiRequest
	/*
		ListOrganizationInvitations Return All Organization Invitations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationInvitationsApiParams - Parameters for the request
		@return ListOrganizationInvitationsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	ListOrganizationInvitationsWithParams(ctx context.Context, args *ListOrganizationInvitationsApiParams) ListOrganizationInvitationsApiRequest

	// Method available only for mocking purposes
	ListOrganizationInvitationsExecute(r ListOrganizationInvitationsApiRequest) ([]OrganizationInvitation, *http.Response, error)

	/*
			ListOrganizationProjects Return All Projects in One Organization

			Returns multiple projects in the specified organization. Each organization can have multiple projects. Use projects to:

		- Isolate different environments, such as development, test, or production environments, from each other.
		- Associate different MongoDB Cloud users or teams with different environments, or give different permission to MongoDB Cloud users in different environments.
		- Maintain separate cluster security configurations.
		- Create different alert settings.

		To use this resource, the requesting Service Account or API Key must have the Organization Member role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrganizationProjectsApiRequest
	*/
	ListOrganizationProjects(ctx context.Context, orgId string) ListOrganizationProjectsApiRequest
	/*
		ListOrganizationProjects Return All Projects in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationProjectsApiParams - Parameters for the request
		@return ListOrganizationProjectsApiRequest
	*/
	ListOrganizationProjectsWithParams(ctx context.Context, args *ListOrganizationProjectsApiParams) ListOrganizationProjectsApiRequest

	// Method available only for mocking purposes
	ListOrganizationProjectsExecute(r ListOrganizationProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error)

	/*
		ListOrganizations Return All Organizations

		Returns all organizations to which the requesting Service Account or API Key has access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListOrganizationsApiRequest
	*/
	ListOrganizations(ctx context.Context) ListOrganizationsApiRequest
	/*
		ListOrganizations Return All Organizations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationsApiParams - Parameters for the request
		@return ListOrganizationsApiRequest
	*/
	ListOrganizationsWithParams(ctx context.Context, args *ListOrganizationsApiParams) ListOrganizationsApiRequest

	// Method available only for mocking purposes
	ListOrganizationsExecute(r ListOrganizationsApiRequest) (*PaginatedOrganization, *http.Response, error)

	/*
		UpdateOrganization Update One Organization

		Updates one organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param atlasOrganization Details to update on the specified organization.
		@return UpdateOrganizationApiRequest
	*/
	UpdateOrganization(ctx context.Context, orgId string, atlasOrganization *AtlasOrganization) UpdateOrganizationApiRequest
	/*
		UpdateOrganization Update One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationApiParams - Parameters for the request
		@return UpdateOrganizationApiRequest
	*/
	UpdateOrganizationWithParams(ctx context.Context, args *UpdateOrganizationApiParams) UpdateOrganizationApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationExecute(r UpdateOrganizationApiRequest) (*AtlasOrganization, *http.Response, error)

	/*
			UpdateOrganizationInvitation Update One Organization Invitation

			Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the username of the invited user. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**:  Invitation management are deprecated. Use Update One MongoDB Cloud User in One Organization to update a pending user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param organizationInvitationRequest Updates the details of one pending invitation to the specified organization.
			@return UpdateOrganizationInvitationApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) UpdateOrganizationInvitationApiRequest
	/*
		UpdateOrganizationInvitation Update One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationInvitationApiParams - Parameters for the request
		@return UpdateOrganizationInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationInvitationWithParams(ctx context.Context, args *UpdateOrganizationInvitationApiParams) UpdateOrganizationInvitationApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationInvitationExecute(r UpdateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
			UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID

			Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the unique identification string for that invitation. Use the Return All Organization Invitations endpoint to retrieve IDs for all pending organization invitations. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: Invitation management APIs are deprecated. Use Update One MongoDB Cloud User in One Organization to update a pending user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
			@param organizationInvitationUpdateRequest Updates the details of one pending invitation to the specified organization.
			@return UpdateOrganizationInvitationByIdApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationInvitationById(ctx context.Context, orgId string, invitationId string, organizationInvitationUpdateRequest *OrganizationInvitationUpdateRequest) UpdateOrganizationInvitationByIdApiRequest
	/*
		UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationInvitationByIdApiParams - Parameters for the request
		@return UpdateOrganizationInvitationByIdApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationInvitationByIdWithParams(ctx context.Context, args *UpdateOrganizationInvitationByIdApiParams) UpdateOrganizationInvitationByIdApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationInvitationByIdExecute(r UpdateOrganizationInvitationByIdApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		UpdateOrganizationRoles Update Organization Roles for One MongoDB Cloud User

		Updates the roles of the specified user in the specified organization. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
		@param updateOrgRolesForUser Roles to update for the specified user.
		@return UpdateOrganizationRolesApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationRoles(ctx context.Context, orgId string, userId string, updateOrgRolesForUser *UpdateOrgRolesForUser) UpdateOrganizationRolesApiRequest
	/*
		UpdateOrganizationRoles Update Organization Roles for One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationRolesApiParams - Parameters for the request
		@return UpdateOrganizationRolesApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for OrganizationsApi
	*/
	UpdateOrganizationRolesWithParams(ctx context.Context, args *UpdateOrganizationRolesApiParams) UpdateOrganizationRolesApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationRolesExecute(r UpdateOrganizationRolesApiRequest) (*UpdateOrgRolesForUser, *http.Response, error)

	/*
		UpdateOrganizationSettings Update Settings for One Organization

		Updates the organization's settings. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param organizationSettings Details to update on the specified organization's settings.
		@return UpdateOrganizationSettingsApiRequest
	*/
	UpdateOrganizationSettings(ctx context.Context, orgId string, organizationSettings *OrganizationSettings) UpdateOrganizationSettingsApiRequest
	/*
		UpdateOrganizationSettings Update Settings for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationSettingsApiParams - Parameters for the request
		@return UpdateOrganizationSettingsApiRequest
	*/
	UpdateOrganizationSettingsWithParams(ctx context.Context, args *UpdateOrganizationSettingsApiParams) UpdateOrganizationSettingsApiRequest

	// Method available only for mocking purposes
	UpdateOrganizationSettingsExecute(r UpdateOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error)
}

// OrganizationsApiService OrganizationsApi service
type OrganizationsApiService service

type CreateOrganizationApiRequest struct {
	ctx                       context.Context
	ApiService                OrganizationsApi
	createOrganizationRequest *CreateOrganizationRequest
}

type CreateOrganizationApiParams struct {
	CreateOrganizationRequest *CreateOrganizationRequest
}

func (a *OrganizationsApiService) CreateOrganizationWithParams(ctx context.Context, args *CreateOrganizationApiParams) CreateOrganizationApiRequest {
	return CreateOrganizationApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		createOrganizationRequest: args.CreateOrganizationRequest,
	}
}

func (r CreateOrganizationApiRequest) Execute() (*CreateOrganizationResponse, *http.Response, error) {
	return r.ApiService.CreateOrganizationExecute(r)
}

/*
CreateOrganization Create One Organization

Creates one organization in MongoDB Cloud and links it to the requesting Service Account's or API Key's organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role. The requesting Service Account's or API Key's organization must be a paying organization. To learn more, see Configure a Paying Organization in the MongoDB Atlas documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateOrganizationApiRequest
*/
func (a *OrganizationsApiService) CreateOrganization(ctx context.Context, createOrganizationRequest *CreateOrganizationRequest) CreateOrganizationApiRequest {
	return CreateOrganizationApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		createOrganizationRequest: createOrganizationRequest,
	}
}

// CreateOrganizationExecute executes the request
//
//	@return CreateOrganizationResponse
func (a *OrganizationsApiService) CreateOrganizationExecute(r CreateOrganizationApiRequest) (*CreateOrganizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CreateOrganizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.CreateOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationRequest
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

type CreateOrganizationInvitationApiRequest struct {
	ctx                           context.Context
	ApiService                    OrganizationsApi
	orgId                         string
	organizationInvitationRequest *OrganizationInvitationRequest
}

type CreateOrganizationInvitationApiParams struct {
	OrgId                         string
	OrganizationInvitationRequest *OrganizationInvitationRequest
}

func (a *OrganizationsApiService) CreateOrganizationInvitationWithParams(ctx context.Context, args *CreateOrganizationInvitationApiParams) CreateOrganizationInvitationApiRequest {
	return CreateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		organizationInvitationRequest: args.OrganizationInvitationRequest,
	}
}

func (r CreateOrganizationInvitationApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.CreateOrganizationInvitationExecute(r)
}

/*
CreateOrganizationInvitation Invite One MongoDB Cloud User to One Atlas Organization

Invites one MongoDB Cloud user to join the specified organization. The user must accept the invitation to access information within the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: Invitation management APIs are deprecated. Use Add One MongoDB Cloud User to One Organization to invite a user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateOrganizationInvitationApiRequest

Deprecated
*/
func (a *OrganizationsApiService) CreateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) CreateOrganizationInvitationApiRequest {
	return CreateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		organizationInvitationRequest: organizationInvitationRequest,
	}
}

// CreateOrganizationInvitationExecute executes the request
//
//	@return OrganizationInvitation
//
// Deprecated
func (a *OrganizationsApiService) CreateOrganizationInvitationExecute(r CreateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.CreateOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("organizationInvitationRequest is required and must be specified")
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
	localVarPostBody = r.organizationInvitationRequest
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

type DeleteOrganizationApiRequest struct {
	ctx        context.Context
	ApiService OrganizationsApi
	orgId      string
}

type DeleteOrganizationApiParams struct {
	OrgId string
}

func (a *OrganizationsApiService) DeleteOrganizationWithParams(ctx context.Context, args *DeleteOrganizationApiParams) DeleteOrganizationApiRequest {
	return DeleteOrganizationApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r DeleteOrganizationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationExecute(r)
}

/*
DeleteOrganization Remove One Organization

Removes one specified organization. MongoDB Cloud imposes the following limits on this resource:

  - Organizations with active projects cannot be removed.

  - All projects in the organization must be removed before you can remove the organization.
    To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
    @return DeleteOrganizationApiRequest
*/
func (a *OrganizationsApiService) DeleteOrganization(ctx context.Context, orgId string) DeleteOrganizationApiRequest {
	return DeleteOrganizationApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// DeleteOrganizationExecute executes the request
func (a *OrganizationsApiService) DeleteOrganizationExecute(r DeleteOrganizationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.DeleteOrganization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

type DeleteOrganizationInvitationApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	orgId        string
	invitationId string
}

type DeleteOrganizationInvitationApiParams struct {
	OrgId        string
	InvitationId string
}

func (a *OrganizationsApiService) DeleteOrganizationInvitationWithParams(ctx context.Context, args *DeleteOrganizationInvitationApiParams) DeleteOrganizationInvitationApiRequest {
	return DeleteOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		invitationId: args.InvitationId,
	}
}

func (r DeleteOrganizationInvitationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationInvitationExecute(r)
}

/*
DeleteOrganizationInvitation Remove One Organization Invitation

Cancels one pending invitation sent to the specified MongoDB Cloud user to join an organization. You can't cancel an invitation that the user accepted. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: Invitation management APIs are deprecated. Use Remove One MongoDB Cloud User From One Organization to remove a pending user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return DeleteOrganizationInvitationApiRequest

Deprecated
*/
func (a *OrganizationsApiService) DeleteOrganizationInvitation(ctx context.Context, orgId string, invitationId string) DeleteOrganizationInvitationApiRequest {
	return DeleteOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		invitationId: invitationId,
	}
}

// DeleteOrganizationInvitationExecute executes the request
// Deprecated
func (a *OrganizationsApiService) DeleteOrganizationInvitationExecute(r DeleteOrganizationInvitationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.DeleteOrganizationInvitation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invitationId == "" {
		return nil, reportError("invitationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(r.invitationId), -1)

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

type GetOrganizationApiRequest struct {
	ctx        context.Context
	ApiService OrganizationsApi
	orgId      string
}

type GetOrganizationApiParams struct {
	OrgId string
}

func (a *OrganizationsApiService) GetOrganizationWithParams(ctx context.Context, args *GetOrganizationApiParams) GetOrganizationApiRequest {
	return GetOrganizationApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetOrganizationApiRequest) Execute() (*AtlasOrganization, *http.Response, error) {
	return r.ApiService.GetOrganizationExecute(r)
}

/*
GetOrganization Return One Organization

Returns one organization to which the requesting Service Account or API Key has access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetOrganizationApiRequest
*/
func (a *OrganizationsApiService) GetOrganization(ctx context.Context, orgId string) GetOrganizationApiRequest {
	return GetOrganizationApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetOrganizationExecute executes the request
//
//	@return AtlasOrganization
func (a *OrganizationsApiService) GetOrganizationExecute(r GetOrganizationApiRequest) (*AtlasOrganization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasOrganization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

type GetOrganizationInvitationApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	orgId        string
	invitationId string
}

type GetOrganizationInvitationApiParams struct {
	OrgId        string
	InvitationId string
}

func (a *OrganizationsApiService) GetOrganizationInvitationWithParams(ctx context.Context, args *GetOrganizationInvitationApiParams) GetOrganizationInvitationApiRequest {
	return GetOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		invitationId: args.InvitationId,
	}
}

func (r GetOrganizationInvitationApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.GetOrganizationInvitationExecute(r)
}

/*
GetOrganizationInvitation Return One Organization Invitation

Returns the details of one pending invitation to the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: Invitation management APIs are deprecated. Use Return One MongoDB Cloud User in One Organization to return a pending user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return GetOrganizationInvitationApiRequest

Deprecated
*/
func (a *OrganizationsApiService) GetOrganizationInvitation(ctx context.Context, orgId string, invitationId string) GetOrganizationInvitationApiRequest {
	return GetOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		invitationId: invitationId,
	}
}

// GetOrganizationInvitationExecute executes the request
//
//	@return OrganizationInvitation
//
// Deprecated
func (a *OrganizationsApiService) GetOrganizationInvitationExecute(r GetOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invitationId == "" {
		return localVarReturnValue, nil, reportError("invitationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(r.invitationId), -1)

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

type GetOrganizationSettingsApiRequest struct {
	ctx        context.Context
	ApiService OrganizationsApi
	orgId      string
}

type GetOrganizationSettingsApiParams struct {
	OrgId string
}

func (a *OrganizationsApiService) GetOrganizationSettingsWithParams(ctx context.Context, args *GetOrganizationSettingsApiParams) GetOrganizationSettingsApiRequest {
	return GetOrganizationSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetOrganizationSettingsApiRequest) Execute() (*OrganizationSettings, *http.Response, error) {
	return r.ApiService.GetOrganizationSettingsExecute(r)
}

/*
GetOrganizationSettings Return Settings for One Organization

Returns details about the specified organization's settings. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetOrganizationSettingsApiRequest
*/
func (a *OrganizationsApiService) GetOrganizationSettings(ctx context.Context, orgId string) GetOrganizationSettingsApiRequest {
	return GetOrganizationSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetOrganizationSettingsExecute executes the request
//
//	@return OrganizationSettings
func (a *OrganizationsApiService) GetOrganizationSettingsExecute(r GetOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganizationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/settings"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

type ListOrganizationInvitationsApiRequest struct {
	ctx        context.Context
	ApiService OrganizationsApi
	orgId      string
	username   *string
}

type ListOrganizationInvitationsApiParams struct {
	OrgId    string
	Username *string
}

func (a *OrganizationsApiService) ListOrganizationInvitationsWithParams(ctx context.Context, args *ListOrganizationInvitationsApiParams) ListOrganizationInvitationsApiRequest {
	return ListOrganizationInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		username:   args.Username,
	}
}

// Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.
func (r ListOrganizationInvitationsApiRequest) Username(username string) ListOrganizationInvitationsApiRequest {
	r.username = &username
	return r
}

func (r ListOrganizationInvitationsApiRequest) Execute() ([]OrganizationInvitation, *http.Response, error) {
	return r.ApiService.ListOrganizationInvitationsExecute(r)
}

/*
ListOrganizationInvitations Return All Organization Invitations

Returns all pending invitations to the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: Invitation management APIs are deprecated. Use Return All MongoDB Cloud Users in One Organization and filter by `orgMembershipStatus` to return all pending users.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationInvitationsApiRequest

Deprecated
*/
func (a *OrganizationsApiService) ListOrganizationInvitations(ctx context.Context, orgId string) ListOrganizationInvitationsApiRequest {
	return ListOrganizationInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrganizationInvitationsExecute executes the request
//
//	@return []OrganizationInvitation
//
// Deprecated
func (a *OrganizationsApiService) ListOrganizationInvitationsExecute(r ListOrganizationInvitationsApiRequest) ([]OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizationInvitations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ListOrganizationProjectsApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	orgId        string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	name         *string
}

type ListOrganizationProjectsApiParams struct {
	OrgId        string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Name         *string
}

func (a *OrganizationsApiService) ListOrganizationProjectsWithParams(ctx context.Context, args *ListOrganizationProjectsApiParams) ListOrganizationProjectsApiRequest {
	return ListOrganizationProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		name:         args.Name,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOrganizationProjectsApiRequest) IncludeCount(includeCount bool) ListOrganizationProjectsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListOrganizationProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListOrganizationProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrganizationProjectsApiRequest) PageNum(pageNum int) ListOrganizationProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.
func (r ListOrganizationProjectsApiRequest) Name(name string) ListOrganizationProjectsApiRequest {
	r.name = &name
	return r
}

func (r ListOrganizationProjectsApiRequest) Execute() (*PaginatedAtlasGroup, *http.Response, error) {
	return r.ApiService.ListOrganizationProjectsExecute(r)
}

/*
ListOrganizationProjects Return All Projects in One Organization

Returns multiple projects in the specified organization. Each organization can have multiple projects. Use projects to:

- Isolate different environments, such as development, test, or production environments, from each other.
- Associate different MongoDB Cloud users or teams with different environments, or give different permission to MongoDB Cloud users in different environments.
- Maintain separate cluster security configurations.
- Create different alert settings.

To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationProjectsApiRequest
*/
func (a *OrganizationsApiService) ListOrganizationProjects(ctx context.Context, orgId string) ListOrganizationProjectsApiRequest {
	return ListOrganizationProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrganizationProjectsExecute executes the request
//
//	@return PaginatedAtlasGroup
func (a *OrganizationsApiService) ListOrganizationProjectsExecute(r ListOrganizationProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizationProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/groups"
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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
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

type ListOrganizationsApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	name         *string
}

type ListOrganizationsApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Name         *string
}

func (a *OrganizationsApiService) ListOrganizationsWithParams(ctx context.Context, args *ListOrganizationsApiParams) ListOrganizationsApiRequest {
	return ListOrganizationsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		name:         args.Name,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOrganizationsApiRequest) IncludeCount(includeCount bool) ListOrganizationsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListOrganizationsApiRequest) ItemsPerPage(itemsPerPage int) ListOrganizationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrganizationsApiRequest) PageNum(pageNum int) ListOrganizationsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.
func (r ListOrganizationsApiRequest) Name(name string) ListOrganizationsApiRequest {
	r.name = &name
	return r
}

func (r ListOrganizationsApiRequest) Execute() (*PaginatedOrganization, *http.Response, error) {
	return r.ApiService.ListOrganizationsExecute(r)
}

/*
ListOrganizations Return All Organizations

Returns all organizations to which the requesting Service Account or API Key has access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListOrganizationsApiRequest
*/
func (a *OrganizationsApiService) ListOrganizations(ctx context.Context) ListOrganizationsApiRequest {
	return ListOrganizationsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListOrganizationsExecute executes the request
//
//	@return PaginatedOrganization
func (a *OrganizationsApiService) ListOrganizationsExecute(r ListOrganizationsApiRequest) (*PaginatedOrganization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedOrganization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs"

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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
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

type UpdateOrganizationApiRequest struct {
	ctx               context.Context
	ApiService        OrganizationsApi
	orgId             string
	atlasOrganization *AtlasOrganization
}

type UpdateOrganizationApiParams struct {
	OrgId             string
	AtlasOrganization *AtlasOrganization
}

func (a *OrganizationsApiService) UpdateOrganizationWithParams(ctx context.Context, args *UpdateOrganizationApiParams) UpdateOrganizationApiRequest {
	return UpdateOrganizationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		orgId:             args.OrgId,
		atlasOrganization: args.AtlasOrganization,
	}
}

func (r UpdateOrganizationApiRequest) Execute() (*AtlasOrganization, *http.Response, error) {
	return r.ApiService.UpdateOrganizationExecute(r)
}

/*
UpdateOrganization Update One Organization

Updates one organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateOrganizationApiRequest
*/
func (a *OrganizationsApiService) UpdateOrganization(ctx context.Context, orgId string, atlasOrganization *AtlasOrganization) UpdateOrganizationApiRequest {
	return UpdateOrganizationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		orgId:             orgId,
		atlasOrganization: atlasOrganization,
	}
}

// UpdateOrganizationExecute executes the request
//
//	@return AtlasOrganization
func (a *OrganizationsApiService) UpdateOrganizationExecute(r UpdateOrganizationApiRequest) (*AtlasOrganization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasOrganization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.atlasOrganization == nil {
		return localVarReturnValue, nil, reportError("atlasOrganization is required and must be specified")
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
	localVarPostBody = r.atlasOrganization
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

type UpdateOrganizationInvitationApiRequest struct {
	ctx                           context.Context
	ApiService                    OrganizationsApi
	orgId                         string
	organizationInvitationRequest *OrganizationInvitationRequest
}

type UpdateOrganizationInvitationApiParams struct {
	OrgId                         string
	OrganizationInvitationRequest *OrganizationInvitationRequest
}

func (a *OrganizationsApiService) UpdateOrganizationInvitationWithParams(ctx context.Context, args *UpdateOrganizationInvitationApiParams) UpdateOrganizationInvitationApiRequest {
	return UpdateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		organizationInvitationRequest: args.OrganizationInvitationRequest,
	}
}

func (r UpdateOrganizationInvitationApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.UpdateOrganizationInvitationExecute(r)
}

/*
UpdateOrganizationInvitation Update One Organization Invitation

Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the username of the invited user. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**:  Invitation management are deprecated. Use Update One MongoDB Cloud User in One Organization to update a pending user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateOrganizationInvitationApiRequest

Deprecated
*/
func (a *OrganizationsApiService) UpdateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) UpdateOrganizationInvitationApiRequest {
	return UpdateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		organizationInvitationRequest: organizationInvitationRequest,
	}
}

// UpdateOrganizationInvitationExecute executes the request
//
//	@return OrganizationInvitation
//
// Deprecated
func (a *OrganizationsApiService) UpdateOrganizationInvitationExecute(r UpdateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("organizationInvitationRequest is required and must be specified")
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
	localVarPostBody = r.organizationInvitationRequest
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

type UpdateOrganizationInvitationByIdApiRequest struct {
	ctx                                 context.Context
	ApiService                          OrganizationsApi
	orgId                               string
	invitationId                        string
	organizationInvitationUpdateRequest *OrganizationInvitationUpdateRequest
}

type UpdateOrganizationInvitationByIdApiParams struct {
	OrgId                               string
	InvitationId                        string
	OrganizationInvitationUpdateRequest *OrganizationInvitationUpdateRequest
}

func (a *OrganizationsApiService) UpdateOrganizationInvitationByIdWithParams(ctx context.Context, args *UpdateOrganizationInvitationByIdApiParams) UpdateOrganizationInvitationByIdApiRequest {
	return UpdateOrganizationInvitationByIdApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		orgId:                               args.OrgId,
		invitationId:                        args.InvitationId,
		organizationInvitationUpdateRequest: args.OrganizationInvitationUpdateRequest,
	}
}

func (r UpdateOrganizationInvitationByIdApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.UpdateOrganizationInvitationByIdExecute(r)
}

/*
UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID

Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the unique identification string for that invitation. Use the Return All Organization Invitations endpoint to retrieve IDs for all pending organization invitations. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: Invitation management APIs are deprecated. Use Update One MongoDB Cloud User in One Organization to update a pending user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return UpdateOrganizationInvitationByIdApiRequest

Deprecated
*/
func (a *OrganizationsApiService) UpdateOrganizationInvitationById(ctx context.Context, orgId string, invitationId string, organizationInvitationUpdateRequest *OrganizationInvitationUpdateRequest) UpdateOrganizationInvitationByIdApiRequest {
	return UpdateOrganizationInvitationByIdApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		orgId:                               orgId,
		invitationId:                        invitationId,
		organizationInvitationUpdateRequest: organizationInvitationUpdateRequest,
	}
}

// UpdateOrganizationInvitationByIdExecute executes the request
//
//	@return OrganizationInvitation
//
// Deprecated
func (a *OrganizationsApiService) UpdateOrganizationInvitationByIdExecute(r UpdateOrganizationInvitationByIdApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationInvitationById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invitationId == "" {
		return localVarReturnValue, nil, reportError("invitationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(r.invitationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationInvitationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("organizationInvitationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.organizationInvitationUpdateRequest
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

type UpdateOrganizationRolesApiRequest struct {
	ctx                   context.Context
	ApiService            OrganizationsApi
	orgId                 string
	userId                string
	updateOrgRolesForUser *UpdateOrgRolesForUser
}

type UpdateOrganizationRolesApiParams struct {
	OrgId                 string
	UserId                string
	UpdateOrgRolesForUser *UpdateOrgRolesForUser
}

func (a *OrganizationsApiService) UpdateOrganizationRolesWithParams(ctx context.Context, args *UpdateOrganizationRolesApiParams) UpdateOrganizationRolesApiRequest {
	return UpdateOrganizationRolesApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		orgId:                 args.OrgId,
		userId:                args.UserId,
		updateOrgRolesForUser: args.UpdateOrgRolesForUser,
	}
}

func (r UpdateOrganizationRolesApiRequest) Execute() (*UpdateOrgRolesForUser, *http.Response, error) {
	return r.ApiService.UpdateOrganizationRolesExecute(r)
}

/*
UpdateOrganizationRoles Update Organization Roles for One MongoDB Cloud User

Updates the roles of the specified user in the specified organization. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization User Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
	@return UpdateOrganizationRolesApiRequest

Deprecated
*/
func (a *OrganizationsApiService) UpdateOrganizationRoles(ctx context.Context, orgId string, userId string, updateOrgRolesForUser *UpdateOrgRolesForUser) UpdateOrganizationRolesApiRequest {
	return UpdateOrganizationRolesApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		orgId:                 orgId,
		userId:                userId,
		updateOrgRolesForUser: updateOrgRolesForUser,
	}
}

// UpdateOrganizationRolesExecute executes the request
//
//	@return UpdateOrgRolesForUser
//
// Deprecated
func (a *OrganizationsApiService) UpdateOrganizationRolesExecute(r UpdateOrganizationRolesApiRequest) (*UpdateOrgRolesForUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UpdateOrgRolesForUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}/roles"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.userId == "" {
		return localVarReturnValue, nil, reportError("userId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(r.userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrgRolesForUser == nil {
		return localVarReturnValue, nil, reportError("updateOrgRolesForUser is required and must be specified")
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
	localVarPostBody = r.updateOrgRolesForUser
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

type UpdateOrganizationSettingsApiRequest struct {
	ctx                  context.Context
	ApiService           OrganizationsApi
	orgId                string
	organizationSettings *OrganizationSettings
}

type UpdateOrganizationSettingsApiParams struct {
	OrgId                string
	OrganizationSettings *OrganizationSettings
}

func (a *OrganizationsApiService) UpdateOrganizationSettingsWithParams(ctx context.Context, args *UpdateOrganizationSettingsApiParams) UpdateOrganizationSettingsApiRequest {
	return UpdateOrganizationSettingsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                args.OrgId,
		organizationSettings: args.OrganizationSettings,
	}
}

func (r UpdateOrganizationSettingsApiRequest) Execute() (*OrganizationSettings, *http.Response, error) {
	return r.ApiService.UpdateOrganizationSettingsExecute(r)
}

/*
UpdateOrganizationSettings Update Settings for One Organization

Updates the organization's settings. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateOrganizationSettingsApiRequest
*/
func (a *OrganizationsApiService) UpdateOrganizationSettings(ctx context.Context, orgId string, organizationSettings *OrganizationSettings) UpdateOrganizationSettingsApiRequest {
	return UpdateOrganizationSettingsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		orgId:                orgId,
		organizationSettings: organizationSettings,
	}
}

// UpdateOrganizationSettingsExecute executes the request
//
//	@return OrganizationSettings
func (a *OrganizationsApiService) UpdateOrganizationSettingsExecute(r UpdateOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/settings"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationSettings == nil {
		return localVarReturnValue, nil, reportError("organizationSettings is required and must be specified")
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
	localVarPostBody = r.organizationSettings
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
