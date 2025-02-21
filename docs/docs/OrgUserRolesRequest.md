# OrgUserRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupRoleAssignments** | Pointer to [**[]GroupRoleAssignment**](GroupRoleAssignment.md) | List of project level role assignments to assign the MongoDB Cloud user. | [optional] 
**OrgRoles** | **[]string** | One or more organization level roles to assign the MongoDB Cloud user. | 

## Methods

### NewOrgUserRolesRequest

`func NewOrgUserRolesRequest(orgRoles []string, ) *OrgUserRolesRequest`

NewOrgUserRolesRequest instantiates a new OrgUserRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUserRolesRequestWithDefaults

`func NewOrgUserRolesRequestWithDefaults() *OrgUserRolesRequest`

NewOrgUserRolesRequestWithDefaults instantiates a new OrgUserRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupRoleAssignments

`func (o *OrgUserRolesRequest) GetGroupRoleAssignments() []GroupRoleAssignment`

GetGroupRoleAssignments returns the GroupRoleAssignments field if non-nil, zero value otherwise.

### GetGroupRoleAssignmentsOk

`func (o *OrgUserRolesRequest) GetGroupRoleAssignmentsOk() (*[]GroupRoleAssignment, bool)`

GetGroupRoleAssignmentsOk returns a tuple with the GroupRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoleAssignments

`func (o *OrgUserRolesRequest) SetGroupRoleAssignments(v []GroupRoleAssignment)`

SetGroupRoleAssignments sets GroupRoleAssignments field to given value.

### HasGroupRoleAssignments

`func (o *OrgUserRolesRequest) HasGroupRoleAssignments() bool`

HasGroupRoleAssignments returns a boolean if a field has been set.
### GetOrgRoles

`func (o *OrgUserRolesRequest) GetOrgRoles() []string`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *OrgUserRolesRequest) GetOrgRolesOk() (*[]string, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *OrgUserRolesRequest) SetOrgRoles(v []string)`

SetOrgRoles sets OrgRoles field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


