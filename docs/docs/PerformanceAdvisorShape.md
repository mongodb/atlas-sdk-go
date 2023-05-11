# PerformanceAdvisorShape

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgMs** | Pointer to **int64** | Average duration in milliseconds for the queries examined that match this shape. | [optional] [readonly] 
**Count** | Pointer to **int64** | Number of queries examined that match this shape. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this shape. This string exists only for the duration of this API request. | [optional] [readonly] 
**InefficiencyScore** | Pointer to **int64** | Average number of documents read for every document that the query returns. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 
**Operations** | Pointer to [**[]PerformanceAdvisorOperation**](PerformanceAdvisorOperation.md) | List that contains specific about individual queries. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorShape

`func NewPerformanceAdvisorShape() *PerformanceAdvisorShape`

NewPerformanceAdvisorShape instantiates a new PerformanceAdvisorShape object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorShapeWithDefaults

`func NewPerformanceAdvisorShapeWithDefaults() *PerformanceAdvisorShape`

NewPerformanceAdvisorShapeWithDefaults instantiates a new PerformanceAdvisorShape object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgMs

`func (o *PerformanceAdvisorShape) GetAvgMs() int64`

GetAvgMs returns the AvgMs field if non-nil, zero value otherwise.

### GetAvgMsOk

`func (o *PerformanceAdvisorShape) GetAvgMsOk() (*int64, bool)`

GetAvgMsOk returns a tuple with the AvgMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMs

`func (o *PerformanceAdvisorShape) SetAvgMs(v int64)`

SetAvgMs sets AvgMs field to given value.

### HasAvgMs

`func (o *PerformanceAdvisorShape) HasAvgMs() bool`

HasAvgMs returns a boolean if a field has been set.

### GetCount

`func (o *PerformanceAdvisorShape) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PerformanceAdvisorShape) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PerformanceAdvisorShape) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PerformanceAdvisorShape) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *PerformanceAdvisorShape) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerformanceAdvisorShape) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerformanceAdvisorShape) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PerformanceAdvisorShape) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInefficiencyScore

`func (o *PerformanceAdvisorShape) GetInefficiencyScore() int64`

GetInefficiencyScore returns the InefficiencyScore field if non-nil, zero value otherwise.

### GetInefficiencyScoreOk

`func (o *PerformanceAdvisorShape) GetInefficiencyScoreOk() (*int64, bool)`

GetInefficiencyScoreOk returns a tuple with the InefficiencyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInefficiencyScore

`func (o *PerformanceAdvisorShape) SetInefficiencyScore(v int64)`

SetInefficiencyScore sets InefficiencyScore field to given value.

### HasInefficiencyScore

`func (o *PerformanceAdvisorShape) HasInefficiencyScore() bool`

HasInefficiencyScore returns a boolean if a field has been set.

### GetNamespace

`func (o *PerformanceAdvisorShape) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PerformanceAdvisorShape) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PerformanceAdvisorShape) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PerformanceAdvisorShape) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperations

`func (o *PerformanceAdvisorShape) GetOperations() []PerformanceAdvisorOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PerformanceAdvisorShape) GetOperationsOk() (*[]PerformanceAdvisorOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PerformanceAdvisorShape) SetOperations(v []PerformanceAdvisorOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *PerformanceAdvisorShape) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


