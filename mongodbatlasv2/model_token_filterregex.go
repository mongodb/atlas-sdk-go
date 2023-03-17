/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the TokenFilterregex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFilterregex{}

// TokenFilterregex Filter that applies a regular expression to each token, replacing matches with a specified string.
type TokenFilterregex struct {
	// Value that indicates whether to replace only the first matching pattern or all matching patterns.
	Matches string `json:"matches"`
	// Regular expression pattern to apply to each token.
	Pattern string `json:"pattern"`
	// Replacement string to substitute wherever a matching pattern occurs.
	Replacement string `json:"replacement"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterregex instantiates a new TokenFilterregex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterregex(matches string, pattern string, replacement string, type_ string) *TokenFilterregex {
	this := TokenFilterregex{}
	this.Matches = matches
	this.Pattern = pattern
	this.Replacement = replacement
	this.Type = type_
	return &this
}

// NewTokenFilterregexWithDefaults instantiates a new TokenFilterregex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterregexWithDefaults() *TokenFilterregex {
	this := TokenFilterregex{}
	return &this
}

// GetMatches returns the Matches field value
func (o *TokenFilterregex) GetMatches() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
func (o *TokenFilterregex) GetMatchesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matches, true
}

// SetMatches sets field value
func (o *TokenFilterregex) SetMatches(v string) {
	o.Matches = v
}

// GetPattern returns the Pattern field value
func (o *TokenFilterregex) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *TokenFilterregex) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *TokenFilterregex) SetPattern(v string) {
	o.Pattern = v
}

// GetReplacement returns the Replacement field value
func (o *TokenFilterregex) GetReplacement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value
// and a boolean to check if the value has been set.
func (o *TokenFilterregex) GetReplacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replacement, true
}

// SetReplacement sets field value
func (o *TokenFilterregex) SetReplacement(v string) {
	o.Replacement = v
}

// GetType returns the Type field value
func (o *TokenFilterregex) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterregex) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterregex) SetType(v string) {
	o.Type = v
}

func (o TokenFilterregex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenFilterregex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matches"] = o.Matches
	toSerialize["pattern"] = o.Pattern
	toSerialize["replacement"] = o.Replacement
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenFilterregex struct {
	value *TokenFilterregex
	isSet bool
}

func (v NullableTokenFilterregex) Get() *TokenFilterregex {
	return v.value
}

func (v *NullableTokenFilterregex) Set(val *TokenFilterregex) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterregex) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterregex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterregex(val *TokenFilterregex) *NullableTokenFilterregex {
	return &NullableTokenFilterregex{value: val, isSet: true}
}

func (v NullableTokenFilterregex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterregex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


