# PerformanceAdvisorSlowQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **string** | Text of the MongoDB log related to this slow query. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


