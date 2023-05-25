# AWSPrivateLinkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Cloud service provider that serves the requested endpoint service. | [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**InterfaceEndpoints** | Pointer to **[]string** | List of strings that identify private endpoint interfaces applied to the specified project. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewAWSPrivateLinkConnection

`func NewAWSPrivateLinkConnection(cloudProvider string, ) *AWSPrivateLinkConnection`

NewAWSPrivateLinkConnection instantiates a new AWSPrivateLinkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSPrivateLinkConnectionWithDefaults

`func NewAWSPrivateLinkConnectionWithDefaults() *AWSPrivateLinkConnection`

NewAWSPrivateLinkConnectionWithDefaults instantiates a new AWSPrivateLinkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AWSPrivateLinkConnection) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AWSPrivateLinkConnection) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AWSPrivateLinkConnection) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetEndpointServiceName

`func (o *AWSPrivateLinkConnection) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *AWSPrivateLinkConnection) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *AWSPrivateLinkConnection) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *AWSPrivateLinkConnection) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AWSPrivateLinkConnection) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AWSPrivateLinkConnection) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AWSPrivateLinkConnection) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AWSPrivateLinkConnection) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *AWSPrivateLinkConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSPrivateLinkConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSPrivateLinkConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSPrivateLinkConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceEndpoints

`func (o *AWSPrivateLinkConnection) GetInterfaceEndpoints() []string`

GetInterfaceEndpoints returns the InterfaceEndpoints field if non-nil, zero value otherwise.

### GetInterfaceEndpointsOk

`func (o *AWSPrivateLinkConnection) GetInterfaceEndpointsOk() (*[]string, bool)`

GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpoints

`func (o *AWSPrivateLinkConnection) SetInterfaceEndpoints(v []string)`

SetInterfaceEndpoints sets InterfaceEndpoints field to given value.

### HasInterfaceEndpoints

`func (o *AWSPrivateLinkConnection) HasInterfaceEndpoints() bool`

HasInterfaceEndpoints returns a boolean if a field has been set.

### GetRegionName

`func (o *AWSPrivateLinkConnection) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AWSPrivateLinkConnection) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AWSPrivateLinkConnection) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *AWSPrivateLinkConnection) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStatus

`func (o *AWSPrivateLinkConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AWSPrivateLinkConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AWSPrivateLinkConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AWSPrivateLinkConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


