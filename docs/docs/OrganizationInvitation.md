# OrganizationInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupRoleAssignments** | Pointer to [**[]GroupRole**](GroupRole.md) | List of projects that the user will be added to when they accept their invitation to the organization. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this invitation. | [optional] [readonly] 
**InviterUsername** | Pointer to **string** | Email address of the MongoDB Cloud user who sent the invitation to join the organization. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization. | [optional] [readonly] 
**OrgName** | **string** | Human-readable label that identifies this organization. | 
**Roles** | Pointer to **[]string** | One or more organization-level roles to assign to the MongoDB Cloud user. | [optional] 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies each team. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address of the MongoDB Cloud user invited to join the organization. | [optional] 

## Methods

### NewOrganizationInvitation

`func NewOrganizationInvitation(orgName string, ) *OrganizationInvitation`

NewOrganizationInvitation instantiates a new OrganizationInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationWithDefaults

`func NewOrganizationInvitationWithDefaults() *OrganizationInvitation`

NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationInvitation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.
### GetExpiresAt

`func (o *OrganizationInvitation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationInvitation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationInvitation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationInvitation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.
### GetGroupRoleAssignments

`func (o *OrganizationInvitation) GetGroupRoleAssignments() []GroupRole`

GetGroupRoleAssignments returns the GroupRoleAssignments field if non-nil, zero value otherwise.

### GetGroupRoleAssignmentsOk

`func (o *OrganizationInvitation) GetGroupRoleAssignmentsOk() (*[]GroupRole, bool)`

GetGroupRoleAssignmentsOk returns a tuple with the GroupRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoleAssignments

`func (o *OrganizationInvitation) SetGroupRoleAssignments(v []GroupRole)`

SetGroupRoleAssignments sets GroupRoleAssignments field to given value.

### HasGroupRoleAssignments

`func (o *OrganizationInvitation) HasGroupRoleAssignments() bool`

HasGroupRoleAssignments returns a boolean if a field has been set.
### GetId

`func (o *OrganizationInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationInvitation) HasId() bool`

HasId returns a boolean if a field has been set.
### GetInviterUsername

`func (o *OrganizationInvitation) GetInviterUsername() string`

GetInviterUsername returns the InviterUsername field if non-nil, zero value otherwise.

### GetInviterUsernameOk

`func (o *OrganizationInvitation) GetInviterUsernameOk() (*string, bool)`

GetInviterUsernameOk returns a tuple with the InviterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUsername

`func (o *OrganizationInvitation) SetInviterUsername(v string)`

SetInviterUsername sets InviterUsername field to given value.

### HasInviterUsername

`func (o *OrganizationInvitation) HasInviterUsername() bool`

HasInviterUsername returns a boolean if a field has been set.
### GetLinks

`func (o *OrganizationInvitation) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationInvitation) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationInvitation) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganizationInvitation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgId

`func (o *OrganizationInvitation) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrganizationInvitation) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrganizationInvitation) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrganizationInvitation) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetOrgName

`func (o *OrganizationInvitation) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrganizationInvitation) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrganizationInvitation) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### GetRoles

`func (o *OrganizationInvitation) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationInvitation) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationInvitation) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrganizationInvitation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetTeamIds

`func (o *OrganizationInvitation) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrganizationInvitation) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrganizationInvitation) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrganizationInvitation) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.
### GetUsername

`func (o *OrganizationInvitation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OrganizationInvitation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OrganizationInvitation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OrganizationInvitation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


