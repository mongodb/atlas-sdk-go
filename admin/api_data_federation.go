// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DataFederationApi interface {

	/*
			CreateDataFederationPrivateEndpoint Create One Federated Database Instance and Online Archive Private Endpoint for One Project

			Adds one private endpoint for Federated Database Instances and Online Archives to the specified projects. If the endpoint ID already exists and the associated comment is unchanged, Atlas Data Federation makes no change to the endpoint ID list. If the endpoint ID already exists and the associated comment is changed, Atlas Data Federation updates the comment value only in the endpoint ID list. If the endpoint ID doesn't exist, Atlas Data Federation appends the new endpoint to the list of endpoints in the endpoint ID list. Each region has an associated service name for the various endpoints in each region.

		 `us-east-1` is `com.amazonaws.vpce.us-east-1.vpce-svc-00e311695874992b4`.

		 `us-west-1` is `com.amazonaws.vpce.us-west-2.vpce-svc-09d86b19e59d1b4bb`.

		 `eu-west-1` is `com.amazonaws.vpce.eu-west-1.vpce-svc-0824460b72e1a420e`.

		 `eu-west-2` is `com.amazonaws.vpce.eu-west-2.vpce-svc-052f1840aa0c4f1f9`.

		 `eu-central-1` is `com.amazonaws.vpce.eu-central-1.vpce-svc-0ac8ce91871138c0d`.

		 `sa-east-1` is `com.amazonaws.vpce.sa-east-1.vpce-svc-0b56e75e8cdf50044`.

		 `ap-southeast-2` is `com.amazonaws.vpce.ap-southeast-2.vpce-svc-036f1de74d761706e`.

		 `ap-south-1` is `com.amazonaws.vpce.ap-south-1.vpce-svc-03eb8a541f96d356d`.

		 To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param privateNetworkEndpointIdEntry Private endpoint for Federated Database Instances and Online Archives to add to the specified project.
			@return CreateDataFederationPrivateEndpointApiRequest
	*/
	CreateDataFederationPrivateEndpoint(ctx context.Context, groupId string, privateNetworkEndpointIdEntry *PrivateNetworkEndpointIdEntry) CreateDataFederationPrivateEndpointApiRequest
	/*
		CreateDataFederationPrivateEndpoint Create One Federated Database Instance and Online Archive Private Endpoint for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateDataFederationPrivateEndpointApiParams - Parameters for the request
		@return CreateDataFederationPrivateEndpointApiRequest
	*/
	CreateDataFederationPrivateEndpointWithParams(ctx context.Context, args *CreateDataFederationPrivateEndpointApiParams) CreateDataFederationPrivateEndpointApiRequest

	// Method available only for mocking purposes
	CreateDataFederationPrivateEndpointExecute(r CreateDataFederationPrivateEndpointApiRequest) (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error)

	/*
		CreateFederatedDatabase Create One Federated Database Instance in One Project

		Creates one federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param dataLakeTenant Details to create one federated database instance in the specified project.
		@return CreateFederatedDatabaseApiRequest
	*/
	CreateFederatedDatabase(ctx context.Context, groupId string, dataLakeTenant *DataLakeTenant) CreateFederatedDatabaseApiRequest
	/*
		CreateFederatedDatabase Create One Federated Database Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateFederatedDatabaseApiParams - Parameters for the request
		@return CreateFederatedDatabaseApiRequest
	*/
	CreateFederatedDatabaseWithParams(ctx context.Context, args *CreateFederatedDatabaseApiParams) CreateFederatedDatabaseApiRequest

	// Method available only for mocking purposes
	CreateFederatedDatabaseExecute(r CreateFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error)

	/*
		CreateOneDataFederationQueryLimit Configure One Query Limit for One Federated Database Instance

		Creates or updates one query limit for one federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
		@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
		@param dataFederationTenantQueryLimit Creates or updates one query limit for one federated database instance.
		@return CreateOneDataFederationQueryLimitApiRequest
	*/
	CreateOneDataFederationQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string, dataFederationTenantQueryLimit *DataFederationTenantQueryLimit) CreateOneDataFederationQueryLimitApiRequest
	/*
		CreateOneDataFederationQueryLimit Configure One Query Limit for One Federated Database Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOneDataFederationQueryLimitApiParams - Parameters for the request
		@return CreateOneDataFederationQueryLimitApiRequest
	*/
	CreateOneDataFederationQueryLimitWithParams(ctx context.Context, args *CreateOneDataFederationQueryLimitApiParams) CreateOneDataFederationQueryLimitApiRequest

	// Method available only for mocking purposes
	CreateOneDataFederationQueryLimitExecute(r CreateOneDataFederationQueryLimitApiRequest) (*DataFederationTenantQueryLimit, *http.Response, error)

	/*
		DeleteDataFederationPrivateEndpoint Remove One Federated Database Instance and Online Archive Private Endpoint from One Project

		Removes one private endpoint for Federated Database Instances and Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param endpointId Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
		@return DeleteDataFederationPrivateEndpointApiRequest
	*/
	DeleteDataFederationPrivateEndpoint(ctx context.Context, groupId string, endpointId string) DeleteDataFederationPrivateEndpointApiRequest
	/*
		DeleteDataFederationPrivateEndpoint Remove One Federated Database Instance and Online Archive Private Endpoint from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteDataFederationPrivateEndpointApiParams - Parameters for the request
		@return DeleteDataFederationPrivateEndpointApiRequest
	*/
	DeleteDataFederationPrivateEndpointWithParams(ctx context.Context, args *DeleteDataFederationPrivateEndpointApiParams) DeleteDataFederationPrivateEndpointApiRequest

	// Method available only for mocking purposes
	DeleteDataFederationPrivateEndpointExecute(r DeleteDataFederationPrivateEndpointApiRequest) (*http.Response, error)

	/*
		DeleteFederatedDatabase Remove One Federated Database Instance from One Project

		Removes one federated database instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance to remove.
		@return DeleteFederatedDatabaseApiRequest
	*/
	DeleteFederatedDatabase(ctx context.Context, groupId string, tenantName string) DeleteFederatedDatabaseApiRequest
	/*
		DeleteFederatedDatabase Remove One Federated Database Instance from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteFederatedDatabaseApiParams - Parameters for the request
		@return DeleteFederatedDatabaseApiRequest
	*/
	DeleteFederatedDatabaseWithParams(ctx context.Context, args *DeleteFederatedDatabaseApiParams) DeleteFederatedDatabaseApiRequest

	// Method available only for mocking purposes
	DeleteFederatedDatabaseExecute(r DeleteFederatedDatabaseApiRequest) (*http.Response, error)

	/*
		DeleteOneDataFederationInstanceQueryLimit Delete One Query Limit for One Federated Database Instance

		Deletes one query limit for one federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
		@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
		@return DeleteOneDataFederationInstanceQueryLimitApiRequest
	*/
	DeleteOneDataFederationInstanceQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string) DeleteOneDataFederationInstanceQueryLimitApiRequest
	/*
		DeleteOneDataFederationInstanceQueryLimit Delete One Query Limit for One Federated Database Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOneDataFederationInstanceQueryLimitApiParams - Parameters for the request
		@return DeleteOneDataFederationInstanceQueryLimitApiRequest
	*/
	DeleteOneDataFederationInstanceQueryLimitWithParams(ctx context.Context, args *DeleteOneDataFederationInstanceQueryLimitApiParams) DeleteOneDataFederationInstanceQueryLimitApiRequest

	// Method available only for mocking purposes
	DeleteOneDataFederationInstanceQueryLimitExecute(r DeleteOneDataFederationInstanceQueryLimitApiRequest) (*http.Response, error)

	/*
		DownloadFederatedDatabaseQueryLogs Download Query Logs for One Federated Database Instance

		Downloads the query logs for the specified federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Data Access Read Write roles. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance for which you want to download query logs.
		@return DownloadFederatedDatabaseQueryLogsApiRequest
	*/
	DownloadFederatedDatabaseQueryLogs(ctx context.Context, groupId string, tenantName string) DownloadFederatedDatabaseQueryLogsApiRequest
	/*
		DownloadFederatedDatabaseQueryLogs Download Query Logs for One Federated Database Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadFederatedDatabaseQueryLogsApiParams - Parameters for the request
		@return DownloadFederatedDatabaseQueryLogsApiRequest
	*/
	DownloadFederatedDatabaseQueryLogsWithParams(ctx context.Context, args *DownloadFederatedDatabaseQueryLogsApiParams) DownloadFederatedDatabaseQueryLogsApiRequest

	// Method available only for mocking purposes
	DownloadFederatedDatabaseQueryLogsExecute(r DownloadFederatedDatabaseQueryLogsApiRequest) (io.ReadCloser, *http.Response, error)

	/*
		GetDataFederationPrivateEndpoint Return One Federated Database Instance and Online Archive Private Endpoint in One Project

		Returns the specified private endpoint for Federated Database Instances or Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param endpointId Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
		@return GetDataFederationPrivateEndpointApiRequest
	*/
	GetDataFederationPrivateEndpoint(ctx context.Context, groupId string, endpointId string) GetDataFederationPrivateEndpointApiRequest
	/*
		GetDataFederationPrivateEndpoint Return One Federated Database Instance and Online Archive Private Endpoint in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDataFederationPrivateEndpointApiParams - Parameters for the request
		@return GetDataFederationPrivateEndpointApiRequest
	*/
	GetDataFederationPrivateEndpointWithParams(ctx context.Context, args *GetDataFederationPrivateEndpointApiParams) GetDataFederationPrivateEndpointApiRequest

	// Method available only for mocking purposes
	GetDataFederationPrivateEndpointExecute(r GetDataFederationPrivateEndpointApiRequest) (*PrivateNetworkEndpointIdEntry, *http.Response, error)

	/*
		GetFederatedDatabase Return One Federated Database Instance in One Project

		Returns the details of one federated database instance within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the Federated Database to return.
		@return GetFederatedDatabaseApiRequest
	*/
	GetFederatedDatabase(ctx context.Context, groupId string, tenantName string) GetFederatedDatabaseApiRequest
	/*
		GetFederatedDatabase Return One Federated Database Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetFederatedDatabaseApiParams - Parameters for the request
		@return GetFederatedDatabaseApiRequest
	*/
	GetFederatedDatabaseWithParams(ctx context.Context, args *GetFederatedDatabaseApiParams) GetFederatedDatabaseApiRequest

	// Method available only for mocking purposes
	GetFederatedDatabaseExecute(r GetFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error)

	/*
		ListDataFederationPrivateEndpoints Return All Federated Database Instance and Online Archive Private Endpoints in One Project

		Returns all private endpoints for Federated Database Instances and Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListDataFederationPrivateEndpointsApiRequest
	*/
	ListDataFederationPrivateEndpoints(ctx context.Context, groupId string) ListDataFederationPrivateEndpointsApiRequest
	/*
		ListDataFederationPrivateEndpoints Return All Federated Database Instance and Online Archive Private Endpoints in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDataFederationPrivateEndpointsApiParams - Parameters for the request
		@return ListDataFederationPrivateEndpointsApiRequest
	*/
	ListDataFederationPrivateEndpointsWithParams(ctx context.Context, args *ListDataFederationPrivateEndpointsApiParams) ListDataFederationPrivateEndpointsApiRequest

	// Method available only for mocking purposes
	ListDataFederationPrivateEndpointsExecute(r ListDataFederationPrivateEndpointsApiRequest) (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error)

	/*
		ListFederatedDatabases Return All Federated Database Instances in One Project

		Returns the details of all federated database instances in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or higher role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListFederatedDatabasesApiRequest
	*/
	ListFederatedDatabases(ctx context.Context, groupId string) ListFederatedDatabasesApiRequest
	/*
		ListFederatedDatabases Return All Federated Database Instances in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListFederatedDatabasesApiParams - Parameters for the request
		@return ListFederatedDatabasesApiRequest
	*/
	ListFederatedDatabasesWithParams(ctx context.Context, args *ListFederatedDatabasesApiParams) ListFederatedDatabasesApiRequest

	// Method available only for mocking purposes
	ListFederatedDatabasesExecute(r ListFederatedDatabasesApiRequest) ([]DataLakeTenant, *http.Response, error)

	/*
		ReturnFederatedDatabaseQueryLimit Return One Federated Database Instance Query Limit for One Project

		Returns the details of one query limit for the specified federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
		@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
		@return ReturnFederatedDatabaseQueryLimitApiRequest
	*/
	ReturnFederatedDatabaseQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string) ReturnFederatedDatabaseQueryLimitApiRequest
	/*
		ReturnFederatedDatabaseQueryLimit Return One Federated Database Instance Query Limit for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnFederatedDatabaseQueryLimitApiParams - Parameters for the request
		@return ReturnFederatedDatabaseQueryLimitApiRequest
	*/
	ReturnFederatedDatabaseQueryLimitWithParams(ctx context.Context, args *ReturnFederatedDatabaseQueryLimitApiParams) ReturnFederatedDatabaseQueryLimitApiRequest

	// Method available only for mocking purposes
	ReturnFederatedDatabaseQueryLimitExecute(r ReturnFederatedDatabaseQueryLimitApiRequest) (*DataFederationTenantQueryLimit, *http.Response, error)

	/*
		ReturnFederatedDatabaseQueryLimits Return All Query Limits for One Federated Database Instance

		Returns query limits for a federated databases instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance for which you want to retrieve query limits.
		@return ReturnFederatedDatabaseQueryLimitsApiRequest
	*/
	ReturnFederatedDatabaseQueryLimits(ctx context.Context, groupId string, tenantName string) ReturnFederatedDatabaseQueryLimitsApiRequest
	/*
		ReturnFederatedDatabaseQueryLimits Return All Query Limits for One Federated Database Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnFederatedDatabaseQueryLimitsApiParams - Parameters for the request
		@return ReturnFederatedDatabaseQueryLimitsApiRequest
	*/
	ReturnFederatedDatabaseQueryLimitsWithParams(ctx context.Context, args *ReturnFederatedDatabaseQueryLimitsApiParams) ReturnFederatedDatabaseQueryLimitsApiRequest

	// Method available only for mocking purposes
	ReturnFederatedDatabaseQueryLimitsExecute(r ReturnFederatedDatabaseQueryLimitsApiRequest) ([]DataFederationTenantQueryLimit, *http.Response, error)

	/*
		UpdateFederatedDatabase Update One Federated Database Instance in One Project

		Updates the details of one federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or higher role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the federated database instance to update.
		@param dataLakeTenant Details of one Federated Database to update in the specified project.
		@return UpdateFederatedDatabaseApiRequest
	*/
	UpdateFederatedDatabase(ctx context.Context, groupId string, tenantName string, dataLakeTenant *DataLakeTenant) UpdateFederatedDatabaseApiRequest
	/*
		UpdateFederatedDatabase Update One Federated Database Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateFederatedDatabaseApiParams - Parameters for the request
		@return UpdateFederatedDatabaseApiRequest
	*/
	UpdateFederatedDatabaseWithParams(ctx context.Context, args *UpdateFederatedDatabaseApiParams) UpdateFederatedDatabaseApiRequest

	// Method available only for mocking purposes
	UpdateFederatedDatabaseExecute(r UpdateFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error)
}

