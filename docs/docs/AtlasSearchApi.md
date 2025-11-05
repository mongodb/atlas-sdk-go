# \AtlasSearchApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterFtsIndex**](AtlasSearchApi.md#CreateClusterFtsIndex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
[**CreateClusterSearchDeployment**](AtlasSearchApi.md#CreateClusterSearchDeployment) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Create Search Nodes
[**CreateClusterSearchIndex**](AtlasSearchApi.md#CreateClusterSearchIndex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Create One Atlas Search Index
[**DeleteClusterFtsIndex**](AtlasSearchApi.md#DeleteClusterFtsIndex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
[**DeleteClusterSearchDeployment**](AtlasSearchApi.md#DeleteClusterSearchDeployment) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Delete Search Nodes
[**DeleteClusterSearchIndex**](AtlasSearchApi.md#DeleteClusterSearchIndex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Remove One Atlas Search Index by ID
[**DeleteIndexByName**](AtlasSearchApi.md#DeleteIndexByName) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Remove One Atlas Search Index by Name
[**GetClusterFtsIndex**](AtlasSearchApi.md#GetClusterFtsIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
[**GetClusterSearchDeployment**](AtlasSearchApi.md#GetClusterSearchDeployment) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Return All Search Nodes
[**GetClusterSearchIndex**](AtlasSearchApi.md#GetClusterSearchIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Return One Atlas Search Index by ID
[**GetIndexByName**](AtlasSearchApi.md#GetIndexByName) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Return One Atlas Search Index by Name
[**ListClusterFtsIndex**](AtlasSearchApi.md#ListClusterFtsIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**ListClusterSearchIndexes**](AtlasSearchApi.md#ListClusterSearchIndexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes | Return All Atlas Search Indexes for One Cluster
[**ListSearchIndex**](AtlasSearchApi.md#ListSearchIndex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**UpdateClusterFtsIndex**](AtlasSearchApi.md#UpdateClusterFtsIndex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index
[**UpdateClusterSearchDeployment**](AtlasSearchApi.md#UpdateClusterSearchDeployment) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment | Update Search Nodes
[**UpdateClusterSearchIndex**](AtlasSearchApi.md#UpdateClusterSearchIndex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId} | Update One Atlas Search Index by ID
[**UpdateIndexByName**](AtlasSearchApi.md#UpdateIndexByName) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName} | Update One Atlas Search Index by Name



## CreateClusterFtsIndex

> ClusterSearchIndex CreateClusterFtsIndex(ctx, groupId, clusterName, clusterSearchIndex ClusterSearchIndex).Execute()

Create One Atlas Search Index


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
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.CreateClusterFtsIndex(context.Background(), groupId, clusterName, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateClusterFtsIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateClusterFtsIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateClusterFtsIndex`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection on which to create an Atlas Search index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterFtsIndexRequest struct via the builder pattern


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


## CreateClusterSearchDeployment

> ApiSearchDeploymentResponse CreateClusterSearchDeployment(ctx, groupId, clusterName, apiSearchDeploymentRequest ApiSearchDeploymentRequest).Execute()

Create Search Nodes


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
    apiSearchDeploymentRequest := *openapiclient.NewApiSearchDeploymentRequest([]openapiclient.ApiSearchDeploymentSpec{*openapiclient.NewApiSearchDeploymentSpec("InstanceSize_example", int(2))}) // ApiSearchDeploymentRequest | 

    resp, r, err := sdk.AtlasSearchApi.CreateClusterSearchDeployment(context.Background(), groupId, clusterName, &apiSearchDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateClusterSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateClusterSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateClusterSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to create Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterSearchDeploymentRequest struct via the builder pattern


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


## CreateClusterSearchIndex

> SearchIndexResponse CreateClusterSearchIndex(ctx, groupId, clusterName, searchIndexCreateRequest SearchIndexCreateRequest).Execute()

Create One Atlas Search Index


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
    searchIndexCreateRequest := *openapiclient.NewSearchIndexCreateRequest("CollectionName_example", "Database_example", "Name_example") // SearchIndexCreateRequest | 

    resp, r, err := sdk.AtlasSearchApi.CreateClusterSearchIndex(context.Background(), groupId, clusterName, &searchIndexCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateClusterSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateClusterSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateClusterSearchIndex`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection on which to create an Atlas Search index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterSearchIndexRequest struct via the builder pattern


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


## DeleteClusterFtsIndex

> DeleteClusterFtsIndex(ctx, groupId, clusterName, indexId).Execute()

Remove One Atlas Search Index


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
    indexId := "indexId_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteClusterFtsIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteClusterFtsIndex`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteClusterFtsIndexRequest struct via the builder pattern


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


## DeleteClusterSearchDeployment

> DeleteClusterSearchDeployment(ctx, groupId, clusterName).Execute()

Delete Search Nodes


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

    r, err := sdk.AtlasSearchApi.DeleteClusterSearchDeployment(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteClusterSearchDeployment`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteClusterSearchDeploymentRequest struct via the builder pattern


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


## DeleteClusterSearchIndex

> DeleteClusterSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Remove One Atlas Search Index by ID


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
    indexId := "indexId_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteClusterSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteClusterSearchIndex`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteClusterSearchIndexRequest struct via the builder pattern


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


## DeleteIndexByName

> DeleteIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName).Execute()

Remove One Atlas Search Index by Name


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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 

    r, err := sdk.AtlasSearchApi.DeleteIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.DeleteIndexByName`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteIndexByNameRequest struct via the builder pattern


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


## GetClusterFtsIndex

> ClusterSearchIndex GetClusterFtsIndex(ctx, groupId, clusterName, indexId).Execute()

Return One Atlas Search Index


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
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetClusterFtsIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetClusterFtsIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterFtsIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetClusterFtsIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiGetClusterFtsIndexRequest struct via the builder pattern


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


## GetClusterSearchDeployment

> ApiSearchDeploymentResponse GetClusterSearchDeployment(ctx, groupId, clusterName).Execute()

Return All Search Nodes


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

    resp, r, err := sdk.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetClusterSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetClusterSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to return the Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSearchDeploymentRequest struct via the builder pattern


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


## GetClusterSearchIndex

> SearchIndexResponse GetClusterSearchIndex(ctx, groupId, clusterName, indexId).Execute()

Return One Atlas Search Index by ID


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
    indexId := "indexId_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetClusterSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetClusterSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetClusterSearchIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiGetClusterSearchIndexRequest struct via the builder pattern


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


## GetIndexByName

> SearchIndexResponse GetIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName).Execute()

Return One Atlas Search Index by Name


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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.GetIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.GetIndexByName`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetIndexByName`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.GetIndexByName`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiGetIndexByNameRequest struct via the builder pattern


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


## ListClusterFtsIndex

> []ClusterSearchIndex ListClusterFtsIndex(ctx, groupId, clusterName, collectionName, databaseName).Execute()

Return All Atlas Search Indexes for One Collection


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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.ListClusterFtsIndex(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListClusterFtsIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListClusterFtsIndex`: []ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListClusterFtsIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiListClusterFtsIndexRequest struct via the builder pattern


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


## ListClusterSearchIndexes

> []SearchIndexResponse ListClusterSearchIndexes(ctx, groupId, clusterName).Execute()

Return All Atlas Search Indexes for One Cluster


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

    resp, r, err := sdk.AtlasSearchApi.ListClusterSearchIndexes(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListClusterSearchIndexes`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListClusterSearchIndexes`: []SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListClusterSearchIndexes`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterSearchIndexesRequest struct via the builder pattern


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


## ListSearchIndex

> []SearchIndexResponse ListSearchIndex(ctx, groupId, clusterName, collectionName, databaseName).Execute()

Return All Atlas Search Indexes for One Collection


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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 

    resp, r, err := sdk.AtlasSearchApi.ListSearchIndex(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ListSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSearchIndex`: []SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ListSearchIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiListSearchIndexRequest struct via the builder pattern


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


## UpdateClusterFtsIndex

> ClusterSearchIndex UpdateClusterFtsIndex(ctx, groupId, clusterName, indexId, clusterSearchIndex ClusterSearchIndex).Execute()

Update One Atlas Search Index


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
    indexId := "indexId_example" // string | 
    clusterSearchIndex := *openapiclient.NewClusterSearchIndex("CollectionName_example", "Database_example", "Name_example") // ClusterSearchIndex | 

    resp, r, err := sdk.AtlasSearchApi.UpdateClusterFtsIndex(context.Background(), groupId, clusterName, indexId, &clusterSearchIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateClusterFtsIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateClusterFtsIndex`: ClusterSearchIndex
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateClusterFtsIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiUpdateClusterFtsIndexRequest struct via the builder pattern


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


## UpdateClusterSearchDeployment

> ApiSearchDeploymentResponse UpdateClusterSearchDeployment(ctx, groupId, clusterName, apiSearchDeploymentRequest ApiSearchDeploymentRequest).Execute()

Update Search Nodes


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
    apiSearchDeploymentRequest := *openapiclient.NewApiSearchDeploymentRequest([]openapiclient.ApiSearchDeploymentSpec{*openapiclient.NewApiSearchDeploymentSpec("InstanceSize_example", int(2))}) // ApiSearchDeploymentRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateClusterSearchDeployment(context.Background(), groupId, clusterName, &apiSearchDeploymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateClusterSearchDeployment`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateClusterSearchDeployment`: ApiSearchDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateClusterSearchDeployment`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Label that identifies the cluster to update the Search Nodes for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterSearchDeploymentRequest struct via the builder pattern


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


## UpdateClusterSearchIndex

> SearchIndexResponse UpdateClusterSearchIndex(ctx, groupId, clusterName, indexId, searchIndexUpdateRequest SearchIndexUpdateRequest).Execute()

Update One Atlas Search Index by ID


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
    indexId := "indexId_example" // string | 
    searchIndexUpdateRequest := *openapiclient.NewSearchIndexUpdateRequest(*openapiclient.NewSearchIndexUpdateRequestDefinition()) // SearchIndexUpdateRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateClusterSearchIndex(context.Background(), groupId, clusterName, indexId, &searchIndexUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateClusterSearchIndex`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateClusterSearchIndex`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateClusterSearchIndex`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiUpdateClusterSearchIndexRequest struct via the builder pattern


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


## UpdateIndexByName

> SearchIndexResponse UpdateIndexByName(ctx, groupId, clusterName, collectionName, databaseName, indexName, searchIndexUpdateRequest SearchIndexUpdateRequest).Execute()

Update One Atlas Search Index by Name


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
    collectionName := "collectionName_example" // string | 
    databaseName := "databaseName_example" // string | 
    indexName := "indexName_example" // string | 
    searchIndexUpdateRequest := *openapiclient.NewSearchIndexUpdateRequest(*openapiclient.NewSearchIndexUpdateRequestDefinition()) // SearchIndexUpdateRequest | 

    resp, r, err := sdk.AtlasSearchApi.UpdateIndexByName(context.Background(), groupId, clusterName, collectionName, databaseName, indexName, &searchIndexUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateIndexByName`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateIndexByName`: SearchIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateIndexByName`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiUpdateIndexByNameRequest struct via the builder pattern


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

