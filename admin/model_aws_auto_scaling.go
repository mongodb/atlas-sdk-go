// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the AWSAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSAutoScaling{}

// AWSAutoScaling struct for AWSAutoScaling
type AWSAutoScaling struct {
	Compute *AWSComputeAutoScaling `json:"compute,omitempty"`
}

// NewAWSAutoScaling instantiates a new AWSAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSAutoScaling() *AWSAutoScaling {
	this := AWSAutoScaling{}
	return &this
}

// NewAWSAutoScalingWithDefaults instantiates a new AWSAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSAutoScalingWithDefaults() *AWSAutoScaling {
	this := AWSAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *AWSAutoScaling) GetCompute() AWSComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret AWSComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAutoScaling) GetComputeOk() (*AWSComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *AWSAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given AWSComputeAutoScaling and assigns it to the Compute field.
func (o *AWSAutoScaling) SetCompute(v AWSComputeAutoScaling) {
	o.Compute = &v
}

func (o AWSAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableAWSAutoScaling struct {
	value *AWSAutoScaling
	isSet bool
}

func (v NullableAWSAutoScaling) Get() *AWSAutoScaling {
	return v.value
}

func (v *NullableAWSAutoScaling) Set(val *AWSAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSAutoScaling(val *AWSAutoScaling) *NullableAWSAutoScaling {
	return &NullableAWSAutoScaling{value: val, isSet: true}
}

func (v NullableAWSAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
