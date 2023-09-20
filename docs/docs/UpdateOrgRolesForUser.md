# UpdateOrgRolesForUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OrgRoles** | Pointer to **[]string** | One or more organization level roles to assign to the MongoDB Cloud user. | [optional] 

## Methods

### NewUpdateOrgRolesForUser

`func NewUpdateOrgRolesForUser() *UpdateOrgRolesForUser`

NewUpdateOrgRolesForUser instantiates a new UpdateOrgRolesForUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgRolesForUserWithDefaults

`func NewUpdateOrgRolesForUserWithDefaults() *UpdateOrgRolesForUser`

NewUpdateOrgRolesForUserWithDefaults instantiates a new UpdateOrgRolesForUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UpdateOrgRolesForUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateOrgRolesForUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateOrgRolesForUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateOrgRolesForUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOrgRoles

`func (o *UpdateOrgRolesForUser) GetOrgRoles() []string`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *UpdateOrgRolesForUser) GetOrgRolesOk() (*[]string, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *UpdateOrgRolesForUser) SetOrgRoles(v []string)`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *UpdateOrgRolesForUser) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


