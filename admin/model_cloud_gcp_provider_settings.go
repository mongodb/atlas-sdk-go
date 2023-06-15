// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudGCPProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudGCPProviderSettings{}

// CloudGCPProviderSettings struct for CloudGCPProviderSettings
type CloudGCPProviderSettings struct {
	AutoScaling *CloudProviderGCPAutoScaling `json:"autoScaling,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Google Compute Regions.
	RegionName   *string `json:"regionName,omitempty"`
	ProviderName string  `json:"providerName"`
}

// NewCloudGCPProviderSettings instantiates a new CloudGCPProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudGCPProviderSettings(providerName string) *CloudGCPProviderSettings {
	this := CloudGCPProviderSettings{}
	this.ProviderName = providerName
	return &this
}

// NewCloudGCPProviderSettingsWithDefaults instantiates a new CloudGCPProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudGCPProviderSettingsWithDefaults() *CloudGCPProviderSettings {
	this := CloudGCPProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *CloudGCPProviderSettings) GetAutoScaling() CloudProviderGCPAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret CloudProviderGCPAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudGCPProviderSettings) GetAutoScalingOk() (*CloudProviderGCPAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *CloudGCPProviderSettings) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given CloudProviderGCPAutoScaling and assigns it to the AutoScaling field.
func (o *CloudGCPProviderSettings) SetAutoScaling(v CloudProviderGCPAutoScaling) {
	o.AutoScaling = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *CloudGCPProviderSettings) GetInstanceSizeName() string {
	if o == nil || IsNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudGCPProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSizeName) {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *CloudGCPProviderSettings) HasInstanceSizeName() bool {
	if o != nil && !IsNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *CloudGCPProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *CloudGCPProviderSettings) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudGCPProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudGCPProviderSettings) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudGCPProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetProviderName returns the ProviderName field value
func (o *CloudGCPProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudGCPProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudGCPProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o CloudGCPProviderSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudGCPProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
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

type NullableCloudGCPProviderSettings struct {
	value *CloudGCPProviderSettings
	isSet bool
}

func (v NullableCloudGCPProviderSettings) Get() *CloudGCPProviderSettings {
	return v.value
}

func (v *NullableCloudGCPProviderSettings) Set(val *CloudGCPProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudGCPProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudGCPProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudGCPProviderSettings(val *CloudGCPProviderSettings) *NullableCloudGCPProviderSettings {
	return &NullableCloudGCPProviderSettings{value: val, isSet: true}
}

func (v NullableCloudGCPProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudGCPProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
