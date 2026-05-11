# CollStatsLatencyNamespaceMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Metrics** | [**[]CollStatsLatencyNamespaceMetric**](CollStatsLatencyNamespaceMetric.md) | List of Coll Stats Latency metric names and their respective units. | [readonly] 

## Methods

### NewCollStatsLatencyNamespaceMetrics

`func NewCollStatsLatencyNamespaceMetrics(groupId string, metrics []CollStatsLatencyNamespaceMetric, ) *CollStatsLatencyNamespaceMetrics`

NewCollStatsLatencyNamespaceMetrics instantiates a new CollStatsLatencyNamespaceMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollStatsLatencyNamespaceMetricsWithDefaults

`func NewCollStatsLatencyNamespaceMetricsWithDefaults() *CollStatsLatencyNamespaceMetrics`

NewCollStatsLatencyNamespaceMetricsWithDefaults instantiates a new CollStatsLatencyNamespaceMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CollStatsLatencyNamespaceMetrics) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollStatsLatencyNamespaceMetrics) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollStatsLatencyNamespaceMetrics) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetLinks

`func (o *CollStatsLatencyNamespaceMetrics) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollStatsLatencyNamespaceMetrics) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollStatsLatencyNamespaceMetrics) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CollStatsLatencyNamespaceMetrics) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMetrics

`func (o *CollStatsLatencyNamespaceMetrics) GetMetrics() []CollStatsLatencyNamespaceMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CollStatsLatencyNamespaceMetrics) GetMetricsOk() (*[]CollStatsLatencyNamespaceMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CollStatsLatencyNamespaceMetrics) SetMetrics(v []CollStatsLatencyNamespaceMetric)`

SetMetrics sets Metrics field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


