# \ThirdPartyIntegrationsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateThirdPartyIntegration**](ThirdPartyIntegrationsApi.md#CreateThirdPartyIntegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration
[**DeleteThirdPartyIntegration**](ThirdPartyIntegrationsApi.md#DeleteThirdPartyIntegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration
[**GetThirdPartyIntegration**](ThirdPartyIntegrationsApi.md#GetThirdPartyIntegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration
[**ListThirdPartyIntegrations**](ThirdPartyIntegrationsApi.md#ListThirdPartyIntegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations
[**UpdateThirdPartyIntegration**](ThirdPartyIntegrationsApi.md#UpdateThirdPartyIntegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration



## CreateThirdPartyIntegration

> PaginatedIntegration CreateThirdPartyIntegration(ctx, integrationType, groupId, integration Integration).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Configure One Third-Party Service Integration



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
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    integrationType := "integrationType_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    integration := openapiclient.Integration{Datadog: openapiclient.NewDatadog("****************************a23c")} // Integration | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ThirdPartyIntegrationsApi.CreateThirdPartyIntegration(context.Background(), integrationType, groupId, &integration).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyIntegrationsApi.CreateThirdPartyIntegration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateThirdPartyIntegration`: PaginatedIntegration
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyIntegrationsApi.CreateThirdPartyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateThirdPartyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integration** | [**Integration**](Integration.md) | Third-party integration that you want to configure for your project. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedIntegration**](PaginatedIntegration.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThirdPartyIntegration

> map[string]interface{} DeleteThirdPartyIntegration(ctx, integrationType, groupId).Execute()

Remove One Third-Party Service Integration



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
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    integrationType := "integrationType_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.ThirdPartyIntegrationsApi.DeleteThirdPartyIntegration(context.Background(), integrationType, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyIntegrationsApi.DeleteThirdPartyIntegration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteThirdPartyIntegration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyIntegrationsApi.DeleteThirdPartyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThirdPartyIntegrationRequest struct via the builder pattern


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


## GetThirdPartyIntegration

> Integration GetThirdPartyIntegration(ctx, groupId, integrationType).Execute()

Return One Third-Party Service Integration



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
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    integrationType := "integrationType_example" // string | 

    resp, r, err := sdk.ThirdPartyIntegrationsApi.GetThirdPartyIntegration(context.Background(), groupId, integrationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyIntegrationsApi.GetThirdPartyIntegration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetThirdPartyIntegration`: Integration
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyIntegrationsApi.GetThirdPartyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Integration**](Integration.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThirdPartyIntegrations

> PaginatedIntegration ListThirdPartyIntegrations(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Active Third-Party Service Integrations



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
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ThirdPartyIntegrationsApi.ListThirdPartyIntegrations(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyIntegrationsApi.ListThirdPartyIntegrations``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListThirdPartyIntegrations`: PaginatedIntegration
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyIntegrationsApi.ListThirdPartyIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListThirdPartyIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedIntegration**](PaginatedIntegration.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThirdPartyIntegration

> PaginatedIntegration UpdateThirdPartyIntegration(ctx, integrationType, groupId, integration Integration).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Update One Third-Party Service Integration



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
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    integrationType := "integrationType_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    integration := openapiclient.Integration{Datadog: openapiclient.NewDatadog("****************************a23c")} // Integration | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.ThirdPartyIntegrationsApi.UpdateThirdPartyIntegration(context.Background(), integrationType, groupId, &integration).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyIntegrationsApi.UpdateThirdPartyIntegration``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateThirdPartyIntegration`: PaginatedIntegration
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyIntegrationsApi.UpdateThirdPartyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThirdPartyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integration** | [**Integration**](Integration.md) | Third-party integration that you want to configure for your project. | 
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedIntegration**](PaginatedIntegration.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

