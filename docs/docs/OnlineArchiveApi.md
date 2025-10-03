# \OnlineArchiveApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOnlineArchive**](OnlineArchiveApi.md#CreateOnlineArchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
[**DeleteOnlineArchive**](OnlineArchiveApi.md#DeleteOnlineArchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
[**DownloadQueryLogs**](OnlineArchiveApi.md#DownloadQueryLogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
[**GetOnlineArchive**](OnlineArchiveApi.md#GetOnlineArchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
[**ListOnlineArchives**](OnlineArchiveApi.md#ListOnlineArchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
[**UpdateOnlineArchive**](OnlineArchiveApi.md#UpdateOnlineArchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive



## CreateOnlineArchive

> BackupOnlineArchive CreateOnlineArchive(ctx, groupId, clusterName, backupOnlineArchiveCreate BackupOnlineArchiveCreate).Execute()

Create One Online Archive


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
    clusterName := "clusterName_example" // string | 
    backupOnlineArchiveCreate := *openapiclient.NewBackupOnlineArchiveCreate("CollName_example", *openapiclient.NewCriteria(), "DbName_example") // BackupOnlineArchiveCreate | 

    resp, r, err := sdk.OnlineArchiveApi.CreateOnlineArchive(context.Background(), groupId, clusterName, &backupOnlineArchiveCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.CreateOnlineArchive`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOnlineArchive`: BackupOnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.CreateOnlineArchive`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupOnlineArchiveCreate** | [**BackupOnlineArchiveCreate**](BackupOnlineArchiveCreate.md) | Creates one online archive. | 

### Return type

[**BackupOnlineArchive**](BackupOnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOnlineArchive

> DeleteOnlineArchive(ctx, groupId, archiveId, clusterName).Execute()

Remove One Online Archive


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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 

    r, err := sdk.OnlineArchiveApi.DeleteOnlineArchive(context.Background(), groupId, archiveId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.DeleteOnlineArchive`: %v (%v)\n", err, r)
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
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to delete. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOnlineArchiveRequest struct via the builder pattern


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


## DownloadQueryLogs

> io.ReadCloser DownloadQueryLogs(ctx, groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()

Download Online Archive Query Logs


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
    clusterName := "clusterName_example" // string | 
    startDate := int64(1636481348) // int64 |  (optional)
    endDate := int64(1636481348) // int64 |  (optional)
    archiveOnly := true // bool |  (optional) (default to false)

    resp, r, err := sdk.OnlineArchiveApi.DownloadQueryLogs(context.Background(), groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.DownloadQueryLogs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadQueryLogs`: io.ReadCloser
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.DownloadQueryLogs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **int64** | Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **endDate** | **int64** | Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **archiveOnly** | **bool** | Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster. | [default to false]

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


## GetOnlineArchive

> BackupOnlineArchive GetOnlineArchive(ctx, groupId, archiveId, clusterName).Execute()

Return One Online Archive


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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.OnlineArchiveApi.GetOnlineArchive(context.Background(), groupId, archiveId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.GetOnlineArchive`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOnlineArchive`: BackupOnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.GetOnlineArchive`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to return. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BackupOnlineArchive**](BackupOnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOnlineArchives

> PaginatedOnlineArchive ListOnlineArchives(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Online Archives for One Cluster


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
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.OnlineArchiveApi.ListOnlineArchives(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.ListOnlineArchives`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOnlineArchives`: PaginatedOnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.ListOnlineArchives`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOnlineArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedOnlineArchive**](PaginatedOnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnlineArchive

> BackupOnlineArchive UpdateOnlineArchive(ctx, groupId, archiveId, clusterName, backupOnlineArchive BackupOnlineArchive).Execute()

Update One Online Archive


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
    archiveId := "archiveId_example" // string | 
    clusterName := "clusterName_example" // string | 
    backupOnlineArchive := *openapiclient.NewBackupOnlineArchive() // BackupOnlineArchive | 

    resp, r, err := sdk.OnlineArchiveApi.UpdateOnlineArchive(context.Background(), groupId, archiveId, clusterName, &backupOnlineArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.UpdateOnlineArchive`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateOnlineArchive`: BackupOnlineArchive
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.UpdateOnlineArchive`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to update. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backupOnlineArchive** | [**BackupOnlineArchive**](BackupOnlineArchive.md) | Updates, pauses, or resumes one online archive. | 

### Return type

[**BackupOnlineArchive**](BackupOnlineArchive.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

