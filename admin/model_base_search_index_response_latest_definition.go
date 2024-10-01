// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// BaseSearchIndexResponseLatestDefinition struct for BaseSearchIndexResponseLatestDefinition
type BaseSearchIndexResponseLatestDefinition struct {
	// Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves making the following changes to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the process you select to build the Atlas Search index.
	Analyzer *string `json:"analyzer,omitempty"`
	// List of user-defined methods to convert database field text into searchable words.
	Analyzers *[]AtlasSearchAnalyzer `json:"analyzers,omitempty"`
	Mappings  *SearchMappings        `json:"mappings,omitempty"`
	// Method applied to identify words when searching this index.
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty"`
	// Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn't store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. To learn more, see Stored Source Fields.
	StoredSource interface{} `json:"storedSource,omitempty"`
	// Rule sets that map words to their synonyms in this index.
	Synonyms *[]SearchSynonymMappingDefinition `json:"synonyms,omitempty"`
	// Settings that configure the fields, one per object, to index. You must define at least one \"vector\" type field. You can optionally define \"filter\" type fields also.
	Fields *[]interface{} `json:"fields,omitempty"`
}

// NewBaseSearchIndexResponseLatestDefinition instantiates a new BaseSearchIndexResponseLatestDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSearchIndexResponseLatestDefinition() *BaseSearchIndexResponseLatestDefinition {
	this := BaseSearchIndexResponseLatestDefinition{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// NewBaseSearchIndexResponseLatestDefinitionWithDefaults instantiates a new BaseSearchIndexResponseLatestDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSearchIndexResponseLatestDefinitionWithDefaults() *BaseSearchIndexResponseLatestDefinition {
	this := BaseSearchIndexResponseLatestDefinition{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}

	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *BaseSearchIndexResponseLatestDefinition) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetAnalyzers returns the Analyzers field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzers() []AtlasSearchAnalyzer {
	if o == nil || IsNil(o.Analyzers) {
		var ret []AtlasSearchAnalyzer
		return ret
	}
	return *o.Analyzers
}

// GetAnalyzersOk returns a tuple with the Analyzers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetAnalyzersOk() (*[]AtlasSearchAnalyzer, bool) {
	if o == nil || IsNil(o.Analyzers) {
		return nil, false
	}

	return o.Analyzers, true
}

// HasAnalyzers returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasAnalyzers() bool {
	if o != nil && !IsNil(o.Analyzers) {
		return true
	}

	return false
}

// SetAnalyzers gets a reference to the given []AtlasSearchAnalyzer and assigns it to the Analyzers field.
func (o *BaseSearchIndexResponseLatestDefinition) SetAnalyzers(v []AtlasSearchAnalyzer) {
	o.Analyzers = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetMappings() SearchMappings {
	if o == nil || IsNil(o.Mappings) {
		var ret SearchMappings
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetMappingsOk() (*SearchMappings, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}

	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given SearchMappings and assigns it to the Mappings field.
func (o *BaseSearchIndexResponseLatestDefinition) SetMappings(v SearchMappings) {
	o.Mappings = &v
}

// GetSearchAnalyzer returns the SearchAnalyzer field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetSearchAnalyzer() string {
	if o == nil || IsNil(o.SearchAnalyzer) {
		var ret string
		return ret
	}
	return *o.SearchAnalyzer
}

// GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetSearchAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.SearchAnalyzer) {
		return nil, false
	}

	return o.SearchAnalyzer, true
}

// HasSearchAnalyzer returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasSearchAnalyzer() bool {
	if o != nil && !IsNil(o.SearchAnalyzer) {
		return true
	}

	return false
}

// SetSearchAnalyzer gets a reference to the given string and assigns it to the SearchAnalyzer field.
func (o *BaseSearchIndexResponseLatestDefinition) SetSearchAnalyzer(v string) {
	o.SearchAnalyzer = &v
}

// GetStoredSource returns the StoredSource field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetStoredSource() interface{} {
	if o == nil || IsNil(o.StoredSource) {
		var ret interface{}
		return ret
	}
	return o.StoredSource
}

// GetStoredSourceOk returns a tuple with the StoredSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetStoredSourceOk() (interface{}, bool) {
	if o == nil || IsNil(o.StoredSource) {
		var ret interface{}
		return ret, false
	}

	return o.StoredSource, true
}

// HasStoredSource returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasStoredSource() bool {
	if o != nil && !IsNil(o.StoredSource) {
		return true
	}

	return false
}

// SetStoredSource gets a reference to the given interface{} and assigns it to the StoredSource field.
func (o *BaseSearchIndexResponseLatestDefinition) SetStoredSource(v interface{}) {
	o.StoredSource = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetSynonyms() []SearchSynonymMappingDefinition {
	if o == nil || IsNil(o.Synonyms) {
		var ret []SearchSynonymMappingDefinition
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}

	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []SearchSynonymMappingDefinition and assigns it to the Synonyms field.
func (o *BaseSearchIndexResponseLatestDefinition) SetSynonyms(v []SearchSynonymMappingDefinition) {
	o.Synonyms = &v
}

// GetFields returns the Fields field value if set, zero value otherwise
func (o *BaseSearchIndexResponseLatestDefinition) GetFields() []interface{} {
	if o == nil || IsNil(o.Fields) {
		var ret []interface{}
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchIndexResponseLatestDefinition) GetFieldsOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}

	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *BaseSearchIndexResponseLatestDefinition) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []interface{} and assigns it to the Fields field.
func (o *BaseSearchIndexResponseLatestDefinition) SetFields(v []interface{}) {
	o.Fields = &v
}

func (o *BaseSearchIndexResponseLatestDefinition) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *BaseSearchIndexResponseLatestDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.Analyzers) {
		toSerialize["analyzers"] = o.Analyzers
	}
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	if !IsNil(o.SearchAnalyzer) {
		toSerialize["searchAnalyzer"] = o.SearchAnalyzer
	}
	if !IsNil(o.StoredSource) {
		toSerialize["storedSource"] = o.StoredSource
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}
