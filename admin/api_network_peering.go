// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NetworkPeeringApi interface {

	/*
		CreatePeeringConnection Create One Network Peering Connection

		Creates one new network peering connection in the specified project. Network peering allows multiple cloud-hosted applications to securely connect to the same project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. To learn more about considerations and prerequisites, see the Network Peering Documentation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param baseNetworkPeeringConnectionSettings Create one network peering connection.
		@return CreatePeeringConnectionApiRequest
	*/
	CreatePeeringConnection(ctx context.Context, groupId string, baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings) CreatePeeringConnectionApiRequest
	/*
		CreatePeeringConnection Create One Network Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePeeringConnectionApiParams - Parameters for the request
		@return CreatePeeringConnectionApiRequest
	*/
	CreatePeeringConnectionWithParams(ctx context.Context, args *CreatePeeringConnectionApiParams) CreatePeeringConnectionApiRequest

	// Method available only for mocking purposes
	CreatePeeringConnectionExecute(r CreatePeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error)

	/*
		CreatePeeringContainer Create One Network Peering Container

		Creates one new network peering container in the specified project. MongoDB Cloud can deploy Network Peering connections in a network peering container. GCP can have one container per project. AWS and Azure can have one container per cloud provider region. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProviderContainer Creates one new network peering container in the specified project.
		@return CreatePeeringContainerApiRequest
	*/
	CreatePeeringContainer(ctx context.Context, groupId string, cloudProviderContainer *CloudProviderContainer) CreatePeeringContainerApiRequest
	/*
		CreatePeeringContainer Create One Network Peering Container


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePeeringContainerApiParams - Parameters for the request
		@return CreatePeeringContainerApiRequest
	*/
	CreatePeeringContainerWithParams(ctx context.Context, args *CreatePeeringContainerApiParams) CreatePeeringContainerApiRequest

	// Method available only for mocking purposes
	CreatePeeringContainerExecute(r CreatePeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error)

	/*
		DeletePeeringConnection Remove One Network Peering Connection

		Removes one network peering connection in the specified project. If you Removes the last network peering connection associated with a project, MongoDB Cloud also removes any AWS security groups from the project IP access list. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.
		@return DeletePeeringConnectionApiRequest
	*/
	DeletePeeringConnection(ctx context.Context, groupId string, peerId string) DeletePeeringConnectionApiRequest
	/*
		DeletePeeringConnection Remove One Network Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePeeringConnectionApiParams - Parameters for the request
		@return DeletePeeringConnectionApiRequest
	*/
	DeletePeeringConnectionWithParams(ctx context.Context, args *DeletePeeringConnectionApiParams) DeletePeeringConnectionApiRequest

	// Method available only for mocking purposes
	DeletePeeringConnectionExecute(r DeletePeeringConnectionApiRequest) (any, *http.Response, error)

	/*
		DeletePeeringContainer Remove One Network Peering Container

		Removes one network peering container in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
		@return DeletePeeringContainerApiRequest
	*/
	DeletePeeringContainer(ctx context.Context, groupId string, containerId string) DeletePeeringContainerApiRequest
	/*
		DeletePeeringContainer Remove One Network Peering Container


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePeeringContainerApiParams - Parameters for the request
		@return DeletePeeringContainerApiRequest
	*/
	DeletePeeringContainerWithParams(ctx context.Context, args *DeletePeeringContainerApiParams) DeletePeeringContainerApiRequest

	// Method available only for mocking purposes
	DeletePeeringContainerExecute(r DeletePeeringContainerApiRequest) (*http.Response, error)

	/*
		DisablePeering Disable Connect via Peering-Only Mode for One Project

		Disables Connect via Peering Only mode for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param privateIPMode Disables Connect via Peering Only mode for the specified project.
		@return DisablePeeringApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for NetworkPeeringApi
	*/
	DisablePeering(ctx context.Context, groupId string, privateIPMode *PrivateIPMode) DisablePeeringApiRequest
	/*
		DisablePeering Disable Connect via Peering-Only Mode for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DisablePeeringApiParams - Parameters for the request
		@return DisablePeeringApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for NetworkPeeringApi
	*/
	DisablePeeringWithParams(ctx context.Context, args *DisablePeeringApiParams) DisablePeeringApiRequest

	// Method available only for mocking purposes
	DisablePeeringExecute(r DisablePeeringApiRequest) (*PrivateIPMode, *http.Response, error)

	/*
		GetPeeringConnection Return One Network Peering Connection in One Project

		Returns details about one specified network peering connection in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.
		@return GetPeeringConnectionApiRequest
	*/
	GetPeeringConnection(ctx context.Context, groupId string, peerId string) GetPeeringConnectionApiRequest
	/*
		GetPeeringConnection Return One Network Peering Connection in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPeeringConnectionApiParams - Parameters for the request
		@return GetPeeringConnectionApiRequest
	*/
	GetPeeringConnectionWithParams(ctx context.Context, args *GetPeeringConnectionApiParams) GetPeeringConnectionApiRequest

	// Method available only for mocking purposes
	GetPeeringConnectionExecute(r GetPeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error)

	/*
		GetPeeringContainer Return One Network Peering Container

		Returns details about one network peering container in one specified project. Network peering containers contain network peering connections. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container.
		@return GetPeeringContainerApiRequest
	*/
	GetPeeringContainer(ctx context.Context, groupId string, containerId string) GetPeeringContainerApiRequest
	/*
		GetPeeringContainer Return One Network Peering Container


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPeeringContainerApiParams - Parameters for the request
		@return GetPeeringContainerApiRequest
	*/
	GetPeeringContainerWithParams(ctx context.Context, args *GetPeeringContainerApiParams) GetPeeringContainerApiRequest

	// Method available only for mocking purposes
	GetPeeringContainerExecute(r GetPeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error)

	/*
		ListPeeringConnections Return All Network Peering Connections in One Project

		Returns details about all network peering connections in the specified project. Network peering allows multiple cloud-hosted applications to securely connect to the same project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListPeeringConnectionsApiRequest
	*/
	ListPeeringConnections(ctx context.Context, groupId string) ListPeeringConnectionsApiRequest
	/*
		ListPeeringConnections Return All Network Peering Connections in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPeeringConnectionsApiParams - Parameters for the request
		@return ListPeeringConnectionsApiRequest
	*/
	ListPeeringConnectionsWithParams(ctx context.Context, args *ListPeeringConnectionsApiParams) ListPeeringConnectionsApiRequest

	// Method available only for mocking purposes
	ListPeeringConnectionsExecute(r ListPeeringConnectionsApiRequest) (*PaginatedContainerPeer, *http.Response, error)

	/*
		ListPeeringContainerByCloudProvider Return All Network Peering Containers in One Project for One Cloud Provider

		Returns details about all network peering containers in the specified project for the specified cloud provider. If you do not specify the cloud provider, MongoDB Cloud returns details about all network peering containers in the project for Amazon Web Services (AWS). To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListPeeringContainerByCloudProviderApiRequest
	*/
	ListPeeringContainerByCloudProvider(ctx context.Context, groupId string) ListPeeringContainerByCloudProviderApiRequest
	/*
		ListPeeringContainerByCloudProvider Return All Network Peering Containers in One Project for One Cloud Provider


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPeeringContainerByCloudProviderApiParams - Parameters for the request
		@return ListPeeringContainerByCloudProviderApiRequest
	*/
	ListPeeringContainerByCloudProviderWithParams(ctx context.Context, args *ListPeeringContainerByCloudProviderApiParams) ListPeeringContainerByCloudProviderApiRequest

	// Method available only for mocking purposes
	ListPeeringContainerByCloudProviderExecute(r ListPeeringContainerByCloudProviderApiRequest) (*PaginatedCloudProviderContainer, *http.Response, error)

	/*
		ListPeeringContainers Return All Network Peering Containers in One Project

		Returns details about all network peering containers in the specified project. Network peering containers contain network peering connections. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListPeeringContainersApiRequest
	*/
	ListPeeringContainers(ctx context.Context, groupId string) ListPeeringContainersApiRequest
	/*
		ListPeeringContainers Return All Network Peering Containers in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPeeringContainersApiParams - Parameters for the request
		@return ListPeeringContainersApiRequest
	*/
	ListPeeringContainersWithParams(ctx context.Context, args *ListPeeringContainersApiParams) ListPeeringContainersApiRequest

	// Method available only for mocking purposes
	ListPeeringContainersExecute(r ListPeeringContainersApiRequest) (*PaginatedCloudProviderContainer, *http.Response, error)

	/*
		UpdatePeeringConnection Update One Network Peering Connection

		Updates one specified network peering connection in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.
		@param baseNetworkPeeringConnectionSettings Modify one network peering connection.
		@return UpdatePeeringConnectionApiRequest
	*/
	UpdatePeeringConnection(ctx context.Context, groupId string, peerId string, baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings) UpdatePeeringConnectionApiRequest
	/*
		UpdatePeeringConnection Update One Network Peering Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdatePeeringConnectionApiParams - Parameters for the request
		@return UpdatePeeringConnectionApiRequest
	*/
	UpdatePeeringConnectionWithParams(ctx context.Context, args *UpdatePeeringConnectionApiParams) UpdatePeeringConnectionApiRequest

	// Method available only for mocking purposes
	UpdatePeeringConnectionExecute(r UpdatePeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error)

	/*
		UpdatePeeringContainer Update One Network Peering Container

		Updates the network details and labels of one specified network peering container in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
		@param cloudProviderContainer Updates the network details and labels of one specified network peering container in the specified project.
		@return UpdatePeeringContainerApiRequest
	*/
	UpdatePeeringContainer(ctx context.Context, groupId string, containerId string, cloudProviderContainer *CloudProviderContainer) UpdatePeeringContainerApiRequest
	/*
		UpdatePeeringContainer Update One Network Peering Container


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdatePeeringContainerApiParams - Parameters for the request
		@return UpdatePeeringContainerApiRequest
	*/
	UpdatePeeringContainerWithParams(ctx context.Context, args *UpdatePeeringContainerApiParams) UpdatePeeringContainerApiRequest

	// Method available only for mocking purposes
	UpdatePeeringContainerExecute(r UpdatePeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error)

	/*
		VerifyConnectViaPeeringOnlyModeForOneProject Verify Connect via Peering-Only Mode for One Project

		Verifies if someone set the specified project to **Connect via Peering Only** mode. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for NetworkPeeringApi
	*/
	VerifyConnectViaPeeringOnlyModeForOneProject(ctx context.Context, groupId string) VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest
	/*
		VerifyConnectViaPeeringOnlyModeForOneProject Verify Connect via Peering-Only Mode for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param VerifyConnectViaPeeringOnlyModeForOneProjectApiParams - Parameters for the request
		@return VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for NetworkPeeringApi
	*/
	VerifyConnectViaPeeringOnlyModeForOneProjectWithParams(ctx context.Context, args *VerifyConnectViaPeeringOnlyModeForOneProjectApiParams) VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest

	// Method available only for mocking purposes
	VerifyConnectViaPeeringOnlyModeForOneProjectExecute(r VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest) (*PrivateIPMode, *http.Response, error)
}

