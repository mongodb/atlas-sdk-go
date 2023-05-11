# IdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsUrl** | Pointer to **string** | URL that points to where to send the SAML response. | [optional] 
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**AssociatedOrgs** | Pointer to [**[]ConnectedOrgConfig**](ConnectedOrgConfig.md) | List that contains the connected organization configurations associated with the identity provider. | [optional] 
**AudienceUri** | Pointer to **string** | Unique string that identifies the intended audience of the SAML assertion. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion. | [optional] 
**OktaIdpId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 
**PemFileInfo** | Pointer to [**PemFileInfo**](PemFileInfo.md) |  | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**SsoDebugEnabled** | Pointer to **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | [optional] 
**SsoUrl** | Pointer to **string** | URL that points to the receiver of the SAML authentication request. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewIdentityProvider

`func NewIdentityProvider(oktaIdpId string, ) *IdentityProvider`

NewIdentityProvider instantiates a new IdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderWithDefaults

`func NewIdentityProviderWithDefaults() *IdentityProvider`

NewIdentityProviderWithDefaults instantiates a new IdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *IdentityProvider) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *IdentityProvider) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *IdentityProvider) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *IdentityProvider) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetAssociatedDomains

`func (o *IdentityProvider) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *IdentityProvider) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *IdentityProvider) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *IdentityProvider) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.

### GetAssociatedOrgs

`func (o *IdentityProvider) GetAssociatedOrgs() []ConnectedOrgConfig`

GetAssociatedOrgs returns the AssociatedOrgs field if non-nil, zero value otherwise.

### GetAssociatedOrgsOk

`func (o *IdentityProvider) GetAssociatedOrgsOk() (*[]ConnectedOrgConfig, bool)`

GetAssociatedOrgsOk returns a tuple with the AssociatedOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedOrgs

`func (o *IdentityProvider) SetAssociatedOrgs(v []ConnectedOrgConfig)`

SetAssociatedOrgs sets AssociatedOrgs field to given value.

### HasAssociatedOrgs

`func (o *IdentityProvider) HasAssociatedOrgs() bool`

HasAssociatedOrgs returns a boolean if a field has been set.

### GetAudienceUri

`func (o *IdentityProvider) GetAudienceUri() string`

GetAudienceUri returns the AudienceUri field if non-nil, zero value otherwise.

### GetAudienceUriOk

`func (o *IdentityProvider) GetAudienceUriOk() (*string, bool)`

GetAudienceUriOk returns a tuple with the AudienceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceUri

`func (o *IdentityProvider) SetAudienceUri(v string)`

SetAudienceUri sets AudienceUri field to given value.

### HasAudienceUri

`func (o *IdentityProvider) HasAudienceUri() bool`

HasAudienceUri returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIssuerUri

`func (o *IdentityProvider) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *IdentityProvider) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *IdentityProvider) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *IdentityProvider) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.

### GetOktaIdpId

`func (o *IdentityProvider) GetOktaIdpId() string`

GetOktaIdpId returns the OktaIdpId field if non-nil, zero value otherwise.

### GetOktaIdpIdOk

`func (o *IdentityProvider) GetOktaIdpIdOk() (*string, bool)`

GetOktaIdpIdOk returns a tuple with the OktaIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaIdpId

`func (o *IdentityProvider) SetOktaIdpId(v string)`

SetOktaIdpId sets OktaIdpId field to given value.


### GetPemFileInfo

`func (o *IdentityProvider) GetPemFileInfo() PemFileInfo`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *IdentityProvider) GetPemFileInfoOk() (*PemFileInfo, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *IdentityProvider) SetPemFileInfo(v PemFileInfo)`

SetPemFileInfo sets PemFileInfo field to given value.

### HasPemFileInfo

`func (o *IdentityProvider) HasPemFileInfo() bool`

HasPemFileInfo returns a boolean if a field has been set.

### GetRequestBinding

`func (o *IdentityProvider) GetRequestBinding() string`

GetRequestBinding returns the RequestBinding field if non-nil, zero value otherwise.

### GetRequestBindingOk

`func (o *IdentityProvider) GetRequestBindingOk() (*string, bool)`

GetRequestBindingOk returns a tuple with the RequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBinding

`func (o *IdentityProvider) SetRequestBinding(v string)`

SetRequestBinding sets RequestBinding field to given value.

### HasRequestBinding

`func (o *IdentityProvider) HasRequestBinding() bool`

HasRequestBinding returns a boolean if a field has been set.

### GetResponseSignatureAlgorithm

`func (o *IdentityProvider) GetResponseSignatureAlgorithm() string`

GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field if non-nil, zero value otherwise.

### GetResponseSignatureAlgorithmOk

`func (o *IdentityProvider) GetResponseSignatureAlgorithmOk() (*string, bool)`

GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSignatureAlgorithm

`func (o *IdentityProvider) SetResponseSignatureAlgorithm(v string)`

SetResponseSignatureAlgorithm sets ResponseSignatureAlgorithm field to given value.

### HasResponseSignatureAlgorithm

`func (o *IdentityProvider) HasResponseSignatureAlgorithm() bool`

HasResponseSignatureAlgorithm returns a boolean if a field has been set.

### GetSsoDebugEnabled

`func (o *IdentityProvider) GetSsoDebugEnabled() bool`

GetSsoDebugEnabled returns the SsoDebugEnabled field if non-nil, zero value otherwise.

### GetSsoDebugEnabledOk

`func (o *IdentityProvider) GetSsoDebugEnabledOk() (*bool, bool)`

GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDebugEnabled

`func (o *IdentityProvider) SetSsoDebugEnabled(v bool)`

SetSsoDebugEnabled sets SsoDebugEnabled field to given value.

### HasSsoDebugEnabled

`func (o *IdentityProvider) HasSsoDebugEnabled() bool`

HasSsoDebugEnabled returns a boolean if a field has been set.

### GetSsoUrl

`func (o *IdentityProvider) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *IdentityProvider) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *IdentityProvider) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *IdentityProvider) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


