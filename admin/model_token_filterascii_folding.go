// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterasciiFolding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterasciiFolding{}

// TokenFilterasciiFolding Filter that converts alphabetic, numeric, and symbolic unicode characters that are not in the Basic Latin Unicode block to their ASCII equivalents, if available.
type TokenFilterasciiFolding struct {
	// Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose `include` if you want to support queries on both the original tokens as well as the converted forms.   Choose `omit` if you want to query only on the converted forms of the original tokens.
	OriginalTokens *string `json:"originalTokens,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterasciiFolding instantiates a new TokenFilterasciiFolding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterasciiFolding(type_ string) *TokenFilterasciiFolding {
	this := TokenFilterasciiFolding{}
	var originalTokens string = "omit"
	this.OriginalTokens = &originalTokens
	this.Type = type_
	return &this
}

// NewTokenFilterasciiFoldingWithDefaults instantiates a new TokenFilterasciiFolding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterasciiFoldingWithDefaults() *TokenFilterasciiFolding {
	this := TokenFilterasciiFolding{}
	var originalTokens string = "omit"
	this.OriginalTokens = &originalTokens
	return &this
}

// GetOriginalTokens returns the OriginalTokens field value if set, zero value otherwise.
func (o *TokenFilterasciiFolding) GetOriginalTokens() string {
	if o == nil || IsNil(o.OriginalTokens) {
		var ret string
		return ret
	}
	return *o.OriginalTokens
}

// GetOriginalTokensOk returns a tuple with the OriginalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterasciiFolding) GetOriginalTokensOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalTokens) {
		return nil, false
	}
	return o.OriginalTokens, true
}

// HasOriginalTokens returns a boolean if a field has been set.
func (o *TokenFilterasciiFolding) HasOriginalTokens() bool {
	if o != nil && !IsNil(o.OriginalTokens) {
		return true
	}

	return false
}

// SetOriginalTokens gets a reference to the given string and assigns it to the OriginalTokens field.
func (o *TokenFilterasciiFolding) SetOriginalTokens(v string) {
	o.OriginalTokens = &v
}

// GetType returns the Type field value
func (o *TokenFilterasciiFolding) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterasciiFolding) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterasciiFolding) SetType(v string) {
	o.Type = v
}

func (o TokenFilterasciiFolding) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterasciiFolding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalTokens) {
		toSerialize["originalTokens"] = o.OriginalTokens
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterasciiFolding struct {
	value *TokenFilterasciiFolding
	isSet bool
}

func (v NullableTokenFilterasciiFolding) Get() *TokenFilterasciiFolding {
	return v.value
}

func (v *NullableTokenFilterasciiFolding) Set(val *TokenFilterasciiFolding) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterasciiFolding) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterasciiFolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterasciiFolding(val *TokenFilterasciiFolding) *NullableTokenFilterasciiFolding {
	return &NullableTokenFilterasciiFolding{value: val, isSet: true}
}

func (v NullableTokenFilterasciiFolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterasciiFolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
