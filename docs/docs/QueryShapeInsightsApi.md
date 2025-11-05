# \QueryShapeInsightsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryShapeDetails**](QueryShapeInsightsApi.md#GetQueryShapeDetails) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/{queryShapeHash}/details | Return Query Shape Details
[**ListQueryShapeSummaries**](QueryShapeInsightsApi.md#ListQueryShapeSummaries) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/queryShapeInsights/summaries | Return Query Statistic Summaries



## GetQueryShapeDetails

> QueryStatsDetailsResponse GetQueryShapeDetails(ctx, groupId, clusterName, queryShapeHash).Since(since).Until(until).ProcessIds(processIds).Execute()

Return Query Shape Details


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312009/admin"
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
    queryShapeHash := "queryShapeHash_example" // string | 
    since := int64(789) // int64 |  (optional)
    until := int64(789) // int64 |  (optional)
    processIds := []string{"Inner_example"} // []string |  (optional)

    resp, r, err := sdk.QueryShapeInsightsApi.GetQueryShapeDetails(context.Background(), groupId, clusterName, queryShapeHash).Since(since).Until(until).ProcessIds(processIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryShapeInsightsApi.GetQueryShapeDetails`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetQueryShapeDetails`: QueryStatsDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryShapeInsightsApi.GetQueryShapeDetails`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**queryShapeHash** | **string** | A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryShapeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **int64** | Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 
 **until** | **int64** | Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you specify the **until** parameter, you must specify the **since** parameter. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 
 **processIds** | **[]string** | ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each processId. | 

### Return type

[**QueryStatsDetailsResponse**](QueryStatsDetailsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryShapeSummaries

> QueryStatsSummaryListResponse ListQueryShapeSummaries(ctx, groupId, clusterName).Since(since).Until(until).ProcessIds(processIds).Namespaces(namespaces).Commands(commands).NSummaries(nSummaries).Series(series).QueryShapeHashes(queryShapeHashes).Execute()

Return Query Statistic Summaries


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312009/admin"
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
    since := int64(789) // int64 |  (optional)
    until := int64(789) // int64 |  (optional)
    processIds := []string{"Inner_example"} // []string |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    commands := []string{"Inner_example"} // []string |  (optional)
    nSummaries := int64(789) // int64 |  (optional) (default to 100)
    series := []string{"Inner_example"} // []string |  (optional)
    queryShapeHashes := []string{"Inner_example"} // []string |  (optional)

    resp, r, err := sdk.QueryShapeInsightsApi.ListQueryShapeSummaries(context.Background(), groupId, clusterName).Since(since).Until(until).ProcessIds(processIds).Namespaces(namespaces).Commands(commands).NSummaries(nSummaries).Series(series).QueryShapeHashes(queryShapeHashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryShapeInsightsApi.ListQueryShapeSummaries`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListQueryShapeSummaries`: QueryStatsSummaryListResponse
    fmt.Fprintf(os.Stdout, "Response from `QueryShapeInsightsApi.ListQueryShapeSummaries`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryShapeSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int64** | Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 
 **until** | **int64** | Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you specify the **until** parameter, you must specify the **since** parameter. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 
 **processIds** | **[]string** | ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each processId. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve query shape statistics. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **commands** | **[]string** | Retrieve query shape statistics matching specified MongoDB commands. To include multiple commands, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each command. The currently supported parameters are find, distinct, and aggregate. Omit this parameter to return results for all supported commands. | 
 **nSummaries** | **int64** | Maximum number of query statistic summaries to return. | [default to 100]
 **series** | **[]string** | Query shape statistics data series to retrieve. A series represents a specific metric about query execution. To include multiple series, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each series. Omit this parameter to return results for all available series. | 
 **queryShapeHashes** | **[]string** | A list of SHA256 hashes of desired query shapes, output by MongoDB commands like $queryStats and $explain or slow query logs. To include multiple series, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each series. Omit this parameter to return results for all available series. | 

### Return type

[**QueryStatsSummaryListResponse**](QueryStatsSummaryListResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

