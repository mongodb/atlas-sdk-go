// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterdaitchMokotoffSoundex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterdaitchMokotoffSoundex{}

// TokenFilterdaitchMokotoffSoundex Filter that creates tokens for words that sound the same based on the Daitch-Mokotoff Soundex phonetic algorithm. This filter can generate multiple encodings for each input, where each encoded token is a 6 digit number.  **NOTE**: Don't use the **daitchMokotoffSoundex** token filter in:  -Synonym or autocomplete mapping definitions - Operators where **fuzzy** is enabled. Atlas Search supports the **fuzzy** option only for the **autocomplete**, **term**, and **text** operators.
type TokenFilterdaitchMokotoffSoundex struct {
	// Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose `include` if you want to support queries on both the original tokens as well as the converted forms.   Choose `omit` if you want to query only on the converted forms of the original tokens.
	OriginalTokens *string `json:"originalTokens,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterdaitchMokotoffSoundex instantiates a new TokenFilterdaitchMokotoffSoundex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterdaitchMokotoffSoundex(type_ string) *TokenFilterdaitchMokotoffSoundex {
	this := TokenFilterdaitchMokotoffSoundex{}
	var originalTokens string = "include"
	this.OriginalTokens = &originalTokens
	this.Type = type_
	return &this
}

// NewTokenFilterdaitchMokotoffSoundexWithDefaults instantiates a new TokenFilterdaitchMokotoffSoundex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterdaitchMokotoffSoundexWithDefaults() *TokenFilterdaitchMokotoffSoundex {
	this := TokenFilterdaitchMokotoffSoundex{}
	var originalTokens string = "include"
	this.OriginalTokens = &originalTokens
	return &this
}

// GetOriginalTokens returns the OriginalTokens field value if set, zero value otherwise.
func (o *TokenFilterdaitchMokotoffSoundex) GetOriginalTokens() string {
	if o == nil || IsNil(o.OriginalTokens) {
		var ret string
		return ret
	}
	return *o.OriginalTokens
}

// GetOriginalTokensOk returns a tuple with the OriginalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterdaitchMokotoffSoundex) GetOriginalTokensOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalTokens) {
		return nil, false
	}
	return o.OriginalTokens, true
}

// HasOriginalTokens returns a boolean if a field has been set.
func (o *TokenFilterdaitchMokotoffSoundex) HasOriginalTokens() bool {
	if o != nil && !IsNil(o.OriginalTokens) {
		return true
	}

	return false
}

// SetOriginalTokens gets a reference to the given string and assigns it to the OriginalTokens field.
func (o *TokenFilterdaitchMokotoffSoundex) SetOriginalTokens(v string) {
	o.OriginalTokens = &v
}

// GetType returns the Type field value
func (o *TokenFilterdaitchMokotoffSoundex) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterdaitchMokotoffSoundex) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterdaitchMokotoffSoundex) SetType(v string) {
	o.Type = v
}

func (o TokenFilterdaitchMokotoffSoundex) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterdaitchMokotoffSoundex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalTokens) {
		toSerialize["originalTokens"] = o.OriginalTokens
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterdaitchMokotoffSoundex struct {
	value *TokenFilterdaitchMokotoffSoundex
	isSet bool
}

func (v NullableTokenFilterdaitchMokotoffSoundex) Get() *TokenFilterdaitchMokotoffSoundex {
	return v.value
}

func (v *NullableTokenFilterdaitchMokotoffSoundex) Set(val *TokenFilterdaitchMokotoffSoundex) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterdaitchMokotoffSoundex) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterdaitchMokotoffSoundex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterdaitchMokotoffSoundex(val *TokenFilterdaitchMokotoffSoundex) *NullableTokenFilterdaitchMokotoffSoundex {
	return &NullableTokenFilterdaitchMokotoffSoundex{value: val, isSet: true}
}

func (v NullableTokenFilterdaitchMokotoffSoundex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterdaitchMokotoffSoundex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
