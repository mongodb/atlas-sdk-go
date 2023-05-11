# PerformanceAdvisorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shapes** | Pointer to [**[]PerformanceAdvisorShape**](PerformanceAdvisorShape.md) | List of query predicates, sorts, and projections that the Performance Advisor suggests. | [optional] [readonly] 
**SuggestedIndexes** | Pointer to [**[]PerformanceAdvisorIndex**](PerformanceAdvisorIndex.md) | List that contains the documents with information about the indexes that the Performance Advisor suggests. | [optional] [readonly] 

## Methods

### NewPerformanceAdvisorResponse

`func NewPerformanceAdvisorResponse() *PerformanceAdvisorResponse`

NewPerformanceAdvisorResponse instantiates a new PerformanceAdvisorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceAdvisorResponseWithDefaults

`func NewPerformanceAdvisorResponseWithDefaults() *PerformanceAdvisorResponse`

NewPerformanceAdvisorResponseWithDefaults instantiates a new PerformanceAdvisorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShapes

`func (o *PerformanceAdvisorResponse) GetShapes() []PerformanceAdvisorShape`

GetShapes returns the Shapes field if non-nil, zero value otherwise.

### GetShapesOk

`func (o *PerformanceAdvisorResponse) GetShapesOk() (*[]PerformanceAdvisorShape, bool)`

GetShapesOk returns a tuple with the Shapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapes

`func (o *PerformanceAdvisorResponse) SetShapes(v []PerformanceAdvisorShape)`

SetShapes sets Shapes field to given value.

### HasShapes

`func (o *PerformanceAdvisorResponse) HasShapes() bool`

HasShapes returns a boolean if a field has been set.

### GetSuggestedIndexes

`func (o *PerformanceAdvisorResponse) GetSuggestedIndexes() []PerformanceAdvisorIndex`

GetSuggestedIndexes returns the SuggestedIndexes field if non-nil, zero value otherwise.

### GetSuggestedIndexesOk

`func (o *PerformanceAdvisorResponse) GetSuggestedIndexesOk() (*[]PerformanceAdvisorIndex, bool)`

GetSuggestedIndexesOk returns a tuple with the SuggestedIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedIndexes

`func (o *PerformanceAdvisorResponse) SetSuggestedIndexes(v []PerformanceAdvisorIndex)`

SetSuggestedIndexes sets SuggestedIndexes field to given value.

### HasSuggestedIndexes

`func (o *PerformanceAdvisorResponse) HasSuggestedIndexes() bool`

HasSuggestedIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


