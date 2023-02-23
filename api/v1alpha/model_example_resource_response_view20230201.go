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

// checks if the ExampleResourceResponseView20230201 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExampleResourceResponseView20230201{}

// ExampleResourceResponseView20230201 struct for ExampleResourceResponseView20230201
type ExampleResourceResponseView20230201 struct {
	// Dummy additional field added to the response.
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	// Array that contains the dummy metadata.
	Data []string `json:"data,omitempty"`
	// Dummy description added as response.
	Description string `json:"description"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewExampleResourceResponseView20230201 instantiates a new ExampleResourceResponseView20230201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleResourceResponseView20230201() *ExampleResourceResponseView20230201 {
	this := ExampleResourceResponseView20230201{}
	return &this
}

// NewExampleResourceResponseView20230201WithDefaults instantiates a new ExampleResourceResponseView20230201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleResourceResponseView20230201WithDefaults() *ExampleResourceResponseView20230201 {
	this := ExampleResourceResponseView20230201{}
	return &this
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *ExampleResourceResponseView20230201) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230201) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *ExampleResourceResponseView20230201) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *ExampleResourceResponseView20230201) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExampleResourceResponseView20230201) GetData() []string {
	if o == nil || IsNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230201) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExampleResourceResponseView20230201) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ExampleResourceResponseView20230201) SetData(v []string) {
	o.Data = v
}

// GetDescription returns the Description field value
func (o *ExampleResourceResponseView20230201) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230201) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExampleResourceResponseView20230201) SetDescription(v string) {
	o.Description = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExampleResourceResponseView20230201) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230201) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExampleResourceResponseView20230201) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ExampleResourceResponseView20230201) SetLinks(v []Link) {
	o.Links = v
}

func (o ExampleResourceResponseView20230201) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExampleResourceResponseView20230201) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["description"] = o.Description
	// skip: links is readOnly
	return toSerialize, nil
}

type NullableExampleResourceResponseView20230201 struct {
	value *ExampleResourceResponseView20230201
	isSet bool
}

func (v NullableExampleResourceResponseView20230201) Get() *ExampleResourceResponseView20230201 {
	return v.value
}

func (v *NullableExampleResourceResponseView20230201) Set(val *ExampleResourceResponseView20230201) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleResourceResponseView20230201) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleResourceResponseView20230201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleResourceResponseView20230201(val *ExampleResourceResponseView20230201) *NullableExampleResourceResponseView20230201 {
	return &NullableExampleResourceResponseView20230201{value: val, isSet: true}
}

func (v NullableExampleResourceResponseView20230201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleResourceResponseView20230201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


