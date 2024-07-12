# StreamsProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the stream processor. | [optional] [readonly] 
**Dlq** | Pointer to [**StreamsDLQ**](StreamsDLQ.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable name of the stream processor. | [optional] 
**Pipeline** | Pointer to **string** | Stream aggregation pipeline you want to apply to your streaming data. | [optional] 

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
### GetDlq

`func (o *StreamsProcessor) GetDlq() StreamsDLQ`

GetDlq returns the Dlq field if non-nil, zero value otherwise.

### GetDlqOk

`func (o *StreamsProcessor) GetDlqOk() (*StreamsDLQ, bool)`

GetDlqOk returns a tuple with the Dlq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlq

`func (o *StreamsProcessor) SetDlq(v StreamsDLQ)`

SetDlq sets Dlq field to given value.

### HasDlq

`func (o *StreamsProcessor) HasDlq() bool`

HasDlq returns a boolean if a field has been set.
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
### GetPipeline

`func (o *StreamsProcessor) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StreamsProcessor) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StreamsProcessor) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *StreamsProcessor) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


