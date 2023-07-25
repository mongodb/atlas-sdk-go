# CreateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique string that identifies the private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] 
**Endpoints** | Pointer to [**[]CreateGCPForwardingRuleRequest**](CreateGCPForwardingRuleRequest.md) | List of individual private endpoints that comprise this endpoint group. | [optional] 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the Google Cloud project in which you created the endpoints. | [optional] 

## Methods

### NewCreateEndpointRequest

`func NewCreateEndpointRequest() *CreateEndpointRequest`

NewCreateEndpointRequest instantiates a new CreateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointRequestWithDefaults

`func NewCreateEndpointRequestWithDefaults() *CreateEndpointRequest`

NewCreateEndpointRequestWithDefaults instantiates a new CreateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEndpointRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEndpointRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEndpointRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEndpointRequest) HasId() bool`

HasId returns a boolean if a field has been set.
### GetPrivateEndpointIPAddress

`func (o *CreateEndpointRequest) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *CreateEndpointRequest) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *CreateEndpointRequest) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *CreateEndpointRequest) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.
### GetEndpointGroupName

`func (o *CreateEndpointRequest) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *CreateEndpointRequest) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *CreateEndpointRequest) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *CreateEndpointRequest) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.
### GetEndpoints

`func (o *CreateEndpointRequest) GetEndpoints() []CreateGCPForwardingRuleRequest`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateEndpointRequest) GetEndpointsOk() (*[]CreateGCPForwardingRuleRequest, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateEndpointRequest) SetEndpoints(v []CreateGCPForwardingRuleRequest)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateEndpointRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.
### GetGcpProjectId

`func (o *CreateEndpointRequest) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *CreateEndpointRequest) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *CreateEndpointRequest) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *CreateEndpointRequest) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


