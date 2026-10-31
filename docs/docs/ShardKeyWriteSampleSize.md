# ShardKeyWriteSampleSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | Pointer to **int64** | Number of sampled &#x60;delete&#x60; operations. | [optional] [readonly] 
**FindAndModify** | Pointer to **int64** | Number of sampled &#x60;findAndModify&#x60; operations. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total number of sampled write operations. | [optional] [readonly] 
**Update** | Pointer to **int64** | Number of sampled &#x60;update&#x60; operations. | [optional] [readonly] 

## Methods

### NewShardKeyWriteSampleSize

`func NewShardKeyWriteSampleSize() *ShardKeyWriteSampleSize`

NewShardKeyWriteSampleSize instantiates a new ShardKeyWriteSampleSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyWriteSampleSizeWithDefaults

`func NewShardKeyWriteSampleSizeWithDefaults() *ShardKeyWriteSampleSize`

NewShardKeyWriteSampleSizeWithDefaults instantiates a new ShardKeyWriteSampleSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *ShardKeyWriteSampleSize) GetDelete() int64`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ShardKeyWriteSampleSize) GetDeleteOk() (*int64, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ShardKeyWriteSampleSize) SetDelete(v int64)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ShardKeyWriteSampleSize) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *ShardKeyWriteSampleSize) SetDeleteNil()`

SetDeleteNil sets Delete to an explicit JSON null when marshaled, overriding any value previously set with SetDelete. Calling SetDelete again clears the null override.

### GetFindAndModify

`func (o *ShardKeyWriteSampleSize) GetFindAndModify() int64`

GetFindAndModify returns the FindAndModify field if non-nil, zero value otherwise.

### GetFindAndModifyOk

`func (o *ShardKeyWriteSampleSize) GetFindAndModifyOk() (*int64, bool)`

GetFindAndModifyOk returns a tuple with the FindAndModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindAndModify

`func (o *ShardKeyWriteSampleSize) SetFindAndModify(v int64)`

SetFindAndModify sets FindAndModify field to given value.

### HasFindAndModify

`func (o *ShardKeyWriteSampleSize) HasFindAndModify() bool`

HasFindAndModify returns a boolean if a field has been set.

### SetFindAndModifyNil

`func (o *ShardKeyWriteSampleSize) SetFindAndModifyNil()`

SetFindAndModifyNil sets FindAndModify to an explicit JSON null when marshaled, overriding any value previously set with SetFindAndModify. Calling SetFindAndModify again clears the null override.

### GetTotal

`func (o *ShardKeyWriteSampleSize) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ShardKeyWriteSampleSize) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ShardKeyWriteSampleSize) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ShardKeyWriteSampleSize) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ShardKeyWriteSampleSize) SetTotalNil()`

SetTotalNil sets Total to an explicit JSON null when marshaled, overriding any value previously set with SetTotal. Calling SetTotal again clears the null override.

### GetUpdate

`func (o *ShardKeyWriteSampleSize) GetUpdate() int64`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ShardKeyWriteSampleSize) GetUpdateOk() (*int64, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ShardKeyWriteSampleSize) SetUpdate(v int64)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ShardKeyWriteSampleSize) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *ShardKeyWriteSampleSize) SetUpdateNil()`

SetUpdateNil sets Update to an explicit JSON null when marshaled, overriding any value previously set with SetUpdate. Calling SetUpdate again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


