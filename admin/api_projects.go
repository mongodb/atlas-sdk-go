// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ProjectsApi interface {

	/*
		AddUserToProject Add One MongoDB Cloud User to One Project

		Adds one MongoDB Cloud user to the specified project. If the MongoDB Cloud user is not a member of the project's organization, then the user must accept their invitation to the organization to access information within the specified project. If the MongoDB Cloud User is already a member of the project's organization, then they will be added to the project immediately and an invitation will not be returned by this resource. To use this resource, the requesting Service Account or API Key must have the Group User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupInvitationRequest Adds one MongoDB Cloud user to the specified project.
		@return AddUserToProjectApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	AddUserToProject(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) AddUserToProjectApiRequest
	/*
		AddUserToProject Add One MongoDB Cloud User to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddUserToProjectApiParams - Parameters for the request
		@return AddUserToProjectApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	AddUserToProjectWithParams(ctx context.Context, args *AddUserToProjectApiParams) AddUserToProjectApiRequest

	// Method available only for mocking purposes
	AddUserToProjectExecute(r AddUserToProjectApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		CreateProject Create One Project

		Creates one project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param group Creates one project.
		@return CreateProjectApiRequest
	*/
	CreateProject(ctx context.Context, group *Group) CreateProjectApiRequest
	/*
		CreateProject Create One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectApiParams - Parameters for the request
		@return CreateProjectApiRequest
	*/
	CreateProjectWithParams(ctx context.Context, args *CreateProjectApiParams) CreateProjectApiRequest

	// Method available only for mocking purposes
	CreateProjectExecute(r CreateProjectApiRequest) (*Group, *http.Response, error)

	/*
		CreateProjectInvitation Invite One MongoDB Cloud User to One Project

		Invites one MongoDB Cloud user to join the specified project. The MongoDB Cloud user must accept the invitation to access information within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupInvitationRequest Invites one MongoDB Cloud user to join the specified project.
		@return CreateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	CreateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) CreateProjectInvitationApiRequest
	/*
		CreateProjectInvitation Invite One MongoDB Cloud User to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectInvitationApiParams - Parameters for the request
		@return CreateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	CreateProjectInvitationWithParams(ctx context.Context, args *CreateProjectInvitationApiParams) CreateProjectInvitationApiRequest

	// Method available only for mocking purposes
	CreateProjectInvitationExecute(r CreateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		DeleteProject Remove One Project

		Removes the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. You can delete a project only if there are no Online Archives for the clusters in the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectApiRequest
	*/
	DeleteProject(ctx context.Context, groupId string) DeleteProjectApiRequest
	/*
		DeleteProject Remove One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectApiParams - Parameters for the request
		@return DeleteProjectApiRequest
	*/
	DeleteProjectWithParams(ctx context.Context, args *DeleteProjectApiParams) DeleteProjectApiRequest

	// Method available only for mocking purposes
	DeleteProjectExecute(r DeleteProjectApiRequest) (*http.Response, error)

	/*
		DeleteProjectInvitation Remove One Project Invitation

		Cancels one pending invitation sent to the specified MongoDB Cloud user to join a project. You can't cancel an invitation that the user accepted. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return DeleteProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	DeleteProjectInvitation(ctx context.Context, groupId string, invitationId string) DeleteProjectInvitationApiRequest
	/*
		DeleteProjectInvitation Remove One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectInvitationApiParams - Parameters for the request
		@return DeleteProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	DeleteProjectInvitationWithParams(ctx context.Context, args *DeleteProjectInvitationApiParams) DeleteProjectInvitationApiRequest

	// Method available only for mocking purposes
	DeleteProjectInvitationExecute(r DeleteProjectInvitationApiRequest) (*http.Response, error)

	/*
		DeleteProjectLimit Remove One Project Limit

		Removes the specified project limit. Depending on the limit, Atlas either resets the limit to its default value or removes the limit entirely. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectLimitApiRequest
	*/
	DeleteProjectLimit(ctx context.Context, limitName string, groupId string) DeleteProjectLimitApiRequest
	/*
		DeleteProjectLimit Remove One Project Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectLimitApiParams - Parameters for the request
		@return DeleteProjectLimitApiRequest
	*/
	DeleteProjectLimitWithParams(ctx context.Context, args *DeleteProjectLimitApiParams) DeleteProjectLimitApiRequest

	// Method available only for mocking purposes
	DeleteProjectLimitExecute(r DeleteProjectLimitApiRequest) (*http.Response, error)

	/*
		GetProject Return One Project

		Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectApiRequest
	*/
	GetProject(ctx context.Context, groupId string) GetProjectApiRequest
	/*
		GetProject Return One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectApiParams - Parameters for the request
		@return GetProjectApiRequest
	*/
	GetProjectWithParams(ctx context.Context, args *GetProjectApiParams) GetProjectApiRequest

	// Method available only for mocking purposes
	GetProjectExecute(r GetProjectApiRequest) (*Group, *http.Response, error)

	/*
		GetProjectByName Return One Project by Name

		Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupName Human-readable label that identifies this project.
		@return GetProjectByNameApiRequest
	*/
	GetProjectByName(ctx context.Context, groupName string) GetProjectByNameApiRequest
	/*
		GetProjectByName Return One Project by Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectByNameApiParams - Parameters for the request
		@return GetProjectByNameApiRequest
	*/
	GetProjectByNameWithParams(ctx context.Context, args *GetProjectByNameApiParams) GetProjectByNameApiRequest

	// Method available only for mocking purposes
	GetProjectByNameExecute(r GetProjectByNameApiRequest) (*Group, *http.Response, error)

	/*
		GetProjectInvitation Return One Project Invitation

		Returns the details of one pending invitation to the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return GetProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	GetProjectInvitation(ctx context.Context, groupId string, invitationId string) GetProjectInvitationApiRequest
	/*
		GetProjectInvitation Return One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectInvitationApiParams - Parameters for the request
		@return GetProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	GetProjectInvitationWithParams(ctx context.Context, args *GetProjectInvitationApiParams) GetProjectInvitationApiRequest

	// Method available only for mocking purposes
	GetProjectInvitationExecute(r GetProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		GetProjectLimit Return One Limit for One Project

		Returns the specified limit for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectLimitApiRequest
	*/
	GetProjectLimit(ctx context.Context, limitName string, groupId string) GetProjectLimitApiRequest
	/*
		GetProjectLimit Return One Limit for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectLimitApiParams - Parameters for the request
		@return GetProjectLimitApiRequest
	*/
	GetProjectLimitWithParams(ctx context.Context, args *GetProjectLimitApiParams) GetProjectLimitApiRequest

	// Method available only for mocking purposes
	GetProjectLimitExecute(r GetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error)

	/*
		GetProjectLtsVersions Return All Available MongoDB LTS Versions for Clusters in One Project

		Returns the MongoDB Long Term Support Major Versions available to new clusters in this project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectLtsVersionsApiRequest
	*/
	GetProjectLtsVersions(ctx context.Context, groupId string) GetProjectLtsVersionsApiRequest
	/*
		GetProjectLtsVersions Return All Available MongoDB LTS Versions for Clusters in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectLtsVersionsApiParams - Parameters for the request
		@return GetProjectLtsVersionsApiRequest
	*/
	GetProjectLtsVersionsWithParams(ctx context.Context, args *GetProjectLtsVersionsApiParams) GetProjectLtsVersionsApiRequest

	// Method available only for mocking purposes
	GetProjectLtsVersionsExecute(r GetProjectLtsVersionsApiRequest) (*PaginatedAvailableVersion, *http.Response, error)

	/*
		GetProjectSettings Return Project Settings

		Returns details about the specified project's settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectSettingsApiRequest
	*/
	GetProjectSettings(ctx context.Context, groupId string) GetProjectSettingsApiRequest
	/*
		GetProjectSettings Return Project Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectSettingsApiParams - Parameters for the request
		@return GetProjectSettingsApiRequest
	*/
	GetProjectSettingsWithParams(ctx context.Context, args *GetProjectSettingsApiParams) GetProjectSettingsApiRequest

	// Method available only for mocking purposes
	GetProjectSettingsExecute(r GetProjectSettingsApiRequest) (*GroupSettings, *http.Response, error)

	/*
		ListProjectInvitations Return All Project Invitations

		Returns all pending invitations to the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectInvitationsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	ListProjectInvitations(ctx context.Context, groupId string) ListProjectInvitationsApiRequest
	/*
		ListProjectInvitations Return All Project Invitations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectInvitationsApiParams - Parameters for the request
		@return ListProjectInvitationsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	ListProjectInvitationsWithParams(ctx context.Context, args *ListProjectInvitationsApiParams) ListProjectInvitationsApiRequest

	// Method available only for mocking purposes
	ListProjectInvitationsExecute(r ListProjectInvitationsApiRequest) ([]GroupInvitation, *http.Response, error)

	/*
		ListProjectLimits Return All Limits for One Project

		Returns all the limits for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectLimitsApiRequest
	*/
	ListProjectLimits(ctx context.Context, groupId string) ListProjectLimitsApiRequest
	/*
		ListProjectLimits Return All Limits for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectLimitsApiParams - Parameters for the request
		@return ListProjectLimitsApiRequest
	*/
	ListProjectLimitsWithParams(ctx context.Context, args *ListProjectLimitsApiParams) ListProjectLimitsApiRequest

	// Method available only for mocking purposes
	ListProjectLimitsExecute(r ListProjectLimitsApiRequest) ([]DataFederationLimit, *http.Response, error)

	/*
		ListProjects Return All Projects

		Returns details about all projects. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Organization Read Only role or higher.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListProjectsApiRequest
	*/
	ListProjects(ctx context.Context) ListProjectsApiRequest
	/*
		ListProjects Return All Projects


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectsApiParams - Parameters for the request
		@return ListProjectsApiRequest
	*/
	ListProjectsWithParams(ctx context.Context, args *ListProjectsApiParams) ListProjectsApiRequest

	// Method available only for mocking purposes
	ListProjectsExecute(r ListProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error)

	/*
		MigrateProjectToAnotherOrg Migrate One Project to Another Organization

		Migrates a project from its current organization to another organization. All project users and their roles will be copied to the same project in the destination organization. You must include an organization API key with the Organization Owner role for the destination organization to verify access to the destination organization when you authenticate with Programmatic API Keys. Otherwise, the requesting user must have the Organization Owner role in both organizations. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupMigrationRequest Migrates a project from its current organization to another organization.
		@return MigrateProjectToAnotherOrgApiRequest
	*/
	MigrateProjectToAnotherOrg(ctx context.Context, groupId string, groupMigrationRequest *GroupMigrationRequest) MigrateProjectToAnotherOrgApiRequest
	/*
		MigrateProjectToAnotherOrg Migrate One Project to Another Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param MigrateProjectToAnotherOrgApiParams - Parameters for the request
		@return MigrateProjectToAnotherOrgApiRequest
	*/
	MigrateProjectToAnotherOrgWithParams(ctx context.Context, args *MigrateProjectToAnotherOrgApiParams) MigrateProjectToAnotherOrgApiRequest

	// Method available only for mocking purposes
	MigrateProjectToAnotherOrgExecute(r MigrateProjectToAnotherOrgApiRequest) (*Group, *http.Response, error)

	/*
		ReturnAllIpAddresses Return All IP Addresses for One Project

		Returns all IP addresses for this project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ReturnAllIpAddressesApiRequest
	*/
	ReturnAllIpAddresses(ctx context.Context, groupId string) ReturnAllIpAddressesApiRequest
	/*
		ReturnAllIpAddresses Return All IP Addresses for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnAllIpAddressesApiParams - Parameters for the request
		@return ReturnAllIpAddressesApiRequest
	*/
	ReturnAllIpAddressesWithParams(ctx context.Context, args *ReturnAllIpAddressesApiParams) ReturnAllIpAddressesApiRequest

	// Method available only for mocking purposes
	ReturnAllIpAddressesExecute(r ReturnAllIpAddressesApiRequest) (*GroupIPAddresses, *http.Response, error)

	/*
			SetProjectLimit Set One Project Limit

			Sets the specified project limit. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		**NOTE**: Increasing the following configuration limits might lead to slower response times in the MongoDB Cloud UI or increased user management overhead leading to authentication or authorization re-architecture. If possible, we recommend that you create additional projects to gain access to more of these resources for a more sustainable growth pattern.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param dataFederationLimit Limit to update.
			@return SetProjectLimitApiRequest
	*/
	SetProjectLimit(ctx context.Context, limitName string, groupId string, dataFederationLimit *DataFederationLimit) SetProjectLimitApiRequest
	/*
		SetProjectLimit Set One Project Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param SetProjectLimitApiParams - Parameters for the request
		@return SetProjectLimitApiRequest
	*/
	SetProjectLimitWithParams(ctx context.Context, args *SetProjectLimitApiParams) SetProjectLimitApiRequest

	// Method available only for mocking purposes
	SetProjectLimitExecute(r SetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error)

	/*
		UpdateProject Update One Project

		Updates the human-readable label that identifies the specified project, or the tags associated with the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupUpdate Project to update.
		@return UpdateProjectApiRequest
	*/
	UpdateProject(ctx context.Context, groupId string, groupUpdate *GroupUpdate) UpdateProjectApiRequest
	/*
		UpdateProject Update One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectApiParams - Parameters for the request
		@return UpdateProjectApiRequest
	*/
	UpdateProjectWithParams(ctx context.Context, args *UpdateProjectApiParams) UpdateProjectApiRequest

	// Method available only for mocking purposes
	UpdateProjectExecute(r UpdateProjectApiRequest) (*Group, *http.Response, error)

	/*
		UpdateProjectInvitation Update One Project Invitation

		Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the username of the invited user. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupInvitationRequest Updates the details of one pending invitation to the specified project.
		@return UpdateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) UpdateProjectInvitationApiRequest
	/*
		UpdateProjectInvitation Update One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectInvitationApiParams - Parameters for the request
		@return UpdateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationWithParams(ctx context.Context, args *UpdateProjectInvitationApiParams) UpdateProjectInvitationApiRequest

	// Method available only for mocking purposes
	UpdateProjectInvitationExecute(r UpdateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		UpdateProjectInvitationById Update One Project Invitation by Invitation ID

		Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the unique identification string for that invitation. Use the Return All Project Invitations endpoint to retrieve IDs for all pending project invitations. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@param groupInvitationUpdateRequest Updates the details of one pending invitation to the specified project.
		@return UpdateProjectInvitationByIdApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationById(ctx context.Context, groupId string, invitationId string, groupInvitationUpdateRequest *GroupInvitationUpdateRequest) UpdateProjectInvitationByIdApiRequest
	/*
		UpdateProjectInvitationById Update One Project Invitation by Invitation ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectInvitationByIdApiParams - Parameters for the request
		@return UpdateProjectInvitationByIdApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationByIdWithParams(ctx context.Context, args *UpdateProjectInvitationByIdApiParams) UpdateProjectInvitationByIdApiRequest

	// Method available only for mocking purposes
	UpdateProjectInvitationByIdExecute(r UpdateProjectInvitationByIdApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		UpdateProjectRoles Update Project Roles for One MongoDB Cloud User

		Updates the roles of the specified user in the specified project. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified project. To use this resource, the requesting Service Account or API Key must have the Group User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
		@param updateGroupRolesForUser Roles to update for the specified user.
		@return UpdateProjectRolesApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectRoles(ctx context.Context, groupId string, userId string, updateGroupRolesForUser *UpdateGroupRolesForUser) UpdateProjectRolesApiRequest
	/*
		UpdateProjectRoles Update Project Roles for One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectRolesApiParams - Parameters for the request
		@return UpdateProjectRolesApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectRolesWithParams(ctx context.Context, args *UpdateProjectRolesApiParams) UpdateProjectRolesApiRequest

	// Method available only for mocking purposes
	UpdateProjectRolesExecute(r UpdateProjectRolesApiRequest) (*UpdateGroupRolesForUser, *http.Response, error)

	/*
		UpdateProjectSettings Update Project Settings

		Updates the settings of the specified project. You can update any of the options available. MongoDB cloud only updates the options provided in the request. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupSettings Settings to update.
		@return UpdateProjectSettingsApiRequest
	*/
	UpdateProjectSettings(ctx context.Context, groupId string, groupSettings *GroupSettings) UpdateProjectSettingsApiRequest
	/*
		UpdateProjectSettings Update Project Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectSettingsApiParams - Parameters for the request
		@return UpdateProjectSettingsApiRequest
	*/
	UpdateProjectSettingsWithParams(ctx context.Context, args *UpdateProjectSettingsApiParams) UpdateProjectSettingsApiRequest

	// Method available only for mocking purposes
	UpdateProjectSettingsExecute(r UpdateProjectSettingsApiRequest) (*GroupSettings, *http.Response, error)
}

