# StreamsProcessorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Mode** | Pointer to **string** | Strategy for the processor: GRACEFUL - attempt to stop the processor, error if processor cannot be stopped. if stop was successful, start the processor in the new region with the latest checkpoint.  FORCED - attempt to stop the processor, proceed to starting the processor in the new region with checkpoints disabled regardless of whether or not the stop succeeds. | [optional] 
**Region** | Pointer to **string** | Name of the region against which to apply the status change. Required when &#x60;status&#x60; is &#x60;FAILED_OVER&#x60;; optional otherwise. | [optional] 
**Status** | **string** | Represents the desired action to apply to stream processors within a workspace, such as starting all processors, stopping all processors, or performing a bulk regional failover. | 

## Methods

### NewStreamsProcessorStatus

`func NewStreamsProcessorStatus(status string, ) *StreamsProcessorStatus`

NewStreamsProcessorStatus instantiates a new StreamsProcessorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsProcessorStatusWithDefaults

`func NewStreamsProcessorStatusWithDefaults() *StreamsProcessorStatus`

NewStreamsProcessorStatusWithDefaults instantiates a new StreamsProcessorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsProcessorStatus) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsProcessorStatus) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsProcessorStatus) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsProcessorStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMode

`func (o *StreamsProcessorStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StreamsProcessorStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StreamsProcessorStatus) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StreamsProcessorStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsProcessorStatus) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsProcessorStatus) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsProcessorStatus) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsProcessorStatus) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetStatus

`func (o *StreamsProcessorStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamsProcessorStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamsProcessorStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


