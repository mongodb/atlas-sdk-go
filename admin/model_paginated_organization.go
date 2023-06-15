// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PaginatedOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedOrganization{}

// PaginatedOrganization struct for PaginatedOrganization
type PaginatedOrganization struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []Organization `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedOrganization instantiates a new PaginatedOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOrganization() *PaginatedOrganization {
	this := PaginatedOrganization{}
	return &this
}

// NewPaginatedOrganizationWithDefaults instantiates a new PaginatedOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOrganizationWithDefaults() *PaginatedOrganization {
	this := PaginatedOrganization{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedOrganization) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOrganization) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedOrganization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedOrganization) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedOrganization) GetResults() []Organization {
	if o == nil || IsNil(o.Results) {
		var ret []Organization
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOrganization) GetResultsOk() ([]Organization, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedOrganization) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Organization and assigns it to the Results field.
func (o *PaginatedOrganization) SetResults(v []Organization) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedOrganization) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOrganization) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedOrganization) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedOrganization) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedOrganization) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedOrganization struct {
	value *PaginatedOrganization
	isSet bool
}

func (v NullablePaginatedOrganization) Get() *PaginatedOrganization {
	return v.value
}

func (v *NullablePaginatedOrganization) Set(val *PaginatedOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOrganization(val *PaginatedOrganization) *NullablePaginatedOrganization {
	return &NullablePaginatedOrganization{value: val, isSet: true}
}

func (v NullablePaginatedOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
