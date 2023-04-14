/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the AzureAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureAutoScaling{}

// AzureAutoScaling Range of instance sizes to which your cluster can scale.
type AzureAutoScaling struct {
	Compute *AzureComputeAutoScaling `json:"compute,omitempty"`
}

// NewAzureAutoScaling instantiates a new AzureAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureAutoScaling() *AzureAutoScaling {
	this := AzureAutoScaling{}
	return &this
}

// NewAzureAutoScalingWithDefaults instantiates a new AzureAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureAutoScalingWithDefaults() *AzureAutoScaling {
	this := AzureAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *AzureAutoScaling) GetCompute() AzureComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret AzureComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAutoScaling) GetComputeOk() (*AzureComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *AzureAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given AzureComputeAutoScaling and assigns it to the Compute field.
func (o *AzureAutoScaling) SetCompute(v AzureComputeAutoScaling) {
	o.Compute = &v
}

func (o AzureAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}

type NullableAzureAutoScaling struct {
	value *AzureAutoScaling
	isSet bool
}

func (v NullableAzureAutoScaling) Get() *AzureAutoScaling {
	return v.value
}

func (v *NullableAzureAutoScaling) Set(val *AzureAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureAutoScaling(val *AzureAutoScaling) *NullableAzureAutoScaling {
	return &NullableAzureAutoScaling{value: val, isSet: true}
}

func (v NullableAzureAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


