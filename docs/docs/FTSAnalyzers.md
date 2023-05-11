# FTSAnalyzers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharFilters** | Pointer to [**[]FTSAnalyzersCharFiltersInner**](FTSAnalyzersCharFiltersInner.md) | Filters that examine text one character at a time and perform filtering operations. | [optional] 
**Name** | **string** | Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - &#x60;lucene.&#x60; - &#x60;builtin.&#x60; - &#x60;mongodb.&#x60; | 
**TokenFilters** | Pointer to [**[]FTSAnalyzersTokenFiltersInner**](FTSAnalyzersTokenFiltersInner.md) | Filter that performs operations such as:  - Stemming, which reduces related words, such as \&quot;talking\&quot;, \&quot;talked\&quot;, and \&quot;talks\&quot; to their root word \&quot;talk\&quot;.  - Redaction, the removal of sensitive information from public documents. | [optional] 
**Tokenizer** | [**FTSAnalyzersTokenizer**](FTSAnalyzersTokenizer.md) |  | 

## Methods

### NewFTSAnalyzers

`func NewFTSAnalyzers(name string, tokenizer FTSAnalyzersTokenizer, ) *FTSAnalyzers`

NewFTSAnalyzers instantiates a new FTSAnalyzers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSAnalyzersWithDefaults

`func NewFTSAnalyzersWithDefaults() *FTSAnalyzers`

NewFTSAnalyzersWithDefaults instantiates a new FTSAnalyzers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharFilters

`func (o *FTSAnalyzers) GetCharFilters() []FTSAnalyzersCharFiltersInner`

GetCharFilters returns the CharFilters field if non-nil, zero value otherwise.

### GetCharFiltersOk

`func (o *FTSAnalyzers) GetCharFiltersOk() (*[]FTSAnalyzersCharFiltersInner, bool)`

GetCharFiltersOk returns a tuple with the CharFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharFilters

`func (o *FTSAnalyzers) SetCharFilters(v []FTSAnalyzersCharFiltersInner)`

SetCharFilters sets CharFilters field to given value.

### HasCharFilters

`func (o *FTSAnalyzers) HasCharFilters() bool`

HasCharFilters returns a boolean if a field has been set.

### GetName

`func (o *FTSAnalyzers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FTSAnalyzers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FTSAnalyzers) SetName(v string)`

SetName sets Name field to given value.


### GetTokenFilters

`func (o *FTSAnalyzers) GetTokenFilters() []FTSAnalyzersTokenFiltersInner`

GetTokenFilters returns the TokenFilters field if non-nil, zero value otherwise.

### GetTokenFiltersOk

`func (o *FTSAnalyzers) GetTokenFiltersOk() (*[]FTSAnalyzersTokenFiltersInner, bool)`

GetTokenFiltersOk returns a tuple with the TokenFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFilters

`func (o *FTSAnalyzers) SetTokenFilters(v []FTSAnalyzersTokenFiltersInner)`

SetTokenFilters sets TokenFilters field to given value.

### HasTokenFilters

`func (o *FTSAnalyzers) HasTokenFilters() bool`

HasTokenFilters returns a boolean if a field has been set.

### GetTokenizer

`func (o *FTSAnalyzers) GetTokenizer() FTSAnalyzersTokenizer`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *FTSAnalyzers) GetTokenizerOk() (*FTSAnalyzersTokenizer, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *FTSAnalyzers) SetTokenizer(v FTSAnalyzersTokenizer)`

SetTokenizer sets Tokenizer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


