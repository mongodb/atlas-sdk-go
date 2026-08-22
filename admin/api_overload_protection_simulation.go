// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OverloadProtectionSimulationAPI interface {

	/*
		CreateClusterOverloadSimulation Create One Overload Protection Simulation

		Starts an overload protection simulation for one cluster. Returns a 409 if a simulation is already in progress or completed; DELETE the existing simulation before starting a new one.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster on which to start the overload protection simulation.
		@param overloadProtectionSimulationRequest Details of the overload protection simulation to start. Valid durations (in seconds): 900, 3600, 28800, 86400.
		@return CreateClusterOverloadSimulationApiRequest
	*/
	CreateClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, overloadProtectionSimulationRequest *OverloadProtectionSimulationRequest) CreateClusterOverloadSimulationApiRequest
	/*
		CreateClusterOverloadSimulation Create One Overload Protection Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateClusterOverloadSimulationApiParams - Parameters for the request
		@return CreateClusterOverloadSimulationApiRequest
	*/
	CreateClusterOverloadSimulationWithParams(ctx context.Context, args *CreateClusterOverloadSimulationApiParams) CreateClusterOverloadSimulationApiRequest

	// Method available only for mocking purposes
	CreateClusterOverloadSimulationExecute(r CreateClusterOverloadSimulationApiRequest) (*OverloadProtectionSimulationResponse, *http.Response, error)

	/*
		DeleteClusterOverloadSimulation Delete One Overload Protection Simulation

		Deletes the overload protection simulation for one cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster on which the overload protection simulation is running.
		@param simulationId Unique 24-hexadecimal digit string that identifies the overload protection simulation.
		@return DeleteClusterOverloadSimulationApiRequest
	*/
	DeleteClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, simulationId string) DeleteClusterOverloadSimulationApiRequest
	/*
		DeleteClusterOverloadSimulation Delete One Overload Protection Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteClusterOverloadSimulationApiParams - Parameters for the request
		@return DeleteClusterOverloadSimulationApiRequest
	*/
	DeleteClusterOverloadSimulationWithParams(ctx context.Context, args *DeleteClusterOverloadSimulationApiParams) DeleteClusterOverloadSimulationApiRequest

	// Method available only for mocking purposes
	DeleteClusterOverloadSimulationExecute(r DeleteClusterOverloadSimulationApiRequest) (*http.Response, error)

	/*
		GetClusterOverloadSimulation Return One Overload Protection Simulation

		Returns the overload protection simulation for one cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster on which the overload protection simulation is running.
		@param simulationId Unique 24-hexadecimal digit string that identifies the overload protection simulation.
		@return GetClusterOverloadSimulationApiRequest
	*/
	GetClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, simulationId string) GetClusterOverloadSimulationApiRequest
	/*
		GetClusterOverloadSimulation Return One Overload Protection Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterOverloadSimulationApiParams - Parameters for the request
		@return GetClusterOverloadSimulationApiRequest
	*/
	GetClusterOverloadSimulationWithParams(ctx context.Context, args *GetClusterOverloadSimulationApiParams) GetClusterOverloadSimulationApiRequest

	// Method available only for mocking purposes
	GetClusterOverloadSimulationExecute(r GetClusterOverloadSimulationApiRequest) (*OverloadProtectionSimulationResponse, *http.Response, error)

	/*
		ListClusterOverloadSimulations Return All Overload Protection Simulations

		Returns all overload protection simulations for one cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster on which the overload protection simulations are running.
		@return ListClusterOverloadSimulationsApiRequest
	*/
	ListClusterOverloadSimulations(ctx context.Context, groupId string, clusterName string) ListClusterOverloadSimulationsApiRequest
	/*
		ListClusterOverloadSimulations Return All Overload Protection Simulations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListClusterOverloadSimulationsApiParams - Parameters for the request
		@return ListClusterOverloadSimulationsApiRequest
	*/
	ListClusterOverloadSimulationsWithParams(ctx context.Context, args *ListClusterOverloadSimulationsApiParams) ListClusterOverloadSimulationsApiRequest

	// Method available only for mocking purposes
	ListClusterOverloadSimulationsExecute(r ListClusterOverloadSimulationsApiRequest) (*PaginatedOverloadProtectionSimulationResponse, *http.Response, error)
}

