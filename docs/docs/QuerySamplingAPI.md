# \QuerySamplingAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterQuerySampling**](QuerySamplingAPI.md#CreateClusterQuerySampling) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling | Create Query Sampling for One Namespace
[**DeleteClusterQuerySampling**](QuerySamplingAPI.md#DeleteClusterQuerySampling) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling/{namespace} | Remove Query Sampling from One Namespace
[**GetClusterQuerySampling**](QuerySamplingAPI.md#GetClusterQuerySampling) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/querySampling/{namespace} | Return Query Sampling State for One Namespace



## CreateClusterQuerySampling

> QuerySamplingResponse CreateClusterQuerySampling(ctx, groupId, clusterName, querySamplingCreateRequest QuerySamplingCreateRequest).Execute()

Create Query Sampling for One Namespace


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
    querySamplingCreateRequest := *admin.NewQuerySamplingCreateRequest("sales.orders") // QuerySamplingCreateRequest | 

    resp, r, err := sdk.QuerySamplingAPI.CreateClusterQuerySampling(context.Background(), groupId, clusterName, &querySamplingCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuerySamplingAPI.CreateClusterQuerySampling`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateClusterQuerySampling`: QuerySamplingResponse
    fmt.Fprintf(os.Stdout, "Response from `QuerySamplingAPI.CreateClusterQuerySampling`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterQuerySamplingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **querySamplingCreateRequest** | [**QuerySamplingCreateRequest**](QuerySamplingCreateRequest.md) | Namespace on which to start query sampling. | 

### Return type

[**QuerySamplingResponse**](QuerySamplingResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterQuerySampling

> DeleteClusterQuerySampling(ctx, groupId, clusterName, namespace).Execute()

Remove Query Sampling from One Namespace


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
    namespace := "sales.orders" // string | 

    r, err := sdk.QuerySamplingAPI.DeleteClusterQuerySampling(context.Background(), groupId, clusterName, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuerySamplingAPI.DeleteClusterQuerySampling`: %v (%v)\n", err, r)
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
**namespace** | **string** | Human-readable label that identifies the namespace (&#x60;&lt;database&gt;.&lt;collection&gt;&#x60;) on which queries are sampled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterQuerySamplingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterQuerySampling

> QuerySamplingResponse GetClusterQuerySampling(ctx, groupId, clusterName, namespace).Execute()

Return Query Sampling State for One Namespace


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
    namespace := "sales.orders" // string | 

    resp, r, err := sdk.QuerySamplingAPI.GetClusterQuerySampling(context.Background(), groupId, clusterName, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuerySamplingAPI.GetClusterQuerySampling`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterQuerySampling`: QuerySamplingResponse
    fmt.Fprintf(os.Stdout, "Response from `QuerySamplingAPI.GetClusterQuerySampling`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**namespace** | **string** | Human-readable label that identifies the namespace (&#x60;&lt;database&gt;.&lt;collection&gt;&#x60;) on which queries are sampled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterQuerySamplingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**QuerySamplingResponse**](QuerySamplingResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

