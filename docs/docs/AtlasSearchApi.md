# \AtlasSearchApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAtlasSearchDeployment**](AtlasSearchApi.md#CreateAtlasSearchDeployment) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Create Search Nodes
[**CreateAtlasSearchIndex**](AtlasSearchApi.md#CreateAtlasSearchIndex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Create One Atlas Search Index
[**CreateAtlasSearchIndexDeprecated**](AtlasSearchApi.md#CreateAtlasSearchIndexDeprecated) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
[**DeleteAtlasSearchDeployment**](AtlasSearchApi.md#DeleteAtlasSearchDeployment) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Delete Search Nodes
[**DeleteAtlasSearchIndex**](AtlasSearchApi.md#DeleteAtlasSearchIndex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Remove One Atlas Search Index by ID
[**DeleteAtlasSearchIndexByName**](AtlasSearchApi.md#DeleteAtlasSearchIndexByName) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Remove One Atlas Search Index by Name
[**DeleteAtlasSearchIndexDeprecated**](AtlasSearchApi.md#DeleteAtlasSearchIndexDeprecated) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
[**GetAtlasSearchDeployment**](AtlasSearchApi.md#GetAtlasSearchDeployment) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Return All Search Nodes
[**GetAtlasSearchIndex**](AtlasSearchApi.md#GetAtlasSearchIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Return One Atlas Search Index by ID
[**GetAtlasSearchIndexByName**](AtlasSearchApi.md#GetAtlasSearchIndexByName) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Return One Atlas Search Index by Name
[**GetAtlasSearchIndexDeprecated**](AtlasSearchApi.md#GetAtlasSearchIndexDeprecated) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
[**ListAtlasSearchIndexes**](AtlasSearchApi.md#ListAtlasSearchIndexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**ListAtlasSearchIndexesCluster**](AtlasSearchApi.md#ListAtlasSearchIndexesCluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Return All Atlas Search Indexes for One Cluster
[**ListAtlasSearchIndexesDeprecated**](AtlasSearchApi.md#ListAtlasSearchIndexesDeprecated) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**UpdateAtlasSearchDeployment**](AtlasSearchApi.md#UpdateAtlasSearchDeployment) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Update Search Nodes
[**UpdateAtlasSearchIndex**](AtlasSearchApi.md#UpdateAtlasSearchIndex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Update One Atlas Search Index by ID
[**UpdateAtlasSearchIndexByName**](AtlasSearchApi.md#UpdateAtlasSearchIndexByName) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Update One Atlas Search Index by Name
[**UpdateAtlasSearchIndexDeprecated**](AtlasSearchApi.md#UpdateAtlasSearchIndexDeprecated) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index



## CreateAtlasSearchDeployment

> ApiSearchDeploymentResponse CreateAtlasSearchDeployment(ctx, groupId, clusterName, apiSearchDeploymentRequest ApiSearchDeploymentRequest).Execute()

Create Search Nodes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    apiSearchDeploymentRequest := *openapiclient.NewApiSearchDeploymentRequest([]openapiclient.ApiSearchDeploymentSpec{*openapiclient.NewApiSearchDeploymentSpec("InstanceSize_example", int(2))}) // ApiSearchDeploymentRequest | 

    resp, r, err := sdk.AtlasSearchApi.CreateAtlasSearchDeployment(context.Background(), groupId, clusterName, &apiSearchDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateAtlasSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAtlasSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateAtlasSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to create Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAtlasSearchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSearchDeploymentRequest** | [**ApiSearchDeploymentRequest**](ApiSearchDeploymentRequest.md) | Creates Search Nodes for the specified cluster. | 

### Return type

[**ApiSearchDeploymentResponse**](ApiSearchDeploymentResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAtlasSearchIndex

> SearchIndexResponse CreateAtlasSearchIndex(ctx, groupId, clusterName, searchIndexCreateRequest SearchIndexCreateRequest).Execute()

Create One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    searchIndexCreateRequest := *openapiclient.NewSearchIndexCreateRequest("CollectionName_example", "Database_example", "Name_example") // SearchIndexCreateRequest | 

    resp, r, err := sdk.AtlasSearchApi.CreateAtlasSearchIndex(context.Background(), groupId, clusterName, &searchIndexCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateAtlasSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAtlasSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateAtlasSearchIndex`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection on which to create an Atlas Search index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchIndexCreateRequest** | [**SearchIndexCreateRequest**](SearchIndexCreateRequest.md) | Creates one Atlas Search index on the specified collection. | 

### Return type

[**SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAtlasSearchIndexDeprecated

> ClusterSearchIndex CreateAtlasSearchIndexDeprecated(ctx, groupId, clusterName, clusterSearchIndex ClusterSearchIndex).Execute()

Create One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.CreateAtlasSearchIndexDeprecated(context.Background(), groupId, clusterName, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateAtlasSearchIndexDeprecated`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAtlasSearchIndexDeprecated`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateAtlasSearchIndexDeprecated`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection on which to create an Atlas Search index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAtlasSearchIndexDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterSearchIndex** | [**ClusterSearchIndex**](ClusterSearchIndex.md) | Creates one Atlas Search index on the specified collection. | 

### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasSearchDeployment

> DeleteAtlasSearchDeployment(ctx, groupId, clusterName).Execute()

Delete Search Nodes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    r, err := sdk.AtlasSearchApi.DeleteAtlasSearchDeployment(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteAtlasSearchDeployment`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Label that identifies the cluster to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAtlasSearchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasSearchIndex

> DeleteAtlasSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Remove One Atlas Search Index by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteAtlasSearchIndex`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Name of the cluster that contains the database and collection with one or more Application Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasSearchIndexByName

> DeleteAtlasSearchIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName).Execute()

Remove One Atlas Search Index by Name


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteAtlasSearchIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteAtlasSearchIndexByName`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Name of the cluster that contains the database and collection with one or more Application Search indexes. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Label that identifies the database that contains the collection with one or more Atlas Search indexes. | 
**indexName** | **string** | Name of the Atlas Search index to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAtlasSearchIndexByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasSearchIndexDeprecated

> DeleteAtlasSearchIndexDeprecated(ctx, groupId, clusterName, indexId).Execute()

Remove One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteAtlasSearchIndexDeprecated(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteAtlasSearchIndexDeprecated`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Name of the cluster that contains the database and collection with one or more Application Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAtlasSearchIndexDeprecatedRequest struct via the builder pattern


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


## GetAtlasSearchDeployment

> ApiSearchDeploymentResponse GetAtlasSearchDeployment(ctx, groupId, clusterName).Execute()

Return All Search Nodes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetAtlasSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetAtlasSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to return the Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasSearchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSearchDeploymentResponse**](ApiSearchDeploymentResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasSearchIndex

> SearchIndexResponse GetAtlasSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Return One Atlas Search Index by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetAtlasSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetAtlasSearchIndex`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasSearchIndexByName

> SearchIndexResponse GetAtlasSearchIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName).Execute()

Return One Atlas Search Index by Name


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetAtlasSearchIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetAtlasSearchIndexByName`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasSearchIndexByName`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetAtlasSearchIndexByName`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Label that identifies the database that contains the collection with one or more Atlas Search indexes. | 
**indexName** | **string** | Name of the Atlas Search index to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasSearchIndexByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasSearchIndexDeprecated

> ClusterSearchIndex GetAtlasSearchIndexDeprecated(ctx, groupId, clusterName, indexId).Execute()

Return One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetAtlasSearchIndexDeprecated(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetAtlasSearchIndexDeprecated`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAtlasSearchIndexDeprecated`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetAtlasSearchIndexDeprecated`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasSearchIndexDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtlasSearchIndexes

> []SearchIndexResponse ListAtlasSearchIndexes(ctx, groupId, clusterName, collectionName, databaseName).Execute()

Return All Atlas Search Indexes for One Collection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.ListAtlasSearchIndexes(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListAtlasSearchIndexes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAtlasSearchIndexes`: []SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListAtlasSearchIndexes`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Label that identifies the database that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtlasSearchIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtlasSearchIndexesCluster

> []SearchIndexResponse ListAtlasSearchIndexesCluster(ctx, groupId, clusterName).Execute()

Return All Atlas Search Indexes for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.AtlasSearchApi.ListAtlasSearchIndexesCluster(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListAtlasSearchIndexesCluster`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAtlasSearchIndexesCluster`: []SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListAtlasSearchIndexesCluster`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtlasSearchIndexesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtlasSearchIndexesDeprecated

> []ClusterSearchIndex ListAtlasSearchIndexesDeprecated(ctx, groupId, clusterName, collectionName, databaseName).Execute()

Return All Atlas Search Indexes for One Collection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.ListAtlasSearchIndexesDeprecated(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListAtlasSearchIndexesDeprecated`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAtlasSearchIndexesDeprecated`: []ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListAtlasSearchIndexesDeprecated`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAtlasSearchIndexesDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasSearchDeployment

> ApiSearchDeploymentResponse UpdateAtlasSearchDeployment(ctx, groupId, clusterName, apiSearchDeploymentRequest ApiSearchDeploymentRequest).Execute()

Update Search Nodes


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    apiSearchDeploymentRequest := *openapiclient.NewApiSearchDeploymentRequest([]openapiclient.ApiSearchDeploymentSpec{*openapiclient.NewApiSearchDeploymentSpec("InstanceSize_example", int(2))}) // ApiSearchDeploymentRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateAtlasSearchDeployment(context.Background(), groupId, clusterName, &apiSearchDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAtlasSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAtlasSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAtlasSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to update the Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasSearchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSearchDeploymentRequest** | [**ApiSearchDeploymentRequest**](ApiSearchDeploymentRequest.md) | Updates the Search Nodes for the specified cluster. | 

### Return type

[**ApiSearchDeploymentResponse**](ApiSearchDeploymentResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasSearchIndex

> SearchIndexResponse UpdateAtlasSearchIndex(ctx, groupId, clusterName, indexId, searchIndexUpdateRequest SearchIndexUpdateRequest).Execute()

Update One Atlas Search Index by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 
    searchIndexUpdateRequest := *openapiclient.NewSearchIndexUpdateRequest(*openapiclient.NewSearchIndexUpdateRequestDefinition()) // SearchIndexUpdateRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateAtlasSearchIndex(context.Background(), groupId, clusterName, indexId, &searchIndexUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAtlasSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAtlasSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAtlasSearchIndex`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection whose Atlas Search index you want to update. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **searchIndexUpdateRequest** | [**SearchIndexUpdateRequest**](SearchIndexUpdateRequest.md) | Details to update on the Atlas Search index. | 

### Return type

[**SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasSearchIndexByName

> SearchIndexResponse UpdateAtlasSearchIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName, searchIndexUpdateRequest SearchIndexUpdateRequest).Execute()

Update One Atlas Search Index by Name


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 
    searchIndexUpdateRequest := *openapiclient.NewSearchIndexUpdateRequest(*openapiclient.NewSearchIndexUpdateRequestDefinition()) // SearchIndexUpdateRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateAtlasSearchIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName, &searchIndexUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAtlasSearchIndexByName`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAtlasSearchIndexByName`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAtlasSearchIndexByName`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection whose Atlas Search index you want to update. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Label that identifies the database that contains the collection with one or more Atlas Search indexes. | 
**indexName** | **string** | Name of the Atlas Search index to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasSearchIndexByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **searchIndexUpdateRequest** | [**SearchIndexUpdateRequest**](SearchIndexUpdateRequest.md) | Details to update the Atlas Search index with. | 

### Return type

[**SearchIndexResponse**](SearchIndexResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasSearchIndexDeprecated

> ClusterSearchIndex UpdateAtlasSearchIndexDeprecated(ctx, groupId, clusterName, indexId, clusterSearchIndex ClusterSearchIndex).Execute()

Update One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    indexId := "indexId_example" // string | 
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.UpdateAtlasSearchIndexDeprecated(context.Background(), groupId, clusterName, indexId, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAtlasSearchIndexDeprecated`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAtlasSearchIndexDeprecated`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAtlasSearchIndexDeprecated`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection whose Atlas Search index to update. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasSearchIndexDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clusterSearchIndex** | [**ClusterSearchIndex**](ClusterSearchIndex.md) | Details to update on the Atlas Search index. | 

### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

