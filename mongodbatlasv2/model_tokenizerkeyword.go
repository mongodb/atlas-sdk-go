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

// checks if the Tokenizerkeyword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tokenizerkeyword{}

// Tokenizerkeyword Tokenizer that combines the entire input as a single token.
type Tokenizerkeyword struct {
	// Human-readable label that identifies this tokenizer type.
	Type string `json:"type"`
}

// NewTokenizerkeyword instantiates a new Tokenizerkeyword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizerkeyword(type_ string) *Tokenizerkeyword {
	this := Tokenizerkeyword{}
	this.Type = type_
	return &this
}

// NewTokenizerkeywordWithDefaults instantiates a new Tokenizerkeyword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizerkeywordWithDefaults() *Tokenizerkeyword {
	this := Tokenizerkeyword{}
	return &this
}

// GetType returns the Type field value
func (o *Tokenizerkeyword) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Tokenizerkeyword) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Tokenizerkeyword) SetType(v string) {
	o.Type = v
}

func (o Tokenizerkeyword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tokenizerkeyword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenizerkeyword struct {
	value *Tokenizerkeyword
	isSet bool
}

func (v NullableTokenizerkeyword) Get() *Tokenizerkeyword {
	return v.value
}

func (v *NullableTokenizerkeyword) Set(val *Tokenizerkeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizerkeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizerkeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizerkeyword(val *Tokenizerkeyword) *NullableTokenizerkeyword {
	return &NullableTokenizerkeyword{value: val, isSet: true}
}

func (v NullableTokenizerkeyword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizerkeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


