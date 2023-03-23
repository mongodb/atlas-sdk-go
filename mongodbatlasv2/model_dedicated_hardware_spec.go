/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the DedicatedHardwareSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedHardwareSpec{}

// DedicatedHardwareSpec Hardware specifications for read-only nodes in the region. Read-only nodes can never become the primary member, but can enable local reads.If you don't specify this parameter, no read-only nodes are deployed to the region.
type DedicatedHardwareSpec struct {
	// Number of read-only nodes for MongoDB Cloud to deploy to the region. Read-only nodes can never become the primary, but can enable local reads.
	NodeCount *int32 `json:"nodeCount,omitempty"`
	// Target throughput desired for storage attached to your AWS-provisioned cluster. Change this parameter only if you:  - set `\"replicationSpecs[n].regionConfigs[m].providerName\" : \"AWS\"`. - set `\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\" : \"M30\"` or greater not including `Mxx_NVME` tiers.  The maximum input/output operations per second (IOPS) depend on the selected **.instanceSize** and **.diskSizeGB**. This parameter defaults to the cluster tier's standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.  - Instance sizes `M10` to `M40` have a ratio of disk capacity to system memory of 60:1. - Instance sizes greater than `M40` have a ratio of 120:1.
	DiskIOPS *int32 `json:"diskIOPS,omitempty"`
	// Type of storage you want to attach to your AWS-provisioned cluster.  - `STANDARD` volume types can't exceed the default input/output operations per second (IOPS) rate for the selected volume size.   - `PROVISIONED` volume types must fall within the allowable IOPS range for the selected volume size.
	EbsVolumeType *string `json:"ebsVolumeType,omitempty"`
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
	InstanceSize *string `json:"instanceSize,omitempty"`
}

// NewDedicatedHardwareSpec instantiates a new DedicatedHardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedHardwareSpec() *DedicatedHardwareSpec {
	this := DedicatedHardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// NewDedicatedHardwareSpecWithDefaults instantiates a new DedicatedHardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedHardwareSpecWithDefaults() *DedicatedHardwareSpec {
	this := DedicatedHardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *DedicatedHardwareSpec) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedHardwareSpec) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *DedicatedHardwareSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *DedicatedHardwareSpec) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise.
func (o *DedicatedHardwareSpec) GetDiskIOPS() int32 {
	if o == nil || IsNil(o.DiskIOPS) {
		var ret int32
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedHardwareSpec) GetDiskIOPSOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskIOPS) {
		return nil, false
	}
	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *DedicatedHardwareSpec) HasDiskIOPS() bool {
	if o != nil && !IsNil(o.DiskIOPS) {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given int32 and assigns it to the DiskIOPS field.
func (o *DedicatedHardwareSpec) SetDiskIOPS(v int32) {
	o.DiskIOPS = &v
}

// GetEbsVolumeType returns the EbsVolumeType field value if set, zero value otherwise.
func (o *DedicatedHardwareSpec) GetEbsVolumeType() string {
	if o == nil || IsNil(o.EbsVolumeType) {
		var ret string
		return ret
	}
	return *o.EbsVolumeType
}

// GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedHardwareSpec) GetEbsVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EbsVolumeType) {
		return nil, false
	}
	return o.EbsVolumeType, true
}

// HasEbsVolumeType returns a boolean if a field has been set.
func (o *DedicatedHardwareSpec) HasEbsVolumeType() bool {
	if o != nil && !IsNil(o.EbsVolumeType) {
		return true
	}

	return false
}

// SetEbsVolumeType gets a reference to the given string and assigns it to the EbsVolumeType field.
func (o *DedicatedHardwareSpec) SetEbsVolumeType(v string) {
	o.EbsVolumeType = &v
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *DedicatedHardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedHardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *DedicatedHardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *DedicatedHardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

func (o DedicatedHardwareSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DedicatedHardwareSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if !IsNil(o.DiskIOPS) {
		toSerialize["diskIOPS"] = o.DiskIOPS
	}
	if !IsNil(o.EbsVolumeType) {
		toSerialize["ebsVolumeType"] = o.EbsVolumeType
	}
	if !IsNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	return toSerialize, nil
}

type NullableDedicatedHardwareSpec struct {
	value *DedicatedHardwareSpec
	isSet bool
}

func (v NullableDedicatedHardwareSpec) Get() *DedicatedHardwareSpec {
	return v.value
}

func (v *NullableDedicatedHardwareSpec) Set(val *DedicatedHardwareSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedHardwareSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedHardwareSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedHardwareSpec(val *DedicatedHardwareSpec) *NullableDedicatedHardwareSpec {
	return &NullableDedicatedHardwareSpec{value: val, isSet: true}
}

func (v NullableDedicatedHardwareSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedHardwareSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

