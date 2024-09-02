# \EncryptionAtRestUsingCustomerKeyManagementApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEncryptionAtRestPrivateEndpoint**](EncryptionAtRestUsingCustomerKeyManagementApi.md#CreateEncryptionAtRestPrivateEndpoint) | **Post** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Create One Private Endpoint in a Specified Region for Encryption at Rest Using Customer Key Management
[**GetEncryptionAtRest**](EncryptionAtRestUsingCustomerKeyManagementApi.md#GetEncryptionAtRest) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
[**GetEncryptionAtRestPrivateEndpoint**](EncryptionAtRestUsingCustomerKeyManagementApi.md#GetEncryptionAtRestPrivateEndpoint) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Return One Private Endpoint for Encryption at Rest Using Customer Key Management
[**GetEncryptionAtRestPrivateEndpointsForCloudProvider**](EncryptionAtRestUsingCustomerKeyManagementApi.md#GetEncryptionAtRestPrivateEndpointsForCloudProvider) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints | Return Private Endpoints of a Cloud Provider for Encryption at Rest Using Customer Key Management for One Project
[**RequestEncryptionAtRestPrivateEndpointDeletion**](EncryptionAtRestUsingCustomerKeyManagementApi.md#RequestEncryptionAtRestPrivateEndpointDeletion) | **Delete** /api/atlas/v2/groups/{groupId}/encryptionAtRest/{cloudProvider}/privateEndpoints/{endpointId} | Delete Private Endpoint for Encryption at Rest Using Customer Key Management
[**UpdateEncryptionAtRest**](EncryptionAtRestUsingCustomerKeyManagementApi.md#UpdateEncryptionAtRest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project



## CreateEncryptionAtRestPrivateEndpoint

> EARPrivateEndpoint CreateEncryptionAtRestPrivateEndpoint(ctx, groupId, cloudProvider, eARPrivateEndpoint EARPrivateEndpoint).Execute()

Create One Private Endpoint in a Specified Region for Encryption at Rest Using Customer Key Management


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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
    eARPrivateEndpoint := *openapiclient.NewEARPrivateEndpoint() // EARPrivateEndpoint | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.CreateEncryptionAtRestPrivateEndpoint(context.Background(), groupId, cloudProvider, &eARPrivateEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.CreateEncryptionAtRestPrivateEndpoint`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateEncryptionAtRestPrivateEndpoint`: EARPrivateEndpoint
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.CreateEncryptionAtRestPrivateEndpoint`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider for the private endpoint to create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEncryptionAtRestPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eARPrivateEndpoint** | [**EARPrivateEndpoint**](EARPrivateEndpoint.md) | Creates a private endpoint in the specified region. | 

### Return type

[**EARPrivateEndpoint**](EARPrivateEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionAtRest

> EncryptionAtRest GetEncryptionAtRest(ctx, groupId).Execute()

Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetEncryptionAtRest`: EncryptionAtRest
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionAtRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EncryptionAtRest**](EncryptionAtRest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionAtRestPrivateEndpoint

> EARPrivateEndpoint GetEncryptionAtRestPrivateEndpoint(ctx, groupId, cloudProvider, endpointId).Execute()

Return One Private Endpoint for Encryption at Rest Using Customer Key Management


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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
    endpointId := "endpointId_example" // string | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpoint(context.Background(), groupId, cloudProvider, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpoint`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetEncryptionAtRestPrivateEndpoint`: EARPrivateEndpoint
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpoint`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider of the private endpoint. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionAtRestPrivateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EARPrivateEndpoint**](EARPrivateEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionAtRestPrivateEndpointsForCloudProvider

> PaginatedApiAtlasEARPrivateEndpoint GetEncryptionAtRestPrivateEndpointsForCloudProvider(ctx, groupId, cloudProvider).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return Private Endpoints of a Cloud Provider for Encryption at Rest Using Customer Key Management for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpointsForCloudProvider(context.Background(), groupId, cloudProvider).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpointsForCloudProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetEncryptionAtRestPrivateEndpointsForCloudProvider`: PaginatedApiAtlasEARPrivateEndpoint
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestPrivateEndpointsForCloudProvider`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider for the private endpoints to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionAtRestPrivateEndpointsForCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasEARPrivateEndpoint**](PaginatedApiAtlasEARPrivateEndpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestEncryptionAtRestPrivateEndpointDeletion

> interface{} RequestEncryptionAtRestPrivateEndpointDeletion(ctx, groupId, cloudProvider, endpointId).Execute()

Delete Private Endpoint for Encryption at Rest Using Customer Key Management


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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
    endpointId := "endpointId_example" // string | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.RequestEncryptionAtRestPrivateEndpointDeletion(context.Background(), groupId, cloudProvider, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.RequestEncryptionAtRestPrivateEndpointDeletion`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RequestEncryptionAtRestPrivateEndpointDeletion`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.RequestEncryptionAtRestPrivateEndpointDeletion`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider of the private endpoint to delete. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestEncryptionAtRestPrivateEndpointDeletionRequest struct via the builder pattern


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


## UpdateEncryptionAtRest

> EncryptionAtRest UpdateEncryptionAtRest(ctx, groupId, encryptionAtRest EncryptionAtRest).Execute()

Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805003/admin"
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
    encryptionAtRest := *openapiclient.NewEncryptionAtRest() // EncryptionAtRest | 

    resp, r, err := sdk.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), groupId, &encryptionAtRest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateEncryptionAtRest`: EncryptionAtRest
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEncryptionAtRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encryptionAtRest** | [**EncryptionAtRest**](EncryptionAtRest.md) | Required parameters depend on whether someone has enabled Encryption at Rest using Customer Key Management:  If you have enabled Encryption at Rest using Customer Key Management (CMK), Atlas requires all of the parameters for the desired encryption provider.  - To use AWS Key Management Service (KMS), MongoDB Cloud requires all the fields in the **awsKms** object. - To use Azure Key Vault, MongoDB Cloud requires all the fields in the **azureKeyVault** object. - To use Google Cloud Key Management Service (KMS), MongoDB Cloud requires all the fields in the **googleCloudKms** object.  If you enabled Encryption at Rest using Customer Key  Management, administrators can pass only the changed fields for the **awsKms**, **azureKeyVault**, or **googleCloudKms** object to update the configuration to this endpoint. | 

### Return type

[**EncryptionAtRest**](EncryptionAtRest.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

