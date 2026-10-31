// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ShardMeasurementDataPoint Value of one measurement at one moment in time. The `value` is `null` when no reading exists for that moment.
type ShardMeasurementDataPoint struct {
	// Date and time of the data point. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Timestamp time.Time `json:"timestamp"`
	// Value of the measurement at this moment, in the `units` of the measurement. `null` when no reading exists for this moment, including the period after a cluster or shard is created before MongoDB Cloud starts collecting its metrics.
	// Read only field.
	Value float64 `json:"value"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardMeasurementDataPoint) MarshalJSON() ([]byte, error) {
	type noMethod ShardMeasurementDataPoint
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardMeasurementDataPoint instantiates a new ShardMeasurementDataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardMeasurementDataPoint(timestamp time.Time, value float64) *ShardMeasurementDataPoint {
	this := ShardMeasurementDataPoint{}
	this.Timestamp = timestamp
	this.Value = value
	return &this
}

// NewShardMeasurementDataPointWithDefaults instantiates a new ShardMeasurementDataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardMeasurementDataPointWithDefaults() *ShardMeasurementDataPoint {
	this := ShardMeasurementDataPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *ShardMeasurementDataPoint) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementDataPoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ShardMeasurementDataPoint) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetValue returns the Value field value
func (o *ShardMeasurementDataPoint) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ShardMeasurementDataPoint) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ShardMeasurementDataPoint) SetValue(v float64) {
	o.Value = v
}
