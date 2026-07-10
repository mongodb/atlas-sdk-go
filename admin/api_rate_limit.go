// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RateLimitAPI interface {

	/*
			GetGroupRatelimits Return Rate Limit State for One Group

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve rate limiting bucket state for the specified group.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return GetGroupRatelimitsApiRequest
	*/
	GetGroupRatelimits(ctx context.Context, groupId string) GetGroupRatelimitsApiRequest
	/*
		GetGroupRatelimits Return Rate Limit State for One Group


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupRatelimitsApiParams - Parameters for the request
		@return GetGroupRatelimitsApiRequest
	*/
	GetGroupRatelimitsWithParams(ctx context.Context, args *GetGroupRatelimitsApiParams) GetGroupRatelimitsApiRequest

	// Method available only for mocking purposes
	GetGroupRatelimitsExecute(r GetGroupRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error)

	/*
			GetOrgRatelimits Return Rate Limit State for One Organization

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve rate limiting bucket state for the specified organization.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return GetOrgRatelimitsApiRequest
	*/
	GetOrgRatelimits(ctx context.Context, orgId string) GetOrgRatelimitsApiRequest
	/*
		GetOrgRatelimits Return Rate Limit State for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrgRatelimitsApiParams - Parameters for the request
		@return GetOrgRatelimitsApiRequest
	*/
	GetOrgRatelimitsWithParams(ctx context.Context, args *GetOrgRatelimitsApiParams) GetOrgRatelimitsApiRequest

	// Method available only for mocking purposes
	GetOrgRatelimitsExecute(r GetOrgRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error)

	/*
			ListRatelimits Return Rate Limit State for Unauthenticated Requests

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve rate limiting bucket state for unauthenticated requests.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ListRatelimitsApiRequest
	*/
	ListRatelimits(ctx context.Context) ListRatelimitsApiRequest
	/*
		ListRatelimits Return Rate Limit State for Unauthenticated Requests


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListRatelimitsApiParams - Parameters for the request
		@return ListRatelimitsApiRequest
	*/
	ListRatelimitsWithParams(ctx context.Context, args *ListRatelimitsApiParams) ListRatelimitsApiRequest

	// Method available only for mocking purposes
	ListRatelimitsExecute(r ListRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error)

	/*
			ListUserRatelimits Return Rate Limit State for One User

			This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

		 Retrieve rate limiting bucket state for the current user.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ListUserRatelimitsApiRequest
	*/
	ListUserRatelimits(ctx context.Context) ListUserRatelimitsApiRequest
	/*
		ListUserRatelimits Return Rate Limit State for One User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListUserRatelimitsApiParams - Parameters for the request
		@return ListUserRatelimitsApiRequest
	*/
	ListUserRatelimitsWithParams(ctx context.Context, args *ListUserRatelimitsApiParams) ListUserRatelimitsApiRequest

	// Method available only for mocking purposes
	ListUserRatelimitsExecute(r ListUserRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error)
}

// RateLimitAPIService RateLimitAPI service
type RateLimitAPIService service

type GetGroupRatelimitsApiRequest struct {
	ctx        context.Context
	ApiService RateLimitAPI
	groupId    string
}

type GetGroupRatelimitsApiParams struct {
	GroupId string
}

func (a *RateLimitAPIService) GetGroupRatelimitsWithParams(ctx context.Context, args *GetGroupRatelimitsApiParams) GetGroupRatelimitsApiRequest {
	return GetGroupRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetGroupRatelimitsApiRequest) Execute() (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	return r.ApiService.GetGroupRatelimitsExecute(r)
}

/*
GetGroupRatelimits Return Rate Limit State for One Group

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve rate limiting bucket state for the specified group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetGroupRatelimitsApiRequest
*/
func (a *RateLimitAPIService) GetGroupRatelimits(ctx context.Context, groupId string) GetGroupRatelimitsApiRequest {
	return GetGroupRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetGroupRatelimitsExecute executes the request
//
//	@return AtlasRateLimitInspectionResponse
func (a *RateLimitAPIService) GetGroupRatelimitsExecute(r GetGroupRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasRateLimitInspectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitAPIService.GetGroupRatelimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/ratelimits"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

type GetOrgRatelimitsApiRequest struct {
	ctx        context.Context
	ApiService RateLimitAPI
	orgId      string
}

type GetOrgRatelimitsApiParams struct {
	OrgId string
}

func (a *RateLimitAPIService) GetOrgRatelimitsWithParams(ctx context.Context, args *GetOrgRatelimitsApiParams) GetOrgRatelimitsApiRequest {
	return GetOrgRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetOrgRatelimitsApiRequest) Execute() (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	return r.ApiService.GetOrgRatelimitsExecute(r)
}

/*
GetOrgRatelimits Return Rate Limit State for One Organization

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve rate limiting bucket state for the specified organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [`/orgs`](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetOrgRatelimitsApiRequest
*/
func (a *RateLimitAPIService) GetOrgRatelimits(ctx context.Context, orgId string) GetOrgRatelimitsApiRequest {
	return GetOrgRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetOrgRatelimitsExecute executes the request
//
//	@return AtlasRateLimitInspectionResponse
func (a *RateLimitAPIService) GetOrgRatelimitsExecute(r GetOrgRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasRateLimitInspectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitAPIService.GetOrgRatelimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/ratelimits"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

type ListRatelimitsApiRequest struct {
	ctx        context.Context
	ApiService RateLimitAPI
}

type ListRatelimitsApiParams struct {
}

func (a *RateLimitAPIService) ListRatelimitsWithParams(ctx context.Context, args *ListRatelimitsApiParams) ListRatelimitsApiRequest {
	return ListRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ListRatelimitsApiRequest) Execute() (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	return r.ApiService.ListRatelimitsExecute(r)
}

/*
ListRatelimits Return Rate Limit State for Unauthenticated Requests

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve rate limiting bucket state for unauthenticated requests.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListRatelimitsApiRequest
*/
func (a *RateLimitAPIService) ListRatelimits(ctx context.Context) ListRatelimitsApiRequest {
	return ListRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListRatelimitsExecute executes the request
//
//	@return AtlasRateLimitInspectionResponse
func (a *RateLimitAPIService) ListRatelimitsExecute(r ListRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasRateLimitInspectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitAPIService.ListRatelimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/unauth/ratelimits"

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

type ListUserRatelimitsApiRequest struct {
	ctx        context.Context
	ApiService RateLimitAPI
}

type ListUserRatelimitsApiParams struct {
}

func (a *RateLimitAPIService) ListUserRatelimitsWithParams(ctx context.Context, args *ListUserRatelimitsApiParams) ListUserRatelimitsApiRequest {
	return ListUserRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ListUserRatelimitsApiRequest) Execute() (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	return r.ApiService.ListUserRatelimitsExecute(r)
}

/*
ListUserRatelimits Return Rate Limit State for One User

This API is in preview. Breaking changes might be introduced before it is released. Don't use preview APIs in production.

	Retrieve rate limiting bucket state for the current user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListUserRatelimitsApiRequest
*/
func (a *RateLimitAPIService) ListUserRatelimits(ctx context.Context) ListUserRatelimitsApiRequest {
	return ListUserRatelimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListUserRatelimitsExecute executes the request
//
//	@return AtlasRateLimitInspectionResponse
func (a *RateLimitAPIService) ListUserRatelimitsExecute(r ListUserRatelimitsApiRequest) (*AtlasRateLimitInspectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AtlasRateLimitInspectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitAPIService.ListUserRatelimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/userRateLimits"

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
