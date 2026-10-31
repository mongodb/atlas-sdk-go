# ShardCpuUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AveragePercent** | **float64** | Average CPU utilization across the shard&#39;s nodes over the last hour. | [readonly] 
**MaximumPercent** | **float64** | Highest CPU utilization that any one of the shard&#39;s nodes reached over the last hour. | [readonly] 

## Methods

### NewShardCpuUtilization

`func NewShardCpuUtilization(averagePercent float64, maximumPercent float64, ) *ShardCpuUtilization`

NewShardCpuUtilization instantiates a new ShardCpuUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardCpuUtilizationWithDefaults

`func NewShardCpuUtilizationWithDefaults() *ShardCpuUtilization`

NewShardCpuUtilizationWithDefaults instantiates a new ShardCpuUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAveragePercent

`func (o *ShardCpuUtilization) GetAveragePercent() float64`

GetAveragePercent returns the AveragePercent field if non-nil, zero value otherwise.

### GetAveragePercentOk

`func (o *ShardCpuUtilization) GetAveragePercentOk() (*float64, bool)`

GetAveragePercentOk returns a tuple with the AveragePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePercent

`func (o *ShardCpuUtilization) SetAveragePercent(v float64)`

SetAveragePercent sets AveragePercent field to given value.

### GetMaximumPercent

`func (o *ShardCpuUtilization) GetMaximumPercent() float64`

GetMaximumPercent returns the MaximumPercent field if non-nil, zero value otherwise.

### GetMaximumPercentOk

`func (o *ShardCpuUtilization) GetMaximumPercentOk() (*float64, bool)`

GetMaximumPercentOk returns a tuple with the MaximumPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPercent

`func (o *ShardCpuUtilization) SetMaximumPercent(v float64)`

SetMaximumPercent sets MaximumPercent field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


