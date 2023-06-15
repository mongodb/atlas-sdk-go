// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TokenFiltericuNormalizer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFiltericuNormalizer{}

// TokenFiltericuNormalizer Filter that normalizes tokens using a standard Unicode Normalization Mode.
type TokenFiltericuNormalizer struct {
	// Normalization form to apply.
	NormalizationForm *string `json:"normalizationForm,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFiltericuNormalizer instantiates a new TokenFiltericuNormalizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFiltericuNormalizer(type_ string) *TokenFiltericuNormalizer {
	this := TokenFiltericuNormalizer{}
	var normalizationForm string = "nfc"
	this.NormalizationForm = &normalizationForm
	this.Type = type_
	return &this
}

// NewTokenFiltericuNormalizerWithDefaults instantiates a new TokenFiltericuNormalizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFiltericuNormalizerWithDefaults() *TokenFiltericuNormalizer {
	this := TokenFiltericuNormalizer{}
	var normalizationForm string = "nfc"
	this.NormalizationForm = &normalizationForm
	return &this
}

// GetNormalizationForm returns the NormalizationForm field value if set, zero value otherwise.
func (o *TokenFiltericuNormalizer) GetNormalizationForm() string {
	if o == nil || IsNil(o.NormalizationForm) {
		var ret string
		return ret
	}
	return *o.NormalizationForm
}

// GetNormalizationFormOk returns a tuple with the NormalizationForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFiltericuNormalizer) GetNormalizationFormOk() (*string, bool) {
	if o == nil || IsNil(o.NormalizationForm) {
		return nil, false
	}
	return o.NormalizationForm, true
}

// HasNormalizationForm returns a boolean if a field has been set.
func (o *TokenFiltericuNormalizer) HasNormalizationForm() bool {
	if o != nil && !IsNil(o.NormalizationForm) {
		return true
	}

	return false
}

// SetNormalizationForm gets a reference to the given string and assigns it to the NormalizationForm field.
func (o *TokenFiltericuNormalizer) SetNormalizationForm(v string) {
	o.NormalizationForm = &v
}

// GetType returns the Type field value
func (o *TokenFiltericuNormalizer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFiltericuNormalizer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFiltericuNormalizer) SetType(v string) {
	o.Type = v
}

func (o TokenFiltericuNormalizer) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TokenFiltericuNormalizer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NormalizationForm) {
		toSerialize["normalizationForm"] = o.NormalizationForm
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFiltericuNormalizer struct {
	value *TokenFiltericuNormalizer
	isSet bool
}

func (v NullableTokenFiltericuNormalizer) Get() *TokenFiltericuNormalizer {
	return v.value
}

func (v *NullableTokenFiltericuNormalizer) Set(val *TokenFiltericuNormalizer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFiltericuNormalizer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFiltericuNormalizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFiltericuNormalizer(val *TokenFiltericuNormalizer) *NullableTokenFiltericuNormalizer {
	return &NullableTokenFiltericuNormalizer{value: val, isSet: true}
}

func (v NullableTokenFiltericuNormalizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFiltericuNormalizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
