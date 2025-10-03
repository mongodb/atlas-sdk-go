# \AlertConfigurationsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertConfig**](AlertConfigurationsApi.md#CreateAlertConfig) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project
[**DeleteAlertConfig**](AlertConfigurationsApi.md#DeleteAlertConfig) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project
[**GetAlertConfig**](AlertConfigurationsApi.md#GetAlertConfig) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project
[**GetAlertConfigs**](AlertConfigurationsApi.md#GetAlertConfigs) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs | Return All Alert Configurations Set for One Alert
[**ListAlertConfigs**](AlertConfigurationsApi.md#ListAlertConfigs) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations in One Project
[**ListMatcherFieldNames**](AlertConfigurationsApi.md#ListMatcherFieldNames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Return All Alert Configuration Matchers Field Names
[**ToggleAlertConfig**](AlertConfigurationsApi.md#ToggleAlertConfig) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle State of One Alert Configuration in One Project
[**UpdateAlertConfig**](AlertConfigurationsApi.md#UpdateAlertConfig) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration in One Project



## CreateAlertConfig

> GroupAlertsConfig CreateAlertConfig(ctx, groupId, groupAlertsConfig GroupAlertsConfig).Execute()

Create One Alert Configuration in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    groupAlertsConfig := *openapiclient.NewGroupAlertsConfig() // GroupAlertsConfig | 

    resp, r, err := sdk.AlertConfigurationsApi.CreateAlertConfig(context.Background(), groupId, &groupAlertsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.CreateAlertConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateAlertConfig`: GroupAlertsConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.CreateAlertConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupAlertsConfig** | [**GroupAlertsConfig**](GroupAlertsConfig.md) | Creates one alert configuration for the specified project. | 

### Return type

[**GroupAlertsConfig**](GroupAlertsConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertConfig

> DeleteAlertConfig(ctx, groupId, alertConfigId).Execute()

Remove One Alert Configuration from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    alertConfigId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.AlertConfigurationsApi.DeleteAlertConfig(context.Background(), groupId, alertConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.DeleteAlertConfig`: %v (%v)\n", err, r)
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
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertConfigRequest struct via the builder pattern


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


## GetAlertConfig

> GroupAlertsConfig GetAlertConfig(ctx, groupId, alertConfigId).Execute()

Return One Alert Configuration from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    alertConfigId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.AlertConfigurationsApi.GetAlertConfig(context.Background(), groupId, alertConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.GetAlertConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAlertConfig`: GroupAlertsConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.GetAlertConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupAlertsConfig**](GroupAlertsConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertConfigs

> PaginatedAlertConfig GetAlertConfigs(ctx, groupId, alertId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Alert Configurations Set for One Alert


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    alertId := "alertId_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.AlertConfigurationsApi.GetAlertConfigs(context.Background(), groupId, alertId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.GetAlertConfigs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAlertConfigs`: PaginatedAlertConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.GetAlertConfigs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertId** | **string** | Unique 24-hexadecimal digit string that identifies the alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAlertConfig**](PaginatedAlertConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertConfigs

> PaginatedAlertConfig ListAlertConfigs(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Alert Configurations in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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

    resp, r, err := sdk.AlertConfigurationsApi.ListAlertConfigs(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ListAlertConfigs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListAlertConfigs`: PaginatedAlertConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ListAlertConfigs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAlertConfig**](PaginatedAlertConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMatcherFieldNames

> []string ListMatcherFieldNames(ctx).Execute()

Return All Alert Configuration Matchers Field Names


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }


    resp, r, err := sdk.AlertConfigurationsApi.ListMatcherFieldNames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ListMatcherFieldNames`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListMatcherFieldNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ListMatcherFieldNames`: %v (%v)\n", resp, r)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMatcherFieldNamesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleAlertConfig

> GroupAlertsConfig ToggleAlertConfig(ctx, groupId, alertConfigId, alertsToggle AlertsToggle).Execute()

Toggle State of One Alert Configuration in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    alertConfigId := "32b6e34b3d91647abb20e7b8" // string | 
    alertsToggle := *openapiclient.NewAlertsToggle() // AlertsToggle | 

    resp, r, err := sdk.AlertConfigurationsApi.ToggleAlertConfig(context.Background(), groupId, alertConfigId, &alertsToggle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ToggleAlertConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ToggleAlertConfig`: GroupAlertsConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ToggleAlertConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertsToggle** | [**AlertsToggle**](AlertsToggle.md) | Enables or disables the specified alert configuration in the specified project. | 

### Return type

[**GroupAlertsConfig**](GroupAlertsConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertConfig

> GroupAlertsConfig UpdateAlertConfig(ctx, groupId, alertConfigId, groupAlertsConfig GroupAlertsConfig).Execute()

Update One Alert Configuration in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312008/admin"
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
    alertConfigId := "32b6e34b3d91647abb20e7b8" // string | 
    groupAlertsConfig := *openapiclient.NewGroupAlertsConfig() // GroupAlertsConfig | 

    resp, r, err := sdk.AlertConfigurationsApi.UpdateAlertConfig(context.Background(), groupId, alertConfigId, &groupAlertsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.UpdateAlertConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateAlertConfig`: GroupAlertsConfig
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.UpdateAlertConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupAlertsConfig** | [**GroupAlertsConfig**](GroupAlertsConfig.md) | Updates one alert configuration in the specified project. | 

### Return type

[**GroupAlertsConfig**](GroupAlertsConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

