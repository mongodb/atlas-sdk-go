// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ApiAtlasFTSAnalyzersCharFiltersInnerMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAtlasFTSAnalyzersCharFiltersInnerMappings{}

// ApiAtlasFTSAnalyzersCharFiltersInnerMappings Comma-separated list of mappings. A mapping indicates that one character or group of characters should be substituted for another, using the following format:  `<original> : <replacement>`
type ApiAtlasFTSAnalyzersCharFiltersInnerMappings struct {
	AdditionalPropertiesField *string `json:"additionalProperties,omitempty"`
}

// NewApiAtlasFTSAnalyzersCharFiltersInnerMappings instantiates a new ApiAtlasFTSAnalyzersCharFiltersInnerMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFTSAnalyzersCharFiltersInnerMappings() *ApiAtlasFTSAnalyzersCharFiltersInnerMappings {
	this := ApiAtlasFTSAnalyzersCharFiltersInnerMappings{}
	return &this
}

// NewApiAtlasFTSAnalyzersCharFiltersInnerMappingsWithDefaults instantiates a new ApiAtlasFTSAnalyzersCharFiltersInnerMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFTSAnalyzersCharFiltersInnerMappingsWithDefaults() *ApiAtlasFTSAnalyzersCharFiltersInnerMappings {
	this := ApiAtlasFTSAnalyzersCharFiltersInnerMappings{}
	return &this
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) GetAdditionalPropertiesField() string {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret string
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) GetAdditionalPropertiesFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given string and assigns it to the AdditionalPropertiesField field.
func (o *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) SetAdditionalPropertiesField(v string) {
	o.AdditionalPropertiesField = &v
}

func (o ApiAtlasFTSAnalyzersCharFiltersInnerMappings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiAtlasFTSAnalyzersCharFiltersInnerMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return toSerialize, nil
}

type NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings struct {
	value *ApiAtlasFTSAnalyzersCharFiltersInnerMappings
	isSet bool
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) Get() *ApiAtlasFTSAnalyzersCharFiltersInnerMappings {
	return v.value
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) Set(val *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasFTSAnalyzersCharFiltersInnerMappings(val *ApiAtlasFTSAnalyzersCharFiltersInnerMappings) *NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings {
	return &NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings{value: val, isSet: true}
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInnerMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