// DataFederationApiService DataFederationApi service
type DataFederationApiService service

type CreateDataFederationPrivateEndpointApiRequest struct {
	ctx                           context.Context
	ApiService                    DataFederationApi
	groupId                       string
	privateNetworkEndpointIdEntry *PrivateNetworkEndpointIdEntry
}

type CreateDataFederationPrivateEndpointApiParams struct {
	GroupId                       string
	PrivateNetworkEndpointIdEntry *PrivateNetworkEndpointIdEntry
}

func (a *DataFederationApiService) CreateDataFederationPrivateEndpointWithParams(ctx context.Context, args *CreateDataFederationPrivateEndpointApiParams) CreateDataFederationPrivateEndpointApiRequest {
	return CreateDataFederationPrivateEndpointApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		privateNetworkEndpointIdEntry: args.PrivateNetworkEndpointIdEntry,
	}
}

func (r CreateDataFederationPrivateEndpointApiRequest) Execute() (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error) {
	return r.ApiService.CreateDataFederationPrivateEndpointExecute(r)
}

/*
CreateDataFederationPrivateEndpoint Create One Federated Database Instance and Online Archive Private Endpoint for One Project

Adds one private endpoint for Federated Database Instances and Online Archives to the specified projects. If the endpoint ID already exists and the associated comment is unchanged, Atlas Data Federation makes no change to the endpoint ID list. If the endpoint ID already exists and the associated comment is changed, Atlas Data Federation updates the comment value only in the endpoint ID list. If the endpoint ID doesn't exist, Atlas Data Federation appends the new endpoint to the list of endpoints in the endpoint ID list. Each region has an associated service name for the various endpoints in each region.

	`us-east-1` is `com.amazonaws.vpce.us-east-1.vpce-svc-00e311695874992b4`.

	`us-west-1` is `com.amazonaws.vpce.us-west-2.vpce-svc-09d86b19e59d1b4bb`.

	`eu-west-1` is `com.amazonaws.vpce.eu-west-1.vpce-svc-0824460b72e1a420e`.

	`eu-west-2` is `com.amazonaws.vpce.eu-west-2.vpce-svc-052f1840aa0c4f1f9`.

	`eu-central-1` is `com.amazonaws.vpce.eu-central-1.vpce-svc-0ac8ce91871138c0d`.

	`sa-east-1` is `com.amazonaws.vpce.sa-east-1.vpce-svc-0b56e75e8cdf50044`.

	`ap-southeast-2` is `com.amazonaws.vpce.ap-southeast-2.vpce-svc-036f1de74d761706e`.

	`ap-south-1` is `com.amazonaws.vpce.ap-south-1.vpce-svc-03eb8a541f96d356d`.

	To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateDataFederationPrivateEndpointApiRequest
*/
func (a *DataFederationApiService) CreateDataFederationPrivateEndpoint(ctx context.Context, groupId string, privateNetworkEndpointIdEntry *PrivateNetworkEndpointIdEntry) CreateDataFederationPrivateEndpointApiRequest {
	return CreateDataFederationPrivateEndpointApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		privateNetworkEndpointIdEntry: privateNetworkEndpointIdEntry,
	}
}

