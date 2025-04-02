# DropIndexSuggestionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HiddenIndexes** | Pointer to [**[]DropIndexSuggestionsIndex**](DropIndexSuggestionsIndex.md) | List that contains the documents with information about the hidden indexes that the Performance Advisor suggests to remove. | [optional] [readonly] 
**RedundantIndexes** | Pointer to [**[]DropIndexSuggestionsIndex**](DropIndexSuggestionsIndex.md) | List that contains the documents with information about the redundant indexes that the Performance Advisor suggests to remove. | [optional] [readonly] 
**UnusedIndexes** | Pointer to [**[]DropIndexSuggestionsIndex**](DropIndexSuggestionsIndex.md) | List that contains the documents with information about the unused indexes that the Performance Advisor suggests to remove. | [optional] [readonly] 

## Methods

### NewDropIndexSuggestionsResponse

`func NewDropIndexSuggestionsResponse() *DropIndexSuggestionsResponse`

NewDropIndexSuggestionsResponse instantiates a new DropIndexSuggestionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropIndexSuggestionsResponseWithDefaults

`func NewDropIndexSuggestionsResponseWithDefaults() *DropIndexSuggestionsResponse`

NewDropIndexSuggestionsResponseWithDefaults instantiates a new DropIndexSuggestionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHiddenIndexes

`func (o *DropIndexSuggestionsResponse) GetHiddenIndexes() []DropIndexSuggestionsIndex`

GetHiddenIndexes returns the HiddenIndexes field if non-nil, zero value otherwise.

### GetHiddenIndexesOk

`func (o *DropIndexSuggestionsResponse) GetHiddenIndexesOk() (*[]DropIndexSuggestionsIndex, bool)`

GetHiddenIndexesOk returns a tuple with the HiddenIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenIndexes

`func (o *DropIndexSuggestionsResponse) SetHiddenIndexes(v []DropIndexSuggestionsIndex)`

SetHiddenIndexes sets HiddenIndexes field to given value.

### HasHiddenIndexes

`func (o *DropIndexSuggestionsResponse) HasHiddenIndexes() bool`

HasHiddenIndexes returns a boolean if a field has been set.
### GetRedundantIndexes

`func (o *DropIndexSuggestionsResponse) GetRedundantIndexes() []DropIndexSuggestionsIndex`

GetRedundantIndexes returns the RedundantIndexes field if non-nil, zero value otherwise.

### GetRedundantIndexesOk

`func (o *DropIndexSuggestionsResponse) GetRedundantIndexesOk() (*[]DropIndexSuggestionsIndex, bool)`

GetRedundantIndexesOk returns a tuple with the RedundantIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantIndexes

`func (o *DropIndexSuggestionsResponse) SetRedundantIndexes(v []DropIndexSuggestionsIndex)`

SetRedundantIndexes sets RedundantIndexes field to given value.

### HasRedundantIndexes

`func (o *DropIndexSuggestionsResponse) HasRedundantIndexes() bool`

HasRedundantIndexes returns a boolean if a field has been set.
### GetUnusedIndexes

`func (o *DropIndexSuggestionsResponse) GetUnusedIndexes() []DropIndexSuggestionsIndex`

GetUnusedIndexes returns the UnusedIndexes field if non-nil, zero value otherwise.

### GetUnusedIndexesOk

`func (o *DropIndexSuggestionsResponse) GetUnusedIndexesOk() (*[]DropIndexSuggestionsIndex, bool)`

GetUnusedIndexesOk returns a tuple with the UnusedIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedIndexes

`func (o *DropIndexSuggestionsResponse) SetUnusedIndexes(v []DropIndexSuggestionsIndex)`

SetUnusedIndexes sets UnusedIndexes field to given value.

### HasUnusedIndexes

`func (o *DropIndexSuggestionsResponse) HasUnusedIndexes() bool`

HasUnusedIndexes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


