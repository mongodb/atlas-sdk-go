// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PaginatedBackupSnapshotExportBuckets struct for PaginatedBackupSnapshotExportBuckets
type PaginatedBackupSnapshotExportBuckets struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	// Read only field.
	Results *[]DiskBackupSnapshotExportBucket `json:"results,omitempty"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedBackupSnapshotExportBuckets instantiates a new PaginatedBackupSnapshotExportBuckets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBackupSnapshotExportBuckets() *PaginatedBackupSnapshotExportBuckets {
	this := PaginatedBackupSnapshotExportBuckets{}
	return &this
}

// NewPaginatedBackupSnapshotExportBucketsWithDefaults instantiates a new PaginatedBackupSnapshotExportBuckets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBackupSnapshotExportBucketsWithDefaults() *PaginatedBackupSnapshotExportBuckets {
	this := PaginatedBackupSnapshotExportBuckets{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBuckets) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBuckets) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBuckets) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedBackupSnapshotExportBuckets) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBuckets) GetResults() []DiskBackupSnapshotExportBucket {
	if o == nil || IsNil(o.Results) {
		var ret []DiskBackupSnapshotExportBucket
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBuckets) GetResultsOk() (*[]DiskBackupSnapshotExportBucket, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBuckets) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiskBackupSnapshotExportBucket and assigns it to the Results field.
func (o *PaginatedBackupSnapshotExportBuckets) SetResults(v []DiskBackupSnapshotExportBucket) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedBackupSnapshotExportBuckets) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBackupSnapshotExportBuckets) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedBackupSnapshotExportBuckets) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedBackupSnapshotExportBuckets) SetTotalCount(v int) {
	o.TotalCount = &v
}
