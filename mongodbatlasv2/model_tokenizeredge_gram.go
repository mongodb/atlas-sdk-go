/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TokenizeredgeGram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizeredgeGram{}

// TokenizeredgeGram Tokenizer that splits input from the left side, or \"edge\", of a text input into n-grams of given sizes. You can't use the edgeGram tokenizer in synonym or autocomplete mapping definitions.
type TokenizeredgeGram struct {
	// Characters to include in the longest token that Atlas Search creates.
	MaxGram int32 `json:"maxGram"`
	// Characters to include in the shortest token that Atlas Search creates.
	MinGram int32 `json:"minGram"`
	// Human-readable label that identifies this tokenizer type.
	Type string `json:"type"`
}

// NewTokenizeredgeGram instantiates a new TokenizeredgeGram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizeredgeGram(maxGram int32, minGram int32, type_ string) *TokenizeredgeGram {
	this := TokenizeredgeGram{}
	this.MaxGram = maxGram
	this.MinGram = minGram
	this.Type = type_
	return &this
}

// NewTokenizeredgeGramWithDefaults instantiates a new TokenizeredgeGram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizeredgeGramWithDefaults() *TokenizeredgeGram {
	this := TokenizeredgeGram{}
	return &this
}

// GetMaxGram returns the MaxGram field value
func (o *TokenizeredgeGram) GetMaxGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxGram
}

// GetMaxGramOk returns a tuple with the MaxGram field value
// and a boolean to check if the value has been set.
func (o *TokenizeredgeGram) GetMaxGramOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGram, true
}

// SetMaxGram sets field value
func (o *TokenizeredgeGram) SetMaxGram(v int32) {
	o.MaxGram = v
}

// GetMinGram returns the MinGram field value
func (o *TokenizeredgeGram) GetMinGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinGram
}

// GetMinGramOk returns a tuple with the MinGram field value
// and a boolean to check if the value has been set.
func (o *TokenizeredgeGram) GetMinGramOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinGram, true
}

// SetMinGram sets field value
func (o *TokenizeredgeGram) SetMinGram(v int32) {
	o.MinGram = v
}

// GetType returns the Type field value
func (o *TokenizeredgeGram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenizeredgeGram) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenizeredgeGram) SetType(v string) {
	o.Type = v
}

func (o TokenizeredgeGram) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenizeredgeGram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxGram"] = o.MaxGram
	toSerialize["minGram"] = o.MinGram
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenizeredgeGram struct {
	value *TokenizeredgeGram
	isSet bool
}

func (v NullableTokenizeredgeGram) Get() *TokenizeredgeGram {
	return v.value
}

func (v *NullableTokenizeredgeGram) Set(val *TokenizeredgeGram) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizeredgeGram) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizeredgeGram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizeredgeGram(val *TokenizeredgeGram) *NullableTokenizeredgeGram {
	return &NullableTokenizeredgeGram{value: val, isSet: true}
}

func (v NullableTokenizeredgeGram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizeredgeGram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