// ProjectsApiService ProjectsApi service
type ProjectsApiService service

type AddUserToProjectApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type AddUserToProjectApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) AddUserToProjectWithParams(ctx context.Context, args *AddUserToProjectApiParams) AddUserToProjectApiRequest {
	return AddUserToProjectApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r AddUserToProjectApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.AddUserToProjectExecute(r)
}

/*
AddUserToProject Add One MongoDB Cloud User to One Project

Adds one MongoDB Cloud user to the specified project. If the MongoDB Cloud user is not a member of the project's organization, then the user must accept their invitation to the organization to access information within the specified project. If the MongoDB Cloud User is already a member of the project's organization, then they will be added to the project immediately and an invitation will not be returned by this resource. To use this resource, the requesting Service Account or API Key must have the Group User Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AddUserToProjectApiRequest

Deprecated
*/
func (a *ProjectsApiService) AddUserToProject(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) AddUserToProjectApiRequest {
	return AddUserToProjectApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// AddUserToProjectExecute executes the request
//
//	@return OrganizationInvitation
//
// Deprecated
func (a *ProjectsApiService) AddUserToProjectExecute(r AddUserToProjectApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.AddUserToProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/access"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupInvitationRequest
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

type CreateProjectApiRequest struct {
	ctx            context.Context
	ApiService     ProjectsApi
	group          *Group
	projectOwnerId *string
}

type CreateProjectApiParams struct {
	Group          *Group
	ProjectOwnerId *string
}

func (a *ProjectsApiService) CreateProjectWithParams(ctx context.Context, args *CreateProjectApiParams) CreateProjectApiRequest {
	return CreateProjectApiRequest{
		ApiService:     a,
		ctx:            ctx,
		group:          args.Group,
		projectOwnerId: args.ProjectOwnerId,
	}
}

// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user to whom to grant the Project Owner role on the specified project. If you set this parameter, it overrides the default value of the oldest Organization Owner.
func (r CreateProjectApiRequest) ProjectOwnerId(projectOwnerId string) CreateProjectApiRequest {
	r.projectOwnerId = &projectOwnerId
	return r
}

func (r CreateProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.CreateProjectExecute(r)
}

/*
CreateProject Create One Project

Creates one project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateProjectApiRequest
*/
func (a *ProjectsApiService) CreateProject(ctx context.Context, group *Group) CreateProjectApiRequest {
	return CreateProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		group:      group,
	}
}

// CreateProjectExecute executes the request
//
//	@return Group
func (a *ProjectsApiService) CreateProjectExecute(r CreateProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreateProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.group == nil {
		return localVarReturnValue, nil, reportError("group is required and must be specified")
	}

	if r.projectOwnerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectOwnerId", r.projectOwnerId, "")
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
	localVarPostBody = r.group
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

type CreateProjectInvitationApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type CreateProjectInvitationApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) CreateProjectInvitationWithParams(ctx context.Context, args *CreateProjectInvitationApiParams) CreateProjectInvitationApiRequest {
	return CreateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r CreateProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.CreateProjectInvitationExecute(r)
}

/*
CreateProjectInvitation Invite One MongoDB Cloud User to One Project

Invites one MongoDB Cloud user to join the specified project. The MongoDB Cloud user must accept the invitation to access information within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) CreateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) CreateProjectInvitationApiRequest {
	return CreateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// CreateProjectInvitationExecute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) CreateProjectInvitationExecute(r CreateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreateProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationRequest
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

type DeleteProjectApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type DeleteProjectApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) DeleteProjectWithParams(ctx context.Context, args *DeleteProjectApiParams) DeleteProjectApiRequest {
	return DeleteProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectExecute(r)
}

/*
DeleteProject Remove One Project

Removes the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. You can delete a project only if there are no Online Archives for the clusters in the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectApiRequest
*/
func (a *ProjectsApiService) DeleteProject(ctx context.Context, groupId string) DeleteProjectApiRequest {
	return DeleteProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// DeleteProjectExecute executes the request
func (a *ProjectsApiService) DeleteProjectExecute(r DeleteProjectApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProject")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type DeleteProjectInvitationApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	groupId      string
	invitationId string
}

type DeleteProjectInvitationApiParams struct {
	GroupId      string
	InvitationId string
}

func (a *ProjectsApiService) DeleteProjectInvitationWithParams(ctx context.Context, args *DeleteProjectInvitationApiParams) DeleteProjectInvitationApiRequest {
	return DeleteProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		invitationId: args.InvitationId,
	}
}

func (r DeleteProjectInvitationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectInvitationExecute(r)
}

/*
DeleteProjectInvitation Remove One Project Invitation

Cancels one pending invitation sent to the specified MongoDB Cloud user to join a project. You can't cancel an invitation that the user accepted. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return DeleteProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) DeleteProjectInvitation(ctx context.Context, groupId string, invitationId string) DeleteProjectInvitationApiRequest {
	return DeleteProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		invitationId: invitationId,
	}
}

// DeleteProjectInvitationExecute executes the request
// Deprecated
func (a *ProjectsApiService) DeleteProjectInvitationExecute(r DeleteProjectInvitationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProjectInvitation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
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

type DeleteProjectLimitApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	limitName  string
	groupId    string
}

type DeleteProjectLimitApiParams struct {
	LimitName string
	GroupId   string
}

func (a *ProjectsApiService) DeleteProjectLimitWithParams(ctx context.Context, args *DeleteProjectLimitApiParams) DeleteProjectLimitApiRequest {
	return DeleteProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  args.LimitName,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectLimitApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectLimitExecute(r)
}

/*
DeleteProjectLimit Remove One Project Limit

Removes the specified project limit. Depending on the limit, Atlas either resets the limit to its default value or removes the limit entirely. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectLimitApiRequest
*/
func (a *ProjectsApiService) DeleteProjectLimit(ctx context.Context, limitName string, groupId string) DeleteProjectLimitApiRequest {
	return DeleteProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  limitName,
		groupId:    groupId,
	}
}

