# \AtlasSearchApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAtlasSearchIndex**](AtlasSearchApi.md#CreateAtlasSearchIndex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
[**DeleteAtlasSearchIndex**](AtlasSearchApi.md#DeleteAtlasSearchIndex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
[**GetAtlasSearchIndex**](AtlasSearchApi.md#GetAtlasSearchIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
[**ListAtlasSearchIndexes**](AtlasSearchApi.md#ListAtlasSearchIndexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**UpdateAtlasSearchIndex**](AtlasSearchApi.md#UpdateAtlasSearchIndex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index



## CreateAtlasSearchIndex

> ClusterSearchIndex CreateAtlasSearchIndex(ctx, groupId, clusterName, clusterSearchIndex ClusterSearchIndex).Execute()

Create One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.CreateAtlasSearchIndex(context.Background(), groupId, clusterName, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateAtlasSearchIndex``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateAtlasSearchIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateAtlasSearchIndex`: %v\n", resp)
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


 **clusterSearchIndex** | [**ClusterSearchIndex**](ClusterSearchIndex.md) | Creates one Atlas Search index on the specified collection. | 

### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAtlasSearchIndex

> map[string]interface{} DeleteAtlasSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Remove One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.DeleteAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteAtlasSearchIndex``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteAtlasSearchIndex`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.DeleteAtlasSearchIndex`: %v\n", resp)
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

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasSearchIndex

> ClusterSearchIndex GetAtlasSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Return One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetAtlasSearchIndex``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetAtlasSearchIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetAtlasSearchIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAtlasSearchIndexes

> []ClusterSearchIndex ListAtlasSearchIndexes(ctx, groupId, clusterName, collectionName, databaseName).Execute()

Return All Atlas Search Indexes for One Collection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.ListAtlasSearchIndexes(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListAtlasSearchIndexes``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListAtlasSearchIndexes`: []ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListAtlasSearchIndexes`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListAtlasSearchIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlasSearchIndex

> ClusterSearchIndex UpdateAtlasSearchIndex(ctx, groupId, clusterName, indexId, clusterSearchIndex ClusterSearchIndex).Execute()

Update One Atlas Search Index


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    indexId := "indexId_example" // string | 
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.UpdateAtlasSearchIndex(context.Background(), groupId, clusterName, indexId, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAtlasSearchIndex``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateAtlasSearchIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAtlasSearchIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection whose Atlas Search index to update. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clusterSearchIndex** | [**ClusterSearchIndex**](ClusterSearchIndex.md) | Details to update on the Atlas Search index. | 

### Return type

[**ClusterSearchIndex**](ClusterSearchIndex.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

