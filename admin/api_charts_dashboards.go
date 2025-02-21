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

		Exports the specified Charts dashboard. To use this resource, the requesting API Key must have the Project Read Only role.

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
	ExportChartsDashboardExecute(r ExportChartsDashboardApiRequest) (io.ReadCloser, *http.Response, error)

	/*
		ImportChartsDashboard Import One Charts Dashboard

		Imports the Charts dashboard that the template specifies. Optionally, you can specify a `dashboardId` with `override=true` to import into an existing dashboard. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param apiChartsDashboardImportRequestView20430101 Details to import or override a dashboard.
		@return ImportChartsDashboardApiRequest
	*/
	ImportChartsDashboard(ctx context.Context, groupId string, apiChartsDashboardImportRequestView20430101 *ApiChartsDashboardImportRequestView20430101) ImportChartsDashboardApiRequest
	/*
		ImportChartsDashboard Import One Charts Dashboard


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ImportChartsDashboardApiParams - Parameters for the request
		@return ImportChartsDashboardApiRequest
	*/
	ImportChartsDashboardWithParams(ctx context.Context, args *ImportChartsDashboardApiParams) ImportChartsDashboardApiRequest

	// Method available only for mocking purposes
	ImportChartsDashboardExecute(r ImportChartsDashboardApiRequest) (*ApiChartsDashboardImportResponseView20430101, *http.Response, error)
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

func (r ExportChartsDashboardApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.ExportChartsDashboardExecute(r)
}

/*
ExportChartsDashboard Export One Charts Dashboard

Exports the specified Charts dashboard. To use this resource, the requesting API Key must have the Project Read Only role.

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
//	@return io.ReadCloser
func (a *ChartsDashboardsApiService) ExportChartsDashboardExecute(r ExportChartsDashboardApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartsDashboardsApiService.ExportChartsDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/chartsDashboards/{dashboardId}:export"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2043-01-01+json"}

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

type ImportChartsDashboardApiRequest struct {
	ctx                                         context.Context
	ApiService                                  ChartsDashboardsApi
	groupId                                     string
	apiChartsDashboardImportRequestView20430101 *ApiChartsDashboardImportRequestView20430101
}

type ImportChartsDashboardApiParams struct {
	GroupId                                     string
	ApiChartsDashboardImportRequestView20430101 *ApiChartsDashboardImportRequestView20430101
}

func (a *ChartsDashboardsApiService) ImportChartsDashboardWithParams(ctx context.Context, args *ImportChartsDashboardApiParams) ImportChartsDashboardApiRequest {
	return ImportChartsDashboardApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		apiChartsDashboardImportRequestView20430101: args.ApiChartsDashboardImportRequestView20430101,
	}
}

func (r ImportChartsDashboardApiRequest) Execute() (*ApiChartsDashboardImportResponseView20430101, *http.Response, error) {
	return r.ApiService.ImportChartsDashboardExecute(r)
}

/*
ImportChartsDashboard Import One Charts Dashboard

Imports the Charts dashboard that the template specifies. Optionally, you can specify a `dashboardId` with `override=true` to import into an existing dashboard. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ImportChartsDashboardApiRequest
*/
func (a *ChartsDashboardsApiService) ImportChartsDashboard(ctx context.Context, groupId string, apiChartsDashboardImportRequestView20430101 *ApiChartsDashboardImportRequestView20430101) ImportChartsDashboardApiRequest {
	return ImportChartsDashboardApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		apiChartsDashboardImportRequestView20430101: apiChartsDashboardImportRequestView20430101,
	}
}

// ImportChartsDashboardExecute executes the request
//
//	@return ApiChartsDashboardImportResponseView20430101
func (a *ChartsDashboardsApiService) ImportChartsDashboardExecute(r ImportChartsDashboardApiRequest) (*ApiChartsDashboardImportResponseView20430101, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiChartsDashboardImportResponseView20430101
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartsDashboardsApiService.ImportChartsDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/chartsDashboards:import"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiChartsDashboardImportRequestView20430101 == nil {
		return localVarReturnValue, nil, reportError("apiChartsDashboardImportRequestView20430101 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2043-01-01+multipart"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2043-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiChartsDashboardImportRequestView20430101
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
