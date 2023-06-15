// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the GCPComputeAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPComputeAutoScaling{}

// GCPComputeAutoScaling Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type GCPComputeAutoScaling struct {
	// Maximum instance size to which your cluster can automatically scale.
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size to which your cluster can automatically scale.
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
}

// NewGCPComputeAutoScaling instantiates a new GCPComputeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPComputeAutoScaling() *GCPComputeAutoScaling {
	this := GCPComputeAutoScaling{}
	return &this
}

// NewGCPComputeAutoScalingWithDefaults instantiates a new GCPComputeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPComputeAutoScalingWithDefaults() *GCPComputeAutoScaling {
	this := GCPComputeAutoScaling{}
	return &this
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise.
func (o *GCPComputeAutoScaling) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}
	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *GCPComputeAutoScaling) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *GCPComputeAutoScaling) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise.
func (o *GCPComputeAutoScaling) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}
	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *GCPComputeAutoScaling) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *GCPComputeAutoScaling) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
}

func (o GCPComputeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GCPComputeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !IsNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	return toSerialize, nil
}

type NullableGCPComputeAutoScaling struct {
	value *GCPComputeAutoScaling
	isSet bool
}

func (v NullableGCPComputeAutoScaling) Get() *GCPComputeAutoScaling {
	return v.value
}

func (v *NullableGCPComputeAutoScaling) Set(val *GCPComputeAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPComputeAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPComputeAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPComputeAutoScaling(val *GCPComputeAutoScaling) *NullableGCPComputeAutoScaling {
	return &NullableGCPComputeAutoScaling{value: val, isSet: true}
}

func (v NullableGCPComputeAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPComputeAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
