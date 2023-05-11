# LessThanTimeThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value. | [optional] 
**Threshold** | Pointer to **int** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to [**TimeMetricUnits**](TimeMetricUnits.md) |  | [optional] [default to TIMEMETRICUNITS_HOURS]

## Methods

### NewLessThanTimeThreshold

`func NewLessThanTimeThreshold() *LessThanTimeThreshold`

NewLessThanTimeThreshold instantiates a new LessThanTimeThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLessThanTimeThresholdWithDefaults

`func NewLessThanTimeThresholdWithDefaults() *LessThanTimeThreshold`

NewLessThanTimeThresholdWithDefaults instantiates a new LessThanTimeThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *LessThanTimeThreshold) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LessThanTimeThreshold) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LessThanTimeThreshold) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LessThanTimeThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *LessThanTimeThreshold) GetThreshold() int`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LessThanTimeThreshold) GetThresholdOk() (*int, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LessThanTimeThreshold) SetThreshold(v int)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *LessThanTimeThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *LessThanTimeThreshold) GetUnits() TimeMetricUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *LessThanTimeThreshold) GetUnitsOk() (*TimeMetricUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *LessThanTimeThreshold) SetUnits(v TimeMetricUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *LessThanTimeThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


