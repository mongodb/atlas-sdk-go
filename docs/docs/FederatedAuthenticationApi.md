# \FederatedAuthenticationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProvider**](FederatedAuthenticationApi.md#CreateIdentityProvider) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Create One Identity Provider
[**CreateRoleMapping**](FederatedAuthenticationApi.md#CreateRoleMapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
[**DeleteFederationApp**](FederatedAuthenticationApi.md#DeleteFederationApp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete One Federation Settings Instance
[**DeleteIdentityProvider**](FederatedAuthenticationApi.md#DeleteIdentityProvider) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Delete One Identity Provider
[**DeleteRoleMapping**](FederatedAuthenticationApi.md#DeleteRoleMapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
[**GetConnectedOrgConfig**](FederatedAuthenticationApi.md#GetConnectedOrgConfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
[**GetFederationSettings**](FederatedAuthenticationApi.md#GetFederationSettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
[**GetIdentityProvider**](FederatedAuthenticationApi.md#GetIdentityProvider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return One Identity Provider by ID
[**GetIdentityProviderMetadata**](FederatedAuthenticationApi.md#GetIdentityProviderMetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the Metadata of One Identity Provider
[**GetRoleMapping**](FederatedAuthenticationApi.md#GetRoleMapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
[**ListConnectedOrgConfigs**](FederatedAuthenticationApi.md#ListConnectedOrgConfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from One Federation
[**ListIdentityProviders**](FederatedAuthenticationApi.md#ListIdentityProviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return All Identity Providers in One Federation
[**ListRoleMappings**](FederatedAuthenticationApi.md#ListRoleMappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
[**RemoveConnectedOrgConfig**](FederatedAuthenticationApi.md#RemoveConnectedOrgConfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
[**RevokeJwksFromIdentityProvider**](FederatedAuthenticationApi.md#RevokeJwksFromIdentityProvider) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/jwks | Revoke the JWKS from One OIDC Identity Provider
[**UpdateConnectedOrgConfig**](FederatedAuthenticationApi.md#UpdateConnectedOrgConfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Org Config Connected to One Federation
[**UpdateIdentityProvider**](FederatedAuthenticationApi.md#UpdateIdentityProvider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update One Identity Provider
[**UpdateRoleMapping**](FederatedAuthenticationApi.md#UpdateRoleMapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization



## CreateIdentityProvider

> FederationOidcIdentityProvider CreateIdentityProvider(ctx, federationSettingsId, federationOidcIdentityProviderUpdate FederationOidcIdentityProviderUpdate).Execute()

Create One Identity Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    federationOidcIdentityProviderUpdate := *openapiclient.NewFederationOidcIdentityProviderUpdate() // FederationOidcIdentityProviderUpdate | 

    resp, r, err := sdk.FederatedAuthenticationApi.CreateIdentityProvider(context.Background(), federationSettingsId, &federationOidcIdentityProviderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.CreateIdentityProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateIdentityProvider`: FederationOidcIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.CreateIdentityProvider`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **federationOidcIdentityProviderUpdate** | [**FederationOidcIdentityProviderUpdate**](FederationOidcIdentityProviderUpdate.md) | The identity provider that you want to create. | 

### Return type

[**FederationOidcIdentityProvider**](FederationOidcIdentityProvider.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-11-15+json
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleMapping

> AuthFederationRoleMapping CreateRoleMapping(ctx, federationSettingsId, orgId, authFederationRoleMapping AuthFederationRoleMapping).Execute()

Add One Role Mapping to One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 
    authFederationRoleMapping := *openapiclient.NewAuthFederationRoleMapping("ExternalGroupName_example") // AuthFederationRoleMapping | 

    resp, r, err := sdk.FederatedAuthenticationApi.CreateRoleMapping(context.Background(), federationSettingsId, orgId, &authFederationRoleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.CreateRoleMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.CreateRoleMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authFederationRoleMapping** | [**AuthFederationRoleMapping**](AuthFederationRoleMapping.md) | The role mapping that you want to create. | 

### Return type

[**AuthFederationRoleMapping**](AuthFederationRoleMapping.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFederationApp

> DeleteFederationApp(ctx, federationSettingsId).Execute()

Delete One Federation Settings Instance


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 

    r, err := sdk.FederatedAuthenticationApi.DeleteFederationApp(context.Background(), federationSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteFederationApp`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFederationAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, federationSettingsId, identityProviderId).Execute()

Delete One Identity Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.FederatedAuthenticationApi.DeleteIdentityProvider(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteIdentityProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 24-hexadecimal digit string that identifies the identity provider to connect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleMapping

> DeleteRoleMapping(ctx, federationSettingsId, id, orgId).Execute()

Remove One Role Mapping from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    r, err := sdk.FederatedAuthenticationApi.DeleteRoleMapping(context.Background(), federationSettingsId, id, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteRoleMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedOrgConfig

> ConnectedOrgConfig GetConnectedOrgConfig(ctx, federationSettingsId, orgId).Execute()

Return One Org Config Connected to One Federation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetConnectedOrgConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetConnectedOrgConfig`: ConnectedOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetConnectedOrgConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectedOrgConfig**](ConnectedOrgConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFederationSettings

> OrgFederationSettings GetFederationSettings(ctx, orgId).Execute()

Return Federation Settings for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetFederationSettings(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetFederationSettings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetFederationSettings`: OrgFederationSettings
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetFederationSettings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFederationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgFederationSettings**](OrgFederationSettings.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProvider

> FederationIdentityProvider GetIdentityProvider(ctx, federationSettingsId, identityProviderId).Execute()

Return One Identity Provider by ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetIdentityProvider(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetIdentityProvider`: FederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProvider`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console &gt; Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FederationIdentityProvider**](FederationIdentityProvider.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviderMetadata

> string GetIdentityProviderMetadata(ctx, federationSettingsId, identityProviderId).Execute()

Return the Metadata of One Identity Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "c2777a9eca931f29fc2f" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetIdentityProviderMetadata(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProviderMetadata`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetIdentityProviderMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProviderMetadata`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Legacy 20-hexadecimal digit string that identifies the identity provider. This id can be found within the Federation Management Console &gt; Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMapping

> AuthFederationRoleMapping GetRoleMapping(ctx, federationSettingsId, id, orgId).Execute()

Return One Role Mapping from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetRoleMapping(context.Background(), federationSettingsId, id, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetRoleMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetRoleMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to return. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AuthFederationRoleMapping**](AuthFederationRoleMapping.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectedOrgConfigs

> PaginatedConnectedOrgConfigs ListConnectedOrgConfigs(ctx, federationSettingsId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Connected Org Configs from One Federation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.FederatedAuthenticationApi.ListConnectedOrgConfigs(context.Background(), federationSettingsId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListConnectedOrgConfigs`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListConnectedOrgConfigs`: PaginatedConnectedOrgConfigs
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListConnectedOrgConfigs`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectedOrgConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedConnectedOrgConfigs**](PaginatedConnectedOrgConfigs.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviders

> PaginatedFederationIdentityProvider ListIdentityProviders(ctx, federationSettingsId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Protocol(protocol).IdpType(idpType).Execute()

Return All Identity Providers in One Federation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    protocol := []string{"Inner_example"} // []string |  (optional)
    idpType := []string{"Inner_example"} // []string |  (optional)

    resp, r, err := sdk.FederatedAuthenticationApi.ListIdentityProviders(context.Background(), federationSettingsId).ItemsPerPage(itemsPerPage).PageNum(pageNum).Protocol(protocol).IdpType(idpType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListIdentityProviders`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListIdentityProviders`: PaginatedFederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListIdentityProviders`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **protocol** | **[]string** | The protocols of the target identity providers. | 
 **idpType** | **[]string** | The types of the target identity providers. | 

### Return type

[**PaginatedFederationIdentityProvider**](PaginatedFederationIdentityProvider.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleMappings

> PaginatedRoleMapping ListRoleMappings(ctx, federationSettingsId, orgId).Execute()

Return All Role Mappings from One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.ListRoleMappings(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListRoleMappings`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListRoleMappings`: PaginatedRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListRoleMappings`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaginatedRoleMapping**](PaginatedRoleMapping.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveConnectedOrgConfig

> any RemoveConnectedOrgConfig(ctx, federationSettingsId, orgId).Execute()

Remove One Org Config Connected to One Federation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.RemoveConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.RemoveConnectedOrgConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `RemoveConnectedOrgConfig`: any
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.RemoveConnectedOrgConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveConnectedOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeJwksFromIdentityProvider

> RevokeJwksFromIdentityProvider(ctx, federationSettingsId, identityProviderId).Execute()

Revoke the JWKS from One OIDC Identity Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "32b6e34b3d91647abb20e7b8" // string | 

    r, err := sdk.FederatedAuthenticationApi.RevokeJwksFromIdentityProvider(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.RevokeJwksFromIdentityProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 24-hexadecimal digit string that identifies the identity provider to connect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeJwksFromIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectedOrgConfig

> ConnectedOrgConfig UpdateConnectedOrgConfig(ctx, federationSettingsId, orgId, connectedOrgConfig ConnectedOrgConfig).Execute()

Update One Org Config Connected to One Federation


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 
    connectedOrgConfig := *openapiclient.NewConnectedOrgConfig(false, "32b6e34b3d91647abb20e7b8") // ConnectedOrgConfig | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateConnectedOrgConfig(context.Background(), federationSettingsId, orgId, &connectedOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateConnectedOrgConfig`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateConnectedOrgConfig`: ConnectedOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateConnectedOrgConfig`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectedOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectedOrgConfig** | [**ConnectedOrgConfig**](ConnectedOrgConfig.md) | The connected organization configuration that you want to update. | 

### Return type

[**ConnectedOrgConfig**](ConnectedOrgConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvider

> FederationIdentityProvider UpdateIdentityProvider(ctx, federationSettingsId, identityProviderId, federationIdentityProviderUpdate FederationIdentityProviderUpdate).Execute()

Update One Identity Provider


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "32b6e34b3d91647abb20e7b8" // string | 
    federationIdentityProviderUpdate := *openapiclient.NewFederationIdentityProviderUpdate() // FederationIdentityProviderUpdate | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateIdentityProvider(context.Background(), federationSettingsId, identityProviderId, &federationIdentityProviderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateIdentityProvider`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateIdentityProvider`: FederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateIdentityProvider`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console &gt; Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **federationIdentityProviderUpdate** | [**FederationIdentityProviderUpdate**](FederationIdentityProviderUpdate.md) | The identity provider that you want to update. | 

### Return type

[**FederationIdentityProvider**](FederationIdentityProvider.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-11-15+json
- **Accept**: application/vnd.atlas.2023-11-15+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleMapping

> AuthFederationRoleMapping UpdateRoleMapping(ctx, federationSettingsId, id, orgId, authFederationRoleMapping AuthFederationRoleMapping).Execute()

Update One Role Mapping in One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 
    authFederationRoleMapping := *openapiclient.NewAuthFederationRoleMapping("ExternalGroupName_example") // AuthFederationRoleMapping | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateRoleMapping(context.Background(), federationSettingsId, id, orgId, &authFederationRoleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateRoleMapping`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `UpdateRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateRoleMapping`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to update. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authFederationRoleMapping** | [**AuthFederationRoleMapping**](AuthFederationRoleMapping.md) | The role mapping that you want to update. | 

### Return type

[**AuthFederationRoleMapping**](AuthFederationRoleMapping.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

