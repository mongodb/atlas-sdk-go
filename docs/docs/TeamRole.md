# TeamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RoleNames** | Pointer to **[]string** | One or more project-level roles to assign to the team. | [optional] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the team. | [optional] 

## Methods

### NewTeamRole

`func NewTeamRole() *TeamRole`

NewTeamRole instantiates a new TeamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRoleWithDefaults

`func NewTeamRoleWithDefaults() *TeamRole`

NewTeamRoleWithDefaults instantiates a new TeamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TeamRole) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamRole) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamRole) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRoleNames

`func (o *TeamRole) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *TeamRole) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *TeamRole) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *TeamRole) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.
### GetTeamId

`func (o *TeamRole) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamRole) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamRole) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamRole) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


