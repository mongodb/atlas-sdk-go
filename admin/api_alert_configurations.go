// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AlertConfigurationsApi interface {

	/*
			CreateAlertConfiguration Create One Alert Configuration in One Project

			Creates one alert configuration for the specified project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return CreateAlertConfigurationApiRequest
	*/
	CreateAlertConfiguration(ctx context.Context, groupId string, groupAlertsConfig *GroupAlertsConfig) CreateAlertConfigurationApiRequest
	/*
		CreateAlertConfiguration Create One Alert Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAlertConfigurationApiParams - Parameters for the request
		@return CreateAlertConfigurationApiRequest
	*/
	CreateAlertConfigurationWithParams(ctx context.Context, args *CreateAlertConfigurationApiParams) CreateAlertConfigurationApiRequest

	// Method available only for mocking purposes
	CreateAlertConfigurationExecute(r CreateAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error)

	/*
			DeleteAlertConfiguration Remove One Alert Configuration from One Project

			Removes one alert configuration from the specified project. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
			@return DeleteAlertConfigurationApiRequest
	*/
	DeleteAlertConfiguration(ctx context.Context, groupId string, alertConfigId string) DeleteAlertConfigurationApiRequest
	/*
		DeleteAlertConfiguration Remove One Alert Configuration from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAlertConfigurationApiParams - Parameters for the request
		@return DeleteAlertConfigurationApiRequest
	*/
	DeleteAlertConfigurationWithParams(ctx context.Context, args *DeleteAlertConfigurationApiParams) DeleteAlertConfigurationApiRequest

	// Method available only for mocking purposes
	DeleteAlertConfigurationExecute(r DeleteAlertConfigurationApiRequest) (*http.Response, error)

	/*
			GetAlertConfiguration Return One Alert Configuration from One Project

			Returns the specified alert configuration from the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
			@return GetAlertConfigurationApiRequest
	*/
	GetAlertConfiguration(ctx context.Context, groupId string, alertConfigId string) GetAlertConfigurationApiRequest
	/*
		GetAlertConfiguration Return One Alert Configuration from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAlertConfigurationApiParams - Parameters for the request
		@return GetAlertConfigurationApiRequest
	*/
	GetAlertConfigurationWithParams(ctx context.Context, args *GetAlertConfigurationApiParams) GetAlertConfigurationApiRequest

	// Method available only for mocking purposes
	GetAlertConfigurationExecute(r GetAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error)

	/*
		ListAlertConfigurationMatchersFieldNames Get All Alert Configuration Matchers Field Names

		Get all field names that the `matchers.fieldName` parameter accepts when you create or update an Alert Configuration. You can successfully call this endpoint with any assigned role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListAlertConfigurationMatchersFieldNamesApiRequest
	*/
	ListAlertConfigurationMatchersFieldNames(ctx context.Context) ListAlertConfigurationMatchersFieldNamesApiRequest
	/*
		ListAlertConfigurationMatchersFieldNames Get All Alert Configuration Matchers Field Names


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAlertConfigurationMatchersFieldNamesApiParams - Parameters for the request
		@return ListAlertConfigurationMatchersFieldNamesApiRequest
	*/
	ListAlertConfigurationMatchersFieldNamesWithParams(ctx context.Context, args *ListAlertConfigurationMatchersFieldNamesApiParams) ListAlertConfigurationMatchersFieldNamesApiRequest

	// Method available only for mocking purposes
	ListAlertConfigurationMatchersFieldNamesExecute(r ListAlertConfigurationMatchersFieldNamesApiRequest) ([]string, *http.Response, error)

	/*
			ListAlertConfigurations Return All Alert Configurations for One Project

			Returns all alert configurations for one project. These alert configurations apply to any component in the project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Project Read Only role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListAlertConfigurationsApiRequest
	*/
	ListAlertConfigurations(ctx context.Context, groupId string) ListAlertConfigurationsApiRequest
	/*
		ListAlertConfigurations Return All Alert Configurations for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAlertConfigurationsApiParams - Parameters for the request
		@return ListAlertConfigurationsApiRequest
	*/
	ListAlertConfigurationsWithParams(ctx context.Context, args *ListAlertConfigurationsApiParams) ListAlertConfigurationsApiRequest

	// Method available only for mocking purposes
	ListAlertConfigurationsExecute(r ListAlertConfigurationsApiRequest) (*PaginatedAlertConfig, *http.Response, error)

	/*
			ListAlertConfigurationsByAlertId Return All Alert Configurations Set for One Alert

			Returns all alert configurations set for the specified alert. To use this resource, the requesting API Key must have the Project Read Only role.

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
			@return ListAlertConfigurationsByAlertIdApiRequest
	*/
	ListAlertConfigurationsByAlertId(ctx context.Context, groupId string, alertId string) ListAlertConfigurationsByAlertIdApiRequest
	/*
		ListAlertConfigurationsByAlertId Return All Alert Configurations Set for One Alert


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAlertConfigurationsByAlertIdApiParams - Parameters for the request
		@return ListAlertConfigurationsByAlertIdApiRequest
	*/
	ListAlertConfigurationsByAlertIdWithParams(ctx context.Context, args *ListAlertConfigurationsByAlertIdApiParams) ListAlertConfigurationsByAlertIdApiRequest

	// Method available only for mocking purposes
	ListAlertConfigurationsByAlertIdExecute(r ListAlertConfigurationsByAlertIdApiRequest) (*PaginatedAlertConfig, *http.Response, error)

	/*
			ToggleAlertConfiguration Toggle One State of One Alert Configuration in One Project

			Enables or disables the specified alert configuration in the specified project. The resource enables the specified alert configuration if currently enabled. The resource disables the specified alert configuration if currently disabled. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

		**NOTE**: This endpoint updates only the enabled/disabled state for the alert configuration. To update more than just this configuration, see [Update One Alert Configuration](#tag/Alert-Configurations/operation/updateAlertConfiguration).

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
			@return ToggleAlertConfigurationApiRequest
	*/
	ToggleAlertConfiguration(ctx context.Context, groupId string, alertConfigId string, alertsToggle *AlertsToggle) ToggleAlertConfigurationApiRequest
	/*
		ToggleAlertConfiguration Toggle One State of One Alert Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ToggleAlertConfigurationApiParams - Parameters for the request
		@return ToggleAlertConfigurationApiRequest
	*/
	ToggleAlertConfigurationWithParams(ctx context.Context, args *ToggleAlertConfigurationApiParams) ToggleAlertConfigurationApiRequest

	// Method available only for mocking purposes
	ToggleAlertConfigurationExecute(r ToggleAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error)

	/*
			UpdateAlertConfiguration Update One Alert Configuration for One Project

			Updates one alert configuration in the specified project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

		**NOTE**: To enable or disable the alert configuration, see [Toggle One State of One Alert Configuration in One Project](#tag/Alert-Configurations/operation/toggleAlertConfiguration).

		This resource remains under revision and may change.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
			@return UpdateAlertConfigurationApiRequest
	*/
	UpdateAlertConfiguration(ctx context.Context, groupId string, alertConfigId string, groupAlertsConfig *GroupAlertsConfig) UpdateAlertConfigurationApiRequest
	/*
		UpdateAlertConfiguration Update One Alert Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAlertConfigurationApiParams - Parameters for the request
		@return UpdateAlertConfigurationApiRequest
	*/
	UpdateAlertConfigurationWithParams(ctx context.Context, args *UpdateAlertConfigurationApiParams) UpdateAlertConfigurationApiRequest

	// Method available only for mocking purposes
	UpdateAlertConfigurationExecute(r UpdateAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error)
}

