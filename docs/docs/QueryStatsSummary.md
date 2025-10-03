# QueryStatsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgWorkingMillis** | Pointer to **float64** | Average total time in milliseconds spent running queries with the given query shape. If the query resulted in getMore commands, this metric includes the time spent processing the getMore requests. This metric does not include time spent waiting for the client. | [optional] 
**BytesRead** | Pointer to **float64** | The number of bytes read by the given query shape from the disk to the cache. | [optional] 
**Command** | Pointer to **string** | The MongoDB command issued for this query shape. | [optional] 
**DocsExamined** | Pointer to **float64** | Total number of documents examined by queries with the given query shape. | [optional] 
**DocsExaminedRatio** | Pointer to **float64** | Ratio of documents examined to documents returned by queries with the given query shape. | [optional] 
**DocsReturned** | Pointer to **float64** | Total number of documents returned by queries with the given query shape. | [optional] 
**ExecCount** | Pointer to **float64** | Total number of times that queries with the given query shape have been executed. | [optional] 
**KeysExamined** | Pointer to **float64** | Total number of in-bounds and out-of-bounds index keys examined by queries with the given query shape. | [optional] 
**KeysExaminedRatio** | Pointer to **float64** | Ratio of in-bounds and out-of-bounds index keys examined to indexes containing documents returned by queries with the given query shape. | [optional] 
**LastExecMicros** | Pointer to **float64** | Execution runtime in microseconds for the most recent query with the given query shape. | [optional] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] 
**P50ExecMicros** | Pointer to **float64** | The 50th percentile value of execution time in microseconds. | [optional] 
**P90ExecMicros** | Pointer to **float64** | The 90th percentile value of execution time in microseconds. | [optional] 
**P99ExecMicros** | Pointer to **float64** | The 99th percentile value of execution time in microseconds. | [optional] 
**QueryShape** | Pointer to **string** | A query shape is a set of specifications that group similar queries together. Specifications can include filters, sorts, projections, aggregation pipeline stages, a namespace, and others. Queries that have similar specifications have the same query shape. | [optional] 
**QueryShapeHash** | Pointer to **string** | A hexadecimal string that represents the hash of a MongoDB query shape. | [optional] 
**SystemQuery** | Pointer to **bool** | Indicates whether this query shape represents a system-initiated query. | [optional] 
**TotalTimeToResponseMicros** | Pointer to **float64** | Time in microseconds spent from the beginning of query processing to the first server response. | [optional] 
**TotalWorkingMillis** | Pointer to **float64** | Total time in milliseconds spent running queries with the given query shape. If the query resulted in &#x60;getMore&#x60; commands, this metric includes the time spent processing the &#x60;getMore&#x60; requests. This metric does not include time spent waiting for the client. | [optional] 

## Methods

### NewQueryStatsSummary

`func NewQueryStatsSummary() *QueryStatsSummary`

NewQueryStatsSummary instantiates a new QueryStatsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatsSummaryWithDefaults

`func NewQueryStatsSummaryWithDefaults() *QueryStatsSummary`

NewQueryStatsSummaryWithDefaults instantiates a new QueryStatsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgWorkingMillis

`func (o *QueryStatsSummary) GetAvgWorkingMillis() float64`

GetAvgWorkingMillis returns the AvgWorkingMillis field if non-nil, zero value otherwise.

### GetAvgWorkingMillisOk

`func (o *QueryStatsSummary) GetAvgWorkingMillisOk() (*float64, bool)`

GetAvgWorkingMillisOk returns a tuple with the AvgWorkingMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgWorkingMillis

`func (o *QueryStatsSummary) SetAvgWorkingMillis(v float64)`

SetAvgWorkingMillis sets AvgWorkingMillis field to given value.

### HasAvgWorkingMillis

`func (o *QueryStatsSummary) HasAvgWorkingMillis() bool`

HasAvgWorkingMillis returns a boolean if a field has been set.
### GetBytesRead

