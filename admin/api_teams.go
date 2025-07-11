// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TeamsApi interface {

	/*
		AddAllTeamsToProject Add One Team to One Project

		Adds one team to the specified project. All members of the team share the same project access. MongoDB Cloud limits the number of users to a maximum of 100 teams per project and a maximum of 250 teams per organization. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param teamRole Team to add to the specified project.
		@return AddAllTeamsToProjectApiRequest
	*/
	AddAllTeamsToProject(ctx context.Context, groupId string, teamRole *[]TeamRole) AddAllTeamsToProjectApiRequest
	/*
		AddAllTeamsToProject Add One Team to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddAllTeamsToProjectApiParams - Parameters for the request
		@return AddAllTeamsToProjectApiRequest
	*/
	AddAllTeamsToProjectWithParams(ctx context.Context, args *AddAllTeamsToProjectApiParams) AddAllTeamsToProjectApiRequest

	// Method available only for mocking purposes
	AddAllTeamsToProjectExecute(r AddAllTeamsToProjectApiRequest) (*PaginatedTeamRole, *http.Response, error)

	/*
			AddTeamUser Assign MongoDB Cloud Users in One Organization to One Team

			Adds one or more MongoDB Cloud users from the specified organization to the specified team. Teams enable you to grant project access roles to MongoDB Cloud users. You can assign up to 250 MongoDB Cloud users from one organization to one team. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: This endpoint is deprecated. Use [Add One MongoDB Cloud User to One Team](#tag/MongoDB-Cloud-Users/operation/addUserToTeam) to add an active or pending user to a team.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param teamId Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.
			@param addUserToTeam One or more MongoDB Cloud users that you want to add to the specified team.
			@return AddTeamUserApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for TeamsApi
	*/
	AddTeamUser(ctx context.Context, orgId string, teamId string, addUserToTeam *[]AddUserToTeam) AddTeamUserApiRequest
	/*
		AddTeamUser Assign MongoDB Cloud Users in One Organization to One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddTeamUserApiParams - Parameters for the request
		@return AddTeamUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for TeamsApi
	*/
	AddTeamUserWithParams(ctx context.Context, args *AddTeamUserApiParams) AddTeamUserApiRequest

	// Method available only for mocking purposes
	AddTeamUserExecute(r AddTeamUserApiRequest) (*PaginatedApiAppUser, *http.Response, error)

	/*
		CreateTeam Create One Team in One Organization

		Creates one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud limits the number of teams to a maximum of 250 teams per organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param team Team that you want to create in the specified organization.
		@return CreateTeamApiRequest
	*/
	CreateTeam(ctx context.Context, orgId string, team *Team) CreateTeamApiRequest
	/*
		CreateTeam Create One Team in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateTeamApiParams - Parameters for the request
		@return CreateTeamApiRequest
	*/
	CreateTeamWithParams(ctx context.Context, args *CreateTeamApiParams) CreateTeamApiRequest

	// Method available only for mocking purposes
	CreateTeamExecute(r CreateTeamApiRequest) (*Team, *http.Response, error)

	/*
		DeleteTeam Remove One Team from One Organization

		Removes one team specified using its unique 24-hexadecimal digit identifier from the organization specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to delete.
		@return DeleteTeamApiRequest
	*/
	DeleteTeam(ctx context.Context, orgId string, teamId string) DeleteTeamApiRequest
	/*
		DeleteTeam Remove One Team from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteTeamApiParams - Parameters for the request
		@return DeleteTeamApiRequest
	*/
	DeleteTeamWithParams(ctx context.Context, args *DeleteTeamApiParams) DeleteTeamApiRequest

	// Method available only for mocking purposes
	DeleteTeamExecute(r DeleteTeamApiRequest) (*http.Response, error)

	/*
		GetProjectTeam Return One Team in One Project

		Returns one team to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param teamId Unique 24-hexadecimal digit string that identifies the team for which you want to get.
		@return GetProjectTeamApiRequest
	*/
	GetProjectTeam(ctx context.Context, groupId string, teamId string) GetProjectTeamApiRequest
	/*
		GetProjectTeam Return One Team in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectTeamApiParams - Parameters for the request
		@return GetProjectTeamApiRequest
	*/
	GetProjectTeamWithParams(ctx context.Context, args *GetProjectTeamApiParams) GetProjectTeamApiRequest

	// Method available only for mocking purposes
	GetProjectTeamExecute(r GetProjectTeamApiRequest) (*TeamRole, *http.Response, error)

	/*
		GetTeamById Return One Team by ID

		Returns one team that you identified using its unique 24-hexadecimal digit ID. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team whose information you want to return.
		@return GetTeamByIdApiRequest
	*/
	GetTeamById(ctx context.Context, orgId string, teamId string) GetTeamByIdApiRequest
	/*
		GetTeamById Return One Team by ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetTeamByIdApiParams - Parameters for the request
		@return GetTeamByIdApiRequest
	*/
	GetTeamByIdWithParams(ctx context.Context, args *GetTeamByIdApiParams) GetTeamByIdApiRequest

	// Method available only for mocking purposes
	GetTeamByIdExecute(r GetTeamByIdApiRequest) (*TeamResponse, *http.Response, error)

	/*
		GetTeamByName Return One Team by Name

		Returns one team that you identified using its human-readable name. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamName Name of the team whose information you want to return.
		@return GetTeamByNameApiRequest
	*/
	GetTeamByName(ctx context.Context, orgId string, teamName string) GetTeamByNameApiRequest
	/*
		GetTeamByName Return One Team by Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetTeamByNameApiParams - Parameters for the request
		@return GetTeamByNameApiRequest
	*/
	GetTeamByNameWithParams(ctx context.Context, args *GetTeamByNameApiParams) GetTeamByNameApiRequest

	// Method available only for mocking purposes
	GetTeamByNameExecute(r GetTeamByNameApiRequest) (*TeamResponse, *http.Response, error)

	/*
		ListOrganizationTeams Return All Teams in One Organization

		Returns all teams that belong to the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud only returns teams for which you have access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrganizationTeamsApiRequest
	*/
	ListOrganizationTeams(ctx context.Context, orgId string) ListOrganizationTeamsApiRequest
	/*
		ListOrganizationTeams Return All Teams in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationTeamsApiParams - Parameters for the request
		@return ListOrganizationTeamsApiRequest
	*/
	ListOrganizationTeamsWithParams(ctx context.Context, args *ListOrganizationTeamsApiParams) ListOrganizationTeamsApiRequest

	// Method available only for mocking purposes
	ListOrganizationTeamsExecute(r ListOrganizationTeamsApiRequest) (*PaginatedTeam, *http.Response, error)

	/*
		ListProjectTeams Return All Teams in One Project

		Returns all teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectTeamsApiRequest
	*/
	ListProjectTeams(ctx context.Context, groupId string) ListProjectTeamsApiRequest
	/*
		ListProjectTeams Return All Teams in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectTeamsApiParams - Parameters for the request
		@return ListProjectTeamsApiRequest
	*/
	ListProjectTeamsWithParams(ctx context.Context, args *ListProjectTeamsApiParams) ListProjectTeamsApiRequest

	// Method available only for mocking purposes
	ListProjectTeamsExecute(r ListProjectTeamsApiRequest) (*PaginatedTeamRole, *http.Response, error)

	/*
		RemoveProjectTeam Remove One Team from One Project

		Removes one team specified using its unique 24-hexadecimal digit identifier from the project specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.
		@return RemoveProjectTeamApiRequest
	*/
	RemoveProjectTeam(ctx context.Context, groupId string, teamId string) RemoveProjectTeamApiRequest
	/*
		RemoveProjectTeam Remove One Team from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveProjectTeamApiParams - Parameters for the request
		@return RemoveProjectTeamApiRequest
	*/
	RemoveProjectTeamWithParams(ctx context.Context, args *RemoveProjectTeamApiParams) RemoveProjectTeamApiRequest

	// Method available only for mocking purposes
	RemoveProjectTeamExecute(r RemoveProjectTeamApiRequest) (*http.Response, error)

	/*
			RemoveTeamUser Remove One MongoDB Cloud User from One Team

			Removes one MongoDB Cloud user from the specified team. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		**Note**: This endpoint is deprecated. Use [Remove One MongoDB Cloud User from One Team](#tag/MongoDB-Cloud-Users/operation/removeUserFromTeam) to remove an active or pending user from a team.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param teamId Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.
			@param userId Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.
			@return RemoveTeamUserApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for TeamsApi
	*/
	RemoveTeamUser(ctx context.Context, orgId string, teamId string, userId string) RemoveTeamUserApiRequest
	/*
		RemoveTeamUser Remove One MongoDB Cloud User from One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveTeamUserApiParams - Parameters for the request
		@return RemoveTeamUserApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for TeamsApi
	*/
	RemoveTeamUserWithParams(ctx context.Context, args *RemoveTeamUserApiParams) RemoveTeamUserApiRequest

	// Method available only for mocking purposes
	RemoveTeamUserExecute(r RemoveTeamUserApiRequest) (*http.Response, error)

	/*
		RenameTeam Rename One Team

		Renames one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to rename.
		@param teamUpdate Details to update on the specified team.
		@return RenameTeamApiRequest
	*/
	RenameTeam(ctx context.Context, orgId string, teamId string, teamUpdate *TeamUpdate) RenameTeamApiRequest
	/*
		RenameTeam Rename One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RenameTeamApiParams - Parameters for the request
		@return RenameTeamApiRequest
	*/
	RenameTeamWithParams(ctx context.Context, args *RenameTeamApiParams) RenameTeamApiRequest

	// Method available only for mocking purposes
	RenameTeamExecute(r RenameTeamApiRequest) (*TeamResponse, *http.Response, error)

	/*
		UpdateTeamRoles Update Team Roles in One Project

		Updates the project roles assigned to the specified team. You can grant team roles for specific projects and grant project access roles to users in the team. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param teamId Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.
		@param teamRole The project roles assigned to the specified team.
		@return UpdateTeamRolesApiRequest
	*/
	UpdateTeamRoles(ctx context.Context, groupId string, teamId string, teamRole *TeamRole) UpdateTeamRolesApiRequest
	/*
		UpdateTeamRoles Update Team Roles in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateTeamRolesApiParams - Parameters for the request
		@return UpdateTeamRolesApiRequest
	*/
	UpdateTeamRolesWithParams(ctx context.Context, args *UpdateTeamRolesApiParams) UpdateTeamRolesApiRequest

	// Method available only for mocking purposes
	UpdateTeamRolesExecute(r UpdateTeamRolesApiRequest) (*PaginatedTeamRole, *http.Response, error)
}

