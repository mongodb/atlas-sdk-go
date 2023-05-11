# FTSMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**HardwareMetrics** | Pointer to [**[]FTSMetric**](FTSMetric.md) | List that contains all host compute, memory, and storage utilization dedicated to Atlas Search when MongoDB Atlas received this request. | [optional] [readonly] 
**IndexMetrics** | Pointer to [**[]FTSMetric**](FTSMetric.md) | List that contains all performance and utilization measurements that Atlas Search index performed by the time MongoDB Atlas received this request. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ProcessId** | **string** | Hostname and port that identifies the process. | [readonly] 
**StatusMetrics** | Pointer to [**[]FTSMetric**](FTSMetric.md) | List that contains all available Atlas Search status metrics when MongoDB Atlas received this request. | [optional] [readonly] 

## Methods

### NewFTSMetrics

`func NewFTSMetrics(groupId string, processId string, ) *FTSMetrics`

NewFTSMetrics instantiates a new FTSMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSMetricsWithDefaults

`func NewFTSMetricsWithDefaults() *FTSMetrics`

NewFTSMetricsWithDefaults instantiates a new FTSMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *FTSMetrics) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FTSMetrics) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FTSMetrics) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHardwareMetrics

`func (o *FTSMetrics) GetHardwareMetrics() []FTSMetric`

GetHardwareMetrics returns the HardwareMetrics field if non-nil, zero value otherwise.

### GetHardwareMetricsOk

`func (o *FTSMetrics) GetHardwareMetricsOk() (*[]FTSMetric, bool)`

GetHardwareMetricsOk returns a tuple with the HardwareMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareMetrics

`func (o *FTSMetrics) SetHardwareMetrics(v []FTSMetric)`

SetHardwareMetrics sets HardwareMetrics field to given value.

### HasHardwareMetrics

`func (o *FTSMetrics) HasHardwareMetrics() bool`

HasHardwareMetrics returns a boolean if a field has been set.

### GetIndexMetrics

`func (o *FTSMetrics) GetIndexMetrics() []FTSMetric`

GetIndexMetrics returns the IndexMetrics field if non-nil, zero value otherwise.

### GetIndexMetricsOk

`func (o *FTSMetrics) GetIndexMetricsOk() (*[]FTSMetric, bool)`

GetIndexMetricsOk returns a tuple with the IndexMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMetrics

`func (o *FTSMetrics) SetIndexMetrics(v []FTSMetric)`

SetIndexMetrics sets IndexMetrics field to given value.

### HasIndexMetrics

`func (o *FTSMetrics) HasIndexMetrics() bool`

HasIndexMetrics returns a boolean if a field has been set.

### GetLinks

`func (o *FTSMetrics) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FTSMetrics) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FTSMetrics) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FTSMetrics) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProcessId

`func (o *FTSMetrics) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *FTSMetrics) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *FTSMetrics) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetStatusMetrics

`func (o *FTSMetrics) GetStatusMetrics() []FTSMetric`

GetStatusMetrics returns the StatusMetrics field if non-nil, zero value otherwise.

### GetStatusMetricsOk

`func (o *FTSMetrics) GetStatusMetricsOk() (*[]FTSMetric, bool)`

GetStatusMetricsOk returns a tuple with the StatusMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMetrics

`func (o *FTSMetrics) SetStatusMetrics(v []FTSMetric)`

SetStatusMetrics sets StatusMetrics field to given value.

### HasStatusMetrics

`func (o *FTSMetrics) HasStatusMetrics() bool`

HasStatusMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


