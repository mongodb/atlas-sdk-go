// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CloudBackupsApi interface {

	/*
		CancelBackupRestoreJob Cancel One Restore Job of One Cluster

		[experimental] Cancels one cloud backup restore job of one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to remove.
		@return CancelBackupRestoreJobApiRequest
	*/
	CancelBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) CancelBackupRestoreJobApiRequest
	/*
		CancelBackupRestoreJob Cancel One Restore Job of One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CancelBackupRestoreJobApiParams - Parameters for the request
		@return CancelBackupRestoreJobApiRequest
	*/
	CancelBackupRestoreJobWithParams(ctx context.Context, args *CancelBackupRestoreJobApiParams) CancelBackupRestoreJobApiRequest

	// Interface only available internally
	cancelBackupRestoreJobExecute(r CancelBackupRestoreJobApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		CreateBackupExportJob Create One Cloud Backup Snapshot Export Job

		Exports one backup snapshot for dedicated Atlas cluster using Cloud Backups to an AWS bucket. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return CreateBackupExportJobApiRequest
	*/
	CreateBackupExportJob(ctx context.Context, groupId string, clusterName string, diskBackupExportJobRequest *DiskBackupExportJobRequest) CreateBackupExportJobApiRequest
	/*
		CreateBackupExportJob Create One Cloud Backup Snapshot Export Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateBackupExportJobApiParams - Parameters for the request
		@return CreateBackupExportJobApiRequest
	*/
	CreateBackupExportJobWithParams(ctx context.Context, args *CreateBackupExportJobApiParams) CreateBackupExportJobApiRequest

	// Interface only available internally
	createBackupExportJobExecute(r CreateBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error)

	/*
		CreateBackupRestoreJob Restore One Snapshot of One Cluster

		Restores one snapshot of one cluster from the specified project. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	 To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return CreateBackupRestoreJobApiRequest
	*/
	CreateBackupRestoreJob(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotRestoreJob *DiskBackupSnapshotRestoreJob) CreateBackupRestoreJobApiRequest
	/*
		CreateBackupRestoreJob Restore One Snapshot of One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateBackupRestoreJobApiParams - Parameters for the request
		@return CreateBackupRestoreJobApiRequest
	*/
	CreateBackupRestoreJobWithParams(ctx context.Context, args *CreateBackupRestoreJobApiParams) CreateBackupRestoreJobApiRequest

	// Interface only available internally
	createBackupRestoreJobExecute(r CreateBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error)

	/*
		CreateExportBucket Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports

		Grants MongoDB Cloud access to the specified AWS S3 bucket. This enables this bucket to receive Atlas Cloud Backup snapshots. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateExportBucketApiRequest
	*/
	CreateExportBucket(ctx context.Context, groupId string, diskBackupSnapshotAWSExportBucket *DiskBackupSnapshotAWSExportBucket) CreateExportBucketApiRequest
	/*
		CreateExportBucket Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateExportBucketApiParams - Parameters for the request
		@return CreateExportBucketApiRequest
	*/
	CreateExportBucketWithParams(ctx context.Context, args *CreateExportBucketApiParams) CreateExportBucketApiRequest

	// Interface only available internally
	createExportBucketExecute(r CreateExportBucketApiRequest) (*DiskBackupSnapshotAWSExportBucket, *http.Response, error)

	/*
		CreateServerlessBackupRestoreJob Restore One Snapshot of One Serverless Instance

		Restores one snapshot of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the serverless instance whose snapshot you want to restore.
		@return CreateServerlessBackupRestoreJobApiRequest
	*/
	CreateServerlessBackupRestoreJob(ctx context.Context, groupId string, clusterName string, serverlessBackupRestoreJob *ServerlessBackupRestoreJob) CreateServerlessBackupRestoreJobApiRequest
	/*
		CreateServerlessBackupRestoreJob Restore One Snapshot of One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServerlessBackupRestoreJobApiParams - Parameters for the request
		@return CreateServerlessBackupRestoreJobApiRequest
	*/
	CreateServerlessBackupRestoreJobWithParams(ctx context.Context, args *CreateServerlessBackupRestoreJobApiParams) CreateServerlessBackupRestoreJobApiRequest

	// Interface only available internally
	createServerlessBackupRestoreJobExecute(r CreateServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error)

	/*
		DeleteAllBackupSchedules Remove All Cloud Backup Schedules

		Removes all cloud backup schedules for the specified cluster. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Atlas Admin role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return DeleteAllBackupSchedulesApiRequest
	*/
	DeleteAllBackupSchedules(ctx context.Context, groupId string, clusterName string) DeleteAllBackupSchedulesApiRequest
	/*
		DeleteAllBackupSchedules Remove All Cloud Backup Schedules


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAllBackupSchedulesApiParams - Parameters for the request
		@return DeleteAllBackupSchedulesApiRequest
	*/
	DeleteAllBackupSchedulesWithParams(ctx context.Context, args *DeleteAllBackupSchedulesApiParams) DeleteAllBackupSchedulesApiRequest

	// Interface only available internally
	deleteAllBackupSchedulesExecute(r DeleteAllBackupSchedulesApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error)

	/*
		DeleteExportBucket Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports

		Revoke MongoDB Cloud access to the specified AWS S3 bucket. This prevents this bucket to receive Atlas Cloud Backup snapshots. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param exportBucketId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
		@return DeleteExportBucketApiRequest
	*/
	DeleteExportBucket(ctx context.Context, groupId string, exportBucketId string) DeleteExportBucketApiRequest
	/*
		DeleteExportBucket Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteExportBucketApiParams - Parameters for the request
		@return DeleteExportBucketApiRequest
	*/
	DeleteExportBucketWithParams(ctx context.Context, args *DeleteExportBucketApiParams) DeleteExportBucketApiRequest

	// Interface only available internally
	deleteExportBucketExecute(r DeleteExportBucketApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteReplicaSetBackup Remove One Replica Set Cloud Backup

		Removes the specified snapshot. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return DeleteReplicaSetBackupApiRequest
	*/
	DeleteReplicaSetBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteReplicaSetBackupApiRequest
	/*
		DeleteReplicaSetBackup Remove One Replica Set Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteReplicaSetBackupApiParams - Parameters for the request
		@return DeleteReplicaSetBackupApiRequest
	*/
	DeleteReplicaSetBackupWithParams(ctx context.Context, args *DeleteReplicaSetBackupApiParams) DeleteReplicaSetBackupApiRequest

	// Interface only available internally
	deleteReplicaSetBackupExecute(r DeleteReplicaSetBackupApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteShardedClusterBackup Remove One Sharded Cluster Cloud Backup

		[experimental] Removes one snapshot of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return DeleteShardedClusterBackupApiRequest
	*/
	DeleteShardedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteShardedClusterBackupApiRequest
	/*
		DeleteShardedClusterBackup Remove One Sharded Cluster Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteShardedClusterBackupApiParams - Parameters for the request
		@return DeleteShardedClusterBackupApiRequest
	*/
	DeleteShardedClusterBackupWithParams(ctx context.Context, args *DeleteShardedClusterBackupApiParams) DeleteShardedClusterBackupApiRequest

	// Interface only available internally
	deleteShardedClusterBackupExecute(r DeleteShardedClusterBackupApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetBackupExportJob Return One Cloud Backup Snapshot Export Job

		Returns one Cloud Backup snapshot export job associated with the specified Atlas cluster. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param exportId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
		@return GetBackupExportJobApiRequest
	*/
	GetBackupExportJob(ctx context.Context, groupId string, clusterName string, exportId string) GetBackupExportJobApiRequest
	/*
		GetBackupExportJob Return One Cloud Backup Snapshot Export Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetBackupExportJobApiParams - Parameters for the request
		@return GetBackupExportJobApiRequest
	*/
	GetBackupExportJobWithParams(ctx context.Context, args *GetBackupExportJobApiParams) GetBackupExportJobApiRequest

	// Interface only available internally
	getBackupExportJobExecute(r GetBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error)

	/*
		GetBackupRestoreJob Return One Restore Job of One Cluster

		Returns one cloud backup restore job for one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the restore jobs you want to return.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
		@return GetBackupRestoreJobApiRequest
	*/
	GetBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) GetBackupRestoreJobApiRequest
	/*
		GetBackupRestoreJob Return One Restore Job of One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetBackupRestoreJobApiParams - Parameters for the request
		@return GetBackupRestoreJobApiRequest
	*/
	GetBackupRestoreJobWithParams(ctx context.Context, args *GetBackupRestoreJobApiParams) GetBackupRestoreJobApiRequest

	// Interface only available internally
	getBackupRestoreJobExecute(r GetBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error)

	/*
		GetBackupSchedule Return One Cloud Backup Schedule

		Returns the cloud backup schedule for the specified cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetBackupScheduleApiRequest
	*/
	GetBackupSchedule(ctx context.Context, groupId string, clusterName string) GetBackupScheduleApiRequest
	/*
		GetBackupSchedule Return One Cloud Backup Schedule


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetBackupScheduleApiParams - Parameters for the request
		@return GetBackupScheduleApiRequest
	*/
	GetBackupScheduleWithParams(ctx context.Context, args *GetBackupScheduleApiParams) GetBackupScheduleApiRequest

	// Interface only available internally
	getBackupScheduleExecute(r GetBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error)

	/*
		GetDataProtectionSettings Return the Backup Compliance Policy settings

		[experimental] Returns the Backup Compliance Policy settings with the specified project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetDataProtectionSettingsApiRequest
	*/
	GetDataProtectionSettings(ctx context.Context, groupId string) GetDataProtectionSettingsApiRequest
	/*
		GetDataProtectionSettings Return the Backup Compliance Policy settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDataProtectionSettingsApiParams - Parameters for the request
		@return GetDataProtectionSettingsApiRequest
	*/
	GetDataProtectionSettingsWithParams(ctx context.Context, args *GetDataProtectionSettingsApiParams) GetDataProtectionSettingsApiRequest

	// Interface only available internally
	getDataProtectionSettingsExecute(r GetDataProtectionSettingsApiRequest) (*DataProtectionSettings, *http.Response, error)

	/*
		GetExportBucket Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports

		Returns one AWS S3 bucket associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param exportBucketId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
		@return GetExportBucketApiRequest
	*/
	GetExportBucket(ctx context.Context, groupId string, exportBucketId string) GetExportBucketApiRequest
	/*
		GetExportBucket Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetExportBucketApiParams - Parameters for the request
		@return GetExportBucketApiRequest
	*/
	GetExportBucketWithParams(ctx context.Context, args *GetExportBucketApiParams) GetExportBucketApiRequest

	// Interface only available internally
	getExportBucketExecute(r GetExportBucketApiRequest) (*DiskBackupSnapshotAWSExportBucket, *http.Response, error)

	/*
		GetReplicaSetBackup Return One Replica Set Cloud Backup

		Returns one snapshot from the specified cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return GetReplicaSetBackupApiRequest
	*/
	GetReplicaSetBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetReplicaSetBackupApiRequest
	/*
		GetReplicaSetBackup Return One Replica Set Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetReplicaSetBackupApiParams - Parameters for the request
		@return GetReplicaSetBackupApiRequest
	*/
	GetReplicaSetBackupWithParams(ctx context.Context, args *GetReplicaSetBackupApiParams) GetReplicaSetBackupApiRequest

	// Interface only available internally
	getReplicaSetBackupExecute(r GetReplicaSetBackupApiRequest) (*DiskBackupReplicaSet, *http.Response, error)

	/*
		GetServerlessBackup Return One Snapshot of One Serverless Instance

		Returns one snapshot of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the serverless instance.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return GetServerlessBackupApiRequest
	*/
	GetServerlessBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetServerlessBackupApiRequest
	/*
		GetServerlessBackup Return One Snapshot of One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServerlessBackupApiParams - Parameters for the request
		@return GetServerlessBackupApiRequest
	*/
	GetServerlessBackupWithParams(ctx context.Context, args *GetServerlessBackupApiParams) GetServerlessBackupApiRequest

	// Interface only available internally
	getServerlessBackupExecute(r GetServerlessBackupApiRequest) (*ServerlessBackupSnapshot, *http.Response, error)

	/*
		GetServerlessBackupRestoreJob Return One Restore Job for One Serverless Instance

		Returns one restore job for one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the serverless instance.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
		@return GetServerlessBackupRestoreJobApiRequest
	*/
	GetServerlessBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) GetServerlessBackupRestoreJobApiRequest
	/*
		GetServerlessBackupRestoreJob Return One Restore Job for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServerlessBackupRestoreJobApiParams - Parameters for the request
		@return GetServerlessBackupRestoreJobApiRequest
	*/
	GetServerlessBackupRestoreJobWithParams(ctx context.Context, args *GetServerlessBackupRestoreJobApiParams) GetServerlessBackupRestoreJobApiRequest

	// Interface only available internally
	getServerlessBackupRestoreJobExecute(r GetServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error)

	/*
		GetShardedClusterBackup Return One Sharded Cluster Cloud Backup

		[experimental] Returns one snapshot of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return GetShardedClusterBackupApiRequest
	*/
	GetShardedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetShardedClusterBackupApiRequest
	/*
		GetShardedClusterBackup Return One Sharded Cluster Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetShardedClusterBackupApiParams - Parameters for the request
		@return GetShardedClusterBackupApiRequest
	*/
	GetShardedClusterBackupWithParams(ctx context.Context, args *GetShardedClusterBackupApiParams) GetShardedClusterBackupApiRequest

	// Interface only available internally
	getShardedClusterBackupExecute(r GetShardedClusterBackupApiRequest) (*DiskBackupShardedClusterSnapshot, *http.Response, error)

	/*
		ListBackupExportJobs Return All Cloud Backup Snapshot Export Jobs

		Returns all Cloud Backup snapshot export jobs associated with the specified Atlas cluster. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListBackupExportJobsApiRequest
	*/
	ListBackupExportJobs(ctx context.Context, groupId string, clusterName string) ListBackupExportJobsApiRequest
	/*
		ListBackupExportJobs Return All Cloud Backup Snapshot Export Jobs


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListBackupExportJobsApiParams - Parameters for the request
		@return ListBackupExportJobsApiRequest
	*/
	ListBackupExportJobsWithParams(ctx context.Context, args *ListBackupExportJobsApiParams) ListBackupExportJobsApiRequest

	// Interface only available internally
	listBackupExportJobsExecute(r ListBackupExportJobsApiRequest) (*PaginatedApiAtlasDiskBackupExportJob, *http.Response, error)

	/*
		ListBackupRestoreJobs Return All Restore Jobs for One Cluster

		Returns all cloud backup restore jobs for one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the restore jobs you want to return.
		@return ListBackupRestoreJobsApiRequest
	*/
	ListBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListBackupRestoreJobsApiRequest
	/*
		ListBackupRestoreJobs Return All Restore Jobs for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListBackupRestoreJobsApiParams - Parameters for the request
		@return ListBackupRestoreJobsApiRequest
	*/
	ListBackupRestoreJobsWithParams(ctx context.Context, args *ListBackupRestoreJobsApiParams) ListBackupRestoreJobsApiRequest

	// Interface only available internally
	listBackupRestoreJobsExecute(r ListBackupRestoreJobsApiRequest) (*PaginatedCloudBackupRestoreJob, *http.Response, error)

	/*
		ListExportBuckets Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports

		Returns all AWS S3 buckets associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListExportBucketsApiRequest
	*/
	ListExportBuckets(ctx context.Context, groupId string) ListExportBucketsApiRequest
	/*
		ListExportBuckets Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListExportBucketsApiParams - Parameters for the request
		@return ListExportBucketsApiRequest
	*/
	ListExportBucketsWithParams(ctx context.Context, args *ListExportBucketsApiParams) ListExportBucketsApiRequest

	// Interface only available internally
	listExportBucketsExecute(r ListExportBucketsApiRequest) (*PaginatedBackupSnapshotExportBucket, *http.Response, error)

	/*
		ListReplicaSetBackups Return All Replica Set Cloud Backups

		Returns all snapshots of one cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListReplicaSetBackupsApiRequest
	*/
	ListReplicaSetBackups(ctx context.Context, groupId string, clusterName string) ListReplicaSetBackupsApiRequest
	/*
		ListReplicaSetBackups Return All Replica Set Cloud Backups


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListReplicaSetBackupsApiParams - Parameters for the request
		@return ListReplicaSetBackupsApiRequest
	*/
	ListReplicaSetBackupsWithParams(ctx context.Context, args *ListReplicaSetBackupsApiParams) ListReplicaSetBackupsApiRequest

	// Interface only available internally
	listReplicaSetBackupsExecute(r ListReplicaSetBackupsApiRequest) (*PaginatedCloudBackupReplicaSet, *http.Response, error)

	/*
		ListServerlessBackupRestoreJobs Return All Restore Jobs for One Serverless Instance

		Returns all restore jobs for one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the serverless instance.
		@return ListServerlessBackupRestoreJobsApiRequest
	*/
	ListServerlessBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListServerlessBackupRestoreJobsApiRequest
	/*
		ListServerlessBackupRestoreJobs Return All Restore Jobs for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServerlessBackupRestoreJobsApiParams - Parameters for the request
		@return ListServerlessBackupRestoreJobsApiRequest
	*/
	ListServerlessBackupRestoreJobsWithParams(ctx context.Context, args *ListServerlessBackupRestoreJobsApiParams) ListServerlessBackupRestoreJobsApiRequest

	// Interface only available internally
	listServerlessBackupRestoreJobsExecute(r ListServerlessBackupRestoreJobsApiRequest) (*PaginatedApiAtlasServerlessBackupRestoreJob, *http.Response, error)

	/*
		ListServerlessBackups Return All Snapshots of One Serverless Instance

		Returns all snapshots of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the serverless instance.
		@return ListServerlessBackupsApiRequest
	*/
	ListServerlessBackups(ctx context.Context, groupId string, clusterName string) ListServerlessBackupsApiRequest
	/*
		ListServerlessBackups Return All Snapshots of One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServerlessBackupsApiParams - Parameters for the request
		@return ListServerlessBackupsApiRequest
	*/
	ListServerlessBackupsWithParams(ctx context.Context, args *ListServerlessBackupsApiParams) ListServerlessBackupsApiRequest

	// Interface only available internally
	listServerlessBackupsExecute(r ListServerlessBackupsApiRequest) (*PaginatedApiAtlasServerlessBackupSnapshot, *http.Response, error)

	/*
		ListShardedClusterBackups Return All Sharded Cluster Cloud Backups

		[experimental] Returns all snapshots of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListShardedClusterBackupsApiRequest
	*/
	ListShardedClusterBackups(ctx context.Context, groupId string, clusterName string) ListShardedClusterBackupsApiRequest
	/*
		ListShardedClusterBackups Return All Sharded Cluster Cloud Backups


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListShardedClusterBackupsApiParams - Parameters for the request
		@return ListShardedClusterBackupsApiRequest
	*/
	ListShardedClusterBackupsWithParams(ctx context.Context, args *ListShardedClusterBackupsApiParams) ListShardedClusterBackupsApiRequest

	// Interface only available internally
	listShardedClusterBackupsExecute(r ListShardedClusterBackupsApiRequest) (*PaginatedCloudBackupShardedClusterSnapshot, *http.Response, error)

	/*
		TakeSnapshot Take One On-Demand Snapshot

		Takes one on-demand snapshot for the specified cluster. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	 To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return TakeSnapshotApiRequest
	*/
	TakeSnapshot(ctx context.Context, groupId string, clusterName string, diskBackupOnDemandSnapshotRequest *DiskBackupOnDemandSnapshotRequest) TakeSnapshotApiRequest
	/*
		TakeSnapshot Take One On-Demand Snapshot


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param TakeSnapshotApiParams - Parameters for the request
		@return TakeSnapshotApiRequest
	*/
	TakeSnapshotWithParams(ctx context.Context, args *TakeSnapshotApiParams) TakeSnapshotApiRequest

	// Interface only available internally
	takeSnapshotExecute(r TakeSnapshotApiRequest) (*DiskBackupSnapshot, *http.Response, error)

	/*
		UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster

		Updates the cloud backup schedule for one cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return UpdateBackupScheduleApiRequest
	*/
	UpdateBackupSchedule(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotSchedule *DiskBackupSnapshotSchedule) UpdateBackupScheduleApiRequest
	/*
		UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateBackupScheduleApiParams - Parameters for the request
		@return UpdateBackupScheduleApiRequest
	*/
	UpdateBackupScheduleWithParams(ctx context.Context, args *UpdateBackupScheduleApiParams) UpdateBackupScheduleApiRequest

	// Interface only available internally
	updateBackupScheduleExecute(r UpdateBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error)

	/*
		UpdateDataProtectionSettings Update or enable the Backup Compliance Policy settings

		[experimental] Updates the Backup Compliance Policy settings for the specified project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpdateDataProtectionSettingsApiRequest
	*/
	UpdateDataProtectionSettings(ctx context.Context, groupId string, dataProtectionSettings *DataProtectionSettings) UpdateDataProtectionSettingsApiRequest
	/*
		UpdateDataProtectionSettings Update or enable the Backup Compliance Policy settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateDataProtectionSettingsApiParams - Parameters for the request
		@return UpdateDataProtectionSettingsApiRequest
	*/
	UpdateDataProtectionSettingsWithParams(ctx context.Context, args *UpdateDataProtectionSettingsApiParams) UpdateDataProtectionSettingsApiRequest

	// Interface only available internally
	updateDataProtectionSettingsExecute(r UpdateDataProtectionSettingsApiRequest) (*DataProtectionSettings, *http.Response, error)

	/*
		UpdateSnapshotRetention Change Expiration Date for One Cloud Backup

		[experimental] Changes the expiration date for one cloud backup snapshot for one cluster in the specified project. This resource doesn't require the API Key to have an Access List.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return UpdateSnapshotRetentionApiRequest
	*/
	UpdateSnapshotRetention(ctx context.Context, groupId string, clusterName string, snapshotId string, backupSnapshotRetention *BackupSnapshotRetention) UpdateSnapshotRetentionApiRequest
	/*
		UpdateSnapshotRetention Change Expiration Date for One Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateSnapshotRetentionApiParams - Parameters for the request
		@return UpdateSnapshotRetentionApiRequest
	*/
	UpdateSnapshotRetentionWithParams(ctx context.Context, args *UpdateSnapshotRetentionApiParams) UpdateSnapshotRetentionApiRequest

	// Interface only available internally
	updateSnapshotRetentionExecute(r UpdateSnapshotRetentionApiRequest) (*DiskBackupReplicaSet, *http.Response, error)
}

// CloudBackupsApiService CloudBackupsApi service
type CloudBackupsApiService service

type CancelBackupRestoreJobApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	restoreJobId string
}

