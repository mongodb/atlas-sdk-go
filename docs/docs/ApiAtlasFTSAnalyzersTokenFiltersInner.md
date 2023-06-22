# ApiAtlasFTSAnalyzersTokenFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalTokens** | Pointer to **string** | Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose &#x60;include&#x60; if you want to support queries on both the original tokens as well as the converted forms.   Choose &#x60;omit&#x60; if you want to query only on the converted forms of the original tokens. | [optional] [default to "include"]
**Type** | Pointer to **string** | Human-readable label that identifies this token filter type. | [optional] 
**MaxGram** | Pointer to **int** | Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**. | [optional] 
**MinGram** | Pointer to **int** | Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**. | [optional] 
**TermNotInBounds** | Pointer to **string** | Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**. | [optional] [default to "omit"]
**NormalizationForm** | Pointer to **string** | Normalization form to apply. | [optional] [default to "nfc"]
**Max** | Pointer to **int** | Number that specifies the maximum length of a token. Value must be greater than or equal to **min**. | [optional] [default to 255]
**Min** | Pointer to **int** | Number that specifies the minimum length of a token. This value must be less than or equal to **max**. | [optional] [default to 0]
**Matches** | Pointer to **string** | Value that indicates whether to replace only the first matching pattern or all matching patterns. | [optional] 
**Pattern** | Pointer to **string** | Regular expression pattern to apply to each token. | [optional] 
**Replacement** | Pointer to **string** | Replacement string to substitute wherever a matching pattern occurs. | [optional] 
**MaxShingleSize** | Pointer to **int** | Value that specifies the maximum number of tokens per shingle. This value must be greater than or equal to **minShingleSize**. | [optional] 
**MinShingleSize** | Pointer to **int** | Value that specifies the minimum number of tokens per shingle. This value must be less than or equal to **maxShingleSize**. | [optional] 
**StemmerName** | Pointer to **string** | Snowball-generated stemmer to use. | [optional] 
**IgnoreCase** | Pointer to **bool** | Flag that indicates whether to ignore the case of stop words when filtering the tokens to remove. | [optional] [default to true]
**Tokens** | Pointer to **[]string** | The stop words that correspond to the tokens to remove. Value must be one or more stop words. | [optional] 

## Methods

### NewApiAtlasFTSAnalyzersTokenFiltersInner

`func NewApiAtlasFTSAnalyzersTokenFiltersInner() *ApiAtlasFTSAnalyzersTokenFiltersInner`

NewApiAtlasFTSAnalyzersTokenFiltersInner instantiates a new ApiAtlasFTSAnalyzersTokenFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersTokenFiltersInnerWithDefaults

`func NewApiAtlasFTSAnalyzersTokenFiltersInnerWithDefaults() *ApiAtlasFTSAnalyzersTokenFiltersInner`

NewApiAtlasFTSAnalyzersTokenFiltersInnerWithDefaults instantiates a new ApiAtlasFTSAnalyzersTokenFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetOriginalTokens() string`

GetOriginalTokens returns the OriginalTokens field if non-nil, zero value otherwise.

### GetOriginalTokensOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetOriginalTokensOk() (*string, bool)`

GetOriginalTokensOk returns a tuple with the OriginalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetOriginalTokens(v string)`

SetOriginalTokens sets OriginalTokens field to given value.

### HasOriginalTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasOriginalTokens() bool`

HasOriginalTokens returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMaxGram() int`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMaxGramOk() (*int, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMaxGram(v int)`

SetMaxGram sets MaxGram field to given value.

### HasMaxGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMaxGram() bool`

HasMaxGram returns a boolean if a field has been set.

### GetMinGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMinGram() int`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMinGramOk() (*int, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMinGram(v int)`

SetMinGram sets MinGram field to given value.

### HasMinGram

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMinGram() bool`

HasMinGram returns a boolean if a field has been set.

### GetTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetTermNotInBounds() string`

GetTermNotInBounds returns the TermNotInBounds field if non-nil, zero value otherwise.

### GetTermNotInBoundsOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetTermNotInBoundsOk() (*string, bool)`

GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetTermNotInBounds(v string)`

SetTermNotInBounds sets TermNotInBounds field to given value.

### HasTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasTermNotInBounds() bool`

HasTermNotInBounds returns a boolean if a field has been set.

### GetNormalizationForm

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetNormalizationForm() string`

GetNormalizationForm returns the NormalizationForm field if non-nil, zero value otherwise.

### GetNormalizationFormOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetNormalizationFormOk() (*string, bool)`

GetNormalizationFormOk returns a tuple with the NormalizationForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationForm

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetNormalizationForm(v string)`

SetNormalizationForm sets NormalizationForm field to given value.

### HasNormalizationForm

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasNormalizationForm() bool`

HasNormalizationForm returns a boolean if a field has been set.

### GetMax

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMax() int`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMaxOk() (*int, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMax(v int)`

SetMax sets Max field to given value.

### HasMax

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMin() int`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMinOk() (*int, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMin(v int)`

SetMin sets Min field to given value.

### HasMin

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMatches

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMatches() string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMatchesOk() (*string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMatches(v string)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetPattern

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetReplacement

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetMaxShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMaxShingleSize() int`

GetMaxShingleSize returns the MaxShingleSize field if non-nil, zero value otherwise.

### GetMaxShingleSizeOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMaxShingleSizeOk() (*int, bool)`

GetMaxShingleSizeOk returns a tuple with the MaxShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMaxShingleSize(v int)`

SetMaxShingleSize sets MaxShingleSize field to given value.

### HasMaxShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMaxShingleSize() bool`

HasMaxShingleSize returns a boolean if a field has been set.

### GetMinShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMinShingleSize() int`

GetMinShingleSize returns the MinShingleSize field if non-nil, zero value otherwise.

### GetMinShingleSizeOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetMinShingleSizeOk() (*int, bool)`

GetMinShingleSizeOk returns a tuple with the MinShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetMinShingleSize(v int)`

SetMinShingleSize sets MinShingleSize field to given value.

### HasMinShingleSize

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasMinShingleSize() bool`

HasMinShingleSize returns a boolean if a field has been set.

### GetStemmerName

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetStemmerName() string`

GetStemmerName returns the StemmerName field if non-nil, zero value otherwise.

### GetStemmerNameOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetStemmerNameOk() (*string, bool)`

GetStemmerNameOk returns a tuple with the StemmerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStemmerName

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetStemmerName(v string)`

SetStemmerName sets StemmerName field to given value.

### HasStemmerName

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasStemmerName() bool`

HasStemmerName returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ApiAtlasFTSAnalyzersTokenFiltersInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


