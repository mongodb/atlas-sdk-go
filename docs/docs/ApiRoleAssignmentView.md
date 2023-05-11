# ApiRoleAssignmentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to which this role belongs. You can set a value for this parameter or **orgId** but not both in the same request. | [optional] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which this role belongs. You can set a value for this parameter or **groupId** but not both in the same request. | [optional] 
**RoleName** | Pointer to **string** | Human-readable label that identifies the collection of privileges that MongoDB Cloud grants a specific API key, MongoDB Cloud user, or MongoDB Cloud team. These roles include organization- and project-level roles.  Organization Roles  * ORG_OWNER * ORG_MEMBER * ORG_GROUP_CREATOR * ORG_BILLING_ADMIN * ORG_READ_ONLY  Project Roles  * GROUP_CLUSTER_MANAGER * GROUP_DATA_ACCESS_ADMIN * GROUP_DATA_ACCESS_READ_ONLY * GROUP_DATA_ACCESS_READ_WRITE * GROUP_OWNER * GROUP_READ_ONLY * GROUP_SEARCH_INDEX_EDITOR   | [optional] 

## Methods

### NewApiRoleAssignmentView

`func NewApiRoleAssignmentView() *ApiRoleAssignmentView`

NewApiRoleAssignmentView instantiates a new ApiRoleAssignmentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleAssignmentViewWithDefaults

`func NewApiRoleAssignmentViewWithDefaults() *ApiRoleAssignmentView`

NewApiRoleAssignmentViewWithDefaults instantiates a new ApiRoleAssignmentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ApiRoleAssignmentView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiRoleAssignmentView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiRoleAssignmentView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiRoleAssignmentView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetOrgId

`func (o *ApiRoleAssignmentView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiRoleAssignmentView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiRoleAssignmentView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiRoleAssignmentView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRoleName

`func (o *ApiRoleAssignmentView) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ApiRoleAssignmentView) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ApiRoleAssignmentView) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *ApiRoleAssignmentView) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


