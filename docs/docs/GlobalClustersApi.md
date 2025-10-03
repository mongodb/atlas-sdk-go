# \GlobalClustersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomZoneMapping**](GlobalClustersApi.md#CreateCustomZoneMapping) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add One Custom Zone Mapping to One Global Cluster
[**CreateManagedNamespace**](GlobalClustersApi.md#CreateManagedNamespace) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Cluster
[**DeleteCustomZoneMapping**](GlobalClustersApi.md#DeleteCustomZoneMapping) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Cluster
[**DeleteManagedNamespaces**](GlobalClustersApi.md#DeleteManagedNamespaces) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Cluster
[**GetClusterGlobalWrites**](GlobalClustersApi.md#GetClusterGlobalWrites) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites | Return One Managed Namespace in One Global Cluster



## CreateCustomZoneMapping

> GeoSharding20240805 CreateCustomZoneMapping(ctx, groupId, clusterName, customZoneMappings CustomZoneMappings).Execute()

Add One Custom Zone Mapping to One Global Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    customZoneMappings := *openapiclient.NewCustomZoneMappings() // CustomZoneMappings | 

    resp, r, err := sdk.GlobalClustersApi.CreateCustomZoneMapping(context.Background(), groupId, clusterName, &customZoneMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.CreateCustomZoneMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCustomZoneMapping`: GeoSharding20240805
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.CreateCustomZoneMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomZoneMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customZoneMappings** | [**CustomZoneMappings**](CustomZoneMappings.md) | Custom zone mapping to add to the specified global cluster. | 

### Return type

[**GeoSharding20240805**](GeoSharding20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagedNamespace

> GeoSharding20240805 CreateManagedNamespace(ctx, groupId, clusterName, managedNamespaces ManagedNamespaces).Execute()

Create One Managed Namespace in One Global Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    managedNamespaces := *openapiclient.NewManagedNamespaces("Collection_example", "CustomShardKey_example", "Db_example") // ManagedNamespaces | 

    resp, r, err := sdk.GlobalClustersApi.CreateManagedNamespace(context.Background(), groupId, clusterName, &managedNamespaces).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.CreateManagedNamespace`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateManagedNamespace`: GeoSharding20240805
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.CreateManagedNamespace`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managedNamespaces** | [**ManagedNamespaces**](ManagedNamespaces.md) | Managed namespace to create within the specified global cluster. | 

### Return type

[**GeoSharding20240805**](GeoSharding20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomZoneMapping

> GeoSharding20240805 DeleteCustomZoneMapping(ctx, groupId, clusterName).Execute()

Remove All Custom Zone Mappings from One Global Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.GlobalClustersApi.DeleteCustomZoneMapping(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.DeleteCustomZoneMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteCustomZoneMapping`: GeoSharding20240805
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.DeleteCustomZoneMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomZoneMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GeoSharding20240805**](GeoSharding20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedNamespaces

> GeoSharding20240805 DeleteManagedNamespaces(ctx, clusterName, groupId).Db(db).Collection(collection).Execute()

Remove One Managed Namespace from One Global Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    clusterName := "clusterName_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    db := "db_example" // string |  (optional)
    collection := "collection_example" // string |  (optional)

    resp, r, err := sdk.GlobalClustersApi.DeleteManagedNamespaces(context.Background(), clusterName, groupId).Db(db).Collection(collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.DeleteManagedNamespaces`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteManagedNamespaces`: GeoSharding20240805
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.DeleteManagedNamespaces`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies this cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **db** | **string** | Human-readable label that identifies the database that contains the collection. | 
 **collection** | **string** | Human-readable label that identifies the collection associated with the managed namespace. | 

### Return type

[**GeoSharding20240805**](GeoSharding20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterGlobalWrites

> GeoSharding20240805 GetClusterGlobalWrites(ctx, groupId, clusterName).Execute()

Return One Managed Namespace in One Global Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.GlobalClustersApi.GetClusterGlobalWrites(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.GetClusterGlobalWrites`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetClusterGlobalWrites`: GeoSharding20240805
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.GetClusterGlobalWrites`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterGlobalWritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GeoSharding20240805**](GeoSharding20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

