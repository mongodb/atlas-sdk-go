// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PaginatedCloudProviderContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCloudProviderContainer{}

// PaginatedCloudProviderContainer List of Network Peering Containers that Amazon Web Services serves.
type PaginatedCloudProviderContainer struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []CloudProviderContainer `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedCloudProviderContainer instantiates a new PaginatedCloudProviderContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudProviderContainer() *PaginatedCloudProviderContainer {
	this := PaginatedCloudProviderContainer{}
	return &this
}

// NewPaginatedCloudProviderContainerWithDefaults instantiates a new PaginatedCloudProviderContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudProviderContainerWithDefaults() *PaginatedCloudProviderContainer {
	this := PaginatedCloudProviderContainer{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainer) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainer) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedCloudProviderContainer) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainer) GetResults() []CloudProviderContainer {
	if o == nil || IsNil(o.Results) {
		var ret []CloudProviderContainer
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainer) GetResultsOk() ([]CloudProviderContainer, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainer) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CloudProviderContainer and assigns it to the Results field.
func (o *PaginatedCloudProviderContainer) SetResults(v []CloudProviderContainer) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainer) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainer) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainer) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedCloudProviderContainer) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedCloudProviderContainer) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedCloudProviderContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedCloudProviderContainer struct {
	value *PaginatedCloudProviderContainer
	isSet bool
}

func (v NullablePaginatedCloudProviderContainer) Get() *PaginatedCloudProviderContainer {
	return v.value
}

func (v *NullablePaginatedCloudProviderContainer) Set(val *PaginatedCloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudProviderContainer(val *PaginatedCloudProviderContainer) *NullablePaginatedCloudProviderContainer {
	return &NullablePaginatedCloudProviderContainer{value: val, isSet: true}
}

func (v NullablePaginatedCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
