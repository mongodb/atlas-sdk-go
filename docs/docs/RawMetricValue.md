# RawMetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **float64** | Amount of the **metricName** recorded at the time of the event. This value triggered the alert. | [optional] [readonly] 
**Units** | Pointer to [**RawMetricUnits**](RawMetricUnits.md) |  | [optional] [default to RAWMETRICUNITS_RAW]

## Methods

### NewRawMetricValue

`func NewRawMetricValue() *RawMetricValue`

NewRawMetricValue instantiates a new RawMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawMetricValueWithDefaults

`func NewRawMetricValueWithDefaults() *RawMetricValue`

NewRawMetricValueWithDefaults instantiates a new RawMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *RawMetricValue) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RawMetricValue) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RawMetricValue) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RawMetricValue) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUnits

`func (o *RawMetricValue) GetUnits() RawMetricUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RawMetricValue) GetUnitsOk() (*RawMetricUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RawMetricValue) SetUnits(v RawMetricUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *RawMetricValue) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


