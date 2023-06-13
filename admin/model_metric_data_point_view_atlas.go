// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"time"
)

// checks if the MetricDataPointViewAtlas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDataPointViewAtlas{}

// MetricDataPointViewAtlas struct for MetricDataPointViewAtlas
type MetricDataPointViewAtlas struct {
	// Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Value that comprises this data point.
	Value *float32 `json:"value,omitempty"`
}

// NewMetricDataPointViewAtlas instantiates a new MetricDataPointViewAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDataPointViewAtlas() *MetricDataPointViewAtlas {
	this := MetricDataPointViewAtlas{}
	return &this
}

// NewMetricDataPointViewAtlasWithDefaults instantiates a new MetricDataPointViewAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDataPointViewAtlasWithDefaults() *MetricDataPointViewAtlas {
	this := MetricDataPointViewAtlas{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MetricDataPointViewAtlas) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointViewAtlas) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetricDataPointViewAtlas) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MetricDataPointViewAtlas) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricDataPointViewAtlas) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointViewAtlas) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricDataPointViewAtlas) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *MetricDataPointViewAtlas) SetValue(v float32) {
	o.Value = &v
}

func (o MetricDataPointViewAtlas) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o MetricDataPointViewAtlas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableMetricDataPointViewAtlas struct {
	value *MetricDataPointViewAtlas
	isSet bool
}

func (v NullableMetricDataPointViewAtlas) Get() *MetricDataPointViewAtlas {
	return v.value
}

func (v *NullableMetricDataPointViewAtlas) Set(val *MetricDataPointViewAtlas) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDataPointViewAtlas) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDataPointViewAtlas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDataPointViewAtlas(val *MetricDataPointViewAtlas) *NullableMetricDataPointViewAtlas {
	return &NullableMetricDataPointViewAtlas{value: val, isSet: true}
}

func (v NullableMetricDataPointViewAtlas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDataPointViewAtlas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
