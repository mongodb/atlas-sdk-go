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

// checks if the CharFiltermappingMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharFiltermappingMappings{}

// CharFiltermappingMappings Comma-separated list of mappings. A mapping indicates that one character or group of characters should be substituted for another, using the following format:  `<original> : <replacement>`
type CharFiltermappingMappings struct {
	AdditionalPropertiesField *string `json:"additionalProperties,omitempty"`
}

// NewCharFiltermappingMappings instantiates a new CharFiltermappingMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharFiltermappingMappings() *CharFiltermappingMappings {
	this := CharFiltermappingMappings{}
	return &this
}

// NewCharFiltermappingMappingsWithDefaults instantiates a new CharFiltermappingMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharFiltermappingMappingsWithDefaults() *CharFiltermappingMappings {
	this := CharFiltermappingMappings{}
	return &this
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *CharFiltermappingMappings) GetAdditionalPropertiesField() string {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret string
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharFiltermappingMappings) GetAdditionalPropertiesFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *CharFiltermappingMappings) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given string and assigns it to the AdditionalPropertiesField field.
func (o *CharFiltermappingMappings) SetAdditionalPropertiesField(v string) {
	o.AdditionalPropertiesField = &v
}

func (o CharFiltermappingMappings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CharFiltermappingMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return toSerialize, nil
}

type NullableCharFiltermappingMappings struct {
	value *CharFiltermappingMappings
	isSet bool
}

func (v NullableCharFiltermappingMappings) Get() *CharFiltermappingMappings {
	return v.value
}

func (v *NullableCharFiltermappingMappings) Set(val *CharFiltermappingMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableCharFiltermappingMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableCharFiltermappingMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharFiltermappingMappings(val *CharFiltermappingMappings) *NullableCharFiltermappingMappings {
	return &NullableCharFiltermappingMappings{value: val, isSet: true}
}

func (v NullableCharFiltermappingMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharFiltermappingMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


