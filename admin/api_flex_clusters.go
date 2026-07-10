// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FlexClustersAPI interface {

	/*
		CreateFlexCluster Create One Flex Cluster in One Project

		Creates one flex cluster in the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param flexClusterDescriptionCreate20241113 Create One Flex Cluster in One Project.
		@return CreateFlexClusterApiRequest
	*/
	CreateFlexCluster(ctx context.Context, groupId string, flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113) CreateFlexClusterApiRequest
	/*
		CreateFlexCluster Create One Flex Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateFlexClusterApiParams - Parameters for the request
		@return CreateFlexClusterApiRequest
	*/
	CreateFlexClusterWithParams(ctx context.Context, args *CreateFlexClusterApiParams) CreateFlexClusterApiRequest

	// Method available only for mocking purposes
	CreateFlexClusterExecute(r CreateFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error)

	/*
		DeleteFlexCluster Remove One Flex Cluster from One Project

		Removes one flex cluster from the specified project. The flex cluster must have termination protection disabled in order to be deleted.

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
	DeleteFlexClusterExecute(r DeleteFlexClusterApiRequest) (*http.Response, error)

	/*
			EstimateCreateFlexCost Estimate the Cost of Creating One Flex Cluster

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 This endpoint validates a supplied cluster configuration and returns an estimated hourly and monthly compute cost without creating or modifying the cluster. It applies the same validation as the Create One Flex Cluster endpoint, so invalid configurations are rejected with the same errors. The estimate excludes backup storage, disk storage, data transfer, and premium feature surcharges, which are billed based on actual usage. Subscription, payment, and resource policy checks are not performed, so a configuration that can be priced here may still be rejected during the actual operation.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param flexClusterDescriptionCreate20241113 Flex cluster configuration to estimate.
			@return EstimateCreateFlexCostApiRequest
	*/
	EstimateCreateFlexCost(ctx context.Context, groupId string, flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113) EstimateCreateFlexCostApiRequest
	/*
		EstimateCreateFlexCost Estimate the Cost of Creating One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param EstimateCreateFlexCostApiParams - Parameters for the request
		@return EstimateCreateFlexCostApiRequest
	*/
	EstimateCreateFlexCostWithParams(ctx context.Context, args *EstimateCreateFlexCostApiParams) EstimateCreateFlexCostApiRequest

	// Method available only for mocking purposes
	EstimateCreateFlexCostExecute(r EstimateCreateFlexCostApiRequest) (*ClusterCostEstimate, *http.Response, error)

	/*
			EstimateUpdateFlexCost Estimate the Cost of Updating One Flex Cluster

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 This endpoint validates a supplied cluster configuration and returns an estimated hourly and monthly compute cost without creating or modifying the cluster. It applies the same validation as the Update One Flex Cluster endpoint, so invalid configurations are rejected with the same errors. The estimate excludes backup storage, disk storage, data transfer, and premium feature surcharges, which are billed based on actual usage. Subscription, payment, and resource policy checks are not performed, so a configuration that can be priced here may still be rejected during the actual operation.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the flex cluster.
			@param flexClusterDescriptionUpdate20241113 Flex cluster update to estimate pricing for.
			@return EstimateUpdateFlexCostApiRequest
	*/
	EstimateUpdateFlexCost(ctx context.Context, groupId string, clusterName string, flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113) EstimateUpdateFlexCostApiRequest
	/*
		EstimateUpdateFlexCost Estimate the Cost of Updating One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param EstimateUpdateFlexCostApiParams - Parameters for the request
		@return EstimateUpdateFlexCostApiRequest
	*/
	EstimateUpdateFlexCostWithParams(ctx context.Context, args *EstimateUpdateFlexCostApiParams) EstimateUpdateFlexCostApiRequest

	// Method available only for mocking purposes
	EstimateUpdateFlexCostExecute(r EstimateUpdateFlexCostApiRequest) (*ClusterCostEstimate, *http.Response, error)

	/*
		GetFlexCluster Return One Flex Cluster from One Project

		Returns details for one flex cluster in the specified project.

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
	GetFlexClusterExecute(r GetFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error)

	/*
		ListFlexClusters Return All Flex Clusters from One Project

		Returns details for all flex clusters in the specified project.

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
	ListFlexClustersExecute(r ListFlexClustersApiRequest) (*PaginatedFlexClusters20241113, *http.Response, error)

	/*
		TenantUpgrade Upgrade One Flex Cluster

		Upgrades a flex cluster to a dedicated cluster (M10+) in the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param atlasTenantClusterUpgradeRequest20240805 Details of the flex cluster upgrade in the specified project.
		@return TenantUpgradeApiRequest
	*/
	TenantUpgrade(ctx context.Context, groupId string, atlasTenantClusterUpgradeRequest20240805 *AtlasTenantClusterUpgradeRequest20240805) TenantUpgradeApiRequest
	/*
		TenantUpgrade Upgrade One Flex Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param TenantUpgradeApiParams - Parameters for the request
		@return TenantUpgradeApiRequest
	*/
	TenantUpgradeWithParams(ctx context.Context, args *TenantUpgradeApiParams) TenantUpgradeApiRequest

	// Method available only for mocking purposes
	TenantUpgradeExecute(r TenantUpgradeApiRequest) (*FlexClusterDescription20241113, *http.Response, error)

	/*
		UpdateFlexCluster Update One Flex Cluster in One Project

		Updates one flex cluster in the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the flex cluster.
		@param flexClusterDescriptionUpdate20241113 Update One Flex Cluster in One Project.
		@return UpdateFlexClusterApiRequest
	*/
	UpdateFlexCluster(ctx context.Context, groupId string, name string, flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113) UpdateFlexClusterApiRequest
	/*
		UpdateFlexCluster Update One Flex Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateFlexClusterApiParams - Parameters for the request
		@return UpdateFlexClusterApiRequest
	*/
	UpdateFlexClusterWithParams(ctx context.Context, args *UpdateFlexClusterApiParams) UpdateFlexClusterApiRequest

	// Method available only for mocking purposes
	UpdateFlexClusterExecute(r UpdateFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error)
}

