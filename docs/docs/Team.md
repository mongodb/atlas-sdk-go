# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this team. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies the team. | 
**Usernames** | Pointer to **[]string** | List that contains the MongoDB Cloud users in this team. | [optional] 

## Methods

### NewTeam

`func NewTeam(name string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *Team) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Team) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Team) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Team) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetUsernames

`func (o *Team) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *Team) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *Team) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.

### HasUsernames

`func (o *Team) HasUsernames() bool`

HasUsernames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


