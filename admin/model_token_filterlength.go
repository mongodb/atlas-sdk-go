// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilterlength type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterlength{}

// TokenFilterlength Filter that removes tokens that are too short or too long.
type TokenFilterlength struct {
	// Number that specifies the maximum length of a token. Value must be greater than or equal to **min**.
	Max *int `json:"max,omitempty"`
	// Number that specifies the minimum length of a token. This value must be less than or equal to **max**.
	Min *int `json:"min,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterlength instantiates a new TokenFilterlength object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterlength(type_ string) *TokenFilterlength {
	this := TokenFilterlength{}
	var max int = 255
	this.Max = &max
	var min int = 0
	this.Min = &min
	this.Type = type_
	return &this
}

// NewTokenFilterlengthWithDefaults instantiates a new TokenFilterlength object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterlengthWithDefaults() *TokenFilterlength {
	this := TokenFilterlength{}
	var max int = 255
	this.Max = &max
	var min int = 0
	this.Min = &min
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *TokenFilterlength) GetMax() int {
	if o == nil || IsNil(o.Max) {
		var ret int
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterlength) GetMaxOk() (*int, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *TokenFilterlength) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int and assigns it to the Max field.
func (o *TokenFilterlength) SetMax(v int) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *TokenFilterlength) GetMin() int {
	if o == nil || IsNil(o.Min) {
		var ret int
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterlength) GetMinOk() (*int, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *TokenFilterlength) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int and assigns it to the Min field.
func (o *TokenFilterlength) SetMin(v int) {
	o.Min = &v
}

// GetType returns the Type field value
func (o *TokenFilterlength) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterlength) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterlength) SetType(v string) {
	o.Type = v
}

func (o TokenFilterlength) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilterlength) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterlength struct {
	value *TokenFilterlength
	isSet bool
}

func (v NullableTokenFilterlength) Get() *TokenFilterlength {
	return v.value
}

func (v *NullableTokenFilterlength) Set(val *TokenFilterlength) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterlength) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterlength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterlength(val *TokenFilterlength) *NullableTokenFilterlength {
	return &NullableTokenFilterlength{value: val, isSet: true}
}

func (v NullableTokenFilterlength) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterlength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
