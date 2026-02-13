# ClusterComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether instance size reactive auto-scaling is enabled.  - Set to &#x60;true&#x60; to enable instance size reactive auto-scaling. If enabled, you must specify a value for &#x60;providerSettings.autoScaling.compute.maxInstanceSize&#x60;. - Set to &#x60;false&#x60; to disable instance size reactive auto-scaling. | [optional] [default to false]
**ScaleDownEnabled** | Pointer to **bool** | Flag that indicates whether the cluster tier can scale down via reactive auto-scaling. This is required if &#x60;autoScaling.compute.enabled&#x60; is &#x60;true&#x60;. If you enable this option, specify a value for &#x60;providerSettings.autoScaling.compute.minInstanceSize&#x60;. | [optional] [default to false]

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


