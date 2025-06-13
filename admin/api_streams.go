// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StreamsApi interface {

	/*
		AcceptVpcPeeringConnection Accept One Incoming VPC Peering Connection

		Requests the acceptance of an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@param vPCPeeringActionChallenge Challenge values for VPC Peering requester account ID, and requester VPC ID.
		@return AcceptVpcPeeringConnectionApiRequest
	*/
	AcceptVpcPeeringConnection(ctx context.Context, groupId string, id string, vPCPeeringActionChallenge *VPCPeeringActionChallenge) AcceptVpcPeeringConnectionApiRequest
	/*
		AcceptVpcPeeringConnection Accept One Incoming VPC Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AcceptVpcPeeringConnectionApiParams - Parameters for the request
		@return AcceptVpcPeeringConnectionApiRequest
	*/
	AcceptVpcPeeringConnectionWithParams(ctx context.Context, args *AcceptVpcPeeringConnectionApiParams) AcceptVpcPeeringConnectionApiRequest

	// Method available only for mocking purposes
	AcceptVpcPeeringConnectionExecute(r AcceptVpcPeeringConnectionApiRequest) (*http.Response, error)

	/*
		CreatePrivateLinkConnection Create One Private Link Connection

		Creates one Private Link in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param streamsPrivateLinkConnection Details to create one Private Link connection for a project. project.
		@return CreatePrivateLinkConnectionApiRequest
	*/
	CreatePrivateLinkConnection(ctx context.Context, groupId string, streamsPrivateLinkConnection *StreamsPrivateLinkConnection) CreatePrivateLinkConnectionApiRequest
	/*
		CreatePrivateLinkConnection Create One Private Link Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePrivateLinkConnectionApiParams - Parameters for the request
		@return CreatePrivateLinkConnectionApiRequest
	*/
	CreatePrivateLinkConnectionWithParams(ctx context.Context, args *CreatePrivateLinkConnectionApiParams) CreatePrivateLinkConnectionApiRequest

	// Method available only for mocking purposes
	CreatePrivateLinkConnectionExecute(r CreatePrivateLinkConnectionApiRequest) (*StreamsPrivateLinkConnection, *http.Response, error)

	/*
		CreateStreamConnection Create One Stream Connection

		Creates one connection for a stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param streamsConnection Details to create one connection for a streams instance in the specified project.
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnection(ctx context.Context, groupId string, tenantName string, streamsConnection *StreamsConnection) CreateStreamConnectionApiRequest
	/*
		CreateStreamConnection Create One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamConnectionApiParams - Parameters for the request
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnectionWithParams(ctx context.Context, args *CreateStreamConnectionApiParams) CreateStreamConnectionApiRequest

	// Method available only for mocking purposes
	CreateStreamConnectionExecute(r CreateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		CreateStreamInstance Create One Stream Instance

		Creates one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param streamsTenant Details to create one streams instance in the specified project.
		@return CreateStreamInstanceApiRequest
	*/
	CreateStreamInstance(ctx context.Context, groupId string, streamsTenant *StreamsTenant) CreateStreamInstanceApiRequest
	/*
		CreateStreamInstance Create One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamInstanceApiParams - Parameters for the request
		@return CreateStreamInstanceApiRequest
	*/
	CreateStreamInstanceWithParams(ctx context.Context, args *CreateStreamInstanceApiParams) CreateStreamInstanceApiRequest

	// Method available only for mocking purposes
	CreateStreamInstanceExecute(r CreateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		CreateStreamInstanceWithSampleConnections Create One Stream Instance with Sample Connections

		Creates one stream instance in the specified project with sample connections. To use this resource the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param body Details to create one streams instance in the specified project.
		@return CreateStreamInstanceWithSampleConnectionsApiRequest
	*/
	CreateStreamInstanceWithSampleConnections(ctx context.Context, groupId string, body *any) CreateStreamInstanceWithSampleConnectionsApiRequest
	/*
		CreateStreamInstanceWithSampleConnections Create One Stream Instance with Sample Connections


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamInstanceWithSampleConnectionsApiParams - Parameters for the request
		@return CreateStreamInstanceWithSampleConnectionsApiRequest
	*/
	CreateStreamInstanceWithSampleConnectionsWithParams(ctx context.Context, args *CreateStreamInstanceWithSampleConnectionsApiParams) CreateStreamInstanceWithSampleConnectionsApiRequest

	// Method available only for mocking purposes
	CreateStreamInstanceWithSampleConnectionsExecute(r CreateStreamInstanceWithSampleConnectionsApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		CreateStreamProcessor Create One Stream Processor

		Create one Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param streamsProcessor Details to create an Atlas Streams Processor.
		@return CreateStreamProcessorApiRequest
	*/
	CreateStreamProcessor(ctx context.Context, groupId string, tenantName string, streamsProcessor *StreamsProcessor) CreateStreamProcessorApiRequest
	/*
		CreateStreamProcessor Create One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamProcessorApiParams - Parameters for the request
		@return CreateStreamProcessorApiRequest
	*/
	CreateStreamProcessorWithParams(ctx context.Context, args *CreateStreamProcessorApiParams) CreateStreamProcessorApiRequest

	// Method available only for mocking purposes
	CreateStreamProcessorExecute(r CreateStreamProcessorApiRequest) (*StreamsProcessor, *http.Response, error)

	/*
		DeletePrivateLinkConnection Delete One Private Link Connection

		Deletes one Private Link in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param connectionId Unique ID that identifies the Private Link connection.
		@return DeletePrivateLinkConnectionApiRequest
	*/
	DeletePrivateLinkConnection(ctx context.Context, groupId string, connectionId string) DeletePrivateLinkConnectionApiRequest
	/*
		DeletePrivateLinkConnection Delete One Private Link Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePrivateLinkConnectionApiParams - Parameters for the request
		@return DeletePrivateLinkConnectionApiRequest
	*/
	DeletePrivateLinkConnectionWithParams(ctx context.Context, args *DeletePrivateLinkConnectionApiParams) DeletePrivateLinkConnectionApiRequest

	// Method available only for mocking purposes
	DeletePrivateLinkConnectionExecute(r DeletePrivateLinkConnectionApiRequest) (*http.Response, error)

	/*
		DeleteStreamConnection Delete One Stream Connection

		Delete one connection of the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param connectionName Human-readable label that identifies the stream connection.
		@return DeleteStreamConnectionApiRequest
	*/
	DeleteStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) DeleteStreamConnectionApiRequest
	/*
		DeleteStreamConnection Delete One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteStreamConnectionApiParams - Parameters for the request
		@return DeleteStreamConnectionApiRequest
	*/
	DeleteStreamConnectionWithParams(ctx context.Context, args *DeleteStreamConnectionApiParams) DeleteStreamConnectionApiRequest

	// Method available only for mocking purposes
	DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (*http.Response, error)

	/*
		DeleteStreamInstance Delete One Stream Instance

		Delete one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to delete.
		@return DeleteStreamInstanceApiRequest
	*/
	DeleteStreamInstance(ctx context.Context, groupId string, tenantName string) DeleteStreamInstanceApiRequest
	/*
		DeleteStreamInstance Delete One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteStreamInstanceApiParams - Parameters for the request
		@return DeleteStreamInstanceApiRequest
	*/
	DeleteStreamInstanceWithParams(ctx context.Context, args *DeleteStreamInstanceApiParams) DeleteStreamInstanceApiRequest

	// Method available only for mocking purposes
	DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (*http.Response, error)

	/*
		DeleteStreamProcessor Delete One Stream Processor

		Delete a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@return DeleteStreamProcessorApiRequest
	*/
	DeleteStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) DeleteStreamProcessorApiRequest
	/*
		DeleteStreamProcessor Delete One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteStreamProcessorApiParams - Parameters for the request
		@return DeleteStreamProcessorApiRequest
	*/
	DeleteStreamProcessorWithParams(ctx context.Context, args *DeleteStreamProcessorApiParams) DeleteStreamProcessorApiRequest

	// Method available only for mocking purposes
	DeleteStreamProcessorExecute(r DeleteStreamProcessorApiRequest) (*http.Response, error)

	/*
		DeleteVpcPeeringConnection Delete One VPC Peering Connection

		Deletes an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@return DeleteVpcPeeringConnectionApiRequest
	*/
	DeleteVpcPeeringConnection(ctx context.Context, groupId string, id string) DeleteVpcPeeringConnectionApiRequest
	/*
		DeleteVpcPeeringConnection Delete One VPC Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteVpcPeeringConnectionApiParams - Parameters for the request
		@return DeleteVpcPeeringConnectionApiRequest
	*/
	DeleteVpcPeeringConnectionWithParams(ctx context.Context, args *DeleteVpcPeeringConnectionApiParams) DeleteVpcPeeringConnectionApiRequest

	// Method available only for mocking purposes
	DeleteVpcPeeringConnectionExecute(r DeleteVpcPeeringConnectionApiRequest) (*http.Response, error)

	/*
		DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance

		Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return DownloadStreamTenantAuditLogsApiRequest
	*/
	DownloadStreamTenantAuditLogs(ctx context.Context, groupId string, tenantName string) DownloadStreamTenantAuditLogsApiRequest
	/*
		DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadStreamTenantAuditLogsApiParams - Parameters for the request
		@return DownloadStreamTenantAuditLogsApiRequest
	*/
	DownloadStreamTenantAuditLogsWithParams(ctx context.Context, args *DownloadStreamTenantAuditLogsApiParams) DownloadStreamTenantAuditLogsApiRequest

	// Method available only for mocking purposes
	DownloadStreamTenantAuditLogsExecute(r DownloadStreamTenantAuditLogsApiRequest) (io.ReadCloser, *http.Response, error)

	/*
		GetAccountDetails Return Account ID and VPC ID for One Project and Region

		Returns the Account ID, and the VPC ID for the group and region specified.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetAccountDetailsApiRequest
	*/
	GetAccountDetails(ctx context.Context, groupId string) GetAccountDetailsApiRequest
	/*
		GetAccountDetails Return Account ID and VPC ID for One Project and Region


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAccountDetailsApiParams - Parameters for the request
		@return GetAccountDetailsApiRequest
	*/
	GetAccountDetailsWithParams(ctx context.Context, args *GetAccountDetailsApiParams) GetAccountDetailsApiRequest

	// Method available only for mocking purposes
	GetAccountDetailsExecute(r GetAccountDetailsApiRequest) (*AccountDetails, *http.Response, error)

	/*
		GetActiveVpcPeeringConnections Return All Active Incoming VPC Peering Connections

		Returns a list of active incoming VPC Peering Connections.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetActiveVpcPeeringConnectionsApiRequest
	*/
	GetActiveVpcPeeringConnections(ctx context.Context, groupId string) GetActiveVpcPeeringConnectionsApiRequest
	/*
		GetActiveVpcPeeringConnections Return All Active Incoming VPC Peering Connections


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetActiveVpcPeeringConnectionsApiParams - Parameters for the request
		@return GetActiveVpcPeeringConnectionsApiRequest
	*/
	GetActiveVpcPeeringConnectionsWithParams(ctx context.Context, args *GetActiveVpcPeeringConnectionsApiParams) GetActiveVpcPeeringConnectionsApiRequest

	// Method available only for mocking purposes
	GetActiveVpcPeeringConnectionsExecute(r GetActiveVpcPeeringConnectionsApiRequest) (*http.Response, error)

	/*
		GetPrivateLinkConnection Return One Private Link Connection

		Returns the details of one Private Link connection within the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param connectionId Unique ID that identifies the Private Link connection.
		@return GetPrivateLinkConnectionApiRequest
	*/
	GetPrivateLinkConnection(ctx context.Context, groupId string, connectionId string) GetPrivateLinkConnectionApiRequest
	/*
		GetPrivateLinkConnection Return One Private Link Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPrivateLinkConnectionApiParams - Parameters for the request
		@return GetPrivateLinkConnectionApiRequest
	*/
	GetPrivateLinkConnectionWithParams(ctx context.Context, args *GetPrivateLinkConnectionApiParams) GetPrivateLinkConnectionApiRequest

	// Method available only for mocking purposes
	GetPrivateLinkConnectionExecute(r GetPrivateLinkConnectionApiRequest) (*StreamsPrivateLinkConnection, *http.Response, error)

	/*
		GetStreamConnection Return One Stream Connection

		Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to return.
		@param connectionName Human-readable label that identifies the stream connection to return.
		@return GetStreamConnectionApiRequest
	*/
	GetStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) GetStreamConnectionApiRequest
	/*
		GetStreamConnection Return One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamConnectionApiParams - Parameters for the request
		@return GetStreamConnectionApiRequest
	*/
	GetStreamConnectionWithParams(ctx context.Context, args *GetStreamConnectionApiParams) GetStreamConnectionApiRequest

	// Method available only for mocking purposes
	GetStreamConnectionExecute(r GetStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		GetStreamInstance Return One Stream Instance

		Returns the details of one stream instance within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to return.
		@return GetStreamInstanceApiRequest
	*/
	GetStreamInstance(ctx context.Context, groupId string, tenantName string) GetStreamInstanceApiRequest
	/*
		GetStreamInstance Return One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamInstanceApiParams - Parameters for the request
		@return GetStreamInstanceApiRequest
	*/
	GetStreamInstanceWithParams(ctx context.Context, args *GetStreamInstanceApiParams) GetStreamInstanceApiRequest

	// Method available only for mocking purposes
	GetStreamInstanceExecute(r GetStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		GetStreamProcessor Return One Stream Processor

		Get one Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@return GetStreamProcessorApiRequest
	*/
	GetStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) GetStreamProcessorApiRequest
	/*
		GetStreamProcessor Return One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamProcessorApiParams - Parameters for the request
		@return GetStreamProcessorApiRequest
	*/
	GetStreamProcessorWithParams(ctx context.Context, args *GetStreamProcessorApiParams) GetStreamProcessorApiRequest

	// Method available only for mocking purposes
	GetStreamProcessorExecute(r GetStreamProcessorApiRequest) (*StreamsProcessorWithStats, *http.Response, error)

	/*
		GetVpcPeeringConnections Return All VPC Peering Connections

		Returns a list of incoming VPC Peering Connections.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetVpcPeeringConnectionsApiRequest
	*/
	GetVpcPeeringConnections(ctx context.Context, groupId string) GetVpcPeeringConnectionsApiRequest
	/*
		GetVpcPeeringConnections Return All VPC Peering Connections


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetVpcPeeringConnectionsApiParams - Parameters for the request
		@return GetVpcPeeringConnectionsApiRequest
	*/
	GetVpcPeeringConnectionsWithParams(ctx context.Context, args *GetVpcPeeringConnectionsApiParams) GetVpcPeeringConnectionsApiRequest

	// Method available only for mocking purposes
	GetVpcPeeringConnectionsExecute(r GetVpcPeeringConnectionsApiRequest) (*http.Response, error)

	/*
		ListPrivateLinkConnections Return All Private Link Connections

		Returns all Private Link connections for the specified project.To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListPrivateLinkConnectionsApiRequest
	*/
	ListPrivateLinkConnections(ctx context.Context, groupId string) ListPrivateLinkConnectionsApiRequest
	/*
		ListPrivateLinkConnections Return All Private Link Connections


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPrivateLinkConnectionsApiParams - Parameters for the request
		@return ListPrivateLinkConnectionsApiRequest
	*/
	ListPrivateLinkConnectionsWithParams(ctx context.Context, args *ListPrivateLinkConnectionsApiParams) ListPrivateLinkConnectionsApiRequest

	// Method available only for mocking purposes
	ListPrivateLinkConnectionsExecute(r ListPrivateLinkConnectionsApiRequest) (*PaginatedApiStreamsPrivateLink, *http.Response, error)

	/*
		ListStreamConnections Return All Connections of the Stream Instances

		Returns all connections of the stream instance for the specified project.To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnections(ctx context.Context, groupId string, tenantName string) ListStreamConnectionsApiRequest
	/*
		ListStreamConnections Return All Connections of the Stream Instances


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamConnectionsApiParams - Parameters for the request
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnectionsWithParams(ctx context.Context, args *ListStreamConnectionsApiParams) ListStreamConnectionsApiRequest

	// Method available only for mocking purposes
	ListStreamConnectionsExecute(r ListStreamConnectionsApiRequest) (*PaginatedApiStreamsConnection, *http.Response, error)

	/*
		ListStreamInstances Return All Stream Instances in One Project

		Returns all stream instances for the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstances(ctx context.Context, groupId string) ListStreamInstancesApiRequest
	/*
		ListStreamInstances Return All Stream Instances in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamInstancesApiParams - Parameters for the request
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstancesWithParams(ctx context.Context, args *ListStreamInstancesApiParams) ListStreamInstancesApiRequest

	// Method available only for mocking purposes
	ListStreamInstancesExecute(r ListStreamInstancesApiRequest) (*PaginatedApiStreamsTenant, *http.Response, error)

	/*
		ListStreamProcessors Return All Stream Processors in One Stream Instance

		Returns all Stream Processors within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return ListStreamProcessorsApiRequest
	*/
	ListStreamProcessors(ctx context.Context, groupId string, tenantName string) ListStreamProcessorsApiRequest
	/*
		ListStreamProcessors Return All Stream Processors in One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamProcessorsApiParams - Parameters for the request
		@return ListStreamProcessorsApiRequest
	*/
	ListStreamProcessorsWithParams(ctx context.Context, args *ListStreamProcessorsApiParams) ListStreamProcessorsApiRequest

	// Method available only for mocking purposes
	ListStreamProcessorsExecute(r ListStreamProcessorsApiRequest) (*PaginatedApiStreamsStreamProcessorWithStats, *http.Response, error)

	/*
		ModifyStreamProcessor Update One Stream Processor

		Modify one existing Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@param streamsModifyStreamProcessor Modifications to apply to the stream processor.
		@return ModifyStreamProcessorApiRequest
	*/
	ModifyStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string, streamsModifyStreamProcessor *StreamsModifyStreamProcessor) ModifyStreamProcessorApiRequest
	/*
		ModifyStreamProcessor Update One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ModifyStreamProcessorApiParams - Parameters for the request
		@return ModifyStreamProcessorApiRequest
	*/
	ModifyStreamProcessorWithParams(ctx context.Context, args *ModifyStreamProcessorApiParams) ModifyStreamProcessorApiRequest

	// Method available only for mocking purposes
	ModifyStreamProcessorExecute(r ModifyStreamProcessorApiRequest) (*StreamsProcessorWithStats, *http.Response, error)

	/*
		RejectVpcPeeringConnection Reject One Incoming VPC Peering Connection

		Requests the rejection of an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@return RejectVpcPeeringConnectionApiRequest
	*/
	RejectVpcPeeringConnection(ctx context.Context, groupId string, id string) RejectVpcPeeringConnectionApiRequest
	/*
		RejectVpcPeeringConnection Reject One Incoming VPC Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RejectVpcPeeringConnectionApiParams - Parameters for the request
		@return RejectVpcPeeringConnectionApiRequest
	*/
	RejectVpcPeeringConnectionWithParams(ctx context.Context, args *RejectVpcPeeringConnectionApiParams) RejectVpcPeeringConnectionApiRequest

	// Method available only for mocking purposes
	RejectVpcPeeringConnectionExecute(r RejectVpcPeeringConnectionApiRequest) (*http.Response, error)

	/*
		StartStreamProcessor Start One Stream Processor

		Start a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@return StartStreamProcessorApiRequest
	*/
	StartStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) StartStreamProcessorApiRequest
	/*
		StartStreamProcessor Start One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param StartStreamProcessorApiParams - Parameters for the request
		@return StartStreamProcessorApiRequest
	*/
	StartStreamProcessorWithParams(ctx context.Context, args *StartStreamProcessorApiParams) StartStreamProcessorApiRequest

	// Method available only for mocking purposes
	StartStreamProcessorExecute(r StartStreamProcessorApiRequest) (*http.Response, error)

	/*
		StartStreamProcessorWith Start One Stream Processor With Options

		Start a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@param streamsStartStreamProcessorWith Options for starting a stream processor.
		@return StartStreamProcessorWithApiRequest
	*/
	StartStreamProcessorWith(ctx context.Context, groupId string, tenantName string, processorName string, streamsStartStreamProcessorWith *StreamsStartStreamProcessorWith) StartStreamProcessorWithApiRequest
	/*
		StartStreamProcessorWith Start One Stream Processor With Options


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param StartStreamProcessorWithApiParams - Parameters for the request
		@return StartStreamProcessorWithApiRequest
	*/
	StartStreamProcessorWithWithParams(ctx context.Context, args *StartStreamProcessorWithApiParams) StartStreamProcessorWithApiRequest

	// Method available only for mocking purposes
	StartStreamProcessorWithExecute(r StartStreamProcessorWithApiRequest) (*http.Response, error)

	/*
		StopStreamProcessor Stop One Stream Processor

		Stop a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@return StopStreamProcessorApiRequest
	*/
	StopStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) StopStreamProcessorApiRequest
	/*
		StopStreamProcessor Stop One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param StopStreamProcessorApiParams - Parameters for the request
		@return StopStreamProcessorApiRequest
	*/
	StopStreamProcessorWithParams(ctx context.Context, args *StopStreamProcessorApiParams) StopStreamProcessorApiRequest

	// Method available only for mocking purposes
	StopStreamProcessorExecute(r StopStreamProcessorApiRequest) (*http.Response, error)

	/*
		UpdateStreamConnection Update One Stream Connection

		Update one connection for the specified stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param connectionName Human-readable label that identifies the stream connection.
		@param streamsConnection Details to update one connection for a streams instance in the specified project.
		@return UpdateStreamConnectionApiRequest
	*/
	UpdateStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string, streamsConnection *StreamsConnection) UpdateStreamConnectionApiRequest
	/*
		UpdateStreamConnection Update One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateStreamConnectionApiParams - Parameters for the request
		@return UpdateStreamConnectionApiRequest
	*/
	UpdateStreamConnectionWithParams(ctx context.Context, args *UpdateStreamConnectionApiParams) UpdateStreamConnectionApiRequest

	// Method available only for mocking purposes
	UpdateStreamConnectionExecute(r UpdateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		UpdateStreamInstance Update One Stream Instance

		Update one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to update.
		@param streamsDataProcessRegion Details of the new data process region to update in the streams instance.
		@return UpdateStreamInstanceApiRequest
	*/
	UpdateStreamInstance(ctx context.Context, groupId string, tenantName string, streamsDataProcessRegion *StreamsDataProcessRegion) UpdateStreamInstanceApiRequest
	/*
		UpdateStreamInstance Update One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateStreamInstanceApiParams - Parameters for the request
		@return UpdateStreamInstanceApiRequest
	*/
	UpdateStreamInstanceWithParams(ctx context.Context, args *UpdateStreamInstanceApiParams) UpdateStreamInstanceApiRequest

	// Method available only for mocking purposes
	UpdateStreamInstanceExecute(r UpdateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)
}

