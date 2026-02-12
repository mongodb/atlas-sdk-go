# StreamsKafkaNetworkingAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | Reserved. Will be used by &#x60;PRIVATE_LINK&#x60; connection type. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Reserved. Will be used by &#x60;PRIVATE_LINK&#x60; connection type. | [optional] 
**TgwRouteId** | Pointer to **string** | Reserved. Will be used by &#x60;TRANSIT_GATEWAY&#x60; connection type. | [optional] 
**Type** | Pointer to **string** | Selected networking type. Either &#x60;PUBLIC&#x60;, &#x60;VPC&#x60;, &#x60;PRIVATE_LINK&#x60;, or &#x60;TRANSIT_GATEWAY&#x60;. Defaults to &#x60;PUBLIC&#x60;. For VPC, ensure that VPC peering exists and connectivity has been established between Atlas VPC and the VPC where Kafka cluster is hosted for the connection to function properly. &#x60;TRANSIT_GATEWAY&#x60; support is coming soon. | [optional] 

## Methods

### NewStreamsKafkaNetworkingAccess

`func NewStreamsKafkaNetworkingAccess() *StreamsKafkaNetworkingAccess`

NewStreamsKafkaNetworkingAccess instantiates a new StreamsKafkaNetworkingAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsKafkaNetworkingAccessWithDefaults

`func NewStreamsKafkaNetworkingAccessWithDefaults() *StreamsKafkaNetworkingAccess`

NewStreamsKafkaNetworkingAccessWithDefaults instantiates a new StreamsKafkaNetworkingAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *StreamsKafkaNetworkingAccess) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *StreamsKafkaNetworkingAccess) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *StreamsKafkaNetworkingAccess) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *StreamsKafkaNetworkingAccess) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsKafkaNetworkingAccess) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsKafkaNetworkingAccess) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsKafkaNetworkingAccess) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsKafkaNetworkingAccess) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetName

`func (o *StreamsKafkaNetworkingAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsKafkaNetworkingAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsKafkaNetworkingAccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsKafkaNetworkingAccess) HasName() bool`

HasName returns a boolean if a field has been set.
### GetTgwRouteId

`func (o *StreamsKafkaNetworkingAccess) GetTgwRouteId() string`

GetTgwRouteId returns the TgwRouteId field if non-nil, zero value otherwise.

### GetTgwRouteIdOk

`func (o *StreamsKafkaNetworkingAccess) GetTgwRouteIdOk() (*string, bool)`

GetTgwRouteIdOk returns a tuple with the TgwRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteId

`func (o *StreamsKafkaNetworkingAccess) SetTgwRouteId(v string)`

SetTgwRouteId sets TgwRouteId field to given value.

### HasTgwRouteId

`func (o *StreamsKafkaNetworkingAccess) HasTgwRouteId() bool`

HasTgwRouteId returns a boolean if a field has been set.
### GetType

`func (o *StreamsKafkaNetworkingAccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsKafkaNetworkingAccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsKafkaNetworkingAccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamsKafkaNetworkingAccess) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