// AlertConfigurationsApiService AlertConfigurationsApi service
type AlertConfigurationsApiService service

type CreateAlertConfigurationApiRequest struct {
	ctx               context.Context
	ApiService        AlertConfigurationsApi
	groupId           string
	groupAlertsConfig *GroupAlertsConfig
}

type CreateAlertConfigurationApiParams struct {
	GroupId           string
	GroupAlertsConfig *GroupAlertsConfig
}

func (a *AlertConfigurationsApiService) CreateAlertConfigurationWithParams(ctx context.Context, args *CreateAlertConfigurationApiParams) CreateAlertConfigurationApiRequest {
	return CreateAlertConfigurationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		groupAlertsConfig: args.GroupAlertsConfig,
	}
}

func (r CreateAlertConfigurationApiRequest) Execute() (*GroupAlertsConfig, *http.Response, error) {
	return r.ApiService.CreateAlertConfigurationExecute(r)
}

/*
CreateAlertConfiguration Create One Alert Configuration in One Project

Creates one alert configuration for the specified project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateAlertConfigurationApiRequest
*/
func (a *AlertConfigurationsApiService) CreateAlertConfiguration(ctx context.Context, groupId string, groupAlertsConfig *GroupAlertsConfig) CreateAlertConfigurationApiRequest {
	return CreateAlertConfigurationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		groupAlertsConfig: groupAlertsConfig,
	}
}

