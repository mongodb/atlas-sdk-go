// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type LDAPConfigurationApi interface {

	/*
		DeleteLdapConfiguration Remove LDAP User to DN Mapping

		Removes the current LDAP Distinguished Name mapping captured in the ``userToDNMapping`` document from the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteLdapConfigurationApiRequest
	*/
	DeleteLdapConfiguration(ctx context.Context, groupId string) DeleteLdapConfigurationApiRequest
	/*
		DeleteLdapConfiguration Remove LDAP User to DN Mapping


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLdapConfigurationApiParams - Parameters for the request
		@return DeleteLdapConfigurationApiRequest
	*/
	DeleteLdapConfigurationWithParams(ctx context.Context, args *DeleteLdapConfigurationApiParams) DeleteLdapConfigurationApiRequest

	// Method available only for mocking purposes
	DeleteLdapConfigurationExecute(r DeleteLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		GetLdapConfiguration Return LDAP or X.509 Configuration

		Returns the current LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetLdapConfigurationApiRequest
	*/
	GetLdapConfiguration(ctx context.Context, groupId string) GetLdapConfigurationApiRequest
	/*
		GetLdapConfiguration Return LDAP or X.509 Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLdapConfigurationApiParams - Parameters for the request
		@return GetLdapConfigurationApiRequest
	*/
	GetLdapConfigurationWithParams(ctx context.Context, args *GetLdapConfigurationApiParams) GetLdapConfigurationApiRequest

	// Method available only for mocking purposes
	GetLdapConfigurationExecute(r GetLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		GetLdapConfigurationStatus Return Status of LDAP Configuration Verification in One Project

		Returns the status of one request to verify one LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param requestId Unique string that identifies the request to verify an Lightweight Directory Access Protocol (LDAP) configuration.
		@return GetLdapConfigurationStatusApiRequest
	*/
	GetLdapConfigurationStatus(ctx context.Context, groupId string, requestId string) GetLdapConfigurationStatusApiRequest
	/*
		GetLdapConfigurationStatus Return Status of LDAP Configuration Verification in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLdapConfigurationStatusApiParams - Parameters for the request
		@return GetLdapConfigurationStatusApiRequest
	*/
	GetLdapConfigurationStatusWithParams(ctx context.Context, args *GetLdapConfigurationStatusApiParams) GetLdapConfigurationStatusApiRequest

	// Method available only for mocking purposes
	GetLdapConfigurationStatusExecute(r GetLdapConfigurationStatusApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error)

	/*
			SaveLdapConfiguration Update LDAP or X.509 Configuration

			Edits the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		Updating this configuration triggers a rolling restart of the database.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param userSecurity Updates the LDAP configuration for the specified project.
			@return SaveLdapConfigurationApiRequest
	*/
	SaveLdapConfiguration(ctx context.Context, groupId string, userSecurity *UserSecurity) SaveLdapConfigurationApiRequest
	/*
		SaveLdapConfiguration Update LDAP or X.509 Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param SaveLdapConfigurationApiParams - Parameters for the request
		@return SaveLdapConfigurationApiRequest
	*/
	SaveLdapConfigurationWithParams(ctx context.Context, args *SaveLdapConfigurationApiParams) SaveLdapConfigurationApiRequest

	// Method available only for mocking purposes
	SaveLdapConfigurationExecute(r SaveLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		VerifyLdapConfiguration Verify LDAP Configuration in One Project

		Verifies the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param lDAPVerifyConnectivityJobRequestParams The LDAP configuration for the specified project that you want to verify.
		@return VerifyLdapConfigurationApiRequest
	*/
	VerifyLdapConfiguration(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams) VerifyLdapConfigurationApiRequest
	/*
		VerifyLdapConfiguration Verify LDAP Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param VerifyLdapConfigurationApiParams - Parameters for the request
		@return VerifyLdapConfigurationApiRequest
	*/
	VerifyLdapConfigurationWithParams(ctx context.Context, args *VerifyLdapConfigurationApiParams) VerifyLdapConfigurationApiRequest

	// Method available only for mocking purposes
	VerifyLdapConfigurationExecute(r VerifyLdapConfigurationApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error)
}

// LDAPConfigurationApiService LDAPConfigurationApi service
type LDAPConfigurationApiService service

type DeleteLdapConfigurationApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
}

type DeleteLdapConfigurationApiParams struct {
	GroupId string
}

func (a *LDAPConfigurationApiService) DeleteLdapConfigurationWithParams(ctx context.Context, args *DeleteLdapConfigurationApiParams) DeleteLdapConfigurationApiRequest {
	return DeleteLdapConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeleteLdapConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.DeleteLdapConfigurationExecute(r)
}

/*
DeleteLdapConfiguration Remove LDAP User to DN Mapping

Removes the current LDAP Distinguished Name mapping captured in the “userToDNMapping“ document from the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteLdapConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) DeleteLdapConfiguration(ctx context.Context, groupId string) DeleteLdapConfigurationApiRequest {
	return DeleteLdapConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// DeleteLdapConfigurationExecute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) DeleteLdapConfigurationExecute(r DeleteLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.DeleteLdapConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping"
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

type GetLdapConfigurationApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
}

type GetLdapConfigurationApiParams struct {
	GroupId string
}

func (a *LDAPConfigurationApiService) GetLdapConfigurationWithParams(ctx context.Context, args *GetLdapConfigurationApiParams) GetLdapConfigurationApiRequest {
	return GetLdapConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetLdapConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.GetLdapConfigurationExecute(r)
}

/*
GetLdapConfiguration Return LDAP or X.509 Configuration

Returns the current LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetLdapConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) GetLdapConfiguration(ctx context.Context, groupId string) GetLdapConfigurationApiRequest {
	return GetLdapConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetLdapConfigurationExecute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) GetLdapConfigurationExecute(r GetLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.GetLdapConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity"
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

type GetLdapConfigurationStatusApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
	requestId  string
}

type GetLdapConfigurationStatusApiParams struct {
	GroupId   string
	RequestId string
}

func (a *LDAPConfigurationApiService) GetLdapConfigurationStatusWithParams(ctx context.Context, args *GetLdapConfigurationStatusApiParams) GetLdapConfigurationStatusApiRequest {
	return GetLdapConfigurationStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		requestId:  args.RequestId,
	}
}

func (r GetLdapConfigurationStatusApiRequest) Execute() (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	return r.ApiService.GetLdapConfigurationStatusExecute(r)
}

/*
GetLdapConfigurationStatus Return Status of LDAP Configuration Verification in One Project

Returns the status of one request to verify one LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param requestId Unique string that identifies the request to verify an Lightweight Directory Access Protocol (LDAP) configuration.
	@return GetLdapConfigurationStatusApiRequest
*/
func (a *LDAPConfigurationApiService) GetLdapConfigurationStatus(ctx context.Context, groupId string, requestId string) GetLdapConfigurationStatusApiRequest {
	return GetLdapConfigurationStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		requestId:  requestId,
	}
}