// OverloadProtectionSimulationAPIService OverloadProtectionSimulationAPI service
type OverloadProtectionSimulationAPIService service

type CreateClusterOverloadSimulationApiRequest struct {
	ctx                                 context.Context
	ApiService                          OverloadProtectionSimulationAPI
	groupId                             string
	clusterName                         string
	overloadProtectionSimulationRequest *OverloadProtectionSimulationRequest
}

type CreateClusterOverloadSimulationApiParams struct {
	GroupId                             string
	ClusterName                         string
	OverloadProtectionSimulationRequest *OverloadProtectionSimulationRequest
}

func (a *OverloadProtectionSimulationAPIService) CreateClusterOverloadSimulationWithParams(ctx context.Context, args *CreateClusterOverloadSimulationApiParams) CreateClusterOverloadSimulationApiRequest {
	return CreateClusterOverloadSimulationApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             args.GroupId,
		clusterName:                         args.ClusterName,
		overloadProtectionSimulationRequest: args.OverloadProtectionSimulationRequest,
	}
}

func (r CreateClusterOverloadSimulationApiRequest) Execute() (*OverloadProtectionSimulationResponse, *http.Response, error) {
	return r.ApiService.CreateClusterOverloadSimulationExecute(r)
}

/*
CreateClusterOverloadSimulation Create One Overload Protection Simulation

Starts an overload protection simulation for one cluster. Returns a 409 if a simulation is already in progress or completed; DELETE the existing simulation before starting a new one.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster on which to start the overload protection simulation.
	@return CreateClusterOverloadSimulationApiRequest
*/
func (a *OverloadProtectionSimulationAPIService) CreateClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, overloadProtectionSimulationRequest *OverloadProtectionSimulationRequest) CreateClusterOverloadSimulationApiRequest {
	return CreateClusterOverloadSimulationApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             groupId,
		clusterName:                         clusterName,
		overloadProtectionSimulationRequest: overloadProtectionSimulationRequest,
	}
}

// CreateClusterOverloadSimulationExecute executes the request
//
//	@return OverloadProtectionSimulationResponse
func (a *OverloadProtectionSimulationAPIService) CreateClusterOverloadSimulationExecute(r CreateClusterOverloadSimulationApiRequest) (*OverloadProtectionSimulationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OverloadProtectionSimulationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OverloadProtectionSimulationAPIService.CreateClusterOverloadSimulation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations"
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
	if r.overloadProtectionSimulationRequest == nil {
		return localVarReturnValue, nil, reportError("overloadProtectionSimulationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.overloadProtectionSimulationRequest
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

type DeleteClusterOverloadSimulationApiRequest struct {
	ctx          context.Context
	ApiService   OverloadProtectionSimulationAPI
	groupId      string
	clusterName  string
	simulationId string
}

type DeleteClusterOverloadSimulationApiParams struct {
	GroupId      string
	ClusterName  string
	SimulationId string
}

func (a *OverloadProtectionSimulationAPIService) DeleteClusterOverloadSimulationWithParams(ctx context.Context, args *DeleteClusterOverloadSimulationApiParams) DeleteClusterOverloadSimulationApiRequest {
	return DeleteClusterOverloadSimulationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		simulationId: args.SimulationId,
	}
}

func (r DeleteClusterOverloadSimulationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClusterOverloadSimulationExecute(r)
}

/*
DeleteClusterOverloadSimulation Delete One Overload Protection Simulation

Deletes the overload protection simulation for one cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster on which the overload protection simulation is running.
	@param simulationId Unique 24-hexadecimal digit string that identifies the overload protection simulation.
	@return DeleteClusterOverloadSimulationApiRequest
*/
func (a *OverloadProtectionSimulationAPIService) DeleteClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, simulationId string) DeleteClusterOverloadSimulationApiRequest {
	return DeleteClusterOverloadSimulationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		clusterName:  clusterName,
		simulationId: simulationId,
	}
}

