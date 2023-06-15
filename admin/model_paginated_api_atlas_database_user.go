// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PaginatedApiAtlasDatabaseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiAtlasDatabaseUser{}

// PaginatedApiAtlasDatabaseUser List of MongoDB Database users granted access to databases in the specified project.
type PaginatedApiAtlasDatabaseUser struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []CloudDatabaseUser `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiAtlasDatabaseUser instantiates a new PaginatedApiAtlasDatabaseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiAtlasDatabaseUser() *PaginatedApiAtlasDatabaseUser {
	this := PaginatedApiAtlasDatabaseUser{}
	return &this
}

// NewPaginatedApiAtlasDatabaseUserWithDefaults instantiates a new PaginatedApiAtlasDatabaseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiAtlasDatabaseUserWithDefaults() *PaginatedApiAtlasDatabaseUser {
	this := PaginatedApiAtlasDatabaseUser{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiAtlasDatabaseUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasDatabaseUser) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiAtlasDatabaseUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiAtlasDatabaseUser) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiAtlasDatabaseUser) GetResults() []CloudDatabaseUser {
	if o == nil || IsNil(o.Results) {
		var ret []CloudDatabaseUser
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasDatabaseUser) GetResultsOk() ([]CloudDatabaseUser, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiAtlasDatabaseUser) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CloudDatabaseUser and assigns it to the Results field.
func (o *PaginatedApiAtlasDatabaseUser) SetResults(v []CloudDatabaseUser) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiAtlasDatabaseUser) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasDatabaseUser) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiAtlasDatabaseUser) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiAtlasDatabaseUser) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiAtlasDatabaseUser) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiAtlasDatabaseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedApiAtlasDatabaseUser struct {
	value *PaginatedApiAtlasDatabaseUser
	isSet bool
}

func (v NullablePaginatedApiAtlasDatabaseUser) Get() *PaginatedApiAtlasDatabaseUser {
	return v.value
}

func (v *NullablePaginatedApiAtlasDatabaseUser) Set(val *PaginatedApiAtlasDatabaseUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiAtlasDatabaseUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiAtlasDatabaseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiAtlasDatabaseUser(val *PaginatedApiAtlasDatabaseUser) *NullablePaginatedApiAtlasDatabaseUser {
	return &NullablePaginatedApiAtlasDatabaseUser{value: val, isSet: true}
}

func (v NullablePaginatedApiAtlasDatabaseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiAtlasDatabaseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
