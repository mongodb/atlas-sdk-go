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

type FederatedAuthenticationApi interface {

	/*
		CreateRoleMapping Add One Role Mapping to One Organization

		[experimental] Adds one role mapping to the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return CreateRoleMappingApiRequest
	*/
	CreateRoleMapping(ctx context.Context, federationSettingsId string, orgId string, authFederationRoleMapping *AuthFederationRoleMapping) CreateRoleMappingApiRequest
	/*
		CreateRoleMapping Add One Role Mapping to One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateRoleMappingApiParams - Parameters for the request
		@return CreateRoleMappingApiRequest
	*/
	CreateRoleMappingWithParams(ctx context.Context, args *CreateRoleMappingApiParams) CreateRoleMappingApiRequest

	// Interface only available internally
	createRoleMappingExecute(r CreateRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error)

	/*
		DeleteFederationApp Delete the federation settings instance.

		[experimental] Deletes the federation settings instance and all associated data, including identity providers and domains. To use this resource, the requesting API Key must have the Organization Owner role in the last remaining connected organization. **Note**: requests to this resource will fail if there is more than one connected organization in the federation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@return DeleteFederationAppApiRequest
	*/
	DeleteFederationApp(ctx context.Context, federationSettingsId string) DeleteFederationAppApiRequest
	/*
		DeleteFederationApp Delete the federation settings instance.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteFederationAppApiParams - Parameters for the request
		@return DeleteFederationAppApiRequest
	*/
	DeleteFederationAppWithParams(ctx context.Context, args *DeleteFederationAppApiParams) DeleteFederationAppApiRequest

	// Interface only available internally
	deleteFederationAppExecute(r DeleteFederationAppApiRequest) (*http.Response, error)

	/*
		DeleteRoleMapping Remove One Role Mapping from One Organization

		[experimental] Removes one role mapping in the specified organization from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return DeleteRoleMappingApiRequest
	*/
	DeleteRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string) DeleteRoleMappingApiRequest
	/*
		DeleteRoleMapping Remove One Role Mapping from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteRoleMappingApiParams - Parameters for the request
		@return DeleteRoleMappingApiRequest
	*/
	DeleteRoleMappingWithParams(ctx context.Context, args *DeleteRoleMappingApiParams) DeleteRoleMappingApiRequest

	// Interface only available internally
	deleteRoleMappingExecute(r DeleteRoleMappingApiRequest) (*http.Response, error)

	/*
		GetConnectedOrgConfig Return One Org Config Connected to One Federation

		[experimental] Returns the specified connected org config from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in the connected org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.
		@return GetConnectedOrgConfigApiRequest
	*/
	GetConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string) GetConnectedOrgConfigApiRequest
	/*
		GetConnectedOrgConfig Return One Org Config Connected to One Federation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetConnectedOrgConfigApiParams - Parameters for the request
		@return GetConnectedOrgConfigApiRequest
	*/
	GetConnectedOrgConfigWithParams(ctx context.Context, args *GetConnectedOrgConfigApiParams) GetConnectedOrgConfigApiRequest

	// Interface only available internally
	getConnectedOrgConfigExecute(r GetConnectedOrgConfigApiRequest) (*ConnectedOrgConfig, *http.Response, error)

	/*
		GetFederationSettings Return Federation Settings for One Organization

		[experimental] Returns information about the federation settings for the specified organization. To use this resource, the requesting API Key must have the Organization Owner role in the connected org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetFederationSettingsApiRequest
	*/
	GetFederationSettings(ctx context.Context, orgId string) GetFederationSettingsApiRequest
	/*
		GetFederationSettings Return Federation Settings for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetFederationSettingsApiParams - Parameters for the request
		@return GetFederationSettingsApiRequest
	*/
	GetFederationSettingsWithParams(ctx context.Context, args *GetFederationSettingsApiParams) GetFederationSettingsApiRequest

	// Interface only available internally
	getFederationSettingsExecute(r GetFederationSettingsApiRequest) (*OrgFederationSettings, *http.Response, error)

	/*
		GetIdentityProvider Return one identity provider from the specified federation by id.

		[experimental] Returns one identity provider in the specified federation by the identity provider's id. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
		@return GetIdentityProviderApiRequest
	*/
	GetIdentityProvider(ctx context.Context, federationSettingsId string, identityProviderId string) GetIdentityProviderApiRequest
	/*
		GetIdentityProvider Return one identity provider from the specified federation by id.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetIdentityProviderApiParams - Parameters for the request
		@return GetIdentityProviderApiRequest
	*/
	GetIdentityProviderWithParams(ctx context.Context, args *GetIdentityProviderApiParams) GetIdentityProviderApiRequest

	// Interface only available internally
	getIdentityProviderExecute(r GetIdentityProviderApiRequest) (*FederationIdentityProvider, *http.Response, error)

	/*
		GetIdentityProviderMetadata Return the metadata of one identity provider in the specified federation.

		[experimental] Returns the metadata of one identity provider in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
		@return GetIdentityProviderMetadataApiRequest
	*/
	GetIdentityProviderMetadata(ctx context.Context, federationSettingsId string, identityProviderId string) GetIdentityProviderMetadataApiRequest
	/*
		GetIdentityProviderMetadata Return the metadata of one identity provider in the specified federation.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetIdentityProviderMetadataApiParams - Parameters for the request
		@return GetIdentityProviderMetadataApiRequest
	*/
	GetIdentityProviderMetadataWithParams(ctx context.Context, args *GetIdentityProviderMetadataApiParams) GetIdentityProviderMetadataApiRequest

	// Interface only available internally
	getIdentityProviderMetadataExecute(r GetIdentityProviderMetadataApiRequest) (string, *http.Response, error)

	/*
		GetRoleMapping Return One Role Mapping from One Organization

		[experimental] Returns one role mapping from the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetRoleMappingApiRequest
	*/
	GetRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string) GetRoleMappingApiRequest
	/*
		GetRoleMapping Return One Role Mapping from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetRoleMappingApiParams - Parameters for the request
		@return GetRoleMappingApiRequest
	*/
	GetRoleMappingWithParams(ctx context.Context, args *GetRoleMappingApiParams) GetRoleMappingApiRequest

	// Interface only available internally
	getRoleMappingExecute(r GetRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error)

	/*
		ListConnectedOrgConfigs Return All Connected Org Configs from the Federation

		[experimental] Returns all connected org configs in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected orgs.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@return ListConnectedOrgConfigsApiRequest
	*/
	ListConnectedOrgConfigs(ctx context.Context, federationSettingsId string) ListConnectedOrgConfigsApiRequest
	/*
		ListConnectedOrgConfigs Return All Connected Org Configs from the Federation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListConnectedOrgConfigsApiParams - Parameters for the request
		@return ListConnectedOrgConfigsApiRequest
	*/
	ListConnectedOrgConfigsWithParams(ctx context.Context, args *ListConnectedOrgConfigsApiParams) ListConnectedOrgConfigsApiRequest

	// Interface only available internally
	listConnectedOrgConfigsExecute(r ListConnectedOrgConfigsApiRequest) ([]ConnectedOrgConfig, *http.Response, error)

	/*
		ListIdentityProviders Return all identity providers from the specified federation.

		[experimental] Returns all identity providers with the provided protocol in the specified federation. If no protocol is specified, only SAML identity providers will be returned. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@return ListIdentityProvidersApiRequest
	*/
	ListIdentityProviders(ctx context.Context, federationSettingsId string) ListIdentityProvidersApiRequest
	/*
		ListIdentityProviders Return all identity providers from the specified federation.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListIdentityProvidersApiParams - Parameters for the request
		@return ListIdentityProvidersApiRequest
	*/
	ListIdentityProvidersWithParams(ctx context.Context, args *ListIdentityProvidersApiParams) ListIdentityProvidersApiRequest

	// Interface only available internally
	listIdentityProvidersExecute(r ListIdentityProvidersApiRequest) (*PaginatedFederationIdentityProvider, *http.Response, error)

	/*
		ListRoleMappings Return All Role Mappings from One Organization

		[experimental] Returns all role mappings from the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListRoleMappingsApiRequest
	*/
	ListRoleMappings(ctx context.Context, federationSettingsId string, orgId string) ListRoleMappingsApiRequest
	/*
		ListRoleMappings Return All Role Mappings from One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListRoleMappingsApiParams - Parameters for the request
		@return ListRoleMappingsApiRequest
	*/
	ListRoleMappingsWithParams(ctx context.Context, args *ListRoleMappingsApiParams) ListRoleMappingsApiRequest

	// Interface only available internally
	listRoleMappingsExecute(r ListRoleMappingsApiRequest) (*PaginatedRoleMapping, *http.Response, error)

	/*
		RemoveConnectedOrgConfig Remove One Org Config Connected to One Federation

		[experimental] Removes one connected organization configuration from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role. Note: This request fails if only one connected organization exists in the federation.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.
		@return RemoveConnectedOrgConfigApiRequest
	*/
	RemoveConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string) RemoveConnectedOrgConfigApiRequest
	/*
		RemoveConnectedOrgConfig Remove One Org Config Connected to One Federation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveConnectedOrgConfigApiParams - Parameters for the request
		@return RemoveConnectedOrgConfigApiRequest
	*/
	RemoveConnectedOrgConfigWithParams(ctx context.Context, args *RemoveConnectedOrgConfigApiParams) RemoveConnectedOrgConfigApiRequest

	// Interface only available internally
	removeConnectedOrgConfigExecute(r RemoveConnectedOrgConfigApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		UpdateConnectedOrgConfig Update One Org Config Connected to One Federation

		[experimental] Updates one connected organization configuration from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	**Note** If the organization configuration has no associated identity provider, you can't use this resource to update role mappings or post authorization role grants.

	**Note**: The domainRestrictionEnabled field defaults to false if not provided in the request.

	**Note**: If the identityProviderId field is not provided, you will disconnect the organization and the identity provider.

	**Note**: Currently connected data access identity providers missing from the dataAccessIdentityProviderIds field will be disconnected.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.
		@return UpdateConnectedOrgConfigApiRequest
	*/
	UpdateConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string, connectedOrgConfig *ConnectedOrgConfig) UpdateConnectedOrgConfigApiRequest
	/*
		UpdateConnectedOrgConfig Update One Org Config Connected to One Federation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateConnectedOrgConfigApiParams - Parameters for the request
		@return UpdateConnectedOrgConfigApiRequest
	*/
	UpdateConnectedOrgConfigWithParams(ctx context.Context, args *UpdateConnectedOrgConfigApiParams) UpdateConnectedOrgConfigApiRequest

	// Interface only available internally
	updateConnectedOrgConfigExecute(r UpdateConnectedOrgConfigApiRequest) (*ConnectedOrgConfig, *http.Response, error)

	/*
		UpdateIdentityProvider Update the identity provider.

		[experimental] Updates one identity provider in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
		@return UpdateIdentityProviderApiRequest
	*/
	UpdateIdentityProvider(ctx context.Context, federationSettingsId string, identityProviderId string, federationIdentityProviderUpdate *FederationIdentityProviderUpdate) UpdateIdentityProviderApiRequest
	/*
		UpdateIdentityProvider Update the identity provider.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateIdentityProviderApiParams - Parameters for the request
		@return UpdateIdentityProviderApiRequest
	*/
	UpdateIdentityProviderWithParams(ctx context.Context, args *UpdateIdentityProviderApiParams) UpdateIdentityProviderApiRequest

	// Interface only available internally
	updateIdentityProviderExecute(r UpdateIdentityProviderApiRequest) (*FederationIdentityProvider, *http.Response, error)

	/*
		UpdateRoleMapping Update One Role Mapping in One Organization

		[experimental] Updates one role mapping in the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
		@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return UpdateRoleMappingApiRequest
	*/
	UpdateRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string, authFederationRoleMapping *AuthFederationRoleMapping) UpdateRoleMappingApiRequest
	/*
		UpdateRoleMapping Update One Role Mapping in One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateRoleMappingApiParams - Parameters for the request
		@return UpdateRoleMappingApiRequest
	*/
	UpdateRoleMappingWithParams(ctx context.Context, args *UpdateRoleMappingApiParams) UpdateRoleMappingApiRequest

	// Interface only available internally
	updateRoleMappingExecute(r UpdateRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error)
}

