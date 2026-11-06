# ShardKeyMostCommonValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **int64** | Number of sampled documents that carry this shard key value. | [optional] [readonly] 
**Truncated** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud omitted &#x60;value&#x60; because it was too large to return. When this flag is &#x60;true&#x60;, &#x60;frequency&#x60; still describes a real shard key value even though the value itself is absent. | [optional] [readonly] 
**Value** | Pointer to [**any**](interface{}.md) | Shard key value, given as one field-value pair per shard key field. This parameter is absent when &#x60;truncated&#x60; is &#x60;true&#x60;. | [optional] [readonly] 

## Methods

### NewShardKeyMostCommonValue

`func NewShardKeyMostCommonValue() *ShardKeyMostCommonValue`

NewShardKeyMostCommonValue instantiates a new ShardKeyMostCommonValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyMostCommonValueWithDefaults

`func NewShardKeyMostCommonValueWithDefaults() *ShardKeyMostCommonValue`

NewShardKeyMostCommonValueWithDefaults instantiates a new ShardKeyMostCommonValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *ShardKeyMostCommonValue) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ShardKeyMostCommonValue) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ShardKeyMostCommonValue) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ShardKeyMostCommonValue) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *ShardKeyMostCommonValue) SetFrequencyNil()`

SetFrequencyNil sets Frequency to an explicit JSON null when marshaled, overriding any value previously set with SetFrequency. Calling SetFrequency again clears the null override.

### GetTruncated

`func (o *ShardKeyMostCommonValue) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ShardKeyMostCommonValue) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ShardKeyMostCommonValue) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ShardKeyMostCommonValue) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### SetTruncatedNil

`func (o *ShardKeyMostCommonValue) SetTruncatedNil()`

SetTruncatedNil sets Truncated to an explicit JSON null when marshaled, overriding any value previously set with SetTruncated. Calling SetTruncated again clears the null override.

### GetValue

`func (o *ShardKeyMostCommonValue) GetValue() any`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShardKeyMostCommonValue) GetValueOk() (*any, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShardKeyMostCommonValue) SetValue(v any)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShardKeyMostCommonValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ShardKeyMostCommonValue) SetValueNil()`

SetValueNil sets Value to an explicit JSON null when marshaled, overriding any value previously set with SetValue. Calling SetValue again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


