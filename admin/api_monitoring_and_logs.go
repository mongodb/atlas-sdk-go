// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type MonitoringAndLogsApi interface {

	/*
		GetAtlasProcess Return One MongoDB Process by ID

		Returns the processes for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return GetAtlasProcessApiRequest
	*/
	GetAtlasProcess(ctx context.Context, groupId string, processId string) GetAtlasProcessApiRequest
	/*
		GetAtlasProcess Return One MongoDB Process by ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasProcessApiParams - Parameters for the request
		@return GetAtlasProcessApiRequest
	*/
	GetAtlasProcessWithParams(ctx context.Context, args *GetAtlasProcessApiParams) GetAtlasProcessApiRequest

	// Method available only for mocking purposes
	GetAtlasProcessExecute(r GetAtlasProcessApiRequest) (*ApiHostViewAtlas, *http.Response, error)

	/*
		GetDatabase Return One Database for One MongoDB Process

		Returns one database running on the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param databaseName Human-readable label that identifies the database that the specified MongoDB process serves.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return GetDatabaseApiRequest
	*/
	GetDatabase(ctx context.Context, groupId string, databaseName string, processId string) GetDatabaseApiRequest
	/*
		GetDatabase Return One Database for One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDatabaseApiParams - Parameters for the request
		@return GetDatabaseApiRequest
	*/
	GetDatabaseWithParams(ctx context.Context, args *GetDatabaseApiParams) GetDatabaseApiRequest

	// Method available only for mocking purposes
	GetDatabaseExecute(r GetDatabaseApiRequest) (*MesurementsDatabase, *http.Response, error)

	/*
		GetDatabaseMeasurements Return Measurements for One Database in One MongoDB Process

		Returns the measurements of one database for the specified host for the specified project. Returns the database's on-disk storage space based on the MongoDB `dbStats` command output. To calculate some metric series, Atlas takes the rate between every two adjacent points. For these metric series, the first data point has a null value because Atlas can't calculate a rate for the first data point given the query time range. Atlas retrieves database metrics every 20 minutes but reduces frequency when necessary to optimize database performance. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param databaseName Human-readable label that identifies the database that the specified MongoDB process serves.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return GetDatabaseMeasurementsApiRequest
	*/
	GetDatabaseMeasurements(ctx context.Context, groupId string, databaseName string, processId string) GetDatabaseMeasurementsApiRequest
	/*
		GetDatabaseMeasurements Return Measurements for One Database in One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDatabaseMeasurementsApiParams - Parameters for the request
		@return GetDatabaseMeasurementsApiRequest
	*/
	GetDatabaseMeasurementsWithParams(ctx context.Context, args *GetDatabaseMeasurementsApiParams) GetDatabaseMeasurementsApiRequest

	// Method available only for mocking purposes
	GetDatabaseMeasurementsExecute(r GetDatabaseMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error)

	/*
			GetDiskMeasurements Return Measurements of One Disk for One MongoDB Process

			Returns the measurements of one disk or partition for the specified host for the specified project. Returned value can be one of the following:
		- Throughput of I/O operations for the disk partition used for the MongoDB process
		- Percentage of time during which requests the partition issued and serviced
		- Latency per operation type of the disk partition used for the MongoDB process
		- Amount of free and used disk space on the disk partition used for the MongoDB process

		To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param partitionName Human-readable label of the disk or partition to which the measurements apply.
			@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
			@return GetDiskMeasurementsApiRequest
	*/
	GetDiskMeasurements(ctx context.Context, groupId string, partitionName string, processId string) GetDiskMeasurementsApiRequest
	/*
		GetDiskMeasurements Return Measurements of One Disk for One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDiskMeasurementsApiParams - Parameters for the request
		@return GetDiskMeasurementsApiRequest
	*/
	GetDiskMeasurementsWithParams(ctx context.Context, args *GetDiskMeasurementsApiParams) GetDiskMeasurementsApiRequest

	// Method available only for mocking purposes
	GetDiskMeasurementsExecute(r GetDiskMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error)

	/*
		GetHostLogs Download Logs for One Cluster Host in One Project

		Returns a compressed (.gz) log file that contains a range of log messages for the specified host for the specified project. MongoDB updates process and audit logs from the cluster backend infrastructure every five minutes. Logs are stored in chunks approximately five minutes in length, but this duration may vary. If you poll the API for log files, we recommend polling every five minutes even though consecutive polls could contain some overlapping logs. This feature isn't available for `M0` free clusters, `M2`, `M5`, flex, or serverless clusters. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Only or higher role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip". Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param hostName Human-readable label that identifies the host that stores the log files that you want to download.
		@param logName Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project.
		@return GetHostLogsApiRequest
	*/
	GetHostLogs(ctx context.Context, groupId string, hostName string, logName string) GetHostLogsApiRequest
	/*
		GetHostLogs Download Logs for One Cluster Host in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetHostLogsApiParams - Parameters for the request
		@return GetHostLogsApiRequest
	*/
	GetHostLogsWithParams(ctx context.Context, args *GetHostLogsApiParams) GetHostLogsApiRequest

	// Method available only for mocking purposes
	GetHostLogsExecute(r GetHostLogsApiRequest) (io.ReadCloser, *http.Response, error)

	/*
			GetHostMeasurements Return Measurements for One MongoDB Process

			Returns disk, partition, or host measurements per process for the specified host for the specified project. Returned value can be one of the following:
		- Throughput of I/O operations for the disk partition used for the MongoDB process
		- Percentage of time during which requests the partition issued and serviced
		- Latency per operation type of the disk partition used for the MongoDB process
		- Amount of free and used disk space on the disk partition used for the MongoDB process
		- Measurements for the host, such as CPU usage or number of I/O operations

		To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
			@return GetHostMeasurementsApiRequest
	*/
	GetHostMeasurements(ctx context.Context, groupId string, processId string) GetHostMeasurementsApiRequest
	/*
		GetHostMeasurements Return Measurements for One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetHostMeasurementsApiParams - Parameters for the request
		@return GetHostMeasurementsApiRequest
	*/
	GetHostMeasurementsWithParams(ctx context.Context, args *GetHostMeasurementsApiParams) GetHostMeasurementsApiRequest

	// Method available only for mocking purposes
	GetHostMeasurementsExecute(r GetHostMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error)

	/*
		GetIndexMetrics Return Atlas Search Metrics for One Index in One Namespace

		Returns the Atlas Search metrics data series within the provided time range for one namespace and index name on the specified process. You must have the Project Read Only or higher role to view the Atlas Search metric types.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@param indexName Human-readable label that identifies the index.
		@param databaseName Human-readable label that identifies the database.
		@param collectionName Human-readable label that identifies the collection.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetIndexMetricsApiRequest
	*/
	GetIndexMetrics(ctx context.Context, processId string, indexName string, databaseName string, collectionName string, groupId string) GetIndexMetricsApiRequest
	/*
		GetIndexMetrics Return Atlas Search Metrics for One Index in One Namespace


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetIndexMetricsApiParams - Parameters for the request
		@return GetIndexMetricsApiRequest
	*/
	GetIndexMetricsWithParams(ctx context.Context, args *GetIndexMetricsApiParams) GetIndexMetricsApiRequest

	// Method available only for mocking purposes
	GetIndexMetricsExecute(r GetIndexMetricsApiRequest) (*MeasurementsIndexes, *http.Response, error)

	/*
		GetMeasurements Return Atlas Search Hardware and Status Metrics

		Returns the Atlas Search hardware and status data series within the provided time range for one process in the specified project. You must have the Project Read Only or higher role to view the Atlas Search metric types.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetMeasurementsApiRequest
	*/
	GetMeasurements(ctx context.Context, processId string, groupId string) GetMeasurementsApiRequest
	/*
		GetMeasurements Return Atlas Search Hardware and Status Metrics


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetMeasurementsApiParams - Parameters for the request
		@return GetMeasurementsApiRequest
	*/
	GetMeasurementsWithParams(ctx context.Context, args *GetMeasurementsApiParams) GetMeasurementsApiRequest

	// Method available only for mocking purposes
	GetMeasurementsExecute(r GetMeasurementsApiRequest) (*MeasurementsNonIndex, *http.Response, error)

	/*
		ListAtlasProcesses Return All MongoDB Processes in One Project

		Returns details of all processes for the specified project. A MongoDB process can be either a `mongod` or `mongos`. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListAtlasProcessesApiRequest
	*/
	ListAtlasProcesses(ctx context.Context, groupId string) ListAtlasProcessesApiRequest
	/*
		ListAtlasProcesses Return All MongoDB Processes in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasProcessesApiParams - Parameters for the request
		@return ListAtlasProcessesApiRequest
	*/
	ListAtlasProcessesWithParams(ctx context.Context, args *ListAtlasProcessesApiParams) ListAtlasProcessesApiRequest

	// Method available only for mocking purposes
	ListAtlasProcessesExecute(r ListAtlasProcessesApiRequest) (*PaginatedHostViewAtlas, *http.Response, error)

	/*
		ListDatabases Return Available Databases for One MongoDB Process

		Returns the list of databases running on the specified host for the specified project. `M0` free clusters, `M2`, `M5`, serverless, and Flex clusters have some operational limits. The MongoDB Cloud process must be a `mongod`. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListDatabasesApiRequest
	*/
	ListDatabases(ctx context.Context, groupId string, processId string) ListDatabasesApiRequest
	/*
		ListDatabases Return Available Databases for One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDatabasesApiParams - Parameters for the request
		@return ListDatabasesApiRequest
	*/
	ListDatabasesWithParams(ctx context.Context, args *ListDatabasesApiParams) ListDatabasesApiRequest

	// Method available only for mocking purposes
	ListDatabasesExecute(r ListDatabasesApiRequest) (*PaginatedDatabase, *http.Response, error)

	/*
		ListDiskMeasurements Return Measurements for One Disk

		Returns measurement details for one disk or partition for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param partitionName Human-readable label of the disk or partition to which the measurements apply.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListDiskMeasurementsApiRequest
	*/
	ListDiskMeasurements(ctx context.Context, partitionName string, groupId string, processId string) ListDiskMeasurementsApiRequest
	/*
		ListDiskMeasurements Return Measurements for One Disk


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDiskMeasurementsApiParams - Parameters for the request
		@return ListDiskMeasurementsApiRequest
	*/
	ListDiskMeasurementsWithParams(ctx context.Context, args *ListDiskMeasurementsApiParams) ListDiskMeasurementsApiRequest

	// Method available only for mocking purposes
	ListDiskMeasurementsExecute(r ListDiskMeasurementsApiRequest) (*MeasurementDiskPartition, *http.Response, error)

	/*
		ListDiskPartitions Return Available Disks for One MongoDB Process

		Returns the list of disks or partitions for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListDiskPartitionsApiRequest
	*/
	ListDiskPartitions(ctx context.Context, groupId string, processId string) ListDiskPartitionsApiRequest
	/*
		ListDiskPartitions Return Available Disks for One MongoDB Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDiskPartitionsApiParams - Parameters for the request
		@return ListDiskPartitionsApiRequest
	*/
	ListDiskPartitionsWithParams(ctx context.Context, args *ListDiskPartitionsApiParams) ListDiskPartitionsApiRequest

	// Method available only for mocking purposes
	ListDiskPartitionsExecute(r ListDiskPartitionsApiRequest) (*PaginatedDiskPartition, *http.Response, error)

	/*
		ListIndexMetrics Return All Atlas Search Index Metrics for One Namespace

		Returns the Atlas Search index metrics within the specified time range for one namespace in the specified process.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@param databaseName Human-readable label that identifies the database.
		@param collectionName Human-readable label that identifies the collection.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListIndexMetricsApiRequest
	*/
	ListIndexMetrics(ctx context.Context, processId string, databaseName string, collectionName string, groupId string) ListIndexMetricsApiRequest
	/*
		ListIndexMetrics Return All Atlas Search Index Metrics for One Namespace


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListIndexMetricsApiParams - Parameters for the request
		@return ListIndexMetricsApiRequest
	*/
	ListIndexMetricsWithParams(ctx context.Context, args *ListIndexMetricsApiParams) ListIndexMetricsApiRequest

	// Method available only for mocking purposes
	ListIndexMetricsExecute(r ListIndexMetricsApiRequest) (*MeasurementsIndexes, *http.Response, error)

	/*
		ListMetricTypes Return All Atlas Search Metric Types for One Process

		Returns all Atlas Search metric types available for one process in the specified project. You must have the Project Read Only or higher role to view the Atlas Search metric types.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListMetricTypesApiRequest
	*/
	ListMetricTypes(ctx context.Context, processId string, groupId string) ListMetricTypesApiRequest
	/*
		ListMetricTypes Return All Atlas Search Metric Types for One Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListMetricTypesApiParams - Parameters for the request
		@return ListMetricTypesApiRequest
	*/
	ListMetricTypesWithParams(ctx context.Context, args *ListMetricTypesApiParams) ListMetricTypesApiRequest

	// Method available only for mocking purposes
	ListMetricTypesExecute(r ListMetricTypesApiRequest) (*CloudSearchMetrics, *http.Response, error)
}

