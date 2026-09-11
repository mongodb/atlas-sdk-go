# AdaptiveSettingsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptiveSettingsOverrides** | Pointer to [**map[string]any**](interface{}.md) | Map of customer-specified overrides for Adaptive Settings, applied on a best-effort basis. Each supported entry that you specify in this object takes precedence over the corresponding Atlas-managed default. For example, if a setting is enabled by default, you can add an override to disable it for your cluster. | [optional] 

## Methods

### NewAdaptiveSettingsUpdateRequest

`func NewAdaptiveSettingsUpdateRequest() *AdaptiveSettingsUpdateRequest`

NewAdaptiveSettingsUpdateRequest instantiates a new AdaptiveSettingsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdaptiveSettingsUpdateRequestWithDefaults

`func NewAdaptiveSettingsUpdateRequestWithDefaults() *AdaptiveSettingsUpdateRequest`

NewAdaptiveSettingsUpdateRequestWithDefaults instantiates a new AdaptiveSettingsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsUpdateRequest) GetAdaptiveSettingsOverrides() map[string]any`

GetAdaptiveSettingsOverrides returns the AdaptiveSettingsOverrides field if non-nil, zero value otherwise.

### GetAdaptiveSettingsOverridesOk

`func (o *AdaptiveSettingsUpdateRequest) GetAdaptiveSettingsOverridesOk() (*map[string]any, bool)`

GetAdaptiveSettingsOverridesOk returns a tuple with the AdaptiveSettingsOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsUpdateRequest) SetAdaptiveSettingsOverrides(v map[string]any)`

SetAdaptiveSettingsOverrides sets AdaptiveSettingsOverrides field to given value.

### HasAdaptiveSettingsOverrides

`func (o *AdaptiveSettingsUpdateRequest) HasAdaptiveSettingsOverrides() bool`

HasAdaptiveSettingsOverrides returns a boolean if a field has been set.

### SetAdaptiveSettingsOverridesNil

`func (o *AdaptiveSettingsUpdateRequest) SetAdaptiveSettingsOverridesNil()`

SetAdaptiveSettingsOverridesNil sets AdaptiveSettingsOverrides to an explicit JSON null when marshaled, overriding any value previously set with SetAdaptiveSettingsOverrides. Calling SetAdaptiveSettingsOverrides again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


