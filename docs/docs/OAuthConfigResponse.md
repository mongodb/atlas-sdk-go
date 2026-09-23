# OAuthConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAuthMethod** | **string** | How the client authenticates to the token endpoint. | 
**ClientId** | **string** | OAuth 2.0 client identifier registered with the token endpoint. | 
**Scopes** | Pointer to **[]string** | OAuth 2.0 scopes requested on the token. | [optional] 
**SigningKeyInfo** | Pointer to [**OAuthSigningKeyInfo**](OAuthSigningKeyInfo.md) |  | [optional] 
**TokenEndpoint** | **string** | OAuth 2.0 token endpoint URL. | 
**TokenRequestParams** | Pointer to **map[string]string** | Provider-specific parameters added to the token request. | [optional] 

## Methods

### NewOAuthConfigResponse

`func NewOAuthConfigResponse(clientAuthMethod string, clientId string, tokenEndpoint string, ) *OAuthConfigResponse`

NewOAuthConfigResponse instantiates a new OAuthConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthConfigResponseWithDefaults

`func NewOAuthConfigResponseWithDefaults() *OAuthConfigResponse`

NewOAuthConfigResponseWithDefaults instantiates a new OAuthConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAuthMethod

`func (o *OAuthConfigResponse) GetClientAuthMethod() string`

GetClientAuthMethod returns the ClientAuthMethod field if non-nil, zero value otherwise.

### GetClientAuthMethodOk

`func (o *OAuthConfigResponse) GetClientAuthMethodOk() (*string, bool)`

GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthMethod

`func (o *OAuthConfigResponse) SetClientAuthMethod(v string)`

SetClientAuthMethod sets ClientAuthMethod field to given value.

### GetClientId

`func (o *OAuthConfigResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthConfigResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthConfigResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### GetScopes

`func (o *OAuthConfigResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthConfigResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthConfigResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthConfigResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *OAuthConfigResponse) SetScopesNil()`

SetScopesNil sets Scopes to an explicit JSON null when marshaled, overriding any value previously set with SetScopes. Calling SetScopes again clears the null override.

### GetSigningKeyInfo

`func (o *OAuthConfigResponse) GetSigningKeyInfo() OAuthSigningKeyInfo`

GetSigningKeyInfo returns the SigningKeyInfo field if non-nil, zero value otherwise.

### GetSigningKeyInfoOk

`func (o *OAuthConfigResponse) GetSigningKeyInfoOk() (*OAuthSigningKeyInfo, bool)`

GetSigningKeyInfoOk returns a tuple with the SigningKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyInfo

`func (o *OAuthConfigResponse) SetSigningKeyInfo(v OAuthSigningKeyInfo)`

SetSigningKeyInfo sets SigningKeyInfo field to given value.

### HasSigningKeyInfo

`func (o *OAuthConfigResponse) HasSigningKeyInfo() bool`

HasSigningKeyInfo returns a boolean if a field has been set.

### SetSigningKeyInfoNil

`func (o *OAuthConfigResponse) SetSigningKeyInfoNil()`

SetSigningKeyInfoNil sets SigningKeyInfo to an explicit JSON null when marshaled, overriding any value previously set with SetSigningKeyInfo. Calling SetSigningKeyInfo again clears the null override.

### GetTokenEndpoint

`func (o *OAuthConfigResponse) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAuthConfigResponse) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAuthConfigResponse) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### GetTokenRequestParams

`func (o *OAuthConfigResponse) GetTokenRequestParams() map[string]string`

GetTokenRequestParams returns the TokenRequestParams field if non-nil, zero value otherwise.

### GetTokenRequestParamsOk

`func (o *OAuthConfigResponse) GetTokenRequestParamsOk() (*map[string]string, bool)`

GetTokenRequestParamsOk returns a tuple with the TokenRequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestParams

`func (o *OAuthConfigResponse) SetTokenRequestParams(v map[string]string)`

SetTokenRequestParams sets TokenRequestParams field to given value.

### HasTokenRequestParams

`func (o *OAuthConfigResponse) HasTokenRequestParams() bool`

HasTokenRequestParams returns a boolean if a field has been set.

### SetTokenRequestParamsNil

`func (o *OAuthConfigResponse) SetTokenRequestParamsNil()`

SetTokenRequestParamsNil sets TokenRequestParams to an explicit JSON null when marshaled, overriding any value previously set with SetTokenRequestParams. Calling SetTokenRequestParams again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


