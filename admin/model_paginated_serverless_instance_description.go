// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PaginatedServerlessInstanceDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedServerlessInstanceDescription{}

// PaginatedServerlessInstanceDescription struct for PaginatedServerlessInstanceDescription
type PaginatedServerlessInstanceDescription struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ServerlessInstanceDescription `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedServerlessInstanceDescription instantiates a new PaginatedServerlessInstanceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedServerlessInstanceDescription() *PaginatedServerlessInstanceDescription {
	this := PaginatedServerlessInstanceDescription{}
	return &this
}

// NewPaginatedServerlessInstanceDescriptionWithDefaults instantiates a new PaginatedServerlessInstanceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedServerlessInstanceDescriptionWithDefaults() *PaginatedServerlessInstanceDescription {
	this := PaginatedServerlessInstanceDescription{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedServerlessInstanceDescription) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessInstanceDescription) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedServerlessInstanceDescription) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedServerlessInstanceDescription) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedServerlessInstanceDescription) GetResults() []ServerlessInstanceDescription {
	if o == nil || IsNil(o.Results) {
		var ret []ServerlessInstanceDescription
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessInstanceDescription) GetResultsOk() ([]ServerlessInstanceDescription, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedServerlessInstanceDescription) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ServerlessInstanceDescription and assigns it to the Results field.
func (o *PaginatedServerlessInstanceDescription) SetResults(v []ServerlessInstanceDescription) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedServerlessInstanceDescription) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessInstanceDescription) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedServerlessInstanceDescription) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedServerlessInstanceDescription) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedServerlessInstanceDescription) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedServerlessInstanceDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedServerlessInstanceDescription struct {
	value *PaginatedServerlessInstanceDescription
	isSet bool
}

func (v NullablePaginatedServerlessInstanceDescription) Get() *PaginatedServerlessInstanceDescription {
	return v.value
}

func (v *NullablePaginatedServerlessInstanceDescription) Set(val *PaginatedServerlessInstanceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedServerlessInstanceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedServerlessInstanceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedServerlessInstanceDescription(val *PaginatedServerlessInstanceDescription) *NullablePaginatedServerlessInstanceDescription {
	return &NullablePaginatedServerlessInstanceDescription{value: val, isSet: true}
}

func (v NullablePaginatedServerlessInstanceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedServerlessInstanceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
