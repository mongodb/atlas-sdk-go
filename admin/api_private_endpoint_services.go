// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PrivateEndpointServicesApi interface {

	/*
		CreatePrivateEndpoint Create One Private Endpoint for One Provider

		Creates one private endpoint for the specified cloud service provider. This cloud service provider manages the private endpoint service, which in turn manages the private endpoints for the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. To learn more about considerations, limitations, and prerequisites, see the MongoDB documentation for setting up a private endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint.
		@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.
		@param createEndpointRequest Creates one private endpoint for the specified cloud service provider.
		@return CreatePrivateEndpointApiRequest
	*/
	CreatePrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string, createEndpointRequest *CreateEndpointRequest) CreatePrivateEndpointApiRequest
	/*
		CreatePrivateEndpoint Create One Private Endpoint for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePrivateEndpointApiParams - Parameters for the request
		@return CreatePrivateEndpointApiRequest
	*/
	CreatePrivateEndpointWithParams(ctx context.Context, args *CreatePrivateEndpointApiParams) CreatePrivateEndpointApiRequest

	// Method available only for mocking purposes
	CreatePrivateEndpointExecute(r CreatePrivateEndpointApiRequest) (*PrivateLinkEndpoint, *http.Response, error)

	/*
		CreatePrivateEndpointService Create One Private Endpoint Service for One Provider

		Creates one private endpoint service for the specified cloud service provider. This cloud service provider manages the private endpoint service for the project. When you create a private endpoint service, MongoDB Cloud creates a network container in the project for the cloud provider for which you create the private endpoint service if one doesn't already exist. To learn more about private endpoint terminology in MongoDB Cloud, see Private Endpoint Concepts. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProviderEndpointServiceRequest Creates one private endpoint for the specified cloud service provider.
		@return CreatePrivateEndpointServiceApiRequest
	*/
	CreatePrivateEndpointService(ctx context.Context, groupId string, cloudProviderEndpointServiceRequest *CloudProviderEndpointServiceRequest) CreatePrivateEndpointServiceApiRequest
	/*
		CreatePrivateEndpointService Create One Private Endpoint Service for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePrivateEndpointServiceApiParams - Parameters for the request
		@return CreatePrivateEndpointServiceApiRequest
	*/
	CreatePrivateEndpointServiceWithParams(ctx context.Context, args *CreatePrivateEndpointServiceApiParams) CreatePrivateEndpointServiceApiRequest

	// Method available only for mocking purposes
	CreatePrivateEndpointServiceExecute(r CreatePrivateEndpointServiceApiRequest) (*EndpointService, *http.Response, error)

	/*
		DeletePrivateEndpoint Remove One Private Endpoint for One Provider

		Removes one private endpoint from the specified project and private endpoint service, as managed by the specified cloud service provider. When the last private endpoint is removed from a given private endpoint service, that private endpoint service is also removed. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint.
		@param endpointId Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
		@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.
		@return DeletePrivateEndpointApiRequest
	*/
	DeletePrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointId string, endpointServiceId string) DeletePrivateEndpointApiRequest
	/*
		DeletePrivateEndpoint Remove One Private Endpoint for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePrivateEndpointApiParams - Parameters for the request
		@return DeletePrivateEndpointApiRequest
	*/
	DeletePrivateEndpointWithParams(ctx context.Context, args *DeletePrivateEndpointApiParams) DeletePrivateEndpointApiRequest

	// Method available only for mocking purposes
	DeletePrivateEndpointExecute(r DeletePrivateEndpointApiRequest) (*http.Response, error)

	/*
		DeletePrivateEndpointService Remove One Private Endpoint Service for One Provider

		Removes one private endpoint service from the specified project. This cloud service provider manages the private endpoint service that belongs to the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint service.
		@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.
		@return DeletePrivateEndpointServiceApiRequest
	*/
	DeletePrivateEndpointService(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string) DeletePrivateEndpointServiceApiRequest
	/*
		DeletePrivateEndpointService Remove One Private Endpoint Service for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePrivateEndpointServiceApiParams - Parameters for the request
		@return DeletePrivateEndpointServiceApiRequest
	*/
	DeletePrivateEndpointServiceWithParams(ctx context.Context, args *DeletePrivateEndpointServiceApiParams) DeletePrivateEndpointServiceApiRequest

	// Method available only for mocking purposes
	DeletePrivateEndpointServiceExecute(r DeletePrivateEndpointServiceApiRequest) (*http.Response, error)

	/*
		GetPrivateEndpoint Return One Private Endpoint for One Provider

		Returns the connection state of the specified private endpoint. The private endpoint service manages this private endpoint which belongs to one project hosted from one cloud service provider. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint.
		@param endpointId Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
		@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.
		@return GetPrivateEndpointApiRequest
	*/
	GetPrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointId string, endpointServiceId string) GetPrivateEndpointApiRequest
	/*
		GetPrivateEndpoint Return One Private Endpoint for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPrivateEndpointApiParams - Parameters for the request
		@return GetPrivateEndpointApiRequest
	*/
	GetPrivateEndpointWithParams(ctx context.Context, args *GetPrivateEndpointApiParams) GetPrivateEndpointApiRequest

	// Method available only for mocking purposes
	GetPrivateEndpointExecute(r GetPrivateEndpointApiRequest) (*PrivateLinkEndpoint, *http.Response, error)

	/*
		GetPrivateEndpointService Return One Private Endpoint Service for One Provider

		Returns the name, interfaces, and state of the specified private endpoint service from one project. The cloud service provider hosted this private endpoint service that belongs to the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint service.
		@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.
		@return GetPrivateEndpointServiceApiRequest
	*/
	GetPrivateEndpointService(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string) GetPrivateEndpointServiceApiRequest
	/*
		GetPrivateEndpointService Return One Private Endpoint Service for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPrivateEndpointServiceApiParams - Parameters for the request
		@return GetPrivateEndpointServiceApiRequest
	*/
	GetPrivateEndpointServiceWithParams(ctx context.Context, args *GetPrivateEndpointServiceApiParams) GetPrivateEndpointServiceApiRequest

	// Method available only for mocking purposes
	GetPrivateEndpointServiceExecute(r GetPrivateEndpointServiceApiRequest) (*EndpointService, *http.Response, error)

	/*
		GetRegionalizedPrivateEndpointSetting Return Regionalized Private Endpoint Status

		Checks whether each region in the specified cloud service provider can create multiple private endpoints per region. The cloud service provider manages the private endpoint for the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetRegionalizedPrivateEndpointSettingApiRequest
	*/
	GetRegionalizedPrivateEndpointSetting(ctx context.Context, groupId string) GetRegionalizedPrivateEndpointSettingApiRequest
	/*
		GetRegionalizedPrivateEndpointSetting Return Regionalized Private Endpoint Status


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetRegionalizedPrivateEndpointSettingApiParams - Parameters for the request
		@return GetRegionalizedPrivateEndpointSettingApiRequest
	*/
	GetRegionalizedPrivateEndpointSettingWithParams(ctx context.Context, args *GetRegionalizedPrivateEndpointSettingApiParams) GetRegionalizedPrivateEndpointSettingApiRequest

	// Method available only for mocking purposes
	GetRegionalizedPrivateEndpointSettingExecute(r GetRegionalizedPrivateEndpointSettingApiRequest) (*ProjectSettingItem, *http.Response, error)

	/*
		ListPrivateEndpointServices Return All Private Endpoint Services for One Provider

		Returns the name, interfaces, and state of all private endpoint services for the specified cloud service provider. This cloud service provider manages the private endpoint service for the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Cloud service provider that manages this private endpoint service.
		@return ListPrivateEndpointServicesApiRequest
	*/
	ListPrivateEndpointServices(ctx context.Context, groupId string, cloudProvider string) ListPrivateEndpointServicesApiRequest
	/*
		ListPrivateEndpointServices Return All Private Endpoint Services for One Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPrivateEndpointServicesApiParams - Parameters for the request
		@return ListPrivateEndpointServicesApiRequest
	*/
	ListPrivateEndpointServicesWithParams(ctx context.Context, args *ListPrivateEndpointServicesApiParams) ListPrivateEndpointServicesApiRequest

	// Method available only for mocking purposes
	ListPrivateEndpointServicesExecute(r ListPrivateEndpointServicesApiRequest) ([]EndpointService, *http.Response, error)

	/*
		ToggleRegionalizedPrivateEndpointSetting Toggle Regionalized Private Endpoint Status

		Enables or disables the ability to create multiple private endpoints per region in all cloud service providers in one project. The cloud service provider manages the private endpoints for the project. Connection strings to existing multi-region and global sharded clusters change when you enable this setting. You must update your applications to use the new connection strings. This might cause downtime. To use this resource, the requesting Service Account or API Key must have the Project Owner role and all clusters in the deployment must be sharded clusters. Once enabled, you cannot create replica sets.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param projectSettingItem Enables or disables the ability to create multiple private endpoints per region in all cloud service providers in one project.
		@return ToggleRegionalizedPrivateEndpointSettingApiRequest
	*/
	ToggleRegionalizedPrivateEndpointSetting(ctx context.Context, groupId string, projectSettingItem *ProjectSettingItem) ToggleRegionalizedPrivateEndpointSettingApiRequest
	/*
		ToggleRegionalizedPrivateEndpointSetting Toggle Regionalized Private Endpoint Status


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ToggleRegionalizedPrivateEndpointSettingApiParams - Parameters for the request
		@return ToggleRegionalizedPrivateEndpointSettingApiRequest
	*/
	ToggleRegionalizedPrivateEndpointSettingWithParams(ctx context.Context, args *ToggleRegionalizedPrivateEndpointSettingApiParams) ToggleRegionalizedPrivateEndpointSettingApiRequest

	// Method available only for mocking purposes
	ToggleRegionalizedPrivateEndpointSettingExecute(r ToggleRegionalizedPrivateEndpointSettingApiRequest) (*ProjectSettingItem, *http.Response, error)
}

