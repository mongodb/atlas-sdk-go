# AzureProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**AzureAutoScaling**](AzureAutoScaling.md) |  | [optional] 
**DiskTypeName** | Pointer to **string** | Disk type that corresponds to the host&#39;s root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies. | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Microsoft Azure Regions. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewAzureProviderSettings

`func NewAzureProviderSettings(providerName string, ) *AzureProviderSettings`

NewAzureProviderSettings instantiates a new AzureProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProviderSettingsWithDefaults

`func NewAzureProviderSettingsWithDefaults() *AzureProviderSettings`

NewAzureProviderSettingsWithDefaults instantiates a new AzureProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *AzureProviderSettings) GetAutoScaling() AzureAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *AzureProviderSettings) GetAutoScalingOk() (*AzureAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *AzureProviderSettings) SetAutoScaling(v AzureAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *AzureProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskTypeName

`func (o *AzureProviderSettings) GetDiskTypeName() string`

GetDiskTypeName returns the DiskTypeName field if non-nil, zero value otherwise.

### GetDiskTypeNameOk

`func (o *AzureProviderSettings) GetDiskTypeNameOk() (*string, bool)`

GetDiskTypeNameOk returns a tuple with the DiskTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypeName

`func (o *AzureProviderSettings) SetDiskTypeName(v string)`

SetDiskTypeName sets DiskTypeName field to given value.

### HasDiskTypeName

`func (o *AzureProviderSettings) HasDiskTypeName() bool`

HasDiskTypeName returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *AzureProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *AzureProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *AzureProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *AzureProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *AzureProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AzureProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AzureProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AzureProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetProviderName

`func (o *AzureProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AzureProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AzureProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


