/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TenantRegionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantRegionConfig{}

// TenantRegionConfig Details that explain how MongoDB Cloud replicates data in one region on the specified MongoDB database.
type TenantRegionConfig struct {
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant cluster. The resource returns this parameter when **providerSettings.providerName** is `TENANT` and **providerSetting.instanceSizeName** is `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
	ElectableSpecs *HardwareSpec `json:"electableSpecs,omitempty"`
	// Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to `0`. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is `7`.  **Example:** If you have three regions, their priorities would be `7`, `6`, and `5` respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be `4` and `3` respectively.
	Priority *int32 `json:"priority,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to `AWS`, `GCP`, `AZURE` or `TENANT`.
	ProviderName *string `json:"providerName,omitempty"`
	// Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC.
	RegionName *string `json:"regionName,omitempty"`
}

// NewTenantRegionConfig instantiates a new TenantRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRegionConfig() *TenantRegionConfig {
	this := TenantRegionConfig{}
	return &this
}

// NewTenantRegionConfigWithDefaults instantiates a new TenantRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRegionConfigWithDefaults() *TenantRegionConfig {
	this := TenantRegionConfig{}
	return &this
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise.
func (o *TenantRegionConfig) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRegionConfig) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}
	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *TenantRegionConfig) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *TenantRegionConfig) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise.
func (o *TenantRegionConfig) GetElectableSpecs() HardwareSpec {
	if o == nil || IsNil(o.ElectableSpecs) {
		var ret HardwareSpec
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool) {
	if o == nil || IsNil(o.ElectableSpecs) {
		return nil, false
	}
	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *TenantRegionConfig) HasElectableSpecs() bool {
	if o != nil && !IsNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given HardwareSpec and assigns it to the ElectableSpecs field.
func (o *TenantRegionConfig) SetElectableSpecs(v HardwareSpec) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TenantRegionConfig) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRegionConfig) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TenantRegionConfig) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TenantRegionConfig) SetPriority(v int32) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *TenantRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *TenantRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *TenantRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *TenantRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *TenantRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *TenantRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}

func (o TenantRegionConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TenantRegionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackingProviderName) {
		toSerialize["backingProviderName"] = o.BackingProviderName
	}
	if !IsNil(o.ElectableSpecs) {
		toSerialize["electableSpecs"] = o.ElectableSpecs
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	return toSerialize, nil
}

type NullableTenantRegionConfig struct {
	value *TenantRegionConfig
	isSet bool
}

func (v NullableTenantRegionConfig) Get() *TenantRegionConfig {
	return v.value
}

func (v *NullableTenantRegionConfig) Set(val *TenantRegionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRegionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRegionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRegionConfig(val *TenantRegionConfig) *NullableTenantRegionConfig {
	return &NullableTenantRegionConfig{value: val, isSet: true}
}

func (v NullableTenantRegionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRegionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


