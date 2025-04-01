# AdvancedComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether instance size reactive auto-scaling is enabled.  - Set to &#x60;true&#x60; to enable instance size reactive auto-scaling. If enabled, you must specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize**. - Set to &#x60;false&#x60; to disable instance size reactive auto-scaling. | [optional] 
**MaxInstanceSize** | Pointer to **string** | Instance size boundary to which your cluster can automatically scale. | [optional] [readonly] 
**MinInstanceSize** | Pointer to **string** | Instance size boundary to which your cluster can automatically scale. | [optional] [readonly] 
**PredictiveEnabled** | Pointer to **bool** | Flag that indicates whether predictive instance size auto-scaling is enabled.  - Set to &#x60;true&#x60; to enable predictive instance size auto-scaling. MongoDB Cloud requires **replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled** to be &#x60;true&#x60; in order to enable this feature. - Set to &#x60;false&#x60; to disable predictive instance size auto-scaling. | [optional] 
**ScaleDownEnabled** | Pointer to **bool** | Flag that indicates whether the instance size may scale down via reactive auto-scaling. MongoDB Cloud requires this parameter if **replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled** is &#x60;true&#x60;. If you enable this option, specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize**. | [optional] 

## Methods

### NewAdvancedComputeAutoScaling

`func NewAdvancedComputeAutoScaling() *AdvancedComputeAutoScaling`

NewAdvancedComputeAutoScaling instantiates a new AdvancedComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedComputeAutoScalingWithDefaults

`func NewAdvancedComputeAutoScalingWithDefaults() *AdvancedComputeAutoScaling`

NewAdvancedComputeAutoScalingWithDefaults instantiates a new AdvancedComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AdvancedComputeAutoScaling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdvancedComputeAutoScaling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdvancedComputeAutoScaling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdvancedComputeAutoScaling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.
### GetMaxInstanceSize

`func (o *AdvancedComputeAutoScaling) GetMaxInstanceSize() string`

GetMaxInstanceSize returns the MaxInstanceSize field if non-nil, zero value otherwise.

### GetMaxInstanceSizeOk

`func (o *AdvancedComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool)`

GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceSize

`func (o *AdvancedComputeAutoScaling) SetMaxInstanceSize(v string)`

SetMaxInstanceSize sets MaxInstanceSize field to given value.

### HasMaxInstanceSize

`func (o *AdvancedComputeAutoScaling) HasMaxInstanceSize() bool`

HasMaxInstanceSize returns a boolean if a field has been set.
### GetMinInstanceSize

`func (o *AdvancedComputeAutoScaling) GetMinInstanceSize() string`

GetMinInstanceSize returns the MinInstanceSize field if non-nil, zero value otherwise.

### GetMinInstanceSizeOk

`func (o *AdvancedComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool)`

GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceSize

`func (o *AdvancedComputeAutoScaling) SetMinInstanceSize(v string)`

SetMinInstanceSize sets MinInstanceSize field to given value.

### HasMinInstanceSize

`func (o *AdvancedComputeAutoScaling) HasMinInstanceSize() bool`

HasMinInstanceSize returns a boolean if a field has been set.
### GetPredictiveEnabled

`func (o *AdvancedComputeAutoScaling) GetPredictiveEnabled() bool`

GetPredictiveEnabled returns the PredictiveEnabled field if non-nil, zero value otherwise.

### GetPredictiveEnabledOk

`func (o *AdvancedComputeAutoScaling) GetPredictiveEnabledOk() (*bool, bool)`

GetPredictiveEnabledOk returns a tuple with the PredictiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictiveEnabled

`func (o *AdvancedComputeAutoScaling) SetPredictiveEnabled(v bool)`

SetPredictiveEnabled sets PredictiveEnabled field to given value.

### HasPredictiveEnabled

`func (o *AdvancedComputeAutoScaling) HasPredictiveEnabled() bool`

HasPredictiveEnabled returns a boolean if a field has been set.
### GetScaleDownEnabled

`func (o *AdvancedComputeAutoScaling) GetScaleDownEnabled() bool`

GetScaleDownEnabled returns the ScaleDownEnabled field if non-nil, zero value otherwise.

### GetScaleDownEnabledOk

`func (o *AdvancedComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool)`

GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownEnabled

`func (o *AdvancedComputeAutoScaling) SetScaleDownEnabled(v bool)`

SetScaleDownEnabled sets ScaleDownEnabled field to given value.

### HasScaleDownEnabled

`func (o *AdvancedComputeAutoScaling) HasScaleDownEnabled() bool`

HasScaleDownEnabled returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


