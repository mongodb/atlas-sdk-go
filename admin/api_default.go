// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type DefaultApi interface {

	/*
		ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses

		[experimental] Returns all control plane IP addresses.

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

// DefaultApiService DefaultApi service
type DefaultApiService service

type ReturnAllControlPlaneIPAddressesApiRequest struct {
	ctx        context.Context
	ApiService DefaultApi
}

type ReturnAllControlPlaneIPAddressesApiParams struct {
}

func (a *DefaultApiService) ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *ReturnAllControlPlaneIPAddressesApiParams) ReturnAllControlPlaneIPAddressesApiRequest {
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

[experimental] Returns all control plane IP addresses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReturnAllControlPlaneIPAddressesApiRequest
*/
func (a *DefaultApiService) ReturnAllControlPlaneIPAddresses(ctx context.Context) ReturnAllControlPlaneIPAddressesApiRequest {
	return ReturnAllControlPlaneIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ControlPlaneIPAddresses
func (a *DefaultApiService) ReturnAllControlPlaneIPAddressesExecute(r ReturnAllControlPlaneIPAddressesApiRequest) (*ControlPlaneIPAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ControlPlaneIPAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ReturnAllControlPlaneIPAddresses")
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
