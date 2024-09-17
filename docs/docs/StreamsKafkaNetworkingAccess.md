# StreamsKafkaNetworkingAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Reserved. Will be used by PRIVATE_LINK connection type. | [optional] 
**Type** | Pointer to **string** | Selected networking type. Either PUBLIC, VPC or PRIVATE_LINK. Defaults to PUBLIC. For VPC, ensure that VPC peering exists and connectivity has been established between Atlas VPC and the VPC where Kafka cluster is hosted for the connection to function properly. PRIVATE_LINK support is coming soon. | [optional] 

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


