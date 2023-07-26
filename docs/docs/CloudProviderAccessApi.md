# \CloudProviderAccessApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeCloudProviderAccessRole**](CloudProviderAccessApi.md#AuthorizeCloudProviderAccessRole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
[**CreateCloudProviderAccessRole**](CloudProviderAccessApi.md#CreateCloudProviderAccessRole) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
[**DeauthorizeCloudProviderAccessRole**](CloudProviderAccessApi.md#DeauthorizeCloudProviderAccessRole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
[**GetCloudProviderAccessRole**](CloudProviderAccessApi.md#GetCloudProviderAccessRole) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return specified Cloud Provider Access Role
[**ListCloudProviderAccessRoles**](CloudProviderAccessApi.md#ListCloudProviderAccessRoles) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles



## AuthorizeCloudProviderAccessRole

> CloudProviderAccessRole AuthorizeCloudProviderAccessRole(ctx, groupId, roleId, cloudProviderAccessRole CloudProviderAccessRole).Execute()

Authorize One Cloud Provider Access Role


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    roleId := "roleId_example" // string | 
    cloudProviderAccessRole := *openapiclient.NewCloudProviderAccessRole("ProviderName_example") // CloudProviderAccessRole | 

    resp, r, err := sdk.CloudProviderAccessApi.AuthorizeCloudProviderAccessRole(context.Background(), groupId, roleId, &cloudProviderAccessRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.AuthorizeCloudProviderAccessRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `AuthorizeCloudProviderAccessRole`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.AuthorizeCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudProviderAccessRole** | [**CloudProviderAccessRole**](CloudProviderAccessRole.md) | Grants access to the specified project for the specified AWS IAM role. | 

### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudProviderAccessRole

> CloudProviderAccessRole CreateCloudProviderAccessRole(ctx, groupId, cloudProviderAccessRole CloudProviderAccessRole).Execute()

Create One Cloud Provider Access Role


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    cloudProviderAccessRole := *openapiclient.NewCloudProviderAccessRole("ProviderName_example") // CloudProviderAccessRole | 

    resp, r, err := sdk.CloudProviderAccessApi.CreateCloudProviderAccessRole(context.Background(), groupId, &cloudProviderAccessRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.CreateCloudProviderAccessRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateCloudProviderAccessRole`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.CreateCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProviderAccessRole** | [**CloudProviderAccessRole**](CloudProviderAccessRole.md) | Creates one AWS IAM role. | 

### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeauthorizeCloudProviderAccessRole

> DeauthorizeCloudProviderAccessRole(ctx, groupId, cloudProvider, roleId).Execute()

Deauthorize One Cloud Provider Access Role


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    cloudProvider := "cloudProvider_example" // string | 
    roleId := "roleId_example" // string | 

    r, err := sdk.CloudProviderAccessApi.DeauthorizeCloudProviderAccessRole(context.Background(), groupId, cloudProvider, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.DeauthorizeCloudProviderAccessRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
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

Other parameters are passed through a pointer to a apiDeauthorizeCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderAccessRole

> CloudProviderAccessRole GetCloudProviderAccessRole(ctx, groupId, roleId).Execute()

Return specified Cloud Provider Access Role


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    roleId := "roleId_example" // string | 

    resp, r, err := sdk.CloudProviderAccessApi.GetCloudProviderAccessRole(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.GetCloudProviderAccessRole``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetCloudProviderAccessRole`: CloudProviderAccessRole
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.GetCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudProviderAccessRole**](CloudProviderAccessRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderAccessRoles

> CloudProviderAccessRoles ListCloudProviderAccessRoles(ctx, groupId).Execute()

Return All Cloud Provider Access Roles


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudProviderAccessApi.ListCloudProviderAccessRoles(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.ListCloudProviderAccessRoles``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListCloudProviderAccessRoles`: CloudProviderAccessRoles
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.ListCloudProviderAccessRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderAccessRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProviderAccessRoles**](CloudProviderAccessRoles.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

