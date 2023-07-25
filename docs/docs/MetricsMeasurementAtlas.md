# MetricsMeasurementAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | Pointer to [**[]MetricDataPointAtlas**](MetricDataPointAtlas.md) | List that contains the value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the &#x60;value&#x60; parameter returns &#x60;null&#x60;. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label of the measurement that this data point covers. | [optional] [readonly] 
**Units** | Pointer to **string** | Element used to quantify the measurement. The resource returns units of throughput, storage, and time. | [optional] [readonly] 

## Methods

### NewMetricsMeasurementAtlas

`func NewMetricsMeasurementAtlas() *MetricsMeasurementAtlas`

NewMetricsMeasurementAtlas instantiates a new MetricsMeasurementAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMeasurementAtlasWithDefaults

`func NewMetricsMeasurementAtlasWithDefaults() *MetricsMeasurementAtlas`

NewMetricsMeasurementAtlasWithDefaults instantiates a new MetricsMeasurementAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *MetricsMeasurementAtlas) GetDataPoints() []MetricDataPointAtlas`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *MetricsMeasurementAtlas) GetDataPointsOk() (*[]MetricDataPointAtlas, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *MetricsMeasurementAtlas) SetDataPoints(v []MetricDataPointAtlas)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *MetricsMeasurementAtlas) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.
### GetName

`func (o *MetricsMeasurementAtlas) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsMeasurementAtlas) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsMeasurementAtlas) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsMeasurementAtlas) HasName() bool`

HasName returns a boolean if a field has been set.
### GetUnits

`func (o *MetricsMeasurementAtlas) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MetricsMeasurementAtlas) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MetricsMeasurementAtlas) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MetricsMeasurementAtlas) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


