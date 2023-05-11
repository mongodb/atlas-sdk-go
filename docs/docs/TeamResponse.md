# TeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this team. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the team. | [optional] 

## Methods

### NewTeamResponse

`func NewTeamResponse() *TeamResponse`

NewTeamResponse instantiates a new TeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamResponseWithDefaults

`func NewTeamResponseWithDefaults() *TeamResponse`

NewTeamResponseWithDefaults instantiates a new TeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *TeamResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TeamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


