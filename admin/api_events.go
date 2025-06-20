// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type EventsApi interface {

	/*
			GetOrganizationEvent Return One Event from One Organization

			Returns one event for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Organization Member role. Use the Return Events from One Organization endpoint to retrieve all events to which the authenticated user has access.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
			@return GetOrganizationEventApiRequest
	*/
	GetOrganizationEvent(ctx context.Context, orgId string, eventId string) GetOrganizationEventApiRequest
	/*
		GetOrganizationEvent Return One Event from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOrganizationEventApiParams - Parameters for the request
		@return GetOrganizationEventApiRequest
	*/
	GetOrganizationEventWithParams(ctx context.Context, args *GetOrganizationEventApiParams) GetOrganizationEventApiRequest

	// Method available only for mocking purposes
	GetOrganizationEventExecute(r GetOrganizationEventApiRequest) (*EventViewForOrg, *http.Response, error)

	/*
			GetProjectEvent Return One Event from One Project

			Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Use the Return Events from One Project endpoint to retrieve all events to which the authenticated user has access.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
			@return GetProjectEventApiRequest
	*/
	GetProjectEvent(ctx context.Context, groupId string, eventId string) GetProjectEventApiRequest
	/*
		GetProjectEvent Return One Event from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectEventApiParams - Parameters for the request
		@return GetProjectEventApiRequest
	*/
	GetProjectEventWithParams(ctx context.Context, args *GetProjectEventApiParams) GetProjectEventApiRequest

	// Method available only for mocking purposes
	GetProjectEventExecute(r GetProjectEventApiRequest) (*EventViewForNdsGroup, *http.Response, error)

	/*
		ListEventTypes Return All Event Types

		Returns a list of all event types, along with a description and additional metadata about each event.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListEventTypesApiRequest
	*/
	ListEventTypes(ctx context.Context) ListEventTypesApiRequest
	/*
		ListEventTypes Return All Event Types


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListEventTypesApiParams - Parameters for the request
		@return ListEventTypesApiRequest
	*/
	ListEventTypesWithParams(ctx context.Context, args *ListEventTypesApiParams) ListEventTypesApiRequest

	// Method available only for mocking purposes
	ListEventTypesExecute(r ListEventTypesApiRequest) (*PaginatedEventTypeDetailsResponse, *http.Response, error)

	/*
			ListOrganizationEvents Return Events from One Organization

			Returns events for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListOrganizationEventsApiRequest
	*/
	ListOrganizationEvents(ctx context.Context, orgId string) ListOrganizationEventsApiRequest
	/*
		ListOrganizationEvents Return Events from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOrganizationEventsApiParams - Parameters for the request
		@return ListOrganizationEventsApiRequest
	*/
	ListOrganizationEventsWithParams(ctx context.Context, args *ListOrganizationEventsApiParams) ListOrganizationEventsApiRequest

	// Method available only for mocking purposes
	ListOrganizationEventsExecute(r ListOrganizationEventsApiRequest) (*OrgPaginatedEvent, *http.Response, error)

	/*
			ListProjectEvents Return Events from One Project

			Returns events for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListProjectEventsApiRequest
	*/
	ListProjectEvents(ctx context.Context, groupId string) ListProjectEventsApiRequest
	/*
		ListProjectEvents Return Events from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectEventsApiParams - Parameters for the request
		@return ListProjectEventsApiRequest
	*/
	ListProjectEventsWithParams(ctx context.Context, args *ListProjectEventsApiParams) ListProjectEventsApiRequest

	// Method available only for mocking purposes
	ListProjectEventsExecute(r ListProjectEventsApiRequest) (*GroupPaginatedEvent, *http.Response, error)
}

// EventsApiService EventsApi service
type EventsApiService service

type GetOrganizationEventApiRequest struct {
	ctx        context.Context
	ApiService EventsApi
	orgId      string
	eventId    string
	includeRaw *bool
}

type GetOrganizationEventApiParams struct {
	OrgId      string
	EventId    string
	IncludeRaw *bool
}

func (a *EventsApiService) GetOrganizationEventWithParams(ctx context.Context, args *GetOrganizationEventApiParams) GetOrganizationEventApiRequest {
	return GetOrganizationEventApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		eventId:    args.EventId,
		includeRaw: args.IncludeRaw,
	}
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r GetOrganizationEventApiRequest) IncludeRaw(includeRaw bool) GetOrganizationEventApiRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r GetOrganizationEventApiRequest) Execute() (*EventViewForOrg, *http.Response, error) {
	return r.ApiService.GetOrganizationEventExecute(r)
}

/*
GetOrganizationEvent Return One Event from One Organization

Returns one event for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Organization Member role. Use the Return Events from One Organization endpoint to retrieve all events to which the authenticated user has access.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
	@return GetOrganizationEventApiRequest
*/
func (a *EventsApiService) GetOrganizationEvent(ctx context.Context, orgId string, eventId string) GetOrganizationEventApiRequest {
	return GetOrganizationEventApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		eventId:    eventId,
	}
}

