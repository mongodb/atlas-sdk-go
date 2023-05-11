# FTSAnalyzersTokenFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalTokens** | Pointer to **string** | Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose &#x60;include&#x60; if you want to support queries on both the original tokens as well as the converted forms.   Choose &#x60;omit&#x60; if you want to query only on the converted forms of the original tokens. | [optional] [default to "include"]
**Type** | **string** | Human-readable label that identifies this token filter type. | 
**MaxGram** | **int** | Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**. | 
**MinGram** | **int** | Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**. | 
**TermNotInBounds** | Pointer to **string** | Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**. | [optional] [default to "omit"]
**NormalizationForm** | Pointer to **string** | Normalization form to apply. | [optional] [default to "nfc"]
**Max** | Pointer to **int** | Number that specifies the maximum length of a token. Value must be greater than or equal to **min**. | [optional] [default to 255]
**Min** | Pointer to **int** | Number that specifies the minimum length of a token. This value must be less than or equal to **max**. | [optional] [default to 0]
**Matches** | **string** | Value that indicates whether to replace only the first matching pattern or all matching patterns. | 
**Pattern** | **string** | Regular expression pattern to apply to each token. | 
**Replacement** | **string** | Replacement string to substitute wherever a matching pattern occurs. | 
**MaxShingleSize** | **int** | Value that specifies the maximum number of tokens per shingle. This value must be greater than or equal to **minShingleSize**. | 
**MinShingleSize** | **int** | Value that specifies the minimum number of tokens per shingle. This value must be less than or equal to **maxShingleSize**. | 
**StemmerName** | **string** | Snowball-generated stemmer to use. | 
**IgnoreCase** | Pointer to **bool** | Flag that indicates whether to ignore the case of stop words when filtering the tokens to remove. | [optional] [default to true]
**Tokens** | **[]string** | The stop words that correspond to the tokens to remove. Value must be one or more stop words. | 

## Methods

### NewFTSAnalyzersTokenFiltersInner

`func NewFTSAnalyzersTokenFiltersInner(type_ string, maxGram int, minGram int, matches string, pattern string, replacement string, maxShingleSize int, minShingleSize int, stemmerName string, tokens []string, ) *FTSAnalyzersTokenFiltersInner`

NewFTSAnalyzersTokenFiltersInner instantiates a new FTSAnalyzersTokenFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSAnalyzersTokenFiltersInnerWithDefaults

`func NewFTSAnalyzersTokenFiltersInnerWithDefaults() *FTSAnalyzersTokenFiltersInner`

NewFTSAnalyzersTokenFiltersInnerWithDefaults instantiates a new FTSAnalyzersTokenFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalTokens

`func (o *FTSAnalyzersTokenFiltersInner) GetOriginalTokens() string`

GetOriginalTokens returns the OriginalTokens field if non-nil, zero value otherwise.

### GetOriginalTokensOk

`func (o *FTSAnalyzersTokenFiltersInner) GetOriginalTokensOk() (*string, bool)`

GetOriginalTokensOk returns a tuple with the OriginalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTokens

`func (o *FTSAnalyzersTokenFiltersInner) SetOriginalTokens(v string)`

SetOriginalTokens sets OriginalTokens field to given value.

### HasOriginalTokens

`func (o *FTSAnalyzersTokenFiltersInner) HasOriginalTokens() bool`

HasOriginalTokens returns a boolean if a field has been set.

### GetType

`func (o *FTSAnalyzersTokenFiltersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FTSAnalyzersTokenFiltersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FTSAnalyzersTokenFiltersInner) SetType(v string)`

SetType sets Type field to given value.


### GetMaxGram

