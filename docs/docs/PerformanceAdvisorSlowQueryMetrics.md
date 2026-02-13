# PerformanceAdvisorSlowQueryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocsExamined** | Pointer to **int64** | The number of documents in the collection that MongoDB scanned in order to carry out the operation. | [optional] [readonly] 
**DocsExaminedReturnedRatio** | Pointer to **float64** | Ratio of documents examined to documents returned. | [optional] [readonly] 
**DocsReturned** | Pointer to **int64** | The number of documents returned by the operation. | [optional] [readonly] 
**FromUserConnection** | Pointer to **bool** | This boolean will be true when the server can identify the query source as non-server. This field is only available for MDB 8.0+. | [optional] [readonly] 
**HasIndexCoverage** | Pointer to **bool** | Indicates if the query has index coverage. | [optional] [readonly] 
**HasSort** | Pointer to **bool** | This boolean will be true when a query cannot use the ordering in the index to return the requested sorted results; i.e. MongoDB must sort the documents after it receives the documents from a cursor. | [optional] [readonly] 
**KeysExamined** | Pointer to **int64** | The number of index keys that MongoDB scanned in order to carry out the operation. | [optional] [readonly] 
**KeysExaminedReturnedRatio** | Pointer to **float64** | Ratio of keys examined to documents returned. | [optional] [readonly] 
**NumYields** | Pointer to **int64** | The number of times the operation yielded to allow other operations to complete. | [optional] [readonly] 
**OperationExecutionTime** | Pointer to **int64** | Total execution time of a query in milliseconds. | [optional] [readonly] 
**ResponseLength** | Pointer to **int64** | The length in bytes of the operation&#39;s result document. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorSlowQueryMetrics

`func NewPerformanceAdvisorSlowQueryMetrics() *PerformanceAdvisorSlowQueryMetrics`

NewPerformanceAdvisorSlowQueryMetrics instantiates a new PerformanceAdvisorSlowQueryMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorSlowQueryMetricsWithDefaults

`func NewPerformanceAdvisorSlowQueryMetricsWithDefaults() *PerformanceAdvisorSlowQueryMetrics`

NewPerformanceAdvisorSlowQueryMetricsWithDefaults instantiates a new PerformanceAdvisorSlowQueryMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocsExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsExamined() int64`

GetDocsExamined returns the DocsExamined field if non-nil, zero value otherwise.

### GetDocsExaminedOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsExaminedOk() (*int64, bool)`

GetDocsExaminedOk returns a tuple with the DocsExamined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsExamined(v int64)`

SetDocsExamined sets DocsExamined field to given value.

### HasDocsExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) HasDocsExamined() bool`

HasDocsExamined returns a boolean if a field has been set.
### GetDocsExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsExaminedReturnedRatio() float64`

GetDocsExaminedReturnedRatio returns the DocsExaminedReturnedRatio field if non-nil, zero value otherwise.

### GetDocsExaminedReturnedRatioOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsExaminedReturnedRatioOk() (*float64, bool)`

GetDocsExaminedReturnedRatioOk returns a tuple with the DocsExaminedReturnedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsExaminedReturnedRatio(v float64)`

SetDocsExaminedReturnedRatio sets DocsExaminedReturnedRatio field to given value.

### HasDocsExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) HasDocsExaminedReturnedRatio() bool`

HasDocsExaminedReturnedRatio returns a boolean if a field has been set.
### GetDocsReturned

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsReturned() int64`

GetDocsReturned returns the DocsReturned field if non-nil, zero value otherwise.

### GetDocsReturnedOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetDocsReturnedOk() (*int64, bool)`

GetDocsReturnedOk returns a tuple with the DocsReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsReturned

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsReturned(v int64)`

SetDocsReturned sets DocsReturned field to given value.

### HasDocsReturned

`func (o *PerformanceAdvisorSlowQueryMetrics) HasDocsReturned() bool`

HasDocsReturned returns a boolean if a field has been set.
### GetFromUserConnection

`func (o *PerformanceAdvisorSlowQueryMetrics) GetFromUserConnection() bool`

GetFromUserConnection returns the FromUserConnection field if non-nil, zero value otherwise.

### GetFromUserConnectionOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetFromUserConnectionOk() (*bool, bool)`

GetFromUserConnectionOk returns a tuple with the FromUserConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserConnection

`func (o *PerformanceAdvisorSlowQueryMetrics) SetFromUserConnection(v bool)`

SetFromUserConnection sets FromUserConnection field to given value.

### HasFromUserConnection

`func (o *PerformanceAdvisorSlowQueryMetrics) HasFromUserConnection() bool`