// Execute executes the request
//
//	@return GroupAlertsConfig
func (a *AlertConfigurationsApiService) CreateAlertConfigurationExecute(r CreateAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupAlertsConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.CreateAlertConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupAlertsConfig == nil {
		return localVarReturnValue, nil, reportError("groupAlertsConfig is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupAlertsConfig
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

type DeleteAlertConfigurationApiRequest struct {
	ctx           context.Context
	ApiService    AlertConfigurationsApi
	groupId       string
	alertConfigId string
}

type DeleteAlertConfigurationApiParams struct {
	GroupId       string
	AlertConfigId string
}

func (a *AlertConfigurationsApiService) DeleteAlertConfigurationWithParams(ctx context.Context, args *DeleteAlertConfigurationApiParams) DeleteAlertConfigurationApiRequest {
	return DeleteAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		alertConfigId: args.AlertConfigId,
	}
}

func (r DeleteAlertConfigurationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAlertConfigurationExecute(r)
}

/*
DeleteAlertConfiguration Remove One Alert Configuration from One Project

Removes one alert configuration from the specified project. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
	@return DeleteAlertConfigurationApiRequest
*/
func (a *AlertConfigurationsApiService) DeleteAlertConfiguration(ctx context.Context, groupId string, alertConfigId string) DeleteAlertConfigurationApiRequest {
	return DeleteAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		alertConfigId: alertConfigId,
	}
}

// Execute executes the request
func (a *AlertConfigurationsApiService) DeleteAlertConfigurationExecute(r DeleteAlertConfigurationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.DeleteAlertConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertConfigId"+"}", url.PathEscape(parameterValueToString(r.alertConfigId, "alertConfigId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type GetAlertConfigurationApiRequest struct {
	ctx           context.Context
	ApiService    AlertConfigurationsApi
	groupId       string
	alertConfigId string
}

type GetAlertConfigurationApiParams struct {
	GroupId       string
	AlertConfigId string
}

func (a *AlertConfigurationsApiService) GetAlertConfigurationWithParams(ctx context.Context, args *GetAlertConfigurationApiParams) GetAlertConfigurationApiRequest {
	return GetAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		alertConfigId: args.AlertConfigId,
	}
}

func (r GetAlertConfigurationApiRequest) Execute() (*GroupAlertsConfig, *http.Response, error) {
	return r.ApiService.GetAlertConfigurationExecute(r)
}

/*
GetAlertConfiguration Return One Alert Configuration from One Project

Returns the specified alert configuration from the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
	@return GetAlertConfigurationApiRequest
*/
func (a *AlertConfigurationsApiService) GetAlertConfiguration(ctx context.Context, groupId string, alertConfigId string) GetAlertConfigurationApiRequest {
	return GetAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		alertConfigId: alertConfigId,
	}
}

// Execute executes the request
//
//	@return GroupAlertsConfig
func (a *AlertConfigurationsApiService) GetAlertConfigurationExecute(r GetAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupAlertsConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.GetAlertConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertConfigId"+"}", url.PathEscape(parameterValueToString(r.alertConfigId, "alertConfigId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListAlertConfigurationMatchersFieldNamesApiRequest struct {
	ctx        context.Context
	ApiService AlertConfigurationsApi
}

type ListAlertConfigurationMatchersFieldNamesApiParams struct {
}

func (a *AlertConfigurationsApiService) ListAlertConfigurationMatchersFieldNamesWithParams(ctx context.Context, args *ListAlertConfigurationMatchersFieldNamesApiParams) ListAlertConfigurationMatchersFieldNamesApiRequest {
	return ListAlertConfigurationMatchersFieldNamesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ListAlertConfigurationMatchersFieldNamesApiRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ListAlertConfigurationMatchersFieldNamesExecute(r)
}

