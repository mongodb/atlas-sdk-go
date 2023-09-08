// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ToggleableFeatureFlag struct for ToggleableFeatureFlag
type ToggleableFeatureFlag struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Flag    *string `json:"flag,omitempty"`
}

// NewToggleableFeatureFlag instantiates a new ToggleableFeatureFlag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToggleableFeatureFlag() *ToggleableFeatureFlag {
	this := ToggleableFeatureFlag{}
	return &this
}

// NewToggleableFeatureFlagWithDefaults instantiates a new ToggleableFeatureFlag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToggleableFeatureFlagWithDefaults() *ToggleableFeatureFlag {
	this := ToggleableFeatureFlag{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ToggleableFeatureFlag) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleableFeatureFlag) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ToggleableFeatureFlag) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ToggleableFeatureFlag) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise
func (o *ToggleableFeatureFlag) GetFlag() string {
	if o == nil || IsNil(o.Flag) {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleableFeatureFlag) GetFlagOk() (*string, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}

	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *ToggleableFeatureFlag) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *ToggleableFeatureFlag) SetFlag(v string) {
	o.Flag = &v
}

func (o ToggleableFeatureFlag) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ToggleableFeatureFlag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	return toSerialize, nil
}
