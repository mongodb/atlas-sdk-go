# \DataFederationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataFederation**](DataFederationApi.md#CreateDataFederation) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project
[**CreatePrivateEndpointId**](DataFederationApi.md#CreatePrivateEndpointId) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project
[**DeleteDataFederation**](DataFederationApi.md#DeleteDataFederation) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project
[**DeleteDataFederationLimit**](DataFederationApi.md#DeleteDataFederationLimit) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Delete One Query Limit for One Federated Database Instance
[**DeletePrivateEndpointId**](DataFederationApi.md#DeletePrivateEndpointId) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project
[**DownloadFederationQueryLogs**](DataFederationApi.md#DownloadFederationQueryLogs) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance
[**GetDataFederation**](DataFederationApi.md#GetDataFederation) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project
[**GetDataFederationLimit**](DataFederationApi.md#GetDataFederationLimit) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Return One Federated Database Instance Query Limit for One Project
[**GetPrivateEndpointId**](DataFederationApi.md#GetPrivateEndpointId) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project
[**ListDataFederation**](DataFederationApi.md#ListDataFederation) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project
[**ListDataFederationLimits**](DataFederationApi.md#ListDataFederationLimits) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits | Return All Query Limits for One Federated Database Instance
[**ListPrivateEndpointIds**](DataFederationApi.md#ListPrivateEndpointIds) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project
[**SetDataFederationLimit**](DataFederationApi.md#SetDataFederationLimit) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Configure One Query Limit for One Federated Database Instance
[**UpdateDataFederation**](DataFederationApi.md#UpdateDataFederation) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project



## CreateDataFederation

> DataLakeTenant CreateDataFederation(ctx, groupId, dataLakeTenant DataLakeTenant).SkipRoleValidation(skipRoleValidation).Execute()

Create One Federated Database Instance in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    dataLakeTenant := *openapiclient.NewDataLakeTenant() // DataLakeTenant | 
    skipRoleValidation := true // bool |  (optional) (default to false)

    resp, r, err := sdk.DataFederationApi.CreateDataFederation(context.Background(), groupId, &dataLakeTenant).SkipRoleValidation(skipRoleValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.CreateDataFederation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateDataFederation`: DataLakeTenant
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.CreateDataFederation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataLakeTenant** | [**DataLakeTenant**](DataLakeTenant.md) | Details to create one federated database instance in the specified project. | 
 **skipRoleValidation** | **bool** | Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check. | [default to false]

### Return type

[**DataLakeTenant**](DataLakeTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateEndpointId

> PaginatedPrivateNetworkEndpointIdEntry CreatePrivateEndpointId(ctx, groupId, privateNetworkEndpointIdEntry PrivateNetworkEndpointIdEntry).Execute()

Create One Federated Database Instance and Online Archive Private Endpoint for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    privateNetworkEndpointIdEntry := *openapiclient.NewPrivateNetworkEndpointIdEntry("vpce-3bf78b0ddee411ba1") // PrivateNetworkEndpointIdEntry | 

    resp, r, err := sdk.DataFederationApi.CreatePrivateEndpointId(context.Background(), groupId, &privateNetworkEndpointIdEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.CreatePrivateEndpointId`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePrivateEndpointId`: PaginatedPrivateNetworkEndpointIdEntry
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.CreatePrivateEndpointId`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkEndpointIdEntry** | [**PrivateNetworkEndpointIdEntry**](PrivateNetworkEndpointIdEntry.md) | Private endpoint for Federated Database Instances and Online Archives to add to the specified project. | 

### Return type

[**PaginatedPrivateNetworkEndpointIdEntry**](PaginatedPrivateNetworkEndpointIdEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataFederation

> DeleteDataFederation(ctx, groupId, tenantName).Execute()

Remove One Federated Database Instance from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 

    r, err := sdk.DataFederationApi.DeleteDataFederation(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.DeleteDataFederation`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Human-readable label that identifies the federated database instance to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataFederationRequest struct via the builder pattern


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


## DeleteDataFederationLimit

> DeleteDataFederationLimit(ctx, groupId, tenantName, limitName).Execute()

Delete One Query Limit for One Federated Database Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 
    limitName := "limitName_example" // string | 

    r, err := sdk.DataFederationApi.DeleteDataFederationLimit(context.Background(), groupId, tenantName, limitName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.DeleteDataFederationLimit`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Human-readable label that identifies the federated database instance to which the query limit applies. | 
**limitName** | **string** | Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataFederationLimitRequest struct via the builder pattern


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


## DeletePrivateEndpointId

> DeletePrivateEndpointId(ctx, groupId, endpointId).Execute()

Remove One Federated Database Instance and Online Archive Private Endpoint from One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    endpointId := "endpointId_example" // string | 

    r, err := sdk.DataFederationApi.DeletePrivateEndpointId(context.Background(), groupId, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.DeletePrivateEndpointId`: %v (%v)\n", err, r)
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
**endpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateEndpointIdRequest struct via the builder pattern


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


## DownloadFederationQueryLogs

> io.ReadCloser DownloadFederationQueryLogs(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()

Download Query Logs for One Federated Database Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 
    endDate := int64(1636481348) // int64 |  (optional)
    startDate := int64(1636466948) // int64 |  (optional)

    resp, r, err := sdk.DataFederationApi.DownloadFederationQueryLogs(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.DownloadFederationQueryLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadFederationQueryLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.DownloadFederationQueryLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the federated database instance for which you want to download query logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFederationQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **int64** | Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFederation

> DataLakeTenant GetDataFederation(ctx, groupId, tenantName).Execute()

Return One Federated Database Instance in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 

    resp, r, err := sdk.DataFederationApi.GetDataFederation(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.GetDataFederation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDataFederation`: DataLakeTenant
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.GetDataFederation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the Federated Database to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataLakeTenant**](DataLakeTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFederationLimit

> DataFederationTenantQueryLimit GetDataFederationLimit(ctx, groupId, tenantName, limitName).Execute()

Return One Federated Database Instance Query Limit for One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 
    limitName := "limitName_example" // string | 

    resp, r, err := sdk.DataFederationApi.GetDataFederationLimit(context.Background(), groupId, tenantName, limitName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.GetDataFederationLimit`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDataFederationLimit`: DataFederationTenantQueryLimit
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.GetDataFederationLimit`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the federated database instance to which the query limit applies. | 
**limitName** | **string** | Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFederationLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DataFederationTenantQueryLimit**](DataFederationTenantQueryLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateEndpointId

> PrivateNetworkEndpointIdEntry GetPrivateEndpointId(ctx, groupId, endpointId).Execute()

Return One Federated Database Instance and Online Archive Private Endpoint in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    endpointId := "endpointId_example" // string | 

    resp, r, err := sdk.DataFederationApi.GetPrivateEndpointId(context.Background(), groupId, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.GetPrivateEndpointId`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPrivateEndpointId`: PrivateNetworkEndpointIdEntry
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.GetPrivateEndpointId`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**endpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateEndpointIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateNetworkEndpointIdEntry**](PrivateNetworkEndpointIdEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataFederation

> []DataLakeTenant ListDataFederation(ctx, groupId).Type_(type_).Execute()

Return All Federated Database Instances in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    type_ := "type__example" // string |  (optional) (default to "USER")

    resp, r, err := sdk.DataFederationApi.ListDataFederation(context.Background(), groupId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.ListDataFederation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDataFederation`: []DataLakeTenant
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.ListDataFederation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of Federated Database Instances to return. | [default to &quot;USER&quot;]

### Return type

[**[]DataLakeTenant**](DataLakeTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataFederationLimits

> []DataFederationTenantQueryLimit ListDataFederationLimits(ctx, groupId, tenantName).Execute()

Return All Query Limits for One Federated Database Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 

    resp, r, err := sdk.DataFederationApi.ListDataFederationLimits(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.ListDataFederationLimits`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListDataFederationLimits`: []DataFederationTenantQueryLimit
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.ListDataFederationLimits`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the federated database instance for which you want to retrieve query limits. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataFederationLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DataFederationTenantQueryLimit**](DataFederationTenantQueryLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateEndpointIds

> PaginatedPrivateNetworkEndpointIdEntry ListPrivateEndpointIds(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Federated Database Instance and Online Archive Private Endpoints in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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

    resp, r, err := sdk.DataFederationApi.ListPrivateEndpointIds(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.ListPrivateEndpointIds`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPrivateEndpointIds`: PaginatedPrivateNetworkEndpointIdEntry
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.ListPrivateEndpointIds`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedPrivateNetworkEndpointIdEntry**](PaginatedPrivateNetworkEndpointIdEntry.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDataFederationLimit

> DataFederationTenantQueryLimit SetDataFederationLimit(ctx, groupId, tenantName, limitName, dataFederationTenantQueryLimit DataFederationTenantQueryLimit).Execute()

Configure One Query Limit for One Federated Database Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 
    limitName := "limitName_example" // string | 
    dataFederationTenantQueryLimit := *openapiclient.NewDataFederationTenantQueryLimit("Name_example", int64(123)) // DataFederationTenantQueryLimit | 

    resp, r, err := sdk.DataFederationApi.SetDataFederationLimit(context.Background(), groupId, tenantName, limitName, &dataFederationTenantQueryLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.SetDataFederationLimit`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `SetDataFederationLimit`: DataFederationTenantQueryLimit
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.SetDataFederationLimit`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the federated database instance to which the query limit applies. | 
**limitName** | **string** | Human-readable label that identifies this data federation instance limit.  | Limit Name | Description | Default | | --- | --- | --- | | bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A | | bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A | | bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A | | bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDataFederationLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dataFederationTenantQueryLimit** | [**DataFederationTenantQueryLimit**](DataFederationTenantQueryLimit.md) | Creates or updates one query limit for one federated database instance. | 

### Return type

[**DataFederationTenantQueryLimit**](DataFederationTenantQueryLimit.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataFederation

> DataLakeTenant UpdateDataFederation(ctx, groupId, tenantName, dataLakeTenant DataLakeTenant).SkipRoleValidation(skipRoleValidation).Execute()

Update One Federated Database Instance in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312006/admin"
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
    tenantName := "tenantName_example" // string | 
    skipRoleValidation := true // bool | 
    dataLakeTenant := *openapiclient.NewDataLakeTenant() // DataLakeTenant | 

    resp, r, err := sdk.DataFederationApi.UpdateDataFederation(context.Background(), groupId, tenantName, &dataLakeTenant).SkipRoleValidation(skipRoleValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.UpdateDataFederation`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateDataFederation`: DataLakeTenant
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.UpdateDataFederation`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the federated database instance to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipRoleValidation** | **bool** | Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check. | 
 **dataLakeTenant** | [**DataLakeTenant**](DataLakeTenant.md) | Details of one Federated Database to update in the specified project. | 

### Return type

[**DataLakeTenant**](DataLakeTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

