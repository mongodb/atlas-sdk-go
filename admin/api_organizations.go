// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OrganizationsApi interface {

	/*
		CreateOrganization Create One Organization

		Creates one organization in MongoDB Cloud and links it to the requesting API Key's organization. To use this resource, the requesting API Key must have the Organization Owner role. The requesting API Key's organization must be a paying organization. To learn more, see [Configure a Paying Organization](https://www.mongodb.com/docs/atlas/billing/#configure-a-paying-organization) in the MongoDB Atlas documentation. This resource doesn't require the API Key to have an API Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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

	// Interface only available internally
	createOrganizationExecute(r CreateOrganizationApiRequest) (*CreateOrganizationResponse, *http.Response, error)

	/*
		CreateOrganizationInvitation Invite One MongoDB Cloud User to Join One Atlas Organization

		Invites one MongoDB Cloud user to join the specified organization. The user must accept the invitation to access information within the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return CreateOrganizationInvitationApiRequest
	*/
	CreateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) CreateOrganizationInvitationApiRequest
	/*
		CreateOrganizationInvitation Invite One MongoDB Cloud User to Join One Atlas Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrganizationInvitationApiParams - Parameters for the request
		@return CreateOrganizationInvitationApiRequest
	*/
	CreateOrganizationInvitationWithParams(ctx context.Context, args *CreateOrganizationInvitationApiParams) CreateOrganizationInvitationApiRequest

	// Interface only available internally
	createOrganizationInvitationExecute(r CreateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
			DeleteOrganization Remove One Organization

			Removes one specified organization. MongoDB Cloud imposes the following limits on this resource:

		 - Organizations with active projects cannot be removed.
		 - All projects in the organization must be removed before you can remove the organization.
		 To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	deleteOrganizationExecute(r DeleteOrganizationApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteOrganizationInvitation Cancel One Organization Invitation

		Cancels one pending invitation sent to the specified MongoDB Cloud user to join an organization. You can't cancel an invitation that the user accepted. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return DeleteOrganizationInvitationApiRequest
	*/
	DeleteOrganizationInvitation(ctx context.Context, orgId string, invitationId string) DeleteOrganizationInvitationApiRequest
	/*
		DeleteOrganizationInvitation Cancel One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrganizationInvitationApiParams - Parameters for the request
		@return DeleteOrganizationInvitationApiRequest
	*/
	DeleteOrganizationInvitationWithParams(ctx context.Context, args *DeleteOrganizationInvitationApiParams) DeleteOrganizationInvitationApiRequest

	// Interface only available internally
	deleteOrganizationInvitationExecute(r DeleteOrganizationInvitationApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetOrganization Return One Organization

		Returns one organization to which the requesting API key has access. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	getOrganizationExecute(r GetOrganizationApiRequest) (*Organization, *http.Response, error)

	/*
		GetOrganizationInvitation Return One Organization Invitation

		Returns the details of one pending invitation to the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return GetOrganizationInvitationApiRequest
	*/
	GetOrganizationInvitation(ctx context.Context, orgId string, invitationId string) GetOrganizationInvitationApiRequest
	/*
		GetOrganizationInvitation Return One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationInvitationApiParams - Parameters for the request
		@return GetOrganizationInvitationApiRequest
	*/
	GetOrganizationInvitationWithParams(ctx context.Context, args *GetOrganizationInvitationApiParams) GetOrganizationInvitationApiRequest

	// Interface only available internally
	getOrganizationInvitationExecute(r GetOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		GetOrganizationSettings Return Settings for One Organization

		[experimental] Returns details about the specified organization's settings. To use this resource, the requesting API Key must have the Organization Owner role. This resource does not require the API Key to have an API access list.

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

	// Interface only available internally
	getOrganizationSettingsExecute(r GetOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error)

	/*
		ListOrganizationInvitations Return All Organization Invitations

		Returns all pending invitations to the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrganizationInvitationsApiRequest
	*/
	ListOrganizationInvitations(ctx context.Context, orgId string) ListOrganizationInvitationsApiRequest
	/*
		ListOrganizationInvitations Return All Organization Invitations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationInvitationsApiParams - Parameters for the request
		@return ListOrganizationInvitationsApiRequest
	*/
	ListOrganizationInvitationsWithParams(ctx context.Context, args *ListOrganizationInvitationsApiParams) ListOrganizationInvitationsApiRequest

	// Interface only available internally
	listOrganizationInvitationsExecute(r ListOrganizationInvitationsApiRequest) ([]OrganizationInvitation, *http.Response, error)

	/*
			ListOrganizationProjects Return One or More Projects in One Organization

			Returns multiple projects in the specified organization. Each organization can have multiple projects. Use projects to:

		- Isolate different environments, such as development, test, or production environments, from each other.
		- Associate different MongoDB Cloud users or teams with different environments, or give different permission to MongoDB Cloud users in different environments.
		- Maintain separate cluster security configurations.
		- Create different alert settings.

		To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrganizationProjectsApiRequest
	*/
	ListOrganizationProjects(ctx context.Context, orgId string) ListOrganizationProjectsApiRequest
	/*
		ListOrganizationProjects Return One or More Projects in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationProjectsApiParams - Parameters for the request
		@return ListOrganizationProjectsApiRequest
	*/
	ListOrganizationProjectsWithParams(ctx context.Context, args *ListOrganizationProjectsApiParams) ListOrganizationProjectsApiRequest

	// Interface only available internally
	listOrganizationProjectsExecute(r ListOrganizationProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error)

	/*
		ListOrganizationUsers Return All MongoDB Cloud Users in One Organization

		Returns details about the MongoDB Cloud users associated with the specified organization. Each MongoDB Cloud user returned must belong to the specified organization or to a project within the specified organization. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrganizationUsersApiRequest
	*/
	ListOrganizationUsers(ctx context.Context, orgId string) ListOrganizationUsersApiRequest
	/*
		ListOrganizationUsers Return All MongoDB Cloud Users in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationUsersApiParams - Parameters for the request
		@return ListOrganizationUsersApiRequest
	*/
	ListOrganizationUsersWithParams(ctx context.Context, args *ListOrganizationUsersApiParams) ListOrganizationUsersApiRequest

	// Interface only available internally
	listOrganizationUsersExecute(r ListOrganizationUsersApiRequest) (*PaginatedAppUser, *http.Response, error)

	/*
		ListOrganizations Return All Organizations

		Returns all organizations to which you belong. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	listOrganizationsExecute(r ListOrganizationsApiRequest) (*PaginatedOrganization, *http.Response, error)

	/*
		RemoveOrganizationUser Remove One MongoDB Cloud User from One Organization

		[experimental] Removes one MongoDB Cloud user from the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param userId Unique 24-hexadecimal digit string that identifies the user to be deleted.
		@return RemoveOrganizationUserApiRequest
	*/
	RemoveOrganizationUser(ctx context.Context, orgId string, userId string) RemoveOrganizationUserApiRequest
	/*
		RemoveOrganizationUser Remove One MongoDB Cloud User from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveOrganizationUserApiParams - Parameters for the request
		@return RemoveOrganizationUserApiRequest
	*/
	RemoveOrganizationUserWithParams(ctx context.Context, args *RemoveOrganizationUserApiParams) RemoveOrganizationUserApiRequest

	// Interface only available internally
	removeOrganizationUserExecute(r RemoveOrganizationUserApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		RenameOrganization Rename One Organization

		[experimental] Renames one organization. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return RenameOrganizationApiRequest
	*/
	RenameOrganization(ctx context.Context, orgId string, organization *Organization) RenameOrganizationApiRequest
	/*
		RenameOrganization Rename One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RenameOrganizationApiParams - Parameters for the request
		@return RenameOrganizationApiRequest
	*/
	RenameOrganizationWithParams(ctx context.Context, args *RenameOrganizationApiParams) RenameOrganizationApiRequest

	// Interface only available internally
	renameOrganizationExecute(r RenameOrganizationApiRequest) (*Organization, *http.Response, error)

	/*
		UpdateOrganizationInvitation Update One Organization Invitation

		Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the username of the invited user. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return UpdateOrganizationInvitationApiRequest
	*/
	UpdateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) UpdateOrganizationInvitationApiRequest
	/*
		UpdateOrganizationInvitation Update One Organization Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationInvitationApiParams - Parameters for the request
		@return UpdateOrganizationInvitationApiRequest
	*/
	UpdateOrganizationInvitationWithParams(ctx context.Context, args *UpdateOrganizationInvitationApiParams) UpdateOrganizationInvitationApiRequest

	// Interface only available internally
	updateOrganizationInvitationExecute(r UpdateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID

		Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the unique identification string for that invitation. Use the Return All Organization Invitations endpoint to retrieve IDs for all pending organization invitations. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return UpdateOrganizationInvitationByIdApiRequest
	*/
	UpdateOrganizationInvitationById(ctx context.Context, orgId string, invitationId string, organizationInvitationUpdateRequest *OrganizationInvitationUpdateRequest) UpdateOrganizationInvitationByIdApiRequest
	/*
		UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrganizationInvitationByIdApiParams - Parameters for the request
		@return UpdateOrganizationInvitationByIdApiRequest
	*/
	UpdateOrganizationInvitationByIdWithParams(ctx context.Context, args *UpdateOrganizationInvitationByIdApiParams) UpdateOrganizationInvitationByIdApiRequest

	// Interface only available internally
	updateOrganizationInvitationByIdExecute(r UpdateOrganizationInvitationByIdApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		UpdateOrganizationSettings Update Settings for One Organization

		[experimental] Updates the organization's settings. To use this resource, the requesting API Key must have the Organization Owner role. This resource does not require the API Key to have an API access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
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

	// Interface only available internally
	updateOrganizationSettingsExecute(r UpdateOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error)
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
	return r.ApiService.createOrganizationExecute(r)
}

/*
CreateOrganization Create One Organization

Creates one organization in MongoDB Cloud and links it to the requesting API Key's organization. To use this resource, the requesting API Key must have the Organization Owner role. The requesting API Key's organization must be a paying organization. To learn more, see [Configure a Paying Organization](https://www.mongodb.com/docs/atlas/billing/#configure-a-paying-organization) in the MongoDB Atlas documentation. This resource doesn't require the API Key to have an API Access List.

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

// Execute executes the request
//
//	@return CreateOrganizationResponse
func (a *OrganizationsApiService) createOrganizationExecute(r CreateOrganizationApiRequest) (*CreateOrganizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.createOrganizationInvitationExecute(r)
}

/*
CreateOrganizationInvitation Invite One MongoDB Cloud User to Join One Atlas Organization

Invites one MongoDB Cloud user to join the specified organization. The user must accept the invitation to access information within the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateOrganizationInvitationApiRequest
*/
func (a *OrganizationsApiService) CreateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) CreateOrganizationInvitationApiRequest {
	return CreateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		organizationInvitationRequest: organizationInvitationRequest,
	}
}

// Execute executes the request
//
//	@return OrganizationInvitation
func (a *OrganizationsApiService) createOrganizationInvitationExecute(r CreateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.CreateOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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

func (r DeleteOrganizationApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteOrganizationExecute(r)
}

/*
DeleteOrganization Remove One Organization

Removes one specified organization. MongoDB Cloud imposes the following limits on this resource:

  - Organizations with active projects cannot be removed.

  - All projects in the organization must be removed before you can remove the organization.
    To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return map[string]interface{}
func (a *OrganizationsApiService) deleteOrganizationExecute(r DeleteOrganizationApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.DeleteOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
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

func (r DeleteOrganizationInvitationApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteOrganizationInvitationExecute(r)
}

/*
DeleteOrganizationInvitation Cancel One Organization Invitation

Cancels one pending invitation sent to the specified MongoDB Cloud user to join an organization. You can't cancel an invitation that the user accepted. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return DeleteOrganizationInvitationApiRequest
*/
func (a *OrganizationsApiService) DeleteOrganizationInvitation(ctx context.Context, orgId string, invitationId string) DeleteOrganizationInvitationApiRequest {
	return DeleteOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		invitationId: invitationId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *OrganizationsApiService) deleteOrganizationInvitationExecute(r DeleteOrganizationInvitationApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.DeleteOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
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

func (r GetOrganizationApiRequest) Execute() (*Organization, *http.Response, error) {
	return r.ApiService.getOrganizationExecute(r)
}

/*
GetOrganization Return One Organization

Returns one organization to which the requesting API key has access. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return Organization
func (a *OrganizationsApiService) getOrganizationExecute(r GetOrganizationApiRequest) (*Organization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Organization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.getOrganizationInvitationExecute(r)
}

/*
GetOrganizationInvitation Return One Organization Invitation

Returns the details of one pending invitation to the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return GetOrganizationInvitationApiRequest
*/
func (a *OrganizationsApiService) GetOrganizationInvitation(ctx context.Context, orgId string, invitationId string) GetOrganizationInvitationApiRequest {
	return GetOrganizationInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		invitationId: invitationId,
	}
}

// Execute executes the request
//
//	@return OrganizationInvitation
func (a *OrganizationsApiService) getOrganizationInvitationExecute(r GetOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.invitationId) < 24 {
		return localVarReturnValue, nil, reportError("invitationId must have at least 24 elements")
	}
	if strlen(r.invitationId) > 24 {
		return localVarReturnValue, nil, reportError("invitationId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.getOrganizationSettingsExecute(r)
}

/*
GetOrganizationSettings Return Settings for One Organization

[experimental] Returns details about the specified organization's settings. To use this resource, the requesting API Key must have the Organization Owner role. This resource does not require the API Key to have an API access list.

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

// Execute executes the request
//
//	@return OrganizationSettings
func (a *OrganizationsApiService) getOrganizationSettingsExecute(r GetOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrganizationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.listOrganizationInvitationsExecute(r)
}

/*
ListOrganizationInvitations Return All Organization Invitations

Returns all pending invitations to the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationInvitationsApiRequest
*/
func (a *OrganizationsApiService) ListOrganizationInvitations(ctx context.Context, orgId string) ListOrganizationInvitationsApiRequest {
	return ListOrganizationInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return []OrganizationInvitation
func (a *OrganizationsApiService) listOrganizationInvitationsExecute(r ListOrganizationInvitationsApiRequest) ([]OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizationInvitations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.listOrganizationProjectsExecute(r)
}

/*
ListOrganizationProjects Return One or More Projects in One Organization

Returns multiple projects in the specified organization. Each organization can have multiple projects. Use projects to:

- Isolate different environments, such as development, test, or production environments, from each other.
- Associate different MongoDB Cloud users or teams with different environments, or give different permission to MongoDB Cloud users in different environments.
- Maintain separate cluster security configurations.
- Create different alert settings.

To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return PaginatedAtlasGroup
func (a *OrganizationsApiService) listOrganizationProjectsExecute(r ListOrganizationProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizationProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListOrganizationUsersApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	orgId        string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListOrganizationUsersApiParams struct {
	OrgId        string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *OrganizationsApiService) ListOrganizationUsersWithParams(ctx context.Context, args *ListOrganizationUsersApiParams) ListOrganizationUsersApiRequest {
	return ListOrganizationUsersApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
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

func (r ListOrganizationUsersApiRequest) Execute() (*PaginatedAppUser, *http.Response, error) {
	return r.ApiService.listOrganizationUsersExecute(r)
}

/*
ListOrganizationUsers Return All MongoDB Cloud Users in One Organization

Returns details about the MongoDB Cloud users associated with the specified organization. Each MongoDB Cloud user returned must belong to the specified organization or to a project within the specified organization. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationUsersApiRequest
*/
func (a *OrganizationsApiService) ListOrganizationUsers(ctx context.Context, orgId string) ListOrganizationUsersApiRequest {
	return ListOrganizationUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return PaginatedAppUser
func (a *OrganizationsApiService) listOrganizationUsersExecute(r ListOrganizationUsersApiRequest) (*PaginatedAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.ListOrganizationUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.listOrganizationsExecute(r)
}

/*
ListOrganizations Return All Organizations

Returns all organizations to which you belong. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListOrganizationsApiRequest
*/
func (a *OrganizationsApiService) ListOrganizations(ctx context.Context) ListOrganizationsApiRequest {
	return ListOrganizationsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedOrganization
func (a *OrganizationsApiService) listOrganizationsExecute(r ListOrganizationsApiRequest) (*PaginatedOrganization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveOrganizationUserApiRequest struct {
	ctx        context.Context
	ApiService OrganizationsApi
	orgId      string
	userId     string
}

type RemoveOrganizationUserApiParams struct {
	OrgId  string
	UserId string
}

func (a *OrganizationsApiService) RemoveOrganizationUserWithParams(ctx context.Context, args *RemoveOrganizationUserApiParams) RemoveOrganizationUserApiRequest {
	return RemoveOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		userId:     args.UserId,
	}
}

func (r RemoveOrganizationUserApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.removeOrganizationUserExecute(r)
}

/*
RemoveOrganizationUser Remove One MongoDB Cloud User from One Organization

[experimental] Removes one MongoDB Cloud user from the specified organization. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param userId Unique 24-hexadecimal digit string that identifies the user to be deleted.
	@return RemoveOrganizationUserApiRequest
*/
func (a *OrganizationsApiService) RemoveOrganizationUser(ctx context.Context, orgId string, userId string) RemoveOrganizationUserApiRequest {
	return RemoveOrganizationUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *OrganizationsApiService) removeOrganizationUserExecute(r RemoveOrganizationUserApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.RemoveOrganizationUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.userId) < 24 {
		return localVarReturnValue, nil, reportError("userId must have at least 24 elements")
	}
	if strlen(r.userId) > 24 {
		return localVarReturnValue, nil, reportError("userId must have less than 24 elements")
	}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RenameOrganizationApiRequest struct {
	ctx          context.Context
	ApiService   OrganizationsApi
	orgId        string
	organization *Organization
}

type RenameOrganizationApiParams struct {
	OrgId        string
	Organization *Organization
}

func (a *OrganizationsApiService) RenameOrganizationWithParams(ctx context.Context, args *RenameOrganizationApiParams) RenameOrganizationApiRequest {
	return RenameOrganizationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		organization: args.Organization,
	}
}

func (r RenameOrganizationApiRequest) Execute() (*Organization, *http.Response, error) {
	return r.ApiService.renameOrganizationExecute(r)
}

/*
RenameOrganization Rename One Organization

[experimental] Renames one organization. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return RenameOrganizationApiRequest
*/
func (a *OrganizationsApiService) RenameOrganization(ctx context.Context, orgId string, organization *Organization) RenameOrganizationApiRequest {
	return RenameOrganizationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		organization: organization,
	}
}

// Execute executes the request
//
//	@return Organization
func (a *OrganizationsApiService) renameOrganizationExecute(r RenameOrganizationApiRequest) (*Organization, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Organization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.RenameOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if r.organization == nil {
		return localVarReturnValue, nil, reportError("organization is required and must be specified")
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
	localVarPostBody = r.organization
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.updateOrganizationInvitationExecute(r)
}

/*
UpdateOrganizationInvitation Update One Organization Invitation

Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the username of the invited user. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateOrganizationInvitationApiRequest
*/
func (a *OrganizationsApiService) UpdateOrganizationInvitation(ctx context.Context, orgId string, organizationInvitationRequest *OrganizationInvitationRequest) UpdateOrganizationInvitationApiRequest {
	return UpdateOrganizationInvitationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		organizationInvitationRequest: organizationInvitationRequest,
	}
}

// Execute executes the request
//
//	@return OrganizationInvitation
func (a *OrganizationsApiService) updateOrganizationInvitationExecute(r UpdateOrganizationInvitationApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.updateOrganizationInvitationByIdExecute(r)
}

/*
UpdateOrganizationInvitationById Update One Organization Invitation by Invitation ID

Updates the details of one pending invitation to the specified organization. To specify which invitation, provide the unique identification string for that invitation. Use the Return All Organization Invitations endpoint to retrieve IDs for all pending organization invitations. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return UpdateOrganizationInvitationByIdApiRequest
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

// Execute executes the request
//
//	@return OrganizationInvitation
func (a *OrganizationsApiService) updateOrganizationInvitationByIdExecute(r UpdateOrganizationInvitationByIdApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationInvitationById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.invitationId) < 24 {
		return localVarReturnValue, nil, reportError("invitationId must have at least 24 elements")
	}
	if strlen(r.invitationId) > 24 {
		return localVarReturnValue, nil, reportError("invitationId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
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
	return r.ApiService.updateOrganizationSettingsExecute(r)
}

/*
UpdateOrganizationSettings Update Settings for One Organization

[experimental] Updates the organization's settings. To use this resource, the requesting API Key must have the Organization Owner role. This resource does not require the API Key to have an API access list.

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

// Execute executes the request
//
//	@return OrganizationSettings
func (a *OrganizationsApiService) updateOrganizationSettingsExecute(r UpdateOrganizationSettingsApiRequest) (*OrganizationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UpdateOrganizationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