// StreamsApiService StreamsApi service
type StreamsApiService service

type AcceptVpcPeeringConnectionApiRequest struct {
	ctx                       context.Context
	ApiService                StreamsApi
	groupId                   string
	id                        string
	vPCPeeringActionChallenge *VPCPeeringActionChallenge
}

type AcceptVpcPeeringConnectionApiParams struct {
	GroupId                   string
	Id                        string
	VPCPeeringActionChallenge *VPCPeeringActionChallenge
}

func (a *StreamsApiService) AcceptVpcPeeringConnectionWithParams(ctx context.Context, args *AcceptVpcPeeringConnectionApiParams) AcceptVpcPeeringConnectionApiRequest {
	return AcceptVpcPeeringConnectionApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		id:                        args.Id,
		vPCPeeringActionChallenge: args.VPCPeeringActionChallenge,
	}
}

func (r AcceptVpcPeeringConnectionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcceptVpcPeeringConnectionExecute(r)
}

/*
AcceptVpcPeeringConnection Accept One Incoming VPC Peering Connection

Requests the acceptance of an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return AcceptVpcPeeringConnectionApiRequest
*/
func (a *StreamsApiService) AcceptVpcPeeringConnection(ctx context.Context, groupId string, id string, vPCPeeringActionChallenge *VPCPeeringActionChallenge) AcceptVpcPeeringConnectionApiRequest {
	return AcceptVpcPeeringConnectionApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		id:                        id,
		vPCPeeringActionChallenge: vPCPeeringActionChallenge,
	}
}