// NetworkPeeringApiService NetworkPeeringApi service
type NetworkPeeringApiService service

type CreatePeeringConnectionApiRequest struct {
	ctx                                  context.Context
	ApiService                           NetworkPeeringApi
	groupId                              string
	baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings
}

type CreatePeeringConnectionApiParams struct {
	GroupId                              string
	BaseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings
}

func (a *NetworkPeeringApiService) CreatePeeringConnectionWithParams(ctx context.Context, args *CreatePeeringConnectionApiParams) CreatePeeringConnectionApiRequest {
	return CreatePeeringConnectionApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		baseNetworkPeeringConnectionSettings: args.BaseNetworkPeeringConnectionSettings,
	}
}

func (r CreatePeeringConnectionApiRequest) Execute() (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	return r.ApiService.CreatePeeringConnectionExecute(r)
}

/*
CreatePeeringConnection Create One Network Peering Connection

Creates one new network peering connection in the specified project. Network peering allows multiple cloud-hosted applications to securely connect to the same project. To use this resource, the requesting Service Account or API Key must have the Project Owner role. To learn more about considerations and prerequisites, see the Network Peering Documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePeeringConnectionApiRequest
*/
func (a *NetworkPeeringApiService) CreatePeeringConnection(ctx context.Context, groupId string, baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings) CreatePeeringConnectionApiRequest {
	return CreatePeeringConnectionApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		baseNetworkPeeringConnectionSettings: baseNetworkPeeringConnectionSettings,
	}
}