// MonitoringAndLogsApiService MonitoringAndLogsApi service
type MonitoringAndLogsApiService service

type GetAtlasProcessApiRequest struct {
	ctx        context.Context
	ApiService MonitoringAndLogsApi
	groupId    string
	processId  string
}

type GetAtlasProcessApiParams struct {
	GroupId   string
	ProcessId string
}

func (a *MonitoringAndLogsApiService) GetAtlasProcessWithParams(ctx context.Context, args *GetAtlasProcessApiParams) GetAtlasProcessApiRequest {
	return GetAtlasProcessApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		processId:  args.ProcessId,
	}
}

func (r GetAtlasProcessApiRequest) Execute() (*ApiHostViewAtlas, *http.Response, error) {
	return r.ApiService.GetAtlasProcessExecute(r)
}

/*
GetAtlasProcess Return One MongoDB Process by ID

Returns the processes for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetAtlasProcessApiRequest
*/
func (a *MonitoringAndLogsApiService) GetAtlasProcess(ctx context.Context, groupId string, processId string) GetAtlasProcessApiRequest {
	return GetAtlasProcessApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// GetAtlasProcessExecute executes the request
//
//	@return ApiHostViewAtlas
func (a *MonitoringAndLogsApiService) GetAtlasProcessExecute(r GetAtlasProcessApiRequest) (*ApiHostViewAtlas, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiHostViewAtlas
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetAtlasProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

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

type GetDatabaseApiRequest struct {
	ctx          context.Context
	ApiService   MonitoringAndLogsApi
	groupId      string
	databaseName string
	processId    string
}

type GetDatabaseApiParams struct {
	GroupId      string
	DatabaseName string
	ProcessId    string
}

func (a *MonitoringAndLogsApiService) GetDatabaseWithParams(ctx context.Context, args *GetDatabaseApiParams) GetDatabaseApiRequest {
	return GetDatabaseApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		databaseName: args.DatabaseName,
		processId:    args.ProcessId,
	}
}

func (r GetDatabaseApiRequest) Execute() (*MesurementsDatabase, *http.Response, error) {
	return r.ApiService.GetDatabaseExecute(r)
}

/*
GetDatabase Return One Database for One MongoDB Process

Returns one database running on the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param databaseName Human-readable label that identifies the database that the specified MongoDB process serves.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetDatabaseApiRequest
*/
func (a *MonitoringAndLogsApiService) GetDatabase(ctx context.Context, groupId string, databaseName string, processId string) GetDatabaseApiRequest {
	return GetDatabaseApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		databaseName: databaseName,
		processId:    processId,
	}
}

