# AdvancedAutoScalingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**AdvancedComputeAutoScaling**](AdvancedComputeAutoScaling.md) |  | [optional] 
**DiskGB** | Pointer to [**DiskGBAutoScaling**](DiskGBAutoScaling.md) |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