// CreatePeeringConnectionExecute executes the request
//
//	@return BaseNetworkPeeringConnectionSettings
func (a *NetworkPeeringApiService) CreatePeeringConnectionExecute(r CreatePeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BaseNetworkPeeringConnectionSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.CreatePeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/peers"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.baseNetworkPeeringConnectionSettings == nil {
		return localVarReturnValue, nil, reportError("baseNetworkPeeringConnectionSettings is required and must be specified")
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
	localVarPostBody = r.baseNetworkPeeringConnectionSettings
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

type CreatePeeringContainerApiRequest struct {
	ctx                    context.Context
	ApiService             NetworkPeeringApi
	groupId                string
	cloudProviderContainer *CloudProviderContainer
}

type CreatePeeringContainerApiParams struct {
	GroupId                string
	CloudProviderContainer *CloudProviderContainer
}

func (a *NetworkPeeringApiService) CreatePeeringContainerWithParams(ctx context.Context, args *CreatePeeringContainerApiParams) CreatePeeringContainerApiRequest {
	return CreatePeeringContainerApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		cloudProviderContainer: args.CloudProviderContainer,
	}
}

func (r CreatePeeringContainerApiRequest) Execute() (*CloudProviderContainer, *http.Response, error) {
	return r.ApiService.CreatePeeringContainerExecute(r)
}