// DeleteProjectLimitExecute executes the request
func (a *ProjectsApiService) DeleteProjectLimitExecute(r DeleteProjectLimitApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProjectLimit")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	if r.limitName == "" {
		return nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type GetProjectApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type GetProjectApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) GetProjectWithParams(ctx context.Context, args *GetProjectApiParams) GetProjectApiRequest {
	return GetProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.GetProjectExecute(r)
}

/*
GetProject Return One Project

Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectApiRequest
*/
func (a *ProjectsApiService) GetProject(ctx context.Context, groupId string) GetProjectApiRequest {
	return GetProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetProjectExecute executes the request
//
//	@return Group
func (a *ProjectsApiService) GetProjectExecute(r GetProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type GetProjectByNameApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupName  string
}

type GetProjectByNameApiParams struct {
	GroupName string
}

func (a *ProjectsApiService) GetProjectByNameWithParams(ctx context.Context, args *GetProjectByNameApiParams) GetProjectByNameApiRequest {
	return GetProjectByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupName:  args.GroupName,
	}
}

func (r GetProjectByNameApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.GetProjectByNameExecute(r)
}

/*
GetProjectByName Return One Project by Name

Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupName Human-readable label that identifies this project.
	@return GetProjectByNameApiRequest
*/
func (a *ProjectsApiService) GetProjectByName(ctx context.Context, groupName string) GetProjectByNameApiRequest {
	return GetProjectByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupName:  groupName,
	}
}

