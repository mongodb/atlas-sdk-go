# StreamsProcessorWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal character string that identifies the stream processor. | [readonly] 
**EligibleForFailover** | Pointer to **bool** | Flag that indicates whether the stream processor is eligible for failover. | [optional] [readonly] 
**FailoverEnabled** | Pointer to **bool** | Flag that enables or disables failover for the stream processor. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | **string** | Human-readable name of the stream processor. | [readonly] 
**Options** | Pointer to [**StreamsOptions**](StreamsOptions.md) |  | [optional] 
**Pipeline** | [**[]any**](any.md) | Stream aggregation pipeline you want to apply to your streaming data. | [readonly] 
**State** | **string** | The state of the stream processor. Commonly occurring states are &#39;CREATED&#39;, &#39;STARTED&#39;, &#39;STOPPED&#39; and &#39;FAILED&#39;. | [readonly] 
**Stats** | Pointer to [**any**](interface{}.md) | The stats associated with the stream processor. | [optional] [readonly] 
**Tier** | Pointer to **string** | Selected tier for the Stream Workspace. Configures Memory / VCPU allowances. | [optional] 

## Methods

### NewStreamsProcessorWithStats

`func NewStreamsProcessorWithStats(id string, name string, pipeline []any, state string, ) *StreamsProcessorWithStats`

NewStreamsProcessorWithStats instantiates a new StreamsProcessorWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsProcessorWithStatsWithDefaults

`func NewStreamsProcessorWithStatsWithDefaults() *StreamsProcessorWithStats`

NewStreamsProcessorWithStatsWithDefaults instantiates a new StreamsProcessorWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsProcessorWithStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsProcessorWithStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsProcessorWithStats) SetId(v string)`

SetId sets Id field to given value.

### GetEligibleForFailover

`func (o *StreamsProcessorWithStats) GetEligibleForFailover() bool`

GetEligibleForFailover returns the EligibleForFailover field if non-nil, zero value otherwise.

### GetEligibleForFailoverOk

`func (o *StreamsProcessorWithStats) GetEligibleForFailoverOk() (*bool, bool)`

GetEligibleForFailoverOk returns a tuple with the EligibleForFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleForFailover

`func (o *StreamsProcessorWithStats) SetEligibleForFailover(v bool)`

SetEligibleForFailover sets EligibleForFailover field to given value.

### HasEligibleForFailover

`func (o *StreamsProcessorWithStats) HasEligibleForFailover() bool`

HasEligibleForFailover returns a boolean if a field has been set.

### SetEligibleForFailoverNil

`func (o *StreamsProcessorWithStats) SetEligibleForFailoverNil()`

SetEligibleForFailoverNil sets EligibleForFailover to an explicit JSON null when marshaled, overriding any value previously set with SetEligibleForFailover. Calling SetEligibleForFailover again clears the null override.

### GetFailoverEnabled

`func (o *StreamsProcessorWithStats) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *StreamsProcessorWithStats) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *StreamsProcessorWithStats) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *StreamsProcessorWithStats) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### SetFailoverEnabledNil

`func (o *StreamsProcessorWithStats) SetFailoverEnabledNil()`

SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled, overriding any value previously set with SetFailoverEnabled. Calling SetFailoverEnabled again clears the null override.

### GetLinks

`func (o *StreamsProcessorWithStats) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsProcessorWithStats) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsProcessorWithStats) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsProcessorWithStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StreamsProcessorWithStats) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetName

`func (o *StreamsProcessorWithStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsProcessorWithStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsProcessorWithStats) SetName(v string)`

SetName sets Name field to given value.

### GetOptions

`func (o *StreamsProcessorWithStats) GetOptions() StreamsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StreamsProcessorWithStats) GetOptionsOk() (*StreamsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StreamsProcessorWithStats) SetOptions(v StreamsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StreamsProcessorWithStats) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *StreamsProcessorWithStats) SetOptionsNil()`

SetOptionsNil sets Options to an explicit JSON null when marshaled, overriding any value previously set with SetOptions. Calling SetOptions again clears the null override.

### GetPipeline

`func (o *StreamsProcessorWithStats) GetPipeline() []any`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsProcessorWithStats) GetPipelineOk() (*[]any, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsProcessorWithStats) SetPipeline(v []any)`

SetPipeline sets Pipeline field to given value.

### GetState

`func (o *StreamsProcessorWithStats) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsProcessorWithStats) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsProcessorWithStats) SetState(v string)`

SetState sets State field to given value.

### GetStats

`func (o *StreamsProcessorWithStats) GetStats() any`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StreamsProcessorWithStats) GetStatsOk() (*any, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StreamsProcessorWithStats) SetStats(v any)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StreamsProcessorWithStats) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *StreamsProcessorWithStats) SetStatsNil()`

SetStatsNil sets Stats to an explicit JSON null when marshaled, overriding any value previously set with SetStats. Calling SetStats again clears the null override.

### GetTier

`func (o *StreamsProcessorWithStats) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *StreamsProcessorWithStats) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *StreamsProcessorWithStats) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *StreamsProcessorWithStats) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *StreamsProcessorWithStats) SetTierNil()`

SetTierNil sets Tier to an explicit JSON null when marshaled, overriding any value previously set with SetTier. Calling SetTier again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


