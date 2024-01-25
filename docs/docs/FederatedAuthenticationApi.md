# \FederatedAuthenticationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleMapping**](FederatedAuthenticationApi.md#CreateRoleMapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
[**DeleteFederationApp**](FederatedAuthenticationApi.md#DeleteFederationApp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete the federation settings instance.
[**DeleteRoleMapping**](FederatedAuthenticationApi.md#DeleteRoleMapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
[**GetConnectedOrgConfig**](FederatedAuthenticationApi.md#GetConnectedOrgConfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
[**GetFederationSettings**](FederatedAuthenticationApi.md#GetFederationSettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
[**GetIdentityProvider**](FederatedAuthenticationApi.md#GetIdentityProvider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return one identity provider from the specified federation by id.
[**GetIdentityProviderMetadata**](FederatedAuthenticationApi.md#GetIdentityProviderMetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the metadata of one identity provider in the specified federation.
[**GetRoleMapping**](FederatedAuthenticationApi.md#GetRoleMapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
[**ListConnectedOrgConfigs**](FederatedAuthenticationApi.md#ListConnectedOrgConfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from the Federation
[**ListIdentityProviders**](FederatedAuthenticationApi.md#ListIdentityProviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return all identity providers from the specified federation.
[**ListRoleMappings**](FederatedAuthenticationApi.md#ListRoleMappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
[**RemoveConnectedOrgConfig**](FederatedAuthenticationApi.md#RemoveConnectedOrgConfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
[**UpdateConnectedOrgConfig**](FederatedAuthenticationApi.md#UpdateConnectedOrgConfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Org Config Connected to One Federation
[**UpdateIdentityProvider**](FederatedAuthenticationApi.md#UpdateIdentityProvider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update the identity provider.
[**UpdateRoleMapping**](FederatedAuthenticationApi.md#UpdateRoleMapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization



## CreateRoleMapping

> AuthFederationRoleMapping CreateRoleMapping(ctx, federationSettingsId, orgId, authFederationRoleMapping AuthFederationRoleMapping).Execute()

Add One Role Mapping to One Organization


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 
    authFederationRoleMapping := *openapiclient.NewAuthFederationRoleMapping("ExternalGroupName_example") // AuthFederationRoleMapping | 

    resp, r, err := sdk.FederatedAuthenticationApi.CreateRoleMapping(context.Background(), federationSettingsId, orgId, &authFederationRoleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.CreateRoleMapping``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `CreateRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.CreateRoleMapping`: %v\n", resp)
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

Delete the federation settings instance.


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 

    r, err := sdk.FederatedAuthenticationApi.DeleteFederationApp(context.Background(), federationSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteFederationApp``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
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


## DeleteRoleMapping

> DeleteRoleMapping(ctx, federationSettingsId, id, orgId).Execute()

Remove One Role Mapping from One Organization


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    r, err := sdk.FederatedAuthenticationApi.DeleteRoleMapping(context.Background(), federationSettingsId, id, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteRoleMapping``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
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


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetConnectedOrgConfig``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetConnectedOrgConfig`: ConnectedOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetConnectedOrgConfig`: %v\n", resp)
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


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetFederationSettings(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetFederationSettings``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetFederationSettings`: OrgFederationSettings
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetFederationSettings`: %v\n", resp)
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

Return one identity provider from the specified federation by id.


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "c2777a9eca931f29fc2f" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetIdentityProvider(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProvider``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetIdentityProvider`: FederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

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

Return the metadata of one identity provider in the specified federation.


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "c2777a9eca931f29fc2f" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetIdentityProviderMetadata(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProviderMetadata``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetIdentityProviderMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProviderMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

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


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.GetRoleMapping(context.Background(), federationSettingsId, id, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetRoleMapping``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `GetRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetRoleMapping`: %v\n", resp)
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

> []ConnectedOrgConfig ListConnectedOrgConfigs(ctx, federationSettingsId).Execute()

Return All Connected Org Configs from the Federation


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.ListConnectedOrgConfigs(context.Background(), federationSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListConnectedOrgConfigs``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListConnectedOrgConfigs`: []ConnectedOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListConnectedOrgConfigs`: %v\n", resp)
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


### Return type

[**[]ConnectedOrgConfig**](ConnectedOrgConfig.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviders

> PaginatedFederationIdentityProvider ListIdentityProviders(ctx, federationSettingsId).Protocol(protocol).Execute()

Return all identity providers from the specified federation.


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    protocol := "protocol_example" // string |  (optional)

    resp, r, err := sdk.FederatedAuthenticationApi.ListIdentityProviders(context.Background(), federationSettingsId).Protocol(protocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListIdentityProviders``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListIdentityProviders`: PaginatedFederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListIdentityProviders`: %v\n", resp)
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

 **protocol** | **string** | The protocols of the target identity providers. | 

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


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.ListRoleMappings(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.ListRoleMappings``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `ListRoleMappings`: PaginatedRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.ListRoleMappings`: %v\n", resp)
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

> map[string]interface{} RemoveConnectedOrgConfig(ctx, federationSettingsId, orgId).Execute()

Remove One Org Config Connected to One Federation


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 

    resp, r, err := sdk.FederatedAuthenticationApi.RemoveConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.RemoveConnectedOrgConfig``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `RemoveConnectedOrgConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.RemoveConnectedOrgConfig`: %v\n", resp)
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

**map[string]interface{}**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectedOrgConfig

> ConnectedOrgConfig UpdateConnectedOrgConfig(ctx, federationSettingsId, orgId, connectedOrgConfig ConnectedOrgConfig).Execute()

Update One Org Config Connected to One Federation


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    orgId := "32b6e34b3d91647abb20e7b8" // string | 
    connectedOrgConfig := *openapiclient.NewConnectedOrgConfig(false, "IdentityProviderId_example", "32b6e34b3d91647abb20e7b8") // ConnectedOrgConfig | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateConnectedOrgConfig(context.Background(), federationSettingsId, orgId, &connectedOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateConnectedOrgConfig``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateConnectedOrgConfig`: ConnectedOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateConnectedOrgConfig`: %v\n", resp)
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

Update the identity provider.


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    identityProviderId := "c2777a9eca931f29fc2f" // string | 
    federationIdentityProviderUpdate := *openapiclient.NewFederationIdentityProviderUpdate() // FederationIdentityProviderUpdate | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateIdentityProvider(context.Background(), federationSettingsId, identityProviderId, &federationIdentityProviderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateIdentityProvider``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateIdentityProvider`: FederationIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

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


## Experimental

This operation is marked as experimental. It might be changed in the future without compatibility guarantees.
For more information see [ExperimentalMethods](../doc_1_concepts.md#experimental-methods)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))

    federationSettingsId := "55fa922fb343282757d9554e" // string | 
    id := "32b6e34b3d91647abb20e7b8" // string | 
    orgId := "4888442a3354817a7320eb61" // string | 
    authFederationRoleMapping := *openapiclient.NewAuthFederationRoleMapping("ExternalGroupName_example") // AuthFederationRoleMapping | 

    resp, r, err := sdk.FederatedAuthenticationApi.UpdateRoleMapping(context.Background(), federationSettingsId, id, orgId, &authFederationRoleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateRoleMapping``: %v\n", err)
        apiError := admin.AsError(err)
        fmt.Fprintf(os.Stderr, "Error obj: %v\n", apiError)
    }
    // response from `UpdateRoleMapping`: AuthFederationRoleMapping
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateRoleMapping`: %v\n", resp)
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

