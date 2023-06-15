// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type EncryptionAtRestUsingCustomerKeyManagementApi interface {

	/*
			GetEncryptionAtRest Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project

			Returns the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role.

		**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return GetEncryptionAtRestApiRequest
	*/
	GetEncryptionAtRest(ctx context.Context, groupId string) GetEncryptionAtRestApiRequest
	/*
		GetEncryptionAtRest Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetEncryptionAtRestApiParams - Parameters for the request
		@return GetEncryptionAtRestApiRequest
	*/
	GetEncryptionAtRestWithParams(ctx context.Context, args *GetEncryptionAtRestApiParams) GetEncryptionAtRestApiRequest

	// Interface only available internally
	getEncryptionAtRestExecute(r GetEncryptionAtRestApiRequest) (*EncryptionAtRest, *http.Response, error)

	/*
			UpdateEncryptionAtRest Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project

			[experimental] Updates the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role. This feature isn't available for `M0` free clusters, `M2`, `M5`, or serverless clusters.

		 After you configure at least one Encryption at Rest using a Customer Key Management provider for the MongoDB Cloud project, Project Owners can enable Encryption at Rest using Customer Key Management for each MongoDB Cloud cluster for which they require encryption. The Encryption at Rest using Customer Key Management provider doesn't have to match the cluster cloud service provider. MongoDB Cloud doesn't automatically rotate user-managed encryption keys. Defer to your preferred Encryption at Rest using Customer Key Management provider's documentation and guidance for best practices on key rotation. MongoDB Cloud automatically creates a 90-day key rotation alert when you configure Encryption at Rest using Customer Key Management using your Key Management in an MongoDB Cloud project. MongoDB Cloud encrypts all storage whether or not you use your own key management.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return UpdateEncryptionAtRestApiRequest
	*/
	UpdateEncryptionAtRest(ctx context.Context, groupId string, encryptionAtRest *EncryptionAtRest) UpdateEncryptionAtRestApiRequest
	/*
		UpdateEncryptionAtRest Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateEncryptionAtRestApiParams - Parameters for the request
		@return UpdateEncryptionAtRestApiRequest
	*/
	UpdateEncryptionAtRestWithParams(ctx context.Context, args *UpdateEncryptionAtRestApiParams) UpdateEncryptionAtRestApiRequest

	// Interface only available internally
	updateEncryptionAtRestExecute(r UpdateEncryptionAtRestApiRequest) (*EncryptionAtRest, *http.Response, error)
}

// EncryptionAtRestUsingCustomerKeyManagementApiService EncryptionAtRestUsingCustomerKeyManagementApi service
type EncryptionAtRestUsingCustomerKeyManagementApiService service

type GetEncryptionAtRestApiRequest struct {
	ctx        context.Context
	ApiService EncryptionAtRestUsingCustomerKeyManagementApi
	groupId    string
}

type GetEncryptionAtRestApiParams struct {
	GroupId string
}

func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) GetEncryptionAtRestWithParams(ctx context.Context, args *GetEncryptionAtRestApiParams) GetEncryptionAtRestApiRequest {
	return GetEncryptionAtRestApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetEncryptionAtRestApiRequest) Execute() (*EncryptionAtRest, *http.Response, error) {
	return r.ApiService.getEncryptionAtRestExecute(r)
}

/*
GetEncryptionAtRest Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project

Returns the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role.

**LIMITED TO M10 OR GREATER:** MongoDB Cloud limits this feature to dedicated cluster tiers of M10 and greater.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetEncryptionAtRestApiRequest
*/
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) GetEncryptionAtRest(ctx context.Context, groupId string) GetEncryptionAtRestApiRequest {
	return GetEncryptionAtRestApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return EncryptionAtRest
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) getEncryptionAtRestExecute(r GetEncryptionAtRestApiRequest) (*EncryptionAtRest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EncryptionAtRest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncryptionAtRestUsingCustomerKeyManagementApiService.GetEncryptionAtRest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/encryptionAtRest"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateEncryptionAtRestApiRequest struct {
	ctx              context.Context
	ApiService       EncryptionAtRestUsingCustomerKeyManagementApi
	groupId          string
	encryptionAtRest *EncryptionAtRest
}

type UpdateEncryptionAtRestApiParams struct {
	GroupId          string
	EncryptionAtRest *EncryptionAtRest
}

func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) UpdateEncryptionAtRestWithParams(ctx context.Context, args *UpdateEncryptionAtRestApiParams) UpdateEncryptionAtRestApiRequest {
	return UpdateEncryptionAtRestApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		encryptionAtRest: args.EncryptionAtRest,
	}
}

func (r UpdateEncryptionAtRestApiRequest) Execute() (*EncryptionAtRest, *http.Response, error) {
	return r.ApiService.updateEncryptionAtRestExecute(r)
}

/*
UpdateEncryptionAtRest Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project

[experimental] Updates the configuration for encryption at rest using the keys you manage through your cloud provider. MongoDB Cloud encrypts all storage even if you don't use your own key management. This resource requires the requesting API Key to have the Project Owner role. This feature isn't available for `M0` free clusters, `M2`, `M5`, or serverless clusters.

	After you configure at least one Encryption at Rest using a Customer Key Management provider for the MongoDB Cloud project, Project Owners can enable Encryption at Rest using Customer Key Management for each MongoDB Cloud cluster for which they require encryption. The Encryption at Rest using Customer Key Management provider doesn't have to match the cluster cloud service provider. MongoDB Cloud doesn't automatically rotate user-managed encryption keys. Defer to your preferred Encryption at Rest using Customer Key Management provider's documentation and guidance for best practices on key rotation. MongoDB Cloud automatically creates a 90-day key rotation alert when you configure Encryption at Rest using Customer Key Management using your Key Management in an MongoDB Cloud project. MongoDB Cloud encrypts all storage whether or not you use your own key management.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateEncryptionAtRestApiRequest
*/
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) UpdateEncryptionAtRest(ctx context.Context, groupId string, encryptionAtRest *EncryptionAtRest) UpdateEncryptionAtRestApiRequest {
	return UpdateEncryptionAtRestApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		encryptionAtRest: encryptionAtRest,
	}
}

// Execute executes the request
//
//	@return EncryptionAtRest
func (a *EncryptionAtRestUsingCustomerKeyManagementApiService) updateEncryptionAtRestExecute(r UpdateEncryptionAtRestApiRequest) (*EncryptionAtRest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EncryptionAtRest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncryptionAtRestUsingCustomerKeyManagementApiService.UpdateEncryptionAtRest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/encryptionAtRest"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.encryptionAtRest == nil {
		return localVarReturnValue, nil, reportError("encryptionAtRest is required and must be specified")
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
	localVarPostBody = r.encryptionAtRest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
