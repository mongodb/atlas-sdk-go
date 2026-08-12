# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to [**GatewayIpAddresses**](GatewayIpAddresses.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the service that this gateway represents. | [optional] [readonly] 

## Methods

### NewGateway

`func NewGateway() *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *Gateway) GetIps() GatewayIpAddresses`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Gateway) GetIpsOk() (*GatewayIpAddresses, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Gateway) SetIps(v GatewayIpAddresses)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Gateway) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *Gateway) SetIpsNil()`

SetIpsNil sets Ips to an explicit JSON null when marshaled, overriding any value previously set with SetIps. Calling SetIps again clears the null override.

### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gateway) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Gateway) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


