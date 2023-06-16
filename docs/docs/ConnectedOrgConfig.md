# ConnectedOrgConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataAccessIdentityProviderIds** | Pointer to **[]string** | The collection of unique ids of the identity providers for org&#39;s data access. | [optional] 
**DomainAllowList** | Pointer to **[]string** | Approved domains that restrict users who can join the organization based on their email address. | [optional] 
**DomainRestrictionEnabled** | **bool** | Value that indicates whether domain restriction is enabled for this connected org. | 
**IdentityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with. | 
**OrgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration. | [readonly] 
**PostAuthRoleGrants** | Pointer to **[]string** | Atlas roles that are granted to a user in this organization after authenticating. | [optional] 
**RoleMappings** | Pointer to [**[]RoleMapping**](RoleMapping.md) | Role mappings that are configured in this organization. | [optional] 
**UserConflicts** | Pointer to [**[]FederatedUser**](FederatedUser.md) | List that contains the users who have an email address that doesn&#39;t match any domain on the allowed list. | [optional] 

## Methods

### NewConnectedOrgConfig

`func NewConnectedOrgConfig(domainRestrictionEnabled bool, identityProviderId string, orgId string, ) *ConnectedOrgConfig`

NewConnectedOrgConfig instantiates a new ConnectedOrgConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedOrgConfigWithDefaults

`func NewConnectedOrgConfigWithDefaults() *ConnectedOrgConfig`

NewConnectedOrgConfigWithDefaults instantiates a new ConnectedOrgConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataAccessIdentityProviderIds

`func (o *ConnectedOrgConfig) GetDataAccessIdentityProviderIds() []string`

GetDataAccessIdentityProviderIds returns the DataAccessIdentityProviderIds field if non-nil, zero value otherwise.

### GetDataAccessIdentityProviderIdsOk

`func (o *ConnectedOrgConfig) GetDataAccessIdentityProviderIdsOk() (*[]string, bool)`

GetDataAccessIdentityProviderIdsOk returns a tuple with the DataAccessIdentityProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessIdentityProviderIds

`func (o *ConnectedOrgConfig) SetDataAccessIdentityProviderIds(v []string)`

SetDataAccessIdentityProviderIds sets DataAccessIdentityProviderIds field to given value.

### HasDataAccessIdentityProviderIds

`func (o *ConnectedOrgConfig) HasDataAccessIdentityProviderIds() bool`

HasDataAccessIdentityProviderIds returns a boolean if a field has been set.

### GetDomainAllowList

`func (o *ConnectedOrgConfig) GetDomainAllowList() []string`

GetDomainAllowList returns the DomainAllowList field if non-nil, zero value otherwise.

### GetDomainAllowListOk

`func (o *ConnectedOrgConfig) GetDomainAllowListOk() (*[]string, bool)`

GetDomainAllowListOk returns a tuple with the DomainAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAllowList

`func (o *ConnectedOrgConfig) SetDomainAllowList(v []string)`

SetDomainAllowList sets DomainAllowList field to given value.

### HasDomainAllowList

`func (o *ConnectedOrgConfig) HasDomainAllowList() bool`

HasDomainAllowList returns a boolean if a field has been set.

### GetDomainRestrictionEnabled

`func (o *ConnectedOrgConfig) GetDomainRestrictionEnabled() bool`

GetDomainRestrictionEnabled returns the DomainRestrictionEnabled field if non-nil, zero value otherwise.

### GetDomainRestrictionEnabledOk

`func (o *ConnectedOrgConfig) GetDomainRestrictionEnabledOk() (*bool, bool)`

GetDomainRestrictionEnabledOk returns a tuple with the DomainRestrictionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRestrictionEnabled

`func (o *ConnectedOrgConfig) SetDomainRestrictionEnabled(v bool)`

SetDomainRestrictionEnabled sets DomainRestrictionEnabled field to given value.


### GetIdentityProviderId

`func (o *ConnectedOrgConfig) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *ConnectedOrgConfig) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *ConnectedOrgConfig) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.


### GetOrgId

`func (o *ConnectedOrgConfig) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConnectedOrgConfig) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConnectedOrgConfig) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPostAuthRoleGrants

`func (o *ConnectedOrgConfig) GetPostAuthRoleGrants() []string`

GetPostAuthRoleGrants returns the PostAuthRoleGrants field if non-nil, zero value otherwise.

### GetPostAuthRoleGrantsOk

`func (o *ConnectedOrgConfig) GetPostAuthRoleGrantsOk() (*[]string, bool)`

GetPostAuthRoleGrantsOk returns a tuple with the PostAuthRoleGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthRoleGrants

`func (o *ConnectedOrgConfig) SetPostAuthRoleGrants(v []string)`

SetPostAuthRoleGrants sets PostAuthRoleGrants field to given value.

### HasPostAuthRoleGrants

`func (o *ConnectedOrgConfig) HasPostAuthRoleGrants() bool`

HasPostAuthRoleGrants returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ConnectedOrgConfig) GetRoleMappings() []RoleMapping`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ConnectedOrgConfig) GetRoleMappingsOk() (*[]RoleMapping, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ConnectedOrgConfig) SetRoleMappings(v []RoleMapping)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ConnectedOrgConfig) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetUserConflicts

`func (o *ConnectedOrgConfig) GetUserConflicts() []FederatedUser`

GetUserConflicts returns the UserConflicts field if non-nil, zero value otherwise.

### GetUserConflictsOk

`func (o *ConnectedOrgConfig) GetUserConflictsOk() (*[]FederatedUser, bool)`

GetUserConflictsOk returns a tuple with the UserConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConflicts

`func (o *ConnectedOrgConfig) SetUserConflicts(v []FederatedUser)`

SetUserConflicts sets UserConflicts field to given value.

### HasUserConflicts

`func (o *ConnectedOrgConfig) HasUserConflicts() bool`

HasUserConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


