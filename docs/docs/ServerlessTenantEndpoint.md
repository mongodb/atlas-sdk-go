# ServerlessTenantEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the private endpoint. | [optional] [readonly] 
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Azure private endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Human-readable error message that indicates error condition associated with establishing the private endpoint connection. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the current operating status of the private endpoint. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Human-readable label that identifies the cloud service provider. | [optional] [readonly] 
**PrivateEndpointIpAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 

## Methods

### NewServerlessTenantEndpoint

`func NewServerlessTenantEndpoint() *ServerlessTenantEndpoint`

NewServerlessTenantEndpoint instantiates a new ServerlessTenantEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessTenantEndpointWithDefaults

`func NewServerlessTenantEndpointWithDefaults() *ServerlessTenantEndpoint`

NewServerlessTenantEndpointWithDefaults instantiates a new ServerlessTenantEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerlessTenantEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessTenantEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessTenantEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerlessTenantEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServerlessTenantEndpoint) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetCloudProviderEndpointId

`func (o *ServerlessTenantEndpoint) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ServerlessTenantEndpoint) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ServerlessTenantEndpoint) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ServerlessTenantEndpoint) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### SetCloudProviderEndpointIdNil

`func (o *ServerlessTenantEndpoint) SetCloudProviderEndpointIdNil()`

SetCloudProviderEndpointIdNil sets CloudProviderEndpointId to an explicit JSON null when marshaled, overriding any value previously set with SetCloudProviderEndpointId. Calling SetCloudProviderEndpointId again clears the null override.

### GetComment

`func (o *ServerlessTenantEndpoint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerlessTenantEndpoint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerlessTenantEndpoint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServerlessTenantEndpoint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ServerlessTenantEndpoint) SetCommentNil()`

SetCommentNil sets Comment to an explicit JSON null when marshaled, overriding any value previously set with SetComment. Calling SetComment again clears the null override.

### GetEndpointServiceName

`func (o *ServerlessTenantEndpoint) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *ServerlessTenantEndpoint) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *ServerlessTenantEndpoint) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *ServerlessTenantEndpoint) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### SetEndpointServiceNameNil

`func (o *ServerlessTenantEndpoint) SetEndpointServiceNameNil()`

SetEndpointServiceNameNil sets EndpointServiceName to an explicit JSON null when marshaled, overriding any value previously set with SetEndpointServiceName. Calling SetEndpointServiceName again clears the null override.

### GetErrorMessage

`func (o *ServerlessTenantEndpoint) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ServerlessTenantEndpoint) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ServerlessTenantEndpoint) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ServerlessTenantEndpoint) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ServerlessTenantEndpoint) SetErrorMessageNil()`

SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled, overriding any value previously set with SetErrorMessage. Calling SetErrorMessage again clears the null override.

### GetStatus

`func (o *ServerlessTenantEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerlessTenantEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerlessTenantEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerlessTenantEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServerlessTenantEndpoint) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.

### GetProviderName

`func (o *ServerlessTenantEndpoint) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessTenantEndpoint) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessTenantEndpoint) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ServerlessTenantEndpoint) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ServerlessTenantEndpoint) SetProviderNameNil()`

SetProviderNameNil sets ProviderName to an explicit JSON null when marshaled, overriding any value previously set with SetProviderName. Calling SetProviderName again clears the null override.

### GetPrivateEndpointIpAddress

`func (o *ServerlessTenantEndpoint) GetPrivateEndpointIpAddress() string`

GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIpAddressOk

`func (o *ServerlessTenantEndpoint) GetPrivateEndpointIpAddressOk() (*string, bool)`

GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIpAddress

`func (o *ServerlessTenantEndpoint) SetPrivateEndpointIpAddress(v string)`

SetPrivateEndpointIpAddress sets PrivateEndpointIpAddress field to given value.

### HasPrivateEndpointIpAddress

`func (o *ServerlessTenantEndpoint) HasPrivateEndpointIpAddress() bool`

HasPrivateEndpointIpAddress returns a boolean if a field has been set.

### SetPrivateEndpointIpAddressNil

`func (o *ServerlessTenantEndpoint) SetPrivateEndpointIpAddressNil()`

SetPrivateEndpointIpAddressNil sets PrivateEndpointIpAddress to an explicit JSON null when marshaled, overriding any value previously set with SetPrivateEndpointIpAddress. Calling SetPrivateEndpointIpAddress again clears the null override.

### GetPrivateLinkServiceResourceId

`func (o *ServerlessTenantEndpoint) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *ServerlessTenantEndpoint) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *ServerlessTenantEndpoint) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *ServerlessTenantEndpoint) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### SetPrivateLinkServiceResourceIdNil

`func (o *ServerlessTenantEndpoint) SetPrivateLinkServiceResourceIdNil()`

SetPrivateLinkServiceResourceIdNil sets PrivateLinkServiceResourceId to an explicit JSON null when marshaled, overriding any value previously set with SetPrivateLinkServiceResourceId. Calling SetPrivateLinkServiceResourceId again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


