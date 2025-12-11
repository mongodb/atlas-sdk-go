# \ProgrammaticAPIKeysApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupApiKey**](ProgrammaticAPIKeysApi.md#AddGroupApiKey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project
[**CreateGroupApiKey**](ProgrammaticAPIKeysApi.md#CreateGroupApiKey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project
[**CreateOrgAccessEntry**](ProgrammaticAPIKeysApi.md#CreateOrgAccessEntry) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create One Access List Entry for One Organization API Key
[**CreateOrgApiKey**](ProgrammaticAPIKeysApi.md#CreateOrgApiKey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key
[**DeleteAccessEntry**](ProgrammaticAPIKeysApi.md#DeleteAccessEntry) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key
[**DeleteOrgApiKey**](ProgrammaticAPIKeysApi.md#DeleteOrgApiKey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key
[**GetOrgAccessEntry**](ProgrammaticAPIKeysApi.md#GetOrgAccessEntry) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key
[**GetOrgApiKey**](ProgrammaticAPIKeysApi.md#GetOrgApiKey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key
[**ListGroupApiKeys**](ProgrammaticAPIKeysApi.md#ListGroupApiKeys) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project
[**ListOrgAccessEntries**](ProgrammaticAPIKeysApi.md#ListOrgAccessEntries) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key
[**ListOrgApiKeys**](ProgrammaticAPIKeysApi.md#ListOrgApiKeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys
[**RemoveGroupApiKey**](ProgrammaticAPIKeysApi.md#RemoveGroupApiKey) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project
[**UpdateApiKeyRoles**](ProgrammaticAPIKeysApi.md#UpdateApiKeyRoles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Organization API Key Roles for One Project
[**UpdateOrgApiKey**](ProgrammaticAPIKeysApi.md#UpdateOrgApiKey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key



## AddGroupApiKey

> AddGroupApiKey(ctx, groupId, apiUserId, userAccessRoleAssignment []UserAccessRoleAssignment).Execute()

Assign One Organization API Key to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    apiUserId := "apiUserId_example" // string | 
    userAccessRoleAssignment := []openapiclient.UserAccessRoleAssignment{*openapiclient.NewUserAccessRoleAssignment()} // []UserAccessRoleAssignment | 

    r, err := sdk.ProgrammaticAPIKeysApi.AddGroupApiKey(context.Background(), groupId, apiUserId, &userAccessRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.AddGroupApiKey`: %v (%v)\n", err, r)
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
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userAccessRoleAssignment** | [**[]UserAccessRoleAssignment**](UserAccessRoleAssignment.md) | Organization API key to be assigned to the specified project. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupApiKey

> ApiKeyUserDetails CreateGroupApiKey(ctx, groupId, createAtlasProjectApiKey CreateAtlasProjectApiKey).Execute()

Create and Assign One Organization API Key to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    createAtlasProjectApiKey := *openapiclient.NewCreateAtlasProjectApiKey("Desc_example", []string{"Roles_example"}) // CreateAtlasProjectApiKey | 

    resp, r, err := sdk.ProgrammaticAPIKeysApi.CreateGroupApiKey(context.Background(), groupId, &createAtlasProjectApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateGroupApiKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupApiKey`: ApiKeyUserDetails
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateGroupApiKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAtlasProjectApiKey** | [**CreateAtlasProjectApiKey**](CreateAtlasProjectApiKey.md) | Organization API key to be created and assigned to the specified project. | 

### Return type

[**ApiKeyUserDetails**](ApiKeyUserDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgAccessEntry

> PaginatedApiUserAccessListResponse CreateOrgAccessEntry(ctx, orgId, apiUserId, userAccessListRequest []UserAccessListRequest).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Create One Access List Entry for One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 
    userAccessListRequest := []openapiclient.UserAccessListRequest{*openapiclient.NewUserAccessListRequest()} // []UserAccessListRequest | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProgrammaticAPIKeysApi.CreateOrgAccessEntry(context.Background(), orgId, apiUserId, &userAccessListRequest).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateOrgAccessEntry`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgAccessEntry`: PaginatedApiUserAccessListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateOrgAccessEntry`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAccessEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userAccessListRequest** | [**[]UserAccessListRequest**](UserAccessListRequest.md) | Access list entries to be created for the specified organization API key. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiUserAccessListResponse**](PaginatedApiUserAccessListResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgApiKey

> ApiKeyUserDetails CreateOrgApiKey(ctx, orgId, createAtlasOrganizationApiKey CreateAtlasOrganizationApiKey).Execute()

Create One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    createAtlasOrganizationApiKey := *openapiclient.NewCreateAtlasOrganizationApiKey("Desc_example", []string{"Roles_example"}) // CreateAtlasOrganizationApiKey | 

    resp, r, err := sdk.ProgrammaticAPIKeysApi.CreateOrgApiKey(context.Background(), orgId, &createAtlasOrganizationApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateOrgApiKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgApiKey`: ApiKeyUserDetails
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateOrgApiKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAtlasOrganizationApiKey** | [**CreateAtlasOrganizationApiKey**](CreateAtlasOrganizationApiKey.md) | Organization API Key to be created. | 

### Return type

[**ApiKeyUserDetails**](ApiKeyUserDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessEntry

> DeleteAccessEntry(ctx, orgId, apiUserId, ipAddress).Execute()

Remove One Access List Entry for One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 
    ipAddress := "192.0.2.0%2F24" // string | 

    r, err := sdk.ProgrammaticAPIKeysApi.DeleteAccessEntry(context.Background(), orgId, apiUserId, ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.DeleteAccessEntry`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessEntryRequest struct via the builder pattern


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


## DeleteOrgApiKey

> DeleteOrgApiKey(ctx, orgId, apiUserId).Execute()

Remove One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 

    r, err := sdk.ProgrammaticAPIKeysApi.DeleteOrgApiKey(context.Background(), orgId, apiUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.DeleteOrgApiKey`: %v (%v)\n", err, r)
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
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgApiKeyRequest struct via the builder pattern


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


## GetOrgAccessEntry

> UserAccessListResponse GetOrgAccessEntry(ctx, orgId, ipAddress, apiUserId).Execute()

Return One Access List Entry for One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    ipAddress := "192.0.2.0%2F24" // string | 
    apiUserId := "apiUserId_example" // string | 

    resp, r, err := sdk.ProgrammaticAPIKeysApi.GetOrgAccessEntry(context.Background(), orgId, ipAddress, apiUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.GetOrgAccessEntry`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgAccessEntry`: UserAccessListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.GetOrgAccessEntry`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAccessEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserAccessListResponse**](UserAccessListResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgApiKey

> ApiKeyUserDetails GetOrgApiKey(ctx, orgId, apiUserId).Execute()

Return One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 

    resp, r, err := sdk.ProgrammaticAPIKeysApi.GetOrgApiKey(context.Background(), orgId, apiUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.GetOrgApiKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgApiKey`: ApiKeyUserDetails
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.GetOrgApiKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiKeyUserDetails**](ApiKeyUserDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupApiKeys

> PaginatedApiApiUser ListGroupApiKeys(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Organization API Keys Assigned to One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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

    resp, r, err := sdk.ProgrammaticAPIKeysApi.ListGroupApiKeys(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ListGroupApiKeys`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupApiKeys`: PaginatedApiApiUser
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ListGroupApiKeys`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiApiUser**](PaginatedApiApiUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAccessEntries

> PaginatedApiUserAccessListResponse ListOrgAccessEntries(ctx, orgId, apiUserId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Access List Entries for One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProgrammaticAPIKeysApi.ListOrgAccessEntries(context.Background(), orgId, apiUserId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ListOrgAccessEntries`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgAccessEntries`: PaginatedApiUserAccessListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ListOrgAccessEntries`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAccessEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiUserAccessListResponse**](PaginatedApiUserAccessListResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgApiKeys

> PaginatedApiApiUser ListOrgApiKeys(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Organization API Keys


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.ProgrammaticAPIKeysApi.ListOrgApiKeys(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ListOrgApiKeys`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgApiKeys`: PaginatedApiApiUser
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ListOrgApiKeys`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiApiUser**](PaginatedApiApiUser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupApiKey

> RemoveGroupApiKey(ctx, groupId, apiUserId).Execute()

Unassign One Organization API Key from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    apiUserId := "apiUserId_example" // string | 

    r, err := sdk.ProgrammaticAPIKeysApi.RemoveGroupApiKey(context.Background(), groupId, apiUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.RemoveGroupApiKey`: %v (%v)\n", err, r)
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
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupApiKeyRequest struct via the builder pattern


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


## UpdateApiKeyRoles

> ApiKeyUserDetails UpdateApiKeyRoles(ctx, groupId, apiUserId, updateAtlasProjectApiKey UpdateAtlasProjectApiKey).PageNum(pageNum).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).Execute()

Update Organization API Key Roles for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
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
    apiUserId := "apiUserId_example" // string | 
    updateAtlasProjectApiKey := *openapiclient.NewUpdateAtlasProjectApiKey() // UpdateAtlasProjectApiKey | 
    pageNum := int(56) // int |  (optional) (default to 1)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    includeCount := true // bool |  (optional) (default to true)

    resp, r, err := sdk.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), groupId, apiUserId, &updateAtlasProjectApiKey).PageNum(pageNum).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.UpdateApiKeyRoles`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateApiKeyRoles`: ApiKeyUserDetails
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.UpdateApiKeyRoles`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAtlasProjectApiKey** | [**UpdateAtlasProjectApiKey**](UpdateAtlasProjectApiKey.md) | Organization API Key to be updated. This request requires a minimum of one of the two body parameters. | 
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]

### Return type

[**ApiKeyUserDetails**](ApiKeyUserDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgApiKey

> ApiKeyUserDetails UpdateOrgApiKey(ctx, orgId, apiUserId, updateAtlasOrganizationApiKey UpdateAtlasOrganizationApiKey).Execute()

Update One Organization API Key


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    apiUserId := "apiUserId_example" // string | 
    updateAtlasOrganizationApiKey := *openapiclient.NewUpdateAtlasOrganizationApiKey() // UpdateAtlasOrganizationApiKey | 

    resp, r, err := sdk.ProgrammaticAPIKeysApi.UpdateOrgApiKey(context.Background(), orgId, apiUserId, &updateAtlasOrganizationApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.UpdateOrgApiKey`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOrgApiKey`: ApiKeyUserDetails
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.UpdateOrgApiKey`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key you  want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAtlasOrganizationApiKey** | [**UpdateAtlasOrganizationApiKey**](UpdateAtlasOrganizationApiKey.md) | Organization API key to be updated. This request requires a minimum of one of the two body parameters. | 

### Return type

[**ApiKeyUserDetails**](ApiKeyUserDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

