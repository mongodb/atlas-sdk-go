# AppServiceMetricThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. | [optional] 
**Mode** | Pointer to **string** | MongoDB Cloud computes the current metric value as an average. | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Threshold** | Pointer to **float64** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to [**DataMetricUnits**](DataMetricUnits.md) |  | [optional] 

## Methods

### NewAppServiceMetricThreshold

`func NewAppServiceMetricThreshold() *AppServiceMetricThreshold`

NewAppServiceMetricThreshold instantiates a new AppServiceMetricThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceMetricThresholdWithDefaults

`func NewAppServiceMetricThresholdWithDefaults() *AppServiceMetricThreshold`

NewAppServiceMetricThresholdWithDefaults instantiates a new AppServiceMetricThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *AppServiceMetricThreshold) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AppServiceMetricThreshold) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AppServiceMetricThreshold) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *AppServiceMetricThreshold) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMode

`func (o *AppServiceMetricThreshold) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AppServiceMetricThreshold) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AppServiceMetricThreshold) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AppServiceMetricThreshold) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperator

`func (o *AppServiceMetricThreshold) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppServiceMetricThreshold) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppServiceMetricThreshold) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppServiceMetricThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *AppServiceMetricThreshold) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AppServiceMetricThreshold) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AppServiceMetricThreshold) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *AppServiceMetricThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *AppServiceMetricThreshold) GetUnits() DataMetricUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *AppServiceMetricThreshold) GetUnitsOk() (*DataMetricUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *AppServiceMetricThreshold) SetUnits(v DataMetricUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *AppServiceMetricThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