// CreateDataFederationPrivateEndpointExecute executes the request
//
//	@return PaginatedPrivateNetworkEndpointIdEntry
func (a *DataFederationApiService) CreateDataFederationPrivateEndpointExecute(r CreateDataFederationPrivateEndpointApiRequest) (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedPrivateNetworkEndpointIdEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.CreateDataFederationPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.privateNetworkEndpointIdEntry == nil {
		return localVarReturnValue, nil, reportError("privateNetworkEndpointIdEntry is required and must be specified")
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
	localVarPostBody = r.privateNetworkEndpointIdEntry
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

type CreateFederatedDatabaseApiRequest struct {
	ctx                context.Context
	ApiService         DataFederationApi
	groupId            string
	dataLakeTenant     *DataLakeTenant
	skipRoleValidation *bool
}

type CreateFederatedDatabaseApiParams struct {
	GroupId            string
	DataLakeTenant     *DataLakeTenant
	SkipRoleValidation *bool
}

func (a *DataFederationApiService) CreateFederatedDatabaseWithParams(ctx context.Context, args *CreateFederatedDatabaseApiParams) CreateFederatedDatabaseApiRequest {
	return CreateFederatedDatabaseApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		dataLakeTenant:     args.DataLakeTenant,
		skipRoleValidation: args.SkipRoleValidation,
	}
}

// Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.
func (r CreateFederatedDatabaseApiRequest) SkipRoleValidation(skipRoleValidation bool) CreateFederatedDatabaseApiRequest {
	r.skipRoleValidation = &skipRoleValidation
	return r
}

func (r CreateFederatedDatabaseApiRequest) Execute() (*DataLakeTenant, *http.Response, error) {
	return r.ApiService.CreateFederatedDatabaseExecute(r)
}

/*
CreateFederatedDatabase Create One Federated Database Instance in One Project

Creates one federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateFederatedDatabaseApiRequest
*/
func (a *DataFederationApiService) CreateFederatedDatabase(ctx context.Context, groupId string, dataLakeTenant *DataLakeTenant) CreateFederatedDatabaseApiRequest {
	return CreateFederatedDatabaseApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		dataLakeTenant: dataLakeTenant,
	}
}

