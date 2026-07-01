// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RemoteMCPConfigurationsApi interface {

	/*
		CreateGroupMcpConfig Create One MCP Configuration for One Project

		Creates an MCP configuration for the specified project. Returns the configuration ID and ingress credentials.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param createGroupMcpConfigRequest MCP configuration to create.
		@return CreateGroupMcpConfigApiRequest
	*/
	CreateGroupMcpConfig(ctx context.Context, groupId string, createGroupMcpConfigRequest *CreateGroupMcpConfigRequest) CreateGroupMcpConfigApiRequest
	/*
		CreateGroupMcpConfig Create One MCP Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupMcpConfigApiParams - Parameters for the request
		@return CreateGroupMcpConfigApiRequest
	*/
	CreateGroupMcpConfigWithParams(ctx context.Context, args *CreateGroupMcpConfigApiParams) CreateGroupMcpConfigApiRequest

	// Method available only for mocking purposes
	CreateGroupMcpConfigExecute(r CreateGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error)

	/*
		CreateGroupMcpSecret Create One Secret for One Project MCP Configuration

		Creates a new secret on the ingress service account of the specified project-level MCP configuration. The plain-text secret value is returned only in this response and is never shown again.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param serviceAccountSecretRequest Secret creation request.
		@return CreateGroupMcpSecretApiRequest
	*/
	CreateGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateGroupMcpSecretApiRequest
	/*
		CreateGroupMcpSecret Create One Secret for One Project MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupMcpSecretApiParams - Parameters for the request
		@return CreateGroupMcpSecretApiRequest
	*/
	CreateGroupMcpSecretWithParams(ctx context.Context, args *CreateGroupMcpSecretApiParams) CreateGroupMcpSecretApiRequest

	// Method available only for mocking purposes
	CreateGroupMcpSecretExecute(r CreateGroupMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		CreateOrgMcpConfig Create One MCP Configuration for One Organization

		Creates an MCP configuration for the specified organization. Returns the configuration ID and ingress credentials.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param createOrgMcpConfigRequest MCP configuration to create.
		@return CreateOrgMcpConfigApiRequest
	*/
	CreateOrgMcpConfig(ctx context.Context, orgId string, createOrgMcpConfigRequest *CreateOrgMcpConfigRequest) CreateOrgMcpConfigApiRequest
	/*
		CreateOrgMcpConfig Create One MCP Configuration for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrgMcpConfigApiParams - Parameters for the request
		@return CreateOrgMcpConfigApiRequest
	*/
	CreateOrgMcpConfigWithParams(ctx context.Context, args *CreateOrgMcpConfigApiParams) CreateOrgMcpConfigApiRequest

	// Method available only for mocking purposes
	CreateOrgMcpConfigExecute(r CreateOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error)

	/*
		CreateOrgMcpSecret Create One Secret for One Organization MCP Configuration

		Creates a new secret on the ingress service account of the specified organization-level MCP configuration. The plain-text secret value is returned only in this response and is never shown again.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param serviceAccountSecretRequest Secret creation request.
		@return CreateOrgMcpSecretApiRequest
	*/
	CreateOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateOrgMcpSecretApiRequest
	/*
		CreateOrgMcpSecret Create One Secret for One Organization MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOrgMcpSecretApiParams - Parameters for the request
		@return CreateOrgMcpSecretApiRequest
	*/
	CreateOrgMcpSecretWithParams(ctx context.Context, args *CreateOrgMcpSecretApiParams) CreateOrgMcpSecretApiRequest

	// Method available only for mocking purposes
	CreateOrgMcpSecretExecute(r CreateOrgMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		DeleteGroupMcpConfig Delete One MCP Configuration for One Project

		Deletes the MCP configuration with the specified ID for the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration to delete.
		@return DeleteGroupMcpConfigApiRequest
	*/
	DeleteGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string) DeleteGroupMcpConfigApiRequest
	/*
		DeleteGroupMcpConfig Delete One MCP Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupMcpConfigApiParams - Parameters for the request
		@return DeleteGroupMcpConfigApiRequest
	*/
	DeleteGroupMcpConfigWithParams(ctx context.Context, args *DeleteGroupMcpConfigApiParams) DeleteGroupMcpConfigApiRequest

	// Method available only for mocking purposes
	DeleteGroupMcpConfigExecute(r DeleteGroupMcpConfigApiRequest) (*http.Response, error)

	/*
		DeleteGroupMcpSecret Delete One Secret for One Project MCP Configuration

		Deletes the specified secret from the ingress service account of the specified project-level MCP configuration.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@return DeleteGroupMcpSecretApiRequest
	*/
	DeleteGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, secretId string) DeleteGroupMcpSecretApiRequest
	/*
		DeleteGroupMcpSecret Delete One Secret for One Project MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupMcpSecretApiParams - Parameters for the request
		@return DeleteGroupMcpSecretApiRequest
	*/
	DeleteGroupMcpSecretWithParams(ctx context.Context, args *DeleteGroupMcpSecretApiParams) DeleteGroupMcpSecretApiRequest

	// Method available only for mocking purposes
	DeleteGroupMcpSecretExecute(r DeleteGroupMcpSecretApiRequest) (*http.Response, error)

	/*
		DeleteOrgMcpConfig Delete One MCP Configuration for One Organization

		Deletes the MCP configuration with the specified ID for the specified organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration to delete.
		@return DeleteOrgMcpConfigApiRequest
	*/
	DeleteOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string) DeleteOrgMcpConfigApiRequest
	/*
		DeleteOrgMcpConfig Delete One MCP Configuration for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrgMcpConfigApiParams - Parameters for the request
		@return DeleteOrgMcpConfigApiRequest
	*/
	DeleteOrgMcpConfigWithParams(ctx context.Context, args *DeleteOrgMcpConfigApiParams) DeleteOrgMcpConfigApiRequest

	// Method available only for mocking purposes
	DeleteOrgMcpConfigExecute(r DeleteOrgMcpConfigApiRequest) (*http.Response, error)

	/*
		DeleteOrgMcpSecret Delete One Secret for One Organization MCP Configuration

		Deletes the specified secret from the ingress service account of the specified organization-level MCP configuration.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@return DeleteOrgMcpSecretApiRequest
	*/
	DeleteOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, secretId string) DeleteOrgMcpSecretApiRequest
	/*
		DeleteOrgMcpSecret Delete One Secret for One Organization MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOrgMcpSecretApiParams - Parameters for the request
		@return DeleteOrgMcpSecretApiRequest
	*/
	DeleteOrgMcpSecretWithParams(ctx context.Context, args *DeleteOrgMcpSecretApiParams) DeleteOrgMcpSecretApiRequest

	// Method available only for mocking purposes
	DeleteOrgMcpSecretExecute(r DeleteOrgMcpSecretApiRequest) (*http.Response, error)

	/*
		GetGroupMcpConfig Return One MCP Configuration for One Project

		Returns the MCP configuration with the specified ID for the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@return GetGroupMcpConfigApiRequest
	*/
	GetGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string) GetGroupMcpConfigApiRequest
	/*
		GetGroupMcpConfig Return One MCP Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupMcpConfigApiParams - Parameters for the request
		@return GetGroupMcpConfigApiRequest
	*/
	GetGroupMcpConfigWithParams(ctx context.Context, args *GetGroupMcpConfigApiParams) GetGroupMcpConfigApiRequest

	// Method available only for mocking purposes
	GetGroupMcpConfigExecute(r GetGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error)

	/*
		GetGroupMcpSecret Return One Secret for One Project MCP Configuration

		Returns metadata for the specified secret on the ingress service account of a project-level MCP configuration. The secret value is never returned.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@return GetGroupMcpSecretApiRequest
	*/
	GetGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, secretId string) GetGroupMcpSecretApiRequest
	/*
		GetGroupMcpSecret Return One Secret for One Project MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupMcpSecretApiParams - Parameters for the request
		@return GetGroupMcpSecretApiRequest
	*/
	GetGroupMcpSecretWithParams(ctx context.Context, args *GetGroupMcpSecretApiParams) GetGroupMcpSecretApiRequest

	// Method available only for mocking purposes
	GetGroupMcpSecretExecute(r GetGroupMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		GetOrgMcpConfig Return One MCP Configuration for One Organization

		Returns the MCP configuration with the specified ID for the specified organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@return GetOrgMcpConfigApiRequest
	*/
	GetOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string) GetOrgMcpConfigApiRequest
	/*
		GetOrgMcpConfig Return One MCP Configuration for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgMcpConfigApiParams - Parameters for the request
		@return GetOrgMcpConfigApiRequest
	*/
	GetOrgMcpConfigWithParams(ctx context.Context, args *GetOrgMcpConfigApiParams) GetOrgMcpConfigApiRequest

	// Method available only for mocking purposes
	GetOrgMcpConfigExecute(r GetOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error)

	/*
		GetOrgMcpSecret Return One Secret for One Organization MCP Configuration

		Returns metadata for the specified secret on the ingress service account of an organization-level MCP configuration. The secret value is never returned.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@param secretId Unique 24-hexadecimal digit string that identifies the secret.
		@return GetOrgMcpSecretApiRequest
	*/
	GetOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, secretId string) GetOrgMcpSecretApiRequest
	/*
		GetOrgMcpSecret Return One Secret for One Organization MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgMcpSecretApiParams - Parameters for the request
		@return GetOrgMcpSecretApiRequest
	*/
	GetOrgMcpSecretWithParams(ctx context.Context, args *GetOrgMcpSecretApiParams) GetOrgMcpSecretApiRequest

	// Method available only for mocking purposes
	GetOrgMcpSecretExecute(r GetOrgMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error)

	/*
		ListGroupMcpConfigs Return All MCP Configurations for One Project

		Returns all MCP configurations associated with the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListGroupMcpConfigsApiRequest
	*/
	ListGroupMcpConfigs(ctx context.Context, groupId string) ListGroupMcpConfigsApiRequest
	/*
		ListGroupMcpConfigs Return All MCP Configurations for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupMcpConfigsApiParams - Parameters for the request
		@return ListGroupMcpConfigsApiRequest
	*/
	ListGroupMcpConfigsWithParams(ctx context.Context, args *ListGroupMcpConfigsApiParams) ListGroupMcpConfigsApiRequest

	// Method available only for mocking purposes
	ListGroupMcpConfigsExecute(r ListGroupMcpConfigsApiRequest) (*PaginatedGroupMcpConfig, *http.Response, error)

	/*
		ListGroupMcpSecrets Return All Secrets for One Project MCP Configuration

		Returns metadata for all secrets on the ingress service account of the specified project-level MCP configuration. Secret values are never returned.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@return ListGroupMcpSecretsApiRequest
	*/
	ListGroupMcpSecrets(ctx context.Context, groupId string, mcpConfigId string) ListGroupMcpSecretsApiRequest
	/*
		ListGroupMcpSecrets Return All Secrets for One Project MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupMcpSecretsApiParams - Parameters for the request
		@return ListGroupMcpSecretsApiRequest
	*/
	ListGroupMcpSecretsWithParams(ctx context.Context, args *ListGroupMcpSecretsApiParams) ListGroupMcpSecretsApiRequest

	// Method available only for mocking purposes
	ListGroupMcpSecretsExecute(r ListGroupMcpSecretsApiRequest) (*PaginatedMcpConfigSecret, *http.Response, error)

	/*
		ListOrgMcpConfigs Return All MCP Configurations for One Organization

		Returns all MCP configurations associated with the specified organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListOrgMcpConfigsApiRequest
	*/
	ListOrgMcpConfigs(ctx context.Context, orgId string) ListOrgMcpConfigsApiRequest
	/*
		ListOrgMcpConfigs Return All MCP Configurations for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgMcpConfigsApiParams - Parameters for the request
		@return ListOrgMcpConfigsApiRequest
	*/
	ListOrgMcpConfigsWithParams(ctx context.Context, args *ListOrgMcpConfigsApiParams) ListOrgMcpConfigsApiRequest

	// Method available only for mocking purposes
	ListOrgMcpConfigsExecute(r ListOrgMcpConfigsApiRequest) (*PaginatedOrgMcpConfig, *http.Response, error)

	/*
		ListOrgMcpSecrets Return All Secrets for One Organization MCP Configuration

		Returns metadata for all secrets on the ingress service account of the specified organization-level MCP configuration. Secret values are never returned.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration.
		@return ListOrgMcpSecretsApiRequest
	*/
	ListOrgMcpSecrets(ctx context.Context, orgId string, mcpConfigId string) ListOrgMcpSecretsApiRequest
	/*
		ListOrgMcpSecrets Return All Secrets for One Organization MCP Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrgMcpSecretsApiParams - Parameters for the request
		@return ListOrgMcpSecretsApiRequest
	*/
	ListOrgMcpSecretsWithParams(ctx context.Context, args *ListOrgMcpSecretsApiParams) ListOrgMcpSecretsApiRequest

	// Method available only for mocking purposes
	ListOrgMcpSecretsExecute(r ListOrgMcpSecretsApiRequest) (*PaginatedMcpConfigSecret, *http.Response, error)

	/*
		UpdateGroupMcpConfig Update One MCP Configuration for One Project

		Updates the specified MCP configuration for the project. Supports partial updates: only provided fields are changed.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param mcpConfigId Unique identifier of the MCP configuration to update.
		@param groupMcpConfigUpdateRequest MCP configuration fields to update.
		@return UpdateGroupMcpConfigApiRequest
	*/
	UpdateGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string, groupMcpConfigUpdateRequest *GroupMcpConfigUpdateRequest) UpdateGroupMcpConfigApiRequest
	/*
		UpdateGroupMcpConfig Update One MCP Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupMcpConfigApiParams - Parameters for the request
		@return UpdateGroupMcpConfigApiRequest
	*/
	UpdateGroupMcpConfigWithParams(ctx context.Context, args *UpdateGroupMcpConfigApiParams) UpdateGroupMcpConfigApiRequest

	// Method available only for mocking purposes
	UpdateGroupMcpConfigExecute(r UpdateGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error)

	/*
		UpdateOrgMcpConfig Update One MCP Configuration for One Organization

		Updates the specified MCP configuration for the organization. Supports partial updates: only provided fields are changed.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param mcpConfigId Unique identifier of the MCP configuration to update.
		@param orgMcpConfigUpdateRequest MCP configuration fields to update.
		@return UpdateOrgMcpConfigApiRequest
	*/
	UpdateOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string, orgMcpConfigUpdateRequest *OrgMcpConfigUpdateRequest) UpdateOrgMcpConfigApiRequest
	/*
		UpdateOrgMcpConfig Update One MCP Configuration for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOrgMcpConfigApiParams - Parameters for the request
		@return UpdateOrgMcpConfigApiRequest
	*/
	UpdateOrgMcpConfigWithParams(ctx context.Context, args *UpdateOrgMcpConfigApiParams) UpdateOrgMcpConfigApiRequest

	// Method available only for mocking purposes
	UpdateOrgMcpConfigExecute(r UpdateOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error)
}

