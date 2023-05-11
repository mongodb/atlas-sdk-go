# ServerlessAzureTenantEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the private endpoint. | [optional] [readonly] 
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Azure private endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Human-readable error message that indicates error condition associated with establishing the private endpoint connection. | [optional] [readonly] 
**PrivateEndpointIpAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Human-readable label that identifies the cloud service provider. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the current operating status of the private endpoint. | [optional] [readonly] 

## Methods

### NewServerlessAzureTenantEndpoint

`func NewServerlessAzureTenantEndpoint() *ServerlessAzureTenantEndpoint`

NewServerlessAzureTenantEndpoint instantiates a new ServerlessAzureTenantEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessAzureTenantEndpointWithDefaults

`func NewServerlessAzureTenantEndpointWithDefaults() *ServerlessAzureTenantEndpoint`

NewServerlessAzureTenantEndpointWithDefaults instantiates a new ServerlessAzureTenantEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerlessAzureTenantEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessAzureTenantEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessAzureTenantEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerlessAzureTenantEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpoint) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ServerlessAzureTenantEndpoint) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpoint) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ServerlessAzureTenantEndpoint) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### GetComment

`func (o *ServerlessAzureTenantEndpoint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerlessAzureTenantEndpoint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerlessAzureTenantEndpoint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServerlessAzureTenantEndpoint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndpointServiceName

`func (o *ServerlessAzureTenantEndpoint) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *ServerlessAzureTenantEndpoint) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *ServerlessAzureTenantEndpoint) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *ServerlessAzureTenantEndpoint) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ServerlessAzureTenantEndpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ServerlessAzureTenantEndpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ServerlessAzureTenantEndpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ServerlessAzureTenantEndpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpoint) GetPrivateEndpointIpAddress() string`

GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIpAddressOk

`func (o *ServerlessAzureTenantEndpoint) GetPrivateEndpointIpAddressOk() (*string, bool)`

GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpoint) SetPrivateEndpointIpAddress(v string)`

SetPrivateEndpointIpAddress sets PrivateEndpointIpAddress field to given value.

### HasPrivateEndpointIpAddress

`func (o *ServerlessAzureTenantEndpoint) HasPrivateEndpointIpAddress() bool`

HasPrivateEndpointIpAddress returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *ServerlessAzureTenantEndpoint) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *ServerlessAzureTenantEndpoint) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *ServerlessAzureTenantEndpoint) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *ServerlessAzureTenantEndpoint) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### GetProviderName

`func (o *ServerlessAzureTenantEndpoint) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessAzureTenantEndpoint) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessAzureTenantEndpoint) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ServerlessAzureTenantEndpoint) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetStatus

`func (o *ServerlessAzureTenantEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerlessAzureTenantEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerlessAzureTenantEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerlessAzureTenantEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


