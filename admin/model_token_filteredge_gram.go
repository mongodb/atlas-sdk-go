// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFilteredgeGram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilteredgeGram{}

// TokenFilteredgeGram Filter that tokenizes input from the left side, or \"edge\", of a text input into n-grams of configured sizes. You can't use this token filter in synonym or autocomplete mapping definitions.
type TokenFilteredgeGram struct {
	// Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**.
	MaxGram int `json:"maxGram"`
	// Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**.
	MinGram int `json:"minGram"`
	// Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**.
	TermNotInBounds *string `json:"termNotInBounds,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilteredgeGram instantiates a new TokenFilteredgeGram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilteredgeGram(maxGram int, minGram int, type_ string) *TokenFilteredgeGram {
	this := TokenFilteredgeGram{}
	this.MaxGram = maxGram
	this.MinGram = minGram
	var termNotInBounds string = "omit"
	this.TermNotInBounds = &termNotInBounds
	this.Type = type_
	return &this
}

// NewTokenFilteredgeGramWithDefaults instantiates a new TokenFilteredgeGram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilteredgeGramWithDefaults() *TokenFilteredgeGram {
	this := TokenFilteredgeGram{}
	var termNotInBounds string = "omit"
	this.TermNotInBounds = &termNotInBounds
	return &this
}

// GetMaxGram returns the MaxGram field value
func (o *TokenFilteredgeGram) GetMaxGram() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.MaxGram
}

// GetMaxGramOk returns a tuple with the MaxGram field value
// and a boolean to check if the value has been set.
func (o *TokenFilteredgeGram) GetMaxGramOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxGram, true
}

// SetMaxGram sets field value
func (o *TokenFilteredgeGram) SetMaxGram(v int) {
	o.MaxGram = v
}

// GetMinGram returns the MinGram field value
func (o *TokenFilteredgeGram) GetMinGram() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.MinGram
}

// GetMinGramOk returns a tuple with the MinGram field value
// and a boolean to check if the value has been set.
func (o *TokenFilteredgeGram) GetMinGramOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinGram, true
}

// SetMinGram sets field value
func (o *TokenFilteredgeGram) SetMinGram(v int) {
	o.MinGram = v
}

// GetTermNotInBounds returns the TermNotInBounds field value if set, zero value otherwise.
func (o *TokenFilteredgeGram) GetTermNotInBounds() string {
	if o == nil || IsNil(o.TermNotInBounds) {
		var ret string
		return ret
	}
	return *o.TermNotInBounds
}

// GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilteredgeGram) GetTermNotInBoundsOk() (*string, bool) {
	if o == nil || IsNil(o.TermNotInBounds) {
		return nil, false
	}
	return o.TermNotInBounds, true
}

// HasTermNotInBounds returns a boolean if a field has been set.
func (o *TokenFilteredgeGram) HasTermNotInBounds() bool {
	if o != nil && !IsNil(o.TermNotInBounds) {
		return true
	}

	return false
}

// SetTermNotInBounds gets a reference to the given string and assigns it to the TermNotInBounds field.
func (o *TokenFilteredgeGram) SetTermNotInBounds(v string) {
	o.TermNotInBounds = &v
}

// GetType returns the Type field value
func (o *TokenFilteredgeGram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilteredgeGram) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilteredgeGram) SetType(v string) {
	o.Type = v
}

func (o TokenFilteredgeGram) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFilteredgeGram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxGram"] = o.MaxGram
	toSerialize["minGram"] = o.MinGram
	if !IsNil(o.TermNotInBounds) {
		toSerialize["termNotInBounds"] = o.TermNotInBounds
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilteredgeGram struct {
	value *TokenFilteredgeGram
	isSet bool
}

func (v NullableTokenFilteredgeGram) Get() *TokenFilteredgeGram {
	return v.value
}

func (v *NullableTokenFilteredgeGram) Set(val *TokenFilteredgeGram) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilteredgeGram) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilteredgeGram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilteredgeGram(val *TokenFilteredgeGram) *NullableTokenFilteredgeGram {
	return &NullableTokenFilteredgeGram{value: val, isSet: true}
}

func (v NullableTokenFilteredgeGram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilteredgeGram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
