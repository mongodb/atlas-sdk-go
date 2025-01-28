// Code based on the AtlasAPI V2 OpenAPI file

package admin

// HardwareSpec Hardware specifications for all electable nodes deployed in the region. Electable nodes can become the primary and can enable local reads. If you don't specify this option, MongoDB Cloud deploys no electable nodes to the region.
type HardwareSpec struct {
	// Target throughput desired for storage attached to your Azure-provisioned cluster. Change this parameter if you:  - set `\"replicationSpecs[n].regionConfigs[m].providerName\" : \"Azure\"`. - set `\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\" : \"M40\"` or greater not including `Mxx_NVME` tiers.  The maximum input/output operations per second (IOPS) depend on the selected **.instanceSize** and **.diskSizeGB**. This parameter defaults to the cluster tier's standard IOPS value. Changing this value impacts cluster cost.
	DiskIOPS *int `json:"diskIOPS,omitempty"`
	// Type of storage you want to attach to your AWS-provisioned cluster.  - `STANDARD` volume types can't exceed the default input/output operations per second (IOPS) rate for the selected volume size.   - `PROVISIONED` volume types must fall within the allowable IOPS range for the selected volume size. You must set this value to (`PROVISIONED`) for NVMe clusters.
	EbsVolumeType *string `json:"ebsVolumeType,omitempty"`
	// Hardware specification for the instances in this M0/M2/M5 tier cluster.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of nodes of the given type for MongoDB Cloud to deploy to the region.
	NodeCount *int `json:"nodeCount,omitempty"`
	// The true tenant instance size. This is present to support backwards compatibility for deprecated provider types and/or instance sizes.
	// Read only field.
	EffectiveInstanceSize *string `json:"effectiveInstanceSize,omitempty"`
}

// NewHardwareSpec instantiates a new HardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareSpec() *HardwareSpec {
	this := HardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// NewHardwareSpecWithDefaults instantiates a new HardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareSpecWithDefaults() *HardwareSpec {
	this := HardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise
func (o *HardwareSpec) GetDiskIOPS() int {
	if o == nil || IsNil(o.DiskIOPS) {
		var ret int
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetDiskIOPSOk() (*int, bool) {
	if o == nil || IsNil(o.DiskIOPS) {
		return nil, false
	}

	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *HardwareSpec) HasDiskIOPS() bool {
	if o != nil && !IsNil(o.DiskIOPS) {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given int and assigns it to the DiskIOPS field.
func (o *HardwareSpec) SetDiskIOPS(v int) {
	o.DiskIOPS = &v
}

// GetEbsVolumeType returns the EbsVolumeType field value if set, zero value otherwise
func (o *HardwareSpec) GetEbsVolumeType() string {
	if o == nil || IsNil(o.EbsVolumeType) {
		var ret string
		return ret
	}
	return *o.EbsVolumeType
}

// GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetEbsVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EbsVolumeType) {
		return nil, false
	}

	return o.EbsVolumeType, true
}

// HasEbsVolumeType returns a boolean if a field has been set.
func (o *HardwareSpec) HasEbsVolumeType() bool {
	if o != nil && !IsNil(o.EbsVolumeType) {
		return true
	}

	return false
}

// SetEbsVolumeType gets a reference to the given string and assigns it to the EbsVolumeType field.
func (o *HardwareSpec) SetEbsVolumeType(v string) {
	o.EbsVolumeType = &v
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *HardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *HardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *HardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *HardwareSpec) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *HardwareSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *HardwareSpec) SetNodeCount(v int) {
	o.NodeCount = &v
}

// GetEffectiveInstanceSize returns the EffectiveInstanceSize field value if set, zero value otherwise
func (o *HardwareSpec) GetEffectiveInstanceSize() string {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		var ret string
		return ret
	}
	return *o.EffectiveInstanceSize
}

// GetEffectiveInstanceSizeOk returns a tuple with the EffectiveInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetEffectiveInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveInstanceSize) {
		return nil, false
	}

	return o.EffectiveInstanceSize, true
}

// HasEffectiveInstanceSize returns a boolean if a field has been set.
func (o *HardwareSpec) HasEffectiveInstanceSize() bool {
	if o != nil && !IsNil(o.EffectiveInstanceSize) {
		return true
	}

	return false
}

// SetEffectiveInstanceSize gets a reference to the given string and assigns it to the EffectiveInstanceSize field.
func (o *HardwareSpec) SetEffectiveInstanceSize(v string) {
	o.EffectiveInstanceSize = &v
}
