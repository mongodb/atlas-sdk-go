# ServerlessAzureTenantEndpointUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface for this private endpoint service. | [optional] 
**PrivateEndpointIpAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewServerlessAzureTenantEndpointUpdate

`func NewServerlessAzureTenantEndpointUpdate(providerName string, ) *ServerlessAzureTenantEndpointUpdate`

NewServerlessAzureTenantEndpointUpdate instantiates a new ServerlessAzureTenantEndpointUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessAzureTenantEndpointUpdateWithDefaults

`func NewServerlessAzureTenantEndpointUpdateWithDefaults() *ServerlessAzureTenantEndpointUpdate`

NewServerlessAzureTenantEndpointUpdateWithDefaults instantiates a new ServerlessAzureTenantEndpointUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpointUpdate) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ServerlessAzureTenantEndpointUpdate) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpointUpdate) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpointUpdate) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### GetPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpointUpdate) GetPrivateEndpointIpAddress() string`

GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIpAddressOk

`func (o *ServerlessAzureTenantEndpointUpdate) GetPrivateEndpointIpAddressOk() (*string, bool)`

GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpointUpdate) SetPrivateEndpointIpAddress(v string)`

SetPrivateEndpointIpAddress sets PrivateEndpointIpAddress field to given value.

### HasPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpointUpdate) HasPrivateEndpointIpAddress() bool`

HasPrivateEndpointIpAddress returns a boolean if a field has been set.

### GetComment

`func (o *ServerlessAzureTenantEndpointUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerlessAzureTenantEndpointUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerlessAzureTenantEndpointUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServerlessAzureTenantEndpointUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetProviderName

`func (o *ServerlessAzureTenantEndpointUpdate) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessAzureTenantEndpointUpdate) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessAzureTenantEndpointUpdate) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