// FederatedAuthenticationApiService FederatedAuthenticationApi service
type FederatedAuthenticationApiService service

type CreateRoleMappingApiRequest struct {
	ctx                       context.Context
	ApiService                FederatedAuthenticationApi
	federationSettingsId      string
	orgId                     string
	authFederationRoleMapping *AuthFederationRoleMapping
}

type CreateRoleMappingApiParams struct {
	FederationSettingsId      string
	OrgId                     string
	AuthFederationRoleMapping *AuthFederationRoleMapping
}

func (a *FederatedAuthenticationApiService) CreateRoleMappingWithParams(ctx context.Context, args *CreateRoleMappingApiParams) CreateRoleMappingApiRequest {
	return CreateRoleMappingApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		federationSettingsId:      args.FederationSettingsId,
		orgId:                     args.OrgId,
		authFederationRoleMapping: args.AuthFederationRoleMapping,
	}
}

func (r CreateRoleMappingApiRequest) Execute() (*AuthFederationRoleMapping, *http.Response, error) {
	return r.ApiService.createRoleMappingExecute(r)
}

/*
CreateRoleMapping Add One Role Mapping to One Organization

[experimental] Adds one role mapping to the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateRoleMappingApiRequest
*/
func (a *FederatedAuthenticationApiService) CreateRoleMapping(ctx context.Context, federationSettingsId string, orgId string, authFederationRoleMapping *AuthFederationRoleMapping) CreateRoleMappingApiRequest {
	return CreateRoleMappingApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		federationSettingsId:      federationSettingsId,
		orgId:                     orgId,
		authFederationRoleMapping: authFederationRoleMapping,
	}
}

