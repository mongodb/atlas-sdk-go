// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterreverse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterreverse{}

// TokenFilterreverse Filter that reverses each string token.
type TokenFilterreverse struct {
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterreverse instantiates a new TokenFilterreverse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterreverse(type_ string) *TokenFilterreverse {
	this := TokenFilterreverse{}
	this.Type = type_
	return &this
}

// NewTokenFilterreverseWithDefaults instantiates a new TokenFilterreverse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterreverseWithDefaults() *TokenFilterreverse {
	this := TokenFilterreverse{}
	return &this
}

// GetType returns the Type field value
func (o *TokenFilterreverse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterreverse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterreverse) SetType(v string) {
	o.Type = v
}

func (o TokenFilterreverse) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterreverse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterreverse struct {
	value *TokenFilterreverse
	isSet bool
}

func (v NullableTokenFilterreverse) Get() *TokenFilterreverse {
	return v.value
}

func (v *NullableTokenFilterreverse) Set(val *TokenFilterreverse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterreverse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterreverse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterreverse(val *TokenFilterreverse) *NullableTokenFilterreverse {
	return &NullableTokenFilterreverse{value: val, isSet: true}
}

func (v NullableTokenFilterreverse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterreverse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
