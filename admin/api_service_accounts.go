// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ServiceAccountsApi interface {

	/*
		AddProjectServiceAccount Assign One Service Account to One Project

		Assigns the specified Service Account to the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountRoleAssignment The Project permissions for the Service Account in the specified Project.
		@return AddProjectServiceAccountApiRequest
	*/
	AddProjectServiceAccount(ctx context.Context, clientId string, groupId string, groupServiceAccountRoleAssignment *GroupServiceAccountRoleAssignment) AddProjectServiceAccountApiRequest
	/*
		AddProjectServiceAccount Assign One Service Account to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddProjectServiceAccountApiParams - Parameters for the request
		@return AddProjectServiceAccountApiRequest
	*/
	AddProjectServiceAccountWithParams(ctx context.Context, args *AddProjectServiceAccountApiParams) AddProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	AddProjectServiceAccountExecute(r AddProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		CreateProjectServiceAccount Create One Project Service Account

		Creates one Service Account for the specified Project. The Service Account will automatically be added as an Organization Member to the Organization that the specified Project is a part of.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountRequest Details of the new Service Account.
		@return CreateProjectServiceAccountApiRequest
	*/
	CreateProjectServiceAccount(ctx context.Context, groupId string, groupServiceAccountRequest *GroupServiceAccountRequest) CreateProjectServiceAccountApiRequest
	/*
		CreateProjectServiceAccount Create One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectServiceAccountApiParams - Parameters for the request
		@return CreateProjectServiceAccountApiRequest
	*/
	CreateProjectServiceAccountWithParams(ctx context.Context, args *CreateProjectServiceAccountApiParams) CreateProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	CreateProjectServiceAccountExecute(r CreateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		CreateProjectServiceAccountAccessList Add Access List Entries for One Project Service Account

		Add Access List Entries for the specified Service Account for the project. Resources require all API requests to originate from IP addresses on the API access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clientId The Client ID of the Service Account.
		@param serviceAccountIPAccessListEntry A list of access list entries to add to the access list of the specified Service Account for the project.
		@return CreateProjectServiceAccountAccessListApiRequest
	*/
	CreateProjectServiceAccountAccessList(ctx context.Context, groupId string, clientId string, serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry) CreateProjectServiceAccountAccessListApiRequest
	/*
		CreateProjectServiceAccountAccessList Add Access List Entries for One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectServiceAccountAccessListApiParams - Parameters for the request
		@return CreateProjectServiceAccountAccessListApiRequest
	*/
	CreateProjectServiceAccountAccessListWithParams(ctx context.Context, args *CreateProjectServiceAccountAccessListApiParams) CreateProjectServiceAccountAccessListApiRequest

	// Method available only for mocking purposes
	CreateProjectServiceAccountAccessListExecute(r CreateProjectServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error)

	/*
		CreateProjectServiceAccountSecret Create One Project Service Account Secret

		Create a secret for the specified Service Account in the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clientId The Client ID of the Service Account.
		@param serviceAccountSecretRequest Details for the new secret.
		@return CreateProjectServiceAccountSecretApiRequest
	*/
	CreateProjectServiceAccountSecret(ctx context.Context, groupId string, clientId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateProjectServiceAccountSecretApiRequest
	/*
		CreateProjectServiceAccountSecret Create One Project Service Account Secret


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectServiceAccountSecretApiParams - Parameters for the request
		@return CreateProjectServiceAccountSecretApiRequest
	*/
	CreateProjectServiceAccountSecretWithParams(ctx context.Context, args *CreateProjectServiceAccountSecretApiParams) CreateProjectServiceAccountSecretApiRequest

	// Method available only for mocking purposes
	CreateProjectServiceAccountSecretExecute(r CreateProjectServiceAccountSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		CreateServiceAccount Create One Organization Service Account

		Creates one Service Account for the specified Organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param orgServiceAccountRequest Details of the new Service Account.
		@return CreateServiceAccountApiRequest
	*/
	CreateServiceAccount(ctx context.Context, orgId string, orgServiceAccountRequest *OrgServiceAccountRequest) CreateServiceAccountApiRequest
	/*
		CreateServiceAccount Create One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServiceAccountApiParams - Parameters for the request
		@return CreateServiceAccountApiRequest
	*/
	CreateServiceAccountWithParams(ctx context.Context, args *CreateServiceAccountApiParams) CreateServiceAccountApiRequest

	// Method available only for mocking purposes
	CreateServiceAccountExecute(r CreateServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error)

	/*
		CreateServiceAccountAccessList Add Access List Entries for One Organization Service Account

		Add Access List Entries for the specified Service Account for the organization. Resources require all API requests to originate from IP addresses on the API access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@param serviceAccountIPAccessListEntry A list of access list entries to add to the access list of the specified Service Account for the organization.
		@return CreateServiceAccountAccessListApiRequest
	*/
	CreateServiceAccountAccessList(ctx context.Context, orgId string, clientId string, serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry) CreateServiceAccountAccessListApiRequest
	/*
		CreateServiceAccountAccessList Add Access List Entries for One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServiceAccountAccessListApiParams - Parameters for the request
		@return CreateServiceAccountAccessListApiRequest
	*/
	CreateServiceAccountAccessListWithParams(ctx context.Context, args *CreateServiceAccountAccessListApiParams) CreateServiceAccountAccessListApiRequest

	// Method available only for mocking purposes
	CreateServiceAccountAccessListExecute(r CreateServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error)

	/*
		CreateServiceAccountSecret Create One Organization Service Account Secret

		Create a secret for the specified Service Account.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@param serviceAccountSecretRequest Details for the new secret.
		@return CreateServiceAccountSecretApiRequest
	*/
	CreateServiceAccountSecret(ctx context.Context, orgId string, clientId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateServiceAccountSecretApiRequest
	/*
		CreateServiceAccountSecret Create One Organization Service Account Secret


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServiceAccountSecretApiParams - Parameters for the request
		@return CreateServiceAccountSecretApiRequest
	*/
	CreateServiceAccountSecretWithParams(ctx context.Context, args *CreateServiceAccountSecretApiParams) CreateServiceAccountSecretApiRequest

	// Method available only for mocking purposes
	CreateServiceAccountSecretExecute(r CreateServiceAccountSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		DeleteProjectServiceAccount Remove One Project Service Account

		Removes the specified Service Account from the specified project. The Service Account will still be a part of the Organization it was created in, and the credentials will remain active until expired or manually revoked.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectServiceAccountApiRequest
	*/
	DeleteProjectServiceAccount(ctx context.Context, clientId string, groupId string) DeleteProjectServiceAccountApiRequest
	/*
		DeleteProjectServiceAccount Remove One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectServiceAccountApiParams - Parameters for the request
		@return DeleteProjectServiceAccountApiRequest
	*/
	DeleteProjectServiceAccountWithParams(ctx context.Context, args *DeleteProjectServiceAccountApiParams) DeleteProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	DeleteProjectServiceAccountExecute(r DeleteProjectServiceAccountApiRequest) (*http.Response, error)

	/*
		DeleteProjectServiceAccountAccessListEntry Remove One Access List Entry from One Project Service Account

		Removes the specified access list entry from the specified Service Account for the project. You can't remove the requesting IP address from the access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clientId The Client ID of the Service Account.
		@param ipAddress One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
		@return DeleteProjectServiceAccountAccessListEntryApiRequest
	*/
	DeleteProjectServiceAccountAccessListEntry(ctx context.Context, groupId string, clientId string, ipAddress string) DeleteProjectServiceAccountAccessListEntryApiRequest
	/*
		DeleteProjectServiceAccountAccessListEntry Remove One Access List Entry from One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectServiceAccountAccessListEntryApiParams - Parameters for the request
		@return DeleteProjectServiceAccountAccessListEntryApiRequest
	*/
	DeleteProjectServiceAccountAccessListEntryWithParams(ctx context.Context, args *DeleteProjectServiceAccountAccessListEntryApiParams) DeleteProjectServiceAccountAccessListEntryApiRequest

	// Method available only for mocking purposes
	DeleteProjectServiceAccountAccessListEntryExecute(r DeleteProjectServiceAccountAccessListEntryApiRequest) (*http.Response, error)

	/*
		DeleteProjectServiceAccountSecret Delete One Project Service Account Secret

		Deletes the specified Service Account secret.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectServiceAccountSecretApiRequest
	*/
	DeleteProjectServiceAccountSecret(ctx context.Context, clientId string, secretId string, groupId string) DeleteProjectServiceAccountSecretApiRequest
	/*
		DeleteProjectServiceAccountSecret Delete One Project Service Account Secret


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectServiceAccountSecretApiParams - Parameters for the request
		@return DeleteProjectServiceAccountSecretApiRequest
	*/
	DeleteProjectServiceAccountSecretWithParams(ctx context.Context, args *DeleteProjectServiceAccountSecretApiParams) DeleteProjectServiceAccountSecretApiRequest

	// Method available only for mocking purposes
	DeleteProjectServiceAccountSecretExecute(r DeleteProjectServiceAccountSecretApiRequest) (*http.Response, error)

	/*
		DeleteServiceAccount Delete One Organization Service Account

		Deletes the specified Service Account.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return DeleteServiceAccountApiRequest
	*/
	DeleteServiceAccount(ctx context.Context, clientId string, orgId string) DeleteServiceAccountApiRequest
	/*
		DeleteServiceAccount Delete One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteServiceAccountApiParams - Parameters for the request
		@return DeleteServiceAccountApiRequest
	*/
	DeleteServiceAccountWithParams(ctx context.Context, args *DeleteServiceAccountApiParams) DeleteServiceAccountApiRequest

	// Method available only for mocking purposes
	DeleteServiceAccountExecute(r DeleteServiceAccountApiRequest) (*http.Response, error)

	/*
		DeleteServiceAccountAccessListEntry Remove One Access List Entry from One Organization Service Account

		Removes the specified access list entry from the specified Service Account for the organization. You can't remove the requesting IP address from the access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@param ipAddress One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
		@return DeleteServiceAccountAccessListEntryApiRequest
	*/
	DeleteServiceAccountAccessListEntry(ctx context.Context, orgId string, clientId string, ipAddress string) DeleteServiceAccountAccessListEntryApiRequest
	/*
		DeleteServiceAccountAccessListEntry Remove One Access List Entry from One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteServiceAccountAccessListEntryApiParams - Parameters for the request
		@return DeleteServiceAccountAccessListEntryApiRequest
	*/
	DeleteServiceAccountAccessListEntryWithParams(ctx context.Context, args *DeleteServiceAccountAccessListEntryApiParams) DeleteServiceAccountAccessListEntryApiRequest

	// Method available only for mocking purposes
	DeleteServiceAccountAccessListEntryExecute(r DeleteServiceAccountAccessListEntryApiRequest) (*http.Response, error)

	/*
		DeleteServiceAccountSecret Delete One Organization Service Account Secret

		Deletes the specified Service Account secret.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return DeleteServiceAccountSecretApiRequest
	*/
	DeleteServiceAccountSecret(ctx context.Context, clientId string, secretId string, orgId string) DeleteServiceAccountSecretApiRequest
	/*
		DeleteServiceAccountSecret Delete One Organization Service Account Secret


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteServiceAccountSecretApiParams - Parameters for the request
		@return DeleteServiceAccountSecretApiRequest
	*/
	DeleteServiceAccountSecretWithParams(ctx context.Context, args *DeleteServiceAccountSecretApiParams) DeleteServiceAccountSecretApiRequest

	// Method available only for mocking purposes
	DeleteServiceAccountSecretExecute(r DeleteServiceAccountSecretApiRequest) (*http.Response, error)

	/*
		GetProjectServiceAccount Return One Project Service Account

		Returns one Service Account in the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clientId The Client ID of the Service Account.
		@return GetProjectServiceAccountApiRequest
	*/
	GetProjectServiceAccount(ctx context.Context, groupId string, clientId string) GetProjectServiceAccountApiRequest
	/*
		GetProjectServiceAccount Return One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectServiceAccountApiParams - Parameters for the request
		@return GetProjectServiceAccountApiRequest
	*/
	GetProjectServiceAccountWithParams(ctx context.Context, args *GetProjectServiceAccountApiParams) GetProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	GetProjectServiceAccountExecute(r GetProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		GetServiceAccount Return One Organization Service Account

		Returns the specified Service Account.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@return GetServiceAccountApiRequest
	*/
	GetServiceAccount(ctx context.Context, orgId string, clientId string) GetServiceAccountApiRequest
	/*
		GetServiceAccount Return One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServiceAccountApiParams - Parameters for the request
		@return GetServiceAccountApiRequest
	*/
	GetServiceAccountWithParams(ctx context.Context, args *GetServiceAccountApiParams) GetServiceAccountApiRequest

	// Method available only for mocking purposes
	GetServiceAccountExecute(r GetServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error)

	/*
		ListProjectServiceAccountAccessList Return All Access List Entries for One Project Service Account

		Returns all access list entries that you configured for the specified Service Account for the project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clientId The Client ID of the Service Account.
		@return ListProjectServiceAccountAccessListApiRequest
	*/
	ListProjectServiceAccountAccessList(ctx context.Context, groupId string, clientId string) ListProjectServiceAccountAccessListApiRequest
	/*
		ListProjectServiceAccountAccessList Return All Access List Entries for One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectServiceAccountAccessListApiParams - Parameters for the request
		@return ListProjectServiceAccountAccessListApiRequest
	*/
	ListProjectServiceAccountAccessListWithParams(ctx context.Context, args *ListProjectServiceAccountAccessListApiParams) ListProjectServiceAccountAccessListApiRequest

	// Method available only for mocking purposes
	ListProjectServiceAccountAccessListExecute(r ListProjectServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error)

	/*
		ListProjectServiceAccounts Return All Project Service Accounts

		Returns all Service Accounts for the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectServiceAccountsApiRequest
	*/
	ListProjectServiceAccounts(ctx context.Context, groupId string) ListProjectServiceAccountsApiRequest
	/*
		ListProjectServiceAccounts Return All Project Service Accounts


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectServiceAccountsApiParams - Parameters for the request
		@return ListProjectServiceAccountsApiRequest
	*/
	ListProjectServiceAccountsWithParams(ctx context.Context, args *ListProjectServiceAccountsApiParams) ListProjectServiceAccountsApiRequest

	// Method available only for mocking purposes
	ListProjectServiceAccountsExecute(r ListProjectServiceAccountsApiRequest) (*PaginatedGroupServiceAccounts, *http.Response, error)

	/*
		ListServiceAccountAccessList Return All Access List Entries for One Organization Service Account

		Returns all access list entries that you configured for the specified Service Account for the organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@return ListServiceAccountAccessListApiRequest
	*/
	ListServiceAccountAccessList(ctx context.Context, orgId string, clientId string) ListServiceAccountAccessListApiRequest
	/*
		ListServiceAccountAccessList Return All Access List Entries for One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServiceAccountAccessListApiParams - Parameters for the request
		@return ListServiceAccountAccessListApiRequest
	*/
	ListServiceAccountAccessListWithParams(ctx context.Context, args *ListServiceAccountAccessListApiParams) ListServiceAccountAccessListApiRequest

	// Method available only for mocking purposes
	ListServiceAccountAccessListExecute(r ListServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error)

	/*
		ListServiceAccountProjects Return All Service Account Project Assignments

		Returns a list of all projects the specified Service Account is a part of.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param clientId The Client ID of the Service Account.
		@return ListServiceAccountProjectsApiRequest
	*/
	ListServiceAccountProjects(ctx context.Context, orgId string, clientId string) ListServiceAccountProjectsApiRequest
	/*
		ListServiceAccountProjects Return All Service Account Project Assignments


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServiceAccountProjectsApiParams - Parameters for the request
		@return ListServiceAccountProjectsApiRequest
	*/
	ListServiceAccountProjectsWithParams(ctx context.Context, args *ListServiceAccountProjectsApiParams) ListServiceAccountProjectsApiRequest

	// Method available only for mocking purposes
	ListServiceAccountProjectsExecute(r ListServiceAccountProjectsApiRequest) (*PaginatedServiceAccountGroup, *http.Response, error)

	/*
		ListServiceAccounts Return All Organization Service Accounts

		Returns all Service Accounts for the specified Organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListServiceAccountsApiRequest
	*/
	ListServiceAccounts(ctx context.Context, orgId string) ListServiceAccountsApiRequest
	/*
		ListServiceAccounts Return All Organization Service Accounts


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServiceAccountsApiParams - Parameters for the request
		@return ListServiceAccountsApiRequest
	*/
	ListServiceAccountsWithParams(ctx context.Context, args *ListServiceAccountsApiParams) ListServiceAccountsApiRequest

	// Method available only for mocking purposes
	ListServiceAccountsExecute(r ListServiceAccountsApiRequest) (*PaginatedOrgServiceAccounts, *http.Response, error)

	/*
		UpdateProjectServiceAccount Update One Project Service Account

		Updates one Service Account in the specified Project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param groupServiceAccountUpdateRequest The new details for the Service Account.
		@return UpdateProjectServiceAccountApiRequest
	*/
	UpdateProjectServiceAccount(ctx context.Context, clientId string, groupId string, groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest) UpdateProjectServiceAccountApiRequest
	/*
		UpdateProjectServiceAccount Update One Project Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectServiceAccountApiParams - Parameters for the request
		@return UpdateProjectServiceAccountApiRequest
	*/
	UpdateProjectServiceAccountWithParams(ctx context.Context, args *UpdateProjectServiceAccountApiParams) UpdateProjectServiceAccountApiRequest

	// Method available only for mocking purposes
	UpdateProjectServiceAccountExecute(r UpdateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error)

	/*
		UpdateServiceAccount Update One Organization Service Account

		Updates the specified Service Account in the specified Organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId The Client ID of the Service Account.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param orgServiceAccountUpdateRequest The new details for the Service Account.
		@return UpdateServiceAccountApiRequest
	*/
	UpdateServiceAccount(ctx context.Context, clientId string, orgId string, orgServiceAccountUpdateRequest *OrgServiceAccountUpdateRequest) UpdateServiceAccountApiRequest
	/*
		UpdateServiceAccount Update One Organization Service Account


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateServiceAccountApiParams - Parameters for the request
		@return UpdateServiceAccountApiRequest
	*/
	UpdateServiceAccountWithParams(ctx context.Context, args *UpdateServiceAccountApiParams) UpdateServiceAccountApiRequest

	// Method available only for mocking purposes
	UpdateServiceAccountExecute(r UpdateServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error)
}

// ServiceAccountsApiService ServiceAccountsApi service
type ServiceAccountsApiService service

type AddProjectServiceAccountApiRequest struct {
	ctx                               context.Context
	ApiService                        ServiceAccountsApi
	clientId                          string
	groupId                           string
	groupServiceAccountRoleAssignment *GroupServiceAccountRoleAssignment
}

type AddProjectServiceAccountApiParams struct {
	ClientId                          string
	GroupId                           string
	GroupServiceAccountRoleAssignment *GroupServiceAccountRoleAssignment
}

func (a *ServiceAccountsApiService) AddProjectServiceAccountWithParams(ctx context.Context, args *AddProjectServiceAccountApiParams) AddProjectServiceAccountApiRequest {
	return AddProjectServiceAccountApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		clientId:                          args.ClientId,
		groupId:                           args.GroupId,
		groupServiceAccountRoleAssignment: args.GroupServiceAccountRoleAssignment,
	}
}

func (r AddProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.AddProjectServiceAccountExecute(r)
}

/*
AddProjectServiceAccount Assign One Service Account to One Project

Assigns the specified Service Account to the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AddProjectServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) AddProjectServiceAccount(ctx context.Context, clientId string, groupId string, groupServiceAccountRoleAssignment *GroupServiceAccountRoleAssignment) AddProjectServiceAccountApiRequest {
	return AddProjectServiceAccountApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		clientId:                          clientId,
		groupId:                           groupId,
		groupServiceAccountRoleAssignment: groupServiceAccountRoleAssignment,
	}
}

// AddProjectServiceAccountExecute executes the request
//
//	@return GroupServiceAccount
func (a *ServiceAccountsApiService) AddProjectServiceAccountExecute(r AddProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.AddProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}:invite"
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountRoleAssignment == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountRoleAssignment is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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
	ApiService                 ServiceAccountsApi
	groupId                    string
	groupServiceAccountRequest *GroupServiceAccountRequest
}

type CreateProjectServiceAccountApiParams struct {
	GroupId                    string
	GroupServiceAccountRequest *GroupServiceAccountRequest
}

func (a *ServiceAccountsApiService) CreateProjectServiceAccountWithParams(ctx context.Context, args *CreateProjectServiceAccountApiParams) CreateProjectServiceAccountApiRequest {
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
CreateProjectServiceAccount Create One Project Service Account

Creates one Service Account for the specified Project. The Service Account will automatically be added as an Organization Member to the Organization that the specified Project is a part of.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) CreateProjectServiceAccount(ctx context.Context, groupId string, groupServiceAccountRequest *GroupServiceAccountRequest) CreateProjectServiceAccountApiRequest {
	return CreateProjectServiceAccountApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		groupServiceAccountRequest: groupServiceAccountRequest,
	}
}

// CreateProjectServiceAccountExecute executes the request
//
//	@return GroupServiceAccount
func (a *ServiceAccountsApiService) CreateProjectServiceAccountExecute(r CreateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountRequest == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type CreateProjectServiceAccountAccessListApiRequest struct {
	ctx                             context.Context
	ApiService                      ServiceAccountsApi
	groupId                         string
	clientId                        string
	serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry
	includeCount                    *bool
	itemsPerPage                    *int
	pageNum                         *int
}

type CreateProjectServiceAccountAccessListApiParams struct {
	GroupId                         string
	ClientId                        string
	ServiceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry
	IncludeCount                    *bool
	ItemsPerPage                    *int
	PageNum                         *int
}

func (a *ServiceAccountsApiService) CreateProjectServiceAccountAccessListWithParams(ctx context.Context, args *CreateProjectServiceAccountAccessListApiParams) CreateProjectServiceAccountAccessListApiRequest {
	return CreateProjectServiceAccountAccessListApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         args.GroupId,
		clientId:                        args.ClientId,
		serviceAccountIPAccessListEntry: args.ServiceAccountIPAccessListEntry,
		includeCount:                    args.IncludeCount,
		itemsPerPage:                    args.ItemsPerPage,
		pageNum:                         args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r CreateProjectServiceAccountAccessListApiRequest) IncludeCount(includeCount bool) CreateProjectServiceAccountAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r CreateProjectServiceAccountAccessListApiRequest) ItemsPerPage(itemsPerPage int) CreateProjectServiceAccountAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r CreateProjectServiceAccountAccessListApiRequest) PageNum(pageNum int) CreateProjectServiceAccountAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r CreateProjectServiceAccountAccessListApiRequest) Execute() (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	return r.ApiService.CreateProjectServiceAccountAccessListExecute(r)
}

/*
CreateProjectServiceAccountAccessList Add Access List Entries for One Project Service Account

Add Access List Entries for the specified Service Account for the project. Resources require all API requests to originate from IP addresses on the API access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clientId The Client ID of the Service Account.
	@return CreateProjectServiceAccountAccessListApiRequest
*/
func (a *ServiceAccountsApiService) CreateProjectServiceAccountAccessList(ctx context.Context, groupId string, clientId string, serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry) CreateProjectServiceAccountAccessListApiRequest {
	return CreateProjectServiceAccountAccessListApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         groupId,
		clientId:                        clientId,
		serviceAccountIPAccessListEntry: serviceAccountIPAccessListEntry,
	}
}

// CreateProjectServiceAccountAccessListExecute executes the request
//
//	@return PaginatedServiceAccountIPAccessEntry
func (a *ServiceAccountsApiService) CreateProjectServiceAccountAccessListExecute(r CreateProjectServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedServiceAccountIPAccessEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateProjectServiceAccountAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountIPAccessListEntry == nil {
		return localVarReturnValue, nil, reportError("serviceAccountIPAccessListEntry is required and must be specified")
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
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serviceAccountIPAccessListEntry
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

type CreateProjectServiceAccountSecretApiRequest struct {
	ctx                         context.Context
	ApiService                  ServiceAccountsApi
	groupId                     string
	clientId                    string
	serviceAccountSecretRequest *ServiceAccountSecretRequest
}

type CreateProjectServiceAccountSecretApiParams struct {
	GroupId                     string
	ClientId                    string
	ServiceAccountSecretRequest *ServiceAccountSecretRequest
}

func (a *ServiceAccountsApiService) CreateProjectServiceAccountSecretWithParams(ctx context.Context, args *CreateProjectServiceAccountSecretApiParams) CreateProjectServiceAccountSecretApiRequest {
	return CreateProjectServiceAccountSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		clientId:                    args.ClientId,
		serviceAccountSecretRequest: args.ServiceAccountSecretRequest,
	}
}

func (r CreateProjectServiceAccountSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.CreateProjectServiceAccountSecretExecute(r)
}

/*
CreateProjectServiceAccountSecret Create One Project Service Account Secret

Create a secret for the specified Service Account in the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clientId The Client ID of the Service Account.
	@return CreateProjectServiceAccountSecretApiRequest
*/
func (a *ServiceAccountsApiService) CreateProjectServiceAccountSecret(ctx context.Context, groupId string, clientId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateProjectServiceAccountSecretApiRequest {
	return CreateProjectServiceAccountSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		clientId:                    clientId,
		serviceAccountSecretRequest: serviceAccountSecretRequest,
	}
}

// CreateProjectServiceAccountSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *ServiceAccountsApiService) CreateProjectServiceAccountSecretExecute(r CreateProjectServiceAccountSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateProjectServiceAccountSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountSecretRequest == nil {
		return localVarReturnValue, nil, reportError("serviceAccountSecretRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serviceAccountSecretRequest
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

type CreateServiceAccountApiRequest struct {
	ctx                      context.Context
	ApiService               ServiceAccountsApi
	orgId                    string
	orgServiceAccountRequest *OrgServiceAccountRequest
}

type CreateServiceAccountApiParams struct {
	OrgId                    string
	OrgServiceAccountRequest *OrgServiceAccountRequest
}

func (a *ServiceAccountsApiService) CreateServiceAccountWithParams(ctx context.Context, args *CreateServiceAccountApiParams) CreateServiceAccountApiRequest {
	return CreateServiceAccountApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		orgId:                    args.OrgId,
		orgServiceAccountRequest: args.OrgServiceAccountRequest,
	}
}

func (r CreateServiceAccountApiRequest) Execute() (*OrgServiceAccount, *http.Response, error) {
	return r.ApiService.CreateServiceAccountExecute(r)
}

/*
CreateServiceAccount Create One Organization Service Account

Creates one Service Account for the specified Organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) CreateServiceAccount(ctx context.Context, orgId string, orgServiceAccountRequest *OrgServiceAccountRequest) CreateServiceAccountApiRequest {
	return CreateServiceAccountApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		orgId:                    orgId,
		orgServiceAccountRequest: orgServiceAccountRequest,
	}
}

// CreateServiceAccountExecute executes the request
//
//	@return OrgServiceAccount
func (a *ServiceAccountsApiService) CreateServiceAccountExecute(r CreateServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgServiceAccountRequest == nil {
		return localVarReturnValue, nil, reportError("orgServiceAccountRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orgServiceAccountRequest
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

type CreateServiceAccountAccessListApiRequest struct {
	ctx                             context.Context
	ApiService                      ServiceAccountsApi
	orgId                           string
	clientId                        string
	serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry
	includeCount                    *bool
	itemsPerPage                    *int
	pageNum                         *int
}

type CreateServiceAccountAccessListApiParams struct {
	OrgId                           string
	ClientId                        string
	ServiceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry
	IncludeCount                    *bool
	ItemsPerPage                    *int
	PageNum                         *int
}

func (a *ServiceAccountsApiService) CreateServiceAccountAccessListWithParams(ctx context.Context, args *CreateServiceAccountAccessListApiParams) CreateServiceAccountAccessListApiRequest {
	return CreateServiceAccountAccessListApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		orgId:                           args.OrgId,
		clientId:                        args.ClientId,
		serviceAccountIPAccessListEntry: args.ServiceAccountIPAccessListEntry,
		includeCount:                    args.IncludeCount,
		itemsPerPage:                    args.ItemsPerPage,
		pageNum:                         args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r CreateServiceAccountAccessListApiRequest) IncludeCount(includeCount bool) CreateServiceAccountAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r CreateServiceAccountAccessListApiRequest) ItemsPerPage(itemsPerPage int) CreateServiceAccountAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r CreateServiceAccountAccessListApiRequest) PageNum(pageNum int) CreateServiceAccountAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r CreateServiceAccountAccessListApiRequest) Execute() (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	return r.ApiService.CreateServiceAccountAccessListExecute(r)
}

/*
CreateServiceAccountAccessList Add Access List Entries for One Organization Service Account

Add Access List Entries for the specified Service Account for the organization. Resources require all API requests to originate from IP addresses on the API access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@return CreateServiceAccountAccessListApiRequest
*/
func (a *ServiceAccountsApiService) CreateServiceAccountAccessList(ctx context.Context, orgId string, clientId string, serviceAccountIPAccessListEntry *[]ServiceAccountIPAccessListEntry) CreateServiceAccountAccessListApiRequest {
	return CreateServiceAccountAccessListApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		orgId:                           orgId,
		clientId:                        clientId,
		serviceAccountIPAccessListEntry: serviceAccountIPAccessListEntry,
	}
}

// CreateServiceAccountAccessListExecute executes the request
//
//	@return PaginatedServiceAccountIPAccessEntry
func (a *ServiceAccountsApiService) CreateServiceAccountAccessListExecute(r CreateServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedServiceAccountIPAccessEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateServiceAccountAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountIPAccessListEntry == nil {
		return localVarReturnValue, nil, reportError("serviceAccountIPAccessListEntry is required and must be specified")
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
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serviceAccountIPAccessListEntry
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

type CreateServiceAccountSecretApiRequest struct {
	ctx                         context.Context
	ApiService                  ServiceAccountsApi
	orgId                       string
	clientId                    string
	serviceAccountSecretRequest *ServiceAccountSecretRequest
}

type CreateServiceAccountSecretApiParams struct {
	OrgId                       string
	ClientId                    string
	ServiceAccountSecretRequest *ServiceAccountSecretRequest
}

func (a *ServiceAccountsApiService) CreateServiceAccountSecretWithParams(ctx context.Context, args *CreateServiceAccountSecretApiParams) CreateServiceAccountSecretApiRequest {
	return CreateServiceAccountSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		orgId:                       args.OrgId,
		clientId:                    args.ClientId,
		serviceAccountSecretRequest: args.ServiceAccountSecretRequest,
	}
}

func (r CreateServiceAccountSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.CreateServiceAccountSecretExecute(r)
}

/*
CreateServiceAccountSecret Create One Organization Service Account Secret

Create a secret for the specified Service Account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@return CreateServiceAccountSecretApiRequest
*/
func (a *ServiceAccountsApiService) CreateServiceAccountSecret(ctx context.Context, orgId string, clientId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateServiceAccountSecretApiRequest {
	return CreateServiceAccountSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		orgId:                       orgId,
		clientId:                    clientId,
		serviceAccountSecretRequest: serviceAccountSecretRequest,
	}
}

// CreateServiceAccountSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *ServiceAccountsApiService) CreateServiceAccountSecretExecute(r CreateServiceAccountSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateServiceAccountSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountSecretRequest == nil {
		return localVarReturnValue, nil, reportError("serviceAccountSecretRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serviceAccountSecretRequest
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
	ctx        context.Context
	ApiService ServiceAccountsApi
	clientId   string
	groupId    string
}

type DeleteProjectServiceAccountApiParams struct {
	ClientId string
	GroupId  string
}

func (a *ServiceAccountsApiService) DeleteProjectServiceAccountWithParams(ctx context.Context, args *DeleteProjectServiceAccountApiParams) DeleteProjectServiceAccountApiRequest {
	return DeleteProjectServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   args.ClientId,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectServiceAccountApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectServiceAccountExecute(r)
}

/*
DeleteProjectServiceAccount Remove One Project Service Account

Removes the specified Service Account from the specified project. The Service Account will still be a part of the Organization it was created in, and the credentials will remain active until expired or manually revoked.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) DeleteProjectServiceAccount(ctx context.Context, clientId string, groupId string) DeleteProjectServiceAccountApiRequest {
	return DeleteProjectServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   clientId,
		groupId:    groupId,
	}
}

// DeleteProjectServiceAccountExecute executes the request
func (a *ServiceAccountsApiService) DeleteProjectServiceAccountExecute(r DeleteProjectServiceAccountApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteProjectServiceAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}"
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type DeleteProjectServiceAccountAccessListEntryApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	groupId    string
	clientId   string
	ipAddress  string
}

type DeleteProjectServiceAccountAccessListEntryApiParams struct {
	GroupId   string
	ClientId  string
	IpAddress string
}

func (a *ServiceAccountsApiService) DeleteProjectServiceAccountAccessListEntryWithParams(ctx context.Context, args *DeleteProjectServiceAccountAccessListEntryApiParams) DeleteProjectServiceAccountAccessListEntryApiRequest {
	return DeleteProjectServiceAccountAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		clientId:   args.ClientId,
		ipAddress:  args.IpAddress,
	}
}

func (r DeleteProjectServiceAccountAccessListEntryApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectServiceAccountAccessListEntryExecute(r)
}

/*
DeleteProjectServiceAccountAccessListEntry Remove One Access List Entry from One Project Service Account

Removes the specified access list entry from the specified Service Account for the project. You can't remove the requesting IP address from the access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clientId The Client ID of the Service Account.
	@param ipAddress One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
	@return DeleteProjectServiceAccountAccessListEntryApiRequest
*/
func (a *ServiceAccountsApiService) DeleteProjectServiceAccountAccessListEntry(ctx context.Context, groupId string, clientId string, ipAddress string) DeleteProjectServiceAccountAccessListEntryApiRequest {
	return DeleteProjectServiceAccountAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		clientId:   clientId,
		ipAddress:  ipAddress,
	}
}

// DeleteProjectServiceAccountAccessListEntryExecute executes the request
func (a *ServiceAccountsApiService) DeleteProjectServiceAccountAccessListEntryExecute(r DeleteProjectServiceAccountAccessListEntryApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteProjectServiceAccountAccessListEntry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList/{ipAddress}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type DeleteProjectServiceAccountSecretApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	clientId   string
	secretId   string
	groupId    string
}

type DeleteProjectServiceAccountSecretApiParams struct {
	ClientId string
	SecretId string
	GroupId  string
}

func (a *ServiceAccountsApiService) DeleteProjectServiceAccountSecretWithParams(ctx context.Context, args *DeleteProjectServiceAccountSecretApiParams) DeleteProjectServiceAccountSecretApiRequest {
	return DeleteProjectServiceAccountSecretApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   args.ClientId,
		secretId:   args.SecretId,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectServiceAccountSecretApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectServiceAccountSecretExecute(r)
}

/*
DeleteProjectServiceAccountSecret Delete One Project Service Account Secret

Deletes the specified Service Account secret.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectServiceAccountSecretApiRequest
*/
func (a *ServiceAccountsApiService) DeleteProjectServiceAccountSecret(ctx context.Context, clientId string, secretId string, groupId string) DeleteProjectServiceAccountSecretApiRequest {
	return DeleteProjectServiceAccountSecretApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   clientId,
		secretId:   secretId,
		groupId:    groupId,
	}
}

// DeleteProjectServiceAccountSecretExecute executes the request
func (a *ServiceAccountsApiService) DeleteProjectServiceAccountSecretExecute(r DeleteProjectServiceAccountSecretApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteProjectServiceAccountSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/secrets/{secretId}"
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
	if r.secretId == "" {
		return nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type DeleteServiceAccountApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	clientId   string
	orgId      string
}

type DeleteServiceAccountApiParams struct {
	ClientId string
	OrgId    string
}

func (a *ServiceAccountsApiService) DeleteServiceAccountWithParams(ctx context.Context, args *DeleteServiceAccountApiParams) DeleteServiceAccountApiRequest {
	return DeleteServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   args.ClientId,
		orgId:      args.OrgId,
	}
}

func (r DeleteServiceAccountApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceAccountExecute(r)
}

/*
DeleteServiceAccount Delete One Organization Service Account

Deletes the specified Service Account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return DeleteServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) DeleteServiceAccount(ctx context.Context, clientId string, orgId string) DeleteServiceAccountApiRequest {
	return DeleteServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   clientId,
		orgId:      orgId,
	}
}

// DeleteServiceAccountExecute executes the request
func (a *ServiceAccountsApiService) DeleteServiceAccountExecute(r DeleteServiceAccountApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteServiceAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}"
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type DeleteServiceAccountAccessListEntryApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	orgId      string
	clientId   string
	ipAddress  string
}

type DeleteServiceAccountAccessListEntryApiParams struct {
	OrgId     string
	ClientId  string
	IpAddress string
}

func (a *ServiceAccountsApiService) DeleteServiceAccountAccessListEntryWithParams(ctx context.Context, args *DeleteServiceAccountAccessListEntryApiParams) DeleteServiceAccountAccessListEntryApiRequest {
	return DeleteServiceAccountAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		clientId:   args.ClientId,
		ipAddress:  args.IpAddress,
	}
}

func (r DeleteServiceAccountAccessListEntryApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceAccountAccessListEntryExecute(r)
}

/*
DeleteServiceAccountAccessListEntry Remove One Access List Entry from One Organization Service Account

Removes the specified access list entry from the specified Service Account for the organization. You can't remove the requesting IP address from the access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@param ipAddress One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
	@return DeleteServiceAccountAccessListEntryApiRequest
*/
func (a *ServiceAccountsApiService) DeleteServiceAccountAccessListEntry(ctx context.Context, orgId string, clientId string, ipAddress string) DeleteServiceAccountAccessListEntryApiRequest {
	return DeleteServiceAccountAccessListEntryApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		clientId:   clientId,
		ipAddress:  ipAddress,
	}
}

// DeleteServiceAccountAccessListEntryExecute executes the request
func (a *ServiceAccountsApiService) DeleteServiceAccountAccessListEntryExecute(r DeleteServiceAccountAccessListEntryApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteServiceAccountAccessListEntry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList/{ipAddress}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type DeleteServiceAccountSecretApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	clientId   string
	secretId   string
	orgId      string
}

type DeleteServiceAccountSecretApiParams struct {
	ClientId string
	SecretId string
	OrgId    string
}

func (a *ServiceAccountsApiService) DeleteServiceAccountSecretWithParams(ctx context.Context, args *DeleteServiceAccountSecretApiParams) DeleteServiceAccountSecretApiRequest {
	return DeleteServiceAccountSecretApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   args.ClientId,
		secretId:   args.SecretId,
		orgId:      args.OrgId,
	}
}

func (r DeleteServiceAccountSecretApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceAccountSecretExecute(r)
}

/*
DeleteServiceAccountSecret Delete One Organization Service Account Secret

Deletes the specified Service Account secret.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return DeleteServiceAccountSecretApiRequest
*/
func (a *ServiceAccountsApiService) DeleteServiceAccountSecret(ctx context.Context, clientId string, secretId string, orgId string) DeleteServiceAccountSecretApiRequest {
	return DeleteServiceAccountSecretApiRequest{
		ApiService: a,
		ctx:        ctx,
		clientId:   clientId,
		secretId:   secretId,
		orgId:      orgId,
	}
}

// DeleteServiceAccountSecretExecute executes the request
func (a *ServiceAccountsApiService) DeleteServiceAccountSecretExecute(r DeleteServiceAccountSecretApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteServiceAccountSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/secrets/{secretId}"
	if r.clientId == "" {
		return nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
	if r.secretId == "" {
		return nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type GetProjectServiceAccountApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	groupId    string
	clientId   string
}

type GetProjectServiceAccountApiParams struct {
	GroupId  string
	ClientId string
}

func (a *ServiceAccountsApiService) GetProjectServiceAccountWithParams(ctx context.Context, args *GetProjectServiceAccountApiParams) GetProjectServiceAccountApiRequest {
	return GetProjectServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		clientId:   args.ClientId,
	}
}

func (r GetProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.GetProjectServiceAccountExecute(r)
}

/*
GetProjectServiceAccount Return One Project Service Account

Returns one Service Account in the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clientId The Client ID of the Service Account.
	@return GetProjectServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) GetProjectServiceAccount(ctx context.Context, groupId string, clientId string) GetProjectServiceAccountApiRequest {
	return GetProjectServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		clientId:   clientId,
	}
}

// GetProjectServiceAccountExecute executes the request
//
//	@return GroupServiceAccount
func (a *ServiceAccountsApiService) GetProjectServiceAccountExecute(r GetProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.GetProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type GetServiceAccountApiRequest struct {
	ctx        context.Context
	ApiService ServiceAccountsApi
	orgId      string
	clientId   string
}

type GetServiceAccountApiParams struct {
	OrgId    string
	ClientId string
}

func (a *ServiceAccountsApiService) GetServiceAccountWithParams(ctx context.Context, args *GetServiceAccountApiParams) GetServiceAccountApiRequest {
	return GetServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		clientId:   args.ClientId,
	}
}

func (r GetServiceAccountApiRequest) Execute() (*OrgServiceAccount, *http.Response, error) {
	return r.ApiService.GetServiceAccountExecute(r)
}

/*
GetServiceAccount Return One Organization Service Account

Returns the specified Service Account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@return GetServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) GetServiceAccount(ctx context.Context, orgId string, clientId string) GetServiceAccountApiRequest {
	return GetServiceAccountApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		clientId:   clientId,
	}
}

// GetServiceAccountExecute executes the request
//
//	@return OrgServiceAccount
func (a *ServiceAccountsApiService) GetServiceAccountExecute(r GetServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.GetServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type ListProjectServiceAccountAccessListApiRequest struct {
	ctx          context.Context
	ApiService   ServiceAccountsApi
	groupId      string
	clientId     string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectServiceAccountAccessListApiParams struct {
	GroupId      string
	ClientId     string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServiceAccountsApiService) ListProjectServiceAccountAccessListWithParams(ctx context.Context, args *ListProjectServiceAccountAccessListApiParams) ListProjectServiceAccountAccessListApiRequest {
	return ListProjectServiceAccountAccessListApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clientId:     args.ClientId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectServiceAccountAccessListApiRequest) IncludeCount(includeCount bool) ListProjectServiceAccountAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectServiceAccountAccessListApiRequest) ItemsPerPage(itemsPerPage int) ListProjectServiceAccountAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectServiceAccountAccessListApiRequest) PageNum(pageNum int) ListProjectServiceAccountAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectServiceAccountAccessListApiRequest) Execute() (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	return r.ApiService.ListProjectServiceAccountAccessListExecute(r)
}

/*
ListProjectServiceAccountAccessList Return All Access List Entries for One Project Service Account

Returns all access list entries that you configured for the specified Service Account for the project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clientId The Client ID of the Service Account.
	@return ListProjectServiceAccountAccessListApiRequest
*/
func (a *ServiceAccountsApiService) ListProjectServiceAccountAccessList(ctx context.Context, groupId string, clientId string) ListProjectServiceAccountAccessListApiRequest {
	return ListProjectServiceAccountAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		clientId:   clientId,
	}
}

// ListProjectServiceAccountAccessListExecute executes the request
//
//	@return PaginatedServiceAccountIPAccessEntry
func (a *ServiceAccountsApiService) ListProjectServiceAccountAccessListExecute(r ListProjectServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedServiceAccountIPAccessEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListProjectServiceAccountAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}/accessList"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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
	ApiService   ServiceAccountsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListProjectServiceAccountsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServiceAccountsApiService) ListProjectServiceAccountsWithParams(ctx context.Context, args *ListProjectServiceAccountsApiParams) ListProjectServiceAccountsApiRequest {
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
ListProjectServiceAccounts Return All Project Service Accounts

Returns all Service Accounts for the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectServiceAccountsApiRequest
*/
func (a *ServiceAccountsApiService) ListProjectServiceAccounts(ctx context.Context, groupId string) ListProjectServiceAccountsApiRequest {
	return ListProjectServiceAccountsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectServiceAccountsExecute executes the request
//
//	@return PaginatedGroupServiceAccounts
func (a *ServiceAccountsApiService) ListProjectServiceAccountsExecute(r ListProjectServiceAccountsApiRequest) (*PaginatedGroupServiceAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedGroupServiceAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListProjectServiceAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type ListServiceAccountAccessListApiRequest struct {
	ctx          context.Context
	ApiService   ServiceAccountsApi
	orgId        string
	clientId     string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListServiceAccountAccessListApiParams struct {
	OrgId        string
	ClientId     string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServiceAccountsApiService) ListServiceAccountAccessListWithParams(ctx context.Context, args *ListServiceAccountAccessListApiParams) ListServiceAccountAccessListApiRequest {
	return ListServiceAccountAccessListApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		clientId:     args.ClientId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListServiceAccountAccessListApiRequest) IncludeCount(includeCount bool) ListServiceAccountAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListServiceAccountAccessListApiRequest) ItemsPerPage(itemsPerPage int) ListServiceAccountAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServiceAccountAccessListApiRequest) PageNum(pageNum int) ListServiceAccountAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServiceAccountAccessListApiRequest) Execute() (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	return r.ApiService.ListServiceAccountAccessListExecute(r)
}

/*
ListServiceAccountAccessList Return All Access List Entries for One Organization Service Account

Returns all access list entries that you configured for the specified Service Account for the organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@return ListServiceAccountAccessListApiRequest
*/
func (a *ServiceAccountsApiService) ListServiceAccountAccessList(ctx context.Context, orgId string, clientId string) ListServiceAccountAccessListApiRequest {
	return ListServiceAccountAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		clientId:   clientId,
	}
}

