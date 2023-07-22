# CloudSearchMetrics

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

### NewCloudSearchMetrics

`func NewCloudSearchMetrics(groupId string, processId string, ) *CloudSearchMetrics`

NewCloudSearchMetrics instantiates a new CloudSearchMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchMetricsWithDefaults

`func NewCloudSearchMetricsWithDefaults() *CloudSearchMetrics`

NewCloudSearchMetricsWithDefaults instantiates a new CloudSearchMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CloudSearchMetrics) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CloudSearchMetrics) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CloudSearchMetrics) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetHardwareMetrics

`func (o *CloudSearchMetrics) GetHardwareMetrics() []FTSMetric`

GetHardwareMetrics returns the HardwareMetrics field if non-nil, zero value otherwise.

### GetHardwareMetricsOk

`func (o *CloudSearchMetrics) GetHardwareMetricsOk() (*[]FTSMetric, bool)`

GetHardwareMetricsOk returns a tuple with the HardwareMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareMetrics

`func (o *CloudSearchMetrics) SetHardwareMetrics(v []FTSMetric)`

SetHardwareMetrics sets HardwareMetrics field to given value.

### HasHardwareMetrics

`func (o *CloudSearchMetrics) HasHardwareMetrics() bool`

HasHardwareMetrics returns a boolean if a field has been set.
### GetIndexMetrics

`func (o *CloudSearchMetrics) GetIndexMetrics() []FTSMetric`

GetIndexMetrics returns the IndexMetrics field if non-nil, zero value otherwise.

### GetIndexMetricsOk

`func (o *CloudSearchMetrics) GetIndexMetricsOk() (*[]FTSMetric, bool)`

GetIndexMetricsOk returns a tuple with the IndexMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMetrics

`func (o *CloudSearchMetrics) SetIndexMetrics(v []FTSMetric)`

SetIndexMetrics sets IndexMetrics field to given value.

### HasIndexMetrics

`func (o *CloudSearchMetrics) HasIndexMetrics() bool`

HasIndexMetrics returns a boolean if a field has been set.
### GetLinks

`func (o *CloudSearchMetrics) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CloudSearchMetrics) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CloudSearchMetrics) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CloudSearchMetrics) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetProcessId

`func (o *CloudSearchMetrics) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *CloudSearchMetrics) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *CloudSearchMetrics) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### GetStatusMetrics

`func (o *CloudSearchMetrics) GetStatusMetrics() []FTSMetric`

GetStatusMetrics returns the StatusMetrics field if non-nil, zero value otherwise.

### GetStatusMetricsOk

`func (o *CloudSearchMetrics) GetStatusMetricsOk() (*[]FTSMetric, bool)`

GetStatusMetricsOk returns a tuple with the StatusMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMetrics

`func (o *CloudSearchMetrics) SetStatusMetrics(v []FTSMetric)`

SetStatusMetrics sets StatusMetrics field to given value.

### HasStatusMetrics

`func (o *CloudSearchMetrics) HasStatusMetrics() bool`

HasStatusMetrics returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


