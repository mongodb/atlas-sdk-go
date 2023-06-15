// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AzureComputeAutoScalingRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureComputeAutoScalingRules{}

// AzureComputeAutoScalingRules Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type AzureComputeAutoScalingRules struct {
	// Maximum instance size to which your cluster can automatically scale.
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size to which your cluster can automatically scale.
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
}

// NewAzureComputeAutoScalingRules instantiates a new AzureComputeAutoScalingRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeAutoScalingRules() *AzureComputeAutoScalingRules {
	this := AzureComputeAutoScalingRules{}
	return &this
}

// NewAzureComputeAutoScalingRulesWithDefaults instantiates a new AzureComputeAutoScalingRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeAutoScalingRulesWithDefaults() *AzureComputeAutoScalingRules {
	this := AzureComputeAutoScalingRules{}
	return &this
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise.
func (o *AzureComputeAutoScalingRules) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeAutoScalingRules) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}
	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *AzureComputeAutoScalingRules) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *AzureComputeAutoScalingRules) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise.
func (o *AzureComputeAutoScalingRules) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeAutoScalingRules) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}
	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *AzureComputeAutoScalingRules) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *AzureComputeAutoScalingRules) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
}

func (o AzureComputeAutoScalingRules) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureComputeAutoScalingRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !IsNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	return toSerialize, nil
}

type NullableAzureComputeAutoScalingRules struct {
	value *AzureComputeAutoScalingRules
	isSet bool
}

func (v NullableAzureComputeAutoScalingRules) Get() *AzureComputeAutoScalingRules {
	return v.value
}

func (v *NullableAzureComputeAutoScalingRules) Set(val *AzureComputeAutoScalingRules) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeAutoScalingRules) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeAutoScalingRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeAutoScalingRules(val *AzureComputeAutoScalingRules) *NullableAzureComputeAutoScalingRules {
	return &NullableAzureComputeAutoScalingRules{value: val, isSet: true}
}

func (v NullableAzureComputeAutoScalingRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeAutoScalingRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
