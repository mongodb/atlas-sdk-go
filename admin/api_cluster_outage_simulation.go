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

type ClusterOutageSimulationAPI interface {

	/*
		EndOutageSimulation End an Outage Simulation

		[experimental] Ends a cluster outage simulation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that is undergoing outage simulation.
		@return EndOutageSimulationApiRequest
	*/
	EndOutageSimulation(ctx context.Context, groupId string, clusterName string) EndOutageSimulationApiRequest
	/*
		EndOutageSimulation End an Outage Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param EndOutageSimulationApiParams - Parameters for the request
		@return EndOutageSimulationApiRequest
	*/
	EndOutageSimulationWithParams(ctx context.Context, args *EndOutageSimulationApiParams) EndOutageSimulationApiRequest

	// Interface only available internally
	endOutageSimulationExecute(r EndOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error)

	/*
		GetOutageSimulation Return One Outage Simulation

		[experimental] Returns one outage simulation for one cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that is undergoing outage simulation.
		@return GetOutageSimulationApiRequest
	*/
	GetOutageSimulation(ctx context.Context, groupId string, clusterName string) GetOutageSimulationApiRequest
	/*
		GetOutageSimulation Return One Outage Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOutageSimulationApiParams - Parameters for the request
		@return GetOutageSimulationApiRequest
	*/
	GetOutageSimulationWithParams(ctx context.Context, args *GetOutageSimulationApiParams) GetOutageSimulationApiRequest

	// Interface only available internally
	getOutageSimulationExecute(r GetOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error)

	/*
		StartOutageSimulation Start an Outage Simulation

		[experimental] Starts a cluster outage simulation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to undergo an outage simulation.
		@return StartOutageSimulationApiRequest
	*/
	StartOutageSimulation(ctx context.Context, groupId string, clusterName string, clusterOutageSimulation *ClusterOutageSimulation) StartOutageSimulationApiRequest
	/*
		StartOutageSimulation Start an Outage Simulation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param StartOutageSimulationApiParams - Parameters for the request
		@return StartOutageSimulationApiRequest
	*/
	StartOutageSimulationWithParams(ctx context.Context, args *StartOutageSimulationApiParams) StartOutageSimulationApiRequest

	// Interface only available internally
	startOutageSimulationExecute(r StartOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error)
}

// ClusterOutageSimulationAPIService ClusterOutageSimulationAPI service
type ClusterOutageSimulationAPIService service

type EndOutageSimulationApiRequest struct {
	ctx         context.Context
	ApiService  ClusterOutageSimulationAPI
	groupId     string
	clusterName string
}

type EndOutageSimulationApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClusterOutageSimulationAPIService) EndOutageSimulationWithParams(ctx context.Context, args *EndOutageSimulationApiParams) EndOutageSimulationApiRequest {
	return EndOutageSimulationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r EndOutageSimulationApiRequest) Execute() (*ClusterOutageSimulation, *http.Response, error) {
	return r.ApiService.endOutageSimulationExecute(r)
}

/*
EndOutageSimulation End an Outage Simulation

[experimental] Ends a cluster outage simulation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that is undergoing outage simulation.
	@return EndOutageSimulationApiRequest
*/
func (a *ClusterOutageSimulationAPIService) EndOutageSimulation(ctx context.Context, groupId string, clusterName string) EndOutageSimulationApiRequest {
	return EndOutageSimulationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ClusterOutageSimulation
func (a *ClusterOutageSimulationAPIService) endOutageSimulationExecute(r EndOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterOutageSimulation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterOutageSimulationAPIService.EndOutageSimulation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation"
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

type GetOutageSimulationApiRequest struct {
	ctx         context.Context
	ApiService  ClusterOutageSimulationAPI
	groupId     string
	clusterName string
}

type GetOutageSimulationApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClusterOutageSimulationAPIService) GetOutageSimulationWithParams(ctx context.Context, args *GetOutageSimulationApiParams) GetOutageSimulationApiRequest {
	return GetOutageSimulationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetOutageSimulationApiRequest) Execute() (*ClusterOutageSimulation, *http.Response, error) {
	return r.ApiService.getOutageSimulationExecute(r)
}

/*
GetOutageSimulation Return One Outage Simulation

[experimental] Returns one outage simulation for one cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that is undergoing outage simulation.
	@return GetOutageSimulationApiRequest
*/
func (a *ClusterOutageSimulationAPIService) GetOutageSimulation(ctx context.Context, groupId string, clusterName string) GetOutageSimulationApiRequest {
	return GetOutageSimulationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ClusterOutageSimulation
func (a *ClusterOutageSimulationAPIService) getOutageSimulationExecute(r GetOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterOutageSimulation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterOutageSimulationAPIService.GetOutageSimulation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation"
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

type StartOutageSimulationApiRequest struct {
	ctx                     context.Context
	ApiService              ClusterOutageSimulationAPI
	groupId                 string
	clusterName             string
	clusterOutageSimulation *ClusterOutageSimulation
}

type StartOutageSimulationApiParams struct {
	GroupId                 string
	ClusterName             string
	ClusterOutageSimulation *ClusterOutageSimulation
}

func (a *ClusterOutageSimulationAPIService) StartOutageSimulationWithParams(ctx context.Context, args *StartOutageSimulationApiParams) StartOutageSimulationApiRequest {
	return StartOutageSimulationApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		clusterName:             args.ClusterName,
		clusterOutageSimulation: args.ClusterOutageSimulation,
	}
}

func (r StartOutageSimulationApiRequest) Execute() (*ClusterOutageSimulation, *http.Response, error) {
	return r.ApiService.startOutageSimulationExecute(r)
}

/*
StartOutageSimulation Start an Outage Simulation

[experimental] Starts a cluster outage simulation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to undergo an outage simulation.
	@return StartOutageSimulationApiRequest
*/
func (a *ClusterOutageSimulationAPIService) StartOutageSimulation(ctx context.Context, groupId string, clusterName string, clusterOutageSimulation *ClusterOutageSimulation) StartOutageSimulationApiRequest {
	return StartOutageSimulationApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		clusterName:             clusterName,
		clusterOutageSimulation: clusterOutageSimulation,
	}
}

// Execute executes the request
//
//	@return ClusterOutageSimulation
func (a *ClusterOutageSimulationAPIService) startOutageSimulationExecute(r StartOutageSimulationApiRequest) (*ClusterOutageSimulation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterOutageSimulation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusterOutageSimulationAPIService.StartOutageSimulation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation"
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
	if r.clusterOutageSimulation == nil {
		return localVarReturnValue, nil, reportError("clusterOutageSimulation is required and must be specified")
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
	localVarPostBody = r.clusterOutageSimulation
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
