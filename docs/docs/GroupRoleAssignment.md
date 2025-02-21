# GroupRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to which these roles belong. | [optional] 
**GroupRoles** | Pointer to **[]string** | One or more project-level roles assigned to the MongoDB Cloud user. | [optional] 

## Methods

### NewGroupRoleAssignment

`func NewGroupRoleAssignment() *GroupRoleAssignment`

NewGroupRoleAssignment instantiates a new GroupRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleAssignmentWithDefaults

`func NewGroupRoleAssignmentWithDefaults() *GroupRoleAssignment`

NewGroupRoleAssignmentWithDefaults instantiates a new GroupRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupRoleAssignment) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupRoleAssignment) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupRoleAssignment) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupRoleAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetGroupRoles

`func (o *GroupRoleAssignment) GetGroupRoles() []string`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### GetGroupRolesOk

`func (o *GroupRoleAssignment) GetGroupRolesOk() (*[]string, bool)`

GetGroupRolesOk returns a tuple with the GroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoles

`func (o *GroupRoleAssignment) SetGroupRoles(v []string)`

SetGroupRoles sets GroupRoles field to given value.

### HasGroupRoles

`func (o *GroupRoleAssignment) HasGroupRoles() bool`

HasGroupRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


