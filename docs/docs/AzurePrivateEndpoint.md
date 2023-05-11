# AzurePrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Human-readable label that MongoDB Cloud generates that identifies the private endpoint connection. | [optional] [readonly] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**PrivateEndpointResourceId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Azure Private Link Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewAzurePrivateEndpoint

`func NewAzurePrivateEndpoint() *AzurePrivateEndpoint`

NewAzurePrivateEndpoint instantiates a new AzurePrivateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePrivateEndpointWithDefaults

`func NewAzurePrivateEndpointWithDefaults() *AzurePrivateEndpoint`

NewAzurePrivateEndpointWithDefaults instantiates a new AzurePrivateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteRequested

`func (o *AzurePrivateEndpoint) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *AzurePrivateEndpoint) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *AzurePrivateEndpoint) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *AzurePrivateEndpoint) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AzurePrivateEndpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AzurePrivateEndpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AzurePrivateEndpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AzurePrivateEndpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPrivateEndpointConnectionName

`func (o *AzurePrivateEndpoint) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *AzurePrivateEndpoint) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *AzurePrivateEndpoint) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *AzurePrivateEndpoint) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.

### GetPrivateEndpointIPAddress

`func (o *AzurePrivateEndpoint) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *AzurePrivateEndpoint) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *AzurePrivateEndpoint) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *AzurePrivateEndpoint) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.

### GetPrivateEndpointResourceId

`func (o *AzurePrivateEndpoint) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *AzurePrivateEndpoint) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *AzurePrivateEndpoint) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.

### HasPrivateEndpointResourceId

`func (o *AzurePrivateEndpoint) HasPrivateEndpointResourceId() bool`

HasPrivateEndpointResourceId returns a boolean if a field has been set.

### GetStatus

`func (o *AzurePrivateEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AzurePrivateEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AzurePrivateEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AzurePrivateEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


