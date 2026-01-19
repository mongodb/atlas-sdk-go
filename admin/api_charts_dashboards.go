// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ChartsDashboardsApi interface {

	/*
		ExportChartsDashboard Export One Charts Dashboard

		Exports the specified Charts dashboard. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param dashboardId ID of the dashboard to export.
		@return ExportChartsDashboardApiRequest
	*/
	ExportChartsDashboard(ctx context.Context, groupId string, dashboardId string) ExportChartsDashboardApiRequest
	/*
		ExportChartsDashboard Export One Charts Dashboard


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ExportChartsDashboardApiParams - Parameters for the request
		@return ExportChartsDashboardApiRequest
	*/
	ExportChartsDashboardWithParams(ctx context.Context, args *ExportChartsDashboardApiParams) ExportChartsDashboardApiRequest

	// Method available only for mocking purposes
	ExportChartsDashboardExecute(r ExportChartsDashboardApiRequest) (string, *http.Response, error)

	/*
		ImportChartsDashboards Import One Charts Dashboard

		Imports the Charts dashboard that the template specifies. Optionally, you can specify `overwrite=true` to import into an existing dashboard. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param body Schema corresponding to the response fetched from an exported dashboard.
		@return ImportChartsDashboardsApiRequest
	*/
	ImportChartsDashboards(ctx context.Context, groupId string, body *any) ImportChartsDashboardsApiRequest
	/*
		ImportChartsDashboards Import One Charts Dashboard


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ImportChartsDashboardsApiParams - Parameters for the request
		@return ImportChartsDashboardsApiRequest
	*/
	ImportChartsDashboardsWithParams(ctx context.Context, args *ImportChartsDashboardsApiParams) ImportChartsDashboardsApiRequest

	// Method available only for mocking purposes
	ImportChartsDashboardsExecute(r ImportChartsDashboardsApiRequest) (*ApiChartsDashboardImportResponse, *http.Response, error)
}

// ChartsDashboardsApiService ChartsDashboardsApi service
type ChartsDashboardsApiService service

type ExportChartsDashboardApiRequest struct {
	ctx         context.Context
	ApiService  ChartsDashboardsApi
	groupId     string
	dashboardId string
}

type ExportChartsDashboardApiParams struct {
	GroupId     string
	DashboardId string
}

func (a *ChartsDashboardsApiService) ExportChartsDashboardWithParams(ctx context.Context, args *ExportChartsDashboardApiParams) ExportChartsDashboardApiRequest {
	return ExportChartsDashboardApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		dashboardId: args.DashboardId,
	}
}

func (r ExportChartsDashboardApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ExportChartsDashboardExecute(r)
}

/*
ExportChartsDashboard Export One Charts Dashboard

Exports the specified Charts dashboard. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param dashboardId ID of the dashboard to export.
	@return ExportChartsDashboardApiRequest
*/
func (a *ChartsDashboardsApiService) ExportChartsDashboard(ctx context.Context, groupId string, dashboardId string) ExportChartsDashboardApiRequest {
	return ExportChartsDashboardApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		dashboardId: dashboardId,
	}
}

// ExportChartsDashboardExecute executes the request
//
//	@return string
func (a *ChartsDashboardsApiService) ExportChartsDashboardExecute(r ExportChartsDashboardApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartsDashboardsApiService.ExportChartsDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/chartsDashboards/{dashboardId}:export"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.dashboardId == "" {
		return localVarReturnValue, nil, reportError("dashboardId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardId"+"}", url.PathEscape(r.dashboardId), -1)

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

type ImportChartsDashboardsApiRequest struct {
	ctx        context.Context
	ApiService ChartsDashboardsApi
	groupId    string
	body       *any
	overwrite  *bool
}

type ImportChartsDashboardsApiParams struct {
	GroupId   string
	Body      *any
	Overwrite *bool
}

func (a *ChartsDashboardsApiService) ImportChartsDashboardsWithParams(ctx context.Context, args *ImportChartsDashboardsApiParams) ImportChartsDashboardsApiRequest {
	return ImportChartsDashboardsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		body:       args.Body,
		overwrite:  args.Overwrite,
	}
}

// Setting this to true enables importing into an existing dashboard.
func (r ImportChartsDashboardsApiRequest) Overwrite(overwrite bool) ImportChartsDashboardsApiRequest {
	r.overwrite = &overwrite
	return r
}

func (r ImportChartsDashboardsApiRequest) Execute() (*ApiChartsDashboardImportResponse, *http.Response, error) {
	return r.ApiService.ImportChartsDashboardsExecute(r)
}

/*
ImportChartsDashboards Import One Charts Dashboard

Imports the Charts dashboard that the template specifies. Optionally, you can specify `overwrite=true` to import into an existing dashboard. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ImportChartsDashboardsApiRequest
*/
func (a *ChartsDashboardsApiService) ImportChartsDashboards(ctx context.Context, groupId string, body *any) ImportChartsDashboardsApiRequest {
	return ImportChartsDashboardsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		body:       body,
	}
}

// ImportChartsDashboardsExecute executes the request
//
//	@return ApiChartsDashboardImportResponse
func (a *ChartsDashboardsApiService) ImportChartsDashboardsExecute(r ImportChartsDashboardsApiRequest) (*ApiChartsDashboardImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiChartsDashboardImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartsDashboardsApiService.ImportChartsDashboards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/chartsDashboards:import"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.overwrite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "overwrite", r.overwrite, "")
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
	localVarPostBody = r.body
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
