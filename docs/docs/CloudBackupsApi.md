# \CloudBackupsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBackupRestoreJob**](CloudBackupsApi.md#CancelBackupRestoreJob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
[**CreateBackupExportJob**](CloudBackupsApi.md#CreateBackupExportJob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Snapshot Export Job
[**CreateBackupRestoreJob**](CloudBackupsApi.md#CreateBackupRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
[**CreateExportBucket**](CloudBackupsApi.md#CreateExportBucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Create One Snapshot Export Bucket
[**CreateServerlessBackupRestoreJob**](CloudBackupsApi.md#CreateServerlessBackupRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
[**DeleteAllBackupSchedules**](CloudBackupsApi.md#DeleteAllBackupSchedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
[**DeleteExportBucket**](CloudBackupsApi.md#DeleteExportBucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Delete One Snapshot Export Bucket
[**DeleteReplicaSetBackup**](CloudBackupsApi.md#DeleteReplicaSetBackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
[**DeleteShardedClusterBackup**](CloudBackupsApi.md#DeleteShardedClusterBackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
[**DisableDataProtectionSettings**](CloudBackupsApi.md#DisableDataProtectionSettings) | **Delete** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Disable the Backup Compliance Policy settings
[**GetBackupExportJob**](CloudBackupsApi.md#GetBackupExportJob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Snapshot Export Job
[**GetBackupRestoreJob**](CloudBackupsApi.md#GetBackupRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster
[**GetBackupSchedule**](CloudBackupsApi.md#GetBackupSchedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
[**GetDataProtectionSettings**](CloudBackupsApi.md#GetDataProtectionSettings) | **Get** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Return the Backup Compliance Policy settings
[**GetExportBucket**](CloudBackupsApi.md#GetExportBucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One Snapshot Export Bucket
[**GetReplicaSetBackup**](CloudBackupsApi.md#GetReplicaSetBackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
[**GetServerlessBackup**](CloudBackupsApi.md#GetServerlessBackup) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
[**GetServerlessBackupRestoreJob**](CloudBackupsApi.md#GetServerlessBackupRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
[**GetShardedClusterBackup**](CloudBackupsApi.md#GetShardedClusterBackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
[**ListBackupExportJobs**](CloudBackupsApi.md#ListBackupExportJobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Snapshot Export Jobs
[**ListBackupRestoreJobs**](CloudBackupsApi.md#ListBackupRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
[**ListExportBuckets**](CloudBackupsApi.md#ListExportBuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All Snapshot Export Buckets
[**ListReplicaSetBackups**](CloudBackupsApi.md#ListReplicaSetBackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
[**ListServerlessBackupRestoreJobs**](CloudBackupsApi.md#ListServerlessBackupRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
[**ListServerlessBackups**](CloudBackupsApi.md#ListServerlessBackups) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
[**ListShardedClusterBackups**](CloudBackupsApi.md#ListShardedClusterBackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
[**TakeSnapshot**](CloudBackupsApi.md#TakeSnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot
[**UpdateBackupSchedule**](CloudBackupsApi.md#UpdateBackupSchedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster
[**UpdateDataProtectionSettings**](CloudBackupsApi.md#UpdateDataProtectionSettings) | **Put** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Update or enable the Backup Compliance Policy settings
[**UpdateSnapshotRetention**](CloudBackupsApi.md#UpdateSnapshotRetention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup



## CancelBackupRestoreJob

> CancelBackupRestoreJob(ctx, groupId, clusterName, restoreJobId).Execute()

Cancel One Restore Job of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    restoreJobId := "restoreJobId_example" // string | 

    r, err := sdk.CloudBackupsApi.CancelBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CancelBackupRestoreJob`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBackupRestoreJobRequest struct via the builder pattern


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


## CreateBackupExportJob

> DiskBackupExportJob CreateBackupExportJob(ctx, groupId, clusterName, diskBackupExportJobRequest DiskBackupExportJobRequest).Execute()

Create One Snapshot Export Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    diskBackupExportJobRequest := *openapiclient.NewDiskBackupExportJobRequest("32b6e34b3d91647abb20e7b8", "32b6e34b3d91647abb20e7b8") // DiskBackupExportJobRequest | 

    resp, r, err := sdk.CloudBackupsApi.CreateBackupExportJob(context.Background(), groupId, clusterName, &diskBackupExportJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateBackupExportJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateBackupExportJob`: DiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateBackupExportJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **diskBackupExportJobRequest** | [**DiskBackupExportJobRequest**](DiskBackupExportJobRequest.md) | Information about the Cloud Backup Snapshot Export Job to create. | 

### Return type

[**DiskBackupExportJob**](DiskBackupExportJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupRestoreJob

> DiskBackupSnapshotRestoreJob CreateBackupRestoreJob(ctx, groupId, clusterName, diskBackupSnapshotRestoreJob DiskBackupSnapshotRestoreJob).Execute()

Restore One Snapshot of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    diskBackupSnapshotRestoreJob := *openapiclient.NewDiskBackupSnapshotRestoreJob("DeliveryType_example") // DiskBackupSnapshotRestoreJob | 

    resp, r, err := sdk.CloudBackupsApi.CreateBackupRestoreJob(context.Background(), groupId, clusterName, &diskBackupSnapshotRestoreJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateBackupRestoreJob`: DiskBackupSnapshotRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **diskBackupSnapshotRestoreJob** | [**DiskBackupSnapshotRestoreJob**](DiskBackupSnapshotRestoreJob.md) | Restores one snapshot of one cluster from the specified project. | 

### Return type

[**DiskBackupSnapshotRestoreJob**](DiskBackupSnapshotRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExportBucket

> DiskBackupSnapshotExportBucketResponse CreateExportBucket(ctx, groupId, diskBackupSnapshotExportBucketRequest DiskBackupSnapshotExportBucketRequest).Execute()

Create One Snapshot Export Bucket


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    diskBackupSnapshotExportBucketRequest := *openapiclient.NewDiskBackupSnapshotExportBucketRequest("CloudProvider_example") // DiskBackupSnapshotExportBucketRequest | 

    resp, r, err := sdk.CloudBackupsApi.CreateExportBucket(context.Background(), groupId, &diskBackupSnapshotExportBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateExportBucket`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateExportBucket`: DiskBackupSnapshotExportBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateExportBucket`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExportBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskBackupSnapshotExportBucketRequest** | [**DiskBackupSnapshotExportBucketRequest**](DiskBackupSnapshotExportBucketRequest.md) | Specifies the role and AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket that the Export Bucket should reference. | 

### Return type

[**DiskBackupSnapshotExportBucketResponse**](DiskBackupSnapshotExportBucketResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-05-30+json
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerlessBackupRestoreJob

> ServerlessBackupRestoreJob CreateServerlessBackupRestoreJob(ctx, groupId, clusterName, serverlessBackupRestoreJob ServerlessBackupRestoreJob).Execute()

Restore One Snapshot of One Serverless Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    serverlessBackupRestoreJob := *openapiclient.NewServerlessBackupRestoreJob("DeliveryType_example", "TargetClusterName_example", "32b6e34b3d91647abb20e7b8") // ServerlessBackupRestoreJob | 

    resp, r, err := sdk.CloudBackupsApi.CreateServerlessBackupRestoreJob(context.Background(), groupId, clusterName, &serverlessBackupRestoreJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateServerlessBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateServerlessBackupRestoreJob`: ServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateServerlessBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance whose snapshot you want to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerlessBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverlessBackupRestoreJob** | [**ServerlessBackupRestoreJob**](ServerlessBackupRestoreJob.md) | Restores one snapshot of one serverless instance from the specified project. | 

### Return type

[**ServerlessBackupRestoreJob**](ServerlessBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllBackupSchedules

> DiskBackupSnapshotSchedule20240805 DeleteAllBackupSchedules(ctx, groupId, clusterName).Execute()

Remove All Cloud Backup Schedules


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.DeleteAllBackupSchedules(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteAllBackupSchedules`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DeleteAllBackupSchedules`: DiskBackupSnapshotSchedule20240805
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.DeleteAllBackupSchedules`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DiskBackupSnapshotSchedule20240805**](DiskBackupSnapshotSchedule20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExportBucket

> DeleteExportBucket(ctx, groupId, exportBucketId).Execute()

Delete One Snapshot Export Bucket


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    exportBucketId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.CloudBackupsApi.DeleteExportBucket(context.Background(), groupId, exportBucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteExportBucket`: %v (%v)\n", err, r)
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
**exportBucketId** | **string** | Unique 24-hexadecimal character string that identifies the Export Bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExportBucketRequest struct via the builder pattern


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


## DeleteReplicaSetBackup

> DeleteReplicaSetBackup(ctx, groupId, clusterName, snapshotId).Execute()

Remove One Replica Set Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    r, err := sdk.CloudBackupsApi.DeleteReplicaSetBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteReplicaSetBackup`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReplicaSetBackupRequest struct via the builder pattern


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


## DeleteShardedClusterBackup

> DeleteShardedClusterBackup(ctx, groupId, clusterName, snapshotId).Execute()

Remove One Sharded Cluster Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    r, err := sdk.CloudBackupsApi.DeleteShardedClusterBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteShardedClusterBackup`: %v (%v)\n", err, r)
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
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShardedClusterBackupRequest struct via the builder pattern


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


## DisableDataProtectionSettings

> DisableDataProtectionSettings(ctx, groupId).Execute()

Disable the Backup Compliance Policy settings


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    r, err := sdk.CloudBackupsApi.DisableDataProtectionSettings(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DisableDataProtectionSettings`: %v (%v)\n", err, r)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDisableDataProtectionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-11-13+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupExportJob

> DiskBackupExportJob GetBackupExportJob(ctx, groupId, clusterName, exportId).Execute()

Return One Snapshot Export Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    exportId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetBackupExportJob(context.Background(), groupId, clusterName, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupExportJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetBackupExportJob`: DiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupExportJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**exportId** | **string** | Unique 24-hexadecimal character string that identifies the Export Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DiskBackupExportJob**](DiskBackupExportJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupRestoreJob

> DiskBackupSnapshotRestoreJob GetBackupRestoreJob(ctx, groupId, clusterName, restoreJobId).Execute()

Return One Restore Job of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetBackupRestoreJob`: DiskBackupSnapshotRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the restore jobs you want to return. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DiskBackupSnapshotRestoreJob**](DiskBackupSnapshotRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupSchedule

> DiskBackupSnapshotSchedule20240805 GetBackupSchedule(ctx, groupId, clusterName).Execute()

Return One Cloud Backup Schedule


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.GetBackupSchedule(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupSchedule`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetBackupSchedule`: DiskBackupSnapshotSchedule20240805
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupSchedule`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DiskBackupSnapshotSchedule20240805**](DiskBackupSnapshotSchedule20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataProtectionSettings

> DataProtectionSettings20231001 GetDataProtectionSettings(ctx, groupId).Execute()

Return the Backup Compliance Policy settings


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.GetDataProtectionSettings(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetDataProtectionSettings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetDataProtectionSettings`: DataProtectionSettings20231001
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetDataProtectionSettings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataProtectionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProtectionSettings20231001**](DataProtectionSettings20231001.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-10-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportBucket

> DiskBackupSnapshotExportBucketResponse GetExportBucket(ctx, groupId, exportBucketId).Execute()

Return One Snapshot Export Bucket


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    exportBucketId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetExportBucket(context.Background(), groupId, exportBucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetExportBucket`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetExportBucket`: DiskBackupSnapshotExportBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetExportBucket`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**exportBucketId** | **string** | Unique 24-hexadecimal character string that identifies the Export Bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DiskBackupSnapshotExportBucketResponse**](DiskBackupSnapshotExportBucketResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicaSetBackup

> DiskBackupReplicaSet GetReplicaSetBackup(ctx, groupId, clusterName, snapshotId).Execute()

Return One Replica Set Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.GetReplicaSetBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetReplicaSetBackup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetReplicaSetBackup`: DiskBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetReplicaSetBackup`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiGetReplicaSetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DiskBackupReplicaSet**](DiskBackupReplicaSet.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerlessBackup

> ServerlessBackupSnapshot GetServerlessBackup(ctx, groupId, clusterName, snapshotId).Execute()

Return One Snapshot of One Serverless Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    snapshotId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetServerlessBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetServerlessBackup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetServerlessBackup`: ServerlessBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetServerlessBackup`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerlessBackupSnapshot**](ServerlessBackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerlessBackupRestoreJob

> ServerlessBackupRestoreJob GetServerlessBackupRestoreJob(ctx, groupId, clusterName, restoreJobId).Execute()

Return One Restore Job for One Serverless Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetServerlessBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetServerlessBackupRestoreJob`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetServerlessBackupRestoreJob`: ServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetServerlessBackupRestoreJob`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerlessBackupRestoreJob**](ServerlessBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShardedClusterBackup

> DiskBackupShardedClusterSnapshot GetShardedClusterBackup(ctx, groupId, clusterName, snapshotId).Execute()

Return One Sharded Cluster Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.GetShardedClusterBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetShardedClusterBackup`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetShardedClusterBackup`: DiskBackupShardedClusterSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetShardedClusterBackup`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiGetShardedClusterBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DiskBackupShardedClusterSnapshot**](DiskBackupShardedClusterSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupExportJobs

> PaginatedApiAtlasDiskBackupExportJob ListBackupExportJobs(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Snapshot Export Jobs


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListBackupExportJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListBackupExportJobs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListBackupExportJobs`: PaginatedApiAtlasDiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListBackupExportJobs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasDiskBackupExportJob**](PaginatedApiAtlasDiskBackupExportJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRestoreJobs

> PaginatedCloudBackupRestoreJob ListBackupRestoreJobs(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Restore Jobs for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListBackupRestoreJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListBackupRestoreJobs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListBackupRestoreJobs`: PaginatedCloudBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListBackupRestoreJobs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the restore jobs you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedCloudBackupRestoreJob**](PaginatedCloudBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportBuckets

> PaginatedBackupSnapshotExportBuckets ListExportBuckets(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Snapshot Export Buckets


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListExportBuckets(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListExportBuckets`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListExportBuckets`: PaginatedBackupSnapshotExportBuckets
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListExportBuckets`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExportBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedBackupSnapshotExportBuckets**](PaginatedBackupSnapshotExportBuckets.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2024-05-30+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplicaSetBackups

> PaginatedCloudBackupReplicaSet ListReplicaSetBackups(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Replica Set Cloud Backups


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListReplicaSetBackups(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListReplicaSetBackups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListReplicaSetBackups`: PaginatedCloudBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListReplicaSetBackups`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReplicaSetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedCloudBackupReplicaSet**](PaginatedCloudBackupReplicaSet.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerlessBackupRestoreJobs

> PaginatedApiAtlasServerlessBackupRestoreJob ListServerlessBackupRestoreJobs(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Restore Jobs for One Serverless Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListServerlessBackupRestoreJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListServerlessBackupRestoreJobs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListServerlessBackupRestoreJobs`: PaginatedApiAtlasServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListServerlessBackupRestoreJobs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerlessBackupRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasServerlessBackupRestoreJob**](PaginatedApiAtlasServerlessBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerlessBackups

> PaginatedApiAtlasServerlessBackupSnapshot ListServerlessBackups(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Snapshots of One Serverless Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListServerlessBackups(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListServerlessBackups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListServerlessBackups`: PaginatedApiAtlasServerlessBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListServerlessBackups`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerlessBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiAtlasServerlessBackupSnapshot**](PaginatedApiAtlasServerlessBackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShardedClusterBackups

> PaginatedCloudBackupShardedClusterSnapshot ListShardedClusterBackups(ctx, groupId, clusterName).Execute()

Return All Sharded Cluster Cloud Backups


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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

    resp, r, err := sdk.CloudBackupsApi.ListShardedClusterBackups(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListShardedClusterBackups`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListShardedClusterBackups`: PaginatedCloudBackupShardedClusterSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListShardedClusterBackups`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShardedClusterBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedCloudBackupShardedClusterSnapshot**](PaginatedCloudBackupShardedClusterSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TakeSnapshot

> DiskBackupSnapshot TakeSnapshot(ctx, groupId, clusterName, diskBackupOnDemandSnapshotRequest DiskBackupOnDemandSnapshotRequest).Execute()

Take One On-Demand Snapshot


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    diskBackupOnDemandSnapshotRequest := *openapiclient.NewDiskBackupOnDemandSnapshotRequest() // DiskBackupOnDemandSnapshotRequest | 

    resp, r, err := sdk.CloudBackupsApi.TakeSnapshot(context.Background(), groupId, clusterName, &diskBackupOnDemandSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.TakeSnapshot`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `TakeSnapshot`: DiskBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.TakeSnapshot`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **diskBackupOnDemandSnapshotRequest** | [**DiskBackupOnDemandSnapshotRequest**](DiskBackupOnDemandSnapshotRequest.md) | Takes one on-demand snapshot. | 

### Return type

[**DiskBackupSnapshot**](DiskBackupSnapshot.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupSchedule

> DiskBackupSnapshotSchedule20240805 UpdateBackupSchedule(ctx, groupId, clusterName, diskBackupSnapshotSchedule20240805 DiskBackupSnapshotSchedule20240805).Execute()

Update Cloud Backup Schedule for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    diskBackupSnapshotSchedule20240805 := *openapiclient.NewDiskBackupSnapshotSchedule20240805() // DiskBackupSnapshotSchedule20240805 | 

    resp, r, err := sdk.CloudBackupsApi.UpdateBackupSchedule(context.Background(), groupId, clusterName, &diskBackupSnapshotSchedule20240805).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateBackupSchedule`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateBackupSchedule`: DiskBackupSnapshotSchedule20240805
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateBackupSchedule`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **diskBackupSnapshotSchedule20240805** | [**DiskBackupSnapshotSchedule20240805**](DiskBackupSnapshotSchedule20240805.md) | Updates the cloud backup schedule for one cluster within the specified project.  **Note**: In the request body, provide only the fields that you want to update. | 

### Return type

[**DiskBackupSnapshotSchedule20240805**](DiskBackupSnapshotSchedule20240805.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataProtectionSettings

> DataProtectionSettings20231001 UpdateDataProtectionSettings(ctx, groupId, dataProtectionSettings20231001 DataProtectionSettings20231001).OverwriteBackupPolicies(overwriteBackupPolicies).Execute()

Update or enable the Backup Compliance Policy settings


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    dataProtectionSettings20231001 := *openapiclient.NewDataProtectionSettings20231001("AuthorizedEmail_example", "AuthorizedUserFirstName_example", "AuthorizedUserLastName_example") // DataProtectionSettings20231001 | 
    overwriteBackupPolicies := true // bool |  (optional) (default to true)

    resp, r, err := sdk.CloudBackupsApi.UpdateDataProtectionSettings(context.Background(), groupId, &dataProtectionSettings20231001).OverwriteBackupPolicies(overwriteBackupPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateDataProtectionSettings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateDataProtectionSettings`: DataProtectionSettings20231001
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateDataProtectionSettings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataProtectionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataProtectionSettings20231001** | [**DataProtectionSettings20231001**](DataProtectionSettings20231001.md) | The new Backup Compliance Policy settings. | 
 **overwriteBackupPolicies** | **bool** | Flag that indicates whether to overwrite non complying backup policies with the new data protection settings or not. | [default to true]

### Return type

[**DataProtectionSettings20231001**](DataProtectionSettings20231001.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-10-01+json
- **Accept**: application/vnd.atlas.2023-10-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshotRetention

> DiskBackupReplicaSet UpdateSnapshotRetention(ctx, groupId, clusterName, snapshotId, backupSnapshotRetention BackupSnapshotRetention).Execute()

Change Expiration Date for One Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312002/admin"
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
    backupSnapshotRetention := *openapiclient.NewBackupSnapshotRetention("RetentionUnit_example", int(5)) // BackupSnapshotRetention | 

    resp, r, err := sdk.CloudBackupsApi.UpdateSnapshotRetention(context.Background(), groupId, clusterName, snapshotId, &backupSnapshotRetention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateSnapshotRetention`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateSnapshotRetention`: DiskBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateSnapshotRetention`: %v (%v)\n", resp, r)
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

Other parameters are passed through a pointer to a apiUpdateSnapshotRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backupSnapshotRetention** | [**BackupSnapshotRetention**](BackupSnapshotRetention.md) | Changes the expiration date for one cloud backup snapshot for one cluster in the specified project. | 

### Return type

[**DiskBackupReplicaSet**](DiskBackupReplicaSet.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

