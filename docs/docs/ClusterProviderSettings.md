# ClusterProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ClusterFreeAutoScaling**](ClusterFreeAutoScaling.md) |  | [optional] 
**DiskIOPS** | Pointer to **int** | Maximum Disk Input/Output Operations per Second (IOPS) that the database host can perform. | [optional] 
**EncryptEBSVolume** | Pointer to **bool** | Flag that indicates whether the Amazon Elastic Block Store (EBS) encryption feature encrypts the host&#39;s root volume for both data at rest within the volume and for data moving between the volume and the cluster. Clusters always have this setting enabled. | [optional] [default to true]
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. You must set **providerSettings.providerName** to &#x60;TENANT&#x60; and specify the cloud service provider in **providerSettings.backingProviderName**. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). For multi-region clusters, see **replicationSpec.{region}**. | [optional] 
**VolumeType** | Pointer to **string** | Disk Input/Output Operations per Second (IOPS) setting for Amazon Web Services (AWS) storage that you configure only for abbr title&#x3D;\&quot;Amazon Web Services\&quot;&gt;AWS&lt;/abbr&gt;. Specify whether Disk Input/Output Operations per Second (IOPS) must not exceed the default Input/Output Operations per Second (IOPS) rate for the selected volume size (&#x60;STANDARD&#x60;), or must fall within the allowable Input/Output Operations per Second (IOPS) range for the selected volume size (&#x60;PROVISIONED&#x60;). | [optional] 
**ProviderName** | **string** |  | 
**DiskTypeName** | Pointer to **string** | Disk type that corresponds to the host&#39;s root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies. | [optional] 
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the multi-tenant host. The resource returns this parameter when **providerSettings.providerName** is &#x60;TENANT&#x60; and **providerSetting.instanceSizeName** is &#x60;M2&#x60; or &#x60;M5&#x60;. | [optional] 

## Methods

### NewClusterProviderSettings

`func NewClusterProviderSettings(providerName string, ) *ClusterProviderSettings`

NewClusterProviderSettings instantiates a new ClusterProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterProviderSettingsWithDefaults

`func NewClusterProviderSettingsWithDefaults() *ClusterProviderSettings`

NewClusterProviderSettingsWithDefaults instantiates a new ClusterProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ClusterProviderSettings) GetAutoScaling() ClusterFreeAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ClusterProviderSettings) GetAutoScalingOk() (*ClusterFreeAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ClusterProviderSettings) SetAutoScaling(v ClusterFreeAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ClusterProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskIOPS

`func (o *ClusterProviderSettings) GetDiskIOPS() int`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *ClusterProviderSettings) GetDiskIOPSOk() (*int, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *ClusterProviderSettings) SetDiskIOPS(v int)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *ClusterProviderSettings) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.

### GetEncryptEBSVolume

`func (o *ClusterProviderSettings) GetEncryptEBSVolume() bool`

GetEncryptEBSVolume returns the EncryptEBSVolume field if non-nil, zero value otherwise.

### GetEncryptEBSVolumeOk

`func (o *ClusterProviderSettings) GetEncryptEBSVolumeOk() (*bool, bool)`

GetEncryptEBSVolumeOk returns a tuple with the EncryptEBSVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEBSVolume

`func (o *ClusterProviderSettings) SetEncryptEBSVolume(v bool)`

SetEncryptEBSVolume sets EncryptEBSVolume field to given value.

### HasEncryptEBSVolume

`func (o *ClusterProviderSettings) HasEncryptEBSVolume() bool`

HasEncryptEBSVolume returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ClusterProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ClusterProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ClusterProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ClusterProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ClusterProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ClusterProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ClusterProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ClusterProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetVolumeType

`func (o *ClusterProviderSettings) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ClusterProviderSettings) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ClusterProviderSettings) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ClusterProviderSettings) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetProviderName

`func (o *ClusterProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ClusterProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ClusterProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetDiskTypeName

`func (o *ClusterProviderSettings) GetDiskTypeName() string`

GetDiskTypeName returns the DiskTypeName field if non-nil, zero value otherwise.

### GetDiskTypeNameOk

`func (o *ClusterProviderSettings) GetDiskTypeNameOk() (*string, bool)`

GetDiskTypeNameOk returns a tuple with the DiskTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypeName

`func (o *ClusterProviderSettings) SetDiskTypeName(v string)`

SetDiskTypeName sets DiskTypeName field to given value.

### HasDiskTypeName

`func (o *ClusterProviderSettings) HasDiskTypeName() bool`

HasDiskTypeName returns a boolean if a field has been set.

### GetBackingProviderName

`func (o *ClusterProviderSettings) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *ClusterProviderSettings) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *ClusterProviderSettings) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *ClusterProviderSettings) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


