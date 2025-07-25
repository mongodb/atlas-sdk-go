// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CloudBackupsApi interface {

	/*
		CancelBackupRestoreJob Cancel One Restore Job for One Cluster

		Cancels one cloud backup restore job of one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to remove.
		@return CancelBackupRestoreJobApiRequest
	*/
	CancelBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) CancelBackupRestoreJobApiRequest
	/*
		CancelBackupRestoreJob Cancel One Restore Job for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CancelBackupRestoreJobApiParams - Parameters for the request
		@return CancelBackupRestoreJobApiRequest
	*/
	CancelBackupRestoreJobWithParams(ctx context.Context, args *CancelBackupRestoreJobApiParams) CancelBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	CancelBackupRestoreJobExecute(r CancelBackupRestoreJobApiRequest) (*http.Response, error)

	/*
		CreateBackupExportJob Create One Snapshot Export Job

		Exports one backup Snapshot for dedicated Atlas cluster using Cloud Backups to an Export Bucket. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param diskBackupExportJobRequest Information about the Cloud Backup Snapshot Export Job to create.
		@return CreateBackupExportJobApiRequest
	*/
	CreateBackupExportJob(ctx context.Context, groupId string, clusterName string, diskBackupExportJobRequest *DiskBackupExportJobRequest) CreateBackupExportJobApiRequest
	/*
		CreateBackupExportJob Create One Snapshot Export Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateBackupExportJobApiParams - Parameters for the request
		@return CreateBackupExportJobApiRequest
	*/
	CreateBackupExportJobWithParams(ctx context.Context, args *CreateBackupExportJobApiParams) CreateBackupExportJobApiRequest

	// Method available only for mocking purposes
	CreateBackupExportJobExecute(r CreateBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error)

	/*
			CreateBackupRestoreJob Restore One Snapshot of One Cluster

			Restores one snapshot of one cluster from the specified project. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

		 To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster.
			@param diskBackupSnapshotRestoreJob Restores one snapshot of one cluster from the specified project.
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

	// Method available only for mocking purposes
	CreateBackupRestoreJobExecute(r CreateBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error)

	/*
		CreateExportBucket Create One Snapshot Export Bucket

		Creates a Snapshot Export Bucket for an AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket. Once created, an snapshots can be exported to the Export Bucket and its referenced AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param diskBackupSnapshotExportBucketRequest Specifies the role and AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket that the Export Bucket should reference.
		@return CreateExportBucketApiRequest
	*/
	CreateExportBucket(ctx context.Context, groupId string, diskBackupSnapshotExportBucketRequest *DiskBackupSnapshotExportBucketRequest) CreateExportBucketApiRequest
	/*
		CreateExportBucket Create One Snapshot Export Bucket


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateExportBucketApiParams - Parameters for the request
		@return CreateExportBucketApiRequest
	*/
	CreateExportBucketWithParams(ctx context.Context, args *CreateExportBucketApiParams) CreateExportBucketApiRequest

	// Method available only for mocking purposes
	CreateExportBucketExecute(r CreateExportBucketApiRequest) (*DiskBackupSnapshotExportBucketResponse, *http.Response, error)

	/*
			CreateServerlessBackupRestoreJob Restore One Snapshot of One Serverless Instance

			Restores one snapshot of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the createFlexBackupRestoreJob endpoint instead.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the serverless instance whose snapshot you want to restore.
			@param serverlessBackupRestoreJob Restores one snapshot of one serverless instance from the specified project.
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

	// Method available only for mocking purposes
	CreateServerlessBackupRestoreJobExecute(r CreateServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error)

	/*
		DeleteAllBackupSchedules Remove All Cloud Backup Schedules

		Removes all cloud backup schedules for the specified cluster. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role. Deprecated versions: v2-{2023-01-01}

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

	// Method available only for mocking purposes
	DeleteAllBackupSchedulesExecute(r DeleteAllBackupSchedulesApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error)

	/*
		DeleteExportBucket Delete One Snapshot Export Bucket

		Deletes an Export Bucket. Auto export must be disabled on all clusters in this Project exporting to this Export Bucket before revoking access. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param exportBucketId Unique 24-hexadecimal character string that identifies the Export Bucket.
		@return DeleteExportBucketApiRequest
	*/
	DeleteExportBucket(ctx context.Context, groupId string, exportBucketId string) DeleteExportBucketApiRequest
	/*
		DeleteExportBucket Delete One Snapshot Export Bucket


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteExportBucketApiParams - Parameters for the request
		@return DeleteExportBucketApiRequest
	*/
	DeleteExportBucketWithParams(ctx context.Context, args *DeleteExportBucketApiParams) DeleteExportBucketApiRequest

	// Method available only for mocking purposes
	DeleteExportBucketExecute(r DeleteExportBucketApiRequest) (*http.Response, error)

	/*
		DeleteReplicaSetBackup Remove One Replica Set Cloud Backup

		Removes the specified snapshot. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

	// Method available only for mocking purposes
	DeleteReplicaSetBackupExecute(r DeleteReplicaSetBackupApiRequest) (*http.Response, error)

	/*
		DeleteShardedClusterBackup Remove One Sharded Cluster Cloud Backup

		Removes one snapshot of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

	// Method available only for mocking purposes
	DeleteShardedClusterBackupExecute(r DeleteShardedClusterBackupApiRequest) (*http.Response, error)

	/*
		DisableDataProtectionSettings Disable Backup Compliance Policy Settings

		Disables the Backup Compliance Policy settings with the specified project. As a prerequisite, a support ticket needs to be file first, instructions in https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DisableDataProtectionSettingsApiRequest
	*/
	DisableDataProtectionSettings(ctx context.Context, groupId string) DisableDataProtectionSettingsApiRequest
	/*
		DisableDataProtectionSettings Disable Backup Compliance Policy Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DisableDataProtectionSettingsApiParams - Parameters for the request
		@return DisableDataProtectionSettingsApiRequest
	*/
	DisableDataProtectionSettingsWithParams(ctx context.Context, args *DisableDataProtectionSettingsApiParams) DisableDataProtectionSettingsApiRequest

	// Method available only for mocking purposes
	DisableDataProtectionSettingsExecute(r DisableDataProtectionSettingsApiRequest) (*http.Response, error)

	/*
		GetBackupExportJob Return One Snapshot Export Job

		Returns one Cloud Backup Snapshot Export Job associated with the specified Atlas cluster. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param exportId Unique 24-hexadecimal character string that identifies the Export Job.
		@return GetBackupExportJobApiRequest
	*/
	GetBackupExportJob(ctx context.Context, groupId string, clusterName string, exportId string) GetBackupExportJobApiRequest
	/*
		GetBackupExportJob Return One Snapshot Export Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetBackupExportJobApiParams - Parameters for the request
		@return GetBackupExportJobApiRequest
	*/
	GetBackupExportJobWithParams(ctx context.Context, args *GetBackupExportJobApiParams) GetBackupExportJobApiRequest

	// Method available only for mocking purposes
	GetBackupExportJobExecute(r GetBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error)

	/*
		GetBackupRestoreJob Return One Restore Job for One Cluster

		Returns one cloud backup restore job for one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the restore jobs you want to return.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
		@return GetBackupRestoreJobApiRequest
	*/
	GetBackupRestoreJob(ctx context.Context, groupId string, clusterName string, restoreJobId string) GetBackupRestoreJobApiRequest
	/*
		GetBackupRestoreJob Return One Restore Job for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetBackupRestoreJobApiParams - Parameters for the request
		@return GetBackupRestoreJobApiRequest
	*/
	GetBackupRestoreJobWithParams(ctx context.Context, args *GetBackupRestoreJobApiParams) GetBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	GetBackupRestoreJobExecute(r GetBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error)

	/*
		GetBackupSchedule Return One Cloud Backup Schedule

		Returns the cloud backup schedule for the specified cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

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

	// Method available only for mocking purposes
	GetBackupScheduleExecute(r GetBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error)

	/*
		GetDataProtectionSettings Return Backup Compliance Policy Settings

		Returns the Backup Compliance Policy settings with the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetDataProtectionSettingsApiRequest
	*/
	GetDataProtectionSettings(ctx context.Context, groupId string) GetDataProtectionSettingsApiRequest
	/*
		GetDataProtectionSettings Return Backup Compliance Policy Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDataProtectionSettingsApiParams - Parameters for the request
		@return GetDataProtectionSettingsApiRequest
	*/
	GetDataProtectionSettingsWithParams(ctx context.Context, args *GetDataProtectionSettingsApiParams) GetDataProtectionSettingsApiRequest

	// Method available only for mocking purposes
	GetDataProtectionSettingsExecute(r GetDataProtectionSettingsApiRequest) (*DataProtectionSettings20231001, *http.Response, error)

	/*
		GetExportBucket Return One Snapshot Export Bucket

		Returns one Export Bucket associated with the specified Project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param exportBucketId Unique 24-hexadecimal character string that identifies the Export Bucket.
		@return GetExportBucketApiRequest
	*/
	GetExportBucket(ctx context.Context, groupId string, exportBucketId string) GetExportBucketApiRequest
	/*
		GetExportBucket Return One Snapshot Export Bucket


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetExportBucketApiParams - Parameters for the request
		@return GetExportBucketApiRequest
	*/
	GetExportBucketWithParams(ctx context.Context, args *GetExportBucketApiParams) GetExportBucketApiRequest

	// Method available only for mocking purposes
	GetExportBucketExecute(r GetExportBucketApiRequest) (*DiskBackupSnapshotExportBucketResponse, *http.Response, error)

	/*
		GetReplicaSetBackup Return One Replica Set Cloud Backup

		Returns one snapshot from the specified cluster. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

	// Method available only for mocking purposes
	GetReplicaSetBackupExecute(r GetReplicaSetBackupApiRequest) (*DiskBackupReplicaSet, *http.Response, error)

	/*
			GetServerlessBackup Return One Snapshot of One Serverless Instance

			Returns one snapshot of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		This endpoint can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) API or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the getFlexBackup endpoint instead.

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

	// Method available only for mocking purposes
	GetServerlessBackupExecute(r GetServerlessBackupApiRequest) (*ServerlessBackupSnapshot, *http.Response, error)

	/*
			GetServerlessBackupRestoreJob Return One Restore Job for One Serverless Instance

			Returns one restore job for one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the getFlexBackupRestoreJob endpoint instead.

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

	// Method available only for mocking purposes
	GetServerlessBackupRestoreJobExecute(r GetServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error)

	/*
		GetShardedClusterBackup Return One Sharded Cluster Cloud Backup

		Returns one snapshot of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

	// Method available only for mocking purposes
	GetShardedClusterBackupExecute(r GetShardedClusterBackupApiRequest) (*DiskBackupShardedClusterSnapshot, *http.Response, error)

	/*
		ListBackupExportJobs Return All Snapshot Export Jobs

		Returns all Cloud Backup Snapshot Export Jobs associated with the specified Atlas cluster. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListBackupExportJobsApiRequest
	*/
	ListBackupExportJobs(ctx context.Context, groupId string, clusterName string) ListBackupExportJobsApiRequest
	/*
		ListBackupExportJobs Return All Snapshot Export Jobs


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListBackupExportJobsApiParams - Parameters for the request
		@return ListBackupExportJobsApiRequest
	*/
	ListBackupExportJobsWithParams(ctx context.Context, args *ListBackupExportJobsApiParams) ListBackupExportJobsApiRequest

	// Method available only for mocking purposes
	ListBackupExportJobsExecute(r ListBackupExportJobsApiRequest) (*PaginatedApiAtlasDiskBackupExportJob, *http.Response, error)

	/*
		ListBackupRestoreJobs Return All Restore Jobs for One Cluster

		Returns all cloud backup restore jobs for one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

	// Method available only for mocking purposes
	ListBackupRestoreJobsExecute(r ListBackupRestoreJobsApiRequest) (*PaginatedCloudBackupRestoreJob, *http.Response, error)

	/*
		ListExportBuckets Return All Snapshot Export Buckets

		Returns all Export Buckets associated with the specified Project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListExportBucketsApiRequest
	*/
	ListExportBuckets(ctx context.Context, groupId string) ListExportBucketsApiRequest
	/*
		ListExportBuckets Return All Snapshot Export Buckets


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListExportBucketsApiParams - Parameters for the request
		@return ListExportBucketsApiRequest
	*/
	ListExportBucketsWithParams(ctx context.Context, args *ListExportBucketsApiParams) ListExportBucketsApiRequest

	// Method available only for mocking purposes
	ListExportBucketsExecute(r ListExportBucketsApiRequest) (*PaginatedBackupSnapshotExportBuckets, *http.Response, error)

	/*
		ListReplicaSetBackups Return All Replica Set Cloud Backups

		Returns all snapshots of one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

	// Method available only for mocking purposes
	ListReplicaSetBackupsExecute(r ListReplicaSetBackupsApiRequest) (*PaginatedCloudBackupReplicaSet, *http.Response, error)

	/*
			ListServerlessBackupRestoreJobs Return All Restore Jobs for One Serverless Instance

			Returns all restore jobs for one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the listFlexBackupRestoreJobs endpoint instead.

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

	// Method available only for mocking purposes
	ListServerlessBackupRestoreJobsExecute(r ListServerlessBackupRestoreJobsApiRequest) (*PaginatedApiAtlasServerlessBackupRestoreJob, *http.Response, error)

	/*
			ListServerlessBackups Return All Snapshots of One Serverless Instance

			Returns all snapshots of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the listFlexBackups endpoint instead.

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

	// Method available only for mocking purposes
	ListServerlessBackupsExecute(r ListServerlessBackupsApiRequest) (*PaginatedApiAtlasServerlessBackupSnapshot, *http.Response, error)

	/*
		ListShardedClusterBackups Return All Sharded Cluster Cloud Backups

		Returns all snapshots of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

	// Method available only for mocking purposes
	ListShardedClusterBackupsExecute(r ListShardedClusterBackupsApiRequest) (*PaginatedCloudBackupShardedClusterSnapshot, *http.Response, error)

	/*
			TakeSnapshot Take One On-Demand Snapshot

			Takes one on-demand snapshot for the specified cluster. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

		 To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster.
			@param diskBackupOnDemandSnapshotRequest Takes one on-demand snapshot.
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

	// Method available only for mocking purposes
	TakeSnapshotExecute(r TakeSnapshotApiRequest) (*DiskBackupSnapshot, *http.Response, error)

	/*
		UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster

		Updates the cloud backup schedule for one cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param diskBackupSnapshotSchedule20240805 Updates the cloud backup schedule for one cluster within the specified project.  **Note**: In the request body, provide only the fields that you want to update.
		@return UpdateBackupScheduleApiRequest
	*/
	UpdateBackupSchedule(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotSchedule20240805 *DiskBackupSnapshotSchedule20240805) UpdateBackupScheduleApiRequest
	/*
		UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateBackupScheduleApiParams - Parameters for the request
		@return UpdateBackupScheduleApiRequest
	*/
	UpdateBackupScheduleWithParams(ctx context.Context, args *UpdateBackupScheduleApiParams) UpdateBackupScheduleApiRequest

	// Method available only for mocking purposes
	UpdateBackupScheduleExecute(r UpdateBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error)

	/*
		UpdateDataProtectionSettings Update Backup Compliance Policy Settings

		Updates the Backup Compliance Policy settings for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param dataProtectionSettings20231001 The new Backup Compliance Policy settings.
		@return UpdateDataProtectionSettingsApiRequest
	*/
	UpdateDataProtectionSettings(ctx context.Context, groupId string, dataProtectionSettings20231001 *DataProtectionSettings20231001) UpdateDataProtectionSettingsApiRequest
	/*
		UpdateDataProtectionSettings Update Backup Compliance Policy Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateDataProtectionSettingsApiParams - Parameters for the request
		@return UpdateDataProtectionSettingsApiRequest
	*/
	UpdateDataProtectionSettingsWithParams(ctx context.Context, args *UpdateDataProtectionSettingsApiParams) UpdateDataProtectionSettingsApiRequest

	// Method available only for mocking purposes
	UpdateDataProtectionSettingsExecute(r UpdateDataProtectionSettingsApiRequest) (*DataProtectionSettings20231001, *http.Response, error)

	/*
		UpdateSnapshotRetention Update Expiration Date for One Cloud Backup

		Changes the expiration date for one cloud backup snapshot for one cluster in the specified project, the requesting Service Account or API Key must have the Project Backup Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@param backupSnapshotRetention Changes the expiration date for one cloud backup snapshot for one cluster in the specified project.
		@return UpdateSnapshotRetentionApiRequest
	*/
	UpdateSnapshotRetention(ctx context.Context, groupId string, clusterName string, snapshotId string, backupSnapshotRetention *BackupSnapshotRetention) UpdateSnapshotRetentionApiRequest
	/*
		UpdateSnapshotRetention Update Expiration Date for One Cloud Backup


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateSnapshotRetentionApiParams - Parameters for the request
		@return UpdateSnapshotRetentionApiRequest
	*/
	UpdateSnapshotRetentionWithParams(ctx context.Context, args *UpdateSnapshotRetentionApiParams) UpdateSnapshotRetentionApiRequest

	// Method available only for mocking purposes
	UpdateSnapshotRetentionExecute(r UpdateSnapshotRetentionApiRequest) (*DiskBackupReplicaSet, *http.Response, error)
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

func (r CancelBackupRestoreJobApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelBackupRestoreJobExecute(r)
}

/*
CancelBackupRestoreJob Cancel One Restore Job for One Cluster

Cancels one cloud backup restore job of one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// CancelBackupRestoreJobExecute executes the request
func (a *CloudBackupsApiService) CancelBackupRestoreJobExecute(r CancelBackupRestoreJobApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CancelBackupRestoreJob")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.restoreJobId == "" {
		return nil, reportError("restoreJobId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(r.restoreJobId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
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
	return r.ApiService.CreateBackupExportJobExecute(r)
}

/*
CreateBackupExportJob Create One Snapshot Export Job

Exports one backup Snapshot for dedicated Atlas cluster using Cloud Backups to an Export Bucket. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

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

// CreateBackupExportJobExecute executes the request
//
//	@return DiskBackupExportJob
func (a *CloudBackupsApiService) CreateBackupExportJobExecute(r CreateBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateBackupExportJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.CreateBackupRestoreJobExecute(r)
}

/*
CreateBackupRestoreJob Restore One Snapshot of One Cluster

Restores one snapshot of one cluster from the specified project. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// CreateBackupRestoreJobExecute executes the request
//
//	@return DiskBackupSnapshotRestoreJob
func (a *CloudBackupsApiService) CreateBackupRestoreJobExecute(r CreateBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateExportBucketApiRequest struct {
	ctx                                   context.Context
	ApiService                            CloudBackupsApi
	groupId                               string
	diskBackupSnapshotExportBucketRequest *DiskBackupSnapshotExportBucketRequest
}

type CreateExportBucketApiParams struct {
	GroupId                               string
	DiskBackupSnapshotExportBucketRequest *DiskBackupSnapshotExportBucketRequest
}

func (a *CloudBackupsApiService) CreateExportBucketWithParams(ctx context.Context, args *CreateExportBucketApiParams) CreateExportBucketApiRequest {
	return CreateExportBucketApiRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		groupId:                               args.GroupId,
		diskBackupSnapshotExportBucketRequest: args.DiskBackupSnapshotExportBucketRequest,
	}
}

func (r CreateExportBucketApiRequest) Execute() (*DiskBackupSnapshotExportBucketResponse, *http.Response, error) {
	return r.ApiService.CreateExportBucketExecute(r)
}

/*
CreateExportBucket Create One Snapshot Export Bucket

Creates a Snapshot Export Bucket for an AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket. Once created, an snapshots can be exported to the Export Bucket and its referenced AWS S3 Bucket, Azure Blob Storage Container, or Google Cloud Storage Bucket. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateExportBucketApiRequest
*/
func (a *CloudBackupsApiService) CreateExportBucket(ctx context.Context, groupId string, diskBackupSnapshotExportBucketRequest *DiskBackupSnapshotExportBucketRequest) CreateExportBucketApiRequest {
	return CreateExportBucketApiRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		groupId:                               groupId,
		diskBackupSnapshotExportBucketRequest: diskBackupSnapshotExportBucketRequest,
	}
}

