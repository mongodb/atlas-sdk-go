# UpdateCustomDBRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]DatabasePrivilegeAction**](DatabasePrivilegeAction.md) | List of the individual privilege actions that the role grants. | [optional] 
**InheritedRoles** | Pointer to [**[]DatabaseInheritedRole**](DatabaseInheritedRole.md) | List of the built-in roles that this custom role inherits. | [optional] 

## Methods

### NewUpdateCustomDBRole

`func NewUpdateCustomDBRole() *UpdateCustomDBRole`

NewUpdateCustomDBRole instantiates a new UpdateCustomDBRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomDBRoleWithDefaults

`func NewUpdateCustomDBRoleWithDefaults() *UpdateCustomDBRole`

NewUpdateCustomDBRoleWithDefaults instantiates a new UpdateCustomDBRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *UpdateCustomDBRole) GetActions() []DatabasePrivilegeAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdateCustomDBRole) GetActionsOk() (*[]DatabasePrivilegeAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdateCustomDBRole) SetActions(v []DatabasePrivilegeAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdateCustomDBRole) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *UpdateCustomDBRole) SetActionsNil()`

SetActionsNil sets Actions to an explicit JSON null when marshaled, overriding any value previously set with SetActions. Calling SetActions again clears the null override.

### GetInheritedRoles

`func (o *UpdateCustomDBRole) GetInheritedRoles() []DatabaseInheritedRole`

GetInheritedRoles returns the InheritedRoles field if non-nil, zero value otherwise.

### GetInheritedRolesOk

`func (o *UpdateCustomDBRole) GetInheritedRolesOk() (*[]DatabaseInheritedRole, bool)`

GetInheritedRolesOk returns a tuple with the InheritedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedRoles

`func (o *UpdateCustomDBRole) SetInheritedRoles(v []DatabaseInheritedRole)`

SetInheritedRoles sets InheritedRoles field to given value.

### HasInheritedRoles

`func (o *UpdateCustomDBRole) HasInheritedRoles() bool`

HasInheritedRoles returns a boolean if a field has been set.

### SetInheritedRolesNil

`func (o *UpdateCustomDBRole) SetInheritedRolesNil()`

SetInheritedRolesNil sets InheritedRoles to an explicit JSON null when marshaled, overriding any value previously set with SetInheritedRoles. Calling SetInheritedRoles again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


