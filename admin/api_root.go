// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type RootApi interface {

	/*
		GetSystemStatus Return the status of this MongoDB application

		This resource returns information about the MongoDB application along with API key meta data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetSystemStatusApiRequest
	*/
	GetSystemStatus(ctx context.Context) GetSystemStatusApiRequest
	/*
		GetSystemStatus Return the status of this MongoDB application


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSystemStatusApiParams - Parameters for the request
		@return GetSystemStatusApiRequest
	*/
	GetSystemStatusWithParams(ctx context.Context, args *GetSystemStatusApiParams) GetSystemStatusApiRequest

	// Method available only for mocking purposes
	GetSystemStatusExecute(r GetSystemStatusApiRequest) (*SystemStatus, *http.Response, error)

	/*
		ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses

		Returns all control plane IP addresses. Currently, inbound Atlas control plane IP addresses are not yet available. The inbound IP address list in your API response is empty. To manually retrieve a list of inbound Atlas control plane IP addresses, see [Required Inbound Access](https://www.mongodb.com/docs/atlas/setup-cluster-security/#std-label-atlas-required-inbound-access).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ReturnAllControlPlaneIPAddressesApiRequest
	*/
	ReturnAllControlPlaneIPAddresses(ctx context.Context) ReturnAllControlPlaneIPAddressesApiRequest
	/*
		ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnAllControlPlaneIPAddressesApiParams - Parameters for the request
		@return ReturnAllControlPlaneIPAddressesApiRequest
	*/
	ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *ReturnAllControlPlaneIPAddressesApiParams) ReturnAllControlPlaneIPAddressesApiRequest

	// Method available only for mocking purposes
	ReturnAllControlPlaneIPAddressesExecute(r ReturnAllControlPlaneIPAddressesApiRequest) (*ControlPlaneIPAddresses, *http.Response, error)
}

// RootApiService RootApi service
type RootApiService service

type GetSystemStatusApiRequest struct {
	ctx        context.Context
	ApiService RootApi
}

type GetSystemStatusApiParams struct {
}

func (a *RootApiService) GetSystemStatusWithParams(ctx context.Context, args *GetSystemStatusApiParams) GetSystemStatusApiRequest {
	return GetSystemStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r GetSystemStatusApiRequest) Execute() (*SystemStatus, *http.Response, error) {
	return r.ApiService.GetSystemStatusExecute(r)
}

/*
GetSystemStatus Return the status of this MongoDB application

This resource returns information about the MongoDB application along with API key meta data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetSystemStatusApiRequest
*/
func (a *RootApiService) GetSystemStatus(ctx context.Context) GetSystemStatusApiRequest {
	return GetSystemStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// GetSystemStatusExecute executes the request
//
//	@return SystemStatus
func (a *RootApiService) GetSystemStatusExecute(r GetSystemStatusApiRequest) (*SystemStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SystemStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootApiService.GetSystemStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2"

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

type ReturnAllControlPlaneIPAddressesApiRequest struct {
	ctx        context.Context
	ApiService RootApi
}

type ReturnAllControlPlaneIPAddressesApiParams struct {
}

func (a *RootApiService) ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *ReturnAllControlPlaneIPAddressesApiParams) ReturnAllControlPlaneIPAddressesApiRequest {
	return ReturnAllControlPlaneIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ReturnAllControlPlaneIPAddressesApiRequest) Execute() (*ControlPlaneIPAddresses, *http.Response, error) {
	return r.ApiService.ReturnAllControlPlaneIPAddressesExecute(r)
}

/*
ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses

Returns all control plane IP addresses. Currently, inbound Atlas control plane IP addresses are not yet available. The inbound IP address list in your API response is empty. To manually retrieve a list of inbound Atlas control plane IP addresses, see [Required Inbound Access](https://www.mongodb.com/docs/atlas/setup-cluster-security/#std-label-atlas-required-inbound-access).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReturnAllControlPlaneIPAddressesApiRequest
*/
func (a *RootApiService) ReturnAllControlPlaneIPAddresses(ctx context.Context) ReturnAllControlPlaneIPAddressesApiRequest {
	return ReturnAllControlPlaneIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ReturnAllControlPlaneIPAddressesExecute executes the request
//
//	@return ControlPlaneIPAddresses
func (a *RootApiService) ReturnAllControlPlaneIPAddressesExecute(r ReturnAllControlPlaneIPAddressesApiRequest) (*ControlPlaneIPAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ControlPlaneIPAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootApiService.ReturnAllControlPlaneIPAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/unauth/controlPlaneIPAddresses"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

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