// TeamsApiService TeamsApi service
type TeamsApiService service

type AddAllTeamsToProjectApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	groupId    string
	teamRole   *[]TeamRole
}

type AddAllTeamsToProjectApiParams struct {
	GroupId  string
	TeamRole *[]TeamRole
}

func (a *TeamsApiService) AddAllTeamsToProjectWithParams(ctx context.Context, args *AddAllTeamsToProjectApiParams) AddAllTeamsToProjectApiRequest {
	return AddAllTeamsToProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		teamRole:   args.TeamRole,
	}
}

func (r AddAllTeamsToProjectApiRequest) Execute() (*PaginatedTeamRole, *http.Response, error) {
	return r.ApiService.AddAllTeamsToProjectExecute(r)
}

/*
AddAllTeamsToProject Add One Team to One Project

Adds one team to the specified project. All members of the team share the same project access. MongoDB Cloud limits the number of users to a maximum of 100 teams per project and a maximum of 250 teams per organization. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AddAllTeamsToProjectApiRequest
*/
func (a *TeamsApiService) AddAllTeamsToProject(ctx context.Context, groupId string, teamRole *[]TeamRole) AddAllTeamsToProjectApiRequest {
	return AddAllTeamsToProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		teamRole:   teamRole,
	}
}

