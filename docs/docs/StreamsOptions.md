# StreamsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dlq** | Pointer to [**StreamsDLQ**](StreamsDLQ.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewStreamsOptions

`func NewStreamsOptions() *StreamsOptions`

NewStreamsOptions instantiates a new StreamsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsOptionsWithDefaults

`func NewStreamsOptionsWithDefaults() *StreamsOptions`

NewStreamsOptionsWithDefaults instantiates a new StreamsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlq

`func (o *StreamsOptions) GetDlq() StreamsDLQ`

GetDlq returns the Dlq field if non-nil, zero value otherwise.

### GetDlqOk

`func (o *StreamsOptions) GetDlqOk() (*StreamsDLQ, bool)`

GetDlqOk returns a tuple with the Dlq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlq

`func (o *StreamsOptions) SetDlq(v StreamsDLQ)`

SetDlq sets Dlq field to given value.

### HasDlq

`func (o *StreamsOptions) HasDlq() bool`

HasDlq returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsOptions) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsOptions) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsOptions) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsOptions) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


