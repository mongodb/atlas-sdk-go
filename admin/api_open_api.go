// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type OpenAPIApi interface {

	/*
		GetApiVersions Provides a list of versions for a given environment.

		API that provides a list of available versionsfor a given environment.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetApiVersionsApiRequest
	*/
	GetApiVersions(ctx context.Context) GetApiVersionsApiRequest
	/*
		GetApiVersions Provides a list of versions for a given environment.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetApiVersionsApiParams - Parameters for the request
		@return GetApiVersionsApiRequest
	*/
	GetApiVersionsWithParams(ctx context.Context, args *GetApiVersionsApiParams) GetApiVersionsApiRequest

	// Method available only for mocking purposes
	GetApiVersionsExecute(r GetApiVersionsApiRequest) (*PaginatedApiVersions, *http.Response, error)
}

// OpenAPIApiService OpenAPIApi service
type OpenAPIApiService service

type GetApiVersionsApiRequest struct {
	ctx          context.Context
	ApiService   OpenAPIApi
	itemsPerPage *int
	pageNum      *int
	env          *string
}

type GetApiVersionsApiParams struct {
	ItemsPerPage *int
	PageNum      *int
	Env          *string
}

func (a *OpenAPIApiService) GetApiVersionsWithParams(ctx context.Context, args *GetApiVersionsApiParams) GetApiVersionsApiRequest {
	return GetApiVersionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		env:          args.Env,
	}
}

// Number of items that the response returns per page.
func (r GetApiVersionsApiRequest) ItemsPerPage(itemsPerPage int) GetApiVersionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r GetApiVersionsApiRequest) PageNum(pageNum int) GetApiVersionsApiRequest {
	r.pageNum = &pageNum
	return r
}

// The environment to get the versions from. If not provided, it returnsthe versions for the given MongoDB URL. (E.g. prod for cloud.mongodb.com)
func (r GetApiVersionsApiRequest) Env(env string) GetApiVersionsApiRequest {
	r.env = &env
	return r
}

func (r GetApiVersionsApiRequest) Execute() (*PaginatedApiVersions, *http.Response, error) {
	return r.ApiService.GetApiVersionsExecute(r)
}

/*
GetApiVersions Provides a list of versions for a given environment.

API that provides a list of available versionsfor a given environment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetApiVersionsApiRequest
*/
func (a *OpenAPIApiService) GetApiVersions(ctx context.Context) GetApiVersionsApiRequest {
	return GetApiVersionsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// GetApiVersionsExecute executes the request
//
//	@return PaginatedApiVersions
func (a *OpenAPIApiService) GetApiVersionsExecute(r GetApiVersionsApiRequest) (*PaginatedApiVersions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiVersions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenAPIApiService.GetApiVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/unauth/openapi/versions"

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
	if r.env != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json", "application/vnd.atlas.2024-08-05+yaml", "application/json"}

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
