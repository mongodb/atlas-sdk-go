// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// MetricDataPointAtlas Value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the `value` parameter returns `null`.
type MetricDataPointAtlas struct {
	// Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Value that comprises this data point.
	// Read only field.
	Value *float32 `json:"value,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MetricDataPointAtlas) MarshalJSON() ([]byte, error) {
	type noMethod MetricDataPointAtlas
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMetricDataPointAtlas instantiates a new MetricDataPointAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDataPointAtlas() *MetricDataPointAtlas {
	this := MetricDataPointAtlas{}
	return &this
}

// NewMetricDataPointAtlasWithDefaults instantiates a new MetricDataPointAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDataPointAtlasWithDefaults() *MetricDataPointAtlas {
	this := MetricDataPointAtlas{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise
func (o *MetricDataPointAtlas) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointAtlas) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}

	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MetricDataPointAtlas) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MetricDataPointAtlas) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// SetTimestampNil sets Timestamp to an explicit JSON null when marshaled.
func (o *MetricDataPointAtlas) SetTimestampNil() {
	o.Timestamp = nil
	o.NullFields = append(o.NullFields, "Timestamp")
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *MetricDataPointAtlas) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDataPointAtlas) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricDataPointAtlas) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *MetricDataPointAtlas) SetValue(v float32) {
	o.Value = &v
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *MetricDataPointAtlas) SetValueNil() {
	o.Value = nil
	o.NullFields = append(o.NullFields, "Value")
}
