// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type TestApi interface {

	/*
		GetExampleInfo Example Resource Information for Versioning of the Atlas API

		Returns some text dummy data for test purposes.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetExampleInfoApiRequest
	*/
	GetExampleInfo(ctx context.Context) GetExampleInfoApiRequest
	/*
		GetExampleInfo Example Resource Information for Versioning of the Atlas API


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetExampleInfoApiParams - Parameters for the request
		@return GetExampleInfoApiRequest
	*/
	GetExampleInfoWithParams(ctx context.Context, args *GetExampleInfoApiParams) GetExampleInfoApiRequest

	// Method available only for mocking purposes
	GetExampleInfoExecute(r GetExampleInfoApiRequest) (*ExampleResourceResponse20230201, *http.Response, error)
}

// TestApiService TestApi service
type TestApiService service

type GetExampleInfoApiRequest struct {
	ctx            context.Context
	ApiService     TestApi
	additionalInfo *string
}

type GetExampleInfoApiParams struct {
	AdditionalInfo *string
}

func (a *TestApiService) GetExampleInfoWithParams(ctx context.Context, args *GetExampleInfoApiParams) GetExampleInfoApiRequest {
	return GetExampleInfoApiRequest{
		ApiService:     a,
		ctx:            ctx,
		additionalInfo: args.AdditionalInfo,
	}
}

// Show more info.
func (r GetExampleInfoApiRequest) AdditionalInfo(additionalInfo string) GetExampleInfoApiRequest {
	r.additionalInfo = &additionalInfo
	return r
}

func (r GetExampleInfoApiRequest) Execute() (*ExampleResourceResponse20230201, *http.Response, error) {
	return r.ApiService.GetExampleInfoExecute(r)
}

/*
GetExampleInfo Example Resource Information for Versioning of the Atlas API

Returns some text dummy data for test purposes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetExampleInfoApiRequest
*/
func (a *TestApiService) GetExampleInfo(ctx context.Context) GetExampleInfoApiRequest {
	return GetExampleInfoApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// GetExampleInfoExecute executes the request
//
//	@return ExampleResourceResponse20230201
func (a *TestApiService) GetExampleInfoExecute(r GetExampleInfoApiRequest) (*ExampleResourceResponse20230201, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ExampleResourceResponse20230201
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestApiService.GetExampleInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/example/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.additionalInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalInfo", r.additionalInfo, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

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
