# FederationIdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL. | [optional] 
**Protocol** | Pointer to **string** | The protocol of the identity provider. Either SAML or OIDC. | [optional] 
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**PemFileInfo** | Pointer to [**PemFileInfoUpdate**](PemFileInfoUpdate.md) |  | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**Slug** | Pointer to **string** | Custom SSO Url for the identity provider. | [optional] 
**SsoDebugEnabled** | Pointer to **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | [optional] 
**SsoUrl** | Pointer to **string** | URL that points to the receiver of the SAML authentication request. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 
**AudienceClaim** | Pointer to **[]string** | Identifier of the intended recipient of the token. | [optional] 
**ClientId** | Pointer to **string** | Client identifier that is assigned to an application by the Identity Provider. | [optional] 
**GroupsClaim** | Pointer to **string** | Identifier of the claim which contains IdP Group IDs in the token. | [optional] 
**RequestedScopes** | Pointer to **[]string** | Scopes that MongoDB applications will request from the authorization endpoint. | [optional] 
**UserClaim** | Pointer to **string** | Identifier of the claim which contains the user ID in the token. | [optional] 

## Methods

### NewFederationIdentityProviderUpdate

`func NewFederationIdentityProviderUpdate() *FederationIdentityProviderUpdate`

NewFederationIdentityProviderUpdate instantiates a new FederationIdentityProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationIdentityProviderUpdateWithDefaults

`func NewFederationIdentityProviderUpdateWithDefaults() *FederationIdentityProviderUpdate`

NewFederationIdentityProviderUpdateWithDefaults instantiates a new FederationIdentityProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FederationIdentityProviderUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationIdentityProviderUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationIdentityProviderUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationIdentityProviderUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetDisplayName

`func (o *FederationIdentityProviderUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FederationIdentityProviderUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FederationIdentityProviderUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FederationIdentityProviderUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.
### GetIssuerUri

`func (o *FederationIdentityProviderUpdate) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *FederationIdentityProviderUpdate) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *FederationIdentityProviderUpdate) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *FederationIdentityProviderUpdate) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.
### GetProtocol

`func (o *FederationIdentityProviderUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationIdentityProviderUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationIdentityProviderUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationIdentityProviderUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.
### GetAssociatedDomains

`func (o *FederationIdentityProviderUpdate) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *FederationIdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *FederationIdentityProviderUpdate) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *FederationIdentityProviderUpdate) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.
### GetPemFileInfo

`func (o *FederationIdentityProviderUpdate) GetPemFileInfo() PemFileInfoUpdate`

GetPemFileInfo returns the PemFileInfo field if non-nil, zero value otherwise.

### GetPemFileInfoOk

`func (o *FederationIdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfoUpdate, bool)`

GetPemFileInfoOk returns a tuple with the PemFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemFileInfo

`func (o *FederationIdentityProviderUpdate) SetPemFileInfo(v PemFileInfoUpdate)`

SetPemFileInfo sets PemFileInfo field to given value.

### HasPemFileInfo

`func (o *FederationIdentityProviderUpdate) HasPemFileInfo() bool`

HasPemFileInfo returns a boolean if a field has been set.
### GetRequestBinding

`func (o *FederationIdentityProviderUpdate) GetRequestBinding() string`

GetRequestBinding returns the RequestBinding field if non-nil, zero value otherwise.

### GetRequestBindingOk

`func (o *FederationIdentityProviderUpdate) GetRequestBindingOk() (*string, bool)`

GetRequestBindingOk returns a tuple with the RequestBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBinding

`func (o *FederationIdentityProviderUpdate) SetRequestBinding(v string)`

SetRequestBinding sets RequestBinding field to given value.

### HasRequestBinding

`func (o *FederationIdentityProviderUpdate) HasRequestBinding() bool`

HasRequestBinding returns a boolean if a field has been set.
### GetResponseSignatureAlgorithm

`func (o *FederationIdentityProviderUpdate) GetResponseSignatureAlgorithm() string`

GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field if non-nil, zero value otherwise.

### GetResponseSignatureAlgorithmOk

`func (o *FederationIdentityProviderUpdate) GetResponseSignatureAlgorithmOk() (*string, bool)`

GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSignatureAlgorithm

`func (o *FederationIdentityProviderUpdate) SetResponseSignatureAlgorithm(v string)`

SetResponseSignatureAlgorithm sets ResponseSignatureAlgorithm field to given value.

### HasResponseSignatureAlgorithm

`func (o *FederationIdentityProviderUpdate) HasResponseSignatureAlgorithm() bool`

HasResponseSignatureAlgorithm returns a boolean if a field has been set.
### GetSlug

`func (o *FederationIdentityProviderUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FederationIdentityProviderUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FederationIdentityProviderUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FederationIdentityProviderUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.
### GetSsoDebugEnabled

`func (o *FederationIdentityProviderUpdate) GetSsoDebugEnabled() bool`

GetSsoDebugEnabled returns the SsoDebugEnabled field if non-nil, zero value otherwise.

### GetSsoDebugEnabledOk

`func (o *FederationIdentityProviderUpdate) GetSsoDebugEnabledOk() (*bool, bool)`

GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDebugEnabled

`func (o *FederationIdentityProviderUpdate) SetSsoDebugEnabled(v bool)`

SetSsoDebugEnabled sets SsoDebugEnabled field to given value.

### HasSsoDebugEnabled

`func (o *FederationIdentityProviderUpdate) HasSsoDebugEnabled() bool`

HasSsoDebugEnabled returns a boolean if a field has been set.
### GetSsoUrl

`func (o *FederationIdentityProviderUpdate) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *FederationIdentityProviderUpdate) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *FederationIdentityProviderUpdate) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *FederationIdentityProviderUpdate) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.
### GetStatus

`func (o *FederationIdentityProviderUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FederationIdentityProviderUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FederationIdentityProviderUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FederationIdentityProviderUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetAudienceClaim

`func (o *FederationIdentityProviderUpdate) GetAudienceClaim() []string`

GetAudienceClaim returns the AudienceClaim field if non-nil, zero value otherwise.

### GetAudienceClaimOk

`func (o *FederationIdentityProviderUpdate) GetAudienceClaimOk() (*[]string, bool)`

GetAudienceClaimOk returns a tuple with the AudienceClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceClaim

`func (o *FederationIdentityProviderUpdate) SetAudienceClaim(v []string)`

SetAudienceClaim sets AudienceClaim field to given value.

### HasAudienceClaim

`func (o *FederationIdentityProviderUpdate) HasAudienceClaim() bool`

HasAudienceClaim returns a boolean if a field has been set.
### GetClientId

`func (o *FederationIdentityProviderUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FederationIdentityProviderUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FederationIdentityProviderUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FederationIdentityProviderUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetGroupsClaim

`func (o *FederationIdentityProviderUpdate) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *FederationIdentityProviderUpdate) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *FederationIdentityProviderUpdate) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *FederationIdentityProviderUpdate) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.
### GetRequestedScopes

`func (o *FederationIdentityProviderUpdate) GetRequestedScopes() []string`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *FederationIdentityProviderUpdate) GetRequestedScopesOk() (*[]string, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *FederationIdentityProviderUpdate) SetRequestedScopes(v []string)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *FederationIdentityProviderUpdate) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.
### GetUserClaim

`func (o *FederationIdentityProviderUpdate) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *FederationIdentityProviderUpdate) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *FederationIdentityProviderUpdate) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *FederationIdentityProviderUpdate) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


