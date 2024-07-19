# \ProjectIPAccessListApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectIpAccessList**](ProjectIPAccessListApi.md#CreateProjectIpAccessList) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List
[**DeleteProjectIpAccessList**](ProjectIPAccessListApi.md#DeleteProjectIpAccessList) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List
[**GetProjectIpAccessListStatus**](ProjectIPAccessListApi.md#GetProjectIpAccessListStatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry
[**GetProjectIpList**](ProjectIPAccessListApi.md#GetProjectIpList) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry
[**ListProjectIpAccessLists**](ProjectIPAccessListApi.md#ListProjectIpAccessLists) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return Project IP Access List



## CreateProjectIpAccessList

> PaginatedNetworkAccess CreateProjectIpAccessList(ctx, groupId, networkPermissionEntry []NetworkPermissionEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Add Entries to Project IP Access List


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240530004/admin"
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
    networkPermissionEntry := []openapiclient.NetworkPermissionEntry{*openapiclient.NewNetworkPermissionEntry()} // []NetworkPermissionEntry | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProjectIPAccessListApi.CreateProjectIpAccessList(context.Background(), groupId, &networkPermissionEntry).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.CreateProjectIpAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateProjectIpAccessList`: PaginatedNetworkAccess
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.CreateProjectIpAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectIpAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkPermissionEntry** | [**[]NetworkPermissionEntry**](NetworkPermissionEntry.md) | One or more access list entries to add to the specified project. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedNetworkAccess**](PaginatedNetworkAccess.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectIpAccessList

> interface{} DeleteProjectIpAccessList(ctx, groupId, entryValue).Execute()

Remove One Entry from One Project IP Access List


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240530004/admin"
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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | 

    resp, r, err := sdk.ProjectIPAccessListApi.DeleteProjectIpAccessList(context.Background(), groupId, entryValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.DeleteProjectIpAccessList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteProjectIpAccessList`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.DeleteProjectIpAccessList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**entryValue** | **string** | Access list entry that you want to remove from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:  - how your application established the connection, - how MongoDB Cloud or the driver using the address behaves, and - which protocol (like TCP or UDP) the connection uses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectIpAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectIpAccessListStatus

> NetworkPermissionEntryStatus GetProjectIpAccessListStatus(ctx, groupId, entryValue).Execute()

Return Status of One Project IP Access List Entry


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240530004/admin"
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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | 

    resp, r, err := sdk.ProjectIPAccessListApi.GetProjectIpAccessListStatus(context.Background(), groupId, entryValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.GetProjectIpAccessListStatus`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetProjectIpAccessListStatus`: NetworkPermissionEntryStatus
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.GetProjectIpAccessListStatus`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**entryValue** | **string** | Network address or cloud provider security construct that identifies which project access list entry to be verified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectIpAccessListStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkPermissionEntryStatus**](NetworkPermissionEntryStatus.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectIpList

> NetworkPermissionEntry GetProjectIpList(ctx, groupId, entryValue).Execute()

Return One Project IP Access List Entry


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240530004/admin"
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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | 

    resp, r, err := sdk.ProjectIPAccessListApi.GetProjectIpList(context.Background(), groupId, entryValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.GetProjectIpList`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetProjectIpList`: NetworkPermissionEntry
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.GetProjectIpList`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**entryValue** | **string** | Access list entry that you want to return from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkPermissionEntry**](NetworkPermissionEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectIpAccessLists

> PaginatedNetworkAccess ListProjectIpAccessLists(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return Project IP Access List


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240530004/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProjectIPAccessListApi.ListProjectIpAccessLists(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.ListProjectIpAccessLists`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListProjectIpAccessLists`: PaginatedNetworkAccess
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.ListProjectIpAccessLists`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectIpAccessListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedNetworkAccess**](PaginatedNetworkAccess.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

