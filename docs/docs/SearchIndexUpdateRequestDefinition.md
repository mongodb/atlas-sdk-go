# SearchIndexUpdateRequestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** | Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves making the following changes to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the process you select to build the Atlas Search index. | [optional] [default to "lucene.standard"]
**Analyzers** | Pointer to [**[]AtlasSearchAnalyzer**](AtlasSearchAnalyzer.md) | List of user-defined methods to convert database field text into searchable words. | [optional] 
**Mappings** | Pointer to [**SearchMappings**](SearchMappings.md) |  | [optional] 
**NumPartitions** | Pointer to **int** | Number of index partitions. Note: This feature is currently in preview. | [optional] [default to 1]
**SearchAnalyzer** | Pointer to **string** | Method applied to identify words when searching this index. | [optional] [default to "lucene.standard"]
**StoredSource** | Pointer to [**any**](interface{}.md) | Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn&#39;t store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. To learn more, see Stored Source Fields. | [optional] 
**Synonyms** | Pointer to [**[]SearchSynonymMappingDefinition**](SearchSynonymMappingDefinition.md) | Rule sets that map words to their synonyms in this index. | [optional] 
**Fields** | Pointer to [**[]any**](any.md) | Settings that configure the fields, one per object, to index. You must define at least one \&quot;vector\&quot; type field. You can optionally define \&quot;filter\&quot; type fields also. | [optional] 

## Methods

### NewSearchIndexUpdateRequestDefinition

`func NewSearchIndexUpdateRequestDefinition() *SearchIndexUpdateRequestDefinition`

NewSearchIndexUpdateRequestDefinition instantiates a new SearchIndexUpdateRequestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexUpdateRequestDefinitionWithDefaults

`func NewSearchIndexUpdateRequestDefinitionWithDefaults() *SearchIndexUpdateRequestDefinition`

NewSearchIndexUpdateRequestDefinitionWithDefaults instantiates a new SearchIndexUpdateRequestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *SearchIndexUpdateRequestDefinition) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.
### GetAnalyzers

`func (o *SearchIndexUpdateRequestDefinition) GetAnalyzers() []AtlasSearchAnalyzer`

GetAnalyzers returns the Analyzers field if non-nil, zero value otherwise.

### GetAnalyzersOk

`func (o *SearchIndexUpdateRequestDefinition) GetAnalyzersOk() (*[]AtlasSearchAnalyzer, bool)`

GetAnalyzersOk returns a tuple with the Analyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzers

`func (o *SearchIndexUpdateRequestDefinition) SetAnalyzers(v []AtlasSearchAnalyzer)`

SetAnalyzers sets Analyzers field to given value.

### HasAnalyzers

`func (o *SearchIndexUpdateRequestDefinition) HasAnalyzers() bool`

HasAnalyzers returns a boolean if a field has been set.
### GetMappings

`func (o *SearchIndexUpdateRequestDefinition) GetMappings() SearchMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *SearchIndexUpdateRequestDefinition) GetMappingsOk() (*SearchMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *SearchIndexUpdateRequestDefinition) SetMappings(v SearchMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *SearchIndexUpdateRequestDefinition) HasMappings() bool`

HasMappings returns a boolean if a field has been set.
### GetNumPartitions

`func (o *SearchIndexUpdateRequestDefinition) GetNumPartitions() int`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *SearchIndexUpdateRequestDefinition) GetNumPartitionsOk() (*int, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *SearchIndexUpdateRequestDefinition) SetNumPartitions(v int)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *SearchIndexUpdateRequestDefinition) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.
### GetSearchAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *SearchIndexUpdateRequestDefinition) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *SearchIndexUpdateRequestDefinition) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.
### GetStoredSource

`func (o *SearchIndexUpdateRequestDefinition) GetStoredSource() any`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *SearchIndexUpdateRequestDefinition) GetStoredSourceOk() (*any, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *SearchIndexUpdateRequestDefinition) SetStoredSource(v any)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *SearchIndexUpdateRequestDefinition) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.
### GetSynonyms

`func (o *SearchIndexUpdateRequestDefinition) GetSynonyms() []SearchSynonymMappingDefinition`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *SearchIndexUpdateRequestDefinition) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *SearchIndexUpdateRequestDefinition) SetSynonyms(v []SearchSynonymMappingDefinition)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *SearchIndexUpdateRequestDefinition) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.
### GetFields

`func (o *SearchIndexUpdateRequestDefinition) GetFields() []any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchIndexUpdateRequestDefinition) GetFieldsOk() (*[]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchIndexUpdateRequestDefinition) SetFields(v []any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchIndexUpdateRequestDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


