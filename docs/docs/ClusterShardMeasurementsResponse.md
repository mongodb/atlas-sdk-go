# ClusterShardMeasurementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Human-readable label that identifies the cluster. | [readonly] 
**End** | **time.Time** | Date and time at which the window ends. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**Granularity** | **string** | Interval between consecutive data points, in ISO 8601 duration format. Either the requested &#x60;granularity&#x60; or, when none was requested, the resolution MongoDB Cloud chose for the width of the window. | [readonly] 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**Measurements** | [**[]ShardMeasurement**](ShardMeasurement.md) | One time series per requested measurement on each shard of the cluster, ordered by measurement and then by &#x60;replicaSetName&#x60;. A shard that MongoDB Cloud has no readings for is still listed, with &#x60;null&#x60; data point values. | [readonly] 
**Start** | **time.Time** | Date and time of the first data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 

## Methods

### NewClusterShardMeasurementsResponse

`func NewClusterShardMeasurementsResponse(clusterName string, end time.Time, granularity string, groupId string, measurements []ShardMeasurement, start time.Time, ) *ClusterShardMeasurementsResponse`

NewClusterShardMeasurementsResponse instantiates a new ClusterShardMeasurementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterShardMeasurementsResponseWithDefaults

`func NewClusterShardMeasurementsResponseWithDefaults() *ClusterShardMeasurementsResponse`

NewClusterShardMeasurementsResponseWithDefaults instantiates a new ClusterShardMeasurementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterShardMeasurementsResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterShardMeasurementsResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterShardMeasurementsResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### GetEnd

`func (o *ClusterShardMeasurementsResponse) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClusterShardMeasurementsResponse) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClusterShardMeasurementsResponse) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### GetGranularity

`func (o *ClusterShardMeasurementsResponse) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ClusterShardMeasurementsResponse) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ClusterShardMeasurementsResponse) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### GetGroupId

`func (o *ClusterShardMeasurementsResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ClusterShardMeasurementsResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ClusterShardMeasurementsResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetMeasurements

`func (o *ClusterShardMeasurementsResponse) GetMeasurements() []ShardMeasurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *ClusterShardMeasurementsResponse) GetMeasurementsOk() (*[]ShardMeasurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *ClusterShardMeasurementsResponse) SetMeasurements(v []ShardMeasurement)`

SetMeasurements sets Measurements field to given value.

### GetStart

`func (o *ClusterShardMeasurementsResponse) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClusterShardMeasurementsResponse) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClusterShardMeasurementsResponse) SetStart(v time.Time)`

SetStart sets Start field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