// AcceptVpcPeeringConnectionExecute executes the request
func (a *StreamsApiService) AcceptVpcPeeringConnectionExecute(r AcceptVpcPeeringConnectionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.AcceptVpcPeeringConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.id == "" {
		return nil, reportError("id is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(r.id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.vPCPeeringActionChallenge == nil {
		return nil, reportError("vPCPeeringActionChallenge is required and must be specified")
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
	localVarPostBody = r.vPCPeeringActionChallenge
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

type CreatePrivateLinkConnectionApiRequest struct {
	ctx                          context.Context
	ApiService                   StreamsApi
	groupId                      string
	streamsPrivateLinkConnection *StreamsPrivateLinkConnection
}

type CreatePrivateLinkConnectionApiParams struct {
	GroupId                      string
	StreamsPrivateLinkConnection *StreamsPrivateLinkConnection
}

func (a *StreamsApiService) CreatePrivateLinkConnectionWithParams(ctx context.Context, args *CreatePrivateLinkConnectionApiParams) CreatePrivateLinkConnectionApiRequest {
	return CreatePrivateLinkConnectionApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      args.GroupId,
		streamsPrivateLinkConnection: args.StreamsPrivateLinkConnection,
	}
}

func (r CreatePrivateLinkConnectionApiRequest) Execute() (*StreamsPrivateLinkConnection, *http.Response, error) {
	return r.ApiService.CreatePrivateLinkConnectionExecute(r)
}

/*
CreatePrivateLinkConnection Create One Private Link Connection

Creates one Private Link in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePrivateLinkConnectionApiRequest
*/
func (a *StreamsApiService) CreatePrivateLinkConnection(ctx context.Context, groupId string, streamsPrivateLinkConnection *StreamsPrivateLinkConnection) CreatePrivateLinkConnectionApiRequest {
	return CreatePrivateLinkConnectionApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      groupId,
		streamsPrivateLinkConnection: streamsPrivateLinkConnection,
	}
}

// CreatePrivateLinkConnectionExecute executes the request
//
//	@return StreamsPrivateLinkConnection
func (a *StreamsApiService) CreatePrivateLinkConnectionExecute(r CreatePrivateLinkConnectionApiRequest) (*StreamsPrivateLinkConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsPrivateLinkConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreatePrivateLinkConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/privateLinkConnections"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsPrivateLinkConnection == nil {
		return localVarReturnValue, nil, reportError("streamsPrivateLinkConnection is required and must be specified")
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
	localVarPostBody = r.streamsPrivateLinkConnection
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

type CreateStreamConnectionApiRequest struct {
	ctx               context.Context
	ApiService        StreamsApi
	groupId           string
	tenantName        string
	streamsConnection *StreamsConnection
}

type CreateStreamConnectionApiParams struct {
	GroupId           string
	TenantName        string
	StreamsConnection *StreamsConnection
}

func (a *StreamsApiService) CreateStreamConnectionWithParams(ctx context.Context, args *CreateStreamConnectionApiParams) CreateStreamConnectionApiRequest {
	return CreateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		tenantName:        args.TenantName,
		streamsConnection: args.StreamsConnection,
	}
}

func (r CreateStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.CreateStreamConnectionExecute(r)
}

/*
CreateStreamConnection Create One Stream Connection

Creates one connection for a stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return CreateStreamConnectionApiRequest
*/
func (a *StreamsApiService) CreateStreamConnection(ctx context.Context, groupId string, tenantName string, streamsConnection *StreamsConnection) CreateStreamConnectionApiRequest {
	return CreateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		tenantName:        tenantName,
		streamsConnection: streamsConnection,
	}
}

