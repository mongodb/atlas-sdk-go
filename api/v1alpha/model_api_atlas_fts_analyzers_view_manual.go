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

// checks if the ApiAtlasFTSAnalyzersViewManual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAtlasFTSAnalyzersViewManual{}

// ApiAtlasFTSAnalyzersViewManual Settings that describe one Atlas Search custom analyzer.
type ApiAtlasFTSAnalyzersViewManual struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters []ApiAtlasFTSAnalyzersViewManualCharFiltersInner `json:"charFilters,omitempty"`
	// Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - `lucene.` - `builtin.` - `mongodb.`
	Name string `json:"name"`
	// Filter that performs operations such as:  - Stemming, which reduces related words, such as \"talking\", \"talked\", and \"talks\" to their root word \"talk\".  - Redaction, the removal of sensitive information from public documents.
	TokenFilters []ApiAtlasFTSAnalyzersViewManualTokenFiltersInner `json:"tokenFilters,omitempty"`
	Tokenizer ApiAtlasFTSAnalyzersViewManualTokenizer `json:"tokenizer"`
}

// NewApiAtlasFTSAnalyzersViewManual instantiates a new ApiAtlasFTSAnalyzersViewManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFTSAnalyzersViewManual() *ApiAtlasFTSAnalyzersViewManual {
	this := ApiAtlasFTSAnalyzersViewManual{}
	return &this
}

// NewApiAtlasFTSAnalyzersViewManualWithDefaults instantiates a new ApiAtlasFTSAnalyzersViewManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFTSAnalyzersViewManualWithDefaults() *ApiAtlasFTSAnalyzersViewManual {
	this := ApiAtlasFTSAnalyzersViewManual{}
	return &this
}

// GetCharFilters returns the CharFilters field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersViewManual) GetCharFilters() []ApiAtlasFTSAnalyzersViewManualCharFiltersInner {
	if o == nil || IsNil(o.CharFilters) {
		var ret []ApiAtlasFTSAnalyzersViewManualCharFiltersInner
		return ret
	}
	return o.CharFilters
}

// GetCharFiltersOk returns a tuple with the CharFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) GetCharFiltersOk() ([]ApiAtlasFTSAnalyzersViewManualCharFiltersInner, bool) {
	if o == nil || IsNil(o.CharFilters) {
		return nil, false
	}
	return o.CharFilters, true
}

// HasCharFilters returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) HasCharFilters() bool {
	if o != nil && !IsNil(o.CharFilters) {
		return true
	}

	return false
}

// SetCharFilters gets a reference to the given []ApiAtlasFTSAnalyzersViewManualCharFiltersInner and assigns it to the CharFilters field.
func (o *ApiAtlasFTSAnalyzersViewManual) SetCharFilters(v []ApiAtlasFTSAnalyzersViewManualCharFiltersInner) {
	o.CharFilters = v
}

// GetName returns the Name field value
func (o *ApiAtlasFTSAnalyzersViewManual) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasFTSAnalyzersViewManual) SetName(v string) {
	o.Name = v
}

// GetTokenFilters returns the TokenFilters field value if set, zero value otherwise.
func (o *ApiAtlasFTSAnalyzersViewManual) GetTokenFilters() []ApiAtlasFTSAnalyzersViewManualTokenFiltersInner {
	if o == nil || IsNil(o.TokenFilters) {
		var ret []ApiAtlasFTSAnalyzersViewManualTokenFiltersInner
		return ret
	}
	return o.TokenFilters
}

// GetTokenFiltersOk returns a tuple with the TokenFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) GetTokenFiltersOk() ([]ApiAtlasFTSAnalyzersViewManualTokenFiltersInner, bool) {
	if o == nil || IsNil(o.TokenFilters) {
		return nil, false
	}
	return o.TokenFilters, true
}

// HasTokenFilters returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) HasTokenFilters() bool {
	if o != nil && !IsNil(o.TokenFilters) {
		return true
	}

	return false
}

// SetTokenFilters gets a reference to the given []ApiAtlasFTSAnalyzersViewManualTokenFiltersInner and assigns it to the TokenFilters field.
func (o *ApiAtlasFTSAnalyzersViewManual) SetTokenFilters(v []ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) {
	o.TokenFilters = v
}

// GetTokenizer returns the Tokenizer field value
func (o *ApiAtlasFTSAnalyzersViewManual) GetTokenizer() ApiAtlasFTSAnalyzersViewManualTokenizer {
	if o == nil {
		var ret ApiAtlasFTSAnalyzersViewManualTokenizer
		return ret
	}

	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzersViewManual) GetTokenizerOk() (*ApiAtlasFTSAnalyzersViewManualTokenizer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokenizer, true
}

// SetTokenizer sets field value
func (o *ApiAtlasFTSAnalyzersViewManual) SetTokenizer(v ApiAtlasFTSAnalyzersViewManualTokenizer) {
	o.Tokenizer = v
}

func (o ApiAtlasFTSAnalyzersViewManual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAtlasFTSAnalyzersViewManual) ToMap() (map[string]interface{}, error) {
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

type NullableApiAtlasFTSAnalyzersViewManual struct {
	value *ApiAtlasFTSAnalyzersViewManual
	isSet bool
}

func (v NullableApiAtlasFTSAnalyzersViewManual) Get() *ApiAtlasFTSAnalyzersViewManual {
	return v.value
}

func (v *NullableApiAtlasFTSAnalyzersViewManual) Set(val *ApiAtlasFTSAnalyzersViewManual) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasFTSAnalyzersViewManual) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasFTSAnalyzersViewManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasFTSAnalyzersViewManual(val *ApiAtlasFTSAnalyzersViewManual) *NullableApiAtlasFTSAnalyzersViewManual {
	return &NullableApiAtlasFTSAnalyzersViewManual{value: val, isSet: true}
}

func (v NullableApiAtlasFTSAnalyzersViewManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasFTSAnalyzersViewManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