// PrivateEndpointServicesApiService PrivateEndpointServicesApi service
type PrivateEndpointServicesApiService service

type CreatePrivateEndpointApiRequest struct {
	ctx                   context.Context
	ApiService            PrivateEndpointServicesApi
	groupId               string
	cloudProvider         string
	endpointServiceId     string
	createEndpointRequest *CreateEndpointRequest
}

type CreatePrivateEndpointApiParams struct {
	GroupId               string
	CloudProvider         string
	EndpointServiceId     string
	CreateEndpointRequest *CreateEndpointRequest
}

func (a *PrivateEndpointServicesApiService) CreatePrivateEndpointWithParams(ctx context.Context, args *CreatePrivateEndpointApiParams) CreatePrivateEndpointApiRequest {
	return CreatePrivateEndpointApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               args.GroupId,
		cloudProvider:         args.CloudProvider,
		endpointServiceId:     args.EndpointServiceId,
		createEndpointRequest: args.CreateEndpointRequest,
	}
}

func (r CreatePrivateEndpointApiRequest) Execute() (*PrivateLinkEndpoint, *http.Response, error) {
	return r.ApiService.CreatePrivateEndpointExecute(r)
}

/*
CreatePrivateEndpoint Create One Private Endpoint for One Provider

Creates one private endpoint for the specified cloud service provider. This cloud service provider manages the private endpoint service, which in turn manages the private endpoints for the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. To learn more about considerations, limitations, and prerequisites, see the MongoDB documentation for setting up a private endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint.
	@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.
	@return CreatePrivateEndpointApiRequest
*/
func (a *PrivateEndpointServicesApiService) CreatePrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string, createEndpointRequest *CreateEndpointRequest) CreatePrivateEndpointApiRequest {
	return CreatePrivateEndpointApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               groupId,
		cloudProvider:         cloudProvider,
		endpointServiceId:     endpointServiceId,
		createEndpointRequest: createEndpointRequest,
	}
}

