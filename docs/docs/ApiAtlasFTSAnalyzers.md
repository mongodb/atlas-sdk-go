# ApiAtlasFTSAnalyzers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharFilters** | Pointer to [**[]ApiAtlasFTSAnalyzersCharFiltersInner**](ApiAtlasFTSAnalyzersCharFiltersInner.md) | Filters that examine text one character at a time and perform filtering operations. | [optional] 
**Name** | **string** | Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - &#x60;lucene.&#x60; - &#x60;builtin.&#x60; - &#x60;mongodb.&#x60; | 
**TokenFilters** | Pointer to [**[]ApiAtlasFTSAnalyzersTokenFiltersInner**](ApiAtlasFTSAnalyzersTokenFiltersInner.md) | Filter that performs operations such as:  - Stemming, which reduces related words, such as \&quot;talking\&quot;, \&quot;talked\&quot;, and \&quot;talks\&quot; to their root word \&quot;talk\&quot;.  - Redaction, the removal of sensitive information from public documents. | [optional] 
**Tokenizer** | [**ApiAtlasFTSAnalyzersTokenizer**](ApiAtlasFTSAnalyzersTokenizer.md) |  | 

## Methods

### NewApiAtlasFTSAnalyzers

`func NewApiAtlasFTSAnalyzers(name string, tokenizer ApiAtlasFTSAnalyzersTokenizer, ) *ApiAtlasFTSAnalyzers`

NewApiAtlasFTSAnalyzers instantiates a new ApiAtlasFTSAnalyzers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersWithDefaults

`func NewApiAtlasFTSAnalyzersWithDefaults() *ApiAtlasFTSAnalyzers`

NewApiAtlasFTSAnalyzersWithDefaults instantiates a new ApiAtlasFTSAnalyzers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharFilters

`func (o *ApiAtlasFTSAnalyzers) GetCharFilters() []ApiAtlasFTSAnalyzersCharFiltersInner`

GetCharFilters returns the CharFilters field if non-nil, zero value otherwise.

### GetCharFiltersOk

`func (o *ApiAtlasFTSAnalyzers) GetCharFiltersOk() (*[]ApiAtlasFTSAnalyzersCharFiltersInner, bool)`

GetCharFiltersOk returns a tuple with the CharFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharFilters

`func (o *ApiAtlasFTSAnalyzers) SetCharFilters(v []ApiAtlasFTSAnalyzersCharFiltersInner)`

SetCharFilters sets CharFilters field to given value.

### HasCharFilters

`func (o *ApiAtlasFTSAnalyzers) HasCharFilters() bool`

HasCharFilters returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasFTSAnalyzers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasFTSAnalyzers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasFTSAnalyzers) SetName(v string)`

SetName sets Name field to given value.


### GetTokenFilters

`func (o *ApiAtlasFTSAnalyzers) GetTokenFilters() []ApiAtlasFTSAnalyzersTokenFiltersInner`

GetTokenFilters returns the TokenFilters field if non-nil, zero value otherwise.

### GetTokenFiltersOk

`func (o *ApiAtlasFTSAnalyzers) GetTokenFiltersOk() (*[]ApiAtlasFTSAnalyzersTokenFiltersInner, bool)`

GetTokenFiltersOk returns a tuple with the TokenFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFilters

`func (o *ApiAtlasFTSAnalyzers) SetTokenFilters(v []ApiAtlasFTSAnalyzersTokenFiltersInner)`

SetTokenFilters sets TokenFilters field to given value.

### HasTokenFilters

`func (o *ApiAtlasFTSAnalyzers) HasTokenFilters() bool`

HasTokenFilters returns a boolean if a field has been set.

### GetTokenizer

`func (o *ApiAtlasFTSAnalyzers) GetTokenizer() ApiAtlasFTSAnalyzersTokenizer`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *ApiAtlasFTSAnalyzers) GetTokenizerOk() (*ApiAtlasFTSAnalyzersTokenizer, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *ApiAtlasFTSAnalyzers) SetTokenizer(v ApiAtlasFTSAnalyzersTokenizer)`

SetTokenizer sets Tokenizer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


