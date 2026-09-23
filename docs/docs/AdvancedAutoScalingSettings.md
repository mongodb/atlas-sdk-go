# AdvancedAutoScalingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**AdvancedComputeAutoScaling**](AdvancedComputeAutoScaling.md) |  | [optional] 
**DiskGB** | Pointer to [**DiskGBAutoScaling**](DiskGBAutoScaling.md) |  | [optional] 
**StorageConfig** | Pointer to [**StorageConfig**](StorageConfig.md) | Available in Public Preview: Settings that determine the per-shard data-size limit for this cluster. Applies only to Atlas INFINITE clusters. MongoDB Cloud accepts these settings only on &#x60;autoScaling&#x60; and rejects them on &#x60;analyticsAutoScaling&#x60;, including when you send them as &#x60;null&#x60;. In a request that includes &#x60;replicationSpecs&#x60;, omitting &#x60;storageConfig&#x60; or sending it as &#x60;null&#x60; or &#x60;{}&#x60; clears the limit. Omitting &#x60;replicationSpecs&#x60; preserves it. | [optional] 

## Methods

### NewAdvancedAutoScalingSettings

`func NewAdvancedAutoScalingSettings() *AdvancedAutoScalingSettings`

NewAdvancedAutoScalingSettings instantiates a new AdvancedAutoScalingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedAutoScalingSettingsWithDefaults

`func NewAdvancedAutoScalingSettingsWithDefaults() *AdvancedAutoScalingSettings`

NewAdvancedAutoScalingSettingsWithDefaults instantiates a new AdvancedAutoScalingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *AdvancedAutoScalingSettings) GetCompute() AdvancedComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *AdvancedAutoScalingSettings) GetComputeOk() (*AdvancedComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *AdvancedAutoScalingSettings) SetCompute(v AdvancedComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *AdvancedAutoScalingSettings) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *AdvancedAutoScalingSettings) SetComputeNil()`

SetComputeNil sets Compute to an explicit JSON null when marshaled, overriding any value previously set with SetCompute. Calling SetCompute again clears the null override.

### GetDiskGB

`func (o *AdvancedAutoScalingSettings) GetDiskGB() DiskGBAutoScaling`

GetDiskGB returns the DiskGB field if non-nil, zero value otherwise.

### GetDiskGBOk

`func (o *AdvancedAutoScalingSettings) GetDiskGBOk() (*DiskGBAutoScaling, bool)`

GetDiskGBOk returns a tuple with the DiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGB

`func (o *AdvancedAutoScalingSettings) SetDiskGB(v DiskGBAutoScaling)`

SetDiskGB sets DiskGB field to given value.

### HasDiskGB

`func (o *AdvancedAutoScalingSettings) HasDiskGB() bool`

HasDiskGB returns a boolean if a field has been set.

### SetDiskGBNil

`func (o *AdvancedAutoScalingSettings) SetDiskGBNil()`

SetDiskGBNil sets DiskGB to an explicit JSON null when marshaled, overriding any value previously set with SetDiskGB. Calling SetDiskGB again clears the null override.

### GetStorageConfig

`func (o *AdvancedAutoScalingSettings) GetStorageConfig() StorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *AdvancedAutoScalingSettings) GetStorageConfigOk() (*StorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *AdvancedAutoScalingSettings) SetStorageConfig(v StorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *AdvancedAutoScalingSettings) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### SetStorageConfigNil

`func (o *AdvancedAutoScalingSettings) SetStorageConfigNil()`

SetStorageConfigNil sets StorageConfig to an explicit JSON null when marshaled, overriding any value previously set with SetStorageConfig. Calling SetStorageConfig again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