// Execute executes the request
//
//	@return AuthFederationRoleMapping
func (a *FederatedAuthenticationApiService) createRoleMappingExecute(r CreateRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AuthFederationRoleMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.CreateRoleMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authFederationRoleMapping == nil {
		return localVarReturnValue, nil, reportError("authFederationRoleMapping is required and must be specified")
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
	localVarPostBody = r.authFederationRoleMapping
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
		var v ApiError
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

type DeleteFederationAppApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
}

type DeleteFederationAppApiParams struct {
	FederationSettingsId string
}

func (a *FederatedAuthenticationApiService) DeleteFederationAppWithParams(ctx context.Context, args *DeleteFederationAppApiParams) DeleteFederationAppApiRequest {
	return DeleteFederationAppApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
	}
}

func (r DeleteFederationAppApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.deleteFederationAppExecute(r)
}

/*
DeleteFederationApp Delete the federation settings instance.

[experimental] Deletes the federation settings instance and all associated data, including identity providers and domains. To use this resource, the requesting API Key must have the Organization Owner role in the last remaining connected organization. **Note**: requests to this resource will fail if there is more than one connected organization in the federation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@return DeleteFederationAppApiRequest
*/
func (a *FederatedAuthenticationApiService) DeleteFederationApp(ctx context.Context, federationSettingsId string) DeleteFederationAppApiRequest {
	return DeleteFederationAppApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
	}
}

