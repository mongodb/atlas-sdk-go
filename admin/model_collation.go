// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Collation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Collation{}

// Collation One or more settings that specify language-specific rules to compare strings within this index.
type Collation struct {
	// Method to handle whitespace and punctuation as base characters for purposes of comparison.  | Value | Evaluate Whitespace and Punctuation as Base Characters | |---|---| | `\"non-ignorable\"` | Yes | | `\"shifted\"` | No. MongoDB Cloud distinguishes these characters when `\"strength\" > 3`. |
	Alternate *string `json:"alternate,omitempty"`
	// Flag that indicates whether strings with diacritics sort from back of the string. Some French dictionary orders strings in this way.  | Value | String Comparison Method |  |---|---| | `true` | Compare from back to front. | | `false` | Compare from front to back. |
	Backwards *bool `json:"backwards,omitempty"`
	// Method to handle sort order of case differences during tertiary level comparisons.  | Value | Sort Order Method |  |---|---|  | `\"upper\"` | Uppercase sorts before lowercase. |  | `\"lower\"` | Lowercase sorts before uppercase. |  | `\"off\"` | Similar to \"lower\" with slight differences. |
	CaseFirst *string `json:"caseFirst,omitempty"`
	// Flag that indicates whether to include case comparison when `\"strength\" : 1` or `\"strength\" : 2`.  | Value | Compare case at level 1 or 2? | Strength Level | Comparisons Include |  |---|---|---|---|  | true | Yes | 1 | Base characters and case. |  |  |  | 2 | Base characters, diacritics (and possible other secondary differences),   and case. |  | false | No |  |  |
	CaseLevel *bool `json:"caseLevel,omitempty"`
	// International Components for Unicode (ICU) code that represents a localized language. To specify simple binary comparison, set `\"locale\" : \"simple\"`.
	Locale string `json:"locale"`
	// Field that indicates which characters can be ignored when `\"alternate\" : \"shifted\"`. This has no affect if `\"alternate\" : \"non-ignorable\"`.  | Value | Ignore |  |---|---| | `\"punct\"` | Both whitespace and punctuation | | `\"space\"` | Whitespace |
	MaxVariable *string `json:"maxVariable,omitempty"`
	// Flag that indicates whether to check if the text requires normalization and then perform it. Most text doesn't require this normalization processing.  | Value | Normalization Method |  |---|---| | `true` | Yes, check if fully normalized and perform normalization to compare text. | | `false` | No, don't check. |
	Normalization *bool `json:"normalization,omitempty"`
	// Flag that indicates whether to compare sequences of digits as numbers or as strings.  | Value | Comparison Method |  |---|---| | `true` | Compare as numbers. This results in `10 > 2`. | | `false` | Compare as strings. This results in `\"10\" < \"2\"`. |
	NumericOrdering *bool `json:"numericOrdering,omitempty"`
	// Degree of comparison to perform when sorting words. MongoDB Cloud accepts the following values:  | Value | Comparison Level | Comparison Method | |---|---|---| | 1 | Primary | Compares the base characters only, ignoring other differences such as diacritics and case. | | 2 | Secondary | Compares base characters (primary) and diacritics (secondary). Primary differences take precedence over secondary differences. | | 3 | Tertiary | Compares base characters (primary), diacritics (secondary), and case and variants (tertiary). Differences between base characters takes precedence over secondary differences which take precedence over tertiary differences. | | 4 | Quaternary | Compares for the specific use case to consider punctuation when levels 1 through 3 ignore punctuation or for processing Japanese text. | | 5 | Identical | Compares for the specific use case of tie breaker. |
	Strength *int `json:"strength,omitempty"`
}

// NewCollation instantiates a new Collation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollation(locale string) *Collation {
	this := Collation{}
	var alternate string = "non-ignorable"
	this.Alternate = &alternate
	var backwards bool = false
	this.Backwards = &backwards
	var caseFirst string = "false"
	this.CaseFirst = &caseFirst
	var caseLevel bool = false
	this.CaseLevel = &caseLevel
	this.Locale = locale
	var normalization bool = false
	this.Normalization = &normalization
	var numericOrdering bool = false
	this.NumericOrdering = &numericOrdering
	var strength int = 3
	this.Strength = &strength
	return &this
}

