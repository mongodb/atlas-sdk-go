# \StreamsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamConnection**](StreamsApi.md#CreateStreamConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Create One Connection
[**CreateStreamInstance**](StreamsApi.md#CreateStreamInstance) | **Post** /api/atlas/v2/groups/{groupId}/streams | Create One Stream Instance
[**CreateStreamInstanceWithSampleConnections**](StreamsApi.md#CreateStreamInstanceWithSampleConnections) | **Post** /api/atlas/v2/groups/{groupId}/streams:withSampleConnections | Create One Stream Instance With Sample Connections
[**CreateStreamProcessor**](StreamsApi.md#CreateStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor | Create One Stream Processor
[**DeleteStreamConnection**](StreamsApi.md#DeleteStreamConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Delete One Stream Connection
[**DeleteStreamInstance**](StreamsApi.md#DeleteStreamInstance) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Delete One Stream Instance
[**DeleteStreamProcessor**](StreamsApi.md#DeleteStreamProcessor) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Delete One Stream Processor
[**DownloadStreamTenantAuditLogs**](StreamsApi.md#DownloadStreamTenantAuditLogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs | Download Audit Logs for One Atlas Stream Processing Instance
[**GetStreamConnection**](StreamsApi.md#GetStreamConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Return One Stream Connection
[**GetStreamInstance**](StreamsApi.md#GetStreamInstance) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Return One Stream Instance
[**GetStreamProcessor**](StreamsApi.md#GetStreamProcessor) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Get One Stream Processor
[**ListStreamConnections**](StreamsApi.md#ListStreamConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Return All Connections Of The Stream Instances
[**ListStreamInstances**](StreamsApi.md#ListStreamInstances) | **Get** /api/atlas/v2/groups/{groupId}/streams | Return All Project Stream Instances
[**ListStreamProcessors**](StreamsApi.md#ListStreamProcessors) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processors | Return All Stream Processors In The Stream Instance.
[**StartStreamProcessor**](StreamsApi.md#StartStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start | Start One Stream Processor
[**StopStreamProcessor**](StreamsApi.md#StopStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop | Stop One Stream Processor
[**UpdateStreamConnection**](StreamsApi.md#UpdateStreamConnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Update One Stream Connection
[**UpdateStreamInstance**](StreamsApi.md#UpdateStreamInstance) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Update One Stream Instance



## CreateStreamConnection

> StreamsConnection CreateStreamConnection(ctx, groupId, tenantName, streamsConnection StreamsConnection).Execute()

Create One Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    streamsConnection := *openapiclient.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsApi.CreateStreamConnection(context.Background(), groupId, tenantName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to create one connection for a streams instance in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamInstance

> StreamsTenant CreateStreamInstance(ctx, groupId, streamsTenant StreamsTenant).Execute()

Create One Stream Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    streamsTenant := *openapiclient.NewStreamsTenant() // StreamsTenant | 

    resp, r, err := sdk.StreamsApi.CreateStreamInstance(context.Background(), groupId, &streamsTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamInstance`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamInstance`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamsTenant** | [**StreamsTenant**](StreamsTenant.md) | Details to create one streams instance in the specified project. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamInstanceWithSampleConnections

> StreamsTenant CreateStreamInstanceWithSampleConnections(ctx, groupId, body any).Execute()

Create One Stream Instance With Sample Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    body := any{ ... } // any | 

    resp, r, err := sdk.StreamsApi.CreateStreamInstanceWithSampleConnections(context.Background(), groupId, &body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamInstanceWithSampleConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamInstanceWithSampleConnections`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamInstanceWithSampleConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamInstanceWithSampleConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **any** | Details to create one streams instance in the specified project. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamProcessor

> StreamsProcessor CreateStreamProcessor(ctx, groupId, tenantName, streamsProcessor StreamsProcessor).Execute()

Create One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    streamsProcessor := *openapiclient.NewStreamsProcessor() // StreamsProcessor | 

    resp, r, err := sdk.StreamsApi.CreateStreamProcessor(context.Background(), groupId, tenantName, &streamsProcessor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamProcessor`: StreamsProcessor
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamsProcessor** | [**StreamsProcessor**](StreamsProcessor.md) | Details to create an Atlas Streams Processor. | 

### Return type

[**StreamsProcessor**](StreamsProcessor.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamConnection

> any DeleteStreamConnection(ctx, groupId, tenantName, connectionName).Execute()

Delete One Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    connectionName := "connectionName_example" // string | 

    resp, r, err := sdk.StreamsApi.DeleteStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteStreamConnection`: any
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DeleteStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**connectionName** | **string** | Human-readable label that identifies the stream connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamInstance

> any DeleteStreamInstance(ctx, groupId, tenantName).Execute()

Delete One Stream Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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

    resp, r, err := sdk.StreamsApi.DeleteStreamInstance(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamInstance`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteStreamInstance`: any
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DeleteStreamInstance`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamProcessor

> DeleteStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Delete One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    processorName := "processorName_example" // string | 

    r, err := sdk.StreamsApi.DeleteStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamProcessor`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**processorName** | **string** | Human-readable label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadStreamTenantAuditLogs

> io.ReadCloser DownloadStreamTenantAuditLogs(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()

Download Audit Logs for One Atlas Stream Processing Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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

    resp, r, err := sdk.StreamsApi.DownloadStreamTenantAuditLogs(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DownloadStreamTenantAuditLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadStreamTenantAuditLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DownloadStreamTenantAuditLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadStreamTenantAuditLogsRequest struct via the builder pattern


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
- **Accept**: application/vnd.atlas.2023-02-01+gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamConnection

> StreamsConnection GetStreamConnection(ctx, groupId, tenantName, connectionName).Execute()

Return One Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    connectionName := "connectionName_example" // string | 

    resp, r, err := sdk.StreamsApi.GetStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance to return. | 
**connectionName** | **string** | Human-readable label that identifies the stream connection to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamInstance

> StreamsTenant GetStreamInstance(ctx, groupId, tenantName).IncludeConnections(includeConnections).Execute()

Return One Stream Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    includeConnections := true // bool |  (optional)

    resp, r, err := sdk.StreamsApi.GetStreamInstance(context.Background(), groupId, tenantName).IncludeConnections(includeConnections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamInstance`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamInstance`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeConnections** | **bool** | Flag to indicate whether connections information should be included in the stream instance. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamProcessor

> StreamsProcessorWithStats GetStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Get One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    processorName := "processorName_example" // string | 

    resp, r, err := sdk.StreamsApi.GetStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamProcessor`: StreamsProcessorWithStats
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**processorName** | **string** | Human-readable label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsProcessorWithStats**](StreamsProcessorWithStats.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamConnections

> PaginatedApiStreamsConnection ListStreamConnections(ctx, groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Connections Of The Stream Instances


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsApi.ListStreamConnections(context.Background(), groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListStreamConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListStreamConnections`: PaginatedApiStreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListStreamConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsConnection**](PaginatedApiStreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamInstances

> PaginatedApiStreamsTenant ListStreamInstances(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Project Stream Instances


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsApi.ListStreamInstances(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListStreamInstances`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListStreamInstances`: PaginatedApiStreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListStreamInstances`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsTenant**](PaginatedApiStreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamProcessors

> PaginatedApiStreamsStreamProcessorWithStats ListStreamProcessors(ctx, groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).IncludeCount(includeCount).Execute()

Return All Stream Processors In The Stream Instance.


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    includeCount := true // bool |  (optional) (default to true)

    resp, r, err := sdk.StreamsApi.ListStreamProcessors(context.Background(), groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).IncludeCount(includeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListStreamProcessors`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListStreamProcessors`: PaginatedApiStreamsStreamProcessorWithStats
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListStreamProcessors`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]

### Return type

[**PaginatedApiStreamsStreamProcessorWithStats**](PaginatedApiStreamsStreamProcessorWithStats.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartStreamProcessor

> any StartStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Start One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    processorName := "processorName_example" // string | 

    resp, r, err := sdk.StreamsApi.StartStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.StartStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `StartStreamProcessor`: any
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.StartStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**processorName** | **string** | Human-readable label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopStreamProcessor

> any StopStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Stop One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    processorName := "processorName_example" // string | 

    resp, r, err := sdk.StreamsApi.StopStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.StopStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `StopStreamProcessor`: any
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.StopStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**processorName** | **string** | Human-readable label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamConnection

> StreamsConnection UpdateStreamConnection(ctx, groupId, tenantName, connectionName, streamsConnection StreamsConnection).Execute()

Update One Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    connectionName := "connectionName_example" // string | 
    streamsConnection := *openapiclient.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsApi.UpdateStreamConnection(context.Background(), groupId, tenantName, connectionName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**connectionName** | **string** | Human-readable label that identifies the stream connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to update one connection for a streams instance in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamInstance

> StreamsTenant UpdateStreamInstance(ctx, groupId, tenantName, streamsDataProcessRegion StreamsDataProcessRegion).Execute()

Update One Stream Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805004/admin"
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
    streamsDataProcessRegion := *openapiclient.NewStreamsDataProcessRegion("CloudProvider_example", "Region_example") // StreamsDataProcessRegion | 

    resp, r, err := sdk.StreamsApi.UpdateStreamInstance(context.Background(), groupId, tenantName, &streamsDataProcessRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamInstance`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamInstance`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Human-readable label that identifies the stream instance to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamsDataProcessRegion** | [**StreamsDataProcessRegion**](StreamsDataProcessRegion.md) | Details of the new data process region to update in the streams instance. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

