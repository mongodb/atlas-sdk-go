# \CloudProviderAccessApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeProviderAccessRole**](CloudProviderAccessApi.md#AuthorizeProviderAccessRole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
[**CreateCloudProviderAccess**](CloudProviderAccessApi.md#CreateCloudProviderAccess) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
[**DeauthorizeProviderAccessRole**](CloudProviderAccessApi.md#DeauthorizeProviderAccessRole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
[**GetCloudProviderAccess**](CloudProviderAccessApi.md#GetCloudProviderAccess) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return One Cloud Provider Access Role
[**ListCloudProviderAccess**](CloudProviderAccessApi.md#ListCloudProviderAccess) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles



## AuthorizeProviderAccessRole

> CloudProviderAccessRole AuthorizeProviderAccessRole(ctx, groupId, roleId, cloudProviderAccessRoleRequestUpdate CloudProviderAccessRoleRequestUpdate).Execute()

Authorize One Cloud Provider Access Role


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
    roleId := "roleId_example" // string | 
    cloudProviderAccessRoleRequestUpdate := *openapiclient.NewCloudProviderAccessRoleRequestUpdate("ProviderName_example") // CloudProviderAccessRoleRequestUpdate | 

    resp, r, err := sdk.CloudProviderAccessApi.AuthorizeProviderAccessRole(context.Background(), groupId, roleId, &cloudProviderAccessRoleRequestUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.AuthorizeProviderAccessRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `AuthorizeProviderAccessRole`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.AuthorizeProviderAccessRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudProviderAccessRoleRequestUpdate** | [**CloudProviderAccessRoleRequestUpdate**](CloudProviderAccessRoleRequestUpdate.md) | Grants access to the specified project for the specified access role. | 

### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudProviderAccess

> CloudProviderAccessRole CreateCloudProviderAccess(ctx, groupId, cloudProviderAccessRoleRequest CloudProviderAccessRoleRequest).Execute()

Create One Cloud Provider Access Role


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
    cloudProviderAccessRoleRequest := *openapiclient.NewCloudProviderAccessRoleRequest("ProviderName_example") // CloudProviderAccessRoleRequest | 

    resp, r, err := sdk.CloudProviderAccessApi.CreateCloudProviderAccess(context.Background(), groupId, &cloudProviderAccessRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.CreateCloudProviderAccess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCloudProviderAccess`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.CreateCloudProviderAccess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudProviderAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProviderAccessRoleRequest** | [**CloudProviderAccessRoleRequest**](CloudProviderAccessRoleRequest.md) | Creates one role for the specified cloud provider. | 

### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeauthorizeProviderAccessRole

> DeauthorizeProviderAccessRole(ctx, groupId, cloudProvider, roleId).Execute()

Deauthorize One Cloud Provider Access Role


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
    cloudProvider := "cloudProvider_example" // string | 
    roleId := "roleId_example" // string | 

    r, err := sdk.CloudProviderAccessApi.DeauthorizeProviderAccessRole(context.Background(), groupId, cloudProvider, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.DeauthorizeProviderAccessRole`: %v (%v)\n", err, r)
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
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider of the role to deauthorize. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeauthorizeProviderAccessRoleRequest struct via the builder pattern


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


## GetCloudProviderAccess

> CloudProviderAccessRole GetCloudProviderAccess(ctx, groupId, roleId).Execute()

Return One Cloud Provider Access Role


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
    roleId := "roleId_example" // string | 

    resp, r, err := sdk.CloudProviderAccessApi.GetCloudProviderAccess(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.GetCloudProviderAccess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetCloudProviderAccess`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.GetCloudProviderAccess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderAccess

> CloudProviderAccessRoles ListCloudProviderAccess(ctx, groupId).Execute()

Return All Cloud Provider Access Roles


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

    resp, r, err := sdk.CloudProviderAccessApi.ListCloudProviderAccess(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.ListCloudProviderAccess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListCloudProviderAccess`: CloudProviderAccessRoles
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.ListCloudProviderAccess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProviderAccessRoles**](CloudProviderAccessRoles.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

