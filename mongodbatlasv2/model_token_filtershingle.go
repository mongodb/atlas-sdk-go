/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TokenFiltershingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFiltershingle{}

// TokenFiltershingle Filter that constructs shingles (token n-grams) from a series of tokens. You can't use this token filter in synonym or autocomplete mapping definitions.
type TokenFiltershingle struct {
	// Value that specifies the maximum number of tokens per shingle. This value must be greater than or equal to **minShingleSize**.
	MaxShingleSize int32 `json:"maxShingleSize"`
	// Value that specifies the minimum number of tokens per shingle. This value must be less than or equal to **maxShingleSize**.
	MinShingleSize int32 `json:"minShingleSize"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltershingle instantiates a new TokenFiltershingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltershingle(maxShingleSize int32, minShingleSize int32, type_ string) *TokenFiltershingle {
	this := TokenFiltershingle{}
	this.MaxShingleSize = maxShingleSize
	this.MinShingleSize = minShingleSize
	this.Type = type_
	return &this
}

// NewTokenFiltershingleWithDefaults instantiates a new TokenFiltershingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltershingleWithDefaults() *TokenFiltershingle {
	this := TokenFiltershingle{}
	return &this
}

// GetMaxShingleSize returns the MaxShingleSize field value
func (o *TokenFiltershingle) GetMaxShingleSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxShingleSize
}

// GetMaxShingleSizeOk returns a tuple with the MaxShingleSize field value
// and a boolean to check if the value has been set.
func (o *TokenFiltershingle) GetMaxShingleSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxShingleSize, true
}

// SetMaxShingleSize sets field value
func (o *TokenFiltershingle) SetMaxShingleSize(v int32) {
	o.MaxShingleSize = v
}

// GetMinShingleSize returns the MinShingleSize field value
func (o *TokenFiltershingle) GetMinShingleSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinShingleSize
}

// GetMinShingleSizeOk returns a tuple with the MinShingleSize field value
// and a boolean to check if the value has been set.
func (o *TokenFiltershingle) GetMinShingleSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinShingleSize, true
}

// SetMinShingleSize sets field value
func (o *TokenFiltershingle) SetMinShingleSize(v int32) {
	o.MinShingleSize = v
}

// GetType returns the Type field value
func (o *TokenFiltershingle) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltershingle) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltershingle) SetType(v string) {
	o.Type = v
}

func (o TokenFiltershingle) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFiltershingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxShingleSize"] = o.MaxShingleSize
	toSerialize["minShingleSize"] = o.MinShingleSize
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFiltershingle struct {
	value *TokenFiltershingle
	isSet bool
}

func (v NullableTokenFiltershingle) Get() *TokenFiltershingle {
	return v.value
}

func (v *NullableTokenFiltershingle) Set(val *TokenFiltershingle) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltershingle) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltershingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltershingle(val *TokenFiltershingle) *NullableTokenFiltershingle {
	return &NullableTokenFiltershingle{value: val, isSet: true}
}

func (v NullableTokenFiltershingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltershingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