// GetLdapConfigurationStatusExecute executes the request
//
//	@return LDAPVerifyConnectivityJobRequest
func (a *LDAPConfigurationApiService) GetLdapConfigurationStatusExecute(r GetLdapConfigurationStatusApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LDAPVerifyConnectivityJobRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.GetLdapConfigurationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.requestId == "" {
		return localVarReturnValue, nil, reportError("requestId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(r.requestId), -1)

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

type SaveLdapConfigurationApiRequest struct {
	ctx          context.Context
	ApiService   LDAPConfigurationApi
	groupId      string
	userSecurity *UserSecurity
}

type SaveLdapConfigurationApiParams struct {
	GroupId      string
	UserSecurity *UserSecurity
}

func (a *LDAPConfigurationApiService) SaveLdapConfigurationWithParams(ctx context.Context, args *SaveLdapConfigurationApiParams) SaveLdapConfigurationApiRequest {
	return SaveLdapConfigurationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		userSecurity: args.UserSecurity,
	}
}

func (r SaveLdapConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.SaveLdapConfigurationExecute(r)
}

/*
SaveLdapConfiguration Update LDAP or X.509 Configuration

Edits the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

Updating this configuration triggers a rolling restart of the database.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SaveLdapConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) SaveLdapConfiguration(ctx context.Context, groupId string, userSecurity *UserSecurity) SaveLdapConfigurationApiRequest {
	return SaveLdapConfigurationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		userSecurity: userSecurity,
	}
}

// SaveLdapConfigurationExecute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) SaveLdapConfigurationExecute(r SaveLdapConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.SaveLdapConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userSecurity == nil {
		return localVarReturnValue, nil, reportError("userSecurity is required and must be specified")
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
	localVarPostBody = r.userSecurity
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

type VerifyLdapConfigurationApiRequest struct {
	ctx                                    context.Context
	ApiService                             LDAPConfigurationApi
	groupId                                string
	lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams
}

type VerifyLdapConfigurationApiParams struct {
	GroupId                                string
	LDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams
}

func (a *LDAPConfigurationApiService) VerifyLdapConfigurationWithParams(ctx context.Context, args *VerifyLdapConfigurationApiParams) VerifyLdapConfigurationApiRequest {
	return VerifyLdapConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		lDAPVerifyConnectivityJobRequestParams: args.LDAPVerifyConnectivityJobRequestParams,
	}
}

func (r VerifyLdapConfigurationApiRequest) Execute() (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	return r.ApiService.VerifyLdapConfigurationExecute(r)
}

/*
VerifyLdapConfiguration Verify LDAP Configuration in One Project

Verifies the LDAP configuration for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return VerifyLdapConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) VerifyLdapConfiguration(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams) VerifyLdapConfigurationApiRequest {
	return VerifyLdapConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		lDAPVerifyConnectivityJobRequestParams: lDAPVerifyConnectivityJobRequestParams,
	}
}

// VerifyLdapConfigurationExecute executes the request
//
//	@return LDAPVerifyConnectivityJobRequest
func (a *LDAPConfigurationApiService) VerifyLdapConfigurationExecute(r VerifyLdapConfigurationApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LDAPVerifyConnectivityJobRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.VerifyLdapConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPVerifyConnectivityJobRequestParams == nil {
		return localVarReturnValue, nil, reportError("lDAPVerifyConnectivityJobRequestParams is required and must be specified")
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
	localVarPostBody = r.lDAPVerifyConnectivityJobRequestParams
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
