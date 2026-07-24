# StreamsTenantUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider. | [optional] 
**FailoverRegions** | Pointer to [**[]StreamsDataProcessRegion**](StreamsDataProcessRegion.md) | Failover regions for the stream workspace. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ProcessorStatus** | Pointer to [**StreamsProcessorStatus**](StreamsProcessorStatus.md) |  | [optional] 
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

### SetCloudProviderNil

`func (o *StreamsTenantUpdateRequest) SetCloudProviderNil()`

SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled, overriding any value previously set with SetCloudProvider. Calling SetCloudProvider again clears the null override.

### GetFailoverRegions

`func (o *StreamsTenantUpdateRequest) GetFailoverRegions() []StreamsDataProcessRegion`

GetFailoverRegions returns the FailoverRegions field if non-nil, zero value otherwise.

### GetFailoverRegionsOk

`func (o *StreamsTenantUpdateRequest) GetFailoverRegionsOk() (*[]StreamsDataProcessRegion, bool)`

GetFailoverRegionsOk returns a tuple with the FailoverRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverRegions

`func (o *StreamsTenantUpdateRequest) SetFailoverRegions(v []StreamsDataProcessRegion)`

SetFailoverRegions sets FailoverRegions field to given value.

### HasFailoverRegions

`func (o *StreamsTenantUpdateRequest) HasFailoverRegions() bool`

HasFailoverRegions returns a boolean if a field has been set.

### SetFailoverRegionsNil

`func (o *StreamsTenantUpdateRequest) SetFailoverRegionsNil()`

SetFailoverRegionsNil sets FailoverRegions to an explicit JSON null when marshaled, overriding any value previously set with SetFailoverRegions. Calling SetFailoverRegions again clears the null override.

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

### SetLinksNil

`func (o *StreamsTenantUpdateRequest) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetProcessorStatus

`func (o *StreamsTenantUpdateRequest) GetProcessorStatus() StreamsProcessorStatus`

GetProcessorStatus returns the ProcessorStatus field if non-nil, zero value otherwise.

### GetProcessorStatusOk

`func (o *StreamsTenantUpdateRequest) GetProcessorStatusOk() (*StreamsProcessorStatus, bool)`

GetProcessorStatusOk returns a tuple with the ProcessorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorStatus

`func (o *StreamsTenantUpdateRequest) SetProcessorStatus(v StreamsProcessorStatus)`

SetProcessorStatus sets ProcessorStatus field to given value.

### HasProcessorStatus

`func (o *StreamsTenantUpdateRequest) HasProcessorStatus() bool`

HasProcessorStatus returns a boolean if a field has been set.

### SetProcessorStatusNil

`func (o *StreamsTenantUpdateRequest) SetProcessorStatusNil()`

SetProcessorStatusNil sets ProcessorStatus to an explicit JSON null when marshaled, overriding any value previously set with SetProcessorStatus. Calling SetProcessorStatus again clears the null override.

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

### SetRegionNil

`func (o *StreamsTenantUpdateRequest) SetRegionNil()`

SetRegionNil sets Region to an explicit JSON null when marshaled, overriding any value previously set with SetRegion. Calling SetRegion again clears the null override.

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

### SetStreamConfigNil

`func (o *StreamsTenantUpdateRequest) SetStreamConfigNil()`

SetStreamConfigNil sets StreamConfig to an explicit JSON null when marshaled, overriding any value previously set with SetStreamConfig. Calling SetStreamConfig again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


