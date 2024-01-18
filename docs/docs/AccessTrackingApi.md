# \AccessTrackingApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccessLogsByClusterName**](AccessTrackingApi.md#ListAccessLogsByClusterName) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster using Its Cluster Name
[**ListAccessLogsByHostname**](AccessTrackingApi.md#ListAccessLogsByHostname) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster using Its Hostname



## ListAccessLogsByClusterName

> MongoDBAccessLogsList ListAccessLogsByClusterName(ctx, groupId, clusterName).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()

Return Database Access History for One Cluster using Its Cluster Name


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    authResult := true // bool |  (optional)
    end := int64(789) // int64 |  (optional)
    ipAddress := "ipAddress_example" // string |  (optional)
    nLogs := int(56) // int |  (optional) (default to 20000)
    start := int64(789) // int64 |  (optional)

    resp, r, err := sdk.AccessTrackingApi.ListAccessLogsByClusterName(context.Background(), groupId, clusterName).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTrackingApi.ListAccessLogsByClusterName``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListAccessLogsByClusterName`: MongoDBAccessLogsList
    fmt.Fprintf(os.Stdout, "Response from `AccessTrackingApi.ListAccessLogsByClusterName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessLogsByClusterNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authResult** | **bool** | Flag that indicates whether the response returns the successful authentication attempts only. | 
 **end** | **int64** | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds. | 
 **ipAddress** | **string** | One Internet Protocol address that attempted to authenticate with the database. | 
 **nLogs** | **int** | Maximum number of lines from the log to return. | [default to 20000]
 **start** | **int64** | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds. | 

### Return type

[**MongoDBAccessLogsList**](MongoDBAccessLogsList.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessLogsByHostname

> MongoDBAccessLogsList ListAccessLogsByHostname(ctx, groupId, hostname).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()

Return Database Access History for One Cluster using Its Hostname


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115004/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    hostname := "hostname_example" // string | 
    authResult := true // bool |  (optional)
    end := int64(789) // int64 |  (optional)
    ipAddress := "ipAddress_example" // string |  (optional)
    nLogs := int(56) // int |  (optional) (default to 20000)
    start := int64(789) // int64 |  (optional)

    resp, r, err := sdk.AccessTrackingApi.ListAccessLogsByHostname(context.Background(), groupId, hostname).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTrackingApi.ListAccessLogsByHostname``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListAccessLogsByHostname`: MongoDBAccessLogsList
    fmt.Fprintf(os.Stdout, "Response from `AccessTrackingApi.ListAccessLogsByHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**hostname** | **string** | Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessLogsByHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authResult** | **bool** | Flag that indicates whether the response returns the successful authentication attempts only. | 
 **end** | **int64** | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds. | 
 **ipAddress** | **string** | One Internet Protocol address that attempted to authenticate with the database. | 
 **nLogs** | **int** | Maximum number of lines from the log to return. | [default to 20000]
 **start** | **int64** | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds. | 

### Return type

[**MongoDBAccessLogsList**](MongoDBAccessLogsList.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

