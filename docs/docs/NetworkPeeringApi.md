# \NetworkPeeringApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupContainer**](NetworkPeeringApi.md#CreateGroupContainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One Network Peering Container
[**CreateGroupPeer**](NetworkPeeringApi.md#CreateGroupPeer) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One Network Peering Connection
[**DeleteGroupContainer**](NetworkPeeringApi.md#DeleteGroupContainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
[**DeleteGroupPeer**](NetworkPeeringApi.md#DeleteGroupPeer) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Network Peering Connection
[**DisablePeering**](NetworkPeeringApi.md#DisablePeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering-Only Mode for One Project
[**GetGroupContainer**](NetworkPeeringApi.md#GetGroupContainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
[**GetGroupPeer**](NetworkPeeringApi.md#GetGroupPeer) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
[**ListGroupContainerAll**](NetworkPeeringApi.md#ListGroupContainerAll) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
[**ListGroupContainers**](NetworkPeeringApi.md#ListGroupContainers) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
[**ListGroupPeers**](NetworkPeeringApi.md#ListGroupPeers) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project
[**UpdateGroupContainer**](NetworkPeeringApi.md#UpdateGroupContainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container
[**UpdateGroupPeer**](NetworkPeeringApi.md#UpdateGroupPeer) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One Network Peering Connection
[**VerifyPrivateIpMode**](NetworkPeeringApi.md#VerifyPrivateIpMode) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering-Only Mode for One Project



## CreateGroupContainer

> CloudProviderContainer CreateGroupContainer(ctx, groupId, cloudProviderContainer CloudProviderContainer).Execute()

Create One Network Peering Container


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
    cloudProviderContainer := *openapiclient.NewCloudProviderContainer() // CloudProviderContainer | 

    resp, r, err := sdk.NetworkPeeringApi.CreateGroupContainer(context.Background(), groupId, &cloudProviderContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.CreateGroupContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.CreateGroupContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupContainerRequest struct via the builder pattern


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


## CreateGroupPeer

> BaseNetworkPeeringConnectionSettings CreateGroupPeer(ctx, groupId, baseNetworkPeeringConnectionSettings BaseNetworkPeeringConnectionSettings).Execute()

Create One Network Peering Connection


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
    baseNetworkPeeringConnectionSettings := *openapiclient.NewBaseNetworkPeeringConnectionSettings("32b6e34b3d91647abb20e7b8") // BaseNetworkPeeringConnectionSettings | 

    resp, r, err := sdk.NetworkPeeringApi.CreateGroupPeer(context.Background(), groupId, &baseNetworkPeeringConnectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.CreateGroupPeer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupPeer`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.CreateGroupPeer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupPeerRequest struct via the builder pattern


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


## DeleteGroupContainer

> DeleteGroupContainer(ctx, groupId, containerId).Execute()

Remove One Network Peering Container


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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.NetworkPeeringApi.DeleteGroupContainer(context.Background(), groupId, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DeleteGroupContainer`: %v (%v)\n", err, r)
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

Other parameters are passed through a pointer to a apiDeleteGroupContainerRequest struct via the builder pattern


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


## DeleteGroupPeer

> any DeleteGroupPeer(ctx, groupId, peerId).Execute()

Remove One Network Peering Connection


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
    peerId := "peerId_example" // string | 

    resp, r, err := sdk.NetworkPeeringApi.DeleteGroupPeer(context.Background(), groupId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DeleteGroupPeer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteGroupPeer`: any
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.DeleteGroupPeer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupPeerRequest struct via the builder pattern


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


## GetGroupContainer

> CloudProviderContainer GetGroupContainer(ctx, groupId, containerId).Execute()

Return One Network Peering Container


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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.NetworkPeeringApi.GetGroupContainer(context.Background(), groupId, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.GetGroupContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.GetGroupContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupContainerRequest struct via the builder pattern


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


## GetGroupPeer

> BaseNetworkPeeringConnectionSettings GetGroupPeer(ctx, groupId, peerId).Execute()

Return One Network Peering Connection in One Project


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
    peerId := "peerId_example" // string | 

    resp, r, err := sdk.NetworkPeeringApi.GetGroupPeer(context.Background(), groupId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.GetGroupPeer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupPeer`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.GetGroupPeer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPeerRequest struct via the builder pattern


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


## ListGroupContainerAll

> PaginatedCloudProviderContainer ListGroupContainerAll(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Network Peering Containers in One Project


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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.NetworkPeeringApi.ListGroupContainerAll(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListGroupContainerAll`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupContainerAll`: PaginatedCloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListGroupContainerAll`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupContainerAllRequest struct via the builder pattern


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


## ListGroupContainers

> PaginatedCloudProviderContainer ListGroupContainers(ctx, groupId).ProviderName(providerName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Network Peering Containers in One Project for One Cloud Provider


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
    providerName := "providerName_example" // string |  (default to "AWS")
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.NetworkPeeringApi.ListGroupContainers(context.Background(), groupId).ProviderName(providerName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListGroupContainers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupContainers`: PaginatedCloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListGroupContainers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupContainersRequest struct via the builder pattern


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


## ListGroupPeers

> PaginatedContainerPeer ListGroupPeers(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ProviderName(providerName).Execute()

Return All Network Peering Connections in One Project


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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    providerName := "providerName_example" // string |  (optional) (default to "AWS")

    resp, r, err := sdk.NetworkPeeringApi.ListGroupPeers(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ProviderName(providerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.ListGroupPeers`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupPeers`: PaginatedContainerPeer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.ListGroupPeers`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupPeersRequest struct via the builder pattern


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


## UpdateGroupContainer

> CloudProviderContainer UpdateGroupContainer(ctx, groupId, containerId, cloudProviderContainer CloudProviderContainer).Execute()

Update One Network Peering Container


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
    containerId := "32b6e34b3d91647abb20e7b8" // string | 
    cloudProviderContainer := *openapiclient.NewCloudProviderContainer() // CloudProviderContainer | 

    resp, r, err := sdk.NetworkPeeringApi.UpdateGroupContainer(context.Background(), groupId, containerId, &cloudProviderContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.UpdateGroupContainer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupContainer`: CloudProviderContainer
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.UpdateGroupContainer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupContainerRequest struct via the builder pattern


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


## UpdateGroupPeer

> BaseNetworkPeeringConnectionSettings UpdateGroupPeer(ctx, groupId, peerId, baseNetworkPeeringConnectionSettings BaseNetworkPeeringConnectionSettings).Execute()

Update One Network Peering Connection


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
    peerId := "peerId_example" // string | 
    baseNetworkPeeringConnectionSettings := *openapiclient.NewBaseNetworkPeeringConnectionSettings("32b6e34b3d91647abb20e7b8") // BaseNetworkPeeringConnectionSettings | 

    resp, r, err := sdk.NetworkPeeringApi.UpdateGroupPeer(context.Background(), groupId, peerId, &baseNetworkPeeringConnectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.UpdateGroupPeer`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupPeer`: BaseNetworkPeeringConnectionSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.UpdateGroupPeer`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupPeerRequest struct via the builder pattern


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


## VerifyPrivateIpMode

> PrivateIPMode VerifyPrivateIpMode(ctx, groupId).Execute()

Verify Connect via Peering-Only Mode for One Project


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

    resp, r, err := sdk.NetworkPeeringApi.VerifyPrivateIpMode(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.VerifyPrivateIpMode`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `VerifyPrivateIpMode`: PrivateIPMode
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.VerifyPrivateIpMode`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPrivateIpModeRequest struct via the builder pattern


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

