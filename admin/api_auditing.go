// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AuditingApi interface {

	/*
		GetAuditingConfiguration Return Auditing Configuration for One Project

		Returns the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetAuditingConfigurationApiRequest
	*/
	GetAuditingConfiguration(ctx context.Context, groupId string) GetAuditingConfigurationApiRequest
	/*
		GetAuditingConfiguration Return Auditing Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAuditingConfigurationApiParams - Parameters for the request
		@return GetAuditingConfigurationApiRequest
	*/
	GetAuditingConfigurationWithParams(ctx context.Context, args *GetAuditingConfigurationApiParams) GetAuditingConfigurationApiRequest

	// Method available only for mocking purposes
	GetAuditingConfigurationExecute(r GetAuditingConfigurationApiRequest) (*AuditLog, *http.Response, error)

	/*
		UpdateAuditingConfiguration Update Auditing Configuration for One Project

		Updates the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param auditLog Updated auditing configuration for the specified project.
		@return UpdateAuditingConfigurationApiRequest
	*/
	UpdateAuditingConfiguration(ctx context.Context, groupId string, auditLog *AuditLog) UpdateAuditingConfigurationApiRequest
	/*
		UpdateAuditingConfiguration Update Auditing Configuration for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAuditingConfigurationApiParams - Parameters for the request
		@return UpdateAuditingConfigurationApiRequest
	*/
	UpdateAuditingConfigurationWithParams(ctx context.Context, args *UpdateAuditingConfigurationApiParams) UpdateAuditingConfigurationApiRequest

	// Method available only for mocking purposes
	UpdateAuditingConfigurationExecute(r UpdateAuditingConfigurationApiRequest) (*AuditLog, *http.Response, error)
}

// AuditingApiService AuditingApi service
type AuditingApiService service

type GetAuditingConfigurationApiRequest struct {
	ctx        context.Context
	ApiService AuditingApi
	groupId    string
}

type GetAuditingConfigurationApiParams struct {
	GroupId string
}

func (a *AuditingApiService) GetAuditingConfigurationWithParams(ctx context.Context, args *GetAuditingConfigurationApiParams) GetAuditingConfigurationApiRequest {
	return GetAuditingConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetAuditingConfigurationApiRequest) Execute() (*AuditLog, *http.Response, error) {
	return r.ApiService.GetAuditingConfigurationExecute(r)
}

/*
GetAuditingConfiguration Return Auditing Configuration for One Project

Returns the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetAuditingConfigurationApiRequest
*/
func (a *AuditingApiService) GetAuditingConfiguration(ctx context.Context, groupId string) GetAuditingConfigurationApiRequest {
	return GetAuditingConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetAuditingConfigurationExecute executes the request
//
//	@return AuditLog
func (a *AuditingApiService) GetAuditingConfigurationExecute(r GetAuditingConfigurationApiRequest) (*AuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditingApiService.GetAuditingConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/auditLog"
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

type UpdateAuditingConfigurationApiRequest struct {
	ctx        context.Context
	ApiService AuditingApi
	groupId    string
	auditLog   *AuditLog
}

type UpdateAuditingConfigurationApiParams struct {
	GroupId  string
	AuditLog *AuditLog
}

func (a *AuditingApiService) UpdateAuditingConfigurationWithParams(ctx context.Context, args *UpdateAuditingConfigurationApiParams) UpdateAuditingConfigurationApiRequest {
	return UpdateAuditingConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		auditLog:   args.AuditLog,
	}
}

func (r UpdateAuditingConfigurationApiRequest) Execute() (*AuditLog, *http.Response, error) {
	return r.ApiService.UpdateAuditingConfigurationExecute(r)
}

/*
UpdateAuditingConfiguration Update Auditing Configuration for One Project

Updates the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateAuditingConfigurationApiRequest
*/
func (a *AuditingApiService) UpdateAuditingConfiguration(ctx context.Context, groupId string, auditLog *AuditLog) UpdateAuditingConfigurationApiRequest {
	return UpdateAuditingConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		auditLog:   auditLog,
	}
}

// UpdateAuditingConfigurationExecute executes the request
//
//	@return AuditLog
func (a *AuditingApiService) UpdateAuditingConfigurationExecute(r UpdateAuditingConfigurationApiRequest) (*AuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *AuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditingApiService.UpdateAuditingConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/auditLog"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.auditLog == nil {
		return localVarReturnValue, nil, reportError("auditLog is required and must be specified")
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
	localVarPostBody = r.auditLog
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
