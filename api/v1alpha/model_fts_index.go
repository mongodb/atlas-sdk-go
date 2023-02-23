/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the FTSIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSIndex{}

// FTSIndex struct for FTSIndex
type FTSIndex struct {
	// Specific pre-defined method chosen to convert database field text into searchable words. This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:  - extracting words - removing punctuation - removing accents - changing to lowercase - removing common words - reducing words to their root form (stemming) - changing words to their base form (lemmatization)  MongoDB Cloud uses the selected process to build the Atlas Search index.
	Analyzer *string `json:"analyzer,omitempty"`
	// List of user-defined methods to convert database field text into searchable words.
	Analyzers []ApiAtlasFTSAnalyzersViewManual `json:"analyzers,omitempty"`
	// Human-readable label that identifies the collection that contains one or more Atlas Search indexes.
	CollectionName string `json:"collectionName"`
	// Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
	Database string `json:"database"`
	// Unique 24-hexadecimal digit string that identifies this Atlas Search index.
	IndexID *string `json:"indexID,omitempty"`
	Mappings *ApiAtlasFTSMappingsViewManual `json:"mappings,omitempty"`
	// Human-readable label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique.
	Name string `json:"name"`
	// Method applied to identify words when searching this index.
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty"`
	// Condition of the search index when you made this request.  | Status | Index Condition |  |---|---|  | IN_PROGRESS | Atlas is building or re-building the index after an edit. |  | STEADY | You can use this search index. |  | FAILED | Atlas could not build the index. |  | MIGRATING | Atlas is upgrading the underlying cluster tier and migrating indexes. |  | PAUSED | The cluster is paused. | 
	Status *string `json:"status,omitempty"`
	// Rule sets that map words to their synonyms in this index.
	Synonyms []FTSSynonymMappingDefinition `json:"synonyms,omitempty"`
}

// NewFTSIndex instantiates a new FTSIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSIndex() *FTSIndex {
	this := FTSIndex{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// NewFTSIndexWithDefaults instantiates a new FTSIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSIndexWithDefaults() *FTSIndex {
	this := FTSIndex{}
	var analyzer string = "lucene.standard"
	this.Analyzer = &analyzer
	var searchAnalyzer string = "lucene.standard"
	this.SearchAnalyzer = &searchAnalyzer
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *FTSIndex) GetAnalyzer() string {
	if o == nil || IsNil(o.Analyzer) {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.Analyzer) {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *FTSIndex) HasAnalyzer() bool {
	if o != nil && !IsNil(o.Analyzer) {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *FTSIndex) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetAnalyzers returns the Analyzers field value if set, zero value otherwise.
func (o *FTSIndex) GetAnalyzers() []ApiAtlasFTSAnalyzersViewManual {
	if o == nil || IsNil(o.Analyzers) {
		var ret []ApiAtlasFTSAnalyzersViewManual
		return ret
	}
	return o.Analyzers
}

// GetAnalyzersOk returns a tuple with the Analyzers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetAnalyzersOk() ([]ApiAtlasFTSAnalyzersViewManual, bool) {
	if o == nil || IsNil(o.Analyzers) {
		return nil, false
	}
	return o.Analyzers, true
}

// HasAnalyzers returns a boolean if a field has been set.
func (o *FTSIndex) HasAnalyzers() bool {
	if o != nil && !IsNil(o.Analyzers) {
		return true
	}

	return false
}

// SetAnalyzers gets a reference to the given []ApiAtlasFTSAnalyzersViewManual and assigns it to the Analyzers field.
func (o *FTSIndex) SetAnalyzers(v []ApiAtlasFTSAnalyzersViewManual) {
	o.Analyzers = v
}

// GetCollectionName returns the CollectionName field value
func (o *FTSIndex) GetCollectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetCollectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionName, true
}

// SetCollectionName sets field value
func (o *FTSIndex) SetCollectionName(v string) {
	o.CollectionName = v
}

// GetDatabase returns the Database field value
func (o *FTSIndex) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value
func (o *FTSIndex) SetDatabase(v string) {
	o.Database = v
}

