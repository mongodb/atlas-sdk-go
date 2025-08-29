// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PushBasedLogExportApi interface {

	/*
		CreateLogExport Create One Push-Based Log Export Configuration in One Project

		Configures the project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param createPushBasedLogExportProjectRequest The project configuration details. The S3 bucket name, IAM role ID, and prefix path fields are required.
		@return CreateLogExportApiRequest
	*/
	CreateLogExport(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest) CreateLogExportApiRequest
	/*
		CreateLogExport Create One Push-Based Log Export Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateLogExportApiParams - Parameters for the request
		@return CreateLogExportApiRequest
	*/
	CreateLogExportWithParams(ctx context.Context, args *CreateLogExportApiParams) CreateLogExportApiRequest

	// Method available only for mocking purposes
	CreateLogExportExecute(r CreateLogExportApiRequest) (*http.Response, error)

	/*
		DeleteLogExport Disable Push-Based Log Export for One Project

		Disables the push-based log export feature by resetting the project level settings to its default configuration.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteLogExportApiRequest
	*/
	DeleteLogExport(ctx context.Context, groupId string) DeleteLogExportApiRequest
	/*
		DeleteLogExport Disable Push-Based Log Export for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLogExportApiParams - Parameters for the request
		@return DeleteLogExportApiRequest
	*/
	DeleteLogExportWithParams(ctx context.Context, args *DeleteLogExportApiParams) DeleteLogExportApiRequest

	// Method available only for mocking purposes
	DeleteLogExportExecute(r DeleteLogExportApiRequest) (*http.Response, error)

	/*
		GetLogExport Return One Push-Based Log Export Configuration in One Project

		Fetches the current project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetLogExportApiRequest
	*/
	GetLogExport(ctx context.Context, groupId string) GetLogExportApiRequest
	/*
		GetLogExport Return One Push-Based Log Export Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLogExportApiParams - Parameters for the request
		@return GetLogExportApiRequest
	*/
	GetLogExportWithParams(ctx context.Context, args *GetLogExportApiParams) GetLogExportApiRequest

	// Method available only for mocking purposes
	GetLogExportExecute(r GetLogExportApiRequest) (*PushBasedLogExportProject, *http.Response, error)

	/*
		UpdateLogExport Update One Push-Based Log Export Configuration in One Project

		Updates the project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pushBasedLogExportProject The project configuration details. The S3 bucket name, IAM role ID, and prefix path fields are the only fields that may be specified. Fields left unspecified will not be modified.
		@return UpdateLogExportApiRequest
	*/
	UpdateLogExport(ctx context.Context, groupId string, pushBasedLogExportProject *PushBasedLogExportProject) UpdateLogExportApiRequest
	/*
		UpdateLogExport Update One Push-Based Log Export Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateLogExportApiParams - Parameters for the request
		@return UpdateLogExportApiRequest
	*/
	UpdateLogExportWithParams(ctx context.Context, args *UpdateLogExportApiParams) UpdateLogExportApiRequest

	// Method available only for mocking purposes
	UpdateLogExportExecute(r UpdateLogExportApiRequest) (*http.Response, error)
}

// PushBasedLogExportApiService PushBasedLogExportApi service
type PushBasedLogExportApiService service

type CreateLogExportApiRequest struct {
	ctx                                    context.Context
	ApiService                             PushBasedLogExportApi
	groupId                                string
	createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest
}

type CreateLogExportApiParams struct {
	GroupId                                string
	CreatePushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest
}

func (a *PushBasedLogExportApiService) CreateLogExportWithParams(ctx context.Context, args *CreateLogExportApiParams) CreateLogExportApiRequest {
	return CreateLogExportApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		createPushBasedLogExportProjectRequest: args.CreatePushBasedLogExportProjectRequest,
	}
}

func (r CreateLogExportApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateLogExportExecute(r)
}

/*
CreateLogExport Create One Push-Based Log Export Configuration in One Project

Configures the project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateLogExportApiRequest
*/
func (a *PushBasedLogExportApiService) CreateLogExport(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest) CreateLogExportApiRequest {
	return CreateLogExportApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		createPushBasedLogExportProjectRequest: createPushBasedLogExportProjectRequest,
	}
}

