// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PaginatedFederationIdentityProvider struct for PaginatedFederationIdentityProvider
type PaginatedFederationIdentityProvider struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results *[]FederationIdentityProvider `json:"results,omitempty"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`. The total number is an estimate and may not be exact.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedFederationIdentityProvider instantiates a new PaginatedFederationIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedFederationIdentityProvider() *PaginatedFederationIdentityProvider {
	this := PaginatedFederationIdentityProvider{}
	return &this
}

// NewPaginatedFederationIdentityProviderWithDefaults instantiates a new PaginatedFederationIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedFederationIdentityProviderWithDefaults() *PaginatedFederationIdentityProvider {
	this := PaginatedFederationIdentityProvider{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedFederationIdentityProvider) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedFederationIdentityProvider) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedFederationIdentityProvider) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedFederationIdentityProvider) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedFederationIdentityProvider) GetResults() []FederationIdentityProvider {
	if o == nil || IsNil(o.Results) {
		var ret []FederationIdentityProvider
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedFederationIdentityProvider) GetResultsOk() (*[]FederationIdentityProvider, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedFederationIdentityProvider) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FederationIdentityProvider and assigns it to the Results field.
func (o *PaginatedFederationIdentityProvider) SetResults(v []FederationIdentityProvider) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedFederationIdentityProvider) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedFederationIdentityProvider) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedFederationIdentityProvider) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedFederationIdentityProvider) SetTotalCount(v int) {
	o.TotalCount = &v
}
