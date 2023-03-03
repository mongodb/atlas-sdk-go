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

// checks if the FTSSynonymMappingDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSSynonymMappingDefinition{}

// FTSSynonymMappingDefinition Synonyms used for this full text index.
type FTSSynonymMappingDefinition struct {
	// Specific pre-defined method chosen to apply to the synonyms to be searched.
	Analyzer string `json:"analyzer"`
	// Human-readable label that identifies the synonym definition. Each **synonym.name** must be unique within the same index definition.
	Name string `json:"name"`
	Source SynonymSource `json:"source"`
}

// NewFTSSynonymMappingDefinition instantiates a new FTSSynonymMappingDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSSynonymMappingDefinition(analyzer string, name string, source SynonymSource) *FTSSynonymMappingDefinition {
	this := FTSSynonymMappingDefinition{}
	this.Analyzer = analyzer
	this.Name = name
	this.Source = source
	return &this
}

// NewFTSSynonymMappingDefinitionWithDefaults instantiates a new FTSSynonymMappingDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSSynonymMappingDefinitionWithDefaults() *FTSSynonymMappingDefinition {
	this := FTSSynonymMappingDefinition{}
	var analyzer string = "lucene.standard"
	this.Analyzer = analyzer
	return &this
}

// GetAnalyzer returns the Analyzer field value
func (o *FTSSynonymMappingDefinition) GetAnalyzer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value
// and a boolean to check if the value has been set.
func (o *FTSSynonymMappingDefinition) GetAnalyzerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Analyzer, true
}

// SetAnalyzer sets field value
func (o *FTSSynonymMappingDefinition) SetAnalyzer(v string) {
	o.Analyzer = v
}

// GetName returns the Name field value
func (o *FTSSynonymMappingDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FTSSynonymMappingDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FTSSynonymMappingDefinition) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value
func (o *FTSSynonymMappingDefinition) GetSource() SynonymSource {
	if o == nil {
		var ret SynonymSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *FTSSynonymMappingDefinition) GetSourceOk() (*SynonymSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *FTSSynonymMappingDefinition) SetSource(v SynonymSource) {
	o.Source = v
}

func (o FTSSynonymMappingDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FTSSynonymMappingDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["analyzer"] = o.Analyzer
	toSerialize["name"] = o.Name
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableFTSSynonymMappingDefinition struct {
	value *FTSSynonymMappingDefinition
	isSet bool
}

func (v NullableFTSSynonymMappingDefinition) Get() *FTSSynonymMappingDefinition {
	return v.value
}

func (v *NullableFTSSynonymMappingDefinition) Set(val *FTSSynonymMappingDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSSynonymMappingDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSSynonymMappingDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSSynonymMappingDefinition(val *FTSSynonymMappingDefinition) *NullableFTSSynonymMappingDefinition {
	return &NullableFTSSynonymMappingDefinition{value: val, isSet: true}
}

func (v NullableFTSSynonymMappingDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSSynonymMappingDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


