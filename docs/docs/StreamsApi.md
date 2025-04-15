# \StreamsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptVpcPeeringConnection**](StreamsApi.md#AcceptVpcPeeringConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:accept | Requests the acceptance of an incoming VPC Peering connection.
[**CreatePrivateLinkConnection**](StreamsApi.md#CreatePrivateLinkConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Create One Private Link connection
[**CreateStreamConnection**](StreamsApi.md#CreateStreamConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Create One Connection
[**CreateStreamInstance**](StreamsApi.md#CreateStreamInstance) | **Post** /api/atlas/v2/groups/{groupId}/streams | Create One Stream Instance
[**DeletePrivateLinkConnection**](StreamsApi.md#DeletePrivateLinkConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Delete One Private Link connection
[**DeleteStreamConnection**](StreamsApi.md#DeleteStreamConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Delete One Stream Connection
[**DeleteStreamInstance**](StreamsApi.md#DeleteStreamInstance) | **Delete** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Delete One Stream Instance
[**DeleteVpcPeeringConnection**](StreamsApi.md#DeleteVpcPeeringConnection) | **Delete** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id} | Deletes an incoming VPC Peering connection.
[**DownloadStreamTenantAuditLogs**](StreamsApi.md#DownloadStreamTenantAuditLogs) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs | Download Audit Logs for One Atlas Stream Processing Instance
[**GetPrivateLinkConnection**](StreamsApi.md#GetPrivateLinkConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections/{connectionId} | Return One PrivateLink Connection
[**GetStreamConnection**](StreamsApi.md#GetStreamConnection) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Return One Stream Connection
[**GetStreamInstance**](StreamsApi.md#GetStreamInstance) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Return One Stream Instance
[**GetVpcPeeringConnections**](StreamsApi.md#GetVpcPeeringConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections | Return All VPC Peering Connections.
[**ListPrivateLinkConnections**](StreamsApi.md#ListPrivateLinkConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/privateLinkConnections | Return All Private Link
[**ListStreamConnections**](StreamsApi.md#ListStreamConnections) | **Get** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections | Return All Connections Of The Stream Instances
[**ListStreamInstances**](StreamsApi.md#ListStreamInstances) | **Get** /api/atlas/v2/groups/{groupId}/streams | Return All Project Stream Instances
[**RejectVpcPeeringConnection**](StreamsApi.md#RejectVpcPeeringConnection) | **Post** /api/atlas/v2/groups/{groupId}/streams/vpcPeeringConnections/{id}:reject | Requests the rejection of an incoming VPC Peering connection.
[**UpdateStreamConnection**](StreamsApi.md#UpdateStreamConnection) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName} | Update One Stream Connection
[**UpdateStreamInstance**](StreamsApi.md#UpdateStreamInstance) | **Patch** /api/atlas/v2/groups/{groupId}/streams/{tenantName} | Update One Stream Instance



## AcceptVpcPeeringConnection

> AcceptVpcPeeringConnection(ctx, groupId, id, vPCPeeringActionChallenge VPCPeeringActionChallenge).Execute()

Requests the acceptance of an incoming VPC Peering connection.


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    id := "id_example" // string | 
    vPCPeeringActionChallenge := *openapiclient.NewVPCPeeringActionChallenge() // VPCPeeringActionChallenge | 

    r, err := sdk.StreamsApi.AcceptVpcPeeringConnection(context.Background(), groupId, id, &vPCPeeringActionChallenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.AcceptVpcPeeringConnection``: %v\n", err)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateLinkConnection

> StreamsPrivateLinkConnection CreatePrivateLinkConnection(ctx, groupId, streamsPrivateLinkConnection StreamsPrivateLinkConnection).Execute()

Create One Private Link connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    streamsPrivateLinkConnection := *openapiclient.NewStreamsPrivateLinkConnection() // StreamsPrivateLinkConnection | 

    resp, r, err := sdk.StreamsApi.CreatePrivateLinkConnection(context.Background(), groupId, &streamsPrivateLinkConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreatePrivateLinkConnection``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreatePrivateLinkConnection`: StreamsPrivateLinkConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreatePrivateLinkConnection`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamConnection

> StreamsConnection CreateStreamConnection(ctx, groupId, tenantName, streamsConnection StreamsConnection).Execute()

Create One Connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    streamsConnection := *openapiclient.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsApi.CreateStreamConnection(context.Background(), groupId, tenantName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamConnection``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamConnection`: %v\n", resp)
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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    streamsTenant := *openapiclient.NewStreamsTenant() // StreamsTenant | 

    resp, r, err := sdk.StreamsApi.CreateStreamInstance(context.Background(), groupId, &streamsTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreamInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreamInstance`: %v\n", resp)
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


## DeletePrivateLinkConnection

> DeletePrivateLinkConnection(ctx, groupId, connectionId).Execute()

Delete One Private Link connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    connectionId := "connectionId_example" // string | 

    r, err := sdk.StreamsApi.DeletePrivateLinkConnection(context.Background(), groupId, connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeletePrivateLinkConnection``: %v\n", err)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamConnection

> DeleteStreamConnection(ctx, groupId, tenantName, connectionName).Execute()

Delete One Stream Connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    connectionName := "connectionName_example" // string | 

    r, err := sdk.StreamsApi.DeleteStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamConnection``: %v\n", err)
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
**tenantName** | **string** | Human-readable label that identifies the stream instance. | 
**connectionName** | **string** | Human-readable label that identifies the stream connection. | 

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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamInstance

> DeleteStreamInstance(ctx, groupId, tenantName).Execute()

Delete One Stream Instance


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 

    r, err := sdk.StreamsApi.DeleteStreamInstance(context.Background(), groupId, tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamInstance``: %v\n", err)
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
**tenantName** | **string** | Human-readable label that identifies the stream instance to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcPeeringConnection

> DeleteVpcPeeringConnection(ctx, groupId, id).Execute()

Deletes an incoming VPC Peering connection.


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    id := "id_example" // string | 

    r, err := sdk.StreamsApi.DeleteVpcPeeringConnection(context.Background(), groupId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteVpcPeeringConnection``: %v\n", err)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadStreamTenantAuditLogs

> io.ReadCloser DownloadStreamTenantAuditLogs(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()

Download Audit Logs for One Atlas Stream Processing Instance


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    endDate := int64(1636481348) // int64 |  (optional)
    startDate := int64(1636466948) // int64 |  (optional)

    resp, r, err := sdk.StreamsApi.DownloadStreamTenantAuditLogs(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DownloadStreamTenantAuditLogs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DownloadStreamTenantAuditLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DownloadStreamTenantAuditLogs`: %v\n", resp)
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


## GetPrivateLinkConnection

> StreamsPrivateLinkConnection GetPrivateLinkConnection(ctx, groupId, connectionId).Execute()

Return One PrivateLink Connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    connectionId := "connectionId_example" // string | 

    resp, r, err := sdk.StreamsApi.GetPrivateLinkConnection(context.Background(), groupId, connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetPrivateLinkConnection``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetPrivateLinkConnection`: StreamsPrivateLinkConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetPrivateLinkConnection`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamConnection

> StreamsConnection GetStreamConnection(ctx, groupId, tenantName, connectionName).Execute()

Return One Stream Connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    connectionName := "connectionName_example" // string | 

    resp, r, err := sdk.StreamsApi.GetStreamConnection(context.Background(), groupId, tenantName, connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamConnection``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamConnection`: %v\n", resp)
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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    includeConnections := true // bool |  (optional)

    resp, r, err := sdk.StreamsApi.GetStreamInstance(context.Background(), groupId, tenantName).IncludeConnections(includeConnections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamInstance`: %v\n", resp)
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


## GetVpcPeeringConnections

> GetVpcPeeringConnections(ctx, groupId).RequesterAccountId(requesterAccountId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All VPC Peering Connections.


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    requesterAccountId := "requesterAccountId_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    r, err := sdk.StreamsApi.GetVpcPeeringConnections(context.Background(), groupId).RequesterAccountId(requesterAccountId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetVpcPeeringConnections``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcPeeringConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requesterAccountId** | **string** | The Account ID of the VPC Peering connection/s. | 

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateLinkConnections

> PaginatedApiStreamsPrivateLink ListPrivateLinkConnections(ctx, groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Private Link


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsApi.ListPrivateLinkConnections(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListPrivateLinkConnections``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListPrivateLinkConnections`: PaginatedApiStreamsPrivateLink
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListPrivateLinkConnections`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamConnections

> PaginatedApiStreamsConnection ListStreamConnections(ctx, groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Connections Of The Stream Instances


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsApi.ListStreamConnections(context.Background(), groupId, tenantName).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListStreamConnections``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListStreamConnections`: PaginatedApiStreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListStreamConnections`: %v\n", resp)
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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.StreamsApi.ListStreamInstances(context.Background(), groupId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.ListStreamInstances``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListStreamInstances`: PaginatedApiStreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.ListStreamInstances`: %v\n", resp)
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


## RejectVpcPeeringConnection

> RejectVpcPeeringConnection(ctx, groupId, id).Execute()

Requests the rejection of an incoming VPC Peering connection.


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    id := "id_example" // string | 

    r, err := sdk.StreamsApi.RejectVpcPeeringConnection(context.Background(), groupId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.RejectVpcPeeringConnection``: %v\n", err)
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
- **Accept**: application/vnd.atlas.2023-02-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamConnection

> StreamsConnection UpdateStreamConnection(ctx, groupId, tenantName, connectionName, streamsConnection StreamsConnection).Execute()

Update One Stream Connection


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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    connectionName := "connectionName_example" // string | 
    streamsConnection := *openapiclient.NewStreamsConnection() // StreamsConnection | 

    resp, r, err := sdk.StreamsApi.UpdateStreamConnection(context.Background(), groupId, tenantName, connectionName, &streamsConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamConnection``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateStreamConnection`: StreamsConnection
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamConnection`: %v\n", resp)
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

    "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    tenantName := "tenantName_example" // string | 
    streamsDataProcessRegion := *openapiclient.NewStreamsDataProcessRegion("CloudProvider_example", "Region_example") // StreamsDataProcessRegion | 

    resp, r, err := sdk.StreamsApi.UpdateStreamInstance(context.Background(), groupId, tenantName, &streamsDataProcessRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamInstance``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateStreamInstance`: StreamsTenant
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamInstance`: %v\n", resp)
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

