// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAzureAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAzureAutoScaling{}

// CloudProviderAzureAutoScaling Range of instance sizes to which your cluster can scale.
type CloudProviderAzureAutoScaling struct {
	Compute *AzureComputeAutoScalingRules `json:"compute,omitempty"`
}

// NewCloudProviderAzureAutoScaling instantiates a new CloudProviderAzureAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAzureAutoScaling() *CloudProviderAzureAutoScaling {
	this := CloudProviderAzureAutoScaling{}
	return &this
}

// NewCloudProviderAzureAutoScalingWithDefaults instantiates a new CloudProviderAzureAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAzureAutoScalingWithDefaults() *CloudProviderAzureAutoScaling {
	this := CloudProviderAzureAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *CloudProviderAzureAutoScaling) GetCompute() AzureComputeAutoScalingRules {
	if o == nil || IsNil(o.Compute) {
		var ret AzureComputeAutoScalingRules
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAzureAutoScaling) GetComputeOk() (*AzureComputeAutoScalingRules, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *CloudProviderAzureAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given AzureComputeAutoScalingRules and assigns it to the Compute field.
func (o *CloudProviderAzureAutoScaling) SetCompute(v AzureComputeAutoScalingRules) {
	o.Compute = &v
}

func (o CloudProviderAzureAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAzureAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableCloudProviderAzureAutoScaling struct {
	value *CloudProviderAzureAutoScaling
	isSet bool
}

func (v NullableCloudProviderAzureAutoScaling) Get() *CloudProviderAzureAutoScaling {
	return v.value
}

func (v *NullableCloudProviderAzureAutoScaling) Set(val *CloudProviderAzureAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAzureAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAzureAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAzureAutoScaling(val *CloudProviderAzureAutoScaling) *NullableCloudProviderAzureAutoScaling {
	return &NullableCloudProviderAzureAutoScaling{value: val, isSet: true}
}

func (v NullableCloudProviderAzureAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAzureAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
