// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type USSInstancesApi interface {

	/*
		CreateUSSInstance Create One USS Instance in One Project

		Creates one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param uSSInstanceDescriptionCreate20250101 Create One USS Instance in One Project.
		@return CreateUSSInstanceApiRequest
	*/
	CreateUSSInstance(ctx context.Context, groupId string, uSSInstanceDescriptionCreate20250101 *USSInstanceDescriptionCreate20250101) CreateUSSInstanceApiRequest
	/*
		CreateUSSInstance Create One USS Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateUSSInstanceApiParams - Parameters for the request
		@return CreateUSSInstanceApiRequest
	*/
	CreateUSSInstanceWithParams(ctx context.Context, args *CreateUSSInstanceApiParams) CreateUSSInstanceApiRequest

	// Method available only for mocking purposes
	CreateUSSInstanceExecute(r CreateUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error)

	/*
		DeleteUSSInstance Remove One USS Instance from One Project

		Removes one USS instance from the specified project. The USS instance must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the USS instance.
		@return DeleteUSSInstanceApiRequest
	*/
	DeleteUSSInstance(ctx context.Context, groupId string, name string) DeleteUSSInstanceApiRequest
	/*
		DeleteUSSInstance Remove One USS Instance from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteUSSInstanceApiParams - Parameters for the request
		@return DeleteUSSInstanceApiRequest
	*/
	DeleteUSSInstanceWithParams(ctx context.Context, args *DeleteUSSInstanceApiParams) DeleteUSSInstanceApiRequest

	// Method available only for mocking purposes
	DeleteUSSInstanceExecute(r DeleteUSSInstanceApiRequest) (interface{}, *http.Response, error)

	/*
		GetUSSInstance Return One USS Instance from One Project

		Returns details for one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the USS instance.
		@return GetUSSInstanceApiRequest
	*/
	GetUSSInstance(ctx context.Context, groupId string, name string) GetUSSInstanceApiRequest
	/*
		GetUSSInstance Return One USS Instance from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetUSSInstanceApiParams - Parameters for the request
		@return GetUSSInstanceApiRequest
	*/
	GetUSSInstanceWithParams(ctx context.Context, args *GetUSSInstanceApiParams) GetUSSInstanceApiRequest

	// Method available only for mocking purposes
	GetUSSInstanceExecute(r GetUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error)

	/*
		ListUSSInstances Return All USS Instances from One Project

		Returns details for all USS instances in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListUSSInstancesApiRequest
	*/
	ListUSSInstances(ctx context.Context, groupId string) ListUSSInstancesApiRequest
	/*
		ListUSSInstances Return All USS Instances from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListUSSInstancesApiParams - Parameters for the request
		@return ListUSSInstancesApiRequest
	*/
	ListUSSInstancesWithParams(ctx context.Context, args *ListUSSInstancesApiParams) ListUSSInstancesApiRequest

	// Method available only for mocking purposes
	ListUSSInstancesExecute(r ListUSSInstancesApiRequest) (*PaginatedUSSInstance20250101, *http.Response, error)

	/*
		UpdateUSSInstance Update One USS Instance in One Project

		Updates one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the USS instance.
		@param uSSInstanceDescription20250101 Update One USS Instance in One Project.
		@return UpdateUSSInstanceApiRequest
	*/
	UpdateUSSInstance(ctx context.Context, groupId string, name string, uSSInstanceDescription20250101 *USSInstanceDescription20250101) UpdateUSSInstanceApiRequest
	/*
		UpdateUSSInstance Update One USS Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateUSSInstanceApiParams - Parameters for the request
		@return UpdateUSSInstanceApiRequest
	*/
	UpdateUSSInstanceWithParams(ctx context.Context, args *UpdateUSSInstanceApiParams) UpdateUSSInstanceApiRequest

	// Method available only for mocking purposes
	UpdateUSSInstanceExecute(r UpdateUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error)

	/*
		UpgradeUSSInstance Upgrade One USS Instance

		Upgrades a USS instance in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param uSSInstanceDescription20250101 Details of the USS instance upgrade in the specified project.
		@return UpgradeUSSInstanceApiRequest
	*/
	UpgradeUSSInstance(ctx context.Context, groupId string, uSSInstanceDescription20250101 *USSInstanceDescription20250101) UpgradeUSSInstanceApiRequest
	/*
		UpgradeUSSInstance Upgrade One USS Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeUSSInstanceApiParams - Parameters for the request
		@return UpgradeUSSInstanceApiRequest
	*/
	UpgradeUSSInstanceWithParams(ctx context.Context, args *UpgradeUSSInstanceApiParams) UpgradeUSSInstanceApiRequest

	// Method available only for mocking purposes
	UpgradeUSSInstanceExecute(r UpgradeUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error)
}

