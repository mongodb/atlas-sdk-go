# \PerformanceAdvisorApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableManagedSlowMs**](PerformanceAdvisorApi.md#DisableManagedSlowMs) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold
[**EnableManagedSlowMs**](PerformanceAdvisorApi.md#EnableManagedSlowMs) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold
[**GetManagedSlowMs**](PerformanceAdvisorApi.md#GetManagedSlowMs) | **Get** /api/atlas/v2/groups/{groupId}/managedSlowMs | Return Managed Slow Operation Threshold Status
[**GetServerlessAutoIndexing**](PerformanceAdvisorApi.md#GetServerlessAutoIndexing) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Return Serverless Auto-Indexing Status
[**ListClusterSuggestedIndexes**](PerformanceAdvisorApi.md#ListClusterSuggestedIndexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes
[**ListDropIndexSuggestions**](PerformanceAdvisorApi.md#ListDropIndexSuggestions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/dropIndexSuggestions | Return All Suggested Indexes to Drop
[**ListPerformanceAdvisorNamespaces**](PerformanceAdvisorApi.md#ListPerformanceAdvisorNamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
[**ListSchemaAdvice**](PerformanceAdvisorApi.md#ListSchemaAdvice) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/performanceAdvisor/schemaAdvice | Return Schema Advice
[**ListSlowQueryLogs**](PerformanceAdvisorApi.md#ListSlowQueryLogs) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
[**ListSuggestedIndexes**](PerformanceAdvisorApi.md#ListSuggestedIndexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return All Suggested Indexes
[**SetServerlessAutoIndexing**](PerformanceAdvisorApi.md#SetServerlessAutoIndexing) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing | Set Serverless Auto-Indexing Status



## DisableManagedSlowMs

> DisableManagedSlowMs(ctx, groupId).Execute()

Disable Managed Slow Operation Threshold


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    r, err := sdk.PerformanceAdvisorApi.DisableManagedSlowMs(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.DisableManagedSlowMs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableManagedSlowMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableManagedSlowMs

> EnableManagedSlowMs(ctx, groupId).Execute()

Enable Managed Slow Operation Threshold


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    r, err := sdk.PerformanceAdvisorApi.EnableManagedSlowMs(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.EnableManagedSlowMs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableManagedSlowMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSlowMs

> bool GetManagedSlowMs(ctx, groupId).Execute()

Return Managed Slow Operation Threshold Status


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    resp, r, err := sdk.PerformanceAdvisorApi.GetManagedSlowMs(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.GetManagedSlowMs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetManagedSlowMs`: bool
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.GetManagedSlowMs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSlowMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerlessAutoIndexing

> bool GetServerlessAutoIndexing(ctx, groupId, clusterName).Execute()

Return Serverless Auto-Indexing Status


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    resp, r, err := sdk.PerformanceAdvisorApi.GetServerlessAutoIndexing(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.GetServerlessAutoIndexing`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetServerlessAutoIndexing`: bool
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.GetServerlessAutoIndexing`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessAutoIndexingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterSuggestedIndexes

> PerformanceAdvisorResponse ListClusterSuggestedIndexes(ctx, groupId, clusterName).ProcessIds(processIds).Namespaces(namespaces).Since(since).Until(until).Execute()

Return All Suggested Indexes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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
    processIds := []string{"Inner_example"} // []string |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    since := int64(789) // int64 |  (optional)
    until := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListClusterSuggestedIndexes(context.Background(), groupId, clusterName).ProcessIds(processIds).Namespaces(namespaces).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListClusterSuggestedIndexes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListClusterSuggestedIndexes`: PerformanceAdvisorResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListClusterSuggestedIndexes`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterSuggestedIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **processIds** | **[]string** | Process IDs from which to retrieve suggested indexes. A &#x60;processId&#x60; is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple &#x60;processIds&#x60;, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each &#x60;processId&#x60;. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 
 **until** | **int64** | Date and time up until which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you specify the **until** parameter, you must specify the **since** parameter. - If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**PerformanceAdvisorResponse**](PerformanceAdvisorResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDropIndexSuggestions

> DropIndexSuggestionsResponse ListDropIndexSuggestions(ctx, groupId, clusterName).Execute()

Return All Suggested Indexes to Drop


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    resp, r, err := sdk.PerformanceAdvisorApi.ListDropIndexSuggestions(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListDropIndexSuggestions`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDropIndexSuggestions`: DropIndexSuggestionsResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListDropIndexSuggestions`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDropIndexSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DropIndexSuggestionsResponse**](DropIndexSuggestionsResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPerformanceAdvisorNamespaces

> Namespaces ListPerformanceAdvisorNamespaces(ctx, groupId, processId).Duration(duration).Since(since).Execute()

Return All Namespaces for One Host


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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
    processId := "processId_example" // string | 
    duration := int64(789) // int64 |  (optional)
    since := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListPerformanceAdvisorNamespaces(context.Background(), groupId, processId).Duration(duration).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListPerformanceAdvisorNamespaces`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPerformanceAdvisorNamespaces`: Namespaces
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListPerformanceAdvisorNamespaces`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPerformanceAdvisorNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **duration** | **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**Namespaces**](Namespaces.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemaAdvice

> SchemaAdvisorResponse ListSchemaAdvice(ctx, groupId, clusterName).Execute()

Return Schema Advice


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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

    resp, r, err := sdk.PerformanceAdvisorApi.ListSchemaAdvice(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSchemaAdvice`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSchemaAdvice`: SchemaAdvisorResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSchemaAdvice`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemaAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SchemaAdvisorResponse**](SchemaAdvisorResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlowQueryLogs

> PerformanceAdvisorSlowQueryList ListSlowQueryLogs(ctx, groupId, processId).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).IncludeMetrics(includeMetrics).IncludeReplicaState(includeReplicaState).IncludeOpType(includeOpType).Execute()

Return Slow Queries


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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
    processId := "processId_example" // string | 
    duration := int64(789) // int64 |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    nLogs := int64(789) // int64 |  (optional) (default to 20000)
    since := int64(789) // int64 |  (optional)
    includeMetrics := true // bool |  (optional) (default to false)
    includeReplicaState := true // bool |  (optional) (default to false)
    includeOpType := true // bool |  (optional) (default to false)

    resp, r, err := sdk.PerformanceAdvisorApi.ListSlowQueryLogs(context.Background(), groupId, processId).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).IncludeMetrics(includeMetrics).IncludeReplicaState(includeReplicaState).IncludeOpType(includeOpType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSlowQueryLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSlowQueryLogs`: PerformanceAdvisorSlowQueryList
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSlowQueryLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlowQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **duration** | **int64** | Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nLogs** | **int64** | Maximum number of lines from the log to return. | [default to 20000]
 **since** | **int64** | Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **includeMetrics** | **bool** | Whether or not to include metrics extracted from the slow query log as separate fields. | [default to false]
 **includeReplicaState** | **bool** | Whether or not to include the replica state of the host when the slow query log was generated as a separate field. | [default to false]
 **includeOpType** | **bool** | Whether or not to include the operation type (read/write/command) extracted from the slow query log as a separate field. | [default to false]

### Return type

[**PerformanceAdvisorSlowQueryList**](PerformanceAdvisorSlowQueryList.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSuggestedIndexes

> PerformanceAdvisorResponse ListSuggestedIndexes(ctx, groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()

Return All Suggested Indexes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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
    processId := "processId_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    duration := int64(789) // int64 |  (optional)
    namespaces := []string{"Inner_example"} // []string |  (optional)
    nExamples := int64(789) // int64 |  (optional) (default to 5)
    nIndexes := int64(789) // int64 |  (optional)
    since := int64(789) // int64 |  (optional)

    resp, r, err := sdk.PerformanceAdvisorApi.ListSuggestedIndexes(context.Background(), groupId, processId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ListSuggestedIndexes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSuggestedIndexes`: PerformanceAdvisorResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ListSuggestedIndexes`: %v (%v)\n", resp, r)
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


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **duration** | **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nExamples** | **int64** | Maximum number of example queries that benefit from the suggested index. | [default to 5]
 **nIndexes** | **int64** | Number that indicates the maximum indexes to suggest. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**PerformanceAdvisorResponse**](PerformanceAdvisorResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServerlessAutoIndexing

> SetServerlessAutoIndexing(ctx, groupId, clusterName).Enable(enable).Execute()

Set Serverless Auto-Indexing Status


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312015/admin"
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
    enable := true // bool | 

    r, err := sdk.PerformanceAdvisorApi.SetServerlessAutoIndexing(context.Background(), groupId, clusterName).Enable(enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.SetServerlessAutoIndexing`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServerlessAutoIndexingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enable** | **bool** | Value that we want to set for the Serverless Auto Indexing toggle. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