// AddAllTeamsToProjectExecute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) AddAllTeamsToProjectExecute(r AddAllTeamsToProjectApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.AddAllTeamsToProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamRole == nil {
		return localVarReturnValue, nil, reportError("teamRole is required and must be specified")
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
	localVarPostBody = r.teamRole
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

type AddTeamUserApiRequest struct {
	ctx           context.Context
	ApiService    TeamsApi
	orgId         string
	teamId        string
	addUserToTeam *[]AddUserToTeam
}

type AddTeamUserApiParams struct {
	OrgId         string
	TeamId        string
	AddUserToTeam *[]AddUserToTeam
}

func (a *TeamsApiService) AddTeamUserWithParams(ctx context.Context, args *AddTeamUserApiParams) AddTeamUserApiRequest {
	return AddTeamUserApiRequest{
		ApiService:    a,
		ctx:           ctx,
		orgId:         args.OrgId,
		teamId:        args.TeamId,
		addUserToTeam: args.AddUserToTeam,
	}
}

func (r AddTeamUserApiRequest) Execute() (*PaginatedApiAppUser, *http.Response, error) {
	return r.ApiService.AddTeamUserExecute(r)
}

/*
AddTeamUser Assign MongoDB Cloud Users in One Organization to One Team

Adds one or more MongoDB Cloud users from the specified organization to the specified team. Teams enable you to grant project access roles to MongoDB Cloud users. You can assign up to 250 MongoDB Cloud users from one organization to one team. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: This endpoint is deprecated. Use [Add One MongoDB Cloud User to One Team](#tag/MongoDB-Cloud-Users/operation/addUserToTeam) to add an active or pending user to a team.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.
	@return AddTeamUserApiRequest

Deprecated
*/
func (a *TeamsApiService) AddTeamUser(ctx context.Context, orgId string, teamId string, addUserToTeam *[]AddUserToTeam) AddTeamUserApiRequest {
	return AddTeamUserApiRequest{
		ApiService:    a,
		ctx:           ctx,
		orgId:         orgId,
		teamId:        teamId,
		addUserToTeam: addUserToTeam,
	}
}