// CreateFederatedDatabaseExecute executes the request
//
//	@return DataLakeTenant
func (a *DataFederationApiService) CreateFederatedDatabaseExecute(r CreateFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataLakeTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.CreateFederatedDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataLakeTenant == nil {
		return localVarReturnValue, nil, reportError("dataLakeTenant is required and must be specified")
	}

	if r.skipRoleValidation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipRoleValidation", r.skipRoleValidation, "")
	} else {
		var defaultValue bool = false
		r.skipRoleValidation = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipRoleValidation", r.skipRoleValidation, "")
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
	localVarPostBody = r.dataLakeTenant
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

type CreateOneDataFederationQueryLimitApiRequest struct {
	ctx                            context.Context
	ApiService                     DataFederationApi
	groupId                        string
	tenantName                     string
	limitName                      string
	dataFederationTenantQueryLimit *DataFederationTenantQueryLimit
}

type CreateOneDataFederationQueryLimitApiParams struct {
	GroupId                        string
	TenantName                     string
	LimitName                      string
	DataFederationTenantQueryLimit *DataFederationTenantQueryLimit
}

func (a *DataFederationApiService) CreateOneDataFederationQueryLimitWithParams(ctx context.Context, args *CreateOneDataFederationQueryLimitApiParams) CreateOneDataFederationQueryLimitApiRequest {
	return CreateOneDataFederationQueryLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		tenantName:                     args.TenantName,
		limitName:                      args.LimitName,
		dataFederationTenantQueryLimit: args.DataFederationTenantQueryLimit,
	}
}