// ListServiceAccountAccessListExecute executes the request
//
//	@return PaginatedServiceAccountIPAccessEntry
func (a *ServiceAccountsApiService) ListServiceAccountAccessListExecute(r ListServiceAccountAccessListApiRequest) (*PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedServiceAccountIPAccessEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListServiceAccountAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/accessList"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type ListServiceAccountProjectsApiRequest struct {
	ctx          context.Context
	ApiService   ServiceAccountsApi
	orgId        string
	clientId     string
	itemsPerPage *int
	pageNum      *int
}

type ListServiceAccountProjectsApiParams struct {
	OrgId        string
	ClientId     string
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServiceAccountsApiService) ListServiceAccountProjectsWithParams(ctx context.Context, args *ListServiceAccountProjectsApiParams) ListServiceAccountProjectsApiRequest {
	return ListServiceAccountProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		clientId:     args.ClientId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListServiceAccountProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListServiceAccountProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServiceAccountProjectsApiRequest) PageNum(pageNum int) ListServiceAccountProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServiceAccountProjectsApiRequest) Execute() (*PaginatedServiceAccountGroup, *http.Response, error) {
	return r.ApiService.ListServiceAccountProjectsExecute(r)
}

/*
ListServiceAccountProjects Return All Service Account Project Assignments

Returns a list of all projects the specified Service Account is a part of.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param clientId The Client ID of the Service Account.
	@return ListServiceAccountProjectsApiRequest
*/
func (a *ServiceAccountsApiService) ListServiceAccountProjects(ctx context.Context, orgId string, clientId string) ListServiceAccountProjectsApiRequest {
	return ListServiceAccountProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		clientId:   clientId,
	}
}