// USSInstancesApiService USSInstancesApi service
type USSInstancesApiService service

type CreateUSSInstanceApiRequest struct {
	ctx                                  context.Context
	ApiService                           USSInstancesApi
	groupId                              string
	uSSInstanceDescriptionCreate20250101 *USSInstanceDescriptionCreate20250101
}

type CreateUSSInstanceApiParams struct {
	GroupId                              string
	USSInstanceDescriptionCreate20250101 *USSInstanceDescriptionCreate20250101
}

func (a *USSInstancesApiService) CreateUSSInstanceWithParams(ctx context.Context, args *CreateUSSInstanceApiParams) CreateUSSInstanceApiRequest {
	return CreateUSSInstanceApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		uSSInstanceDescriptionCreate20250101: args.USSInstanceDescriptionCreate20250101,
	}
}

func (r CreateUSSInstanceApiRequest) Execute() (*USSInstanceDescription20250101, *http.Response, error) {
	return r.ApiService.CreateUSSInstanceExecute(r)
}

/*
CreateUSSInstance Create One USS Instance in One Project

Creates one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateUSSInstanceApiRequest
*/
func (a *USSInstancesApiService) CreateUSSInstance(ctx context.Context, groupId string, uSSInstanceDescriptionCreate20250101 *USSInstanceDescriptionCreate20250101) CreateUSSInstanceApiRequest {
	return CreateUSSInstanceApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		uSSInstanceDescriptionCreate20250101: uSSInstanceDescriptionCreate20250101,
	}
}