// GetDatabaseExecute executes the request
//
//	@return MesurementsDatabase
func (a *MonitoringAndLogsApiService) GetDatabaseExecute(r GetDatabaseApiRequest) (*MesurementsDatabase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MesurementsDatabase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

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

type GetDatabaseMeasurementsApiRequest struct {
	ctx          context.Context
	ApiService   MonitoringAndLogsApi
	groupId      string
	databaseName string
	processId    string
	granularity  *string
	m            *[]string
	period       *string
	start        *time.Time
	end          *time.Time
}

type GetDatabaseMeasurementsApiParams struct {
	GroupId      string
	DatabaseName string
	ProcessId    string
	Granularity  *string
	M            *[]string
	Period       *string
	Start        *time.Time
	End          *time.Time
}

func (a *MonitoringAndLogsApiService) GetDatabaseMeasurementsWithParams(ctx context.Context, args *GetDatabaseMeasurementsApiParams) GetDatabaseMeasurementsApiRequest {
	return GetDatabaseMeasurementsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		databaseName: args.DatabaseName,
		processId:    args.ProcessId,
		granularity:  args.Granularity,
		m:            args.M,
		period:       args.Period,
		start:        args.Start,
		end:          args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r GetDatabaseMeasurementsApiRequest) Granularity(granularity string) GetDatabaseMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.
func (r GetDatabaseMeasurementsApiRequest) M(m []string) GetDatabaseMeasurementsApiRequest {
	r.m = &m
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetDatabaseMeasurementsApiRequest) Period(period string) GetDatabaseMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetDatabaseMeasurementsApiRequest) Start(start time.Time) GetDatabaseMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetDatabaseMeasurementsApiRequest) End(end time.Time) GetDatabaseMeasurementsApiRequest {
	r.end = &end
	return r
}

