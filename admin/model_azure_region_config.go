// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the AzureRegionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureRegionConfig{}

// AzureRegionConfig Details that explain how MongoDB Cloud replicates data in one region on the specified MongoDB database.
type AzureRegionConfig struct {
	AnalyticsAutoScaling *AutoScalingV15        `json:"analyticsAutoScaling,omitempty"`
	AnalyticsSpecs       *DedicatedHardwareSpec `json:"analyticsSpecs,omitempty"`
	AutoScaling          *AutoScalingV15        `json:"autoScaling,omitempty"`
	ReadOnlySpecs        *DedicatedHardwareSpec `json:"readOnlySpecs,omitempty"`
	ElectableSpecs       *HardwareSpec          `json:"electableSpecs,omitempty"`
	// Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to `0`. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is `7`.  **Example:** If you have three regions, their priorities would be `7`, `6`, and `5` respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be `4` and `3` respectively.
	Priority *int32 `json:"priority,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to `AWS`, `GCP`, `AZURE` or `TENANT`.
	ProviderName *string `json:"providerName,omitempty"`
	// Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC.
	RegionName *string `json:"regionName,omitempty"`
}

// NewAzureRegionConfig instantiates a new AzureRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureRegionConfig() *AzureRegionConfig {
	this := AzureRegionConfig{}
	return &this
}

// NewAzureRegionConfigWithDefaults instantiates a new AzureRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureRegionConfigWithDefaults() *AzureRegionConfig {
	this := AzureRegionConfig{}
	return &this
}

// GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetAnalyticsAutoScaling() AutoScalingV15 {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		var ret AutoScalingV15
		return ret
	}
	return *o.AnalyticsAutoScaling
}

// GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetAnalyticsAutoScalingOk() (*AutoScalingV15, bool) {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		return nil, false
	}
	return o.AnalyticsAutoScaling, true
}

// HasAnalyticsAutoScaling returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasAnalyticsAutoScaling() bool {
	if o != nil && !IsNil(o.AnalyticsAutoScaling) {
		return true
	}

	return false
}

// SetAnalyticsAutoScaling gets a reference to the given AutoScalingV15 and assigns it to the AnalyticsAutoScaling field.
func (o *AzureRegionConfig) SetAnalyticsAutoScaling(v AutoScalingV15) {
	o.AnalyticsAutoScaling = &v
}

// GetAnalyticsSpecs returns the AnalyticsSpecs field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetAnalyticsSpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.AnalyticsSpecs
}

// GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetAnalyticsSpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		return nil, false
	}
	return o.AnalyticsSpecs, true
}

// HasAnalyticsSpecs returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasAnalyticsSpecs() bool {
	if o != nil && !IsNil(o.AnalyticsSpecs) {
		return true
	}

	return false
}

// SetAnalyticsSpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the AnalyticsSpecs field.
func (o *AzureRegionConfig) SetAnalyticsSpecs(v DedicatedHardwareSpec) {
	o.AnalyticsSpecs = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetAutoScaling() AutoScalingV15 {
	if o == nil || IsNil(o.AutoScaling) {
		var ret AutoScalingV15
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetAutoScalingOk() (*AutoScalingV15, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given AutoScalingV15 and assigns it to the AutoScaling field.
func (o *AzureRegionConfig) SetAutoScaling(v AutoScalingV15) {
	o.AutoScaling = &v
}

// GetReadOnlySpecs returns the ReadOnlySpecs field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetReadOnlySpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.ReadOnlySpecs
}

// GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetReadOnlySpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		return nil, false
	}
	return o.ReadOnlySpecs, true
}

// HasReadOnlySpecs returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasReadOnlySpecs() bool {
	if o != nil && !IsNil(o.ReadOnlySpecs) {
		return true
	}

	return false
}

// SetReadOnlySpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the ReadOnlySpecs field.
func (o *AzureRegionConfig) SetReadOnlySpecs(v DedicatedHardwareSpec) {
	o.ReadOnlySpecs = &v
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetElectableSpecs() HardwareSpec {
	if o == nil || IsNil(o.ElectableSpecs) {
		var ret HardwareSpec
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool) {
	if o == nil || IsNil(o.ElectableSpecs) {
		return nil, false
	}
	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasElectableSpecs() bool {
	if o != nil && !IsNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given HardwareSpec and assigns it to the ElectableSpecs field.
func (o *AzureRegionConfig) SetElectableSpecs(v HardwareSpec) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AzureRegionConfig) SetPriority(v int32) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AzureRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AzureRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AzureRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AzureRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}

func (o AzureRegionConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureRegionConfig) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableAzureRegionConfig struct {
	value *AzureRegionConfig
	isSet bool
}

func (v NullableAzureRegionConfig) Get() *AzureRegionConfig {
	return v.value
}

func (v *NullableAzureRegionConfig) Set(val *AzureRegionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureRegionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureRegionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureRegionConfig(val *AzureRegionConfig) *NullableAzureRegionConfig {
	return &NullableAzureRegionConfig{value: val, isSet: true}
}

func (v NullableAzureRegionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureRegionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
