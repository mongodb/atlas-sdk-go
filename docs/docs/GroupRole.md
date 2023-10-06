# GroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to which this role belongs. | [optional] 
**GroupRole** | Pointer to **string** | Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include project-level roles.  Project Roles  * GROUP_CLUSTER_MANAGER * GROUP_DATA_ACCESS_ADMIN * GROUP_DATA_ACCESS_READ_ONLY * GROUP_DATA_ACCESS_READ_WRITE * GROUP_OWNER * GROUP_READ_ONLY * GROUP_SEARCH_INDEX_EDITOR   | [optional] 

## Methods

### NewGroupRole

`func NewGroupRole() *GroupRole`

NewGroupRole instantiates a new GroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleWithDefaults

`func NewGroupRoleWithDefaults() *GroupRole`

NewGroupRoleWithDefaults instantiates a new GroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupRole) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupRole) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupRole) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupRole) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetGroupRole

`func (o *GroupRole) GetGroupRole() string`

GetGroupRole returns the GroupRole field if non-nil, zero value otherwise.

### GetGroupRoleOk

`func (o *GroupRole) GetGroupRoleOk() (*string, bool)`

GetGroupRoleOk returns a tuple with the GroupRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRole

`func (o *GroupRole) SetGroupRole(v string)`

SetGroupRole sets GroupRole field to given value.

### HasGroupRole

`func (o *GroupRole) HasGroupRole() bool`

HasGroupRole returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