func (r GetDatabaseMeasurementsApiRequest) Execute() (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	return r.ApiService.GetDatabaseMeasurementsExecute(r)
}

/*
GetDatabaseMeasurements Return Measurements for One Database in One MongoDB Process

Returns the measurements of one database for the specified host for the specified project. Returns the database's on-disk storage space based on the MongoDB `dbStats` command output. To calculate some metric series, Atlas takes the rate between every two adjacent points. For these metric series, the first data point has a null value because Atlas can't calculate a rate for the first data point given the query time range. Atlas retrieves database metrics every 20 minutes but reduces frequency when necessary to optimize database performance. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param databaseName Human-readable label that identifies the database that the specified MongoDB process serves.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetDatabaseMeasurementsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetDatabaseMeasurements(ctx context.Context, groupId string, databaseName string, processId string) GetDatabaseMeasurementsApiRequest {
	return GetDatabaseMeasurementsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		databaseName: databaseName,
		processId:    processId,
	}
}

// GetDatabaseMeasurementsExecute executes the request
//
//	@return ApiMeasurementsGeneralViewAtlas
func (a *MonitoringAndLogsApiService) GetDatabaseMeasurementsExecute(r GetDatabaseMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiMeasurementsGeneralViewAtlas
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetDatabaseMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}

	if r.m != nil {
		t := *r.m
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "m", t, "multi")

	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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

type GetDiskMeasurementsApiRequest struct {
	ctx           context.Context
	ApiService    MonitoringAndLogsApi
	groupId       string
	partitionName string
	processId     string
	granularity   *string
	m             *[]string
	period        *string
	start         *time.Time
	end           *time.Time
}

type GetDiskMeasurementsApiParams struct {
	GroupId       string
	PartitionName string
	ProcessId     string
	Granularity   *string
	M             *[]string
	Period        *string
	Start         *time.Time
	End           *time.Time
}

func (a *MonitoringAndLogsApiService) GetDiskMeasurementsWithParams(ctx context.Context, args *GetDiskMeasurementsApiParams) GetDiskMeasurementsApiRequest {
	return GetDiskMeasurementsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		partitionName: args.PartitionName,
		processId:     args.ProcessId,
		granularity:   args.Granularity,
		m:             args.M,
		period:        args.Period,
		start:         args.Start,
		end:           args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r GetDiskMeasurementsApiRequest) Granularity(granularity string) GetDiskMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.
func (r GetDiskMeasurementsApiRequest) M(m []string) GetDiskMeasurementsApiRequest {
	r.m = &m
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetDiskMeasurementsApiRequest) Period(period string) GetDiskMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetDiskMeasurementsApiRequest) Start(start time.Time) GetDiskMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetDiskMeasurementsApiRequest) End(end time.Time) GetDiskMeasurementsApiRequest {
	r.end = &end
	return r
}

func (r GetDiskMeasurementsApiRequest) Execute() (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	return r.ApiService.GetDiskMeasurementsExecute(r)
}

/*
GetDiskMeasurements Return Measurements of One Disk for One MongoDB Process

Returns the measurements of one disk or partition for the specified host for the specified project. Returned value can be one of the following:
- Throughput of I/O operations for the disk partition used for the MongoDB process
- Percentage of time during which requests the partition issued and serviced
- Latency per operation type of the disk partition used for the MongoDB process
- Amount of free and used disk space on the disk partition used for the MongoDB process

To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param partitionName Human-readable label of the disk or partition to which the measurements apply.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetDiskMeasurementsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetDiskMeasurements(ctx context.Context, groupId string, partitionName string, processId string) GetDiskMeasurementsApiRequest {
	return GetDiskMeasurementsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		partitionName: partitionName,
		processId:     processId,
	}
}

// GetDiskMeasurementsExecute executes the request
//
//	@return ApiMeasurementsGeneralViewAtlas
func (a *MonitoringAndLogsApiService) GetDiskMeasurementsExecute(r GetDiskMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiMeasurementsGeneralViewAtlas
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetDiskMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.partitionName == "" {
		return localVarReturnValue, nil, reportError("partitionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"partitionName"+"}", url.PathEscape(r.partitionName), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}

	if r.m != nil {
		t := *r.m
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "m", t, "multi")

	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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

type GetHostLogsApiRequest struct {
	ctx        context.Context
	ApiService MonitoringAndLogsApi
	groupId    string
	hostName   string
	logName    string
	endDate    *int64
	startDate  *int64
}

type GetHostLogsApiParams struct {
	GroupId   string
	HostName  string
	LogName   string
	EndDate   *int64
	StartDate *int64
}

func (a *MonitoringAndLogsApiService) GetHostLogsWithParams(ctx context.Context, args *GetHostLogsApiParams) GetHostLogsApiRequest {
	return GetHostLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		hostName:   args.HostName,
		logName:    args.LogName,
		endDate:    args.EndDate,
		startDate:  args.StartDate,
	}
}

// Specifies the date and time for the ending point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours after the start date. If the start date is also unspecified, the value will default to the time of the request.
func (r GetHostLogsApiRequest) EndDate(endDate int64) GetHostLogsApiRequest {
	r.endDate = &endDate
	return r
}

// Specifies the date and time for the starting point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours prior to the end date. If the end date is also unspecified, the value will default to 24 hours prior to the time of the request.
func (r GetHostLogsApiRequest) StartDate(startDate int64) GetHostLogsApiRequest {
	r.startDate = &startDate
	return r
}

func (r GetHostLogsApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.GetHostLogsExecute(r)
}

/*
GetHostLogs Download Logs for One Cluster Host in One Project

Returns a compressed (.gz) log file that contains a range of log messages for the specified host for the specified project. MongoDB updates process and audit logs from the cluster backend infrastructure every five minutes. Logs are stored in chunks approximately five minutes in length, but this duration may vary. If you poll the API for log files, we recommend polling every five minutes even though consecutive polls could contain some overlapping logs. This feature isn't available for `M0` free clusters, `M2`, `M5`, flex, or serverless clusters. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Only or higher role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip". Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param hostName Human-readable label that identifies the host that stores the log files that you want to download.
	@param logName Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project.
	@return GetHostLogsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetHostLogs(ctx context.Context, groupId string, hostName string, logName string) GetHostLogsApiRequest {
	return GetHostLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		hostName:   hostName,
		logName:    logName,
	}
}

// GetHostLogsExecute executes the request
//
//	@return io.ReadCloser
func (a *MonitoringAndLogsApiService) GetHostLogsExecute(r GetHostLogsApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetHostLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.hostName == "" {
		return localVarReturnValue, nil, reportError("hostName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"hostName"+"}", url.PathEscape(r.hostName), -1)
	if r.logName == "" {
		return localVarReturnValue, nil, reportError("logName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"logName"+"}", url.PathEscape(r.logName), -1)

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

type GetHostMeasurementsApiRequest struct {
	ctx         context.Context
	ApiService  MonitoringAndLogsApi
	groupId     string
	processId   string
	granularity *string
	m           *[]string
	period      *string
	start       *time.Time
	end         *time.Time
}

type GetHostMeasurementsApiParams struct {
	GroupId     string
	ProcessId   string
	Granularity *string
	M           *[]string
	Period      *string
	Start       *time.Time
	End         *time.Time
}

func (a *MonitoringAndLogsApiService) GetHostMeasurementsWithParams(ctx context.Context, args *GetHostMeasurementsApiParams) GetHostMeasurementsApiRequest {
	return GetHostMeasurementsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		processId:   args.ProcessId,
		granularity: args.Granularity,
		m:           args.M,
		period:      args.Period,
		start:       args.Start,
		end:         args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r GetHostMeasurementsApiRequest) Granularity(granularity string) GetHostMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.
func (r GetHostMeasurementsApiRequest) M(m []string) GetHostMeasurementsApiRequest {
	r.m = &m
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetHostMeasurementsApiRequest) Period(period string) GetHostMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetHostMeasurementsApiRequest) Start(start time.Time) GetHostMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetHostMeasurementsApiRequest) End(end time.Time) GetHostMeasurementsApiRequest {
	r.end = &end
	return r
}

func (r GetHostMeasurementsApiRequest) Execute() (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	return r.ApiService.GetHostMeasurementsExecute(r)
}

/*
GetHostMeasurements Return Measurements for One MongoDB Process

Returns disk, partition, or host measurements per process for the specified host for the specified project. Returned value can be one of the following:
- Throughput of I/O operations for the disk partition used for the MongoDB process
- Percentage of time during which requests the partition issued and serviced
- Latency per operation type of the disk partition used for the MongoDB process
- Amount of free and used disk space on the disk partition used for the MongoDB process
- Measurements for the host, such as CPU usage or number of I/O operations

To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetHostMeasurementsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetHostMeasurements(ctx context.Context, groupId string, processId string) GetHostMeasurementsApiRequest {
	return GetHostMeasurementsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// GetHostMeasurementsExecute executes the request
//
//	@return ApiMeasurementsGeneralViewAtlas
func (a *MonitoringAndLogsApiService) GetHostMeasurementsExecute(r GetHostMeasurementsApiRequest) (*ApiMeasurementsGeneralViewAtlas, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiMeasurementsGeneralViewAtlas
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetHostMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/measurements"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}

	if r.m != nil {
		t := *r.m
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "m", t, "multi")

	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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

type GetIndexMetricsApiRequest struct {
	ctx            context.Context
	ApiService     MonitoringAndLogsApi
	processId      string
	indexName      string
	databaseName   string
	collectionName string
	groupId        string
	granularity    *string
	metrics        *[]string
	period         *string
	start          *time.Time
	end            *time.Time
}

type GetIndexMetricsApiParams struct {
	ProcessId      string
	IndexName      string
	DatabaseName   string
	CollectionName string
	GroupId        string
	Granularity    *string
	Metrics        *[]string
	Period         *string
	Start          *time.Time
	End            *time.Time
}

func (a *MonitoringAndLogsApiService) GetIndexMetricsWithParams(ctx context.Context, args *GetIndexMetricsApiParams) GetIndexMetricsApiRequest {
	return GetIndexMetricsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		processId:      args.ProcessId,
		indexName:      args.IndexName,
		databaseName:   args.DatabaseName,
		collectionName: args.CollectionName,
		groupId:        args.GroupId,
		granularity:    args.Granularity,
		metrics:        args.Metrics,
		period:         args.Period,
		start:          args.Start,
		end:            args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r GetIndexMetricsApiRequest) Granularity(granularity string) GetIndexMetricsApiRequest {
	r.granularity = &granularity
	return r
}

// List that contains the measurements that MongoDB Atlas reports for the associated data series.
func (r GetIndexMetricsApiRequest) Metrics(metrics []string) GetIndexMetricsApiRequest {
	r.metrics = &metrics
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetIndexMetricsApiRequest) Period(period string) GetIndexMetricsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetIndexMetricsApiRequest) Start(start time.Time) GetIndexMetricsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetIndexMetricsApiRequest) End(end time.Time) GetIndexMetricsApiRequest {
	r.end = &end
	return r
}

func (r GetIndexMetricsApiRequest) Execute() (*MeasurementsIndexes, *http.Response, error) {
	return r.ApiService.GetIndexMetricsExecute(r)
}

/*
GetIndexMetrics Return Atlas Search Metrics for One Index in One Namespace

Returns the Atlas Search metrics data series within the provided time range for one namespace and index name on the specified process. You must have the Project Read Only or higher role to view the Atlas Search metric types.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@param indexName Human-readable label that identifies the index.
	@param databaseName Human-readable label that identifies the database.
	@param collectionName Human-readable label that identifies the collection.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetIndexMetricsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetIndexMetrics(ctx context.Context, processId string, indexName string, databaseName string, collectionName string, groupId string) GetIndexMetricsApiRequest {
	return GetIndexMetricsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		processId:      processId,
		indexName:      indexName,
		databaseName:   databaseName,
		collectionName: collectionName,
		groupId:        groupId,
	}
}

// GetIndexMetricsExecute executes the request
//
//	@return MeasurementsIndexes
func (a *MonitoringAndLogsApiService) GetIndexMetricsExecute(r GetIndexMetricsApiRequest) (*MeasurementsIndexes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MeasurementsIndexes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetIndexMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements"
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)
	if r.indexName == "" {
		return localVarReturnValue, nil, reportError("indexName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexName"+"}", url.PathEscape(r.indexName), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, reportError("metrics is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	{
		t := *r.metrics
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", t, "multi")
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

type GetMeasurementsApiRequest struct {
	ctx         context.Context
	ApiService  MonitoringAndLogsApi
	processId   string
	groupId     string
	granularity *string
	metrics     *[]string
	period      *string
	start       *time.Time
	end         *time.Time
}

type GetMeasurementsApiParams struct {
	ProcessId   string
	GroupId     string
	Granularity *string
	Metrics     *[]string
	Period      *string
	Start       *time.Time
	End         *time.Time
}

func (a *MonitoringAndLogsApiService) GetMeasurementsWithParams(ctx context.Context, args *GetMeasurementsApiParams) GetMeasurementsApiRequest {
	return GetMeasurementsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		processId:   args.ProcessId,
		groupId:     args.GroupId,
		granularity: args.Granularity,
		metrics:     args.Metrics,
		period:      args.Period,
		start:       args.Start,
		end:         args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r GetMeasurementsApiRequest) Granularity(granularity string) GetMeasurementsApiRequest {
	r.granularity = &granularity
	return r
}

// List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don&#39;t set this parameter, this resource returns all hardware and status metrics for the associated data series.
func (r GetMeasurementsApiRequest) Metrics(metrics []string) GetMeasurementsApiRequest {
	r.metrics = &metrics
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetMeasurementsApiRequest) Period(period string) GetMeasurementsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetMeasurementsApiRequest) Start(start time.Time) GetMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetMeasurementsApiRequest) End(end time.Time) GetMeasurementsApiRequest {
	r.end = &end
	return r
}

func (r GetMeasurementsApiRequest) Execute() (*MeasurementsNonIndex, *http.Response, error) {
	return r.ApiService.GetMeasurementsExecute(r)
}

/*
GetMeasurements Return Atlas Search Hardware and Status Metrics

Returns the Atlas Search hardware and status data series within the provided time range for one process in the specified project. You must have the Project Read Only or higher role to view the Atlas Search metric types.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetMeasurementsApiRequest
*/
func (a *MonitoringAndLogsApiService) GetMeasurements(ctx context.Context, processId string, groupId string) GetMeasurementsApiRequest {
	return GetMeasurementsApiRequest{
		ApiService: a,
		ctx:        ctx,
		processId:  processId,
		groupId:    groupId,
	}
}

