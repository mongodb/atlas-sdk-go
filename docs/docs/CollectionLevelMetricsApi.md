# \CollectionLevelMetricsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCollStatsLatencyNamespaceClusterMeasurements**](CollectionLevelMetricsApi.md#GetCollStatsLatencyNamespaceClusterMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements | Return Cluster-Level Query Latency
[**GetCollStatsLatencyNamespaceHostMeasurements**](CollectionLevelMetricsApi.md#GetCollStatsLatencyNamespaceHostMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/{databaseName}/{collectionName}/collStats/measurements | Return Host-Level Query Latency
[**GetCollStatsLatencyNamespaceMetrics**](CollectionLevelMetricsApi.md#GetCollStatsLatencyNamespaceMetrics) | **Get** /api/atlas/v2/groups/{groupId}/collStats/metrics | Return all metric names
[**GetCollStatsLatencyNamespacesForCluster**](CollectionLevelMetricsApi.md#GetCollStatsLatencyNamespacesForCluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/collStats/namespaces | Return Ranked Namespaces from a Cluster
[**GetCollStatsLatencyNamespacesForHost**](CollectionLevelMetricsApi.md#GetCollStatsLatencyNamespacesForHost) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/collStats/namespaces | Return Ranked Namespaces from a Host
[**GetPinnedNamespaces**](CollectionLevelMetricsApi.md#GetPinnedNamespaces) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Return Pinned Namespaces
[**PinNamespacesPatch**](CollectionLevelMetricsApi.md#PinNamespacesPatch) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Add Pinned Namespaces
[**PinNamespacesPut**](CollectionLevelMetricsApi.md#PinNamespacesPut) | **Put** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned | Pin Namespaces
[**UnpinNamespaces**](CollectionLevelMetricsApi.md#UnpinNamespaces) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/unpin | Unpin namespaces



## GetCollStatsLatencyNamespaceClusterMeasurements

> MeasurementsCollStatsLatencyCluster GetCollStatsLatencyNamespaceClusterMeasurements(ctx, groupId, clusterName, clusterView, databaseName, collectionName).Metrics(metrics).Start(start).End(end).Period(period).Execute()

Return Cluster-Level Query Latency


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    clusterView := "clusterView_example" // string | 
    databaseName := "mydb" // string | 
    collectionName := "mycoll" // string | 
    metrics := []string{"Inner_example"} // []string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    period := "PT10H" // string |  (optional)

    resp, r, err := sdk.CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceClusterMeasurements(context.Background(), groupId, clusterName, clusterView, databaseName, collectionName).Metrics(metrics).Start(start).End(end).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceClusterMeasurements``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCollStatsLatencyNamespaceClusterMeasurements`: MeasurementsCollStatsLatencyCluster
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceClusterMeasurements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to retrieve metrics for. | 
**clusterView** | **string** | Human-readable label that identifies the cluster topology to retrieve metrics for. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollStatsLatencyNamespaceClusterMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **metrics** | **[]string** | List that contains the metrics that you want to retrieve for the associated data series. If you don&#39;t set this parameter, this resource returns data series for all Coll Stats Latency metrics. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 

### Return type

[**MeasurementsCollStatsLatencyCluster**](MeasurementsCollStatsLatencyCluster.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollStatsLatencyNamespaceHostMeasurements

> MeasurementsCollStatsLatencyHost GetCollStatsLatencyNamespaceHostMeasurements(ctx, groupId, processId, databaseName, collectionName).Metrics(metrics).Start(start).End(end).Period(period).Execute()

Return Host-Level Query Latency


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "my.host.name.com:27017" // string | 
    databaseName := "mydb" // string | 
    collectionName := "mycoll" // string | 
    metrics := []string{"Inner_example"} // []string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    period := "PT10H" // string |  (optional)

    resp, r, err := sdk.CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceHostMeasurements(context.Background(), groupId, processId, databaseName, collectionName).Metrics(metrics).Start(start).End(end).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceHostMeasurements``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCollStatsLatencyNamespaceHostMeasurements`: MeasurementsCollStatsLatencyHost
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceHostMeasurements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollStatsLatencyNamespaceHostMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **metrics** | **[]string** | List that contains the metrics that you want to retrieve for the associated data series. If you don&#39;t set this parameter, this resource returns data series for all Coll Stats Latency metrics. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 

### Return type

[**MeasurementsCollStatsLatencyHost**](MeasurementsCollStatsLatencyHost.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollStatsLatencyNamespaceMetrics

> map[string]interface{} GetCollStatsLatencyNamespaceMetrics(ctx, groupId).Execute()

Return all metric names


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceMetrics(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceMetrics``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCollStatsLatencyNamespaceMetrics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetCollStatsLatencyNamespaceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollStatsLatencyNamespaceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollStatsLatencyNamespacesForCluster

> CollStatsRankedNamespaces GetCollStatsLatencyNamespacesForCluster(ctx, groupId, clusterName, clusterView).Start(start).End(end).Period(period).Execute()

Return Ranked Namespaces from a Cluster


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    clusterView := "clusterView_example" // string | 
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    period := "PT10H" // string |  (optional)

    resp, r, err := sdk.CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForCluster(context.Background(), groupId, clusterName, clusterView).Start(start).End(end).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForCluster``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCollStatsLatencyNamespacesForCluster`: CollStatsRankedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to pin namespaces to. | 
**clusterView** | **string** | Human-readable label that identifies the cluster topology to retrieve metrics for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollStatsLatencyNamespacesForClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 

### Return type

[**CollStatsRankedNamespaces**](CollStatsRankedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollStatsLatencyNamespacesForHost

> CollStatsRankedNamespaces GetCollStatsLatencyNamespacesForHost(ctx, groupId, processId).Start(start).End(end).Period(period).Execute()

Return Ranked Namespaces from a Host


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "my.host.name.com:27017" // string | 
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    period := "PT10H" // string |  (optional)

    resp, r, err := sdk.CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForHost(context.Background(), groupId, processId).Start(start).End(end).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForHost``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCollStatsLatencyNamespacesForHost`: CollStatsRankedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetCollStatsLatencyNamespacesForHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollStatsLatencyNamespacesForHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 

### Return type

[**CollStatsRankedNamespaces**](CollStatsRankedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPinnedNamespaces

> PinnedNamespaces GetPinnedNamespaces(ctx, groupId, clusterName).Execute()

Return Pinned Namespaces


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.CollectionLevelMetricsApi.GetPinnedNamespaces(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.GetPinnedNamespaces``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetPinnedNamespaces`: PinnedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.GetPinnedNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to retrieve pinned namespaces for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinnedNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PinnedNamespaces**](PinnedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinNamespacesPatch

> PinnedNamespaces PinNamespacesPatch(ctx, groupId, clusterName, namespacesRequest NamespacesRequest).Execute()

Add Pinned Namespaces


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    namespacesRequest := *openapiclient.NewNamespacesRequest() // NamespacesRequest | 

    resp, r, err := sdk.CollectionLevelMetricsApi.PinNamespacesPatch(context.Background(), groupId, clusterName, &namespacesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.PinNamespacesPatch``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `PinNamespacesPatch`: PinnedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.PinNamespacesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to pin namespaces to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinNamespacesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespacesRequest** | [**NamespacesRequest**](NamespacesRequest.md) | List of namespace strings (combination of database and collection name) to pin for query latency metric collection. | 

### Return type

[**PinnedNamespaces**](PinnedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-11-15+json
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinNamespacesPut

> PinnedNamespaces PinNamespacesPut(ctx, groupId, clusterName, namespacesRequest NamespacesRequest).Execute()

Pin Namespaces


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    namespacesRequest := *openapiclient.NewNamespacesRequest() // NamespacesRequest | 

    resp, r, err := sdk.CollectionLevelMetricsApi.PinNamespacesPut(context.Background(), groupId, clusterName, &namespacesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.PinNamespacesPut``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `PinNamespacesPut`: PinnedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.PinNamespacesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to pin namespaces to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinNamespacesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespacesRequest** | [**NamespacesRequest**](NamespacesRequest.md) | List of namespace strings (combination of database and collection name) to pin for query latency metric collection. | 

### Return type

[**PinnedNamespaces**](PinnedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-11-15+json
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpinNamespaces

> PinnedNamespaces UnpinNamespaces(ctx, groupId, clusterName, namespacesRequest NamespacesRequest).Execute()

Unpin namespaces


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    namespacesRequest := *openapiclient.NewNamespacesRequest() // NamespacesRequest | 

    resp, r, err := sdk.CollectionLevelMetricsApi.UnpinNamespaces(context.Background(), groupId, clusterName, &namespacesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionLevelMetricsApi.UnpinNamespaces``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UnpinNamespaces`: PinnedNamespaces
    fmt.Fprintf(os.Stdout, "Response from `CollectionLevelMetricsApi.UnpinNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to unpin namespaces from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpinNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespacesRequest** | [**NamespacesRequest**](NamespacesRequest.md) | List of namespace strings (combination of database and collection name) to pin for query latency metric collection. | 

### Return type

[**PinnedNamespaces**](PinnedNamespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-11-15+json
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

