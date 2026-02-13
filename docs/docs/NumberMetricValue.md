# NumberMetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **float64** | Amount of the &#x60;metricName&#x60; recorded at the time of the event. This value triggered the alert. | [optional] [readonly] 
**Units** | Pointer to **string** | Element used to express the quantity. This can be an element of time, storage capacity, and the like. | [optional] 

## Methods

### NewNumberMetricValue

`func NewNumberMetricValue() *NumberMetricValue`

NewNumberMetricValue instantiates a new NumberMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberMetricValueWithDefaults

`func NewNumberMetricValueWithDefaults() *NumberMetricValue`

NewNumberMetricValueWithDefaults instantiates a new NumberMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *NumberMetricValue) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NumberMetricValue) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NumberMetricValue) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *NumberMetricValue) HasNumber() bool`

HasNumber returns a boolean if a field has been set.
### GetUnits

`func (o *NumberMetricValue) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *NumberMetricValue) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *NumberMetricValue) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *NumberMetricValue) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