// CreateStreamConnectionExecute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) CreateStreamConnectionExecute(r CreateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections"
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
	if r.streamsConnection == nil {
		return localVarReturnValue, nil, reportError("streamsConnection is required and must be specified")
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
	localVarPostBody = r.streamsConnection
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

type CreateStreamInstanceApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	streamsTenant *StreamsTenant
}

type CreateStreamInstanceApiParams struct {
	GroupId       string
	StreamsTenant *StreamsTenant
}

func (a *StreamsApiService) CreateStreamInstanceWithParams(ctx context.Context, args *CreateStreamInstanceApiParams) CreateStreamInstanceApiRequest {
	return CreateStreamInstanceApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		streamsTenant: args.StreamsTenant,
	}
}

func (r CreateStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.CreateStreamInstanceExecute(r)
}

/*
CreateStreamInstance Create One Stream Instance

Creates one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateStreamInstanceApiRequest
*/
func (a *StreamsApiService) CreateStreamInstance(ctx context.Context, groupId string, streamsTenant *StreamsTenant) CreateStreamInstanceApiRequest {
	return CreateStreamInstanceApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		streamsTenant: streamsTenant,
	}
}

// CreateStreamInstanceExecute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) CreateStreamInstanceExecute(r CreateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsTenant == nil {
		return localVarReturnValue, nil, reportError("streamsTenant is required and must be specified")
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
	localVarPostBody = r.streamsTenant
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

type CreateStreamInstanceWithSampleConnectionsApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	body       *any
}

type CreateStreamInstanceWithSampleConnectionsApiParams struct {
	GroupId string
	Body    *any
}

func (a *StreamsApiService) CreateStreamInstanceWithSampleConnectionsWithParams(ctx context.Context, args *CreateStreamInstanceWithSampleConnectionsApiParams) CreateStreamInstanceWithSampleConnectionsApiRequest {
	return CreateStreamInstanceWithSampleConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		body:       args.Body,
	}
}

func (r CreateStreamInstanceWithSampleConnectionsApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.CreateStreamInstanceWithSampleConnectionsExecute(r)
}

/*
CreateStreamInstanceWithSampleConnections Create One Stream Instance with Sample Connections

Creates one stream instance in the specified project with sample connections. To use this resource the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateStreamInstanceWithSampleConnectionsApiRequest
*/
func (a *StreamsApiService) CreateStreamInstanceWithSampleConnections(ctx context.Context, groupId string, body *any) CreateStreamInstanceWithSampleConnectionsApiRequest {
	return CreateStreamInstanceWithSampleConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		body:       body,
	}
}

// CreateStreamInstanceWithSampleConnectionsExecute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) CreateStreamInstanceWithSampleConnectionsExecute(r CreateStreamInstanceWithSampleConnectionsApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamInstanceWithSampleConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams:withSampleConnections"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type CreateStreamProcessorApiRequest struct {
	ctx              context.Context
	ApiService       StreamsApi
	groupId          string
	tenantName       string
	streamsProcessor *StreamsProcessor
}

type CreateStreamProcessorApiParams struct {
	GroupId          string
	TenantName       string
	StreamsProcessor *StreamsProcessor
}

func (a *StreamsApiService) CreateStreamProcessorWithParams(ctx context.Context, args *CreateStreamProcessorApiParams) CreateStreamProcessorApiRequest {
	return CreateStreamProcessorApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		tenantName:       args.TenantName,
		streamsProcessor: args.StreamsProcessor,
	}
}

func (r CreateStreamProcessorApiRequest) Execute() (*StreamsProcessor, *http.Response, error) {
	return r.ApiService.CreateStreamProcessorExecute(r)
}

/*
CreateStreamProcessor Create One Stream Processor

Create one Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return CreateStreamProcessorApiRequest
*/
func (a *StreamsApiService) CreateStreamProcessor(ctx context.Context, groupId string, tenantName string, streamsProcessor *StreamsProcessor) CreateStreamProcessorApiRequest {
	return CreateStreamProcessorApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		tenantName:       tenantName,
		streamsProcessor: streamsProcessor,
	}
}

// CreateStreamProcessorExecute executes the request
//
//	@return StreamsProcessor
func (a *StreamsApiService) CreateStreamProcessorExecute(r CreateStreamProcessorApiRequest) (*StreamsProcessor, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsProcessor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor"
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
	if r.streamsProcessor == nil {
		return localVarReturnValue, nil, reportError("streamsProcessor is required and must be specified")
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
	localVarPostBody = r.streamsProcessor
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

type DeletePrivateLinkConnectionApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	connectionId string
}

type DeletePrivateLinkConnectionApiParams struct {
	GroupId      string
	ConnectionId string
}

func (a *StreamsApiService) DeletePrivateLinkConnectionWithParams(ctx context.Context, args *DeletePrivateLinkConnectionApiParams) DeletePrivateLinkConnectionApiRequest {
	return DeletePrivateLinkConnectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		connectionId: args.ConnectionId,
	}
}

func (r DeletePrivateLinkConnectionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePrivateLinkConnectionExecute(r)
}

/*
DeletePrivateLinkConnection Delete One Private Link Connection

Deletes one Private Link in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param connectionId Unique ID that identifies the Private Link connection.
	@return DeletePrivateLinkConnectionApiRequest
*/
func (a *StreamsApiService) DeletePrivateLinkConnection(ctx context.Context, groupId string, connectionId string) DeletePrivateLinkConnectionApiRequest {
	return DeletePrivateLinkConnectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		connectionId: connectionId,
	}
}

