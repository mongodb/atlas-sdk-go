# ShardMeasurementSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | [**[]ShardMeasurementDataPoint**](ShardMeasurementDataPoint.md) | Data points of the measurement at &#x60;granularity&#x60; intervals, starting at &#x60;start&#x60; and not passing &#x60;end&#x60;, in ascending time order. | [readonly] 
**Name** | **string** | Measurement that the data points quantify. | [readonly] 
**Units** | **string** | Unit in which the data points express the measurement. | [readonly] 

## Methods

### NewShardMeasurementSeries

`func NewShardMeasurementSeries(dataPoints []ShardMeasurementDataPoint, name string, units string, ) *ShardMeasurementSeries`

NewShardMeasurementSeries instantiates a new ShardMeasurementSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardMeasurementSeriesWithDefaults

`func NewShardMeasurementSeriesWithDefaults() *ShardMeasurementSeries`

NewShardMeasurementSeriesWithDefaults instantiates a new ShardMeasurementSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *ShardMeasurementSeries) GetDataPoints() []ShardMeasurementDataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *ShardMeasurementSeries) GetDataPointsOk() (*[]ShardMeasurementDataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *ShardMeasurementSeries) SetDataPoints(v []ShardMeasurementDataPoint)`

SetDataPoints sets DataPoints field to given value.

### GetName

`func (o *ShardMeasurementSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShardMeasurementSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShardMeasurementSeries) SetName(v string)`

SetName sets Name field to given value.

### GetUnits

`func (o *ShardMeasurementSeries) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ShardMeasurementSeries) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ShardMeasurementSeries) SetUnits(v string)`

SetUnits sets Units field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


