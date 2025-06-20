// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ThirdPartyIntegrationsApi interface {

	/*
		CreateThirdPartyIntegration Configure One Third-Party Service Integration

		Adds the settings for configuring one third-party service integration. These settings apply to all databases managed in the specified MongoDB Cloud project. Each project can have only one configuration per `{INTEGRATION-TYPE}`. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param thirdPartyIntegration Third-party integration that you want to configure for your project.
		@return CreateThirdPartyIntegrationApiRequest
	*/
	CreateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *ThirdPartyIntegration) CreateThirdPartyIntegrationApiRequest
	/*
		CreateThirdPartyIntegration Configure One Third-Party Service Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateThirdPartyIntegrationApiParams - Parameters for the request
		@return CreateThirdPartyIntegrationApiRequest
	*/
	CreateThirdPartyIntegrationWithParams(ctx context.Context, args *CreateThirdPartyIntegrationApiParams) CreateThirdPartyIntegrationApiRequest

	// Method available only for mocking purposes
	CreateThirdPartyIntegrationExecute(r CreateThirdPartyIntegrationApiRequest) (*PaginatedIntegration, *http.Response, error)

	/*
		DeleteThirdPartyIntegration Remove One Third-Party Service Integration

		Removes the settings that permit configuring one third-party service integration. These settings apply to all databases managed in one MongoDB Cloud project. If you delete an integration from a project, you remove that integration configuration only for that project. This action doesn't affect any other project or organization's configured `{INTEGRATION-TYPE}` integrations. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteThirdPartyIntegrationApiRequest
	*/
	DeleteThirdPartyIntegration(ctx context.Context, integrationType string, groupId string) DeleteThirdPartyIntegrationApiRequest
	/*
		DeleteThirdPartyIntegration Remove One Third-Party Service Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteThirdPartyIntegrationApiParams - Parameters for the request
		@return DeleteThirdPartyIntegrationApiRequest
	*/
	DeleteThirdPartyIntegrationWithParams(ctx context.Context, args *DeleteThirdPartyIntegrationApiParams) DeleteThirdPartyIntegrationApiRequest

	// Method available only for mocking purposes
	DeleteThirdPartyIntegrationExecute(r DeleteThirdPartyIntegrationApiRequest) (*http.Response, error)

	/*
		GetThirdPartyIntegration Return One Third-Party Service Integration

		Returns the settings for configuring integration with one third-party service. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
		@return GetThirdPartyIntegrationApiRequest
	*/
	GetThirdPartyIntegration(ctx context.Context, groupId string, integrationType string) GetThirdPartyIntegrationApiRequest
	/*
		GetThirdPartyIntegration Return One Third-Party Service Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetThirdPartyIntegrationApiParams - Parameters for the request
		@return GetThirdPartyIntegrationApiRequest
	*/
	GetThirdPartyIntegrationWithParams(ctx context.Context, args *GetThirdPartyIntegrationApiParams) GetThirdPartyIntegrationApiRequest

	// Method available only for mocking purposes
	GetThirdPartyIntegrationExecute(r GetThirdPartyIntegrationApiRequest) (*ThirdPartyIntegration, *http.Response, error)

	/*
		ListThirdPartyIntegrations Return All Active Third-Party Service Integrations

		Returns the settings that permit integrations with all configured third-party services. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListThirdPartyIntegrationsApiRequest
	*/
	ListThirdPartyIntegrations(ctx context.Context, groupId string) ListThirdPartyIntegrationsApiRequest
	/*
		ListThirdPartyIntegrations Return All Active Third-Party Service Integrations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListThirdPartyIntegrationsApiParams - Parameters for the request
		@return ListThirdPartyIntegrationsApiRequest
	*/
	ListThirdPartyIntegrationsWithParams(ctx context.Context, args *ListThirdPartyIntegrationsApiParams) ListThirdPartyIntegrationsApiRequest

	// Method available only for mocking purposes
	ListThirdPartyIntegrationsExecute(r ListThirdPartyIntegrationsApiRequest) (*PaginatedIntegration, *http.Response, error)

	/*
		UpdateThirdPartyIntegration Update One Third-Party Service Integration

		Updates the settings for configuring integration with one third-party service. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param thirdPartyIntegration Third-party integration that you want to configure for your project.
		@return UpdateThirdPartyIntegrationApiRequest
	*/
	UpdateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *ThirdPartyIntegration) UpdateThirdPartyIntegrationApiRequest
	/*
		UpdateThirdPartyIntegration Update One Third-Party Service Integration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateThirdPartyIntegrationApiParams - Parameters for the request
		@return UpdateThirdPartyIntegrationApiRequest
	*/
	UpdateThirdPartyIntegrationWithParams(ctx context.Context, args *UpdateThirdPartyIntegrationApiParams) UpdateThirdPartyIntegrationApiRequest

	// Method available only for mocking purposes
	UpdateThirdPartyIntegrationExecute(r UpdateThirdPartyIntegrationApiRequest) (*PaginatedIntegration, *http.Response, error)
}