// DeleteClusterOverloadSimulationExecute executes the request
func (a *OverloadProtectionSimulationAPIService) DeleteClusterOverloadSimulationExecute(r DeleteClusterOverloadSimulationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OverloadProtectionSimulationAPIService.DeleteClusterOverloadSimulation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations/{simulationId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.simulationId == "" {
		return nil, reportError("simulationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"simulationId"+"}", url.PathEscape(r.simulationId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetClusterOverloadSimulationApiRequest struct {
	ctx          context.Context
	ApiService   OverloadProtectionSimulationAPI
	groupId      string
	clusterName  string
	simulationId string
}

type GetClusterOverloadSimulationApiParams struct {
	GroupId      string
	ClusterName  string
	SimulationId string
}

func (a *OverloadProtectionSimulationAPIService) GetClusterOverloadSimulationWithParams(ctx context.Context, args *GetClusterOverloadSimulationApiParams) GetClusterOverloadSimulationApiRequest {
	return GetClusterOverloadSimulationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		simulationId: args.SimulationId,
	}
}

func (r GetClusterOverloadSimulationApiRequest) Execute() (*OverloadProtectionSimulationResponse, *http.Response, error) {
	return r.ApiService.GetClusterOverloadSimulationExecute(r)
}

/*
GetClusterOverloadSimulation Return One Overload Protection Simulation

Returns the overload protection simulation for one cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster on which the overload protection simulation is running.
	@param simulationId Unique 24-hexadecimal digit string that identifies the overload protection simulation.
	@return GetClusterOverloadSimulationApiRequest
*/
func (a *OverloadProtectionSimulationAPIService) GetClusterOverloadSimulation(ctx context.Context, groupId string, clusterName string, simulationId string) GetClusterOverloadSimulationApiRequest {
	return GetClusterOverloadSimulationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		clusterName:  clusterName,
		simulationId: simulationId,
	}
}

// GetClusterOverloadSimulationExecute executes the request
//
//	@return OverloadProtectionSimulationResponse
func (a *OverloadProtectionSimulationAPIService) GetClusterOverloadSimulationExecute(r GetClusterOverloadSimulationApiRequest) (*OverloadProtectionSimulationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OverloadProtectionSimulationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OverloadProtectionSimulationAPIService.GetClusterOverloadSimulation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations/{simulationId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.simulationId == "" {
		return localVarReturnValue, nil, reportError("simulationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"simulationId"+"}", url.PathEscape(r.simulationId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type ListClusterOverloadSimulationsApiRequest struct {
	ctx          context.Context
	ApiService   OverloadProtectionSimulationAPI
	groupId      string
	clusterName  string
	itemsPerPage *int
	pageNum      *int
}

type ListClusterOverloadSimulationsApiParams struct {
	GroupId      string
	ClusterName  string
	ItemsPerPage *int
	PageNum      *int
}

func (a *OverloadProtectionSimulationAPIService) ListClusterOverloadSimulationsWithParams(ctx context.Context, args *ListClusterOverloadSimulationsApiParams) ListClusterOverloadSimulationsApiRequest {
	return ListClusterOverloadSimulationsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListClusterOverloadSimulationsApiRequest) ItemsPerPage(itemsPerPage int) ListClusterOverloadSimulationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListClusterOverloadSimulationsApiRequest) PageNum(pageNum int) ListClusterOverloadSimulationsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListClusterOverloadSimulationsApiRequest) Execute() (*PaginatedOverloadProtectionSimulationResponse, *http.Response, error) {
	return r.ApiService.ListClusterOverloadSimulationsExecute(r)
}

/*
ListClusterOverloadSimulations Return All Overload Protection Simulations

Returns all overload protection simulations for one cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster on which the overload protection simulations are running.
	@return ListClusterOverloadSimulationsApiRequest
*/
func (a *OverloadProtectionSimulationAPIService) ListClusterOverloadSimulations(ctx context.Context, groupId string, clusterName string) ListClusterOverloadSimulationsApiRequest {
	return ListClusterOverloadSimulationsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListClusterOverloadSimulationsExecute executes the request
//
//	@return PaginatedOverloadProtectionSimulationResponse
func (a *OverloadProtectionSimulationAPIService) ListClusterOverloadSimulationsExecute(r ListClusterOverloadSimulationsApiRequest) (*PaginatedOverloadProtectionSimulationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedOverloadProtectionSimulationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OverloadProtectionSimulationAPIService.ListClusterOverloadSimulations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/overloadSimulations"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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
