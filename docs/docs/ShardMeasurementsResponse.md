# ShardMeasurementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Human-readable label that identifies the cluster. | [readonly] 
**End** | **time.Time** | Date and time at which the window ends. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**Granularity** | **string** | Interval between consecutive data points, in ISO 8601 duration format. Either the requested &#x60;granularity&#x60; or, when none was requested, the resolution MongoDB Cloud chose for the width of the window. | [readonly] 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**Measurements** | [**[]ShardMeasurementSeries**](ShardMeasurementSeries.md) | One time series per requested measurement, in the order of the measurement names. A measurement MongoDB Cloud has no readings for is still listed, with &#x60;null&#x60; data point values. | [readonly] 
**ReplicaSetName** | **string** | Human-readable label that identifies the shard, as the replica set name of its members. | [readonly] 
**Start** | **time.Time** | Date and time of the first data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 

## Methods

### NewShardMeasurementsResponse

`func NewShardMeasurementsResponse(clusterName string, end time.Time, granularity string, groupId string, measurements []ShardMeasurementSeries, replicaSetName string, start time.Time, ) *ShardMeasurementsResponse`

NewShardMeasurementsResponse instantiates a new ShardMeasurementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardMeasurementsResponseWithDefaults

`func NewShardMeasurementsResponseWithDefaults() *ShardMeasurementsResponse`

NewShardMeasurementsResponseWithDefaults instantiates a new ShardMeasurementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ShardMeasurementsResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ShardMeasurementsResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ShardMeasurementsResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### GetEnd

`func (o *ShardMeasurementsResponse) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ShardMeasurementsResponse) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ShardMeasurementsResponse) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### GetGranularity

`func (o *ShardMeasurementsResponse) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ShardMeasurementsResponse) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ShardMeasurementsResponse) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### GetGroupId

`func (o *ShardMeasurementsResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ShardMeasurementsResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ShardMeasurementsResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetMeasurements

`func (o *ShardMeasurementsResponse) GetMeasurements() []ShardMeasurementSeries`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *ShardMeasurementsResponse) GetMeasurementsOk() (*[]ShardMeasurementSeries, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *ShardMeasurementsResponse) SetMeasurements(v []ShardMeasurementSeries)`

SetMeasurements sets Measurements field to given value.

### GetReplicaSetName

`func (o *ShardMeasurementsResponse) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ShardMeasurementsResponse) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ShardMeasurementsResponse) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### GetStart

`func (o *ShardMeasurementsResponse) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ShardMeasurementsResponse) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ShardMeasurementsResponse) SetStart(v time.Time)`

SetStart sets Start field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


