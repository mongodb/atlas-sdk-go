# ClusterAutoScalingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**ClusterComputeAutoScaling**](ClusterComputeAutoScaling.md) |  | [optional] 
**DiskGBEnabled** | Pointer to **bool** | Flag that indicates whether someone enabled disk auto-scaling for this cluster. | [optional] 

## Methods

### NewClusterAutoScalingSettings

`func NewClusterAutoScalingSettings() *ClusterAutoScalingSettings`

NewClusterAutoScalingSettings instantiates a new ClusterAutoScalingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAutoScalingSettingsWithDefaults

`func NewClusterAutoScalingSettingsWithDefaults() *ClusterAutoScalingSettings`

NewClusterAutoScalingSettingsWithDefaults instantiates a new ClusterAutoScalingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *ClusterAutoScalingSettings) GetCompute() ClusterComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *ClusterAutoScalingSettings) GetComputeOk() (*ClusterComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *ClusterAutoScalingSettings) SetCompute(v ClusterComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *ClusterAutoScalingSettings) HasCompute() bool`

HasCompute returns a boolean if a field has been set.
### GetDiskGBEnabled

`func (o *ClusterAutoScalingSettings) GetDiskGBEnabled() bool`

GetDiskGBEnabled returns the DiskGBEnabled field if non-nil, zero value otherwise.

### GetDiskGBEnabledOk

`func (o *ClusterAutoScalingSettings) GetDiskGBEnabledOk() (*bool, bool)`

GetDiskGBEnabledOk returns a tuple with the DiskGBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGBEnabled

`func (o *ClusterAutoScalingSettings) SetDiskGBEnabled(v bool)`

SetDiskGBEnabled sets DiskGBEnabled field to given value.

### HasDiskGBEnabled

`func (o *ClusterAutoScalingSettings) HasDiskGBEnabled() bool`

HasDiskGBEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


