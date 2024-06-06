# AtlasSearchAnalyzer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharFilters** | Pointer to **[]map[string]interface{}** | Filters that examine text one character at a time and perform filtering operations. | [optional] 
**Name** | **string** | Name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - &#x60;lucene.&#x60; - &#x60;builtin.&#x60; - &#x60;mongodb.&#x60; | 
**TokenFilters** | Pointer to **[]map[string]interface{}** | Filter that performs operations such as:  - Stemming, which reduces related words, such as \&quot;talking\&quot;, \&quot;talked\&quot;, and \&quot;talks\&quot; to their root word \&quot;talk\&quot;.  - Redaction, which is the removal of sensitive information from public documents. | [optional] 
**Tokenizer** | **map[string]interface{}** | Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing. | 

## Methods

### NewAtlasSearchAnalyzer

`func NewAtlasSearchAnalyzer(name string, tokenizer map[string]interface{}, ) *AtlasSearchAnalyzer`

NewAtlasSearchAnalyzer instantiates a new AtlasSearchAnalyzer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtlasSearchAnalyzerWithDefaults

`func NewAtlasSearchAnalyzerWithDefaults() *AtlasSearchAnalyzer`

NewAtlasSearchAnalyzerWithDefaults instantiates a new AtlasSearchAnalyzer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharFilters

`func (o *AtlasSearchAnalyzer) GetCharFilters() []map[string]interface{}`

GetCharFilters returns the CharFilters field if non-nil, zero value otherwise.

### GetCharFiltersOk

`func (o *AtlasSearchAnalyzer) GetCharFiltersOk() (*[]map[string]interface{}, bool)`

GetCharFiltersOk returns a tuple with the CharFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharFilters

`func (o *AtlasSearchAnalyzer) SetCharFilters(v []map[string]interface{})`

SetCharFilters sets CharFilters field to given value.

### HasCharFilters

`func (o *AtlasSearchAnalyzer) HasCharFilters() bool`

HasCharFilters returns a boolean if a field has been set.
### GetName

`func (o *AtlasSearchAnalyzer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtlasSearchAnalyzer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtlasSearchAnalyzer) SetName(v string)`

SetName sets Name field to given value.

### GetTokenFilters

`func (o *AtlasSearchAnalyzer) GetTokenFilters() []map[string]interface{}`

GetTokenFilters returns the TokenFilters field if non-nil, zero value otherwise.

### GetTokenFiltersOk

`func (o *AtlasSearchAnalyzer) GetTokenFiltersOk() (*[]map[string]interface{}, bool)`

GetTokenFiltersOk returns a tuple with the TokenFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFilters

`func (o *AtlasSearchAnalyzer) SetTokenFilters(v []map[string]interface{})`

SetTokenFilters sets TokenFilters field to given value.

### HasTokenFilters

`func (o *AtlasSearchAnalyzer) HasTokenFilters() bool`

HasTokenFilters returns a boolean if a field has been set.
### GetTokenizer

`func (o *AtlasSearchAnalyzer) GetTokenizer() map[string]interface{}`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *AtlasSearchAnalyzer) GetTokenizerOk() (*map[string]interface{}, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *AtlasSearchAnalyzer) SetTokenizer(v map[string]interface{})`

SetTokenizer sets Tokenizer field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


