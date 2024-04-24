# \ServerlessInstancesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerlessInstance**](ServerlessInstancesApi.md#CreateServerlessInstance) | **Post** /api/atlas/v2/groups/{groupId}/serverless | Create One Serverless Instance in One Project
[**DeleteServerlessInstance**](ServerlessInstancesApi.md#DeleteServerlessInstance) | **Delete** /api/atlas/v2/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project
[**GetServerlessInstance**](ServerlessInstancesApi.md#GetServerlessInstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project
[**ListServerlessInstances**](ServerlessInstancesApi.md#ListServerlessInstances) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances from One Project
[**UpdateServerlessInstance**](ServerlessInstancesApi.md#UpdateServerlessInstance) | **Patch** /api/atlas/v2/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project



## CreateServerlessInstance

> ServerlessInstanceDescription CreateServerlessInstance(ctx, groupId, serverlessInstanceDescriptionCreate ServerlessInstanceDescriptionCreate).Execute()

Create One Serverless Instance in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    serverlessInstanceDescriptionCreate := *openapiclient.NewServerlessInstanceDescriptionCreate("Name_example", *openapiclient.NewServerlessProviderSettings("BackingProviderName_example", "RegionName_example")) // ServerlessInstanceDescriptionCreate | 

    resp, r, err := sdk.ServerlessInstancesApi.CreateServerlessInstance(context.Background(), groupId, &serverlessInstanceDescriptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.CreateServerlessInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateServerlessInstance`: ServerlessInstanceDescription
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.CreateServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverlessInstanceDescriptionCreate** | [**ServerlessInstanceDescriptionCreate**](ServerlessInstanceDescriptionCreate.md) | Create One Serverless Instance in One Project. | 

### Return type

[**ServerlessInstanceDescription**](ServerlessInstanceDescription.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerlessInstance

> map[string]interface{} DeleteServerlessInstance(ctx, groupId, name).Execute()

Remove One Serverless Instance from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 

    resp, r, err := sdk.ServerlessInstancesApi.DeleteServerlessInstance(context.Background(), groupId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.DeleteServerlessInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteServerlessInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.DeleteServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerlessInstance

> ServerlessInstanceDescription GetServerlessInstance(ctx, groupId, name).Execute()

Return One Serverless Instance from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 

    resp, r, err := sdk.ServerlessInstancesApi.GetServerlessInstance(context.Background(), groupId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.GetServerlessInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetServerlessInstance`: ServerlessInstanceDescription
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.GetServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerlessInstanceDescription**](ServerlessInstanceDescription.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerlessInstances

> PaginatedServerlessInstanceDescription ListServerlessInstances(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Serverless Instances from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ServerlessInstancesApi.ListServerlessInstances(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.ListServerlessInstances``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListServerlessInstances`: PaginatedServerlessInstanceDescription
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.ListServerlessInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerlessInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedServerlessInstanceDescription**](PaginatedServerlessInstanceDescription.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerlessInstance

> ServerlessInstanceDescription UpdateServerlessInstance(ctx, groupId, name, serverlessInstanceDescriptionUpdate ServerlessInstanceDescriptionUpdate).Execute()

Update One Serverless Instance in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115011/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    name := "name_example" // string | 
    serverlessInstanceDescriptionUpdate := *openapiclient.NewServerlessInstanceDescriptionUpdate() // ServerlessInstanceDescriptionUpdate | 

    resp, r, err := sdk.ServerlessInstancesApi.UpdateServerlessInstance(context.Background(), groupId, name, &serverlessInstanceDescriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.UpdateServerlessInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateServerlessInstance`: ServerlessInstanceDescription
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.UpdateServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverlessInstanceDescriptionUpdate** | [**ServerlessInstanceDescriptionUpdate**](ServerlessInstanceDescriptionUpdate.md) | Update One Serverless Instance in One Project. | 

### Return type

[**ServerlessInstanceDescription**](ServerlessInstanceDescription.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

