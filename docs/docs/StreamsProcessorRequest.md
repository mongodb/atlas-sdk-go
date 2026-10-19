# StreamsProcessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverEnabled** | Pointer to **bool** | Flag that enables or disables failover for the stream processor. | [optional] 
**Name** | **string** | Human-readable name of the stream processor. | 
**Options** | Pointer to [**StreamsOptions**](StreamsOptions.md) |  | [optional] 
**Pipeline** | [**[]any**](any.md) | Stream aggregation pipeline you want to apply to your streaming data. | 
**Tier** | Pointer to **string** | Baseline processing tier requested by the user. | [optional] 

## Methods

### NewStreamsProcessorRequest

`func NewStreamsProcessorRequest(name string, pipeline []any, ) *StreamsProcessorRequest`

NewStreamsProcessorRequest instantiates a new StreamsProcessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsProcessorRequestWithDefaults

`func NewStreamsProcessorRequestWithDefaults() *StreamsProcessorRequest`

NewStreamsProcessorRequestWithDefaults instantiates a new StreamsProcessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverEnabled

`func (o *StreamsProcessorRequest) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *StreamsProcessorRequest) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *StreamsProcessorRequest) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *StreamsProcessorRequest) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### SetFailoverEnabledNil

`func (o *StreamsProcessorRequest) SetFailoverEnabledNil()`

SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled, overriding any value previously set with SetFailoverEnabled. Calling SetFailoverEnabled again clears the null override.

### GetName

`func (o *StreamsProcessorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsProcessorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsProcessorRequest) SetName(v string)`

SetName sets Name field to given value.

### GetOptions

`func (o *StreamsProcessorRequest) GetOptions() StreamsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StreamsProcessorRequest) GetOptionsOk() (*StreamsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StreamsProcessorRequest) SetOptions(v StreamsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StreamsProcessorRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *StreamsProcessorRequest) SetOptionsNil()`

SetOptionsNil sets Options to an explicit JSON null when marshaled, overriding any value previously set with SetOptions. Calling SetOptions again clears the null override.

### GetPipeline

`func (o *StreamsProcessorRequest) GetPipeline() []any`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsProcessorRequest) GetPipelineOk() (*[]any, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsProcessorRequest) SetPipeline(v []any)`

SetPipeline sets Pipeline field to given value.

### GetTier

`func (o *StreamsProcessorRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *StreamsProcessorRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *StreamsProcessorRequest) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *StreamsProcessorRequest) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *StreamsProcessorRequest) SetTierNil()`

SetTierNil sets Tier to an explicit JSON null when marshaled, overriding any value previously set with SetTier. Calling SetTier again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


