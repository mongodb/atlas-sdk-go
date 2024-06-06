# PrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Cloud service provider that serves the requested endpoint. | [readonly] 
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request. | [optional] [readonly] 
**InterfaceEndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the interface endpoint. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Human-readable label that MongoDB Cloud generates that identifies the private endpoint connection. | [optional] [readonly] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**PrivateEndpointResourceId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Google Cloud network endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] [readonly] 
**Endpoints** | Pointer to [**[]GCPConsumerForwardingRule**](GCPConsumerForwardingRule.md) | List of individual private endpoints that comprise this endpoint group. | [optional] [readonly] 

## Methods

### NewPrivateLinkEndpoint

`func NewPrivateLinkEndpoint(cloudProvider string, ) *PrivateLinkEndpoint`

NewPrivateLinkEndpoint instantiates a new PrivateLinkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateLinkEndpointWithDefaults

`func NewPrivateLinkEndpointWithDefaults() *PrivateLinkEndpoint`

NewPrivateLinkEndpointWithDefaults instantiates a new PrivateLinkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *PrivateLinkEndpoint) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *PrivateLinkEndpoint) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *PrivateLinkEndpoint) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetDeleteRequested

`func (o *PrivateLinkEndpoint) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *PrivateLinkEndpoint) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *PrivateLinkEndpoint) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *PrivateLinkEndpoint) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.
### GetErrorMessage

`func (o *PrivateLinkEndpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PrivateLinkEndpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PrivateLinkEndpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PrivateLinkEndpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetConnectionStatus

`func (o *PrivateLinkEndpoint) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *PrivateLinkEndpoint) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *PrivateLinkEndpoint) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *PrivateLinkEndpoint) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.
### GetInterfaceEndpointId

`func (o *PrivateLinkEndpoint) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *PrivateLinkEndpoint) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *PrivateLinkEndpoint) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *PrivateLinkEndpoint) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.
### GetPrivateEndpointConnectionName

`func (o *PrivateLinkEndpoint) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *PrivateLinkEndpoint) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *PrivateLinkEndpoint) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *PrivateLinkEndpoint) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.
### GetPrivateEndpointIPAddress

`func (o *PrivateLinkEndpoint) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *PrivateLinkEndpoint) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *PrivateLinkEndpoint) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *PrivateLinkEndpoint) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.
### GetPrivateEndpointResourceId

`func (o *PrivateLinkEndpoint) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *PrivateLinkEndpoint) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *PrivateLinkEndpoint) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.

### HasPrivateEndpointResourceId

`func (o *PrivateLinkEndpoint) HasPrivateEndpointResourceId() bool`

HasPrivateEndpointResourceId returns a boolean if a field has been set.
### GetStatus

`func (o *PrivateLinkEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateLinkEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateLinkEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateLinkEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetEndpointGroupName

`func (o *PrivateLinkEndpoint) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *PrivateLinkEndpoint) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *PrivateLinkEndpoint) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *PrivateLinkEndpoint) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.
### GetEndpoints

`func (o *PrivateLinkEndpoint) GetEndpoints() []GCPConsumerForwardingRule`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PrivateLinkEndpoint) GetEndpointsOk() (*[]GCPConsumerForwardingRule, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PrivateLinkEndpoint) SetEndpoints(v []GCPConsumerForwardingRule)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PrivateLinkEndpoint) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


