# \PerformanceAdvisorApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableSlowOperationThresholding**](PerformanceAdvisorApi.md#DisableSlowOperationThresholding) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold
[**EnableSlowOperationThresholding**](PerformanceAdvisorApi.md#EnableSlowOperationThresholding) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold
[**ListSlowQueries**](PerformanceAdvisorApi.md#ListSlowQueries) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
[**ListSlowQueryNamespaces**](PerformanceAdvisorApi.md#ListSlowQueryNamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
[**ListSuggestedIndexes**](PerformanceAdvisorApi.md#ListSuggestedIndexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes



## DisableSlowOperationThresholding

> DisableSlowOperationThresholding(ctx, groupId).Execute()

Disable Managed Slow Operation Threshold


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.PerformanceAdvisorApi.DisableSlowOperationThresholding(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.DisableSlowOperationThresholding``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSlowOperationThresholdingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableSlowOperationThresholding

> EnableSlowOperationThresholding(ctx, groupId).Execute()

Enable Managed Slow Operation Threshold


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.PerformanceAdvisorApi.EnableSlowOperationThresholding(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.EnableSlowOperationThresholding``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSlowOperationThresholdingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlowQueries

> PerformanceAdvisorSlowQueryList ListSlowQueries(ctx, groupId, processId).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()

Return Slow Queries


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "processId_example" // string | 
    duration := int64(789) // int64 |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    nLogs := int64(789) // int64 |  (optional) (default to 20000)
    since := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListSlowQueries(context.Background(), groupId, processId).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSlowQueries``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListSlowQueries`: PerformanceAdvisorSlowQueryList
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSlowQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **duration** | **int64** | Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nLogs** | **int64** | Maximum number of lines from the log to return. | [default to 20000]
 **since** | **int64** | Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**PerformanceAdvisorSlowQueryList**](PerformanceAdvisorSlowQueryList.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlowQueryNamespaces

> Namespaces ListSlowQueryNamespaces(ctx, groupId, processId).Duration(duration).Since(since).Execute()

Return All Namespaces for One Host


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "processId_example" // string | 
    duration := int64(789) // int64 |  (optional)
    since := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListSlowQueryNamespaces(context.Background(), groupId, processId).Duration(duration).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSlowQueryNamespaces``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListSlowQueryNamespaces`: Namespaces
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSlowQueryNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlowQueryNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **duration** | **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**Namespaces**](Namespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSuggestedIndexes

> PerformanceAdvisorResponse ListSuggestedIndexes(ctx, groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()

Return Suggested Indexes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    processId := "processId_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    duration := int64(789) // int64 |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    nExamples := int64(789) // int64 |  (optional) (default to 5)
    nIndexes := int64(789) // int64 |  (optional)
    since := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListSuggestedIndexes(context.Background(), groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSuggestedIndexes``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListSuggestedIndexes`: PerformanceAdvisorResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSuggestedIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSuggestedIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **duration** | **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nExamples** | **int64** | Maximum number of example queries that benefit from the suggested index. | [default to 5]
 **nIndexes** | **int64** | Number that indicates the maximum indexes to suggest. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**PerformanceAdvisorResponse**](PerformanceAdvisorResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

