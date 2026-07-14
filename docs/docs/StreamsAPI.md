# \StreamsAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptVpcPeeringConnection**](StreamsAPI.md#AcceptVpcPeeringConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept | Accept One Incoming VPC Peering Connection
[**CreateFailoverConnection**](StreamsAPI.md#CreateFailoverConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections | Create One Failover Stream Connection
[**CreatePrivateLinkConnection**](StreamsAPI.md#CreatePrivateLinkConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Create One Private Link Connection
[**CreateStreamConnection**](StreamsAPI.md#CreateStreamConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Create One Stream Connection
[**CreateStreamProcessor**](StreamsAPI.md#CreateStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor | Create One Stream Processor
[**CreateStreamWorkspace**](StreamsAPI.md#CreateStreamWorkspace) | **Post** /api/atlas/v2/groups/{groupId}/streams | Create One Stream Workspace
[**DeletePrivateLinkConnection**](StreamsAPI.md#DeletePrivateLinkConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Delete One Private Link Connection
[**DeleteStreamConnection**](StreamsAPI.md#DeleteStreamConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Delete One Stream Connection
[**DeleteStreamFailoverConnection**](StreamsAPI.md#DeleteStreamFailoverConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Delete One Stream Failover Connection
[**DeleteStreamProcessor**](StreamsAPI.md#DeleteStreamProcessor) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Delete One Stream Processor
[**DeleteStreamWorkspace**](StreamsAPI.md#DeleteStreamWorkspace) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Delete One Stream Workspace
[**DeleteVpcPeeringConnection**](StreamsAPI.md#DeleteVpcPeeringConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id} | Delete One VPC Peering Connection
[**DownloadAuditLogs**](StreamsAPI.md#DownloadAuditLogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs | Download Audit Logs for One Atlas Stream Processing Workspace
[**DownloadOperationalLogs**](StreamsAPI.md#DownloadOperationalLogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}:downloadOperationalLogs | Download Operational Logs for One Atlas Stream Processing Workspace
[**GetAccountDetails**](StreamsAPI.md#GetAccountDetails) | **Get** /api/atlas/v2/groups/{groupId}/streams/accountDetails | Return Account ID and VPC ID for One Project and Region
[**GetPrivateLinkConnection**](StreamsAPI.md#GetPrivateLinkConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Return One Private Link Connection
[**GetStreamConnection**](StreamsAPI.md#GetStreamConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Return One Stream Connection
[**GetStreamFailoverConnection**](StreamsAPI.md#GetStreamFailoverConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Return One Stream Failover Connection
[**GetStreamProcessor**](StreamsAPI.md#GetStreamProcessor) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Return One Stream Processor
[**GetStreamProcessors**](StreamsAPI.md#GetStreamProcessors) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processors | Return All Stream Processors in One Stream Workspace
[**GetStreamWorkspace**](StreamsAPI.md#GetStreamWorkspace) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Return One Stream Workspace
[**ListActivePeeringConnections**](StreamsAPI.md#ListActivePeeringConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/activeVpcPeeringConnections | Return All Active Incoming VPC Peering Connections
[**ListFailoverConnections**](StreamsAPI.md#ListFailoverConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections | Return All Stream Failover Connections
[**ListPrivateLinkConnections**](StreamsAPI.md#ListPrivateLinkConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Return All Private Link Connections
[**ListStreamConnections**](StreamsAPI.md#ListStreamConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Return All Connections of the Stream Workspaces
[**ListStreamWorkspaces**](StreamsAPI.md#ListStreamWorkspaces) | **Get** /api/atlas/v2/groups/{groupId}/streams | Return All Stream Workspaces in One Project
[**ListVpcPeeringConnections**](StreamsAPI.md#ListVpcPeeringConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections | Return All VPC Peering Connections
[**RejectVpcPeeringConnection**](StreamsAPI.md#RejectVpcPeeringConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject | Reject One Incoming VPC Peering Connection
[**StartStreamProcessor**](StreamsAPI.md#StartStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:start | Start One Stream Processor
[**StartStreamProcessorWith**](StreamsAPI.md#StartStreamProcessorWith) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:startWith | Start One Stream Processor With Options
[**StopStreamProcessor**](StreamsAPI.md#StopStreamProcessor) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName}:stop | Stop One Stream Processor
[**UpdateStreamConnection**](StreamsAPI.md#UpdateStreamConnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Update One Stream Connection
[**UpdateStreamFailoverConnection**](StreamsAPI.md#UpdateStreamFailoverConnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}/failoverConnections/{failoverConnectionId} | Update One Stream Failover Connection
[**UpdateStreamProcessor**](StreamsAPI.md#UpdateStreamProcessor) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/processor/{processorName} | Update One Stream Processor
[**UpdateStreamWorkspace**](StreamsAPI.md#UpdateStreamWorkspace) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Update One Stream Workspace
[**WithStreamSampleConnections**](StreamsAPI.md#WithStreamSampleConnections) | **Post** /api/atlas/v2/groups/{groupId}/streams:withSampleConnections | Create One Stream Workspace with Sample Connections



## AcceptVpcPeeringConnection

> AcceptVpcPeeringConnection(ctx, groupId, id, vPCPeeringActionChallenge VPCPeeringActionChallenge).Execute()

Accept One Incoming VPC Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    id := "id_example" // string | 
    vPCPeeringActionChallenge := *admin.NewVPCPeeringActionChallenge() // VPCPeeringActionChallenge | 

    r, err := sdk.StreamsAPI.AcceptVpcPeeringConnection(context.Background(), groupId, id, &vPCPeeringActionChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.AcceptVpcPeeringConnection`: %v (%v)\n", err, r)
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
**id** | **string** | The VPC Peering Connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptVpcPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vPCPeeringActionChallenge** | [**VPCPeeringActionChallenge**](VPCPeeringActionChallenge.md) | Challenge values for VPC Peering requester account ID, and requester VPC ID. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFailoverConnection

> StreamsConnection CreateFailoverConnection(ctx, groupId, tenantName, connectionName, streamsConnection StreamsConnection).Execute()

Create One Failover Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsConnection := *admin.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsAPI.CreateFailoverConnection(context.Background(), groupId, tenantName, connectionName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreateFailoverConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateFailoverConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreateFailoverConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFailoverConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to create one failover connection for a streams workspace in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateLinkConnection

> StreamsPrivateLinkConnection CreatePrivateLinkConnection(ctx, groupId, streamsPrivateLinkConnection StreamsPrivateLinkConnection).Execute()

Create One Private Link Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsPrivateLinkConnection := *admin.NewStreamsPrivateLinkConnection("Provider_example") // StreamsPrivateLinkConnection | 

    resp, r, err := sdk.StreamsAPI.CreatePrivateLinkConnection(context.Background(), groupId, &streamsPrivateLinkConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreatePrivateLinkConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreatePrivateLinkConnection`: StreamsPrivateLinkConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreatePrivateLinkConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateLinkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamsPrivateLinkConnection** | [**StreamsPrivateLinkConnection**](StreamsPrivateLinkConnection.md) | Details to create one Private Link connection for a project. project. | 

### Return type

[**StreamsPrivateLinkConnection**](StreamsPrivateLinkConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamConnection

> StreamsConnection CreateStreamConnection(ctx, groupId, tenantName, streamsConnection StreamsConnection).Execute()

Create One Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsConnection := *admin.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsAPI.CreateStreamConnection(context.Background(), groupId, tenantName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreateStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreateStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to create one connection for a streams workspace in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

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

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsProcessor := *admin.NewStreamsProcessor() // StreamsProcessor | 

    resp, r, err := sdk.StreamsAPI.CreateStreamProcessor(context.Background(), groupId, tenantName, &streamsProcessor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreateStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamProcessor`: StreamsProcessor
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreateStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

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
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamWorkspace

> StreamsTenant CreateStreamWorkspace(ctx, groupId, streamsTenant StreamsTenant).Execute()

Create One Stream Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsTenant := *admin.NewStreamsTenant() // StreamsTenant | 

    resp, r, err := sdk.StreamsAPI.CreateStreamWorkspace(context.Background(), groupId, &streamsTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreateStreamWorkspace`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateStreamWorkspace`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreateStreamWorkspace`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamsTenant** | [**StreamsTenant**](StreamsTenant.md) | Details to create one streams workspace in the specified project. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateLinkConnection

> DeletePrivateLinkConnection(ctx, groupId, connectionId).Execute()

Delete One Private Link Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    connectionId := "connectionId_example" // string | 

    r, err := sdk.StreamsAPI.DeletePrivateLinkConnection(context.Background(), groupId, connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeletePrivateLinkConnection`: %v (%v)\n", err, r)
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
**connectionId** | **string** | Unique ID that identifies the Private Link connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateLinkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamConnection

> DeleteStreamConnection(ctx, groupId, tenantName, connectionName).Execute()

Delete One Stream Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    r, err := sdk.StreamsAPI.DeleteStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStreamConnection`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamFailoverConnection

> DeleteStreamFailoverConnection(ctx, groupId, tenantName, connectionName, failoverConnectionId).Execute()

Delete One Stream Failover Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    failoverConnectionId := "failoverConnectionId_example" // string | 

    r, err := sdk.StreamsAPI.DeleteStreamFailoverConnection(context.Background(), groupId, tenantName, connectionName, failoverConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStreamFailoverConnection`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 
**failoverConnectionId** | **string** | Label that identifies the stream failover connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamFailoverConnectionRequest struct via the builder pattern


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

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    r, err := sdk.StreamsAPI.DeleteStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStreamProcessor`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

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
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamWorkspace

> DeleteStreamWorkspace(ctx, groupId, tenantName).Execute()

Delete One Stream Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    r, err := sdk.StreamsAPI.DeleteStreamWorkspace(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStreamWorkspace`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcPeeringConnection

> DeleteVpcPeeringConnection(ctx, groupId, id).Execute()

Delete One VPC Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    id := "id_example" // string | 

    r, err := sdk.StreamsAPI.DeleteVpcPeeringConnection(context.Background(), groupId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteVpcPeeringConnection`: %v (%v)\n", err, r)
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
**id** | **string** | The VPC Peering Connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAuditLogs

> io.ReadCloser DownloadAuditLogs(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).SpName(spName).Execute()

Download Audit Logs for One Atlas Stream Processing Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    spName := "spName_example" // string |  (optional)

    resp, r, err := sdk.StreamsAPI.DownloadAuditLogs(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).SpName(spName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DownloadAuditLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadAuditLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.DownloadAuditLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **int64** | Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **spName** | **string** | Name of the stream processor to download logs for. An empty string will download logs for all stream processors in the workspace. | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOperationalLogs

> io.ReadCloser DownloadOperationalLogs(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).SpName(spName).Execute()

Download Operational Logs for One Atlas Stream Processing Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    spName := "spName_example" // string |  (optional)

    resp, r, err := sdk.StreamsAPI.DownloadOperationalLogs(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).SpName(spName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DownloadOperationalLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadOperationalLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.DownloadOperationalLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOperationalLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **int64** | Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **spName** | **string** | Name of the stream processor to download logs for. An empty string will download logs for all stream processors in the workspace. | 

### Return type

[**io.ReadCloser**](io.ReadCloser.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDetails

> AccountDetails GetAccountDetails(ctx, groupId).CloudProvider(cloudProvider).RegionName(regionName).Execute()

Return Account ID and VPC ID for One Project and Region


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    regionName := "regionName_example" // string | 

    resp, r, err := sdk.StreamsAPI.GetAccountDetails(context.Background(), groupId).CloudProvider(cloudProvider).RegionName(regionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetAccountDetails`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetAccountDetails`: AccountDetails
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetAccountDetails`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProvider** | **string** | One of \&quot;aws\&quot;, \&quot;azure\&quot; or \&quot;gcp\&quot;. | 
 **regionName** | **string** | The cloud provider specific region name, i.e. \&quot;US_EAST_1\&quot; for cloud provider \&quot;aws\&quot;. | 

### Return type

[**AccountDetails**](AccountDetails.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateLinkConnection

> StreamsPrivateLinkConnection GetPrivateLinkConnection(ctx, groupId, connectionId).Execute()

Return One Private Link Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    connectionId := "connectionId_example" // string | 

    resp, r, err := sdk.StreamsAPI.GetPrivateLinkConnection(context.Background(), groupId, connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetPrivateLinkConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetPrivateLinkConnection`: StreamsPrivateLinkConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetPrivateLinkConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**connectionId** | **string** | Unique ID that identifies the Private Link connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateLinkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamsPrivateLinkConnection**](StreamsPrivateLinkConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

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

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.StreamsAPI.GetStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace to return. | 
**connectionName** | **string** | Label that identifies the stream connection to return. | 

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
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamFailoverConnection

> StreamsConnection GetStreamFailoverConnection(ctx, groupId, tenantName, connectionName, failoverConnectionId).Execute()

Return One Stream Failover Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    failoverConnectionId := "failoverConnectionId_example" // string | 

    resp, r, err := sdk.StreamsAPI.GetStreamFailoverConnection(context.Background(), groupId, tenantName, connectionName, failoverConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamFailoverConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamFailoverConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamFailoverConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 
**failoverConnectionId** | **string** | Label that identifies the stream failover connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamFailoverConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamProcessor

> StreamsProcessorWithStats GetStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Return One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.StreamsAPI.GetStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamProcessor`: StreamsProcessorWithStats
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

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
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamProcessors

> PaginatedApiStreamsStreamProcessorWithStats GetStreamProcessors(ctx, groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).IncludeCount(includeCount).Execute()

Return All Stream Processors in One Stream Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    includeCount := true // bool |  (optional) (default to true)

    resp, r, err := sdk.StreamsAPI.GetStreamProcessors(context.Background(), groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).IncludeCount(includeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamProcessors`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamProcessors`: PaginatedApiStreamsStreamProcessorWithStats
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamProcessors`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]

### Return type

[**PaginatedApiStreamsStreamProcessorWithStats**](PaginatedApiStreamsStreamProcessorWithStats.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamWorkspace

> StreamsTenant GetStreamWorkspace(ctx, groupId, tenantName).IncludeConnections(includeConnections).Execute()

Return One Stream Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.StreamsAPI.GetStreamWorkspace(context.Background(), groupId, tenantName).IncludeConnections(includeConnections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamWorkspace`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetStreamWorkspace`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamWorkspace`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeConnections** | **bool** | Flag to indicate whether connections information should be included in the stream workspace. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivePeeringConnections

> PaginatedApiStreamsVPCPeeringConnection ListActivePeeringConnections(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Active Incoming VPC Peering Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListActivePeeringConnections(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListActivePeeringConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListActivePeeringConnections`: PaginatedApiStreamsVPCPeeringConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListActivePeeringConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActivePeeringConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsVPCPeeringConnection**](PaginatedApiStreamsVPCPeeringConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFailoverConnections

> PaginatedApiStreamsFailoverConnection ListFailoverConnections(ctx, groupId, tenantName, connectionName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Stream Failover Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListFailoverConnections(context.Background(), groupId, tenantName, connectionName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListFailoverConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListFailoverConnections`: PaginatedApiStreamsFailoverConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListFailoverConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFailoverConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsFailoverConnection**](PaginatedApiStreamsFailoverConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateLinkConnections

> PaginatedApiStreamsPrivateLink ListPrivateLinkConnections(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Private Link Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListPrivateLinkConnections(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListPrivateLinkConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPrivateLinkConnections`: PaginatedApiStreamsPrivateLink
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListPrivateLinkConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateLinkConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsPrivateLink**](PaginatedApiStreamsPrivateLink.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamConnections

> PaginatedApiStreamsConnection ListStreamConnections(ctx, groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Connections of the Stream Workspaces


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListStreamConnections(context.Background(), groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListStreamConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListStreamConnections`: PaginatedApiStreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListStreamConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 

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
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamWorkspaces

> PaginatedApiStreamsTenant ListStreamWorkspaces(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Stream Workspaces in One Project


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListStreamWorkspaces(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListStreamWorkspaces`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListStreamWorkspaces`: PaginatedApiStreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListStreamWorkspaces`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamWorkspacesRequest struct via the builder pattern


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
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcPeeringConnections

> PaginatedApiStreamsVPCPeeringConnection ListVpcPeeringConnections(ctx, groupId).RequesterAccountId(requesterAccountId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All VPC Peering Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    requesterAccountId := "requesterAccountId_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsAPI.ListVpcPeeringConnections(context.Background(), groupId).RequesterAccountId(requesterAccountId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.ListVpcPeeringConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListVpcPeeringConnections`: PaginatedApiStreamsVPCPeeringConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.ListVpcPeeringConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeeringConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requesterAccountId** | **string** | The Account ID of the VPC Peering connection/s. | 

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiStreamsVPCPeeringConnection**](PaginatedApiStreamsVPCPeeringConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectVpcPeeringConnection

> RejectVpcPeeringConnection(ctx, groupId, id).Execute()

Reject One Incoming VPC Peering Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    id := "id_example" // string | 

    r, err := sdk.StreamsAPI.RejectVpcPeeringConnection(context.Background(), groupId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.RejectVpcPeeringConnection`: %v (%v)\n", err, r)
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
**id** | **string** | The VPC Peering Connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectVpcPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartStreamProcessor

> StartStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Start One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    r, err := sdk.StreamsAPI.StartStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StartStreamProcessor`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartStreamProcessorWith

> StartStreamProcessorWith(ctx, groupId, tenantName, processorName, streamsStartStreamProcessorWith StreamsStartStreamProcessorWith).Execute()

Start One Stream Processor With Options


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsStartStreamProcessorWith := *admin.NewStreamsStartStreamProcessorWith() // StreamsStartStreamProcessorWith |  (optional)

    r, err := sdk.StreamsAPI.StartStreamProcessorWith(context.Background(), groupId, tenantName, processorName, &streamsStartStreamProcessorWith).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StartStreamProcessorWith`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartStreamProcessorWithRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamsStartStreamProcessorWith** | [**StreamsStartStreamProcessorWith**](StreamsStartStreamProcessorWith.md) | Options for starting a stream processor. | 

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopStreamProcessor

> StopStreamProcessor(ctx, groupId, tenantName, processorName).Execute()

Stop One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    r, err := sdk.StreamsAPI.StopStreamProcessor(context.Background(), groupId, tenantName, processorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StopStreamProcessor`: %v (%v)\n", err, r)
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
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

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

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsConnection := *admin.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsAPI.UpdateStreamConnection(context.Background(), groupId, tenantName, connectionName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStreamConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStreamConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to update one connection for a streams workspace in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamFailoverConnection

> StreamsConnection UpdateStreamFailoverConnection(ctx, groupId, tenantName, connectionName, failoverConnectionId, streamsConnection StreamsConnection).Execute()

Update One Stream Failover Connection


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    failoverConnectionId := "failoverConnectionId_example" // string | 
    streamsConnection := *admin.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsAPI.UpdateStreamFailoverConnection(context.Background(), groupId, tenantName, connectionName, failoverConnectionId, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStreamFailoverConnection`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamFailoverConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStreamFailoverConnection`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**connectionName** | **string** | Label that identifies the stream connection. | 
**failoverConnectionId** | **string** | Label that identifies the stream failover connection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamFailoverConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **streamsConnection** | [**StreamsConnection**](StreamsConnection.md) | Details to update one failover connection for a streams workspace in the specified project. | 

### Return type

[**StreamsConnection**](StreamsConnection.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamProcessor

> StreamsProcessorWithStats UpdateStreamProcessor(ctx, groupId, tenantName, processorName, streamsModifyStreamProcessor StreamsModifyStreamProcessor).Execute()

Update One Stream Processor


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsModifyStreamProcessor := *admin.NewStreamsModifyStreamProcessor() // StreamsModifyStreamProcessor | 

    resp, r, err := sdk.StreamsAPI.UpdateStreamProcessor(context.Background(), groupId, tenantName, processorName, &streamsModifyStreamProcessor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStreamProcessor`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamProcessor`: StreamsProcessorWithStats
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStreamProcessor`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace. | 
**processorName** | **string** | Label that identifies the stream processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamsModifyStreamProcessor** | [**StreamsModifyStreamProcessor**](StreamsModifyStreamProcessor.md) | Modifications to apply to the stream processor. | 

### Return type

[**StreamsProcessorWithStats**](StreamsProcessorWithStats.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamWorkspace

> StreamsTenant UpdateStreamWorkspace(ctx, groupId, tenantName, streamsTenantUpdateRequest StreamsTenantUpdateRequest).Execute()

Update One Stream Workspace


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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
    streamsTenantUpdateRequest := *admin.NewStreamsTenantUpdateRequest() // StreamsTenantUpdateRequest | 

    resp, r, err := sdk.StreamsAPI.UpdateStreamWorkspace(context.Background(), groupId, tenantName, &streamsTenantUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStreamWorkspace`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateStreamWorkspace`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStreamWorkspace`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**tenantName** | **string** | Label that identifies the stream workspace to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamsTenantUpdateRequest** | [**StreamsTenantUpdateRequest**](StreamsTenantUpdateRequest.md) | Details to update in the streams workspace. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-02-01+json
- **Accept**: application/vnd.atlas.2023-02-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithStreamSampleConnections

> StreamsTenant WithStreamSampleConnections(ctx, groupId, body any).Execute()

Create One Stream Workspace with Sample Connections


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312021/admin"
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

    resp, r, err := sdk.StreamsAPI.WithStreamSampleConnections(context.Background(), groupId, &body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.WithStreamSampleConnections`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `WithStreamSampleConnections`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.WithStreamSampleConnections`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithStreamSampleConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **any** | Details to create one streams workspace in the specified project. | 

### Return type

[**StreamsTenant**](StreamsTenant.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