// DeletePrivateLinkConnectionExecute executes the request
func (a *StreamsApiService) DeletePrivateLinkConnectionExecute(r DeletePrivateLinkConnectionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeletePrivateLinkConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.connectionId == "" {
		return nil, reportError("connectionId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(r.connectionId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type DeleteStreamConnectionApiRequest struct {
	ctx            context.Context
	ApiService     StreamsApi
	groupId        string
	tenantName     string
	connectionName string
}

type DeleteStreamConnectionApiParams struct {
	GroupId        string
	TenantName     string
	ConnectionName string
}

func (a *StreamsApiService) DeleteStreamConnectionWithParams(ctx context.Context, args *DeleteStreamConnectionApiParams) DeleteStreamConnectionApiRequest {
	return DeleteStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		tenantName:     args.TenantName,
		connectionName: args.ConnectionName,
	}
}

func (r DeleteStreamConnectionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamConnectionExecute(r)
}

/*
DeleteStreamConnection Delete One Stream Connection

Delete one connection of the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param connectionName Human-readable label that identifies the stream connection.
	@return DeleteStreamConnectionApiRequest
*/
func (a *StreamsApiService) DeleteStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) DeleteStreamConnectionApiRequest {
	return DeleteStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		tenantName:     tenantName,
		connectionName: connectionName,
	}
}

// DeleteStreamConnectionExecute executes the request
func (a *StreamsApiService) DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.connectionName == "" {
		return nil, reportError("connectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(r.connectionName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type DeleteStreamInstanceApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	tenantName string
}

type DeleteStreamInstanceApiParams struct {
	GroupId    string
	TenantName string
}

func (a *StreamsApiService) DeleteStreamInstanceWithParams(ctx context.Context, args *DeleteStreamInstanceApiParams) DeleteStreamInstanceApiRequest {
	return DeleteStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
	}
}

func (r DeleteStreamInstanceApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamInstanceExecute(r)
}

/*
DeleteStreamInstance Delete One Stream Instance

Delete one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to delete.
	@return DeleteStreamInstanceApiRequest
*/
func (a *StreamsApiService) DeleteStreamInstance(ctx context.Context, groupId string, tenantName string) DeleteStreamInstanceApiRequest {
	return DeleteStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// DeleteStreamInstanceExecute executes the request
func (a *StreamsApiService) DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type DeleteStreamProcessorApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	tenantName    string
	processorName string
}

type DeleteStreamProcessorApiParams struct {
	GroupId       string
	TenantName    string
	ProcessorName string
}

func (a *StreamsApiService) DeleteStreamProcessorWithParams(ctx context.Context, args *DeleteStreamProcessorApiParams) DeleteStreamProcessorApiRequest {
	return DeleteStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		tenantName:    args.TenantName,
		processorName: args.ProcessorName,
	}
}

func (r DeleteStreamProcessorApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteStreamProcessorExecute(r)
}

/*
DeleteStreamProcessor Delete One Stream Processor

Delete a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return DeleteStreamProcessorApiRequest
*/
func (a *StreamsApiService) DeleteStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) DeleteStreamProcessorApiRequest {
	return DeleteStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		tenantName:    tenantName,
		processorName: processorName,
	}
}

// DeleteStreamProcessorExecute executes the request
func (a *StreamsApiService) DeleteStreamProcessorExecute(r DeleteStreamProcessorApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamProcessor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

type DeleteVpcPeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	id         string
}

type DeleteVpcPeeringConnectionApiParams struct {
	GroupId string
	Id      string
}

func (a *StreamsApiService) DeleteVpcPeeringConnectionWithParams(ctx context.Context, args *DeleteVpcPeeringConnectionApiParams) DeleteVpcPeeringConnectionApiRequest {
	return DeleteVpcPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		id:         args.Id,
	}
}

func (r DeleteVpcPeeringConnectionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVpcPeeringConnectionExecute(r)
}

/*
DeleteVpcPeeringConnection Delete One VPC Peering Connection

Deletes an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return DeleteVpcPeeringConnectionApiRequest
*/
func (a *StreamsApiService) DeleteVpcPeeringConnection(ctx context.Context, groupId string, id string) DeleteVpcPeeringConnectionApiRequest {
	return DeleteVpcPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		id:         id,
	}
}

// DeleteVpcPeeringConnectionExecute executes the request
func (a *StreamsApiService) DeleteVpcPeeringConnectionExecute(r DeleteVpcPeeringConnectionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteVpcPeeringConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.id == "" {
		return nil, reportError("id is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(r.id), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type DownloadStreamTenantAuditLogsApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	tenantName string
	endDate    *int64
	startDate  *int64
}

type DownloadStreamTenantAuditLogsApiParams struct {
	GroupId    string
	TenantName string
	EndDate    *int64
	StartDate  *int64
}

func (a *StreamsApiService) DownloadStreamTenantAuditLogsWithParams(ctx context.Context, args *DownloadStreamTenantAuditLogsApiParams) DownloadStreamTenantAuditLogsApiRequest {
	return DownloadStreamTenantAuditLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
		endDate:    args.EndDate,
		startDate:  args.StartDate,
	}
}

// Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadStreamTenantAuditLogsApiRequest) EndDate(endDate int64) DownloadStreamTenantAuditLogsApiRequest {
	r.endDate = &endDate
	return r
}

// Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadStreamTenantAuditLogsApiRequest) StartDate(startDate int64) DownloadStreamTenantAuditLogsApiRequest {
	r.startDate = &startDate
	return r
}

func (r DownloadStreamTenantAuditLogsApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.DownloadStreamTenantAuditLogsExecute(r)
}

/*
DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance

Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return DownloadStreamTenantAuditLogsApiRequest
*/
func (a *StreamsApiService) DownloadStreamTenantAuditLogs(ctx context.Context, groupId string, tenantName string) DownloadStreamTenantAuditLogsApiRequest {
	return DownloadStreamTenantAuditLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// DownloadStreamTenantAuditLogsExecute executes the request
//
//	@return io.ReadCloser
func (a *StreamsApiService) DownloadStreamTenantAuditLogsExecute(r DownloadStreamTenantAuditLogsApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DownloadStreamTenantAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+gzip"}

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

type GetAccountDetailsApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	cloudProvider *string
	regionName    *string
}

type GetAccountDetailsApiParams struct {
	GroupId       string
	CloudProvider *string
	RegionName    *string
}

func (a *StreamsApiService) GetAccountDetailsWithParams(ctx context.Context, args *GetAccountDetailsApiParams) GetAccountDetailsApiRequest {
	return GetAccountDetailsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		cloudProvider: args.CloudProvider,
		regionName:    args.RegionName,
	}
}

// One of \&quot;aws\&quot;, \&quot;azure\&quot; or \&quot;gcp\&quot;.
func (r GetAccountDetailsApiRequest) CloudProvider(cloudProvider string) GetAccountDetailsApiRequest {
	r.cloudProvider = &cloudProvider
	return r
}

// The cloud provider specific region name, i.e. \&quot;US_EAST_1\&quot; for cloud provider \&quot;aws\&quot;.
func (r GetAccountDetailsApiRequest) RegionName(regionName string) GetAccountDetailsApiRequest {
	r.regionName = &regionName
	return r
}

func (r GetAccountDetailsApiRequest) Execute() (*AccountDetails, *http.Response, error) {
	return r.ApiService.GetAccountDetailsExecute(r)
}

/*
GetAccountDetails Return Account ID and VPC ID for One Project and Region

Returns the Account ID, and the VPC ID for the group and region specified.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetAccountDetailsApiRequest
*/
func (a *StreamsApiService) GetAccountDetails(ctx context.Context, groupId string) GetAccountDetailsApiRequest {
	return GetAccountDetailsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetAccountDetailsExecute executes the request
//
//	@return AccountDetails
func (a *StreamsApiService) GetAccountDetailsExecute(r GetAccountDetailsApiRequest) (*AccountDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AccountDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetAccountDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/accountDetails"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProvider == nil {
		return localVarReturnValue, nil, reportError("cloudProvider is required and must be specified")
	}
	if r.regionName == nil {
		return localVarReturnValue, nil, reportError("regionName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cloudProvider", r.cloudProvider, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "regionName", r.regionName, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-11-13+json"}

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

type GetActiveVpcPeeringConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type GetActiveVpcPeeringConnectionsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) GetActiveVpcPeeringConnectionsWithParams(ctx context.Context, args *GetActiveVpcPeeringConnectionsApiParams) GetActiveVpcPeeringConnectionsApiRequest {
	return GetActiveVpcPeeringConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r GetActiveVpcPeeringConnectionsApiRequest) ItemsPerPage(itemsPerPage int) GetActiveVpcPeeringConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetActiveVpcPeeringConnectionsApiRequest) PageNum(pageNum int) GetActiveVpcPeeringConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetActiveVpcPeeringConnectionsApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetActiveVpcPeeringConnectionsExecute(r)
}

