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
		GetOpenapiInfo Return General Information on MongoDB Atlas Administration API OpenAPI Specification

		This resource returns general information about the MongoDB Atlas Administration API OpenAPI Specification. Deprecated versions: v2-{2024-05-30}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetOpenapiInfoApiRequest
	*/
	GetOpenapiInfo(ctx context.Context) GetOpenapiInfoApiRequest
	/*
		GetOpenapiInfo Return General Information on MongoDB Atlas Administration API OpenAPI Specification


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOpenapiInfoApiParams - Parameters for the request
		@return GetOpenapiInfoApiRequest
	*/
	GetOpenapiInfoWithParams(ctx context.Context, args *GetOpenapiInfoApiParams) GetOpenapiInfoApiRequest

	// Method available only for mocking purposes
	GetOpenapiInfoExecute(r GetOpenapiInfoApiRequest) (*OpenApiInfo, *http.Response, error)

	/*
		ListOpenapiVersions Return All Versions for One Environment

		API that provides a list of available versionsfor a given environment.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListOpenapiVersionsApiRequest
	*/
	ListOpenapiVersions(ctx context.Context) ListOpenapiVersionsApiRequest
	/*
		ListOpenapiVersions Return All Versions for One Environment


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOpenapiVersionsApiParams - Parameters for the request
		@return ListOpenapiVersionsApiRequest
	*/
	ListOpenapiVersionsWithParams(ctx context.Context, args *ListOpenapiVersionsApiParams) ListOpenapiVersionsApiRequest

	// Method available only for mocking purposes
	ListOpenapiVersionsExecute(r ListOpenapiVersionsApiRequest) (*PaginatedApiVersions, *http.Response, error)
}

// OpenAPIApiService OpenAPIApi service
type OpenAPIApiService service

type GetOpenapiInfoApiRequest struct {
	ctx        context.Context
	ApiService OpenAPIApi
}

type GetOpenapiInfoApiParams struct {
}

func (a *OpenAPIApiService) GetOpenapiInfoWithParams(ctx context.Context, args *GetOpenapiInfoApiParams) GetOpenapiInfoApiRequest {
	return GetOpenapiInfoApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r GetOpenapiInfoApiRequest) Execute() (*OpenApiInfo, *http.Response, error) {
	return r.ApiService.GetOpenapiInfoExecute(r)
}

/*
GetOpenapiInfo Return General Information on MongoDB Atlas Administration API OpenAPI Specification

This resource returns general information about the MongoDB Atlas Administration API OpenAPI Specification. Deprecated versions: v2-{2024-05-30}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetOpenapiInfoApiRequest
*/
func (a *OpenAPIApiService) GetOpenapiInfo(ctx context.Context) GetOpenapiInfoApiRequest {
	return GetOpenapiInfoApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// GetOpenapiInfoExecute executes the request
//
//	@return OpenApiInfo
func (a *OpenAPIApiService) GetOpenapiInfoExecute(r GetOpenapiInfoApiRequest) (*OpenApiInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OpenApiInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenAPIApiService.GetOpenapiInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/openapi/info"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type ListOpenapiVersionsApiRequest struct {
	ctx          context.Context
	ApiService   OpenAPIApi
	itemsPerPage *int
	pageNum      *int
	env          *string
}

type ListOpenapiVersionsApiParams struct {
	ItemsPerPage *int
	PageNum      *int
	Env          *string
}

func (a *OpenAPIApiService) ListOpenapiVersionsWithParams(ctx context.Context, args *ListOpenapiVersionsApiParams) ListOpenapiVersionsApiRequest {
	return ListOpenapiVersionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		env:          args.Env,
	}
}

// Number of items that the response returns per page.
func (r ListOpenapiVersionsApiRequest) ItemsPerPage(itemsPerPage int) ListOpenapiVersionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOpenapiVersionsApiRequest) PageNum(pageNum int) ListOpenapiVersionsApiRequest {
	r.pageNum = &pageNum
	return r
}

// The environment to get the versions from. If not provided, it returnsthe versions for the given MongoDB URL. (E.g. prod for cloud.mongodb.com).
func (r ListOpenapiVersionsApiRequest) Env(env string) ListOpenapiVersionsApiRequest {
	r.env = &env
	return r
}

func (r ListOpenapiVersionsApiRequest) Execute() (*PaginatedApiVersions, *http.Response, error) {
	return r.ApiService.ListOpenapiVersionsExecute(r)
}

/*
ListOpenapiVersions Return All Versions for One Environment

API that provides a list of available versionsfor a given environment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListOpenapiVersionsApiRequest
*/
func (a *OpenAPIApiService) ListOpenapiVersions(ctx context.Context) ListOpenapiVersionsApiRequest {
	return ListOpenapiVersionsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListOpenapiVersionsExecute executes the request
//
//	@return PaginatedApiVersions
func (a *OpenAPIApiService) ListOpenapiVersionsExecute(r ListOpenapiVersionsApiRequest) (*PaginatedApiVersions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiVersions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenAPIApiService.ListOpenapiVersions")
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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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
