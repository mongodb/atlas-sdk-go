# \FlexClustersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlexCluster**](FlexClustersApi.md#CreateFlexCluster) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters | Create One Flex Cluster in One Project
[**DeleteFlexCluster**](FlexClustersApi.md#DeleteFlexCluster) | **Delete** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Remove One Flex Cluster from One Project
[**GetFlexCluster**](FlexClustersApi.md#GetFlexCluster) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Return One Flex Cluster from One Project
[**ListFlexClusters**](FlexClustersApi.md#ListFlexClusters) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters | Return All Flex Clusters from One Project
[**TenantUpgrade**](FlexClustersApi.md#TenantUpgrade) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters:tenantUpgrade | Upgrade One Flex Cluster
[**UpdateFlexCluster**](FlexClustersApi.md#UpdateFlexCluster) | **Patch** /api/atlas/v2/groups/{groupId}/flexClusters/{name} | Update One Flex Cluster in One Project



## CreateFlexCluster

> FlexClusterDescription20241113 CreateFlexCluster(ctx, groupId, flexClusterDescriptionCreate20241113 FlexClusterDescriptionCreate20241113).Execute()

Create One Flex Cluster in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    flexClusterDescriptionCreate20241113 := *openapiclient.NewFlexClusterDescriptionCreate20241113("Name_example", *openapiclient.NewFlexProviderSettingsCreate20241113("BackingProviderName_example", "RegionName_example")) // FlexClusterDescriptionCreate20241113 | 

    resp, r, err := sdk.FlexClustersApi.CreateFlexCluster(context.Background(), groupId, &flexClusterDescriptionCreate20241113).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.CreateFlexCluster`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateFlexCluster`: FlexClusterDescription20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexClustersApi.CreateFlexCluster`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlexClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flexClusterDescriptionCreate20241113** | [**FlexClusterDescriptionCreate20241113**](FlexClusterDescriptionCreate20241113.md) | Create One Flex Cluster in One Project. | 

### Return type

[**FlexClusterDescription20241113**](FlexClusterDescription20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-11-13+json
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlexCluster

> DeleteFlexCluster(ctx, groupId, name).Execute()

Remove One Flex Cluster from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    name := "name_example" // string | 

    r, err := sdk.FlexClustersApi.DeleteFlexCluster(context.Background(), groupId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.DeleteFlexCluster`: %v (%v)\n", err, r)
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
**name** | **string** | Human-readable label that identifies the flex cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlexClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexCluster

> FlexClusterDescription20241113 GetFlexCluster(ctx, groupId, name).Execute()

Return One Flex Cluster from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    name := "name_example" // string | 

    resp, r, err := sdk.FlexClustersApi.GetFlexCluster(context.Background(), groupId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.GetFlexCluster`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetFlexCluster`: FlexClusterDescription20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexClustersApi.GetFlexCluster`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlexClusterDescription20241113**](FlexClusterDescription20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexClusters

> PaginatedFlexClusters20241113 ListFlexClusters(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Flex Clusters from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.FlexClustersApi.ListFlexClusters(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.ListFlexClusters`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListFlexClusters`: PaginatedFlexClusters20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexClustersApi.ListFlexClusters`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFlexClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedFlexClusters20241113**](PaginatedFlexClusters20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantUpgrade

> FlexClusterDescription20241113 TenantUpgrade(ctx, groupId, atlasTenantClusterUpgradeRequest20240805 AtlasTenantClusterUpgradeRequest20240805).Execute()

Upgrade One Flex Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    atlasTenantClusterUpgradeRequest20240805 := *openapiclient.NewAtlasTenantClusterUpgradeRequest20240805("Name_example") // AtlasTenantClusterUpgradeRequest20240805 | 

    resp, r, err := sdk.FlexClustersApi.TenantUpgrade(context.Background(), groupId, &atlasTenantClusterUpgradeRequest20240805).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.TenantUpgrade`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `TenantUpgrade`: FlexClusterDescription20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexClustersApi.TenantUpgrade`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **atlasTenantClusterUpgradeRequest20240805** | [**AtlasTenantClusterUpgradeRequest20240805**](AtlasTenantClusterUpgradeRequest20240805.md) | Details of the flex cluster upgrade in the specified project. | 

### Return type

[**FlexClusterDescription20241113**](FlexClusterDescription20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-11-13+json
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexCluster

> FlexClusterDescription20241113 UpdateFlexCluster(ctx, groupId, name, flexClusterDescriptionUpdate20241113 FlexClusterDescriptionUpdate20241113).Execute()

Update One Flex Cluster in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312007/admin"
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
    name := "name_example" // string | 
    flexClusterDescriptionUpdate20241113 := *openapiclient.NewFlexClusterDescriptionUpdate20241113() // FlexClusterDescriptionUpdate20241113 | 

    resp, r, err := sdk.FlexClustersApi.UpdateFlexCluster(context.Background(), groupId, name, &flexClusterDescriptionUpdate20241113).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexClustersApi.UpdateFlexCluster`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateFlexCluster`: FlexClusterDescription20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexClustersApi.UpdateFlexCluster`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlexClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flexClusterDescriptionUpdate20241113** | [**FlexClusterDescriptionUpdate20241113**](FlexClusterDescriptionUpdate20241113.md) | Update One Flex Cluster in One Project. | 

### Return type

[**FlexClusterDescription20241113**](FlexClusterDescription20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-11-13+json
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

