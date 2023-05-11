# ThresholdViewInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Threshold** | Pointer to **int** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to **string** | Element used to express the quantity. This can be an element of time, storage capacity, and the like. | [optional] 

## Methods

### NewThresholdViewInteger

`func NewThresholdViewInteger() *ThresholdViewInteger`

NewThresholdViewInteger instantiates a new ThresholdViewInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdViewIntegerWithDefaults

`func NewThresholdViewIntegerWithDefaults() *ThresholdViewInteger`

NewThresholdViewIntegerWithDefaults instantiates a new ThresholdViewInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ThresholdViewInteger) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ThresholdViewInteger) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ThresholdViewInteger) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ThresholdViewInteger) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *ThresholdViewInteger) GetThreshold() int`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdViewInteger) GetThresholdOk() (*int, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdViewInteger) SetThreshold(v int)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ThresholdViewInteger) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *ThresholdViewInteger) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ThresholdViewInteger) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ThresholdViewInteger) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ThresholdViewInteger) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


