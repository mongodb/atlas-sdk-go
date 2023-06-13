# \CloudBackupsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBackupRestoreJob**](CloudBackupsApi.md#CancelBackupRestoreJob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
[**CreateBackupExportJob**](CloudBackupsApi.md#CreateBackupExportJob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Cloud Backup Snapshot Export Job
[**CreateBackupRestoreJob**](CloudBackupsApi.md#CreateBackupRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
[**CreateExportBucket**](CloudBackupsApi.md#CreateExportBucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
[**CreateServerlessBackupRestoreJob**](CloudBackupsApi.md#CreateServerlessBackupRestoreJob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
[**DeleteAllBackupSchedules**](CloudBackupsApi.md#DeleteAllBackupSchedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
[**DeleteExportBucket**](CloudBackupsApi.md#DeleteExportBucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
[**DeleteReplicaSetBackup**](CloudBackupsApi.md#DeleteReplicaSetBackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
[**DeleteShardedClusterBackup**](CloudBackupsApi.md#DeleteShardedClusterBackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
[**GetBackupExportJob**](CloudBackupsApi.md#GetBackupExportJob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Cloud Backup Snapshot Export Job
[**GetBackupRestoreJob**](CloudBackupsApi.md#GetBackupRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster
[**GetBackupSchedule**](CloudBackupsApi.md#GetBackupSchedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
[**GetDataProtectionSettings**](CloudBackupsApi.md#GetDataProtectionSettings) | **Get** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Return the Backup Compliance Policy settings
[**GetExportBucket**](CloudBackupsApi.md#GetExportBucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports
[**GetReplicaSetBackup**](CloudBackupsApi.md#GetReplicaSetBackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
[**GetServerlessBackup**](CloudBackupsApi.md#GetServerlessBackup) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
[**GetServerlessBackupRestoreJob**](CloudBackupsApi.md#GetServerlessBackupRestoreJob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
[**GetShardedClusterBackup**](CloudBackupsApi.md#GetShardedClusterBackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
[**ListBackupExportJobs**](CloudBackupsApi.md#ListBackupExportJobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Cloud Backup Snapshot Export Jobs
[**ListBackupRestoreJobs**](CloudBackupsApi.md#ListBackupRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
[**ListExportBuckets**](CloudBackupsApi.md#ListExportBuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports
[**ListReplicaSetBackups**](CloudBackupsApi.md#ListReplicaSetBackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
[**ListServerlessBackupRestoreJobs**](CloudBackupsApi.md#ListServerlessBackupRestoreJobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
[**ListServerlessBackups**](CloudBackupsApi.md#ListServerlessBackups) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
[**ListShardedClusterBackups**](CloudBackupsApi.md#ListShardedClusterBackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
[**TakeSnapshot**](CloudBackupsApi.md#TakeSnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot
[**UpdateBackupSchedule**](CloudBackupsApi.md#UpdateBackupSchedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster
[**UpdateDataProtectionSettings**](CloudBackupsApi.md#UpdateDataProtectionSettings) | **Put** /api/atlas/v2/groups/{groupId}/backupCompliancePolicy | Update or enable the Backup Compliance Policy settings
[**UpdateSnapshotRetention**](CloudBackupsApi.md#UpdateSnapshotRetention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup



## CancelBackupRestoreJob

> map[string]interface{} CancelBackupRestoreJob(ctx, groupId, clusterName, restoreJobId).Execute()

Cancel One Restore Job of One Cluster


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.CancelBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CancelBackupRestoreJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CancelBackupRestoreJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CancelBackupRestoreJob`: %v\n", resp)
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

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupExportJob

> DiskBackupExportJob CreateBackupExportJob(ctx, groupId, clusterName, diskBackupExportJobRequest DiskBackupExportJobRequest).Execute()

Create One Cloud Backup Snapshot Export Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    diskBackupExportJobRequest := *openapiclient.NewDiskBackupExportJobRequest("32b6e34b3d91647abb20e7b8", "32b6e34b3d91647abb20e7b8") // DiskBackupExportJobRequest | 

    resp, r, err := sdk.CloudBackupsApi.CreateBackupExportJob(context.Background(), groupId, clusterName, &diskBackupExportJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateBackupExportJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateBackupExportJob`: DiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateBackupExportJob`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupRestoreJob

> DiskBackupRestoreJob CreateBackupRestoreJob(ctx, groupId, clusterName, diskBackupRestoreJob DiskBackupRestoreJob).Execute()

Restore One Snapshot of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    diskBackupRestoreJob := *openapiclient.NewDiskBackupRestoreJob("DeliveryType_example", "TargetClusterName_example", "32b6e34b3d91647abb20e7b8") // DiskBackupRestoreJob | 

    resp, r, err := sdk.CloudBackupsApi.CreateBackupRestoreJob(context.Background(), groupId, clusterName, &diskBackupRestoreJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateBackupRestoreJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateBackupRestoreJob`: DiskBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateBackupRestoreJob`: %v\n", resp)
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


 **diskBackupRestoreJob** | [**DiskBackupRestoreJob**](DiskBackupRestoreJob.md) | Restores one snapshot of one cluster from the specified project. | 

### Return type

[**DiskBackupRestoreJob**](DiskBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExportBucket

> DiskBackupSnapshotAWSExportBucket CreateExportBucket(ctx, groupId, diskBackupSnapshotAWSExportBucket DiskBackupSnapshotAWSExportBucket).Execute()

Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    diskBackupSnapshotAWSExportBucket := *openapiclient.NewDiskBackupSnapshotAWSExportBucket() // DiskBackupSnapshotAWSExportBucket | 

    resp, r, err := sdk.CloudBackupsApi.CreateExportBucket(context.Background(), groupId, &diskBackupSnapshotAWSExportBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateExportBucket``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateExportBucket`: DiskBackupSnapshotAWSExportBucket
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateExportBucket`: %v\n", resp)
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

 **diskBackupSnapshotAWSExportBucket** | [**DiskBackupSnapshotAWSExportBucket**](DiskBackupSnapshotAWSExportBucket.md) | Grants MongoDB Cloud access to the specified AWS S3 bucket. | 

### Return type

[**DiskBackupSnapshotAWSExportBucket**](DiskBackupSnapshotAWSExportBucket.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    serverlessBackupRestoreJob := *openapiclient.NewServerlessBackupRestoreJob("DeliveryType_example", "TargetClusterName_example", "32b6e34b3d91647abb20e7b8") // ServerlessBackupRestoreJob | 

    resp, r, err := sdk.CloudBackupsApi.CreateServerlessBackupRestoreJob(context.Background(), groupId, clusterName, &serverlessBackupRestoreJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.CreateServerlessBackupRestoreJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateServerlessBackupRestoreJob`: ServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.CreateServerlessBackupRestoreJob`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllBackupSchedules

> DiskBackupSnapshotSchedule DeleteAllBackupSchedules(ctx, groupId, clusterName).Execute()

Remove All Cloud Backup Schedules


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.DeleteAllBackupSchedules(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteAllBackupSchedules``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteAllBackupSchedules`: DiskBackupSnapshotSchedule
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.DeleteAllBackupSchedules`: %v\n", resp)
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

[**DiskBackupSnapshotSchedule**](DiskBackupSnapshotSchedule.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExportBucket

> map[string]interface{} DeleteExportBucket(ctx, groupId, exportBucketId).Execute()

Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    exportBucketId := "exportBucketId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.DeleteExportBucket(context.Background(), groupId, exportBucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteExportBucket``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteExportBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.DeleteExportBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**exportBucketId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExportBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplicaSetBackup

> map[string]interface{} DeleteReplicaSetBackup(ctx, groupId, clusterName, snapshotId).Execute()

Remove One Replica Set Cloud Backup


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.DeleteReplicaSetBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteReplicaSetBackup``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteReplicaSetBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.DeleteReplicaSetBackup`: %v\n", resp)
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

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShardedClusterBackup

> map[string]interface{} DeleteShardedClusterBackup(ctx, groupId, clusterName, snapshotId).Execute()

Remove One Sharded Cluster Cloud Backup


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.DeleteShardedClusterBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.DeleteShardedClusterBackup``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `DeleteShardedClusterBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.DeleteShardedClusterBackup`: %v\n", resp)
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

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupExportJob

> DiskBackupExportJob GetBackupExportJob(ctx, groupId, clusterName, exportId).Execute()

Return One Cloud Backup Snapshot Export Job


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    exportId := "exportId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetBackupExportJob(context.Background(), groupId, clusterName, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupExportJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetBackupExportJob`: DiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupExportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**exportId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupRestoreJob

> DiskBackupRestoreJob GetBackupRestoreJob(ctx, groupId, clusterName, restoreJobId).Execute()

Return One Restore Job of One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupRestoreJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetBackupRestoreJob`: DiskBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupRestoreJob`: %v\n", resp)
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

[**DiskBackupRestoreJob**](DiskBackupRestoreJob.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupSchedule

> DiskBackupSnapshotSchedule GetBackupSchedule(ctx, groupId, clusterName).Execute()

Return One Cloud Backup Schedule


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetBackupSchedule(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetBackupSchedule``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetBackupSchedule`: DiskBackupSnapshotSchedule
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetBackupSchedule`: %v\n", resp)
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

[**DiskBackupSnapshotSchedule**](DiskBackupSnapshotSchedule.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataProtectionSettings

> DataProtectionSettings GetDataProtectionSettings(ctx, groupId).Execute()

Return the Backup Compliance Policy settings


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetDataProtectionSettings(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetDataProtectionSettings``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetDataProtectionSettings`: DataProtectionSettings
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetDataProtectionSettings`: %v\n", resp)
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

[**DataProtectionSettings**](DataProtectionSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportBucket

> DiskBackupSnapshotAWSExportBucket GetExportBucket(ctx, groupId, exportBucketId).Execute()

Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    exportBucketId := "exportBucketId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetExportBucket(context.Background(), groupId, exportBucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetExportBucket``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetExportBucket`: DiskBackupSnapshotAWSExportBucket
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetExportBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups. | 
**exportBucketId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DiskBackupSnapshotAWSExportBucket**](DiskBackupSnapshotAWSExportBucket.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetReplicaSetBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetReplicaSetBackup``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetReplicaSetBackup`: DiskBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetReplicaSetBackup`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetServerlessBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetServerlessBackup``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetServerlessBackup`: ServerlessBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetServerlessBackup`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    restoreJobId := "restoreJobId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetServerlessBackupRestoreJob(context.Background(), groupId, clusterName, restoreJobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetServerlessBackupRestoreJob``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetServerlessBackupRestoreJob`: ServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetServerlessBackupRestoreJob`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShardedClusterBackup

> DiskBackupShardedClusterSnapshot GetShardedClusterBackup(ctx, groupId, clusterName, snapshotId).Execute()

Return One Sharded Cluster Cloud Backup


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "snapshotId_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.GetShardedClusterBackup(context.Background(), groupId, clusterName, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.GetShardedClusterBackup``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetShardedClusterBackup`: DiskBackupShardedClusterSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.GetShardedClusterBackup`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupExportJobs

> PaginatedApiAtlasDiskBackupExportJob ListBackupExportJobs(ctx, groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Cloud Backup Snapshot Export Jobs


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListBackupExportJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListBackupExportJobs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListBackupExportJobs`: PaginatedApiAtlasDiskBackupExportJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListBackupExportJobs`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListBackupRestoreJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListBackupRestoreJobs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListBackupRestoreJobs`: PaginatedCloudBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListBackupRestoreJobs`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportBuckets

> PaginatedBackupSnapshotExportBucket ListExportBuckets(ctx, groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListExportBuckets(context.Background(), groupId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListExportBuckets``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListExportBuckets`: PaginatedBackupSnapshotExportBucket
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListExportBuckets`: %v\n", resp)
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

[**PaginatedBackupSnapshotExportBucket**](PaginatedBackupSnapshotExportBucket.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListReplicaSetBackups(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListReplicaSetBackups``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListReplicaSetBackups`: PaginatedCloudBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListReplicaSetBackups`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListServerlessBackupRestoreJobs(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListServerlessBackupRestoreJobs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListServerlessBackupRestoreJobs`: PaginatedApiAtlasServerlessBackupRestoreJob
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListServerlessBackupRestoreJobs`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.CloudBackupsApi.ListServerlessBackups(context.Background(), groupId, clusterName).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListServerlessBackups``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListServerlessBackups`: PaginatedApiAtlasServerlessBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListServerlessBackups`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShardedClusterBackups

> PaginatedCloudBackupShardedClusterSnapshot ListShardedClusterBackups(ctx, groupId, clusterName).Execute()

Return All Sharded Cluster Cloud Backups


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 

    resp, r, err := sdk.CloudBackupsApi.ListShardedClusterBackups(context.Background(), groupId, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ListShardedClusterBackups``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListShardedClusterBackups`: PaginatedCloudBackupShardedClusterSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ListShardedClusterBackups`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    diskBackupOnDemandSnapshotRequest := *openapiclient.NewDiskBackupOnDemandSnapshotRequest() // DiskBackupOnDemandSnapshotRequest | 

    resp, r, err := sdk.CloudBackupsApi.TakeSnapshot(context.Background(), groupId, clusterName, &diskBackupOnDemandSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.TakeSnapshot``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `TakeSnapshot`: DiskBackupSnapshot
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.TakeSnapshot`: %v\n", resp)
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
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupSchedule

> DiskBackupSnapshotSchedule UpdateBackupSchedule(ctx, groupId, clusterName, diskBackupSnapshotSchedule DiskBackupSnapshotSchedule).Execute()

Update Cloud Backup Schedule for One Cluster


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    diskBackupSnapshotSchedule := *openapiclient.NewDiskBackupSnapshotSchedule() // DiskBackupSnapshotSchedule | 

    resp, r, err := sdk.CloudBackupsApi.UpdateBackupSchedule(context.Background(), groupId, clusterName, &diskBackupSnapshotSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateBackupSchedule``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateBackupSchedule`: DiskBackupSnapshotSchedule
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateBackupSchedule`: %v\n", resp)
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


 **diskBackupSnapshotSchedule** | [**DiskBackupSnapshotSchedule**](DiskBackupSnapshotSchedule.md) | Updates the cloud backup schedule for one cluster within the specified project.  **Note**: In the request body, provide only the fields that you want to update. | 

### Return type

[**DiskBackupSnapshotSchedule**](DiskBackupSnapshotSchedule.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataProtectionSettings

> DataProtectionSettings UpdateDataProtectionSettings(ctx, groupId, dataProtectionSettings DataProtectionSettings).Execute()

Update or enable the Backup Compliance Policy settings


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    dataProtectionSettings := *openapiclient.NewDataProtectionSettings() // DataProtectionSettings | 

    resp, r, err := sdk.CloudBackupsApi.UpdateDataProtectionSettings(context.Background(), groupId, &dataProtectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateDataProtectionSettings``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateDataProtectionSettings`: DataProtectionSettings
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateDataProtectionSettings`: %v\n", resp)
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

 **dataProtectionSettings** | [**DataProtectionSettings**](DataProtectionSettings.md) | The new Backup Compliance Policy settings. | 

### Return type

[**DataProtectionSettings**](DataProtectionSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshotRetention

> DiskBackupReplicaSet UpdateSnapshotRetention(ctx, groupId, clusterName, snapshotId, snapshotRetention SnapshotRetention).Execute()

Change Expiration Date for One Cloud Backup


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    groupId := "32b6e34b3d91647abb20e7b8" // string | 
    clusterName := "clusterName_example" // string | 
    snapshotId := "snapshotId_example" // string | 
    snapshotRetention := *openapiclient.NewSnapshotRetention("RetentionUnit_example", int(5)) // SnapshotRetention | 

    resp, r, err := sdk.CloudBackupsApi.UpdateSnapshotRetention(context.Background(), groupId, clusterName, snapshotId, &snapshotRetention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.UpdateSnapshotRetention``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateSnapshotRetention`: DiskBackupReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.UpdateSnapshotRetention`: %v\n", resp)
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



 **snapshotRetention** | [**SnapshotRetention**](SnapshotRetention.md) | Changes the expiration date for one cloud backup snapshot for one cluster in the specified project. | 

### Return type

[**DiskBackupReplicaSet**](DiskBackupReplicaSet.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

