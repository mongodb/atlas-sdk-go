# CreateGCPForwardingRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** | Human-readable label that identifies the Google Cloud consumer forwarding rule that you created. | [optional] 
**IpAddress** | Pointer to **string** | One Private Internet Protocol version 4 (IPv4) address to which this Google Cloud consumer forwarding rule resolves. | [optional] 

## Methods

### NewCreateGCPForwardingRuleRequest

`func NewCreateGCPForwardingRuleRequest() *CreateGCPForwardingRuleRequest`

NewCreateGCPForwardingRuleRequest instantiates a new CreateGCPForwardingRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGCPForwardingRuleRequestWithDefaults

`func NewCreateGCPForwardingRuleRequestWithDefaults() *CreateGCPForwardingRuleRequest`

NewCreateGCPForwardingRuleRequestWithDefaults instantiates a new CreateGCPForwardingRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *CreateGCPForwardingRuleRequest) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *CreateGCPForwardingRuleRequest) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *CreateGCPForwardingRuleRequest) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *CreateGCPForwardingRuleRequest) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.
### GetIpAddress

`func (o *CreateGCPForwardingRuleRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateGCPForwardingRuleRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateGCPForwardingRuleRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateGCPForwardingRuleRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


