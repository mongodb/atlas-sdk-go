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
		AcceptVPCPeeringConnection Requests the acceptance of an incoming VPC Peering connection.

		Requests the acceptance of an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@param vPCPeeringActionChallenge Challenge values for VPC Peering requester account ID, and requester VPC ID.
		@return AcceptVPCPeeringConnectionApiRequest
	*/
	AcceptVPCPeeringConnection(ctx context.Context, groupId string, id string, vPCPeeringActionChallenge *VPCPeeringActionChallenge) AcceptVPCPeeringConnectionApiRequest
	/*
		AcceptVPCPeeringConnection Requests the acceptance of an incoming VPC Peering connection.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AcceptVPCPeeringConnectionApiParams - Parameters for the request
		@return AcceptVPCPeeringConnectionApiRequest
	*/
	AcceptVPCPeeringConnectionWithParams(ctx context.Context, args *AcceptVPCPeeringConnectionApiParams) AcceptVPCPeeringConnectionApiRequest

	// Method available only for mocking purposes
	AcceptVPCPeeringConnectionExecute(r AcceptVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error)

	/*
		CreateStreamConnection Create One Connection

		Creates one connection for a stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param streamsConnection Details to create one connection for a streams instance in the specified project.
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnection(ctx context.Context, groupId string, tenantName string, streamsConnection *StreamsConnection) CreateStreamConnectionApiRequest
	/*
		CreateStreamConnection Create One Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamConnectionApiParams - Parameters for the request
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnectionWithParams(ctx context.Context, args *CreateStreamConnectionApiParams) CreateStreamConnectionApiRequest

	// Method available only for mocking purposes
	CreateStreamConnectionExecute(r CreateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		CreateStreamInstance Create One Stream Instance

		Creates one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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
		CreateStreamInstanceWithSampleConnections Create One Stream Instance With Sample Connections

		Creates one stream instance in the specified project with sample connections. To use this resource the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param body Details to create one streams instance in the specified project.
		@return CreateStreamInstanceWithSampleConnectionsApiRequest
	*/
	CreateStreamInstanceWithSampleConnections(ctx context.Context, groupId string, body *interface{}) CreateStreamInstanceWithSampleConnectionsApiRequest
	/*
		CreateStreamInstanceWithSampleConnections Create One Stream Instance With Sample Connections


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamInstanceWithSampleConnectionsApiParams - Parameters for the request
		@return CreateStreamInstanceWithSampleConnectionsApiRequest
	*/
	CreateStreamInstanceWithSampleConnectionsWithParams(ctx context.Context, args *CreateStreamInstanceWithSampleConnectionsApiParams) CreateStreamInstanceWithSampleConnectionsApiRequest

	// Method available only for mocking purposes
	CreateStreamInstanceWithSampleConnectionsExecute(r CreateStreamInstanceWithSampleConnectionsApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		CreateStreamProcessor Create One Stream Processor

		Create one Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
		DeleteStreamConnection Delete One Stream Connection

		Delete one connection of the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (interface{}, *http.Response, error)

	/*
		DeleteStreamInstance Delete One Stream Instance

		Delete one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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
	DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (interface{}, *http.Response, error)

	/*
		DeleteStreamProcessor Delete One Stream Processor

		Delete a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
		DeleteVPCPeeringConnection Deletes an incoming VPC Peering connection.

		Deletes an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@return DeleteVPCPeeringConnectionApiRequest
	*/
	DeleteVPCPeeringConnection(ctx context.Context, groupId string, id string) DeleteVPCPeeringConnectionApiRequest
	/*
		DeleteVPCPeeringConnection Deletes an incoming VPC Peering connection.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteVPCPeeringConnectionApiParams - Parameters for the request
		@return DeleteVPCPeeringConnectionApiRequest
	*/
	DeleteVPCPeeringConnectionWithParams(ctx context.Context, args *DeleteVPCPeeringConnectionApiParams) DeleteVPCPeeringConnectionApiRequest

	// Method available only for mocking purposes
	DeleteVPCPeeringConnectionExecute(r DeleteVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error)

	/*
		DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance

		Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

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
		GetStreamConnection Return One Stream Connection

		Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting API Key must have the Project Read Only role.

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

		Returns the details of one stream instance within the specified project. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

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
		GetStreamProcessor Get One Stream Processor

		Get one Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param processorName Human-readable label that identifies the stream processor.
		@return GetStreamProcessorApiRequest
	*/
	GetStreamProcessor(ctx context.Context, groupId string, tenantName string, processorName string) GetStreamProcessorApiRequest
	/*
		GetStreamProcessor Get One Stream Processor


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamProcessorApiParams - Parameters for the request
		@return GetStreamProcessorApiRequest
	*/
	GetStreamProcessorWithParams(ctx context.Context, args *GetStreamProcessorApiParams) GetStreamProcessorApiRequest

	// Method available only for mocking purposes
	GetStreamProcessorExecute(r GetStreamProcessorApiRequest) (*StreamsProcessorWithStats, *http.Response, error)

	/*
		GetVPCPeeringConnections Return All VPC Peering Connections.

		Returns a list of incoming VPC Peering Connections.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetVPCPeeringConnectionsApiRequest
	*/
	GetVPCPeeringConnections(ctx context.Context, groupId string) GetVPCPeeringConnectionsApiRequest
	/*
		GetVPCPeeringConnections Return All VPC Peering Connections.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetVPCPeeringConnectionsApiParams - Parameters for the request
		@return GetVPCPeeringConnectionsApiRequest
	*/
	GetVPCPeeringConnectionsWithParams(ctx context.Context, args *GetVPCPeeringConnectionsApiParams) GetVPCPeeringConnectionsApiRequest

	// Method available only for mocking purposes
	GetVPCPeeringConnectionsExecute(r GetVPCPeeringConnectionsApiRequest) (*http.Response, error)

	/*
		ListStreamConnections Return All Connections Of The Stream Instances

		Returns all connections of the stream instance for the specified project.To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnections(ctx context.Context, groupId string, tenantName string) ListStreamConnectionsApiRequest
	/*
		ListStreamConnections Return All Connections Of The Stream Instances


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamConnectionsApiParams - Parameters for the request
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnectionsWithParams(ctx context.Context, args *ListStreamConnectionsApiParams) ListStreamConnectionsApiRequest

	// Method available only for mocking purposes
	ListStreamConnectionsExecute(r ListStreamConnectionsApiRequest) (*PaginatedApiStreamsConnection, *http.Response, error)

	/*
		ListStreamInstances Return All Project Stream Instances

		Returns all stream instances for the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstances(ctx context.Context, groupId string) ListStreamInstancesApiRequest
	/*
		ListStreamInstances Return All Project Stream Instances


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamInstancesApiParams - Parameters for the request
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstancesWithParams(ctx context.Context, args *ListStreamInstancesApiParams) ListStreamInstancesApiRequest

	// Method available only for mocking purposes
	ListStreamInstancesExecute(r ListStreamInstancesApiRequest) (*PaginatedApiStreamsTenant, *http.Response, error)

	/*
		ListStreamProcessors Return All Stream Processors In The Stream Instance.

		Returns all Stream Processors within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return ListStreamProcessorsApiRequest
	*/
	ListStreamProcessors(ctx context.Context, groupId string, tenantName string) ListStreamProcessorsApiRequest
	/*
		ListStreamProcessors Return All Stream Processors In The Stream Instance.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamProcessorsApiParams - Parameters for the request
		@return ListStreamProcessorsApiRequest
	*/
	ListStreamProcessorsWithParams(ctx context.Context, args *ListStreamProcessorsApiParams) ListStreamProcessorsApiRequest

	// Method available only for mocking purposes
	ListStreamProcessorsExecute(r ListStreamProcessorsApiRequest) (*PaginatedApiStreamsStreamProcessorWithStats, *http.Response, error)

	/*
		RejectVPCPeeringConnection Requests the rejection of an incoming VPC Peering connection.

		Requests the rejection of an incoming VPC Peering connection.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param id The VPC Peering Connection id.
		@return RejectVPCPeeringConnectionApiRequest
	*/
	RejectVPCPeeringConnection(ctx context.Context, groupId string, id string) RejectVPCPeeringConnectionApiRequest
	/*
		RejectVPCPeeringConnection Requests the rejection of an incoming VPC Peering connection.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RejectVPCPeeringConnectionApiParams - Parameters for the request
		@return RejectVPCPeeringConnectionApiRequest
	*/
	RejectVPCPeeringConnectionWithParams(ctx context.Context, args *RejectVPCPeeringConnectionApiParams) RejectVPCPeeringConnectionApiRequest

	// Method available only for mocking purposes
	RejectVPCPeeringConnectionExecute(r RejectVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error)

	/*
		StartStreamProcessor Start One Stream Processor

		Start a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	StartStreamProcessorExecute(r StartStreamProcessorApiRequest) (interface{}, *http.Response, error)

	/*
		StopStreamProcessor Stop One Stream Processor

		Stop a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	StopStreamProcessorExecute(r StopStreamProcessorApiRequest) (interface{}, *http.Response, error)

	/*
		UpdateStreamConnection Update One Stream Connection

		Update one connection for the specified stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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

		Update one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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

type AcceptVPCPeeringConnectionApiRequest struct {
	ctx                       context.Context
	ApiService                StreamsApi
	groupId                   string
	id                        string
	vPCPeeringActionChallenge *VPCPeeringActionChallenge
}

type AcceptVPCPeeringConnectionApiParams struct {
	GroupId                   string
	Id                        string
	VPCPeeringActionChallenge *VPCPeeringActionChallenge
}

func (a *StreamsApiService) AcceptVPCPeeringConnectionWithParams(ctx context.Context, args *AcceptVPCPeeringConnectionApiParams) AcceptVPCPeeringConnectionApiRequest {
	return AcceptVPCPeeringConnectionApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		id:                        args.Id,
		vPCPeeringActionChallenge: args.VPCPeeringActionChallenge,
	}
}

func (r AcceptVPCPeeringConnectionApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.AcceptVPCPeeringConnectionExecute(r)
}

/*
AcceptVPCPeeringConnection Requests the acceptance of an incoming VPC Peering connection.

Requests the acceptance of an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return AcceptVPCPeeringConnectionApiRequest
*/
func (a *StreamsApiService) AcceptVPCPeeringConnection(ctx context.Context, groupId string, id string, vPCPeeringActionChallenge *VPCPeeringActionChallenge) AcceptVPCPeeringConnectionApiRequest {
	return AcceptVPCPeeringConnectionApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		id:                        id,
		vPCPeeringActionChallenge: vPCPeeringActionChallenge,
	}
}

