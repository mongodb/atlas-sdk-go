# PerformanceAdvisorSlowQueryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocsExamined** | Pointer to **int64** | The number of documents in the collection that MongoDB scanned in order to carry out the operation. | [optional] [readonly] 
**DocsExaminedReturnedRatio** | Pointer to **float64** | Ratio of documents examined to documents returned. | [optional] [readonly] 
**DocsReturned** | Pointer to **int64** | The number of documents returned by the operation. | [optional] [readonly] 
**FromUserConnection** | Pointer to **bool** | This boolean will be true when the server can identify the query source as non-server. This field is only available for MDB 8.0+. | [optional] [readonly] 
**HasAutoEmbedding** | Pointer to **bool** | Flag that indicates whether the slow query used automated embedding, where MongoDB Cloud generates embeddings from raw text at query time instead of the client supplying a precomputed vector. | [optional] [readonly] 
**HasIndexCoverage** | Pointer to **bool** | Indicates if the query has index coverage. | [optional] [readonly] 
**HasRerank** | Pointer to **bool** | Flag that indicates whether the slow query used the &#x60;$rerank&#x60; aggregation stage, which reorders results using Voyage AI reranking models. Always &#x60;false&#x60; for MongoDB deployments earlier than 8.3. | [optional] [readonly] 
**HasSort** | Pointer to **bool** | This boolean will be true when a query cannot use the ordering in the index to return the requested sorted results; i.e. MongoDB must sort the documents after it receives the documents from a cursor. | [optional] [readonly] 
**KeysExamined** | Pointer to **int64** | The number of index keys that MongoDB scanned in order to carry out the operation. | [optional] [readonly] 
**KeysExaminedReturnedRatio** | Pointer to **float64** | Ratio of keys examined to documents returned. | [optional] [readonly] 
**NumYields** | Pointer to **int64** | The number of times the operation yielded to allow other operations to complete. | [optional] [readonly] 
**OperationExecutionTime** | Pointer to **int64** | Total execution time of a query in milliseconds. | [optional] [readonly] 
**ResponseLength** | Pointer to **int64** | The length in bytes of the operation&#39;s result document. | [optional] [readonly] 
**TokensUsed** | Pointer to **int64** | The total inference tokens consumed by this operation, including tokens used by &#x60;$rerank&#x60;. Returned only for inference queries that consumed tokens; it is omitted otherwise, including for MongoDB deployments earlier than 8.3. | [optional] [readonly] 

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

### SetDocsExaminedNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsExaminedNil()`

SetDocsExaminedNil sets DocsExamined to an explicit JSON null when marshaled, overriding any value previously set with SetDocsExamined. Calling SetDocsExamined again clears the null override.

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

### SetDocsExaminedReturnedRatioNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsExaminedReturnedRatioNil()`

SetDocsExaminedReturnedRatioNil sets DocsExaminedReturnedRatio to an explicit JSON null when marshaled, overriding any value previously set with SetDocsExaminedReturnedRatio. Calling SetDocsExaminedReturnedRatio again clears the null override.

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

### SetDocsReturnedNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetDocsReturnedNil()`

SetDocsReturnedNil sets DocsReturned to an explicit JSON null when marshaled, overriding any value previously set with SetDocsReturned. Calling SetDocsReturned again clears the null override.

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

### SetFromUserConnectionNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetFromUserConnectionNil()`

SetFromUserConnectionNil sets FromUserConnection to an explicit JSON null when marshaled, overriding any value previously set with SetFromUserConnection. Calling SetFromUserConnection again clears the null override.

### GetHasAutoEmbedding

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasAutoEmbedding() bool`

GetHasAutoEmbedding returns the HasAutoEmbedding field if non-nil, zero value otherwise.

### GetHasAutoEmbeddingOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasAutoEmbeddingOk() (*bool, bool)`

GetHasAutoEmbeddingOk returns a tuple with the HasAutoEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoEmbedding

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasAutoEmbedding(v bool)`

SetHasAutoEmbedding sets HasAutoEmbedding field to given value.

### HasHasAutoEmbedding

`func (o *PerformanceAdvisorSlowQueryMetrics) HasHasAutoEmbedding() bool`

HasHasAutoEmbedding returns a boolean if a field has been set.

### SetHasAutoEmbeddingNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasAutoEmbeddingNil()`