// AddTeamUserExecute executes the request
//
//	@return PaginatedApiAppUser
//
// Deprecated
func (a *TeamsApiService) AddTeamUserExecute(r AddTeamUserApiRequest) (*PaginatedApiAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.AddTeamUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamId == "" {
		return localVarReturnValue, nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addUserToTeam == nil {
		return localVarReturnValue, nil, reportError("addUserToTeam is required and must be specified")
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
	localVarPostBody = r.addUserToTeam
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

type CreateTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	team       *Team
}

type CreateTeamApiParams struct {
	OrgId string
	Team  *Team
}

func (a *TeamsApiService) CreateTeamWithParams(ctx context.Context, args *CreateTeamApiParams) CreateTeamApiRequest {
	return CreateTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		team:       args.Team,
	}
}

func (r CreateTeamApiRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.CreateTeamExecute(r)
}

/*
CreateTeam Create One Team in One Organization

Creates one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud limits the number of teams to a maximum of 250 teams per organization. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateTeamApiRequest
*/
func (a *TeamsApiService) CreateTeam(ctx context.Context, orgId string, team *Team) CreateTeamApiRequest {
	return CreateTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		team:       team,
	}
}

// CreateTeamExecute executes the request
//
//	@return Team
func (a *TeamsApiService) CreateTeamExecute(r CreateTeamApiRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.CreateTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.team == nil {
		return localVarReturnValue, nil, reportError("team is required and must be specified")
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
	localVarPostBody = r.team
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

type DeleteTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamId     string
}

type DeleteTeamApiParams struct {
	OrgId  string
	TeamId string
}

func (a *TeamsApiService) DeleteTeamWithParams(ctx context.Context, args *DeleteTeamApiParams) DeleteTeamApiRequest {
	return DeleteTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamId:     args.TeamId,
	}
}

func (r DeleteTeamApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTeamExecute(r)
}