// FlexClustersAPIService FlexClustersAPI service
type FlexClustersAPIService service

type CreateFlexClusterApiRequest struct {
	ctx                                  context.Context
	ApiService                           FlexClustersAPI
	groupId                              string
	flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113
}

type CreateFlexClusterApiParams struct {
	GroupId                              string
	FlexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113
}

func (a *FlexClustersAPIService) CreateFlexClusterWithParams(ctx context.Context, args *CreateFlexClusterApiParams) CreateFlexClusterApiRequest {
	return CreateFlexClusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		flexClusterDescriptionCreate20241113: args.FlexClusterDescriptionCreate20241113,
	}
}

func (r CreateFlexClusterApiRequest) Execute() (*FlexClusterDescription20241113, *http.Response, error) {
	return r.ApiService.CreateFlexClusterExecute(r)
}

/*
CreateFlexCluster Create One Flex Cluster in One Project

Creates one flex cluster in the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateFlexClusterApiRequest
*/
func (a *FlexClustersAPIService) CreateFlexCluster(ctx context.Context, groupId string, flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113) CreateFlexClusterApiRequest {
	return CreateFlexClusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		flexClusterDescriptionCreate20241113: flexClusterDescriptionCreate20241113,
	}
}

// CreateFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20241113
func (a *FlexClustersAPIService) CreateFlexClusterExecute(r CreateFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.CreateFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexClusterDescriptionCreate20241113 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescriptionCreate20241113 is required and must be specified")
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
	localVarPostBody = r.flexClusterDescriptionCreate20241113
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
	ApiService FlexClustersAPI
	groupId    string
	name       string
}

type DeleteFlexClusterApiParams struct {
	GroupId string
	Name    string
}

