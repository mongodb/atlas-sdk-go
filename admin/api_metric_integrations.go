// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MetricIntegrationsApi interface {

	/*
		CreateGroupMetricIntegration Create One Metric Integration

		Creates a new metric integration configuration identified by a unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param metricIntegrationRequest Metric integration configuration to create.
		@return CreateGroupMetricIntegrationApiRequest
	*/
	CreateGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationRequest *MetricIntegrationRequest) CreateGroupMetricIntegrationApiRequest
	/*
		CreateGroupMetricIntegration Create One Metric Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupMetricIntegrationApiParams - Parameters for the request
		@return CreateGroupMetricIntegrationApiRequest
	*/
	CreateGroupMetricIntegrationWithParams(ctx context.Context, args *CreateGroupMetricIntegrationApiParams) CreateGroupMetricIntegrationApiRequest

	// Method available only for mocking purposes
	CreateGroupMetricIntegrationExecute(r CreateGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error)

	/*
		DeleteGroupMetricIntegration Remove One Metric Integration

		Removes the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param metricIntegrationId Unique identifier of the metric integration configuration.
		@return DeleteGroupMetricIntegrationApiRequest
	*/
	DeleteGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string) DeleteGroupMetricIntegrationApiRequest
	/*
		DeleteGroupMetricIntegration Remove One Metric Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupMetricIntegrationApiParams - Parameters for the request
		@return DeleteGroupMetricIntegrationApiRequest
	*/
	DeleteGroupMetricIntegrationWithParams(ctx context.Context, args *DeleteGroupMetricIntegrationApiParams) DeleteGroupMetricIntegrationApiRequest

	// Method available only for mocking purposes
	DeleteGroupMetricIntegrationExecute(r DeleteGroupMetricIntegrationApiRequest) (*http.Response, error)

	/*
		GetGroupMetricIntegration Return One Metric Integration

		Returns the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param metricIntegrationId Unique identifier of the metric integration configuration.
		@return GetGroupMetricIntegrationApiRequest
	*/
	GetGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string) GetGroupMetricIntegrationApiRequest
	/*
		GetGroupMetricIntegration Return One Metric Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupMetricIntegrationApiParams - Parameters for the request
		@return GetGroupMetricIntegrationApiRequest
	*/
	GetGroupMetricIntegrationWithParams(ctx context.Context, args *GetGroupMetricIntegrationApiParams) GetGroupMetricIntegrationApiRequest

	// Method available only for mocking purposes
	GetGroupMetricIntegrationExecute(r GetGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error)

	/*
		ListGroupMetricIntegrations Return All Active Metric Integrations

		Returns all metric integration configurations for the project. Optionally filter by integration type and provider type. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListGroupMetricIntegrationsApiRequest
	*/
	ListGroupMetricIntegrations(ctx context.Context, groupId string) ListGroupMetricIntegrationsApiRequest
	/*
		ListGroupMetricIntegrations Return All Active Metric Integrations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupMetricIntegrationsApiParams - Parameters for the request
		@return ListGroupMetricIntegrationsApiRequest
	*/
	ListGroupMetricIntegrationsWithParams(ctx context.Context, args *ListGroupMetricIntegrationsApiParams) ListGroupMetricIntegrationsApiRequest

	// Method available only for mocking purposes
	ListGroupMetricIntegrationsExecute(r ListGroupMetricIntegrationsApiRequest) (*PaginatedMetricIntegrationResponse, *http.Response, error)

	/*
		UpdateGroupMetricIntegration Update One Metric Integration

		Updates the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param metricIntegrationId Unique identifier of the metric integration configuration.
		@param metricIntegrationUpdateRequest Updated metric integration configuration.
		@return UpdateGroupMetricIntegrationApiRequest
	*/
	UpdateGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string, metricIntegrationUpdateRequest *MetricIntegrationUpdateRequest) UpdateGroupMetricIntegrationApiRequest
	/*
		UpdateGroupMetricIntegration Update One Metric Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateGroupMetricIntegrationApiParams - Parameters for the request
		@return UpdateGroupMetricIntegrationApiRequest
	*/
	UpdateGroupMetricIntegrationWithParams(ctx context.Context, args *UpdateGroupMetricIntegrationApiParams) UpdateGroupMetricIntegrationApiRequest

	// Method available only for mocking purposes
	UpdateGroupMetricIntegrationExecute(r UpdateGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error)
}

// MetricIntegrationsApiService MetricIntegrationsApi service
type MetricIntegrationsApiService service

type CreateGroupMetricIntegrationApiRequest struct {
	ctx                      context.Context
	ApiService               MetricIntegrationsApi
	groupId                  string
	metricIntegrationRequest *MetricIntegrationRequest
}

type CreateGroupMetricIntegrationApiParams struct {
	GroupId                  string
	MetricIntegrationRequest *MetricIntegrationRequest
}

func (a *MetricIntegrationsApiService) CreateGroupMetricIntegrationWithParams(ctx context.Context, args *CreateGroupMetricIntegrationApiParams) CreateGroupMetricIntegrationApiRequest {
	return CreateGroupMetricIntegrationApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		metricIntegrationRequest: args.MetricIntegrationRequest,
	}
}