// Execute executes the request
func (a *FederatedAuthenticationApiService) deleteFederationAppExecute(r DeleteFederationAppApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.DeleteFederationApp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DeleteRoleMappingApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	id                   string
	orgId                string
}

type DeleteRoleMappingApiParams struct {
	FederationSettingsId string
	Id                   string
	OrgId                string
}

func (a *FederatedAuthenticationApiService) DeleteRoleMappingWithParams(ctx context.Context, args *DeleteRoleMappingApiParams) DeleteRoleMappingApiRequest {
	return DeleteRoleMappingApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		id:                   args.Id,
		orgId:                args.OrgId,
	}
}

func (r DeleteRoleMappingApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.deleteRoleMappingExecute(r)
}

/*
DeleteRoleMapping Remove One Role Mapping from One Organization

[experimental] Removes one role mapping in the specified organization from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return DeleteRoleMappingApiRequest
*/
func (a *FederatedAuthenticationApiService) DeleteRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string) DeleteRoleMappingApiRequest {
	return DeleteRoleMappingApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		id:                   id,
		orgId:                orgId,
	}
}

// Execute executes the request
func (a *FederatedAuthenticationApiService) deleteRoleMappingExecute(r DeleteRoleMappingApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.DeleteRoleMapping")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetConnectedOrgConfigApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	orgId                string
}

type GetConnectedOrgConfigApiParams struct {
	FederationSettingsId string
	OrgId                string
}

func (a *FederatedAuthenticationApiService) GetConnectedOrgConfigWithParams(ctx context.Context, args *GetConnectedOrgConfigApiParams) GetConnectedOrgConfigApiRequest {
	return GetConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		orgId:                args.OrgId,
	}
}

func (r GetConnectedOrgConfigApiRequest) Execute() (*ConnectedOrgConfig, *http.Response, error) {
	return r.ApiService.getConnectedOrgConfigExecute(r)
}

/*
GetConnectedOrgConfig Return One Org Config Connected to One Federation

[experimental] Returns the specified connected org config from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in the connected org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.
	@return GetConnectedOrgConfigApiRequest
*/
func (a *FederatedAuthenticationApiService) GetConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string) GetConnectedOrgConfigApiRequest {
	return GetConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		orgId:                orgId,
	}
}

