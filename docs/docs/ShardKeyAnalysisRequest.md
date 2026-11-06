# ShardKeyAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeKeyCharacteristics** | Pointer to **bool** | Flag that indicates whether to return the cardinality, frequency and monotonicity of the candidate shard key. | [optional] [default to true]
**IncludeReadWriteDistribution** | Pointer to **bool** | Flag that indicates whether to return how reads and writes would be routed across shards under the candidate shard key. This requires query sampling to have collected operations for the namespace. | [optional] [default to false]
**Namespace** | **string** | Combination of the database and collection to analyze, written as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | 
**ShardKey** | **[]map[string]string** | Candidate shard key, given as one object per shard key field in the order the fields form the key. Field order is significant: &#x60;[{\&quot;a\&quot;: \&quot;1\&quot;}, {\&quot;b\&quot;: \&quot;1\&quot;}]&#x60; and &#x60;[{\&quot;b\&quot;: \&quot;1\&quot;}, {\&quot;a\&quot;: \&quot;1\&quot;}]&#x60; are different shard keys. At most one field may be &#x60;hashed&#x60;. | 

## Methods

### NewShardKeyAnalysisRequest

`func NewShardKeyAnalysisRequest(namespace string, shardKey []map[string]string, ) *ShardKeyAnalysisRequest`

NewShardKeyAnalysisRequest instantiates a new ShardKeyAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyAnalysisRequestWithDefaults

`func NewShardKeyAnalysisRequestWithDefaults() *ShardKeyAnalysisRequest`

NewShardKeyAnalysisRequestWithDefaults instantiates a new ShardKeyAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeKeyCharacteristics

`func (o *ShardKeyAnalysisRequest) GetIncludeKeyCharacteristics() bool`

GetIncludeKeyCharacteristics returns the IncludeKeyCharacteristics field if non-nil, zero value otherwise.

### GetIncludeKeyCharacteristicsOk

`func (o *ShardKeyAnalysisRequest) GetIncludeKeyCharacteristicsOk() (*bool, bool)`

GetIncludeKeyCharacteristicsOk returns a tuple with the IncludeKeyCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeKeyCharacteristics

`func (o *ShardKeyAnalysisRequest) SetIncludeKeyCharacteristics(v bool)`

SetIncludeKeyCharacteristics sets IncludeKeyCharacteristics field to given value.

### HasIncludeKeyCharacteristics

`func (o *ShardKeyAnalysisRequest) HasIncludeKeyCharacteristics() bool`

HasIncludeKeyCharacteristics returns a boolean if a field has been set.

### SetIncludeKeyCharacteristicsNil

`func (o *ShardKeyAnalysisRequest) SetIncludeKeyCharacteristicsNil()`

SetIncludeKeyCharacteristicsNil sets IncludeKeyCharacteristics to an explicit JSON null when marshaled, overriding any value previously set with SetIncludeKeyCharacteristics. Calling SetIncludeKeyCharacteristics again clears the null override.

### GetIncludeReadWriteDistribution

`func (o *ShardKeyAnalysisRequest) GetIncludeReadWriteDistribution() bool`

GetIncludeReadWriteDistribution returns the IncludeReadWriteDistribution field if non-nil, zero value otherwise.

### GetIncludeReadWriteDistributionOk

`func (o *ShardKeyAnalysisRequest) GetIncludeReadWriteDistributionOk() (*bool, bool)`

GetIncludeReadWriteDistributionOk returns a tuple with the IncludeReadWriteDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReadWriteDistribution

`func (o *ShardKeyAnalysisRequest) SetIncludeReadWriteDistribution(v bool)`

SetIncludeReadWriteDistribution sets IncludeReadWriteDistribution field to given value.

### HasIncludeReadWriteDistribution

`func (o *ShardKeyAnalysisRequest) HasIncludeReadWriteDistribution() bool`

HasIncludeReadWriteDistribution returns a boolean if a field has been set.

### SetIncludeReadWriteDistributionNil

`func (o *ShardKeyAnalysisRequest) SetIncludeReadWriteDistributionNil()`

SetIncludeReadWriteDistributionNil sets IncludeReadWriteDistribution to an explicit JSON null when marshaled, overriding any value previously set with SetIncludeReadWriteDistribution. Calling SetIncludeReadWriteDistribution again clears the null override.

### GetNamespace

`func (o *ShardKeyAnalysisRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ShardKeyAnalysisRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ShardKeyAnalysisRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### GetShardKey

`func (o *ShardKeyAnalysisRequest) GetShardKey() []map[string]string`

GetShardKey returns the ShardKey field if non-nil, zero value otherwise.

### GetShardKeyOk

`func (o *ShardKeyAnalysisRequest) GetShardKeyOk() (*[]map[string]string, bool)`

GetShardKeyOk returns a tuple with the ShardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardKey

`func (o *ShardKeyAnalysisRequest) SetShardKey(v []map[string]string)`

SetShardKey sets ShardKey field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


