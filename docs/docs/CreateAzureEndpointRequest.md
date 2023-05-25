# CreateAzureEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique string that identifies the private endpoint&#39;s network interface that someone added to this private endpoint service. | 
**PrivateEndpointIPAddress** | **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | 

## Methods

### NewCreateAzureEndpointRequest

`func NewCreateAzureEndpointRequest(id string, privateEndpointIPAddress string, ) *CreateAzureEndpointRequest`

NewCreateAzureEndpointRequest instantiates a new CreateAzureEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureEndpointRequestWithDefaults

`func NewCreateAzureEndpointRequestWithDefaults() *CreateAzureEndpointRequest`

NewCreateAzureEndpointRequestWithDefaults instantiates a new CreateAzureEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAzureEndpointRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAzureEndpointRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAzureEndpointRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPrivateEndpointIPAddress

`func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *CreateAzureEndpointRequest) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *CreateAzureEndpointRequest) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