// ThirdPartyIntegrationsApiService ThirdPartyIntegrationsApi service
type ThirdPartyIntegrationsApiService service

type CreateThirdPartyIntegrationApiRequest struct {
	ctx                   context.Context
	ApiService            ThirdPartyIntegrationsApi
	integrationType       string
	groupId               string
	thirdPartyIntegration *ThirdPartyIntegration
	includeCount          *bool
	itemsPerPage          *int
	pageNum               *int
}

type CreateThirdPartyIntegrationApiParams struct {
	IntegrationType       string
	GroupId               string
	ThirdPartyIntegration *ThirdPartyIntegration
	IncludeCount          *bool
	ItemsPerPage          *int
	PageNum               *int
}

func (a *ThirdPartyIntegrationsApiService) CreateThirdPartyIntegrationWithParams(ctx context.Context, args *CreateThirdPartyIntegrationApiParams) CreateThirdPartyIntegrationApiRequest {
	return CreateThirdPartyIntegrationApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		integrationType:       args.IntegrationType,
		groupId:               args.GroupId,
		thirdPartyIntegration: args.ThirdPartyIntegration,
		includeCount:          args.IncludeCount,
		itemsPerPage:          args.ItemsPerPage,
		pageNum:               args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r CreateThirdPartyIntegrationApiRequest) IncludeCount(includeCount bool) CreateThirdPartyIntegrationApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r CreateThirdPartyIntegrationApiRequest) ItemsPerPage(itemsPerPage int) CreateThirdPartyIntegrationApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r CreateThirdPartyIntegrationApiRequest) PageNum(pageNum int) CreateThirdPartyIntegrationApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r CreateThirdPartyIntegrationApiRequest) Execute() (*PaginatedIntegration, *http.Response, error) {
	return r.ApiService.CreateThirdPartyIntegrationExecute(r)
}

/*
CreateThirdPartyIntegration Configure One Third-Party Service Integration

Adds the settings for configuring one third-party service integration. These settings apply to all databases managed in the specified MongoDB Cloud project. Each project can have only one configuration per `{INTEGRATION-TYPE}`. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateThirdPartyIntegrationApiRequest
*/
func (a *ThirdPartyIntegrationsApiService) CreateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *ThirdPartyIntegration) CreateThirdPartyIntegrationApiRequest {
	return CreateThirdPartyIntegrationApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		integrationType:       integrationType,
		groupId:               groupId,
		thirdPartyIntegration: thirdPartyIntegration,
	}
}