// GetMeasurementsExecute executes the request
//
//	@return MeasurementsNonIndex
func (a *MonitoringAndLogsApiService) GetMeasurementsExecute(r GetMeasurementsApiRequest) (*MeasurementsNonIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MeasurementsNonIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.GetMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements"
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, reportError("metrics is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	{
		t := *r.metrics
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", t, "multi")
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

type ListAtlasProcessesApiRequest struct {
	ctx          context.Context
	ApiService   MonitoringAndLogsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListAtlasProcessesApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *MonitoringAndLogsApiService) ListAtlasProcessesWithParams(ctx context.Context, args *ListAtlasProcessesApiParams) ListAtlasProcessesApiRequest {
	return ListAtlasProcessesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListAtlasProcessesApiRequest) IncludeCount(includeCount bool) ListAtlasProcessesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAtlasProcessesApiRequest) ItemsPerPage(itemsPerPage int) ListAtlasProcessesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAtlasProcessesApiRequest) PageNum(pageNum int) ListAtlasProcessesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListAtlasProcessesApiRequest) Execute() (*PaginatedHostViewAtlas, *http.Response, error) {
	return r.ApiService.ListAtlasProcessesExecute(r)
}

/*
ListAtlasProcesses Return All MongoDB Processes in One Project

Returns details of all processes for the specified project. A MongoDB process can be either a `mongod` or `mongos`. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListAtlasProcessesApiRequest
*/
func (a *MonitoringAndLogsApiService) ListAtlasProcesses(ctx context.Context, groupId string) ListAtlasProcessesApiRequest {
	return ListAtlasProcessesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListAtlasProcessesExecute executes the request
//
//	@return PaginatedHostViewAtlas
func (a *MonitoringAndLogsApiService) ListAtlasProcessesExecute(r ListAtlasProcessesApiRequest) (*PaginatedHostViewAtlas, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedHostViewAtlas
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListAtlasProcesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes"
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

type ListDatabasesApiRequest struct {
	ctx          context.Context
	ApiService   MonitoringAndLogsApi
	groupId      string
	processId    string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListDatabasesApiParams struct {
	GroupId      string
	ProcessId    string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *MonitoringAndLogsApiService) ListDatabasesWithParams(ctx context.Context, args *ListDatabasesApiParams) ListDatabasesApiRequest {
	return ListDatabasesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		processId:    args.ProcessId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListDatabasesApiRequest) IncludeCount(includeCount bool) ListDatabasesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListDatabasesApiRequest) ItemsPerPage(itemsPerPage int) ListDatabasesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListDatabasesApiRequest) PageNum(pageNum int) ListDatabasesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListDatabasesApiRequest) Execute() (*PaginatedDatabase, *http.Response, error) {
	return r.ApiService.ListDatabasesExecute(r)
}