// Execute executes the request
//
//	@return ConnectedOrgConfig
func (a *FederatedAuthenticationApiService) getConnectedOrgConfigExecute(r GetConnectedOrgConfigApiRequest) (*ConnectedOrgConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConnectedOrgConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.GetConnectedOrgConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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
		var v ApiError
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

type GetFederationSettingsApiRequest struct {
	ctx        context.Context
	ApiService FederatedAuthenticationApi
	orgId      string
}

type GetFederationSettingsApiParams struct {
	OrgId string
}

func (a *FederatedAuthenticationApiService) GetFederationSettingsWithParams(ctx context.Context, args *GetFederationSettingsApiParams) GetFederationSettingsApiRequest {
	return GetFederationSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetFederationSettingsApiRequest) Execute() (*OrgFederationSettings, *http.Response, error) {
	return r.ApiService.getFederationSettingsExecute(r)
}

/*
GetFederationSettings Return Federation Settings for One Organization

[experimental] Returns information about the federation settings for the specified organization. To use this resource, the requesting API Key must have the Organization Owner role in the connected org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetFederationSettingsApiRequest
*/
func (a *FederatedAuthenticationApiService) GetFederationSettings(ctx context.Context, orgId string) GetFederationSettingsApiRequest {
	return GetFederationSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return OrgFederationSettings
func (a *FederatedAuthenticationApiService) getFederationSettingsExecute(r GetFederationSettingsApiRequest) (*OrgFederationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrgFederationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.GetFederationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/federationSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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
		var v ApiError
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

type GetIdentityProviderApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	identityProviderId   string
}

type GetIdentityProviderApiParams struct {
	FederationSettingsId string
	IdentityProviderId   string
}

func (a *FederatedAuthenticationApiService) GetIdentityProviderWithParams(ctx context.Context, args *GetIdentityProviderApiParams) GetIdentityProviderApiRequest {
	return GetIdentityProviderApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		identityProviderId:   args.IdentityProviderId,
	}
}

func (r GetIdentityProviderApiRequest) Execute() (*FederationIdentityProvider, *http.Response, error) {
	return r.ApiService.getIdentityProviderExecute(r)
}

/*
GetIdentityProvider Return one identity provider from the specified federation by id.

[experimental] Returns one identity provider in the specified federation by the identity provider's id. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
	@return GetIdentityProviderApiRequest
*/
func (a *FederatedAuthenticationApiService) GetIdentityProvider(ctx context.Context, federationSettingsId string, identityProviderId string) GetIdentityProviderApiRequest {
	return GetIdentityProviderApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		identityProviderId:   identityProviderId,
	}
}

// Execute executes the request
//
//	@return FederationIdentityProvider
func (a *FederatedAuthenticationApiService) getIdentityProviderExecute(r GetIdentityProviderApiRequest) (*FederationIdentityProvider, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FederationIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.GetIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderId"+"}", url.PathEscape(parameterValueToString(r.identityProviderId, "identityProviderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

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
		var v ApiError
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

type GetIdentityProviderMetadataApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	identityProviderId   string
}

type GetIdentityProviderMetadataApiParams struct {
	FederationSettingsId string
	IdentityProviderId   string
}

func (a *FederatedAuthenticationApiService) GetIdentityProviderMetadataWithParams(ctx context.Context, args *GetIdentityProviderMetadataApiParams) GetIdentityProviderMetadataApiRequest {
	return GetIdentityProviderMetadataApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		identityProviderId:   args.IdentityProviderId,
	}
}

func (r GetIdentityProviderMetadataApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.getIdentityProviderMetadataExecute(r)
}

/*
GetIdentityProviderMetadata Return the metadata of one identity provider in the specified federation.

[experimental] Returns the metadata of one identity provider in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
	@return GetIdentityProviderMetadataApiRequest
*/
func (a *FederatedAuthenticationApiService) GetIdentityProviderMetadata(ctx context.Context, federationSettingsId string, identityProviderId string) GetIdentityProviderMetadataApiRequest {
	return GetIdentityProviderMetadataApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		identityProviderId:   identityProviderId,
	}
}

// Execute executes the request
//
//	@return string
func (a *FederatedAuthenticationApiService) getIdentityProviderMetadataExecute(r GetIdentityProviderMetadataApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.GetIdentityProviderMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderId"+"}", url.PathEscape(parameterValueToString(r.identityProviderId, "identityProviderId")), -1)

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
		var v ApiError
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

type GetRoleMappingApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	id                   string
	orgId                string
}

type GetRoleMappingApiParams struct {
	FederationSettingsId string
	Id                   string
	OrgId                string
}

func (a *FederatedAuthenticationApiService) GetRoleMappingWithParams(ctx context.Context, args *GetRoleMappingApiParams) GetRoleMappingApiRequest {
	return GetRoleMappingApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		id:                   args.Id,
		orgId:                args.OrgId,
	}
}