// CreatePrivateEndpointExecute executes the request
//
//	@return PrivateLinkEndpoint
func (a *PrivateEndpointServicesApiService) CreatePrivateEndpointExecute(r CreatePrivateEndpointApiRequest) (*PrivateLinkEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PrivateLinkEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.CreatePrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return localVarReturnValue, nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)
	if r.endpointServiceId == "" {
		return localVarReturnValue, nil, reportError("endpointServiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointServiceId"+"}", url.PathEscape(r.endpointServiceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("createEndpointRequest is required and must be specified")
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
	localVarPostBody = r.createEndpointRequest
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

type CreatePrivateEndpointServiceApiRequest struct {
	ctx                                 context.Context
	ApiService                          PrivateEndpointServicesApi
	groupId                             string
	cloudProviderEndpointServiceRequest *CloudProviderEndpointServiceRequest
}

type CreatePrivateEndpointServiceApiParams struct {
	GroupId                             string
	CloudProviderEndpointServiceRequest *CloudProviderEndpointServiceRequest
}

func (a *PrivateEndpointServicesApiService) CreatePrivateEndpointServiceWithParams(ctx context.Context, args *CreatePrivateEndpointServiceApiParams) CreatePrivateEndpointServiceApiRequest {
	return CreatePrivateEndpointServiceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             args.GroupId,
		cloudProviderEndpointServiceRequest: args.CloudProviderEndpointServiceRequest,
	}
}

func (r CreatePrivateEndpointServiceApiRequest) Execute() (*EndpointService, *http.Response, error) {
	return r.ApiService.CreatePrivateEndpointServiceExecute(r)
}

/*
CreatePrivateEndpointService Create One Private Endpoint Service for One Provider

Creates one private endpoint service for the specified cloud service provider. This cloud service provider manages the private endpoint service for the project. When you create a private endpoint service, MongoDB Cloud creates a network container in the project for the cloud provider for which you create the private endpoint service if one doesn't already exist. To learn more about private endpoint terminology in MongoDB Cloud, see Private Endpoint Concepts. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePrivateEndpointServiceApiRequest
*/
func (a *PrivateEndpointServicesApiService) CreatePrivateEndpointService(ctx context.Context, groupId string, cloudProviderEndpointServiceRequest *CloudProviderEndpointServiceRequest) CreatePrivateEndpointServiceApiRequest {
	return CreatePrivateEndpointServiceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             groupId,
		cloudProviderEndpointServiceRequest: cloudProviderEndpointServiceRequest,
	}
}

// CreatePrivateEndpointServiceExecute executes the request
//
//	@return EndpointService
func (a *PrivateEndpointServicesApiService) CreatePrivateEndpointServiceExecute(r CreatePrivateEndpointServiceApiRequest) (*EndpointService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EndpointService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.CreatePrivateEndpointService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProviderEndpointServiceRequest == nil {
		return localVarReturnValue, nil, reportError("cloudProviderEndpointServiceRequest is required and must be specified")
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
	localVarPostBody = r.cloudProviderEndpointServiceRequest
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

type DeletePrivateEndpointApiRequest struct {
	ctx               context.Context
	ApiService        PrivateEndpointServicesApi
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
}

type DeletePrivateEndpointApiParams struct {
	GroupId           string
	CloudProvider     string
	EndpointId        string
	EndpointServiceId string
}

func (a *PrivateEndpointServicesApiService) DeletePrivateEndpointWithParams(ctx context.Context, args *DeletePrivateEndpointApiParams) DeletePrivateEndpointApiRequest {
	return DeletePrivateEndpointApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		cloudProvider:     args.CloudProvider,
		endpointId:        args.EndpointId,
		endpointServiceId: args.EndpointServiceId,
	}
}

func (r DeletePrivateEndpointApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePrivateEndpointExecute(r)
}

/*
DeletePrivateEndpoint Remove One Private Endpoint for One Provider

Removes one private endpoint from the specified project and private endpoint service, as managed by the specified cloud service provider. When the last private endpoint is removed from a given private endpoint service, that private endpoint service is also removed. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint.
	@param endpointId Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
	@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.
	@return DeletePrivateEndpointApiRequest
*/
func (a *PrivateEndpointServicesApiService) DeletePrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointId string, endpointServiceId string) DeletePrivateEndpointApiRequest {
	return DeletePrivateEndpointApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		cloudProvider:     cloudProvider,
		endpointId:        endpointId,
		endpointServiceId: endpointServiceId,
	}
}

