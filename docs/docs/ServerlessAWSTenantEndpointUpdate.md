# ServerlessAWSTenantEndpointUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the private endpoint&#39;s network interface. | [optional] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] 
**ProviderName** | **string** |  | 

## Methods

### NewServerlessAWSTenantEndpointUpdate

`func NewServerlessAWSTenantEndpointUpdate(providerName string, ) *ServerlessAWSTenantEndpointUpdate`

NewServerlessAWSTenantEndpointUpdate instantiates a new ServerlessAWSTenantEndpointUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessAWSTenantEndpointUpdateWithDefaults

`func NewServerlessAWSTenantEndpointUpdateWithDefaults() *ServerlessAWSTenantEndpointUpdate`

NewServerlessAWSTenantEndpointUpdateWithDefaults instantiates a new ServerlessAWSTenantEndpointUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderEndpointId

`func (o *ServerlessAWSTenantEndpointUpdate) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ServerlessAWSTenantEndpointUpdate) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ServerlessAWSTenantEndpointUpdate) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ServerlessAWSTenantEndpointUpdate) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### GetComment

`func (o *ServerlessAWSTenantEndpointUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerlessAWSTenantEndpointUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerlessAWSTenantEndpointUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServerlessAWSTenantEndpointUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetProviderName

`func (o *ServerlessAWSTenantEndpointUpdate) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessAWSTenantEndpointUpdate) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessAWSTenantEndpointUpdate) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