// NewCollationWithDefaults instantiates a new Collation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollationWithDefaults() *Collation {
	this := Collation{}
	var alternate string = "non-ignorable"
	this.Alternate = &alternate
	var backwards bool = false
	this.Backwards = &backwards
	var caseFirst string = "false"
	this.CaseFirst = &caseFirst
	var caseLevel bool = false
	this.CaseLevel = &caseLevel
	var normalization bool = false
	this.Normalization = &normalization
	var numericOrdering bool = false
	this.NumericOrdering = &numericOrdering
	var strength int = 3
	this.Strength = &strength
	return &this
}

// GetAlternate returns the Alternate field value if set, zero value otherwise.
func (o *Collation) GetAlternate() string {
	if o == nil || IsNil(o.Alternate) {
		var ret string
		return ret
	}
	return *o.Alternate
}

// GetAlternateOk returns a tuple with the Alternate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetAlternateOk() (*string, bool) {
	if o == nil || IsNil(o.Alternate) {
		return nil, false
	}
	return o.Alternate, true
}

// HasAlternate returns a boolean if a field has been set.
func (o *Collation) HasAlternate() bool {
	if o != nil && !IsNil(o.Alternate) {
		return true
	}

	return false
}

// SetAlternate gets a reference to the given string and assigns it to the Alternate field.
func (o *Collation) SetAlternate(v string) {
	o.Alternate = &v
}

// GetBackwards returns the Backwards field value if set, zero value otherwise.
func (o *Collation) GetBackwards() bool {
	if o == nil || IsNil(o.Backwards) {
		var ret bool
		return ret
	}
	return *o.Backwards
}

// GetBackwardsOk returns a tuple with the Backwards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetBackwardsOk() (*bool, bool) {
	if o == nil || IsNil(o.Backwards) {
		return nil, false
	}
	return o.Backwards, true
}

// HasBackwards returns a boolean if a field has been set.
func (o *Collation) HasBackwards() bool {
	if o != nil && !IsNil(o.Backwards) {
		return true
	}

	return false
}

// SetBackwards gets a reference to the given bool and assigns it to the Backwards field.
func (o *Collation) SetBackwards(v bool) {
	o.Backwards = &v
}

// GetCaseFirst returns the CaseFirst field value if set, zero value otherwise.
func (o *Collation) GetCaseFirst() string {
	if o == nil || IsNil(o.CaseFirst) {
		var ret string
		return ret
	}
	return *o.CaseFirst
}

// GetCaseFirstOk returns a tuple with the CaseFirst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetCaseFirstOk() (*string, bool) {
	if o == nil || IsNil(o.CaseFirst) {
		return nil, false
	}
	return o.CaseFirst, true
}

// HasCaseFirst returns a boolean if a field has been set.
func (o *Collation) HasCaseFirst() bool {
	if o != nil && !IsNil(o.CaseFirst) {
		return true
	}

	return false
}

// SetCaseFirst gets a reference to the given string and assigns it to the CaseFirst field.
func (o *Collation) SetCaseFirst(v string) {
	o.CaseFirst = &v
}

// GetCaseLevel returns the CaseLevel field value if set, zero value otherwise.
func (o *Collation) GetCaseLevel() bool {
	if o == nil || IsNil(o.CaseLevel) {
		var ret bool
		return ret
	}
	return *o.CaseLevel
}

// GetCaseLevelOk returns a tuple with the CaseLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetCaseLevelOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseLevel) {
		return nil, false
	}
	return o.CaseLevel, true
}

// HasCaseLevel returns a boolean if a field has been set.
func (o *Collation) HasCaseLevel() bool {
	if o != nil && !IsNil(o.CaseLevel) {
		return true
	}

	return false
}

// SetCaseLevel gets a reference to the given bool and assigns it to the CaseLevel field.
func (o *Collation) SetCaseLevel(v bool) {
	o.CaseLevel = &v
}

// GetLocale returns the Locale field value
func (o *Collation) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *Collation) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *Collation) SetLocale(v string) {
	o.Locale = v
}

