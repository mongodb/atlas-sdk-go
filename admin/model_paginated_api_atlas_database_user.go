// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PaginatedApiAtlasDatabaseUser List of MongoDB Database users granted access to databases in the specified project.
type PaginatedApiAtlasDatabaseUser struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results []CloudDatabaseUser `json:"results"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`. The total number is an estimate and may not be exact.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PaginatedApiAtlasDatabaseUser) MarshalJSON() ([]byte, error) {
	type noMethod PaginatedApiAtlasDatabaseUser
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPaginatedApiAtlasDatabaseUser instantiates a new PaginatedApiAtlasDatabaseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiAtlasDatabaseUser(results []CloudDatabaseUser) *PaginatedApiAtlasDatabaseUser {
	this := PaginatedApiAtlasDatabaseUser{}
	this.Results = results
	return &this
}

// NewPaginatedApiAtlasDatabaseUserWithDefaults instantiates a new PaginatedApiAtlasDatabaseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiAtlasDatabaseUserWithDefaults() *PaginatedApiAtlasDatabaseUser {
	this := PaginatedApiAtlasDatabaseUser{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedApiAtlasDatabaseUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasDatabaseUser) GetLinksOk() (*[]Link, bool) {
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
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *PaginatedApiAtlasDatabaseUser) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetResults returns the Results field value
func (o *PaginatedApiAtlasDatabaseUser) GetResults() []CloudDatabaseUser {
	if o == nil {
		var ret []CloudDatabaseUser
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedApiAtlasDatabaseUser) GetResultsOk() (*[]CloudDatabaseUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedApiAtlasDatabaseUser) SetResults(v []CloudDatabaseUser) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
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
	o.NullFields = removeNullField(o.NullFields, "TotalCount")
}

// SetTotalCountNil sets TotalCount to an explicit JSON null when marshaled.
func (o *PaginatedApiAtlasDatabaseUser) SetTotalCountNil() {
	o.TotalCount = nil
	o.NullFields = addNullField(o.NullFields, "TotalCount")
}
