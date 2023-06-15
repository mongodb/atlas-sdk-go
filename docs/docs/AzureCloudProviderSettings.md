# AzureCloudProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**CloudProviderAzureAutoScaling**](CloudProviderAzureAutoScaling.md) |  | [optional] 
**DiskTypeName** | Pointer to **string** | Disk type that corresponds to the host&#39;s root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies. | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Microsoft Azure Regions. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewAzureCloudProviderSettings

`func NewAzureCloudProviderSettings(providerName string, ) *AzureCloudProviderSettings`

NewAzureCloudProviderSettings instantiates a new AzureCloudProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudProviderSettingsWithDefaults

`func NewAzureCloudProviderSettingsWithDefaults() *AzureCloudProviderSettings`

NewAzureCloudProviderSettingsWithDefaults instantiates a new AzureCloudProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *AzureCloudProviderSettings) GetAutoScaling() CloudProviderAzureAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *AzureCloudProviderSettings) GetAutoScalingOk() (*CloudProviderAzureAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *AzureCloudProviderSettings) SetAutoScaling(v CloudProviderAzureAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *AzureCloudProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskTypeName

`func (o *AzureCloudProviderSettings) GetDiskTypeName() string`

GetDiskTypeName returns the DiskTypeName field if non-nil, zero value otherwise.

### GetDiskTypeNameOk

`func (o *AzureCloudProviderSettings) GetDiskTypeNameOk() (*string, bool)`

GetDiskTypeNameOk returns a tuple with the DiskTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypeName

`func (o *AzureCloudProviderSettings) SetDiskTypeName(v string)`

SetDiskTypeName sets DiskTypeName field to given value.

### HasDiskTypeName

`func (o *AzureCloudProviderSettings) HasDiskTypeName() bool`

HasDiskTypeName returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *AzureCloudProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *AzureCloudProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *AzureCloudProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *AzureCloudProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *AzureCloudProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AzureCloudProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AzureCloudProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AzureCloudProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetProviderName

`func (o *AzureCloudProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AzureCloudProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AzureCloudProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