// ListServiceAccountProjectsExecute executes the request
//
//	@return PaginatedServiceAccountGroup
func (a *ServiceAccountsApiService) ListServiceAccountProjectsExecute(r ListServiceAccountProjectsApiRequest) (*PaginatedServiceAccountGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedServiceAccountGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListServiceAccountProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}/groups"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type ListServiceAccountsApiRequest struct {
	ctx          context.Context
	ApiService   ServiceAccountsApi
	orgId        string
	itemsPerPage *int
	pageNum      *int
}

type ListServiceAccountsApiParams struct {
	OrgId        string
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServiceAccountsApiService) ListServiceAccountsWithParams(ctx context.Context, args *ListServiceAccountsApiParams) ListServiceAccountsApiRequest {
	return ListServiceAccountsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListServiceAccountsApiRequest) ItemsPerPage(itemsPerPage int) ListServiceAccountsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServiceAccountsApiRequest) PageNum(pageNum int) ListServiceAccountsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServiceAccountsApiRequest) Execute() (*PaginatedOrgServiceAccounts, *http.Response, error) {
	return r.ApiService.ListServiceAccountsExecute(r)
}

/*
ListServiceAccounts Return All Organization Service Accounts

Returns all Service Accounts for the specified Organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListServiceAccountsApiRequest
*/
func (a *ServiceAccountsApiService) ListServiceAccounts(ctx context.Context, orgId string) ListServiceAccountsApiRequest {
	return ListServiceAccountsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListServiceAccountsExecute executes the request
//
//	@return PaginatedOrgServiceAccounts
func (a *ServiceAccountsApiService) ListServiceAccountsExecute(r ListServiceAccountsApiRequest) (*PaginatedOrgServiceAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedOrgServiceAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListServiceAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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
	ApiService                       ServiceAccountsApi
	clientId                         string
	groupId                          string
	groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest
}

type UpdateProjectServiceAccountApiParams struct {
	ClientId                         string
	GroupId                          string
	GroupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest
}

func (a *ServiceAccountsApiService) UpdateProjectServiceAccountWithParams(ctx context.Context, args *UpdateProjectServiceAccountApiParams) UpdateProjectServiceAccountApiRequest {
	return UpdateProjectServiceAccountApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		clientId:                         args.ClientId,
		groupId:                          args.GroupId,
		groupServiceAccountUpdateRequest: args.GroupServiceAccountUpdateRequest,
	}
}

