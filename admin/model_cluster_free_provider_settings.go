// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the ClusterFreeProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFreeProviderSettings{}

// ClusterFreeProviderSettings struct for ClusterFreeProviderSettings
type ClusterFreeProviderSettings struct {
	AutoScaling *ClusterFreeAutoScaling `json:"autoScaling,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant host. The resource returns this parameter when **providerSettings.providerName** is `TENANT` and **providerSetting.instanceSizeName** is `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. You must set **providerSettings.providerName** to `TENANT` and specify the cloud service provider in **providerSettings.backingProviderName**.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). For multi-region clusters, see **replicationSpec.{region}**.
	RegionName   *string `json:"regionName,omitempty"`
	ProviderName string  `json:"providerName"`
}

// NewClusterFreeProviderSettings instantiates a new ClusterFreeProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFreeProviderSettings(providerName string) *ClusterFreeProviderSettings {
	this := ClusterFreeProviderSettings{}
	this.ProviderName = providerName
	return &this
}

// NewClusterFreeProviderSettingsWithDefaults instantiates a new ClusterFreeProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFreeProviderSettingsWithDefaults() *ClusterFreeProviderSettings {
	this := ClusterFreeProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *ClusterFreeProviderSettings) GetAutoScaling() ClusterFreeAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret ClusterFreeAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFreeProviderSettings) GetAutoScalingOk() (*ClusterFreeAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ClusterFreeProviderSettings) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ClusterFreeAutoScaling and assigns it to the AutoScaling field.
func (o *ClusterFreeProviderSettings) SetAutoScaling(v ClusterFreeAutoScaling) {
	o.AutoScaling = &v
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise.
func (o *ClusterFreeProviderSettings) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFreeProviderSettings) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}
	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *ClusterFreeProviderSettings) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *ClusterFreeProviderSettings) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *ClusterFreeProviderSettings) GetInstanceSizeName() string {
	if o == nil || IsNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFreeProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSizeName) {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *ClusterFreeProviderSettings) HasInstanceSizeName() bool {
	if o != nil && !IsNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *ClusterFreeProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ClusterFreeProviderSettings) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFreeProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ClusterFreeProviderSettings) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ClusterFreeProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetProviderName returns the ProviderName field value
func (o *ClusterFreeProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *ClusterFreeProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *ClusterFreeProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o ClusterFreeProviderSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterFreeProviderSettings) ToMap() (map[string]interface{}, error) {
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

type NullableClusterFreeProviderSettings struct {
	value *ClusterFreeProviderSettings
	isSet bool
}

func (v NullableClusterFreeProviderSettings) Get() *ClusterFreeProviderSettings {
	return v.value
}

func (v *NullableClusterFreeProviderSettings) Set(val *ClusterFreeProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFreeProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFreeProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFreeProviderSettings(val *ClusterFreeProviderSettings) *NullableClusterFreeProviderSettings {
	return &NullableClusterFreeProviderSettings{value: val, isSet: true}
}

func (v NullableClusterFreeProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFreeProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
