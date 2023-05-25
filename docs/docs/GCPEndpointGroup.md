# GCPEndpointGroup

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

### NewGCPEndpointGroup

`func NewGCPEndpointGroup(cloudProvider string, ) *GCPEndpointGroup`

NewGCPEndpointGroup instantiates a new GCPEndpointGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPEndpointGroupWithDefaults

`func NewGCPEndpointGroupWithDefaults() *GCPEndpointGroup`

NewGCPEndpointGroupWithDefaults instantiates a new GCPEndpointGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *GCPEndpointGroup) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *GCPEndpointGroup) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *GCPEndpointGroup) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDeleteRequested

`func (o *GCPEndpointGroup) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *GCPEndpointGroup) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *GCPEndpointGroup) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *GCPEndpointGroup) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *GCPEndpointGroup) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *GCPEndpointGroup) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *GCPEndpointGroup) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *GCPEndpointGroup) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *GCPEndpointGroup) GetEndpoints() []GCPConsumerForwardingRule`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *GCPEndpointGroup) GetEndpointsOk() (*[]GCPConsumerForwardingRule, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *GCPEndpointGroup) SetEndpoints(v []GCPConsumerForwardingRule)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *GCPEndpointGroup) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GCPEndpointGroup) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GCPEndpointGroup) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GCPEndpointGroup) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GCPEndpointGroup) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GCPEndpointGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GCPEndpointGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GCPEndpointGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GCPEndpointGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


