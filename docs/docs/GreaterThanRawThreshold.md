# GreaterThanRawThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value. | [optional] 
**Threshold** | Pointer to **int** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to **string** | Element used to express the quantity. This can be an element of time, storage capacity, and the like. | [optional] [default to "RAW"]

## Methods

### NewGreaterThanRawThreshold

`func NewGreaterThanRawThreshold() *GreaterThanRawThreshold`

NewGreaterThanRawThreshold instantiates a new GreaterThanRawThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGreaterThanRawThresholdWithDefaults

`func NewGreaterThanRawThresholdWithDefaults() *GreaterThanRawThreshold`

NewGreaterThanRawThresholdWithDefaults instantiates a new GreaterThanRawThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *GreaterThanRawThreshold) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GreaterThanRawThreshold) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GreaterThanRawThreshold) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *GreaterThanRawThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.
### GetThreshold

`func (o *GreaterThanRawThreshold) GetThreshold() int`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GreaterThanRawThreshold) GetThresholdOk() (*int, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GreaterThanRawThreshold) SetThreshold(v int)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GreaterThanRawThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.
### GetUnits

`func (o *GreaterThanRawThreshold) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GreaterThanRawThreshold) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GreaterThanRawThreshold) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GreaterThanRawThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


