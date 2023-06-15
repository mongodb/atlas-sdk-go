# ClusterFreeProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ClusterFreeAutoScaling**](ClusterFreeAutoScaling.md) |  | [optional] 
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the multi-tenant host. The resource returns this parameter when **providerSettings.providerName** is &#x60;TENANT&#x60; and **providerSetting.instanceSizeName** is &#x60;M2&#x60; or &#x60;M5&#x60;. | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. You must set **providerSettings.providerName** to &#x60;TENANT&#x60; and specify the cloud service provider in **providerSettings.backingProviderName**. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). For multi-region clusters, see **replicationSpec.{region}**. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewClusterFreeProviderSettings

`func NewClusterFreeProviderSettings(providerName string, ) *ClusterFreeProviderSettings`

NewClusterFreeProviderSettings instantiates a new ClusterFreeProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFreeProviderSettingsWithDefaults

`func NewClusterFreeProviderSettingsWithDefaults() *ClusterFreeProviderSettings`

NewClusterFreeProviderSettingsWithDefaults instantiates a new ClusterFreeProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ClusterFreeProviderSettings) GetAutoScaling() ClusterFreeAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ClusterFreeProviderSettings) GetAutoScalingOk() (*ClusterFreeAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ClusterFreeProviderSettings) SetAutoScaling(v ClusterFreeAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ClusterFreeProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetBackingProviderName

`func (o *ClusterFreeProviderSettings) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *ClusterFreeProviderSettings) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *ClusterFreeProviderSettings) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *ClusterFreeProviderSettings) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ClusterFreeProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ClusterFreeProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ClusterFreeProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ClusterFreeProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ClusterFreeProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ClusterFreeProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ClusterFreeProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ClusterFreeProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetProviderName

`func (o *ClusterFreeProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ClusterFreeProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ClusterFreeProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