`func (o *FTSAnalyzersTokenFiltersInner) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *FTSAnalyzersTokenFiltersInner) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *FTSAnalyzersTokenFiltersInner) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *FTSAnalyzersTokenFiltersInner) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.


### GetTermNotInBounds

`func (o *FTSAnalyzersTokenFiltersInner) GetTermNotInBounds() string`

GetTermNotInBounds returns the TermNotInBounds field if non-nil, zero value otherwise.

### GetTermNotInBoundsOk

`func (o *FTSAnalyzersTokenFiltersInner) GetTermNotInBoundsOk() (*string, bool)`

GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotInBounds

`func (o *FTSAnalyzersTokenFiltersInner) SetTermNotInBounds(v string)`

SetTermNotInBounds sets TermNotInBounds field to given value.

### HasTermNotInBounds

`func (o *FTSAnalyzersTokenFiltersInner) HasTermNotInBounds() bool`

HasTermNotInBounds returns a boolean if a field has been set.

### GetNormalizationForm

`func (o *FTSAnalyzersTokenFiltersInner) GetNormalizationForm() string`

GetNormalizationForm returns the NormalizationForm field if non-nil, zero value otherwise.

### GetNormalizationFormOk

`func (o *FTSAnalyzersTokenFiltersInner) GetNormalizationFormOk() (*string, bool)`

GetNormalizationFormOk returns a tuple with the NormalizationForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationForm

`func (o *FTSAnalyzersTokenFiltersInner) SetNormalizationForm(v string)`

SetNormalizationForm sets NormalizationForm field to given value.

### HasNormalizationForm

`func (o *FTSAnalyzersTokenFiltersInner) HasNormalizationForm() bool`

HasNormalizationForm returns a boolean if a field has been set.

### GetMax

`func (o *FTSAnalyzersTokenFiltersInner) GetMax() int`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMaxOk() (*int, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FTSAnalyzersTokenFiltersInner) SetMax(v int)`

SetMax sets Max field to given value.

### HasMax

`func (o *FTSAnalyzersTokenFiltersInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *FTSAnalyzersTokenFiltersInner) GetMin() int`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMinOk() (*int, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *FTSAnalyzersTokenFiltersInner) SetMin(v int)`

SetMin sets Min field to given value.

### HasMin

`func (o *FTSAnalyzersTokenFiltersInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMatches

`func (o *FTSAnalyzersTokenFiltersInner) GetMatches() string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMatchesOk() (*string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *FTSAnalyzersTokenFiltersInner) SetMatches(v string)`

SetMatches sets Matches field to given value.


### GetPattern

`func (o *FTSAnalyzersTokenFiltersInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *FTSAnalyzersTokenFiltersInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *FTSAnalyzersTokenFiltersInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReplacement

`func (o *FTSAnalyzersTokenFiltersInner) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *FTSAnalyzersTokenFiltersInner) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *FTSAnalyzersTokenFiltersInner) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.


### GetMaxShingleSize

`func (o *FTSAnalyzersTokenFiltersInner) GetMaxShingleSize() int`

GetMaxShingleSize returns the MaxShingleSize field if non-nil, zero value otherwise.

### GetMaxShingleSizeOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMaxShingleSizeOk() (*int, bool)`

GetMaxShingleSizeOk returns a tuple with the MaxShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxShingleSize

`func (o *FTSAnalyzersTokenFiltersInner) SetMaxShingleSize(v int)`

SetMaxShingleSize sets MaxShingleSize field to given value.


### GetMinShingleSize

`func (o *FTSAnalyzersTokenFiltersInner) GetMinShingleSize() int`

GetMinShingleSize returns the MinShingleSize field if non-nil, zero value otherwise.

### GetMinShingleSizeOk

`func (o *FTSAnalyzersTokenFiltersInner) GetMinShingleSizeOk() (*int, bool)`

GetMinShingleSizeOk returns a tuple with the MinShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinShingleSize

`func (o *FTSAnalyzersTokenFiltersInner) SetMinShingleSize(v int)`

SetMinShingleSize sets MinShingleSize field to given value.


### GetStemmerName

`func (o *FTSAnalyzersTokenFiltersInner) GetStemmerName() string`

GetStemmerName returns the StemmerName field if non-nil, zero value otherwise.

### GetStemmerNameOk

`func (o *FTSAnalyzersTokenFiltersInner) GetStemmerNameOk() (*string, bool)`

GetStemmerNameOk returns a tuple with the StemmerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStemmerName

`func (o *FTSAnalyzersTokenFiltersInner) SetStemmerName(v string)`

SetStemmerName sets StemmerName field to given value.


### GetIgnoreCase

`func (o *FTSAnalyzersTokenFiltersInner) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *FTSAnalyzersTokenFiltersInner) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *FTSAnalyzersTokenFiltersInner) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *FTSAnalyzersTokenFiltersInner) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetTokens

`func (o *FTSAnalyzersTokenFiltersInner) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *FTSAnalyzersTokenFiltersInner) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *FTSAnalyzersTokenFiltersInner) SetTokens(v []string)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


