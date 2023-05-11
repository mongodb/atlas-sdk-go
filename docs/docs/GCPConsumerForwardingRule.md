# GCPConsumerForwardingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** | Human-readable label that identifies the Google Cloud consumer forwarding rule that you created. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | One Private Internet Protocol version 4 (IPv4) address to which this Google Cloud consumer forwarding rule resolves. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the MongoDB Cloud endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewGCPConsumerForwardingRule

`func NewGCPConsumerForwardingRule() *GCPConsumerForwardingRule`

NewGCPConsumerForwardingRule instantiates a new GCPConsumerForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPConsumerForwardingRuleWithDefaults

`func NewGCPConsumerForwardingRuleWithDefaults() *GCPConsumerForwardingRule`

NewGCPConsumerForwardingRuleWithDefaults instantiates a new GCPConsumerForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *GCPConsumerForwardingRule) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *GCPConsumerForwardingRule) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *GCPConsumerForwardingRule) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *GCPConsumerForwardingRule) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetIpAddress

`func (o *GCPConsumerForwardingRule) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GCPConsumerForwardingRule) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GCPConsumerForwardingRule) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GCPConsumerForwardingRule) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetStatus

`func (o *GCPConsumerForwardingRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GCPConsumerForwardingRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GCPConsumerForwardingRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GCPConsumerForwardingRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