// GetProjectByNameExecute executes the request
//
//	@return Group
func (a *ProjectsApiService) GetProjectByNameExecute(r GetProjectByNameApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/byName/{groupName}"
	if r.groupName == "" {
		return localVarReturnValue, nil, reportError("groupName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", url.PathEscape(r.groupName), -1)

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

type GetProjectInvitationApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	groupId      string
	invitationId string
}

type GetProjectInvitationApiParams struct {
	GroupId      string
	InvitationId string
}

func (a *ProjectsApiService) GetProjectInvitationWithParams(ctx context.Context, args *GetProjectInvitationApiParams) GetProjectInvitationApiRequest {
	return GetProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		invitationId: args.InvitationId,
	}
}

func (r GetProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.GetProjectInvitationExecute(r)
}

/*
GetProjectInvitation Return One Project Invitation

Returns the details of one pending invitation to the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return GetProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) GetProjectInvitation(ctx context.Context, groupId string, invitationId string) GetProjectInvitationApiRequest {
	return GetProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		invitationId: invitationId,
	}
}

// GetProjectInvitationExecute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) GetProjectInvitationExecute(r GetProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
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

type GetProjectLimitApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	limitName  string
	groupId    string
}

type GetProjectLimitApiParams struct {
	LimitName string
	GroupId   string
}

func (a *ProjectsApiService) GetProjectLimitWithParams(ctx context.Context, args *GetProjectLimitApiParams) GetProjectLimitApiRequest {
	return GetProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  args.LimitName,
		groupId:    args.GroupId,
	}
}

func (r GetProjectLimitApiRequest) Execute() (*DataFederationLimit, *http.Response, error) {
	return r.ApiService.GetProjectLimitExecute(r)
}

/*
GetProjectLimit Return One Limit for One Project

Returns the specified limit for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectLimitApiRequest
*/
func (a *ProjectsApiService) GetProjectLimit(ctx context.Context, limitName string, groupId string) GetProjectLimitApiRequest {
	return GetProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  limitName,
		groupId:    groupId,
	}
}

