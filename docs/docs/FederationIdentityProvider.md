# FederationIdentityProvider

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

### NewFederationIdentityProvider

`func NewFederationIdentityProvider(oktaIdpId string, ) *FederationIdentityProvider`

NewFederationIdentityProvider instantiates a new FederationIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationIdentityProviderWithDefaults

`func NewFederationIdentityProviderWithDefaults() *FederationIdentityProvider`

NewFederationIdentityProviderWithDefaults instantiates a new FederationIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *FederationIdentityProvider) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *FederationIdentityProvider) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *FederationIdentityProvider) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *FederationIdentityProvider) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetAssociatedDomains

`func (o *FederationIdentityProvider) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *FederationIdentityProvider) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *FederationIdentityProvider) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *FederationIdentityProvider) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.

### GetAssociatedOrgs

`func (o *FederationIdentityProvider) GetAssociatedOrgs() []ConnectedOrgConfig`

GetAssociatedOrgs returns the AssociatedOrgs field if non-nil, zero value otherwise.

### GetAssociatedOrgsOk

`func (o *FederationIdentityProvider) GetAssociatedOrgsOk() (*[]ConnectedOrgConfig, bool)`

GetAssociatedOrgsOk returns a tuple with the AssociatedOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedOrgs

`func (o *FederationIdentityProvider) SetAssociatedOrgs(v []ConnectedOrgConfig)`

SetAssociatedOrgs sets AssociatedOrgs field to given value.

### HasAssociatedOrgs

`func (o *FederationIdentityProvider) HasAssociatedOrgs() bool`

HasAssociatedOrgs returns a boolean if a field has been set.

### GetAudienceUri

`func (o *FederationIdentityProvider) GetAudienceUri() string`

GetAudienceUri returns the AudienceUri field if non-nil, zero value otherwise.

### GetAudienceUriOk

`func (o *FederationIdentityProvider) GetAudienceUriOk() (*string, bool)`

GetAudienceUriOk returns a tuple with the AudienceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceUri

`func (o *FederationIdentityProvider) SetAudienceUri(v string)`

SetAudienceUri sets AudienceUri field to given value.

### HasAudienceUri

`func (o *FederationIdentityProvider) HasAudienceUri() bool`

HasAudienceUri returns a boolean if a field has been set.

### GetDisplayName

`func (o *FederationIdentityProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FederationIdentityProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FederationIdentityProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FederationIdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIssuerUri

`func (o *FederationIdentityProvider) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *FederationIdentityProvider) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *FederationIdentityProvider) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *FederationIdentityProvider) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.

### GetOktaIdpId

`func (o *FederationIdentityProvider) GetOktaIdpId() string`

GetOktaIdpId returns the OktaIdpId field if non-nil, zero value otherwise.

### GetOktaIdpIdOk

`func (o *FederationIdentityProvider) GetOktaIdpIdOk() (*string, bool)`

GetOktaIdpIdOk returns a tuple with the OktaIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaIdpId

`func (o *FederationIdentityProvider) SetOktaIdpId(v string)`

SetOktaIdpId sets OktaIdpId field to given value.


### GetPemFileInfo

`func (o *FederationIdentityProvider) GetPemFileInfo() PemFileInfo`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *FederationIdentityProvider) GetPemFileInfoOk() (*PemFileInfo, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *FederationIdentityProvider) SetPemFileInfo(v PemFileInfo)`

SetPemFileInfo sets PemFileInfo field to given value.

### HasPemFileInfo

`func (o *FederationIdentityProvider) HasPemFileInfo() bool`

HasPemFileInfo returns a boolean if a field has been set.

### GetRequestBinding

`func (o *FederationIdentityProvider) GetRequestBinding() string`

GetRequestBinding returns the RequestBinding field if non-nil, zero value otherwise.

### GetRequestBindingOk

`func (o *FederationIdentityProvider) GetRequestBindingOk() (*string, bool)`

GetRequestBindingOk returns a tuple with the RequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBinding

`func (o *FederationIdentityProvider) SetRequestBinding(v string)`

SetRequestBinding sets RequestBinding field to given value.

### HasRequestBinding

`func (o *FederationIdentityProvider) HasRequestBinding() bool`

HasRequestBinding returns a boolean if a field has been set.

### GetResponseSignatureAlgorithm

`func (o *FederationIdentityProvider) GetResponseSignatureAlgorithm() string`

GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field if non-nil, zero value otherwise.

### GetResponseSignatureAlgorithmOk

`func (o *FederationIdentityProvider) GetResponseSignatureAlgorithmOk() (*string, bool)`

GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSignatureAlgorithm

`func (o *FederationIdentityProvider) SetResponseSignatureAlgorithm(v string)`

SetResponseSignatureAlgorithm sets ResponseSignatureAlgorithm field to given value.

### HasResponseSignatureAlgorithm

`func (o *FederationIdentityProvider) HasResponseSignatureAlgorithm() bool`

HasResponseSignatureAlgorithm returns a boolean if a field has been set.

### GetSsoDebugEnabled

`func (o *FederationIdentityProvider) GetSsoDebugEnabled() bool`

GetSsoDebugEnabled returns the SsoDebugEnabled field if non-nil, zero value otherwise.

### GetSsoDebugEnabledOk

`func (o *FederationIdentityProvider) GetSsoDebugEnabledOk() (*bool, bool)`

GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDebugEnabled

`func (o *FederationIdentityProvider) SetSsoDebugEnabled(v bool)`

SetSsoDebugEnabled sets SsoDebugEnabled field to given value.

### HasSsoDebugEnabled

`func (o *FederationIdentityProvider) HasSsoDebugEnabled() bool`

HasSsoDebugEnabled returns a boolean if a field has been set.

### GetSsoUrl

`func (o *FederationIdentityProvider) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *FederationIdentityProvider) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *FederationIdentityProvider) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *FederationIdentityProvider) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FederationIdentityProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FederationIdentityProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FederationIdentityProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FederationIdentityProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


