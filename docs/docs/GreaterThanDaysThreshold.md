# GreaterThanDaysThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value. | [optional] 
**Threshold** | Pointer to **int** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to **string** | Element used to express the quantity. This can be an element of time, storage capacity, and the like. | [optional] 

## Methods

### NewGreaterThanDaysThreshold

`func NewGreaterThanDaysThreshold() *GreaterThanDaysThreshold`

NewGreaterThanDaysThreshold instantiates a new GreaterThanDaysThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGreaterThanDaysThresholdWithDefaults

`func NewGreaterThanDaysThresholdWithDefaults() *GreaterThanDaysThreshold`

NewGreaterThanDaysThresholdWithDefaults instantiates a new GreaterThanDaysThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *GreaterThanDaysThreshold) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GreaterThanDaysThreshold) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GreaterThanDaysThreshold) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *GreaterThanDaysThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *GreaterThanDaysThreshold) GetThreshold() int`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GreaterThanDaysThreshold) GetThresholdOk() (*int, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GreaterThanDaysThreshold) SetThreshold(v int)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GreaterThanDaysThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *GreaterThanDaysThreshold) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GreaterThanDaysThreshold) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GreaterThanDaysThreshold) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GreaterThanDaysThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


