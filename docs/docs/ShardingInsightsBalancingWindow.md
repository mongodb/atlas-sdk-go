# ShardingInsightsBalancingWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | Time of day at which the balancing window ends, in 24-hour &#x60;HH:MM&#x60; format as configured on the cluster. Present only when &#x60;status&#x60; is &#x60;CONFIGURED&#x60;. | [optional] [readonly] 
**StartTime** | Pointer to **string** | Time of day at which the balancing window starts, in 24-hour &#x60;HH:MM&#x60; format as configured on the cluster. Present only when &#x60;status&#x60; is &#x60;CONFIGURED&#x60;. | [optional] [readonly] 
**Status** | **string** | State of the balancing window. &#x60;CONFIGURED&#x60; means the cluster restricts balancing to the interval between &#x60;startTime&#x60; and &#x60;endTime&#x60;. &#x60;NOT_CONFIGURED&#x60; means the balancer may run at any time. &#x60;FAILED_INITIAL_METRIC_PENDING&#x60; means MongoDB Cloud has not yet collected balancer metrics for a recently created or converted cluster. &#x60;FAILED_TO_DETERMINE&#x60; means the window could not be resolved. | [readonly] 

## Methods

### NewShardingInsightsBalancingWindow

`func NewShardingInsightsBalancingWindow(status string, ) *ShardingInsightsBalancingWindow`

NewShardingInsightsBalancingWindow instantiates a new ShardingInsightsBalancingWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingInsightsBalancingWindowWithDefaults

`func NewShardingInsightsBalancingWindowWithDefaults() *ShardingInsightsBalancingWindow`

NewShardingInsightsBalancingWindowWithDefaults instantiates a new ShardingInsightsBalancingWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ShardingInsightsBalancingWindow) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ShardingInsightsBalancingWindow) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ShardingInsightsBalancingWindow) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ShardingInsightsBalancingWindow) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ShardingInsightsBalancingWindow) SetEndTimeNil()`

SetEndTimeNil sets EndTime to an explicit JSON null when marshaled, overriding any value previously set with SetEndTime. Calling SetEndTime again clears the null override.

### GetStartTime

`func (o *ShardingInsightsBalancingWindow) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ShardingInsightsBalancingWindow) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ShardingInsightsBalancingWindow) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ShardingInsightsBalancingWindow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ShardingInsightsBalancingWindow) SetStartTimeNil()`

SetStartTimeNil sets StartTime to an explicit JSON null when marshaled, overriding any value previously set with SetStartTime. Calling SetStartTime again clears the null override.

### GetStatus

`func (o *ShardingInsightsBalancingWindow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShardingInsightsBalancingWindow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShardingInsightsBalancingWindow) SetStatus(v string)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


