// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ProjectIPAccessListApi interface {

	/*
		CreateProjectIpAccessList Add Entries to Project IP Access List

		Adds one or more access list entries to the specified project. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. Write each entry as either one IP address or one CIDR-notated block of IP addresses. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations. This endpoint doesn't support concurrent `POST` requests. You must submit multiple `POST` requests synchronously.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param networkPermissionEntry One or more access list entries to add to the specified project.
		@return CreateProjectIpAccessListApiRequest
	*/
	CreateProjectIpAccessList(ctx context.Context, groupId string, networkPermissionEntry *[]NetworkPermissionEntry) CreateProjectIpAccessListApiRequest
	/*
		CreateProjectIpAccessList Add Entries to Project IP Access List


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectIpAccessListApiParams - Parameters for the request
		@return CreateProjectIpAccessListApiRequest
	*/
	CreateProjectIpAccessListWithParams(ctx context.Context, args *CreateProjectIpAccessListApiParams) CreateProjectIpAccessListApiRequest

	// Method available only for mocking purposes
	CreateProjectIpAccessListExecute(r CreateProjectIpAccessListApiRequest) (*PaginatedNetworkAccess, *http.Response, error)

	/*
		DeleteProjectIpAccessList Remove One Entry from One Project IP Access List

		Removes one access list entry from the specified project's IP access list. Each entry in the project's IP access list contains one IP address, one CIDR-notated block of IP addresses, or one AWS Security Group ID. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param entryValue Access list entry that you want to remove from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:  - how your application established the connection, - how MongoDB Cloud or the driver using the address behaves, and - which protocol (like TCP or UDP) the connection uses.
		@return DeleteProjectIpAccessListApiRequest
	*/
	DeleteProjectIpAccessList(ctx context.Context, groupId string, entryValue string) DeleteProjectIpAccessListApiRequest
	/*
		DeleteProjectIpAccessList Remove One Entry from One Project IP Access List


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectIpAccessListApiParams - Parameters for the request
		@return DeleteProjectIpAccessListApiRequest
	*/
	DeleteProjectIpAccessListWithParams(ctx context.Context, args *DeleteProjectIpAccessListApiParams) DeleteProjectIpAccessListApiRequest

	// Method available only for mocking purposes
	DeleteProjectIpAccessListExecute(r DeleteProjectIpAccessListApiRequest) (*http.Response, error)

	/*
		GetProjectIpAccessListStatus Return Status of One Project IP Access List Entry

		Returns the status of one project IP access list entry. This resource checks if the provided project IP access list entry applies to all cloud providers serving clusters from the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param entryValue Network address or cloud provider security construct that identifies which project access list entry to be verified.
		@return GetProjectIpAccessListStatusApiRequest
	*/
	GetProjectIpAccessListStatus(ctx context.Context, groupId string, entryValue string) GetProjectIpAccessListStatusApiRequest
	/*
		GetProjectIpAccessListStatus Return Status of One Project IP Access List Entry


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectIpAccessListStatusApiParams - Parameters for the request
		@return GetProjectIpAccessListStatusApiRequest
	*/
	GetProjectIpAccessListStatusWithParams(ctx context.Context, args *GetProjectIpAccessListStatusApiParams) GetProjectIpAccessListStatusApiRequest

	// Method available only for mocking purposes
	GetProjectIpAccessListStatusExecute(r GetProjectIpAccessListStatusApiRequest) (*NetworkPermissionEntryStatus, *http.Response, error)

	/*
		GetProjectIpList Return One Project IP Access List Entry

		Returns one access list entry from the specified project's IP access list. Each entry in the project's IP access list contains either one IP address or one CIDR-notated block of IP addresses. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. This endpoint (`/groups/{GROUP-ID}/accessList`) manages the Project IP Access List. It doesn't manage the access list for MongoDB Cloud organizations. TheProgrammatic API Keys endpoint (`/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist`) manages those access lists.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param entryValue Access list entry that you want to return from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`).
		@return GetProjectIpListApiRequest
	*/
	GetProjectIpList(ctx context.Context, groupId string, entryValue string) GetProjectIpListApiRequest
	/*
		GetProjectIpList Return One Project IP Access List Entry


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectIpListApiParams - Parameters for the request
		@return GetProjectIpListApiRequest
	*/
	GetProjectIpListWithParams(ctx context.Context, args *GetProjectIpListApiParams) GetProjectIpListApiRequest

	// Method available only for mocking purposes
	GetProjectIpListExecute(r GetProjectIpListApiRequest) (*NetworkPermissionEntry, *http.Response, error)

	/*
		ListProjectIpAccessLists Return All Project IP Access List Entries

		Returns all access list entries from the specified project's IP access list. Each entry in the project's IP access list contains either one IP address or one CIDR-notated block of IP addresses. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectIpAccessListsApiRequest
	*/
	ListProjectIpAccessLists(ctx context.Context, groupId string) ListProjectIpAccessListsApiRequest
	/*
		ListProjectIpAccessLists Return All Project IP Access List Entries


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectIpAccessListsApiParams - Parameters for the request
		@return ListProjectIpAccessListsApiRequest
	*/
	ListProjectIpAccessListsWithParams(ctx context.Context, args *ListProjectIpAccessListsApiParams) ListProjectIpAccessListsApiRequest

	// Method available only for mocking purposes
	ListProjectIpAccessListsExecute(r ListProjectIpAccessListsApiRequest) (*PaginatedNetworkAccess, *http.Response, error)
}