// GetProjectLimitExecute executes the request
//
//	@return DataFederationLimit
func (a *ProjectsApiService) GetProjectLimitExecute(r GetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	if r.limitName == "" {
		return localVarReturnValue, nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type GetProjectLtsVersionsApiRequest struct {
	ctx           context.Context
	ApiService    ProjectsApi
	groupId       string
	cloudProvider *string
	instanceSize  *string
	defaultStatus *string
	itemsPerPage  *int64
	pageNum       *int
}

type GetProjectLtsVersionsApiParams struct {
	GroupId       string
	CloudProvider *string
	InstanceSize  *string
	DefaultStatus *string
	ItemsPerPage  *int64
	PageNum       *int
}

func (a *ProjectsApiService) GetProjectLtsVersionsWithParams(ctx context.Context, args *GetProjectLtsVersionsApiParams) GetProjectLtsVersionsApiRequest {
	return GetProjectLtsVersionsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		cloudProvider: args.CloudProvider,
		instanceSize:  args.InstanceSize,
		defaultStatus: args.DefaultStatus,
		itemsPerPage:  args.ItemsPerPage,
		pageNum:       args.PageNum,
	}
}

// Filter results to only one cloud provider.
func (r GetProjectLtsVersionsApiRequest) CloudProvider(cloudProvider string) GetProjectLtsVersionsApiRequest {
	r.cloudProvider = &cloudProvider
	return r
}

// Filter results to only one instance size.
func (r GetProjectLtsVersionsApiRequest) InstanceSize(instanceSize string) GetProjectLtsVersionsApiRequest {
	r.instanceSize = &instanceSize
	return r
}

// Filter results to only the default values per tier. This value must be DEFAULT.
func (r GetProjectLtsVersionsApiRequest) DefaultStatus(defaultStatus string) GetProjectLtsVersionsApiRequest {
	r.defaultStatus = &defaultStatus
	return r
}

// Number of items that the response returns per page.
func (r GetProjectLtsVersionsApiRequest) ItemsPerPage(itemsPerPage int64) GetProjectLtsVersionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetProjectLtsVersionsApiRequest) PageNum(pageNum int) GetProjectLtsVersionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetProjectLtsVersionsApiRequest) Execute() (*PaginatedAvailableVersion, *http.Response, error) {
	return r.ApiService.GetProjectLtsVersionsExecute(r)
}

