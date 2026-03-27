# StreamsPublicPrivateLinkNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**StreamsPublicPrivateLinkNetworkingAccess**](StreamsPublicPrivateLinkNetworkingAccess.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewStreamsPublicPrivateLinkNetworking

`func NewStreamsPublicPrivateLinkNetworking() *StreamsPublicPrivateLinkNetworking`

NewStreamsPublicPrivateLinkNetworking instantiates a new StreamsPublicPrivateLinkNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPublicPrivateLinkNetworkingWithDefaults

`func NewStreamsPublicPrivateLinkNetworkingWithDefaults() *StreamsPublicPrivateLinkNetworking`

NewStreamsPublicPrivateLinkNetworkingWithDefaults instantiates a new StreamsPublicPrivateLinkNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *StreamsPublicPrivateLinkNetworking) GetAccess() StreamsPublicPrivateLinkNetworkingAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *StreamsPublicPrivateLinkNetworking) GetAccessOk() (*StreamsPublicPrivateLinkNetworkingAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *StreamsPublicPrivateLinkNetworking) SetAccess(v StreamsPublicPrivateLinkNetworkingAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *StreamsPublicPrivateLinkNetworking) HasAccess() bool`

HasAccess returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsPublicPrivateLinkNetworking) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsPublicPrivateLinkNetworking) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsPublicPrivateLinkNetworking) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsPublicPrivateLinkNetworking) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


