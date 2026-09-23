# OAuthConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAuthMethod** | **string** | How the client authenticates to the token endpoint. &#x60;CLIENT_SECRET&#x60; sends a shared secret. &#x60;PRIVATE_KEY_JWT&#x60; signs a client assertion with an Atlas-generated, Atlas-managed key. Register the returned JWKS URL with your identity provider. | 
**ClientId** | **string** | OAuth 2.0 client identifier registered with the token endpoint. | 
**ClientSecret** | Pointer to **string** | Shared client secret. Required when &#x60;clientAuthMethod&#x60; is &#x60;CLIENT_SECRET&#x60;, and rejected for &#x60;PRIVATE_KEY_JWT&#x60;. Encrypted at rest and never returned. | [optional] 
**Scopes** | Pointer to **[]string** | Optional OAuth 2.0 scopes requested on the token, sent as a space delimited &#x60;scope&#x60; parameter. Applies to both client authentication methods. | [optional] 
**TokenEndpoint** | **string** | OAuth 2.0 token endpoint URL. Must use HTTPS. | 
**TokenRequestParams** | Pointer to **map[string]string** | Optional provider-specific parameters added to the token request, for example a resource indicator. Applies to both client authentication methods. | [optional] 

## Methods

### NewOAuthConfigRequest

`func NewOAuthConfigRequest(clientAuthMethod string, clientId string, tokenEndpoint string, ) *OAuthConfigRequest`

NewOAuthConfigRequest instantiates a new OAuthConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthConfigRequestWithDefaults

`func NewOAuthConfigRequestWithDefaults() *OAuthConfigRequest`

NewOAuthConfigRequestWithDefaults instantiates a new OAuthConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAuthMethod

`func (o *OAuthConfigRequest) GetClientAuthMethod() string`

GetClientAuthMethod returns the ClientAuthMethod field if non-nil, zero value otherwise.

### GetClientAuthMethodOk

`func (o *OAuthConfigRequest) GetClientAuthMethodOk() (*string, bool)`

GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthMethod

`func (o *OAuthConfigRequest) SetClientAuthMethod(v string)`

SetClientAuthMethod sets ClientAuthMethod field to given value.

### GetClientId

`func (o *OAuthConfigRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthConfigRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthConfigRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### GetClientSecret

`func (o *OAuthConfigRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthConfigRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthConfigRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthConfigRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuthConfigRequest) SetClientSecretNil()`

SetClientSecretNil sets ClientSecret to an explicit JSON null when marshaled, overriding any value previously set with SetClientSecret. Calling SetClientSecret again clears the null override.

### GetScopes

`func (o *OAuthConfigRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthConfigRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthConfigRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthConfigRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *OAuthConfigRequest) SetScopesNil()`

SetScopesNil sets Scopes to an explicit JSON null when marshaled, overriding any value previously set with SetScopes. Calling SetScopes again clears the null override.

### GetTokenEndpoint

`func (o *OAuthConfigRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAuthConfigRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAuthConfigRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### GetTokenRequestParams

`func (o *OAuthConfigRequest) GetTokenRequestParams() map[string]string`

GetTokenRequestParams returns the TokenRequestParams field if non-nil, zero value otherwise.

### GetTokenRequestParamsOk

`func (o *OAuthConfigRequest) GetTokenRequestParamsOk() (*map[string]string, bool)`

GetTokenRequestParamsOk returns a tuple with the TokenRequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestParams

`func (o *OAuthConfigRequest) SetTokenRequestParams(v map[string]string)`

SetTokenRequestParams sets TokenRequestParams field to given value.

### HasTokenRequestParams

`func (o *OAuthConfigRequest) HasTokenRequestParams() bool`

HasTokenRequestParams returns a boolean if a field has been set.

### SetTokenRequestParamsNil

`func (o *OAuthConfigRequest) SetTokenRequestParamsNil()`

SetTokenRequestParamsNil sets TokenRequestParams to an explicit JSON null when marshaled, overriding any value previously set with SetTokenRequestParams. Calling SetTokenRequestParams again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