/*
GetProjectLtsVersions Return All Available MongoDB LTS Versions for Clusters in One Project

Returns the MongoDB Long Term Support Major Versions available to new clusters in this project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectLtsVersionsApiRequest
*/
func (a *ProjectsApiService) GetProjectLtsVersions(ctx context.Context, groupId string) GetProjectLtsVersionsApiRequest {
	return GetProjectLtsVersionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetProjectLtsVersionsExecute executes the request
//
//	@return PaginatedAvailableVersion
func (a *ProjectsApiService) GetProjectLtsVersionsExecute(r GetProjectLtsVersionsApiRequest) (*PaginatedAvailableVersion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAvailableVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectLtsVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mongoDBVersions"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cloudProvider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cloudProvider", r.cloudProvider, "")
	}
	if r.instanceSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "instanceSize", r.instanceSize, "")
	}
	if r.defaultStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "defaultStatus", r.defaultStatus, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int64 = 100
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

type GetProjectSettingsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type GetProjectSettingsApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) GetProjectSettingsWithParams(ctx context.Context, args *GetProjectSettingsApiParams) GetProjectSettingsApiRequest {
	return GetProjectSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetProjectSettingsApiRequest) Execute() (*GroupSettings, *http.Response, error) {
	return r.ApiService.GetProjectSettingsExecute(r)
}

/*
GetProjectSettings Return Project Settings

Returns details about the specified project's settings. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectSettingsApiRequest
*/
func (a *ProjectsApiService) GetProjectSettings(ctx context.Context, groupId string) GetProjectSettingsApiRequest {
	return GetProjectSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetProjectSettingsExecute executes the request
//
//	@return GroupSettings
func (a *ProjectsApiService) GetProjectSettingsExecute(r GetProjectSettingsApiRequest) (*GroupSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/settings"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type ListProjectInvitationsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
	username   *string
}

type ListProjectInvitationsApiParams struct {
	GroupId  string
	Username *string
}

func (a *ProjectsApiService) ListProjectInvitationsWithParams(ctx context.Context, args *ListProjectInvitationsApiParams) ListProjectInvitationsApiRequest {
	return ListProjectInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		username:   args.Username,
	}
}

// Email address of the user account invited to this project.
func (r ListProjectInvitationsApiRequest) Username(username string) ListProjectInvitationsApiRequest {
	r.username = &username
	return r
}

func (r ListProjectInvitationsApiRequest) Execute() ([]GroupInvitation, *http.Response, error) {
	return r.ApiService.ListProjectInvitationsExecute(r)
}

/*
ListProjectInvitations Return All Project Invitations

Returns all pending invitations to the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectInvitationsApiRequest

Deprecated
*/
func (a *ProjectsApiService) ListProjectInvitations(ctx context.Context, groupId string) ListProjectInvitationsApiRequest {
	return ListProjectInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectInvitationsExecute executes the request
//
//	@return []GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) ListProjectInvitationsExecute(r ListProjectInvitationsApiRequest) ([]GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjectInvitations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type ListProjectLimitsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type ListProjectLimitsApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) ListProjectLimitsWithParams(ctx context.Context, args *ListProjectLimitsApiParams) ListProjectLimitsApiRequest {
	return ListProjectLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ListProjectLimitsApiRequest) Execute() ([]DataFederationLimit, *http.Response, error) {
	return r.ApiService.ListProjectLimitsExecute(r)
}

/*
ListProjectLimits Return All Limits for One Project

Returns all the limits for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectLimitsApiRequest
*/
func (a *ProjectsApiService) ListProjectLimits(ctx context.Context, groupId string) ListProjectLimitsApiRequest {
	return ListProjectLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectLimitsExecute executes the request
//
//	@return []DataFederationLimit
func (a *ProjectsApiService) ListProjectLimitsExecute(r ListProjectLimitsApiRequest) ([]DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjectLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type ListProjectsApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectsApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProjectsApiService) ListProjectsWithParams(ctx context.Context, args *ListProjectsApiParams) ListProjectsApiRequest {
	return ListProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectsApiRequest) IncludeCount(includeCount bool) ListProjectsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectsApiRequest) PageNum(pageNum int) ListProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectsApiRequest) Execute() (*PaginatedAtlasGroup, *http.Response, error) {
	return r.ApiService.ListProjectsExecute(r)
}

/*
ListProjects Return All Projects

Returns details about all projects. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting Service Account or API Key must have the Organization Read Only role or higher.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListProjectsApiRequest
*/
func (a *ProjectsApiService) ListProjects(ctx context.Context) ListProjectsApiRequest {
	return ListProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListProjectsExecute executes the request
//
//	@return PaginatedAtlasGroup
func (a *ProjectsApiService) ListProjectsExecute(r ListProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups"

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

type MigrateProjectToAnotherOrgApiRequest struct {
	ctx                   context.Context
	ApiService            ProjectsApi
	groupId               string
	groupMigrationRequest *GroupMigrationRequest
}

type MigrateProjectToAnotherOrgApiParams struct {
	GroupId               string
	GroupMigrationRequest *GroupMigrationRequest
}

func (a *ProjectsApiService) MigrateProjectToAnotherOrgWithParams(ctx context.Context, args *MigrateProjectToAnotherOrgApiParams) MigrateProjectToAnotherOrgApiRequest {
	return MigrateProjectToAnotherOrgApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               args.GroupId,
		groupMigrationRequest: args.GroupMigrationRequest,
	}
}

func (r MigrateProjectToAnotherOrgApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.MigrateProjectToAnotherOrgExecute(r)
}

/*
MigrateProjectToAnotherOrg Migrate One Project to Another Organization

Migrates a project from its current organization to another organization. All project users and their roles will be copied to the same project in the destination organization. You must include an organization API key with the Organization Owner role for the destination organization to verify access to the destination organization when you authenticate with Programmatic API Keys. Otherwise, the requesting user must have the Organization Owner role in both organizations. To use this resource, the requesting Service Account or API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return MigrateProjectToAnotherOrgApiRequest
*/
func (a *ProjectsApiService) MigrateProjectToAnotherOrg(ctx context.Context, groupId string, groupMigrationRequest *GroupMigrationRequest) MigrateProjectToAnotherOrgApiRequest {
	return MigrateProjectToAnotherOrgApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               groupId,
		groupMigrationRequest: groupMigrationRequest,
	}
}

