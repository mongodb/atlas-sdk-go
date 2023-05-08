/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TokenizeruaxUrlEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizeruaxUrlEmail{}

// TokenizeruaxUrlEmail Tokenizer that creates tokens from URLs and email addresses. Although this tokenizer uses word break rules from the Unicode Text Segmentation algorithm, we recommend using it only when the indexed field value includes URLs and email addresses. For fields that don't include URLs or email addresses, use the **standard** tokenizer to create tokens based on word break rules.
type TokenizeruaxUrlEmail struct {
	// Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens.
	MaxTokenLength *int32 `json:"maxTokenLength,omitempty"`
	// Human-readable label that identifies this tokenizer type.
	Type string `json:"type"`
}

// NewTokenizeruaxUrlEmail instantiates a new TokenizeruaxUrlEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizeruaxUrlEmail(type_ string) *TokenizeruaxUrlEmail {
	this := TokenizeruaxUrlEmail{}
	var maxTokenLength int32 = 255
	this.MaxTokenLength = &maxTokenLength
	this.Type = type_
	return &this
}

// NewTokenizeruaxUrlEmailWithDefaults instantiates a new TokenizeruaxUrlEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizeruaxUrlEmailWithDefaults() *TokenizeruaxUrlEmail {
	this := TokenizeruaxUrlEmail{}
	var maxTokenLength int32 = 255
	this.MaxTokenLength = &maxTokenLength
	return &this
}

// GetMaxTokenLength returns the MaxTokenLength field value if set, zero value otherwise.
func (o *TokenizeruaxUrlEmail) GetMaxTokenLength() int32 {
	if o == nil || IsNil(o.MaxTokenLength) {
		var ret int32
		return ret
	}
	return *o.MaxTokenLength
}

// GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizeruaxUrlEmail) GetMaxTokenLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTokenLength) {
		return nil, false
	}
	return o.MaxTokenLength, true
}

// HasMaxTokenLength returns a boolean if a field has been set.
func (o *TokenizeruaxUrlEmail) HasMaxTokenLength() bool {
	if o != nil && !IsNil(o.MaxTokenLength) {
		return true
	}

	return false
}

// SetMaxTokenLength gets a reference to the given int32 and assigns it to the MaxTokenLength field.
func (o *TokenizeruaxUrlEmail) SetMaxTokenLength(v int32) {
	o.MaxTokenLength = &v
}

// GetType returns the Type field value
func (o *TokenizeruaxUrlEmail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenizeruaxUrlEmail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenizeruaxUrlEmail) SetType(v string) {
	o.Type = v
}

func (o TokenizeruaxUrlEmail) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenizeruaxUrlEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxTokenLength) {
		toSerialize["maxTokenLength"] = o.MaxTokenLength
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenizeruaxUrlEmail struct {
	value *TokenizeruaxUrlEmail
	isSet bool
}

func (v NullableTokenizeruaxUrlEmail) Get() *TokenizeruaxUrlEmail {
	return v.value
}

func (v *NullableTokenizeruaxUrlEmail) Set(val *TokenizeruaxUrlEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizeruaxUrlEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizeruaxUrlEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizeruaxUrlEmail(val *TokenizeruaxUrlEmail) *NullableTokenizeruaxUrlEmail {
	return &NullableTokenizeruaxUrlEmail{value: val, isSet: true}
}

func (v NullableTokenizeruaxUrlEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizeruaxUrlEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


