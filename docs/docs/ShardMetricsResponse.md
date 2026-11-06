# ShardMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]ShardAlert**](ShardAlert.md) | Alerts open on this shard&#39;s nodes, sorted by creation date with the most recently created first. MongoDB Cloud returns this property when you request a single shard, and returns at most 100 alerts. | [optional] [readonly] 
**CacheThroughput** | Pointer to [**ShardCacheThroughput**](ShardCacheThroughput.md) |  | [optional] 
**CpuUtilization** | [**ShardCpuUtilization**](ShardCpuUtilization.md) |  | 
**DataSizeBytes** | **int64** | Total size of the user data stored on this shard. The resource expresses this value in bytes. | [readonly] 
**RangeDeleterTasks** | Pointer to **float64** | Average number of range deletion tasks queued across the shard&#39;s nodes, measured at the time of the request. | [optional] [readonly] 
**ReadLatencyMicroseconds** | Pointer to **float64** | Average latency of the read operations this shard served over the last hour, in microseconds. | [optional] [readonly] 
**ReadOperationsPerSecond** | [**ShardOperations**](ShardOperations.md) |  | 
**ReplicaSetName** | **string** | Name of the shard&#39;s replica set. This is the name that identifies the same shard on the processes, alerts, events, and backup resources. | [readonly] 
**ShardedCollections** | **int** | Number of sharded collections with data on this shard. | [readonly] 
**Warnings** | **[]string** | Conditions that MongoDB Cloud observed on this shard relative to the rest of the cluster. &#x60;HOTSPOT&#x60; means this shard&#39;s average CPU utilization is more than twice the average across all shards of the cluster. This is a cluster-scoped comparison, so its value depends on how loaded the other shards are and not only on this shard. | [readonly] 
**WriteOperationsPerSecond** | [**ShardOperations**](ShardOperations.md) |  | 

## Methods

### NewShardMetricsResponse

`func NewShardMetricsResponse(cpuUtilization ShardCpuUtilization, dataSizeBytes int64, readOperationsPerSecond ShardOperations, replicaSetName string, shardedCollections int, warnings []string, writeOperationsPerSecond ShardOperations, ) *ShardMetricsResponse`

NewShardMetricsResponse instantiates a new ShardMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardMetricsResponseWithDefaults

`func NewShardMetricsResponseWithDefaults() *ShardMetricsResponse`

NewShardMetricsResponseWithDefaults instantiates a new ShardMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *ShardMetricsResponse) GetAlerts() []ShardAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ShardMetricsResponse) GetAlertsOk() (*[]ShardAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ShardMetricsResponse) SetAlerts(v []ShardAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *ShardMetricsResponse) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### SetAlertsNil

`func (o *ShardMetricsResponse) SetAlertsNil()`

SetAlertsNil sets Alerts to an explicit JSON null when marshaled, overriding any value previously set with SetAlerts. Calling SetAlerts again clears the null override.

### GetCacheThroughput

`func (o *ShardMetricsResponse) GetCacheThroughput() ShardCacheThroughput`

GetCacheThroughput returns the CacheThroughput field if non-nil, zero value otherwise.

### GetCacheThroughputOk

`func (o *ShardMetricsResponse) GetCacheThroughputOk() (*ShardCacheThroughput, bool)`

GetCacheThroughputOk returns a tuple with the CacheThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheThroughput

`func (o *ShardMetricsResponse) SetCacheThroughput(v ShardCacheThroughput)`

SetCacheThroughput sets CacheThroughput field to given value.

### HasCacheThroughput

`func (o *ShardMetricsResponse) HasCacheThroughput() bool`

HasCacheThroughput returns a boolean if a field has been set.

### SetCacheThroughputNil

`func (o *ShardMetricsResponse) SetCacheThroughputNil()`

SetCacheThroughputNil sets CacheThroughput to an explicit JSON null when marshaled, overriding any value previously set with SetCacheThroughput. Calling SetCacheThroughput again clears the null override.

### GetCpuUtilization

