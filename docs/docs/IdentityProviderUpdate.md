# IdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**Description** | Pointer to **string** | The description for the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion. | [optional] 
**Protocol** | Pointer to **string** | The protocol for the identity provider. | [optional] 
**AudienceClaim** | Pointer to **[]string** | Audience claim for the identity provider. | [optional] 
**ClientId** | Pointer to **string** | Client ID for the identity provider. | [optional] 
**GroupsClaim** | Pointer to **string** | Groups claim for the identity provider. | [optional] 
**RequestedScopes** | Pointer to **[]string** | Requested scopes for the identity provider. | [optional] 
**UserClaim** | Pointer to **string** | User claim for the identity provider. | [optional] 
**PemFileInfo** | Pointer to [**PemFileInfoUpdate**](PemFileInfoUpdate.md) |  | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**Slug** | Pointer to **string** | Custom SSO Url for identity provider. | [optional] 
**SsoDebugEnabled** | Pointer to **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | [optional] 
**SsoUrl** | Pointer to **string** | Unique string that identifies the intended audience of the SAML assertion. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewIdentityProviderUpdate

`func NewIdentityProviderUpdate() *IdentityProviderUpdate`

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
### GetDescription

`func (o *IdentityProviderUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
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
### GetProtocol

`func (o *IdentityProviderUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IdentityProviderUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IdentityProviderUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IdentityProviderUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.
### GetAudienceClaim

`func (o *IdentityProviderUpdate) GetAudienceClaim() []string`

GetAudienceClaim returns the AudienceClaim field if non-nil, zero value otherwise.

### GetAudienceClaimOk

`func (o *IdentityProviderUpdate) GetAudienceClaimOk() (*[]string, bool)`

GetAudienceClaimOk returns a tuple with the AudienceClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceClaim

`func (o *IdentityProviderUpdate) SetAudienceClaim(v []string)`

SetAudienceClaim sets AudienceClaim field to given value.

### HasAudienceClaim

`func (o *IdentityProviderUpdate) HasAudienceClaim() bool`

HasAudienceClaim returns a boolean if a field has been set.
### GetClientId

`func (o *IdentityProviderUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityProviderUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetGroupsClaim

`func (o *IdentityProviderUpdate) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *IdentityProviderUpdate) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *IdentityProviderUpdate) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *IdentityProviderUpdate) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.
### GetRequestedScopes

`func (o *IdentityProviderUpdate) GetRequestedScopes() []string`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *IdentityProviderUpdate) GetRequestedScopesOk() (*[]string, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *IdentityProviderUpdate) SetRequestedScopes(v []string)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *IdentityProviderUpdate) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.
### GetUserClaim

`func (o *IdentityProviderUpdate) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *IdentityProviderUpdate) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *IdentityProviderUpdate) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *IdentityProviderUpdate) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.
### GetPemFileInfo

`func (o *IdentityProviderUpdate) GetPemFileInfo() PemFileInfoUpdate`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *IdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfoUpdate, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *IdentityProviderUpdate) SetPemFileInfo(v PemFileInfoUpdate)`

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
### GetSlug

`func (o *IdentityProviderUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *IdentityProviderUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *IdentityProviderUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *IdentityProviderUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.
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

### HasSsoDebugEnabled

`func (o *IdentityProviderUpdate) HasSsoDebugEnabled() bool`

HasSsoDebugEnabled returns a boolean if a field has been set.
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