`func (o *QueryStatsSummary) GetBytesRead() float64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *QueryStatsSummary) GetBytesReadOk() (*float64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *QueryStatsSummary) SetBytesRead(v float64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *QueryStatsSummary) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.
### GetCommand

`func (o *QueryStatsSummary) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *QueryStatsSummary) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *QueryStatsSummary) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *QueryStatsSummary) HasCommand() bool`

HasCommand returns a boolean if a field has been set.
### GetDocsExamined

`func (o *QueryStatsSummary) GetDocsExamined() float64`

GetDocsExamined returns the DocsExamined field if non-nil, zero value otherwise.

### GetDocsExaminedOk

`func (o *QueryStatsSummary) GetDocsExaminedOk() (*float64, bool)`

GetDocsExaminedOk returns a tuple with the DocsExamined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsExamined

`func (o *QueryStatsSummary) SetDocsExamined(v float64)`

SetDocsExamined sets DocsExamined field to given value.

### HasDocsExamined

`func (o *QueryStatsSummary) HasDocsExamined() bool`

HasDocsExamined returns a boolean if a field has been set.
### GetDocsExaminedRatio

`func (o *QueryStatsSummary) GetDocsExaminedRatio() float64`

GetDocsExaminedRatio returns the DocsExaminedRatio field if non-nil, zero value otherwise.

### GetDocsExaminedRatioOk

`func (o *QueryStatsSummary) GetDocsExaminedRatioOk() (*float64, bool)`

GetDocsExaminedRatioOk returns a tuple with the DocsExaminedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsExaminedRatio

`func (o *QueryStatsSummary) SetDocsExaminedRatio(v float64)`

SetDocsExaminedRatio sets DocsExaminedRatio field to given value.

### HasDocsExaminedRatio

`func (o *QueryStatsSummary) HasDocsExaminedRatio() bool`

HasDocsExaminedRatio returns a boolean if a field has been set.
### GetDocsReturned

`func (o *QueryStatsSummary) GetDocsReturned() float64`

GetDocsReturned returns the DocsReturned field if non-nil, zero value otherwise.

### GetDocsReturnedOk

`func (o *QueryStatsSummary) GetDocsReturnedOk() (*float64, bool)`

GetDocsReturnedOk returns a tuple with the DocsReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsReturned

`func (o *QueryStatsSummary) SetDocsReturned(v float64)`

SetDocsReturned sets DocsReturned field to given value.

### HasDocsReturned

`func (o *QueryStatsSummary) HasDocsReturned() bool`

HasDocsReturned returns a boolean if a field has been set.
### GetExecCount

`func (o *QueryStatsSummary) GetExecCount() float64`

GetExecCount returns the ExecCount field if non-nil, zero value otherwise.

### GetExecCountOk

`func (o *QueryStatsSummary) GetExecCountOk() (*float64, bool)`

GetExecCountOk returns a tuple with the ExecCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecCount

`func (o *QueryStatsSummary) SetExecCount(v float64)`

SetExecCount sets ExecCount field to given value.

### HasExecCount

`func (o *QueryStatsSummary) HasExecCount() bool`

HasExecCount returns a boolean if a field has been set.
### GetKeysExamined

`func (o *QueryStatsSummary) GetKeysExamined() float64`

GetKeysExamined returns the KeysExamined field if non-nil, zero value otherwise.

### GetKeysExaminedOk

`func (o *QueryStatsSummary) GetKeysExaminedOk() (*float64, bool)`

GetKeysExaminedOk returns a tuple with the KeysExamined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysExamined

`func (o *QueryStatsSummary) SetKeysExamined(v float64)`

SetKeysExamined sets KeysExamined field to given value.

### HasKeysExamined

`func (o *QueryStatsSummary) HasKeysExamined() bool`

HasKeysExamined returns a boolean if a field has been set.
### GetKeysExaminedRatio

`func (o *QueryStatsSummary) GetKeysExaminedRatio() float64`

GetKeysExaminedRatio returns the KeysExaminedRatio field if non-nil, zero value otherwise.

### GetKeysExaminedRatioOk

`func (o *QueryStatsSummary) GetKeysExaminedRatioOk() (*float64, bool)`

GetKeysExaminedRatioOk returns a tuple with the KeysExaminedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysExaminedRatio

`func (o *QueryStatsSummary) SetKeysExaminedRatio(v float64)`

SetKeysExaminedRatio sets KeysExaminedRatio field to given value.

### HasKeysExaminedRatio

`func (o *QueryStatsSummary) HasKeysExaminedRatio() bool`

HasKeysExaminedRatio returns a boolean if a field has been set.
### GetLastExecMicros

`func (o *QueryStatsSummary) GetLastExecMicros() float64`

GetLastExecMicros returns the LastExecMicros field if non-nil, zero value otherwise.

### GetLastExecMicrosOk

`func (o *QueryStatsSummary) GetLastExecMicrosOk() (*float64, bool)`

GetLastExecMicrosOk returns a tuple with the LastExecMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecMicros

`func (o *QueryStatsSummary) SetLastExecMicros(v float64)`

SetLastExecMicros sets LastExecMicros field to given value.

### HasLastExecMicros

`func (o *QueryStatsSummary) HasLastExecMicros() bool`

HasLastExecMicros returns a boolean if a field has been set.
### GetNamespace

`func (o *QueryStatsSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *QueryStatsSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *QueryStatsSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *QueryStatsSummary) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetP50ExecMicros

`func (o *QueryStatsSummary) GetP50ExecMicros() float64`

GetP50ExecMicros returns the P50ExecMicros field if non-nil, zero value otherwise.

### GetP50ExecMicrosOk

`func (o *QueryStatsSummary) GetP50ExecMicrosOk() (*float64, bool)`

GetP50ExecMicrosOk returns a tuple with the P50ExecMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50ExecMicros

`func (o *QueryStatsSummary) SetP50ExecMicros(v float64)`

SetP50ExecMicros sets P50ExecMicros field to given value.

### HasP50ExecMicros

`func (o *QueryStatsSummary) HasP50ExecMicros() bool`

HasP50ExecMicros returns a boolean if a field has been set.
### GetP90ExecMicros