func (r GetRoleMappingApiRequest) Execute() (*AuthFederationRoleMapping, *http.Response, error) {
	return r.ApiService.getRoleMappingExecute(r)
}

/*
GetRoleMapping Return One Role Mapping from One Organization

[experimental] Returns one role mapping from the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetRoleMappingApiRequest
*/
func (a *FederatedAuthenticationApiService) GetRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string) GetRoleMappingApiRequest {
	return GetRoleMappingApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		id:                   id,
		orgId:                orgId,
	}
}

// Execute executes the request
//
//	@return AuthFederationRoleMapping
func (a *FederatedAuthenticationApiService) getRoleMappingExecute(r GetRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AuthFederationRoleMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.GetRoleMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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
		var v ApiError
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

type ListConnectedOrgConfigsApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
}

type ListConnectedOrgConfigsApiParams struct {
	FederationSettingsId string
}

func (a *FederatedAuthenticationApiService) ListConnectedOrgConfigsWithParams(ctx context.Context, args *ListConnectedOrgConfigsApiParams) ListConnectedOrgConfigsApiRequest {
	return ListConnectedOrgConfigsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
	}
}

func (r ListConnectedOrgConfigsApiRequest) Execute() ([]ConnectedOrgConfig, *http.Response, error) {
	return r.ApiService.listConnectedOrgConfigsExecute(r)
}

/*
ListConnectedOrgConfigs Return All Connected Org Configs from the Federation

[experimental] Returns all connected org configs in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected orgs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@return ListConnectedOrgConfigsApiRequest
*/
func (a *FederatedAuthenticationApiService) ListConnectedOrgConfigs(ctx context.Context, federationSettingsId string) ListConnectedOrgConfigsApiRequest {
	return ListConnectedOrgConfigsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
	}
}

// Execute executes the request
//
//	@return []ConnectedOrgConfig
func (a *FederatedAuthenticationApiService) listConnectedOrgConfigsExecute(r ListConnectedOrgConfigsApiRequest) ([]ConnectedOrgConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ConnectedOrgConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.ListConnectedOrgConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)

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
		var v ApiError
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

type ListIdentityProvidersApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	protocol             *string
}

type ListIdentityProvidersApiParams struct {
	FederationSettingsId string
	Protocol             *string
}

func (a *FederatedAuthenticationApiService) ListIdentityProvidersWithParams(ctx context.Context, args *ListIdentityProvidersApiParams) ListIdentityProvidersApiRequest {
	return ListIdentityProvidersApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		protocol:             args.Protocol,
	}
}

// The protocols of the target identity providers.
func (r ListIdentityProvidersApiRequest) Protocol(protocol string) ListIdentityProvidersApiRequest {
	r.protocol = &protocol
	return r
}

func (r ListIdentityProvidersApiRequest) Execute() (*PaginatedFederationIdentityProvider, *http.Response, error) {
	return r.ApiService.listIdentityProvidersExecute(r)
}

/*
ListIdentityProviders Return all identity providers from the specified federation.

[experimental] Returns all identity providers with the provided protocol in the specified federation. If no protocol is specified, only SAML identity providers will be returned. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@return ListIdentityProvidersApiRequest
*/
func (a *FederatedAuthenticationApiService) ListIdentityProviders(ctx context.Context, federationSettingsId string) ListIdentityProvidersApiRequest {
	return ListIdentityProvidersApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
	}
}

// Execute executes the request
//
//	@return PaginatedFederationIdentityProvider
func (a *FederatedAuthenticationApiService) listIdentityProvidersExecute(r ListIdentityProvidersApiRequest) (*PaginatedFederationIdentityProvider, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedFederationIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.ListIdentityProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.protocol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "protocol", r.protocol, "")
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
		var v ApiError
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

type ListRoleMappingsApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	orgId                string
}

type ListRoleMappingsApiParams struct {
	FederationSettingsId string
	OrgId                string
}

func (a *FederatedAuthenticationApiService) ListRoleMappingsWithParams(ctx context.Context, args *ListRoleMappingsApiParams) ListRoleMappingsApiRequest {
	return ListRoleMappingsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		orgId:                args.OrgId,
	}
}

func (r ListRoleMappingsApiRequest) Execute() (*PaginatedRoleMapping, *http.Response, error) {
	return r.ApiService.listRoleMappingsExecute(r)
}

/*
ListRoleMappings Return All Role Mappings from One Organization

[experimental] Returns all role mappings from the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListRoleMappingsApiRequest
*/
func (a *FederatedAuthenticationApiService) ListRoleMappings(ctx context.Context, federationSettingsId string, orgId string) ListRoleMappingsApiRequest {
	return ListRoleMappingsApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		orgId:                orgId,
	}
}

