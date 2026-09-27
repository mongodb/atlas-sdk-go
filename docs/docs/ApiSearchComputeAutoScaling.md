# ApiSearchComputeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether Atlas raises the Search Node tier when the nodes are under sustained load. If set to &#x60;true&#x60;, you must also set &#x60;maxInstanceTier&#x60;. | [optional] [default to false]
**MaxInstanceTier** | Pointer to **string** | Highest Search Node tier that Atlas can scale up to. Required when &#x60;enabled&#x60; is &#x60;true&#x60;. | [optional] 
**MinInstanceTier** | Pointer to **string** | Lowest Search Node tier that Atlas can scale down to. Required when &#x60;scaleDownEnabled&#x60; is &#x60;true&#x60;. Scaling down is not supported yet, so setting this returns an error. | [optional] 
**ScaleDownEnabled** | Pointer to **bool** | Flag that indicates whether Atlas lowers the Search Node tier when load drops. Takes effect only when &#x60;enabled&#x60; is &#x60;true&#x60;. If set to &#x60;true&#x60;, you must also set &#x60;minInstanceTier&#x60;. Scaling down is not supported yet, so setting this to &#x60;true&#x60; returns an error. | [optional] [default to false]

## Methods

### NewApiSearchComputeAutoScaling

`func NewApiSearchComputeAutoScaling() *ApiSearchComputeAutoScaling`

NewApiSearchComputeAutoScaling instantiates a new ApiSearchComputeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchComputeAutoScalingWithDefaults

`func NewApiSearchComputeAutoScalingWithDefaults() *ApiSearchComputeAutoScaling`

NewApiSearchComputeAutoScalingWithDefaults instantiates a new ApiSearchComputeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiSearchComputeAutoScaling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiSearchComputeAutoScaling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiSearchComputeAutoScaling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiSearchComputeAutoScaling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ApiSearchComputeAutoScaling) SetEnabledNil()`

SetEnabledNil sets Enabled to an explicit JSON null when marshaled, overriding any value previously set with SetEnabled. Calling SetEnabled again clears the null override.

### GetMaxInstanceTier

`func (o *ApiSearchComputeAutoScaling) GetMaxInstanceTier() string`

GetMaxInstanceTier returns the MaxInstanceTier field if non-nil, zero value otherwise.

### GetMaxInstanceTierOk

`func (o *ApiSearchComputeAutoScaling) GetMaxInstanceTierOk() (*string, bool)`

GetMaxInstanceTierOk returns a tuple with the MaxInstanceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceTier

`func (o *ApiSearchComputeAutoScaling) SetMaxInstanceTier(v string)`

SetMaxInstanceTier sets MaxInstanceTier field to given value.

### HasMaxInstanceTier

`func (o *ApiSearchComputeAutoScaling) HasMaxInstanceTier() bool`

HasMaxInstanceTier returns a boolean if a field has been set.

### SetMaxInstanceTierNil

`func (o *ApiSearchComputeAutoScaling) SetMaxInstanceTierNil()`

SetMaxInstanceTierNil sets MaxInstanceTier to an explicit JSON null when marshaled, overriding any value previously set with SetMaxInstanceTier. Calling SetMaxInstanceTier again clears the null override.

### GetMinInstanceTier

`func (o *ApiSearchComputeAutoScaling) GetMinInstanceTier() string`

GetMinInstanceTier returns the MinInstanceTier field if non-nil, zero value otherwise.

### GetMinInstanceTierOk

`func (o *ApiSearchComputeAutoScaling) GetMinInstanceTierOk() (*string, bool)`

GetMinInstanceTierOk returns a tuple with the MinInstanceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceTier

`func (o *ApiSearchComputeAutoScaling) SetMinInstanceTier(v string)`

SetMinInstanceTier sets MinInstanceTier field to given value.

### HasMinInstanceTier

`func (o *ApiSearchComputeAutoScaling) HasMinInstanceTier() bool`

HasMinInstanceTier returns a boolean if a field has been set.

### SetMinInstanceTierNil

`func (o *ApiSearchComputeAutoScaling) SetMinInstanceTierNil()`

SetMinInstanceTierNil sets MinInstanceTier to an explicit JSON null when marshaled, overriding any value previously set with SetMinInstanceTier. Calling SetMinInstanceTier again clears the null override.

### GetScaleDownEnabled

`func (o *ApiSearchComputeAutoScaling) GetScaleDownEnabled() bool`

GetScaleDownEnabled returns the ScaleDownEnabled field if non-nil, zero value otherwise.

### GetScaleDownEnabledOk

`func (o *ApiSearchComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool)`

GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownEnabled

`func (o *ApiSearchComputeAutoScaling) SetScaleDownEnabled(v bool)`

SetScaleDownEnabled sets ScaleDownEnabled field to given value.

### HasScaleDownEnabled

`func (o *ApiSearchComputeAutoScaling) HasScaleDownEnabled() bool`

HasScaleDownEnabled returns a boolean if a field has been set.

### SetScaleDownEnabledNil

`func (o *ApiSearchComputeAutoScaling) SetScaleDownEnabledNil()`

SetScaleDownEnabledNil sets ScaleDownEnabled to an explicit JSON null when marshaled, overriding any value previously set with SetScaleDownEnabled. Calling SetScaleDownEnabled again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


