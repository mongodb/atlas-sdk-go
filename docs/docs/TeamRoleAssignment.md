# TeamRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamRoleAssignment

`func NewTeamRoleAssignment() *TeamRoleAssignment`

NewTeamRoleAssignment instantiates a new TeamRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRoleAssignmentWithDefaults

`func NewTeamRoleAssignmentWithDefaults() *TeamRoleAssignment`

NewTeamRoleAssignmentWithDefaults instantiates a new TeamRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *TeamRoleAssignment) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TeamRoleAssignment) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TeamRoleAssignment) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TeamRoleAssignment) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetTeamId

`func (o *TeamRoleAssignment) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamRoleAssignment) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamRoleAssignment) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamRoleAssignment) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


