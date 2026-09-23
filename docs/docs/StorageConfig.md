# StorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardSizeLimitGB** | Pointer to **int** | Available in Public Preview: Maximum data size that MongoDB Cloud allows each shard of this cluster to reach, expressed in gigabytes. MongoDB Cloud rejects writes to a shard that reaches the limit that it enforces. In &#x60;replicationSpecs&#x60;, this field reports the limit that you configured, and MongoDB Cloud omits it when you never configured one. In &#x60;effectiveReplicationSpecs&#x60;, this field reports the limit that MongoDB Cloud enforces: usually the limit that you configured, otherwise the default limit that MongoDB Cloud assigns when it creates or updates the cluster. This value may differ from the limit that you configured due to system-managed changes. This limit applies to every shard of the cluster; set the same value on each region configuration&#39;s &#x60;autoScaling&#x60;, as MongoDB Cloud rejects requests that specify differing values. You can set this only on Atlas INFINITE clusters: MongoDB Cloud rejects any request that names this field, including as &#x60;null&#x60;, for a cluster or node type that doesn&#39;t support it. In a request that includes &#x60;replicationSpecs&#x60;, omitting &#x60;shardSizeLimitGB&#x60; or sending it as &#x60;null&#x60; clears the limit. Omitting &#x60;replicationSpecs&#x60; preserves it. | [optional] 

## Methods

### NewStorageConfig

`func NewStorageConfig() *StorageConfig`

NewStorageConfig instantiates a new StorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigWithDefaults

`func NewStorageConfigWithDefaults() *StorageConfig`

NewStorageConfigWithDefaults instantiates a new StorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardSizeLimitGB

`func (o *StorageConfig) GetShardSizeLimitGB() int`

GetShardSizeLimitGB returns the ShardSizeLimitGB field if non-nil, zero value otherwise.

### GetShardSizeLimitGBOk

`func (o *StorageConfig) GetShardSizeLimitGBOk() (*int, bool)`

GetShardSizeLimitGBOk returns a tuple with the ShardSizeLimitGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardSizeLimitGB

`func (o *StorageConfig) SetShardSizeLimitGB(v int)`

SetShardSizeLimitGB sets ShardSizeLimitGB field to given value.

### HasShardSizeLimitGB

`func (o *StorageConfig) HasShardSizeLimitGB() bool`

HasShardSizeLimitGB returns a boolean if a field has been set.

### SetShardSizeLimitGBNil

`func (o *StorageConfig) SetShardSizeLimitGBNil()`

SetShardSizeLimitGBNil sets ShardSizeLimitGB to an explicit JSON null when marshaled, overriding any value previously set with SetShardSizeLimitGB. Calling SetShardSizeLimitGB again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


