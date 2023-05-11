# NetPeerRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ProviderName** | **string** | Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection. | 

## Methods

### NewNetPeerRequestBase

`func NewNetPeerRequestBase(containerId string, providerName string, ) *NetPeerRequestBase`

NewNetPeerRequestBase instantiates a new NetPeerRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetPeerRequestBaseWithDefaults

`func NewNetPeerRequestBaseWithDefaults() *NetPeerRequestBase`

NewNetPeerRequestBaseWithDefaults instantiates a new NetPeerRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *NetPeerRequestBase) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *NetPeerRequestBase) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *NetPeerRequestBase) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetProviderName

`func (o *NetPeerRequestBase) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *NetPeerRequestBase) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *NetPeerRequestBase) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


