# ShardMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | [**[]ShardMeasurementDataPoint**](ShardMeasurementDataPoint.md) | Data points of the measurement at &#x60;granularity&#x60; intervals, starting at &#x60;start&#x60; and not passing &#x60;end&#x60;, in ascending time order. | [readonly] 
**Name** | **string** | Measurement that the data points quantify. | [readonly] 
**ReplicaSetName** | **string** | Human-readable label that identifies the shard, as the replica set name of its members. | [readonly] 
**Units** | **string** | Unit in which the data points express the measurement. | [readonly] 

## Methods

### NewShardMeasurement

`func NewShardMeasurement(dataPoints []ShardMeasurementDataPoint, name string, replicaSetName string, units string, ) *ShardMeasurement`

NewShardMeasurement instantiates a new ShardMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardMeasurementWithDefaults

`func NewShardMeasurementWithDefaults() *ShardMeasurement`

NewShardMeasurementWithDefaults instantiates a new ShardMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *ShardMeasurement) GetDataPoints() []ShardMeasurementDataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *ShardMeasurement) GetDataPointsOk() (*[]ShardMeasurementDataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *ShardMeasurement) SetDataPoints(v []ShardMeasurementDataPoint)`

SetDataPoints sets DataPoints field to given value.

### GetName

`func (o *ShardMeasurement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShardMeasurement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShardMeasurement) SetName(v string)`

SetName sets Name field to given value.

### GetReplicaSetName

`func (o *ShardMeasurement) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ShardMeasurement) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ShardMeasurement) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### GetUnits

`func (o *ShardMeasurement) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ShardMeasurement) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ShardMeasurement) SetUnits(v string)`

SetUnits sets Units field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