/*
CreatePeeringContainer Create One Network Peering Container

Creates one new network peering container in the specified project. MongoDB Cloud can deploy Network Peering connections in a network peering container. GCP can have one container per project. AWS and Azure can have one container per cloud provider region. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePeeringContainerApiRequest
*/
func (a *NetworkPeeringApiService) CreatePeeringContainer(ctx context.Context, groupId string, cloudProviderContainer *CloudProviderContainer) CreatePeeringContainerApiRequest {
	return CreatePeeringContainerApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		cloudProviderContainer: cloudProviderContainer,
	}
}

// CreatePeeringContainerExecute executes the request
//
//	@return CloudProviderContainer
func (a *NetworkPeeringApiService) CreatePeeringContainerExecute(r CreatePeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudProviderContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.CreatePeeringContainer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProviderContainer == nil {
		return localVarReturnValue, nil, reportError("cloudProviderContainer is required and must be specified")
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
	localVarPostBody = r.cloudProviderContainer
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

type DeletePeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService NetworkPeeringApi
	groupId    string
	peerId     string
}

type DeletePeeringConnectionApiParams struct {
	GroupId string
	PeerId  string
}

func (a *NetworkPeeringApiService) DeletePeeringConnectionWithParams(ctx context.Context, args *DeletePeeringConnectionApiParams) DeletePeeringConnectionApiRequest {
	return DeletePeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		peerId:     args.PeerId,
	}
}

func (r DeletePeeringConnectionApiRequest) Execute() (any, *http.Response, error) {
	return r.ApiService.DeletePeeringConnectionExecute(r)
}

/*
DeletePeeringConnection Remove One Network Peering Connection

Removes one network peering connection in the specified project. If you Removes the last network peering connection associated with a project, MongoDB Cloud also removes any AWS security groups from the project IP access list. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.
	@return DeletePeeringConnectionApiRequest
*/
func (a *NetworkPeeringApiService) DeletePeeringConnection(ctx context.Context, groupId string, peerId string) DeletePeeringConnectionApiRequest {
	return DeletePeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		peerId:     peerId,
	}
}

