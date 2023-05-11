# FTSIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** | Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the selected process to build the Atlas Search index. | [optional] [default to "lucene.standard"]
**Analyzers** | Pointer to [**[]FTSAnalyzers**](FTSAnalyzers.md) | List of user-defined methods to convert database field text into searchable words. | [optional] 
**CollectionName** | **string** | Human-readable label that identifies the collection that contains one or more Atlas Search indexes. | 
**Database** | **string** | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes. | 
**IndexID** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this Atlas Search index. | [optional] [readonly] 
**Mappings** | Pointer to [**FTSMappings**](FTSMappings.md) |  | [optional] 
**Name** | **string** | Human-readable label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique. | 
**SearchAnalyzer** | Pointer to **string** | Method applied to identify words when searching this index. | [optional] [default to "lucene.standard"]
**Status** | Pointer to **string** | Condition of the search index when you made this request.  | Status | Index Condition |  |---|---|  | IN_PROGRESS | Atlas is building or re-building the index after an edit. |  | STEADY | You can use this search index. |  | FAILED | Atlas could not build the index. |  | MIGRATING | Atlas is upgrading the underlying cluster tier and migrating indexes. |  | PAUSED | The cluster is paused. |  | [optional] [readonly] 
**Synonyms** | Pointer to [**[]FTSSynonymMappingDefinition**](FTSSynonymMappingDefinition.md) | Rule sets that map words to their synonyms in this index. | [optional] 

## Methods

### NewFTSIndex

`func NewFTSIndex(collectionName string, database string, name string, ) *FTSIndex`

NewFTSIndex instantiates a new FTSIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSIndexWithDefaults

`func NewFTSIndexWithDefaults() *FTSIndex`

NewFTSIndexWithDefaults instantiates a new FTSIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *FTSIndex) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *FTSIndex) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *FTSIndex) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *FTSIndex) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetAnalyzers

`func (o *FTSIndex) GetAnalyzers() []FTSAnalyzers`

GetAnalyzers returns the Analyzers field if non-nil, zero value otherwise.

### GetAnalyzersOk

`func (o *FTSIndex) GetAnalyzersOk() (*[]FTSAnalyzers, bool)`

GetAnalyzersOk returns a tuple with the Analyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzers

`func (o *FTSIndex) SetAnalyzers(v []FTSAnalyzers)`

SetAnalyzers sets Analyzers field to given value.

### HasAnalyzers

`func (o *FTSIndex) HasAnalyzers() bool`

HasAnalyzers returns a boolean if a field has been set.

### GetCollectionName

`func (o *FTSIndex) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *FTSIndex) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *FTSIndex) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetDatabase

`func (o *FTSIndex) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *FTSIndex) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *FTSIndex) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetIndexID

`func (o *FTSIndex) GetIndexID() string`

GetIndexID returns the IndexID field if non-nil, zero value otherwise.

### GetIndexIDOk

`func (o *FTSIndex) GetIndexIDOk() (*string, bool)`

GetIndexIDOk returns a tuple with the IndexID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexID

`func (o *FTSIndex) SetIndexID(v string)`

SetIndexID sets IndexID field to given value.

### HasIndexID

`func (o *FTSIndex) HasIndexID() bool`

HasIndexID returns a boolean if a field has been set.

### GetMappings

`func (o *FTSIndex) GetMappings() FTSMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *FTSIndex) GetMappingsOk() (*FTSMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *FTSIndex) SetMappings(v FTSMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *FTSIndex) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetName

`func (o *FTSIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FTSIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FTSIndex) SetName(v string)`

SetName sets Name field to given value.


### GetSearchAnalyzer

`func (o *FTSIndex) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *FTSIndex) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *FTSIndex) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *FTSIndex) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.

### GetStatus

`func (o *FTSIndex) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FTSIndex) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FTSIndex) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FTSIndex) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSynonyms

`func (o *FTSIndex) GetSynonyms() []FTSSynonymMappingDefinition`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *FTSIndex) GetSynonymsOk() (*[]FTSSynonymMappingDefinition, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *FTSIndex) SetSynonyms(v []FTSSynonymMappingDefinition)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *FTSIndex) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


