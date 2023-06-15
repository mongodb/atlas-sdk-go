# PrivateGCPEndpointGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Cloud service provider that serves the requested endpoint. | [readonly] 
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] [readonly] 
**Endpoints** | Pointer to [**[]GCPConsumerForwardingRule**](GCPConsumerForwardingRule.md) | List of individual private endpoints that comprise this endpoint group. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Google Cloud network endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewPrivateGCPEndpointGroup

`func NewPrivateGCPEndpointGroup(cloudProvider string, ) *PrivateGCPEndpointGroup`

NewPrivateGCPEndpointGroup instantiates a new PrivateGCPEndpointGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateGCPEndpointGroupWithDefaults

`func NewPrivateGCPEndpointGroupWithDefaults() *PrivateGCPEndpointGroup`

NewPrivateGCPEndpointGroupWithDefaults instantiates a new PrivateGCPEndpointGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *PrivateGCPEndpointGroup) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *PrivateGCPEndpointGroup) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *PrivateGCPEndpointGroup) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDeleteRequested

`func (o *PrivateGCPEndpointGroup) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *PrivateGCPEndpointGroup) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *PrivateGCPEndpointGroup) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *PrivateGCPEndpointGroup) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *PrivateGCPEndpointGroup) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *PrivateGCPEndpointGroup) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *PrivateGCPEndpointGroup) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *PrivateGCPEndpointGroup) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *PrivateGCPEndpointGroup) GetEndpoints() []GCPConsumerForwardingRule`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PrivateGCPEndpointGroup) GetEndpointsOk() (*[]GCPConsumerForwardingRule, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PrivateGCPEndpointGroup) SetEndpoints(v []GCPConsumerForwardingRule)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PrivateGCPEndpointGroup) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PrivateGCPEndpointGroup) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PrivateGCPEndpointGroup) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PrivateGCPEndpointGroup) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PrivateGCPEndpointGroup) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *PrivateGCPEndpointGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateGCPEndpointGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateGCPEndpointGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateGCPEndpointGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


