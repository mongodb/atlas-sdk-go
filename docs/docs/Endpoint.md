# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request. | [optional] [readonly] 
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**InterfaceEndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the interface endpoint. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Human-readable label that MongoDB Cloud generates that identifies the private endpoint connection. | [optional] [readonly] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**PrivateEndpointResourceId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Google Cloud network endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] [readonly] 
**Endpoints** | Pointer to [**[]GCPConsumerForwardingRule**](GCPConsumerForwardingRule.md) | List of individual private endpoints that comprise this endpoint group. | [optional] [readonly] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *Endpoint) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *Endpoint) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *Endpoint) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *Endpoint) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeleteRequested

`func (o *Endpoint) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *Endpoint) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *Endpoint) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *Endpoint) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Endpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Endpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Endpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Endpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInterfaceEndpointId

`func (o *Endpoint) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *Endpoint) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *Endpoint) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *Endpoint) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.

### GetPrivateEndpointConnectionName

`func (o *Endpoint) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *Endpoint) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *Endpoint) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *Endpoint) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.

### GetPrivateEndpointIPAddress

`func (o *Endpoint) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *Endpoint) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *Endpoint) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *Endpoint) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.

### GetPrivateEndpointResourceId

`func (o *Endpoint) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *Endpoint) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *Endpoint) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.

### HasPrivateEndpointResourceId

`func (o *Endpoint) HasPrivateEndpointResourceId() bool`

HasPrivateEndpointResourceId returns a boolean if a field has been set.

### GetStatus

`func (o *Endpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Endpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Endpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Endpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *Endpoint) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *Endpoint) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *Endpoint) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *Endpoint) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *Endpoint) GetEndpoints() []GCPConsumerForwardingRule`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *Endpoint) GetEndpointsOk() (*[]GCPConsumerForwardingRule, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *Endpoint) SetEndpoints(v []GCPConsumerForwardingRule)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *Endpoint) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


