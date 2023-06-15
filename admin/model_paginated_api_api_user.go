// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PaginatedApiApiUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiApiUser{}

// PaginatedApiApiUser struct for PaginatedApiApiUser
type PaginatedApiApiUser struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []KeyUser `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiApiUser instantiates a new PaginatedApiApiUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiApiUser() *PaginatedApiApiUser {
	this := PaginatedApiApiUser{}
	return &this
}

// NewPaginatedApiApiUserWithDefaults instantiates a new PaginatedApiApiUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiApiUserWithDefaults() *PaginatedApiApiUser {
	this := PaginatedApiApiUser{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiApiUser) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetResults() []KeyUser {
	if o == nil || IsNil(o.Results) {
		var ret []KeyUser
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetResultsOk() ([]KeyUser, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []KeyUser and assigns it to the Results field.
func (o *PaginatedApiApiUser) SetResults(v []KeyUser) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiApiUser) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiApiUser) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiApiUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedApiApiUser struct {
	value *PaginatedApiApiUser
	isSet bool
}

func (v NullablePaginatedApiApiUser) Get() *PaginatedApiApiUser {
	return v.value
}

func (v *NullablePaginatedApiApiUser) Set(val *PaginatedApiApiUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiApiUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiApiUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiApiUser(val *PaginatedApiApiUser) *NullablePaginatedApiApiUser {
	return &NullablePaginatedApiApiUser{value: val, isSet: true}
}

func (v NullablePaginatedApiApiUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiApiUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