type CancelBackupRestoreJobApiParams struct {
	GroupId      string
	ClusterName  string
	RestoreJobId string
}

func (a *CloudBackupsApiService) CancelBackupRestoreJobWithParams(ctx context.Context, args *CancelBackupRestoreJobApiParams) CancelBackupRestoreJobApiRequest {
	return CancelBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		restoreJobId: args.RestoreJobId,
	}
}

func (r CancelBackupRestoreJobApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.cancelBackupRestoreJobExecute(r)
}

/*
CancelBackupRestoreJob Cancel One Restore Job of One Cluster

[experimental] Cancels one cloud backup restore job of one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to remove.
	@return CancelBackupRestoreJobApiRequest
*/
func (a *CloudBackupsApiService) CancelBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) CancelBackupRestoreJobApiRequest {
	return CancelBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		clusterName:  clusterName,
		restoreJobId: restoreJobId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CloudBackupsApiService) cancelBackupRestoreJobExecute(r CancelBackupRestoreJobApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CancelBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(parameterValueToString(r.restoreJobId, "restoreJobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.restoreJobId) < 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have at least 24 elements")
	}
	if strlen(r.restoreJobId) > 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateBackupExportJobApiRequest struct {
	ctx                        context.Context
	ApiService                 CloudBackupsApi
	groupId                    string
	clusterName                string
	diskBackupExportJobRequest *DiskBackupExportJobRequest
}

type CreateBackupExportJobApiParams struct {
	GroupId                    string
	ClusterName                string
	DiskBackupExportJobRequest *DiskBackupExportJobRequest
}

func (a *CloudBackupsApiService) CreateBackupExportJobWithParams(ctx context.Context, args *CreateBackupExportJobApiParams) CreateBackupExportJobApiRequest {
	return CreateBackupExportJobApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		diskBackupExportJobRequest: args.DiskBackupExportJobRequest,
	}
}

