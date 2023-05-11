# \ServerlessPrivateEndpointsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerlessPrivateEndpoint**](ServerlessPrivateEndpointsApi.md#CreateServerlessPrivateEndpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance
[**DeleteServerlessPrivateEndpoint**](ServerlessPrivateEndpointsApi.md#DeleteServerlessPrivateEndpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance
[**GetServerlessPrivateEndpoint**](ServerlessPrivateEndpointsApi.md#GetServerlessPrivateEndpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance
[**ListServerlessPrivateEndpoints**](ServerlessPrivateEndpointsApi.md#ListServerlessPrivateEndpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance
[**UpdateServerlessPrivateEndpoint**](ServerlessPrivateEndpointsApi.md#UpdateServerlessPrivateEndpoint) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance



## CreateServerlessPrivateEndpoint

> ServerlessTenantEndpoint CreateServerlessPrivateEndpoint(ctx, groupId, instanceName).ServerlessTenantEndpointCreate(serverlessTenantEndpointCreate).Execute()

Create One Private Endpoint for One Serverless Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    instanceName := "instanceName_example" // string | 
    serverlessTenantEndpointCreate := *openapiclient.NewServerlessTenantEndpointCreate() // ServerlessTenantEndpointCreate | 

    resp, r, err := sdk.ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpoint(context.Background(), groupId, instanceName).ServerlessTenantEndpointCreate(serverlessTenantEndpointCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpoint``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateServerlessPrivateEndpoint`: ServerlessTenantEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance for which the tenant endpoint will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerlessPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverlessTenantEndpointCreate** | [**ServerlessTenantEndpointCreate**](ServerlessTenantEndpointCreate.md) | Information about the Private Endpoint to create for the Serverless Instance. | 

### Return type

[**ServerlessTenantEndpoint**](ServerlessTenantEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerlessPrivateEndpoint

> DeleteServerlessPrivateEndpoint(ctx, groupId, instanceName, endpointId).Execute()

Remove One Private Endpoint for One Serverless Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    instanceName := "instanceName_example" // string | 
    endpointId := "endpointId_example" // string | 

    r, err := sdk.ServerlessPrivateEndpointsApi.DeleteServerlessPrivateEndpoint(context.Background(), groupId, instanceName, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.DeleteServerlessPrivateEndpoint``: %v\n", err)
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
**instanceName** | **string** | Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerlessPrivateEndpointRequest struct via the builder pattern


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


## GetServerlessPrivateEndpoint

> ServerlessTenantEndpoint GetServerlessPrivateEndpoint(ctx, groupId, instanceName, endpointId).Execute()

Return One Private Endpoint for One Serverless Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    instanceName := "instanceName_example" // string | 
    endpointId := "endpointId_example" // string | 

    resp, r, err := sdk.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(), groupId, instanceName, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetServerlessPrivateEndpoint`: ServerlessTenantEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerlessTenantEndpoint**](ServerlessTenantEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerlessPrivateEndpoints

> []ServerlessTenantEndpoint ListServerlessPrivateEndpoints(ctx, groupId, instanceName).Execute()

Return All Private Endpoints for One Serverless Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    instanceName := "instanceName_example" // string | 

    resp, r, err := sdk.ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints(context.Background(), groupId, instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListServerlessPrivateEndpoints`: []ServerlessTenantEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerlessPrivateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ServerlessTenantEndpoint**](ServerlessTenantEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerlessPrivateEndpoint

> ServerlessTenantEndpoint UpdateServerlessPrivateEndpoint(ctx, groupId, instanceName, endpointId).ServerlessTenantEndpointUpdate(serverlessTenantEndpointUpdate).Execute()

Update One Private Endpoint for One Serverless Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MDB_API_KEY")
    apiSecret := os.Getenv("MDB_API_SECRET")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    instanceName := "instanceName_example" // string | 
    endpointId := "endpointId_example" // string | 
    serverlessTenantEndpointUpdate := openapiclient.ServerlessTenantEndpointUpdate{ServerlessAWSTenantEndpointUpdate: openapiclient.NewServerlessAWSTenantEndpointUpdate("ProviderName_example")} // ServerlessTenantEndpointUpdate |  (optional)

    resp, r, err := sdk.ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint(context.Background(), groupId, instanceName, endpointId).ServerlessTenantEndpointUpdate(serverlessTenantEndpointUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateServerlessPrivateEndpoint`: ServerlessTenantEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerlessPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serverlessTenantEndpointUpdate** | [**ServerlessTenantEndpointUpdate**](ServerlessTenantEndpointUpdate.md) |  | 

### Return type

[**ServerlessTenantEndpoint**](ServerlessTenantEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

