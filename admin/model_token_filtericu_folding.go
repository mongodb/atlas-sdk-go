// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFiltericuFolding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFiltericuFolding{}

// TokenFiltericuFolding Filter that applies character folding from Unicode Technical Report #30.
type TokenFiltericuFolding struct {
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltericuFolding instantiates a new TokenFiltericuFolding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltericuFolding(type_ string) *TokenFiltericuFolding {
	this := TokenFiltericuFolding{}
	this.Type = type_
	return &this
}

// NewTokenFiltericuFoldingWithDefaults instantiates a new TokenFiltericuFolding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltericuFoldingWithDefaults() *TokenFiltericuFolding {
	this := TokenFiltericuFolding{}
	return &this
}

// GetType returns the Type field value
func (o *TokenFiltericuFolding) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltericuFolding) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltericuFolding) SetType(v string) {
	o.Type = v
}

func (o TokenFiltericuFolding) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFiltericuFolding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFiltericuFolding struct {
	value *TokenFiltericuFolding
	isSet bool
}

func (v NullableTokenFiltericuFolding) Get() *TokenFiltericuFolding {
	return v.value
}

func (v *NullableTokenFiltericuFolding) Set(val *TokenFiltericuFolding) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltericuFolding) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltericuFolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltericuFolding(val *TokenFiltericuFolding) *NullableTokenFiltericuFolding {
	return &NullableTokenFiltericuFolding{value: val, isSet: true}
}

func (v NullableTokenFiltericuFolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltericuFolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
