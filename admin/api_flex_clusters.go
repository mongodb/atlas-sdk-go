// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FlexClustersApi interface {

	/*
		CreateFlexcluster Create One Flex Cluster in One Project

		Creates one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param flexClusterDescriptionCreate20250101 Create One Flex Cluster in One Project.
		@return CreateFlexclusterApiRequest
	*/
	CreateFlexcluster(ctx context.Context, groupId string, flexClusterDescriptionCreate20250101 *FlexClusterDescriptionCreate20250101) CreateFlexclusterApiRequest
	/*
		CreateFlexcluster Create One Flex Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateFlexclusterApiParams - Parameters for the request
		@return CreateFlexclusterApiRequest
	*/
	CreateFlexclusterWithParams(ctx context.Context, args *CreateFlexclusterApiParams) CreateFlexclusterApiRequest

	// Method available only for mocking purposes
	CreateFlexclusterExecute(r CreateFlexclusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error)

	/*
		DeleteFlexCluster Remove One Flex Cluster from One Project

		Removes one flex cluster from the specified project. The flex cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@return DeleteFlexClusterApiRequest
	*/
	DeleteFlexCluster(ctx context.Context, groupId string, name string) DeleteFlexClusterApiRequest
	/*
		DeleteFlexCluster Remove One Flex Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteFlexClusterApiParams - Parameters for the request
		@return DeleteFlexClusterApiRequest
	*/
	DeleteFlexClusterWithParams(ctx context.Context, args *DeleteFlexClusterApiParams) DeleteFlexClusterApiRequest

	// Method available only for mocking purposes
	DeleteFlexClusterExecute(r DeleteFlexClusterApiRequest) (any, *http.Response, error)

	/*
		GetFlexCluster Return One Flex Cluster from One Project

		Returns details for one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@return GetFlexClusterApiRequest
	*/
	GetFlexCluster(ctx context.Context, groupId string, name string) GetFlexClusterApiRequest
	/*
		GetFlexCluster Return One Flex Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetFlexClusterApiParams - Parameters for the request
		@return GetFlexClusterApiRequest
	*/
	GetFlexClusterWithParams(ctx context.Context, args *GetFlexClusterApiParams) GetFlexClusterApiRequest

	// Method available only for mocking purposes
	GetFlexClusterExecute(r GetFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error)

	/*
		ListFlexClusters Return All Flex Clusters from One Project

		Returns details for all flex clusters in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListFlexClustersApiRequest
	*/
	ListFlexClusters(ctx context.Context, groupId string) ListFlexClustersApiRequest
	/*
		ListFlexClusters Return All Flex Clusters from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListFlexClustersApiParams - Parameters for the request
		@return ListFlexClustersApiRequest
	*/
	ListFlexClustersWithParams(ctx context.Context, args *ListFlexClustersApiParams) ListFlexClustersApiRequest

	// Method available only for mocking purposes
	ListFlexClustersExecute(r ListFlexClustersApiRequest) (*PaginatedFlexClusters20250101, *http.Response, error)

	/*
		UpdateFlexCluster Update One Flex Cluster in One Project

		Updates one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@param flexClusterDescription20250101 Update One Flex Cluster in One Project.
		@return UpdateFlexClusterApiRequest
	*/
	UpdateFlexCluster(ctx context.Context, groupId string, name string, flexClusterDescription20250101 *FlexClusterDescription20250101) UpdateFlexClusterApiRequest
	/*
		UpdateFlexCluster Update One Flex Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateFlexClusterApiParams - Parameters for the request
		@return UpdateFlexClusterApiRequest
	*/
	UpdateFlexClusterWithParams(ctx context.Context, args *UpdateFlexClusterApiParams) UpdateFlexClusterApiRequest

	// Method available only for mocking purposes
	UpdateFlexClusterExecute(r UpdateFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error)

	/*
		UpgradeFlexCluster Upgrade One Flex Cluster

		Upgrades a flex cluster in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param flexClusterDescription20250101 Details of the flex cluster upgrade in the specified project.
		@return UpgradeFlexClusterApiRequest
	*/
	UpgradeFlexCluster(ctx context.Context, groupId string, flexClusterDescription20250101 *FlexClusterDescription20250101) UpgradeFlexClusterApiRequest
	/*
		UpgradeFlexCluster Upgrade One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeFlexClusterApiParams - Parameters for the request
		@return UpgradeFlexClusterApiRequest
	*/
	UpgradeFlexClusterWithParams(ctx context.Context, args *UpgradeFlexClusterApiParams) UpgradeFlexClusterApiRequest

	// Method available only for mocking purposes
	UpgradeFlexClusterExecute(r UpgradeFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error)
}

