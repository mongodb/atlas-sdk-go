/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ComputeAutoScalingV15 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputeAutoScalingV15{}

// ComputeAutoScalingV15 Options that determine how this cluster handles CPU scaling.
type ComputeAutoScalingV15 struct {
	// Flag that indicates whether someone enabled instance size auto-scaling.  - Set to `true` to enable instance size auto-scaling. If enabled, you must specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize**. - Set to `false` to disable instance size automatic scaling.
	Enabled *bool `json:"enabled,omitempty"`
	MaxInstanceSize *InstanceSize `json:"maxInstanceSize,omitempty"`
	MinInstanceSize *InstanceSize `json:"minInstanceSize,omitempty"`
	// Flag that indicates whether the instance size may scale down. MongoDB Cloud requires this parameter if `\"replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled\" : true`. If you enable this option, specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize**.
	ScaleDownEnabled *bool `json:"scaleDownEnabled,omitempty"`
}

// NewComputeAutoScalingV15 instantiates a new ComputeAutoScalingV15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAutoScalingV15() *ComputeAutoScalingV15 {
	this := ComputeAutoScalingV15{}
	return &this
}

// NewComputeAutoScalingV15WithDefaults instantiates a new ComputeAutoScalingV15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAutoScalingV15WithDefaults() *ComputeAutoScalingV15 {
	this := ComputeAutoScalingV15{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComputeAutoScalingV15) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScalingV15) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComputeAutoScalingV15) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComputeAutoScalingV15) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise.
func (o *ComputeAutoScalingV15) GetMaxInstanceSize() InstanceSize {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret InstanceSize
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScalingV15) GetMaxInstanceSizeOk() (*InstanceSize, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}
	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *ComputeAutoScalingV15) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given InstanceSize and assigns it to the MaxInstanceSize field.
func (o *ComputeAutoScalingV15) SetMaxInstanceSize(v InstanceSize) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise.
func (o *ComputeAutoScalingV15) GetMinInstanceSize() InstanceSize {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret InstanceSize
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScalingV15) GetMinInstanceSizeOk() (*InstanceSize, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}
	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *ComputeAutoScalingV15) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given InstanceSize and assigns it to the MinInstanceSize field.
func (o *ComputeAutoScalingV15) SetMinInstanceSize(v InstanceSize) {
	o.MinInstanceSize = &v
}

// GetScaleDownEnabled returns the ScaleDownEnabled field value if set, zero value otherwise.
func (o *ComputeAutoScalingV15) GetScaleDownEnabled() bool {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		var ret bool
		return ret
	}
	return *o.ScaleDownEnabled
}

// GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAutoScalingV15) GetScaleDownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		return nil, false
	}
	return o.ScaleDownEnabled, true
}

// HasScaleDownEnabled returns a boolean if a field has been set.
func (o *ComputeAutoScalingV15) HasScaleDownEnabled() bool {
	if o != nil && !IsNil(o.ScaleDownEnabled) {
		return true
	}

	return false
}

// SetScaleDownEnabled gets a reference to the given bool and assigns it to the ScaleDownEnabled field.
func (o *ComputeAutoScalingV15) SetScaleDownEnabled(v bool) {
	o.ScaleDownEnabled = &v
}

func (o ComputeAutoScalingV15) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ComputeAutoScalingV15) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !IsNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	if !IsNil(o.ScaleDownEnabled) {
		toSerialize["scaleDownEnabled"] = o.ScaleDownEnabled
	}
	return toSerialize, nil
}

type NullableComputeAutoScalingV15 struct {
	value *ComputeAutoScalingV15
	isSet bool
}

func (v NullableComputeAutoScalingV15) Get() *ComputeAutoScalingV15 {
	return v.value
}

func (v *NullableComputeAutoScalingV15) Set(val *ComputeAutoScalingV15) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeAutoScalingV15) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeAutoScalingV15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeAutoScalingV15(val *ComputeAutoScalingV15) *NullableComputeAutoScalingV15 {
	return &NullableComputeAutoScalingV15{value: val, isSet: true}
}

func (v NullableComputeAutoScalingV15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeAutoScalingV15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


