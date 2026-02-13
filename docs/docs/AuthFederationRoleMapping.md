# AuthFederationRoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalGroupName** | **string** | Unique human-readable label that identifies the identity provider group to which this role mapping applies. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this role mapping. | [optional] [readonly] 
**RoleAssignments** | Pointer to [**[]ConnectedOrgConfigRoleAssignment**](ConnectedOrgConfigRoleAssignment.md) | Atlas roles and the unique identifiers of the groups and organizations associated with each role. The array must include at least one element with an Organization role and its respective &#x60;orgId&#x60;. Each element in the array can have a value for &#x60;orgId&#x60; or &#x60;groupId&#x60;, but not both. | [optional] 

## Methods

### NewAuthFederationRoleMapping

`func NewAuthFederationRoleMapping(externalGroupName string, ) *AuthFederationRoleMapping`

NewAuthFederationRoleMapping instantiates a new AuthFederationRoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthFederationRoleMappingWithDefaults

`func NewAuthFederationRoleMappingWithDefaults() *AuthFederationRoleMapping`

NewAuthFederationRoleMappingWithDefaults instantiates a new AuthFederationRoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalGroupName

`func (o *AuthFederationRoleMapping) GetExternalGroupName() string`

GetExternalGroupName returns the ExternalGroupName field if non-nil, zero value otherwise.

### GetExternalGroupNameOk

`func (o *AuthFederationRoleMapping) GetExternalGroupNameOk() (*string, bool)`

GetExternalGroupNameOk returns a tuple with the ExternalGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupName

`func (o *AuthFederationRoleMapping) SetExternalGroupName(v string)`

SetExternalGroupName sets ExternalGroupName field to given value.

### GetId

`func (o *AuthFederationRoleMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthFederationRoleMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthFederationRoleMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthFederationRoleMapping) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRoleAssignments

`func (o *AuthFederationRoleMapping) GetRoleAssignments() []ConnectedOrgConfigRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *AuthFederationRoleMapping) GetRoleAssignmentsOk() (*[]ConnectedOrgConfigRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *AuthFederationRoleMapping) SetRoleAssignments(v []ConnectedOrgConfigRoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *AuthFederationRoleMapping) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


