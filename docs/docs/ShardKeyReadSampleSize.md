# ShardKeyReadSampleSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to **int64** | Number of sampled &#x60;aggregate&#x60; operations. | [optional] [readonly] 
**Count** | Pointer to **int64** | Number of sampled &#x60;count&#x60; operations. | [optional] [readonly] 
**Distinct** | Pointer to **int64** | Number of sampled &#x60;distinct&#x60; operations. | [optional] [readonly] 
**Find** | Pointer to **int64** | Number of sampled &#x60;find&#x60; operations. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total number of sampled read operations. | [optional] [readonly] 

## Methods

### NewShardKeyReadSampleSize

`func NewShardKeyReadSampleSize() *ShardKeyReadSampleSize`

NewShardKeyReadSampleSize instantiates a new ShardKeyReadSampleSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyReadSampleSizeWithDefaults

`func NewShardKeyReadSampleSizeWithDefaults() *ShardKeyReadSampleSize`

NewShardKeyReadSampleSizeWithDefaults instantiates a new ShardKeyReadSampleSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *ShardKeyReadSampleSize) GetAggregate() int64`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *ShardKeyReadSampleSize) GetAggregateOk() (*int64, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *ShardKeyReadSampleSize) SetAggregate(v int64)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *ShardKeyReadSampleSize) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### SetAggregateNil

`func (o *ShardKeyReadSampleSize) SetAggregateNil()`

SetAggregateNil sets Aggregate to an explicit JSON null when marshaled, overriding any value previously set with SetAggregate. Calling SetAggregate again clears the null override.

### GetCount

`func (o *ShardKeyReadSampleSize) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ShardKeyReadSampleSize) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ShardKeyReadSampleSize) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ShardKeyReadSampleSize) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ShardKeyReadSampleSize) SetCountNil()`

SetCountNil sets Count to an explicit JSON null when marshaled, overriding any value previously set with SetCount. Calling SetCount again clears the null override.

### GetDistinct

`func (o *ShardKeyReadSampleSize) GetDistinct() int64`

GetDistinct returns the Distinct field if non-nil, zero value otherwise.

### GetDistinctOk

`func (o *ShardKeyReadSampleSize) GetDistinctOk() (*int64, bool)`

GetDistinctOk returns a tuple with the Distinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinct

`func (o *ShardKeyReadSampleSize) SetDistinct(v int64)`

SetDistinct sets Distinct field to given value.

### HasDistinct

`func (o *ShardKeyReadSampleSize) HasDistinct() bool`

HasDistinct returns a boolean if a field has been set.

### SetDistinctNil

`func (o *ShardKeyReadSampleSize) SetDistinctNil()`

SetDistinctNil sets Distinct to an explicit JSON null when marshaled, overriding any value previously set with SetDistinct. Calling SetDistinct again clears the null override.

### GetFind

`func (o *ShardKeyReadSampleSize) GetFind() int64`

GetFind returns the Find field if non-nil, zero value otherwise.

### GetFindOk

`func (o *ShardKeyReadSampleSize) GetFindOk() (*int64, bool)`

GetFindOk returns a tuple with the Find field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFind

`func (o *ShardKeyReadSampleSize) SetFind(v int64)`

SetFind sets Find field to given value.

### HasFind

`func (o *ShardKeyReadSampleSize) HasFind() bool`

HasFind returns a boolean if a field has been set.

### SetFindNil

`func (o *ShardKeyReadSampleSize) SetFindNil()`

SetFindNil sets Find to an explicit JSON null when marshaled, overriding any value previously set with SetFind. Calling SetFind again clears the null override.

### GetTotal

`func (o *ShardKeyReadSampleSize) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ShardKeyReadSampleSize) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ShardKeyReadSampleSize) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ShardKeyReadSampleSize) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ShardKeyReadSampleSize) SetTotalNil()`

SetTotalNil sets Total to an explicit JSON null when marshaled, overriding any value previously set with SetTotal. Calling SetTotal again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