// ProjectIPAccessListApiService ProjectIPAccessListApi service
type ProjectIPAccessListApiService service

type CreateProjectIpAccessListApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectIPAccessListApi
	groupId                string
	networkPermissionEntry *[]NetworkPermissionEntry
	includeCount           *bool
	itemsPerPage           *int
	pageNum                *int
}

type CreateProjectIpAccessListApiParams struct {
	GroupId                string
	NetworkPermissionEntry *[]NetworkPermissionEntry
	IncludeCount           *bool
	ItemsPerPage           *int
	PageNum                *int
}

func (a *ProjectIPAccessListApiService) CreateProjectIpAccessListWithParams(ctx context.Context, args *CreateProjectIpAccessListApiParams) CreateProjectIpAccessListApiRequest {
	return CreateProjectIpAccessListApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		networkPermissionEntry: args.NetworkPermissionEntry,
		includeCount:           args.IncludeCount,
		itemsPerPage:           args.ItemsPerPage,
		pageNum:                args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r CreateProjectIpAccessListApiRequest) IncludeCount(includeCount bool) CreateProjectIpAccessListApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r CreateProjectIpAccessListApiRequest) ItemsPerPage(itemsPerPage int) CreateProjectIpAccessListApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r CreateProjectIpAccessListApiRequest) PageNum(pageNum int) CreateProjectIpAccessListApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r CreateProjectIpAccessListApiRequest) Execute() (*PaginatedNetworkAccess, *http.Response, error) {
	return r.ApiService.CreateProjectIpAccessListExecute(r)
}

/*
CreateProjectIpAccessList Add Entries to Project IP Access List

Adds one or more access list entries to the specified project. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. Write each entry as either one IP address or one CIDR-notated block of IP addresses. To use this resource, the requesting Service Account or API Key must have the Project Owner or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations. This endpoint doesn't support concurrent `POST` requests. You must submit multiple `POST` requests synchronously.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectIpAccessListApiRequest
*/
func (a *ProjectIPAccessListApiService) CreateProjectIpAccessList(ctx context.Context, groupId string, networkPermissionEntry *[]NetworkPermissionEntry) CreateProjectIpAccessListApiRequest {
	return CreateProjectIpAccessListApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		networkPermissionEntry: networkPermissionEntry,
	}
}