/*
ListAlertConfigurationMatchersFieldNames Get All Alert Configuration Matchers Field Names

Get all field names that the `matchers.fieldName` parameter accepts when you create or update an Alert Configuration. You can successfully call this endpoint with any assigned role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListAlertConfigurationMatchersFieldNamesApiRequest
*/
func (a *AlertConfigurationsApiService) ListAlertConfigurationMatchersFieldNames(ctx context.Context) ListAlertConfigurationMatchersFieldNamesApiRequest {
	return ListAlertConfigurationMatchersFieldNamesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *AlertConfigurationsApiService) ListAlertConfigurationMatchersFieldNamesExecute(r ListAlertConfigurationMatchersFieldNamesApiRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.ListAlertConfigurationMatchersFieldNames")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/alertConfigs/matchers/fieldNames"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListAlertConfigurationsApiRequest struct {
	ctx          context.Context
	ApiService   AlertConfigurationsApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListAlertConfigurationsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *AlertConfigurationsApiService) ListAlertConfigurationsWithParams(ctx context.Context, args *ListAlertConfigurationsApiParams) ListAlertConfigurationsApiRequest {
	return ListAlertConfigurationsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListAlertConfigurationsApiRequest) IncludeCount(includeCount bool) ListAlertConfigurationsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAlertConfigurationsApiRequest) ItemsPerPage(itemsPerPage int) ListAlertConfigurationsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAlertConfigurationsApiRequest) PageNum(pageNum int) ListAlertConfigurationsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListAlertConfigurationsApiRequest) Execute() (*PaginatedAlertConfig, *http.Response, error) {
	return r.ApiService.ListAlertConfigurationsExecute(r)
}

/*
ListAlertConfigurations Return All Alert Configurations for One Project

Returns all alert configurations for one project. These alert configurations apply to any component in the project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListAlertConfigurationsApiRequest
*/
func (a *AlertConfigurationsApiService) ListAlertConfigurations(ctx context.Context, groupId string) ListAlertConfigurationsApiRequest {
	return ListAlertConfigurationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedAlertConfig
func (a *AlertConfigurationsApiService) ListAlertConfigurationsExecute(r ListAlertConfigurationsApiRequest) (*PaginatedAlertConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAlertConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.ListAlertConfigurations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListAlertConfigurationsByAlertIdApiRequest struct {
	ctx          context.Context
	ApiService   AlertConfigurationsApi
	groupId      string
	alertId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListAlertConfigurationsByAlertIdApiParams struct {
	GroupId      string
	AlertId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *AlertConfigurationsApiService) ListAlertConfigurationsByAlertIdWithParams(ctx context.Context, args *ListAlertConfigurationsByAlertIdApiParams) ListAlertConfigurationsByAlertIdApiRequest {
	return ListAlertConfigurationsByAlertIdApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		alertId:      args.AlertId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListAlertConfigurationsByAlertIdApiRequest) IncludeCount(includeCount bool) ListAlertConfigurationsByAlertIdApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListAlertConfigurationsByAlertIdApiRequest) ItemsPerPage(itemsPerPage int) ListAlertConfigurationsByAlertIdApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListAlertConfigurationsByAlertIdApiRequest) PageNum(pageNum int) ListAlertConfigurationsByAlertIdApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListAlertConfigurationsByAlertIdApiRequest) Execute() (*PaginatedAlertConfig, *http.Response, error) {
	return r.ApiService.ListAlertConfigurationsByAlertIdExecute(r)
}

/*
ListAlertConfigurationsByAlertId Return All Alert Configurations Set for One Alert

Returns all alert configurations set for the specified alert. To use this resource, the requesting API Key must have the Project Read Only role.

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertId Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.
	@return ListAlertConfigurationsByAlertIdApiRequest
*/
func (a *AlertConfigurationsApiService) ListAlertConfigurationsByAlertId(ctx context.Context, groupId string, alertId string) ListAlertConfigurationsByAlertIdApiRequest {
	return ListAlertConfigurationsByAlertIdApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		alertId:    alertId,
	}
}

// Execute executes the request
//
//	@return PaginatedAlertConfig
func (a *AlertConfigurationsApiService) ListAlertConfigurationsByAlertIdExecute(r ListAlertConfigurationsByAlertIdApiRequest) (*PaginatedAlertConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAlertConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.ListAlertConfigurationsByAlertId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertId"+"}", url.PathEscape(parameterValueToString(r.alertId, "alertId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ToggleAlertConfigurationApiRequest struct {
	ctx           context.Context
	ApiService    AlertConfigurationsApi
	groupId       string
	alertConfigId string
	alertsToggle  *AlertsToggle
}

