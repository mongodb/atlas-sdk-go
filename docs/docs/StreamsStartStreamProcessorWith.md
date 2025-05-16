# StreamsStartStreamProcessorWith

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**StartAtOperationTime** | Pointer to **time.Time** | The operation time after which the change stream source should begin reporting. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] 

## Methods

### NewStreamsStartStreamProcessorWith

`func NewStreamsStartStreamProcessorWith() *StreamsStartStreamProcessorWith`

NewStreamsStartStreamProcessorWith instantiates a new StreamsStartStreamProcessorWith object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsStartStreamProcessorWithWithDefaults

`func NewStreamsStartStreamProcessorWithWithDefaults() *StreamsStartStreamProcessorWith`

NewStreamsStartStreamProcessorWithWithDefaults instantiates a new StreamsStartStreamProcessorWith object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsStartStreamProcessorWith) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsStartStreamProcessorWith) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsStartStreamProcessorWith) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsStartStreamProcessorWith) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetStartAtOperationTime

`func (o *StreamsStartStreamProcessorWith) GetStartAtOperationTime() time.Time`

GetStartAtOperationTime returns the StartAtOperationTime field if non-nil, zero value otherwise.

### GetStartAtOperationTimeOk

`func (o *StreamsStartStreamProcessorWith) GetStartAtOperationTimeOk() (*time.Time, bool)`

GetStartAtOperationTimeOk returns a tuple with the StartAtOperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAtOperationTime

`func (o *StreamsStartStreamProcessorWith) SetStartAtOperationTime(v time.Time)`

SetStartAtOperationTime sets StartAtOperationTime field to given value.

### HasStartAtOperationTime

`func (o *StreamsStartStreamProcessorWith) HasStartAtOperationTime() bool`

HasStartAtOperationTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