func (r CreateGroupMetricIntegrationApiRequest) Execute() (*MetricIntegrationResponse, *http.Response, error) {
	return r.ApiService.CreateGroupMetricIntegrationExecute(r)
}

/*
CreateGroupMetricIntegration Create One Metric Integration

Creates a new metric integration configuration identified by a unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateGroupMetricIntegrationApiRequest
*/
func (a *MetricIntegrationsApiService) CreateGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationRequest *MetricIntegrationRequest) CreateGroupMetricIntegrationApiRequest {
	return CreateGroupMetricIntegrationApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		metricIntegrationRequest: metricIntegrationRequest,
	}
}

// CreateGroupMetricIntegrationExecute executes the request
//
//	@return MetricIntegrationResponse
func (a *MetricIntegrationsApiService) CreateGroupMetricIntegrationExecute(r CreateGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MetricIntegrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricIntegrationsApiService.CreateGroupMetricIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/metricIntegrations"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metricIntegrationRequest == nil {
		return localVarReturnValue, nil, reportError("metricIntegrationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

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
	// body params
	localVarPostBody = r.metricIntegrationRequest
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

type DeleteGroupMetricIntegrationApiRequest struct {
	ctx                 context.Context
	ApiService          MetricIntegrationsApi
	groupId             string
	metricIntegrationId string
}

type DeleteGroupMetricIntegrationApiParams struct {
	GroupId             string
	MetricIntegrationId string
}

func (a *MetricIntegrationsApiService) DeleteGroupMetricIntegrationWithParams(ctx context.Context, args *DeleteGroupMetricIntegrationApiParams) DeleteGroupMetricIntegrationApiRequest {
	return DeleteGroupMetricIntegrationApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		metricIntegrationId: args.MetricIntegrationId,
	}
}

func (r DeleteGroupMetricIntegrationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupMetricIntegrationExecute(r)
}

/*
DeleteGroupMetricIntegration Remove One Metric Integration

Removes the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param metricIntegrationId Unique identifier of the metric integration configuration.
	@return DeleteGroupMetricIntegrationApiRequest
*/
func (a *MetricIntegrationsApiService) DeleteGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string) DeleteGroupMetricIntegrationApiRequest {
	return DeleteGroupMetricIntegrationApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             groupId,
		metricIntegrationId: metricIntegrationId,
	}
}

// DeleteGroupMetricIntegrationExecute executes the request
func (a *MetricIntegrationsApiService) DeleteGroupMetricIntegrationExecute(r DeleteGroupMetricIntegrationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricIntegrationsApiService.DeleteGroupMetricIntegration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.metricIntegrationId == "" {
		return nil, reportError("metricIntegrationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"metricIntegrationId"+"}", url.PathEscape(r.metricIntegrationId), -1)

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

type GetGroupMetricIntegrationApiRequest struct {
	ctx                 context.Context
	ApiService          MetricIntegrationsApi
	groupId             string
	metricIntegrationId string
}

type GetGroupMetricIntegrationApiParams struct {
	GroupId             string
	MetricIntegrationId string
}

func (a *MetricIntegrationsApiService) GetGroupMetricIntegrationWithParams(ctx context.Context, args *GetGroupMetricIntegrationApiParams) GetGroupMetricIntegrationApiRequest {
	return GetGroupMetricIntegrationApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		metricIntegrationId: args.MetricIntegrationId,
	}
}

func (r GetGroupMetricIntegrationApiRequest) Execute() (*MetricIntegrationResponse, *http.Response, error) {
	return r.ApiService.GetGroupMetricIntegrationExecute(r)
}

/*
GetGroupMetricIntegration Return One Metric Integration

Returns the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param metricIntegrationId Unique identifier of the metric integration configuration.
	@return GetGroupMetricIntegrationApiRequest
*/
func (a *MetricIntegrationsApiService) GetGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string) GetGroupMetricIntegrationApiRequest {
	return GetGroupMetricIntegrationApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             groupId,
		metricIntegrationId: metricIntegrationId,
	}
}

