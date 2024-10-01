// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PaginatedPublicApiUsageDetailsLineItem struct for PaginatedPublicApiUsageDetailsLineItem
type PaginatedPublicApiUsageDetailsLineItem struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results *[]PublicApiUsageDetailsLineItem `json:"results,omitempty"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedPublicApiUsageDetailsLineItem instantiates a new PaginatedPublicApiUsageDetailsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPublicApiUsageDetailsLineItem() *PaginatedPublicApiUsageDetailsLineItem {
	this := PaginatedPublicApiUsageDetailsLineItem{}
	return &this
}

// NewPaginatedPublicApiUsageDetailsLineItemWithDefaults instantiates a new PaginatedPublicApiUsageDetailsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPublicApiUsageDetailsLineItemWithDefaults() *PaginatedPublicApiUsageDetailsLineItem {
	this := PaginatedPublicApiUsageDetailsLineItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedPublicApiUsageDetailsLineItem) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedPublicApiUsageDetailsLineItem) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedPublicApiUsageDetailsLineItem) GetResults() []PublicApiUsageDetailsLineItem {
	if o == nil || IsNil(o.Results) {
		var ret []PublicApiUsageDetailsLineItem
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) GetResultsOk() (*[]PublicApiUsageDetailsLineItem, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PublicApiUsageDetailsLineItem and assigns it to the Results field.
func (o *PaginatedPublicApiUsageDetailsLineItem) SetResults(v []PublicApiUsageDetailsLineItem) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedPublicApiUsageDetailsLineItem) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedPublicApiUsageDetailsLineItem) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedPublicApiUsageDetailsLineItem) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o *PaginatedPublicApiUsageDetailsLineItem) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *PaginatedPublicApiUsageDetailsLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
