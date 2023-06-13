// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the GCPAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPAutoScaling{}

// GCPAutoScaling struct for GCPAutoScaling
type GCPAutoScaling struct {
	Compute *GCPComputeAutoScaling `json:"compute,omitempty"`
}

// NewGCPAutoScaling instantiates a new GCPAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPAutoScaling() *GCPAutoScaling {
	this := GCPAutoScaling{}
	return &this
}

// NewGCPAutoScalingWithDefaults instantiates a new GCPAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPAutoScalingWithDefaults() *GCPAutoScaling {
	this := GCPAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *GCPAutoScaling) GetCompute() GCPComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret GCPComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAutoScaling) GetComputeOk() (*GCPComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *GCPAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given GCPComputeAutoScaling and assigns it to the Compute field.
func (o *GCPAutoScaling) SetCompute(v GCPComputeAutoScaling) {
	o.Compute = &v
}

func (o GCPAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GCPAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableGCPAutoScaling struct {
	value *GCPAutoScaling
	isSet bool
}

func (v NullableGCPAutoScaling) Get() *GCPAutoScaling {
	return v.value
}

func (v *NullableGCPAutoScaling) Set(val *GCPAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPAutoScaling(val *GCPAutoScaling) *NullableGCPAutoScaling {
	return &NullableGCPAutoScaling{value: val, isSet: true}
}

func (v NullableGCPAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