// MigrateProjectToAnotherOrgExecute executes the request
//
//	@return Group
func (a *ProjectsApiService) MigrateProjectToAnotherOrgExecute(r MigrateProjectToAnotherOrgApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.MigrateProjectToAnotherOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}:migrate"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupMigrationRequest == nil {
		return localVarReturnValue, nil, reportError("groupMigrationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupMigrationRequest
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

type ReturnAllIpAddressesApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type ReturnAllIpAddressesApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) ReturnAllIpAddressesWithParams(ctx context.Context, args *ReturnAllIpAddressesApiParams) ReturnAllIpAddressesApiRequest {
	return ReturnAllIpAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ReturnAllIpAddressesApiRequest) Execute() (*GroupIPAddresses, *http.Response, error) {
	return r.ApiService.ReturnAllIpAddressesExecute(r)
}

/*
ReturnAllIpAddresses Return All IP Addresses for One Project

Returns all IP addresses for this project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ReturnAllIpAddressesApiRequest
*/
func (a *ProjectsApiService) ReturnAllIpAddresses(ctx context.Context, groupId string) ReturnAllIpAddressesApiRequest {
	return ReturnAllIpAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ReturnAllIpAddressesExecute executes the request
//
//	@return GroupIPAddresses
func (a *ProjectsApiService) ReturnAllIpAddressesExecute(r ReturnAllIpAddressesApiRequest) (*GroupIPAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupIPAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ReturnAllIpAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/ipAddresses"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type SetProjectLimitApiRequest struct {
	ctx                 context.Context
	ApiService          ProjectsApi
	limitName           string
	groupId             string
	dataFederationLimit *DataFederationLimit
}

type SetProjectLimitApiParams struct {
	LimitName           string
	GroupId             string
	DataFederationLimit *DataFederationLimit
}

func (a *ProjectsApiService) SetProjectLimitWithParams(ctx context.Context, args *SetProjectLimitApiParams) SetProjectLimitApiRequest {
	return SetProjectLimitApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		limitName:           args.LimitName,
		groupId:             args.GroupId,
		dataFederationLimit: args.DataFederationLimit,
	}
}

func (r SetProjectLimitApiRequest) Execute() (*DataFederationLimit, *http.Response, error) {
	return r.ApiService.SetProjectLimitExecute(r)
}

/*
SetProjectLimit Set One Project Limit

Sets the specified project limit. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

**NOTE**: Increasing the following configuration limits might lead to slower response times in the MongoDB Cloud UI or increased user management overhead leading to authentication or authorization re-architecture. If possible, we recommend that you create additional projects to gain access to more of these resources for a more sustainable growth pattern.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27| | atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SetProjectLimitApiRequest
*/
func (a *ProjectsApiService) SetProjectLimit(ctx context.Context, limitName string, groupId string, dataFederationLimit *DataFederationLimit) SetProjectLimitApiRequest {
	return SetProjectLimitApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		limitName:           limitName,
		groupId:             groupId,
		dataFederationLimit: dataFederationLimit,
	}
}

