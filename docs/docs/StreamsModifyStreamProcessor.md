# StreamsModifyStreamProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | New name for the stream processor. | [optional] 
**Options** | Pointer to [**StreamsModifyStreamProcessorOptions**](StreamsModifyStreamProcessorOptions.md) |  | [optional] 
**Pipeline** | Pointer to [**[]any**](any.md) | New pipeline for the stream processor. | [optional] 

## Methods

### NewStreamsModifyStreamProcessor

`func NewStreamsModifyStreamProcessor() *StreamsModifyStreamProcessor`

NewStreamsModifyStreamProcessor instantiates a new StreamsModifyStreamProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsModifyStreamProcessorWithDefaults

`func NewStreamsModifyStreamProcessorWithDefaults() *StreamsModifyStreamProcessor`

NewStreamsModifyStreamProcessorWithDefaults instantiates a new StreamsModifyStreamProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsModifyStreamProcessor) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsModifyStreamProcessor) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsModifyStreamProcessor) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsModifyStreamProcessor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *StreamsModifyStreamProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsModifyStreamProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsModifyStreamProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsModifyStreamProcessor) HasName() bool`

HasName returns a boolean if a field has been set.
### GetOptions

`func (o *StreamsModifyStreamProcessor) GetOptions() StreamsModifyStreamProcessorOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StreamsModifyStreamProcessor) GetOptionsOk() (*StreamsModifyStreamProcessorOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StreamsModifyStreamProcessor) SetOptions(v StreamsModifyStreamProcessorOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StreamsModifyStreamProcessor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.
### GetPipeline

`func (o *StreamsModifyStreamProcessor) GetPipeline() []any`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsModifyStreamProcessor) GetPipelineOk() (*[]any, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsModifyStreamProcessor) SetPipeline(v []any)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *StreamsModifyStreamProcessor) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


