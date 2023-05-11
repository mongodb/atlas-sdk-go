# FTSSynonymMappingDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | **string** | Specific pre-defined method chosen to apply to the synonyms to be searched. | [default to "lucene.standard"]
**Name** | **string** | Human-readable label that identifies the synonym definition. Each **synonym.name** must be unique within the same index definition. | 
**Source** | [**SynonymSource**](SynonymSource.md) |  | 

## Methods

### NewFTSSynonymMappingDefinition

`func NewFTSSynonymMappingDefinition(analyzer string, name string, source SynonymSource, ) *FTSSynonymMappingDefinition`

NewFTSSynonymMappingDefinition instantiates a new FTSSynonymMappingDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTSSynonymMappingDefinitionWithDefaults

`func NewFTSSynonymMappingDefinitionWithDefaults() *FTSSynonymMappingDefinition`

NewFTSSynonymMappingDefinitionWithDefaults instantiates a new FTSSynonymMappingDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *FTSSynonymMappingDefinition) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *FTSSynonymMappingDefinition) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *FTSSynonymMappingDefinition) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.


### GetName

`func (o *FTSSynonymMappingDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FTSSynonymMappingDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FTSSynonymMappingDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *FTSSynonymMappingDefinition) GetSource() SynonymSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FTSSynonymMappingDefinition) GetSourceOk() (*SynonymSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FTSSynonymMappingDefinition) SetSource(v SynonymSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


