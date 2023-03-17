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

// checks if the TokenizerregexCaptureGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizerregexCaptureGroup{}

// TokenizerregexCaptureGroup Tokenizer that uses a regular expression pattern to extract tokens.
type TokenizerregexCaptureGroup struct {
	// Index of the character group within the matching expression to extract into tokens. Use `0` to extract all character groups.
	Group int32 `json:"group"`
	// Regular expression to match against.
	Pattern string `json:"pattern"`
	// Human-readable label that identifies this tokenizer type.
	Type string `json:"type"`
}

// NewTokenizerregexCaptureGroup instantiates a new TokenizerregexCaptureGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizerregexCaptureGroup(group int32, pattern string, type_ string) *TokenizerregexCaptureGroup {
	this := TokenizerregexCaptureGroup{}
	this.Group = group
	this.Pattern = pattern
	this.Type = type_
	return &this
}

// NewTokenizerregexCaptureGroupWithDefaults instantiates a new TokenizerregexCaptureGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizerregexCaptureGroupWithDefaults() *TokenizerregexCaptureGroup {
	this := TokenizerregexCaptureGroup{}
	return &this
}

// GetGroup returns the Group field value
func (o *TokenizerregexCaptureGroup) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *TokenizerregexCaptureGroup) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *TokenizerregexCaptureGroup) SetGroup(v int32) {
	o.Group = v
}

// GetPattern returns the Pattern field value
func (o *TokenizerregexCaptureGroup) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *TokenizerregexCaptureGroup) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *TokenizerregexCaptureGroup) SetPattern(v string) {
	o.Pattern = v
}

// GetType returns the Type field value
func (o *TokenizerregexCaptureGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenizerregexCaptureGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenizerregexCaptureGroup) SetType(v string) {
	o.Type = v
}

func (o TokenizerregexCaptureGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizerregexCaptureGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["pattern"] = o.Pattern
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenizerregexCaptureGroup struct {
	value *TokenizerregexCaptureGroup
	isSet bool
}

func (v NullableTokenizerregexCaptureGroup) Get() *TokenizerregexCaptureGroup {
	return v.value
}

func (v *NullableTokenizerregexCaptureGroup) Set(val *TokenizerregexCaptureGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizerregexCaptureGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizerregexCaptureGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizerregexCaptureGroup(val *TokenizerregexCaptureGroup) *NullableTokenizerregexCaptureGroup {
	return &NullableTokenizerregexCaptureGroup{value: val, isSet: true}
}

func (v NullableTokenizerregexCaptureGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizerregexCaptureGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


