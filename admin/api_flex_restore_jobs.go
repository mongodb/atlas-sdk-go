// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FlexRestoreJobsApi interface {

	/*
		CreateFlexBackupRestoreJob Restore One Snapshot of One Flex Cluster

		Restores one snapshot of one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster whose snapshot you want to restore.
		@param flexBackupRestoreJobCreate20241113 Restores one snapshot of one flex cluster from the specified project.
		@return CreateFlexBackupRestoreJobApiRequest
	*/
	CreateFlexBackupRestoreJob(ctx context.Context, groupId string, name string, flexBackupRestoreJobCreate20241113 *FlexBackupRestoreJobCreate20241113) CreateFlexBackupRestoreJobApiRequest
	/*
		CreateFlexBackupRestoreJob Restore One Snapshot of One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateFlexBackupRestoreJobApiParams - Parameters for the request
		@return CreateFlexBackupRestoreJobApiRequest
	*/
	CreateFlexBackupRestoreJobWithParams(ctx context.Context, args *CreateFlexBackupRestoreJobApiParams) CreateFlexBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	CreateFlexBackupRestoreJobExecute(r CreateFlexBackupRestoreJobApiRequest) (*FlexBackupRestoreJob20241113, *http.Response, error)

	/*
		GetFlexBackupRestoreJob Return One Restore Job for One Flex Cluster

		Returns one restore job for one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
		@return GetFlexBackupRestoreJobApiRequest
	*/
	GetFlexBackupRestoreJob(ctx context.Context, groupId string, name string, restoreJobId string) GetFlexBackupRestoreJobApiRequest
	/*
		GetFlexBackupRestoreJob Return One Restore Job for One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetFlexBackupRestoreJobApiParams - Parameters for the request
		@return GetFlexBackupRestoreJobApiRequest
	*/
	GetFlexBackupRestoreJobWithParams(ctx context.Context, args *GetFlexBackupRestoreJobApiParams) GetFlexBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	GetFlexBackupRestoreJobExecute(r GetFlexBackupRestoreJobApiRequest) (*FlexBackupRestoreJob20241113, *http.Response, error)

	/*
		ListFlexBackupRestoreJobs Return All Restore Jobs for One Flex Cluster

		Returns all restore jobs for one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@return ListFlexBackupRestoreJobsApiRequest
	*/
	ListFlexBackupRestoreJobs(ctx context.Context, groupId string, name string) ListFlexBackupRestoreJobsApiRequest
	/*
		ListFlexBackupRestoreJobs Return All Restore Jobs for One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListFlexBackupRestoreJobsApiParams - Parameters for the request
		@return ListFlexBackupRestoreJobsApiRequest
	*/
	ListFlexBackupRestoreJobsWithParams(ctx context.Context, args *ListFlexBackupRestoreJobsApiParams) ListFlexBackupRestoreJobsApiRequest

	// Method available only for mocking purposes
	ListFlexBackupRestoreJobsExecute(r ListFlexBackupRestoreJobsApiRequest) (*PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error)
}

// FlexRestoreJobsApiService FlexRestoreJobsApi service
type FlexRestoreJobsApiService service

type CreateFlexBackupRestoreJobApiRequest struct {
	ctx                                context.Context
	ApiService                         FlexRestoreJobsApi
	groupId                            string
	name                               string
	flexBackupRestoreJobCreate20241113 *FlexBackupRestoreJobCreate20241113
}

type CreateFlexBackupRestoreJobApiParams struct {
	GroupId                            string
	Name                               string
	FlexBackupRestoreJobCreate20241113 *FlexBackupRestoreJobCreate20241113
}

func (a *FlexRestoreJobsApiService) CreateFlexBackupRestoreJobWithParams(ctx context.Context, args *CreateFlexBackupRestoreJobApiParams) CreateFlexBackupRestoreJobApiRequest {
	return CreateFlexBackupRestoreJobApiRequest{
		ApiService:                         a,
		ctx:                                ctx,
		groupId:                            args.GroupId,
		name:                               args.Name,
		flexBackupRestoreJobCreate20241113: args.FlexBackupRestoreJobCreate20241113,
	}
}

func (r CreateFlexBackupRestoreJobApiRequest) Execute() (*FlexBackupRestoreJob20241113, *http.Response, error) {
	return r.ApiService.CreateFlexBackupRestoreJobExecute(r)
}

/*
CreateFlexBackupRestoreJob Restore One Snapshot of One Flex Cluster

Restores one snapshot of one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster whose snapshot you want to restore.
	@return CreateFlexBackupRestoreJobApiRequest
*/
func (a *FlexRestoreJobsApiService) CreateFlexBackupRestoreJob(ctx context.Context, groupId string, name string, flexBackupRestoreJobCreate20241113 *FlexBackupRestoreJobCreate20241113) CreateFlexBackupRestoreJobApiRequest {
	return CreateFlexBackupRestoreJobApiRequest{
		ApiService:                         a,
		ctx:                                ctx,
		groupId:                            groupId,
		name:                               name,
		flexBackupRestoreJobCreate20241113: flexBackupRestoreJobCreate20241113,
	}
}

