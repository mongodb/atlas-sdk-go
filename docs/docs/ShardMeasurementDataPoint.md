# ShardMeasurementDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | Date and time of the data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**Value** | **float64** | Value of the measurement at this moment, in the &#x60;units&#x60; of the measurement. &#x60;null&#x60; when no reading exists for this moment, including the period after a cluster or shard is created before MongoDB Cloud starts collecting its metrics. | [readonly] 

## Methods

### NewShardMeasurementDataPoint

`func NewShardMeasurementDataPoint(timestamp time.Time, value float64, ) *ShardMeasurementDataPoint`

NewShardMeasurementDataPoint instantiates a new ShardMeasurementDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardMeasurementDataPointWithDefaults

`func NewShardMeasurementDataPointWithDefaults() *ShardMeasurementDataPoint`

NewShardMeasurementDataPointWithDefaults instantiates a new ShardMeasurementDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ShardMeasurementDataPoint) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShardMeasurementDataPoint) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShardMeasurementDataPoint) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### GetValue

`func (o *ShardMeasurementDataPoint) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShardMeasurementDataPoint) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShardMeasurementDataPoint) SetValue(v float64)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


