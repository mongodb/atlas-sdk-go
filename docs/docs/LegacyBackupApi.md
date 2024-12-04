# \LegacyBackupApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLegacyBackupRestoreJob**](LegacyBackupApi.md#CreateLegacyBackupRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job
[**DeleteLegacySnapshot**](LegacyBackupApi.md#DeleteLegacySnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot
[**GetLegacyBackupCheckpoint**](LegacyBackupApi.md#GetLegacyBackupCheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint
[**GetLegacyBackupRestoreJob**](LegacyBackupApi.md#GetLegacyBackupRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job
[**GetLegacySnapshot**](LegacyBackupApi.md#GetLegacySnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot
[**GetLegacySnapshotSchedule**](LegacyBackupApi.md#GetLegacySnapshotSchedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule
[**ListLegacyBackupCheckpoints**](LegacyBackupApi.md#ListLegacyBackupCheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints
[**ListLegacyBackupRestoreJobs**](LegacyBackupApi.md#ListLegacyBackupRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs
[**ListLegacySnapshots**](LegacyBackupApi.md#ListLegacySnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots
[**UpdateLegacySnapshotRetention**](LegacyBackupApi.md#UpdateLegacySnapshotRetention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Change One Legacy Backup Snapshot Expiration
[**UpdateLegacySnapshotSchedule**](LegacyBackupApi.md#UpdateLegacySnapshotSchedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster



## CreateLegacyBackupRestoreJob

> PaginatedRestoreJob CreateLegacyBackupRestoreJob(ctx, groupId, clusterName, backupRestoreJob BackupRestoreJob).Execute()

Create One Legacy Backup Restore Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    backupRestoreJob := *openapiclient.NewBackupRestoreJob(*openapiclient.NewBackupRestoreJobDelivery("MethodName_example")) // BackupRestoreJob | 

    resp, r, err := sdk.LegacyBackupApi.CreateLegacyBackupRestoreJob(context.Background(), groupId, clusterName, &backupRestoreJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.CreateLegacyBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateLegacyBackupRestoreJob`: PaginatedRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.CreateLegacyBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLegacyBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupRestoreJob** | [**BackupRestoreJob**](BackupRestoreJob.md) | Legacy backup to restore to one cluster in the specified project. | 

### Return type

[**PaginatedRestoreJob**](PaginatedRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLegacySnapshot

> any DeleteLegacySnapshot(ctx, groupId, clusterName, snapshotId).Execute()

Remove One Legacy Backup Snapshot


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.LegacyBackupApi.DeleteLegacySnapshot(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.DeleteLegacySnapshot`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteLegacySnapshot`: any
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.DeleteLegacySnapshot`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLegacySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacyBackupCheckpoint

> ApiAtlasCheckpoint GetLegacyBackupCheckpoint(ctx, groupId, checkpointId, clusterName).Execute()

Return One Legacy Backup Checkpoint


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    checkpointId := "checkpointId_example" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.LegacyBackupApi.GetLegacyBackupCheckpoint(context.Background(), groupId, checkpointId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.GetLegacyBackupCheckpoint`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLegacyBackupCheckpoint`: ApiAtlasCheckpoint
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.GetLegacyBackupCheckpoint`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**checkpointId** | **string** | Unique 24-hexadecimal digit string that identifies the checkpoint. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the checkpoints that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyBackupCheckpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiAtlasCheckpoint**](ApiAtlasCheckpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacyBackupRestoreJob

> BackupRestoreJob GetLegacyBackupRestoreJob(ctx, groupId, clusterName, jobId).Execute()

Return One Legacy Backup Restore Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    jobId := "jobId_example" // string | 

    resp, r, err := sdk.LegacyBackupApi.GetLegacyBackupRestoreJob(context.Background(), groupId, clusterName, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.GetLegacyBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLegacyBackupRestoreJob`: BackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.GetLegacyBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 
**jobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BackupRestoreJob**](BackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacySnapshot

> BackupSnapshot GetLegacySnapshot(ctx, groupId, clusterName, snapshotId).Execute()

Return One Legacy Backup Snapshot


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.LegacyBackupApi.GetLegacySnapshot(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.GetLegacySnapshot`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLegacySnapshot`: BackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.GetLegacySnapshot`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BackupSnapshot**](BackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacySnapshotSchedule

> ApiAtlasSnapshotSchedule GetLegacySnapshotSchedule(ctx, groupId, clusterName).Execute()

Return One Snapshot Schedule


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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

    resp, r, err := sdk.LegacyBackupApi.GetLegacySnapshotSchedule(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.GetLegacySnapshotSchedule`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetLegacySnapshotSchedule`: ApiAtlasSnapshotSchedule
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.GetLegacySnapshotSchedule`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacySnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAtlasSnapshotSchedule**](ApiAtlasSnapshotSchedule.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyBackupCheckpoints

> PaginatedApiAtlasCheckpoint ListLegacyBackupCheckpoints(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Legacy Backup Checkpoints


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.LegacyBackupApi.ListLegacyBackupCheckpoints(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.ListLegacyBackupCheckpoints`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListLegacyBackupCheckpoints`: PaginatedApiAtlasCheckpoint
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.ListLegacyBackupCheckpoints`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the checkpoints that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyBackupCheckpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasCheckpoint**](PaginatedApiAtlasCheckpoint.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyBackupRestoreJobs

> PaginatedRestoreJob ListLegacyBackupRestoreJobs(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).BatchId(batchId).Execute()

Return All Legacy Backup Restore Jobs


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    batchId := "batchId_example" // string |  (optional)

    resp, r, err := sdk.LegacyBackupApi.ListLegacyBackupRestoreJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).BatchId(batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.ListLegacyBackupRestoreJobs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListLegacyBackupRestoreJobs`: PaginatedRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.ListLegacyBackupRestoreJobs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyBackupRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **batchId** | **string** | Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch. | 

### Return type

[**PaginatedRestoreJob**](PaginatedRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacySnapshots

> PaginatedSnapshot ListLegacySnapshots(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Completed(completed).Execute()

Return All Legacy Backup Snapshots


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    completed := "completed_example" // string |  (optional) (default to "true")

    resp, r, err := sdk.LegacyBackupApi.ListLegacySnapshots(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Completed(completed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.ListLegacySnapshots`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListLegacySnapshots`: PaginatedSnapshot
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.ListLegacySnapshots`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacySnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **completed** | **string** | Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots. | [default to &quot;true&quot;]

### Return type

[**PaginatedSnapshot**](PaginatedSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLegacySnapshotRetention

> BackupSnapshot UpdateLegacySnapshotRetention(ctx, groupId, clusterName, snapshotId, backupSnapshot BackupSnapshot).Execute()

Change One Legacy Backup Snapshot Expiration


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    snapshotId := "snapshotId_example" // string | 
    backupSnapshot := *openapiclient.NewBackupSnapshot() // BackupSnapshot | 

    resp, r, err := sdk.LegacyBackupApi.UpdateLegacySnapshotRetention(context.Background(), groupId, clusterName, snapshotId, &backupSnapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.UpdateLegacySnapshotRetention`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateLegacySnapshotRetention`: BackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.UpdateLegacySnapshotRetention`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLegacySnapshotRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backupSnapshot** | [**BackupSnapshot**](BackupSnapshot.md) | Changes One Legacy Backup Snapshot Expiration. | 

### Return type

[**BackupSnapshot**](BackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLegacySnapshotSchedule

> ApiAtlasSnapshotSchedule UpdateLegacySnapshotSchedule(ctx, groupId, clusterName, apiAtlasSnapshotSchedule ApiAtlasSnapshotSchedule).Execute()

Update Snapshot Schedule for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241113002/admin"
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
    apiAtlasSnapshotSchedule := *openapiclient.NewApiAtlasSnapshotSchedule(int(123), "32b6e34b3d91647abb20e7b8", int(123), "32b6e34b3d91647abb20e7b8", int(123), int(123), int(123), int(123), int(123)) // ApiAtlasSnapshotSchedule | 

    resp, r, err := sdk.LegacyBackupApi.UpdateLegacySnapshotSchedule(context.Background(), groupId, clusterName, &apiAtlasSnapshotSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupApi.UpdateLegacySnapshotSchedule`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateLegacySnapshotSchedule`: ApiAtlasSnapshotSchedule
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupApi.UpdateLegacySnapshotSchedule`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLegacySnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasSnapshotSchedule** | [**ApiAtlasSnapshotSchedule**](ApiAtlasSnapshotSchedule.md) | Update the snapshot schedule for one cluster in the specified project. | 

### Return type

[**ApiAtlasSnapshotSchedule**](ApiAtlasSnapshotSchedule.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