// Execute executes the request
//
//	@return PaginatedRoleMapping
func (a *FederatedAuthenticationApiService) listRoleMappingsExecute(r ListRoleMappingsApiRequest) (*PaginatedRoleMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedRoleMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.ListRoleMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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
		var v ApiError
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

type RemoveConnectedOrgConfigApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	orgId                string
}

type RemoveConnectedOrgConfigApiParams struct {
	FederationSettingsId string
	OrgId                string
}

func (a *FederatedAuthenticationApiService) RemoveConnectedOrgConfigWithParams(ctx context.Context, args *RemoveConnectedOrgConfigApiParams) RemoveConnectedOrgConfigApiRequest {
	return RemoveConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		orgId:                args.OrgId,
	}
}

func (r RemoveConnectedOrgConfigApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.removeConnectedOrgConfigExecute(r)
}

/*
RemoveConnectedOrgConfig Remove One Org Config Connected to One Federation

[experimental] Removes one connected organization configuration from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role. Note: This request fails if only one connected organization exists in the federation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.
	@return RemoveConnectedOrgConfigApiRequest
*/
func (a *FederatedAuthenticationApiService) RemoveConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string) RemoveConnectedOrgConfigApiRequest {
	return RemoveConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		orgId:                orgId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *FederatedAuthenticationApiService) removeConnectedOrgConfigExecute(r RemoveConnectedOrgConfigApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.RemoveConnectedOrgConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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
		var v ApiError
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

type UpdateConnectedOrgConfigApiRequest struct {
	ctx                  context.Context
	ApiService           FederatedAuthenticationApi
	federationSettingsId string
	orgId                string
	connectedOrgConfig   *ConnectedOrgConfig
}

type UpdateConnectedOrgConfigApiParams struct {
	FederationSettingsId string
	OrgId                string
	ConnectedOrgConfig   *ConnectedOrgConfig
}

func (a *FederatedAuthenticationApiService) UpdateConnectedOrgConfigWithParams(ctx context.Context, args *UpdateConnectedOrgConfigApiParams) UpdateConnectedOrgConfigApiRequest {
	return UpdateConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: args.FederationSettingsId,
		orgId:                args.OrgId,
		connectedOrgConfig:   args.ConnectedOrgConfig,
	}
}

func (r UpdateConnectedOrgConfigApiRequest) Execute() (*ConnectedOrgConfig, *http.Response, error) {
	return r.ApiService.updateConnectedOrgConfigExecute(r)
}

/*
UpdateConnectedOrgConfig Update One Org Config Connected to One Federation

[experimental] Updates one connected organization configuration from the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

**Note** If the organization configuration has no associated identity provider, you can't use this resource to update role mappings or post authorization role grants.

**Note**: The domainRestrictionEnabled field defaults to false if not provided in the request.

**Note**: If the identityProviderId field is not provided, you will disconnect the organization and the identity provider.

**Note**: Currently connected data access identity providers missing from the dataAccessIdentityProviderIds field will be disconnected.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param orgId Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.
	@return UpdateConnectedOrgConfigApiRequest
*/
func (a *FederatedAuthenticationApiService) UpdateConnectedOrgConfig(ctx context.Context, federationSettingsId string, orgId string, connectedOrgConfig *ConnectedOrgConfig) UpdateConnectedOrgConfigApiRequest {
	return UpdateConnectedOrgConfigApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		federationSettingsId: federationSettingsId,
		orgId:                orgId,
		connectedOrgConfig:   connectedOrgConfig,
	}
}

