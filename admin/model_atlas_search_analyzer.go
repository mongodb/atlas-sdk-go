// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// AtlasSearchAnalyzer struct for AtlasSearchAnalyzer
type AtlasSearchAnalyzer struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters *[]map[string]interface{} `json:"charFilters,omitempty"`
	// Name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - `lucene.` - `builtin.` - `mongodb.`
	Name string `json:"name"`
	// Filter that performs operations such as:  - Stemming, which reduces related words, such as \"talking\", \"talked\", and \"talks\" to their root word \"talk\".  - Redaction, which is the removal of sensitive information from public documents.
	TokenFilters *[]map[string]interface{} `json:"tokenFilters,omitempty"`
	// Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing.
	Tokenizer map[string]interface{} `json:"tokenizer"`
}

// NewAtlasSearchAnalyzer instantiates a new AtlasSearchAnalyzer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasSearchAnalyzer(name string, tokenizer map[string]interface{}) *AtlasSearchAnalyzer {
	this := AtlasSearchAnalyzer{}
	this.Name = name
	this.Tokenizer = tokenizer
	return &this
}

// NewAtlasSearchAnalyzerWithDefaults instantiates a new AtlasSearchAnalyzer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasSearchAnalyzerWithDefaults() *AtlasSearchAnalyzer {
	this := AtlasSearchAnalyzer{}
	return &this
}

// GetCharFilters returns the CharFilters field value if set, zero value otherwise
func (o *AtlasSearchAnalyzer) GetCharFilters() []map[string]interface{} {
	if o == nil || IsNil(o.CharFilters) {
		var ret []map[string]interface{}
		return ret
	}
	return *o.CharFilters
}

// GetCharFiltersOk returns a tuple with the CharFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasSearchAnalyzer) GetCharFiltersOk() (*[]map[string]interface{}, bool) {
	if o == nil || IsNil(o.CharFilters) {
		return nil, false
	}

	return o.CharFilters, true
}

// HasCharFilters returns a boolean if a field has been set.
func (o *AtlasSearchAnalyzer) HasCharFilters() bool {
	if o != nil && !IsNil(o.CharFilters) {
		return true
	}

	return false
}

// SetCharFilters gets a reference to the given []map[string]interface{} and assigns it to the CharFilters field.
func (o *AtlasSearchAnalyzer) SetCharFilters(v []map[string]interface{}) {
	o.CharFilters = &v
}

// GetName returns the Name field value
func (o *AtlasSearchAnalyzer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AtlasSearchAnalyzer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AtlasSearchAnalyzer) SetName(v string) {
	o.Name = v
}

// GetTokenFilters returns the TokenFilters field value if set, zero value otherwise
func (o *AtlasSearchAnalyzer) GetTokenFilters() []map[string]interface{} {
	if o == nil || IsNil(o.TokenFilters) {
		var ret []map[string]interface{}
		return ret
	}
	return *o.TokenFilters
}

// GetTokenFiltersOk returns a tuple with the TokenFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasSearchAnalyzer) GetTokenFiltersOk() (*[]map[string]interface{}, bool) {
	if o == nil || IsNil(o.TokenFilters) {
		return nil, false
	}

	return o.TokenFilters, true
}

// HasTokenFilters returns a boolean if a field has been set.
func (o *AtlasSearchAnalyzer) HasTokenFilters() bool {
	if o != nil && !IsNil(o.TokenFilters) {
		return true
	}

	return false
}

// SetTokenFilters gets a reference to the given []map[string]interface{} and assigns it to the TokenFilters field.
func (o *AtlasSearchAnalyzer) SetTokenFilters(v []map[string]interface{}) {
	o.TokenFilters = &v
}

// GetTokenizer returns the Tokenizer field value
func (o *AtlasSearchAnalyzer) GetTokenizer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value
// and a boolean to check if the value has been set.
func (o *AtlasSearchAnalyzer) GetTokenizerOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Tokenizer, true
}

// SetTokenizer sets field value
func (o *AtlasSearchAnalyzer) SetTokenizer(v map[string]interface{}) {
	o.Tokenizer = v
}

func (o AtlasSearchAnalyzer) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AtlasSearchAnalyzer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CharFilters) {
		toSerialize["charFilters"] = o.CharFilters
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TokenFilters) {
		toSerialize["tokenFilters"] = o.TokenFilters
	}
	toSerialize["tokenizer"] = o.Tokenizer
	return toSerialize, nil
}
