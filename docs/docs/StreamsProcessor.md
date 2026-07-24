# StreamsProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the stream processor. | [optional] [readonly] 
**FailoverEnabled** | Pointer to **bool** | Flag that enables or disables failover for the stream processor. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable name of the stream processor. | [optional] 
**Options** | Pointer to [**StreamsOptions**](StreamsOptions.md) |  | [optional] 
**Pipeline** | Pointer to [**[]any**](any.md) | Stream aggregation pipeline you want to apply to your streaming data. | [optional] 
**Tier** | Pointer to **string** | Selected tier for the Stream Workspace. Configures Memory / VCPU allowances. | [optional] 

## Methods

### NewStreamsProcessor

`func NewStreamsProcessor() *StreamsProcessor`

NewStreamsProcessor instantiates a new StreamsProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsProcessorWithDefaults

`func NewStreamsProcessorWithDefaults() *StreamsProcessor`

NewStreamsProcessorWithDefaults instantiates a new StreamsProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsProcessor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsProcessor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsProcessor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsProcessor) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamsProcessor) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetFailoverEnabled

`func (o *StreamsProcessor) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *StreamsProcessor) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *StreamsProcessor) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *StreamsProcessor) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### SetFailoverEnabledNil

`func (o *StreamsProcessor) SetFailoverEnabledNil()`

SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled, overriding any value previously set with SetFailoverEnabled. Calling SetFailoverEnabled again clears the null override.

### GetLinks

`func (o *StreamsProcessor) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsProcessor) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsProcessor) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsProcessor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StreamsProcessor) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetName

`func (o *StreamsProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StreamsProcessor) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetOptions

`func (o *StreamsProcessor) GetOptions() StreamsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StreamsProcessor) GetOptionsOk() (*StreamsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StreamsProcessor) SetOptions(v StreamsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StreamsProcessor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *StreamsProcessor) SetOptionsNil()`

SetOptionsNil sets Options to an explicit JSON null when marshaled, overriding any value previously set with SetOptions. Calling SetOptions again clears the null override.

### GetPipeline

`func (o *StreamsProcessor) GetPipeline() []any`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsProcessor) GetPipelineOk() (*[]any, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsProcessor) SetPipeline(v []any)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *StreamsProcessor) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *StreamsProcessor) SetPipelineNil()`

SetPipelineNil sets Pipeline to an explicit JSON null when marshaled, overriding any value previously set with SetPipeline. Calling SetPipeline again clears the null override.

### GetTier

`func (o *StreamsProcessor) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *StreamsProcessor) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *StreamsProcessor) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *StreamsProcessor) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *StreamsProcessor) SetTierNil()`

SetTierNil sets Tier to an explicit JSON null when marshaled, overriding any value previously set with SetTier. Calling SetTier again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