// CreateLogExportExecute executes the request
func (a *PushBasedLogExportApiService) CreateLogExportExecute(r CreateLogExportApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.CreateLogExport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createPushBasedLogExportProjectRequest == nil {
		return nil, reportError("createPushBasedLogExportProjectRequest is required and must be specified")
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
	localVarPostBody = r.createPushBasedLogExportProjectRequest
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

type DeleteLogExportApiRequest struct {
	ctx        context.Context
	ApiService PushBasedLogExportApi
	groupId    string
}

type DeleteLogExportApiParams struct {
	GroupId string
}

func (a *PushBasedLogExportApiService) DeleteLogExportWithParams(ctx context.Context, args *DeleteLogExportApiParams) DeleteLogExportApiRequest {
	return DeleteLogExportApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeleteLogExportApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLogExportExecute(r)
}

/*
DeleteLogExport Disable Push-Based Log Export for One Project

Disables the push-based log export feature by resetting the project level settings to its default configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteLogExportApiRequest
*/
func (a *PushBasedLogExportApiService) DeleteLogExport(ctx context.Context, groupId string) DeleteLogExportApiRequest {
	return DeleteLogExportApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// DeleteLogExportExecute executes the request
func (a *PushBasedLogExportApiService) DeleteLogExportExecute(r DeleteLogExportApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.DeleteLogExport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
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

type GetLogExportApiRequest struct {
	ctx        context.Context
	ApiService PushBasedLogExportApi
	groupId    string
}

type GetLogExportApiParams struct {
	GroupId string
}

func (a *PushBasedLogExportApiService) GetLogExportWithParams(ctx context.Context, args *GetLogExportApiParams) GetLogExportApiRequest {
	return GetLogExportApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetLogExportApiRequest) Execute() (*PushBasedLogExportProject, *http.Response, error) {
	return r.ApiService.GetLogExportExecute(r)
}

/*
GetLogExport Return One Push-Based Log Export Configuration in One Project

Fetches the current project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetLogExportApiRequest
*/
func (a *PushBasedLogExportApiService) GetLogExport(ctx context.Context, groupId string) GetLogExportApiRequest {
	return GetLogExportApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetLogExportExecute executes the request
//
//	@return PushBasedLogExportProject
func (a *PushBasedLogExportApiService) GetLogExportExecute(r GetLogExportApiRequest) (*PushBasedLogExportProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PushBasedLogExportProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.GetLogExport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
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

type UpdateLogExportApiRequest struct {
	ctx                       context.Context
	ApiService                PushBasedLogExportApi
	groupId                   string
	pushBasedLogExportProject *PushBasedLogExportProject
}

type UpdateLogExportApiParams struct {
	GroupId                   string
	PushBasedLogExportProject *PushBasedLogExportProject
}

func (a *PushBasedLogExportApiService) UpdateLogExportWithParams(ctx context.Context, args *UpdateLogExportApiParams) UpdateLogExportApiRequest {
	return UpdateLogExportApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		pushBasedLogExportProject: args.PushBasedLogExportProject,
	}
}

func (r UpdateLogExportApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateLogExportExecute(r)
}

/*
UpdateLogExport Update One Push-Based Log Export Configuration in One Project

Updates the project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateLogExportApiRequest
*/
func (a *PushBasedLogExportApiService) UpdateLogExport(ctx context.Context, groupId string, pushBasedLogExportProject *PushBasedLogExportProject) UpdateLogExportApiRequest {
	return UpdateLogExportApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		pushBasedLogExportProject: pushBasedLogExportProject,
	}
}

// UpdateLogExportExecute executes the request
func (a *PushBasedLogExportApiService) UpdateLogExportExecute(r UpdateLogExportApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.UpdateLogExport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pushBasedLogExportProject == nil {
		return nil, reportError("pushBasedLogExportProject is required and must be specified")
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
	localVarPostBody = r.pushBasedLogExportProject
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