type ToggleAlertConfigurationApiParams struct {
	GroupId       string
	AlertConfigId string
	AlertsToggle  *AlertsToggle
}

func (a *AlertConfigurationsApiService) ToggleAlertConfigurationWithParams(ctx context.Context, args *ToggleAlertConfigurationApiParams) ToggleAlertConfigurationApiRequest {
	return ToggleAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		alertConfigId: args.AlertConfigId,
		alertsToggle:  args.AlertsToggle,
	}
}

func (r ToggleAlertConfigurationApiRequest) Execute() (*GroupAlertsConfig, *http.Response, error) {
	return r.ApiService.ToggleAlertConfigurationExecute(r)
}

/*
ToggleAlertConfiguration Toggle One State of One Alert Configuration in One Project

Enables or disables the specified alert configuration in the specified project. The resource enables the specified alert configuration if currently enabled. The resource disables the specified alert configuration if currently disabled. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

**NOTE**: This endpoint updates only the enabled/disabled state for the alert configuration. To update more than just this configuration, see [Update One Alert Configuration](#tag/Alert-Configurations/operation/updateAlertConfiguration).

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
	@return ToggleAlertConfigurationApiRequest
*/
func (a *AlertConfigurationsApiService) ToggleAlertConfiguration(ctx context.Context, groupId string, alertConfigId string, alertsToggle *AlertsToggle) ToggleAlertConfigurationApiRequest {
	return ToggleAlertConfigurationApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		alertConfigId: alertConfigId,
		alertsToggle:  alertsToggle,
	}
}

// Execute executes the request
//
//	@return GroupAlertsConfig
func (a *AlertConfigurationsApiService) ToggleAlertConfigurationExecute(r ToggleAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupAlertsConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.ToggleAlertConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertConfigId"+"}", url.PathEscape(parameterValueToString(r.alertConfigId, "alertConfigId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alertsToggle == nil {
		return localVarReturnValue, nil, reportError("alertsToggle is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.alertsToggle
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

type UpdateAlertConfigurationApiRequest struct {
	ctx               context.Context
	ApiService        AlertConfigurationsApi
	groupId           string
	alertConfigId     string
	groupAlertsConfig *GroupAlertsConfig
}

type UpdateAlertConfigurationApiParams struct {
	GroupId           string
	AlertConfigId     string
	GroupAlertsConfig *GroupAlertsConfig
}

func (a *AlertConfigurationsApiService) UpdateAlertConfigurationWithParams(ctx context.Context, args *UpdateAlertConfigurationApiParams) UpdateAlertConfigurationApiRequest {
	return UpdateAlertConfigurationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		alertConfigId:     args.AlertConfigId,
		groupAlertsConfig: args.GroupAlertsConfig,
	}
}

func (r UpdateAlertConfigurationApiRequest) Execute() (*GroupAlertsConfig, *http.Response, error) {
	return r.ApiService.UpdateAlertConfigurationExecute(r)
}

/*
UpdateAlertConfiguration Update One Alert Configuration for One Project

Updates one alert configuration in the specified project. Alert configurations define the triggers and notification methods for alerts. To use this resource, the requesting API Key must have the Organization Owner or Project Owner role.

**NOTE**: To enable or disable the alert configuration, see [Toggle One State of One Alert Configuration in One Project](#tag/Alert-Configurations/operation/toggleAlertConfiguration).

This resource remains under revision and may change.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param alertConfigId Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.
	@return UpdateAlertConfigurationApiRequest
*/
func (a *AlertConfigurationsApiService) UpdateAlertConfiguration(ctx context.Context, groupId string, alertConfigId string, groupAlertsConfig *GroupAlertsConfig) UpdateAlertConfigurationApiRequest {
	return UpdateAlertConfigurationApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		alertConfigId:     alertConfigId,
		groupAlertsConfig: groupAlertsConfig,
	}
}

// Execute executes the request
//
//	@return GroupAlertsConfig
func (a *AlertConfigurationsApiService) UpdateAlertConfigurationExecute(r UpdateAlertConfigurationApiRequest) (*GroupAlertsConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupAlertsConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertConfigurationsApiService.UpdateAlertConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"alertConfigId"+"}", url.PathEscape(parameterValueToString(r.alertConfigId, "alertConfigId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupAlertsConfig == nil {
		return localVarReturnValue, nil, reportError("groupAlertsConfig is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupAlertsConfig
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
