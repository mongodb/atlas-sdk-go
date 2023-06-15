// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the FTSAnalyzersCharFiltersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSAnalyzersCharFiltersInner{}

// FTSAnalyzersCharFiltersInner struct for FTSAnalyzersCharFiltersInner
type FTSAnalyzersCharFiltersInner struct {
	// The HTML tags that you want to exclude from filtering.
	IgnoredTags []string `json:"ignoredTags,omitempty"`
	// Human-readable label that identifies this character filter type.
	Type     *string                               `json:"type,omitempty"`
	Mappings *FTSAnalyzersCharFiltersInnerMappings `json:"mappings,omitempty"`
}

// NewFTSAnalyzersCharFiltersInner instantiates a new FTSAnalyzersCharFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSAnalyzersCharFiltersInner() *FTSAnalyzersCharFiltersInner {
	this := FTSAnalyzersCharFiltersInner{}
	return &this
}

// NewFTSAnalyzersCharFiltersInnerWithDefaults instantiates a new FTSAnalyzersCharFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSAnalyzersCharFiltersInnerWithDefaults() *FTSAnalyzersCharFiltersInner {
	this := FTSAnalyzersCharFiltersInner{}
	return &this
}

// GetIgnoredTags returns the IgnoredTags field value if set, zero value otherwise.
func (o *FTSAnalyzersCharFiltersInner) GetIgnoredTags() []string {
	if o == nil || IsNil(o.IgnoredTags) {
		var ret []string
		return ret
	}
	return o.IgnoredTags
}

// GetIgnoredTagsOk returns a tuple with the IgnoredTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersCharFiltersInner) GetIgnoredTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoredTags) {
		return nil, false
	}
	return o.IgnoredTags, true
}

// HasIgnoredTags returns a boolean if a field has been set.
func (o *FTSAnalyzersCharFiltersInner) HasIgnoredTags() bool {
	if o != nil && !IsNil(o.IgnoredTags) {
		return true
	}

	return false
}

// SetIgnoredTags gets a reference to the given []string and assigns it to the IgnoredTags field.
func (o *FTSAnalyzersCharFiltersInner) SetIgnoredTags(v []string) {
	o.IgnoredTags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FTSAnalyzersCharFiltersInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersCharFiltersInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FTSAnalyzersCharFiltersInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FTSAnalyzersCharFiltersInner) SetType(v string) {
	o.Type = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *FTSAnalyzersCharFiltersInner) GetMappings() FTSAnalyzersCharFiltersInnerMappings {
	if o == nil || IsNil(o.Mappings) {
		var ret FTSAnalyzersCharFiltersInnerMappings
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersCharFiltersInner) GetMappingsOk() (*FTSAnalyzersCharFiltersInnerMappings, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *FTSAnalyzersCharFiltersInner) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given FTSAnalyzersCharFiltersInnerMappings and assigns it to the Mappings field.
func (o *FTSAnalyzersCharFiltersInner) SetMappings(v FTSAnalyzersCharFiltersInnerMappings) {
	o.Mappings = &v
}

func (o FTSAnalyzersCharFiltersInner) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSAnalyzersCharFiltersInner) ToMap() (map[string]interface{}, error) {
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

type NullableFTSAnalyzersCharFiltersInner struct {
	value *FTSAnalyzersCharFiltersInner
	isSet bool
}

func (v NullableFTSAnalyzersCharFiltersInner) Get() *FTSAnalyzersCharFiltersInner {
	return v.value
}

func (v *NullableFTSAnalyzersCharFiltersInner) Set(val *FTSAnalyzersCharFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersCharFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersCharFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersCharFiltersInner(val *FTSAnalyzersCharFiltersInner) *NullableFTSAnalyzersCharFiltersInner {
	return &NullableFTSAnalyzersCharFiltersInner{value: val, isSet: true}
}

func (v NullableFTSAnalyzersCharFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersCharFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
