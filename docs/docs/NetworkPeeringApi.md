# \NetworkPeeringApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePeeringConnection**](NetworkPeeringApi.md#CreatePeeringConnection) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One Network Peering Connection
[**CreatePeeringContainer**](NetworkPeeringApi.md#CreatePeeringContainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One Network Peering Container
[**DeletePeeringConnection**](NetworkPeeringApi.md#DeletePeeringConnection) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Network Peering Connection
[**DeletePeeringContainer**](NetworkPeeringApi.md#DeletePeeringContainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
[**DisablePeering**](NetworkPeeringApi.md#DisablePeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering-Only Mode for One Project
[**GetPeeringConnection**](NetworkPeeringApi.md#GetPeeringConnection) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
[**GetPeeringContainer**](NetworkPeeringApi.md#GetPeeringContainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
[**ListPeeringConnections**](NetworkPeeringApi.md#ListPeeringConnections) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project
[**ListPeeringContainerByCloudProvider**](NetworkPeeringApi.md#ListPeeringContainerByCloudProvider) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
[**ListPeeringContainers**](NetworkPeeringApi.md#ListPeeringContainers) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
[**UpdatePeeringConnection**](NetworkPeeringApi.md#UpdatePeeringConnection) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One Network Peering Connection
[**UpdatePeeringContainer**](NetworkPeeringApi.md#UpdatePeeringContainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container
[**VerifyConnectViaPeeringOnlyModeForOneProject**](NetworkPeeringApi.md#VerifyConnectViaPeeringOnlyModeForOneProject) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering-Only Mode for One Project



## CreatePeeringConnection

> BaseNetworkPeeringConnectionSettings CreatePeeringConnection(ctx, groupId, baseNetworkPeeringConnectionSettings BaseNetworkPeeringConnectionSettings).Execute()

Create One Network Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    baseNetworkPeeringConnectionSettings := *openapiclient.NewBaseNetworkPeeringConnectionSettings("32b6e34b3d91647abb20e7b8") // BaseNetworkPeeringConnectionSettings | 

    resp, r, err := sdk.NetworkPeeringApi.CreatePeeringConnection(context.Background(), groupId, &baseNetworkPeeringConnectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.CreatePeeringConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePeeringConnection`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.CreatePeeringConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseNetworkPeeringConnectionSettings** | [**BaseNetworkPeeringConnectionSettings**](BaseNetworkPeeringConnectionSettings.md) | Create one network peering connection. | 

### Return type

[**BaseNetworkPeeringConnectionSettings**](BaseNetworkPeeringConnectionSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePeeringContainer

> CloudProviderContainer CreatePeeringContainer(ctx, groupId, cloudProviderContainer CloudProviderContainer).Execute()

Create One Network Peering Container


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    cloudProviderContainer := *openapiclient.NewCloudProviderContainer() // CloudProviderContainer | 

    resp, r, err := sdk.NetworkPeeringApi.CreatePeeringContainer(context.Background(), groupId, &cloudProviderContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.CreatePeeringContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePeeringContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.CreatePeeringContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProviderContainer** | [**CloudProviderContainer**](CloudProviderContainer.md) | Creates one new network peering container in the specified project. | 

### Return type

[**CloudProviderContainer**](CloudProviderContainer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePeeringConnection

> any DeletePeeringConnection(ctx, groupId, peerId).Execute()

Remove One Network Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    peerId := "peerId_example" // string | 

    resp, r, err := sdk.NetworkPeeringApi.DeletePeeringConnection(context.Background(), groupId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DeletePeeringConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeletePeeringConnection`: any
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.DeletePeeringConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePeeringContainer

> DeletePeeringContainer(ctx, groupId, containerId).Execute()

Remove One Network Peering Container


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.NetworkPeeringApi.DeletePeeringContainer(context.Background(), groupId, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DeletePeeringContainer`: %v (%v)\n", err, r)
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
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePeeringContainerRequest struct via the builder pattern


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


## DisablePeering

> PrivateIPMode DisablePeering(ctx, groupId, privateIPMode PrivateIPMode).Execute()

Disable Connect via Peering-Only Mode for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    privateIPMode := *openapiclient.NewPrivateIPMode(false) // PrivateIPMode | 

    resp, r, err := sdk.NetworkPeeringApi.DisablePeering(context.Background(), groupId, &privateIPMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DisablePeering`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DisablePeering`: PrivateIPMode
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.DisablePeering`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateIPMode** | [**PrivateIPMode**](PrivateIPMode.md) | Disables Connect via Peering Only mode for the specified project. | 

### Return type

[**PrivateIPMode**](PrivateIPMode.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeeringConnection

> BaseNetworkPeeringConnectionSettings GetPeeringConnection(ctx, groupId, peerId).Execute()

Return One Network Peering Connection in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    peerId := "peerId_example" // string | 

    resp, r, err := sdk.NetworkPeeringApi.GetPeeringConnection(context.Background(), groupId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.GetPeeringConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPeeringConnection`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.GetPeeringConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseNetworkPeeringConnectionSettings**](BaseNetworkPeeringConnectionSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeeringContainer

> CloudProviderContainer GetPeeringContainer(ctx, groupId, containerId).Execute()

Return One Network Peering Container


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.NetworkPeeringApi.GetPeeringContainer(context.Background(), groupId, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.GetPeeringContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPeeringContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.GetPeeringContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudProviderContainer**](CloudProviderContainer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPeeringConnections

> PaginatedContainerPeer ListPeeringConnections(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ProviderName(providerName).Execute()

Return All Network Peering Connections in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    providerName := "providerName_example" // string |  (optional) (default to "AWS")

    resp, r, err := sdk.NetworkPeeringApi.ListPeeringConnections(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ProviderName(providerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListPeeringConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPeeringConnections`: PaginatedContainerPeer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListPeeringConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPeeringConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **providerName** | **string** | Cloud service provider to use for this VPC peering connection. | [default to &quot;AWS&quot;]

### Return type

[**PaginatedContainerPeer**](PaginatedContainerPeer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPeeringContainerByCloudProvider

> PaginatedCloudProviderContainer ListPeeringContainerByCloudProvider(ctx, groupId).ProviderName(providerName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Network Peering Containers in One Project for One Cloud Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    providerName := "providerName_example" // string |  (default to "AWS")
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.NetworkPeeringApi.ListPeeringContainerByCloudProvider(context.Background(), groupId).ProviderName(providerName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListPeeringContainerByCloudProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPeeringContainerByCloudProvider`: PaginatedCloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListPeeringContainerByCloudProvider`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPeeringContainerByCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerName** | **string** | Cloud service provider that serves the desired network peering containers. | [default to &quot;AWS&quot;]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedCloudProviderContainer**](PaginatedCloudProviderContainer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPeeringContainers

> PaginatedCloudProviderContainer ListPeeringContainers(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Network Peering Containers in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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

    resp, r, err := sdk.NetworkPeeringApi.ListPeeringContainers(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListPeeringContainers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPeeringContainers`: PaginatedCloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListPeeringContainers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPeeringContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedCloudProviderContainer**](PaginatedCloudProviderContainer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePeeringConnection

> BaseNetworkPeeringConnectionSettings UpdatePeeringConnection(ctx, groupId, peerId, baseNetworkPeeringConnectionSettings BaseNetworkPeeringConnectionSettings).Execute()

Update One Network Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    peerId := "peerId_example" // string | 
    baseNetworkPeeringConnectionSettings := *openapiclient.NewBaseNetworkPeeringConnectionSettings("32b6e34b3d91647abb20e7b8") // BaseNetworkPeeringConnectionSettings | 

    resp, r, err := sdk.NetworkPeeringApi.UpdatePeeringConnection(context.Background(), groupId, peerId, &baseNetworkPeeringConnectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.UpdatePeeringConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdatePeeringConnection`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.UpdatePeeringConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **baseNetworkPeeringConnectionSettings** | [**BaseNetworkPeeringConnectionSettings**](BaseNetworkPeeringConnectionSettings.md) | Modify one network peering connection. | 

### Return type

[**BaseNetworkPeeringConnectionSettings**](BaseNetworkPeeringConnectionSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePeeringContainer

> CloudProviderContainer UpdatePeeringContainer(ctx, groupId, containerId, cloudProviderContainer CloudProviderContainer).Execute()

Update One Network Peering Container


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 
    cloudProviderContainer := *openapiclient.NewCloudProviderContainer() // CloudProviderContainer | 

    resp, r, err := sdk.NetworkPeeringApi.UpdatePeeringContainer(context.Background(), groupId, containerId, &cloudProviderContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.UpdatePeeringContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdatePeeringContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.UpdatePeeringContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudProviderContainer** | [**CloudProviderContainer**](CloudProviderContainer.md) | Updates the network details and labels of one specified network peering container in the specified project. | 

### Return type

[**CloudProviderContainer**](CloudProviderContainer.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyConnectViaPeeringOnlyModeForOneProject

> PrivateIPMode VerifyConnectViaPeeringOnlyModeForOneProject(ctx, groupId).Execute()

Verify Connect via Peering-Only Mode for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312005/admin"
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

    resp, r, err := sdk.NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `VerifyConnectViaPeeringOnlyModeForOneProject`: PrivateIPMode
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnectViaPeeringOnlyModeForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateIPMode**](PrivateIPMode.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

