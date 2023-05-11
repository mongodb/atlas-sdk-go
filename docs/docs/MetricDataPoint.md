# MetricDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Value** | Pointer to **float32** | Value that comprises this data point. | [optional] [readonly] 

## Methods

### NewMetricDataPoint

`func NewMetricDataPoint() *MetricDataPoint`

NewMetricDataPoint instantiates a new MetricDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataPointWithDefaults

`func NewMetricDataPointWithDefaults() *MetricDataPoint`

NewMetricDataPointWithDefaults instantiates a new MetricDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MetricDataPoint) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricDataPoint) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricDataPoint) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricDataPoint) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *MetricDataPoint) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricDataPoint) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricDataPoint) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricDataPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


