# HardwareSpec20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskSizeGB** | Pointer to **float64** | Storage capacity of instance data volumes expressed in gigabytes. Increase this number to add capacity.   This value must be equal for all shards and node types.   This value is not configurable on M0/M2/M5 clusters.   MongoDB Cloud requires this parameter if you set &#x60;replicationSpecs&#x60;.   If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value.    Storage charge calculations depend on whether you choose the default value or a custom value.   The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**DiskIOPS** | Pointer to **int** | Target throughput desired for storage attached to your Azure-provisioned cluster. Change this parameter if you:  - set &#x60;replicationSpecs[n].regionConfigs[m].providerName&#x60; : &#x60;Azure&#x60;. - set &#x60;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize&#x60; : &#x60;M40&#x60; or greater not including &#x60;Mxx_NVME&#x60; tiers.  The maximum input/output operations per second (IOPS) depend on the selected &#x60;.instanceSize&#x60; and &#x60;.diskSizeGB&#x60;. This parameter defaults to the cluster tier&#39;s standard IOPS value. Changing this value impacts cluster cost. | [optional] 
**EbsVolumeType** | Pointer to **string** | Type of storage you want to attach to your AWS-provisioned cluster.  - &#x60;STANDARD&#x60; volume types can&#39;t exceed the default input/output operations per second (IOPS) rate for the selected volume size.   - &#x60;PROVISIONED&#x60; volume types must fall within the allowable IOPS range for the selected volume size. You must set this value to (&#x60;PROVISIONED&#x60;) for NVMe clusters. | [optional] [default to "STANDARD"]
**InstanceSize** | Pointer to **string** | Hardware specification for the instances in this M0/M2/M5 tier cluster. | [optional] 
**NodeCount** | Pointer to **int** | Number of nodes of the given type for MongoDB Cloud to deploy to the region. | [optional] 
**EffectiveInstanceSize** | Pointer to **string** | The true tenant instance size. This is present to support backwards compatibility for deprecated provider types and/or instance sizes. | [optional] [readonly] 

## Methods

### NewHardwareSpec20240805

`func NewHardwareSpec20240805() *HardwareSpec20240805`

NewHardwareSpec20240805 instantiates a new HardwareSpec20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareSpec20240805WithDefaults

`func NewHardwareSpec20240805WithDefaults() *HardwareSpec20240805`

NewHardwareSpec20240805WithDefaults instantiates a new HardwareSpec20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskSizeGB

`func (o *HardwareSpec20240805) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *HardwareSpec20240805) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *HardwareSpec20240805) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *HardwareSpec20240805) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.
### GetDiskIOPS

`func (o *HardwareSpec20240805) GetDiskIOPS() int`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *HardwareSpec20240805) GetDiskIOPSOk() (*int, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *HardwareSpec20240805) SetDiskIOPS(v int)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *HardwareSpec20240805) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.
### GetEbsVolumeType

`func (o *HardwareSpec20240805) GetEbsVolumeType() string`

GetEbsVolumeType returns the EbsVolumeType field if non-nil, zero value otherwise.

### GetEbsVolumeTypeOk

`func (o *HardwareSpec20240805) GetEbsVolumeTypeOk() (*string, bool)`

GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeType

`func (o *HardwareSpec20240805) SetEbsVolumeType(v string)`

SetEbsVolumeType sets EbsVolumeType field to given value.

### HasEbsVolumeType

`func (o *HardwareSpec20240805) HasEbsVolumeType() bool`

HasEbsVolumeType returns a boolean if a field has been set.
### GetInstanceSize

`func (o *HardwareSpec20240805) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *HardwareSpec20240805) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *HardwareSpec20240805) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *HardwareSpec20240805) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.
### GetNodeCount

`func (o *HardwareSpec20240805) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *HardwareSpec20240805) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *HardwareSpec20240805) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *HardwareSpec20240805) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.
### GetEffectiveInstanceSize

`func (o *HardwareSpec20240805) GetEffectiveInstanceSize() string`

GetEffectiveInstanceSize returns the EffectiveInstanceSize field if non-nil, zero value otherwise.

### GetEffectiveInstanceSizeOk

`func (o *HardwareSpec20240805) GetEffectiveInstanceSizeOk() (*string, bool)`

GetEffectiveInstanceSizeOk returns a tuple with the EffectiveInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveInstanceSize

`func (o *HardwareSpec20240805) SetEffectiveInstanceSize(v string)`

SetEffectiveInstanceSize sets EffectiveInstanceSize field to given value.

### HasEffectiveInstanceSize

`func (o *HardwareSpec20240805) HasEffectiveInstanceSize() bool`

HasEffectiveInstanceSize returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


