# BaseSearchIndexResponseLatestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** | Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves making the following changes to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the process you select to build the Atlas Search index. | [optional] [default to "lucene.standard"]
**Analyzers** | Pointer to [**[]AtlasSearchAnalyzer**](AtlasSearchAnalyzer.md) | List of user-defined methods to convert database field text into searchable words. | [optional] 
**Mappings** | Pointer to [**SearchMappings**](SearchMappings.md) |  | [optional] 
**NumPartitions** | Pointer to **int** | Number of index partitions. Allowed values are [1, 2, 4]. | [optional] [default to 1]
**SearchAnalyzer** | Pointer to **string** | Method applied to identify words when searching this index. | [optional] [default to "lucene.standard"]
**StoredSource** | Pointer to [**any**](interface{}.md) | Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn&#39;t store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. To learn more, see Stored Source Fields. | [optional] 
**Synonyms** | Pointer to [**[]SearchSynonymMappingDefinition**](SearchSynonymMappingDefinition.md) | Rule sets that map words to their synonyms in this index. | [optional] 
**TypeSets** | Pointer to [**[]SearchTypeSets**](SearchTypeSets.md) | Type sets for the index. | [optional] 
**Fields** | Pointer to [**[]any**](any.md) | Settings that configure the fields, one per object, to index. You must define at least one \&quot;vector\&quot; type field. You can optionally define \&quot;filter\&quot; type fields also. | [optional] 

## Methods

### NewBaseSearchIndexResponseLatestDefinition

`func NewBaseSearchIndexResponseLatestDefinition() *BaseSearchIndexResponseLatestDefinition`

NewBaseSearchIndexResponseLatestDefinition instantiates a new BaseSearchIndexResponseLatestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSearchIndexResponseLatestDefinitionWithDefaults

`func NewBaseSearchIndexResponseLatestDefinitionWithDefaults() *BaseSearchIndexResponseLatestDefinition`

NewBaseSearchIndexResponseLatestDefinitionWithDefaults instantiates a new BaseSearchIndexResponseLatestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.
### GetAnalyzers

`func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzers() []AtlasSearchAnalyzer`

GetAnalyzers returns the Analyzers field if non-nil, zero value otherwise.

### GetAnalyzersOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzersOk() (*[]AtlasSearchAnalyzer, bool)`

GetAnalyzersOk returns a tuple with the Analyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzers

`func (o *BaseSearchIndexResponseLatestDefinition) SetAnalyzers(v []AtlasSearchAnalyzer)`

SetAnalyzers sets Analyzers field to given value.

### HasAnalyzers

`func (o *BaseSearchIndexResponseLatestDefinition) HasAnalyzers() bool`

HasAnalyzers returns a boolean if a field has been set.
### GetMappings

`func (o *BaseSearchIndexResponseLatestDefinition) GetMappings() SearchMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetMappingsOk() (*SearchMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *BaseSearchIndexResponseLatestDefinition) SetMappings(v SearchMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *BaseSearchIndexResponseLatestDefinition) HasMappings() bool`

HasMappings returns a boolean if a field has been set.
### GetNumPartitions

`func (o *BaseSearchIndexResponseLatestDefinition) GetNumPartitions() int`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetNumPartitionsOk() (*int, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *BaseSearchIndexResponseLatestDefinition) SetNumPartitions(v int)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *BaseSearchIndexResponseLatestDefinition) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.
### GetSearchAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *BaseSearchIndexResponseLatestDefinition) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.
### GetStoredSource

`func (o *BaseSearchIndexResponseLatestDefinition) GetStoredSource() any`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetStoredSourceOk() (*any, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *BaseSearchIndexResponseLatestDefinition) SetStoredSource(v any)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *BaseSearchIndexResponseLatestDefinition) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.
### GetSynonyms

`func (o *BaseSearchIndexResponseLatestDefinition) GetSynonyms() []SearchSynonymMappingDefinition`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *BaseSearchIndexResponseLatestDefinition) SetSynonyms(v []SearchSynonymMappingDefinition)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *BaseSearchIndexResponseLatestDefinition) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.
### GetTypeSets

`func (o *BaseSearchIndexResponseLatestDefinition) GetTypeSets() []SearchTypeSets`

GetTypeSets returns the TypeSets field if non-nil, zero value otherwise.

### GetTypeSetsOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetTypeSetsOk() (*[]SearchTypeSets, bool)`

GetTypeSetsOk returns a tuple with the TypeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSets

`func (o *BaseSearchIndexResponseLatestDefinition) SetTypeSets(v []SearchTypeSets)`

SetTypeSets sets TypeSets field to given value.

### HasTypeSets

`func (o *BaseSearchIndexResponseLatestDefinition) HasTypeSets() bool`

HasTypeSets returns a boolean if a field has been set.
### GetFields

`func (o *BaseSearchIndexResponseLatestDefinition) GetFields() []any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BaseSearchIndexResponseLatestDefinition) GetFieldsOk() (*[]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BaseSearchIndexResponseLatestDefinition) SetFields(v []any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BaseSearchIndexResponseLatestDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