// Execute executes the request
//
//	@return ConnectedOrgConfig
func (a *FederatedAuthenticationApiService) updateConnectedOrgConfigExecute(r UpdateConnectedOrgConfigApiRequest) (*ConnectedOrgConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConnectedOrgConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.UpdateConnectedOrgConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.connectedOrgConfig == nil {
		return localVarReturnValue, nil, reportError("connectedOrgConfig is required and must be specified")
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
	localVarPostBody = r.connectedOrgConfig
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
		var v ApiError
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

type UpdateIdentityProviderApiRequest struct {
	ctx                              context.Context
	ApiService                       FederatedAuthenticationApi
	federationSettingsId             string
	identityProviderId               string
	federationIdentityProviderUpdate *FederationIdentityProviderUpdate
}

type UpdateIdentityProviderApiParams struct {
	FederationSettingsId             string
	IdentityProviderId               string
	FederationIdentityProviderUpdate *FederationIdentityProviderUpdate
}

func (a *FederatedAuthenticationApiService) UpdateIdentityProviderWithParams(ctx context.Context, args *UpdateIdentityProviderApiParams) UpdateIdentityProviderApiRequest {
	return UpdateIdentityProviderApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		federationSettingsId:             args.FederationSettingsId,
		identityProviderId:               args.IdentityProviderId,
		federationIdentityProviderUpdate: args.FederationIdentityProviderUpdate,
	}
}

func (r UpdateIdentityProviderApiRequest) Execute() (*FederationIdentityProvider, *http.Response, error) {
	return r.ApiService.updateIdentityProviderExecute(r)
}

/*
UpdateIdentityProvider Update the identity provider.

[experimental] Updates one identity provider in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role in one of the connected organizations. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param identityProviderId Unique 20-hexadecimal digit string that identifies the identity provider.
	@return UpdateIdentityProviderApiRequest
*/
func (a *FederatedAuthenticationApiService) UpdateIdentityProvider(ctx context.Context, federationSettingsId string, identityProviderId string, federationIdentityProviderUpdate *FederationIdentityProviderUpdate) UpdateIdentityProviderApiRequest {
	return UpdateIdentityProviderApiRequest{
		ApiService:                       a,
		ctx:                              ctx,
		federationSettingsId:             federationSettingsId,
		identityProviderId:               identityProviderId,
		federationIdentityProviderUpdate: federationIdentityProviderUpdate,
	}
}

// Execute executes the request
//
//	@return FederationIdentityProvider
func (a *FederatedAuthenticationApiService) updateIdentityProviderExecute(r UpdateIdentityProviderApiRequest) (*FederationIdentityProvider, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FederationIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.UpdateIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderId"+"}", url.PathEscape(parameterValueToString(r.identityProviderId, "identityProviderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.federationIdentityProviderUpdate == nil {
		return localVarReturnValue, nil, reportError("federationIdentityProviderUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-11-15+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.federationIdentityProviderUpdate
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
		var v ApiError
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

type UpdateRoleMappingApiRequest struct {
	ctx                       context.Context
	ApiService                FederatedAuthenticationApi
	federationSettingsId      string
	id                        string
	orgId                     string
	authFederationRoleMapping *AuthFederationRoleMapping
}

type UpdateRoleMappingApiParams struct {
	FederationSettingsId      string
	Id                        string
	OrgId                     string
	AuthFederationRoleMapping *AuthFederationRoleMapping
}

func (a *FederatedAuthenticationApiService) UpdateRoleMappingWithParams(ctx context.Context, args *UpdateRoleMappingApiParams) UpdateRoleMappingApiRequest {
	return UpdateRoleMappingApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		federationSettingsId:      args.FederationSettingsId,
		id:                        args.Id,
		orgId:                     args.OrgId,
		authFederationRoleMapping: args.AuthFederationRoleMapping,
	}
}

func (r UpdateRoleMappingApiRequest) Execute() (*AuthFederationRoleMapping, *http.Response, error) {
	return r.ApiService.updateRoleMappingExecute(r)
}

/*
UpdateRoleMapping Update One Role Mapping in One Organization

[experimental] Updates one role mapping in the specified organization in the specified federation. To use this resource, the requesting API Key must have the Organization Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param federationSettingsId Unique 24-hexadecimal digit string that identifies your federation.
	@param id Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return UpdateRoleMappingApiRequest
*/
func (a *FederatedAuthenticationApiService) UpdateRoleMapping(ctx context.Context, federationSettingsId string, id string, orgId string, authFederationRoleMapping *AuthFederationRoleMapping) UpdateRoleMappingApiRequest {
	return UpdateRoleMappingApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		federationSettingsId:      federationSettingsId,
		id:                        id,
		orgId:                     orgId,
		authFederationRoleMapping: authFederationRoleMapping,
	}
}

// Execute executes the request
//
//	@return AuthFederationRoleMapping
func (a *FederatedAuthenticationApiService) updateRoleMappingExecute(r UpdateRoleMappingApiRequest) (*AuthFederationRoleMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AuthFederationRoleMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FederatedAuthenticationApiService.UpdateRoleMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"federationSettingsId"+"}", url.PathEscape(parameterValueToString(r.federationSettingsId, "federationSettingsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authFederationRoleMapping == nil {
		return localVarReturnValue, nil, reportError("authFederationRoleMapping is required and must be specified")
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
	localVarPostBody = r.authFederationRoleMapping
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
		var v ApiError
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
