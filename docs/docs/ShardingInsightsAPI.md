# \ShardingInsightsAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterShardMetric**](ShardingInsightsAPI.md#GetClusterShardMetric) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics/{replicaSetName} | Return One Shard in One Cluster
[**GetClusterShardingInsights**](ShardingInsightsAPI.md#GetClusterShardingInsights) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardingInsights | Return Sharding Insights for One Cluster
[**GetShardMetricMeasurements**](ShardingInsightsAPI.md#GetShardMetricMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics/{replicaSetName}/measurements | Return Measurements for One Shard
[**ListClusterShardMetrics**](ShardingInsightsAPI.md#ListClusterShardMetrics) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics | Return All Shards in One Cluster
[**ListShardMetricMeasurements**](ShardingInsightsAPI.md#ListShardMetricMeasurements) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardMetrics:listMeasurements | Return Measurements for All Shards of One Cluster



## GetClusterShardMetric

> ShardMetricsResponse GetClusterShardMetric(ctx, groupId, clusterName, replicaSetName).Execute()

Return One Shard in One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312028/admin"
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
    clusterName := "clusterName_example" // string | 
    replicaSetName := "replicaSetName_example" // string | 

    resp, r, err := sdk.ShardingInsightsAPI.GetClusterShardMetric(context.Background(), groupId, clusterName, replicaSetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardingInsightsAPI.GetClusterShardMetric`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterShardMetric`: ShardMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardingInsightsAPI.GetClusterShardMetric`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**replicaSetName** | **string** | Name of the replica set that identifies the shard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterShardMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ShardMetricsResponse**](ShardMetricsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterShardingInsights

> ShardingInsightsResponse GetClusterShardingInsights(ctx, groupId, clusterName).Execute()

Return Sharding Insights for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312028/admin"
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
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.ShardingInsightsAPI.GetClusterShardingInsights(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardingInsightsAPI.GetClusterShardingInsights`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterShardingInsights`: ShardingInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardingInsightsAPI.GetClusterShardingInsights`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterShardingInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShardingInsightsResponse**](ShardingInsightsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShardMetricMeasurements

> ShardMeasurementsResponse GetShardMetricMeasurements(ctx, groupId, clusterName, replicaSetName).M(m).Granularity(granularity).Period(period).Start(start).End(end).NodeType(nodeType).Execute()

Return Measurements for One Shard


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312028/admin"
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
    clusterName := "clusterName_example" // string | 
    replicaSetName := "atlas-12vorz-shard-0" // string | 
    m := []string{"Inner_example"} // []string |  (optional)
    granularity := "PT1M" // string |  (optional)
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    nodeType := "nodeType_example" // string |  (optional) (default to "ALL")

    resp, r, err := sdk.ShardingInsightsAPI.GetShardMetricMeasurements(context.Background(), groupId, clusterName, replicaSetName).M(m).Granularity(granularity).Period(period).Start(start).End(end).NodeType(nodeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardingInsightsAPI.GetShardMetricMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetShardMetricMeasurements`: ShardMeasurementsResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardingInsightsAPI.GetShardMetricMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**replicaSetName** | **string** | Human-readable label that identifies the shard, as the replica set name of its members. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShardMetricMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **m** | **[]string** | One or more measurements to return. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. | 
 **granularity** | **string** | Interval between consecutive data points, in ISO 8601 duration format. If omitted, MongoDB Cloud chooses the resolution from the width of the window. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **nodeType** | **string** | Members of the shard whose readings are averaged into the shard&#39;s series. | [default to &quot;ALL&quot;]

### Return type

[**ShardMeasurementsResponse**](ShardMeasurementsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterShardMetrics

> PaginatedShardMetricsResponse ListClusterShardMetrics(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Shards in One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312028/admin"
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
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ShardingInsightsAPI.ListClusterShardMetrics(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardingInsightsAPI.ListClusterShardMetrics`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListClusterShardMetrics`: PaginatedShardMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardingInsightsAPI.ListClusterShardMetrics`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterShardMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether MongoDB Cloud calculates the total number of items for the response. When set to &#x60;false&#x60;, MongoDB Cloud may skip an additional count operation. The response may still include &#x60;totalCount&#x60; when the count is available without additional calculation. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedShardMetricsResponse**](PaginatedShardMetricsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShardMetricMeasurements

> ClusterShardMeasurementsResponse ListShardMetricMeasurements(ctx, groupId, clusterName).M(m).Granularity(granularity).Period(period).Start(start).End(end).NodeType(nodeType).Execute()

Return Measurements for All Shards of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312028/admin"
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
    clusterName := "clusterName_example" // string | 
    m := []string{"Inner_example"} // []string |  (optional)
    granularity := "PT1M" // string |  (optional)
    period := "PT10H" // string |  (optional)
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    nodeType := "nodeType_example" // string |  (optional) (default to "ALL")

    resp, r, err := sdk.ShardingInsightsAPI.ListShardMetricMeasurements(context.Background(), groupId, clusterName).M(m).Granularity(granularity).Period(period).Start(start).End(end).NodeType(nodeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardingInsightsAPI.ListShardMetricMeasurements`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListShardMetricMeasurements`: ClusterShardMeasurementsResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardingInsightsAPI.ListShardMetricMeasurements`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShardMetricMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **m** | **[]string** | One or more measurements to return. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. | 
 **granularity** | **string** | Interval between consecutive data points, in ISO 8601 duration format. If omitted, MongoDB Cloud chooses the resolution from the width of the window. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **nodeType** | **string** | Members of each shard whose readings are averaged into the shard&#39;s series. | [default to &quot;ALL&quot;]

### Return type

[**ClusterShardMeasurementsResponse**](ClusterShardMeasurementsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

