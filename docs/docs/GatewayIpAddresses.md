# GatewayIpAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | Pointer to [**InboundControlPlaneCloudProviderIPAddresses**](InboundControlPlaneCloudProviderIPAddresses.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundControlPlaneCloudProviderIPAddresses**](OutboundControlPlaneCloudProviderIPAddresses.md) |  | [optional] 

## Methods

### NewGatewayIpAddresses

`func NewGatewayIpAddresses() *GatewayIpAddresses`

NewGatewayIpAddresses instantiates a new GatewayIpAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayIpAddressesWithDefaults

`func NewGatewayIpAddressesWithDefaults() *GatewayIpAddresses`

NewGatewayIpAddressesWithDefaults instantiates a new GatewayIpAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *GatewayIpAddresses) GetInbound() InboundControlPlaneCloudProviderIPAddresses`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *GatewayIpAddresses) GetInboundOk() (*InboundControlPlaneCloudProviderIPAddresses, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *GatewayIpAddresses) SetInbound(v InboundControlPlaneCloudProviderIPAddresses)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *GatewayIpAddresses) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### SetInboundNil

`func (o *GatewayIpAddresses) SetInboundNil()`

SetInboundNil sets Inbound to an explicit JSON null when marshaled, overriding any value previously set with SetInbound. Calling SetInbound again clears the null override.

### GetOutbound

`func (o *GatewayIpAddresses) GetOutbound() OutboundControlPlaneCloudProviderIPAddresses`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *GatewayIpAddresses) GetOutboundOk() (*OutboundControlPlaneCloudProviderIPAddresses, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *GatewayIpAddresses) SetOutbound(v OutboundControlPlaneCloudProviderIPAddresses)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *GatewayIpAddresses) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### SetOutboundNil

`func (o *GatewayIpAddresses) SetOutboundNil()`

SetOutboundNil sets Outbound to an explicit JSON null when marshaled, overriding any value previously set with SetOutbound. Calling SetOutbound again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


