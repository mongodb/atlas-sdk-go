# StreamsModifyStreamProcessorOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dlq** | Pointer to [**StreamsDLQ**](StreamsDLQ.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ResumeFromCheckpoint** | Pointer to **bool** | When true, the modified stream processor resumes from its last checkpoint. | [optional] 

## Methods

### NewStreamsModifyStreamProcessorOptions

`func NewStreamsModifyStreamProcessorOptions() *StreamsModifyStreamProcessorOptions`

NewStreamsModifyStreamProcessorOptions instantiates a new StreamsModifyStreamProcessorOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsModifyStreamProcessorOptionsWithDefaults

`func NewStreamsModifyStreamProcessorOptionsWithDefaults() *StreamsModifyStreamProcessorOptions`

NewStreamsModifyStreamProcessorOptionsWithDefaults instantiates a new StreamsModifyStreamProcessorOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlq

`func (o *StreamsModifyStreamProcessorOptions) GetDlq() StreamsDLQ`

GetDlq returns the Dlq field if non-nil, zero value otherwise.

### GetDlqOk

`func (o *StreamsModifyStreamProcessorOptions) GetDlqOk() (*StreamsDLQ, bool)`

GetDlqOk returns a tuple with the Dlq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlq

`func (o *StreamsModifyStreamProcessorOptions) SetDlq(v StreamsDLQ)`

SetDlq sets Dlq field to given value.

### HasDlq

`func (o *StreamsModifyStreamProcessorOptions) HasDlq() bool`

HasDlq returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsModifyStreamProcessorOptions) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsModifyStreamProcessorOptions) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsModifyStreamProcessorOptions) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsModifyStreamProcessorOptions) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResumeFromCheckpoint

`func (o *StreamsModifyStreamProcessorOptions) GetResumeFromCheckpoint() bool`

GetResumeFromCheckpoint returns the ResumeFromCheckpoint field if non-nil, zero value otherwise.

### GetResumeFromCheckpointOk

`func (o *StreamsModifyStreamProcessorOptions) GetResumeFromCheckpointOk() (*bool, bool)`

GetResumeFromCheckpointOk returns a tuple with the ResumeFromCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeFromCheckpoint

`func (o *StreamsModifyStreamProcessorOptions) SetResumeFromCheckpoint(v bool)`

SetResumeFromCheckpoint sets ResumeFromCheckpoint field to given value.

### HasResumeFromCheckpoint

`func (o *StreamsModifyStreamProcessorOptions) HasResumeFromCheckpoint() bool`

HasResumeFromCheckpoint returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