func (r CreateBackupExportJobApiRequest) Execute() (*DiskBackupExportJob, *http.Response, error) {
	return r.ApiService.createBackupExportJobExecute(r)
}

/*
CreateBackupExportJob Create One Cloud Backup Snapshot Export Job

Exports one backup snapshot for dedicated Atlas cluster using Cloud Backups to an AWS bucket. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return CreateBackupExportJobApiRequest
*/
func (a *CloudBackupsApiService) CreateBackupExportJob(ctx context.Context, groupId string, clusterName string, diskBackupExportJobRequest *DiskBackupExportJobRequest) CreateBackupExportJobApiRequest {
	return CreateBackupExportJobApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		diskBackupExportJobRequest: diskBackupExportJobRequest,
	}
}

// Execute executes the request
//
//	@return DiskBackupExportJob
func (a *CloudBackupsApiService) createBackupExportJobExecute(r CreateBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateBackupExportJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.diskBackupExportJobRequest == nil {
		return localVarReturnValue, nil, reportError("diskBackupExportJobRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupExportJobRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateBackupRestoreJobApiRequest struct {
	ctx                          context.Context
	ApiService                   CloudBackupsApi
	groupId                      string
	clusterName                  string
	diskBackupSnapshotRestoreJob *DiskBackupSnapshotRestoreJob
}

type CreateBackupRestoreJobApiParams struct {
	GroupId                      string
	ClusterName                  string
	DiskBackupSnapshotRestoreJob *DiskBackupSnapshotRestoreJob
}

func (a *CloudBackupsApiService) CreateBackupRestoreJobWithParams(ctx context.Context, args *CreateBackupRestoreJobApiParams) CreateBackupRestoreJobApiRequest {
	return CreateBackupRestoreJobApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      args.GroupId,
		clusterName:                  args.ClusterName,
		diskBackupSnapshotRestoreJob: args.DiskBackupSnapshotRestoreJob,
	}
}

func (r CreateBackupRestoreJobApiRequest) Execute() (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	return r.ApiService.createBackupRestoreJobExecute(r)
}

/*
CreateBackupRestoreJob Restore One Snapshot of One Cluster

Restores one snapshot of one cluster from the specified project. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return CreateBackupRestoreJobApiRequest
*/
func (a *CloudBackupsApiService) CreateBackupRestoreJob(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotRestoreJob *DiskBackupSnapshotRestoreJob) CreateBackupRestoreJobApiRequest {
	return CreateBackupRestoreJobApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      groupId,
		clusterName:                  clusterName,
		diskBackupSnapshotRestoreJob: diskBackupSnapshotRestoreJob,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotRestoreJob
func (a *CloudBackupsApiService) createBackupRestoreJobExecute(r CreateBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.diskBackupSnapshotRestoreJob == nil {
		return localVarReturnValue, nil, reportError("diskBackupSnapshotRestoreJob is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupSnapshotRestoreJob
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateExportBucketApiRequest struct {
	ctx                               context.Context
	ApiService                        CloudBackupsApi
	groupId                           string
	diskBackupSnapshotAWSExportBucket *DiskBackupSnapshotAWSExportBucket
}

type CreateExportBucketApiParams struct {
	GroupId                           string
	DiskBackupSnapshotAWSExportBucket *DiskBackupSnapshotAWSExportBucket
}

func (a *CloudBackupsApiService) CreateExportBucketWithParams(ctx context.Context, args *CreateExportBucketApiParams) CreateExportBucketApiRequest {
	return CreateExportBucketApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           args.GroupId,
		diskBackupSnapshotAWSExportBucket: args.DiskBackupSnapshotAWSExportBucket,
	}
}

func (r CreateExportBucketApiRequest) Execute() (*DiskBackupSnapshotAWSExportBucket, *http.Response, error) {
	return r.ApiService.createExportBucketExecute(r)
}

/*
CreateExportBucket Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports

Grants MongoDB Cloud access to the specified AWS S3 bucket. This enables this bucket to receive Atlas Cloud Backup snapshots. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateExportBucketApiRequest
*/
func (a *CloudBackupsApiService) CreateExportBucket(ctx context.Context, groupId string, diskBackupSnapshotAWSExportBucket *DiskBackupSnapshotAWSExportBucket) CreateExportBucketApiRequest {
	return CreateExportBucketApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           groupId,
		diskBackupSnapshotAWSExportBucket: diskBackupSnapshotAWSExportBucket,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotAWSExportBucket
func (a *CloudBackupsApiService) createExportBucketExecute(r CreateExportBucketApiRequest) (*DiskBackupSnapshotAWSExportBucket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotAWSExportBucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateExportBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.diskBackupSnapshotAWSExportBucket == nil {
		return localVarReturnValue, nil, reportError("diskBackupSnapshotAWSExportBucket is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupSnapshotAWSExportBucket
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateServerlessBackupRestoreJobApiRequest struct {
	ctx                        context.Context
	ApiService                 CloudBackupsApi
	groupId                    string
	clusterName                string
	serverlessBackupRestoreJob *ServerlessBackupRestoreJob
}

type CreateServerlessBackupRestoreJobApiParams struct {
	GroupId                    string
	ClusterName                string
	ServerlessBackupRestoreJob *ServerlessBackupRestoreJob
}

func (a *CloudBackupsApiService) CreateServerlessBackupRestoreJobWithParams(ctx context.Context, args *CreateServerlessBackupRestoreJobApiParams) CreateServerlessBackupRestoreJobApiRequest {
	return CreateServerlessBackupRestoreJobApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		serverlessBackupRestoreJob: args.ServerlessBackupRestoreJob,
	}
}

func (r CreateServerlessBackupRestoreJobApiRequest) Execute() (*ServerlessBackupRestoreJob, *http.Response, error) {
	return r.ApiService.createServerlessBackupRestoreJobExecute(r)
}

/*
CreateServerlessBackupRestoreJob Restore One Snapshot of One Serverless Instance

Restores one snapshot of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the serverless instance whose snapshot you want to restore.
	@return CreateServerlessBackupRestoreJobApiRequest
*/
func (a *CloudBackupsApiService) CreateServerlessBackupRestoreJob(ctx context.Context, groupId string, clusterName string, serverlessBackupRestoreJob *ServerlessBackupRestoreJob) CreateServerlessBackupRestoreJobApiRequest {
	return CreateServerlessBackupRestoreJobApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		serverlessBackupRestoreJob: serverlessBackupRestoreJob,
	}
}

// Execute executes the request
//
//	@return ServerlessBackupRestoreJob
func (a *CloudBackupsApiService) createServerlessBackupRestoreJobExecute(r CreateServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateServerlessBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.serverlessBackupRestoreJob == nil {
		return localVarReturnValue, nil, reportError("serverlessBackupRestoreJob is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serverlessBackupRestoreJob
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteAllBackupSchedulesApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
}

type DeleteAllBackupSchedulesApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *CloudBackupsApiService) DeleteAllBackupSchedulesWithParams(ctx context.Context, args *DeleteAllBackupSchedulesApiParams) DeleteAllBackupSchedulesApiRequest {
	return DeleteAllBackupSchedulesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r DeleteAllBackupSchedulesApiRequest) Execute() (*DiskBackupSnapshotSchedule, *http.Response, error) {
	return r.ApiService.deleteAllBackupSchedulesExecute(r)
}

/*
DeleteAllBackupSchedules Remove All Cloud Backup Schedules

Removes all cloud backup schedules for the specified cluster. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Atlas Admin role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return DeleteAllBackupSchedulesApiRequest
*/
func (a *CloudBackupsApiService) DeleteAllBackupSchedules(ctx context.Context, groupId string, clusterName string) DeleteAllBackupSchedulesApiRequest {
	return DeleteAllBackupSchedulesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotSchedule
func (a *CloudBackupsApiService) deleteAllBackupSchedulesExecute(r DeleteAllBackupSchedulesApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteAllBackupSchedules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteExportBucketApiRequest struct {
	ctx            context.Context
	ApiService     CloudBackupsApi
	groupId        string
	exportBucketId string
}

type DeleteExportBucketApiParams struct {
	GroupId        string
	ExportBucketId string
}

func (a *CloudBackupsApiService) DeleteExportBucketWithParams(ctx context.Context, args *DeleteExportBucketApiParams) DeleteExportBucketApiRequest {
	return DeleteExportBucketApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		exportBucketId: args.ExportBucketId,
	}
}

func (r DeleteExportBucketApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteExportBucketExecute(r)
}

/*
DeleteExportBucket Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports

Revoke MongoDB Cloud access to the specified AWS S3 bucket. This prevents this bucket to receive Atlas Cloud Backup snapshots. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param exportBucketId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
	@return DeleteExportBucketApiRequest
*/
func (a *CloudBackupsApiService) DeleteExportBucket(ctx context.Context, groupId string, exportBucketId string) DeleteExportBucketApiRequest {
	return DeleteExportBucketApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		exportBucketId: exportBucketId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CloudBackupsApiService) deleteExportBucketExecute(r DeleteExportBucketApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteExportBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exportBucketId"+"}", url.PathEscape(parameterValueToString(r.exportBucketId, "exportBucketId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.exportBucketId) < 3 {
		return localVarReturnValue, nil, reportError("exportBucketId must have at least 3 elements")
	}
	if strlen(r.exportBucketId) > 63 {
		return localVarReturnValue, nil, reportError("exportBucketId must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteReplicaSetBackupApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type DeleteReplicaSetBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *CloudBackupsApiService) DeleteReplicaSetBackupWithParams(ctx context.Context, args *DeleteReplicaSetBackupApiParams) DeleteReplicaSetBackupApiRequest {
	return DeleteReplicaSetBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r DeleteReplicaSetBackupApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteReplicaSetBackupExecute(r)
}

/*
DeleteReplicaSetBackup Remove One Replica Set Cloud Backup

Removes the specified snapshot. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return DeleteReplicaSetBackupApiRequest
*/
func (a *CloudBackupsApiService) DeleteReplicaSetBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteReplicaSetBackupApiRequest {
	return DeleteReplicaSetBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CloudBackupsApiService) deleteReplicaSetBackupExecute(r DeleteReplicaSetBackupApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteReplicaSetBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteShardedClusterBackupApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type DeleteShardedClusterBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *CloudBackupsApiService) DeleteShardedClusterBackupWithParams(ctx context.Context, args *DeleteShardedClusterBackupApiParams) DeleteShardedClusterBackupApiRequest {
	return DeleteShardedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r DeleteShardedClusterBackupApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.deleteShardedClusterBackupExecute(r)
}

/*
DeleteShardedClusterBackup Remove One Sharded Cluster Cloud Backup

[experimental] Removes one snapshot of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return DeleteShardedClusterBackupApiRequest
*/
func (a *CloudBackupsApiService) DeleteShardedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteShardedClusterBackupApiRequest {
	return DeleteShardedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CloudBackupsApiService) deleteShardedClusterBackupExecute(r DeleteShardedClusterBackupApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteShardedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetBackupExportJobApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	exportId    string
}

type GetBackupExportJobApiParams struct {
	GroupId     string
	ClusterName string
	ExportId    string
}

func (a *CloudBackupsApiService) GetBackupExportJobWithParams(ctx context.Context, args *GetBackupExportJobApiParams) GetBackupExportJobApiRequest {
	return GetBackupExportJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		exportId:    args.ExportId,
	}
}

func (r GetBackupExportJobApiRequest) Execute() (*DiskBackupExportJob, *http.Response, error) {
	return r.ApiService.getBackupExportJobExecute(r)
}

/*
GetBackupExportJob Return One Cloud Backup Snapshot Export Job

Returns one Cloud Backup snapshot export job associated with the specified Atlas cluster. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param exportId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
	@return GetBackupExportJobApiRequest
*/
func (a *CloudBackupsApiService) GetBackupExportJob(ctx context.Context, groupId string, clusterName string, exportId string) GetBackupExportJobApiRequest {
	return GetBackupExportJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		exportId:    exportId,
	}
}

// Execute executes the request
//
//	@return DiskBackupExportJob
func (a *CloudBackupsApiService) getBackupExportJobExecute(r GetBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupExportJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exportId"+"}", url.PathEscape(parameterValueToString(r.exportId, "exportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetBackupRestoreJobApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	restoreJobId string
}

type GetBackupRestoreJobApiParams struct {
	GroupId      string
	ClusterName  string
	RestoreJobId string
}

func (a *CloudBackupsApiService) GetBackupRestoreJobWithParams(ctx context.Context, args *GetBackupRestoreJobApiParams) GetBackupRestoreJobApiRequest {
	return GetBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		restoreJobId: args.RestoreJobId,
	}
}

func (r GetBackupRestoreJobApiRequest) Execute() (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	return r.ApiService.getBackupRestoreJobExecute(r)
}

/*
GetBackupRestoreJob Return One Restore Job of One Cluster

Returns one cloud backup restore job for one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the restore jobs you want to return.
	@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return GetBackupRestoreJobApiRequest
*/
func (a *CloudBackupsApiService) GetBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) GetBackupRestoreJobApiRequest {
	return GetBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		clusterName:  clusterName,
		restoreJobId: restoreJobId,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotRestoreJob
func (a *CloudBackupsApiService) getBackupRestoreJobExecute(r GetBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(parameterValueToString(r.restoreJobId, "restoreJobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.restoreJobId) < 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have at least 24 elements")
	}
	if strlen(r.restoreJobId) > 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetBackupScheduleApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
}

type GetBackupScheduleApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *CloudBackupsApiService) GetBackupScheduleWithParams(ctx context.Context, args *GetBackupScheduleApiParams) GetBackupScheduleApiRequest {
	return GetBackupScheduleApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetBackupScheduleApiRequest) Execute() (*DiskBackupSnapshotSchedule, *http.Response, error) {
	return r.ApiService.getBackupScheduleExecute(r)
}

/*
GetBackupSchedule Return One Cloud Backup Schedule

Returns the cloud backup schedule for the specified cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetBackupScheduleApiRequest
*/
func (a *CloudBackupsApiService) GetBackupSchedule(ctx context.Context, groupId string, clusterName string) GetBackupScheduleApiRequest {
	return GetBackupScheduleApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotSchedule
func (a *CloudBackupsApiService) getBackupScheduleExecute(r GetBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetDataProtectionSettingsApiRequest struct {
	ctx        context.Context
	ApiService CloudBackupsApi
	groupId    string
}

type GetDataProtectionSettingsApiParams struct {
	GroupId string
}

func (a *CloudBackupsApiService) GetDataProtectionSettingsWithParams(ctx context.Context, args *GetDataProtectionSettingsApiParams) GetDataProtectionSettingsApiRequest {
	return GetDataProtectionSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetDataProtectionSettingsApiRequest) Execute() (*DataProtectionSettings, *http.Response, error) {
	return r.ApiService.getDataProtectionSettingsExecute(r)
}

/*
GetDataProtectionSettings Return the Backup Compliance Policy settings

[experimental] Returns the Backup Compliance Policy settings with the specified project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetDataProtectionSettingsApiRequest
*/
func (a *CloudBackupsApiService) GetDataProtectionSettings(ctx context.Context, groupId string) GetDataProtectionSettingsApiRequest {
	return GetDataProtectionSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return DataProtectionSettings
func (a *CloudBackupsApiService) getDataProtectionSettingsExecute(r GetDataProtectionSettingsApiRequest) (*DataProtectionSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataProtectionSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetDataProtectionSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backupCompliancePolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetExportBucketApiRequest struct {
	ctx            context.Context
	ApiService     CloudBackupsApi
	groupId        string
	exportBucketId string
}

type GetExportBucketApiParams struct {
	GroupId        string
	ExportBucketId string
}

func (a *CloudBackupsApiService) GetExportBucketWithParams(ctx context.Context, args *GetExportBucketApiParams) GetExportBucketApiRequest {
	return GetExportBucketApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		exportBucketId: args.ExportBucketId,
	}
}

func (r GetExportBucketApiRequest) Execute() (*DiskBackupSnapshotAWSExportBucket, *http.Response, error) {
	return r.ApiService.getExportBucketExecute(r)
}

/*
GetExportBucket Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports

Returns one AWS S3 bucket associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param exportBucketId Unique string that identifies the AWS S3 bucket to which you export your snapshots.
	@return GetExportBucketApiRequest
*/
func (a *CloudBackupsApiService) GetExportBucket(ctx context.Context, groupId string, exportBucketId string) GetExportBucketApiRequest {
	return GetExportBucketApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		exportBucketId: exportBucketId,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotAWSExportBucket
func (a *CloudBackupsApiService) getExportBucketExecute(r GetExportBucketApiRequest) (*DiskBackupSnapshotAWSExportBucket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotAWSExportBucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetExportBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exportBucketId"+"}", url.PathEscape(parameterValueToString(r.exportBucketId, "exportBucketId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.exportBucketId) < 3 {
		return localVarReturnValue, nil, reportError("exportBucketId must have at least 3 elements")
	}
	if strlen(r.exportBucketId) > 63 {
		return localVarReturnValue, nil, reportError("exportBucketId must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetReplicaSetBackupApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type GetReplicaSetBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *CloudBackupsApiService) GetReplicaSetBackupWithParams(ctx context.Context, args *GetReplicaSetBackupApiParams) GetReplicaSetBackupApiRequest {
	return GetReplicaSetBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r GetReplicaSetBackupApiRequest) Execute() (*DiskBackupReplicaSet, *http.Response, error) {
	return r.ApiService.getReplicaSetBackupExecute(r)
}

/*
GetReplicaSetBackup Return One Replica Set Cloud Backup

Returns one snapshot from the specified cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return GetReplicaSetBackupApiRequest
*/
func (a *CloudBackupsApiService) GetReplicaSetBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetReplicaSetBackupApiRequest {
	return GetReplicaSetBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return DiskBackupReplicaSet
func (a *CloudBackupsApiService) getReplicaSetBackupExecute(r GetReplicaSetBackupApiRequest) (*DiskBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetReplicaSetBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetServerlessBackupApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type GetServerlessBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *CloudBackupsApiService) GetServerlessBackupWithParams(ctx context.Context, args *GetServerlessBackupApiParams) GetServerlessBackupApiRequest {
	return GetServerlessBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r GetServerlessBackupApiRequest) Execute() (*ServerlessBackupSnapshot, *http.Response, error) {
	return r.ApiService.getServerlessBackupExecute(r)
}

/*
GetServerlessBackup Return One Snapshot of One Serverless Instance

Returns one snapshot of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the serverless instance.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return GetServerlessBackupApiRequest
*/
func (a *CloudBackupsApiService) GetServerlessBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetServerlessBackupApiRequest {
	return GetServerlessBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return ServerlessBackupSnapshot
func (a *CloudBackupsApiService) getServerlessBackupExecute(r GetServerlessBackupApiRequest) (*ServerlessBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetServerlessBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetServerlessBackupRestoreJobApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	restoreJobId string
}

type GetServerlessBackupRestoreJobApiParams struct {
	GroupId      string
	ClusterName  string
	RestoreJobId string
}

func (a *CloudBackupsApiService) GetServerlessBackupRestoreJobWithParams(ctx context.Context, args *GetServerlessBackupRestoreJobApiParams) GetServerlessBackupRestoreJobApiRequest {
	return GetServerlessBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		restoreJobId: args.RestoreJobId,
	}
}

func (r GetServerlessBackupRestoreJobApiRequest) Execute() (*ServerlessBackupRestoreJob, *http.Response, error) {
	return r.ApiService.getServerlessBackupRestoreJobExecute(r)
}

/*
GetServerlessBackupRestoreJob Return One Restore Job for One Serverless Instance

Returns one restore job for one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the serverless instance.
	@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return GetServerlessBackupRestoreJobApiRequest
*/
func (a *CloudBackupsApiService) GetServerlessBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) GetServerlessBackupRestoreJobApiRequest {
	return GetServerlessBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		clusterName:  clusterName,
		restoreJobId: restoreJobId,
	}
}

// Execute executes the request
//
//	@return ServerlessBackupRestoreJob
func (a *CloudBackupsApiService) getServerlessBackupRestoreJobExecute(r GetServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetServerlessBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(parameterValueToString(r.restoreJobId, "restoreJobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.restoreJobId) < 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have at least 24 elements")
	}
	if strlen(r.restoreJobId) > 24 {
		return localVarReturnValue, nil, reportError("restoreJobId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetShardedClusterBackupApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type GetShardedClusterBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *CloudBackupsApiService) GetShardedClusterBackupWithParams(ctx context.Context, args *GetShardedClusterBackupApiParams) GetShardedClusterBackupApiRequest {
	return GetShardedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r GetShardedClusterBackupApiRequest) Execute() (*DiskBackupShardedClusterSnapshot, *http.Response, error) {
	return r.ApiService.getShardedClusterBackupExecute(r)
}

/*
GetShardedClusterBackup Return One Sharded Cluster Cloud Backup

[experimental] Returns one snapshot of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return GetShardedClusterBackupApiRequest
*/
func (a *CloudBackupsApiService) GetShardedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetShardedClusterBackupApiRequest {
	return GetShardedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return DiskBackupShardedClusterSnapshot
func (a *CloudBackupsApiService) getShardedClusterBackupExecute(r GetShardedClusterBackupApiRequest) (*DiskBackupShardedClusterSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupShardedClusterSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetShardedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListBackupExportJobsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListBackupExportJobsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListBackupExportJobsWithParams(ctx context.Context, args *ListBackupExportJobsApiParams) ListBackupExportJobsApiRequest {
	return ListBackupExportJobsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListBackupExportJobsApiRequest) IncludeCount(includeCount bool) ListBackupExportJobsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListBackupExportJobsApiRequest) ItemsPerPage(itemsPerPage int) ListBackupExportJobsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListBackupExportJobsApiRequest) PageNum(pageNum int) ListBackupExportJobsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListBackupExportJobsApiRequest) Execute() (*PaginatedApiAtlasDiskBackupExportJob, *http.Response, error) {
	return r.ApiService.listBackupExportJobsExecute(r)
}

/*
ListBackupExportJobs Return All Cloud Backup Snapshot Export Jobs

Returns all Cloud Backup snapshot export jobs associated with the specified Atlas cluster. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListBackupExportJobsApiRequest
*/
func (a *CloudBackupsApiService) ListBackupExportJobs(ctx context.Context, groupId string, clusterName string) ListBackupExportJobsApiRequest {
	return ListBackupExportJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAtlasDiskBackupExportJob
func (a *CloudBackupsApiService) listBackupExportJobsExecute(r ListBackupExportJobsApiRequest) (*PaginatedApiAtlasDiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasDiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListBackupExportJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListBackupRestoreJobsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListBackupRestoreJobsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListBackupRestoreJobsWithParams(ctx context.Context, args *ListBackupRestoreJobsApiParams) ListBackupRestoreJobsApiRequest {
	return ListBackupRestoreJobsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListBackupRestoreJobsApiRequest) IncludeCount(includeCount bool) ListBackupRestoreJobsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListBackupRestoreJobsApiRequest) ItemsPerPage(itemsPerPage int) ListBackupRestoreJobsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListBackupRestoreJobsApiRequest) PageNum(pageNum int) ListBackupRestoreJobsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListBackupRestoreJobsApiRequest) Execute() (*PaginatedCloudBackupRestoreJob, *http.Response, error) {
	return r.ApiService.listBackupRestoreJobsExecute(r)
}

/*
ListBackupRestoreJobs Return All Restore Jobs for One Cluster

Returns all cloud backup restore jobs for one cluster from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the restore jobs you want to return.
	@return ListBackupRestoreJobsApiRequest
*/
func (a *CloudBackupsApiService) ListBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListBackupRestoreJobsApiRequest {
	return ListBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedCloudBackupRestoreJob
func (a *CloudBackupsApiService) listBackupRestoreJobsExecute(r ListBackupRestoreJobsApiRequest) (*PaginatedCloudBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListExportBucketsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListExportBucketsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListExportBucketsWithParams(ctx context.Context, args *ListExportBucketsApiParams) ListExportBucketsApiRequest {
	return ListExportBucketsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListExportBucketsApiRequest) IncludeCount(includeCount bool) ListExportBucketsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListExportBucketsApiRequest) ItemsPerPage(itemsPerPage int) ListExportBucketsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListExportBucketsApiRequest) PageNum(pageNum int) ListExportBucketsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListExportBucketsApiRequest) Execute() (*PaginatedBackupSnapshotExportBucket, *http.Response, error) {
	return r.ApiService.listExportBucketsExecute(r)
}

/*
ListExportBuckets Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports

Returns all AWS S3 buckets associated with the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListExportBucketsApiRequest
*/
func (a *CloudBackupsApiService) ListExportBuckets(ctx context.Context, groupId string) ListExportBucketsApiRequest {
	return ListExportBucketsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedBackupSnapshotExportBucket
func (a *CloudBackupsApiService) listExportBucketsExecute(r ListExportBucketsApiRequest) (*PaginatedBackupSnapshotExportBucket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedBackupSnapshotExportBucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListExportBuckets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListReplicaSetBackupsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListReplicaSetBackupsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListReplicaSetBackupsWithParams(ctx context.Context, args *ListReplicaSetBackupsApiParams) ListReplicaSetBackupsApiRequest {
	return ListReplicaSetBackupsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListReplicaSetBackupsApiRequest) IncludeCount(includeCount bool) ListReplicaSetBackupsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListReplicaSetBackupsApiRequest) ItemsPerPage(itemsPerPage int) ListReplicaSetBackupsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListReplicaSetBackupsApiRequest) PageNum(pageNum int) ListReplicaSetBackupsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListReplicaSetBackupsApiRequest) Execute() (*PaginatedCloudBackupReplicaSet, *http.Response, error) {
	return r.ApiService.listReplicaSetBackupsExecute(r)
}

/*
ListReplicaSetBackups Return All Replica Set Cloud Backups

Returns all snapshots of one cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListReplicaSetBackupsApiRequest
*/
func (a *CloudBackupsApiService) ListReplicaSetBackups(ctx context.Context, groupId string, clusterName string) ListReplicaSetBackupsApiRequest {
	return ListReplicaSetBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedCloudBackupReplicaSet
func (a *CloudBackupsApiService) listReplicaSetBackupsExecute(r ListReplicaSetBackupsApiRequest) (*PaginatedCloudBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListReplicaSetBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListServerlessBackupRestoreJobsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListServerlessBackupRestoreJobsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListServerlessBackupRestoreJobsWithParams(ctx context.Context, args *ListServerlessBackupRestoreJobsApiParams) ListServerlessBackupRestoreJobsApiRequest {
	return ListServerlessBackupRestoreJobsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListServerlessBackupRestoreJobsApiRequest) IncludeCount(includeCount bool) ListServerlessBackupRestoreJobsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListServerlessBackupRestoreJobsApiRequest) ItemsPerPage(itemsPerPage int) ListServerlessBackupRestoreJobsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServerlessBackupRestoreJobsApiRequest) PageNum(pageNum int) ListServerlessBackupRestoreJobsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServerlessBackupRestoreJobsApiRequest) Execute() (*PaginatedApiAtlasServerlessBackupRestoreJob, *http.Response, error) {
	return r.ApiService.listServerlessBackupRestoreJobsExecute(r)
}

/*
ListServerlessBackupRestoreJobs Return All Restore Jobs for One Serverless Instance

Returns all restore jobs for one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the serverless instance.
	@return ListServerlessBackupRestoreJobsApiRequest
*/
func (a *CloudBackupsApiService) ListServerlessBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListServerlessBackupRestoreJobsApiRequest {
	return ListServerlessBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAtlasServerlessBackupRestoreJob
func (a *CloudBackupsApiService) listServerlessBackupRestoreJobsExecute(r ListServerlessBackupRestoreJobsApiRequest) (*PaginatedApiAtlasServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListServerlessBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListServerlessBackupsApiRequest struct {
	ctx          context.Context
	ApiService   CloudBackupsApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListServerlessBackupsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *CloudBackupsApiService) ListServerlessBackupsWithParams(ctx context.Context, args *ListServerlessBackupsApiParams) ListServerlessBackupsApiRequest {
	return ListServerlessBackupsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListServerlessBackupsApiRequest) IncludeCount(includeCount bool) ListServerlessBackupsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListServerlessBackupsApiRequest) ItemsPerPage(itemsPerPage int) ListServerlessBackupsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServerlessBackupsApiRequest) PageNum(pageNum int) ListServerlessBackupsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServerlessBackupsApiRequest) Execute() (*PaginatedApiAtlasServerlessBackupSnapshot, *http.Response, error) {
	return r.ApiService.listServerlessBackupsExecute(r)
}

/*
ListServerlessBackups Return All Snapshots of One Serverless Instance

Returns all snapshots of one serverless instance from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the serverless instance.
	@return ListServerlessBackupsApiRequest
*/
func (a *CloudBackupsApiService) ListServerlessBackups(ctx context.Context, groupId string, clusterName string) ListServerlessBackupsApiRequest {
	return ListServerlessBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAtlasServerlessBackupSnapshot
func (a *CloudBackupsApiService) listServerlessBackupsExecute(r ListServerlessBackupsApiRequest) (*PaginatedApiAtlasServerlessBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasServerlessBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListServerlessBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListShardedClusterBackupsApiRequest struct {
	ctx         context.Context
	ApiService  CloudBackupsApi
	groupId     string
	clusterName string
}

type ListShardedClusterBackupsApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *CloudBackupsApiService) ListShardedClusterBackupsWithParams(ctx context.Context, args *ListShardedClusterBackupsApiParams) ListShardedClusterBackupsApiRequest {
	return ListShardedClusterBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r ListShardedClusterBackupsApiRequest) Execute() (*PaginatedCloudBackupShardedClusterSnapshot, *http.Response, error) {
	return r.ApiService.listShardedClusterBackupsExecute(r)
}

/*
ListShardedClusterBackups Return All Sharded Cluster Cloud Backups

[experimental] Returns all snapshots of one sharded cluster from the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListShardedClusterBackupsApiRequest
*/
func (a *CloudBackupsApiService) ListShardedClusterBackups(ctx context.Context, groupId string, clusterName string) ListShardedClusterBackupsApiRequest {
	return ListShardedClusterBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedCloudBackupShardedClusterSnapshot
func (a *CloudBackupsApiService) listShardedClusterBackupsExecute(r ListShardedClusterBackupsApiRequest) (*PaginatedCloudBackupShardedClusterSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupShardedClusterSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListShardedClusterBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TakeSnapshotApiRequest struct {
	ctx                               context.Context
	ApiService                        CloudBackupsApi
	groupId                           string
	clusterName                       string
	diskBackupOnDemandSnapshotRequest *DiskBackupOnDemandSnapshotRequest
}

type TakeSnapshotApiParams struct {
	GroupId                           string
	ClusterName                       string
	DiskBackupOnDemandSnapshotRequest *DiskBackupOnDemandSnapshotRequest
}

func (a *CloudBackupsApiService) TakeSnapshotWithParams(ctx context.Context, args *TakeSnapshotApiParams) TakeSnapshotApiRequest {
	return TakeSnapshotApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           args.GroupId,
		clusterName:                       args.ClusterName,
		diskBackupOnDemandSnapshotRequest: args.DiskBackupOnDemandSnapshotRequest,
	}
}

func (r TakeSnapshotApiRequest) Execute() (*DiskBackupSnapshot, *http.Response, error) {
	return r.ApiService.takeSnapshotExecute(r)
}

/*
TakeSnapshot Take One On-Demand Snapshot

Takes one on-demand snapshot for the specified cluster. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return TakeSnapshotApiRequest
*/
func (a *CloudBackupsApiService) TakeSnapshot(ctx context.Context, groupId string, clusterName string, diskBackupOnDemandSnapshotRequest *DiskBackupOnDemandSnapshotRequest) TakeSnapshotApiRequest {
	return TakeSnapshotApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           groupId,
		clusterName:                       clusterName,
		diskBackupOnDemandSnapshotRequest: diskBackupOnDemandSnapshotRequest,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshot
func (a *CloudBackupsApiService) takeSnapshotExecute(r TakeSnapshotApiRequest) (*DiskBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.TakeSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.diskBackupOnDemandSnapshotRequest == nil {
		return localVarReturnValue, nil, reportError("diskBackupOnDemandSnapshotRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupOnDemandSnapshotRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateBackupScheduleApiRequest struct {
	ctx                        context.Context
	ApiService                 CloudBackupsApi
	groupId                    string
	clusterName                string
	diskBackupSnapshotSchedule *DiskBackupSnapshotSchedule
}

type UpdateBackupScheduleApiParams struct {
	GroupId                    string
	ClusterName                string
	DiskBackupSnapshotSchedule *DiskBackupSnapshotSchedule
}

func (a *CloudBackupsApiService) UpdateBackupScheduleWithParams(ctx context.Context, args *UpdateBackupScheduleApiParams) UpdateBackupScheduleApiRequest {
	return UpdateBackupScheduleApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		diskBackupSnapshotSchedule: args.DiskBackupSnapshotSchedule,
	}
}

func (r UpdateBackupScheduleApiRequest) Execute() (*DiskBackupSnapshotSchedule, *http.Response, error) {
	return r.ApiService.updateBackupScheduleExecute(r)
}

/*
UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster

Updates the cloud backup schedule for one cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateBackupScheduleApiRequest
*/
func (a *CloudBackupsApiService) UpdateBackupSchedule(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotSchedule *DiskBackupSnapshotSchedule) UpdateBackupScheduleApiRequest {
	return UpdateBackupScheduleApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		diskBackupSnapshotSchedule: diskBackupSnapshotSchedule,
	}
}

// Execute executes the request
//
//	@return DiskBackupSnapshotSchedule
func (a *CloudBackupsApiService) updateBackupScheduleExecute(r UpdateBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateBackupSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.diskBackupSnapshotSchedule == nil {
		return localVarReturnValue, nil, reportError("diskBackupSnapshotSchedule is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupSnapshotSchedule
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateDataProtectionSettingsApiRequest struct {
	ctx                    context.Context
	ApiService             CloudBackupsApi
	groupId                string
	dataProtectionSettings *DataProtectionSettings
}

type UpdateDataProtectionSettingsApiParams struct {
	GroupId                string
	DataProtectionSettings *DataProtectionSettings
}

func (a *CloudBackupsApiService) UpdateDataProtectionSettingsWithParams(ctx context.Context, args *UpdateDataProtectionSettingsApiParams) UpdateDataProtectionSettingsApiRequest {
	return UpdateDataProtectionSettingsApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		dataProtectionSettings: args.DataProtectionSettings,
	}
}

func (r UpdateDataProtectionSettingsApiRequest) Execute() (*DataProtectionSettings, *http.Response, error) {
	return r.ApiService.updateDataProtectionSettingsExecute(r)
}

/*
UpdateDataProtectionSettings Update or enable the Backup Compliance Policy settings

[experimental] Updates the Backup Compliance Policy settings for the specified project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateDataProtectionSettingsApiRequest
*/
func (a *CloudBackupsApiService) UpdateDataProtectionSettings(ctx context.Context, groupId string, dataProtectionSettings *DataProtectionSettings) UpdateDataProtectionSettingsApiRequest {
	return UpdateDataProtectionSettingsApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		dataProtectionSettings: dataProtectionSettings,
	}
}

// Execute executes the request
//
//	@return DataProtectionSettings
func (a *CloudBackupsApiService) updateDataProtectionSettingsExecute(r UpdateDataProtectionSettingsApiRequest) (*DataProtectionSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataProtectionSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateDataProtectionSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backupCompliancePolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.dataProtectionSettings == nil {
		return localVarReturnValue, nil, reportError("dataProtectionSettings is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.dataProtectionSettings
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateSnapshotRetentionApiRequest struct {
	ctx                     context.Context
	ApiService              CloudBackupsApi
	groupId                 string
	clusterName             string
	snapshotId              string
	backupSnapshotRetention *BackupSnapshotRetention
}

type UpdateSnapshotRetentionApiParams struct {
	GroupId                 string
	ClusterName             string
	SnapshotId              string
	BackupSnapshotRetention *BackupSnapshotRetention
}

func (a *CloudBackupsApiService) UpdateSnapshotRetentionWithParams(ctx context.Context, args *UpdateSnapshotRetentionApiParams) UpdateSnapshotRetentionApiRequest {
	return UpdateSnapshotRetentionApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		clusterName:             args.ClusterName,
		snapshotId:              args.SnapshotId,
		backupSnapshotRetention: args.BackupSnapshotRetention,
	}
}

func (r UpdateSnapshotRetentionApiRequest) Execute() (*DiskBackupReplicaSet, *http.Response, error) {
	return r.ApiService.updateSnapshotRetentionExecute(r)
}

/*
UpdateSnapshotRetention Change Expiration Date for One Cloud Backup

[experimental] Changes the expiration date for one cloud backup snapshot for one cluster in the specified project. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return UpdateSnapshotRetentionApiRequest
*/
func (a *CloudBackupsApiService) UpdateSnapshotRetention(ctx context.Context, groupId string, clusterName string, snapshotId string, backupSnapshotRetention *BackupSnapshotRetention) UpdateSnapshotRetentionApiRequest {
	return UpdateSnapshotRetentionApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		clusterName:             clusterName,
		snapshotId:              snapshotId,
		backupSnapshotRetention: backupSnapshotRetention,
	}
}

// Execute executes the request
//
//	@return DiskBackupReplicaSet
func (a *CloudBackupsApiService) updateSnapshotRetentionExecute(r UpdateSnapshotRetentionApiRequest) (*DiskBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiskBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateSnapshotRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
	}
	if r.backupSnapshotRetention == nil {
		return localVarReturnValue, nil, reportError("backupSnapshotRetention is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.backupSnapshotRetention
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
