/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the FTSMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSMetric{}

// FTSMetric Measurement of one Atlas Search status when MongoDB Atlas received this request.
type FTSMetric struct {
	// Human-readable label that identifies this Atlas Search hardware, status, or index measurement.
	MetricName string `json:"metricName"`
	// Unit of measurement that applies to this Atlas Search metric.
	Units string `json:"units"`
}

// NewFTSMetric instantiates a new FTSMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSMetric(metricName string, units string) *FTSMetric {
	this := FTSMetric{}
	this.MetricName = metricName
	this.Units = units
	return &this
}

// NewFTSMetricWithDefaults instantiates a new FTSMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSMetricWithDefaults() *FTSMetric {
	this := FTSMetric{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *FTSMetric) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *FTSMetric) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *FTSMetric) SetMetricName(v string) {
	o.MetricName = v
}

// GetUnits returns the Units field value
func (o *FTSMetric) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *FTSMetric) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *FTSMetric) SetUnits(v string) {
	o.Units = v
}

func (o FTSMetric) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableFTSMetric struct {
	value *FTSMetric
	isSet bool
}

func (v NullableFTSMetric) Get() *FTSMetric {
	return v.value
}

func (v *NullableFTSMetric) Set(val *FTSMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSMetric(val *FTSMetric) *NullableFTSMetric {
	return &NullableFTSMetric{value: val, isSet: true}
}

func (v NullableFTSMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