// AcceptVPCPeeringConnectionExecute executes the request
//
//	@return interface{}
func (a *StreamsApiService) AcceptVPCPeeringConnectionExecute(r AcceptVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.AcceptVPCPeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.vPCPeeringActionChallenge == nil {
		return localVarReturnValue, nil, reportError("vPCPeeringActionChallenge is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.vPCPeeringActionChallenge
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
CreateStreamConnection Create One Connection

Creates one connection for a stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

Creates one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
	body       *interface{}
}

type CreateStreamInstanceWithSampleConnectionsApiParams struct {
	GroupId string
	Body    *interface{}
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
CreateStreamInstanceWithSampleConnections Create One Stream Instance With Sample Connections

Creates one stream instance in the specified project with sample connections. To use this resource the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateStreamInstanceWithSampleConnectionsApiRequest
*/
func (a *StreamsApiService) CreateStreamInstanceWithSampleConnections(ctx context.Context, groupId string, body *interface{}) CreateStreamInstanceWithSampleConnectionsApiRequest {
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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json", "application/json"}

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

Create one Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

func (r DeleteStreamConnectionApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteStreamConnectionExecute(r)
}

/*
DeleteStreamConnection Delete One Stream Connection

Delete one connection of the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
//
//	@return interface{}
func (a *StreamsApiService) DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

func (r DeleteStreamInstanceApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteStreamInstanceExecute(r)
}

/*
DeleteStreamInstance Delete One Stream Instance

Delete one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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
//
//	@return interface{}
func (a *StreamsApiService) DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

Delete a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(parameterValueToString(r.processorName, "processorName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

type DeleteVPCPeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	id         string
}

type DeleteVPCPeeringConnectionApiParams struct {
	GroupId string
	Id      string
}

func (a *StreamsApiService) DeleteVPCPeeringConnectionWithParams(ctx context.Context, args *DeleteVPCPeeringConnectionApiParams) DeleteVPCPeeringConnectionApiRequest {
	return DeleteVPCPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		id:         args.Id,
	}
}

func (r DeleteVPCPeeringConnectionApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteVPCPeeringConnectionExecute(r)
}

/*
DeleteVPCPeeringConnection Deletes an incoming VPC Peering connection.

Deletes an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return DeleteVPCPeeringConnectionApiRequest
*/
func (a *StreamsApiService) DeleteVPCPeeringConnection(ctx context.Context, groupId string, id string) DeleteVPCPeeringConnectionApiRequest {
	return DeleteVPCPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		id:         id,
	}
}

// DeleteVPCPeeringConnectionExecute executes the request
//
//	@return interface{}
func (a *StreamsApiService) DeleteVPCPeeringConnectionExecute(r DeleteVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteVPCPeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+gzip", "application/json"}

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

Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting API Key must have the Project Read Only role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

Returns the details of one stream instance within the specified project. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
GetStreamProcessor Get One Stream Processor

Get one Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(parameterValueToString(r.processorName, "processorName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

type GetVPCPeeringConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type GetVPCPeeringConnectionsApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) GetVPCPeeringConnectionsWithParams(ctx context.Context, args *GetVPCPeeringConnectionsApiParams) GetVPCPeeringConnectionsApiRequest {
	return GetVPCPeeringConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r GetVPCPeeringConnectionsApiRequest) ItemsPerPage(itemsPerPage int) GetVPCPeeringConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetVPCPeeringConnectionsApiRequest) PageNum(pageNum int) GetVPCPeeringConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r GetVPCPeeringConnectionsApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetVPCPeeringConnectionsExecute(r)
}

/*
GetVPCPeeringConnections Return All VPC Peering Connections.

Returns a list of incoming VPC Peering Connections.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetVPCPeeringConnectionsApiRequest
*/
func (a *StreamsApiService) GetVPCPeeringConnections(ctx context.Context, groupId string) GetVPCPeeringConnectionsApiRequest {
	return GetVPCPeeringConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetVPCPeeringConnectionsExecute executes the request
func (a *StreamsApiService) GetVPCPeeringConnectionsExecute(r GetVPCPeeringConnectionsApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetVPCPeeringConnections")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
ListStreamConnections Return All Connections Of The Stream Instances

Returns all connections of the stream instance for the specified project.To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
ListStreamInstances Return All Project Stream Instances

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
ListStreamProcessors Return All Stream Processors In The Stream Instance.

Returns all Stream Processors within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

type RejectVPCPeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	id         string
}

type RejectVPCPeeringConnectionApiParams struct {
	GroupId string
	Id      string
}

func (a *StreamsApiService) RejectVPCPeeringConnectionWithParams(ctx context.Context, args *RejectVPCPeeringConnectionApiParams) RejectVPCPeeringConnectionApiRequest {
	return RejectVPCPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		id:         args.Id,
	}
}

func (r RejectVPCPeeringConnectionApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.RejectVPCPeeringConnectionExecute(r)
}

/*
RejectVPCPeeringConnection Requests the rejection of an incoming VPC Peering connection.

Requests the rejection of an incoming VPC Peering connection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param id The VPC Peering Connection id.
	@return RejectVPCPeeringConnectionApiRequest
*/
func (a *StreamsApiService) RejectVPCPeeringConnection(ctx context.Context, groupId string, id string) RejectVPCPeeringConnectionApiRequest {
	return RejectVPCPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		id:         id,
	}
}

// RejectVPCPeeringConnectionExecute executes the request
//
//	@return interface{}
func (a *StreamsApiService) RejectVPCPeeringConnectionExecute(r RejectVPCPeeringConnectionApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.RejectVPCPeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

func (r StartStreamProcessorApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StartStreamProcessorExecute(r)
}

/*
StartStreamProcessor Start One Stream Processor

Start a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
//
//	@return interface{}
func (a *StreamsApiService) StartStreamProcessorExecute(r StartStreamProcessorApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.StartStreamProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(parameterValueToString(r.processorName, "processorName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

func (r StopStreamProcessorApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StopStreamProcessorExecute(r)
}

/*
StopStreamProcessor Stop One Stream Processor

Stop a Stream Processor within the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
//
//	@return interface{}
func (a *StreamsApiService) StopStreamProcessorExecute(r StopStreamProcessorApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.StopStreamProcessor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorName"+"}", url.PathEscape(parameterValueToString(r.processorName, "processorName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json", "application/json"}

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

Update one connection for the specified stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

Update one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

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
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
