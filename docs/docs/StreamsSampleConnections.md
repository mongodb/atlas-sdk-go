# StreamsSampleConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Solar** | Pointer to **bool** | Flag that indicates whether to add a &#x60;sample_stream_solar&#x60; connection. | [optional] [default to false]

## Methods

### NewStreamsSampleConnections

`func NewStreamsSampleConnections() *StreamsSampleConnections`

NewStreamsSampleConnections instantiates a new StreamsSampleConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsSampleConnectionsWithDefaults

`func NewStreamsSampleConnectionsWithDefaults() *StreamsSampleConnections`

NewStreamsSampleConnectionsWithDefaults instantiates a new StreamsSampleConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsSampleConnections) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsSampleConnections) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsSampleConnections) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsSampleConnections) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetSolar

`func (o *StreamsSampleConnections) GetSolar() bool`

GetSolar returns the Solar field if non-nil, zero value otherwise.

### GetSolarOk

`func (o *StreamsSampleConnections) GetSolarOk() (*bool, bool)`

GetSolarOk returns a tuple with the Solar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolar

`func (o *StreamsSampleConnections) SetSolar(v bool)`

SetSolar sets Solar field to given value.

### HasSolar

`func (o *StreamsSampleConnections) HasSolar() bool`

HasSolar returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


