# UserCustomDBRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]DatabasePrivilegeAction**](DatabasePrivilegeAction.md) | List of the individual privilege actions that the role grants. | [optional] 
**InheritedRoles** | Pointer to [**[]DatabaseInheritedRole**](DatabaseInheritedRole.md) | List of the built-in roles that this custom role inherits. | [optional] 
**RoleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

## Methods

### NewUserCustomDBRole

`func NewUserCustomDBRole(roleName string, ) *UserCustomDBRole`

NewUserCustomDBRole instantiates a new UserCustomDBRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCustomDBRoleWithDefaults

`func NewUserCustomDBRoleWithDefaults() *UserCustomDBRole`

NewUserCustomDBRoleWithDefaults instantiates a new UserCustomDBRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *UserCustomDBRole) GetActions() []DatabasePrivilegeAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UserCustomDBRole) GetActionsOk() (*[]DatabasePrivilegeAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UserCustomDBRole) SetActions(v []DatabasePrivilegeAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UserCustomDBRole) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetInheritedRoles

`func (o *UserCustomDBRole) GetInheritedRoles() []DatabaseInheritedRole`

GetInheritedRoles returns the InheritedRoles field if non-nil, zero value otherwise.

### GetInheritedRolesOk

`func (o *UserCustomDBRole) GetInheritedRolesOk() (*[]DatabaseInheritedRole, bool)`

GetInheritedRolesOk returns a tuple with the InheritedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedRoles

`func (o *UserCustomDBRole) SetInheritedRoles(v []DatabaseInheritedRole)`

SetInheritedRoles sets InheritedRoles field to given value.

### HasInheritedRoles

`func (o *UserCustomDBRole) HasInheritedRoles() bool`

HasInheritedRoles returns a boolean if a field has been set.

### GetRoleName

`func (o *UserCustomDBRole) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UserCustomDBRole) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UserCustomDBRole) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


