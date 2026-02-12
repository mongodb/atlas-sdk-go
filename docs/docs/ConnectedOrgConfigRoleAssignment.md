# ConnectedOrgConfigRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to which this role belongs. Each element within &#x60;roleAssignments&#x60; can have a value for &#x60;groupId&#x60; or &#x60;orgId&#x60;, but not both. | [optional] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which this role belongs. Each element within &#x60;roleAssignments&#x60; can have a value for &#x60;orgId&#x60; or &#x60;groupId&#x60;, but not both. | [optional] 
**Role** | Pointer to **string** | Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include organization- and project-level roles. | [optional] 

## Methods

### NewConnectedOrgConfigRoleAssignment

`func NewConnectedOrgConfigRoleAssignment() *ConnectedOrgConfigRoleAssignment`

NewConnectedOrgConfigRoleAssignment instantiates a new ConnectedOrgConfigRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedOrgConfigRoleAssignmentWithDefaults

`func NewConnectedOrgConfigRoleAssignmentWithDefaults() *ConnectedOrgConfigRoleAssignment`

NewConnectedOrgConfigRoleAssignmentWithDefaults instantiates a new ConnectedOrgConfigRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ConnectedOrgConfigRoleAssignment) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectedOrgConfigRoleAssignment) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectedOrgConfigRoleAssignment) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ConnectedOrgConfigRoleAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetOrgId

`func (o *ConnectedOrgConfigRoleAssignment) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConnectedOrgConfigRoleAssignment) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConnectedOrgConfigRoleAssignment) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConnectedOrgConfigRoleAssignment) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.
### GetRole

`func (o *ConnectedOrgConfigRoleAssignment) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConnectedOrgConfigRoleAssignment) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConnectedOrgConfigRoleAssignment) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConnectedOrgConfigRoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