/*
DeleteTeam Remove One Team from One Organization

Removes one team specified using its unique 24-hexadecimal digit identifier from the organization specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to delete.
	@return DeleteTeamApiRequest
*/
func (a *TeamsApiService) DeleteTeam(ctx context.Context, orgId string, teamId string) DeleteTeamApiRequest {
	return DeleteTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
	}
}

// DeleteTeamExecute executes the request
func (a *TeamsApiService) DeleteTeamExecute(r DeleteTeamApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.DeleteTeam")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamId == "" {
		return nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

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

type GetProjectTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	groupId    string
	teamId     string
}

type GetProjectTeamApiParams struct {
	GroupId string
	TeamId  string
}

func (a *TeamsApiService) GetProjectTeamWithParams(ctx context.Context, args *GetProjectTeamApiParams) GetProjectTeamApiRequest {
	return GetProjectTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		teamId:     args.TeamId,
	}
}

func (r GetProjectTeamApiRequest) Execute() (*TeamRole, *http.Response, error) {
	return r.ApiService.GetProjectTeamExecute(r)
}

/*
GetProjectTeam Return One Team in One Project

Returns one team to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param teamId Unique 24-hexadecimal digit string that identifies the team for which you want to get.
	@return GetProjectTeamApiRequest
*/
func (a *TeamsApiService) GetProjectTeam(ctx context.Context, groupId string, teamId string) GetProjectTeamApiRequest {
	return GetProjectTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		teamId:     teamId,
	}
}

// GetProjectTeamExecute executes the request
//
//	@return TeamRole
func (a *TeamsApiService) GetProjectTeamExecute(r GetProjectTeamApiRequest) (*TeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetProjectTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams/{teamId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.teamId == "" {
		return localVarReturnValue, nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

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

type GetTeamByIdApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamId     string
}

type GetTeamByIdApiParams struct {
	OrgId  string
	TeamId string
}

func (a *TeamsApiService) GetTeamByIdWithParams(ctx context.Context, args *GetTeamByIdApiParams) GetTeamByIdApiRequest {
	return GetTeamByIdApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamId:     args.TeamId,
	}
}

func (r GetTeamByIdApiRequest) Execute() (*TeamResponse, *http.Response, error) {
	return r.ApiService.GetTeamByIdExecute(r)
}

/*
GetTeamById Return One Team by ID

Returns one team that you identified using its unique 24-hexadecimal digit ID. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team whose information you want to return.
	@return GetTeamByIdApiRequest
*/
func (a *TeamsApiService) GetTeamById(ctx context.Context, orgId string, teamId string) GetTeamByIdApiRequest {
	return GetTeamByIdApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
	}
}

// GetTeamByIdExecute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) GetTeamByIdExecute(r GetTeamByIdApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetTeamById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamId == "" {
		return localVarReturnValue, nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

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

type GetTeamByNameApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamName   string
}

type GetTeamByNameApiParams struct {
	OrgId    string
	TeamName string
}

func (a *TeamsApiService) GetTeamByNameWithParams(ctx context.Context, args *GetTeamByNameApiParams) GetTeamByNameApiRequest {
	return GetTeamByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamName:   args.TeamName,
	}
}

func (r GetTeamByNameApiRequest) Execute() (*TeamResponse, *http.Response, error) {
	return r.ApiService.GetTeamByNameExecute(r)
}

/*
GetTeamByName Return One Team by Name

Returns one team that you identified using its human-readable name. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamName Name of the team whose information you want to return.
	@return GetTeamByNameApiRequest
*/
func (a *TeamsApiService) GetTeamByName(ctx context.Context, orgId string, teamName string) GetTeamByNameApiRequest {
	return GetTeamByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamName:   teamName,
	}
}

// GetTeamByNameExecute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) GetTeamByNameExecute(r GetTeamByNameApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetTeamByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/byName/{teamName}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamName == "" {
		return localVarReturnValue, nil, reportError("teamName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamName"+"}", url.PathEscape(r.teamName), -1)

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

type ListOrganizationTeamsApiRequest struct {
	ctx          context.Context
	ApiService   TeamsApi
	orgId        string
	itemsPerPage *int
	includeCount *bool
	pageNum      *int
}

