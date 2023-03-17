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

// checks if the TokenizernGram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizernGram{}

// TokenizernGram Tokenizer that splits input into text chunks, or \"n-grams\", of into given sizes. You can't use the nGram tokenizer in synonym or autocomplete mapping definitions.
type TokenizernGram struct {
	// Characters to include in the longest token that Atlas Search creates.
	MaxGram int32 `json:"maxGram"`
	// Characters to include in the shortest token that Atlas Search creates.
	MinGram int32 `json:"minGram"`
	// Human-readable label that identifies this tokenizer type.
	Type string `json:"type"`
}

// NewTokenizernGram instantiates a new TokenizernGram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizernGram(maxGram int32, minGram int32, type_ string) *TokenizernGram {
	this := TokenizernGram{}
	this.MaxGram = maxGram
	this.MinGram = minGram
	this.Type = type_
	return &this
}

// NewTokenizernGramWithDefaults instantiates a new TokenizernGram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizernGramWithDefaults() *TokenizernGram {
	this := TokenizernGram{}
	return &this
}

// GetMaxGram returns the MaxGram field value
func (o *TokenizernGram) GetMaxGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxGram
}

// GetMaxGramOk returns a tuple with the MaxGram field value
// and a boolean to check if the value has been set.
func (o *TokenizernGram) GetMaxGramOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGram, true
}

// SetMaxGram sets field value
func (o *TokenizernGram) SetMaxGram(v int32) {
	o.MaxGram = v
}

// GetMinGram returns the MinGram field value
func (o *TokenizernGram) GetMinGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinGram
}

// GetMinGramOk returns a tuple with the MinGram field value
// and a boolean to check if the value has been set.
func (o *TokenizernGram) GetMinGramOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinGram, true
}

// SetMinGram sets field value
func (o *TokenizernGram) SetMinGram(v int32) {
	o.MinGram = v
}

// GetType returns the Type field value
func (o *TokenizernGram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenizernGram) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenizernGram) SetType(v string) {
	o.Type = v
}

func (o TokenizernGram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizernGram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxGram"] = o.MaxGram
	toSerialize["minGram"] = o.MinGram
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenizernGram struct {
	value *TokenizernGram
	isSet bool
}

func (v NullableTokenizernGram) Get() *TokenizernGram {
	return v.value
}

func (v *NullableTokenizernGram) Set(val *TokenizernGram) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizernGram) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizernGram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizernGram(val *TokenizernGram) *NullableTokenizernGram {
	return &NullableTokenizernGram{value: val, isSet: true}
}

func (v NullableTokenizernGram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizernGram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


