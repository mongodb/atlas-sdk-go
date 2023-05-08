/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the PaginatedAlertConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAlertConfig{}

// PaginatedAlertConfig struct for PaginatedAlertConfig
type PaginatedAlertConfig struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []AlertConfigViewForNdsGroup `json:"results,omitempty"`
	// Number of documents returned in this response if **includeCount** query param is true.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedAlertConfig instantiates a new PaginatedAlertConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAlertConfig() *PaginatedAlertConfig {
	this := PaginatedAlertConfig{}
	return &this
}

// NewPaginatedAlertConfigWithDefaults instantiates a new PaginatedAlertConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAlertConfigWithDefaults() *PaginatedAlertConfig {
	this := PaginatedAlertConfig{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedAlertConfig) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfig) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedAlertConfig) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedAlertConfig) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedAlertConfig) GetResults() []AlertConfigViewForNdsGroup {
	if o == nil || IsNil(o.Results) {
		var ret []AlertConfigViewForNdsGroup
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfig) GetResultsOk() ([]AlertConfigViewForNdsGroup, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedAlertConfig) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AlertConfigViewForNdsGroup and assigns it to the Results field.
func (o *PaginatedAlertConfig) SetResults(v []AlertConfigViewForNdsGroup) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedAlertConfig) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfig) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedAlertConfig) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedAlertConfig) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedAlertConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedAlertConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedAlertConfig struct {
	value *PaginatedAlertConfig
	isSet bool
}

func (v NullablePaginatedAlertConfig) Get() *PaginatedAlertConfig {
	return v.value
}

func (v *NullablePaginatedAlertConfig) Set(val *PaginatedAlertConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAlertConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAlertConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAlertConfig(val *PaginatedAlertConfig) *NullablePaginatedAlertConfig {
	return &NullablePaginatedAlertConfig{value: val, isSet: true}
}

func (v NullablePaginatedAlertConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAlertConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


