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

type TeamsApi interface {

	/*
		AddAllTeamsToProject Add One or More Teams to One Project

		Adds one team to the specified project. All members of the team share the same project access. MongoDB Cloud limits the number of users to a maximum of 100 teams per project and a maximum of 250 teams per organization. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return AddAllTeamsToProjectApiRequest
	*/
	AddAllTeamsToProject(ctx context.Context, groupId string, teamRole *[]TeamRole) AddAllTeamsToProjectApiRequest
	/*
		AddAllTeamsToProject Add One or More Teams to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddAllTeamsToProjectApiParams - Parameters for the request
		@return AddAllTeamsToProjectApiRequest
	*/
	AddAllTeamsToProjectWithParams(ctx context.Context, args *AddAllTeamsToProjectApiParams) AddAllTeamsToProjectApiRequest

	// Interface only available internally
	addAllTeamsToProjectExecute(r AddAllTeamsToProjectApiRequest) (*PaginatedTeamRole, *http.Response, error)

	/*
		AddTeamUser Assign MongoDB Cloud Users from One Organization to One Team

		Adds one or more MongoDB Cloud users from the specified organization to the specified team. Teams enable you to grant project access roles to MongoDB Cloud users. You can assign up to 250 MongoDB Cloud users from one organization to one team. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.
		@return AddTeamUserApiRequest
	*/
	AddTeamUser(ctx context.Context, orgId string, teamId string, addUserToTeam *[]AddUserToTeam) AddTeamUserApiRequest
	/*
		AddTeamUser Assign MongoDB Cloud Users from One Organization to One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddTeamUserApiParams - Parameters for the request
		@return AddTeamUserApiRequest
	*/
	AddTeamUserWithParams(ctx context.Context, args *AddTeamUserApiParams) AddTeamUserApiRequest

	// Interface only available internally
	addTeamUserExecute(r AddTeamUserApiRequest) (*PaginatedApiAppUser, *http.Response, error)

	/*
		CreateTeam Create One Team in One Organization

		Creates one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud limits the number of teams to a maximum of 250 teams per organization. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
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

	// Interface only available internally
	createTeamExecute(r CreateTeamApiRequest) (*Team, *http.Response, error)

	/*
		DeleteTeam Remove One Team from One Organization

		Removes one team specified using its unique 24-hexadecimal digit identifier from the organization specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	deleteTeamExecute(r DeleteTeamApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetTeamById Return One Team using its ID

		Returns one team that you identified using its unique 24-hexadecimal digit ID. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an  Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team whose information you want to return.
		@return GetTeamByIdApiRequest
	*/
	GetTeamById(ctx context.Context, orgId string, teamId string) GetTeamByIdApiRequest
	/*
		GetTeamById Return One Team using its ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetTeamByIdApiParams - Parameters for the request
		@return GetTeamByIdApiRequest
	*/
	GetTeamByIdWithParams(ctx context.Context, args *GetTeamByIdApiParams) GetTeamByIdApiRequest

	// Interface only available internally
	getTeamByIdExecute(r GetTeamByIdApiRequest) (*TeamResponse, *http.Response, error)

	/*
		GetTeamByName Return One Team using its Name

		Returns one team that you identified using its human-readable name. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamName Name of the team whose information you want to return.
		@return GetTeamByNameApiRequest
	*/
	GetTeamByName(ctx context.Context, orgId string, teamName string) GetTeamByNameApiRequest
	/*
		GetTeamByName Return One Team using its Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetTeamByNameApiParams - Parameters for the request
		@return GetTeamByNameApiRequest
	*/
	GetTeamByNameWithParams(ctx context.Context, args *GetTeamByNameApiParams) GetTeamByNameApiRequest

	// Interface only available internally
	getTeamByNameExecute(r GetTeamByNameApiRequest) (*TeamResponse, *http.Response, error)

	/*
		ListOrganizationTeams Return All Teams in One Organization

		Returns all teams that belong to the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud only returns teams for which you have access. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	listOrganizationTeamsExecute(r ListOrganizationTeamsApiRequest) (*PaginatedTeam, *http.Response, error)

	/*
		ListProjectTeams Return All Teams in One Project

		Returns all teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

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

	// Interface only available internally
	listProjectTeamsExecute(r ListProjectTeamsApiRequest) (*PaginatedTeamRole, *http.Response, error)

	/*
		ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team

		Returns all MongoDB Cloud users assigned to the team specified using its unique 24-hexadecimal digit identifier. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.
		@return ListTeamUsersApiRequest
	*/
	ListTeamUsers(ctx context.Context, orgId string, teamId string) ListTeamUsersApiRequest
	/*
		ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListTeamUsersApiParams - Parameters for the request
		@return ListTeamUsersApiRequest
	*/
	ListTeamUsersWithParams(ctx context.Context, args *ListTeamUsersApiParams) ListTeamUsersApiRequest

	// Interface only available internally
	listTeamUsersExecute(r ListTeamUsersApiRequest) (*PaginatedApiAppUser, *http.Response, error)

	/*
		RemoveProjectTeam Remove One Team from One Project

		Removes one team specified using its unique 24-hexadecimal digit identifier from the project specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting API Key must have the Project User Admin role.

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

	// Interface only available internally
	removeProjectTeamExecute(r RemoveProjectTeamApiRequest) (*http.Response, error)

	/*
		RemoveTeamUser Remove One MongoDB Cloud User from One Team

		Removes one MongoDB Cloud user from the specified team. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.
		@param userId Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.
		@return RemoveTeamUserApiRequest
	*/
	RemoveTeamUser(ctx context.Context, orgId string, teamId string, userId string) RemoveTeamUserApiRequest
	/*
		RemoveTeamUser Remove One MongoDB Cloud User from One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveTeamUserApiParams - Parameters for the request
		@return RemoveTeamUserApiRequest
	*/
	RemoveTeamUserWithParams(ctx context.Context, args *RemoveTeamUserApiParams) RemoveTeamUserApiRequest

	// Interface only available internally
	removeTeamUserExecute(r RemoveTeamUserApiRequest) (*http.Response, error)

	/*
		RenameTeam Rename One Team

		[experimental] Renames one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to rename.
		@return RenameTeamApiRequest
	*/
	RenameTeam(ctx context.Context, orgId string, teamId string, team *Team) RenameTeamApiRequest
	/*
		RenameTeam Rename One Team


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RenameTeamApiParams - Parameters for the request
		@return RenameTeamApiRequest
	*/
	RenameTeamWithParams(ctx context.Context, args *RenameTeamApiParams) RenameTeamApiRequest

	// Interface only available internally
	renameTeamExecute(r RenameTeamApiRequest) (*TeamResponse, *http.Response, error)

	/*
		UpdateTeamRoles Update Team Roles in One Project

		Updates the project roles assigned to the specified team. You can grant team roles for specific projects and grant project access roles to users in the team. All members of the team share the same project access. To use this resource, the requesting API Key must have the Project User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param teamId Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.
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

	// Interface only available internally
	updateTeamRolesExecute(r UpdateTeamRolesApiRequest) (*PaginatedTeamRole, *http.Response, error)
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
	return r.ApiService.addAllTeamsToProjectExecute(r)
}

/*
AddAllTeamsToProject Add One or More Teams to One Project

Adds one team to the specified project. All members of the team share the same project access. MongoDB Cloud limits the number of users to a maximum of 100 teams per project and a maximum of 250 teams per organization. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) addAllTeamsToProjectExecute(r AddAllTeamsToProjectApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.AddAllTeamsToProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
	return r.ApiService.addTeamUserExecute(r)
}

/*
AddTeamUser Assign MongoDB Cloud Users from One Organization to One Team

Adds one or more MongoDB Cloud users from the specified organization to the specified team. Teams enable you to grant project access roles to MongoDB Cloud users. You can assign up to 250 MongoDB Cloud users from one organization to one team. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.
	@return AddTeamUserApiRequest
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

// Execute executes the request
//
//	@return PaginatedApiAppUser
func (a *TeamsApiService) addTeamUserExecute(r AddTeamUserApiRequest) (*PaginatedApiAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.AddTeamUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
	return r.ApiService.createTeamExecute(r)
}

/*
CreateTeam Create One Team in One Organization

Creates one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud limits the number of teams to a maximum of 250 teams per organization. To use this resource, the requesting API Key must have the Organization Owner role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return Team
func (a *TeamsApiService) createTeamExecute(r CreateTeamApiRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.CreateTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams"
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

func (r DeleteTeamApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteTeamExecute(r)
}

/*
DeleteTeam Remove One Team from One Organization

Removes one team specified using its unique 24-hexadecimal digit identifier from the organization specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return map[string]interface{}
func (a *TeamsApiService) deleteTeamExecute(r DeleteTeamApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.DeleteTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
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
	return r.ApiService.getTeamByIdExecute(r)
}

/*
GetTeamById Return One Team using its ID

Returns one team that you identified using its unique 24-hexadecimal digit ID. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an  Access List.

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

// Execute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) getTeamByIdExecute(r GetTeamByIdApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetTeamById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
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
	return r.ApiService.getTeamByNameExecute(r)
}

/*
GetTeamByName Return One Team using its Name

Returns one team that you identified using its human-readable name. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) getTeamByNameExecute(r GetTeamByNameApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetTeamByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/byName/{teamName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamName"+"}", url.PathEscape(parameterValueToString(r.teamName, "teamName")), -1)

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
	return r.ApiService.listOrganizationTeamsExecute(r)
}

/*
ListOrganizationTeams Return All Teams in One Organization

Returns all teams that belong to the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. MongoDB Cloud only returns teams for which you have access. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return PaginatedTeam
func (a *TeamsApiService) listOrganizationTeamsExecute(r ListOrganizationTeamsApiRequest) (*PaginatedTeam, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListOrganizationTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams"
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
	return r.ApiService.listProjectTeamsExecute(r)
}

/*
ListProjectTeams Return All Teams in One Project

Returns all teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier. All members of the team share the same project access. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

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

// Execute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) listProjectTeamsExecute(r ListProjectTeamsApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListProjectTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
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

type ListTeamUsersApiRequest struct {
	ctx          context.Context
	ApiService   TeamsApi
	orgId        string
	teamId       string
	itemsPerPage *int
	pageNum      *int
}

type ListTeamUsersApiParams struct {
	OrgId        string
	TeamId       string
	ItemsPerPage *int
	PageNum      *int
}

func (a *TeamsApiService) ListTeamUsersWithParams(ctx context.Context, args *ListTeamUsersApiParams) ListTeamUsersApiRequest {
	return ListTeamUsersApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		teamId:       args.TeamId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
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

func (r ListTeamUsersApiRequest) Execute() (*PaginatedApiAppUser, *http.Response, error) {
	return r.ApiService.listTeamUsersExecute(r)
}

/*
ListTeamUsers Return All MongoDB Cloud Users Assigned to One Team

Returns all MongoDB Cloud users assigned to the team specified using its unique 24-hexadecimal digit identifier. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.
	@return ListTeamUsersApiRequest
*/
func (a *TeamsApiService) ListTeamUsers(ctx context.Context, orgId string, teamId string) ListTeamUsersApiRequest {
	return ListTeamUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAppUser
func (a *TeamsApiService) listTeamUsersExecute(r ListTeamUsersApiRequest) (*PaginatedApiAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListTeamUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
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
	return r.ApiService.removeProjectTeamExecute(r)
}

/*
RemoveProjectTeam Remove One Team from One Project

Removes one team specified using its unique 24-hexadecimal digit identifier from the project specified using its unique 24-hexadecimal digit identifier. To use this resource, the requesting API Key must have the Project User Admin role.

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

// Execute executes the request
func (a *TeamsApiService) removeProjectTeamExecute(r RemoveProjectTeamApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RemoveProjectTeam")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return nil, reportError("teamId must have less than 24 elements")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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
	return r.ApiService.removeTeamUserExecute(r)
}

/*
RemoveTeamUser Remove One MongoDB Cloud User from One Team

Removes one MongoDB Cloud user from the specified team. This team belongs to one organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.
	@param userId Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.
	@return RemoveTeamUserApiRequest
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

// Execute executes the request
func (a *TeamsApiService) removeTeamUserExecute(r RemoveTeamUserApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RemoveTeamUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return nil, reportError("teamId must have less than 24 elements")
	}
	if strlen(r.userId) < 24 {
		return nil, reportError("userId must have at least 24 elements")
	}
	if strlen(r.userId) > 24 {
		return nil, reportError("userId must have less than 24 elements")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RenameTeamApiRequest struct {
	ctx        context.Context
	ApiService TeamsApi
	orgId      string
	teamId     string
	team       *Team
}

type RenameTeamApiParams struct {
	OrgId  string
	TeamId string
	Team   *Team
}

func (a *TeamsApiService) RenameTeamWithParams(ctx context.Context, args *RenameTeamApiParams) RenameTeamApiRequest {
	return RenameTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		teamId:     args.TeamId,
		team:       args.Team,
	}
}

func (r RenameTeamApiRequest) Execute() (*TeamResponse, *http.Response, error) {
	return r.ApiService.renameTeamExecute(r)
}

/*
RenameTeam Rename One Team

[experimental] Renames one team in the specified organization. Teams enable you to grant project access roles to MongoDB Cloud users. To use this resource, the requesting API Key must have the Organization User Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param teamId Unique 24-hexadecimal digit string that identifies the team that you want to rename.
	@return RenameTeamApiRequest
*/
func (a *TeamsApiService) RenameTeam(ctx context.Context, orgId string, teamId string, team *Team) RenameTeamApiRequest {
	return RenameTeamApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		teamId:     teamId,
		team:       team,
	}
}

// Execute executes the request
//
//	@return TeamResponse
func (a *TeamsApiService) renameTeamExecute(r RenameTeamApiRequest) (*TeamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TeamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.RenameTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/teams/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
	return r.ApiService.updateTeamRolesExecute(r)
}

/*
UpdateTeamRoles Update Team Roles in One Project

Updates the project roles assigned to the specified team. You can grant team roles for specific projects and grant project access roles to users in the team. All members of the team share the same project access. To use this resource, the requesting API Key must have the Project User Admin role.

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

// Execute executes the request
//
//	@return PaginatedTeamRole
func (a *TeamsApiService) updateTeamRolesExecute(r UpdateTeamRolesApiRequest) (*PaginatedTeamRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTeamRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.UpdateTeamRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/teams/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.teamId) < 24 {
		return localVarReturnValue, nil, reportError("teamId must have at least 24 elements")
	}
	if strlen(r.teamId) > 24 {
		return localVarReturnValue, nil, reportError("teamId must have less than 24 elements")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
