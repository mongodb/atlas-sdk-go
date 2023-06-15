// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFiltertrim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFiltertrim{}

// TokenFiltertrim Filter that trims leading and trailing whitespace from tokens.
type TokenFiltertrim struct {
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltertrim instantiates a new TokenFiltertrim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltertrim(type_ string) *TokenFiltertrim {
	this := TokenFiltertrim{}
	this.Type = type_
	return &this
}

// NewTokenFiltertrimWithDefaults instantiates a new TokenFiltertrim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltertrimWithDefaults() *TokenFiltertrim {
	this := TokenFiltertrim{}
	return &this
}

// GetType returns the Type field value
func (o *TokenFiltertrim) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltertrim) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltertrim) SetType(v string) {
	o.Type = v
}

func (o TokenFiltertrim) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFiltertrim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFiltertrim struct {
	value *TokenFiltertrim
	isSet bool
}

func (v NullableTokenFiltertrim) Get() *TokenFiltertrim {
	return v.value
}

func (v *NullableTokenFiltertrim) Set(val *TokenFiltertrim) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltertrim) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltertrim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltertrim(val *TokenFiltertrim) *NullableTokenFiltertrim {
	return &NullableTokenFiltertrim{value: val, isSet: true}
}

func (v NullableTokenFiltertrim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltertrim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
