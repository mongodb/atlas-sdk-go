// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the PaginatedTeamRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedTeamRole{}

// PaginatedTeamRole struct for PaginatedTeamRole
type PaginatedTeamRole struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []TeamRole `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedTeamRole instantiates a new PaginatedTeamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTeamRole() *PaginatedTeamRole {
	this := PaginatedTeamRole{}
	return &this
}

// NewPaginatedTeamRoleWithDefaults instantiates a new PaginatedTeamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTeamRoleWithDefaults() *PaginatedTeamRole {
	this := PaginatedTeamRole{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedTeamRole) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamRole) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedTeamRole) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedTeamRole) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTeamRole) GetResults() []TeamRole {
	if o == nil || IsNil(o.Results) {
		var ret []TeamRole
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamRole) GetResultsOk() ([]TeamRole, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTeamRole) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TeamRole and assigns it to the Results field.
func (o *PaginatedTeamRole) SetResults(v []TeamRole) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedTeamRole) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTeamRole) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedTeamRole) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedTeamRole) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedTeamRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedTeamRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedTeamRole struct {
	value *PaginatedTeamRole
	isSet bool
}

func (v NullablePaginatedTeamRole) Get() *PaginatedTeamRole {
	return v.value
}

func (v *NullablePaginatedTeamRole) Set(val *PaginatedTeamRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTeamRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTeamRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTeamRole(val *PaginatedTeamRole) *NullablePaginatedTeamRole {
	return &NullablePaginatedTeamRole{value: val, isSet: true}
}

func (v NullablePaginatedTeamRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTeamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
