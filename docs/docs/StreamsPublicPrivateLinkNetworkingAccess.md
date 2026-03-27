# StreamsPublicPrivateLinkNetworkingAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | The ID of the Private Link connection. Required for &#x60;PRIVATE_LINK&#x60; type. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Type** | Pointer to **string** | Selected networking type. Either &#x60;PUBLIC&#x60; or &#x60;PRIVATE_LINK&#x60;. Defaults to &#x60;PUBLIC&#x60;. | [optional] 

## Methods

### NewStreamsPublicPrivateLinkNetworkingAccess

`func NewStreamsPublicPrivateLinkNetworkingAccess() *StreamsPublicPrivateLinkNetworkingAccess`

NewStreamsPublicPrivateLinkNetworkingAccess instantiates a new StreamsPublicPrivateLinkNetworkingAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPublicPrivateLinkNetworkingAccessWithDefaults

`func NewStreamsPublicPrivateLinkNetworkingAccessWithDefaults() *StreamsPublicPrivateLinkNetworkingAccess`

NewStreamsPublicPrivateLinkNetworkingAccessWithDefaults instantiates a new StreamsPublicPrivateLinkNetworkingAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *StreamsPublicPrivateLinkNetworkingAccess) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *StreamsPublicPrivateLinkNetworkingAccess) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsPublicPrivateLinkNetworkingAccess) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsPublicPrivateLinkNetworkingAccess) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetType

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsPublicPrivateLinkNetworkingAccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsPublicPrivateLinkNetworkingAccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamsPublicPrivateLinkNetworkingAccess) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