// DeletePrivateEndpointExecute executes the request
func (a *PrivateEndpointServicesApiService) DeletePrivateEndpointExecute(r DeletePrivateEndpointApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.DeletePrivateEndpoint")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)
	if r.endpointId == "" {
		return nil, reportError("endpointId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(r.endpointId), -1)
	if r.endpointServiceId == "" {
		return nil, reportError("endpointServiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointServiceId"+"}", url.PathEscape(r.endpointServiceId), -1)

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

type DeletePrivateEndpointServiceApiRequest struct {
	ctx               context.Context
	ApiService        PrivateEndpointServicesApi
	groupId           string
	cloudProvider     string
	endpointServiceId string
}

type DeletePrivateEndpointServiceApiParams struct {
	GroupId           string
	CloudProvider     string
	EndpointServiceId string
}

func (a *PrivateEndpointServicesApiService) DeletePrivateEndpointServiceWithParams(ctx context.Context, args *DeletePrivateEndpointServiceApiParams) DeletePrivateEndpointServiceApiRequest {
	return DeletePrivateEndpointServiceApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		cloudProvider:     args.CloudProvider,
		endpointServiceId: args.EndpointServiceId,
	}
}

func (r DeletePrivateEndpointServiceApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePrivateEndpointServiceExecute(r)
}

/*
DeletePrivateEndpointService Remove One Private Endpoint Service for One Provider

Removes one private endpoint service from the specified project. This cloud service provider manages the private endpoint service that belongs to the project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint service.
	@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.
	@return DeletePrivateEndpointServiceApiRequest
*/
func (a *PrivateEndpointServicesApiService) DeletePrivateEndpointService(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string) DeletePrivateEndpointServiceApiRequest {
	return DeletePrivateEndpointServiceApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		cloudProvider:     cloudProvider,
		endpointServiceId: endpointServiceId,
	}
}

// DeletePrivateEndpointServiceExecute executes the request
func (a *PrivateEndpointServicesApiService) DeletePrivateEndpointServiceExecute(r DeletePrivateEndpointServiceApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.DeletePrivateEndpointService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)
	if r.endpointServiceId == "" {
		return nil, reportError("endpointServiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointServiceId"+"}", url.PathEscape(r.endpointServiceId), -1)

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

type GetPrivateEndpointApiRequest struct {
	ctx               context.Context
	ApiService        PrivateEndpointServicesApi
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
}

type GetPrivateEndpointApiParams struct {
	GroupId           string
	CloudProvider     string
	EndpointId        string
	EndpointServiceId string
}

func (a *PrivateEndpointServicesApiService) GetPrivateEndpointWithParams(ctx context.Context, args *GetPrivateEndpointApiParams) GetPrivateEndpointApiRequest {
	return GetPrivateEndpointApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		cloudProvider:     args.CloudProvider,
		endpointId:        args.EndpointId,
		endpointServiceId: args.EndpointServiceId,
	}
}

func (r GetPrivateEndpointApiRequest) Execute() (*PrivateLinkEndpoint, *http.Response, error) {
	return r.ApiService.GetPrivateEndpointExecute(r)
}

/*
GetPrivateEndpoint Return One Private Endpoint for One Provider

Returns the connection state of the specified private endpoint. The private endpoint service manages this private endpoint which belongs to one project hosted from one cloud service provider. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint.
	@param endpointId Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
	@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.
	@return GetPrivateEndpointApiRequest
*/
func (a *PrivateEndpointServicesApiService) GetPrivateEndpoint(ctx context.Context, groupId string, cloudProvider string, endpointId string, endpointServiceId string) GetPrivateEndpointApiRequest {
	return GetPrivateEndpointApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		cloudProvider:     cloudProvider,
		endpointId:        endpointId,
		endpointServiceId: endpointServiceId,
	}
}

// GetPrivateEndpointExecute executes the request
//
//	@return PrivateLinkEndpoint
func (a *PrivateEndpointServicesApiService) GetPrivateEndpointExecute(r GetPrivateEndpointApiRequest) (*PrivateLinkEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PrivateLinkEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.GetPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return localVarReturnValue, nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)
	if r.endpointId == "" {
		return localVarReturnValue, nil, reportError("endpointId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(r.endpointId), -1)
	if r.endpointServiceId == "" {
		return localVarReturnValue, nil, reportError("endpointServiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointServiceId"+"}", url.PathEscape(r.endpointServiceId), -1)

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

type GetPrivateEndpointServiceApiRequest struct {
	ctx               context.Context
	ApiService        PrivateEndpointServicesApi
	groupId           string
	cloudProvider     string
	endpointServiceId string
}

type GetPrivateEndpointServiceApiParams struct {
	GroupId           string
	CloudProvider     string
	EndpointServiceId string
}

func (a *PrivateEndpointServicesApiService) GetPrivateEndpointServiceWithParams(ctx context.Context, args *GetPrivateEndpointServiceApiParams) GetPrivateEndpointServiceApiRequest {
	return GetPrivateEndpointServiceApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		cloudProvider:     args.CloudProvider,
		endpointServiceId: args.EndpointServiceId,
	}
}

func (r GetPrivateEndpointServiceApiRequest) Execute() (*EndpointService, *http.Response, error) {
	return r.ApiService.GetPrivateEndpointServiceExecute(r)
}

/*
GetPrivateEndpointService Return One Private Endpoint Service for One Provider

Returns the name, interfaces, and state of the specified private endpoint service from one project. The cloud service provider hosted this private endpoint service that belongs to the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint service.
	@param endpointServiceId Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.
	@return GetPrivateEndpointServiceApiRequest
*/
func (a *PrivateEndpointServicesApiService) GetPrivateEndpointService(ctx context.Context, groupId string, cloudProvider string, endpointServiceId string) GetPrivateEndpointServiceApiRequest {
	return GetPrivateEndpointServiceApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		cloudProvider:     cloudProvider,
		endpointServiceId: endpointServiceId,
	}
}

// GetPrivateEndpointServiceExecute executes the request
//
//	@return EndpointService
func (a *PrivateEndpointServicesApiService) GetPrivateEndpointServiceExecute(r GetPrivateEndpointServiceApiRequest) (*EndpointService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EndpointService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.GetPrivateEndpointService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return localVarReturnValue, nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)
	if r.endpointServiceId == "" {
		return localVarReturnValue, nil, reportError("endpointServiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"endpointServiceId"+"}", url.PathEscape(r.endpointServiceId), -1)

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

type GetRegionalizedPrivateEndpointSettingApiRequest struct {
	ctx        context.Context
	ApiService PrivateEndpointServicesApi
	groupId    string
}

type GetRegionalizedPrivateEndpointSettingApiParams struct {
	GroupId string
}

func (a *PrivateEndpointServicesApiService) GetRegionalizedPrivateEndpointSettingWithParams(ctx context.Context, args *GetRegionalizedPrivateEndpointSettingApiParams) GetRegionalizedPrivateEndpointSettingApiRequest {
	return GetRegionalizedPrivateEndpointSettingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetRegionalizedPrivateEndpointSettingApiRequest) Execute() (*ProjectSettingItem, *http.Response, error) {
	return r.ApiService.GetRegionalizedPrivateEndpointSettingExecute(r)
}

/*
GetRegionalizedPrivateEndpointSetting Return Regionalized Private Endpoint Status

Checks whether each region in the specified cloud service provider can create multiple private endpoints per region. The cloud service provider manages the private endpoint for the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetRegionalizedPrivateEndpointSettingApiRequest
*/
func (a *PrivateEndpointServicesApiService) GetRegionalizedPrivateEndpointSetting(ctx context.Context, groupId string) GetRegionalizedPrivateEndpointSettingApiRequest {
	return GetRegionalizedPrivateEndpointSettingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// GetRegionalizedPrivateEndpointSettingExecute executes the request
//
//	@return ProjectSettingItem
func (a *PrivateEndpointServicesApiService) GetRegionalizedPrivateEndpointSettingExecute(r GetRegionalizedPrivateEndpointSettingApiRequest) (*ProjectSettingItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ProjectSettingItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.GetRegionalizedPrivateEndpointSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode"
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

type ListPrivateEndpointServicesApiRequest struct {
	ctx           context.Context
	ApiService    PrivateEndpointServicesApi
	groupId       string
	cloudProvider string
}

type ListPrivateEndpointServicesApiParams struct {
	GroupId       string
	CloudProvider string
}

func (a *PrivateEndpointServicesApiService) ListPrivateEndpointServicesWithParams(ctx context.Context, args *ListPrivateEndpointServicesApiParams) ListPrivateEndpointServicesApiRequest {
	return ListPrivateEndpointServicesApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		cloudProvider: args.CloudProvider,
	}
}

func (r ListPrivateEndpointServicesApiRequest) Execute() ([]EndpointService, *http.Response, error) {
	return r.ApiService.ListPrivateEndpointServicesExecute(r)
}

/*
ListPrivateEndpointServices Return All Private Endpoint Services for One Provider

Returns the name, interfaces, and state of all private endpoint services for the specified cloud service provider. This cloud service provider manages the private endpoint service for the project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Cloud service provider that manages this private endpoint service.
	@return ListPrivateEndpointServicesApiRequest
*/
func (a *PrivateEndpointServicesApiService) ListPrivateEndpointServices(ctx context.Context, groupId string, cloudProvider string) ListPrivateEndpointServicesApiRequest {
	return ListPrivateEndpointServicesApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		cloudProvider: cloudProvider,
	}
}

// ListPrivateEndpointServicesExecute executes the request
//
//	@return []EndpointService
func (a *PrivateEndpointServicesApiService) ListPrivateEndpointServicesExecute(r ListPrivateEndpointServicesApiRequest) ([]EndpointService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []EndpointService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.ListPrivateEndpointServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.cloudProvider == "" {
		return localVarReturnValue, nil, reportError("cloudProvider is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(r.cloudProvider), -1)

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

type ToggleRegionalizedPrivateEndpointSettingApiRequest struct {
	ctx                context.Context
	ApiService         PrivateEndpointServicesApi
	groupId            string
	projectSettingItem *ProjectSettingItem
}

type ToggleRegionalizedPrivateEndpointSettingApiParams struct {
	GroupId            string
	ProjectSettingItem *ProjectSettingItem
}

func (a *PrivateEndpointServicesApiService) ToggleRegionalizedPrivateEndpointSettingWithParams(ctx context.Context, args *ToggleRegionalizedPrivateEndpointSettingApiParams) ToggleRegionalizedPrivateEndpointSettingApiRequest {
	return ToggleRegionalizedPrivateEndpointSettingApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		projectSettingItem: args.ProjectSettingItem,
	}
}

func (r ToggleRegionalizedPrivateEndpointSettingApiRequest) Execute() (*ProjectSettingItem, *http.Response, error) {
	return r.ApiService.ToggleRegionalizedPrivateEndpointSettingExecute(r)
}

/*
ToggleRegionalizedPrivateEndpointSetting Toggle Regionalized Private Endpoint Status

Enables or disables the ability to create multiple private endpoints per region in all cloud service providers in one project. The cloud service provider manages the private endpoints for the project. Connection strings to existing multi-region and global sharded clusters change when you enable this setting. You must update your applications to use the new connection strings. This might cause downtime. To use this resource, the requesting Service Account or API Key must have the Project Owner role and all clusters in the deployment must be sharded clusters. Once enabled, you cannot create replica sets.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ToggleRegionalizedPrivateEndpointSettingApiRequest
*/
func (a *PrivateEndpointServicesApiService) ToggleRegionalizedPrivateEndpointSetting(ctx context.Context, groupId string, projectSettingItem *ProjectSettingItem) ToggleRegionalizedPrivateEndpointSettingApiRequest {
	return ToggleRegionalizedPrivateEndpointSettingApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		projectSettingItem: projectSettingItem,
	}
}

// ToggleRegionalizedPrivateEndpointSettingExecute executes the request
//
//	@return ProjectSettingItem
func (a *PrivateEndpointServicesApiService) ToggleRegionalizedPrivateEndpointSettingExecute(r ToggleRegionalizedPrivateEndpointSettingApiRequest) (*ProjectSettingItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ProjectSettingItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateEndpointServicesApiService.ToggleRegionalizedPrivateEndpointSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.projectSettingItem == nil {
		return localVarReturnValue, nil, reportError("projectSettingItem is required and must be specified")
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
	localVarPostBody = r.projectSettingItem
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
