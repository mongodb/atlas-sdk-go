# AWSCloudProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**CloudProviderAWSAutoScaling**](CloudProviderAWSAutoScaling.md) |  | [optional] 
**DiskIOPS** | Pointer to **int** | Maximum Disk Input/Output Operations per Second (IOPS) that the database host can perform. | [optional] 
**EncryptEBSVolume** | Pointer to **bool** | Flag that indicates whether the Amazon Elastic Block Store (EBS) encryption feature encrypts the host&#39;s root volume for both data at rest within the volume and for data moving between the volume and the cluster. Clusters always have this setting enabled. | [optional] [default to true]
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** |  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 
**VolumeType** | Pointer to **string** | Disk Input/Output Operations per Second (IOPS) setting for Amazon Web Services (AWS) storage that you configure only for abbr title&#x3D;\&quot;Amazon Web Services\&quot;&gt;AWS&lt;/abbr&gt;. Specify whether Disk Input/Output Operations per Second (IOPS) must not exceed the default Input/Output Operations per Second (IOPS) rate for the selected volume size (&#x60;STANDARD&#x60;), or must fall within the allowable Input/Output Operations per Second (IOPS) range for the selected volume size (&#x60;PROVISIONED&#x60;). | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewAWSCloudProviderSettings

`func NewAWSCloudProviderSettings(providerName string, ) *AWSCloudProviderSettings`

NewAWSCloudProviderSettings instantiates a new AWSCloudProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCloudProviderSettingsWithDefaults

`func NewAWSCloudProviderSettingsWithDefaults() *AWSCloudProviderSettings`

NewAWSCloudProviderSettingsWithDefaults instantiates a new AWSCloudProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *AWSCloudProviderSettings) GetAutoScaling() CloudProviderAWSAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *AWSCloudProviderSettings) GetAutoScalingOk() (*CloudProviderAWSAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *AWSCloudProviderSettings) SetAutoScaling(v CloudProviderAWSAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *AWSCloudProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskIOPS

`func (o *AWSCloudProviderSettings) GetDiskIOPS() int`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *AWSCloudProviderSettings) GetDiskIOPSOk() (*int, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *AWSCloudProviderSettings) SetDiskIOPS(v int)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *AWSCloudProviderSettings) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.

### GetEncryptEBSVolume

`func (o *AWSCloudProviderSettings) GetEncryptEBSVolume() bool`

GetEncryptEBSVolume returns the EncryptEBSVolume field if non-nil, zero value otherwise.

### GetEncryptEBSVolumeOk

`func (o *AWSCloudProviderSettings) GetEncryptEBSVolumeOk() (*bool, bool)`

GetEncryptEBSVolumeOk returns a tuple with the EncryptEBSVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEBSVolume

`func (o *AWSCloudProviderSettings) SetEncryptEBSVolume(v bool)`

SetEncryptEBSVolume sets EncryptEBSVolume field to given value.

### HasEncryptEBSVolume

`func (o *AWSCloudProviderSettings) HasEncryptEBSVolume() bool`

HasEncryptEBSVolume returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *AWSCloudProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *AWSCloudProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *AWSCloudProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *AWSCloudProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *AWSCloudProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AWSCloudProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AWSCloudProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AWSCloudProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetVolumeType

`func (o *AWSCloudProviderSettings) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AWSCloudProviderSettings) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AWSCloudProviderSettings) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AWSCloudProviderSettings) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetProviderName

`func (o *AWSCloudProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AWSCloudProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AWSCloudProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