/*
ListDatabases Return Available Databases for One MongoDB Process

Returns the list of databases running on the specified host for the specified project. `M0` free clusters, `M2`, `M5`, serverless, and Flex clusters have some operational limits. The MongoDB Cloud process must be a `mongod`. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListDatabasesApiRequest
*/
func (a *MonitoringAndLogsApiService) ListDatabases(ctx context.Context, groupId string, processId string) ListDatabasesApiRequest {
	return ListDatabasesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// ListDatabasesExecute executes the request
//
//	@return PaginatedDatabase
func (a *MonitoringAndLogsApiService) ListDatabasesExecute(r ListDatabasesApiRequest) (*PaginatedDatabase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedDatabase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListDatabases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/databases"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

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

type ListDiskMeasurementsApiRequest struct {
	ctx           context.Context
	ApiService    MonitoringAndLogsApi
	partitionName string
	groupId       string
	processId     string
}

type ListDiskMeasurementsApiParams struct {
	PartitionName string
	GroupId       string
	ProcessId     string
}

func (a *MonitoringAndLogsApiService) ListDiskMeasurementsWithParams(ctx context.Context, args *ListDiskMeasurementsApiParams) ListDiskMeasurementsApiRequest {
	return ListDiskMeasurementsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		partitionName: args.PartitionName,
		groupId:       args.GroupId,
		processId:     args.ProcessId,
	}
}

func (r ListDiskMeasurementsApiRequest) Execute() (*MeasurementDiskPartition, *http.Response, error) {
	return r.ApiService.ListDiskMeasurementsExecute(r)
}

/*
ListDiskMeasurements Return Measurements for One Disk

Returns measurement details for one disk or partition for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param partitionName Human-readable label of the disk or partition to which the measurements apply.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListDiskMeasurementsApiRequest
*/
func (a *MonitoringAndLogsApiService) ListDiskMeasurements(ctx context.Context, partitionName string, groupId string, processId string) ListDiskMeasurementsApiRequest {
	return ListDiskMeasurementsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		partitionName: partitionName,
		groupId:       groupId,
		processId:     processId,
	}
}

