# ClusterSearchIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** | Human-readable label that identifies the collection that contains one or more Atlas Search indexes. | 
**Database** | **string** | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes. | 
**IndexID** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this Atlas Search index. | [optional] [readonly] 
**Name** | **string** | Human-readable label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique. | 
**NumPartitions** | Pointer to **int** | Number of index partitions. Allowed values are [1, 2, 4]. | [optional] [default to 1]
**Status** | Pointer to **string** | Condition of the search index when you made this request.  - &#x60;IN_PROGRESS&#x60;: Atlas is building or re-building the index after an edit. - &#x60;STEADY&#x60;: You can use this search index. - &#x60;FAILED&#x60;: Atlas could not build the index. - &#x60;MIGRATING&#x60;: Atlas is upgrading the underlying cluster tier and migrating indexes. - &#x60;PAUSED&#x60;: The cluster is paused. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the index. Default type is search. | [optional] 
**Analyzer** | Pointer to **string** | Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the selected process to build the Atlas Search index. | [optional] [default to "lucene.standard"]
**Analyzers** | Pointer to [**[]ApiAtlasFTSAnalyzers**](ApiAtlasFTSAnalyzers.md) | List of user-defined methods to convert database field text into searchable words. | [optional] 
**Mappings** | Pointer to [**ApiAtlasFTSMappings**](ApiAtlasFTSMappings.md) |  | [optional] 
**SearchAnalyzer** | Pointer to **string** | Method applied to identify words when searching this index. | [optional] [default to "lucene.standard"]
**StoredSource** | Pointer to [**any**](interface{}.md) | Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn&#39;t store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. To learn more, see documentation. | [optional] 
**Synonyms** | Pointer to [**[]SearchSynonymMappingDefinition**](SearchSynonymMappingDefinition.md) | Rule sets that map words to their synonyms in this index. | [optional] 
**Fields** | Pointer to [**[]any**](any.md) | Settings that configure the fields, one per object, to index. You must define at least one \&quot;vector\&quot; type field. You can optionally define \&quot;filter\&quot; type fields also. | [optional] 

## Methods

### NewClusterSearchIndex

`func NewClusterSearchIndex(collectionName string, database string, name string, ) *ClusterSearchIndex`

NewClusterSearchIndex instantiates a new ClusterSearchIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSearchIndexWithDefaults

`func NewClusterSearchIndexWithDefaults() *ClusterSearchIndex`

NewClusterSearchIndexWithDefaults instantiates a new ClusterSearchIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *ClusterSearchIndex) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ClusterSearchIndex) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ClusterSearchIndex) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### GetDatabase

`func (o *ClusterSearchIndex) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ClusterSearchIndex) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ClusterSearchIndex) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### GetIndexID

`func (o *ClusterSearchIndex) GetIndexID() string`

GetIndexID returns the IndexID field if non-nil, zero value otherwise.

### GetIndexIDOk

`func (o *ClusterSearchIndex) GetIndexIDOk() (*string, bool)`

GetIndexIDOk returns a tuple with the IndexID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexID

`func (o *ClusterSearchIndex) SetIndexID(v string)`

SetIndexID sets IndexID field to given value.

### HasIndexID

`func (o *ClusterSearchIndex) HasIndexID() bool`

HasIndexID returns a boolean if a field has been set.
### GetName

`func (o *ClusterSearchIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterSearchIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterSearchIndex) SetName(v string)`

SetName sets Name field to given value.

### GetNumPartitions

`func (o *ClusterSearchIndex) GetNumPartitions() int`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *ClusterSearchIndex) GetNumPartitionsOk() (*int, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *ClusterSearchIndex) SetNumPartitions(v int)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *ClusterSearchIndex) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.
### GetStatus

`func (o *ClusterSearchIndex) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterSearchIndex) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterSearchIndex) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterSearchIndex) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetType

`func (o *ClusterSearchIndex) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterSearchIndex) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterSearchIndex) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterSearchIndex) HasType() bool`

HasType returns a boolean if a field has been set.
### GetAnalyzer

`func (o *ClusterSearchIndex) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *ClusterSearchIndex) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *ClusterSearchIndex) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *ClusterSearchIndex) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.
### GetAnalyzers

`func (o *ClusterSearchIndex) GetAnalyzers() []ApiAtlasFTSAnalyzers`

GetAnalyzers returns the Analyzers field if non-nil, zero value otherwise.

### GetAnalyzersOk

`func (o *ClusterSearchIndex) GetAnalyzersOk() (*[]ApiAtlasFTSAnalyzers, bool)`

GetAnalyzersOk returns a tuple with the Analyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzers

`func (o *ClusterSearchIndex) SetAnalyzers(v []ApiAtlasFTSAnalyzers)`

SetAnalyzers sets Analyzers field to given value.

### HasAnalyzers

`func (o *ClusterSearchIndex) HasAnalyzers() bool`

HasAnalyzers returns a boolean if a field has been set.
### GetMappings

`func (o *ClusterSearchIndex) GetMappings() ApiAtlasFTSMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ClusterSearchIndex) GetMappingsOk() (*ApiAtlasFTSMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ClusterSearchIndex) SetMappings(v ApiAtlasFTSMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ClusterSearchIndex) HasMappings() bool`

HasMappings returns a boolean if a field has been set.
### GetSearchAnalyzer

`func (o *ClusterSearchIndex) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *ClusterSearchIndex) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *ClusterSearchIndex) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *ClusterSearchIndex) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.
### GetStoredSource

`func (o *ClusterSearchIndex) GetStoredSource() any`

GetStoredSource returns the StoredSource field if non-nil, zero value otherwise.

### GetStoredSourceOk

`func (o *ClusterSearchIndex) GetStoredSourceOk() (*any, bool)`

GetStoredSourceOk returns a tuple with the StoredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredSource

`func (o *ClusterSearchIndex) SetStoredSource(v any)`

SetStoredSource sets StoredSource field to given value.

### HasStoredSource

`func (o *ClusterSearchIndex) HasStoredSource() bool`

HasStoredSource returns a boolean if a field has been set.
### GetSynonyms

`func (o *ClusterSearchIndex) GetSynonyms() []SearchSynonymMappingDefinition`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *ClusterSearchIndex) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *ClusterSearchIndex) SetSynonyms(v []SearchSynonymMappingDefinition)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *ClusterSearchIndex) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.
### GetFields

`func (o *ClusterSearchIndex) GetFields() []any`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ClusterSearchIndex) GetFieldsOk() (*[]any, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ClusterSearchIndex) SetFields(v []any)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ClusterSearchIndex) HasFields() bool`

HasFields returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


