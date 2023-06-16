// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the ApiAtlasFTSAnalyzersCharFiltersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAtlasFTSAnalyzersCharFiltersInner{}

// ApiAtlasFTSAnalyzersCharFiltersInner struct for ApiAtlasFTSAnalyzersCharFiltersInner
type ApiAtlasFTSAnalyzersCharFiltersInner struct {
	// The HTML tags that you want to exclude from filtering.
	IgnoredTags []string `json:"ignoredTags,omitempty"`
	// Human-readable label that identifies this character filter type.
	Type     *string                                       `json:"type,omitempty"`
	Mappings *ApiAtlasFTSAnalyzersCharFiltersInnerMappings `json:"mappings,omitempty"`
}

// NewApiAtlasFTSAnalyzersCharFiltersInner instantiates a new ApiAtlasFTSAnalyzersCharFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFTSAnalyzersCharFiltersInner() *ApiAtlasFTSAnalyzersCharFiltersInner {
	this := ApiAtlasFTSAnalyzersCharFiltersInner{}
	return &this
}

// NewApiAtlasFTSAnalyzersCharFiltersInnerWithDefaults instantiates a new ApiAtlasFTSAnalyzersCharFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFTSAnalyzersCharFiltersInnerWithDefaults() *ApiAtlasFTSAnalyzersCharFiltersInner {
	this := ApiAtlasFTSAnalyzersCharFiltersInner{}
	return &this
}

// GetIgnoredTags returns the IgnoredTags field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetIgnoredTags() []string {
	if o == nil || IsNil(o.IgnoredTags) {
		var ret []string
		return ret
	}
	return o.IgnoredTags
}

// GetIgnoredTagsOk returns a tuple with the IgnoredTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetIgnoredTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoredTags) {
		return nil, false
	}
	return o.IgnoredTags, true
}

// HasIgnoredTags returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) HasIgnoredTags() bool {
	if o != nil && !IsNil(o.IgnoredTags) {
		return true
	}

	return false
}

// SetIgnoredTags gets a reference to the given []string and assigns it to the IgnoredTags field.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) SetIgnoredTags(v []string) {
	o.IgnoredTags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) SetType(v string) {
	o.Type = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetMappings() ApiAtlasFTSAnalyzersCharFiltersInnerMappings {
	if o == nil || IsNil(o.Mappings) {
		var ret ApiAtlasFTSAnalyzersCharFiltersInnerMappings
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) GetMappingsOk() (*ApiAtlasFTSAnalyzersCharFiltersInnerMappings, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given ApiAtlasFTSAnalyzersCharFiltersInnerMappings and assigns it to the Mappings field.
func (o *ApiAtlasFTSAnalyzersCharFiltersInner) SetMappings(v ApiAtlasFTSAnalyzersCharFiltersInnerMappings) {
	o.Mappings = &v
}

func (o ApiAtlasFTSAnalyzersCharFiltersInner) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiAtlasFTSAnalyzersCharFiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoredTags) {
		toSerialize["ignoredTags"] = o.IgnoredTags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return toSerialize, nil
}

type NullableApiAtlasFTSAnalyzersCharFiltersInner struct {
	value *ApiAtlasFTSAnalyzersCharFiltersInner
	isSet bool
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInner) Get() *ApiAtlasFTSAnalyzersCharFiltersInner {
	return v.value
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInner) Set(val *ApiAtlasFTSAnalyzersCharFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasFTSAnalyzersCharFiltersInner(val *ApiAtlasFTSAnalyzersCharFiltersInner) *NullableApiAtlasFTSAnalyzersCharFiltersInner {
	return &NullableApiAtlasFTSAnalyzersCharFiltersInner{value: val, isSet: true}
}

func (v NullableApiAtlasFTSAnalyzersCharFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasFTSAnalyzersCharFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