// Execute executes the request
//
//	@return USSInstanceDescription20250101
func (a *USSInstancesApiService) CreateUSSInstanceExecute(r CreateUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *USSInstanceDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.CreateUSSInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uSSInstanceDescriptionCreate20250101 == nil {
		return localVarReturnValue, nil, reportError("uSSInstanceDescriptionCreate20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uSSInstanceDescriptionCreate20250101
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

type DeleteUSSInstanceApiRequest struct {
	ctx        context.Context
	ApiService USSInstancesApi
	groupId    string
	name       string
}

type DeleteUSSInstanceApiParams struct {
	GroupId string
	Name    string
}

func (a *USSInstancesApiService) DeleteUSSInstanceWithParams(ctx context.Context, args *DeleteUSSInstanceApiParams) DeleteUSSInstanceApiRequest {
	return DeleteUSSInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r DeleteUSSInstanceApiRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteUSSInstanceExecute(r)
}

/*
DeleteUSSInstance Remove One USS Instance from One Project

Removes one USS instance from the specified project. The USS instance must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the USS instance.
	@return DeleteUSSInstanceApiRequest
*/
func (a *USSInstancesApiService) DeleteUSSInstance(ctx context.Context, groupId string, name string) DeleteUSSInstanceApiRequest {
	return DeleteUSSInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *USSInstancesApiService) DeleteUSSInstanceExecute(r DeleteUSSInstanceApiRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.DeleteUSSInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type GetUSSInstanceApiRequest struct {
	ctx        context.Context
	ApiService USSInstancesApi
	groupId    string
	name       string
}

type GetUSSInstanceApiParams struct {
	GroupId string
	Name    string
}

func (a *USSInstancesApiService) GetUSSInstanceWithParams(ctx context.Context, args *GetUSSInstanceApiParams) GetUSSInstanceApiRequest {
	return GetUSSInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r GetUSSInstanceApiRequest) Execute() (*USSInstanceDescription20250101, *http.Response, error) {
	return r.ApiService.GetUSSInstanceExecute(r)
}

/*
GetUSSInstance Return One USS Instance from One Project

Returns details for one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the USS instance.
	@return GetUSSInstanceApiRequest
*/
func (a *USSInstancesApiService) GetUSSInstance(ctx context.Context, groupId string, name string) GetUSSInstanceApiRequest {
	return GetUSSInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// Execute executes the request
//
//	@return USSInstanceDescription20250101
func (a *USSInstancesApiService) GetUSSInstanceExecute(r GetUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *USSInstanceDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.GetUSSInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type ListUSSInstancesApiRequest struct {
	ctx          context.Context
	ApiService   USSInstancesApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListUSSInstancesApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *USSInstancesApiService) ListUSSInstancesWithParams(ctx context.Context, args *ListUSSInstancesApiParams) ListUSSInstancesApiRequest {
	return ListUSSInstancesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListUSSInstancesApiRequest) IncludeCount(includeCount bool) ListUSSInstancesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListUSSInstancesApiRequest) ItemsPerPage(itemsPerPage int) ListUSSInstancesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListUSSInstancesApiRequest) PageNum(pageNum int) ListUSSInstancesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListUSSInstancesApiRequest) Execute() (*PaginatedUSSInstance20250101, *http.Response, error) {
	return r.ApiService.ListUSSInstancesExecute(r)
}

/*
ListUSSInstances Return All USS Instances from One Project

Returns details for all USS instances in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListUSSInstancesApiRequest
*/
func (a *USSInstancesApiService) ListUSSInstances(ctx context.Context, groupId string) ListUSSInstancesApiRequest {
	return ListUSSInstancesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedUSSInstance20250101
func (a *USSInstancesApiService) ListUSSInstancesExecute(r ListUSSInstancesApiRequest) (*PaginatedUSSInstance20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedUSSInstance20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.ListUSSInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

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

type UpdateUSSInstanceApiRequest struct {
	ctx                            context.Context
	ApiService                     USSInstancesApi
	groupId                        string
	name                           string
	uSSInstanceDescription20250101 *USSInstanceDescription20250101
}

type UpdateUSSInstanceApiParams struct {
	GroupId                        string
	Name                           string
	USSInstanceDescription20250101 *USSInstanceDescription20250101
}

func (a *USSInstancesApiService) UpdateUSSInstanceWithParams(ctx context.Context, args *UpdateUSSInstanceApiParams) UpdateUSSInstanceApiRequest {
	return UpdateUSSInstanceApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		name:                           args.Name,
		uSSInstanceDescription20250101: args.USSInstanceDescription20250101,
	}
}

func (r UpdateUSSInstanceApiRequest) Execute() (*USSInstanceDescription20250101, *http.Response, error) {
	return r.ApiService.UpdateUSSInstanceExecute(r)
}

/*
UpdateUSSInstance Update One USS Instance in One Project

Updates one USS instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the USS instance.
	@return UpdateUSSInstanceApiRequest
*/
func (a *USSInstancesApiService) UpdateUSSInstance(ctx context.Context, groupId string, name string, uSSInstanceDescription20250101 *USSInstanceDescription20250101) UpdateUSSInstanceApiRequest {
	return UpdateUSSInstanceApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		name:                           name,
		uSSInstanceDescription20250101: uSSInstanceDescription20250101,
	}
}

// Execute executes the request
//
//	@return USSInstanceDescription20250101
func (a *USSInstancesApiService) UpdateUSSInstanceExecute(r UpdateUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *USSInstanceDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.UpdateUSSInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uSSInstanceDescription20250101 == nil {
		return localVarReturnValue, nil, reportError("uSSInstanceDescription20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uSSInstanceDescription20250101
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

type UpgradeUSSInstanceApiRequest struct {
	ctx                            context.Context
	ApiService                     USSInstancesApi
	groupId                        string
	uSSInstanceDescription20250101 *USSInstanceDescription20250101
}

type UpgradeUSSInstanceApiParams struct {
	GroupId                        string
	USSInstanceDescription20250101 *USSInstanceDescription20250101
}

func (a *USSInstancesApiService) UpgradeUSSInstanceWithParams(ctx context.Context, args *UpgradeUSSInstanceApiParams) UpgradeUSSInstanceApiRequest {
	return UpgradeUSSInstanceApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		uSSInstanceDescription20250101: args.USSInstanceDescription20250101,
	}
}

func (r UpgradeUSSInstanceApiRequest) Execute() (*USSInstanceDescription20250101, *http.Response, error) {
	return r.ApiService.UpgradeUSSInstanceExecute(r)
}

/*
UpgradeUSSInstance Upgrade One USS Instance

Upgrades a USS instance in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeUSSInstanceApiRequest
*/
func (a *USSInstancesApiService) UpgradeUSSInstance(ctx context.Context, groupId string, uSSInstanceDescription20250101 *USSInstanceDescription20250101) UpgradeUSSInstanceApiRequest {
	return UpgradeUSSInstanceApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		uSSInstanceDescription20250101: uSSInstanceDescription20250101,
	}
}

// Execute executes the request
//
//	@return USSInstanceDescription20250101
func (a *USSInstancesApiService) UpgradeUSSInstanceExecute(r UpgradeUSSInstanceApiRequest) (*USSInstanceDescription20250101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *USSInstanceDescription20250101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "USSInstancesApiService.UpgradeUSSInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/uss:tenantUpgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uSSInstanceDescription20250101 == nil {
		return localVarReturnValue, nil, reportError("uSSInstanceDescription20250101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uSSInstanceDescription20250101
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
