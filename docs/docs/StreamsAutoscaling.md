# StreamsAutoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether autoscaling is enabled.  - **Omitted, &#x60;null&#x60;, or &#x60;false&#x60;:**   - On &#x60;CREATE&#x60;: a no-op, there is no persisted setting yet to disable or clear.   - On &#x60;MODIFY&#x60; or &#x60;:startWith&#x60;: omitted preserves the current setting. &#x60;null&#x60; or &#x60;false&#x60; disables autoscaling and clears its configuration. - **&#x60;true&#x60;** on &#x60;CREATE&#x60;, &#x60;MODIFY&#x60;, or &#x60;:startWith&#x60;: enables autoscaling. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MaxTier** | Pointer to **string** | Tier ceiling for autoscaling (scale-up limit).  - **Omitted:**   - On &#x60;CREATE&#x60;: falls back to the workspace max tier (there is no current bound to preserve).   - On &#x60;MODIFY&#x60; or &#x60;:startWith&#x60;: the current bound is preserved. - **&#x60;null&#x60;** on &#x60;CREATE&#x60;, &#x60;MODIFY&#x60;, or &#x60;:startWith&#x60;: resets the bound to the workspace max tier. - **A tier value** on &#x60;CREATE&#x60;, &#x60;MODIFY&#x60;, or &#x60;:startWith&#x60;: sets the bound to that tier. | [optional] 
**MinTier** | Pointer to **string** | Tier floor for autoscaling (scale-down limit).  - **Omitted:**   - On &#x60;CREATE&#x60;: falls back to the workspace default tier (there is no current bound to preserve).   - On &#x60;MODIFY&#x60; or &#x60;:startWith&#x60;: the current bound is preserved. - **&#x60;null&#x60;** on &#x60;CREATE&#x60;, &#x60;MODIFY&#x60;, or &#x60;:startWith&#x60;: resets the bound to the workspace default tier. - **A tier value** on &#x60;CREATE&#x60;, &#x60;MODIFY&#x60;, or &#x60;:startWith&#x60;: sets the bound to that tier. | [optional] 

## Methods

### NewStreamsAutoscaling

`func NewStreamsAutoscaling() *StreamsAutoscaling`

NewStreamsAutoscaling instantiates a new StreamsAutoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsAutoscalingWithDefaults

`func NewStreamsAutoscalingWithDefaults() *StreamsAutoscaling`

NewStreamsAutoscalingWithDefaults instantiates a new StreamsAutoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StreamsAutoscaling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamsAutoscaling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamsAutoscaling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamsAutoscaling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *StreamsAutoscaling) SetEnabledNil()`

SetEnabledNil sets Enabled to an explicit JSON null when marshaled, overriding any value previously set with SetEnabled. Calling SetEnabled again clears the null override.

### GetLinks

`func (o *StreamsAutoscaling) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsAutoscaling) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsAutoscaling) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsAutoscaling) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StreamsAutoscaling) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetMaxTier

`func (o *StreamsAutoscaling) GetMaxTier() string`

GetMaxTier returns the MaxTier field if non-nil, zero value otherwise.

### GetMaxTierOk

`func (o *StreamsAutoscaling) GetMaxTierOk() (*string, bool)`

GetMaxTierOk returns a tuple with the MaxTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTier

`func (o *StreamsAutoscaling) SetMaxTier(v string)`

SetMaxTier sets MaxTier field to given value.

### HasMaxTier

`func (o *StreamsAutoscaling) HasMaxTier() bool`

HasMaxTier returns a boolean if a field has been set.

### SetMaxTierNil

`func (o *StreamsAutoscaling) SetMaxTierNil()`

SetMaxTierNil sets MaxTier to an explicit JSON null when marshaled, overriding any value previously set with SetMaxTier. Calling SetMaxTier again clears the null override.

### GetMinTier

`func (o *StreamsAutoscaling) GetMinTier() string`

GetMinTier returns the MinTier field if non-nil, zero value otherwise.

### GetMinTierOk

`func (o *StreamsAutoscaling) GetMinTierOk() (*string, bool)`

GetMinTierOk returns a tuple with the MinTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTier

`func (o *StreamsAutoscaling) SetMinTier(v string)`

SetMinTier sets MinTier field to given value.

### HasMinTier

`func (o *StreamsAutoscaling) HasMinTier() bool`

HasMinTier returns a boolean if a field has been set.

### SetMinTierNil

`func (o *StreamsAutoscaling) SetMinTierNil()`

SetMinTierNil sets MinTier to an explicit JSON null when marshaled, overriding any value previously set with SetMinTier. Calling SetMinTier again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


