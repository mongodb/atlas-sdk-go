# CharFilterhtmlStrip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredTags** | Pointer to **[]string** | The HTML tags that you want to exclude from filtering. | [optional] 
**Type** | **string** | Human-readable label that identifies this character filter type. | 

## Methods

### NewCharFilterhtmlStrip

`func NewCharFilterhtmlStrip(type_ string, ) *CharFilterhtmlStrip`

NewCharFilterhtmlStrip instantiates a new CharFilterhtmlStrip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharFilterhtmlStripWithDefaults

`func NewCharFilterhtmlStripWithDefaults() *CharFilterhtmlStrip`

NewCharFilterhtmlStripWithDefaults instantiates a new CharFilterhtmlStrip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoredTags

`func (o *CharFilterhtmlStrip) GetIgnoredTags() []string`

GetIgnoredTags returns the IgnoredTags field if non-nil, zero value otherwise.

### GetIgnoredTagsOk

`func (o *CharFilterhtmlStrip) GetIgnoredTagsOk() (*[]string, bool)`

GetIgnoredTagsOk returns a tuple with the IgnoredTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredTags

`func (o *CharFilterhtmlStrip) SetIgnoredTags(v []string)`

SetIgnoredTags sets IgnoredTags field to given value.

### HasIgnoredTags

`func (o *CharFilterhtmlStrip) HasIgnoredTags() bool`

HasIgnoredTags returns a boolean if a field has been set.

### GetType

`func (o *CharFilterhtmlStrip) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharFilterhtmlStrip) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharFilterhtmlStrip) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


