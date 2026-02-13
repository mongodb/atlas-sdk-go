# StreamsTenantUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Region** | Pointer to **string** | Name of the cloud provider region hosting Atlas Stream Processing. | [optional] 
**StreamConfig** | Pointer to [**StreamConfig**](StreamConfig.md) |  | [optional] 

## Methods

### NewStreamsTenantUpdateRequest

`func NewStreamsTenantUpdateRequest() *StreamsTenantUpdateRequest`

NewStreamsTenantUpdateRequest instantiates a new StreamsTenantUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsTenantUpdateRequestWithDefaults

`func NewStreamsTenantUpdateRequestWithDefaults() *StreamsTenantUpdateRequest`

NewStreamsTenantUpdateRequestWithDefaults instantiates a new StreamsTenantUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *StreamsTenantUpdateRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *StreamsTenantUpdateRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *StreamsTenantUpdateRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *StreamsTenantUpdateRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsTenantUpdateRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsTenantUpdateRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsTenantUpdateRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsTenantUpdateRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsTenantUpdateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsTenantUpdateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsTenantUpdateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsTenantUpdateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetStreamConfig

`func (o *StreamsTenantUpdateRequest) GetStreamConfig() StreamConfig`

GetStreamConfig returns the StreamConfig field if non-nil, zero value otherwise.

### GetStreamConfigOk

`func (o *StreamsTenantUpdateRequest) GetStreamConfigOk() (*StreamConfig, bool)`

GetStreamConfigOk returns a tuple with the StreamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamConfig

`func (o *StreamsTenantUpdateRequest) SetStreamConfig(v StreamConfig)`

SetStreamConfig sets StreamConfig field to given value.

### HasStreamConfig

`func (o *StreamsTenantUpdateRequest) HasStreamConfig() bool`

HasStreamConfig returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


