// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AlertsThresholdInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsThresholdInteger{}

// AlertsThresholdInteger A Limit that triggers an alert when  exceeded. The resource returns this parameter when **eventTypeName** has not been set to `OUTSIDE_METRIC_THRESHOLD`.
type AlertsThresholdInteger struct {
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *int `json:"threshold,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.
	Units *string `json:"units,omitempty"`
}

// NewAlertsThresholdInteger instantiates a new AlertsThresholdInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsThresholdInteger() *AlertsThresholdInteger {
	this := AlertsThresholdInteger{}
	return &this
}

// NewAlertsThresholdIntegerWithDefaults instantiates a new AlertsThresholdInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsThresholdIntegerWithDefaults() *AlertsThresholdInteger {
	this := AlertsThresholdInteger{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *AlertsThresholdInteger) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsThresholdInteger) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *AlertsThresholdInteger) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *AlertsThresholdInteger) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *AlertsThresholdInteger) GetThreshold() int {
	if o == nil || IsNil(o.Threshold) {
		var ret int
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsThresholdInteger) GetThresholdOk() (*int, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *AlertsThresholdInteger) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int and assigns it to the Threshold field.
func (o *AlertsThresholdInteger) SetThreshold(v int) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *AlertsThresholdInteger) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsThresholdInteger) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *AlertsThresholdInteger) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *AlertsThresholdInteger) SetUnits(v string) {
	o.Units = &v
}

func (o AlertsThresholdInteger) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AlertsThresholdInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableAlertsThresholdInteger struct {
	value *AlertsThresholdInteger
	isSet bool
}

func (v NullableAlertsThresholdInteger) Get() *AlertsThresholdInteger {
	return v.value
}

func (v *NullableAlertsThresholdInteger) Set(val *AlertsThresholdInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsThresholdInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsThresholdInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsThresholdInteger(val *AlertsThresholdInteger) *NullableAlertsThresholdInteger {
	return &NullableAlertsThresholdInteger{value: val, isSet: true}
}

func (v NullableAlertsThresholdInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsThresholdInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
