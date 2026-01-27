# OrganizationInvitationGroupRoleAssignmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to which these roles belong. | [optional] 
**Roles** | Pointer to **[]string** | One or more project-level roles to assign to the MongoDB Cloud user. | [optional] 

## Methods

### NewOrganizationInvitationGroupRoleAssignmentsRequest

`func NewOrganizationInvitationGroupRoleAssignmentsRequest() *OrganizationInvitationGroupRoleAssignmentsRequest`

NewOrganizationInvitationGroupRoleAssignmentsRequest instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults

`func NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults() *OrganizationInvitationGroupRoleAssignmentsRequest`

NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetRoles

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


