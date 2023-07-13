# PerformanceAdvisorOpStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ms** | Pointer to **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds. This parameter relates to the **duration** query parameter. | [optional] [readonly] 
**NReturned** | Pointer to **int64** | Number of results that the query returns. | [optional] [readonly] 
**NScanned** | Pointer to **int64** | Number of documents that the query read. | [optional] [readonly] 
**Ts** | Pointer to **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). This parameter relates to the **since** query parameter. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorOpStats

`func NewPerformanceAdvisorOpStats() *PerformanceAdvisorOpStats`

NewPerformanceAdvisorOpStats instantiates a new PerformanceAdvisorOpStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorOpStatsWithDefaults

`func NewPerformanceAdvisorOpStatsWithDefaults() *PerformanceAdvisorOpStats`

NewPerformanceAdvisorOpStatsWithDefaults instantiates a new PerformanceAdvisorOpStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMs

`func (o *PerformanceAdvisorOpStats) GetMs() int64`

GetMs returns the Ms field if non-nil, zero value otherwise.

### GetMsOk

`func (o *PerformanceAdvisorOpStats) GetMsOk() (*int64, bool)`

GetMsOk returns a tuple with the Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMs

`func (o *PerformanceAdvisorOpStats) SetMs(v int64)`

SetMs sets Ms field to given value.

### HasMs

`func (o *PerformanceAdvisorOpStats) HasMs() bool`

HasMs returns a boolean if a field has been set.
### GetNReturned

`func (o *PerformanceAdvisorOpStats) GetNReturned() int64`

GetNReturned returns the NReturned field if non-nil, zero value otherwise.

### GetNReturnedOk

`func (o *PerformanceAdvisorOpStats) GetNReturnedOk() (*int64, bool)`

GetNReturnedOk returns a tuple with the NReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNReturned

`func (o *PerformanceAdvisorOpStats) SetNReturned(v int64)`

SetNReturned sets NReturned field to given value.

### HasNReturned

`func (o *PerformanceAdvisorOpStats) HasNReturned() bool`

HasNReturned returns a boolean if a field has been set.
### GetNScanned

`func (o *PerformanceAdvisorOpStats) GetNScanned() int64`

GetNScanned returns the NScanned field if non-nil, zero value otherwise.

### GetNScannedOk

`func (o *PerformanceAdvisorOpStats) GetNScannedOk() (*int64, bool)`

GetNScannedOk returns a tuple with the NScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNScanned

`func (o *PerformanceAdvisorOpStats) SetNScanned(v int64)`

SetNScanned sets NScanned field to given value.

### HasNScanned

`func (o *PerformanceAdvisorOpStats) HasNScanned() bool`

HasNScanned returns a boolean if a field has been set.
### GetTs

`func (o *PerformanceAdvisorOpStats) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *PerformanceAdvisorOpStats) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *PerformanceAdvisorOpStats) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *PerformanceAdvisorOpStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


