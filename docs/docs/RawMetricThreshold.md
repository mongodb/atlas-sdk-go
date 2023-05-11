# RawMetricThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. | [optional] 
**Mode** | Pointer to **string** | MongoDB Cloud computes the current metric value as an average. | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Threshold** | Pointer to **float64** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to [**RawMetricUnits**](RawMetricUnits.md) |  | [optional] [default to RAWMETRICUNITS_RAW]

## Methods

### NewRawMetricThreshold

`func NewRawMetricThreshold() *RawMetricThreshold`

NewRawMetricThreshold instantiates a new RawMetricThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawMetricThresholdWithDefaults

`func NewRawMetricThresholdWithDefaults() *RawMetricThreshold`

NewRawMetricThresholdWithDefaults instantiates a new RawMetricThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *RawMetricThreshold) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RawMetricThreshold) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RawMetricThreshold) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RawMetricThreshold) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMode

`func (o *RawMetricThreshold) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RawMetricThreshold) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RawMetricThreshold) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RawMetricThreshold) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperator

`func (o *RawMetricThreshold) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RawMetricThreshold) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RawMetricThreshold) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RawMetricThreshold) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *RawMetricThreshold) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RawMetricThreshold) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RawMetricThreshold) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RawMetricThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *RawMetricThreshold) GetUnits() RawMetricUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RawMetricThreshold) GetUnitsOk() (*RawMetricUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RawMetricThreshold) SetUnits(v RawMetricUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *RawMetricThreshold) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


