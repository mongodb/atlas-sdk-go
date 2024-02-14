# \AlertsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeAlert**](AlertsApi.md#AcknowledgeAlert) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project
[**GetAlert**](AlertsApi.md#GetAlert) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project
[**ListAlerts**](AlertsApi.md#ListAlerts) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project
[**ListAlertsByAlertConfigurationId**](AlertsApi.md#ListAlertsByAlertConfigurationId) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for Alert Configuration



## AcknowledgeAlert

> AlertViewForNdsGroup AcknowledgeAlert(ctx, groupId, alertId, alertViewForNdsGroup AlertViewForNdsGroup).Execute()

Acknowledge One Alert from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    alertId := "alertId_example" // string | 
    alertViewForNdsGroup := *openapiclient.NewAlertViewForNdsGroup() // AlertViewForNdsGroup | 

    resp, r, err := sdk.AlertsApi.AcknowledgeAlert(context.Background(), groupId, alertId, &alertViewForNdsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AcknowledgeAlert``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `AcknowledgeAlert`: AlertViewForNdsGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AcknowledgeAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertId** | **string** | Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertViewForNdsGroup** | [**AlertViewForNdsGroup**](AlertViewForNdsGroup.md) | Confirm one alert. | 

### Return type

[**AlertViewForNdsGroup**](AlertViewForNdsGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert

> AlertViewForNdsGroup GetAlert(ctx, groupId, alertId).Execute()

Return One Alert from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    alertId := "alertId_example" // string | 

    resp, r, err := sdk.AlertsApi.GetAlert(context.Background(), groupId, alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlert``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetAlert`: AlertViewForNdsGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertId** | **string** | Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertViewForNdsGroup**](AlertViewForNdsGroup.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> PaginatedAlert ListAlerts(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Status(status).Execute()

Return All Alerts from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    status := "status_example" // string |  (optional)

    resp, r, err := sdk.AlertsApi.ListAlerts(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlerts``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListAlerts`: PaginatedAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **status** | **string** | Status of the alerts to return. Omit to return all alerts in all statuses. | 

### Return type

[**PaginatedAlert**](PaginatedAlert.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertsByAlertConfigurationId

> PaginatedAlert ListAlertsByAlertConfigurationId(ctx, groupId, alertConfigId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Open Alerts for Alert Configuration


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

    "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    alertConfigId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.AlertsApi.ListAlertsByAlertConfigurationId(context.Background(), groupId, alertConfigId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertsByAlertConfigurationId``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListAlertsByAlertConfigurationId`: PaginatedAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertsByAlertConfigurationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsByAlertConfigurationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedAlert**](PaginatedAlert.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