// GetOrganizationEventExecute executes the request
//
//	@return EventViewForOrg
func (a *EventsApiService) GetOrganizationEventExecute(r GetOrganizationEventApiRequest) (*EventViewForOrg, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EventViewForOrg
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetOrganizationEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/events/{eventId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.eventId == "" {
		return localVarReturnValue, nil, reportError("eventId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(r.eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
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

type GetProjectEventApiRequest struct {
	ctx        context.Context
	ApiService EventsApi
	groupId    string
	eventId    string
	includeRaw *bool
}

type GetProjectEventApiParams struct {
	GroupId    string
	EventId    string
	IncludeRaw *bool
}

func (a *EventsApiService) GetProjectEventWithParams(ctx context.Context, args *GetProjectEventApiParams) GetProjectEventApiRequest {
	return GetProjectEventApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		eventId:    args.EventId,
		includeRaw: args.IncludeRaw,
	}
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r GetProjectEventApiRequest) IncludeRaw(includeRaw bool) GetProjectEventApiRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r GetProjectEventApiRequest) Execute() (*EventViewForNdsGroup, *http.Response, error) {
	return r.ApiService.GetProjectEventExecute(r)
}

/*
GetProjectEvent Return One Event from One Project

Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Use the Return Events from One Project endpoint to retrieve all events to which the authenticated user has access.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
	@return GetProjectEventApiRequest
*/
func (a *EventsApiService) GetProjectEvent(ctx context.Context, groupId string, eventId string) GetProjectEventApiRequest {
	return GetProjectEventApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		eventId:    eventId,
	}
}

// GetProjectEventExecute executes the request
//
//	@return EventViewForNdsGroup
func (a *EventsApiService) GetProjectEventExecute(r GetProjectEventApiRequest) (*EventViewForNdsGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EventViewForNdsGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetProjectEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/events/{eventId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.eventId == "" {
		return localVarReturnValue, nil, reportError("eventId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(r.eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
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

type ListEventTypesApiRequest struct {
	ctx          context.Context
	ApiService   EventsApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListEventTypesApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *EventsApiService) ListEventTypesWithParams(ctx context.Context, args *ListEventTypesApiParams) ListEventTypesApiRequest {
	return ListEventTypesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListEventTypesApiRequest) IncludeCount(includeCount bool) ListEventTypesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListEventTypesApiRequest) ItemsPerPage(itemsPerPage int) ListEventTypesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListEventTypesApiRequest) PageNum(pageNum int) ListEventTypesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListEventTypesApiRequest) Execute() (*PaginatedEventTypeDetailsResponse, *http.Response, error) {
	return r.ApiService.ListEventTypesExecute(r)
}

/*
ListEventTypes Return All Event Types

Returns a list of all event types, along with a description and additional metadata about each event.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListEventTypesApiRequest
*/
func (a *EventsApiService) ListEventTypes(ctx context.Context) ListEventTypesApiRequest {
	return ListEventTypesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListEventTypesExecute executes the request
//
//	@return PaginatedEventTypeDetailsResponse
func (a *EventsApiService) ListEventTypesExecute(r ListEventTypesApiRequest) (*PaginatedEventTypeDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedEventTypeDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListEventTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/eventTypes"

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

type ListOrganizationEventsApiRequest struct {
	ctx          context.Context
	ApiService   EventsApi
	orgId        string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	eventType    *[]string
	includeRaw   *bool
	maxDate      *time.Time
	minDate      *time.Time
}

type ListOrganizationEventsApiParams struct {
	OrgId        string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	EventType    *[]string
	IncludeRaw   *bool
	MaxDate      *time.Time
	MinDate      *time.Time
}

func (a *EventsApiService) ListOrganizationEventsWithParams(ctx context.Context, args *ListOrganizationEventsApiParams) ListOrganizationEventsApiRequest {
	return ListOrganizationEventsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		eventType:    args.EventType,
		includeRaw:   args.IncludeRaw,
		maxDate:      args.MaxDate,
		minDate:      args.MinDate,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOrganizationEventsApiRequest) IncludeCount(includeCount bool) ListOrganizationEventsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListOrganizationEventsApiRequest) ItemsPerPage(itemsPerPage int) ListOrganizationEventsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOrganizationEventsApiRequest) PageNum(pageNum int) ListOrganizationEventsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.
func (r ListOrganizationEventsApiRequest) EventType(eventType []string) ListOrganizationEventsApiRequest {
	r.eventType = &eventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ListOrganizationEventsApiRequest) IncludeRaw(includeRaw bool) ListOrganizationEventsApiRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ListOrganizationEventsApiRequest) MaxDate(maxDate time.Time) ListOrganizationEventsApiRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ListOrganizationEventsApiRequest) MinDate(minDate time.Time) ListOrganizationEventsApiRequest {
	r.minDate = &minDate
	return r
}