type ListOrganizationTeamsApiParams struct {
	OrgId        string
	ItemsPerPage *int
	IncludeCount *bool
	PageNum      *int
}

func (a *TeamsApiService) ListOrganizationTeamsWithParams(ctx context.Context, args *ListOrganizationTeamsApiParams) ListOrganizationTeamsApiRequest {
	return ListOrganizationTeamsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		includeCount: args.IncludeCount,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrganizationTeamsApiRequest) ItemsPerPage(itemsPerPage int) ListOrganizationTeamsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOrganizationTeamsApiRequest) IncludeCount(includeCount bool) ListOrganizationTeamsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrganizationTeamsApiRequest) PageNum(pageNum int) ListOrganizationTeamsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrganizationTeamsApiRequest) Execute() (*PaginatedTeam, *http.Response, error) {
	return r.ApiService.ListOrganizationTeamsExecute(r)
}

/*
ListOrganizationTeams Return All Teams in One Organization

Returns all teams that belong to the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud only returns teams for which you have access. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationTeamsApiRequest
*/
func (a *TeamsApiService) ListOrganizationTeams(ctx context.Context, orgId string) ListOrganizationTeamsApiRequest {
	return ListOrganizationTeamsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrganizationTeamsExecute executes the request
//
//	@return PaginatedTeam
func (a *TeamsApiService) ListOrganizationTeamsExecute(r ListOrganizationTeamsApiRequest) (*PaginatedTeam, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListOrganizationTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams"
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
	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
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

type ListProjectTeamsApiRequest struct {
	ctx          context.Context
	ApiService   TeamsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectTeamsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *TeamsApiService) ListProjectTeamsWithParams(ctx context.Context, args *ListProjectTeamsApiParams) ListProjectTeamsApiRequest {
	return ListProjectTeamsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectTeamsApiRequest) IncludeCount(includeCount bool) ListProjectTeamsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectTeamsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectTeamsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectTeamsApiRequest) PageNum(pageNum int) ListProjectTeamsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectTeamsApiRequest) Execute() (*PaginatedTeamRole, *http.Response, error) {
	return r.ApiService.ListProjectTeamsExecute(r)
}

/*
ListProjectTeams Return All Teams in One Project

Returns all teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectTeamsApiRequest
*/
func (a *TeamsApiService) ListProjectTeams(ctx context.Context, groupId string) ListProjectTeamsApiRequest {
	return ListProjectTeamsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectTeamsExecute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) ListProjectTeamsExecute(r ListProjectTeamsApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListProjectTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams"
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

type RemoveProjectTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	groupId    string
	teamId     string
}

type RemoveProjectTeamApiParams struct {
	GroupId string
	TeamId  string
}

func (a *TeamsApiService) RemoveProjectTeamWithParams(ctx context.Context, args *RemoveProjectTeamApiParams) RemoveProjectTeamApiRequest {
	return RemoveProjectTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		teamId:     args.TeamId,
	}
}

func (r RemoveProjectTeamApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveProjectTeamExecute(r)
}

/*
RemoveProjectTeam Remove One Team from One Project

Removes one team specified using its unique 24-hexadecimal digit identifier from the project specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.
	@return RemoveProjectTeamApiRequest
*/
func (a *TeamsApiService) RemoveProjectTeam(ctx context.Context, groupId string, teamId string) RemoveProjectTeamApiRequest {
	return RemoveProjectTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		teamId:     teamId,
	}
}

// RemoveProjectTeamExecute executes the request
func (a *TeamsApiService) RemoveProjectTeamExecute(r RemoveProjectTeamApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RemoveProjectTeam")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams/{teamId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.teamId == "" {
		return nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

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

type RemoveTeamUserApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamId     string
	userId     string
}

type RemoveTeamUserApiParams struct {
	OrgId  string
	TeamId string
	UserId string
}

func (a *TeamsApiService) RemoveTeamUserWithParams(ctx context.Context, args *RemoveTeamUserApiParams) RemoveTeamUserApiRequest {
	return RemoveTeamUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamId:     args.TeamId,
		userId:     args.UserId,
	}
}

