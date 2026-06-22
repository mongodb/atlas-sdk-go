# StreamsStartProcessorFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, simulates the operation without making any changes. | [optional] 
**Mode** | Pointer to **string** | Strategy for the processor: GRACEFUL - attempt to stop the processor, error if processor cannot be stopped. if stop was successful, start the processor in the new region with the latest checkpoint.  FORCED - attempt to stop the processor, proceed to starting the processor in the new region with checkpoints disabled regardless of whether or not the stop succeeds. | [optional] 
**Region** | **string** | Cloud provider region where the stream processor should be started in failover mode. The region must be a valid region for the stream processor&#39;s cloud provider and must be included in the tenant&#39;s configured failover regions, or it may be the tenant&#39;s default (primary) region. | 

## Methods

### NewStreamsStartProcessorFailover

`func NewStreamsStartProcessorFailover(region string, ) *StreamsStartProcessorFailover`

NewStreamsStartProcessorFailover instantiates a new StreamsStartProcessorFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsStartProcessorFailoverWithDefaults

`func NewStreamsStartProcessorFailoverWithDefaults() *StreamsStartProcessorFailover`

NewStreamsStartProcessorFailoverWithDefaults instantiates a new StreamsStartProcessorFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *StreamsStartProcessorFailover) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *StreamsStartProcessorFailover) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *StreamsStartProcessorFailover) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *StreamsStartProcessorFailover) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.
### GetMode

`func (o *StreamsStartProcessorFailover) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StreamsStartProcessorFailover) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StreamsStartProcessorFailover) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StreamsStartProcessorFailover) HasMode() bool`

HasMode returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsStartProcessorFailover) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsStartProcessorFailover) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsStartProcessorFailover) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


