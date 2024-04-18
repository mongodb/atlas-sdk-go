# FederationOidcIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedOrgs** | Pointer to [**[]ConnectedOrgConfig**](ConnectedOrgConfig.md) | List that contains the connected organization configurations associated with the identity provider. | [optional] 
**Audience** | Pointer to **string** | Identifier of the intended recipient of the token. | [optional] 
**AuthorizationType** | Pointer to **string** | Indicates whether authorization is granted based on group membership or user ID. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date that the identity provider was created on. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**GroupsClaim** | Pointer to **string** | Identifier of the claim which contains IdP Group IDs in the token. | [optional] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the identity provider. | [readonly] 
**IdpType** | Pointer to **string** | String enum that indicates the type of the identity provider. Default is WORKFORCE. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL. | [optional] 
**OktaIdpId** | **string** | Legacy 20-hexadecimal digit string that identifies the identity provider. | 
**Protocol** | Pointer to **string** | The protocol of the identity provider. Either SAML or OIDC. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date that the identity provider was last updated on. | [optional] [readonly] 
**UserClaim** | Pointer to **string** | Identifier of the claim which contains the user ID in the token. | [optional] 
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**ClientId** | Pointer to **string** | Client identifier that is assigned to an application by the Identity Provider. | [optional] 
**RequestedScopes** | Pointer to **[]string** | Scopes that MongoDB applications will request from the authorization endpoint. | [optional] 

## Methods

### NewFederationOidcIdentityProvider

`func NewFederationOidcIdentityProvider(id string, oktaIdpId string, ) *FederationOidcIdentityProvider`

NewFederationOidcIdentityProvider instantiates a new FederationOidcIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationOidcIdentityProviderWithDefaults

`func NewFederationOidcIdentityProviderWithDefaults() *FederationOidcIdentityProvider`

NewFederationOidcIdentityProviderWithDefaults instantiates a new FederationOidcIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedOrgs

`func (o *FederationOidcIdentityProvider) GetAssociatedOrgs() []ConnectedOrgConfig`

GetAssociatedOrgs returns the AssociatedOrgs field if non-nil, zero value otherwise.

### GetAssociatedOrgsOk

`func (o *FederationOidcIdentityProvider) GetAssociatedOrgsOk() (*[]ConnectedOrgConfig, bool)`

GetAssociatedOrgsOk returns a tuple with the AssociatedOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedOrgs

`func (o *FederationOidcIdentityProvider) SetAssociatedOrgs(v []ConnectedOrgConfig)`

SetAssociatedOrgs sets AssociatedOrgs field to given value.

### HasAssociatedOrgs

`func (o *FederationOidcIdentityProvider) HasAssociatedOrgs() bool`

HasAssociatedOrgs returns a boolean if a field has been set.
### GetAudience

`func (o *FederationOidcIdentityProvider) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *FederationOidcIdentityProvider) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *FederationOidcIdentityProvider) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *FederationOidcIdentityProvider) HasAudience() bool`

HasAudience returns a boolean if a field has been set.
### GetAuthorizationType

`func (o *FederationOidcIdentityProvider) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *FederationOidcIdentityProvider) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *FederationOidcIdentityProvider) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *FederationOidcIdentityProvider) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.
### GetCreatedAt

`func (o *FederationOidcIdentityProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationOidcIdentityProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationOidcIdentityProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederationOidcIdentityProvider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetDescription

`func (o *FederationOidcIdentityProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationOidcIdentityProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationOidcIdentityProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationOidcIdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetDisplayName

`func (o *FederationOidcIdentityProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FederationOidcIdentityProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FederationOidcIdentityProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FederationOidcIdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.
### GetGroupsClaim

`func (o *FederationOidcIdentityProvider) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *FederationOidcIdentityProvider) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *FederationOidcIdentityProvider) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *FederationOidcIdentityProvider) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.
### GetId

`func (o *FederationOidcIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationOidcIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationOidcIdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### GetIdpType

`func (o *FederationOidcIdentityProvider) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FederationOidcIdentityProvider) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FederationOidcIdentityProvider) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *FederationOidcIdentityProvider) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.
### GetIssuerUri

`func (o *FederationOidcIdentityProvider) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *FederationOidcIdentityProvider) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *FederationOidcIdentityProvider) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *FederationOidcIdentityProvider) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.
### GetOktaIdpId

`func (o *FederationOidcIdentityProvider) GetOktaIdpId() string`

GetOktaIdpId returns the OktaIdpId field if non-nil, zero value otherwise.

### GetOktaIdpIdOk

`func (o *FederationOidcIdentityProvider) GetOktaIdpIdOk() (*string, bool)`

GetOktaIdpIdOk returns a tuple with the OktaIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaIdpId

`func (o *FederationOidcIdentityProvider) SetOktaIdpId(v string)`

SetOktaIdpId sets OktaIdpId field to given value.

### GetProtocol

`func (o *FederationOidcIdentityProvider) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationOidcIdentityProvider) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationOidcIdentityProvider) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationOidcIdentityProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.
### GetUpdatedAt

`func (o *FederationOidcIdentityProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederationOidcIdentityProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederationOidcIdentityProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederationOidcIdentityProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.
### GetUserClaim

`func (o *FederationOidcIdentityProvider) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *FederationOidcIdentityProvider) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *FederationOidcIdentityProvider) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *FederationOidcIdentityProvider) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.
### GetAssociatedDomains

`func (o *FederationOidcIdentityProvider) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *FederationOidcIdentityProvider) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *FederationOidcIdentityProvider) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *FederationOidcIdentityProvider) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.
### GetClientId

`func (o *FederationOidcIdentityProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FederationOidcIdentityProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FederationOidcIdentityProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FederationOidcIdentityProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetRequestedScopes

`func (o *FederationOidcIdentityProvider) GetRequestedScopes() []string`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *FederationOidcIdentityProvider) GetRequestedScopesOk() (*[]string, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *FederationOidcIdentityProvider) SetRequestedScopes(v []string)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *FederationOidcIdentityProvider) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


