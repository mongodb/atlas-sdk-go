// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SearchIndexUpdateRequestDefinition The index definition to update the search index to.
type SearchIndexUpdateRequestDefinition struct {
	// Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves making the following changes to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the process you select to build the Atlas Search index.
	Analyzer *string `json:"analyzer,omitempty"`
	// List of user-defined methods to convert database field text into searchable words.
	Analyzers *[]AtlasSearchAnalyzer `json:"analyzers,omitempty"`
	Mappings  *SearchMappings        `json:"mappings,omitempty"`
	// Number of index partitions. Allowed values are [1, 2, 4].
	NumPartitions *int `json:"numPartitions,omitempty"`
	// Method applied to identify words when searching this index.
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty"`
	// Flag that indicates whether to store all fields (true) on Atlas Search. By default, Atlas doesn't store (false) the fields on Atlas Search.  Alternatively, you can specify an object that only contains the list of fields to store (include) or not store (exclude) on Atlas Search. To learn more, see Stored Source Fields.
	StoredSource any `json:"storedSource,omitempty"`
	// Rule sets that map words to their synonyms in this index.
	Synonyms *[]SearchSynonymMappingDefinition `json:"synonyms,omitempty"`
	// Settings that configure the fields, one per object, to index. You must define at least one \"vector\" type field. You can optionally define \"filter\" type fields also.
	Fields *[]any `json:"fields,omitempty"`
}

// NewSearchIndexUpdateRequestDefinition instantiates a new SearchIndexUpdateRequestDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexUpdateRequestDefinition() *SearchIndexUpdateRequestDefinition {
	this := SearchIndexUpdateRequestDefinition{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var numPartitions int = 1
	this.NumPartitions = &numPartitions
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// NewSearchIndexUpdateRequestDefinitionWithDefaults instantiates a new SearchIndexUpdateRequestDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexUpdateRequestDefinitionWithDefaults() *SearchIndexUpdateRequestDefinition {
	this := SearchIndexUpdateRequestDefinition{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var numPartitions int = 1
	this.NumPartitions = &numPartitions
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}

	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *SearchIndexUpdateRequestDefinition) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetAnalyzers returns the Analyzers field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetAnalyzers() []AtlasSearchAnalyzer {
	if o == nil || IsNil(o.Analyzers) {
		var ret []AtlasSearchAnalyzer
		return ret
	}
	return *o.Analyzers
}

// GetAnalyzersOk returns a tuple with the Analyzers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetAnalyzersOk() (*[]AtlasSearchAnalyzer, bool) {
	if o == nil || IsNil(o.Analyzers) {
		return nil, false
	}

	return o.Analyzers, true
}

// HasAnalyzers returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasAnalyzers() bool {
	if o != nil && !IsNil(o.Analyzers) {
		return true
	}

	return false
}

// SetAnalyzers gets a reference to the given []AtlasSearchAnalyzer and assigns it to the Analyzers field.
func (o *SearchIndexUpdateRequestDefinition) SetAnalyzers(v []AtlasSearchAnalyzer) {
	o.Analyzers = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetMappings() SearchMappings {
	if o == nil || IsNil(o.Mappings) {
		var ret SearchMappings
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetMappingsOk() (*SearchMappings, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}

	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given SearchMappings and assigns it to the Mappings field.
func (o *SearchIndexUpdateRequestDefinition) SetMappings(v SearchMappings) {
	o.Mappings = &v
}

// GetNumPartitions returns the NumPartitions field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetNumPartitions() int {
	if o == nil || IsNil(o.NumPartitions) {
		var ret int
		return ret
	}
	return *o.NumPartitions
}

// GetNumPartitionsOk returns a tuple with the NumPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetNumPartitionsOk() (*int, bool) {
	if o == nil || IsNil(o.NumPartitions) {
		return nil, false
	}

	return o.NumPartitions, true
}

// HasNumPartitions returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasNumPartitions() bool {
	if o != nil && !IsNil(o.NumPartitions) {
		return true
	}

	return false
}

// SetNumPartitions gets a reference to the given int and assigns it to the NumPartitions field.
func (o *SearchIndexUpdateRequestDefinition) SetNumPartitions(v int) {
	o.NumPartitions = &v
}

// GetSearchAnalyzer returns the SearchAnalyzer field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetSearchAnalyzer() string {
	if o == nil || IsNil(o.SearchAnalyzer) {
		var ret string
		return ret
	}
	return *o.SearchAnalyzer
}

// GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetSearchAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.SearchAnalyzer) {
		return nil, false
	}

	return o.SearchAnalyzer, true
}

// HasSearchAnalyzer returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasSearchAnalyzer() bool {
	if o != nil && !IsNil(o.SearchAnalyzer) {
		return true
	}

	return false
}

// SetSearchAnalyzer gets a reference to the given string and assigns it to the SearchAnalyzer field.
func (o *SearchIndexUpdateRequestDefinition) SetSearchAnalyzer(v string) {
	o.SearchAnalyzer = &v
}

// GetStoredSource returns the StoredSource field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetStoredSource() any {
	if o == nil || IsNil(o.StoredSource) {
		var ret any
		return ret
	}
	return o.StoredSource
}

// GetStoredSourceOk returns a tuple with the StoredSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetStoredSourceOk() (any, bool) {
	if o == nil || IsNil(o.StoredSource) {
		var ret any
		return ret, false
	}

	return o.StoredSource, true
}

// HasStoredSource returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasStoredSource() bool {
	if o != nil && !IsNil(o.StoredSource) {
		return true
	}

	return false
}

// SetStoredSource gets a reference to the given any and assigns it to the StoredSource field.
func (o *SearchIndexUpdateRequestDefinition) SetStoredSource(v any) {
	o.StoredSource = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetSynonyms() []SearchSynonymMappingDefinition {
	if o == nil || IsNil(o.Synonyms) {
		var ret []SearchSynonymMappingDefinition
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetSynonymsOk() (*[]SearchSynonymMappingDefinition, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}

	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []SearchSynonymMappingDefinition and assigns it to the Synonyms field.
func (o *SearchIndexUpdateRequestDefinition) SetSynonyms(v []SearchSynonymMappingDefinition) {
	o.Synonyms = &v
}

// GetFields returns the Fields field value if set, zero value otherwise
func (o *SearchIndexUpdateRequestDefinition) GetFields() []any {
	if o == nil || IsNil(o.Fields) {
		var ret []any
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequestDefinition) GetFieldsOk() (*[]any, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}

	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SearchIndexUpdateRequestDefinition) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []any and assigns it to the Fields field.
func (o *SearchIndexUpdateRequestDefinition) SetFields(v []any) {
	o.Fields = &v
}
