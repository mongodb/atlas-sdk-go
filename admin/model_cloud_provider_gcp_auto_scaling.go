// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderGCPAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderGCPAutoScaling{}

// CloudProviderGCPAutoScaling Range of instance sizes to which your cluster can scale.
type CloudProviderGCPAutoScaling struct {
	Compute *GCPComputeAutoScaling `json:"compute,omitempty"`
}

// NewCloudProviderGCPAutoScaling instantiates a new CloudProviderGCPAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderGCPAutoScaling() *CloudProviderGCPAutoScaling {
	this := CloudProviderGCPAutoScaling{}
	return &this
}

// NewCloudProviderGCPAutoScalingWithDefaults instantiates a new CloudProviderGCPAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderGCPAutoScalingWithDefaults() *CloudProviderGCPAutoScaling {
	this := CloudProviderGCPAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *CloudProviderGCPAutoScaling) GetCompute() GCPComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret GCPComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderGCPAutoScaling) GetComputeOk() (*GCPComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *CloudProviderGCPAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given GCPComputeAutoScaling and assigns it to the Compute field.
func (o *CloudProviderGCPAutoScaling) SetCompute(v GCPComputeAutoScaling) {
	o.Compute = &v
}

func (o CloudProviderGCPAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderGCPAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableCloudProviderGCPAutoScaling struct {
	value *CloudProviderGCPAutoScaling
	isSet bool
}

func (v NullableCloudProviderGCPAutoScaling) Get() *CloudProviderGCPAutoScaling {
	return v.value
}

func (v *NullableCloudProviderGCPAutoScaling) Set(val *CloudProviderGCPAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderGCPAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderGCPAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderGCPAutoScaling(val *CloudProviderGCPAutoScaling) *NullableCloudProviderGCPAutoScaling {
	return &NullableCloudProviderGCPAutoScaling{value: val, isSet: true}
}

func (v NullableCloudProviderGCPAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderGCPAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