SetHasAutoEmbeddingNil sets HasAutoEmbedding to an explicit JSON null when marshaled, overriding any value previously set with SetHasAutoEmbedding. Calling SetHasAutoEmbedding again clears the null override.

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

### SetHasIndexCoverageNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasIndexCoverageNil()`

SetHasIndexCoverageNil sets HasIndexCoverage to an explicit JSON null when marshaled, overriding any value previously set with SetHasIndexCoverage. Calling SetHasIndexCoverage again clears the null override.

### GetHasRerank

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasRerank() bool`

GetHasRerank returns the HasRerank field if non-nil, zero value otherwise.

### GetHasRerankOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetHasRerankOk() (*bool, bool)`

GetHasRerankOk returns a tuple with the HasRerank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRerank

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasRerank(v bool)`

SetHasRerank sets HasRerank field to given value.

### HasHasRerank

`func (o *PerformanceAdvisorSlowQueryMetrics) HasHasRerank() bool`

HasHasRerank returns a boolean if a field has been set.

### SetHasRerankNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasRerankNil()`

SetHasRerankNil sets HasRerank to an explicit JSON null when marshaled, overriding any value previously set with SetHasRerank. Calling SetHasRerank again clears the null override.

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

### SetHasSortNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetHasSortNil()`

SetHasSortNil sets HasSort to an explicit JSON null when marshaled, overriding any value previously set with SetHasSort. Calling SetHasSort again clears the null override.

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

### SetKeysExaminedNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetKeysExaminedNil()`

SetKeysExaminedNil sets KeysExamined to an explicit JSON null when marshaled, overriding any value previously set with SetKeysExamined. Calling SetKeysExamined again clears the null override.

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

### SetKeysExaminedReturnedRatioNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetKeysExaminedReturnedRatioNil()`

SetKeysExaminedReturnedRatioNil sets KeysExaminedReturnedRatio to an explicit JSON null when marshaled, overriding any value previously set with SetKeysExaminedReturnedRatio. Calling SetKeysExaminedReturnedRatio again clears the null override.

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

### SetNumYieldsNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetNumYieldsNil()`

SetNumYieldsNil sets NumYields to an explicit JSON null when marshaled, overriding any value previously set with SetNumYields. Calling SetNumYields again clears the null override.

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

### SetOperationExecutionTimeNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetOperationExecutionTimeNil()`

SetOperationExecutionTimeNil sets OperationExecutionTime to an explicit JSON null when marshaled, overriding any value previously set with SetOperationExecutionTime. Calling SetOperationExecutionTime again clears the null override.

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

### SetResponseLengthNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetResponseLengthNil()`

SetResponseLengthNil sets ResponseLength to an explicit JSON null when marshaled, overriding any value previously set with SetResponseLength. Calling SetResponseLength again clears the null override.

### GetTokensUsed

`func (o *PerformanceAdvisorSlowQueryMetrics) GetTokensUsed() int64`

GetTokensUsed returns the TokensUsed field if non-nil, zero value otherwise.

### GetTokensUsedOk

`func (o *PerformanceAdvisorSlowQueryMetrics) GetTokensUsedOk() (*int64, bool)`

GetTokensUsedOk returns a tuple with the TokensUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensUsed

`func (o *PerformanceAdvisorSlowQueryMetrics) SetTokensUsed(v int64)`

SetTokensUsed sets TokensUsed field to given value.

### HasTokensUsed

`func (o *PerformanceAdvisorSlowQueryMetrics) HasTokensUsed() bool`

HasTokensUsed returns a boolean if a field has been set.

### SetTokensUsedNil

`func (o *PerformanceAdvisorSlowQueryMetrics) SetTokensUsedNil()`

SetTokensUsedNil sets TokensUsed to an explicit JSON null when marshaled, overriding any value previously set with SetTokensUsed. Calling SetTokensUsed again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


