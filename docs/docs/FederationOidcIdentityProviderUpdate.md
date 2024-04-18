# FederationOidcIdentityProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Identifier of the intended recipient of the token. | [optional] 
**AuthorizationType** | Pointer to **string** | Indicates whether authorization is granted based on group membership or user ID. | [optional] 
**Description** | Pointer to **string** | The description of the identity provider. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label that identifies the identity provider. | [optional] 
**GroupsClaim** | Pointer to **string** | Identifier of the claim which contains IdP Group IDs in the token. | [optional] 
**IdpType** | Pointer to **string** | String enum that indicates the type of the identity provider. Default is WORKFORCE. | [optional] 
**IssuerUri** | Pointer to **string** | Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL. | [optional] 
**Protocol** | Pointer to **string** | The protocol of the identity provider. Either SAML or OIDC. | [optional] 
**UserClaim** | Pointer to **string** | Identifier of the claim which contains the user ID in the token. | [optional] 
**AssociatedDomains** | Pointer to **[]string** | List that contains the domains associated with the identity provider. | [optional] 
**ClientId** | Pointer to **string** | Client identifier that is assigned to an application by the Identity Provider. | [optional] 
**RequestedScopes** | Pointer to **[]string** | Scopes that MongoDB applications will request from the authorization endpoint. | [optional] 

## Methods

### NewFederationOidcIdentityProviderUpdate

`func NewFederationOidcIdentityProviderUpdate() *FederationOidcIdentityProviderUpdate`

NewFederationOidcIdentityProviderUpdate instantiates a new FederationOidcIdentityProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationOidcIdentityProviderUpdateWithDefaults

`func NewFederationOidcIdentityProviderUpdateWithDefaults() *FederationOidcIdentityProviderUpdate`

NewFederationOidcIdentityProviderUpdateWithDefaults instantiates a new FederationOidcIdentityProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *FederationOidcIdentityProviderUpdate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *FederationOidcIdentityProviderUpdate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *FederationOidcIdentityProviderUpdate) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *FederationOidcIdentityProviderUpdate) HasAudience() bool`

HasAudience returns a boolean if a field has been set.
### GetAuthorizationType

`func (o *FederationOidcIdentityProviderUpdate) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *FederationOidcIdentityProviderUpdate) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *FederationOidcIdentityProviderUpdate) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *FederationOidcIdentityProviderUpdate) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.
### GetDescription

`func (o *FederationOidcIdentityProviderUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationOidcIdentityProviderUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationOidcIdentityProviderUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationOidcIdentityProviderUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetDisplayName

`func (o *FederationOidcIdentityProviderUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FederationOidcIdentityProviderUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FederationOidcIdentityProviderUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FederationOidcIdentityProviderUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.
### GetGroupsClaim

`func (o *FederationOidcIdentityProviderUpdate) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *FederationOidcIdentityProviderUpdate) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *FederationOidcIdentityProviderUpdate) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *FederationOidcIdentityProviderUpdate) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.
### GetIdpType

`func (o *FederationOidcIdentityProviderUpdate) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FederationOidcIdentityProviderUpdate) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FederationOidcIdentityProviderUpdate) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *FederationOidcIdentityProviderUpdate) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.
### GetIssuerUri

`func (o *FederationOidcIdentityProviderUpdate) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *FederationOidcIdentityProviderUpdate) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *FederationOidcIdentityProviderUpdate) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.

### HasIssuerUri

`func (o *FederationOidcIdentityProviderUpdate) HasIssuerUri() bool`

HasIssuerUri returns a boolean if a field has been set.
### GetProtocol

`func (o *FederationOidcIdentityProviderUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationOidcIdentityProviderUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationOidcIdentityProviderUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationOidcIdentityProviderUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.
### GetUserClaim

`func (o *FederationOidcIdentityProviderUpdate) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *FederationOidcIdentityProviderUpdate) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *FederationOidcIdentityProviderUpdate) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *FederationOidcIdentityProviderUpdate) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.
### GetAssociatedDomains

`func (o *FederationOidcIdentityProviderUpdate) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *FederationOidcIdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *FederationOidcIdentityProviderUpdate) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *FederationOidcIdentityProviderUpdate) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.
### GetClientId

`func (o *FederationOidcIdentityProviderUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FederationOidcIdentityProviderUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FederationOidcIdentityProviderUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FederationOidcIdentityProviderUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.
### GetRequestedScopes

`func (o *FederationOidcIdentityProviderUpdate) GetRequestedScopes() []string`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *FederationOidcIdentityProviderUpdate) GetRequestedScopesOk() (*[]string, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *FederationOidcIdentityProviderUpdate) SetRequestedScopes(v []string)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *FederationOidcIdentityProviderUpdate) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


