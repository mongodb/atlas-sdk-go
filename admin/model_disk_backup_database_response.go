// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DiskBackupDatabaseResponse struct for DiskBackupDatabaseResponse
type DiskBackupDatabaseResponse struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable label that identifies the database within the snapshot.
	// Read only field.
	Name string `json:"name"`
}

// NewDiskBackupDatabaseResponse instantiates a new DiskBackupDatabaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupDatabaseResponse(name string) *DiskBackupDatabaseResponse {
	this := DiskBackupDatabaseResponse{}
	this.Name = name
	return &this
}

// NewDiskBackupDatabaseResponseWithDefaults instantiates a new DiskBackupDatabaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupDatabaseResponseWithDefaults() *DiskBackupDatabaseResponse {
	this := DiskBackupDatabaseResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupDatabaseResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupDatabaseResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupDatabaseResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupDatabaseResponse) SetLinks(v []Link) {
	o.Links = &v
}

// GetName returns the Name field value
func (o *DiskBackupDatabaseResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiskBackupDatabaseResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DiskBackupDatabaseResponse) SetName(v string) {
	o.Name = v
}
