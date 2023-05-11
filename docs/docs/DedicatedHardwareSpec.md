# DedicatedHardwareSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int** | Number of nodes of the given type for MongoDB Cloud to deploy to the region. | [optional] 
**DiskIOPS** | Pointer to **int** | Target throughput desired for storage attached to your AWS-provisioned cluster. Change this parameter only if you:  - set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].providerName\&quot; : \&quot;AWS\&quot;&#x60;. - set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\&quot; : \&quot;M30\&quot;&#x60; or greater not including &#x60;Mxx_NVME&#x60; tiers.  The maximum input/output operations per second (IOPS) depend on the selected **.instanceSize** and **.diskSizeGB**. This parameter defaults to the cluster tier&#39;s standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.  - Instance sizes &#x60;M10&#x60; to &#x60;M40&#x60; have a ratio of disk capacity to system memory of 60:1. - Instance sizes greater than &#x60;M40&#x60; have a ratio of 120:1. | [optional] 
**EbsVolumeType** | Pointer to **string** | Type of storage you want to attach to your AWS-provisioned cluster.  - &#x60;STANDARD&#x60; volume types can&#39;t exceed the default input/output operations per second (IOPS) rate for the selected volume size.   - &#x60;PROVISIONED&#x60; volume types must fall within the allowable IOPS range for the selected volume size. | [optional] [default to "STANDARD"]
**InstanceSize** | Pointer to **string** | Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. | [optional] 

## Methods

### NewDedicatedHardwareSpec

`func NewDedicatedHardwareSpec() *DedicatedHardwareSpec`

NewDedicatedHardwareSpec instantiates a new DedicatedHardwareSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedHardwareSpecWithDefaults

`func NewDedicatedHardwareSpecWithDefaults() *DedicatedHardwareSpec`

NewDedicatedHardwareSpecWithDefaults instantiates a new DedicatedHardwareSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *DedicatedHardwareSpec) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *DedicatedHardwareSpec) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *DedicatedHardwareSpec) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *DedicatedHardwareSpec) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetDiskIOPS

`func (o *DedicatedHardwareSpec) GetDiskIOPS() int`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *DedicatedHardwareSpec) GetDiskIOPSOk() (*int, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *DedicatedHardwareSpec) SetDiskIOPS(v int)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *DedicatedHardwareSpec) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.

### GetEbsVolumeType

`func (o *DedicatedHardwareSpec) GetEbsVolumeType() string`

GetEbsVolumeType returns the EbsVolumeType field if non-nil, zero value otherwise.

### GetEbsVolumeTypeOk

`func (o *DedicatedHardwareSpec) GetEbsVolumeTypeOk() (*string, bool)`

GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeType

`func (o *DedicatedHardwareSpec) SetEbsVolumeType(v string)`

SetEbsVolumeType sets EbsVolumeType field to given value.

### HasEbsVolumeType

`func (o *DedicatedHardwareSpec) HasEbsVolumeType() bool`

HasEbsVolumeType returns a boolean if a field has been set.

### GetInstanceSize

`func (o *DedicatedHardwareSpec) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *DedicatedHardwareSpec) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *DedicatedHardwareSpec) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *DedicatedHardwareSpec) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


