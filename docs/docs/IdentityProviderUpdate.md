# IdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion. | [optional] 
**PemFileInfo** | Pointer to [**PemFileInfo**](PemFileInfo.md) |  | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**SsoDebugEnabled** | **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | 
**SsoUrl** | Pointer to **string** | Unique string that identifies the intended audience of the SAML assertion. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewIdentityProviderUpdate

`func NewIdentityProviderUpdate(ssoDebugEnabled bool, ) *IdentityProviderUpdate`

NewIdentityProviderUpdate instantiates a new IdentityProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderUpdateWithDefaults

`func NewIdentityProviderUpdateWithDefaults() *IdentityProviderUpdate`

NewIdentityProviderUpdateWithDefaults instantiates a new IdentityProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedDomains

`func (o *IdentityProviderUpdate) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *IdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *IdentityProviderUpdate) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *IdentityProviderUpdate) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityProviderUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIssuerUri

`func (o *IdentityProviderUpdate) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *IdentityProviderUpdate) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *IdentityProviderUpdate) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *IdentityProviderUpdate) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.

### GetPemFileInfo

`func (o *IdentityProviderUpdate) GetPemFileInfo() PemFileInfo`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *IdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfo, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *IdentityProviderUpdate) SetPemFileInfo(v PemFileInfo)`

SetPemFileInfo sets PemFileInfo field to given value.

### HasPemFileInfo

`func (o *IdentityProviderUpdate) HasPemFileInfo() bool`

HasPemFileInfo returns a boolean if a field has been set.

### GetRequestBinding

`func (o *IdentityProviderUpdate) GetRequestBinding() string`

GetRequestBinding returns the RequestBinding field if non-nil, zero value otherwise.

### GetRequestBindingOk

`func (o *IdentityProviderUpdate) GetRequestBindingOk() (*string, bool)`

GetRequestBindingOk returns a tuple with the RequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBinding

`func (o *IdentityProviderUpdate) SetRequestBinding(v string)`

SetRequestBinding sets RequestBinding field to given value.

### HasRequestBinding

`func (o *IdentityProviderUpdate) HasRequestBinding() bool`

HasRequestBinding returns a boolean if a field has been set.

### GetResponseSignatureAlgorithm

`func (o *IdentityProviderUpdate) GetResponseSignatureAlgorithm() string`

GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field if non-nil, zero value otherwise.

### GetResponseSignatureAlgorithmOk

`func (o *IdentityProviderUpdate) GetResponseSignatureAlgorithmOk() (*string, bool)`

GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSignatureAlgorithm

`func (o *IdentityProviderUpdate) SetResponseSignatureAlgorithm(v string)`

SetResponseSignatureAlgorithm sets ResponseSignatureAlgorithm field to given value.

### HasResponseSignatureAlgorithm

`func (o *IdentityProviderUpdate) HasResponseSignatureAlgorithm() bool`

HasResponseSignatureAlgorithm returns a boolean if a field has been set.

### GetSsoDebugEnabled

`func (o *IdentityProviderUpdate) GetSsoDebugEnabled() bool`

GetSsoDebugEnabled returns the SsoDebugEnabled field if non-nil, zero value otherwise.

### GetSsoDebugEnabledOk

`func (o *IdentityProviderUpdate) GetSsoDebugEnabledOk() (*bool, bool)`

GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDebugEnabled

`func (o *IdentityProviderUpdate) SetSsoDebugEnabled(v bool)`

SetSsoDebugEnabled sets SsoDebugEnabled field to given value.


### GetSsoUrl

`func (o *IdentityProviderUpdate) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *IdentityProviderUpdate) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *IdentityProviderUpdate) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *IdentityProviderUpdate) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityProviderUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProviderUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProviderUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProviderUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


