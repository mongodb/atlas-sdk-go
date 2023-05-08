/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the PaginatedClusterDescriptionV15 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedClusterDescriptionV15{}

// PaginatedClusterDescriptionV15 struct for PaginatedClusterDescriptionV15
type PaginatedClusterDescriptionV15 struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ClusterDescriptionV15 `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedClusterDescriptionV15 instantiates a new PaginatedClusterDescriptionV15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedClusterDescriptionV15() *PaginatedClusterDescriptionV15 {
	this := PaginatedClusterDescriptionV15{}
	return &this
}

// NewPaginatedClusterDescriptionV15WithDefaults instantiates a new PaginatedClusterDescriptionV15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedClusterDescriptionV15WithDefaults() *PaginatedClusterDescriptionV15 {
	this := PaginatedClusterDescriptionV15{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedClusterDescriptionV15) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedClusterDescriptionV15) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedClusterDescriptionV15) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedClusterDescriptionV15) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedClusterDescriptionV15) GetResults() []ClusterDescriptionV15 {
	if o == nil || IsNil(o.Results) {
		var ret []ClusterDescriptionV15
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedClusterDescriptionV15) GetResultsOk() ([]ClusterDescriptionV15, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedClusterDescriptionV15) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterDescriptionV15 and assigns it to the Results field.
func (o *PaginatedClusterDescriptionV15) SetResults(v []ClusterDescriptionV15) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedClusterDescriptionV15) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedClusterDescriptionV15) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedClusterDescriptionV15) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedClusterDescriptionV15) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedClusterDescriptionV15) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedClusterDescriptionV15) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedClusterDescriptionV15 struct {
	value *PaginatedClusterDescriptionV15
	isSet bool
}

func (v NullablePaginatedClusterDescriptionV15) Get() *PaginatedClusterDescriptionV15 {
	return v.value
}

func (v *NullablePaginatedClusterDescriptionV15) Set(val *PaginatedClusterDescriptionV15) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedClusterDescriptionV15) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedClusterDescriptionV15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedClusterDescriptionV15(val *PaginatedClusterDescriptionV15) *NullablePaginatedClusterDescriptionV15 {
	return &NullablePaginatedClusterDescriptionV15{value: val, isSet: true}
}

func (v NullablePaginatedClusterDescriptionV15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedClusterDescriptionV15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


