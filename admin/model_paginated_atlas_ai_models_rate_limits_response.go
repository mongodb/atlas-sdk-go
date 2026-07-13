// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PaginatedAtlasAiModelsRateLimitsResponse List response for AI Model Rate Limits at the organization and project level.
type PaginatedAtlasAiModelsRateLimitsResponse struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results []AiModelsRateLimitResponse `json:"results"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`. The total number is an estimate and may not be exact.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedAtlasAiModelsRateLimitsResponse instantiates a new PaginatedAtlasAiModelsRateLimitsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAtlasAiModelsRateLimitsResponse(results []AiModelsRateLimitResponse) *PaginatedAtlasAiModelsRateLimitsResponse {
	this := PaginatedAtlasAiModelsRateLimitsResponse{}
	this.Results = results
	return &this
}

// NewPaginatedAtlasAiModelsRateLimitsResponseWithDefaults instantiates a new PaginatedAtlasAiModelsRateLimitsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAtlasAiModelsRateLimitsResponseWithDefaults() *PaginatedAtlasAiModelsRateLimitsResponse {
	this := PaginatedAtlasAiModelsRateLimitsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetResults() []AiModelsRateLimitResponse {
	if o == nil {
		var ret []AiModelsRateLimitResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetResultsOk() (*[]AiModelsRateLimitResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedAtlasAiModelsRateLimitsResponse) SetResults(v []AiModelsRateLimitResponse) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedAtlasAiModelsRateLimitsResponse) SetTotalCount(v int) {
	o.TotalCount = &v
}