// GetGroupMetricIntegrationExecute executes the request
//
//	@return MetricIntegrationResponse
func (a *MetricIntegrationsApiService) GetGroupMetricIntegrationExecute(r GetGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MetricIntegrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricIntegrationsApiService.GetGroupMetricIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.metricIntegrationId == "" {
		return localVarReturnValue, nil, reportError("metricIntegrationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"metricIntegrationId"+"}", url.PathEscape(r.metricIntegrationId), -1)

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

type ListGroupMetricIntegrationsApiRequest struct {
	ctx             context.Context
	ApiService      MetricIntegrationsApi
	groupId         string
	includeCount    *bool
	itemsPerPage    *int
	pageNum         *int
	integrationType *string
	providerType    *string
}

type ListGroupMetricIntegrationsApiParams struct {
	GroupId         string
	IncludeCount    *bool
	ItemsPerPage    *int
	PageNum         *int
	IntegrationType *string
	ProviderType    *string
}

func (a *MetricIntegrationsApiService) ListGroupMetricIntegrationsWithParams(ctx context.Context, args *ListGroupMetricIntegrationsApiParams) ListGroupMetricIntegrationsApiRequest {
	return ListGroupMetricIntegrationsApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		includeCount:    args.IncludeCount,
		itemsPerPage:    args.ItemsPerPage,
		pageNum:         args.PageNum,
		integrationType: args.IntegrationType,
		providerType:    args.ProviderType,
	}
}

// Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response.
func (r ListGroupMetricIntegrationsApiRequest) IncludeCount(includeCount bool) ListGroupMetricIntegrationsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListGroupMetricIntegrationsApiRequest) ItemsPerPage(itemsPerPage int) ListGroupMetricIntegrationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupMetricIntegrationsApiRequest) PageNum(pageNum int) ListGroupMetricIntegrationsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Optional filter by integration type (e.g., &#x60;OTEL&#x60;). When specified, &#x60;providerType&#x60; must also be specified.
func (r ListGroupMetricIntegrationsApiRequest) IntegrationType(integrationType string) ListGroupMetricIntegrationsApiRequest {
	r.integrationType = &integrationType
	return r
}

// Optional filter by provider type (e.g., &#x60;CUSTOM&#x60;). When specified, &#x60;integrationType&#x60; must also be specified.
func (r ListGroupMetricIntegrationsApiRequest) ProviderType(providerType string) ListGroupMetricIntegrationsApiRequest {
	r.providerType = &providerType
	return r
}

func (r ListGroupMetricIntegrationsApiRequest) Execute() (*PaginatedMetricIntegrationResponse, *http.Response, error) {
	return r.ApiService.ListGroupMetricIntegrationsExecute(r)
}

/*
ListGroupMetricIntegrations Return All Active Metric Integrations

Returns all metric integration configurations for the project. Optionally filter by integration type and provider type. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListGroupMetricIntegrationsApiRequest
*/
func (a *MetricIntegrationsApiService) ListGroupMetricIntegrations(ctx context.Context, groupId string) ListGroupMetricIntegrationsApiRequest {
	return ListGroupMetricIntegrationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupMetricIntegrationsExecute executes the request
//
//	@return PaginatedMetricIntegrationResponse
func (a *MetricIntegrationsApiService) ListGroupMetricIntegrationsExecute(r ListGroupMetricIntegrationsApiRequest) (*PaginatedMetricIntegrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedMetricIntegrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricIntegrationsApiService.ListGroupMetricIntegrations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/metricIntegrations"
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
	if r.integrationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "integrationType", r.integrationType, "")
	}
	if r.providerType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "providerType", r.providerType, "")
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

type UpdateGroupMetricIntegrationApiRequest struct {
	ctx                            context.Context
	ApiService                     MetricIntegrationsApi
	groupId                        string
	metricIntegrationId            string
	metricIntegrationUpdateRequest *MetricIntegrationUpdateRequest
}

type UpdateGroupMetricIntegrationApiParams struct {
	GroupId                        string
	MetricIntegrationId            string
	MetricIntegrationUpdateRequest *MetricIntegrationUpdateRequest
}

func (a *MetricIntegrationsApiService) UpdateGroupMetricIntegrationWithParams(ctx context.Context, args *UpdateGroupMetricIntegrationApiParams) UpdateGroupMetricIntegrationApiRequest {
	return UpdateGroupMetricIntegrationApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		metricIntegrationId:            args.MetricIntegrationId,
		metricIntegrationUpdateRequest: args.MetricIntegrationUpdateRequest,
	}
}

func (r UpdateGroupMetricIntegrationApiRequest) Execute() (*MetricIntegrationResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupMetricIntegrationExecute(r)
}

/*
UpdateGroupMetricIntegration Update One Metric Integration

Updates the configuration for one metric integration identified by its unique ID. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param metricIntegrationId Unique identifier of the metric integration configuration.
	@return UpdateGroupMetricIntegrationApiRequest
*/
func (a *MetricIntegrationsApiService) UpdateGroupMetricIntegration(ctx context.Context, groupId string, metricIntegrationId string, metricIntegrationUpdateRequest *MetricIntegrationUpdateRequest) UpdateGroupMetricIntegrationApiRequest {
	return UpdateGroupMetricIntegrationApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		metricIntegrationId:            metricIntegrationId,
		metricIntegrationUpdateRequest: metricIntegrationUpdateRequest,
	}
}

// UpdateGroupMetricIntegrationExecute executes the request
//
//	@return MetricIntegrationResponse
func (a *MetricIntegrationsApiService) UpdateGroupMetricIntegrationExecute(r UpdateGroupMetricIntegrationApiRequest) (*MetricIntegrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *MetricIntegrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricIntegrationsApiService.UpdateGroupMetricIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/metricIntegrations/{metricIntegrationId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.metricIntegrationId == "" {
		return localVarReturnValue, nil, reportError("metricIntegrationId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"metricIntegrationId"+"}", url.PathEscape(r.metricIntegrationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metricIntegrationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("metricIntegrationUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

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
	// body params
	localVarPostBody = r.metricIntegrationUpdateRequest
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
