# CustomDBRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]DBAction**](DBAction.md) | List of the individual privilege actions that the role grants. | [optional] 
**InheritedRoles** | Pointer to [**[]InheritedRole**](InheritedRole.md) | List of the built-in roles that this custom role inherits. | [optional] 
**RoleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

## Methods

### NewCustomDBRole

`func NewCustomDBRole(roleName string, ) *CustomDBRole`

NewCustomDBRole instantiates a new CustomDBRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDBRoleWithDefaults

`func NewCustomDBRoleWithDefaults() *CustomDBRole`

NewCustomDBRoleWithDefaults instantiates a new CustomDBRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CustomDBRole) GetActions() []DBAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CustomDBRole) GetActionsOk() (*[]DBAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CustomDBRole) SetActions(v []DBAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CustomDBRole) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetInheritedRoles

`func (o *CustomDBRole) GetInheritedRoles() []InheritedRole`

GetInheritedRoles returns the InheritedRoles field if non-nil, zero value otherwise.

### GetInheritedRolesOk

`func (o *CustomDBRole) GetInheritedRolesOk() (*[]InheritedRole, bool)`

GetInheritedRolesOk returns a tuple with the InheritedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedRoles

`func (o *CustomDBRole) SetInheritedRoles(v []InheritedRole)`

SetInheritedRoles sets InheritedRoles field to given value.

### HasInheritedRoles

`func (o *CustomDBRole) HasInheritedRoles() bool`

HasInheritedRoles returns a boolean if a field has been set.

### GetRoleName

`func (o *CustomDBRole) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *CustomDBRole) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *CustomDBRole) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


