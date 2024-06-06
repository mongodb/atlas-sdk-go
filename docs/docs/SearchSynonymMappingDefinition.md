# SearchSynonymMappingDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | **string** | Specific pre-defined method chosen to apply to the synonyms to be searched. | [default to "lucene.standard"]
**Name** | **string** | Label that identifies the synonym definition. Each **synonym.name** must be unique within the same index definition. | 
**Source** | [**SynonymSource**](SynonymSource.md) |  | 

## Methods

### NewSearchSynonymMappingDefinition

`func NewSearchSynonymMappingDefinition(analyzer string, name string, source SynonymSource, ) *SearchSynonymMappingDefinition`

NewSearchSynonymMappingDefinition instantiates a new SearchSynonymMappingDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSynonymMappingDefinitionWithDefaults

`func NewSearchSynonymMappingDefinitionWithDefaults() *SearchSynonymMappingDefinition`

NewSearchSynonymMappingDefinitionWithDefaults instantiates a new SearchSynonymMappingDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *SearchSynonymMappingDefinition) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *SearchSynonymMappingDefinition) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *SearchSynonymMappingDefinition) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### GetName

`func (o *SearchSynonymMappingDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchSynonymMappingDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchSynonymMappingDefinition) SetName(v string)`

SetName sets Name field to given value.

### GetSource

`func (o *SearchSynonymMappingDefinition) GetSource() SynonymSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchSynonymMappingDefinition) GetSourceOk() (*SynonymSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchSynonymMappingDefinition) SetSource(v SynonymSource)`

SetSource sets Source field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


