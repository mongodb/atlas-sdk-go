# ResourceTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Constant that defines the set of the tag. For example, &#x60;environment&#x60; in the &#x60;environment : production&#x60; tag. | 
**Value** | **string** | Variable that belongs to the set of the tag. For example, &#x60;production&#x60; in the &#x60;environment : production&#x60; tag. | 

## Methods

### NewResourceTag

`func NewResourceTag(key string, value string, ) *ResourceTag`

NewResourceTag instantiates a new ResourceTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTagWithDefaults

`func NewResourceTagWithDefaults() *ResourceTag`

NewResourceTagWithDefaults instantiates a new ResourceTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ResourceTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResourceTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResourceTag) SetKey(v string)`

SetKey sets Key field to given value.

### GetValue

`func (o *ResourceTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceTag) SetValue(v string)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


