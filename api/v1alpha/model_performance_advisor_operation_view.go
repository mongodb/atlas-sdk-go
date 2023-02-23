/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the PerformanceAdvisorOperationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceAdvisorOperationView{}

// PerformanceAdvisorOperationView struct for PerformanceAdvisorOperationView
type PerformanceAdvisorOperationView struct {
	// List that contains the search criteria that the query uses. To use the values in key-value pairs in these predicates requires **Project Data Access Read Only** permissions or greater. Otherwise, MongoDB Cloud redacts these values.
	Predicates []map[string]interface{} `json:"predicates,omitempty"`
	Stats *PerformanceAdvisorOpStatsView `json:"stats,omitempty"`
}

// NewPerformanceAdvisorOperationView instantiates a new PerformanceAdvisorOperationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorOperationView() *PerformanceAdvisorOperationView {
	this := PerformanceAdvisorOperationView{}
	return &this
}

// NewPerformanceAdvisorOperationViewWithDefaults instantiates a new PerformanceAdvisorOperationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorOperationViewWithDefaults() *PerformanceAdvisorOperationView {
	this := PerformanceAdvisorOperationView{}
	return &this
}

// GetPredicates returns the Predicates field value if set, zero value otherwise.
func (o *PerformanceAdvisorOperationView) GetPredicates() []map[string]interface{} {
	if o == nil || IsNil(o.Predicates) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Predicates
}

// GetPredicatesOk returns a tuple with the Predicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOperationView) GetPredicatesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Predicates) {
		return nil, false
	}
	return o.Predicates, true
}

// HasPredicates returns a boolean if a field has been set.
func (o *PerformanceAdvisorOperationView) HasPredicates() bool {
	if o != nil && !IsNil(o.Predicates) {
		return true
	}

	return false
}

// SetPredicates gets a reference to the given []map[string]interface{} and assigns it to the Predicates field.
func (o *PerformanceAdvisorOperationView) SetPredicates(v []map[string]interface{}) {
	o.Predicates = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *PerformanceAdvisorOperationView) GetStats() PerformanceAdvisorOpStatsView {
	if o == nil || IsNil(o.Stats) {
		var ret PerformanceAdvisorOpStatsView
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorOperationView) GetStatsOk() (*PerformanceAdvisorOpStatsView, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *PerformanceAdvisorOperationView) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given PerformanceAdvisorOpStatsView and assigns it to the Stats field.
func (o *PerformanceAdvisorOperationView) SetStats(v PerformanceAdvisorOpStatsView) {
	o.Stats = &v
}

func (o PerformanceAdvisorOperationView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceAdvisorOperationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: predicates is readOnly
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	return toSerialize, nil
}

type NullablePerformanceAdvisorOperationView struct {
	value *PerformanceAdvisorOperationView
	isSet bool
}

func (v NullablePerformanceAdvisorOperationView) Get() *PerformanceAdvisorOperationView {
	return v.value
}

func (v *NullablePerformanceAdvisorOperationView) Set(val *PerformanceAdvisorOperationView) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceAdvisorOperationView) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceAdvisorOperationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceAdvisorOperationView(val *PerformanceAdvisorOperationView) *NullablePerformanceAdvisorOperationView {
	return &NullablePerformanceAdvisorOperationView{value: val, isSet: true}
}

func (v NullablePerformanceAdvisorOperationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceAdvisorOperationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


