# PerformanceAdvisorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Explanation of Performance Advisor availability limitations on the current tier. | [optional] [readonly] 
**Scale** | Pointer to [**ApiClusterScale**](ApiClusterScale.md) |  | [optional] 
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

### GetMessage

`func (o *PerformanceAdvisorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PerformanceAdvisorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PerformanceAdvisorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PerformanceAdvisorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PerformanceAdvisorResponse) SetMessageNil()`

SetMessageNil sets Message to an explicit JSON null when marshaled, overriding any value previously set with SetMessage. Calling SetMessage again clears the null override.

### GetScale

`func (o *PerformanceAdvisorResponse) GetScale() ApiClusterScale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *PerformanceAdvisorResponse) GetScaleOk() (*ApiClusterScale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *PerformanceAdvisorResponse) SetScale(v ApiClusterScale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *PerformanceAdvisorResponse) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScaleNil

`func (o *PerformanceAdvisorResponse) SetScaleNil()`

SetScaleNil sets Scale to an explicit JSON null when marshaled, overriding any value previously set with SetScale. Calling SetScale again clears the null override.

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

### SetShapesNil

`func (o *PerformanceAdvisorResponse) SetShapesNil()`

SetShapesNil sets Shapes to an explicit JSON null when marshaled, overriding any value previously set with SetShapes. Calling SetShapes again clears the null override.

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

### SetSuggestedIndexesNil

`func (o *PerformanceAdvisorResponse) SetSuggestedIndexesNil()`

SetSuggestedIndexesNil sets SuggestedIndexes to an explicit JSON null when marshaled, overriding any value previously set with SetSuggestedIndexes. Calling SetSuggestedIndexes again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