// CreateFlexBackupRestoreJobExecute executes the request
//
//	@return FlexBackupRestoreJob20241113
func (a *FlexRestoreJobsApiService) CreateFlexBackupRestoreJobExecute(r CreateFlexBackupRestoreJobApiRequest) (*FlexBackupRestoreJob20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexBackupRestoreJob20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexRestoreJobsApiService.CreateFlexBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.name == "" {
		return localVarReturnValue, nil, reportError("name is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexBackupRestoreJobCreate20241113 == nil {
		return localVarReturnValue, nil, reportError("flexBackupRestoreJobCreate20241113 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-11-13+json"}

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
	// body params
	localVarPostBody = r.flexBackupRestoreJobCreate20241113
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

type GetFlexBackupRestoreJobApiRequest struct {
	ctx          context.Context
	ApiService   FlexRestoreJobsApi
	groupId      string
	name         string
	restoreJobId string
}

type GetFlexBackupRestoreJobApiParams struct {
	GroupId      string
	Name         string
	RestoreJobId string
}

func (a *FlexRestoreJobsApiService) GetFlexBackupRestoreJobWithParams(ctx context.Context, args *GetFlexBackupRestoreJobApiParams) GetFlexBackupRestoreJobApiRequest {
	return GetFlexBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		name:         args.Name,
		restoreJobId: args.RestoreJobId,
	}
}

func (r GetFlexBackupRestoreJobApiRequest) Execute() (*FlexBackupRestoreJob20241113, *http.Response, error) {
	return r.ApiService.GetFlexBackupRestoreJobExecute(r)
}

/*
GetFlexBackupRestoreJob Return One Restore Job for One Flex Cluster

Returns one restore job for one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@param restoreJobId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return GetFlexBackupRestoreJobApiRequest
*/
func (a *FlexRestoreJobsApiService) GetFlexBackupRestoreJob(ctx context.Context, groupId string, name string, restoreJobId string) GetFlexBackupRestoreJobApiRequest {
	return GetFlexBackupRestoreJobApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		name:         name,
		restoreJobId: restoreJobId,
	}
}

// GetFlexBackupRestoreJobExecute executes the request
//
//	@return FlexBackupRestoreJob20241113
func (a *FlexRestoreJobsApiService) GetFlexBackupRestoreJobExecute(r GetFlexBackupRestoreJobApiRequest) (*FlexBackupRestoreJob20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexBackupRestoreJob20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexRestoreJobsApiService.GetFlexBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs/{restoreJobId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.name == "" {
		return localVarReturnValue, nil, reportError("name is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-11-13+json"}

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

type ListFlexBackupRestoreJobsApiRequest struct {
	ctx          context.Context
	ApiService   FlexRestoreJobsApi
	groupId      string
	name         string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListFlexBackupRestoreJobsApiParams struct {
	GroupId      string
	Name         string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *FlexRestoreJobsApiService) ListFlexBackupRestoreJobsWithParams(ctx context.Context, args *ListFlexBackupRestoreJobsApiParams) ListFlexBackupRestoreJobsApiRequest {
	return ListFlexBackupRestoreJobsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		name:         args.Name,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListFlexBackupRestoreJobsApiRequest) IncludeCount(includeCount bool) ListFlexBackupRestoreJobsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListFlexBackupRestoreJobsApiRequest) ItemsPerPage(itemsPerPage int) ListFlexBackupRestoreJobsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListFlexBackupRestoreJobsApiRequest) PageNum(pageNum int) ListFlexBackupRestoreJobsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListFlexBackupRestoreJobsApiRequest) Execute() (*PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error) {
	return r.ApiService.ListFlexBackupRestoreJobsExecute(r)
}

/*
ListFlexBackupRestoreJobs Return All Restore Jobs for One Flex Cluster

Returns all restore jobs for one flex cluster from the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return ListFlexBackupRestoreJobsApiRequest
*/
func (a *FlexRestoreJobsApiService) ListFlexBackupRestoreJobs(ctx context.Context, groupId string, name string) ListFlexBackupRestoreJobsApiRequest {
	return ListFlexBackupRestoreJobsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// ListFlexBackupRestoreJobsExecute executes the request
//
//	@return PaginatedApiAtlasFlexBackupRestoreJob20241113
func (a *FlexRestoreJobsApiService) ListFlexBackupRestoreJobsExecute(r ListFlexBackupRestoreJobsApiRequest) (*PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasFlexBackupRestoreJob20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexRestoreJobsApiService.ListFlexBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}/backup/restoreJobs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.name == "" {
		return localVarReturnValue, nil, reportError("name is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-11-13+json"}

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
