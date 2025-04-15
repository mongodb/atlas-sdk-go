# StreamsKafkaNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**StreamsKafkaNetworkingAccess**](StreamsKafkaNetworkingAccess.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewStreamsKafkaNetworking

`func NewStreamsKafkaNetworking() *StreamsKafkaNetworking`

NewStreamsKafkaNetworking instantiates a new StreamsKafkaNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsKafkaNetworkingWithDefaults

`func NewStreamsKafkaNetworkingWithDefaults() *StreamsKafkaNetworking`

NewStreamsKafkaNetworkingWithDefaults instantiates a new StreamsKafkaNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *StreamsKafkaNetworking) GetAccess() StreamsKafkaNetworkingAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *StreamsKafkaNetworking) GetAccessOk() (*StreamsKafkaNetworkingAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *StreamsKafkaNetworking) SetAccess(v StreamsKafkaNetworkingAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *StreamsKafkaNetworking) HasAccess() bool`

HasAccess returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsKafkaNetworking) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsKafkaNetworking) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsKafkaNetworking) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsKafkaNetworking) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


