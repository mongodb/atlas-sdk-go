// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type TestApi interface {

	/*
		VersionedExample Example resource info for versioning of the Atlas API

		Returns some text dummy data for test purposes. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return VersionedExampleApiRequest
	*/
	VersionedExample(ctx context.Context) VersionedExampleApiRequest
	/*
		VersionedExample Example resource info for versioning of the Atlas API


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param VersionedExampleApiParams - Parameters for the request
		@return VersionedExampleApiRequest
	*/
	VersionedExampleWithParams(ctx context.Context, args *VersionedExampleApiParams) VersionedExampleApiRequest

	// Interface only available internally
	versionedExampleExecute(r VersionedExampleApiRequest) (*ExampleResourceResponseView20230201, *http.Response, error)
}

// TestApiService TestApi service
type TestApiService service

type VersionedExampleApiRequest struct {
	ctx            context.Context
	ApiService     TestApi
	additionalInfo *bool
}

type VersionedExampleApiParams struct {
	AdditionalInfo *bool
}

func (a *TestApiService) VersionedExampleWithParams(ctx context.Context, args *VersionedExampleApiParams) VersionedExampleApiRequest {
	return VersionedExampleApiRequest{
		ApiService:     a,
		ctx:            ctx,
		additionalInfo: args.AdditionalInfo,
	}
}

func (r VersionedExampleApiRequest) AdditionalInfo(additionalInfo bool) VersionedExampleApiRequest {
	r.additionalInfo = &additionalInfo
	return r
}

func (r VersionedExampleApiRequest) Execute() (*ExampleResourceResponseView20230201, *http.Response, error) {
	return r.ApiService.versionedExampleExecute(r)
}

/*
VersionedExample Example resource info for versioning of the Atlas API

Returns some text dummy data for test purposes. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return VersionedExampleApiRequest
*/
func (a *TestApiService) VersionedExample(ctx context.Context) VersionedExampleApiRequest {
	return VersionedExampleApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExampleResourceResponseView20230201
func (a *TestApiService) versionedExampleExecute(r VersionedExampleApiRequest) (*ExampleResourceResponseView20230201, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExampleResourceResponseView20230201
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestApiService.VersionedExample")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/example/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.additionalInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalInfo", r.additionalInfo, "")
	} else {
		var defaultValue bool = false
		r.additionalInfo = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalInfo", r.additionalInfo, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
