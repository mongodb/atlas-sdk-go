/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the FreeAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeAutoScaling{}

// FreeAutoScaling Range of instance sizes to which your cluster can scale.
type FreeAutoScaling struct {
	// Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down.
	Compute *string `json:"compute,omitempty"`
}

// NewFreeAutoScaling instantiates a new FreeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeAutoScaling() *FreeAutoScaling {
	this := FreeAutoScaling{}
	return &this
}

// NewFreeAutoScalingWithDefaults instantiates a new FreeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeAutoScalingWithDefaults() *FreeAutoScaling {
	this := FreeAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *FreeAutoScaling) GetCompute() string {
	if o == nil || IsNil(o.Compute) {
		var ret string
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeAutoScaling) GetComputeOk() (*string, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *FreeAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given string and assigns it to the Compute field.
func (o *FreeAutoScaling) SetCompute(v string) {
	o.Compute = &v
}

func (o FreeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FreeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableFreeAutoScaling struct {
	value *FreeAutoScaling
	isSet bool
}

func (v NullableFreeAutoScaling) Get() *FreeAutoScaling {
	return v.value
}

func (v *NullableFreeAutoScaling) Set(val *FreeAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeAutoScaling(val *FreeAutoScaling) *NullableFreeAutoScaling {
	return &NullableFreeAutoScaling{value: val, isSet: true}
}

func (v NullableFreeAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


