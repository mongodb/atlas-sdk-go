# PerformanceAdvisorSlowQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **string** | Text of the MongoDB log related to this slow query. | [optional] [readonly] 
**Metrics** | Pointer to [**PerformanceAdvisorSlowQueryMetrics**](PerformanceAdvisorSlowQueryMetrics.md) |  | [optional] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 
**OpType** | Pointer to **string** | Operation type (read/write/command) associated with this slow query log. | [optional] [readonly] 
**ReplicaState** | Pointer to **string** | Replica state associated with this slow query log. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorSlowQuery

`func NewPerformanceAdvisorSlowQuery() *PerformanceAdvisorSlowQuery`

NewPerformanceAdvisorSlowQuery instantiates a new PerformanceAdvisorSlowQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorSlowQueryWithDefaults

`func NewPerformanceAdvisorSlowQueryWithDefaults() *PerformanceAdvisorSlowQuery`

NewPerformanceAdvisorSlowQueryWithDefaults instantiates a new PerformanceAdvisorSlowQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *PerformanceAdvisorSlowQuery) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PerformanceAdvisorSlowQuery) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PerformanceAdvisorSlowQuery) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *PerformanceAdvisorSlowQuery) HasLine() bool`

HasLine returns a boolean if a field has been set.
### GetMetrics

`func (o *PerformanceAdvisorSlowQuery) GetMetrics() PerformanceAdvisorSlowQueryMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PerformanceAdvisorSlowQuery) GetMetricsOk() (*PerformanceAdvisorSlowQueryMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PerformanceAdvisorSlowQuery) SetMetrics(v PerformanceAdvisorSlowQueryMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PerformanceAdvisorSlowQuery) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.
### GetNamespace

`func (o *PerformanceAdvisorSlowQuery) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PerformanceAdvisorSlowQuery) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PerformanceAdvisorSlowQuery) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PerformanceAdvisorSlowQuery) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetOpType

`func (o *PerformanceAdvisorSlowQuery) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *PerformanceAdvisorSlowQuery) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *PerformanceAdvisorSlowQuery) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *PerformanceAdvisorSlowQuery) HasOpType() bool`

HasOpType returns a boolean if a field has been set.
### GetReplicaState

`func (o *PerformanceAdvisorSlowQuery) GetReplicaState() string`

GetReplicaState returns the ReplicaState field if non-nil, zero value otherwise.

### GetReplicaStateOk

`func (o *PerformanceAdvisorSlowQuery) GetReplicaStateOk() (*string, bool)`

GetReplicaStateOk returns a tuple with the ReplicaState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaState

`func (o *PerformanceAdvisorSlowQuery) SetReplicaState(v string)`

SetReplicaState sets ReplicaState field to given value.

### HasReplicaState

`func (o *PerformanceAdvisorSlowQuery) HasReplicaState() bool`

HasReplicaState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