`func (o *QueryStatsSummary) GetP90ExecMicros() float64`

GetP90ExecMicros returns the P90ExecMicros field if non-nil, zero value otherwise.

### GetP90ExecMicrosOk

`func (o *QueryStatsSummary) GetP90ExecMicrosOk() (*float64, bool)`

GetP90ExecMicrosOk returns a tuple with the P90ExecMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90ExecMicros

`func (o *QueryStatsSummary) SetP90ExecMicros(v float64)`

SetP90ExecMicros sets P90ExecMicros field to given value.

### HasP90ExecMicros

`func (o *QueryStatsSummary) HasP90ExecMicros() bool`

HasP90ExecMicros returns a boolean if a field has been set.
### GetP99ExecMicros

`func (o *QueryStatsSummary) GetP99ExecMicros() float64`

GetP99ExecMicros returns the P99ExecMicros field if non-nil, zero value otherwise.

### GetP99ExecMicrosOk

`func (o *QueryStatsSummary) GetP99ExecMicrosOk() (*float64, bool)`

GetP99ExecMicrosOk returns a tuple with the P99ExecMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99ExecMicros

`func (o *QueryStatsSummary) SetP99ExecMicros(v float64)`

SetP99ExecMicros sets P99ExecMicros field to given value.

### HasP99ExecMicros

`func (o *QueryStatsSummary) HasP99ExecMicros() bool`

HasP99ExecMicros returns a boolean if a field has been set.
### GetQueryShape

`func (o *QueryStatsSummary) GetQueryShape() string`

GetQueryShape returns the QueryShape field if non-nil, zero value otherwise.

### GetQueryShapeOk

`func (o *QueryStatsSummary) GetQueryShapeOk() (*string, bool)`

GetQueryShapeOk returns a tuple with the QueryShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryShape

`func (o *QueryStatsSummary) SetQueryShape(v string)`

SetQueryShape sets QueryShape field to given value.

### HasQueryShape

`func (o *QueryStatsSummary) HasQueryShape() bool`

HasQueryShape returns a boolean if a field has been set.
### GetQueryShapeHash

`func (o *QueryStatsSummary) GetQueryShapeHash() string`

GetQueryShapeHash returns the QueryShapeHash field if non-nil, zero value otherwise.

### GetQueryShapeHashOk

`func (o *QueryStatsSummary) GetQueryShapeHashOk() (*string, bool)`

GetQueryShapeHashOk returns a tuple with the QueryShapeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryShapeHash

`func (o *QueryStatsSummary) SetQueryShapeHash(v string)`

SetQueryShapeHash sets QueryShapeHash field to given value.

### HasQueryShapeHash

`func (o *QueryStatsSummary) HasQueryShapeHash() bool`

HasQueryShapeHash returns a boolean if a field has been set.
### GetSystemQuery

`func (o *QueryStatsSummary) GetSystemQuery() bool`

GetSystemQuery returns the SystemQuery field if non-nil, zero value otherwise.

### GetSystemQueryOk

`func (o *QueryStatsSummary) GetSystemQueryOk() (*bool, bool)`

GetSystemQueryOk returns a tuple with the SystemQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemQuery

`func (o *QueryStatsSummary) SetSystemQuery(v bool)`

SetSystemQuery sets SystemQuery field to given value.

### HasSystemQuery

`func (o *QueryStatsSummary) HasSystemQuery() bool`

HasSystemQuery returns a boolean if a field has been set.
### GetTotalTimeToResponseMicros

`func (o *QueryStatsSummary) GetTotalTimeToResponseMicros() float64`

GetTotalTimeToResponseMicros returns the TotalTimeToResponseMicros field if non-nil, zero value otherwise.

### GetTotalTimeToResponseMicrosOk

`func (o *QueryStatsSummary) GetTotalTimeToResponseMicrosOk() (*float64, bool)`

GetTotalTimeToResponseMicrosOk returns a tuple with the TotalTimeToResponseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeToResponseMicros

`func (o *QueryStatsSummary) SetTotalTimeToResponseMicros(v float64)`

SetTotalTimeToResponseMicros sets TotalTimeToResponseMicros field to given value.

### HasTotalTimeToResponseMicros

`func (o *QueryStatsSummary) HasTotalTimeToResponseMicros() bool`

HasTotalTimeToResponseMicros returns a boolean if a field has been set.
### GetTotalWorkingMillis

`func (o *QueryStatsSummary) GetTotalWorkingMillis() float64`

GetTotalWorkingMillis returns the TotalWorkingMillis field if non-nil, zero value otherwise.

### GetTotalWorkingMillisOk

`func (o *QueryStatsSummary) GetTotalWorkingMillisOk() (*float64, bool)`

GetTotalWorkingMillisOk returns a tuple with the TotalWorkingMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkingMillis

`func (o *QueryStatsSummary) SetTotalWorkingMillis(v float64)`

SetTotalWorkingMillis sets TotalWorkingMillis field to given value.

### HasTotalWorkingMillis

`func (o *QueryStatsSummary) HasTotalWorkingMillis() bool`

HasTotalWorkingMillis returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


