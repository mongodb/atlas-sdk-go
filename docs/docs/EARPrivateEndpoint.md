# EARPrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider for the Encryption At Rest private endpoint. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message for failures associated with the Encryption At Rest private endpoint. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region in which the Encryption At Rest private endpoint is located. | [optional] 
**Status** | Pointer to **string** | State of the Encryption At Rest private endpoint. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Connection name of the Azure Private Endpoint.  Alternatively: Resource Id of the Aws Private Endpoint. | [optional] [readonly] 

## Methods

### NewEARPrivateEndpoint

`func NewEARPrivateEndpoint() *EARPrivateEndpoint`

NewEARPrivateEndpoint instantiates a new EARPrivateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEARPrivateEndpointWithDefaults

`func NewEARPrivateEndpointWithDefaults() *EARPrivateEndpoint`

NewEARPrivateEndpointWithDefaults instantiates a new EARPrivateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *EARPrivateEndpoint) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EARPrivateEndpoint) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EARPrivateEndpoint) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *EARPrivateEndpoint) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetErrorMessage

`func (o *EARPrivateEndpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EARPrivateEndpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EARPrivateEndpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EARPrivateEndpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetId

`func (o *EARPrivateEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EARPrivateEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EARPrivateEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EARPrivateEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.
### GetRegionName

`func (o *EARPrivateEndpoint) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *EARPrivateEndpoint) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *EARPrivateEndpoint) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *EARPrivateEndpoint) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetStatus

`func (o *EARPrivateEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EARPrivateEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EARPrivateEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EARPrivateEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetPrivateEndpointConnectionName

`func (o *EARPrivateEndpoint) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *EARPrivateEndpoint) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *EARPrivateEndpoint) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *EARPrivateEndpoint) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