// CreateProjectIpAccessListExecute executes the request
//
//	@return PaginatedNetworkAccess
func (a *ProjectIPAccessListApiService) CreateProjectIpAccessListExecute(r CreateProjectIpAccessListApiRequest) (*PaginatedNetworkAccess, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedNetworkAccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectIPAccessListApiService.CreateProjectIpAccessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/accessList"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.networkPermissionEntry == nil {
		return localVarReturnValue, nil, reportError("networkPermissionEntry is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.networkPermissionEntry
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

type DeleteProjectIpAccessListApiRequest struct {
	ctx        context.Context
	ApiService ProjectIPAccessListApi
	groupId    string
	entryValue string
}

type DeleteProjectIpAccessListApiParams struct {
	GroupId    string
	EntryValue string
}

func (a *ProjectIPAccessListApiService) DeleteProjectIpAccessListWithParams(ctx context.Context, args *DeleteProjectIpAccessListApiParams) DeleteProjectIpAccessListApiRequest {
	return DeleteProjectIpAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		entryValue: args.EntryValue,
	}
}

func (r DeleteProjectIpAccessListApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectIpAccessListExecute(r)
}

/*
DeleteProjectIpAccessList Remove One Entry from One Project IP Access List

Removes one access list entry from the specified project's IP access list. Each entry in the project's IP access list contains one IP address, one CIDR-notated block of IP addresses, or one AWS Security Group ID. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param entryValue Access list entry that you want to remove from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:  - how your application established the connection, - how MongoDB Cloud or the driver using the address behaves, and - which protocol (like TCP or UDP) the connection uses.
	@return DeleteProjectIpAccessListApiRequest
*/
func (a *ProjectIPAccessListApiService) DeleteProjectIpAccessList(ctx context.Context, groupId string, entryValue string) DeleteProjectIpAccessListApiRequest {
	return DeleteProjectIpAccessListApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		entryValue: entryValue,
	}
}

// DeleteProjectIpAccessListExecute executes the request
func (a *ProjectIPAccessListApiService) DeleteProjectIpAccessListExecute(r DeleteProjectIpAccessListApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectIPAccessListApiService.DeleteProjectIpAccessList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/accessList/{entryValue}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.entryValue == "" {
		return nil, reportError("entryValue is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"entryValue"+"}", url.PathEscape(r.entryValue), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetProjectIpAccessListStatusApiRequest struct {
	ctx        context.Context
	ApiService ProjectIPAccessListApi
	groupId    string
	entryValue string
}

type GetProjectIpAccessListStatusApiParams struct {
	GroupId    string
	EntryValue string
}

func (a *ProjectIPAccessListApiService) GetProjectIpAccessListStatusWithParams(ctx context.Context, args *GetProjectIpAccessListStatusApiParams) GetProjectIpAccessListStatusApiRequest {
	return GetProjectIpAccessListStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		entryValue: args.EntryValue,
	}
}

func (r GetProjectIpAccessListStatusApiRequest) Execute() (*NetworkPermissionEntryStatus, *http.Response, error) {
	return r.ApiService.GetProjectIpAccessListStatusExecute(r)
}

/*
GetProjectIpAccessListStatus Return Status of One Project IP Access List Entry

Returns the status of one project IP access list entry. This resource checks if the provided project IP access list entry applies to all cloud providers serving clusters from the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param entryValue Network address or cloud provider security construct that identifies which project access list entry to be verified.
	@return GetProjectIpAccessListStatusApiRequest
*/
func (a *ProjectIPAccessListApiService) GetProjectIpAccessListStatus(ctx context.Context, groupId string, entryValue string) GetProjectIpAccessListStatusApiRequest {
	return GetProjectIpAccessListStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		entryValue: entryValue,
	}
}

// GetProjectIpAccessListStatusExecute executes the request
//
//	@return NetworkPermissionEntryStatus
func (a *ProjectIPAccessListApiService) GetProjectIpAccessListStatusExecute(r GetProjectIpAccessListStatusApiRequest) (*NetworkPermissionEntryStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkPermissionEntryStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectIPAccessListApiService.GetProjectIpAccessListStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.entryValue == "" {
		return localVarReturnValue, nil, reportError("entryValue is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"entryValue"+"}", url.PathEscape(r.entryValue), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type GetProjectIpListApiRequest struct {
	ctx        context.Context
	ApiService ProjectIPAccessListApi
	groupId    string
	entryValue string
}

type GetProjectIpListApiParams struct {
	GroupId    string
	EntryValue string
}

func (a *ProjectIPAccessListApiService) GetProjectIpListWithParams(ctx context.Context, args *GetProjectIpListApiParams) GetProjectIpListApiRequest {
	return GetProjectIpListApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		entryValue: args.EntryValue,
	}
}