func (a *FlexClustersAPIService) DeleteFlexClusterWithParams(ctx context.Context, args *DeleteFlexClusterApiParams) DeleteFlexClusterApiRequest {
	return DeleteFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r DeleteFlexClusterApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFlexClusterExecute(r)
}

/*
DeleteFlexCluster Remove One Flex Cluster from One Project

Removes one flex cluster from the specified project. The flex cluster must have termination protection disabled in order to be deleted.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return DeleteFlexClusterApiRequest
*/
func (a *FlexClustersAPIService) DeleteFlexCluster(ctx context.Context, groupId string, name string) DeleteFlexClusterApiRequest {
	return DeleteFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// DeleteFlexClusterExecute executes the request
func (a *FlexClustersAPIService) DeleteFlexClusterExecute(r DeleteFlexClusterApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.DeleteFlexCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.name == "" {
		return nil, reportError("name is empty and must be specified")
	}
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

type EstimateCreateFlexCostApiRequest struct {
	ctx                                  context.Context
	ApiService                           FlexClustersAPI
	groupId                              string
	flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113
}

type EstimateCreateFlexCostApiParams struct {
	GroupId                              string
	FlexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113
}

func (a *FlexClustersAPIService) EstimateCreateFlexCostWithParams(ctx context.Context, args *EstimateCreateFlexCostApiParams) EstimateCreateFlexCostApiRequest {
	return EstimateCreateFlexCostApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		flexClusterDescriptionCreate20241113: args.FlexClusterDescriptionCreate20241113,
	}
}

func (r EstimateCreateFlexCostApiRequest) Execute() (*ClusterCostEstimate, *http.Response, error) {
	return r.ApiService.EstimateCreateFlexCostExecute(r)
}

/*
EstimateCreateFlexCost Estimate the Cost of Creating One Flex Cluster

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	This endpoint validates a supplied cluster configuration and returns an estimated hourly and monthly compute cost without creating or modifying the cluster. It applies the same validation as the Create One Flex Cluster endpoint, so invalid configurations are rejected with the same errors. The estimate excludes backup storage, disk storage, data transfer, and premium feature surcharges, which are billed based on actual usage. Subscription, payment, and resource policy checks are not performed, so a configuration that can be priced here may still be rejected during the actual operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return EstimateCreateFlexCostApiRequest
*/
func (a *FlexClustersAPIService) EstimateCreateFlexCost(ctx context.Context, groupId string, flexClusterDescriptionCreate20241113 *FlexClusterDescriptionCreate20241113) EstimateCreateFlexCostApiRequest {
	return EstimateCreateFlexCostApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		flexClusterDescriptionCreate20241113: flexClusterDescriptionCreate20241113,
	}
}

// EstimateCreateFlexCostExecute executes the request
//
//	@return ClusterCostEstimate
func (a *FlexClustersAPIService) EstimateCreateFlexCostExecute(r EstimateCreateFlexCostApiRequest) (*ClusterCostEstimate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterCostEstimate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.EstimateCreateFlexCost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters:estimateCreateCost"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexClusterDescriptionCreate20241113 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescriptionCreate20241113 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.flexClusterDescriptionCreate20241113
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

type EstimateUpdateFlexCostApiRequest struct {
	ctx                                  context.Context
	ApiService                           FlexClustersAPI
	groupId                              string
	clusterName                          string
	flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113
}

type EstimateUpdateFlexCostApiParams struct {
	GroupId                              string
	ClusterName                          string
	FlexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113
}

func (a *FlexClustersAPIService) EstimateUpdateFlexCostWithParams(ctx context.Context, args *EstimateUpdateFlexCostApiParams) EstimateUpdateFlexCostApiRequest {
	return EstimateUpdateFlexCostApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		clusterName:                          args.ClusterName,
		flexClusterDescriptionUpdate20241113: args.FlexClusterDescriptionUpdate20241113,
	}
}

func (r EstimateUpdateFlexCostApiRequest) Execute() (*ClusterCostEstimate, *http.Response, error) {
	return r.ApiService.EstimateUpdateFlexCostExecute(r)
}

