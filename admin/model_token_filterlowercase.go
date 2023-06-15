// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterlowercase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterlowercase{}

// TokenFilterlowercase Filter that normalizes token text to lowercase.
type TokenFilterlowercase struct {
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterlowercase instantiates a new TokenFilterlowercase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterlowercase(type_ string) *TokenFilterlowercase {
	this := TokenFilterlowercase{}
	this.Type = type_
	return &this
}

// NewTokenFilterlowercaseWithDefaults instantiates a new TokenFilterlowercase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterlowercaseWithDefaults() *TokenFilterlowercase {
	this := TokenFilterlowercase{}
	return &this
}

// GetType returns the Type field value
func (o *TokenFilterlowercase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterlowercase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterlowercase) SetType(v string) {
	o.Type = v
}

func (o TokenFilterlowercase) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterlowercase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterlowercase struct {
	value *TokenFilterlowercase
	isSet bool
}

func (v NullableTokenFilterlowercase) Get() *TokenFilterlowercase {
	return v.value
}

func (v *NullableTokenFilterlowercase) Set(val *TokenFilterlowercase) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterlowercase) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterlowercase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterlowercase(val *TokenFilterlowercase) *NullableTokenFilterlowercase {
	return &NullableTokenFilterlowercase{value: val, isSet: true}
}

func (v NullableTokenFilterlowercase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterlowercase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
