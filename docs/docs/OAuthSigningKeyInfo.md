# OAuthSigningKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Signing algorithm of the Atlas-managed key. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | When the currently active signing key was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**JwksUri** | Pointer to **string** | Public JWKS URL serving this integration&#39;s signing keys. Fixed for the lifetime of the integration. | [optional] [readonly] 
**Kid** | Pointer to **string** | Key ID stamped on client assertions, the &#x60;SHA-1&#x60; thumbprint of the key certificate in uppercase hexadecimal. Changes when Atlas rotates the key. | [optional] [readonly] 

## Methods

### NewOAuthSigningKeyInfo

`func NewOAuthSigningKeyInfo() *OAuthSigningKeyInfo`

NewOAuthSigningKeyInfo instantiates a new OAuthSigningKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthSigningKeyInfoWithDefaults

`func NewOAuthSigningKeyInfoWithDefaults() *OAuthSigningKeyInfo`

NewOAuthSigningKeyInfoWithDefaults instantiates a new OAuthSigningKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *OAuthSigningKeyInfo) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *OAuthSigningKeyInfo) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *OAuthSigningKeyInfo) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *OAuthSigningKeyInfo) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *OAuthSigningKeyInfo) SetAlgorithmNil()`

SetAlgorithmNil sets Algorithm to an explicit JSON null when marshaled, overriding any value previously set with SetAlgorithm. Calling SetAlgorithm again clears the null override.

### GetCreatedAt

`func (o *OAuthSigningKeyInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OAuthSigningKeyInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OAuthSigningKeyInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OAuthSigningKeyInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *OAuthSigningKeyInfo) SetCreatedAtNil()`

SetCreatedAtNil sets CreatedAt to an explicit JSON null when marshaled, overriding any value previously set with SetCreatedAt. Calling SetCreatedAt again clears the null override.

### GetJwksUri

`func (o *OAuthSigningKeyInfo) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OAuthSigningKeyInfo) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OAuthSigningKeyInfo) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OAuthSigningKeyInfo) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### SetJwksUriNil

`func (o *OAuthSigningKeyInfo) SetJwksUriNil()`

SetJwksUriNil sets JwksUri to an explicit JSON null when marshaled, overriding any value previously set with SetJwksUri. Calling SetJwksUri again clears the null override.

### GetKid

`func (o *OAuthSigningKeyInfo) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuthSigningKeyInfo) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuthSigningKeyInfo) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuthSigningKeyInfo) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuthSigningKeyInfo) SetKidNil()`

SetKidNil sets Kid to an explicit JSON null when marshaled, overriding any value previously set with SetKid. Calling SetKid again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