// RemoteMCPConfigurationsApiService RemoteMCPConfigurationsApi service
type RemoteMCPConfigurationsApiService service

type CreateGroupMcpConfigApiRequest struct {
	ctx                         context.Context
	ApiService                  RemoteMCPConfigurationsApi
	groupId                     string
	createGroupMcpConfigRequest *CreateGroupMcpConfigRequest
}

type CreateGroupMcpConfigApiParams struct {
	GroupId                     string
	CreateGroupMcpConfigRequest *CreateGroupMcpConfigRequest
}

func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpConfigWithParams(ctx context.Context, args *CreateGroupMcpConfigApiParams) CreateGroupMcpConfigApiRequest {
	return CreateGroupMcpConfigApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		createGroupMcpConfigRequest: args.CreateGroupMcpConfigRequest,
	}
}

func (r CreateGroupMcpConfigApiRequest) Execute() (*GroupMcpConfigResponse, *http.Response, error) {
	return r.ApiService.CreateGroupMcpConfigExecute(r)
}

/*
CreateGroupMcpConfig Create One MCP Configuration for One Project

Creates an MCP configuration for the specified project. Returns the configuration ID and ingress credentials.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateGroupMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpConfig(ctx context.Context, groupId string, createGroupMcpConfigRequest *CreateGroupMcpConfigRequest) CreateGroupMcpConfigApiRequest {
	return CreateGroupMcpConfigApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		createGroupMcpConfigRequest: createGroupMcpConfigRequest,
	}
}

// CreateGroupMcpConfigExecute executes the request
//
//	@return GroupMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpConfigExecute(r CreateGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.CreateGroupMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createGroupMcpConfigRequest == nil {
		return localVarReturnValue, nil, reportError("createGroupMcpConfigRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createGroupMcpConfigRequest
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

type CreateGroupMcpSecretApiRequest struct {
	ctx                         context.Context
	ApiService                  RemoteMCPConfigurationsApi
	groupId                     string
	mcpConfigId                 string
	serviceAccountSecretRequest *ServiceAccountSecretRequest
}

type CreateGroupMcpSecretApiParams struct {
	GroupId                     string
	McpConfigId                 string
	ServiceAccountSecretRequest *ServiceAccountSecretRequest
}

func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpSecretWithParams(ctx context.Context, args *CreateGroupMcpSecretApiParams) CreateGroupMcpSecretApiRequest {
	return CreateGroupMcpSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		mcpConfigId:                 args.McpConfigId,
		serviceAccountSecretRequest: args.ServiceAccountSecretRequest,
	}
}

func (r CreateGroupMcpSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.CreateGroupMcpSecretExecute(r)
}

/*
CreateGroupMcpSecret Create One Secret for One Project MCP Configuration

Creates a new secret on the ingress service account of the specified project-level MCP configuration. The plain-text secret value is returned only in this response and is never shown again.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return CreateGroupMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateGroupMcpSecretApiRequest {
	return CreateGroupMcpSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		mcpConfigId:                 mcpConfigId,
		serviceAccountSecretRequest: serviceAccountSecretRequest,
	}
}

// CreateGroupMcpSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *RemoteMCPConfigurationsApiService) CreateGroupMcpSecretExecute(r CreateGroupMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.CreateGroupMcpSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountSecretRequest == nil {
		return localVarReturnValue, nil, reportError("serviceAccountSecretRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type CreateOrgMcpConfigApiRequest struct {
	ctx                       context.Context
	ApiService                RemoteMCPConfigurationsApi
	orgId                     string
	createOrgMcpConfigRequest *CreateOrgMcpConfigRequest
}

type CreateOrgMcpConfigApiParams struct {
	OrgId                     string
	CreateOrgMcpConfigRequest *CreateOrgMcpConfigRequest
}

func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpConfigWithParams(ctx context.Context, args *CreateOrgMcpConfigApiParams) CreateOrgMcpConfigApiRequest {
	return CreateOrgMcpConfigApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		orgId:                     args.OrgId,
		createOrgMcpConfigRequest: args.CreateOrgMcpConfigRequest,
	}
}

func (r CreateOrgMcpConfigApiRequest) Execute() (*OrgMcpConfigResponse, *http.Response, error) {
	return r.ApiService.CreateOrgMcpConfigExecute(r)
}

/*
CreateOrgMcpConfig Create One MCP Configuration for One Organization

Creates an MCP configuration for the specified organization. Returns the configuration ID and ingress credentials.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateOrgMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpConfig(ctx context.Context, orgId string, createOrgMcpConfigRequest *CreateOrgMcpConfigRequest) CreateOrgMcpConfigApiRequest {
	return CreateOrgMcpConfigApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		orgId:                     orgId,
		createOrgMcpConfigRequest: createOrgMcpConfigRequest,
	}
}

// CreateOrgMcpConfigExecute executes the request
//
//	@return OrgMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpConfigExecute(r CreateOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.CreateOrgMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrgMcpConfigRequest == nil {
		return localVarReturnValue, nil, reportError("createOrgMcpConfigRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createOrgMcpConfigRequest
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

type CreateOrgMcpSecretApiRequest struct {
	ctx                         context.Context
	ApiService                  RemoteMCPConfigurationsApi
	orgId                       string
	mcpConfigId                 string
	serviceAccountSecretRequest *ServiceAccountSecretRequest
}

type CreateOrgMcpSecretApiParams struct {
	OrgId                       string
	McpConfigId                 string
	ServiceAccountSecretRequest *ServiceAccountSecretRequest
}

func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpSecretWithParams(ctx context.Context, args *CreateOrgMcpSecretApiParams) CreateOrgMcpSecretApiRequest {
	return CreateOrgMcpSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		orgId:                       args.OrgId,
		mcpConfigId:                 args.McpConfigId,
		serviceAccountSecretRequest: args.ServiceAccountSecretRequest,
	}
}

func (r CreateOrgMcpSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.CreateOrgMcpSecretExecute(r)
}

/*
CreateOrgMcpSecret Create One Secret for One Organization MCP Configuration

Creates a new secret on the ingress service account of the specified organization-level MCP configuration. The plain-text secret value is returned only in this response and is never shown again.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return CreateOrgMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, serviceAccountSecretRequest *ServiceAccountSecretRequest) CreateOrgMcpSecretApiRequest {
	return CreateOrgMcpSecretApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		orgId:                       orgId,
		mcpConfigId:                 mcpConfigId,
		serviceAccountSecretRequest: serviceAccountSecretRequest,
	}
}

// CreateOrgMcpSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *RemoteMCPConfigurationsApiService) CreateOrgMcpSecretExecute(r CreateOrgMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.CreateOrgMcpSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceAccountSecretRequest == nil {
		return localVarReturnValue, nil, reportError("serviceAccountSecretRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type DeleteGroupMcpConfigApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	groupId     string
	mcpConfigId string
	cascading   *bool
}

type DeleteGroupMcpConfigApiParams struct {
	GroupId     string
	McpConfigId string
	Cascading   *bool
}

func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpConfigWithParams(ctx context.Context, args *DeleteGroupMcpConfigApiParams) DeleteGroupMcpConfigApiRequest {
	return DeleteGroupMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		mcpConfigId: args.McpConfigId,
		cascading:   args.Cascading,
	}
}

// Flag that indicates whether to delete the MCP configuration even if it has active secrets. If false and active secrets exist, the request returns an error. Defaults to false.
func (r DeleteGroupMcpConfigApiRequest) Cascading(cascading bool) DeleteGroupMcpConfigApiRequest {
	r.cascading = &cascading
	return r
}

func (r DeleteGroupMcpConfigApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupMcpConfigExecute(r)
}

/*
DeleteGroupMcpConfig Delete One MCP Configuration for One Project

Deletes the MCP configuration with the specified ID for the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration to delete.
	@return DeleteGroupMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string) DeleteGroupMcpConfigApiRequest {
	return DeleteGroupMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		mcpConfigId: mcpConfigId,
	}
}

// DeleteGroupMcpConfigExecute executes the request
func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpConfigExecute(r DeleteGroupMcpConfigApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.DeleteGroupMcpConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cascading != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cascading", r.cascading, "")
	} else {
		var defaultValue bool = false
		r.cascading = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "cascading", r.cascading, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type DeleteGroupMcpSecretApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	groupId     string
	mcpConfigId string
	secretId    string
}

type DeleteGroupMcpSecretApiParams struct {
	GroupId     string
	McpConfigId string
	SecretId    string
}

func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpSecretWithParams(ctx context.Context, args *DeleteGroupMcpSecretApiParams) DeleteGroupMcpSecretApiRequest {
	return DeleteGroupMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		mcpConfigId: args.McpConfigId,
		secretId:    args.SecretId,
	}
}

func (r DeleteGroupMcpSecretApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupMcpSecretExecute(r)
}

/*
DeleteGroupMcpSecret Delete One Secret for One Project MCP Configuration

Deletes the specified secret from the ingress service account of the specified project-level MCP configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@return DeleteGroupMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, secretId string) DeleteGroupMcpSecretApiRequest {
	return DeleteGroupMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		mcpConfigId: mcpConfigId,
		secretId:    secretId,
	}
}

// DeleteGroupMcpSecretExecute executes the request
func (a *RemoteMCPConfigurationsApiService) DeleteGroupMcpSecretExecute(r DeleteGroupMcpSecretApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.DeleteGroupMcpSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets/{secretId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)
	if r.secretId == "" {
		return nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type DeleteOrgMcpConfigApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	orgId       string
	mcpConfigId string
	cascading   *bool
}

type DeleteOrgMcpConfigApiParams struct {
	OrgId       string
	McpConfigId string
	Cascading   *bool
}

func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpConfigWithParams(ctx context.Context, args *DeleteOrgMcpConfigApiParams) DeleteOrgMcpConfigApiRequest {
	return DeleteOrgMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       args.OrgId,
		mcpConfigId: args.McpConfigId,
		cascading:   args.Cascading,
	}
}

// Flag that indicates whether to delete the MCP configuration even if it has active secrets. If false and active secrets exist, the request returns an error. Defaults to false.
func (r DeleteOrgMcpConfigApiRequest) Cascading(cascading bool) DeleteOrgMcpConfigApiRequest {
	r.cascading = &cascading
	return r
}

func (r DeleteOrgMcpConfigApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrgMcpConfigExecute(r)
}

/*
DeleteOrgMcpConfig Delete One MCP Configuration for One Organization

Deletes the MCP configuration with the specified ID for the specified organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration to delete.
	@return DeleteOrgMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string) DeleteOrgMcpConfigApiRequest {
	return DeleteOrgMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       orgId,
		mcpConfigId: mcpConfigId,
	}
}

// DeleteOrgMcpConfigExecute executes the request
func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpConfigExecute(r DeleteOrgMcpConfigApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.DeleteOrgMcpConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cascading != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cascading", r.cascading, "")
	} else {
		var defaultValue bool = false
		r.cascading = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "cascading", r.cascading, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type DeleteOrgMcpSecretApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	orgId       string
	mcpConfigId string
	secretId    string
}

type DeleteOrgMcpSecretApiParams struct {
	OrgId       string
	McpConfigId string
	SecretId    string
}

func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpSecretWithParams(ctx context.Context, args *DeleteOrgMcpSecretApiParams) DeleteOrgMcpSecretApiRequest {
	return DeleteOrgMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       args.OrgId,
		mcpConfigId: args.McpConfigId,
		secretId:    args.SecretId,
	}
}

func (r DeleteOrgMcpSecretApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrgMcpSecretExecute(r)
}

/*
DeleteOrgMcpSecret Delete One Secret for One Organization MCP Configuration

Deletes the specified secret from the ingress service account of the specified organization-level MCP configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@return DeleteOrgMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, secretId string) DeleteOrgMcpSecretApiRequest {
	return DeleteOrgMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       orgId,
		mcpConfigId: mcpConfigId,
		secretId:    secretId,
	}
}

// DeleteOrgMcpSecretExecute executes the request
func (a *RemoteMCPConfigurationsApiService) DeleteOrgMcpSecretExecute(r DeleteOrgMcpSecretApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.DeleteOrgMcpSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets/{secretId}"
	if r.orgId == "" {
		return nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)
	if r.secretId == "" {
		return nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetGroupMcpConfigApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	groupId     string
	mcpConfigId string
}

type GetGroupMcpConfigApiParams struct {
	GroupId     string
	McpConfigId string
}

func (a *RemoteMCPConfigurationsApiService) GetGroupMcpConfigWithParams(ctx context.Context, args *GetGroupMcpConfigApiParams) GetGroupMcpConfigApiRequest {
	return GetGroupMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		mcpConfigId: args.McpConfigId,
	}
}

func (r GetGroupMcpConfigApiRequest) Execute() (*GroupMcpConfigResponse, *http.Response, error) {
	return r.ApiService.GetGroupMcpConfigExecute(r)
}

/*
GetGroupMcpConfig Return One MCP Configuration for One Project

Returns the MCP configuration with the specified ID for the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return GetGroupMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) GetGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string) GetGroupMcpConfigApiRequest {
	return GetGroupMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		mcpConfigId: mcpConfigId,
	}
}

// GetGroupMcpConfigExecute executes the request
//
//	@return GroupMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) GetGroupMcpConfigExecute(r GetGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.GetGroupMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetGroupMcpSecretApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	groupId     string
	mcpConfigId string
	secretId    string
}

type GetGroupMcpSecretApiParams struct {
	GroupId     string
	McpConfigId string
	SecretId    string
}

func (a *RemoteMCPConfigurationsApiService) GetGroupMcpSecretWithParams(ctx context.Context, args *GetGroupMcpSecretApiParams) GetGroupMcpSecretApiRequest {
	return GetGroupMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		mcpConfigId: args.McpConfigId,
		secretId:    args.SecretId,
	}
}

func (r GetGroupMcpSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.GetGroupMcpSecretExecute(r)
}

/*
GetGroupMcpSecret Return One Secret for One Project MCP Configuration

Returns metadata for the specified secret on the ingress service account of a project-level MCP configuration. The secret value is never returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@return GetGroupMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) GetGroupMcpSecret(ctx context.Context, groupId string, mcpConfigId string, secretId string) GetGroupMcpSecretApiRequest {
	return GetGroupMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		mcpConfigId: mcpConfigId,
		secretId:    secretId,
	}
}

// GetGroupMcpSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *RemoteMCPConfigurationsApiService) GetGroupMcpSecretExecute(r GetGroupMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.GetGroupMcpSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets/{secretId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)
	if r.secretId == "" {
		return localVarReturnValue, nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetOrgMcpConfigApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	orgId       string
	mcpConfigId string
}

type GetOrgMcpConfigApiParams struct {
	OrgId       string
	McpConfigId string
}

func (a *RemoteMCPConfigurationsApiService) GetOrgMcpConfigWithParams(ctx context.Context, args *GetOrgMcpConfigApiParams) GetOrgMcpConfigApiRequest {
	return GetOrgMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       args.OrgId,
		mcpConfigId: args.McpConfigId,
	}
}

func (r GetOrgMcpConfigApiRequest) Execute() (*OrgMcpConfigResponse, *http.Response, error) {
	return r.ApiService.GetOrgMcpConfigExecute(r)
}

/*
GetOrgMcpConfig Return One MCP Configuration for One Organization

Returns the MCP configuration with the specified ID for the specified organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return GetOrgMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) GetOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string) GetOrgMcpConfigApiRequest {
	return GetOrgMcpConfigApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       orgId,
		mcpConfigId: mcpConfigId,
	}
}

// GetOrgMcpConfigExecute executes the request
//
//	@return OrgMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) GetOrgMcpConfigExecute(r GetOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.GetOrgMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetOrgMcpSecretApiRequest struct {
	ctx         context.Context
	ApiService  RemoteMCPConfigurationsApi
	orgId       string
	mcpConfigId string
	secretId    string
}

type GetOrgMcpSecretApiParams struct {
	OrgId       string
	McpConfigId string
	SecretId    string
}

func (a *RemoteMCPConfigurationsApiService) GetOrgMcpSecretWithParams(ctx context.Context, args *GetOrgMcpSecretApiParams) GetOrgMcpSecretApiRequest {
	return GetOrgMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       args.OrgId,
		mcpConfigId: args.McpConfigId,
		secretId:    args.SecretId,
	}
}

func (r GetOrgMcpSecretApiRequest) Execute() (*ServiceAccountSecret, *http.Response, error) {
	return r.ApiService.GetOrgMcpSecretExecute(r)
}

/*
GetOrgMcpSecret Return One Secret for One Organization MCP Configuration

Returns metadata for the specified secret on the ingress service account of an organization-level MCP configuration. The secret value is never returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@param secretId Unique 24-hexadecimal digit string that identifies the secret.
	@return GetOrgMcpSecretApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) GetOrgMcpSecret(ctx context.Context, orgId string, mcpConfigId string, secretId string) GetOrgMcpSecretApiRequest {
	return GetOrgMcpSecretApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       orgId,
		mcpConfigId: mcpConfigId,
		secretId:    secretId,
	}
}

// GetOrgMcpSecretExecute executes the request
//
//	@return ServiceAccountSecret
func (a *RemoteMCPConfigurationsApiService) GetOrgMcpSecretExecute(r GetOrgMcpSecretApiRequest) (*ServiceAccountSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServiceAccountSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.GetOrgMcpSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets/{secretId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)
	if r.secretId == "" {
		return localVarReturnValue, nil, reportError("secretId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(r.secretId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type ListGroupMcpConfigsApiRequest struct {
	ctx          context.Context
	ApiService   RemoteMCPConfigurationsApi
	groupId      string
	itemsPerPage *int
	includeCount *bool
	pageNum      *int
}

type ListGroupMcpConfigsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	IncludeCount *bool
	PageNum      *int
}

func (a *RemoteMCPConfigurationsApiService) ListGroupMcpConfigsWithParams(ctx context.Context, args *ListGroupMcpConfigsApiParams) ListGroupMcpConfigsApiRequest {
	return ListGroupMcpConfigsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		includeCount: args.IncludeCount,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupMcpConfigsApiRequest) ItemsPerPage(itemsPerPage int) ListGroupMcpConfigsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
func (r ListGroupMcpConfigsApiRequest) IncludeCount(includeCount bool) ListGroupMcpConfigsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupMcpConfigsApiRequest) PageNum(pageNum int) ListGroupMcpConfigsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupMcpConfigsApiRequest) Execute() (*PaginatedGroupMcpConfig, *http.Response, error) {
	return r.ApiService.ListGroupMcpConfigsExecute(r)
}

/*
ListGroupMcpConfigs Return All MCP Configurations for One Project

Returns all MCP configurations associated with the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListGroupMcpConfigsApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) ListGroupMcpConfigs(ctx context.Context, groupId string) ListGroupMcpConfigsApiRequest {
	return ListGroupMcpConfigsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupMcpConfigsExecute executes the request
//
//	@return PaginatedGroupMcpConfig
func (a *RemoteMCPConfigurationsApiService) ListGroupMcpConfigsExecute(r ListGroupMcpConfigsApiRequest) (*PaginatedGroupMcpConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedGroupMcpConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.ListGroupMcpConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type ListGroupMcpSecretsApiRequest struct {
	ctx          context.Context
	ApiService   RemoteMCPConfigurationsApi
	groupId      string
	mcpConfigId  string
	itemsPerPage *int
	includeCount *bool
	pageNum      *int
}

type ListGroupMcpSecretsApiParams struct {
	GroupId      string
	McpConfigId  string
	ItemsPerPage *int
	IncludeCount *bool
	PageNum      *int
}

func (a *RemoteMCPConfigurationsApiService) ListGroupMcpSecretsWithParams(ctx context.Context, args *ListGroupMcpSecretsApiParams) ListGroupMcpSecretsApiRequest {
	return ListGroupMcpSecretsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		mcpConfigId:  args.McpConfigId,
		itemsPerPage: args.ItemsPerPage,
		includeCount: args.IncludeCount,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupMcpSecretsApiRequest) ItemsPerPage(itemsPerPage int) ListGroupMcpSecretsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
func (r ListGroupMcpSecretsApiRequest) IncludeCount(includeCount bool) ListGroupMcpSecretsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupMcpSecretsApiRequest) PageNum(pageNum int) ListGroupMcpSecretsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupMcpSecretsApiRequest) Execute() (*PaginatedMcpConfigSecret, *http.Response, error) {
	return r.ApiService.ListGroupMcpSecretsExecute(r)
}

/*
ListGroupMcpSecrets Return All Secrets for One Project MCP Configuration

Returns metadata for all secrets on the ingress service account of the specified project-level MCP configuration. Secret values are never returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return ListGroupMcpSecretsApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) ListGroupMcpSecrets(ctx context.Context, groupId string, mcpConfigId string) ListGroupMcpSecretsApiRequest {
	return ListGroupMcpSecretsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		mcpConfigId: mcpConfigId,
	}
}

// ListGroupMcpSecretsExecute executes the request
//
//	@return PaginatedMcpConfigSecret
func (a *RemoteMCPConfigurationsApiService) ListGroupMcpSecretsExecute(r ListGroupMcpSecretsApiRequest) (*PaginatedMcpConfigSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedMcpConfigSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.ListGroupMcpSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}/secrets"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type ListOrgMcpConfigsApiRequest struct {
	ctx          context.Context
	ApiService   RemoteMCPConfigurationsApi
	orgId        string
	itemsPerPage *int
	includeCount *bool
	pageNum      *int
}

type ListOrgMcpConfigsApiParams struct {
	OrgId        string
	ItemsPerPage *int
	IncludeCount *bool
	PageNum      *int
}

func (a *RemoteMCPConfigurationsApiService) ListOrgMcpConfigsWithParams(ctx context.Context, args *ListOrgMcpConfigsApiParams) ListOrgMcpConfigsApiRequest {
	return ListOrgMcpConfigsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		itemsPerPage: args.ItemsPerPage,
		includeCount: args.IncludeCount,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgMcpConfigsApiRequest) ItemsPerPage(itemsPerPage int) ListOrgMcpConfigsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
func (r ListOrgMcpConfigsApiRequest) IncludeCount(includeCount bool) ListOrgMcpConfigsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgMcpConfigsApiRequest) PageNum(pageNum int) ListOrgMcpConfigsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgMcpConfigsApiRequest) Execute() (*PaginatedOrgMcpConfig, *http.Response, error) {
	return r.ApiService.ListOrgMcpConfigsExecute(r)
}

/*
ListOrgMcpConfigs Return All MCP Configurations for One Organization

Returns all MCP configurations associated with the specified organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrgMcpConfigsApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) ListOrgMcpConfigs(ctx context.Context, orgId string) ListOrgMcpConfigsApiRequest {
	return ListOrgMcpConfigsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrgMcpConfigsExecute executes the request
//
//	@return PaginatedOrgMcpConfig
func (a *RemoteMCPConfigurationsApiService) ListOrgMcpConfigsExecute(r ListOrgMcpConfigsApiRequest) (*PaginatedOrgMcpConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedOrgMcpConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.ListOrgMcpConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type ListOrgMcpSecretsApiRequest struct {
	ctx          context.Context
	ApiService   RemoteMCPConfigurationsApi
	orgId        string
	mcpConfigId  string
	itemsPerPage *int
	includeCount *bool
	pageNum      *int
}

type ListOrgMcpSecretsApiParams struct {
	OrgId        string
	McpConfigId  string
	ItemsPerPage *int
	IncludeCount *bool
	PageNum      *int
}

func (a *RemoteMCPConfigurationsApiService) ListOrgMcpSecretsWithParams(ctx context.Context, args *ListOrgMcpSecretsApiParams) ListOrgMcpSecretsApiRequest {
	return ListOrgMcpSecretsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		mcpConfigId:  args.McpConfigId,
		itemsPerPage: args.ItemsPerPage,
		includeCount: args.IncludeCount,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListOrgMcpSecretsApiRequest) ItemsPerPage(itemsPerPage int) ListOrgMcpSecretsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
func (r ListOrgMcpSecretsApiRequest) IncludeCount(includeCount bool) ListOrgMcpSecretsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrgMcpSecretsApiRequest) PageNum(pageNum int) ListOrgMcpSecretsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOrgMcpSecretsApiRequest) Execute() (*PaginatedMcpConfigSecret, *http.Response, error) {
	return r.ApiService.ListOrgMcpSecretsExecute(r)
}

/*
ListOrgMcpSecrets Return All Secrets for One Organization MCP Configuration

Returns metadata for all secrets on the ingress service account of the specified organization-level MCP configuration. Secret values are never returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration.
	@return ListOrgMcpSecretsApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) ListOrgMcpSecrets(ctx context.Context, orgId string, mcpConfigId string) ListOrgMcpSecretsApiRequest {
	return ListOrgMcpSecretsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		orgId:       orgId,
		mcpConfigId: mcpConfigId,
	}
}

// ListOrgMcpSecretsExecute executes the request
//
//	@return PaginatedMcpConfigSecret
func (a *RemoteMCPConfigurationsApiService) ListOrgMcpSecretsExecute(r ListOrgMcpSecretsApiRequest) (*PaginatedMcpConfigSecret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedMcpConfigSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.ListOrgMcpSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}/secrets"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type UpdateGroupMcpConfigApiRequest struct {
	ctx                         context.Context
	ApiService                  RemoteMCPConfigurationsApi
	groupId                     string
	mcpConfigId                 string
	groupMcpConfigUpdateRequest *GroupMcpConfigUpdateRequest
}

type UpdateGroupMcpConfigApiParams struct {
	GroupId                     string
	McpConfigId                 string
	GroupMcpConfigUpdateRequest *GroupMcpConfigUpdateRequest
}

func (a *RemoteMCPConfigurationsApiService) UpdateGroupMcpConfigWithParams(ctx context.Context, args *UpdateGroupMcpConfigApiParams) UpdateGroupMcpConfigApiRequest {
	return UpdateGroupMcpConfigApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     args.GroupId,
		mcpConfigId:                 args.McpConfigId,
		groupMcpConfigUpdateRequest: args.GroupMcpConfigUpdateRequest,
	}
}

func (r UpdateGroupMcpConfigApiRequest) Execute() (*GroupMcpConfigResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupMcpConfigExecute(r)
}

/*
UpdateGroupMcpConfig Update One MCP Configuration for One Project

Updates the specified MCP configuration for the project. Supports partial updates: only provided fields are changed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param mcpConfigId Unique identifier of the MCP configuration to update.
	@return UpdateGroupMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) UpdateGroupMcpConfig(ctx context.Context, groupId string, mcpConfigId string, groupMcpConfigUpdateRequest *GroupMcpConfigUpdateRequest) UpdateGroupMcpConfigApiRequest {
	return UpdateGroupMcpConfigApiRequest{
		ApiService:                  a,
		ctx:                         ctx,
		groupId:                     groupId,
		mcpConfigId:                 mcpConfigId,
		groupMcpConfigUpdateRequest: groupMcpConfigUpdateRequest,
	}
}

// UpdateGroupMcpConfigExecute executes the request
//
//	@return GroupMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) UpdateGroupMcpConfigExecute(r UpdateGroupMcpConfigApiRequest) (*GroupMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.UpdateGroupMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/mcpConfigs/{mcpConfigId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupMcpConfigUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("groupMcpConfigUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupMcpConfigUpdateRequest
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

type UpdateOrgMcpConfigApiRequest struct {
	ctx                       context.Context
	ApiService                RemoteMCPConfigurationsApi
	orgId                     string
	mcpConfigId               string
	orgMcpConfigUpdateRequest *OrgMcpConfigUpdateRequest
}

type UpdateOrgMcpConfigApiParams struct {
	OrgId                     string
	McpConfigId               string
	OrgMcpConfigUpdateRequest *OrgMcpConfigUpdateRequest
}

func (a *RemoteMCPConfigurationsApiService) UpdateOrgMcpConfigWithParams(ctx context.Context, args *UpdateOrgMcpConfigApiParams) UpdateOrgMcpConfigApiRequest {
	return UpdateOrgMcpConfigApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		orgId:                     args.OrgId,
		mcpConfigId:               args.McpConfigId,
		orgMcpConfigUpdateRequest: args.OrgMcpConfigUpdateRequest,
	}
}

func (r UpdateOrgMcpConfigApiRequest) Execute() (*OrgMcpConfigResponse, *http.Response, error) {
	return r.ApiService.UpdateOrgMcpConfigExecute(r)
}

/*
UpdateOrgMcpConfig Update One MCP Configuration for One Organization

Updates the specified MCP configuration for the organization. Supports partial updates: only provided fields are changed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param mcpConfigId Unique identifier of the MCP configuration to update.
	@return UpdateOrgMcpConfigApiRequest
*/
func (a *RemoteMCPConfigurationsApiService) UpdateOrgMcpConfig(ctx context.Context, orgId string, mcpConfigId string, orgMcpConfigUpdateRequest *OrgMcpConfigUpdateRequest) UpdateOrgMcpConfigApiRequest {
	return UpdateOrgMcpConfigApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		orgId:                     orgId,
		mcpConfigId:               mcpConfigId,
		orgMcpConfigUpdateRequest: orgMcpConfigUpdateRequest,
	}
}

// UpdateOrgMcpConfigExecute executes the request
//
//	@return OrgMcpConfigResponse
func (a *RemoteMCPConfigurationsApiService) UpdateOrgMcpConfigExecute(r UpdateOrgMcpConfigApiRequest) (*OrgMcpConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgMcpConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteMCPConfigurationsApiService.UpdateOrgMcpConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/mcpConfigs/{mcpConfigId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.mcpConfigId == "" {
		return localVarReturnValue, nil, reportError("mcpConfigId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"mcpConfigId"+"}", url.PathEscape(r.mcpConfigId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgMcpConfigUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("orgMcpConfigUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orgMcpConfigUpdateRequest
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
