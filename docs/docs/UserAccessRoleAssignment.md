# UserAccessRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization API key. | [optional] [readonly] 
**Roles** | Pointer to **[]string** | List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project. | [optional] 

## Methods

### NewUserAccessRoleAssignment

`func NewUserAccessRoleAssignment() *UserAccessRoleAssignment`

NewUserAccessRoleAssignment instantiates a new UserAccessRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessRoleAssignmentWithDefaults

`func NewUserAccessRoleAssignmentWithDefaults() *UserAccessRoleAssignment`

NewUserAccessRoleAssignmentWithDefaults instantiates a new UserAccessRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUserId

`func (o *UserAccessRoleAssignment) GetApiUserId() string`

GetApiUserId returns the ApiUserId field if non-nil, zero value otherwise.

### GetApiUserIdOk

`func (o *UserAccessRoleAssignment) GetApiUserIdOk() (*string, bool)`

GetApiUserIdOk returns a tuple with the ApiUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUserId

`func (o *UserAccessRoleAssignment) SetApiUserId(v string)`

SetApiUserId sets ApiUserId field to given value.

### HasApiUserId

`func (o *UserAccessRoleAssignment) HasApiUserId() bool`

HasApiUserId returns a boolean if a field has been set.

### GetRoles

`func (o *UserAccessRoleAssignment) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAccessRoleAssignment) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAccessRoleAssignment) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserAccessRoleAssignment) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


