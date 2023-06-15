// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the FTSAnalyzers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSAnalyzers{}

// FTSAnalyzers Settings that describe one Atlas Search custom analyzer.
type FTSAnalyzers struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters []FTSAnalyzersCharFiltersInner `json:"charFilters,omitempty"`
	// Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - `lucene.` - `builtin.` - `mongodb.`
	Name string `json:"name"`
	// Filter that performs operations such as:  - Stemming, which reduces related words, such as \"talking\", \"talked\", and \"talks\" to their root word \"talk\".  - Redaction, the removal of sensitive information from public documents.
	TokenFilters []FTSAnalyzersTokenFiltersInner `json:"tokenFilters,omitempty"`
	Tokenizer    FTSAnalyzersTokenizer           `json:"tokenizer"`
}

// NewFTSAnalyzers instantiates a new FTSAnalyzers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSAnalyzers(name string, tokenizer FTSAnalyzersTokenizer) *FTSAnalyzers {
	this := FTSAnalyzers{}
	this.Name = name
	this.Tokenizer = tokenizer
	return &this
}

// NewFTSAnalyzersWithDefaults instantiates a new FTSAnalyzers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSAnalyzersWithDefaults() *FTSAnalyzers {
	this := FTSAnalyzers{}
	return &this
}

// GetCharFilters returns the CharFilters field value if set, zero value otherwise.
func (o *FTSAnalyzers) GetCharFilters() []FTSAnalyzersCharFiltersInner {
	if o == nil || IsNil(o.CharFilters) {
		var ret []FTSAnalyzersCharFiltersInner
		return ret
	}
	return o.CharFilters
}

// GetCharFiltersOk returns a tuple with the CharFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzers) GetCharFiltersOk() ([]FTSAnalyzersCharFiltersInner, bool) {
	if o == nil || IsNil(o.CharFilters) {
		return nil, false
	}
	return o.CharFilters, true
}

// HasCharFilters returns a boolean if a field has been set.
func (o *FTSAnalyzers) HasCharFilters() bool {
	if o != nil && !IsNil(o.CharFilters) {
		return true
	}

	return false
}

// SetCharFilters gets a reference to the given []FTSAnalyzersCharFiltersInner and assigns it to the CharFilters field.
func (o *FTSAnalyzers) SetCharFilters(v []FTSAnalyzersCharFiltersInner) {
	o.CharFilters = v
}

// GetName returns the Name field value
func (o *FTSAnalyzers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FTSAnalyzers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FTSAnalyzers) SetName(v string) {
	o.Name = v
}

// GetTokenFilters returns the TokenFilters field value if set, zero value otherwise.
func (o *FTSAnalyzers) GetTokenFilters() []FTSAnalyzersTokenFiltersInner {
	if o == nil || IsNil(o.TokenFilters) {
		var ret []FTSAnalyzersTokenFiltersInner
		return ret
	}
	return o.TokenFilters
}

// GetTokenFiltersOk returns a tuple with the TokenFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzers) GetTokenFiltersOk() ([]FTSAnalyzersTokenFiltersInner, bool) {
	if o == nil || IsNil(o.TokenFilters) {
		return nil, false
	}
	return o.TokenFilters, true
}

// HasTokenFilters returns a boolean if a field has been set.
func (o *FTSAnalyzers) HasTokenFilters() bool {
	if o != nil && !IsNil(o.TokenFilters) {
		return true
	}

	return false
}

// SetTokenFilters gets a reference to the given []FTSAnalyzersTokenFiltersInner and assigns it to the TokenFilters field.
func (o *FTSAnalyzers) SetTokenFilters(v []FTSAnalyzersTokenFiltersInner) {
	o.TokenFilters = v
}

// GetTokenizer returns the Tokenizer field value
func (o *FTSAnalyzers) GetTokenizer() FTSAnalyzersTokenizer {
	if o == nil {
		var ret FTSAnalyzersTokenizer
		return ret
	}

	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value
// and a boolean to check if the value has been set.
func (o *FTSAnalyzers) GetTokenizerOk() (*FTSAnalyzersTokenizer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokenizer, true
}

// SetTokenizer sets field value
func (o *FTSAnalyzers) SetTokenizer(v FTSAnalyzersTokenizer) {
	o.Tokenizer = v
}

func (o FTSAnalyzers) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSAnalyzers) ToMap() (map[string]interface{}, error) {
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

type NullableFTSAnalyzers struct {
	value *FTSAnalyzers
	isSet bool
}

func (v NullableFTSAnalyzers) Get() *FTSAnalyzers {
	return v.value
}

func (v *NullableFTSAnalyzers) Set(val *FTSAnalyzers) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzers) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzers(val *FTSAnalyzers) *NullableFTSAnalyzers {
	return &NullableFTSAnalyzers{value: val, isSet: true}
}

func (v NullableFTSAnalyzers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
