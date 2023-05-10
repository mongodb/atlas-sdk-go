// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the CharFiltermapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharFiltermapping{}

// CharFiltermapping Filter that applies normalization mappings that you specify to characters.
type CharFiltermapping struct {
	Mappings CharFiltermappingMappings `json:"mappings"`
	// Human-readable label that identifies this character filter type.
	Type string `json:"type"`
}

// NewCharFiltermapping instantiates a new CharFiltermapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharFiltermapping(mappings CharFiltermappingMappings, type_ string) *CharFiltermapping {
	this := CharFiltermapping{}
	this.Mappings = mappings
	this.Type = type_
	return &this
}

// NewCharFiltermappingWithDefaults instantiates a new CharFiltermapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharFiltermappingWithDefaults() *CharFiltermapping {
	this := CharFiltermapping{}
	return &this
}

// GetMappings returns the Mappings field value
func (o *CharFiltermapping) GetMappings() CharFiltermappingMappings {
	if o == nil {
		var ret CharFiltermappingMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *CharFiltermapping) GetMappingsOk() (*CharFiltermappingMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// SetMappings sets field value
func (o *CharFiltermapping) SetMappings(v CharFiltermappingMappings) {
	o.Mappings = v
}

// GetType returns the Type field value
func (o *CharFiltermapping) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharFiltermapping) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharFiltermapping) SetType(v string) {
	o.Type = v
}

func (o CharFiltermapping) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CharFiltermapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mappings"] = o.Mappings
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCharFiltermapping struct {
	value *CharFiltermapping
	isSet bool
}

func (v NullableCharFiltermapping) Get() *CharFiltermapping {
	return v.value
}

func (v *NullableCharFiltermapping) Set(val *CharFiltermapping) {
	v.value = val
	v.isSet = true
}

func (v NullableCharFiltermapping) IsSet() bool {
	return v.isSet
}

func (v *NullableCharFiltermapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharFiltermapping(val *CharFiltermapping) *NullableCharFiltermapping {
	return &NullableCharFiltermapping{value: val, isSet: true}
}

func (v NullableCharFiltermapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharFiltermapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
