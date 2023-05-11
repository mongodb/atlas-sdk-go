# TokenFilternGram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | **int** | Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**. | 
**MinGram** | **int** | Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**. | 
**TermNotInBounds** | Pointer to **string** | Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**. | [optional] [default to "omit"]
**Type** | **string** | Human-readable label that identifies this token filter type. | 

## Methods

### NewTokenFilternGram

`func NewTokenFilternGram(maxGram int, minGram int, type_ string, ) *TokenFilternGram`

NewTokenFilternGram instantiates a new TokenFilternGram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFilternGramWithDefaults

`func NewTokenFilternGramWithDefaults() *TokenFilternGram`

NewTokenFilternGramWithDefaults instantiates a new TokenFilternGram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *TokenFilternGram) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *TokenFilternGram) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *TokenFilternGram) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *TokenFilternGram) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *TokenFilternGram) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *TokenFilternGram) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.


### GetTermNotInBounds

`func (o *TokenFilternGram) GetTermNotInBounds() string`

GetTermNotInBounds returns the TermNotInBounds field if non-nil, zero value otherwise.

### GetTermNotInBoundsOk

`func (o *TokenFilternGram) GetTermNotInBoundsOk() (*string, bool)`

GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotInBounds

`func (o *TokenFilternGram) SetTermNotInBounds(v string)`

SetTermNotInBounds sets TermNotInBounds field to given value.

### HasTermNotInBounds

`func (o *TokenFilternGram) HasTermNotInBounds() bool`

HasTermNotInBounds returns a boolean if a field has been set.

### GetType

`func (o *TokenFilternGram) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenFilternGram) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenFilternGram) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


