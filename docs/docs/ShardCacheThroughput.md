# ShardCacheThroughput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadIntoBytesPerSecond** | **float64** | Average bytes per second read into the cache, across the shard&#39;s nodes over the last hour. | [readonly] 
**WrittenFromBytesPerSecond** | **float64** | Average bytes per second written from the cache, across the shard&#39;s nodes over the last hour. | [readonly] 

## Methods

### NewShardCacheThroughput

`func NewShardCacheThroughput(readIntoBytesPerSecond float64, writtenFromBytesPerSecond float64, ) *ShardCacheThroughput`

NewShardCacheThroughput instantiates a new ShardCacheThroughput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardCacheThroughputWithDefaults

`func NewShardCacheThroughputWithDefaults() *ShardCacheThroughput`

NewShardCacheThroughputWithDefaults instantiates a new ShardCacheThroughput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadIntoBytesPerSecond

`func (o *ShardCacheThroughput) GetReadIntoBytesPerSecond() float64`

GetReadIntoBytesPerSecond returns the ReadIntoBytesPerSecond field if non-nil, zero value otherwise.

### GetReadIntoBytesPerSecondOk

`func (o *ShardCacheThroughput) GetReadIntoBytesPerSecondOk() (*float64, bool)`

GetReadIntoBytesPerSecondOk returns a tuple with the ReadIntoBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIntoBytesPerSecond

`func (o *ShardCacheThroughput) SetReadIntoBytesPerSecond(v float64)`

SetReadIntoBytesPerSecond sets ReadIntoBytesPerSecond field to given value.

### GetWrittenFromBytesPerSecond

`func (o *ShardCacheThroughput) GetWrittenFromBytesPerSecond() float64`

GetWrittenFromBytesPerSecond returns the WrittenFromBytesPerSecond field if non-nil, zero value otherwise.

### GetWrittenFromBytesPerSecondOk

`func (o *ShardCacheThroughput) GetWrittenFromBytesPerSecondOk() (*float64, bool)`

GetWrittenFromBytesPerSecondOk returns a tuple with the WrittenFromBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenFromBytesPerSecond

`func (o *ShardCacheThroughput) SetWrittenFromBytesPerSecond(v float64)`

SetWrittenFromBytesPerSecond sets WrittenFromBytesPerSecond field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


