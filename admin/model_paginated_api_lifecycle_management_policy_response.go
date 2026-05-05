// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PaginatedApiLifecycleManagementPolicyResponse struct for PaginatedApiLifecycleManagementPolicyResponse
type PaginatedApiLifecycleManagementPolicyResponse struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results []ApiLifecycleManagementPolicyResponse `json:"results"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`. The total number is an estimate and may not be exact.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiLifecycleManagementPolicyResponse instantiates a new PaginatedApiLifecycleManagementPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiLifecycleManagementPolicyResponse(results []ApiLifecycleManagementPolicyResponse) *PaginatedApiLifecycleManagementPolicyResponse {
	this := PaginatedApiLifecycleManagementPolicyResponse{}
	this.Results = results
	return &this
}

// NewPaginatedApiLifecycleManagementPolicyResponseWithDefaults instantiates a new PaginatedApiLifecycleManagementPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiLifecycleManagementPolicyResponseWithDefaults() *PaginatedApiLifecycleManagementPolicyResponse {
	this := PaginatedApiLifecycleManagementPolicyResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiLifecycleManagementPolicyResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiLifecycleManagementPolicyResponse) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetResults() []ApiLifecycleManagementPolicyResponse {
	if o == nil {
		var ret []ApiLifecycleManagementPolicyResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetResultsOk() (*[]ApiLifecycleManagementPolicyResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedApiLifecycleManagementPolicyResponse) SetResults(v []ApiLifecycleManagementPolicyResponse) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiLifecycleManagementPolicyResponse) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiLifecycleManagementPolicyResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiLifecycleManagementPolicyResponse) SetTotalCount(v int) {
	o.TotalCount = &v
}
