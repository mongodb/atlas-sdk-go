# \MetricIntegrationsAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupMetricIntegration**](MetricIntegrationsAPI.md#CreateGroupMetricIntegration) | **Post** /api/atlas/v2/groups/{groupId}/metricIntegrations | Create One Metric Integration
[**DeleteGroupMetricIntegration**](MetricIntegrationsAPI.md#DeleteGroupMetricIntegration) | **Delete** /api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId} | Remove One Metric Integration
[**GetGroupMetricIntegration**](MetricIntegrationsAPI.md#GetGroupMetricIntegration) | **Get** /api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId} | Return One Metric Integration
[**ListGroupMetricIntegrations**](MetricIntegrationsAPI.md#ListGroupMetricIntegrations) | **Get** /api/atlas/v2/groups/{groupId}/metricIntegrations | Return All Active Metric Integrations
[**UpdateGroupMetricIntegration**](MetricIntegrationsAPI.md#UpdateGroupMetricIntegration) | **Put** /api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId} | Update One Metric Integration



## CreateGroupMetricIntegration

> MetricIntegrationResponse CreateGroupMetricIntegration(ctx, groupId, metricIntegrationRequest MetricIntegrationRequest).Execute()

Create One Metric Integration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    metricIntegrationRequest := *admin.NewMetricIntegrationRequest("AggregationTemporality_example", "AuthType_example", "https://otel-collector.example.com:4318/v1/metrics", "IntegrationType_example", []string{"MetricSelection_example"}, "ProviderType_example") // MetricIntegrationRequest | 

    resp, r, err := sdk.MetricIntegrationsAPI.CreateGroupMetricIntegration(context.Background(), groupId, &metricIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricIntegrationsAPI.CreateGroupMetricIntegration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateGroupMetricIntegration`: MetricIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricIntegrationsAPI.CreateGroupMetricIntegration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupMetricIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricIntegrationRequest** | [**MetricIntegrationRequest**](MetricIntegrationRequest.md) | Metric integration configuration to create. | 

### Return type

[**MetricIntegrationResponse**](MetricIntegrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupMetricIntegration

> DeleteGroupMetricIntegration(ctx, groupId, metricIntegrationId).Execute()

Remove One Metric Integration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    metricIntegrationId := "metricIntegrationId_example" // string | 

    r, err := sdk.MetricIntegrationsAPI.DeleteGroupMetricIntegration(context.Background(), groupId, metricIntegrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricIntegrationsAPI.DeleteGroupMetricIntegration`: %v (%v)\n", err, r)
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
**metricIntegrationId** | **string** | Unique identifier of the metric integration configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMetricIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMetricIntegration

> MetricIntegrationResponse GetGroupMetricIntegration(ctx, groupId, metricIntegrationId).Execute()

Return One Metric Integration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    metricIntegrationId := "metricIntegrationId_example" // string | 

    resp, r, err := sdk.MetricIntegrationsAPI.GetGroupMetricIntegration(context.Background(), groupId, metricIntegrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricIntegrationsAPI.GetGroupMetricIntegration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetGroupMetricIntegration`: MetricIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricIntegrationsAPI.GetGroupMetricIntegration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**metricIntegrationId** | **string** | Unique identifier of the metric integration configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMetricIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetricIntegrationResponse**](MetricIntegrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupMetricIntegrations

> PaginatedMetricIntegrationResponse ListGroupMetricIntegrations(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).IntegrationType(integrationType).ProviderType(providerType).Execute()

Return All Active Metric Integrations


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    integrationType := "integrationType_example" // string |  (optional)
    providerType := "providerType_example" // string |  (optional)

    resp, r, err := sdk.MetricIntegrationsAPI.ListGroupMetricIntegrations(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).IntegrationType(integrationType).ProviderType(providerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricIntegrationsAPI.ListGroupMetricIntegrations`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListGroupMetricIntegrations`: PaginatedMetricIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricIntegrationsAPI.ListGroupMetricIntegrations`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupMetricIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **integrationType** | **string** | Optional filter by integration type (e.g., &#x60;OTEL&#x60;). | 
 **providerType** | **string** | Optional filter by provider type (e.g., &#x60;CUSTOM&#x60;). When specified, &#x60;integrationType&#x60; must also be specified. | 

### Return type

[**PaginatedMetricIntegrationResponse**](PaginatedMetricIntegrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMetricIntegration

> MetricIntegrationResponse UpdateGroupMetricIntegration(ctx, groupId, metricIntegrationId, metricIntegrationUpdateRequest MetricIntegrationUpdateRequest).Execute()

Update One Metric Integration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
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
    metricIntegrationId := "metricIntegrationId_example" // string | 
    metricIntegrationUpdateRequest := *admin.NewMetricIntegrationUpdateRequest("AggregationTemporality_example", "AuthType_example", "https://otel-collector.example.com:4318/v1/metrics", "IntegrationType_example", []string{"MetricSelection_example"}, "ProviderType_example") // MetricIntegrationUpdateRequest | 

    resp, r, err := sdk.MetricIntegrationsAPI.UpdateGroupMetricIntegration(context.Background(), groupId, metricIntegrationId, &metricIntegrationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricIntegrationsAPI.UpdateGroupMetricIntegration`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateGroupMetricIntegration`: MetricIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricIntegrationsAPI.UpdateGroupMetricIntegration`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**metricIntegrationId** | **string** | Unique identifier of the metric integration configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMetricIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricIntegrationUpdateRequest** | [**MetricIntegrationUpdateRequest**](MetricIntegrationUpdateRequest.md) | Updated metric integration configuration. | 

### Return type

[**MetricIntegrationResponse**](MetricIntegrationResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

