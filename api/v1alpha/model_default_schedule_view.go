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

// checks if the DefaultScheduleView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultScheduleView{}

// DefaultScheduleView struct for DefaultScheduleView
type DefaultScheduleView struct {
	Type string `json:"type"`
}

// NewDefaultScheduleView instantiates a new DefaultScheduleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultScheduleView() *DefaultScheduleView {
	this := DefaultScheduleView{}
	return &this
}

// NewDefaultScheduleViewWithDefaults instantiates a new DefaultScheduleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultScheduleViewWithDefaults() *DefaultScheduleView {
	this := DefaultScheduleView{}
	return &this
}

// GetType returns the Type field value
func (o *DefaultScheduleView) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DefaultScheduleView) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DefaultScheduleView) SetType(v string) {
	o.Type = v
}

func (o DefaultScheduleView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultScheduleView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDefaultScheduleView struct {
	value *DefaultScheduleView
	isSet bool
}

func (v NullableDefaultScheduleView) Get() *DefaultScheduleView {
	return v.value
}

func (v *NullableDefaultScheduleView) Set(val *DefaultScheduleView) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultScheduleView) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultScheduleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultScheduleView(val *DefaultScheduleView) *NullableDefaultScheduleView {
	return &NullableDefaultScheduleView{value: val, isSet: true}
}

func (v NullableDefaultScheduleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultScheduleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