func (r CreateOneDataFederationQueryLimitApiRequest) Execute() (*DataFederationTenantQueryLimit, *http.Response, error) {
	return r.ApiService.CreateOneDataFederationQueryLimitExecute(r)
}

/*
CreateOneDataFederationQueryLimit Configure One Query Limit for One Federated Database Instance

Creates or updates one query limit for one federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
	@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
	@return CreateOneDataFederationQueryLimitApiRequest
*/
func (a *DataFederationApiService) CreateOneDataFederationQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string, dataFederationTenantQueryLimit *DataFederationTenantQueryLimit) CreateOneDataFederationQueryLimitApiRequest {
	return CreateOneDataFederationQueryLimitApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		tenantName:                     tenantName,
		limitName:                      limitName,
		dataFederationTenantQueryLimit: dataFederationTenantQueryLimit,
	}
}

// CreateOneDataFederationQueryLimitExecute executes the request
//
//	@return DataFederationTenantQueryLimit
func (a *DataFederationApiService) CreateOneDataFederationQueryLimitExecute(r CreateOneDataFederationQueryLimitApiRequest) (*DataFederationTenantQueryLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataFederationTenantQueryLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.CreateOneDataFederationQueryLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.limitName == "" {
		return localVarReturnValue, nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataFederationTenantQueryLimit == nil {
		return localVarReturnValue, nil, reportError("dataFederationTenantQueryLimit is required and must be specified")
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
	localVarPostBody = r.dataFederationTenantQueryLimit
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

type DeleteDataFederationPrivateEndpointApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	endpointId string
}

type DeleteDataFederationPrivateEndpointApiParams struct {
	GroupId    string
	EndpointId string
}

func (a *DataFederationApiService) DeleteDataFederationPrivateEndpointWithParams(ctx context.Context, args *DeleteDataFederationPrivateEndpointApiParams) DeleteDataFederationPrivateEndpointApiRequest {
	return DeleteDataFederationPrivateEndpointApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		endpointId: args.EndpointId,
	}
}

func (r DeleteDataFederationPrivateEndpointApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDataFederationPrivateEndpointExecute(r)
}

/*
DeleteDataFederationPrivateEndpoint Remove One Federated Database Instance and Online Archive Private Endpoint from One Project

Removes one private endpoint for Federated Database Instances and Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param endpointId Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
	@return DeleteDataFederationPrivateEndpointApiRequest
*/
func (a *DataFederationApiService) DeleteDataFederationPrivateEndpoint(ctx context.Context, groupId string, endpointId string) DeleteDataFederationPrivateEndpointApiRequest {
	return DeleteDataFederationPrivateEndpointApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		endpointId: endpointId,
	}
}

// DeleteDataFederationPrivateEndpointExecute executes the request
func (a *DataFederationApiService) DeleteDataFederationPrivateEndpointExecute(r DeleteDataFederationPrivateEndpointApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.DeleteDataFederationPrivateEndpoint")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.endpointId == "" {
		return nil, reportError("endpointId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(r.endpointId), -1)

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

type DeleteFederatedDatabaseApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
}

type DeleteFederatedDatabaseApiParams struct {
	GroupId    string
	TenantName string
}

func (a *DataFederationApiService) DeleteFederatedDatabaseWithParams(ctx context.Context, args *DeleteFederatedDatabaseApiParams) DeleteFederatedDatabaseApiRequest {
	return DeleteFederatedDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
	}
}

func (r DeleteFederatedDatabaseApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFederatedDatabaseExecute(r)
}

/*
DeleteFederatedDatabase Remove One Federated Database Instance from One Project

Removes one federated database instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance to remove.
	@return DeleteFederatedDatabaseApiRequest
*/
func (a *DataFederationApiService) DeleteFederatedDatabase(ctx context.Context, groupId string, tenantName string) DeleteFederatedDatabaseApiRequest {
	return DeleteFederatedDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// DeleteFederatedDatabaseExecute executes the request
func (a *DataFederationApiService) DeleteFederatedDatabaseExecute(r DeleteFederatedDatabaseApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.DeleteFederatedDatabase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)

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

type DeleteOneDataFederationInstanceQueryLimitApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
	limitName  string
}

