# \MonitoringAndLogsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAtlasProcess**](MonitoringAndLogsApi.md#GetAtlasProcess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID
[**GetDatabase**](MonitoringAndLogsApi.md#GetDatabase) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for a MongoDB Process
[**GetDatabaseMeasurements**](MonitoringAndLogsApi.md#GetDatabaseMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements of One Database for One MongoDB Process
[**GetDiskMeasurements**](MonitoringAndLogsApi.md#GetDiskMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process
[**GetHostLogs**](MonitoringAndLogsApi.md#GetHostLogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Cluster Host in One Project
[**GetHostMeasurements**](MonitoringAndLogsApi.md#GetHostMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process
[**GetIndexMetrics**](MonitoringAndLogsApi.md#GetIndexMetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Specified Namespace
[**GetMeasurements**](MonitoringAndLogsApi.md#GetMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics
[**ListAtlasProcesses**](MonitoringAndLogsApi.md#ListAtlasProcesses) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project
[**ListDatabases**](MonitoringAndLogsApi.md#ListDatabases) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process
[**ListDiskMeasurements**](MonitoringAndLogsApi.md#ListDiskMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements of One Disk
[**ListDiskPartitions**](MonitoringAndLogsApi.md#ListDiskPartitions) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process
[**ListIndexMetrics**](MonitoringAndLogsApi.md#ListIndexMetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace
[**ListMetricTypes**](MonitoringAndLogsApi.md#ListMetricTypes) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process



## GetAtlasProcess

> ApiHostViewAtlas GetAtlasProcess(ctx, groupId, processId).Execute()

Return One MongoDB Process by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "mongodb.example.com:27017" // string | 

    resp, r, err := sdk.MonitoringAndLogsApi.GetAtlasProcess(context.Background(), groupId, processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetAtlasProcess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasProcess`: ApiHostViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetAtlasProcess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiHostViewAtlas**](ApiHostViewAtlas.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> MesurementsDatabase GetDatabase(ctx, groupId, databaseName, processId).Execute()

Return One Database for a MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    databaseName := "databaseName_example" // string | 
    processId := "mongodb.example.com:27017" // string | 

    resp, r, err := sdk.MonitoringAndLogsApi.GetDatabase(context.Background(), groupId, databaseName, processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetDatabase`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDatabase`: MesurementsDatabase
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetDatabase`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**databaseName** | **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MesurementsDatabase**](MesurementsDatabase.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMeasurements

> ApiMeasurementsGeneralViewAtlas GetDatabaseMeasurements(ctx, groupId, databaseName, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()

Return Measurements of One Database for One MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    databaseName := "databaseName_example" // string | 
    processId := "mongodb.example.com:27017" // string | 
    granularity := "PT1M" // string | 
    m := []string{"Inner_example"} // []string |  (optional)
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetDatabaseMeasurements(context.Background(), groupId, databaseName, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetDatabaseMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDatabaseMeasurements`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetDatabaseMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**databaseName** | **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiskMeasurements

> ApiMeasurementsGeneralViewAtlas GetDiskMeasurements(ctx, groupId, partitionName, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()

Return Measurements of One Disk for One MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    partitionName := "partitionName_example" // string | 
    processId := "mongodb.example.com:27017" // string | 
    granularity := "PT1M" // string | 
    m := []string{"Inner_example"} // []string |  (optional)
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetDiskMeasurements(context.Background(), groupId, partitionName, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetDiskMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDiskMeasurements`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetDiskMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**partitionName** | **string** | Human-readable label of the disk or partition to which the measurements apply. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostLogs

> io.ReadCloser GetHostLogs(ctx, groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()

Download Logs for One Cluster Host in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    hostName := "hostName_example" // string | 
    logName := "logName_example" // string | 
    endDate := int64(789) // int64 |  (optional)
    startDate := int64(789) // int64 |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetHostLogs(context.Background(), groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetHostLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetHostLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetHostLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**hostName** | **string** | Human-readable label that identifies the host that stores the log files that you want to download. | 
**logName** | **string** | Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **endDate** | **int64** | Date and time when the period specifies the inclusive ending point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Date and time when the period specifies the inclusive starting point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostMeasurements

> ApiMeasurementsGeneralViewAtlas GetHostMeasurements(ctx, groupId, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()

Return Measurements for One MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "mongodb.example.com:27017" // string | 
    granularity := "PT1M" // string | 
    m := []string{"Inner_example"} // []string |  (optional)
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetHostMeasurements(context.Background(), groupId, processId).Granularity(granularity).M(m).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetHostMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetHostMeasurements`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetHostMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexMetrics

> MeasurementsIndexes GetIndexMetrics(ctx, processId, indexName, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()

Return Atlas Search Metrics for One Index in One Specified Namespace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    processId := "my.host.name.com:27017" // string | 
    indexName := "myindex" // string | 
    databaseName := "mydb" // string | 
    collectionName := "mycoll" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    granularity := "PT1M" // string | 
    metrics := []string{"Inner_example"} // []string | 
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetIndexMetrics(context.Background(), processId, indexName, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetIndexMetrics`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetIndexMetrics`: MeasurementsIndexes
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetIndexMetrics`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**indexName** | **string** | Human-readable label that identifies the index. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the measurements that MongoDB Atlas reports for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**MeasurementsIndexes**](MeasurementsIndexes.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeasurements

> MeasurementsNonIndex GetMeasurements(ctx, processId, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()

Return Atlas Search Hardware and Status Metrics


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    processId := "my.host.name.com:27017" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    granularity := "PT1M" // string | 
    metrics := []string{"Inner_example"} // []string | 
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.GetMeasurements(context.Background(), processId, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetMeasurements`: MeasurementsNonIndex
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don&#39;t set this parameter, this resource returns all hardware and status metrics for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**MeasurementsNonIndex**](MeasurementsNonIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtlasProcesses

> PaginatedHostViewAtlas ListAtlasProcesses(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All MongoDB Processes in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.MonitoringAndLogsApi.ListAtlasProcesses(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListAtlasProcesses`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAtlasProcesses`: PaginatedHostViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListAtlasProcesses`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtlasProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedHostViewAtlas**](PaginatedHostViewAtlas.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> PaginatedDatabase ListDatabases(ctx, groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return Available Databases for One MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "mongodb.example.com:27017" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.MonitoringAndLogsApi.ListDatabases(context.Background(), groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListDatabases`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDatabases`: PaginatedDatabase
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListDatabases`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedDatabase**](PaginatedDatabase.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskMeasurements

> MeasurementDiskPartition ListDiskMeasurements(ctx, partitionName, groupId, processId).Execute()

Return Measurements of One Disk


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    partitionName := "partitionName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "mongodb.example.com:27017" // string | 

    resp, r, err := sdk.MonitoringAndLogsApi.ListDiskMeasurements(context.Background(), partitionName, groupId, processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListDiskMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDiskMeasurements`: MeasurementDiskPartition
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListDiskMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Human-readable label of the disk or partition to which the measurements apply. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MeasurementDiskPartition**](MeasurementDiskPartition.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskPartitions

> PaginatedDiskPartition ListDiskPartitions(ctx, groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return Available Disks for One MongoDB Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "mongodb.example.com:27017" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.MonitoringAndLogsApi.ListDiskPartitions(context.Background(), groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListDiskPartitions`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDiskPartitions`: PaginatedDiskPartition
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListDiskPartitions`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedDiskPartition**](PaginatedDiskPartition.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexMetrics

> MeasurementsIndexes ListIndexMetrics(ctx, processId, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()

Return All Atlas Search Index Metrics for One Namespace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    processId := "my.host.name.com:27017" // string | 
    databaseName := "mydb" // string | 
    collectionName := "mycoll" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    granularity := "PT1M" // string | 
    metrics := []string{"Inner_example"} // []string | 
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)

    resp, r, err := sdk.MonitoringAndLogsApi.ListIndexMetrics(context.Background(), processId, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListIndexMetrics`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListIndexMetrics`: MeasurementsIndexes
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListIndexMetrics`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the measurements that MongoDB Atlas reports for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 

### Return type

[**MeasurementsIndexes**](MeasurementsIndexes.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricTypes

> CloudSearchMetrics ListMetricTypes(ctx, processId, groupId).Execute()

Return All Atlas Search Metric Types for One Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    processId := "my.host.name.com:27017" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.MonitoringAndLogsApi.ListMetricTypes(context.Background(), processId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ListMetricTypes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListMetricTypes`: CloudSearchMetrics
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ListMetricTypes`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetricTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudSearchMetrics**](CloudSearchMetrics.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