// CreateThirdPartyIntegrationExecute executes the request
//
//	@return PaginatedIntegration
func (a *ThirdPartyIntegrationsApiService) CreateThirdPartyIntegrationExecute(r CreateThirdPartyIntegrationApiRequest) (*PaginatedIntegration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedIntegration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThirdPartyIntegrationsApiService.CreateThirdPartyIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/integrations/{integrationType}"
	if r.integrationType == "" {
		return localVarReturnValue, nil, reportError("integrationType is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"integrationType"+"}", url.PathEscape(r.integrationType), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.thirdPartyIntegration == nil {
		return localVarReturnValue, nil, reportError("thirdPartyIntegration is required and must be specified")
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
	localVarPostBody = r.thirdPartyIntegration
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

type DeleteThirdPartyIntegrationApiRequest struct {
	ctx             context.Context
	ApiService      ThirdPartyIntegrationsApi
	integrationType string
	groupId         string
}

type DeleteThirdPartyIntegrationApiParams struct {
	IntegrationType string
	GroupId         string
}

func (a *ThirdPartyIntegrationsApiService) DeleteThirdPartyIntegrationWithParams(ctx context.Context, args *DeleteThirdPartyIntegrationApiParams) DeleteThirdPartyIntegrationApiRequest {
	return DeleteThirdPartyIntegrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		integrationType: args.IntegrationType,
		groupId:         args.GroupId,
	}
}

func (r DeleteThirdPartyIntegrationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteThirdPartyIntegrationExecute(r)
}

/*
DeleteThirdPartyIntegration Remove One Third-Party Service Integration

Removes the settings that permit configuring one third-party service integration. These settings apply to all databases managed in one MongoDB Cloud project. If you delete an integration from a project, you remove that integration configuration only for that project. This action doesn't affect any other project or organization's configured `{INTEGRATION-TYPE}` integrations. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteThirdPartyIntegrationApiRequest
*/
func (a *ThirdPartyIntegrationsApiService) DeleteThirdPartyIntegration(ctx context.Context, integrationType string, groupId string) DeleteThirdPartyIntegrationApiRequest {
	return DeleteThirdPartyIntegrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		integrationType: integrationType,
		groupId:         groupId,
	}
}

// DeleteThirdPartyIntegrationExecute executes the request
func (a *ThirdPartyIntegrationsApiService) DeleteThirdPartyIntegrationExecute(r DeleteThirdPartyIntegrationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThirdPartyIntegrationsApiService.DeleteThirdPartyIntegration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/integrations/{integrationType}"
	if r.integrationType == "" {
		return nil, reportError("integrationType is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"integrationType"+"}", url.PathEscape(r.integrationType), -1)
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
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

type GetThirdPartyIntegrationApiRequest struct {
	ctx             context.Context
	ApiService      ThirdPartyIntegrationsApi
	groupId         string
	integrationType string
}

type GetThirdPartyIntegrationApiParams struct {
	GroupId         string
	IntegrationType string
}

func (a *ThirdPartyIntegrationsApiService) GetThirdPartyIntegrationWithParams(ctx context.Context, args *GetThirdPartyIntegrationApiParams) GetThirdPartyIntegrationApiRequest {
	return GetThirdPartyIntegrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		integrationType: args.IntegrationType,
	}
}

func (r GetThirdPartyIntegrationApiRequest) Execute() (*ThirdPartyIntegration, *http.Response, error) {
	return r.ApiService.GetThirdPartyIntegrationExecute(r)
}

/*
GetThirdPartyIntegration Return One Third-Party Service Integration

Returns the settings for configuring integration with one third-party service. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
	@return GetThirdPartyIntegrationApiRequest
*/
func (a *ThirdPartyIntegrationsApiService) GetThirdPartyIntegration(ctx context.Context, groupId string, integrationType string) GetThirdPartyIntegrationApiRequest {
	return GetThirdPartyIntegrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		integrationType: integrationType,
	}
}

// GetThirdPartyIntegrationExecute executes the request
//
//	@return ThirdPartyIntegration
func (a *ThirdPartyIntegrationsApiService) GetThirdPartyIntegrationExecute(r GetThirdPartyIntegrationApiRequest) (*ThirdPartyIntegration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ThirdPartyIntegration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThirdPartyIntegrationsApiService.GetThirdPartyIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/integrations/{integrationType}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.integrationType == "" {
		return localVarReturnValue, nil, reportError("integrationType is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"integrationType"+"}", url.PathEscape(r.integrationType), -1)

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

type ListThirdPartyIntegrationsApiRequest struct {
	ctx          context.Context
	ApiService   ThirdPartyIntegrationsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListThirdPartyIntegrationsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ThirdPartyIntegrationsApiService) ListThirdPartyIntegrationsWithParams(ctx context.Context, args *ListThirdPartyIntegrationsApiParams) ListThirdPartyIntegrationsApiRequest {
	return ListThirdPartyIntegrationsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListThirdPartyIntegrationsApiRequest) IncludeCount(includeCount bool) ListThirdPartyIntegrationsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListThirdPartyIntegrationsApiRequest) ItemsPerPage(itemsPerPage int) ListThirdPartyIntegrationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListThirdPartyIntegrationsApiRequest) PageNum(pageNum int) ListThirdPartyIntegrationsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListThirdPartyIntegrationsApiRequest) Execute() (*PaginatedIntegration, *http.Response, error) {
	return r.ApiService.ListThirdPartyIntegrationsExecute(r)
}

