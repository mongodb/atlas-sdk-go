# \CustomDatabaseRolesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomDatabaseRole**](CustomDatabaseRolesApi.md#CreateCustomDatabaseRole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role
[**DeleteCustomDatabaseRole**](CustomDatabaseRolesApi.md#DeleteCustomDatabaseRole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
[**GetCustomDatabaseRole**](CustomDatabaseRolesApi.md#GetCustomDatabaseRole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
[**ListCustomDatabaseRoles**](CustomDatabaseRolesApi.md#ListCustomDatabaseRoles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
[**UpdateCustomDatabaseRole**](CustomDatabaseRolesApi.md#UpdateCustomDatabaseRole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project



## CreateCustomDatabaseRole

> UserCustomDBRole CreateCustomDatabaseRole(ctx, groupId, userCustomDBRole UserCustomDBRole).Execute()

Create One Custom Role


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    userCustomDBRole := *openapiclient.NewUserCustomDBRole("RoleName_example") // UserCustomDBRole | 

    resp, r, err := sdk.CustomDatabaseRolesApi.CreateCustomDatabaseRole(context.Background(), groupId, &userCustomDBRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.CreateCustomDatabaseRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCustomDatabaseRole`: UserCustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.CreateCustomDatabaseRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCustomDBRole** | [**UserCustomDBRole**](UserCustomDBRole.md) | Creates one custom role in the specified project. | 

### Return type

[**UserCustomDBRole**](UserCustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDatabaseRole

> DeleteCustomDatabaseRole(ctx, groupId, roleName).Execute()

Remove One Custom Role from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    roleName := "roleName_example" // string | 

    r, err := sdk.CustomDatabaseRolesApi.DeleteCustomDatabaseRole(context.Background(), groupId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.DeleteCustomDatabaseRole`: %v (%v)\n", err, r)
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
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomDatabaseRoleRequest struct via the builder pattern


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


## GetCustomDatabaseRole

> UserCustomDBRole GetCustomDatabaseRole(ctx, groupId, roleName).Execute()

Return One Custom Role in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    roleName := "roleName_example" // string | 

    resp, r, err := sdk.CustomDatabaseRolesApi.GetCustomDatabaseRole(context.Background(), groupId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.GetCustomDatabaseRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetCustomDatabaseRole`: UserCustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.GetCustomDatabaseRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserCustomDBRole**](UserCustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomDatabaseRoles

> []UserCustomDBRole ListCustomDatabaseRoles(ctx, groupId).Execute()

Return All Custom Roles in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CustomDatabaseRolesApi.ListCustomDatabaseRoles(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.ListCustomDatabaseRoles`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListCustomDatabaseRoles`: []UserCustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.ListCustomDatabaseRoles`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomDatabaseRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserCustomDBRole**](UserCustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomDatabaseRole

> UserCustomDBRole UpdateCustomDatabaseRole(ctx, groupId, roleName, updateCustomDBRole UpdateCustomDBRole).Execute()

Update One Custom Role in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    roleName := "roleName_example" // string | 
    updateCustomDBRole := *openapiclient.NewUpdateCustomDBRole() // UpdateCustomDBRole | 

    resp, r, err := sdk.CustomDatabaseRolesApi.UpdateCustomDatabaseRole(context.Background(), groupId, roleName, &updateCustomDBRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.UpdateCustomDatabaseRole`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateCustomDatabaseRole`: UserCustomDBRole
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.UpdateCustomDatabaseRole`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomDatabaseRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCustomDBRole** | [**UpdateCustomDBRole**](UpdateCustomDBRole.md) | Updates one custom role in the specified project. | 

### Return type

[**UserCustomDBRole**](UserCustomDBRole.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

