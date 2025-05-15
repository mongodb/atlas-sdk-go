# FederationIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedOrgs** | Pointer to [**[]ConnectedOrgConfig**](ConnectedOrgConfig.md) | List that contains the connected organization configurations associated with the identity provider. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date that the identity provider was created on. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the identity provider. | [readonly] 
**IdpType** | Pointer to **string** | String enum that indicates the type of the identity provider. Default is WORKFORCE. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL. | [optional] 
**OktaIdpId** | **string** | Legacy 20-hexadecimal digit string that identifies the identity provider. | 
**Protocol** | Pointer to **string** | String enum that indicates the protocol of the identity provider. Either SAML or OIDC. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date that the identity provider was last updated on. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**AcsUrl** | Pointer to **string** | URL that points to where to send the SAML response. | [optional] 
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**AudienceUri** | Pointer to **string** | Unique string that identifies the intended audience of the SAML assertion. | [optional] 
**PemFileInfo** | Pointer to [**PemFileInfo**](PemFileInfo.md) |  | [optional] 
**RequestBinding** | Pointer to **string** | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. | [optional] 
**ResponseSignatureAlgorithm** | Pointer to **string** | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. | [optional] 
**Slug** | Pointer to **string** | Custom SSO Url for the identity provider. | [optional] 
**SsoDebugEnabled** | Pointer to **bool** | Flag that indicates whether the identity provider has SSO debug enabled. | [optional] 
**SsoUrl** | Pointer to **string** | URL that points to the receiver of the SAML authentication request. | [optional] 
**Status** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 
**Audience** | Pointer to **string** | Identifier of the intended recipient of the token. | [optional] 
**AuthorizationType** | Pointer to **string** | Indicates whether authorization is granted based on group membership or user ID. | [optional] 
**ClientId** | Pointer to **string** | Client identifier that is assigned to an application by the Identity Provider. | [optional] 
**GroupsClaim** | Pointer to **string** | Identifier of the claim which contains IdP Group IDs in the token. | [optional] 
**RequestedScopes** | Pointer to **[]string** | Scopes that MongoDB applications will request from the authorization endpoint. | [optional] 
**UserClaim** | Pointer to **string** | Identifier of the claim which contains the user ID in the token. | [optional] 

## Methods

### NewFederationIdentityProvider

`func NewFederationIdentityProvider(id string, oktaIdpId string, ) *FederationIdentityProvider`

NewFederationIdentityProvider instantiates a new FederationIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationIdentityProviderWithDefaults

`func NewFederationIdentityProviderWithDefaults() *FederationIdentityProvider`

NewFederationIdentityProviderWithDefaults instantiates a new FederationIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetCreatedAt

`func (o *FederationIdentityProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationIdentityProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationIdentityProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederationIdentityProvider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *FederationIdentityProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationIdentityProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationIdentityProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationIdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
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
### GetId

`func (o *FederationIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationIdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### GetIdpType

`func (o *FederationIdentityProvider) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FederationIdentityProvider) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FederationIdentityProvider) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *FederationIdentityProvider) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.
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

### GetProtocol

`func (o *FederationIdentityProvider) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationIdentityProvider) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationIdentityProvider) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationIdentityProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.
### GetUpdatedAt

`func (o *FederationIdentityProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederationIdentityProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederationIdentityProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederationIdentityProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.
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
### GetSlug

`func (o *FederationIdentityProvider) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FederationIdentityProvider) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FederationIdentityProvider) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FederationIdentityProvider) HasSlug() bool`

HasSlug returns a boolean if a field has been set.
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
### GetAudience

`func (o *FederationIdentityProvider) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *FederationIdentityProvider) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *FederationIdentityProvider) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *FederationIdentityProvider) HasAudience() bool`

HasAudience returns a boolean if a field has been set.
### GetAuthorizationType

`func (o *FederationIdentityProvider) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *FederationIdentityProvider) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *FederationIdentityProvider) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *FederationIdentityProvider) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.
### GetClientId

`func (o *FederationIdentityProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FederationIdentityProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FederationIdentityProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FederationIdentityProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetGroupsClaim

`func (o *FederationIdentityProvider) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *FederationIdentityProvider) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *FederationIdentityProvider) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *FederationIdentityProvider) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.
### GetRequestedScopes

`func (o *FederationIdentityProvider) GetRequestedScopes() []string`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *FederationIdentityProvider) GetRequestedScopesOk() (*[]string, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *FederationIdentityProvider) SetRequestedScopes(v []string)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *FederationIdentityProvider) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.
### GetUserClaim

`func (o *FederationIdentityProvider) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *FederationIdentityProvider) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *FederationIdentityProvider) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *FederationIdentityProvider) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