/*
ListThirdPartyIntegrations Return All Active Third-Party Service Integrations

Returns the settings that permit integrations with all configured third-party services. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListThirdPartyIntegrationsApiRequest
*/
func (a *ThirdPartyIntegrationsApiService) ListThirdPartyIntegrations(ctx context.Context, groupId string) ListThirdPartyIntegrationsApiRequest {
	return ListThirdPartyIntegrationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListThirdPartyIntegrationsExecute executes the request
//
//	@return PaginatedIntegration
func (a *ThirdPartyIntegrationsApiService) ListThirdPartyIntegrationsExecute(r ListThirdPartyIntegrationsApiRequest) (*PaginatedIntegration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedIntegration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThirdPartyIntegrationsApiService.ListThirdPartyIntegrations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/integrations"
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

type UpdateThirdPartyIntegrationApiRequest struct {
	ctx                   context.Context
	ApiService            ThirdPartyIntegrationsApi
	integrationType       string
	groupId               string
	thirdPartyIntegration *ThirdPartyIntegration
	includeCount          *bool
	itemsPerPage          *int
	pageNum               *int
}

type UpdateThirdPartyIntegrationApiParams struct {
	IntegrationType       string
	GroupId               string
	ThirdPartyIntegration *ThirdPartyIntegration
	IncludeCount          *bool
	ItemsPerPage          *int
	PageNum               *int
}

func (a *ThirdPartyIntegrationsApiService) UpdateThirdPartyIntegrationWithParams(ctx context.Context, args *UpdateThirdPartyIntegrationApiParams) UpdateThirdPartyIntegrationApiRequest {
	return UpdateThirdPartyIntegrationApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		integrationType:       args.IntegrationType,
		groupId:               args.GroupId,
		thirdPartyIntegration: args.ThirdPartyIntegration,
		includeCount:          args.IncludeCount,
		itemsPerPage:          args.ItemsPerPage,
		pageNum:               args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r UpdateThirdPartyIntegrationApiRequest) IncludeCount(includeCount bool) UpdateThirdPartyIntegrationApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r UpdateThirdPartyIntegrationApiRequest) ItemsPerPage(itemsPerPage int) UpdateThirdPartyIntegrationApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r UpdateThirdPartyIntegrationApiRequest) PageNum(pageNum int) UpdateThirdPartyIntegrationApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r UpdateThirdPartyIntegrationApiRequest) Execute() (*PaginatedIntegration, *http.Response, error) {
	return r.ApiService.UpdateThirdPartyIntegrationExecute(r)
}

/*
UpdateThirdPartyIntegration Update One Third-Party Service Integration

Updates the settings for configuring integration with one third-party service. These settings apply to all databases managed in one MongoDB Cloud project. To use this resource, the requesting Service Account or API Key must have the Organization Owner or Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationType Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateThirdPartyIntegrationApiRequest
*/
func (a *ThirdPartyIntegrationsApiService) UpdateThirdPartyIntegration(ctx context.Context, integrationType string, groupId string, thirdPartyIntegration *ThirdPartyIntegration) UpdateThirdPartyIntegrationApiRequest {
	return UpdateThirdPartyIntegrationApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		integrationType:       integrationType,
		groupId:               groupId,
		thirdPartyIntegration: thirdPartyIntegration,
	}
}

// UpdateThirdPartyIntegrationExecute executes the request
//
//	@return PaginatedIntegration
func (a *ThirdPartyIntegrationsApiService) UpdateThirdPartyIntegrationExecute(r UpdateThirdPartyIntegrationApiRequest) (*PaginatedIntegration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedIntegration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThirdPartyIntegrationsApiService.UpdateThirdPartyIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/integrations/{integrationType}"
	if r.integrationType == "" {
		return localVarReturnValue, nil, reportError("integrationType is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"integrationType"+"}", url.PathEscape(r.integrationType), -1)
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.thirdPartyIntegration == nil {
		return localVarReturnValue, nil, reportError("thirdPartyIntegration is required and must be specified")
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
	localVarPostBody = r.thirdPartyIntegration
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
