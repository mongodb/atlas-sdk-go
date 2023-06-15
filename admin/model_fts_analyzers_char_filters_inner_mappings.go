// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the FTSAnalyzersCharFiltersInnerMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSAnalyzersCharFiltersInnerMappings{}

// FTSAnalyzersCharFiltersInnerMappings Comma-separated list of mappings. A mapping indicates that one character or group of characters should be substituted for another, using the following format:  `<original> : <replacement>`
type FTSAnalyzersCharFiltersInnerMappings struct {
	AdditionalPropertiesField *string `json:"additionalProperties,omitempty"`
}

// NewFTSAnalyzersCharFiltersInnerMappings instantiates a new FTSAnalyzersCharFiltersInnerMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSAnalyzersCharFiltersInnerMappings() *FTSAnalyzersCharFiltersInnerMappings {
	this := FTSAnalyzersCharFiltersInnerMappings{}
	return &this
}

// NewFTSAnalyzersCharFiltersInnerMappingsWithDefaults instantiates a new FTSAnalyzersCharFiltersInnerMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSAnalyzersCharFiltersInnerMappingsWithDefaults() *FTSAnalyzersCharFiltersInnerMappings {
	this := FTSAnalyzersCharFiltersInnerMappings{}
	return &this
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *FTSAnalyzersCharFiltersInnerMappings) GetAdditionalPropertiesField() string {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret string
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersCharFiltersInnerMappings) GetAdditionalPropertiesFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *FTSAnalyzersCharFiltersInnerMappings) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given string and assigns it to the AdditionalPropertiesField field.
func (o *FTSAnalyzersCharFiltersInnerMappings) SetAdditionalPropertiesField(v string) {
	o.AdditionalPropertiesField = &v
}

func (o FTSAnalyzersCharFiltersInnerMappings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSAnalyzersCharFiltersInnerMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return toSerialize, nil
}

type NullableFTSAnalyzersCharFiltersInnerMappings struct {
	value *FTSAnalyzersCharFiltersInnerMappings
	isSet bool
}

func (v NullableFTSAnalyzersCharFiltersInnerMappings) Get() *FTSAnalyzersCharFiltersInnerMappings {
	return v.value
}

func (v *NullableFTSAnalyzersCharFiltersInnerMappings) Set(val *FTSAnalyzersCharFiltersInnerMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersCharFiltersInnerMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersCharFiltersInnerMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersCharFiltersInnerMappings(val *FTSAnalyzersCharFiltersInnerMappings) *NullableFTSAnalyzersCharFiltersInnerMappings {
	return &NullableFTSAnalyzersCharFiltersInnerMappings{value: val, isSet: true}
}

func (v NullableFTSAnalyzersCharFiltersInnerMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersCharFiltersInnerMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
