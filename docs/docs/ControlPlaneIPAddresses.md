# ControlPlaneIPAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]Gateway**](Gateway.md) | List of gateways, each representing a group of service-specific IP addresses that customers can add to their allow lists independently. Includes the Atlas Gateway (data plane) group when present. | [optional] [readonly] 
**Inbound** | Pointer to [**InboundControlPlaneCloudProviderIPAddresses**](InboundControlPlaneCloudProviderIPAddresses.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundControlPlaneCloudProviderIPAddresses**](OutboundControlPlaneCloudProviderIPAddresses.md) |  | [optional] 

## Methods

### NewControlPlaneIPAddresses

`func NewControlPlaneIPAddresses() *ControlPlaneIPAddresses`

NewControlPlaneIPAddresses instantiates a new ControlPlaneIPAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlPlaneIPAddressesWithDefaults

`func NewControlPlaneIPAddressesWithDefaults() *ControlPlaneIPAddresses`

NewControlPlaneIPAddressesWithDefaults instantiates a new ControlPlaneIPAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *ControlPlaneIPAddresses) GetGateways() []Gateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ControlPlaneIPAddresses) GetGatewaysOk() (*[]Gateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ControlPlaneIPAddresses) SetGateways(v []Gateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ControlPlaneIPAddresses) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### SetGatewaysNil

`func (o *ControlPlaneIPAddresses) SetGatewaysNil()`

SetGatewaysNil sets Gateways to an explicit JSON null when marshaled, overriding any value previously set with SetGateways. Calling SetGateways again clears the null override.

### GetInbound

`func (o *ControlPlaneIPAddresses) GetInbound() InboundControlPlaneCloudProviderIPAddresses`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *ControlPlaneIPAddresses) GetInboundOk() (*InboundControlPlaneCloudProviderIPAddresses, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *ControlPlaneIPAddresses) SetInbound(v InboundControlPlaneCloudProviderIPAddresses)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *ControlPlaneIPAddresses) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### SetInboundNil

`func (o *ControlPlaneIPAddresses) SetInboundNil()`

SetInboundNil sets Inbound to an explicit JSON null when marshaled, overriding any value previously set with SetInbound. Calling SetInbound again clears the null override.

### GetOutbound

`func (o *ControlPlaneIPAddresses) GetOutbound() OutboundControlPlaneCloudProviderIPAddresses`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *ControlPlaneIPAddresses) GetOutboundOk() (*OutboundControlPlaneCloudProviderIPAddresses, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *ControlPlaneIPAddresses) SetOutbound(v OutboundControlPlaneCloudProviderIPAddresses)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *ControlPlaneIPAddresses) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### SetOutboundNil

`func (o *ControlPlaneIPAddresses) SetOutboundNil()`

SetOutboundNil sets Outbound to an explicit JSON null when marshaled, overriding any value previously set with SetOutbound. Calling SetOutbound again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


