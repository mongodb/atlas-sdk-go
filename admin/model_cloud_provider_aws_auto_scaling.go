// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAWSAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAWSAutoScaling{}

// CloudProviderAWSAutoScaling Range of instance sizes to which your cluster can scale.
type CloudProviderAWSAutoScaling struct {
	Compute *AWSComputeAutoScaling `json:"compute,omitempty"`
}

// NewCloudProviderAWSAutoScaling instantiates a new CloudProviderAWSAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAWSAutoScaling() *CloudProviderAWSAutoScaling {
	this := CloudProviderAWSAutoScaling{}
	return &this
}

// NewCloudProviderAWSAutoScalingWithDefaults instantiates a new CloudProviderAWSAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAWSAutoScalingWithDefaults() *CloudProviderAWSAutoScaling {
	this := CloudProviderAWSAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *CloudProviderAWSAutoScaling) GetCompute() AWSComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret AWSComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAWSAutoScaling) GetComputeOk() (*AWSComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *CloudProviderAWSAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given AWSComputeAutoScaling and assigns it to the Compute field.
func (o *CloudProviderAWSAutoScaling) SetCompute(v AWSComputeAutoScaling) {
	o.Compute = &v
}

func (o CloudProviderAWSAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAWSAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableCloudProviderAWSAutoScaling struct {
	value *CloudProviderAWSAutoScaling
	isSet bool
}

func (v NullableCloudProviderAWSAutoScaling) Get() *CloudProviderAWSAutoScaling {
	return v.value
}

func (v *NullableCloudProviderAWSAutoScaling) Set(val *CloudProviderAWSAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAWSAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAWSAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAWSAutoScaling(val *CloudProviderAWSAutoScaling) *NullableCloudProviderAWSAutoScaling {
	return &NullableCloudProviderAWSAutoScaling{value: val, isSet: true}
}

func (v NullableCloudProviderAWSAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAWSAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