func (r UpdateProjectServiceAccountApiRequest) Execute() (*GroupServiceAccount, *http.Response, error) {
	return r.ApiService.UpdateProjectServiceAccountExecute(r)
}

/*
UpdateProjectServiceAccount Update One Project Service Account

Updates one Service Account in the specified Project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) UpdateProjectServiceAccount(ctx context.Context, clientId string, groupId string, groupServiceAccountUpdateRequest *GroupServiceAccountUpdateRequest) UpdateProjectServiceAccountApiRequest {
	return UpdateProjectServiceAccountApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		clientId:                         clientId,
		groupId:                          groupId,
		groupServiceAccountUpdateRequest: groupServiceAccountUpdateRequest,
	}
}

// UpdateProjectServiceAccountExecute executes the request
//
//	@return GroupServiceAccount
func (a *ServiceAccountsApiService) UpdateProjectServiceAccountExecute(r UpdateProjectServiceAccountApiRequest) (*GroupServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.UpdateProjectServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serviceAccounts/{clientId}"
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupServiceAccountUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("groupServiceAccountUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type UpdateServiceAccountApiRequest struct {
	ctx                            context.Context
	ApiService                     ServiceAccountsApi
	clientId                       string
	orgId                          string
	orgServiceAccountUpdateRequest *OrgServiceAccountUpdateRequest
}

type UpdateServiceAccountApiParams struct {
	ClientId                       string
	OrgId                          string
	OrgServiceAccountUpdateRequest *OrgServiceAccountUpdateRequest
}

func (a *ServiceAccountsApiService) UpdateServiceAccountWithParams(ctx context.Context, args *UpdateServiceAccountApiParams) UpdateServiceAccountApiRequest {
	return UpdateServiceAccountApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		clientId:                       args.ClientId,
		orgId:                          args.OrgId,
		orgServiceAccountUpdateRequest: args.OrgServiceAccountUpdateRequest,
	}
}

func (r UpdateServiceAccountApiRequest) Execute() (*OrgServiceAccount, *http.Response, error) {
	return r.ApiService.UpdateServiceAccountExecute(r)
}

/*
UpdateServiceAccount Update One Organization Service Account

Updates the specified Service Account in the specified Organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The Client ID of the Service Account.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateServiceAccountApiRequest
*/
func (a *ServiceAccountsApiService) UpdateServiceAccount(ctx context.Context, clientId string, orgId string, orgServiceAccountUpdateRequest *OrgServiceAccountUpdateRequest) UpdateServiceAccountApiRequest {
	return UpdateServiceAccountApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		clientId:                       clientId,
		orgId:                          orgId,
		orgServiceAccountUpdateRequest: orgServiceAccountUpdateRequest,
	}
}

// UpdateServiceAccountExecute executes the request
//
//	@return OrgServiceAccount
func (a *ServiceAccountsApiService) UpdateServiceAccountExecute(r UpdateServiceAccountApiRequest) (*OrgServiceAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgServiceAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.UpdateServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/serviceAccounts/{clientId}"
	if r.clientId == "" {
		return localVarReturnValue, nil, reportError("clientId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(r.clientId), -1)
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgServiceAccountUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("orgServiceAccountUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orgServiceAccountUpdateRequest
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