// GetIndexID returns the IndexID field value if set, zero value otherwise.
func (o *FTSIndex) GetIndexID() string {
	if o == nil || IsNil(o.IndexID) {
		var ret string
		return ret
	}
	return *o.IndexID
}

// GetIndexIDOk returns a tuple with the IndexID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetIndexIDOk() (*string, bool) {
	if o == nil || IsNil(o.IndexID) {
		return nil, false
	}
	return o.IndexID, true
}

// HasIndexID returns a boolean if a field has been set.
func (o *FTSIndex) HasIndexID() bool {
	if o != nil && !IsNil(o.IndexID) {
		return true
	}

	return false
}

// SetIndexID gets a reference to the given string and assigns it to the IndexID field.
func (o *FTSIndex) SetIndexID(v string) {
	o.IndexID = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *FTSIndex) GetMappings() ApiAtlasFTSMappingsViewManual {
	if o == nil || IsNil(o.Mappings) {
		var ret ApiAtlasFTSMappingsViewManual
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetMappingsOk() (*ApiAtlasFTSMappingsViewManual, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *FTSIndex) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given ApiAtlasFTSMappingsViewManual and assigns it to the Mappings field.
func (o *FTSIndex) SetMappings(v ApiAtlasFTSMappingsViewManual) {
	o.Mappings = &v
}

// GetName returns the Name field value
func (o *FTSIndex) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FTSIndex) SetName(v string) {
	o.Name = v
}

// GetSearchAnalyzer returns the SearchAnalyzer field value if set, zero value otherwise.
func (o *FTSIndex) GetSearchAnalyzer() string {
	if o == nil || IsNil(o.SearchAnalyzer) {
		var ret string
		return ret
	}
	return *o.SearchAnalyzer
}

// GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetSearchAnalyzerOk() (*string, bool) {
	if o == nil || IsNil(o.SearchAnalyzer) {
		return nil, false
	}
	return o.SearchAnalyzer, true
}

// HasSearchAnalyzer returns a boolean if a field has been set.
func (o *FTSIndex) HasSearchAnalyzer() bool {
	if o != nil && !IsNil(o.SearchAnalyzer) {
		return true
	}

	return false
}

// SetSearchAnalyzer gets a reference to the given string and assigns it to the SearchAnalyzer field.
func (o *FTSIndex) SetSearchAnalyzer(v string) {
	o.SearchAnalyzer = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FTSIndex) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FTSIndex) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FTSIndex) SetStatus(v string) {
	o.Status = &v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *FTSIndex) GetSynonyms() []FTSSynonymMappingDefinition {
	if o == nil || IsNil(o.Synonyms) {
		var ret []FTSSynonymMappingDefinition
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSIndex) GetSynonymsOk() ([]FTSSynonymMappingDefinition, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *FTSIndex) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []FTSSynonymMappingDefinition and assigns it to the Synonyms field.
func (o *FTSIndex) SetSynonyms(v []FTSSynonymMappingDefinition) {
	o.Synonyms = v
}

func (o FTSIndex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FTSIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analyzer) {
		toSerialize["analyzer"] = o.Analyzer
	}
	if !IsNil(o.Analyzers) {
		toSerialize["analyzers"] = o.Analyzers
	}
	toSerialize["collectionName"] = o.CollectionName
	toSerialize["database"] = o.Database
	// skip: indexID is readOnly
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SearchAnalyzer) {
		toSerialize["searchAnalyzer"] = o.SearchAnalyzer
	}
	// skip: status is readOnly
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	return toSerialize, nil
}

type NullableFTSIndex struct {
	value *FTSIndex
	isSet bool
}

func (v NullableFTSIndex) Get() *FTSIndex {
	return v.value
}

func (v *NullableFTSIndex) Set(val *FTSIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSIndex(val *FTSIndex) *NullableFTSIndex {
	return &NullableFTSIndex{value: val, isSet: true}
}

func (v NullableFTSIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


