# PerformanceAdvisorIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgObjSize** | Pointer to **float64** | The average size of an object in the collection of this index. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this index. | [optional] [readonly] 
**Impact** | Pointer to **[]string** | List that contains unique 24-hexadecimal character string that identifies the query shapes in this response that the Performance Advisor suggests. | [optional] [readonly] 
**Index** | Pointer to **[]map[string]int** | List that contains documents that specify a key in the index and its sort order. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 
**Weight** | Pointer to **float64** | Estimated performance improvement that the suggested index provides. This value corresponds to **Impact** in the Performance Advisor user interface. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorIndex

`func NewPerformanceAdvisorIndex() *PerformanceAdvisorIndex`

NewPerformanceAdvisorIndex instantiates a new PerformanceAdvisorIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorIndexWithDefaults

`func NewPerformanceAdvisorIndexWithDefaults() *PerformanceAdvisorIndex`

NewPerformanceAdvisorIndexWithDefaults instantiates a new PerformanceAdvisorIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgObjSize

`func (o *PerformanceAdvisorIndex) GetAvgObjSize() float64`

GetAvgObjSize returns the AvgObjSize field if non-nil, zero value otherwise.

### GetAvgObjSizeOk

`func (o *PerformanceAdvisorIndex) GetAvgObjSizeOk() (*float64, bool)`

GetAvgObjSizeOk returns a tuple with the AvgObjSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgObjSize

`func (o *PerformanceAdvisorIndex) SetAvgObjSize(v float64)`

SetAvgObjSize sets AvgObjSize field to given value.

### HasAvgObjSize

`func (o *PerformanceAdvisorIndex) HasAvgObjSize() bool`

HasAvgObjSize returns a boolean if a field has been set.
### GetId

`func (o *PerformanceAdvisorIndex) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerformanceAdvisorIndex) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerformanceAdvisorIndex) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PerformanceAdvisorIndex) HasId() bool`

HasId returns a boolean if a field has been set.
### GetImpact

`func (o *PerformanceAdvisorIndex) GetImpact() []string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *PerformanceAdvisorIndex) GetImpactOk() (*[]string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *PerformanceAdvisorIndex) SetImpact(v []string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *PerformanceAdvisorIndex) HasImpact() bool`

HasImpact returns a boolean if a field has been set.
### GetIndex

`func (o *PerformanceAdvisorIndex) GetIndex() []map[string]int`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PerformanceAdvisorIndex) GetIndexOk() (*[]map[string]int, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PerformanceAdvisorIndex) SetIndex(v []map[string]int)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PerformanceAdvisorIndex) HasIndex() bool`

HasIndex returns a boolean if a field has been set.
### GetNamespace

`func (o *PerformanceAdvisorIndex) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PerformanceAdvisorIndex) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PerformanceAdvisorIndex) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PerformanceAdvisorIndex) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetWeight

`func (o *PerformanceAdvisorIndex) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PerformanceAdvisorIndex) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PerformanceAdvisorIndex) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PerformanceAdvisorIndex) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


