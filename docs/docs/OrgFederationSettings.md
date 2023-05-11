# OrgFederationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FederatedDomains** | Pointer to **[]string** | List of domains associated with the organization&#39;s identity provider. | [optional] 
**HasRoleMappings** | Pointer to **bool** | Flag that indicates whether this organization has role mappings configured. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this federation. | [optional] [readonly] 
**IdentityProviderId** | Pointer to **string** | Unique 20-hexadecimal digit string that identifies the identity provider connected to this organization. | [optional] 
**IdentityProviderStatus** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewOrgFederationSettings

`func NewOrgFederationSettings() *OrgFederationSettings`

NewOrgFederationSettings instantiates a new OrgFederationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgFederationSettingsWithDefaults

`func NewOrgFederationSettingsWithDefaults() *OrgFederationSettings`

NewOrgFederationSettingsWithDefaults instantiates a new OrgFederationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFederatedDomains

`func (o *OrgFederationSettings) GetFederatedDomains() []string`

GetFederatedDomains returns the FederatedDomains field if non-nil, zero value otherwise.

### GetFederatedDomainsOk

`func (o *OrgFederationSettings) GetFederatedDomainsOk() (*[]string, bool)`

GetFederatedDomainsOk returns a tuple with the FederatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedDomains

`func (o *OrgFederationSettings) SetFederatedDomains(v []string)`

SetFederatedDomains sets FederatedDomains field to given value.

### HasFederatedDomains

`func (o *OrgFederationSettings) HasFederatedDomains() bool`

HasFederatedDomains returns a boolean if a field has been set.

### GetHasRoleMappings

`func (o *OrgFederationSettings) GetHasRoleMappings() bool`

GetHasRoleMappings returns the HasRoleMappings field if non-nil, zero value otherwise.

### GetHasRoleMappingsOk

`func (o *OrgFederationSettings) GetHasRoleMappingsOk() (*bool, bool)`

GetHasRoleMappingsOk returns a tuple with the HasRoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRoleMappings

`func (o *OrgFederationSettings) SetHasRoleMappings(v bool)`

SetHasRoleMappings sets HasRoleMappings field to given value.

### HasHasRoleMappings

`func (o *OrgFederationSettings) HasHasRoleMappings() bool`

HasHasRoleMappings returns a boolean if a field has been set.

### GetId

`func (o *OrgFederationSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgFederationSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgFederationSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgFederationSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProviderId

`func (o *OrgFederationSettings) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *OrgFederationSettings) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *OrgFederationSettings) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *OrgFederationSettings) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### GetIdentityProviderStatus

`func (o *OrgFederationSettings) GetIdentityProviderStatus() string`

GetIdentityProviderStatus returns the IdentityProviderStatus field if non-nil, zero value otherwise.

### GetIdentityProviderStatusOk

`func (o *OrgFederationSettings) GetIdentityProviderStatusOk() (*string, bool)`

GetIdentityProviderStatusOk returns a tuple with the IdentityProviderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderStatus

`func (o *OrgFederationSettings) SetIdentityProviderStatus(v string)`

SetIdentityProviderStatus sets IdentityProviderStatus field to given value.

### HasIdentityProviderStatus

`func (o *OrgFederationSettings) HasIdentityProviderStatus() bool`

HasIdentityProviderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