/*
EstimateUpdateFlexCost Estimate the Cost of Updating One Flex Cluster

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	This endpoint validates a supplied cluster configuration and returns an estimated hourly and monthly compute cost without creating or modifying the cluster. It applies the same validation as the Update One Flex Cluster endpoint, so invalid configurations are rejected with the same errors. The estimate excludes backup storage, disk storage, data transfer, and premium feature surcharges, which are billed based on actual usage. Subscription, payment, and resource policy checks are not performed, so a configuration that can be priced here may still be rejected during the actual operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the flex cluster.
	@return EstimateUpdateFlexCostApiRequest
*/
func (a *FlexClustersAPIService) EstimateUpdateFlexCost(ctx context.Context, groupId string, clusterName string, flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113) EstimateUpdateFlexCostApiRequest {
	return EstimateUpdateFlexCostApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		clusterName:                          clusterName,
		flexClusterDescriptionUpdate20241113: flexClusterDescriptionUpdate20241113,
	}
}

// EstimateUpdateFlexCostExecute executes the request
//
//	@return ClusterCostEstimate
func (a *FlexClustersAPIService) EstimateUpdateFlexCostExecute(r EstimateUpdateFlexCostApiRequest) (*ClusterCostEstimate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterCostEstimate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.EstimateUpdateFlexCost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{clusterName}:estimateUpdateCost"
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
	if r.flexClusterDescriptionUpdate20241113 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescriptionUpdate20241113 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.flexClusterDescriptionUpdate20241113
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
	ApiService FlexClustersAPI
	groupId    string
	name       string
}

type GetFlexClusterApiParams struct {
	GroupId string
	Name    string
}

func (a *FlexClustersAPIService) GetFlexClusterWithParams(ctx context.Context, args *GetFlexClusterApiParams) GetFlexClusterApiRequest {
	return GetFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r GetFlexClusterApiRequest) Execute() (*FlexClusterDescription20241113, *http.Response, error) {
	return r.ApiService.GetFlexClusterExecute(r)
}

