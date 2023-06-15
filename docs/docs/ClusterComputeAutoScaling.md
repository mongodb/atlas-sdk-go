# ClusterComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether cluster tier auto-scaling is enabled. Set to &#x60;true&#x60; to enable cluster tier auto-scaling. If enabled, you must specify a value for **providerSettings.autoScaling.compute.maxInstanceSize** also. Set to &#x60;false&#x60; to disable cluster tier auto-scaling. | [optional] 
**ScaleDownEnabled** | Pointer to **bool** | Flag that indicates whether the cluster tier can scale down. This is required if **autoScaling.compute.enabled** is &#x60;true&#x60;. If you enable this option, specify a value for **providerSettings.autoScaling.compute.minInstanceSize**. | [optional] 

## Methods

### NewClusterComputeAutoScaling

`func NewClusterComputeAutoScaling() *ClusterComputeAutoScaling`

NewClusterComputeAutoScaling instantiates a new ClusterComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterComputeAutoScalingWithDefaults

`func NewClusterComputeAutoScalingWithDefaults() *ClusterComputeAutoScaling`

NewClusterComputeAutoScalingWithDefaults instantiates a new ClusterComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterComputeAutoScaling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterComputeAutoScaling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterComputeAutoScaling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterComputeAutoScaling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScaleDownEnabled

`func (o *ClusterComputeAutoScaling) GetScaleDownEnabled() bool`

GetScaleDownEnabled returns the ScaleDownEnabled field if non-nil, zero value otherwise.

### GetScaleDownEnabledOk

`func (o *ClusterComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool)`

GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownEnabled

`func (o *ClusterComputeAutoScaling) SetScaleDownEnabled(v bool)`

SetScaleDownEnabled sets ScaleDownEnabled field to given value.

### HasScaleDownEnabled

`func (o *ClusterComputeAutoScaling) HasScaleDownEnabled() bool`

HasScaleDownEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


