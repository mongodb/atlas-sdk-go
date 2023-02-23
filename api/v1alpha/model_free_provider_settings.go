/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the FreeProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeProviderSettings{}

// FreeProviderSettings struct for FreeProviderSettings
type FreeProviderSettings struct {
	AutoScaling *FreeAutoScaling `json:"autoScaling,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant host. The resource returns this parameter when **providerSettings.providerName** is `TENANT` and **providerSetting.instanceSizeName** is `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. You must set **providerSettings.providerName** to `TENANT` and specify the cloud service provider in **providerSettings.backingProviderName**.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). For multi-region clusters, see **replicationSpec.{region}**.
	RegionName *string `json:"regionName,omitempty"`
	ProviderName string `json:"providerName"`
}

// NewFreeProviderSettings instantiates a new FreeProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeProviderSettings() *FreeProviderSettings {
	this := FreeProviderSettings{}
	return &this
}

// NewFreeProviderSettingsWithDefaults instantiates a new FreeProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeProviderSettingsWithDefaults() *FreeProviderSettings {
	this := FreeProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *FreeProviderSettings) GetAutoScaling() FreeAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret FreeAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeProviderSettings) GetAutoScalingOk() (*FreeAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *FreeProviderSettings) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given FreeAutoScaling and assigns it to the AutoScaling field.
func (o *FreeProviderSettings) SetAutoScaling(v FreeAutoScaling) {
	o.AutoScaling = &v
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise.
func (o *FreeProviderSettings) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeProviderSettings) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}
	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *FreeProviderSettings) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *FreeProviderSettings) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *FreeProviderSettings) GetInstanceSizeName() string {
	if o == nil || IsNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSizeName) {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *FreeProviderSettings) HasInstanceSizeName() bool {
	if o != nil && !IsNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *FreeProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *FreeProviderSettings) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *FreeProviderSettings) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *FreeProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetProviderName returns the ProviderName field value
func (o *FreeProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *FreeProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *FreeProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o FreeProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.BackingProviderName) {
		toSerialize["backingProviderName"] = o.BackingProviderName
	}
	if !IsNil(o.InstanceSizeName) {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableFreeProviderSettings struct {
	value *FreeProviderSettings
	isSet bool
}

func (v NullableFreeProviderSettings) Get() *FreeProviderSettings {
	return v.value
}

func (v *NullableFreeProviderSettings) Set(val *FreeProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeProviderSettings(val *FreeProviderSettings) *NullableFreeProviderSettings {
	return &NullableFreeProviderSettings{value: val, isSet: true}
}

func (v NullableFreeProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


