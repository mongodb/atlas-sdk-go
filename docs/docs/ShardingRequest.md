# ShardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateSupportingIndexes** | **bool** | Flag that lets the migration create supporting indexes for the shard keys, if none exists, as the destination cluster also needs compatible indexes for the specified shard keys. | 
**ShardingEntries** | Pointer to [**[]ShardEntry**](ShardEntry.md) | List of shard configurations to shard destination collections. Atlas shards only those collections that you include in the sharding entries array. | [optional] 

## Methods

### NewShardingRequest

`func NewShardingRequest(createSupportingIndexes bool, ) *ShardingRequest`

NewShardingRequest instantiates a new ShardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingRequestWithDefaults

`func NewShardingRequestWithDefaults() *ShardingRequest`

NewShardingRequestWithDefaults instantiates a new ShardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateSupportingIndexes

`func (o *ShardingRequest) GetCreateSupportingIndexes() bool`

GetCreateSupportingIndexes returns the CreateSupportingIndexes field if non-nil, zero value otherwise.

### GetCreateSupportingIndexesOk

`func (o *ShardingRequest) GetCreateSupportingIndexesOk() (*bool, bool)`

GetCreateSupportingIndexesOk returns a tuple with the CreateSupportingIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSupportingIndexes

`func (o *ShardingRequest) SetCreateSupportingIndexes(v bool)`

SetCreateSupportingIndexes sets CreateSupportingIndexes field to given value.

### GetShardingEntries

`func (o *ShardingRequest) GetShardingEntries() []ShardEntry`

GetShardingEntries returns the ShardingEntries field if non-nil, zero value otherwise.

### GetShardingEntriesOk

`func (o *ShardingRequest) GetShardingEntriesOk() (*[]ShardEntry, bool)`

GetShardingEntriesOk returns a tuple with the ShardingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardingEntries

`func (o *ShardingRequest) SetShardingEntries(v []ShardEntry)`

SetShardingEntries sets ShardingEntries field to given value.

### HasShardingEntries

`func (o *ShardingRequest) HasShardingEntries() bool`

HasShardingEntries returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


