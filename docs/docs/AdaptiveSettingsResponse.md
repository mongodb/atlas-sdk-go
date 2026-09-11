# AdaptiveSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptiveSettingsOverrides** | Pointer to [**map[string]any**](interface{}.md) | Map of customer-specified overrides for Adaptive Settings, applied on a best-effort basis. Each supported entry that you specify in this object takes precedence over the corresponding Atlas-managed default. For example, if a setting is enabled by default, you can add an override to disable it for your cluster. | [optional] 
**EffectiveAdaptiveSettings** | [**map[string]any**](interface{}.md) | The effective state of Adaptive Settings currently applied to your cluster, based on your overrides and Atlas-managed defaults. Atlas-managed defaults can vary by MongoDB version, so the same setting may default differently across clusters running different versions. If you set an override for a setting that your cluster&#39;s current MongoDB version doesn&#39;t support, the override doesn&#39;t take effect and the effective value reflects the Atlas-managed default instead. | [readonly] 

## Methods

### NewAdaptiveSettingsResponse

`func NewAdaptiveSettingsResponse(effectiveAdaptiveSettings map[string]any, ) *AdaptiveSettingsResponse`

NewAdaptiveSettingsResponse instantiates a new AdaptiveSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdaptiveSettingsResponseWithDefaults

`func NewAdaptiveSettingsResponseWithDefaults() *AdaptiveSettingsResponse`

NewAdaptiveSettingsResponseWithDefaults instantiates a new AdaptiveSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsResponse) GetAdaptiveSettingsOverrides() map[string]any`

GetAdaptiveSettingsOverrides returns the AdaptiveSettingsOverrides field if non-nil, zero value otherwise.

### GetAdaptiveSettingsOverridesOk

`func (o *AdaptiveSettingsResponse) GetAdaptiveSettingsOverridesOk() (*map[string]any, bool)`

GetAdaptiveSettingsOverridesOk returns a tuple with the AdaptiveSettingsOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsResponse) SetAdaptiveSettingsOverrides(v map[string]any)`

SetAdaptiveSettingsOverrides sets AdaptiveSettingsOverrides field to given value.

### HasAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsResponse) HasAdaptiveSettingsOverrides() bool`

HasAdaptiveSettingsOverrides returns a boolean if a field has been set.

### SetAdaptiveSettingsOverridesNil

`func (o *AdaptiveSettingsResponse) SetAdaptiveSettingsOverridesNil()`

SetAdaptiveSettingsOverridesNil sets AdaptiveSettingsOverrides to an explicit JSON null when marshaled, overriding any value previously set with SetAdaptiveSettingsOverrides. Calling SetAdaptiveSettingsOverrides again clears the null override.

### GetEffectiveAdaptiveSettings

`func (o *AdaptiveSettingsResponse) GetEffectiveAdaptiveSettings() map[string]any`

GetEffectiveAdaptiveSettings returns the EffectiveAdaptiveSettings field if non-nil, zero value otherwise.

### GetEffectiveAdaptiveSettingsOk

`func (o *AdaptiveSettingsResponse) GetEffectiveAdaptiveSettingsOk() (*map[string]any, bool)`

GetEffectiveAdaptiveSettingsOk returns a tuple with the EffectiveAdaptiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAdaptiveSettings

`func (o *AdaptiveSettingsResponse) SetEffectiveAdaptiveSettings(v map[string]any)`

SetEffectiveAdaptiveSettings sets EffectiveAdaptiveSettings field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