type DeleteOneDataFederationInstanceQueryLimitApiParams struct {
	GroupId    string
	TenantName string
	LimitName  string
}

func (a *DataFederationApiService) DeleteOneDataFederationInstanceQueryLimitWithParams(ctx context.Context, args *DeleteOneDataFederationInstanceQueryLimitApiParams) DeleteOneDataFederationInstanceQueryLimitApiRequest {
	return DeleteOneDataFederationInstanceQueryLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
		limitName:  args.LimitName,
	}
}

func (r DeleteOneDataFederationInstanceQueryLimitApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOneDataFederationInstanceQueryLimitExecute(r)
}

/*
DeleteOneDataFederationInstanceQueryLimit Delete One Query Limit for One Federated Database Instance

Deletes one query limit for one federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
	@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
	@return DeleteOneDataFederationInstanceQueryLimitApiRequest
*/
func (a *DataFederationApiService) DeleteOneDataFederationInstanceQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string) DeleteOneDataFederationInstanceQueryLimitApiRequest {
	return DeleteOneDataFederationInstanceQueryLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
		limitName:  limitName,
	}
}

// DeleteOneDataFederationInstanceQueryLimitExecute executes the request
func (a *DataFederationApiService) DeleteOneDataFederationInstanceQueryLimitExecute(r DeleteOneDataFederationInstanceQueryLimitApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.DeleteOneDataFederationInstanceQueryLimit")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.limitName == "" {
		return nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)

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

type DownloadFederatedDatabaseQueryLogsApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
	endDate    *int64
	startDate  *int64
}

type DownloadFederatedDatabaseQueryLogsApiParams struct {
	GroupId    string
	TenantName string
	EndDate    *int64
	StartDate  *int64
}

func (a *DataFederationApiService) DownloadFederatedDatabaseQueryLogsWithParams(ctx context.Context, args *DownloadFederatedDatabaseQueryLogsApiParams) DownloadFederatedDatabaseQueryLogsApiRequest {
	return DownloadFederatedDatabaseQueryLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
		endDate:    args.EndDate,
		startDate:  args.StartDate,
	}
}

// Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadFederatedDatabaseQueryLogsApiRequest) EndDate(endDate int64) DownloadFederatedDatabaseQueryLogsApiRequest {
	r.endDate = &endDate
	return r
}

// Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadFederatedDatabaseQueryLogsApiRequest) StartDate(startDate int64) DownloadFederatedDatabaseQueryLogsApiRequest {
	r.startDate = &startDate
	return r
}

func (r DownloadFederatedDatabaseQueryLogsApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.DownloadFederatedDatabaseQueryLogsExecute(r)
}

/*
DownloadFederatedDatabaseQueryLogs Download Query Logs for One Federated Database Instance

Downloads the query logs for the specified federated database instance. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Data Access Read Write roles. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance for which you want to download query logs.
	@return DownloadFederatedDatabaseQueryLogsApiRequest
*/
func (a *DataFederationApiService) DownloadFederatedDatabaseQueryLogs(ctx context.Context, groupId string, tenantName string) DownloadFederatedDatabaseQueryLogsApiRequest {
	return DownloadFederatedDatabaseQueryLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// DownloadFederatedDatabaseQueryLogsExecute executes the request
//
//	@return io.ReadCloser
func (a *DataFederationApiService) DownloadFederatedDatabaseQueryLogsExecute(r DownloadFederatedDatabaseQueryLogsApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.DownloadFederatedDatabaseQueryLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+gzip"}

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

type GetDataFederationPrivateEndpointApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	endpointId string
}

type GetDataFederationPrivateEndpointApiParams struct {
	GroupId    string
	EndpointId string
}

func (a *DataFederationApiService) GetDataFederationPrivateEndpointWithParams(ctx context.Context, args *GetDataFederationPrivateEndpointApiParams) GetDataFederationPrivateEndpointApiRequest {
	return GetDataFederationPrivateEndpointApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		endpointId: args.EndpointId,
	}
}

func (r GetDataFederationPrivateEndpointApiRequest) Execute() (*PrivateNetworkEndpointIdEntry, *http.Response, error) {
	return r.ApiService.GetDataFederationPrivateEndpointExecute(r)
}

/*
GetDataFederationPrivateEndpoint Return One Federated Database Instance and Online Archive Private Endpoint in One Project

Returns the specified private endpoint for Federated Database Instances or Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param endpointId Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
	@return GetDataFederationPrivateEndpointApiRequest
*/
func (a *DataFederationApiService) GetDataFederationPrivateEndpoint(ctx context.Context, groupId string, endpointId string) GetDataFederationPrivateEndpointApiRequest {
	return GetDataFederationPrivateEndpointApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		endpointId: endpointId,
	}
}