// CreateExportBucketExecute executes the request
//
//	@return DiskBackupSnapshotExportBucketResponse
func (a *CloudBackupsApiService) CreateExportBucketExecute(r CreateExportBucketApiRequest) (*DiskBackupSnapshotExportBucketResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotExportBucketResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateExportBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.diskBackupSnapshotExportBucketRequest == nil {
		return localVarReturnValue, nil, reportError("diskBackupSnapshotExportBucketRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupSnapshotExportBucketRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.CreateServerlessBackupRestoreJobExecute(r)
}

/*
CreateServerlessBackupRestoreJob Restore One Snapshot of One Serverless Instance

Restores one snapshot of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the createFlexBackupRestoreJob endpoint instead.

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

// CreateServerlessBackupRestoreJobExecute executes the request
//
//	@return ServerlessBackupRestoreJob
func (a *CloudBackupsApiService) CreateServerlessBackupRestoreJobExecute(r CreateServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.CreateServerlessBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r DeleteAllBackupSchedulesApiRequest) Execute() (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	return r.ApiService.DeleteAllBackupSchedulesExecute(r)
}

/*
DeleteAllBackupSchedules Remove All Cloud Backup Schedules

Removes all cloud backup schedules for the specified cluster. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role. Deprecated versions: v2-{2023-01-01}

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

// DeleteAllBackupSchedulesExecute executes the request
//
//	@return DiskBackupSnapshotSchedule20240805
func (a *CloudBackupsApiService) DeleteAllBackupSchedulesExecute(r DeleteAllBackupSchedulesApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteAllBackupSchedules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r DeleteExportBucketApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExportBucketExecute(r)
}

/*
DeleteExportBucket Delete One Snapshot Export Bucket

Deletes an Export Bucket. Auto export must be disabled on all clusters in this Project exporting to this Export Bucket before revoking access. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param exportBucketId Unique 24-hexadecimal character string that identifies the Export Bucket.
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

// DeleteExportBucketExecute executes the request
func (a *CloudBackupsApiService) DeleteExportBucketExecute(r DeleteExportBucketApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteExportBucket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.exportBucketId == "" {
		return nil, reportError("exportBucketId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"exportBucketId"+"}", url.PathEscape(r.exportBucketId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
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

func (r DeleteReplicaSetBackupApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteReplicaSetBackupExecute(r)
}

/*
DeleteReplicaSetBackup Remove One Replica Set Cloud Backup

Removes the specified snapshot. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// DeleteReplicaSetBackupExecute executes the request
func (a *CloudBackupsApiService) DeleteReplicaSetBackupExecute(r DeleteReplicaSetBackupApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteReplicaSetBackup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
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

func (r DeleteShardedClusterBackupApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteShardedClusterBackupExecute(r)
}

/*
DeleteShardedClusterBackup Remove One Sharded Cluster Cloud Backup

Removes one snapshot of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// DeleteShardedClusterBackupExecute executes the request
func (a *CloudBackupsApiService) DeleteShardedClusterBackupExecute(r DeleteShardedClusterBackupApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DeleteShardedClusterBackup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DisableDataProtectionSettingsApiRequest struct {
	ctx        context.Context
	ApiService CloudBackupsApi
	groupId    string
}

type DisableDataProtectionSettingsApiParams struct {
	GroupId string
}

func (a *CloudBackupsApiService) DisableDataProtectionSettingsWithParams(ctx context.Context, args *DisableDataProtectionSettingsApiParams) DisableDataProtectionSettingsApiRequest {
	return DisableDataProtectionSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DisableDataProtectionSettingsApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DisableDataProtectionSettingsExecute(r)
}

/*
DisableDataProtectionSettings Disable Backup Compliance Policy Settings

Disables the Backup Compliance Policy settings with the specified project. As a prerequisite, a support ticket needs to be file first, instructions in https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DisableDataProtectionSettingsApiRequest
*/
func (a *CloudBackupsApiService) DisableDataProtectionSettings(ctx context.Context, groupId string) DisableDataProtectionSettingsApiRequest {
	return DisableDataProtectionSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// DisableDataProtectionSettingsExecute executes the request
func (a *CloudBackupsApiService) DisableDataProtectionSettingsExecute(r DisableDataProtectionSettingsApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.DisableDataProtectionSettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backupCompliancePolicy"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-11-13+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
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
	return r.ApiService.GetBackupExportJobExecute(r)
}

/*
GetBackupExportJob Return One Snapshot Export Job

Returns one Cloud Backup Snapshot Export Job associated with the specified Atlas cluster. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param exportId Unique 24-hexadecimal character string that identifies the Export Job.
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

// GetBackupExportJobExecute executes the request
//
//	@return DiskBackupExportJob
func (a *CloudBackupsApiService) GetBackupExportJobExecute(r GetBackupExportJobApiRequest) (*DiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupExportJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.exportId == "" {
		return localVarReturnValue, nil, reportError("exportId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"exportId"+"}", url.PathEscape(r.exportId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.GetBackupRestoreJobExecute(r)
}

/*
GetBackupRestoreJob Return One Restore Job for One Cluster

Returns one cloud backup restore job for one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// GetBackupRestoreJobExecute executes the request
//
//	@return DiskBackupSnapshotRestoreJob
func (a *CloudBackupsApiService) GetBackupRestoreJobExecute(r GetBackupRestoreJobApiRequest) (*DiskBackupSnapshotRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.restoreJobId == "" {
		return localVarReturnValue, nil, reportError("restoreJobId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(r.restoreJobId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r GetBackupScheduleApiRequest) Execute() (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	return r.ApiService.GetBackupScheduleExecute(r)
}

/*
GetBackupSchedule Return One Cloud Backup Schedule

Returns the cloud backup schedule for the specified cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

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

// GetBackupScheduleExecute executes the request
//
//	@return DiskBackupSnapshotSchedule20240805
func (a *CloudBackupsApiService) GetBackupScheduleExecute(r GetBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetBackupSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r GetDataProtectionSettingsApiRequest) Execute() (*DataProtectionSettings20231001, *http.Response, error) {
	return r.ApiService.GetDataProtectionSettingsExecute(r)
}

/*
GetDataProtectionSettings Return Backup Compliance Policy Settings

Returns the Backup Compliance Policy settings with the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

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

// GetDataProtectionSettingsExecute executes the request
//
//	@return DataProtectionSettings20231001
func (a *CloudBackupsApiService) GetDataProtectionSettingsExecute(r GetDataProtectionSettingsApiRequest) (*DataProtectionSettings20231001, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataProtectionSettings20231001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetDataProtectionSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backupCompliancePolicy"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r GetExportBucketApiRequest) Execute() (*DiskBackupSnapshotExportBucketResponse, *http.Response, error) {
	return r.ApiService.GetExportBucketExecute(r)
}

/*
GetExportBucket Return One Snapshot Export Bucket

Returns one Export Bucket associated with the specified Project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param exportBucketId Unique 24-hexadecimal character string that identifies the Export Bucket.
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

// GetExportBucketExecute executes the request
//
//	@return DiskBackupSnapshotExportBucketResponse
func (a *CloudBackupsApiService) GetExportBucketExecute(r GetExportBucketApiRequest) (*DiskBackupSnapshotExportBucketResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotExportBucketResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetExportBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.exportBucketId == "" {
		return localVarReturnValue, nil, reportError("exportBucketId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"exportBucketId"+"}", url.PathEscape(r.exportBucketId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.GetReplicaSetBackupExecute(r)
}

/*
GetReplicaSetBackup Return One Replica Set Cloud Backup

Returns one snapshot from the specified cluster. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

// GetReplicaSetBackupExecute executes the request
//
//	@return DiskBackupReplicaSet
func (a *CloudBackupsApiService) GetReplicaSetBackupExecute(r GetReplicaSetBackupApiRequest) (*DiskBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetReplicaSetBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return localVarReturnValue, nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.GetServerlessBackupExecute(r)
}

/*
GetServerlessBackup Return One Snapshot of One Serverless Instance

Returns one snapshot of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

This endpoint can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) API or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the getFlexBackup endpoint instead.

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

// GetServerlessBackupExecute executes the request
//
//	@return ServerlessBackupSnapshot
func (a *CloudBackupsApiService) GetServerlessBackupExecute(r GetServerlessBackupApiRequest) (*ServerlessBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetServerlessBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return localVarReturnValue, nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.GetServerlessBackupRestoreJobExecute(r)
}

/*
GetServerlessBackupRestoreJob Return One Restore Job for One Serverless Instance

Returns one restore job for one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the getFlexBackupRestoreJob endpoint instead.

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

// GetServerlessBackupRestoreJobExecute executes the request
//
//	@return ServerlessBackupRestoreJob
func (a *CloudBackupsApiService) GetServerlessBackupRestoreJobExecute(r GetServerlessBackupRestoreJobApiRequest) (*ServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetServerlessBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.restoreJobId == "" {
		return localVarReturnValue, nil, reportError("restoreJobId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"restoreJobId"+"}", url.PathEscape(r.restoreJobId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.GetShardedClusterBackupExecute(r)
}

/*
GetShardedClusterBackup Return One Sharded Cluster Cloud Backup

Returns one snapshot of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

// GetShardedClusterBackupExecute executes the request
//
//	@return DiskBackupShardedClusterSnapshot
func (a *CloudBackupsApiService) GetShardedClusterBackupExecute(r GetShardedClusterBackupApiRequest) (*DiskBackupShardedClusterSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupShardedClusterSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.GetShardedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return localVarReturnValue, nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListBackupExportJobsExecute(r)
}

/*
ListBackupExportJobs Return All Snapshot Export Jobs

Returns all Cloud Backup Snapshot Export Jobs associated with the specified Atlas cluster. To use this resource, the requesting Service Account or API Key must have the Project Atlas Admin role.

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

// ListBackupExportJobsExecute executes the request
//
//	@return PaginatedApiAtlasDiskBackupExportJob
func (a *CloudBackupsApiService) ListBackupExportJobsExecute(r ListBackupExportJobsApiRequest) (*PaginatedApiAtlasDiskBackupExportJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasDiskBackupExportJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListBackupExportJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListBackupRestoreJobsExecute(r)
}

/*
ListBackupRestoreJobs Return All Restore Jobs for One Cluster

Returns all cloud backup restore jobs for one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// ListBackupRestoreJobsExecute executes the request
//
//	@return PaginatedCloudBackupRestoreJob
func (a *CloudBackupsApiService) ListBackupRestoreJobsExecute(r ListBackupRestoreJobsApiRequest) (*PaginatedCloudBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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

func (r ListExportBucketsApiRequest) Execute() (*PaginatedBackupSnapshotExportBuckets, *http.Response, error) {
	return r.ApiService.ListExportBucketsExecute(r)
}

/*
ListExportBuckets Return All Snapshot Export Buckets

Returns all Export Buckets associated with the specified Project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

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

// ListExportBucketsExecute executes the request
//
//	@return PaginatedBackupSnapshotExportBuckets
func (a *CloudBackupsApiService) ListExportBucketsExecute(r ListExportBucketsApiRequest) (*PaginatedBackupSnapshotExportBuckets, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedBackupSnapshotExportBuckets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListExportBuckets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backup/exportBuckets"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListReplicaSetBackupsExecute(r)
}

/*
ListReplicaSetBackups Return All Replica Set Cloud Backups

Returns all snapshots of one cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

// ListReplicaSetBackupsExecute executes the request
//
//	@return PaginatedCloudBackupReplicaSet
func (a *CloudBackupsApiService) ListReplicaSetBackupsExecute(r ListReplicaSetBackupsApiRequest) (*PaginatedCloudBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListReplicaSetBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListServerlessBackupRestoreJobsExecute(r)
}

/*
ListServerlessBackupRestoreJobs Return All Restore Jobs for One Serverless Instance

Returns all restore jobs for one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the listFlexBackupRestoreJobs endpoint instead.

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

// ListServerlessBackupRestoreJobsExecute executes the request
//
//	@return PaginatedApiAtlasServerlessBackupRestoreJob
func (a *CloudBackupsApiService) ListServerlessBackupRestoreJobsExecute(r ListServerlessBackupRestoreJobsApiRequest) (*PaginatedApiAtlasServerlessBackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasServerlessBackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListServerlessBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListServerlessBackupsExecute(r)
}

/*
ListServerlessBackups Return All Snapshots of One Serverless Instance

Returns all snapshots of one serverless instance from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

This API can also be used on Flex clusters that were created with the [createServerlessInstance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Instances/operation/createServerlessInstance) endpoint or Flex clusters that were migrated from Serverless instances. This endpoint will be sunset in January 2026. Please use the listFlexBackups endpoint instead.

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

// ListServerlessBackupsExecute executes the request
//
//	@return PaginatedApiAtlasServerlessBackupSnapshot
func (a *CloudBackupsApiService) ListServerlessBackupsExecute(r ListServerlessBackupsApiRequest) (*PaginatedApiAtlasServerlessBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasServerlessBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListServerlessBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.ListShardedClusterBackupsExecute(r)
}

/*
ListShardedClusterBackups Return All Sharded Cluster Cloud Backups

Returns all snapshots of one sharded cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role or Project Backup Manager role.

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

// ListShardedClusterBackupsExecute executes the request
//
//	@return PaginatedCloudBackupShardedClusterSnapshot
func (a *CloudBackupsApiService) ListShardedClusterBackupsExecute(r ListShardedClusterBackupsApiRequest) (*PaginatedCloudBackupShardedClusterSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudBackupShardedClusterSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.ListShardedClusterBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.TakeSnapshotExecute(r)
}

/*
TakeSnapshot Take One On-Demand Snapshot

Takes one on-demand snapshot for the specified cluster. Atlas takes on-demand snapshots immediately and scheduled snapshots at regular intervals. If an on-demand snapshot with a status of **queued** or **inProgress** exists, before taking another snapshot, wait until Atlas completes completes processing the previously taken on-demand snapshot.

	To use this resource, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// TakeSnapshotExecute executes the request
//
//	@return DiskBackupSnapshot
func (a *CloudBackupsApiService) TakeSnapshotExecute(r TakeSnapshotApiRequest) (*DiskBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.TakeSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateBackupScheduleApiRequest struct {
	ctx                                context.Context
	ApiService                         CloudBackupsApi
	groupId                            string
	clusterName                        string
	diskBackupSnapshotSchedule20240805 *DiskBackupSnapshotSchedule20240805
}

type UpdateBackupScheduleApiParams struct {
	GroupId                            string
	ClusterName                        string
	DiskBackupSnapshotSchedule20240805 *DiskBackupSnapshotSchedule20240805
}

func (a *CloudBackupsApiService) UpdateBackupScheduleWithParams(ctx context.Context, args *UpdateBackupScheduleApiParams) UpdateBackupScheduleApiRequest {
	return UpdateBackupScheduleApiRequest{
		ApiService:                         a,
		ctx:                                ctx,
		groupId:                            args.GroupId,
		clusterName:                        args.ClusterName,
		diskBackupSnapshotSchedule20240805: args.DiskBackupSnapshotSchedule20240805,
	}
}

func (r UpdateBackupScheduleApiRequest) Execute() (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	return r.ApiService.UpdateBackupScheduleExecute(r)
}

/*
UpdateBackupSchedule Update Cloud Backup Schedule for One Cluster

Updates the cloud backup schedule for one cluster within the specified project. This schedule defines when MongoDB Cloud takes scheduled snapshots and how long it stores those snapshots. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateBackupScheduleApiRequest
*/
func (a *CloudBackupsApiService) UpdateBackupSchedule(ctx context.Context, groupId string, clusterName string, diskBackupSnapshotSchedule20240805 *DiskBackupSnapshotSchedule20240805) UpdateBackupScheduleApiRequest {
	return UpdateBackupScheduleApiRequest{
		ApiService:                         a,
		ctx:                                ctx,
		groupId:                            groupId,
		clusterName:                        clusterName,
		diskBackupSnapshotSchedule20240805: diskBackupSnapshotSchedule20240805,
	}
}

// UpdateBackupScheduleExecute executes the request
//
//	@return DiskBackupSnapshotSchedule20240805
func (a *CloudBackupsApiService) UpdateBackupScheduleExecute(r UpdateBackupScheduleApiRequest) (*DiskBackupSnapshotSchedule20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupSnapshotSchedule20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateBackupSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.diskBackupSnapshotSchedule20240805 == nil {
		return localVarReturnValue, nil, reportError("diskBackupSnapshotSchedule20240805 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.diskBackupSnapshotSchedule20240805
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateDataProtectionSettingsApiRequest struct {
	ctx                            context.Context
	ApiService                     CloudBackupsApi
	groupId                        string
	dataProtectionSettings20231001 *DataProtectionSettings20231001
	overwriteBackupPolicies        *bool
}

type UpdateDataProtectionSettingsApiParams struct {
	GroupId                        string
	DataProtectionSettings20231001 *DataProtectionSettings20231001
	OverwriteBackupPolicies        *bool
}

func (a *CloudBackupsApiService) UpdateDataProtectionSettingsWithParams(ctx context.Context, args *UpdateDataProtectionSettingsApiParams) UpdateDataProtectionSettingsApiRequest {
	return UpdateDataProtectionSettingsApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		dataProtectionSettings20231001: args.DataProtectionSettings20231001,
		overwriteBackupPolicies:        args.OverwriteBackupPolicies,
	}
}

// Flag that indicates whether to overwrite non complying backup policies with the new data protection settings or not.
func (r UpdateDataProtectionSettingsApiRequest) OverwriteBackupPolicies(overwriteBackupPolicies bool) UpdateDataProtectionSettingsApiRequest {
	r.overwriteBackupPolicies = &overwriteBackupPolicies
	return r
}

func (r UpdateDataProtectionSettingsApiRequest) Execute() (*DataProtectionSettings20231001, *http.Response, error) {
	return r.ApiService.UpdateDataProtectionSettingsExecute(r)
}

/*
UpdateDataProtectionSettings Update Backup Compliance Policy Settings

Updates the Backup Compliance Policy settings for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateDataProtectionSettingsApiRequest
*/
func (a *CloudBackupsApiService) UpdateDataProtectionSettings(ctx context.Context, groupId string, dataProtectionSettings20231001 *DataProtectionSettings20231001) UpdateDataProtectionSettingsApiRequest {
	return UpdateDataProtectionSettingsApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		dataProtectionSettings20231001: dataProtectionSettings20231001,
	}
}

// UpdateDataProtectionSettingsExecute executes the request
//
//	@return DataProtectionSettings20231001
func (a *CloudBackupsApiService) UpdateDataProtectionSettingsExecute(r UpdateDataProtectionSettingsApiRequest) (*DataProtectionSettings20231001, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DataProtectionSettings20231001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateDataProtectionSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/backupCompliancePolicy"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataProtectionSettings20231001 == nil {
		return localVarReturnValue, nil, reportError("dataProtectionSettings20231001 is required and must be specified")
	}

	if r.overwriteBackupPolicies != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "overwriteBackupPolicies", r.overwriteBackupPolicies, "")
	} else {
		var defaultValue bool = true
		r.overwriteBackupPolicies = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "overwriteBackupPolicies", r.overwriteBackupPolicies, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-10-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-10-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.dataProtectionSettings20231001
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
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
	return r.ApiService.UpdateSnapshotRetentionExecute(r)
}

/*
UpdateSnapshotRetention Update Expiration Date for One Cloud Backup

Changes the expiration date for one cloud backup snapshot for one cluster in the specified project, the requesting Service Account or API Key must have the Project Backup Manager role.

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

// UpdateSnapshotRetentionExecute executes the request
//
//	@return DiskBackupReplicaSet
func (a *CloudBackupsApiService) UpdateSnapshotRetentionExecute(r UpdateSnapshotRetentionApiRequest) (*DiskBackupReplicaSet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DiskBackupReplicaSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudBackupsApiService.UpdateSnapshotRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.snapshotId == "" {
		return localVarReturnValue, nil, reportError("snapshotId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(r.snapshotId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
