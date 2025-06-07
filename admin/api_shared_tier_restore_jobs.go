// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SharedTierRestoreJobsApi interface {

	/*
			CreateSharedClusterBackupRestoreJob Create One Restore Job for One M2 or M5 Cluster

			Restores the specified M2 or M5 cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the createFlexBackupRestoreJob endpoint instead.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clusterName Human-readable label that identifies the cluster.
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param tenantRestore The restore job details.
			@return CreateSharedClusterBackupRestoreJobApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	CreateSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) CreateSharedClusterBackupRestoreJobApiRequest
	/*
		CreateSharedClusterBackupRestoreJob Create One Restore Job for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateSharedClusterBackupRestoreJobApiParams - Parameters for the request
		@return CreateSharedClusterBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	CreateSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *CreateSharedClusterBackupRestoreJobApiParams) CreateSharedClusterBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	CreateSharedClusterBackupRestoreJobExecute(r CreateSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error)

	/*
			GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster

			Returns the specified restore job for the specified M2 or M5 cluster. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the getFlexBackupRestoreJob endpoint instead.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clusterName Human-readable label that identifies the cluster.
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
			@return GetSharedClusterBackupRestoreJobApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	GetSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, restoreId string) GetSharedClusterBackupRestoreJobApiRequest
	/*
		GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSharedClusterBackupRestoreJobApiParams - Parameters for the request
		@return GetSharedClusterBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	GetSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *GetSharedClusterBackupRestoreJobApiParams) GetSharedClusterBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	GetSharedClusterBackupRestoreJobExecute(r GetSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error)

	/*
			ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster

			Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the listFlexBackupRestoreJobs endpoint instead.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clusterName Human-readable label that identifies the cluster.
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListSharedClusterBackupRestoreJobsApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	ListSharedClusterBackupRestoreJobs(ctx context.Context, clusterName string, groupId string) ListSharedClusterBackupRestoreJobsApiRequest
	/*
		ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSharedClusterBackupRestoreJobsApiParams - Parameters for the request
		@return ListSharedClusterBackupRestoreJobsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for SharedTierRestoreJobsApi
	*/
	ListSharedClusterBackupRestoreJobsWithParams(ctx context.Context, args *ListSharedClusterBackupRestoreJobsApiParams) ListSharedClusterBackupRestoreJobsApiRequest

	// Method available only for mocking purposes
	ListSharedClusterBackupRestoreJobsExecute(r ListSharedClusterBackupRestoreJobsApiRequest) (*PaginatedTenantRestore, *http.Response, error)
}

// SharedTierRestoreJobsApiService SharedTierRestoreJobsApi service
type SharedTierRestoreJobsApiService service

type CreateSharedClusterBackupRestoreJobApiRequest struct {
	ctx           context.Context
	ApiService    SharedTierRestoreJobsApi
	clusterName   string
	groupId       string
	tenantRestore *TenantRestore
}

type CreateSharedClusterBackupRestoreJobApiParams struct {
	ClusterName   string
	GroupId       string
	TenantRestore *TenantRestore
}

func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *CreateSharedClusterBackupRestoreJobApiParams) CreateSharedClusterBackupRestoreJobApiRequest {
	return CreateSharedClusterBackupRestoreJobApiRequest{
		ApiService:    a,
		ctx:           ctx,
		clusterName:   args.ClusterName,
		groupId:       args.GroupId,
		tenantRestore: args.TenantRestore,
	}
}

func (r CreateSharedClusterBackupRestoreJobApiRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.CreateSharedClusterBackupRestoreJobExecute(r)
}

/*
CreateSharedClusterBackupRestoreJob Create One Restore Job for One M2 or M5 Cluster

Restores the specified M2 or M5 cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the createFlexBackupRestoreJob endpoint instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateSharedClusterBackupRestoreJobApiRequest

Deprecated
*/
func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) CreateSharedClusterBackupRestoreJobApiRequest {
	return CreateSharedClusterBackupRestoreJobApiRequest{
		ApiService:    a,
		ctx:           ctx,
		clusterName:   clusterName,
		groupId:       groupId,
		tenantRestore: tenantRestore,
	}
}

// CreateSharedClusterBackupRestoreJobExecute executes the request
//
//	@return TenantRestore
//
// Deprecated
func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJobExecute(r CreateSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.CreateSharedClusterBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore"
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantRestore == nil {
		return localVarReturnValue, nil, reportError("tenantRestore is required and must be specified")
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
	localVarPostBody = r.tenantRestore
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

type GetSharedClusterBackupRestoreJobApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierRestoreJobsApi
	clusterName string
	groupId     string
	restoreId   string
}

type GetSharedClusterBackupRestoreJobApiParams struct {
	ClusterName string
	GroupId     string
	RestoreId   string
}

func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *GetSharedClusterBackupRestoreJobApiParams) GetSharedClusterBackupRestoreJobApiRequest {
	return GetSharedClusterBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: args.ClusterName,
		groupId:     args.GroupId,
		restoreId:   args.RestoreId,
	}
}

func (r GetSharedClusterBackupRestoreJobApiRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.GetSharedClusterBackupRestoreJobExecute(r)
}

/*
GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster

Returns the specified restore job for the specified M2 or M5 cluster. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the getFlexBackupRestoreJob endpoint instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return GetSharedClusterBackupRestoreJobApiRequest

Deprecated
*/
func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, restoreId string) GetSharedClusterBackupRestoreJobApiRequest {
	return GetSharedClusterBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: clusterName,
		groupId:     groupId,
		restoreId:   restoreId,
	}
}

// GetSharedClusterBackupRestoreJobExecute executes the request
//
//	@return TenantRestore
//
// Deprecated
func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJobExecute(r GetSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.GetSharedClusterBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId}"
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.restoreId == "" {
		return localVarReturnValue, nil, reportError("restoreId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"restoreId"+"}", url.PathEscape(r.restoreId), -1)

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

type ListSharedClusterBackupRestoreJobsApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierRestoreJobsApi
	clusterName string
	groupId     string
}

type ListSharedClusterBackupRestoreJobsApiParams struct {
	ClusterName string
	GroupId     string
}

func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobsWithParams(ctx context.Context, args *ListSharedClusterBackupRestoreJobsApiParams) ListSharedClusterBackupRestoreJobsApiRequest {
	return ListSharedClusterBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: args.ClusterName,
		groupId:     args.GroupId,
	}
}

func (r ListSharedClusterBackupRestoreJobsApiRequest) Execute() (*PaginatedTenantRestore, *http.Response, error) {
	return r.ApiService.ListSharedClusterBackupRestoreJobsExecute(r)
}

/*
ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster

Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026, after which this endpoint will be sunset. Please use the listFlexBackupRestoreJobs endpoint instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListSharedClusterBackupRestoreJobsApiRequest

Deprecated
*/
func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobs(ctx context.Context, clusterName string, groupId string) ListSharedClusterBackupRestoreJobsApiRequest {
	return ListSharedClusterBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: clusterName,
		groupId:     groupId,
	}
}

// ListSharedClusterBackupRestoreJobsExecute executes the request
//
//	@return PaginatedTenantRestore
//
// Deprecated
func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobsExecute(r ListSharedClusterBackupRestoreJobsApiRequest) (*PaginatedTenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedTenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.ListSharedClusterBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores"
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
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
