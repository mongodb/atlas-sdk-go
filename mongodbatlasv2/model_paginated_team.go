/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the PaginatedTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedTeam{}

// PaginatedTeam struct for PaginatedTeam
type PaginatedTeam struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []TeamResponse `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedTeam instantiates a new PaginatedTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTeam() *PaginatedTeam {
	this := PaginatedTeam{}
	return &this
}

// NewPaginatedTeamWithDefaults instantiates a new PaginatedTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTeamWithDefaults() *PaginatedTeam {
	this := PaginatedTeam{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedTeam) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeam) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedTeam) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedTeam) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTeam) GetResults() []TeamResponse {
	if o == nil || IsNil(o.Results) {
		var ret []TeamResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeam) GetResultsOk() ([]TeamResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTeam) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TeamResponse and assigns it to the Results field.
func (o *PaginatedTeam) SetResults(v []TeamResponse) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedTeam) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeam) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedTeam) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedTeam) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedTeam) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedTeam struct {
	value *PaginatedTeam
	isSet bool
}

func (v NullablePaginatedTeam) Get() *PaginatedTeam {
	return v.value
}

func (v *NullablePaginatedTeam) Set(val *PaginatedTeam) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTeam) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTeam(val *PaginatedTeam) *NullablePaginatedTeam {
	return &NullablePaginatedTeam{value: val, isSet: true}
}

func (v NullablePaginatedTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


