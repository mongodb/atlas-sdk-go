// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type LegacyBackupRestoreJobsApi interface {

	/*
		CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job

		Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). This endpoint doesn't support creating checkpoint restore jobs for sharded clusters, or creating restore jobs for queryable backup snapshots. If you create an automated restore job by specifying `delivery.methodName` of `AUTOMATED_RESTORE` in your request body, MongoDB Cloud removes all existing data on the target cluster prior to the restore.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
		@return CreateLegacyBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupRestoreJobsApi
	*/
	CreateLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string) CreateLegacyBackupRestoreJobApiRequest
	/*
		CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateLegacyBackupRestoreJobApiParams - Parameters for the request
		@return CreateLegacyBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupRestoreJobsApi
	*/
	CreateLegacyBackupRestoreJobWithParams(ctx context.Context, args *CreateLegacyBackupRestoreJobApiParams) CreateLegacyBackupRestoreJobApiRequest

	// Interface only available internally
	createLegacyBackupRestoreJobExecute(r CreateLegacyBackupRestoreJobApiRequest) (*PaginatedRestoreJob, *http.Response, error)
}

// LegacyBackupRestoreJobsApiService LegacyBackupRestoreJobsApi service
type LegacyBackupRestoreJobsApiService service

type CreateLegacyBackupRestoreJobApiRequest struct {
	ctx         context.Context
	ApiService  LegacyBackupRestoreJobsApi
	groupId     string
	clusterName string
	restoreJob  *RestoreJob
}

type CreateLegacyBackupRestoreJobApiParams struct {
	GroupId     string
	ClusterName string
	RestoreJob  *RestoreJob
}

func (a *LegacyBackupRestoreJobsApiService) CreateLegacyBackupRestoreJobWithParams(ctx context.Context, args *CreateLegacyBackupRestoreJobApiParams) CreateLegacyBackupRestoreJobApiRequest {
	return CreateLegacyBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		restoreJob:  args.RestoreJob,
	}
}

// Legacy backup to restore to one cluster in the specified project.
func (r CreateLegacyBackupRestoreJobApiRequest) RestoreJob(restoreJob RestoreJob) CreateLegacyBackupRestoreJobApiRequest {
	r.restoreJob = &restoreJob
	return r
}

func (r CreateLegacyBackupRestoreJobApiRequest) Execute() (*PaginatedRestoreJob, *http.Response, error) {
	return r.ApiService.createLegacyBackupRestoreJobExecute(r)
}

/*
CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job

Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role and an entry for the project access list. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). This endpoint doesn't support creating checkpoint restore jobs for sharded clusters, or creating restore jobs for queryable backup snapshots. If you create an automated restore job by specifying `delivery.methodName` of `AUTOMATED_RESTORE` in your request body, MongoDB Cloud removes all existing data on the target cluster prior to the restore.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return CreateLegacyBackupRestoreJobApiRequest

Deprecated
*/
func (a *LegacyBackupRestoreJobsApiService) CreateLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string) CreateLegacyBackupRestoreJobApiRequest {
	return CreateLegacyBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedRestoreJob
//
// Deprecated
func (a *LegacyBackupRestoreJobsApiService) createLegacyBackupRestoreJobExecute(r CreateLegacyBackupRestoreJobApiRequest) (*PaginatedRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupRestoreJobsApiService.CreateLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs"
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
	if r.restoreJob == nil {
		return localVarReturnValue, nil, reportError("restoreJob is required and must be specified")
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
	localVarPostBody = r.restoreJob
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
		var v Error
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
