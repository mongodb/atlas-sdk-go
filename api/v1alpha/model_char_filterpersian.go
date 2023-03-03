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

// checks if the CharFilterpersian type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharFilterpersian{}

// CharFilterpersian Filter that replaces instances of a zero-width non-joiner with an ordinary space. It is based on Lucene's PersianCharFilter.
type CharFilterpersian struct {
	// Human-readable label that identifies this character filter type.
	Type string `json:"type"`
}

// NewCharFilterpersian instantiates a new CharFilterpersian object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharFilterpersian(type_ string) *CharFilterpersian {
	this := CharFilterpersian{}
	this.Type = type_
	return &this
}

// NewCharFilterpersianWithDefaults instantiates a new CharFilterpersian object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharFilterpersianWithDefaults() *CharFilterpersian {
	this := CharFilterpersian{}
	return &this
}

// GetType returns the Type field value
func (o *CharFilterpersian) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharFilterpersian) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharFilterpersian) SetType(v string) {
	o.Type = v
}

func (o CharFilterpersian) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharFilterpersian) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCharFilterpersian struct {
	value *CharFilterpersian
	isSet bool
}

func (v NullableCharFilterpersian) Get() *CharFilterpersian {
	return v.value
}

func (v *NullableCharFilterpersian) Set(val *CharFilterpersian) {
	v.value = val
	v.isSet = true
}

func (v NullableCharFilterpersian) IsSet() bool {
	return v.isSet
}

func (v *NullableCharFilterpersian) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharFilterpersian(val *CharFilterpersian) *NullableCharFilterpersian {
	return &NullableCharFilterpersian{value: val, isSet: true}
}

func (v NullableCharFilterpersian) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharFilterpersian) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