// SetProjectLimitExecute executes the request
//
//	@return DataFederationLimit
func (a *ProjectsApiService) SetProjectLimitExecute(r SetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.SetProjectLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	if r.limitName == "" {
		return localVarReturnValue, nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataFederationLimit == nil {
		return localVarReturnValue, nil, reportError("dataFederationLimit is required and must be specified")
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
	localVarPostBody = r.dataFederationLimit
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

type UpdateProjectApiRequest struct {
	ctx         context.Context
	ApiService  ProjectsApi
	groupId     string
	groupUpdate *GroupUpdate
}

type UpdateProjectApiParams struct {
	GroupId     string
	GroupUpdate *GroupUpdate
}

func (a *ProjectsApiService) UpdateProjectWithParams(ctx context.Context, args *UpdateProjectApiParams) UpdateProjectApiRequest {
	return UpdateProjectApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		groupUpdate: args.GroupUpdate,
	}
}

func (r UpdateProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.UpdateProjectExecute(r)
}

/*
UpdateProject Update One Project

Updates the human-readable label that identifies the specified project, or the tags associated with the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectApiRequest
*/
func (a *ProjectsApiService) UpdateProject(ctx context.Context, groupId string, groupUpdate *GroupUpdate) UpdateProjectApiRequest {
	return UpdateProjectApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		groupUpdate: groupUpdate,
	}
}

// UpdateProjectExecute executes the request
//
//	@return Group
func (a *ProjectsApiService) UpdateProjectExecute(r UpdateProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupUpdate == nil {
		return localVarReturnValue, nil, reportError("groupUpdate is required and must be specified")
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
	localVarPostBody = r.groupUpdate
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

type UpdateProjectInvitationApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type UpdateProjectInvitationApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) UpdateProjectInvitationWithParams(ctx context.Context, args *UpdateProjectInvitationApiParams) UpdateProjectInvitationApiRequest {
	return UpdateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r UpdateProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.UpdateProjectInvitationExecute(r)
}

/*
UpdateProjectInvitation Update One Project Invitation

Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the username of the invited user. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) UpdateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) UpdateProjectInvitationApiRequest {
	return UpdateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// UpdateProjectInvitationExecute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) UpdateProjectInvitationExecute(r UpdateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationRequest
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

type UpdateProjectInvitationByIdApiRequest struct {
	ctx                          context.Context
	ApiService                   ProjectsApi
	groupId                      string
	invitationId                 string
	groupInvitationUpdateRequest *GroupInvitationUpdateRequest
}

type UpdateProjectInvitationByIdApiParams struct {
	GroupId                      string
	InvitationId                 string
	GroupInvitationUpdateRequest *GroupInvitationUpdateRequest
}

func (a *ProjectsApiService) UpdateProjectInvitationByIdWithParams(ctx context.Context, args *UpdateProjectInvitationByIdApiParams) UpdateProjectInvitationByIdApiRequest {
	return UpdateProjectInvitationByIdApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      args.GroupId,
		invitationId:                 args.InvitationId,
		groupInvitationUpdateRequest: args.GroupInvitationUpdateRequest,
	}
}

func (r UpdateProjectInvitationByIdApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.UpdateProjectInvitationByIdExecute(r)
}

/*
UpdateProjectInvitationById Update One Project Invitation by Invitation ID

Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the unique identification string for that invitation. Use the Return All Project Invitations endpoint to retrieve IDs for all pending project invitations. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return UpdateProjectInvitationByIdApiRequest

Deprecated
*/
func (a *ProjectsApiService) UpdateProjectInvitationById(ctx context.Context, groupId string, invitationId string, groupInvitationUpdateRequest *GroupInvitationUpdateRequest) UpdateProjectInvitationByIdApiRequest {
	return UpdateProjectInvitationByIdApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      groupId,
		invitationId:                 invitationId,
		groupInvitationUpdateRequest: groupInvitationUpdateRequest,
	}
}

// UpdateProjectInvitationByIdExecute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) UpdateProjectInvitationByIdExecute(r UpdateProjectInvitationByIdApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectInvitationById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.invitationId == "" {
		return localVarReturnValue, nil, reportError("invitationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(r.invitationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationUpdateRequest
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

type UpdateProjectRolesApiRequest struct {
	ctx                     context.Context
	ApiService              ProjectsApi
	groupId                 string
	userId                  string
	updateGroupRolesForUser *UpdateGroupRolesForUser
}

type UpdateProjectRolesApiParams struct {
	GroupId                 string
	UserId                  string
	UpdateGroupRolesForUser *UpdateGroupRolesForUser
}

func (a *ProjectsApiService) UpdateProjectRolesWithParams(ctx context.Context, args *UpdateProjectRolesApiParams) UpdateProjectRolesApiRequest {
	return UpdateProjectRolesApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		userId:                  args.UserId,
		updateGroupRolesForUser: args.UpdateGroupRolesForUser,
	}
}

func (r UpdateProjectRolesApiRequest) Execute() (*UpdateGroupRolesForUser, *http.Response, error) {
	return r.ApiService.UpdateProjectRolesExecute(r)
}

/*
UpdateProjectRoles Update Project Roles for One MongoDB Cloud User

Updates the roles of the specified user in the specified project. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified project. To use this resource, the requesting Service Account or API Key must have the Group User Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
	@return UpdateProjectRolesApiRequest

Deprecated
*/
func (a *ProjectsApiService) UpdateProjectRoles(ctx context.Context, groupId string, userId string, updateGroupRolesForUser *UpdateGroupRolesForUser) UpdateProjectRolesApiRequest {
	return UpdateProjectRolesApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		userId:                  userId,
		updateGroupRolesForUser: updateGroupRolesForUser,
	}
}

// UpdateProjectRolesExecute executes the request
//
//	@return UpdateGroupRolesForUser
//
// Deprecated
func (a *ProjectsApiService) UpdateProjectRolesExecute(r UpdateProjectRolesApiRequest) (*UpdateGroupRolesForUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UpdateGroupRolesForUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users/{userId}/roles"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.userId == "" {
		return localVarReturnValue, nil, reportError("userId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(r.userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateGroupRolesForUser == nil {
		return localVarReturnValue, nil, reportError("updateGroupRolesForUser is required and must be specified")
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
	localVarPostBody = r.updateGroupRolesForUser
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

type UpdateProjectSettingsApiRequest struct {
	ctx           context.Context
	ApiService    ProjectsApi
	groupId       string
	groupSettings *GroupSettings
}

type UpdateProjectSettingsApiParams struct {
	GroupId       string
	GroupSettings *GroupSettings
}

func (a *ProjectsApiService) UpdateProjectSettingsWithParams(ctx context.Context, args *UpdateProjectSettingsApiParams) UpdateProjectSettingsApiRequest {
	return UpdateProjectSettingsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		groupSettings: args.GroupSettings,
	}
}

func (r UpdateProjectSettingsApiRequest) Execute() (*GroupSettings, *http.Response, error) {
	return r.ApiService.UpdateProjectSettingsExecute(r)
}

/*
UpdateProjectSettings Update Project Settings

Updates the settings of the specified project. You can update any of the options available. MongoDB cloud only updates the options provided in the request. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectSettingsApiRequest
*/
func (a *ProjectsApiService) UpdateProjectSettings(ctx context.Context, groupId string, groupSettings *GroupSettings) UpdateProjectSettingsApiRequest {
	return UpdateProjectSettingsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		groupSettings: groupSettings,
	}
}

// UpdateProjectSettingsExecute executes the request
//
//	@return GroupSettings
func (a *ProjectsApiService) UpdateProjectSettingsExecute(r UpdateProjectSettingsApiRequest) (*GroupSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/settings"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupSettings == nil {
		return localVarReturnValue, nil, reportError("groupSettings is required and must be specified")
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
	localVarPostBody = r.groupSettings
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
