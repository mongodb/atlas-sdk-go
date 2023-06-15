// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the ClusterOutageSimulationOutageFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterOutageSimulationOutageFilter{}

// ClusterOutageSimulationOutageFilter struct for ClusterOutageSimulationOutageFilter
type ClusterOutageSimulationOutageFilter struct {
	// The cloud provider of the region that undergoes the outage simulation.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// The name of the region to undergo an outage simulation.
	RegionName *string `json:"regionName,omitempty"`
	// The type of cluster outage to simulate.  | Type       | Description | |------------|-------------| | `REGION`   | Simulates a cluster outage for a region.|
	Type *string `json:"type,omitempty"`
}

// NewClusterOutageSimulationOutageFilter instantiates a new ClusterOutageSimulationOutageFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOutageSimulationOutageFilter() *ClusterOutageSimulationOutageFilter {
	this := ClusterOutageSimulationOutageFilter{}
	return &this
}

// NewClusterOutageSimulationOutageFilterWithDefaults instantiates a new ClusterOutageSimulationOutageFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOutageSimulationOutageFilterWithDefaults() *ClusterOutageSimulationOutageFilter {
	this := ClusterOutageSimulationOutageFilter{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ClusterOutageSimulationOutageFilter) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOutageSimulationOutageFilter) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ClusterOutageSimulationOutageFilter) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ClusterOutageSimulationOutageFilter) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ClusterOutageSimulationOutageFilter) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOutageSimulationOutageFilter) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ClusterOutageSimulationOutageFilter) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ClusterOutageSimulationOutageFilter) SetRegionName(v string) {
	o.RegionName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterOutageSimulationOutageFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOutageSimulationOutageFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterOutageSimulationOutageFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterOutageSimulationOutageFilter) SetType(v string) {
	o.Type = &v
}

func (o ClusterOutageSimulationOutageFilter) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterOutageSimulationOutageFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableClusterOutageSimulationOutageFilter struct {
	value *ClusterOutageSimulationOutageFilter
	isSet bool
}

func (v NullableClusterOutageSimulationOutageFilter) Get() *ClusterOutageSimulationOutageFilter {
	return v.value
}

func (v *NullableClusterOutageSimulationOutageFilter) Set(val *ClusterOutageSimulationOutageFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOutageSimulationOutageFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOutageSimulationOutageFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOutageSimulationOutageFilter(val *ClusterOutageSimulationOutageFilter) *NullableClusterOutageSimulationOutageFilter {
	return &NullableClusterOutageSimulationOutageFilter{value: val, isSet: true}
}

func (v NullableClusterOutageSimulationOutageFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOutageSimulationOutageFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
