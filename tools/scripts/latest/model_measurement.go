/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
)

// checks if the Measurement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Measurement{}

// Measurement struct for Measurement
type Measurement struct {
	// List that contains the value of, and metadata provided for, one data point generated at a particular moment in time. If no data point exists for a particular moment in time, the `value` parameter returns `null`.
	DataPoints []MetricDataPoint `json:"dataPoints,omitempty"`
	// Human-readable label of the measurement that this data point covers.
	Name *string `json:"name,omitempty"`
	// Element used to quantify the measurement. The resource returns units of throughput, storage, and time.
	Units *string `json:"units,omitempty"`
}

// NewMeasurement instantiates a new Measurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurement() *Measurement {
	this := Measurement{}
	return &this
}

// NewMeasurementWithDefaults instantiates a new Measurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementWithDefaults() *Measurement {
	this := Measurement{}
	return &this
}

// GetDataPoints returns the DataPoints field value if set, zero value otherwise.
func (o *Measurement) GetDataPoints() []MetricDataPoint {
	if o == nil || IsNil(o.DataPoints) {
		var ret []MetricDataPoint
		return ret
	}
	return o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetDataPointsOk() ([]MetricDataPoint, bool) {
	if o == nil || IsNil(o.DataPoints) {
		return nil, false
	}
	return o.DataPoints, true
}

// HasDataPoints returns a boolean if a field has been set.
func (o *Measurement) HasDataPoints() bool {
	if o != nil && !IsNil(o.DataPoints) {
		return true
	}

	return false
}

// SetDataPoints gets a reference to the given []MetricDataPoint and assigns it to the DataPoints field.
func (o *Measurement) SetDataPoints(v []MetricDataPoint) {
	o.DataPoints = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Measurement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Measurement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Measurement) SetName(v string) {
	o.Name = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *Measurement) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Measurement) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *Measurement) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *Measurement) SetUnits(v string) {
	o.Units = &v
}

func (o Measurement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Measurement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: dataPoints is readOnly
	// skip: name is readOnly
	// skip: units is readOnly
	return toSerialize, nil
}

type NullableMeasurement struct {
	value *Measurement
	isSet bool
}

func (v NullableMeasurement) Get() *Measurement {
	return v.value
}

func (v *NullableMeasurement) Set(val *Measurement) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurement(val *Measurement) *NullableMeasurement {
	return &NullableMeasurement{value: val, isSet: true}
}

func (v NullableMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