func (r GetProjectIpListApiRequest) Execute() (*NetworkPermissionEntry, *http.Response, error) {
	return r.ApiService.GetProjectIpListExecute(r)
}

/*
GetProjectIpList Return One Project IP Access List Entry

Returns one access list entry from the specified project's IP access list. Each entry in the project's IP access list contains either one IP address or one CIDR-notated block of IP addresses. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. This endpoint (`/groups/{GROUP-ID}/accessList`) manages the Project IP Access List. It doesn't manage the access list for MongoDB Cloud organizations. TheProgrammatic API Keys endpoint (`/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist`) manages those access lists.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param entryValue Access list entry that you want to return from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`).
	@return GetProjectIpListApiRequest
*/
func (a *ProjectIPAccessListApiService) GetProjectIpList(ctx context.Context, groupId string, entryValue string) GetProjectIpListApiRequest {
	return GetProjectIpListApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		entryValue: entryValue,
	}
}

// GetProjectIpListExecute executes the request
//
//	@return NetworkPermissionEntry
func (a *ProjectIPAccessListApiService) GetProjectIpListExecute(r GetProjectIpListApiRequest) (*NetworkPermissionEntry, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *NetworkPermissionEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectIPAccessListApiService.GetProjectIpList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/accessList/{entryValue}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.entryValue == "" {
		return localVarReturnValue, nil, reportError("entryValue is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"entryValue"+"}", url.PathEscape(r.entryValue), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListProjectIpAccessListsApiRequest struct {
	ctx          context.Context
	ApiService   ProjectIPAccessListApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectIpAccessListsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProjectIPAccessListApiService) ListProjectIpAccessListsWithParams(ctx context.Context, args *ListProjectIpAccessListsApiParams) ListProjectIpAccessListsApiRequest {
	return ListProjectIpAccessListsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectIpAccessListsApiRequest) IncludeCount(includeCount bool) ListProjectIpAccessListsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectIpAccessListsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectIpAccessListsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectIpAccessListsApiRequest) PageNum(pageNum int) ListProjectIpAccessListsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectIpAccessListsApiRequest) Execute() (*PaginatedNetworkAccess, *http.Response, error) {
	return r.ApiService.ListProjectIpAccessListsExecute(r)
}

/*
ListProjectIpAccessLists Return All Project IP Access List Entries

Returns all access list entries from the specified project's IP access list. Each entry in the project's IP access list contains either one IP address or one CIDR-notated block of IP addresses. MongoDB Cloud only allows client connections to the cluster from entries in the project's IP access list. To use this resource, the requesting Service Account or API Key must have the Project Read Only or Project Charts Admin roles. This resource replaces the whitelist resource. MongoDB Cloud removed whitelists in July 2021. Update your applications to use this new resource. The `/groups/{GROUP-ID}/accessList` endpoint manages the database IP access list. This endpoint is distinct from the `orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist` endpoint, which manages the access list for MongoDB Cloud organizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectIpAccessListsApiRequest
*/
func (a *ProjectIPAccessListApiService) ListProjectIpAccessLists(ctx context.Context, groupId string) ListProjectIpAccessListsApiRequest {
	return ListProjectIpAccessListsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectIpAccessListsExecute executes the request
//
//	@return PaginatedNetworkAccess
func (a *ProjectIPAccessListApiService) ListProjectIpAccessListsExecute(r ListProjectIpAccessListsApiRequest) (*PaginatedNetworkAccess, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedNetworkAccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectIPAccessListApiService.ListProjectIpAccessLists")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/accessList"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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