func (r ListOrganizationEventsApiRequest) Execute() (*OrgPaginatedEvent, *http.Response, error) {
	return r.ApiService.ListOrganizationEventsExecute(r)
}

/*
ListOrganizationEvents Return Events from One Organization

Returns events for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Organization Member role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListOrganizationEventsApiRequest
*/
func (a *EventsApiService) ListOrganizationEvents(ctx context.Context, orgId string) ListOrganizationEventsApiRequest {
	return ListOrganizationEventsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListOrganizationEventsExecute executes the request
//
//	@return OrgPaginatedEvent
func (a *EventsApiService) ListOrganizationEventsExecute(r ListOrganizationEventsApiRequest) (*OrgPaginatedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *OrgPaginatedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListOrganizationEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/events"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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
	if r.eventType != nil {
		t := *r.eventType
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "multi")

	}
	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	}
	if r.maxDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxDate", r.maxDate, "")
	}
	if r.minDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minDate", r.minDate, "")
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

type ListProjectEventsApiRequest struct {
	ctx               context.Context
	ApiService        EventsApi
	groupId           string
	includeCount      *bool
	itemsPerPage      *int
	pageNum           *int
	clusterNames      *[]string
	eventType         *[]string
	excludedEventType *[]string
	includeRaw        *bool
	maxDate           *time.Time
	minDate           *time.Time
}

type ListProjectEventsApiParams struct {
	GroupId           string
	IncludeCount      *bool
	ItemsPerPage      *int
	PageNum           *int
	ClusterNames      *[]string
	EventType         *[]string
	ExcludedEventType *[]string
	IncludeRaw        *bool
	MaxDate           *time.Time
	MinDate           *time.Time
}

func (a *EventsApiService) ListProjectEventsWithParams(ctx context.Context, args *ListProjectEventsApiParams) ListProjectEventsApiRequest {
	return ListProjectEventsApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		includeCount:      args.IncludeCount,
		itemsPerPage:      args.ItemsPerPage,
		pageNum:           args.PageNum,
		clusterNames:      args.ClusterNames,
		eventType:         args.EventType,
		excludedEventType: args.ExcludedEventType,
		includeRaw:        args.IncludeRaw,
		maxDate:           args.MaxDate,
		minDate:           args.MinDate,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectEventsApiRequest) IncludeCount(includeCount bool) ListProjectEventsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectEventsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectEventsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectEventsApiRequest) PageNum(pageNum int) ListProjectEventsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Human-readable label that identifies the cluster.
func (r ListProjectEventsApiRequest) ClusterNames(clusterNames []string) ListProjectEventsApiRequest {
	r.clusterNames = &clusterNames
	return r
}

// Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.
func (r ListProjectEventsApiRequest) EventType(eventType []string) ListProjectEventsApiRequest {
	r.eventType = &eventType
	return r
}

// Category of event that you would like to exclude from query results, such as CLUSTER_CREATED  **IMPORTANT**: Event type names change frequently. Verify that you specify the event type correctly by checking the complete list of event types.
func (r ListProjectEventsApiRequest) ExcludedEventType(excludedEventType []string) ListProjectEventsApiRequest {
	r.excludedEventType = &excludedEventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ListProjectEventsApiRequest) IncludeRaw(includeRaw bool) ListProjectEventsApiRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ListProjectEventsApiRequest) MaxDate(maxDate time.Time) ListProjectEventsApiRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ListProjectEventsApiRequest) MinDate(minDate time.Time) ListProjectEventsApiRequest {
	r.minDate = &minDate
	return r
}

func (r ListProjectEventsApiRequest) Execute() (*GroupPaginatedEvent, *http.Response, error) {
	return r.ApiService.ListProjectEventsExecute(r)
}

/*
ListProjectEvents Return Events from One Project

Returns events for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectEventsApiRequest
*/
func (a *EventsApiService) ListProjectEvents(ctx context.Context, groupId string) ListProjectEventsApiRequest {
	return ListProjectEventsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListProjectEventsExecute executes the request
//
//	@return GroupPaginatedEvent
func (a *EventsApiService) ListProjectEventsExecute(r ListProjectEventsApiRequest) (*GroupPaginatedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GroupPaginatedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListProjectEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/events"
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
	if r.clusterNames != nil {
		t := *r.clusterNames
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "clusterNames", t, "multi")

	}
	if r.eventType != nil {
		t := *r.eventType
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "multi")

	}
	if r.excludedEventType != nil {
		t := *r.excludedEventType
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludedEventType", t, "multi")

	}
	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	}
	if r.maxDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxDate", r.maxDate, "")
	}
	if r.minDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minDate", r.minDate, "")
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
