# SamlIdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**Description** | Pointer to **string** | The description for the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion. | [optional] 
**PemFileInfo** | Pointer to [**PemFileInfo**](PemFileInfo.md) |  | [optional] 
**Protocol** | Pointer to **string** | The protocol for the identity provider. | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**SsoDebugEnabled** | **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | 
**SsoUrl** | Pointer to **string** | Unique string that identifies the intended audience of the SAML assertion. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewSamlIdentityProviderUpdate

`func NewSamlIdentityProviderUpdate(ssoDebugEnabled bool, ) *SamlIdentityProviderUpdate`

NewSamlIdentityProviderUpdate instantiates a new SamlIdentityProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlIdentityProviderUpdateWithDefaults

`func NewSamlIdentityProviderUpdateWithDefaults() *SamlIdentityProviderUpdate`

NewSamlIdentityProviderUpdateWithDefaults instantiates a new SamlIdentityProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedDomains

`func (o *SamlIdentityProviderUpdate) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *SamlIdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *SamlIdentityProviderUpdate) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *SamlIdentityProviderUpdate) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.

### GetDescription

`func (o *SamlIdentityProviderUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SamlIdentityProviderUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SamlIdentityProviderUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SamlIdentityProviderUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SamlIdentityProviderUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SamlIdentityProviderUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SamlIdentityProviderUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SamlIdentityProviderUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIssuerUri

`func (o *SamlIdentityProviderUpdate) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *SamlIdentityProviderUpdate) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *SamlIdentityProviderUpdate) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *SamlIdentityProviderUpdate) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.

### GetPemFileInfo

`func (o *SamlIdentityProviderUpdate) GetPemFileInfo() PemFileInfo`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *SamlIdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfo, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *SamlIdentityProviderUpdate) SetPemFileInfo(v PemFileInfo)`

SetPemFileInfo sets PemFileInfo field to given value.

### HasPemFileInfo

`func (o *SamlIdentityProviderUpdate) HasPemFileInfo() bool`

HasPemFileInfo returns a boolean if a field has been set.

### GetProtocol

`func (o *SamlIdentityProviderUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SamlIdentityProviderUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SamlIdentityProviderUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SamlIdentityProviderUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRequestBinding

`func (o *SamlIdentityProviderUpdate) GetRequestBinding() string`

GetRequestBinding returns the RequestBinding field if non-nil, zero value otherwise.

### GetRequestBindingOk

`func (o *SamlIdentityProviderUpdate) GetRequestBindingOk() (*string, bool)`

GetRequestBindingOk returns a tuple with the RequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBinding

`func (o *SamlIdentityProviderUpdate) SetRequestBinding(v string)`

SetRequestBinding sets RequestBinding field to given value.

### HasRequestBinding

`func (o *SamlIdentityProviderUpdate) HasRequestBinding() bool`

HasRequestBinding returns a boolean if a field has been set.

### GetResponseSignatureAlgorithm

`func (o *SamlIdentityProviderUpdate) GetResponseSignatureAlgorithm() string`

GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field if non-nil, zero value otherwise.

### GetResponseSignatureAlgorithmOk

`func (o *SamlIdentityProviderUpdate) GetResponseSignatureAlgorithmOk() (*string, bool)`

GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSignatureAlgorithm

`func (o *SamlIdentityProviderUpdate) SetResponseSignatureAlgorithm(v string)`

SetResponseSignatureAlgorithm sets ResponseSignatureAlgorithm field to given value.

### HasResponseSignatureAlgorithm

`func (o *SamlIdentityProviderUpdate) HasResponseSignatureAlgorithm() bool`

HasResponseSignatureAlgorithm returns a boolean if a field has been set.

### GetSsoDebugEnabled

`func (o *SamlIdentityProviderUpdate) GetSsoDebugEnabled() bool`

GetSsoDebugEnabled returns the SsoDebugEnabled field if non-nil, zero value otherwise.

### GetSsoDebugEnabledOk

`func (o *SamlIdentityProviderUpdate) GetSsoDebugEnabledOk() (*bool, bool)`

GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDebugEnabled

`func (o *SamlIdentityProviderUpdate) SetSsoDebugEnabled(v bool)`

SetSsoDebugEnabled sets SsoDebugEnabled field to given value.


### GetSsoUrl

`func (o *SamlIdentityProviderUpdate) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SamlIdentityProviderUpdate) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SamlIdentityProviderUpdate) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *SamlIdentityProviderUpdate) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *SamlIdentityProviderUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SamlIdentityProviderUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SamlIdentityProviderUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SamlIdentityProviderUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


