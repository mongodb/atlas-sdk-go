# GroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable label that identifies the project included in the MongoDB Cloud organization. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the project. | [optional] 

## Methods

### NewGroupUpdate

`func NewGroupUpdate() *GroupUpdate`

NewGroupUpdate instantiates a new GroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateWithDefaults

`func NewGroupUpdateWithDefaults() *GroupUpdate`

NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.
### GetTags

`func (o *GroupUpdate) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GroupUpdate) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GroupUpdate) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GroupUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


