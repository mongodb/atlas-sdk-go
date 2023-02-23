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

// checks if the CharFilterhtmlStrip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharFilterhtmlStrip{}

// CharFilterhtmlStrip Filter that strips out HTML constructs.
type CharFilterhtmlStrip struct {
	// The HTML tags that you want to exclude from filtering.
	IgnoredTags []string `json:"ignoredTags,omitempty"`
	// Human-readable label that identifies this character filter type.
	Type string `json:"type"`
}

// NewCharFilterhtmlStrip instantiates a new CharFilterhtmlStrip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharFilterhtmlStrip() *CharFilterhtmlStrip {
	this := CharFilterhtmlStrip{}
	return &this
}

// NewCharFilterhtmlStripWithDefaults instantiates a new CharFilterhtmlStrip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharFilterhtmlStripWithDefaults() *CharFilterhtmlStrip {
	this := CharFilterhtmlStrip{}
	return &this
}

// GetIgnoredTags returns the IgnoredTags field value if set, zero value otherwise.
func (o *CharFilterhtmlStrip) GetIgnoredTags() []string {
	if o == nil || IsNil(o.IgnoredTags) {
		var ret []string
		return ret
	}
	return o.IgnoredTags
}

// GetIgnoredTagsOk returns a tuple with the IgnoredTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharFilterhtmlStrip) GetIgnoredTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoredTags) {
		return nil, false
	}
	return o.IgnoredTags, true
}

// HasIgnoredTags returns a boolean if a field has been set.
func (o *CharFilterhtmlStrip) HasIgnoredTags() bool {
	if o != nil && !IsNil(o.IgnoredTags) {
		return true
	}

	return false
}

// SetIgnoredTags gets a reference to the given []string and assigns it to the IgnoredTags field.
func (o *CharFilterhtmlStrip) SetIgnoredTags(v []string) {
	o.IgnoredTags = v
}

// GetType returns the Type field value
func (o *CharFilterhtmlStrip) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharFilterhtmlStrip) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharFilterhtmlStrip) SetType(v string) {
	o.Type = v
}

func (o CharFilterhtmlStrip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharFilterhtmlStrip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoredTags) {
		toSerialize["ignoredTags"] = o.IgnoredTags
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCharFilterhtmlStrip struct {
	value *CharFilterhtmlStrip
	isSet bool
}

func (v NullableCharFilterhtmlStrip) Get() *CharFilterhtmlStrip {
	return v.value
}

func (v *NullableCharFilterhtmlStrip) Set(val *CharFilterhtmlStrip) {
	v.value = val
	v.isSet = true
}

func (v NullableCharFilterhtmlStrip) IsSet() bool {
	return v.isSet
}

func (v *NullableCharFilterhtmlStrip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharFilterhtmlStrip(val *CharFilterhtmlStrip) *NullableCharFilterhtmlStrip {
	return &NullableCharFilterhtmlStrip{value: val, isSet: true}
}

func (v NullableCharFilterhtmlStrip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharFilterhtmlStrip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