HasFromUserConnection returns a boolean if a field has been set.
### GetHasIndexCoverage

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasIndexCoverage() bool`

GetHasIndexCoverage returns the HasIndexCoverage field if non-nil, zero value otherwise.

### GetHasIndexCoverageOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasIndexCoverageOk() (*bool, bool)`

GetHasIndexCoverageOk returns a tuple with the HasIndexCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIndexCoverage

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasIndexCoverage(v bool)`

SetHasIndexCoverage sets HasIndexCoverage field to given value.

### HasHasIndexCoverage

`func (o *PerformanceAdvisorSlowQueryMetrics) HasHasIndexCoverage() bool`

HasHasIndexCoverage returns a boolean if a field has been set.
### GetHasSort

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasSort() bool`

GetHasSort returns the HasSort field if non-nil, zero value otherwise.

### GetHasSortOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasSortOk() (*bool, bool)`

GetHasSortOk returns a tuple with the HasSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSort

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasSort(v bool)`

SetHasSort sets HasSort field to given value.

### HasHasSort

`func (o *PerformanceAdvisorSlowQueryMetrics) HasHasSort() bool`

HasHasSort returns a boolean if a field has been set.
### GetKeysExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) GetKeysExamined() int64`

GetKeysExamined returns the KeysExamined field if non-nil, zero value otherwise.

### GetKeysExaminedOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetKeysExaminedOk() (*int64, bool)`

GetKeysExaminedOk returns a tuple with the KeysExamined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) SetKeysExamined(v int64)`

SetKeysExamined sets KeysExamined field to given value.

### HasKeysExamined

`func (o *PerformanceAdvisorSlowQueryMetrics) HasKeysExamined() bool`

HasKeysExamined returns a boolean if a field has been set.
### GetKeysExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) GetKeysExaminedReturnedRatio() float64`

GetKeysExaminedReturnedRatio returns the KeysExaminedReturnedRatio field if non-nil, zero value otherwise.

### GetKeysExaminedReturnedRatioOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetKeysExaminedReturnedRatioOk() (*float64, bool)`

GetKeysExaminedReturnedRatioOk returns a tuple with the KeysExaminedReturnedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) SetKeysExaminedReturnedRatio(v float64)`

SetKeysExaminedReturnedRatio sets KeysExaminedReturnedRatio field to given value.

### HasKeysExaminedReturnedRatio

`func (o *PerformanceAdvisorSlowQueryMetrics) HasKeysExaminedReturnedRatio() bool`

HasKeysExaminedReturnedRatio returns a boolean if a field has been set.
### GetNumYields

`func (o *PerformanceAdvisorSlowQueryMetrics) GetNumYields() int64`

GetNumYields returns the NumYields field if non-nil, zero value otherwise.

### GetNumYieldsOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetNumYieldsOk() (*int64, bool)`

GetNumYieldsOk returns a tuple with the NumYields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumYields

`func (o *PerformanceAdvisorSlowQueryMetrics) SetNumYields(v int64)`

SetNumYields sets NumYields field to given value.

### HasNumYields

`func (o *PerformanceAdvisorSlowQueryMetrics) HasNumYields() bool`

HasNumYields returns a boolean if a field has been set.
### GetOperationExecutionTime

`func (o *PerformanceAdvisorSlowQueryMetrics) GetOperationExecutionTime() int64`

GetOperationExecutionTime returns the OperationExecutionTime field if non-nil, zero value otherwise.

### GetOperationExecutionTimeOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetOperationExecutionTimeOk() (*int64, bool)`

GetOperationExecutionTimeOk returns a tuple with the OperationExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationExecutionTime

`func (o *PerformanceAdvisorSlowQueryMetrics) SetOperationExecutionTime(v int64)`

SetOperationExecutionTime sets OperationExecutionTime field to given value.

### HasOperationExecutionTime

`func (o *PerformanceAdvisorSlowQueryMetrics) HasOperationExecutionTime() bool`

HasOperationExecutionTime returns a boolean if a field has been set.
### GetResponseLength

`func (o *PerformanceAdvisorSlowQueryMetrics) GetResponseLength() int64`

GetResponseLength returns the ResponseLength field if non-nil, zero value otherwise.

### GetResponseLengthOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetResponseLengthOk() (*int64, bool)`

GetResponseLengthOk returns a tuple with the ResponseLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLength

`func (o *PerformanceAdvisorSlowQueryMetrics) SetResponseLength(v int64)`

SetResponseLength sets ResponseLength field to given value.

### HasResponseLength

`func (o *PerformanceAdvisorSlowQueryMetrics) HasResponseLength() bool`

HasResponseLength returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


