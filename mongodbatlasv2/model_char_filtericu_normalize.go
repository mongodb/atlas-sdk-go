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

// checks if the CharFiltericuNormalize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharFiltericuNormalize{}

// CharFiltericuNormalize Filter that processes normalized text with the ICU Normalizer. It is based on Lucene's ICUNormalizer2CharFilter.
type CharFiltericuNormalize struct {
	// Human-readable label that identifies this character filter type.
	Type string `json:"type"`
}

// NewCharFiltericuNormalize instantiates a new CharFiltericuNormalize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharFiltericuNormalize(type_ string) *CharFiltericuNormalize {
	this := CharFiltericuNormalize{}
	this.Type = type_
	return &this
}

// NewCharFiltericuNormalizeWithDefaults instantiates a new CharFiltericuNormalize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharFiltericuNormalizeWithDefaults() *CharFiltericuNormalize {
	this := CharFiltericuNormalize{}
	return &this
}

// GetType returns the Type field value
func (o *CharFiltericuNormalize) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharFiltericuNormalize) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharFiltericuNormalize) SetType(v string) {
	o.Type = v
}

func (o CharFiltericuNormalize) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CharFiltericuNormalize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCharFiltericuNormalize struct {
	value *CharFiltericuNormalize
	isSet bool
}

func (v NullableCharFiltericuNormalize) Get() *CharFiltericuNormalize {
	return v.value
}

func (v *NullableCharFiltericuNormalize) Set(val *CharFiltericuNormalize) {
	v.value = val
	v.isSet = true
}

func (v NullableCharFiltericuNormalize) IsSet() bool {
	return v.isSet
}

func (v *NullableCharFiltericuNormalize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharFiltericuNormalize(val *CharFiltericuNormalize) *NullableCharFiltericuNormalize {
	return &NullableCharFiltericuNormalize{value: val, isSet: true}
}

func (v NullableCharFiltericuNormalize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharFiltericuNormalize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