// ListDiskMeasurementsExecute executes the request
//
//	@return MeasurementDiskPartition
func (a *MonitoringAndLogsApiService) ListDiskMeasurementsExecute(r ListDiskMeasurementsApiRequest) (*MeasurementDiskPartition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MeasurementDiskPartition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListDiskMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}"
	if r.partitionName == "" {
		return localVarReturnValue, nil, reportError("partitionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"partitionName"+"}", url.PathEscape(r.partitionName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

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

type ListDiskPartitionsApiRequest struct {
	ctx          context.Context
	ApiService   MonitoringAndLogsApi
	groupId      string
	processId    string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListDiskPartitionsApiParams struct {
	GroupId      string
	ProcessId    string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *MonitoringAndLogsApiService) ListDiskPartitionsWithParams(ctx context.Context, args *ListDiskPartitionsApiParams) ListDiskPartitionsApiRequest {
	return ListDiskPartitionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		processId:    args.ProcessId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListDiskPartitionsApiRequest) IncludeCount(includeCount bool) ListDiskPartitionsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListDiskPartitionsApiRequest) ItemsPerPage(itemsPerPage int) ListDiskPartitionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListDiskPartitionsApiRequest) PageNum(pageNum int) ListDiskPartitionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListDiskPartitionsApiRequest) Execute() (*PaginatedDiskPartition, *http.Response, error) {
	return r.ApiService.ListDiskPartitionsExecute(r)
}

