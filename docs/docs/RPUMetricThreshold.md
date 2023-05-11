# RPUMetricThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. | [optional] 
**Mode** | Pointer to **string** | MongoDB Cloud computes the current metric value as an average. | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Threshold** | Pointer to **float64** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to [**ServerlessMetricUnits**](ServerlessMetricUnits.md) |  | [optional] 

## Methods

### NewRPUMetricThreshold

`func NewRPUMetricThreshold() *RPUMetricThreshold`

NewRPUMetricThreshold instantiates a new RPUMetricThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRPUMetricThresholdWithDefaults

`func NewRPUMetricThresholdWithDefaults() *RPUMetricThreshold`

NewRPUMetricThresholdWithDefaults instantiates a new RPUMetricThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *RPUMetricThreshold) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RPUMetricThreshold) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RPUMetricThreshold) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RPUMetricThreshold) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMode

`func (o *RPUMetricThreshold) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RPUMetricThreshold) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RPUMetricThreshold) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RPUMetricThreshold) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperator

`func (o *RPUMetricThreshold) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RPUMetricThreshold) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RPUMetricThreshold) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RPUMetricThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *RPUMetricThreshold) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RPUMetricThreshold) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RPUMetricThreshold) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RPUMetricThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *RPUMetricThreshold) GetUnits() ServerlessMetricUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RPUMetricThreshold) GetUnitsOk() (*ServerlessMetricUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RPUMetricThreshold) SetUnits(v ServerlessMetricUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *RPUMetricThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


