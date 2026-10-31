# ShardingInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancerState** | **string** | State of the cluster balancer. &#x60;BALANCING&#x60; means a balancer round is in progress. &#x60;IDLE&#x60; means the balancer is enabled and no round is in progress. &#x60;DISABLED&#x60; means the balancer is turned off. &#x60;FAILED_TO_DETERMINE&#x60; means the balancer state could not be resolved. | [readonly] 
**BalancingWindow** | [**ShardingInsightsBalancingWindow**](ShardingInsightsBalancingWindow.md) |  | 
**ConfigServerStatus** | **string** | Availability of the cluster&#39;s config servers. &#x60;AVAILABLE&#x60; means at least one config server reported recently. &#x60;UNAVAILABLE&#x60; means config servers exist but none has reported recently. &#x60;FAILED_TO_DETERMINE&#x60; means availability could not be resolved. | [readonly] 
**TotalShards** | **int** | Number of shards in the cluster. | [readonly] 

## Methods

### NewShardingInsightsResponse

`func NewShardingInsightsResponse(balancerState string, balancingWindow ShardingInsightsBalancingWindow, configServerStatus string, totalShards int, ) *ShardingInsightsResponse`

NewShardingInsightsResponse instantiates a new ShardingInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardingInsightsResponseWithDefaults

`func NewShardingInsightsResponseWithDefaults() *ShardingInsightsResponse`

NewShardingInsightsResponseWithDefaults instantiates a new ShardingInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancerState

`func (o *ShardingInsightsResponse) GetBalancerState() string`

GetBalancerState returns the BalancerState field if non-nil, zero value otherwise.

### GetBalancerStateOk

`func (o *ShardingInsightsResponse) GetBalancerStateOk() (*string, bool)`

GetBalancerStateOk returns a tuple with the BalancerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancerState

`func (o *ShardingInsightsResponse) SetBalancerState(v string)`

SetBalancerState sets BalancerState field to given value.

### GetBalancingWindow

`func (o *ShardingInsightsResponse) GetBalancingWindow() ShardingInsightsBalancingWindow`

GetBalancingWindow returns the BalancingWindow field if non-nil, zero value otherwise.

### GetBalancingWindowOk

`func (o *ShardingInsightsResponse) GetBalancingWindowOk() (*ShardingInsightsBalancingWindow, bool)`

GetBalancingWindowOk returns a tuple with the BalancingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancingWindow

`func (o *ShardingInsightsResponse) SetBalancingWindow(v ShardingInsightsBalancingWindow)`

SetBalancingWindow sets BalancingWindow field to given value.

### GetConfigServerStatus

`func (o *ShardingInsightsResponse) GetConfigServerStatus() string`

GetConfigServerStatus returns the ConfigServerStatus field if non-nil, zero value otherwise.

### GetConfigServerStatusOk

`func (o *ShardingInsightsResponse) GetConfigServerStatusOk() (*string, bool)`

GetConfigServerStatusOk returns a tuple with the ConfigServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigServerStatus

`func (o *ShardingInsightsResponse) SetConfigServerStatus(v string)`

SetConfigServerStatus sets ConfigServerStatus field to given value.

### GetTotalShards

`func (o *ShardingInsightsResponse) GetTotalShards() int`

GetTotalShards returns the TotalShards field if non-nil, zero value otherwise.

### GetTotalShardsOk

`func (o *ShardingInsightsResponse) GetTotalShardsOk() (*int, bool)`

GetTotalShardsOk returns a tuple with the TotalShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShards

`func (o *ShardingInsightsResponse) SetTotalShards(v int)`

SetTotalShards sets TotalShards field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