// FlexClustersApiService FlexClustersApi service
type FlexClustersApiService service

type CreateFlexclusterApiRequest struct {
	ctx                                  context.Context
	ApiService                           FlexClustersApi
	groupId                              string
	flexClusterDescriptionCreate20250101 *FlexClusterDescriptionCreate20250101
}

type CreateFlexclusterApiParams struct {
	GroupId                              string
	FlexClusterDescriptionCreate20250101 *FlexClusterDescriptionCreate20250101
}

func (a *FlexClustersApiService) CreateFlexclusterWithParams(ctx context.Context, args *CreateFlexclusterApiParams) CreateFlexclusterApiRequest {
	return CreateFlexclusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		flexClusterDescriptionCreate20250101: args.FlexClusterDescriptionCreate20250101,
	}
}

func (r CreateFlexclusterApiRequest) Execute() (*FlexClusterDescription20250101, *http.Response, error) {
	return r.ApiService.CreateFlexclusterExecute(r)
}

/*
CreateFlexcluster Create One Flex Cluster in One Project

Creates one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateFlexclusterApiRequest
*/
func (a *FlexClustersApiService) CreateFlexcluster(ctx context.Context, groupId string, flexClusterDescriptionCreate20250101 *FlexClusterDescriptionCreate20250101) CreateFlexclusterApiRequest {
	return CreateFlexclusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		flexClusterDescriptionCreate20250101: flexClusterDescriptionCreate20250101,
	}
}

