# \ShardKeyAnalyzerAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShardKeyAnalysis**](ShardKeyAnalyzerAPI.md#CreateShardKeyAnalysis) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses | Create One Shard Key Analysis for One Cluster
[**GetAnalysisOperation**](ShardKeyAnalyzerAPI.md#GetAnalysisOperation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/operations/{operationId} | Return One Shard Key Analysis Operation for One Cluster
[**GetShardKeyAnalysis**](ShardKeyAnalyzerAPI.md#GetShardKeyAnalysis) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/{analysisId} | Return One Shard Key Analysis for One Cluster
[**ListAnalysisOperations**](ShardKeyAnalyzerAPI.md#ListAnalysisOperations) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses/operations | Return All Shard Key Analysis Operations for One Cluster
[**ListShardKeyAnalyses**](ShardKeyAnalyzerAPI.md#ListShardKeyAnalyses) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/shardKeyAnalyses | Return All Shard Key Analyses for One Cluster



## CreateShardKeyAnalysis

> CreateShardKeyAnalysis(ctx, groupId, clusterName, shardKeyAnalysisRequest ShardKeyAnalysisRequest).Execute()

Create One Shard Key Analysis for One Cluster


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
    shardKeyAnalysisRequest := *admin.NewShardKeyAnalysisRequest("sales.orders", []map[string]string{map[string]string{"key": "Inner_example"}}) // ShardKeyAnalysisRequest | 

    r, err := sdk.ShardKeyAnalyzerAPI.CreateShardKeyAnalysis(context.Background(), groupId, clusterName, &shardKeyAnalysisRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardKeyAnalyzerAPI.CreateShardKeyAnalysis`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection to analyze. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShardKeyAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shardKeyAnalysisRequest** | [**ShardKeyAnalysisRequest**](ShardKeyAnalysisRequest.md) | Shard key to analyze, and which metrics to compute for it. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisOperation

> ShardKeyAnalysisOperationResponse GetAnalysisOperation(ctx, groupId, clusterName, operationId).Execute()

Return One Shard Key Analysis Operation for One Cluster


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
    operationId := "operationId_example" // string | 

    resp, r, err := sdk.ShardKeyAnalyzerAPI.GetAnalysisOperation(context.Background(), groupId, clusterName, operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardKeyAnalyzerAPI.GetAnalysisOperation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAnalysisOperation`: ShardKeyAnalysisOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardKeyAnalyzerAPI.GetAnalysisOperation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the shard key analysis operation you want to return. | 
**operationId** | **string** | Unique 24-hexadecimal digit string that identifies the shard key analysis operation to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ShardKeyAnalysisOperationResponse**](ShardKeyAnalysisOperationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShardKeyAnalysis

> ShardKeyAnalysisResponse GetShardKeyAnalysis(ctx, groupId, clusterName, analysisId).Execute()

Return One Shard Key Analysis for One Cluster


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
    analysisId := "analysisId_example" // string | 

    resp, r, err := sdk.ShardKeyAnalyzerAPI.GetShardKeyAnalysis(context.Background(), groupId, clusterName, analysisId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardKeyAnalyzerAPI.GetShardKeyAnalysis`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetShardKeyAnalysis`: ShardKeyAnalysisResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardKeyAnalyzerAPI.GetShardKeyAnalysis`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the shard key analysis you want to return. | 
**analysisId** | **string** | Unique 24-hexadecimal digit string that identifies the shard key analysis to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShardKeyAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ShardKeyAnalysisResponse**](ShardKeyAnalysisResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisOperations

> PaginatedShardKeyAnalysisOperationResponse ListAnalysisOperations(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Shard Key Analysis Operations for One Cluster


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

    resp, r, err := sdk.ShardKeyAnalyzerAPI.ListAnalysisOperations(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardKeyAnalyzerAPI.ListAnalysisOperations`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAnalysisOperations`: PaginatedShardKeyAnalysisOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardKeyAnalyzerAPI.ListAnalysisOperations`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the shard key analysis operations you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysisOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether MongoDB Cloud calculates the total number of items for the response. When set to &#x60;false&#x60;, MongoDB Cloud may skip an additional count operation. The response may still include &#x60;totalCount&#x60; when the count is available without additional calculation. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedShardKeyAnalysisOperationResponse**](PaginatedShardKeyAnalysisOperationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShardKeyAnalyses

> PaginatedShardKeyAnalysisResponse ListShardKeyAnalyses(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Shard Key Analyses for One Cluster


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

    resp, r, err := sdk.ShardKeyAnalyzerAPI.ListShardKeyAnalyses(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShardKeyAnalyzerAPI.ListShardKeyAnalyses`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListShardKeyAnalyses`: PaginatedShardKeyAnalysisResponse
    fmt.Fprintf(os.Stdout, "Response from `ShardKeyAnalyzerAPI.ListShardKeyAnalyses`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the shard key analyses you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShardKeyAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether MongoDB Cloud calculates the total number of items for the response. When set to &#x60;false&#x60;, MongoDB Cloud may skip an additional count operation. The response may still include &#x60;totalCount&#x60; when the count is available without additional calculation. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedShardKeyAnalysisResponse**](PaginatedShardKeyAnalysisResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

