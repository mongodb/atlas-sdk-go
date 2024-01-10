# ControlPlaneIPAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