// GetDataFederationPrivateEndpointExecute executes the request
//
//	@return PrivateNetworkEndpointIdEntry
func (a *DataFederationApiService) GetDataFederationPrivateEndpointExecute(r GetDataFederationPrivateEndpointApiRequest) (*PrivateNetworkEndpointIdEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PrivateNetworkEndpointIdEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.GetDataFederationPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.endpointId == "" {
		return localVarReturnValue, nil, reportError("endpointId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(r.endpointId), -1)

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

type GetFederatedDatabaseApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
}

type GetFederatedDatabaseApiParams struct {
	GroupId    string
	TenantName string
}

func (a *DataFederationApiService) GetFederatedDatabaseWithParams(ctx context.Context, args *GetFederatedDatabaseApiParams) GetFederatedDatabaseApiRequest {
	return GetFederatedDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
	}
}

func (r GetFederatedDatabaseApiRequest) Execute() (*DataLakeTenant, *http.Response, error) {
	return r.ApiService.GetFederatedDatabaseExecute(r)
}

/*
GetFederatedDatabase Return One Federated Database Instance in One Project

Returns the details of one federated database instance within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the Federated Database to return.
	@return GetFederatedDatabaseApiRequest
*/
func (a *DataFederationApiService) GetFederatedDatabase(ctx context.Context, groupId string, tenantName string) GetFederatedDatabaseApiRequest {
	return GetFederatedDatabaseApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// GetFederatedDatabaseExecute executes the request
//
//	@return DataLakeTenant
func (a *DataFederationApiService) GetFederatedDatabaseExecute(r GetFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataLakeTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.GetFederatedDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)

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

type ListDataFederationPrivateEndpointsApiRequest struct {
	ctx          context.Context
	ApiService   DataFederationApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListDataFederationPrivateEndpointsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *DataFederationApiService) ListDataFederationPrivateEndpointsWithParams(ctx context.Context, args *ListDataFederationPrivateEndpointsApiParams) ListDataFederationPrivateEndpointsApiRequest {
	return ListDataFederationPrivateEndpointsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListDataFederationPrivateEndpointsApiRequest) IncludeCount(includeCount bool) ListDataFederationPrivateEndpointsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListDataFederationPrivateEndpointsApiRequest) ItemsPerPage(itemsPerPage int) ListDataFederationPrivateEndpointsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListDataFederationPrivateEndpointsApiRequest) PageNum(pageNum int) ListDataFederationPrivateEndpointsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListDataFederationPrivateEndpointsApiRequest) Execute() (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error) {
	return r.ApiService.ListDataFederationPrivateEndpointsExecute(r)
}

/*
ListDataFederationPrivateEndpoints Return All Federated Database Instance and Online Archive Private Endpoints in One Project

Returns all private endpoints for Federated Database Instances and Online Archives in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListDataFederationPrivateEndpointsApiRequest
*/
func (a *DataFederationApiService) ListDataFederationPrivateEndpoints(ctx context.Context, groupId string) ListDataFederationPrivateEndpointsApiRequest {
	return ListDataFederationPrivateEndpointsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListDataFederationPrivateEndpointsExecute executes the request
//
//	@return PaginatedPrivateNetworkEndpointIdEntry
func (a *DataFederationApiService) ListDataFederationPrivateEndpointsExecute(r ListDataFederationPrivateEndpointsApiRequest) (*PaginatedPrivateNetworkEndpointIdEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedPrivateNetworkEndpointIdEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.ListDataFederationPrivateEndpoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds"
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

type ListFederatedDatabasesApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	type_      *string
}

type ListFederatedDatabasesApiParams struct {
	GroupId string
	Type_   *string
}

func (a *DataFederationApiService) ListFederatedDatabasesWithParams(ctx context.Context, args *ListFederatedDatabasesApiParams) ListFederatedDatabasesApiRequest {
	return ListFederatedDatabasesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		type_:      args.Type_,
	}
}

// Type of Federated Database Instances to return.
func (r ListFederatedDatabasesApiRequest) Type_(type_ string) ListFederatedDatabasesApiRequest {
	r.type_ = &type_
	return r
}

func (r ListFederatedDatabasesApiRequest) Execute() ([]DataLakeTenant, *http.Response, error) {
	return r.ApiService.ListFederatedDatabasesExecute(r)
}

/*
ListFederatedDatabases Return All Federated Database Instances in One Project

Returns the details of all federated database instances in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only or higher role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListFederatedDatabasesApiRequest
*/
func (a *DataFederationApiService) ListFederatedDatabases(ctx context.Context, groupId string) ListFederatedDatabasesApiRequest {
	return ListFederatedDatabasesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListFederatedDatabasesExecute executes the request
//
//	@return []DataLakeTenant
func (a *DataFederationApiService) ListFederatedDatabasesExecute(r ListFederatedDatabasesApiRequest) ([]DataLakeTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []DataLakeTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.ListFederatedDatabases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	} else {
		var defaultValue string = "USER"
		r.type_ = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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

type ReturnFederatedDatabaseQueryLimitApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
	limitName  string
}

type ReturnFederatedDatabaseQueryLimitApiParams struct {
	GroupId    string
	TenantName string
	LimitName  string
}

func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimitWithParams(ctx context.Context, args *ReturnFederatedDatabaseQueryLimitApiParams) ReturnFederatedDatabaseQueryLimitApiRequest {
	return ReturnFederatedDatabaseQueryLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
		limitName:  args.LimitName,
	}
}