func (r RemoveTeamUserApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveTeamUserExecute(r)
}

/*
RemoveTeamUser Remove One MongoDB Cloud User from One Team

Removes one MongoDB Cloud user from the specified team. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

**Note**: This endpoint is deprecated. Use [Remove One MongoDB Cloud User from One Team](#tag/MongoDB-Cloud-Users/operation/removeUserFromTeam) to remove an active or pending user from a team.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.
	@param userId Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.
	@return RemoveTeamUserApiRequest

Deprecated
*/
func (a *TeamsApiService) RemoveTeamUser(ctx context.Context, orgId string, teamId string, userId string) RemoveTeamUserApiRequest {
	return RemoveTeamUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
		userId:     userId,
	}
}

// RemoveTeamUserExecute executes the request
// Deprecated
func (a *TeamsApiService) RemoveTeamUserExecute(r RemoveTeamUserApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RemoveTeamUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamId == "" {
		return nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)
	if r.userId == "" {
		return nil, reportError("userId is empty and must be specified")
	}
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

type RenameTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamId     string
	teamUpdate *TeamUpdate
}

type RenameTeamApiParams struct {
	OrgId      string
	TeamId     string
	TeamUpdate *TeamUpdate
}

func (a *TeamsApiService) RenameTeamWithParams(ctx context.Context, args *RenameTeamApiParams) RenameTeamApiRequest {
	return RenameTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamId:     args.TeamId,
		teamUpdate: args.TeamUpdate,
	}
}

func (r RenameTeamApiRequest) Execute() (*TeamResponse, *http.Response, error) {
	return r.ApiService.RenameTeamExecute(r)
}

/*
RenameTeam Rename One Team

Renames one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to rename.
	@return RenameTeamApiRequest
*/
func (a *TeamsApiService) RenameTeam(ctx context.Context, orgId string, teamId string, teamUpdate *TeamUpdate) RenameTeamApiRequest {
	return RenameTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
		teamUpdate: teamUpdate,
	}
}

// RenameTeamExecute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) RenameTeamExecute(r RenameTeamApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RenameTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.teamId == "" {
		return localVarReturnValue, nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamUpdate == nil {
		return localVarReturnValue, nil, reportError("teamUpdate is required and must be specified")
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
	localVarPostBody = r.teamUpdate
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

type UpdateTeamRolesApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	groupId    string
	teamId     string
	teamRole   *TeamRole
}

type UpdateTeamRolesApiParams struct {
	GroupId  string
	TeamId   string
	TeamRole *TeamRole
}

func (a *TeamsApiService) UpdateTeamRolesWithParams(ctx context.Context, args *UpdateTeamRolesApiParams) UpdateTeamRolesApiRequest {
	return UpdateTeamRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		teamId:     args.TeamId,
		teamRole:   args.TeamRole,
	}
}

func (r UpdateTeamRolesApiRequest) Execute() (*PaginatedTeamRole, *http.Response, error) {
	return r.ApiService.UpdateTeamRolesExecute(r)
}

/*
UpdateTeamRoles Update Team Roles in One Project

Updates the project roles assigned to the specified team. You can grant team roles for specific projects and grant project access roles to users in the team. All members of the team share the same project access. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param teamId Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.
	@return UpdateTeamRolesApiRequest
*/
func (a *TeamsApiService) UpdateTeamRoles(ctx context.Context, groupId string, teamId string, teamRole *TeamRole) UpdateTeamRolesApiRequest {
	return UpdateTeamRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		teamId:     teamId,
		teamRole:   teamRole,
	}
}

// UpdateTeamRolesExecute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) UpdateTeamRolesExecute(r UpdateTeamRolesApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.UpdateTeamRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams/{teamId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.teamId == "" {
		return localVarReturnValue, nil, reportError("teamId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(r.teamId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamRole == nil {
		return localVarReturnValue, nil, reportError("teamRole is required and must be specified")
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
	localVarPostBody = r.teamRole
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
