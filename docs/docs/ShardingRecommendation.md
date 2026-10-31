# ShardingRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | [**[]ShardingRecommendationTrigger**](ShardingRecommendationTrigger.md) | Monitored conditions that fired and produced this recommendation. | [readonly] 
**Type** | **string** | Recommendation type that produced this recommendation. | [readonly] 
**RecommendedShardCount** | Pointer to **int** | Recommended total number of shards for the cluster. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Namespace of the collection to shard, in &#x60;&lt;database&gt;.&lt;collection&gt;&#x60; form. | [optional] [readonly] 

## Methods

### NewShardingRecommendation

`func NewShardingRecommendation(triggers []ShardingRecommendationTrigger, type_ string, ) *ShardingRecommendation`

NewShardingRecommendation instantiates a new ShardingRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingRecommendationWithDefaults

`func NewShardingRecommendationWithDefaults() *ShardingRecommendation`

NewShardingRecommendationWithDefaults instantiates a new ShardingRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *ShardingRecommendation) GetTriggers() []ShardingRecommendationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ShardingRecommendation) GetTriggersOk() (*[]ShardingRecommendationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ShardingRecommendation) SetTriggers(v []ShardingRecommendationTrigger)`

SetTriggers sets Triggers field to given value.

### GetType

`func (o *ShardingRecommendation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShardingRecommendation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShardingRecommendation) SetType(v string)`

SetType sets Type field to given value.

### GetRecommendedShardCount

`func (o *ShardingRecommendation) GetRecommendedShardCount() int`

GetRecommendedShardCount returns the RecommendedShardCount field if non-nil, zero value otherwise.

### GetRecommendedShardCountOk

`func (o *ShardingRecommendation) GetRecommendedShardCountOk() (*int, bool)`

GetRecommendedShardCountOk returns a tuple with the RecommendedShardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedShardCount

`func (o *ShardingRecommendation) SetRecommendedShardCount(v int)`

SetRecommendedShardCount sets RecommendedShardCount field to given value.

### HasRecommendedShardCount

`func (o *ShardingRecommendation) HasRecommendedShardCount() bool`

HasRecommendedShardCount returns a boolean if a field has been set.

### SetRecommendedShardCountNil

`func (o *ShardingRecommendation) SetRecommendedShardCountNil()`

SetRecommendedShardCountNil sets RecommendedShardCount to an explicit JSON null when marshaled, overriding any value previously set with SetRecommendedShardCount. Calling SetRecommendedShardCount again clears the null override.

### GetNamespace

`func (o *ShardingRecommendation) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ShardingRecommendation) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ShardingRecommendation) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ShardingRecommendation) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ShardingRecommendation) SetNamespaceNil()`

SetNamespaceNil sets Namespace to an explicit JSON null when marshaled, overriding any value previously set with SetNamespace. Calling SetNamespace again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


