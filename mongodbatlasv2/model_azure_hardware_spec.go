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

// checks if the AzureHardwareSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureHardwareSpec{}

// AzureHardwareSpec struct for AzureHardwareSpec
type AzureHardwareSpec struct {
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of nodes of the given type for MongoDB Cloud to deploy to the region.
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// NewAzureHardwareSpec instantiates a new AzureHardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureHardwareSpec() *AzureHardwareSpec {
	this := AzureHardwareSpec{}
	return &this
}

// NewAzureHardwareSpecWithDefaults instantiates a new AzureHardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureHardwareSpecWithDefaults() *AzureHardwareSpec {
	this := AzureHardwareSpec{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *AzureHardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureHardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *AzureHardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *AzureHardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *AzureHardwareSpec) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureHardwareSpec) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *AzureHardwareSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *AzureHardwareSpec) SetNodeCount(v int32) {
	o.NodeCount = &v
}

func (o AzureHardwareSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureHardwareSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	if !IsNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	return toSerialize, nil
}

type NullableAzureHardwareSpec struct {
	value *AzureHardwareSpec
	isSet bool
}

func (v NullableAzureHardwareSpec) Get() *AzureHardwareSpec {
	return v.value
}

func (v *NullableAzureHardwareSpec) Set(val *AzureHardwareSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureHardwareSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureHardwareSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureHardwareSpec(val *AzureHardwareSpec) *NullableAzureHardwareSpec {
	return &NullableAzureHardwareSpec{value: val, isSet: true}
}

func (v NullableAzureHardwareSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureHardwareSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