/*
GetActiveVpcPeeringConnections Return All Active Incoming VPC Peering Connections

Returns a list of active incoming VPC Peering Connections.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetActiveVpcPeeringConnectionsApiRequest
*/
func (a *StreamsApiService) GetActiveVpcPeeringConnections(ctx context.Context, groupId string) GetActiveVpcPeeringConnectionsApiRequest {
	return GetActiveVpcPeeringConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetActiveVpcPeeringConnectionsExecute executes the request
func (a *StreamsApiService) GetActiveVpcPeeringConnectionsExecute(r GetActiveVpcPeeringConnectionsApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetActiveVpcPeeringConnections")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/activeVpcPeeringConnections"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-11-13+json"}

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

type GetPrivateLinkConnectionApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	connectionId string
}

type GetPrivateLinkConnectionApiParams struct {
	GroupId      string
	ConnectionId string
}

func (a *StreamsApiService) GetPrivateLinkConnectionWithParams(ctx context.Context, args *GetPrivateLinkConnectionApiParams) GetPrivateLinkConnectionApiRequest {
	return GetPrivateLinkConnectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		connectionId: args.ConnectionId,
	}
}

func (r GetPrivateLinkConnectionApiRequest) Execute() (*StreamsPrivateLinkConnection, *http.Response, error) {
	return r.ApiService.GetPrivateLinkConnectionExecute(r)
}

/*
GetPrivateLinkConnection Return One Private Link Connection

Returns the details of one Private Link connection within the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param connectionId Unique ID that identifies the Private Link connection.
	@return GetPrivateLinkConnectionApiRequest
*/
func (a *StreamsApiService) GetPrivateLinkConnection(ctx context.Context, groupId string, connectionId string) GetPrivateLinkConnectionApiRequest {
	return GetPrivateLinkConnectionApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		connectionId: connectionId,
	}
}

// GetPrivateLinkConnectionExecute executes the request
//
//	@return StreamsPrivateLinkConnection
func (a *StreamsApiService) GetPrivateLinkConnectionExecute(r GetPrivateLinkConnectionApiRequest) (*StreamsPrivateLinkConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsPrivateLinkConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetPrivateLinkConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.connectionId == "" {
		return localVarReturnValue, nil, reportError("connectionId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(r.connectionId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type GetStreamConnectionApiRequest struct {
	ctx            context.Context
	ApiService     StreamsApi
	groupId        string
	tenantName     string
	connectionName string
}

type GetStreamConnectionApiParams struct {
	GroupId        string
	TenantName     string
	ConnectionName string
}

func (a *StreamsApiService) GetStreamConnectionWithParams(ctx context.Context, args *GetStreamConnectionApiParams) GetStreamConnectionApiRequest {
	return GetStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		tenantName:     args.TenantName,
		connectionName: args.ConnectionName,
	}
}

func (r GetStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.GetStreamConnectionExecute(r)
}

/*
GetStreamConnection Return One Stream Connection

Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to return.
	@param connectionName Human-readable label that identifies the stream connection to return.
	@return GetStreamConnectionApiRequest
*/
func (a *StreamsApiService) GetStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) GetStreamConnectionApiRequest {
	return GetStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		tenantName:     tenantName,
		connectionName: connectionName,
	}
}

// GetStreamConnectionExecute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) GetStreamConnectionExecute(r GetStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.connectionName == "" {
		return localVarReturnValue, nil, reportError("connectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(r.connectionName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type GetStreamInstanceApiRequest struct {
	ctx                context.Context
	ApiService         StreamsApi
	groupId            string
	tenantName         string
	includeConnections *bool
}

type GetStreamInstanceApiParams struct {
	GroupId            string
	TenantName         string
	IncludeConnections *bool
}

func (a *StreamsApiService) GetStreamInstanceWithParams(ctx context.Context, args *GetStreamInstanceApiParams) GetStreamInstanceApiRequest {
	return GetStreamInstanceApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		tenantName:         args.TenantName,
		includeConnections: args.IncludeConnections,
	}
}

// Flag to indicate whether connections information should be included in the stream instance.
func (r GetStreamInstanceApiRequest) IncludeConnections(includeConnections bool) GetStreamInstanceApiRequest {
	r.includeConnections = &includeConnections
	return r
}

func (r GetStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.GetStreamInstanceExecute(r)
}

/*
GetStreamInstance Return One Stream Instance

Returns the details of one stream instance within the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to return.
	@return GetStreamInstanceApiRequest
*/
func (a *StreamsApiService) GetStreamInstance(ctx context.Context, groupId string, tenantName string) GetStreamInstanceApiRequest {
	return GetStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// GetStreamInstanceExecute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) GetStreamInstanceExecute(r GetStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
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

	if r.includeConnections != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeConnections", r.includeConnections, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type GetStreamProcessorApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	tenantName    string
	processorName string
}

type GetStreamProcessorApiParams struct {
	GroupId       string
	TenantName    string
	ProcessorName string
}

func (a *StreamsApiService) GetStreamProcessorWithParams(ctx context.Context, args *GetStreamProcessorApiParams) GetStreamProcessorApiRequest {
	return GetStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		tenantName:    args.TenantName,
		processorName: args.ProcessorName,
	}
}

func (r GetStreamProcessorApiRequest) Execute() (*StreamsProcessorWithStats, *http.Response, error) {
	return r.ApiService.GetStreamProcessorExecute(r)
}

/*
GetStreamProcessor Return One Stream Processor

Get one Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return GetStreamProcessorApiRequest
*/
func (a *StreamsApiService) GetStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) GetStreamProcessorApiRequest {
	return GetStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		tenantName:    tenantName,
		processorName: processorName,
	}
}

// GetStreamProcessorExecute executes the request
//
//	@return StreamsProcessorWithStats
func (a *StreamsApiService) GetStreamProcessorExecute(r GetStreamProcessorApiRequest) (*StreamsProcessorWithStats, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsProcessorWithStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetStreamProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return localVarReturnValue, nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

type GetVpcPeeringConnectionsApiRequest struct {
	ctx                context.Context
	ApiService         StreamsApi
	requesterAccountId *string
	groupId            string
	itemsPerPage       *int
	pageNum            *int
}

type GetVpcPeeringConnectionsApiParams struct {
	RequesterAccountId *string
	GroupId            string
	ItemsPerPage       *int
	PageNum            *int
}

func (a *StreamsApiService) GetVpcPeeringConnectionsWithParams(ctx context.Context, args *GetVpcPeeringConnectionsApiParams) GetVpcPeeringConnectionsApiRequest {
	return GetVpcPeeringConnectionsApiRequest{
		ApiService:         a,
		ctx:                ctx,
		requesterAccountId: args.RequesterAccountId,
		groupId:            args.GroupId,
		itemsPerPage:       args.ItemsPerPage,
		pageNum:            args.PageNum,
	}
}

// The Account ID of the VPC Peering connection/s.
func (r GetVpcPeeringConnectionsApiRequest) RequesterAccountId(requesterAccountId string) GetVpcPeeringConnectionsApiRequest {
	r.requesterAccountId = &requesterAccountId
	return r
}

// Number of items that the response returns per page.
func (r GetVpcPeeringConnectionsApiRequest) ItemsPerPage(itemsPerPage int) GetVpcPeeringConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetVpcPeeringConnectionsApiRequest) PageNum(pageNum int) GetVpcPeeringConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetVpcPeeringConnectionsApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetVpcPeeringConnectionsExecute(r)
}

