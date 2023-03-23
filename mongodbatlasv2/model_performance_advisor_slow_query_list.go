/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the PerformanceAdvisorSlowQueryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceAdvisorSlowQueryList{}

// PerformanceAdvisorSlowQueryList struct for PerformanceAdvisorSlowQueryList
type PerformanceAdvisorSlowQueryList struct {
	// List of operations that the Performance Advisor detected that took longer to execute than a specified threshold.
	SlowQueries []PerformanceAdvisorSlowQuery `json:"slowQueries,omitempty"`
}

// NewPerformanceAdvisorSlowQueryList instantiates a new PerformanceAdvisorSlowQueryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorSlowQueryList() *PerformanceAdvisorSlowQueryList {
	this := PerformanceAdvisorSlowQueryList{}
	return &this
}

// NewPerformanceAdvisorSlowQueryListWithDefaults instantiates a new PerformanceAdvisorSlowQueryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorSlowQueryListWithDefaults() *PerformanceAdvisorSlowQueryList {
	this := PerformanceAdvisorSlowQueryList{}
	return &this
}

// GetSlowQueries returns the SlowQueries field value if set, zero value otherwise.
func (o *PerformanceAdvisorSlowQueryList) GetSlowQueries() []PerformanceAdvisorSlowQuery {
	if o == nil || IsNil(o.SlowQueries) {
		var ret []PerformanceAdvisorSlowQuery
		return ret
	}
	return o.SlowQueries
}

// GetSlowQueriesOk returns a tuple with the SlowQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQueryList) GetSlowQueriesOk() ([]PerformanceAdvisorSlowQuery, bool) {
	if o == nil || IsNil(o.SlowQueries) {
		return nil, false
	}
	return o.SlowQueries, true
}

// HasSlowQueries returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQueryList) HasSlowQueries() bool {
	if o != nil && !IsNil(o.SlowQueries) {
		return true
	}

	return false
}

// SetSlowQueries gets a reference to the given []PerformanceAdvisorSlowQuery and assigns it to the SlowQueries field.
func (o *PerformanceAdvisorSlowQueryList) SetSlowQueries(v []PerformanceAdvisorSlowQuery) {
	o.SlowQueries = v
}

func (o PerformanceAdvisorSlowQueryList) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PerformanceAdvisorSlowQueryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePerformanceAdvisorSlowQueryList struct {
	value *PerformanceAdvisorSlowQueryList
	isSet bool
}

func (v NullablePerformanceAdvisorSlowQueryList) Get() *PerformanceAdvisorSlowQueryList {
	return v.value
}

func (v *NullablePerformanceAdvisorSlowQueryList) Set(val *PerformanceAdvisorSlowQueryList) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceAdvisorSlowQueryList) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceAdvisorSlowQueryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceAdvisorSlowQueryList(val *PerformanceAdvisorSlowQueryList) *NullablePerformanceAdvisorSlowQueryList {
	return &NullablePerformanceAdvisorSlowQueryList{value: val, isSet: true}
}

func (v NullablePerformanceAdvisorSlowQueryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceAdvisorSlowQueryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

