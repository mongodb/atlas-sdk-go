# AzurePrivateLinkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Cloud service provider that serves the requested endpoint service. | [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**PrivateEndpoints** | Pointer to **[]string** | List of private endpoints assigned to this Azure Private Link Service. | [optional] [readonly] 
**PrivateLinkServiceName** | Pointer to **string** | Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewAzurePrivateLinkConnection

`func NewAzurePrivateLinkConnection(cloudProvider string, ) *AzurePrivateLinkConnection`

NewAzurePrivateLinkConnection instantiates a new AzurePrivateLinkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePrivateLinkConnectionWithDefaults

`func NewAzurePrivateLinkConnectionWithDefaults() *AzurePrivateLinkConnection`

NewAzurePrivateLinkConnectionWithDefaults instantiates a new AzurePrivateLinkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AzurePrivateLinkConnection) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AzurePrivateLinkConnection) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AzurePrivateLinkConnection) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetErrorMessage

`func (o *AzurePrivateLinkConnection) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AzurePrivateLinkConnection) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AzurePrivateLinkConnection) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AzurePrivateLinkConnection) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *AzurePrivateLinkConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzurePrivateLinkConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzurePrivateLinkConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzurePrivateLinkConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivateEndpoints

`func (o *AzurePrivateLinkConnection) GetPrivateEndpoints() []string`

GetPrivateEndpoints returns the PrivateEndpoints field if non-nil, zero value otherwise.

### GetPrivateEndpointsOk

`func (o *AzurePrivateLinkConnection) GetPrivateEndpointsOk() (*[]string, bool)`

GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoints

`func (o *AzurePrivateLinkConnection) SetPrivateEndpoints(v []string)`

SetPrivateEndpoints sets PrivateEndpoints field to given value.

### HasPrivateEndpoints

`func (o *AzurePrivateLinkConnection) HasPrivateEndpoints() bool`

HasPrivateEndpoints returns a boolean if a field has been set.

### GetPrivateLinkServiceName

`func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceName() string`

GetPrivateLinkServiceName returns the PrivateLinkServiceName field if non-nil, zero value otherwise.

### GetPrivateLinkServiceNameOk

`func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceNameOk() (*string, bool)`

GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceName

`func (o *AzurePrivateLinkConnection) SetPrivateLinkServiceName(v string)`

SetPrivateLinkServiceName sets PrivateLinkServiceName field to given value.

### HasPrivateLinkServiceName

`func (o *AzurePrivateLinkConnection) HasPrivateLinkServiceName() bool`

HasPrivateLinkServiceName returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *AzurePrivateLinkConnection) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *AzurePrivateLinkConnection) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *AzurePrivateLinkConnection) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### GetRegionName

`func (o *AzurePrivateLinkConnection) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AzurePrivateLinkConnection) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AzurePrivateLinkConnection) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AzurePrivateLinkConnection) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStatus

`func (o *AzurePrivateLinkConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AzurePrivateLinkConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AzurePrivateLinkConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AzurePrivateLinkConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


