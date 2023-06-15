// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the SynonymSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SynonymSource{}

// SynonymSource Data set that stores the mapping one or more words map to one or more synonyms of those words.
type SynonymSource struct {
	// Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms.
	Collection string `json:"collection"`
}

// NewSynonymSource instantiates a new SynonymSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynonymSource(collection string) *SynonymSource {
	this := SynonymSource{}
	this.Collection = collection
	return &this
}

// NewSynonymSourceWithDefaults instantiates a new SynonymSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynonymSourceWithDefaults() *SynonymSource {
	this := SynonymSource{}
	return &this
}

// GetCollection returns the Collection field value
func (o *SynonymSource) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *SynonymSource) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *SynonymSource) SetCollection(v string) {
	o.Collection = v
}

func (o SynonymSource) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SynonymSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	return toSerialize, nil
}

type NullableSynonymSource struct {
	value *SynonymSource
	isSet bool
}

func (v NullableSynonymSource) Get() *SynonymSource {
	return v.value
}

func (v *NullableSynonymSource) Set(val *SynonymSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSynonymSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSynonymSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynonymSource(val *SynonymSource) *NullableSynonymSource {
	return &NullableSynonymSource{value: val, isSet: true}
}

func (v NullableSynonymSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynonymSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