// DeletePeeringConnectionExecute executes the request
//
//	@return any
func (a *NetworkPeeringApiService) DeletePeeringConnectionExecute(r DeletePeeringConnectionApiRequest) (any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.DeletePeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/peers/{peerId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.peerId == "" {
		return localVarReturnValue, nil, reportError("peerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"peerId"+"}", url.PathEscape(r.peerId), -1)

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

type DeletePeeringContainerApiRequest struct {
	ctx         context.Context
	ApiService  NetworkPeeringApi
	groupId     string
	containerId string
}

type DeletePeeringContainerApiParams struct {
	GroupId     string
	ContainerId string
}

func (a *NetworkPeeringApiService) DeletePeeringContainerWithParams(ctx context.Context, args *DeletePeeringContainerApiParams) DeletePeeringContainerApiRequest {
	return DeletePeeringContainerApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		containerId: args.ContainerId,
	}
}

func (r DeletePeeringContainerApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePeeringContainerExecute(r)
}

/*
DeletePeeringContainer Remove One Network Peering Container

Removes one network peering container in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
	@return DeletePeeringContainerApiRequest
*/
func (a *NetworkPeeringApiService) DeletePeeringContainer(ctx context.Context, groupId string, containerId string) DeletePeeringContainerApiRequest {
	return DeletePeeringContainerApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		containerId: containerId,
	}
}

// DeletePeeringContainerExecute executes the request
func (a *NetworkPeeringApiService) DeletePeeringContainerExecute(r DeletePeeringContainerApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.DeletePeeringContainer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers/{containerId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.containerId == "" {
		return nil, reportError("containerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(r.containerId), -1)

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

type DisablePeeringApiRequest struct {
	ctx           context.Context
	ApiService    NetworkPeeringApi
	groupId       string
	privateIPMode *PrivateIPMode
}

type DisablePeeringApiParams struct {
	GroupId       string
	PrivateIPMode *PrivateIPMode
}

func (a *NetworkPeeringApiService) DisablePeeringWithParams(ctx context.Context, args *DisablePeeringApiParams) DisablePeeringApiRequest {
	return DisablePeeringApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		privateIPMode: args.PrivateIPMode,
	}
}

func (r DisablePeeringApiRequest) Execute() (*PrivateIPMode, *http.Response, error) {
	return r.ApiService.DisablePeeringExecute(r)
}

/*
DisablePeering Disable Connect via Peering-Only Mode for One Project

Disables Connect via Peering Only mode for the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DisablePeeringApiRequest

Deprecated
*/
func (a *NetworkPeeringApiService) DisablePeering(ctx context.Context, groupId string, privateIPMode *PrivateIPMode) DisablePeeringApiRequest {
	return DisablePeeringApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		privateIPMode: privateIPMode,
	}
}

// DisablePeeringExecute executes the request
//
//	@return PrivateIPMode
//
// Deprecated
func (a *NetworkPeeringApiService) DisablePeeringExecute(r DisablePeeringApiRequest) (*PrivateIPMode, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PrivateIPMode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.DisablePeering")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateIpMode"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.privateIPMode == nil {
		return localVarReturnValue, nil, reportError("privateIPMode is required and must be specified")
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
	localVarPostBody = r.privateIPMode
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

type GetPeeringConnectionApiRequest struct {
	ctx        context.Context
	ApiService NetworkPeeringApi
	groupId    string
	peerId     string
}

type GetPeeringConnectionApiParams struct {
	GroupId string
	PeerId  string
}

func (a *NetworkPeeringApiService) GetPeeringConnectionWithParams(ctx context.Context, args *GetPeeringConnectionApiParams) GetPeeringConnectionApiRequest {
	return GetPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		peerId:     args.PeerId,
	}
}

func (r GetPeeringConnectionApiRequest) Execute() (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	return r.ApiService.GetPeeringConnectionExecute(r)
}

/*
GetPeeringConnection Return One Network Peering Connection in One Project

Returns details about one specified network peering connection in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.
	@return GetPeeringConnectionApiRequest
*/
func (a *NetworkPeeringApiService) GetPeeringConnection(ctx context.Context, groupId string, peerId string) GetPeeringConnectionApiRequest {
	return GetPeeringConnectionApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		peerId:     peerId,
	}
}

// GetPeeringConnectionExecute executes the request
//
//	@return BaseNetworkPeeringConnectionSettings
func (a *NetworkPeeringApiService) GetPeeringConnectionExecute(r GetPeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BaseNetworkPeeringConnectionSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.GetPeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/peers/{peerId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.peerId == "" {
		return localVarReturnValue, nil, reportError("peerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"peerId"+"}", url.PathEscape(r.peerId), -1)

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

type GetPeeringContainerApiRequest struct {
	ctx         context.Context
	ApiService  NetworkPeeringApi
	groupId     string
	containerId string
}

type GetPeeringContainerApiParams struct {
	GroupId     string
	ContainerId string
}

func (a *NetworkPeeringApiService) GetPeeringContainerWithParams(ctx context.Context, args *GetPeeringContainerApiParams) GetPeeringContainerApiRequest {
	return GetPeeringContainerApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		containerId: args.ContainerId,
	}
}

func (r GetPeeringContainerApiRequest) Execute() (*CloudProviderContainer, *http.Response, error) {
	return r.ApiService.GetPeeringContainerExecute(r)
}

/*
GetPeeringContainer Return One Network Peering Container

Returns details about one network peering container in one specified project. Network peering containers contain network peering connections. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container.
	@return GetPeeringContainerApiRequest
*/
func (a *NetworkPeeringApiService) GetPeeringContainer(ctx context.Context, groupId string, containerId string) GetPeeringContainerApiRequest {
	return GetPeeringContainerApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		containerId: containerId,
	}
}

// GetPeeringContainerExecute executes the request
//
//	@return CloudProviderContainer
func (a *NetworkPeeringApiService) GetPeeringContainerExecute(r GetPeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudProviderContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.GetPeeringContainer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers/{containerId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.containerId == "" {
		return localVarReturnValue, nil, reportError("containerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(r.containerId), -1)

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

type ListPeeringConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   NetworkPeeringApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	providerName *string
}

type ListPeeringConnectionsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	ProviderName *string
}

func (a *NetworkPeeringApiService) ListPeeringConnectionsWithParams(ctx context.Context, args *ListPeeringConnectionsApiParams) ListPeeringConnectionsApiRequest {
	return ListPeeringConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		providerName: args.ProviderName,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListPeeringConnectionsApiRequest) IncludeCount(includeCount bool) ListPeeringConnectionsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListPeeringConnectionsApiRequest) ItemsPerPage(itemsPerPage int) ListPeeringConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPeeringConnectionsApiRequest) PageNum(pageNum int) ListPeeringConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Cloud service provider to use for this VPC peering connection.
func (r ListPeeringConnectionsApiRequest) ProviderName(providerName string) ListPeeringConnectionsApiRequest {
	r.providerName = &providerName
	return r
}

func (r ListPeeringConnectionsApiRequest) Execute() (*PaginatedContainerPeer, *http.Response, error) {
	return r.ApiService.ListPeeringConnectionsExecute(r)
}

/*
ListPeeringConnections Return All Network Peering Connections in One Project

Returns details about all network peering connections in the specified project. Network peering allows multiple cloud-hosted applications to securely connect to the same project. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListPeeringConnectionsApiRequest
*/
func (a *NetworkPeeringApiService) ListPeeringConnections(ctx context.Context, groupId string) ListPeeringConnectionsApiRequest {
	return ListPeeringConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListPeeringConnectionsExecute executes the request
//
//	@return PaginatedContainerPeer
func (a *NetworkPeeringApiService) ListPeeringConnectionsExecute(r ListPeeringConnectionsApiRequest) (*PaginatedContainerPeer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedContainerPeer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.ListPeeringConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/peers"
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
	if r.providerName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "providerName", r.providerName, "")
	} else {
		var defaultValue string = "AWS"
		r.providerName = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "providerName", r.providerName, "")
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

type ListPeeringContainerByCloudProviderApiRequest struct {
	ctx          context.Context
	ApiService   NetworkPeeringApi
	groupId      string
	providerName *string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListPeeringContainerByCloudProviderApiParams struct {
	GroupId      string
	ProviderName *string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *NetworkPeeringApiService) ListPeeringContainerByCloudProviderWithParams(ctx context.Context, args *ListPeeringContainerByCloudProviderApiParams) ListPeeringContainerByCloudProviderApiRequest {
	return ListPeeringContainerByCloudProviderApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		providerName: args.ProviderName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Cloud service provider that serves the desired network peering containers.
func (r ListPeeringContainerByCloudProviderApiRequest) ProviderName(providerName string) ListPeeringContainerByCloudProviderApiRequest {
	r.providerName = &providerName
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListPeeringContainerByCloudProviderApiRequest) IncludeCount(includeCount bool) ListPeeringContainerByCloudProviderApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListPeeringContainerByCloudProviderApiRequest) ItemsPerPage(itemsPerPage int) ListPeeringContainerByCloudProviderApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPeeringContainerByCloudProviderApiRequest) PageNum(pageNum int) ListPeeringContainerByCloudProviderApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListPeeringContainerByCloudProviderApiRequest) Execute() (*PaginatedCloudProviderContainer, *http.Response, error) {
	return r.ApiService.ListPeeringContainerByCloudProviderExecute(r)
}

/*
ListPeeringContainerByCloudProvider Return All Network Peering Containers in One Project for One Cloud Provider

Returns details about all network peering containers in the specified project for the specified cloud provider. If you do not specify the cloud provider, MongoDB Cloud returns details about all network peering containers in the project for Amazon Web Services (AWS). To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListPeeringContainerByCloudProviderApiRequest
*/
func (a *NetworkPeeringApiService) ListPeeringContainerByCloudProvider(ctx context.Context, groupId string) ListPeeringContainerByCloudProviderApiRequest {
	return ListPeeringContainerByCloudProviderApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListPeeringContainerByCloudProviderExecute executes the request
//
//	@return PaginatedCloudProviderContainer
func (a *NetworkPeeringApiService) ListPeeringContainerByCloudProviderExecute(r ListPeeringContainerByCloudProviderApiRequest) (*PaginatedCloudProviderContainer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudProviderContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.ListPeeringContainerByCloudProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.providerName == nil {
		return localVarReturnValue, nil, reportError("providerName is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarQueryParams, "providerName", r.providerName, "")
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

type ListPeeringContainersApiRequest struct {
	ctx          context.Context
	ApiService   NetworkPeeringApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListPeeringContainersApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *NetworkPeeringApiService) ListPeeringContainersWithParams(ctx context.Context, args *ListPeeringContainersApiParams) ListPeeringContainersApiRequest {
	return ListPeeringContainersApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListPeeringContainersApiRequest) IncludeCount(includeCount bool) ListPeeringContainersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListPeeringContainersApiRequest) ItemsPerPage(itemsPerPage int) ListPeeringContainersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPeeringContainersApiRequest) PageNum(pageNum int) ListPeeringContainersApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListPeeringContainersApiRequest) Execute() (*PaginatedCloudProviderContainer, *http.Response, error) {
	return r.ApiService.ListPeeringContainersExecute(r)
}

/*
ListPeeringContainers Return All Network Peering Containers in One Project

Returns details about all network peering containers in the specified project. Network peering containers contain network peering connections. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListPeeringContainersApiRequest
*/
func (a *NetworkPeeringApiService) ListPeeringContainers(ctx context.Context, groupId string) ListPeeringContainersApiRequest {
	return ListPeeringContainersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListPeeringContainersExecute executes the request
//
//	@return PaginatedCloudProviderContainer
func (a *NetworkPeeringApiService) ListPeeringContainersExecute(r ListPeeringContainersApiRequest) (*PaginatedCloudProviderContainer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedCloudProviderContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.ListPeeringContainers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers/all"
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

type UpdatePeeringConnectionApiRequest struct {
	ctx                                  context.Context
	ApiService                           NetworkPeeringApi
	groupId                              string
	peerId                               string
	baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings
}

type UpdatePeeringConnectionApiParams struct {
	GroupId                              string
	PeerId                               string
	BaseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings
}

func (a *NetworkPeeringApiService) UpdatePeeringConnectionWithParams(ctx context.Context, args *UpdatePeeringConnectionApiParams) UpdatePeeringConnectionApiRequest {
	return UpdatePeeringConnectionApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              args.GroupId,
		peerId:                               args.PeerId,
		baseNetworkPeeringConnectionSettings: args.BaseNetworkPeeringConnectionSettings,
	}
}

func (r UpdatePeeringConnectionApiRequest) Execute() (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	return r.ApiService.UpdatePeeringConnectionExecute(r)
}

/*
UpdatePeeringConnection Update One Network Peering Connection

Updates one specified network peering connection in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param peerId Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.
	@return UpdatePeeringConnectionApiRequest
*/
func (a *NetworkPeeringApiService) UpdatePeeringConnection(ctx context.Context, groupId string, peerId string, baseNetworkPeeringConnectionSettings *BaseNetworkPeeringConnectionSettings) UpdatePeeringConnectionApiRequest {
	return UpdatePeeringConnectionApiRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		groupId:                              groupId,
		peerId:                               peerId,
		baseNetworkPeeringConnectionSettings: baseNetworkPeeringConnectionSettings,
	}
}

// UpdatePeeringConnectionExecute executes the request
//
//	@return BaseNetworkPeeringConnectionSettings
func (a *NetworkPeeringApiService) UpdatePeeringConnectionExecute(r UpdatePeeringConnectionApiRequest) (*BaseNetworkPeeringConnectionSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BaseNetworkPeeringConnectionSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.UpdatePeeringConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/peers/{peerId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.peerId == "" {
		return localVarReturnValue, nil, reportError("peerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"peerId"+"}", url.PathEscape(r.peerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.baseNetworkPeeringConnectionSettings == nil {
		return localVarReturnValue, nil, reportError("baseNetworkPeeringConnectionSettings is required and must be specified")
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
	localVarPostBody = r.baseNetworkPeeringConnectionSettings
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

type UpdatePeeringContainerApiRequest struct {
	ctx                    context.Context
	ApiService             NetworkPeeringApi
	groupId                string
	containerId            string
	cloudProviderContainer *CloudProviderContainer
}

type UpdatePeeringContainerApiParams struct {
	GroupId                string
	ContainerId            string
	CloudProviderContainer *CloudProviderContainer
}

func (a *NetworkPeeringApiService) UpdatePeeringContainerWithParams(ctx context.Context, args *UpdatePeeringContainerApiParams) UpdatePeeringContainerApiRequest {
	return UpdatePeeringContainerApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		containerId:            args.ContainerId,
		cloudProviderContainer: args.CloudProviderContainer,
	}
}

func (r UpdatePeeringContainerApiRequest) Execute() (*CloudProviderContainer, *http.Response, error) {
	return r.ApiService.UpdatePeeringContainerExecute(r)
}

/*
UpdatePeeringContainer Update One Network Peering Container

Updates the network details and labels of one specified network peering container in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param containerId Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
	@return UpdatePeeringContainerApiRequest
*/
func (a *NetworkPeeringApiService) UpdatePeeringContainer(ctx context.Context, groupId string, containerId string, cloudProviderContainer *CloudProviderContainer) UpdatePeeringContainerApiRequest {
	return UpdatePeeringContainerApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		containerId:            containerId,
		cloudProviderContainer: cloudProviderContainer,
	}
}

// UpdatePeeringContainerExecute executes the request
//
//	@return CloudProviderContainer
func (a *NetworkPeeringApiService) UpdatePeeringContainerExecute(r UpdatePeeringContainerApiRequest) (*CloudProviderContainer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CloudProviderContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.UpdatePeeringContainer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/containers/{containerId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.containerId == "" {
		return localVarReturnValue, nil, reportError("containerId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(r.containerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProviderContainer == nil {
		return localVarReturnValue, nil, reportError("cloudProviderContainer is required and must be specified")
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
	localVarPostBody = r.cloudProviderContainer
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

type VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest struct {
	ctx        context.Context
	ApiService NetworkPeeringApi
	groupId    string
}

type VerifyConnectViaPeeringOnlyModeForOneProjectApiParams struct {
	GroupId string
}

func (a *NetworkPeeringApiService) VerifyConnectViaPeeringOnlyModeForOneProjectWithParams(ctx context.Context, args *VerifyConnectViaPeeringOnlyModeForOneProjectApiParams) VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest {
	return VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest) Execute() (*PrivateIPMode, *http.Response, error) {
	return r.ApiService.VerifyConnectViaPeeringOnlyModeForOneProjectExecute(r)
}

/*
VerifyConnectViaPeeringOnlyModeForOneProject Verify Connect via Peering-Only Mode for One Project

Verifies if someone set the specified project to **Connect via Peering Only** mode. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest

Deprecated
*/
func (a *NetworkPeeringApiService) VerifyConnectViaPeeringOnlyModeForOneProject(ctx context.Context, groupId string) VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest {
	return VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// VerifyConnectViaPeeringOnlyModeForOneProjectExecute executes the request
//
//	@return PrivateIPMode
//
// Deprecated
func (a *NetworkPeeringApiService) VerifyConnectViaPeeringOnlyModeForOneProjectExecute(r VerifyConnectViaPeeringOnlyModeForOneProjectApiRequest) (*PrivateIPMode, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PrivateIPMode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkPeeringApiService.VerifyConnectViaPeeringOnlyModeForOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateIpMode"
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
