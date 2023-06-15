// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AWSCloudProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSCloudProviderSettings{}

// AWSCloudProviderSettings struct for AWSCloudProviderSettings
type AWSCloudProviderSettings struct {
	AutoScaling *CloudProviderAWSAutoScaling `json:"autoScaling,omitempty"`
	// Maximum Disk Input/Output Operations per Second (IOPS) that the database host can perform.
	DiskIOPS *int `json:"diskIOPS,omitempty"`
	// Flag that indicates whether the Amazon Elastic Block Store (EBS) encryption feature encrypts the host's root volume for both data at rest within the volume and for data moving between the volume and the cluster. Clusters always have this setting enabled.
	// Deprecated
	EncryptEBSVolume *bool `json:"encryptEBSVolume,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	//  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	RegionName *string `json:"regionName,omitempty"`
	// Disk Input/Output Operations per Second (IOPS) setting for Amazon Web Services (AWS) storage that you configure only for abbr title=\"Amazon Web Services\">AWS</abbr>. Specify whether Disk Input/Output Operations per Second (IOPS) must not exceed the default Input/Output Operations per Second (IOPS) rate for the selected volume size (`STANDARD`), or must fall within the allowable Input/Output Operations per Second (IOPS) range for the selected volume size (`PROVISIONED`).
	VolumeType   *string `json:"volumeType,omitempty"`
	ProviderName string  `json:"providerName"`
}

// NewAWSCloudProviderSettings instantiates a new AWSCloudProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCloudProviderSettings(providerName string) *AWSCloudProviderSettings {
	this := AWSCloudProviderSettings{}
	var encryptEBSVolume bool = true
	this.EncryptEBSVolume = &encryptEBSVolume
	this.ProviderName = providerName
	return &this
}

// NewAWSCloudProviderSettingsWithDefaults instantiates a new AWSCloudProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCloudProviderSettingsWithDefaults() *AWSCloudProviderSettings {
	this := AWSCloudProviderSettings{}
	var encryptEBSVolume bool = true
	this.EncryptEBSVolume = &encryptEBSVolume
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *AWSCloudProviderSettings) GetAutoScaling() CloudProviderAWSAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret CloudProviderAWSAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetAutoScalingOk() (*CloudProviderAWSAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given CloudProviderAWSAutoScaling and assigns it to the AutoScaling field.
func (o *AWSCloudProviderSettings) SetAutoScaling(v CloudProviderAWSAutoScaling) {
	o.AutoScaling = &v
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise.
func (o *AWSCloudProviderSettings) GetDiskIOPS() int {
	if o == nil || IsNil(o.DiskIOPS) {
		var ret int
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetDiskIOPSOk() (*int, bool) {
	if o == nil || IsNil(o.DiskIOPS) {
		return nil, false
	}
	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasDiskIOPS() bool {
	if o != nil && !IsNil(o.DiskIOPS) {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given int and assigns it to the DiskIOPS field.
func (o *AWSCloudProviderSettings) SetDiskIOPS(v int) {
	o.DiskIOPS = &v
}

// GetEncryptEBSVolume returns the EncryptEBSVolume field value if set, zero value otherwise.
// Deprecated
func (o *AWSCloudProviderSettings) GetEncryptEBSVolume() bool {
	if o == nil || IsNil(o.EncryptEBSVolume) {
		var ret bool
		return ret
	}
	return *o.EncryptEBSVolume
}

// GetEncryptEBSVolumeOk returns a tuple with the EncryptEBSVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AWSCloudProviderSettings) GetEncryptEBSVolumeOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptEBSVolume) {
		return nil, false
	}
	return o.EncryptEBSVolume, true
}

// HasEncryptEBSVolume returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasEncryptEBSVolume() bool {
	if o != nil && !IsNil(o.EncryptEBSVolume) {
		return true
	}

	return false
}

// SetEncryptEBSVolume gets a reference to the given bool and assigns it to the EncryptEBSVolume field.
// Deprecated
func (o *AWSCloudProviderSettings) SetEncryptEBSVolume(v bool) {
	o.EncryptEBSVolume = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *AWSCloudProviderSettings) GetInstanceSizeName() string {
	if o == nil || IsNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSizeName) {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasInstanceSizeName() bool {
	if o != nil && !IsNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *AWSCloudProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AWSCloudProviderSettings) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AWSCloudProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *AWSCloudProviderSettings) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *AWSCloudProviderSettings) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *AWSCloudProviderSettings) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetProviderName returns the ProviderName field value
func (o *AWSCloudProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AWSCloudProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AWSCloudProviderSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSCloudProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.DiskIOPS) {
		toSerialize["diskIOPS"] = o.DiskIOPS
	}
	if !IsNil(o.EncryptEBSVolume) {
		toSerialize["encryptEBSVolume"] = o.EncryptEBSVolume
	}
	if !IsNil(o.InstanceSizeName) {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAWSCloudProviderSettings struct {
	value *AWSCloudProviderSettings
	isSet bool
}

func (v NullableAWSCloudProviderSettings) Get() *AWSCloudProviderSettings {
	return v.value
}

func (v *NullableAWSCloudProviderSettings) Set(val *AWSCloudProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSCloudProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSCloudProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSCloudProviderSettings(val *AWSCloudProviderSettings) *NullableAWSCloudProviderSettings {
	return &NullableAWSCloudProviderSettings{value: val, isSet: true}
}

func (v NullableAWSCloudProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSCloudProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
