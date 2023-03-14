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

// checks if the DefaultSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultSchedule{}

// DefaultSchedule struct for DefaultSchedule
type DefaultSchedule struct {
	Type string `json:"type"`
}

// NewDefaultSchedule instantiates a new DefaultSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultSchedule(type_ string) *DefaultSchedule {
	this := DefaultSchedule{}
	this.Type = type_
	return &this
}

// NewDefaultScheduleWithDefaults instantiates a new DefaultSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultScheduleWithDefaults() *DefaultSchedule {
	this := DefaultSchedule{}
	return &this
}

// GetType returns the Type field value
func (o *DefaultSchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DefaultSchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DefaultSchedule) SetType(v string) {
	o.Type = v
}

func (o DefaultSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDefaultSchedule struct {
	value *DefaultSchedule
	isSet bool
}

func (v NullableDefaultSchedule) Get() *DefaultSchedule {
	return v.value
}

func (v *NullableDefaultSchedule) Set(val *DefaultSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultSchedule(val *DefaultSchedule) *NullableDefaultSchedule {
	return &NullableDefaultSchedule{value: val, isSet: true}
}

func (v NullableDefaultSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


