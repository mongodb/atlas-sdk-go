// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type LimitDescriptionApi interface {

	/*
		GetDefaultGroupLimit Return One User-Configurable Project Limit and Description

		Returns the description of one user-configurable, project-level limit, along with its default and maximum values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param limitName Identifies the user-managed limit.
		@return GetDefaultGroupLimitApiRequest
	*/
	GetDefaultGroupLimit(ctx context.Context, limitName string) GetDefaultGroupLimitApiRequest
	/*
		GetDefaultGroupLimit Return One User-Configurable Project Limit and Description


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetDefaultGroupLimitApiParams - Parameters for the request
		@return GetDefaultGroupLimitApiRequest
	*/
	GetDefaultGroupLimitWithParams(ctx context.Context, args *GetDefaultGroupLimitApiParams) GetDefaultGroupLimitApiRequest

	// Method available only for mocking purposes
	GetDefaultGroupLimitExecute(r GetDefaultGroupLimitApiRequest) (*DefaultGroupLimitResponse, *http.Response, error)

	/*
		ListDefaultGroupLimits Return Descriptions of User-Configurable Project Limits

		Returns a list of all user-configurable, project-level limits, along with a description and their default and maximum values.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListDefaultGroupLimitsApiRequest
	*/
	ListDefaultGroupLimits(ctx context.Context) ListDefaultGroupLimitsApiRequest
	/*
		ListDefaultGroupLimits Return Descriptions of User-Configurable Project Limits


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDefaultGroupLimitsApiParams - Parameters for the request
		@return ListDefaultGroupLimitsApiRequest
	*/
	ListDefaultGroupLimitsWithParams(ctx context.Context, args *ListDefaultGroupLimitsApiParams) ListDefaultGroupLimitsApiRequest

	// Method available only for mocking purposes
	ListDefaultGroupLimitsExecute(r ListDefaultGroupLimitsApiRequest) (*PaginatedDefaultGroupLimitResponse, *http.Response, error)
}

// LimitDescriptionApiService LimitDescriptionApi service
type LimitDescriptionApiService service

type GetDefaultGroupLimitApiRequest struct {
	ctx        context.Context
	ApiService LimitDescriptionApi
	limitName  string
}

type GetDefaultGroupLimitApiParams struct {
	LimitName string
}

func (a *LimitDescriptionApiService) GetDefaultGroupLimitWithParams(ctx context.Context, args *GetDefaultGroupLimitApiParams) GetDefaultGroupLimitApiRequest {
	return GetDefaultGroupLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  args.LimitName,
	}
}

func (r GetDefaultGroupLimitApiRequest) Execute() (*DefaultGroupLimitResponse, *http.Response, error) {
	return r.ApiService.GetDefaultGroupLimitExecute(r)
}

/*
GetDefaultGroupLimit Return One User-Configurable Project Limit and Description

Returns the description of one user-configurable, project-level limit, along with its default and maximum values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Identifies the user-managed limit.
	@return GetDefaultGroupLimitApiRequest
*/
func (a *LimitDescriptionApiService) GetDefaultGroupLimit(ctx context.Context, limitName string) GetDefaultGroupLimitApiRequest {
	return GetDefaultGroupLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  limitName,
	}
}

// GetDefaultGroupLimitExecute executes the request
//
//	@return DefaultGroupLimitResponse
func (a *LimitDescriptionApiService) GetDefaultGroupLimitExecute(r GetDefaultGroupLimitApiRequest) (*DefaultGroupLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *DefaultGroupLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LimitDescriptionApiService.GetDefaultGroupLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/defaultGroupLimits/{limitName}"
	if r.limitName == "" {
		return localVarReturnValue, nil, reportError("limitName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(r.limitName), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

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

type ListDefaultGroupLimitsApiRequest struct {
	ctx          context.Context
	ApiService   LimitDescriptionApi
	itemsPerPage *int
	pageNum      *int
}

type ListDefaultGroupLimitsApiParams struct {
	ItemsPerPage *int
	PageNum      *int
}

func (a *LimitDescriptionApiService) ListDefaultGroupLimitsWithParams(ctx context.Context, args *ListDefaultGroupLimitsApiParams) ListDefaultGroupLimitsApiRequest {
	return ListDefaultGroupLimitsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListDefaultGroupLimitsApiRequest) ItemsPerPage(itemsPerPage int) ListDefaultGroupLimitsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListDefaultGroupLimitsApiRequest) PageNum(pageNum int) ListDefaultGroupLimitsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListDefaultGroupLimitsApiRequest) Execute() (*PaginatedDefaultGroupLimitResponse, *http.Response, error) {
	return r.ApiService.ListDefaultGroupLimitsExecute(r)
}

/*
ListDefaultGroupLimits Return Descriptions of User-Configurable Project Limits

Returns a list of all user-configurable, project-level limits, along with a description and their default and maximum values.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListDefaultGroupLimitsApiRequest
*/
func (a *LimitDescriptionApiService) ListDefaultGroupLimits(ctx context.Context) ListDefaultGroupLimitsApiRequest {
	return ListDefaultGroupLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListDefaultGroupLimitsExecute executes the request
//
//	@return PaginatedDefaultGroupLimitResponse
func (a *LimitDescriptionApiService) ListDefaultGroupLimitsExecute(r ListDefaultGroupLimitsApiRequest) (*PaginatedDefaultGroupLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedDefaultGroupLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LimitDescriptionApiService.ListDefaultGroupLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/defaultGroupLimits"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

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