`func (o *ShardMetricsResponse) GetCpuUtilization() ShardCpuUtilization`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *ShardMetricsResponse) GetCpuUtilizationOk() (*ShardCpuUtilization, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *ShardMetricsResponse) SetCpuUtilization(v ShardCpuUtilization)`

SetCpuUtilization sets CpuUtilization field to given value.

### GetDataSizeBytes

`func (o *ShardMetricsResponse) GetDataSizeBytes() int64`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *ShardMetricsResponse) GetDataSizeBytesOk() (*int64, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *ShardMetricsResponse) SetDataSizeBytes(v int64)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### GetRangeDeleterTasks

`func (o *ShardMetricsResponse) GetRangeDeleterTasks() float64`

GetRangeDeleterTasks returns the RangeDeleterTasks field if non-nil, zero value otherwise.

### GetRangeDeleterTasksOk

`func (o *ShardMetricsResponse) GetRangeDeleterTasksOk() (*float64, bool)`

GetRangeDeleterTasksOk returns a tuple with the RangeDeleterTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDeleterTasks

`func (o *ShardMetricsResponse) SetRangeDeleterTasks(v float64)`

SetRangeDeleterTasks sets RangeDeleterTasks field to given value.

### HasRangeDeleterTasks

`func (o *ShardMetricsResponse) HasRangeDeleterTasks() bool`

HasRangeDeleterTasks returns a boolean if a field has been set.

### SetRangeDeleterTasksNil

`func (o *ShardMetricsResponse) SetRangeDeleterTasksNil()`

SetRangeDeleterTasksNil sets RangeDeleterTasks to an explicit JSON null when marshaled, overriding any value previously set with SetRangeDeleterTasks. Calling SetRangeDeleterTasks again clears the null override.

### GetReadLatencyMicroseconds

`func (o *ShardMetricsResponse) GetReadLatencyMicroseconds() float64`

GetReadLatencyMicroseconds returns the ReadLatencyMicroseconds field if non-nil, zero value otherwise.

### GetReadLatencyMicrosecondsOk

`func (o *ShardMetricsResponse) GetReadLatencyMicrosecondsOk() (*float64, bool)`

GetReadLatencyMicrosecondsOk returns a tuple with the ReadLatencyMicroseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyMicroseconds

`func (o *ShardMetricsResponse) SetReadLatencyMicroseconds(v float64)`

SetReadLatencyMicroseconds sets ReadLatencyMicroseconds field to given value.

### HasReadLatencyMicroseconds

`func (o *ShardMetricsResponse) HasReadLatencyMicroseconds() bool`

HasReadLatencyMicroseconds returns a boolean if a field has been set.

### SetReadLatencyMicrosecondsNil

`func (o *ShardMetricsResponse) SetReadLatencyMicrosecondsNil()`

SetReadLatencyMicrosecondsNil sets ReadLatencyMicroseconds to an explicit JSON null when marshaled, overriding any value previously set with SetReadLatencyMicroseconds. Calling SetReadLatencyMicroseconds again clears the null override.

### GetReadOperationsPerSecond

`func (o *ShardMetricsResponse) GetReadOperationsPerSecond() ShardOperations`

GetReadOperationsPerSecond returns the ReadOperationsPerSecond field if non-nil, zero value otherwise.

### GetReadOperationsPerSecondOk

`func (o *ShardMetricsResponse) GetReadOperationsPerSecondOk() (*ShardOperations, bool)`

GetReadOperationsPerSecondOk returns a tuple with the ReadOperationsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOperationsPerSecond

`func (o *ShardMetricsResponse) SetReadOperationsPerSecond(v ShardOperations)`

SetReadOperationsPerSecond sets ReadOperationsPerSecond field to given value.

### GetReplicaSetName

`func (o *ShardMetricsResponse) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ShardMetricsResponse) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ShardMetricsResponse) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### GetShardedCollections

`func (o *ShardMetricsResponse) GetShardedCollections() int`

GetShardedCollections returns the ShardedCollections field if non-nil, zero value otherwise.

### GetShardedCollectionsOk

`func (o *ShardMetricsResponse) GetShardedCollectionsOk() (*int, bool)`

GetShardedCollectionsOk returns a tuple with the ShardedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardedCollections

`func (o *ShardMetricsResponse) SetShardedCollections(v int)`

SetShardedCollections sets ShardedCollections field to given value.

### GetWarnings

`func (o *ShardMetricsResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ShardMetricsResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ShardMetricsResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### GetWriteOperationsPerSecond

`func (o *ShardMetricsResponse) GetWriteOperationsPerSecond() ShardOperations`

GetWriteOperationsPerSecond returns the WriteOperationsPerSecond field if non-nil, zero value otherwise.

### GetWriteOperationsPerSecondOk

`func (o *ShardMetricsResponse) GetWriteOperationsPerSecondOk() (*ShardOperations, bool)`

GetWriteOperationsPerSecondOk returns a tuple with the WriteOperationsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOperationsPerSecond

`func (o *ShardMetricsResponse) SetWriteOperationsPerSecond(v ShardOperations)`

SetWriteOperationsPerSecond sets WriteOperationsPerSecond field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


