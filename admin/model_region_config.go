// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the RegionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionConfig{}

// RegionConfig struct for RegionConfig
type RegionConfig struct {
	AnalyticsAutoScaling *AutoScalingV15        `json:"analyticsAutoScaling,omitempty"`
	AnalyticsSpecs       *DedicatedHardwareSpec `json:"analyticsSpecs,omitempty"`
	AutoScaling          *AutoScalingV15        `json:"autoScaling,omitempty"`
	ReadOnlySpecs        *DedicatedHardwareSpec `json:"readOnlySpecs,omitempty"`
	ElectableSpecs       *HardwareSpec          `json:"electableSpecs,omitempty"`
	// Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to `0`. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is `7`.  **Example:** If you have three regions, their priorities would be `7`, `6`, and `5` respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be `4` and `3` respectively.
	Priority *int `json:"priority,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to `AWS`, `GCP`, `AZURE` or `TENANT`.
	ProviderName *string `json:"providerName,omitempty"`
	// Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC.
	RegionName *string `json:"regionName,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant cluster. The resource returns this parameter when **providerSettings.providerName** is `TENANT` and **providerSetting.instanceSizeName** is `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
}

// NewRegionConfig instantiates a new RegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionConfig() *RegionConfig {
	this := RegionConfig{}
	return &this
}

// NewRegionConfigWithDefaults instantiates a new RegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionConfigWithDefaults() *RegionConfig {
	this := RegionConfig{}
	return &this
}

// GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field value if set, zero value otherwise.
func (o *RegionConfig) GetAnalyticsAutoScaling() AutoScalingV15 {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		var ret AutoScalingV15
		return ret
	}
	return *o.AnalyticsAutoScaling
}

// GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetAnalyticsAutoScalingOk() (*AutoScalingV15, bool) {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		return nil, false
	}
	return o.AnalyticsAutoScaling, true
}

// HasAnalyticsAutoScaling returns a boolean if a field has been set.
func (o *RegionConfig) HasAnalyticsAutoScaling() bool {
	if o != nil && !IsNil(o.AnalyticsAutoScaling) {
		return true
	}

	return false
}

// SetAnalyticsAutoScaling gets a reference to the given AutoScalingV15 and assigns it to the AnalyticsAutoScaling field.
func (o *RegionConfig) SetAnalyticsAutoScaling(v AutoScalingV15) {
	o.AnalyticsAutoScaling = &v
}

// GetAnalyticsSpecs returns the AnalyticsSpecs field value if set, zero value otherwise.
func (o *RegionConfig) GetAnalyticsSpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.AnalyticsSpecs
}

// GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetAnalyticsSpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		return nil, false
	}
	return o.AnalyticsSpecs, true
}

// HasAnalyticsSpecs returns a boolean if a field has been set.
func (o *RegionConfig) HasAnalyticsSpecs() bool {
	if o != nil && !IsNil(o.AnalyticsSpecs) {
		return true
	}

	return false
}

// SetAnalyticsSpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the AnalyticsSpecs field.
func (o *RegionConfig) SetAnalyticsSpecs(v DedicatedHardwareSpec) {
	o.AnalyticsSpecs = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *RegionConfig) GetAutoScaling() AutoScalingV15 {
	if o == nil || IsNil(o.AutoScaling) {
		var ret AutoScalingV15
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetAutoScalingOk() (*AutoScalingV15, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *RegionConfig) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given AutoScalingV15 and assigns it to the AutoScaling field.
func (o *RegionConfig) SetAutoScaling(v AutoScalingV15) {
	o.AutoScaling = &v
}

// GetReadOnlySpecs returns the ReadOnlySpecs field value if set, zero value otherwise.
func (o *RegionConfig) GetReadOnlySpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.ReadOnlySpecs
}

// GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetReadOnlySpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		return nil, false
	}
	return o.ReadOnlySpecs, true
}

// HasReadOnlySpecs returns a boolean if a field has been set.
func (o *RegionConfig) HasReadOnlySpecs() bool {
	if o != nil && !IsNil(o.ReadOnlySpecs) {
		return true
	}

	return false
}

// SetReadOnlySpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the ReadOnlySpecs field.
func (o *RegionConfig) SetReadOnlySpecs(v DedicatedHardwareSpec) {
	o.ReadOnlySpecs = &v
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise.
func (o *RegionConfig) GetElectableSpecs() HardwareSpec {
	if o == nil || IsNil(o.ElectableSpecs) {
		var ret HardwareSpec
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool) {
	if o == nil || IsNil(o.ElectableSpecs) {
		return nil, false
	}
	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *RegionConfig) HasElectableSpecs() bool {
	if o != nil && !IsNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given HardwareSpec and assigns it to the ElectableSpecs field.
func (o *RegionConfig) SetElectableSpecs(v HardwareSpec) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RegionConfig) GetPriority() int {
	if o == nil || IsNil(o.Priority) {
		var ret int
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetPriorityOk() (*int, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RegionConfig) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int and assigns it to the Priority field.
func (o *RegionConfig) SetPriority(v int) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *RegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *RegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *RegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *RegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *RegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *RegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise.
func (o *RegionConfig) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionConfig) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}
	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *RegionConfig) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *RegionConfig) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

func (o RegionConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o RegionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnalyticsAutoScaling) {
		toSerialize["analyticsAutoScaling"] = o.AnalyticsAutoScaling
	}
	if !IsNil(o.AnalyticsSpecs) {
		toSerialize["analyticsSpecs"] = o.AnalyticsSpecs
	}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.ReadOnlySpecs) {
		toSerialize["readOnlySpecs"] = o.ReadOnlySpecs
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
	if !IsNil(o.BackingProviderName) {
		toSerialize["backingProviderName"] = o.BackingProviderName
	}
	return toSerialize, nil
}

type NullableRegionConfig struct {
	value *RegionConfig
	isSet bool
}

func (v NullableRegionConfig) Get() *RegionConfig {
	return v.value
}

func (v *NullableRegionConfig) Set(val *RegionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionConfig(val *RegionConfig) *NullableRegionConfig {
	return &NullableRegionConfig{value: val, isSet: true}
}

func (v NullableRegionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
