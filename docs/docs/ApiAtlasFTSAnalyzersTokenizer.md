# ApiAtlasFTSAnalyzersTokenizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | Pointer to **int** | Characters to include in the longest token that Atlas Search creates. | [optional] 
**MinGram** | Pointer to **int** | Characters to include in the shortest token that Atlas Search creates. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies this tokenizer type. | [optional] 
**Group** | Pointer to **int** | Index of the character group within the matching expression to extract into tokens. Use &#x60;0&#x60; to extract all character groups. | [optional] 
**Pattern** | Pointer to **string** | Regular expression to match against. | [optional] 
**MaxTokenLength** | Pointer to **int** | Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens. | [optional] [default to 255]

## Methods

### NewApiAtlasFTSAnalyzersTokenizer

`func NewApiAtlasFTSAnalyzersTokenizer() *ApiAtlasFTSAnalyzersTokenizer`

NewApiAtlasFTSAnalyzersTokenizer instantiates a new ApiAtlasFTSAnalyzersTokenizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersTokenizerWithDefaults

`func NewApiAtlasFTSAnalyzersTokenizerWithDefaults() *ApiAtlasFTSAnalyzersTokenizer`

NewApiAtlasFTSAnalyzersTokenizerWithDefaults instantiates a new ApiAtlasFTSAnalyzersTokenizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.

### HasMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasMaxGram() bool`

HasMaxGram returns a boolean if a field has been set.

### GetMinGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.

### HasMinGram

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasMinGram() bool`

HasMinGram returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetGroup() int`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetGroupOk() (*int, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetGroup(v int)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPattern

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMaxTokenLength() int`

GetMaxTokenLength returns the MaxTokenLength field if non-nil, zero value otherwise.

### GetMaxTokenLengthOk

`func (o *ApiAtlasFTSAnalyzersTokenizer) GetMaxTokenLengthOk() (*int, bool)`

GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersTokenizer) SetMaxTokenLength(v int)`

SetMaxTokenLength sets MaxTokenLength field to given value.

### HasMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersTokenizer) HasMaxTokenLength() bool`

HasMaxTokenLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


