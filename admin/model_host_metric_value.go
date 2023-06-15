// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the HostMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostMetricValue{}

// HostMetricValue Value of the metric that triggered the alert. The resource returns this parameter for alerts of events impacting hosts.
type HostMetricValue struct {
	// Amount of the **metricName** recorded at the time of the event. This value triggered the alert.
	Number *float64 `json:"number,omitempty"`
	Units  *string  `json:"units,omitempty"`
}

// NewHostMetricValue instantiates a new HostMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMetricValue() *HostMetricValue {
	this := HostMetricValue{}
	return &this
}

// NewHostMetricValueWithDefaults instantiates a new HostMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMetricValueWithDefaults() *HostMetricValue {
	this := HostMetricValue{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *HostMetricValue) GetNumber() float64 {
	if o == nil || IsNil(o.Number) {
		var ret float64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricValue) GetNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *HostMetricValue) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float64 and assigns it to the Number field.
func (o *HostMetricValue) SetNumber(v float64) {
	o.Number = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *HostMetricValue) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMetricValue) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *HostMetricValue) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *HostMetricValue) SetUnits(v string) {
	o.Units = &v
}

func (o HostMetricValue) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o HostMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableHostMetricValue struct {
	value *HostMetricValue
	isSet bool
}

func (v NullableHostMetricValue) Get() *HostMetricValue {
	return v.value
}

func (v *NullableHostMetricValue) Set(val *HostMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMetricValue(val *HostMetricValue) *NullableHostMetricValue {
	return &NullableHostMetricValue{value: val, isSet: true}
}

func (v NullableHostMetricValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
