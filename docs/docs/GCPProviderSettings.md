# GCPProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**GCPAutoScaling**](GCPAutoScaling.md) |  | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Google Compute Regions. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewGCPProviderSettings

`func NewGCPProviderSettings(providerName string, ) *GCPProviderSettings`

NewGCPProviderSettings instantiates a new GCPProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPProviderSettingsWithDefaults

`func NewGCPProviderSettingsWithDefaults() *GCPProviderSettings`

NewGCPProviderSettingsWithDefaults instantiates a new GCPProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *GCPProviderSettings) GetAutoScaling() GCPAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *GCPProviderSettings) GetAutoScalingOk() (*GCPAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *GCPProviderSettings) SetAutoScaling(v GCPAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *GCPProviderSettings) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *GCPProviderSettings) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *GCPProviderSettings) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *GCPProviderSettings) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *GCPProviderSettings) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *GCPProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GCPProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GCPProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GCPProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetProviderName

`func (o *GCPProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GCPProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GCPProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