func (r ReturnFederatedDatabaseQueryLimitApiRequest) Execute() (*DataFederationTenantQueryLimit, *http.Response, error) {
	return r.ApiService.ReturnFederatedDatabaseQueryLimitExecute(r)
}

/*
ReturnFederatedDatabaseQueryLimit Return One Federated Database Instance Query Limit for One Project

Returns the details of one query limit for the specified federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance to which the query limit applies.
	@param limitName Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
	@return ReturnFederatedDatabaseQueryLimitApiRequest
*/
func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimit(ctx context.Context, groupId string, tenantName string, limitName string) ReturnFederatedDatabaseQueryLimitApiRequest {
	return ReturnFederatedDatabaseQueryLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
		limitName:  limitName,
	}
}

// ReturnFederatedDatabaseQueryLimitExecute executes the request
//
//	@return DataFederationTenantQueryLimit
func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimitExecute(r ReturnFederatedDatabaseQueryLimitApiRequest) (*DataFederationTenantQueryLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataFederationTenantQueryLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.ReturnFederatedDatabaseQueryLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.limitName == "" {
		return localVarReturnValue, nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)

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

type ReturnFederatedDatabaseQueryLimitsApiRequest struct {
	ctx        context.Context
	ApiService DataFederationApi
	groupId    string
	tenantName string
}

type ReturnFederatedDatabaseQueryLimitsApiParams struct {
	GroupId    string
	TenantName string
}

func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimitsWithParams(ctx context.Context, args *ReturnFederatedDatabaseQueryLimitsApiParams) ReturnFederatedDatabaseQueryLimitsApiRequest {
	return ReturnFederatedDatabaseQueryLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
	}
}

func (r ReturnFederatedDatabaseQueryLimitsApiRequest) Execute() ([]DataFederationTenantQueryLimit, *http.Response, error) {
	return r.ApiService.ReturnFederatedDatabaseQueryLimitsExecute(r)
}

/*
ReturnFederatedDatabaseQueryLimits Return All Query Limits for One Federated Database Instance

Returns query limits for a federated databases instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance for which you want to retrieve query limits.
	@return ReturnFederatedDatabaseQueryLimitsApiRequest
*/
func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimits(ctx context.Context, groupId string, tenantName string) ReturnFederatedDatabaseQueryLimitsApiRequest {
	return ReturnFederatedDatabaseQueryLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// ReturnFederatedDatabaseQueryLimitsExecute executes the request
//
//	@return []DataFederationTenantQueryLimit
func (a *DataFederationApiService) ReturnFederatedDatabaseQueryLimitsExecute(r ReturnFederatedDatabaseQueryLimitsApiRequest) ([]DataFederationTenantQueryLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []DataFederationTenantQueryLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.ReturnFederatedDatabaseQueryLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)

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

type UpdateFederatedDatabaseApiRequest struct {
	ctx                context.Context
	ApiService         DataFederationApi
	groupId            string
	tenantName         string
	skipRoleValidation *bool
	dataLakeTenant     *DataLakeTenant
}

type UpdateFederatedDatabaseApiParams struct {
	GroupId            string
	TenantName         string
	SkipRoleValidation *bool
	DataLakeTenant     *DataLakeTenant
}

func (a *DataFederationApiService) UpdateFederatedDatabaseWithParams(ctx context.Context, args *UpdateFederatedDatabaseApiParams) UpdateFederatedDatabaseApiRequest {
	return UpdateFederatedDatabaseApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		tenantName:         args.TenantName,
		skipRoleValidation: args.SkipRoleValidation,
		dataLakeTenant:     args.DataLakeTenant,
	}
}

// Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.
func (r UpdateFederatedDatabaseApiRequest) SkipRoleValidation(skipRoleValidation bool) UpdateFederatedDatabaseApiRequest {
	r.skipRoleValidation = &skipRoleValidation
	return r
}

func (r UpdateFederatedDatabaseApiRequest) Execute() (*DataLakeTenant, *http.Response, error) {
	return r.ApiService.UpdateFederatedDatabaseExecute(r)
}

/*
UpdateFederatedDatabase Update One Federated Database Instance in One Project

Updates the details of one federated database instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or higher role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the federated database instance to update.
	@return UpdateFederatedDatabaseApiRequest
*/
func (a *DataFederationApiService) UpdateFederatedDatabase(ctx context.Context, groupId string, tenantName string, dataLakeTenant *DataLakeTenant) UpdateFederatedDatabaseApiRequest {
	return UpdateFederatedDatabaseApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		tenantName:     tenantName,
		dataLakeTenant: dataLakeTenant,
	}
}

// UpdateFederatedDatabaseExecute executes the request
//
//	@return DataLakeTenant
func (a *DataFederationApiService) UpdateFederatedDatabaseExecute(r UpdateFederatedDatabaseApiRequest) (*DataLakeTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataLakeTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataFederationApiService.UpdateFederatedDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.skipRoleValidation == nil {
		return localVarReturnValue, nil, reportError("skipRoleValidation is required and must be specified")
	}
	if r.dataLakeTenant == nil {
		return localVarReturnValue, nil, reportError("dataLakeTenant is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "skipRoleValidation", r.skipRoleValidation, "")
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
	localVarPostBody = r.dataLakeTenant
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
