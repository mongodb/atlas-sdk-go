# FTSMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | Human-readable label that identifies this Atlas Search hardware, status, or index measurement. | [readonly] 
**Units** | **string** | Unit of measurement that applies to this Atlas Search metric. | [readonly] 

## Methods

### NewFTSMetric

`func NewFTSMetric(metricName string, units string, ) *FTSMetric`

NewFTSMetric instantiates a new FTSMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSMetricWithDefaults

`func NewFTSMetricWithDefaults() *FTSMetric`

NewFTSMetricWithDefaults instantiates a new FTSMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *FTSMetric) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *FTSMetric) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *FTSMetric) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetUnits

`func (o *FTSMetric) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FTSMetric) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FTSMetric) SetUnits(v string)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


