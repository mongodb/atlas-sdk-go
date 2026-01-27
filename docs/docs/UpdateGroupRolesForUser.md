# UpdateGroupRolesForUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupRoles** | Pointer to **[]string** | One or more project-level roles to assign to the MongoDB Cloud user. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewUpdateGroupRolesForUser

`func NewUpdateGroupRolesForUser() *UpdateGroupRolesForUser`

NewUpdateGroupRolesForUser instantiates a new UpdateGroupRolesForUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupRolesForUserWithDefaults

`func NewUpdateGroupRolesForUserWithDefaults() *UpdateGroupRolesForUser`

NewUpdateGroupRolesForUserWithDefaults instantiates a new UpdateGroupRolesForUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupRoles

`func (o *UpdateGroupRolesForUser) GetGroupRoles() []string`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### GetGroupRolesOk

`func (o *UpdateGroupRolesForUser) GetGroupRolesOk() (*[]string, bool)`

GetGroupRolesOk returns a tuple with the GroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoles

`func (o *UpdateGroupRolesForUser) SetGroupRoles(v []string)`

SetGroupRoles sets GroupRoles field to given value.

### HasGroupRoles

`func (o *UpdateGroupRolesForUser) HasGroupRoles() bool`

HasGroupRoles returns a boolean if a field has been set.
### GetLinks

`func (o *UpdateGroupRolesForUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateGroupRolesForUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateGroupRolesForUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateGroupRolesForUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


