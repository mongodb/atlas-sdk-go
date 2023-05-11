# HostMetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **float64** | Amount of the **metricName** recorded at the time of the event. This value triggered the alert. | [optional] [readonly] 
**Units** | Pointer to **string** |  | [optional] 

## Methods

### NewHostMetricValue

`func NewHostMetricValue() *HostMetricValue`

NewHostMetricValue instantiates a new HostMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMetricValueWithDefaults

`func NewHostMetricValueWithDefaults() *HostMetricValue`

NewHostMetricValueWithDefaults instantiates a new HostMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *HostMetricValue) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *HostMetricValue) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *HostMetricValue) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *HostMetricValue) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUnits

`func (o *HostMetricValue) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *HostMetricValue) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *HostMetricValue) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *HostMetricValue) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


