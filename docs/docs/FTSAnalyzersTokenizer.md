# FTSAnalyzersTokenizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | **int** | Characters to include in the longest token that Atlas Search creates. | 
**MinGram** | **int** | Characters to include in the shortest token that Atlas Search creates. | 
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 
**Group** | **int** | Index of the character group within the matching expression to extract into tokens. Use &#x60;0&#x60; to extract all character groups. | 
**Pattern** | **string** | Regular expression to match against. | 
**MaxTokenLength** | Pointer to **int** | Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens. | [optional] [default to 255]

## Methods

### NewFTSAnalyzersTokenizer

`func NewFTSAnalyzersTokenizer(maxGram int, minGram int, type_ string, group int, pattern string, ) *FTSAnalyzersTokenizer`

NewFTSAnalyzersTokenizer instantiates a new FTSAnalyzersTokenizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSAnalyzersTokenizerWithDefaults

`func NewFTSAnalyzersTokenizerWithDefaults() *FTSAnalyzersTokenizer`

NewFTSAnalyzersTokenizerWithDefaults instantiates a new FTSAnalyzersTokenizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *FTSAnalyzersTokenizer) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *FTSAnalyzersTokenizer) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *FTSAnalyzersTokenizer) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *FTSAnalyzersTokenizer) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *FTSAnalyzersTokenizer) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *FTSAnalyzersTokenizer) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.


### GetType

`func (o *FTSAnalyzersTokenizer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FTSAnalyzersTokenizer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FTSAnalyzersTokenizer) SetType(v string)`

SetType sets Type field to given value.


### GetGroup

`func (o *FTSAnalyzersTokenizer) GetGroup() int`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FTSAnalyzersTokenizer) GetGroupOk() (*int, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FTSAnalyzersTokenizer) SetGroup(v int)`

SetGroup sets Group field to given value.


### GetPattern

`func (o *FTSAnalyzersTokenizer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *FTSAnalyzersTokenizer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *FTSAnalyzersTokenizer) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMaxTokenLength

`func (o *FTSAnalyzersTokenizer) GetMaxTokenLength() int`

GetMaxTokenLength returns the MaxTokenLength field if non-nil, zero value otherwise.

### GetMaxTokenLengthOk

`func (o *FTSAnalyzersTokenizer) GetMaxTokenLengthOk() (*int, bool)`

GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenLength

`func (o *FTSAnalyzersTokenizer) SetMaxTokenLength(v int)`

SetMaxTokenLength sets MaxTokenLength field to given value.

### HasMaxTokenLength

`func (o *FTSAnalyzersTokenizer) HasMaxTokenLength() bool`

HasMaxTokenLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


