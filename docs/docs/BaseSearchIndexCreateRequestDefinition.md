# BaseSearchIndexCreateRequestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** | Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves making the following changes to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the process you select to build the Atlas Search index. | [optional] [default to "lucene.standard"]
**Analyzers** | Pointer to [**[]AtlasSearchAnalyzer**](AtlasSearchAnalyzer.md) | List of user-defined methods to convert database field text into searchable words. | [optional] 
**Mappings** | Pointer to [**SearchMappings**](SearchMappings.md) |  | [optional] 
**NumPartitions** | Pointer to **int** | Number of index partitions. Allowed values are [1, 2, 4]. | [optional] [default to 1]
**SearchAnalyzer** | Pointer to **string** | Method applied to identify words when searching this index. | [optional] [default to "lucene.standard"]
**Sort** | Pointer to [**any**](interface{}.md) | Sort definition for the index. When defined, the index will be pre-sorted on the specified fields, which improves query sort performance for those fields. Supports two formats: simple format with field name and direction, or complex format with additional options. The &#x60;order&#x60; field is required (1&#x3D;ascending, -1&#x3D;descending).The &#x60;noData&#x60; field is optional and controls how missing values are sorted(default: \&quot;lowest\&quot;). | [optional] 
**StoredSource** | Pointer to [**any**](interface{}.md) | Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn&#39;t store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. Note that storing all fields (true) is not allowed for vector search indexes. To learn more, see Stored Source Fields. | [optional] 
**Synonyms** | Pointer to [**[]SearchSynonymMappingDefinition**](SearchSynonymMappingDefinition.md) | Rule sets that map words to their synonyms in this index. | [optional] 
**TypeSets** | Pointer to [**[]SearchTypeSets**](SearchTypeSets.md) | Type sets for the index. | [optional] 
**Fields** | Pointer to [**[]any**](any.md) | Settings that configure the fields, one per object, to index. You must define at least one \&quot;vector\&quot; type field. You can optionally define \&quot;filter\&quot; type fields also. | [optional] 

## Methods

### NewBaseSearchIndexCreateRequestDefinition

`func NewBaseSearchIndexCreateRequestDefinition() *BaseSearchIndexCreateRequestDefinition`

NewBaseSearchIndexCreateRequestDefinition instantiates a new BaseSearchIndexCreateRequestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSearchIndexCreateRequestDefinitionWithDefaults

`func NewBaseSearchIndexCreateRequestDefinitionWithDefaults() *BaseSearchIndexCreateRequestDefinition`

NewBaseSearchIndexCreateRequestDefinitionWithDefaults instantiates a new BaseSearchIndexCreateRequestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.
### GetAnalyzers

`func (o *BaseSearchIndexCreateRequestDefinition) GetAnalyzers() []AtlasSearchAnalyzer`

GetAnalyzers returns the Analyzers field if non-nil, zero value otherwise.

### GetAnalyzersOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetAnalyzersOk() (*[]AtlasSearchAnalyzer, bool)`

GetAnalyzersOk returns a tuple with the Analyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzers

`func (o *BaseSearchIndexCreateRequestDefinition) SetAnalyzers(v []AtlasSearchAnalyzer)`

SetAnalyzers sets Analyzers field to given value.

### HasAnalyzers

`func (o *BaseSearchIndexCreateRequestDefinition) HasAnalyzers() bool`

HasAnalyzers returns a boolean if a field has been set.
### GetMappings

`func (o *BaseSearchIndexCreateRequestDefinition) GetMappings() SearchMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetMappingsOk() (*SearchMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *BaseSearchIndexCreateRequestDefinition) SetMappings(v SearchMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *BaseSearchIndexCreateRequestDefinition) HasMappings() bool`

HasMappings returns a boolean if a field has been set.
### GetNumPartitions

`func (o *BaseSearchIndexCreateRequestDefinition) GetNumPartitions() int`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetNumPartitionsOk() (*int, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *BaseSearchIndexCreateRequestDefinition) SetNumPartitions(v int)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *BaseSearchIndexCreateRequestDefinition) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.
### GetSearchAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *BaseSearchIndexCreateRequestDefinition) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.
### GetSort

`func (o *BaseSearchIndexCreateRequestDefinition) GetSort() any`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetSortOk() (*any, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BaseSearchIndexCreateRequestDefinition) SetSort(v any)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BaseSearchIndexCreateRequestDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.
### GetStoredSource

`func (o *BaseSearchIndexCreateRequestDefinition) GetStoredSource() any`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetStoredSourceOk() (*any, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *BaseSearchIndexCreateRequestDefinition) SetStoredSource(v any)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *BaseSearchIndexCreateRequestDefinition) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.
### GetSynonyms

`func (o *BaseSearchIndexCreateRequestDefinition) GetSynonyms() []SearchSynonymMappingDefinition`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *BaseSearchIndexCreateRequestDefinition) SetSynonyms(v []SearchSynonymMappingDefinition)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *BaseSearchIndexCreateRequestDefinition) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.
### GetTypeSets

`func (o *BaseSearchIndexCreateRequestDefinition) GetTypeSets() []SearchTypeSets`

GetTypeSets returns the TypeSets field if non-nil, zero value otherwise.

### GetTypeSetsOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetTypeSetsOk() (*[]SearchTypeSets, bool)`

GetTypeSetsOk returns a tuple with the TypeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSets

`func (o *BaseSearchIndexCreateRequestDefinition) SetTypeSets(v []SearchTypeSets)`

SetTypeSets sets TypeSets field to given value.

### HasTypeSets

`func (o *BaseSearchIndexCreateRequestDefinition) HasTypeSets() bool`

HasTypeSets returns a boolean if a field has been set.
### GetFields

`func (o *BaseSearchIndexCreateRequestDefinition) GetFields() []any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BaseSearchIndexCreateRequestDefinition) GetFieldsOk() (*[]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BaseSearchIndexCreateRequestDefinition) SetFields(v []any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BaseSearchIndexCreateRequestDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


