# GCPEndpointService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointGroupNames** | Pointer to **[]string** | List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service. | [optional] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**ServiceAttachmentNames** | Pointer to **[]string** | List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network. | [optional] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewGCPEndpointService

`func NewGCPEndpointService() *GCPEndpointService`

NewGCPEndpointService instantiates a new GCPEndpointService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPEndpointServiceWithDefaults

`func NewGCPEndpointServiceWithDefaults() *GCPEndpointService`

NewGCPEndpointServiceWithDefaults instantiates a new GCPEndpointService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointGroupNames

`func (o *GCPEndpointService) GetEndpointGroupNames() []string`

GetEndpointGroupNames returns the EndpointGroupNames field if non-nil, zero value otherwise.

### GetEndpointGroupNamesOk

`func (o *GCPEndpointService) GetEndpointGroupNamesOk() (*[]string, bool)`

GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupNames

`func (o *GCPEndpointService) SetEndpointGroupNames(v []string)`

SetEndpointGroupNames sets EndpointGroupNames field to given value.

### HasEndpointGroupNames

`func (o *GCPEndpointService) HasEndpointGroupNames() bool`

HasEndpointGroupNames returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GCPEndpointService) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GCPEndpointService) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GCPEndpointService) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GCPEndpointService) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *GCPEndpointService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCPEndpointService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCPEndpointService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GCPEndpointService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegionName

`func (o *GCPEndpointService) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GCPEndpointService) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GCPEndpointService) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GCPEndpointService) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetServiceAttachmentNames

`func (o *GCPEndpointService) GetServiceAttachmentNames() []string`

GetServiceAttachmentNames returns the ServiceAttachmentNames field if non-nil, zero value otherwise.

### GetServiceAttachmentNamesOk

`func (o *GCPEndpointService) GetServiceAttachmentNamesOk() (*[]string, bool)`

GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachmentNames

`func (o *GCPEndpointService) SetServiceAttachmentNames(v []string)`

SetServiceAttachmentNames sets ServiceAttachmentNames field to given value.

### HasServiceAttachmentNames

`func (o *GCPEndpointService) HasServiceAttachmentNames() bool`

HasServiceAttachmentNames returns a boolean if a field has been set.

### GetStatus

`func (o *GCPEndpointService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GCPEndpointService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GCPEndpointService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GCPEndpointService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


