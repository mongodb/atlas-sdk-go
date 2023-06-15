// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataMetricThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataMetricThreshold{}

// DataMetricThreshold struct for DataMetricThreshold
type DataMetricThreshold struct {
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.
	MetricName *string `json:"metricName,omitempty"`
	// MongoDB Cloud computes the current metric value as an average.
	Mode *string `json:"mode,omitempty"`
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `json:"threshold,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.
	Units *string `json:"units,omitempty"`
}

// NewDataMetricThreshold instantiates a new DataMetricThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataMetricThreshold() *DataMetricThreshold {
	this := DataMetricThreshold{}
	return &this
}

// NewDataMetricThresholdWithDefaults instantiates a new DataMetricThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataMetricThresholdWithDefaults() *DataMetricThreshold {
	this := DataMetricThreshold{}
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *DataMetricThreshold) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMetricThreshold) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *DataMetricThreshold) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *DataMetricThreshold) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DataMetricThreshold) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMetricThreshold) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DataMetricThreshold) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DataMetricThreshold) SetMode(v string) {
	o.Mode = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *DataMetricThreshold) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMetricThreshold) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *DataMetricThreshold) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *DataMetricThreshold) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *DataMetricThreshold) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMetricThreshold) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *DataMetricThreshold) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *DataMetricThreshold) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *DataMetricThreshold) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMetricThreshold) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *DataMetricThreshold) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *DataMetricThreshold) SetUnits(v string) {
	o.Units = &v
}

func (o DataMetricThreshold) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataMetricThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableDataMetricThreshold struct {
	value *DataMetricThreshold
	isSet bool
}

func (v NullableDataMetricThreshold) Get() *DataMetricThreshold {
	return v.value
}

func (v *NullableDataMetricThreshold) Set(val *DataMetricThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableDataMetricThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableDataMetricThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataMetricThreshold(val *DataMetricThreshold) *NullableDataMetricThreshold {
	return &NullableDataMetricThreshold{value: val, isSet: true}
}

func (v NullableDataMetricThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataMetricThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