// GetMaxVariable returns the MaxVariable field value if set, zero value otherwise.
func (o *Collation) GetMaxVariable() string {
	if o == nil || IsNil(o.MaxVariable) {
		var ret string
		return ret
	}
	return *o.MaxVariable
}

// GetMaxVariableOk returns a tuple with the MaxVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetMaxVariableOk() (*string, bool) {
	if o == nil || IsNil(o.MaxVariable) {
		return nil, false
	}
	return o.MaxVariable, true
}

// HasMaxVariable returns a boolean if a field has been set.
func (o *Collation) HasMaxVariable() bool {
	if o != nil && !IsNil(o.MaxVariable) {
		return true
	}

	return false
}

// SetMaxVariable gets a reference to the given string and assigns it to the MaxVariable field.
func (o *Collation) SetMaxVariable(v string) {
	o.MaxVariable = &v
}

// GetNormalization returns the Normalization field value if set, zero value otherwise.
func (o *Collation) GetNormalization() bool {
	if o == nil || IsNil(o.Normalization) {
		var ret bool
		return ret
	}
	return *o.Normalization
}

// GetNormalizationOk returns a tuple with the Normalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetNormalizationOk() (*bool, bool) {
	if o == nil || IsNil(o.Normalization) {
		return nil, false
	}
	return o.Normalization, true
}

// HasNormalization returns a boolean if a field has been set.
func (o *Collation) HasNormalization() bool {
	if o != nil && !IsNil(o.Normalization) {
		return true
	}

	return false
}

// SetNormalization gets a reference to the given bool and assigns it to the Normalization field.
func (o *Collation) SetNormalization(v bool) {
	o.Normalization = &v
}

// GetNumericOrdering returns the NumericOrdering field value if set, zero value otherwise.
func (o *Collation) GetNumericOrdering() bool {
	if o == nil || IsNil(o.NumericOrdering) {
		var ret bool
		return ret
	}
	return *o.NumericOrdering
}

// GetNumericOrderingOk returns a tuple with the NumericOrdering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetNumericOrderingOk() (*bool, bool) {
	if o == nil || IsNil(o.NumericOrdering) {
		return nil, false
	}
	return o.NumericOrdering, true
}

// HasNumericOrdering returns a boolean if a field has been set.
func (o *Collation) HasNumericOrdering() bool {
	if o != nil && !IsNil(o.NumericOrdering) {
		return true
	}

	return false
}

// SetNumericOrdering gets a reference to the given bool and assigns it to the NumericOrdering field.
func (o *Collation) SetNumericOrdering(v bool) {
	o.NumericOrdering = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *Collation) GetStrength() int {
	if o == nil || IsNil(o.Strength) {
		var ret int
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collation) GetStrengthOk() (*int, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *Collation) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given int and assigns it to the Strength field.
func (o *Collation) SetStrength(v int) {
	o.Strength = &v
}

func (o Collation) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Collation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alternate) {
		toSerialize["alternate"] = o.Alternate
	}
	if !IsNil(o.Backwards) {
		toSerialize["backwards"] = o.Backwards
	}
	if !IsNil(o.CaseFirst) {
		toSerialize["caseFirst"] = o.CaseFirst
	}
	if !IsNil(o.CaseLevel) {
		toSerialize["caseLevel"] = o.CaseLevel
	}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.MaxVariable) {
		toSerialize["maxVariable"] = o.MaxVariable
	}
	if !IsNil(o.Normalization) {
		toSerialize["normalization"] = o.Normalization
	}
	if !IsNil(o.NumericOrdering) {
		toSerialize["numericOrdering"] = o.NumericOrdering
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	return toSerialize, nil
}

type NullableCollation struct {
	value *Collation
	isSet bool
}

func (v NullableCollation) Get() *Collation {
	return v.value
}

func (v *NullableCollation) Set(val *Collation) {
	v.value = val
	v.isSet = true
}

func (v NullableCollation) IsSet() bool {
	return v.isSet
}

func (v *NullableCollation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollation(val *Collation) *NullableCollation {
	return &NullableCollation{value: val, isSet: true}
}

func (v NullableCollation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
