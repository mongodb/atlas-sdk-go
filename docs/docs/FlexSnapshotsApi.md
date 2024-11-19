# \FlexSnapshotsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadFlexBackup**](FlexSnapshotsApi.md#DownloadFlexBackup) | **Post** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/download | Download One Flex Cluster Snapshot
[**GetFlexBackup**](FlexSnapshotsApi.md#GetFlexBackup) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots/{snapshotId} | Return One Snapshot of One Flex Cluster
[**ListFlexBackups**](FlexSnapshotsApi.md#ListFlexBackups) | **Get** /api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/snapshots | Return All Snapshots of One Flex Cluster



## DownloadFlexBackup

> FlexBackupRestoreJob20241113 DownloadFlexBackup(ctx, name, groupId, flexBackupSnapshotDownloadCreate20241113 FlexBackupSnapshotDownloadCreate20241113).Execute()

Download One Flex Cluster Snapshot


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    name := "name_example" // string | 
    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    flexBackupSnapshotDownloadCreate20241113 := *openapiclient.NewFlexBackupSnapshotDownloadCreate20241113("32b6e34b3d91647abb20e7b8") // FlexBackupSnapshotDownloadCreate20241113 | 

    resp, r, err := sdk.FlexSnapshotsApi.DownloadFlexBackup(context.Background(), name, groupId, &flexBackupSnapshotDownloadCreate20241113).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexSnapshotsApi.DownloadFlexBackup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadFlexBackup`: FlexBackupRestoreJob20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexSnapshotsApi.DownloadFlexBackup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Human-readable label that identifies the flex cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFlexBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flexBackupSnapshotDownloadCreate20241113** | [**FlexBackupSnapshotDownloadCreate20241113**](FlexBackupSnapshotDownloadCreate20241113.md) | Snapshot to be downloaded. | 

### Return type

[**FlexBackupRestoreJob20241113**](FlexBackupRestoreJob20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-11-13+json
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexBackup

> FlexBackupSnapshot20241113 GetFlexBackup(ctx, groupId, name, snapshotId).Execute()

Return One Snapshot of One Flex Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113001/admin"
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
    name := "name_example" // string | 
    snapshotId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FlexSnapshotsApi.GetFlexBackup(context.Background(), groupId, name, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexSnapshotsApi.GetFlexBackup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetFlexBackup`: FlexBackupSnapshot20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexSnapshotsApi.GetFlexBackup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlexBackupSnapshot20241113**](FlexBackupSnapshot20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexBackups

> PaginatedApiAtlasFlexBackupSnapshot20241113 ListFlexBackups(ctx, groupId, name).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Snapshots of One Flex Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113001/admin"
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
    name := "name_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.FlexSnapshotsApi.ListFlexBackups(context.Background(), groupId, name).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexSnapshotsApi.ListFlexBackups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListFlexBackups`: PaginatedApiAtlasFlexBackupSnapshot20241113
    fmt.Fprintf(os.Stdout, "Response from `FlexSnapshotsApi.ListFlexBackups`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**name** | **string** | Human-readable label that identifies the flex cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFlexBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasFlexBackupSnapshot20241113**](PaginatedApiAtlasFlexBackupSnapshot20241113.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

