// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AzureCloudProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCloudProviderSettings{}

// AzureCloudProviderSettings struct for AzureCloudProviderSettings
type AzureCloudProviderSettings struct {
	AutoScaling *CloudProviderAzureAutoScaling `json:"autoScaling,omitempty"`
	// Disk type that corresponds to the host's root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies.
	DiskTypeName *string `json:"diskTypeName,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Microsoft Azure Regions.
	RegionName   *string `json:"regionName,omitempty"`
	ProviderName string  `json:"providerName"`
}

// NewAzureCloudProviderSettings instantiates a new AzureCloudProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCloudProviderSettings(providerName string) *AzureCloudProviderSettings {
	this := AzureCloudProviderSettings{}
	this.ProviderName = providerName
	return &this
}

// NewAzureCloudProviderSettingsWithDefaults instantiates a new AzureCloudProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCloudProviderSettingsWithDefaults() *AzureCloudProviderSettings {
	this := AzureCloudProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *AzureCloudProviderSettings) GetAutoScaling() CloudProviderAzureAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret CloudProviderAzureAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderSettings) GetAutoScalingOk() (*CloudProviderAzureAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *AzureCloudProviderSettings) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given CloudProviderAzureAutoScaling and assigns it to the AutoScaling field.
func (o *AzureCloudProviderSettings) SetAutoScaling(v CloudProviderAzureAutoScaling) {
	o.AutoScaling = &v
}

// GetDiskTypeName returns the DiskTypeName field value if set, zero value otherwise.
func (o *AzureCloudProviderSettings) GetDiskTypeName() string {
	if o == nil || IsNil(o.DiskTypeName) {
		var ret string
		return ret
	}
	return *o.DiskTypeName
}

// GetDiskTypeNameOk returns a tuple with the DiskTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderSettings) GetDiskTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.DiskTypeName) {
		return nil, false
	}
	return o.DiskTypeName, true
}

// HasDiskTypeName returns a boolean if a field has been set.
func (o *AzureCloudProviderSettings) HasDiskTypeName() bool {
	if o != nil && !IsNil(o.DiskTypeName) {
		return true
	}

	return false
}

// SetDiskTypeName gets a reference to the given string and assigns it to the DiskTypeName field.
func (o *AzureCloudProviderSettings) SetDiskTypeName(v string) {
	o.DiskTypeName = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *AzureCloudProviderSettings) GetInstanceSizeName() string {
	if o == nil || IsNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSizeName) {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *AzureCloudProviderSettings) HasInstanceSizeName() bool {
	if o != nil && !IsNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *AzureCloudProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AzureCloudProviderSettings) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AzureCloudProviderSettings) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AzureCloudProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetProviderName returns the ProviderName field value
func (o *AzureCloudProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AzureCloudProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AzureCloudProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AzureCloudProviderSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureCloudProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.DiskTypeName) {
		toSerialize["diskTypeName"] = o.DiskTypeName
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

type NullableAzureCloudProviderSettings struct {
	value *AzureCloudProviderSettings
	isSet bool
}

func (v NullableAzureCloudProviderSettings) Get() *AzureCloudProviderSettings {
	return v.value
}

func (v *NullableAzureCloudProviderSettings) Set(val *AzureCloudProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCloudProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCloudProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCloudProviderSettings(val *AzureCloudProviderSettings) *NullableAzureCloudProviderSettings {
	return &NullableAzureCloudProviderSettings{value: val, isSet: true}
}

func (v NullableAzureCloudProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCloudProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