/*
GetVpcPeeringConnections Return All VPC Peering Connections

Returns a list of incoming VPC Peering Connections.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetVpcPeeringConnectionsApiRequest
*/
func (a *StreamsApiService) GetVpcPeeringConnections(ctx context.Context, groupId string) GetVpcPeeringConnectionsApiRequest {
	return GetVpcPeeringConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetVpcPeeringConnectionsExecute executes the request
func (a *StreamsApiService) GetVpcPeeringConnectionsExecute(r GetVpcPeeringConnectionsApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetVpcPeeringConnections")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requesterAccountId == nil {
		return nil, reportError("requesterAccountId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "requesterAccountId", r.requesterAccountId, "")
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type ListPrivateLinkConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListPrivateLinkConnectionsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) ListPrivateLinkConnectionsWithParams(ctx context.Context, args *ListPrivateLinkConnectionsApiParams) ListPrivateLinkConnectionsApiRequest {
	return ListPrivateLinkConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListPrivateLinkConnectionsApiRequest) ItemsPerPage(itemsPerPage int) ListPrivateLinkConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPrivateLinkConnectionsApiRequest) PageNum(pageNum int) ListPrivateLinkConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListPrivateLinkConnectionsApiRequest) Execute() (*PaginatedApiStreamsPrivateLink, *http.Response, error) {
	return r.ApiService.ListPrivateLinkConnectionsExecute(r)
}

/*
ListPrivateLinkConnections Return All Private Link Connections

Returns all Private Link connections for the specified project.To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListPrivateLinkConnectionsApiRequest
*/
func (a *StreamsApiService) ListPrivateLinkConnections(ctx context.Context, groupId string) ListPrivateLinkConnectionsApiRequest {
	return ListPrivateLinkConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListPrivateLinkConnectionsExecute executes the request
//
//	@return PaginatedApiStreamsPrivateLink
func (a *StreamsApiService) ListPrivateLinkConnectionsExecute(r ListPrivateLinkConnectionsApiRequest) (*PaginatedApiStreamsPrivateLink, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsPrivateLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListPrivateLinkConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/privateLinkConnections"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type ListStreamConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	tenantName   string
	itemsPerPage *int
	pageNum      *int
}

type ListStreamConnectionsApiParams struct {
	GroupId      string
	TenantName   string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) ListStreamConnectionsWithParams(ctx context.Context, args *ListStreamConnectionsApiParams) ListStreamConnectionsApiRequest {
	return ListStreamConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		tenantName:   args.TenantName,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListStreamConnectionsApiRequest) ItemsPerPage(itemsPerPage int) ListStreamConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStreamConnectionsApiRequest) PageNum(pageNum int) ListStreamConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListStreamConnectionsApiRequest) Execute() (*PaginatedApiStreamsConnection, *http.Response, error) {
	return r.ApiService.ListStreamConnectionsExecute(r)
}

/*
ListStreamConnections Return All Connections of the Stream Instances

Returns all connections of the stream instance for the specified project.To use this resource, the requesting Service Account or API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return ListStreamConnectionsApiRequest
*/
func (a *StreamsApiService) ListStreamConnections(ctx context.Context, groupId string, tenantName string) ListStreamConnectionsApiRequest {
	return ListStreamConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// ListStreamConnectionsExecute executes the request
//
//	@return PaginatedApiStreamsConnection
func (a *StreamsApiService) ListStreamConnectionsExecute(r ListStreamConnectionsApiRequest) (*PaginatedApiStreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListStreamConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type ListStreamInstancesApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListStreamInstancesApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) ListStreamInstancesWithParams(ctx context.Context, args *ListStreamInstancesApiParams) ListStreamInstancesApiRequest {
	return ListStreamInstancesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListStreamInstancesApiRequest) ItemsPerPage(itemsPerPage int) ListStreamInstancesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStreamInstancesApiRequest) PageNum(pageNum int) ListStreamInstancesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListStreamInstancesApiRequest) Execute() (*PaginatedApiStreamsTenant, *http.Response, error) {
	return r.ApiService.ListStreamInstancesExecute(r)
}

/*
ListStreamInstances Return All Stream Instances in One Project

Returns all stream instances for the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListStreamInstancesApiRequest
*/
func (a *StreamsApiService) ListStreamInstances(ctx context.Context, groupId string) ListStreamInstancesApiRequest {
	return ListStreamInstancesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListStreamInstancesExecute executes the request
//
//	@return PaginatedApiStreamsTenant
func (a *StreamsApiService) ListStreamInstancesExecute(r ListStreamInstancesApiRequest) (*PaginatedApiStreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListStreamInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type ListStreamProcessorsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	tenantName   string
	itemsPerPage *int
	pageNum      *int
	includeCount *bool
}

type ListStreamProcessorsApiParams struct {
	GroupId      string
	TenantName   string
	ItemsPerPage *int
	PageNum      *int
	IncludeCount *bool
}

func (a *StreamsApiService) ListStreamProcessorsWithParams(ctx context.Context, args *ListStreamProcessorsApiParams) ListStreamProcessorsApiRequest {
	return ListStreamProcessorsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		tenantName:   args.TenantName,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		includeCount: args.IncludeCount,
	}
}

// Number of items that the response returns per page.
func (r ListStreamProcessorsApiRequest) ItemsPerPage(itemsPerPage int) ListStreamProcessorsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStreamProcessorsApiRequest) PageNum(pageNum int) ListStreamProcessorsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListStreamProcessorsApiRequest) IncludeCount(includeCount bool) ListStreamProcessorsApiRequest {
	r.includeCount = &includeCount
	return r
}

func (r ListStreamProcessorsApiRequest) Execute() (*PaginatedApiStreamsStreamProcessorWithStats, *http.Response, error) {
	return r.ApiService.ListStreamProcessorsExecute(r)
}

/*
ListStreamProcessors Return All Stream Processors in One Stream Instance

Returns all Stream Processors within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return ListStreamProcessorsApiRequest
*/
func (a *StreamsApiService) ListStreamProcessors(ctx context.Context, groupId string, tenantName string) ListStreamProcessorsApiRequest {
	return ListStreamProcessorsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// ListStreamProcessorsExecute executes the request
//
//	@return PaginatedApiStreamsStreamProcessorWithStats
func (a *StreamsApiService) ListStreamProcessorsExecute(r ListStreamProcessorsApiRequest) (*PaginatedApiStreamsStreamProcessorWithStats, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsStreamProcessorWithStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListStreamProcessors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processors"
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
	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ModifyStreamProcessorApiRequest struct {
	ctx                          context.Context
	ApiService                   StreamsApi
	groupId                      string
	tenantName                   string
	processorName                string
	streamsModifyStreamProcessor *StreamsModifyStreamProcessor
}

type ModifyStreamProcessorApiParams struct {
	GroupId                      string
	TenantName                   string
	ProcessorName                string
	StreamsModifyStreamProcessor *StreamsModifyStreamProcessor
}

func (a *StreamsApiService) ModifyStreamProcessorWithParams(ctx context.Context, args *ModifyStreamProcessorApiParams) ModifyStreamProcessorApiRequest {
	return ModifyStreamProcessorApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      args.GroupId,
		tenantName:                   args.TenantName,
		processorName:                args.ProcessorName,
		streamsModifyStreamProcessor: args.StreamsModifyStreamProcessor,
	}
}

func (r ModifyStreamProcessorApiRequest) Execute() (*StreamsProcessorWithStats, *http.Response, error) {
	return r.ApiService.ModifyStreamProcessorExecute(r)
}

/*
ModifyStreamProcessor Update One Stream Processor

Modify one existing Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return ModifyStreamProcessorApiRequest
*/
func (a *StreamsApiService) ModifyStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string, streamsModifyStreamProcessor *StreamsModifyStreamProcessor) ModifyStreamProcessorApiRequest {
	return ModifyStreamProcessorApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      groupId,
		tenantName:                   tenantName,
		processorName:                processorName,
		streamsModifyStreamProcessor: streamsModifyStreamProcessor,
	}
}

// ModifyStreamProcessorExecute executes the request
//
//	@return StreamsProcessorWithStats
func (a *StreamsApiService) ModifyStreamProcessorExecute(r ModifyStreamProcessorApiRequest) (*StreamsProcessorWithStats, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsProcessorWithStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ModifyStreamProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return localVarReturnValue, nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsModifyStreamProcessor == nil {
		return localVarReturnValue, nil, reportError("streamsModifyStreamProcessor is required and must be specified")
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
	localVarPostBody = r.streamsModifyStreamProcessor
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

type RejectVpcPeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	id         string
}

type RejectVpcPeeringConnectionApiParams struct {
	GroupId string
	Id      string
}

func (a *StreamsApiService) RejectVpcPeeringConnectionWithParams(ctx context.Context, args *RejectVpcPeeringConnectionApiParams) RejectVpcPeeringConnectionApiRequest {
	return RejectVpcPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		id:         args.Id,
	}
}

func (r RejectVpcPeeringConnectionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RejectVpcPeeringConnectionExecute(r)
}

/*
RejectVpcPeeringConnection Reject One Incoming VPC Peering Connection

Requests the rejection of an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return RejectVpcPeeringConnectionApiRequest
*/
func (a *StreamsApiService) RejectVpcPeeringConnection(ctx context.Context, groupId string, id string) RejectVpcPeeringConnectionApiRequest {
	return RejectVpcPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		id:         id,
	}
}

// RejectVpcPeeringConnectionExecute executes the request
func (a *StreamsApiService) RejectVpcPeeringConnectionExecute(r RejectVpcPeeringConnectionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.RejectVpcPeeringConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.id == "" {
		return nil, reportError("id is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(r.id), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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

type StartStreamProcessorApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	tenantName    string
	processorName string
}

type StartStreamProcessorApiParams struct {
	GroupId       string
	TenantName    string
	ProcessorName string
}

func (a *StreamsApiService) StartStreamProcessorWithParams(ctx context.Context, args *StartStreamProcessorApiParams) StartStreamProcessorApiRequest {
	return StartStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		tenantName:    args.TenantName,
		processorName: args.ProcessorName,
	}
}

func (r StartStreamProcessorApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartStreamProcessorExecute(r)
}

/*
StartStreamProcessor Start One Stream Processor

Start a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return StartStreamProcessorApiRequest
*/
func (a *StreamsApiService) StartStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) StartStreamProcessorApiRequest {
	return StartStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		tenantName:    tenantName,
		processorName: processorName,
	}
}

// StartStreamProcessorExecute executes the request
func (a *StreamsApiService) StartStreamProcessorExecute(r StartStreamProcessorApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.StartStreamProcessor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

type StartStreamProcessorWithApiRequest struct {
	ctx                             context.Context
	ApiService                      StreamsApi
	groupId                         string
	tenantName                      string
	processorName                   string
	streamsStartStreamProcessorWith *StreamsStartStreamProcessorWith
}

type StartStreamProcessorWithApiParams struct {
	GroupId                         string
	TenantName                      string
	ProcessorName                   string
	StreamsStartStreamProcessorWith *StreamsStartStreamProcessorWith
}

func (a *StreamsApiService) StartStreamProcessorWithWithParams(ctx context.Context, args *StartStreamProcessorWithApiParams) StartStreamProcessorWithApiRequest {
	return StartStreamProcessorWithApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         args.GroupId,
		tenantName:                      args.TenantName,
		processorName:                   args.ProcessorName,
		streamsStartStreamProcessorWith: args.StreamsStartStreamProcessorWith,
	}
}

func (r StartStreamProcessorWithApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartStreamProcessorWithExecute(r)
}

/*
StartStreamProcessorWith Start One Stream Processor With Options

Start a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return StartStreamProcessorWithApiRequest
*/
func (a *StreamsApiService) StartStreamProcessorWith(ctx context.Context, groupId string, tenantName string, processorName string, streamsStartStreamProcessorWith *StreamsStartStreamProcessorWith) StartStreamProcessorWithApiRequest {
	return StartStreamProcessorWithApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         groupId,
		tenantName:                      tenantName,
		processorName:                   processorName,
		streamsStartStreamProcessorWith: streamsStartStreamProcessorWith,
	}
}

// StartStreamProcessorWithExecute executes the request
func (a *StreamsApiService) StartStreamProcessorWithExecute(r StartStreamProcessorWithApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.StartStreamProcessorWith")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:startWith"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.streamsStartStreamProcessorWith
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

type StopStreamProcessorApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	tenantName    string
	processorName string
}

type StopStreamProcessorApiParams struct {
	GroupId       string
	TenantName    string
	ProcessorName string
}

func (a *StreamsApiService) StopStreamProcessorWithParams(ctx context.Context, args *StopStreamProcessorApiParams) StopStreamProcessorApiRequest {
	return StopStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		tenantName:    args.TenantName,
		processorName: args.ProcessorName,
	}
}

func (r StopStreamProcessorApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.StopStreamProcessorExecute(r)
}

/*
StopStreamProcessor Stop One Stream Processor

Stop a Stream Processor within the specified stream instance. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param processorName Human-readable label that identifies the stream processor.
	@return StopStreamProcessorApiRequest
*/
func (a *StreamsApiService) StopStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) StopStreamProcessorApiRequest {
	return StopStreamProcessorApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		tenantName:    tenantName,
		processorName: processorName,
	}
}

// StopStreamProcessorExecute executes the request
func (a *StreamsApiService) StopStreamProcessorExecute(r StopStreamProcessorApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.StopStreamProcessor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.processorName == "" {
		return nil, reportError("processorName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(r.processorName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

type UpdateStreamConnectionApiRequest struct {
	ctx               context.Context
	ApiService        StreamsApi
	groupId           string
	tenantName        string
	connectionName    string
	streamsConnection *StreamsConnection
}

type UpdateStreamConnectionApiParams struct {
	GroupId           string
	TenantName        string
	ConnectionName    string
	StreamsConnection *StreamsConnection
}

func (a *StreamsApiService) UpdateStreamConnectionWithParams(ctx context.Context, args *UpdateStreamConnectionApiParams) UpdateStreamConnectionApiRequest {
	return UpdateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		tenantName:        args.TenantName,
		connectionName:    args.ConnectionName,
		streamsConnection: args.StreamsConnection,
	}
}

func (r UpdateStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.UpdateStreamConnectionExecute(r)
}

/*
UpdateStreamConnection Update One Stream Connection

Update one connection for the specified stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param connectionName Human-readable label that identifies the stream connection.
	@return UpdateStreamConnectionApiRequest
*/
func (a *StreamsApiService) UpdateStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string, streamsConnection *StreamsConnection) UpdateStreamConnectionApiRequest {
	return UpdateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		tenantName:        tenantName,
		connectionName:    connectionName,
		streamsConnection: streamsConnection,
	}
}

// UpdateStreamConnectionExecute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) UpdateStreamConnectionExecute(r UpdateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.UpdateStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.tenantName == "" {
		return localVarReturnValue, nil, reportError("tenantName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(r.tenantName), -1)
	if r.connectionName == "" {
		return localVarReturnValue, nil, reportError("connectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(r.connectionName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsConnection == nil {
		return localVarReturnValue, nil, reportError("streamsConnection is required and must be specified")
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
	localVarPostBody = r.streamsConnection
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

type UpdateStreamInstanceApiRequest struct {
	ctx                      context.Context
	ApiService               StreamsApi
	groupId                  string
	tenantName               string
	streamsDataProcessRegion *StreamsDataProcessRegion
}

type UpdateStreamInstanceApiParams struct {
	GroupId                  string
	TenantName               string
	StreamsDataProcessRegion *StreamsDataProcessRegion
}

func (a *StreamsApiService) UpdateStreamInstanceWithParams(ctx context.Context, args *UpdateStreamInstanceApiParams) UpdateStreamInstanceApiRequest {
	return UpdateStreamInstanceApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		tenantName:               args.TenantName,
		streamsDataProcessRegion: args.StreamsDataProcessRegion,
	}
}

func (r UpdateStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.UpdateStreamInstanceExecute(r)
}

/*
UpdateStreamInstance Update One Stream Instance

Update one stream instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to update.
	@return UpdateStreamInstanceApiRequest
*/
func (a *StreamsApiService) UpdateStreamInstance(ctx context.Context, groupId string, tenantName string, streamsDataProcessRegion *StreamsDataProcessRegion) UpdateStreamInstanceApiRequest {
	return UpdateStreamInstanceApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		tenantName:               tenantName,
		streamsDataProcessRegion: streamsDataProcessRegion,
	}
}

// UpdateStreamInstanceExecute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) UpdateStreamInstanceExecute(r UpdateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.UpdateStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
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
	if r.streamsDataProcessRegion == nil {
		return localVarReturnValue, nil, reportError("streamsDataProcessRegion is required and must be specified")
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
	localVarPostBody = r.streamsDataProcessRegion
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
