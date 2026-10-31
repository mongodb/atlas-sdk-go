# ShardingRecommendationTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Identifier of the monitored condition that fired. | [readonly] 
**Target** | Pointer to **string** | Entity the condition was observed on: a namespace in &#x60;&lt;database&gt;.&lt;collection&gt;&#x60; form, a host name, or a replica set (shard) name, as identified by &#x60;targetType&#x60;. Present if and only if &#x60;targetType&#x60; is present; both are absent for cluster-wide conditions, such as &#x60;UNSHARDED_COLLECTION_COUNT&#x60;, which are not tied to a single entity. The cluster is already identified in the request path. | [optional] [readonly] 
**TargetType** | Pointer to **string** | Kind of entity that &#x60;target&#x60; identifies. Present if and only if &#x60;target&#x60; is present; both are absent for cluster-wide conditions. | [optional] [readonly] 
**Unit** | **string** | Unit of the measured trigger value. | [readonly] 
**Value** | Pointer to **float64** | Measured value of the trigger metric for this entity. This parameter is absent when no finite measurement was available. | [optional] [readonly] 

## Methods

### NewShardingRecommendationTrigger

`func NewShardingRecommendationTrigger(name string, unit string, ) *ShardingRecommendationTrigger`

NewShardingRecommendationTrigger instantiates a new ShardingRecommendationTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingRecommendationTriggerWithDefaults

`func NewShardingRecommendationTriggerWithDefaults() *ShardingRecommendationTrigger`

NewShardingRecommendationTriggerWithDefaults instantiates a new ShardingRecommendationTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShardingRecommendationTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShardingRecommendationTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShardingRecommendationTrigger) SetName(v string)`

SetName sets Name field to given value.

### GetTarget

`func (o *ShardingRecommendationTrigger) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ShardingRecommendationTrigger) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ShardingRecommendationTrigger) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ShardingRecommendationTrigger) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ShardingRecommendationTrigger) SetTargetNil()`

SetTargetNil sets Target to an explicit JSON null when marshaled, overriding any value previously set with SetTarget. Calling SetTarget again clears the null override.

### GetTargetType

`func (o *ShardingRecommendationTrigger) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ShardingRecommendationTrigger) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ShardingRecommendationTrigger) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ShardingRecommendationTrigger) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ShardingRecommendationTrigger) SetTargetTypeNil()`

SetTargetTypeNil sets TargetType to an explicit JSON null when marshaled, overriding any value previously set with SetTargetType. Calling SetTargetType again clears the null override.

### GetUnit

`func (o *ShardingRecommendationTrigger) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ShardingRecommendationTrigger) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ShardingRecommendationTrigger) SetUnit(v string)`

SetUnit sets Unit field to given value.

### GetValue

`func (o *ShardingRecommendationTrigger) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShardingRecommendationTrigger) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShardingRecommendationTrigger) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShardingRecommendationTrigger) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ShardingRecommendationTrigger) SetValueNil()`

SetValueNil sets Value to an explicit JSON null when marshaled, overriding any value previously set with SetValue. Calling SetValue again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