/*
ListDiskPartitions Return Available Disks for One MongoDB Process

Returns the list of disks or partitions for the specified host for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListDiskPartitionsApiRequest
*/
func (a *MonitoringAndLogsApiService) ListDiskPartitions(ctx context.Context, groupId string, processId string) ListDiskPartitionsApiRequest {
	return ListDiskPartitionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// ListDiskPartitionsExecute executes the request
//
//	@return PaginatedDiskPartition
func (a *MonitoringAndLogsApiService) ListDiskPartitionsExecute(r ListDiskPartitionsApiRequest) (*PaginatedDiskPartition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedDiskPartition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListDiskPartitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/disks"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)

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

type ListIndexMetricsApiRequest struct {
	ctx            context.Context
	ApiService     MonitoringAndLogsApi
	processId      string
	databaseName   string
	collectionName string
	groupId        string
	granularity    *string
	metrics        *[]string
	period         *string
	start          *time.Time
	end            *time.Time
}

type ListIndexMetricsApiParams struct {
	ProcessId      string
	DatabaseName   string
	CollectionName string
	GroupId        string
	Granularity    *string
	Metrics        *[]string
	Period         *string
	Start          *time.Time
	End            *time.Time
}

func (a *MonitoringAndLogsApiService) ListIndexMetricsWithParams(ctx context.Context, args *ListIndexMetricsApiParams) ListIndexMetricsApiRequest {
	return ListIndexMetricsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		processId:      args.ProcessId,
		databaseName:   args.DatabaseName,
		collectionName: args.CollectionName,
		groupId:        args.GroupId,
		granularity:    args.Granularity,
		metrics:        args.Metrics,
		period:         args.Period,
		start:          args.Start,
		end:            args.End,
	}
}

// Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
func (r ListIndexMetricsApiRequest) Granularity(granularity string) ListIndexMetricsApiRequest {
	r.granularity = &granularity
	return r
}

// List that contains the measurements that MongoDB Atlas reports for the associated data series.
func (r ListIndexMetricsApiRequest) Metrics(metrics []string) ListIndexMetricsApiRequest {
	r.metrics = &metrics
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r ListIndexMetricsApiRequest) Period(period string) ListIndexMetricsApiRequest {
	r.period = &period
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r ListIndexMetricsApiRequest) Start(start time.Time) ListIndexMetricsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r ListIndexMetricsApiRequest) End(end time.Time) ListIndexMetricsApiRequest {
	r.end = &end
	return r
}

func (r ListIndexMetricsApiRequest) Execute() (*MeasurementsIndexes, *http.Response, error) {
	return r.ApiService.ListIndexMetricsExecute(r)
}

/*
ListIndexMetrics Return All Atlas Search Index Metrics for One Namespace

Returns the Atlas Search index metrics within the specified time range for one namespace in the specified process.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@param databaseName Human-readable label that identifies the database.
	@param collectionName Human-readable label that identifies the collection.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListIndexMetricsApiRequest
*/
func (a *MonitoringAndLogsApiService) ListIndexMetrics(ctx context.Context, processId string, databaseName string, collectionName string, groupId string) ListIndexMetricsApiRequest {
	return ListIndexMetricsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		processId:      processId,
		databaseName:   databaseName,
		collectionName: collectionName,
		groupId:        groupId,
	}
}

// ListIndexMetricsExecute executes the request
//
//	@return MeasurementsIndexes
func (a *MonitoringAndLogsApiService) ListIndexMetricsExecute(r ListIndexMetricsApiRequest) (*MeasurementsIndexes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MeasurementsIndexes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListIndexMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements"
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError("granularity is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, reportError("metrics is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	{
		t := *r.metrics
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", t, "multi")
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

type ListMetricTypesApiRequest struct {
	ctx        context.Context
	ApiService MonitoringAndLogsApi
	processId  string
	groupId    string
}

type ListMetricTypesApiParams struct {
	ProcessId string
	GroupId   string
}

func (a *MonitoringAndLogsApiService) ListMetricTypesWithParams(ctx context.Context, args *ListMetricTypesApiParams) ListMetricTypesApiRequest {
	return ListMetricTypesApiRequest{
		ApiService: a,
		ctx:        ctx,
		processId:  args.ProcessId,
		groupId:    args.GroupId,
	}
}

func (r ListMetricTypesApiRequest) Execute() (*CloudSearchMetrics, *http.Response, error) {
	return r.ApiService.ListMetricTypesExecute(r)
}

/*
ListMetricTypes Return All Atlas Search Metric Types for One Process

Returns all Atlas Search metric types available for one process in the specified project. You must have the Project Read Only or higher role to view the Atlas Search metric types.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListMetricTypesApiRequest
*/
func (a *MonitoringAndLogsApiService) ListMetricTypes(ctx context.Context, processId string, groupId string) ListMetricTypesApiRequest {
	return ListMetricTypesApiRequest{
		ApiService: a,
		ctx:        ctx,
		processId:  processId,
		groupId:    groupId,
	}
}

// ListMetricTypesExecute executes the request
//
//	@return CloudSearchMetrics
func (a *MonitoringAndLogsApiService) ListMetricTypesExecute(r ListMetricTypesApiRequest) (*CloudSearchMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudSearchMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringAndLogsApiService.ListMetricTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics"
	if r.processId == "" {
		return localVarReturnValue, nil, reportError("processId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(r.processId), -1)
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