/*
GetFlexCluster Return One Flex Cluster from One Project

Returns details for one flex cluster in the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return GetFlexClusterApiRequest
*/
func (a *FlexClustersAPIService) GetFlexCluster(ctx context.Context, groupId string, name string) GetFlexClusterApiRequest {
	return GetFlexClusterApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// GetFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20241113
func (a *FlexClustersAPIService) GetFlexClusterExecute(r GetFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.GetFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
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

type ListFlexClustersApiRequest struct {
	ctx          context.Context
	ApiService   FlexClustersAPI
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

func (a *FlexClustersAPIService) ListFlexClustersWithParams(ctx context.Context, args *ListFlexClustersApiParams) ListFlexClustersApiRequest {
	return ListFlexClustersApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
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

func (r ListFlexClustersApiRequest) Execute() (*PaginatedFlexClusters20241113, *http.Response, error) {
	return r.ApiService.ListFlexClustersExecute(r)
}

/*
ListFlexClusters Return All Flex Clusters from One Project

Returns details for all flex clusters in the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListFlexClustersApiRequest
*/
func (a *FlexClustersAPIService) ListFlexClusters(ctx context.Context, groupId string) ListFlexClustersApiRequest {
	return ListFlexClustersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListFlexClustersExecute executes the request
//
//	@return PaginatedFlexClusters20241113
func (a *FlexClustersAPIService) ListFlexClustersExecute(r ListFlexClustersApiRequest) (*PaginatedFlexClusters20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedFlexClusters20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.ListFlexClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters"
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

type TenantUpgradeApiRequest struct {
	ctx                                      context.Context
	ApiService                               FlexClustersAPI
	groupId                                  string
	atlasTenantClusterUpgradeRequest20240805 *AtlasTenantClusterUpgradeRequest20240805
}

type TenantUpgradeApiParams struct {
	GroupId                                  string
	AtlasTenantClusterUpgradeRequest20240805 *AtlasTenantClusterUpgradeRequest20240805
}

func (a *FlexClustersAPIService) TenantUpgradeWithParams(ctx context.Context, args *TenantUpgradeApiParams) TenantUpgradeApiRequest {
	return TenantUpgradeApiRequest{
		ApiService:                               a,
		ctx:                                      ctx,
		groupId:                                  args.GroupId,
		atlasTenantClusterUpgradeRequest20240805: args.AtlasTenantClusterUpgradeRequest20240805,
	}
}

func (r TenantUpgradeApiRequest) Execute() (*FlexClusterDescription20241113, *http.Response, error) {
	return r.ApiService.TenantUpgradeExecute(r)
}

/*
TenantUpgrade Upgrade One Flex Cluster

Upgrades a flex cluster to a dedicated cluster (M10+) in the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return TenantUpgradeApiRequest
*/
func (a *FlexClustersAPIService) TenantUpgrade(ctx context.Context, groupId string, atlasTenantClusterUpgradeRequest20240805 *AtlasTenantClusterUpgradeRequest20240805) TenantUpgradeApiRequest {
	return TenantUpgradeApiRequest{
		ApiService:                               a,
		ctx:                                      ctx,
		groupId:                                  groupId,
		atlasTenantClusterUpgradeRequest20240805: atlasTenantClusterUpgradeRequest20240805,
	}
}

// TenantUpgradeExecute executes the request
//
//	@return FlexClusterDescription20241113
func (a *FlexClustersAPIService) TenantUpgradeExecute(r TenantUpgradeApiRequest) (*FlexClusterDescription20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.TenantUpgrade")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters:tenantUpgrade"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.atlasTenantClusterUpgradeRequest20240805 == nil {
		return localVarReturnValue, nil, reportError("atlasTenantClusterUpgradeRequest20240805 is required and must be specified")
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
	localVarPostBody = r.atlasTenantClusterUpgradeRequest20240805
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
	ctx                                  context.Context
	ApiService                           FlexClustersAPI
	groupId                              string
	name                                 string
	flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113
}

type UpdateFlexClusterApiParams struct {
	GroupId                              string
	Name                                 string
	FlexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113
}

func (a *FlexClustersAPIService) UpdateFlexClusterWithParams(ctx context.Context, args *UpdateFlexClusterApiParams) UpdateFlexClusterApiRequest {
	return UpdateFlexClusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		name:                                 args.Name,
		flexClusterDescriptionUpdate20241113: args.FlexClusterDescriptionUpdate20241113,
	}
}

func (r UpdateFlexClusterApiRequest) Execute() (*FlexClusterDescription20241113, *http.Response, error) {
	return r.ApiService.UpdateFlexClusterExecute(r)
}

/*
UpdateFlexCluster Update One Flex Cluster in One Project

Updates one flex cluster in the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the flex cluster.
	@return UpdateFlexClusterApiRequest
*/
func (a *FlexClustersAPIService) UpdateFlexCluster(ctx context.Context, groupId string, name string, flexClusterDescriptionUpdate20241113 *FlexClusterDescriptionUpdate20241113) UpdateFlexClusterApiRequest {
	return UpdateFlexClusterApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		name:                                 name,
		flexClusterDescriptionUpdate20241113: flexClusterDescriptionUpdate20241113,
	}
}

// UpdateFlexClusterExecute executes the request
//
//	@return FlexClusterDescription20241113
func (a *FlexClustersAPIService) UpdateFlexClusterExecute(r UpdateFlexClusterApiRequest) (*FlexClusterDescription20241113, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *FlexClusterDescription20241113
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexClustersAPIService.UpdateFlexCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/flexClusters/{name}"
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
	if r.flexClusterDescriptionUpdate20241113 == nil {
		return localVarReturnValue, nil, reportError("flexClusterDescriptionUpdate20241113 is required and must be specified")
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
	localVarPostBody = r.flexClusterDescriptionUpdate20241113
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