// CreateFlexclusterExecute executes the request
//
//	@return FlexClusterDescription20250101
func (a *FlexClustersApiService) CreateFlexclusterExecute(r CreateFlexclusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.CreateFlexcluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexClusterDescriptionCreate20250101 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescriptionCreate20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.flexClusterDescriptionCreate20250101
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

type DeleteFlexClusterApiRequest struct {
	ctx        context.Context
	ApiService FlexClustersApi
	groupId    string
	name       string
}

type DeleteFlexClusterApiParams struct {
	GroupId string
	Name    string
}

func (a *FlexClustersApiService) DeleteFlexClusterWithParams(ctx context.Context, args *DeleteFlexClusterApiParams) DeleteFlexClusterApiRequest {
	return DeleteFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r DeleteFlexClusterApiRequest) Execute() (any, *http.Response, error) {
	return r.ApiService.DeleteFlexClusterExecute(r)
}

/*
DeleteFlexCluster Remove One Flex Cluster from One Project

Removes one flex cluster from the specified project. The flex cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return DeleteFlexClusterApiRequest
*/
func (a *FlexClustersApiService) DeleteFlexCluster(ctx context.Context, groupId string, name string) DeleteFlexClusterApiRequest {
	return DeleteFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// DeleteFlexClusterExecute executes the request
//
//	@return any
func (a *FlexClustersApiService) DeleteFlexClusterExecute(r DeleteFlexClusterApiRequest) (any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.DeleteFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type GetFlexClusterApiRequest struct {
	ctx        context.Context
	ApiService FlexClustersApi
	groupId    string
	name       string
}

type GetFlexClusterApiParams struct {
	GroupId string
	Name    string
}

func (a *FlexClustersApiService) GetFlexClusterWithParams(ctx context.Context, args *GetFlexClusterApiParams) GetFlexClusterApiRequest {
	return GetFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r GetFlexClusterApiRequest) Execute() (*FlexClusterDescription20250101, *http.Response, error) {
	return r.ApiService.GetFlexClusterExecute(r)
}

/*
GetFlexCluster Return One Flex Cluster from One Project

Returns details for one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return GetFlexClusterApiRequest
*/
func (a *FlexClustersApiService) GetFlexCluster(ctx context.Context, groupId string, name string) GetFlexClusterApiRequest {
	return GetFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// GetFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20250101
func (a *FlexClustersApiService) GetFlexClusterExecute(r GetFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.GetFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type ListFlexClustersApiRequest struct {
	ctx          context.Context
	ApiService   FlexClustersApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListFlexClustersApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *FlexClustersApiService) ListFlexClustersWithParams(ctx context.Context, args *ListFlexClustersApiParams) ListFlexClustersApiRequest {
	return ListFlexClustersApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListFlexClustersApiRequest) IncludeCount(includeCount bool) ListFlexClustersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListFlexClustersApiRequest) ItemsPerPage(itemsPerPage int) ListFlexClustersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListFlexClustersApiRequest) PageNum(pageNum int) ListFlexClustersApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListFlexClustersApiRequest) Execute() (*PaginatedFlexClusters20250101, *http.Response, error) {
	return r.ApiService.ListFlexClustersExecute(r)
}

/*
ListFlexClusters Return All Flex Clusters from One Project

Returns details for all flex clusters in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListFlexClustersApiRequest
*/
func (a *FlexClustersApiService) ListFlexClusters(ctx context.Context, groupId string) ListFlexClustersApiRequest {
	return ListFlexClustersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListFlexClustersExecute executes the request
//
//	@return PaginatedFlexClusters20250101
func (a *FlexClustersApiService) ListFlexClustersExecute(r ListFlexClustersApiRequest) (*PaginatedFlexClusters20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedFlexClusters20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.ListFlexClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters"
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type UpdateFlexClusterApiRequest struct {
	ctx                            context.Context
	ApiService                     FlexClustersApi
	groupId                        string
	name                           string
	flexClusterDescription20250101 *FlexClusterDescription20250101
}

type UpdateFlexClusterApiParams struct {
	GroupId                        string
	Name                           string
	FlexClusterDescription20250101 *FlexClusterDescription20250101
}

func (a *FlexClustersApiService) UpdateFlexClusterWithParams(ctx context.Context, args *UpdateFlexClusterApiParams) UpdateFlexClusterApiRequest {
	return UpdateFlexClusterApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		name:                           args.Name,
		flexClusterDescription20250101: args.FlexClusterDescription20250101,
	}
}

func (r UpdateFlexClusterApiRequest) Execute() (*FlexClusterDescription20250101, *http.Response, error) {
	return r.ApiService.UpdateFlexClusterExecute(r)
}

/*
UpdateFlexCluster Update One Flex Cluster in One Project

Updates one flex cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return UpdateFlexClusterApiRequest
*/
func (a *FlexClustersApiService) UpdateFlexCluster(ctx context.Context, groupId string, name string, flexClusterDescription20250101 *FlexClusterDescription20250101) UpdateFlexClusterApiRequest {
	return UpdateFlexClusterApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		name:                           name,
		flexClusterDescription20250101: flexClusterDescription20250101,
	}
}

// UpdateFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20250101
func (a *FlexClustersApiService) UpdateFlexClusterExecute(r UpdateFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.UpdateFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexClusterDescription20250101 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescription20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.flexClusterDescription20250101
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

type UpgradeFlexClusterApiRequest struct {
	ctx                            context.Context
	ApiService                     FlexClustersApi
	groupId                        string
	flexClusterDescription20250101 *FlexClusterDescription20250101
}

type UpgradeFlexClusterApiParams struct {
	GroupId                        string
	FlexClusterDescription20250101 *FlexClusterDescription20250101
}

func (a *FlexClustersApiService) UpgradeFlexClusterWithParams(ctx context.Context, args *UpgradeFlexClusterApiParams) UpgradeFlexClusterApiRequest {
	return UpgradeFlexClusterApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		flexClusterDescription20250101: args.FlexClusterDescription20250101,
	}
}

func (r UpgradeFlexClusterApiRequest) Execute() (*FlexClusterDescription20250101, *http.Response, error) {
	return r.ApiService.UpgradeFlexClusterExecute(r)
}

/*
UpgradeFlexCluster Upgrade One Flex Cluster

Upgrades a flex cluster in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeFlexClusterApiRequest
*/
func (a *FlexClustersApiService) UpgradeFlexCluster(ctx context.Context, groupId string, flexClusterDescription20250101 *FlexClusterDescription20250101) UpgradeFlexClusterApiRequest {
	return UpgradeFlexClusterApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		flexClusterDescription20250101: flexClusterDescription20250101,
	}
}

// UpgradeFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20250101
func (a *FlexClustersApiService) UpgradeFlexClusterExecute(r UpgradeFlexClusterApiRequest) (*FlexClusterDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersApiService.UpgradeFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters:tenantUpgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexClusterDescription20250101 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescription20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.flexClusterDescription20250101
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
