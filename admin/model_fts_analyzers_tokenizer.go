// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the FTSAnalyzersTokenizer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FTSAnalyzersTokenizer{}

// FTSAnalyzersTokenizer Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing.
type FTSAnalyzersTokenizer struct {
	// Characters to include in the longest token that Atlas Search creates.
	MaxGram *int `json:"maxGram,omitempty"`
	// Characters to include in the shortest token that Atlas Search creates.
	MinGram *int `json:"minGram,omitempty"`
	// Human-readable label that identifies this tokenizer type.
	Type *string `json:"type,omitempty"`
	// Index of the character group within the matching expression to extract into tokens. Use `0` to extract all character groups.
	Group *int `json:"group,omitempty"`
	// Regular expression to match against.
	Pattern *string `json:"pattern,omitempty"`
	// Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens.
	MaxTokenLength *int `json:"maxTokenLength,omitempty"`
}

// NewFTSAnalyzersTokenizer instantiates a new FTSAnalyzersTokenizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFTSAnalyzersTokenizer() *FTSAnalyzersTokenizer {
	this := FTSAnalyzersTokenizer{}
	var maxTokenLength int = 255
	this.MaxTokenLength = &maxTokenLength
	return &this
}

// NewFTSAnalyzersTokenizerWithDefaults instantiates a new FTSAnalyzersTokenizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFTSAnalyzersTokenizerWithDefaults() *FTSAnalyzersTokenizer {
	this := FTSAnalyzersTokenizer{}
	var maxTokenLength int = 255
	this.MaxTokenLength = &maxTokenLength
	return &this
}

// GetMaxGram returns the MaxGram field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetMaxGram() int {
	if o == nil || IsNil(o.MaxGram) {
		var ret int
		return ret
	}
	return *o.MaxGram
}

// GetMaxGramOk returns a tuple with the MaxGram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetMaxGramOk() (*int, bool) {
	if o == nil || IsNil(o.MaxGram) {
		return nil, false
	}
	return o.MaxGram, true
}

// HasMaxGram returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasMaxGram() bool {
	if o != nil && !IsNil(o.MaxGram) {
		return true
	}

	return false
}

// SetMaxGram gets a reference to the given int and assigns it to the MaxGram field.
func (o *FTSAnalyzersTokenizer) SetMaxGram(v int) {
	o.MaxGram = &v
}

// GetMinGram returns the MinGram field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetMinGram() int {
	if o == nil || IsNil(o.MinGram) {
		var ret int
		return ret
	}
	return *o.MinGram
}

// GetMinGramOk returns a tuple with the MinGram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetMinGramOk() (*int, bool) {
	if o == nil || IsNil(o.MinGram) {
		return nil, false
	}
	return o.MinGram, true
}

// HasMinGram returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasMinGram() bool {
	if o != nil && !IsNil(o.MinGram) {
		return true
	}

	return false
}

// SetMinGram gets a reference to the given int and assigns it to the MinGram field.
func (o *FTSAnalyzersTokenizer) SetMinGram(v int) {
	o.MinGram = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FTSAnalyzersTokenizer) SetType(v string) {
	o.Type = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetGroup() int {
	if o == nil || IsNil(o.Group) {
		var ret int
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetGroupOk() (*int, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given int and assigns it to the Group field.
func (o *FTSAnalyzersTokenizer) SetGroup(v int) {
	o.Group = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *FTSAnalyzersTokenizer) SetPattern(v string) {
	o.Pattern = &v
}

// GetMaxTokenLength returns the MaxTokenLength field value if set, zero value otherwise.
func (o *FTSAnalyzersTokenizer) GetMaxTokenLength() int {
	if o == nil || IsNil(o.MaxTokenLength) {
		var ret int
		return ret
	}
	return *o.MaxTokenLength
}

// GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FTSAnalyzersTokenizer) GetMaxTokenLengthOk() (*int, bool) {
	if o == nil || IsNil(o.MaxTokenLength) {
		return nil, false
	}
	return o.MaxTokenLength, true
}

// HasMaxTokenLength returns a boolean if a field has been set.
func (o *FTSAnalyzersTokenizer) HasMaxTokenLength() bool {
	if o != nil && !IsNil(o.MaxTokenLength) {
		return true
	}

	return false
}

// SetMaxTokenLength gets a reference to the given int and assigns it to the MaxTokenLength field.
func (o *FTSAnalyzersTokenizer) SetMaxTokenLength(v int) {
	o.MaxTokenLength = &v
}

func (o FTSAnalyzersTokenizer) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FTSAnalyzersTokenizer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxGram) {
		toSerialize["maxGram"] = o.MaxGram
	}
	if !IsNil(o.MinGram) {
		toSerialize["minGram"] = o.MinGram
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.MaxTokenLength) {
		toSerialize["maxTokenLength"] = o.MaxTokenLength
	}
	return toSerialize, nil
}

type NullableFTSAnalyzersTokenizer struct {
	value *FTSAnalyzersTokenizer
	isSet bool
}

func (v NullableFTSAnalyzersTokenizer) Get() *FTSAnalyzersTokenizer {
	return v.value
}

func (v *NullableFTSAnalyzersTokenizer) Set(val *FTSAnalyzersTokenizer) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersTokenizer) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersTokenizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersTokenizer(val *FTSAnalyzersTokenizer) *NullableFTSAnalyzersTokenizer {
	return &NullableFTSAnalyzersTokenizer{value: val, isSet: true}
}

func (v NullableFTSAnalyzersTokenizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersTokenizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
